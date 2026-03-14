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

var _ resource.Resource = &IPsecPSKResource{}
var _ resource.ResourceWithImportState = &IPsecPSKResource{}

func NewIPsecPSKResource() resource.Resource {
	return &IPsecPSKResource{}
}

type IPsecPSKResource struct {
	client opnsense.Client
}

func (r *IPsecPSKResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsec_psk"
}

func (r *IPsecPSKResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = IPsecPSKResourceSchema()
}

func (r *IPsecPSKResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *IPsecPSKResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *IPsecPSKResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := convertIPsecPSKSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to parse IPsec pre-shared key, got error: %s", err))
		return
	}
	id, err := r.client.Ipsec().AddIPsecPSK(ctx, obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create IPsec pre-shared key, got error: %s", err))
		return
	}
	data.Id = types.StringValue(id)
	tflog.Trace(ctx, "created a resource")
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPsecPSKResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *IPsecPSKResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := r.client.Ipsec().GetIPsecPSK(ctx, data.Id.ValueString())
	if err != nil {
		var notFoundError *errs.NotFoundError
		if errors.As(err, &notFoundError) {
			tflog.Warn(ctx, "IPsec pre-shared key not present in remote, removing from state")
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read IPsec pre-shared key, got error: %s", err))
		return
	}
	model, err := convertIPsecPSKStructToSchema(obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read IPsec pre-shared key, got error: %s", err))
		return
	}
	model.Id = data.Id
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *IPsecPSKResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *IPsecPSKResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := convertIPsecPSKSchemaToStruct(data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to parse IPsec pre-shared key, got error: %s", err))
		return
	}
	err = r.client.Ipsec().UpdateIPsecPSK(ctx, data.Id.ValueString(), obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update IPsec pre-shared key, got error: %s", err))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IPsecPSKResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *IPsecPSKResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := r.client.Ipsec().DeleteIPsecPSK(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete IPsec pre-shared key, got error: %s", err))
		return
	}
}

func (r *IPsecPSKResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
