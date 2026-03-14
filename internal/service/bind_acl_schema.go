package service

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/bind"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"terraform-provider-opnsense/internal/tools"
)

type BindAclResourceModel struct {
	Name     types.String `tfsdk:"name"`
	Enabled  types.Bool   `tfsdk:"enabled"`
	Networks types.Set    `tfsdk:"networks"`
	Id       types.String `tfsdk:"id"`
}

func BindAclResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a BIND ACL in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Name of the ACL.",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this ACL. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"networks": schema.SetAttribute{
				MarkdownDescription: "Networks included in the ACL.",
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

func BindAclDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a BIND ACL in OPNsense.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "Name of the ACL.",
				Computed:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether this ACL is enabled.",
				Computed:            true,
			},
			"networks": dschema.SetAttribute{
				MarkdownDescription: "Networks included in the ACL.",
				Computed:            true,
				ElementType:         types.StringType,
			},
		},
	}
}

func convertBindAclSchemaToStruct(d *BindAclResourceModel) (*bind.Acl, error) {
	var networksList []string
	d.Networks.ElementsAs(context.Background(), &networksList, false)

	return &bind.Acl{
		Name:     d.Name.ValueString(),
		Enabled:  tools.BoolToString(d.Enabled.ValueBool()),
		Networks: networksList,
	}, nil
}

func convertBindAclStructToSchema(d *bind.Acl) (*BindAclResourceModel, error) {
	model := &BindAclResourceModel{
		Name:    types.StringValue(d.Name),
		Enabled: types.BoolValue(tools.StringToBool(d.Enabled)),
	}

	var networkList []attr.Value
	for _, i := range d.Networks {
		if i == "" {
			continue
		}
		networkList = append(networkList, basetypes.NewStringValue(i))
	}
	typeList, _ := types.SetValue(types.StringType, networkList)
	model.Networks = typeList

	return model, nil
}
