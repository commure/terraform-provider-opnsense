package service

import (
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

type InterfacesAssignResourceModel struct {
	Interface        types.String `tfsdk:"interface"`
	Device           types.String `tfsdk:"device"`
	Ip               types.String `tfsdk:"ip"`
	Subnet           types.String `tfsdk:"subnet"`
	Gateway          types.String `tfsdk:"gateway"`
	GatewayInterface types.String `tfsdk:"gateway_interface"`
	Enable           types.Bool   `tfsdk:"enable"`
	Description      types.String `tfsdk:"description"`
	Id               types.String `tfsdk:"id"`
}

func InterfacesAssignResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense interface assignment.",
		Attributes: map[string]schema.Attribute{
			"interface":         schema.StringAttribute{MarkdownDescription: "Interface identifier.", Optional: true},
			"device":            schema.StringAttribute{MarkdownDescription: "Device name.", Required: true},
			"ip":                schema.StringAttribute{MarkdownDescription: "IP address.", Optional: true},
			"subnet":            schema.StringAttribute{MarkdownDescription: "Subnet mask.", Optional: true},
			"gateway":           schema.StringAttribute{MarkdownDescription: "Gateway address.", Optional: true},
			"gateway_interface": schema.StringAttribute{MarkdownDescription: "Gateway interface.", Optional: true},
			"enable": schema.BoolAttribute{
				MarkdownDescription: "Enable this interface. Defaults to `true`.",
				Optional: true, Computed: true, Default: booldefault.StaticBool(true),
			},
			"description": schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func InterfacesAssignDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an OPNsense interface assignment.",
		Attributes: map[string]dschema.Attribute{
			"id":                dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"interface":         dschema.StringAttribute{Computed: true, MarkdownDescription: "Interface identifier."},
			"device":            dschema.StringAttribute{Computed: true, MarkdownDescription: "Device name."},
			"ip":                dschema.StringAttribute{Computed: true, MarkdownDescription: "IP address."},
			"subnet":            dschema.StringAttribute{Computed: true, MarkdownDescription: "Subnet mask."},
			"gateway":           dschema.StringAttribute{Computed: true, MarkdownDescription: "Gateway."},
			"gateway_interface": dschema.StringAttribute{Computed: true, MarkdownDescription: "Gateway interface."},
			"enable":            dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether enabled."},
			"description":       dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertInterfacesAssignSchemaToStruct(d *InterfacesAssignResourceModel) (*interfaces.Assign, error) {
	return &interfaces.Assign{
		Interface:        d.Interface.ValueString(),
		Device:           d.Device.ValueString(),
		Ip:               d.Ip.ValueString(),
		Subnet:           d.Subnet.ValueString(),
		Gateway:          d.Gateway.ValueString(),
		GatewayInterface: d.GatewayInterface.ValueString(),
		Enable:           tools.BoolToString(d.Enable.ValueBool()),
		Description:      d.Description.ValueString(),
	}, nil
}

func convertInterfacesAssignStructToSchema(d *interfaces.Assign) (*InterfacesAssignResourceModel, error) {
	return &InterfacesAssignResourceModel{
		Interface:        tools.StringOrNull(d.Interface),
		Device:           types.StringValue(d.Device),
		Ip:               tools.StringOrNull(d.Ip),
		Subnet:           tools.StringOrNull(d.Subnet),
		Gateway:          tools.StringOrNull(d.Gateway),
		GatewayInterface: tools.StringOrNull(d.GatewayInterface),
		Enable:           types.BoolValue(tools.StringToBool(d.Enable)),
		Description:      tools.StringOrNull(d.Description),
	}, nil
}
