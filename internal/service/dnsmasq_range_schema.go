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

type DnsmasqRangeResourceModel struct {
	Interface        types.String `tfsdk:"interface"`
	Tag              types.String `tfsdk:"tag"`
	StartAddress     types.String `tfsdk:"start_address"`
	EndAddress       types.String `tfsdk:"end_address"`
	SubnetMask       types.String `tfsdk:"subnet_mask"`
	Constructor      types.String `tfsdk:"constructor"`
	Mode             types.String `tfsdk:"mode"`
	PrefixLength     types.String `tfsdk:"prefix_length"`
	LeaseTime        types.String `tfsdk:"lease_time"`
	DomainType       types.String `tfsdk:"domain_type"`
	Domain           types.String `tfsdk:"domain"`
	DisableHASync    types.Bool   `tfsdk:"disable_ha_sync"`
	RaMode           types.String `tfsdk:"ra_mode"`
	RaPriority       types.String `tfsdk:"ra_priority"`
	RaMTU            types.String `tfsdk:"ra_mtu"`
	RaInterval       types.String `tfsdk:"ra_interval"`
	RaRouterLifetime types.String `tfsdk:"ra_router_lifetime"`
	Description      types.String `tfsdk:"description"`
	Id               types.String `tfsdk:"id"`
}

func DnsmasqRangeResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a Dnsmasq DHCP range in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"interface":          schema.StringAttribute{MarkdownDescription: "Interface.", Optional: true, Computed: true},
			"tag":                schema.StringAttribute{MarkdownDescription: "Tag.", Optional: true, Computed: true},
			"start_address":      schema.StringAttribute{MarkdownDescription: "Start address.", Optional: true},
			"end_address":        schema.StringAttribute{MarkdownDescription: "End address.", Optional: true},
			"subnet_mask":        schema.StringAttribute{MarkdownDescription: "Subnet mask.", Optional: true},
			"constructor":        schema.StringAttribute{MarkdownDescription: "Constructor.", Optional: true, Computed: true},
			"mode":               schema.StringAttribute{MarkdownDescription: "Mode.", Optional: true, Computed: true},
			"prefix_length":      schema.StringAttribute{MarkdownDescription: "IPv6 prefix length.", Optional: true},
			"lease_time":         schema.StringAttribute{MarkdownDescription: "Lease time.", Optional: true},
			"domain_type":        schema.StringAttribute{MarkdownDescription: "Domain type.", Optional: true, Computed: true},
			"domain":             schema.StringAttribute{MarkdownDescription: "Domain.", Optional: true},
			"disable_ha_sync":    schema.BoolAttribute{MarkdownDescription: "Disable HA sync.", Optional: true, Computed: true},
			"ra_mode":            schema.StringAttribute{MarkdownDescription: "Router advertisement mode.", Optional: true, Computed: true},
			"ra_priority":        schema.StringAttribute{MarkdownDescription: "Router advertisement priority.", Optional: true, Computed: true},
			"ra_mtu":             schema.StringAttribute{MarkdownDescription: "Router advertisement MTU.", Optional: true},
			"ra_interval":        schema.StringAttribute{MarkdownDescription: "Router advertisement interval.", Optional: true},
			"ra_router_lifetime": schema.StringAttribute{MarkdownDescription: "Router advertisement router lifetime.", Optional: true},
			"description":        schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func DnsmasqRangeDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a Dnsmasq DHCP range.",
		Attributes: map[string]dschema.Attribute{
			"id":                 dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"interface":          dschema.StringAttribute{Computed: true, MarkdownDescription: "Interface."},
			"tag":                dschema.StringAttribute{Computed: true, MarkdownDescription: "Tag."},
			"start_address":      dschema.StringAttribute{Computed: true, MarkdownDescription: "Start address."},
			"end_address":        dschema.StringAttribute{Computed: true, MarkdownDescription: "End address."},
			"subnet_mask":        dschema.StringAttribute{Computed: true, MarkdownDescription: "Subnet mask."},
			"constructor":        dschema.StringAttribute{Computed: true, MarkdownDescription: "Constructor."},
			"mode":               dschema.StringAttribute{Computed: true, MarkdownDescription: "Mode."},
			"prefix_length":      dschema.StringAttribute{Computed: true, MarkdownDescription: "Prefix length."},
			"lease_time":         dschema.StringAttribute{Computed: true, MarkdownDescription: "Lease time."},
			"domain_type":        dschema.StringAttribute{Computed: true, MarkdownDescription: "Domain type."},
			"domain":             dschema.StringAttribute{Computed: true, MarkdownDescription: "Domain."},
			"disable_ha_sync":    dschema.BoolAttribute{Computed: true, MarkdownDescription: "Disable HA sync."},
			"ra_mode":            dschema.StringAttribute{Computed: true, MarkdownDescription: "RA mode."},
			"ra_priority":        dschema.StringAttribute{Computed: true, MarkdownDescription: "RA priority."},
			"ra_mtu":             dschema.StringAttribute{Computed: true, MarkdownDescription: "RA MTU."},
			"ra_interval":        dschema.StringAttribute{Computed: true, MarkdownDescription: "RA interval."},
			"ra_router_lifetime": dschema.StringAttribute{Computed: true, MarkdownDescription: "RA router lifetime."},
			"description":        dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertDnsmasqRangeSchemaToStruct(d *DnsmasqRangeResourceModel) (*dnsmasq.Range, error) {
	return &dnsmasq.Range{
		Interface:        api.SelectedMap(d.Interface.ValueString()),
		Tag:              api.SelectedMap(d.Tag.ValueString()),
		StartAddress:     d.StartAddress.ValueString(),
		EndAddress:       d.EndAddress.ValueString(),
		SubnetMask:       d.SubnetMask.ValueString(),
		Constructor:      api.SelectedMap(d.Constructor.ValueString()),
		Mode:             api.SelectedMap(d.Mode.ValueString()),
		PrefixLength:     d.PrefixLength.ValueString(),
		LeaseTime:        d.LeaseTime.ValueString(),
		DomainType:       api.SelectedMap(d.DomainType.ValueString()),
		Domain:           d.Domain.ValueString(),
		DisableHASync:    tools.BoolToString(d.DisableHASync.ValueBool()),
		RaMode:           api.SelectedMap(d.RaMode.ValueString()),
		RaPriority:       api.SelectedMap(d.RaPriority.ValueString()),
		RaMTU:            d.RaMTU.ValueString(),
		RaInterval:       d.RaInterval.ValueString(),
		RaRouterLifetime: d.RaRouterLifetime.ValueString(),
		Description:      d.Description.ValueString(),
	}, nil
}

func convertDnsmasqRangeStructToSchema(d *dnsmasq.Range) (*DnsmasqRangeResourceModel, error) {
	return &DnsmasqRangeResourceModel{
		Interface:        types.StringValue(d.Interface.String()),
		Tag:              types.StringValue(d.Tag.String()),
		StartAddress:     tools.StringOrNull(d.StartAddress),
		EndAddress:       tools.StringOrNull(d.EndAddress),
		SubnetMask:       tools.StringOrNull(d.SubnetMask),
		Constructor:      types.StringValue(d.Constructor.String()),
		Mode:             types.StringValue(d.Mode.String()),
		PrefixLength:     tools.StringOrNull(d.PrefixLength),
		LeaseTime:        tools.StringOrNull(d.LeaseTime),
		DomainType:       types.StringValue(d.DomainType.String()),
		Domain:           tools.StringOrNull(d.Domain),
		DisableHASync:    types.BoolValue(tools.StringToBool(d.DisableHASync)),
		RaMode:           types.StringValue(d.RaMode.String()),
		RaPriority:       types.StringValue(d.RaPriority.String()),
		RaMTU:            tools.StringOrNull(d.RaMTU),
		RaInterval:       tools.StringOrNull(d.RaInterval),
		RaRouterLifetime: tools.StringOrNull(d.RaRouterLifetime),
		Description:      tools.StringOrNull(d.Description),
	}, nil
}
