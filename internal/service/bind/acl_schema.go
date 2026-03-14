package bind

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/bind"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// aclResourceModel describes the resource data model.
type aclResourceModel struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Name     types.String `tfsdk:"name"`
	Networks types.Set    `tfsdk:"networks"`

	Id types.String `tfsdk:"id"`
}

func aclResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense BIND ACL.",

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this ACL. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the ACL.",
				Required:            true,
			},
			"networks": schema.SetAttribute{
				MarkdownDescription: "Set of networks for this ACL. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
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

func aclDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense BIND ACL.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether this ACL is enabled.",
				Computed:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "The name of the ACL.",
				Computed:            true,
			},
			"networks": dschema.SetAttribute{
				MarkdownDescription: "Set of networks for this ACL.",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func convertAclSchemaToStruct(d *aclResourceModel) (*bind.Acl, error) {
	var networksList []string
	d.Networks.ElementsAs(context.Background(), &networksList, false)

	return &bind.Acl{
		Enabled:  tools.BoolToString(d.Enabled.ValueBool()),
		Name:     d.Name.ValueString(),
		Networks: networksList,
	}, nil
}

func convertAclStructToSchema(d *bind.Acl) (*aclResourceModel, error) {
	model := &aclResourceModel{
		Enabled: types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:    types.StringValue(d.Name),
	}

	var networksList []attr.Value
	for _, i := range d.Networks {
		if i == "" {
			continue
		}
		networksList = append(networksList, basetypes.NewStringValue(i))
	}
	networksTypeList, _ := types.SetValue(types.StringType, networksList)
	model.Networks = networksTypeList

	return model, nil
}
