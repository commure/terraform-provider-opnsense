package interfaces

import (
	"github.com/browningluke/opnsense-go/pkg/interfaces"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type assignResourceModel struct {
	Interface        types.String `tfsdk:"interface"`
	Device           types.String `tfsdk:"device"`
	Ip               types.String `tfsdk:"ip"`
	Subnet           types.String `tfsdk:"subnet"`
	Gateway          types.String `tfsdk:"gateway"`
	GatewayInterface types.String `tfsdk:"gateway_interface"`
	Enable           types.Bool   `tfsdk:"enable"`
	Description      types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func assignResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense interface assignment.",

		Attributes: map[string]schema.Attribute{
			"interface": schema.StringAttribute{
				MarkdownDescription: "The interface name.",
				Optional:            true,
			},
			"device": schema.StringAttribute{
				MarkdownDescription: "The device.",
				Optional:            true,
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: "The IP address.",
				Optional:            true,
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: "The subnet.",
				Optional:            true,
			},
			"gateway": schema.StringAttribute{
				MarkdownDescription: "The gateway.",
				Optional:            true,
			},
			"gateway_interface": schema.StringAttribute{
				MarkdownDescription: "The gateway interface.",
				Optional:            true,
			},
			"enable": schema.BoolAttribute{
				MarkdownDescription: "Enable this interface. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this assignment.",
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

func assignDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense interface assignment.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"interface": dschema.StringAttribute{
				MarkdownDescription: "The interface name.",
				Computed:            true,
			},
			"device": dschema.StringAttribute{
				MarkdownDescription: "The device.",
				Computed:            true,
			},
			"ip": dschema.StringAttribute{
				MarkdownDescription: "The IP address.",
				Computed:            true,
			},
			"subnet": dschema.StringAttribute{
				MarkdownDescription: "The subnet.",
				Computed:            true,
			},
			"gateway": dschema.StringAttribute{
				MarkdownDescription: "The gateway.",
				Computed:            true,
			},
			"gateway_interface": dschema.StringAttribute{
				MarkdownDescription: "The gateway interface.",
				Computed:            true,
			},
			"enable": dschema.BoolAttribute{
				MarkdownDescription: "Whether this interface is enabled.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this assignment.",
				Computed:            true,
			},
		},
	}
}

func convertAssignSchemaToStruct(d *assignResourceModel) (*interfaces.Assign, error) {
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

func convertAssignStructToSchema(d *interfaces.Assign) (*assignResourceModel, error) {
	return &assignResourceModel{
		Interface:        tools.StringOrNull(d.Interface),
		Device:           tools.StringOrNull(d.Device),
		Ip:               tools.StringOrNull(d.Ip),
		Subnet:           tools.StringOrNull(d.Subnet),
		Gateway:          tools.StringOrNull(d.Gateway),
		GatewayInterface: tools.StringOrNull(d.GatewayInterface),
		Enable:           types.BoolValue(tools.StringToBool(d.Enable)),
		Description:      tools.StringOrNull(d.Description),
	}, nil
}
