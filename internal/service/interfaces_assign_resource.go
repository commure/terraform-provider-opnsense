package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

var _ resource.Resource = &InterfacesAssignResource{}
var _ resource.ResourceWithImportState = &InterfacesAssignResource{}

func NewInterfacesAssignResource() resource.Resource {
	return &InterfacesAssignResource{}
}

type InterfacesAssignResource struct {
	client opnsense.Client
}

func (r *InterfacesAssignResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interfaces_assign"
}

func (r *InterfacesAssignResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = InterfacesAssignResourceSchema()
}

func (r *InterfacesAssignResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *opnsense.Client, got: %T.", req.ProviderData))
		return
	}
	r.client = opnsense.NewClient(apiClient)
}

func (r *InterfacesAssignResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *InterfacesAssignResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := convertInterfacesAssignSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to parse interface assignment, got error: %s", err))
		return
	}
	id, err := r.client.Interfaces().AddAssign(ctx, obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create interface assignment, got error: %s", err))
		return
	}
	data.Id = types.StringValue(id)
	tflog.Trace(ctx, "created a resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfacesAssignResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *InterfacesAssignResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := r.client.Interfaces().GetAssign(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, "interface assignment not present in remote, removing from state")
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read interface assignment, got error: %s", err))
		return
	}
	model, err := convertInterfacesAssignStructToSchema(obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read interface assignment, got error: %s", err))
		return
	}
	model.Id = data.Id
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *InterfacesAssignResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *InterfacesAssignResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := convertInterfacesAssignSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to parse interface assignment, got error: %s", err))
		return
	}
	err = r.client.Interfaces().UpdateAssign(ctx, data.Id.ValueString(), obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update interface assignment, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfacesAssignResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *InterfacesAssignResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.client.Interfaces().DeleteAssign(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete interface assignment, got error: %s", err))
		return
	}
}

func (r *InterfacesAssignResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
