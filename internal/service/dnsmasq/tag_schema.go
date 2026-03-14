package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type tagResourceModel struct {
	Tag types.String `tfsdk:"tag"`

	Id types.String `tfsdk:"id"`
}

func tagResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense Dnsmasq tag.",

		Attributes: map[string]schema.Attribute{
			"tag": schema.StringAttribute{
				MarkdownDescription: "The tag value.",
				Required:            true,
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

func tagDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense Dnsmasq tag.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"tag": dschema.StringAttribute{
				MarkdownDescription: "The tag value.",
				Computed:            true,
			},
		},
	}
}

func convertTagSchemaToStruct(d *tagResourceModel) (*dnsmasq.Tag, error) {
	return &dnsmasq.Tag{
		Tag: d.Tag.ValueString(),
	}, nil
}

func convertTagStructToSchema(d *dnsmasq.Tag) (*tagResourceModel, error) {
	return &tagResourceModel{
		Tag: types.StringValue(d.Tag),
	}, nil
}
