package service

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

type InterfacesVipResourceModel struct {
	Interface   types.String `tfsdk:"interface"`
	Mode        types.String `tfsdk:"mode"`
	Network     types.String `tfsdk:"network"`
	Description types.String `tfsdk:"description"`
	Gateway     types.String `tfsdk:"gateway"`
	Id          types.String `tfsdk:"id"`
}

func InterfacesVipResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense virtual IP (VIP).",
		Attributes: map[string]schema.Attribute{
			"interface":   schema.StringAttribute{MarkdownDescription: "Interface.", Required: true},
			"mode":        schema.StringAttribute{MarkdownDescription: "VIP mode (e.g. ipalias, carp, proxyarp).", Required: true},
			"network":     schema.StringAttribute{MarkdownDescription: "Network/IP address with CIDR.", Required: true},
			"description": schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"gateway":     schema.StringAttribute{MarkdownDescription: "Gateway address.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func InterfacesVipDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an OPNsense virtual IP.",
		Attributes: map[string]dschema.Attribute{
			"id":          dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"interface":   dschema.StringAttribute{Computed: true, MarkdownDescription: "Interface."},
			"mode":        dschema.StringAttribute{Computed: true, MarkdownDescription: "VIP mode."},
			"network":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Network/IP."},
			"description": dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
			"gateway":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Gateway."},
		},
	}
}

func convertInterfacesVipSchemaToStruct(d *InterfacesVipResourceModel) (*interfaces.Vip, error) {
	return &interfaces.Vip{
		Interface:   api.SelectedMap(d.Interface.ValueString()),
		Mode:        api.SelectedMap(d.Mode.ValueString()),
		Network:     d.Network.ValueString(),
		Description: d.Description.ValueString(),
		Gateway:     d.Gateway.ValueString(),
	}, nil
}

func convertInterfacesVipStructToSchema(d *interfaces.Vip) (*InterfacesVipResourceModel, error) {
	return &InterfacesVipResourceModel{
		Interface:   types.StringValue(d.Interface.String()),
		Mode:        types.StringValue(d.Mode.String()),
		Network:     types.StringValue(d.Network),
		Description: tools.StringOrNull(d.Description),
		Gateway:     tools.StringOrNull(d.Gateway),
	}, nil
}
