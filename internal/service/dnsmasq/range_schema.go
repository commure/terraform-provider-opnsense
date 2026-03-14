package dnsmasq

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type rangeResourceModel struct {
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

	Id types.String `tfsdk:"id"`
}

func rangeResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense Dnsmasq DHCP range.",

		Attributes: map[string]schema.Attribute{
			"interface": schema.StringAttribute{
				MarkdownDescription: "The interface.",
				Optional:            true,
			},
			"tag": schema.StringAttribute{
				MarkdownDescription: "The tag.",
				Optional:            true,
			},
			"start_address": schema.StringAttribute{
				MarkdownDescription: "The start address of the range.",
				Optional:            true,
			},
			"end_address": schema.StringAttribute{
				MarkdownDescription: "The end address of the range.",
				Optional:            true,
			},
			"subnet_mask": schema.StringAttribute{
				MarkdownDescription: "The subnet mask.",
				Optional:            true,
			},
			"constructor": schema.StringAttribute{
				MarkdownDescription: "The constructor.",
				Optional:            true,
			},
			"mode": schema.StringAttribute{
				MarkdownDescription: "The mode.",
				Optional:            true,
			},
			"prefix_length": schema.StringAttribute{
				MarkdownDescription: "The prefix length.",
				Optional:            true,
			},
			"lease_time": schema.StringAttribute{
				MarkdownDescription: "The lease time.",
				Optional:            true,
			},
			"domain_type": schema.StringAttribute{
				MarkdownDescription: "The domain type.",
				Optional:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The domain.",
				Optional:            true,
			},
			"disable_ha_sync": schema.BoolAttribute{
				MarkdownDescription: "Disable HA sync. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ra_mode": schema.StringAttribute{
				MarkdownDescription: "The RA mode.",
				Optional:            true,
			},
			"ra_priority": schema.StringAttribute{
				MarkdownDescription: "The RA priority.",
				Optional:            true,
			},
			"ra_mtu": schema.StringAttribute{
				MarkdownDescription: "The RA MTU.",
				Optional:            true,
			},
			"ra_interval": schema.StringAttribute{
				MarkdownDescription: "The RA interval.",
				Optional:            true,
			},
			"ra_router_lifetime": schema.StringAttribute{
				MarkdownDescription: "The RA router lifetime.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this range.",
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

func rangeDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense Dnsmasq DHCP range.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"interface": dschema.StringAttribute{
				MarkdownDescription: "The interface.",
				Computed:            true,
			},
			"tag": dschema.StringAttribute{
				MarkdownDescription: "The tag.",
				Computed:            true,
			},
			"start_address": dschema.StringAttribute{
				MarkdownDescription: "The start address of the range.",
				Computed:            true,
			},
			"end_address": dschema.StringAttribute{
				MarkdownDescription: "The end address of the range.",
				Computed:            true,
			},
			"subnet_mask": dschema.StringAttribute{
				MarkdownDescription: "The subnet mask.",
				Computed:            true,
			},
			"constructor": dschema.StringAttribute{
				MarkdownDescription: "The constructor.",
				Computed:            true,
			},
			"mode": dschema.StringAttribute{
				MarkdownDescription: "The mode.",
				Computed:            true,
			},
			"prefix_length": dschema.StringAttribute{
				MarkdownDescription: "The prefix length.",
				Computed:            true,
			},
			"lease_time": dschema.StringAttribute{
				MarkdownDescription: "The lease time.",
				Computed:            true,
			},
			"domain_type": dschema.StringAttribute{
				MarkdownDescription: "The domain type.",
				Computed:            true,
			},
			"domain": dschema.StringAttribute{
				MarkdownDescription: "The domain.",
				Computed:            true,
			},
			"disable_ha_sync": dschema.BoolAttribute{
				MarkdownDescription: "Whether HA sync is disabled.",
				Computed:            true,
			},
			"ra_mode": dschema.StringAttribute{
				MarkdownDescription: "The RA mode.",
				Computed:            true,
			},
			"ra_priority": dschema.StringAttribute{
				MarkdownDescription: "The RA priority.",
				Computed:            true,
			},
			"ra_mtu": dschema.StringAttribute{
				MarkdownDescription: "The RA MTU.",
				Computed:            true,
			},
			"ra_interval": dschema.StringAttribute{
				MarkdownDescription: "The RA interval.",
				Computed:            true,
			},
			"ra_router_lifetime": dschema.StringAttribute{
				MarkdownDescription: "The RA router lifetime.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this range.",
				Computed:            true,
			},
		},
	}
}

func convertRangeSchemaToStruct(d *rangeResourceModel) (*dnsmasq.Range, error) {
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

func convertRangeStructToSchema(d *dnsmasq.Range) (*rangeResourceModel, error) {
	return &rangeResourceModel{
		Interface:        tools.StringOrNull(d.Interface.String()),
		Tag:              tools.StringOrNull(d.Tag.String()),
		StartAddress:     tools.StringOrNull(d.StartAddress),
		EndAddress:       tools.StringOrNull(d.EndAddress),
		SubnetMask:       tools.StringOrNull(d.SubnetMask),
		Constructor:      tools.StringOrNull(d.Constructor.String()),
		Mode:             tools.StringOrNull(d.Mode.String()),
		PrefixLength:     tools.StringOrNull(d.PrefixLength),
		LeaseTime:        tools.StringOrNull(d.LeaseTime),
		DomainType:       tools.StringOrNull(d.DomainType.String()),
		Domain:           tools.StringOrNull(d.Domain),
		DisableHASync:    types.BoolValue(tools.StringToBool(d.DisableHASync)),
		RaMode:           tools.StringOrNull(d.RaMode.String()),
		RaPriority:       tools.StringOrNull(d.RaPriority.String()),
		RaMTU:            tools.StringOrNull(d.RaMTU),
		RaInterval:       tools.StringOrNull(d.RaInterval),
		RaRouterLifetime: tools.StringOrNull(d.RaRouterLifetime),
		Description:      tools.StringOrNull(d.Description),
	}, nil
}
