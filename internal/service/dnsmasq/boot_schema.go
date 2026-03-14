package dnsmasq

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
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

type bootResourceModel struct {
	Interface     types.String `tfsdk:"interface"`
	Tag           types.Set    `tfsdk:"tag"`
	Filename      types.String `tfsdk:"filename"`
	Servername    types.String `tfsdk:"servername"`
	ServerAddress types.String `tfsdk:"server_address"`
	Description   types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func bootResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense Dnsmasq DHCP boot entry.",

		Attributes: map[string]schema.Attribute{
			"interface": schema.StringAttribute{
				MarkdownDescription: "The interface for this boot entry.",
				Optional:            true,
			},
			"tag": schema.SetAttribute{
				MarkdownDescription: "Tags associated with this boot entry. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"filename": schema.StringAttribute{
				MarkdownDescription: "The boot filename.",
				Optional:            true,
			},
			"servername": schema.StringAttribute{
				MarkdownDescription: "The boot server name.",
				Optional:            true,
			},
			"server_address": schema.StringAttribute{
				MarkdownDescription: "The boot server address.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this boot entry.",
				Optional:            true,
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

func bootDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense Dnsmasq DHCP boot entry.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"interface": dschema.StringAttribute{
				MarkdownDescription: "The interface for this boot entry.",
				Computed:            true,
			},
			"tag": dschema.SetAttribute{
				MarkdownDescription: "Tags associated with this boot entry.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"filename": dschema.StringAttribute{
				MarkdownDescription: "The boot filename.",
				Computed:            true,
			},
			"servername": dschema.StringAttribute{
				MarkdownDescription: "The boot server name.",
				Computed:            true,
			},
			"server_address": dschema.StringAttribute{
				MarkdownDescription: "The boot server address.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this boot entry.",
				Computed:            true,
			},
		},
	}
}

func convertBootSchemaToStruct(d *bootResourceModel) (*dnsmasq.Boot, error) {
	var tagList []string
	d.Tag.ElementsAs(context.Background(), &tagList, false)

	return &dnsmasq.Boot{
		Interface:     api.SelectedMap(d.Interface.ValueString()),
		Tag:           tagList,
		Filename:      d.Filename.ValueString(),
		Servername:    d.Servername.ValueString(),
		ServerAddress: d.ServerAddress.ValueString(),
		Description:   d.Description.ValueString(),
	}, nil
}

func convertBootStructToSchema(d *dnsmasq.Boot) (*bootResourceModel, error) {
	model := &bootResourceModel{
		Interface:     types.StringValue(d.Interface.String()),
		Filename:      tools.StringOrNull(d.Filename),
		Servername:    tools.StringOrNull(d.Servername),
		ServerAddress: tools.StringOrNull(d.ServerAddress),
		Description:   tools.StringOrNull(d.Description),
	}

	var tagList []attr.Value
	for _, i := range d.Tag {
		if i == "" {
			continue
		}
		tagList = append(tagList, basetypes.NewStringValue(i))
	}
	tagTypeList, _ := types.SetValue(types.StringType, tagList)
	model.Tag = tagTypeList

	return model, nil
}
