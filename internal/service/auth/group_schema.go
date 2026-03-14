package auth

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/auth"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// groupResourceModel describes the resource data model.
type groupResourceModel struct {
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
	Scope          types.String `tfsdk:"scope"`
	Privilege      types.Set    `tfsdk:"privilege"`
	Member         types.String `tfsdk:"member"`
	SourceNetworks types.String `tfsdk:"source_networks"`

	GroupId types.String `tfsdk:"group_id"`
	Id      types.String `tfsdk:"id"`
}

func groupResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense auth group.",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the group.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this group.",
				Optional:            true,
			},
			"scope": schema.StringAttribute{
				MarkdownDescription: "The group scope.",
				Optional:            true,
				Computed:            true,
			},
			"privilege": schema.SetAttribute{
				MarkdownDescription: "Set of privileges assigned to this group. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"member": schema.StringAttribute{
				MarkdownDescription: "Group member.",
				Optional:            true,
			},
			"source_networks": schema.StringAttribute{
				MarkdownDescription: "Source networks.",
				Optional:            true,
			},
			"group_id": schema.StringAttribute{
				MarkdownDescription: "The numeric group ID (GID).",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func groupDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense auth group.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "The name of the group.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this group.",
				Computed:            true,
			},
			"scope": dschema.StringAttribute{
				MarkdownDescription: "The group scope.",
				Computed:            true,
			},
			"privilege": dschema.SetAttribute{
				MarkdownDescription: "Set of privileges assigned to this group.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"member": dschema.StringAttribute{
				MarkdownDescription: "Group member.",
				Computed:            true,
			},
			"source_networks": dschema.StringAttribute{
				MarkdownDescription: "Source networks.",
				Computed:            true,
			},
			"group_id": dschema.StringAttribute{
				MarkdownDescription: "The numeric group ID (GID).",
				Computed:            true,
			},
		},
	}
}

func convertGroupSchemaToStruct(d *groupResourceModel) (*auth.Group, error) {
	// Parse 'Privilege'
	var privilegeList []string
	d.Privilege.ElementsAs(context.Background(), &privilegeList, false)

	return &auth.Group{
		Name:           d.Name.ValueString(),
		Description:    d.Description.ValueString(),
		Scope:          d.Scope.ValueString(),
		Privilege:      privilegeList,
		Member:         api.SelectedMap(d.Member.ValueString()),
		SourceNetworks: api.SelectedMap(d.SourceNetworks.ValueString()),
	}, nil
}

func convertGroupStructToSchema(d *auth.Group) (*groupResourceModel, error) {
	model := &groupResourceModel{
		Name:           types.StringValue(d.Name),
		Description:    tools.StringOrNull(d.Description),
		Scope:          types.StringValue(d.Scope),
		Member:         tools.StringOrNull(d.Member.String()),
		SourceNetworks: tools.StringOrNull(d.SourceNetworks.String()),
		GroupId:        tools.StringOrNull(d.GroupId),
	}

	// Parse 'Privilege'
	var privilegeList []attr.Value
	for _, i := range d.Privilege {
		if i == "" {
			continue
		}
		privilegeList = append(privilegeList, basetypes.NewStringValue(i))
	}
	privilegeTypeList, _ := types.SetValue(types.StringType, privilegeList)
	model.Privilege = privilegeTypeList

	return model, nil
}
