package service

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/auth"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"terraform-provider-opnsense/internal/tools"
)

type AuthGroupResourceModel struct {
	GroupId        types.String `tfsdk:"group_id"`
	Name           types.String `tfsdk:"name"`
	Scope          types.String `tfsdk:"scope"`
	Description    types.String `tfsdk:"description"`
	Privilege      types.Set    `tfsdk:"privilege"`
	Member         types.String `tfsdk:"member"`
	SourceNetworks types.String `tfsdk:"source_networks"`

	Id types.String `tfsdk:"id"`
}

func AuthGroupResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense authentication group.",

		Attributes: map[string]schema.Attribute{
			"group_id":    schema.StringAttribute{MarkdownDescription: "The numeric group ID.", Optional: true, Computed: true},
			"name":        schema.StringAttribute{MarkdownDescription: "The group name.", Required: true},
			"scope":       schema.StringAttribute{MarkdownDescription: "The group scope.", Optional: true, Computed: true},
			"description": schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"privilege": schema.SetAttribute{
				MarkdownDescription: "Privileges.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"member":          schema.StringAttribute{MarkdownDescription: "Group member selection.", Optional: true, Computed: true},
			"source_networks": schema.StringAttribute{MarkdownDescription: "Source networks.", Optional: true, Computed: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func AuthGroupDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an OPNsense authentication group.",

		Attributes: map[string]dschema.Attribute{
			"id":              dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"group_id":        dschema.StringAttribute{Computed: true, MarkdownDescription: "The numeric group ID."},
			"name":            dschema.StringAttribute{Computed: true, MarkdownDescription: "The group name."},
			"scope":           dschema.StringAttribute{Computed: true, MarkdownDescription: "The group scope."},
			"description":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
			"privilege":       dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Privileges."},
			"member":          dschema.StringAttribute{Computed: true, MarkdownDescription: "Member."},
			"source_networks": dschema.StringAttribute{Computed: true, MarkdownDescription: "Source networks."},
		},
	}
}

func convertAuthGroupSchemaToStruct(d *AuthGroupResourceModel) (*auth.Group, error) {
	var privList []string
	d.Privilege.ElementsAs(context.Background(), &privList, false)

	return &auth.Group{
		GroupId:        d.GroupId.ValueString(),
		Name:           d.Name.ValueString(),
		Scope:          d.Scope.ValueString(),
		Description:    d.Description.ValueString(),
		Privilege:      privList,
		Member:         api.SelectedMap(d.Member.ValueString()),
		SourceNetworks: api.SelectedMap(d.SourceNetworks.ValueString()),
	}, nil
}

func convertAuthGroupStructToSchema(d *auth.Group) (*AuthGroupResourceModel, error) {
	model := &AuthGroupResourceModel{
		GroupId:        tools.StringOrNull(d.GroupId),
		Name:           types.StringValue(d.Name),
		Scope:          tools.StringOrNull(d.Scope),
		Description:    tools.StringOrNull(d.Description),
		Member:         types.StringValue(d.Member.String()),
		SourceNetworks: types.StringValue(d.SourceNetworks.String()),
	}

	var privList []attr.Value
	for _, i := range d.Privilege {
		if i == "" { continue }
		privList = append(privList, basetypes.NewStringValue(i))
	}
	privSet, _ := types.SetValue(types.StringType, privList)
	model.Privilege = privSet

	return model, nil
}
