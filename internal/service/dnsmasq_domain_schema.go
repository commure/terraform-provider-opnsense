package service

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

type DnsmasqDomainResourceModel struct {
	Sequence      types.String `tfsdk:"sequence"`
	Domain        types.String `tfsdk:"domain"`
	FirewallAlias types.String `tfsdk:"firewall_alias"`
	SourceIp      types.String `tfsdk:"source_ip"`
	Port          types.String `tfsdk:"port"`
	Ip            types.String `tfsdk:"ip"`
	Description   types.String `tfsdk:"description"`
	Id            types.String `tfsdk:"id"`
}

func DnsmasqDomainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a Dnsmasq domain override in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"sequence":       schema.StringAttribute{MarkdownDescription: "Sequence order.", Optional: true},
			"domain":         schema.StringAttribute{MarkdownDescription: "Domain name.", Required: true},
			"firewall_alias": schema.StringAttribute{MarkdownDescription: "Firewall alias (ipset).", Optional: true, Computed: true},
			"source_ip":      schema.StringAttribute{MarkdownDescription: "Source IP.", Optional: true},
			"port":           schema.StringAttribute{MarkdownDescription: "Port.", Optional: true},
			"ip":             schema.StringAttribute{MarkdownDescription: "IP address.", Optional: true},
			"description":    schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func DnsmasqDomainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a Dnsmasq domain override.",
		Attributes: map[string]dschema.Attribute{
			"id":             dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"sequence":       dschema.StringAttribute{Computed: true, MarkdownDescription: "Sequence."},
			"domain":         dschema.StringAttribute{Computed: true, MarkdownDescription: "Domain name."},
			"firewall_alias": dschema.StringAttribute{Computed: true, MarkdownDescription: "Firewall alias."},
			"source_ip":      dschema.StringAttribute{Computed: true, MarkdownDescription: "Source IP."},
			"port":           dschema.StringAttribute{Computed: true, MarkdownDescription: "Port."},
			"ip":             dschema.StringAttribute{Computed: true, MarkdownDescription: "IP address."},
			"description":    dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertDnsmasqDomainSchemaToStruct(d *DnsmasqDomainResourceModel) (*dnsmasq.Domain, error) {
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

func convertDnsmasqDomainStructToSchema(d *dnsmasq.Domain) (*DnsmasqDomainResourceModel, error) {
	return &DnsmasqDomainResourceModel{
		Sequence:      tools.StringOrNull(d.Sequence),
		Domain:        types.StringValue(d.Domain),
		FirewallAlias: types.StringValue(d.FirewallAlias.String()),
		SourceIp:      tools.StringOrNull(d.SourceIp),
		Port:          tools.StringOrNull(d.Port),
		Ip:            tools.StringOrNull(d.Ip),
		Description:   tools.StringOrNull(d.Description),
	}, nil
}
