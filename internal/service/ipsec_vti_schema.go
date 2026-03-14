package service

import (
	"github.com/browningluke/opnsense-go/pkg/ipsec"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

type IPsecVTIResourceModel struct {
	Enabled         types.Bool   `tfsdk:"enabled"`
	RequestID       types.String `tfsdk:"request_id"`
	LocalIP         types.String `tfsdk:"local_ip"`
	RemoteIP        types.String `tfsdk:"remote_ip"`
	TunnelLocalIP   types.String `tfsdk:"tunnel_local_ip"`
	TunnelRemoteIP  types.String `tfsdk:"tunnel_remote_ip"`
	TunnelLocalIP2  types.String `tfsdk:"tunnel_local_ip2"`
	TunnelRemoteIP2 types.String `tfsdk:"tunnel_remote_ip2"`
	Description     types.String `tfsdk:"description"`
	Id              types.String `tfsdk:"id"`
}

func IPsecVTIResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an IPsec Virtual Tunnel Interface (VTI) in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"enabled":           schema.BoolAttribute{MarkdownDescription: "Enable. Defaults to `true`.", Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
			"request_id":        schema.StringAttribute{MarkdownDescription: "Request ID.", Optional: true},
			"local_ip":          schema.StringAttribute{MarkdownDescription: "Local endpoint IP.", Required: true},
			"remote_ip":         schema.StringAttribute{MarkdownDescription: "Remote endpoint IP.", Required: true},
			"tunnel_local_ip":   schema.StringAttribute{MarkdownDescription: "Tunnel local IP.", Optional: true},
			"tunnel_remote_ip":  schema.StringAttribute{MarkdownDescription: "Tunnel remote IP.", Optional: true},
			"tunnel_local_ip2":  schema.StringAttribute{MarkdownDescription: "Tunnel local IP (secondary).", Optional: true},
			"tunnel_remote_ip2": schema.StringAttribute{MarkdownDescription: "Tunnel remote IP (secondary).", Optional: true},
			"description":       schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{Computed: true, MarkdownDescription: "UUID of the resource.", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
		},
	}
}

func IPsecVTIDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an IPsec VTI.",
		Attributes: map[string]dschema.Attribute{
			"id":                dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"enabled":           dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether enabled."},
			"request_id":        dschema.StringAttribute{Computed: true, MarkdownDescription: "Request ID."},
			"local_ip":          dschema.StringAttribute{Computed: true, MarkdownDescription: "Local IP."},
			"remote_ip":         dschema.StringAttribute{Computed: true, MarkdownDescription: "Remote IP."},
			"tunnel_local_ip":   dschema.StringAttribute{Computed: true, MarkdownDescription: "Tunnel local IP."},
			"tunnel_remote_ip":  dschema.StringAttribute{Computed: true, MarkdownDescription: "Tunnel remote IP."},
			"tunnel_local_ip2":  dschema.StringAttribute{Computed: true, MarkdownDescription: "Tunnel local IP (secondary)."},
			"tunnel_remote_ip2": dschema.StringAttribute{Computed: true, MarkdownDescription: "Tunnel remote IP (secondary)."},
			"description":       dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertIPsecVTISchemaToStruct(d *IPsecVTIResourceModel) (*ipsec.IPsecVTI, error) {
	return &ipsec.IPsecVTI{
		Enabled:         tools.BoolToString(d.Enabled.ValueBool()),
		RequestID:       d.RequestID.ValueString(),
		LocalIP:         d.LocalIP.ValueString(),
		RemoteIP:        d.RemoteIP.ValueString(),
		TunnelLocalIP:   d.TunnelLocalIP.ValueString(),
		TunnelRemoteIP:  d.TunnelRemoteIP.ValueString(),
		TunnelLocalIP2:  d.TunnelLocalIP2.ValueString(),
		TunnelRemoteIP2: d.TunnelRemoteIP2.ValueString(),
		Description:     d.Description.ValueString(),
	}, nil
}

func convertIPsecVTIStructToSchema(d *ipsec.IPsecVTI) (*IPsecVTIResourceModel, error) {
	return &IPsecVTIResourceModel{
		Enabled:         types.BoolValue(tools.StringToBool(d.Enabled)),
		RequestID:       tools.StringOrNull(d.RequestID),
		LocalIP:         types.StringValue(d.LocalIP),
		RemoteIP:        types.StringValue(d.RemoteIP),
		TunnelLocalIP:   tools.StringOrNull(d.TunnelLocalIP),
		TunnelRemoteIP:  tools.StringOrNull(d.TunnelRemoteIP),
		TunnelLocalIP2:  tools.StringOrNull(d.TunnelLocalIP2),
		TunnelRemoteIP2: tools.StringOrNull(d.TunnelRemoteIP2),
		Description:     tools.StringOrNull(d.Description),
	}, nil
}
