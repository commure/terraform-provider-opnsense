package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type domainResourceModel struct {
	Sequence      types.String `tfsdk:"sequence"`
	Domain        types.String `tfsdk:"domain"`
	FirewallAlias types.String `tfsdk:"firewall_alias"`
	SourceIp      types.String `tfsdk:"source_ip"`
	Port          types.String `tfsdk:"port"`
	Ip            types.String `tfsdk:"ip"`
	Description   types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func domainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense Dnsmasq domain override.",

		Attributes: map[string]schema.Attribute{
			"sequence": schema.StringAttribute{
				MarkdownDescription: "The sequence number.",
				Optional:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The domain name.",
				Optional:            true,
			},
			"firewall_alias": schema.StringAttribute{
				MarkdownDescription: "The firewall alias (ipset).",
				Optional:            true,
			},
			"source_ip": schema.StringAttribute{
				MarkdownDescription: "The source IP.",
				Optional:            true,
			},
			"port": schema.StringAttribute{
				MarkdownDescription: "The port.",
				Optional:            true,
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: "The IP address.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this domain override.",
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

func domainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense Dnsmasq domain override.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"sequence": dschema.StringAttribute{
				MarkdownDescription: "The sequence number.",
				Computed:            true,
			},
			"domain": dschema.StringAttribute{
				MarkdownDescription: "The domain name.",
				Computed:            true,
			},
			"firewall_alias": dschema.StringAttribute{
				MarkdownDescription: "The firewall alias (ipset).",
				Computed:            true,
			},
			"source_ip": dschema.StringAttribute{
				MarkdownDescription: "The source IP.",
				Computed:            true,
			},
			"port": dschema.StringAttribute{
				MarkdownDescription: "The port.",
				Computed:            true,
			},
			"ip": dschema.StringAttribute{
				MarkdownDescription: "The IP address.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this domain override.",
				Computed:            true,
			},
		},
	}
}

func convertDomainSchemaToStruct(d *domainResourceModel) (*dnsmasq.Domain, error) {
	return &dnsmasq.Domain{
		Sequence:      d.Sequence.ValueString(),
		Domain:        d.Domain.ValueString(),
		FirewallAlias: api.SelectedMap(d.FirewallAlias.ValueString()),
		SourceIp:      d.SourceIp.ValueString(),
		Port:          d.Port.ValueString(),
		Ip:            d.Ip.ValueString(),
		Description:   d.Description.ValueString(),
	}, nil
}

func convertDomainStructToSchema(d *dnsmasq.Domain) (*domainResourceModel, error) {
	return &domainResourceModel{
		Sequence:      tools.StringOrNull(d.Sequence),
		Domain:        tools.StringOrNull(d.Domain),
		FirewallAlias: tools.StringOrNull(d.FirewallAlias.String()),
		SourceIp:      tools.StringOrNull(d.SourceIp),
		Port:          tools.StringOrNull(d.Port),
		Ip:            tools.StringOrNull(d.Ip),
		Description:   tools.StringOrNull(d.Description),
	}, nil
}
