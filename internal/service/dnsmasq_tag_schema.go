package service

import (
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DnsmasqTagResourceModel struct {
	Tag types.String `tfsdk:"tag"`
	Id  types.String `tfsdk:"id"`
}

func DnsmasqTagResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a Dnsmasq tag in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"tag": schema.StringAttribute{
				MarkdownDescription: "Tag name.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func DnsmasqTagDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a Dnsmasq tag.",
		Attributes: map[string]dschema.Attribute{
			"id":  dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"tag": dschema.StringAttribute{Computed: true, MarkdownDescription: "Tag name."},
		},
	}
}

func convertDnsmasqTagSchemaToStruct(d *DnsmasqTagResourceModel) (*dnsmasq.Tag, error) {
	return &dnsmasq.Tag{
		Tag: d.Tag.ValueString(),
	}, nil
}

func convertDnsmasqTagStructToSchema(d *dnsmasq.Tag) (*DnsmasqTagResourceModel, error) {
	return &DnsmasqTagResourceModel{
		Tag: types.StringValue(d.Tag),
	}, nil
}
