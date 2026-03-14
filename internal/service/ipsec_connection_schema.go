package service

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ipsec"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"terraform-provider-opnsense/internal/tools"
)

type IPsecConnectionResourceModel struct {
	Enabled                types.Bool   `tfsdk:"enabled"`
	Proposals              types.Set    `tfsdk:"proposals"`
	Unique                 types.String `tfsdk:"unique"`
	Aggressive             types.Bool   `tfsdk:"aggressive"`
	Version                types.String `tfsdk:"version"`
	Mobike                 types.Bool   `tfsdk:"mobike"`
	LocalAddresses         types.Set    `tfsdk:"local_addresses"`
	RemoteAddresses        types.Set    `tfsdk:"remote_addresses"`
	LocalPort              types.String `tfsdk:"local_port"`
	RemotePort             types.String `tfsdk:"remote_port"`
	UDPEncapsulation       types.Bool   `tfsdk:"udp_encapsulation"`
	ReauthenticationTime   types.String `tfsdk:"reauthentication_time"`
	RekeyTime              types.String `tfsdk:"rekey_time"`
	IKELifetime            types.String `tfsdk:"ike_lifetime"`
	DPDDelay               types.String `tfsdk:"dpd_delay"`
	DPDTimeout             types.String `tfsdk:"dpd_timeout"`
	IPPools                types.Set    `tfsdk:"ip_pools"`
	SendCertificateRequest types.Bool   `tfsdk:"send_certificate_request"`
	SendCertificate        types.String `tfsdk:"send_certificate"`
	KeyingTries            types.String `tfsdk:"keying_tries"`
	Description            types.String `tfsdk:"description"`
	Id                     types.String `tfsdk:"id"`
}

func IPsecConnectionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an IPsec connection in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this connection. Defaults to `true`.",
				Optional: true, Computed: true, Default: booldefault.StaticBool(true),
			},
			"proposals": schema.SetAttribute{
				MarkdownDescription: "IKE proposals.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"unique":     schema.StringAttribute{MarkdownDescription: "Connection uniqueness policy.", Optional: true, Computed: true},
			"aggressive": schema.BoolAttribute{MarkdownDescription: "Use aggressive mode. Defaults to `false`.", Optional: true, Computed: true, Default: booldefault.StaticBool(false)},
			"version":    schema.StringAttribute{MarkdownDescription: "IKE version.", Optional: true, Computed: true},
			"mobike":     schema.BoolAttribute{MarkdownDescription: "Enable MOBIKE. Defaults to `true`.", Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
			"local_addresses": schema.SetAttribute{
				MarkdownDescription: "Local addresses.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"remote_addresses": schema.SetAttribute{
				MarkdownDescription: "Remote addresses.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"local_port":              schema.StringAttribute{MarkdownDescription: "Local port.", Optional: true, Computed: true},
			"remote_port":             schema.StringAttribute{MarkdownDescription: "Remote port.", Optional: true, Computed: true},
			"udp_encapsulation":       schema.BoolAttribute{MarkdownDescription: "Force UDP encapsulation. Defaults to `false`.", Optional: true, Computed: true, Default: booldefault.StaticBool(false)},
			"reauthentication_time":   schema.StringAttribute{MarkdownDescription: "Reauthentication time.", Optional: true},
			"rekey_time":              schema.StringAttribute{MarkdownDescription: "Rekey time.", Optional: true},
			"ike_lifetime":            schema.StringAttribute{MarkdownDescription: "IKE SA lifetime.", Optional: true},
			"dpd_delay":               schema.StringAttribute{MarkdownDescription: "DPD delay.", Optional: true},
			"dpd_timeout":             schema.StringAttribute{MarkdownDescription: "DPD timeout.", Optional: true},
			"ip_pools": schema.SetAttribute{
				MarkdownDescription: "IP pools.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"send_certificate_request": schema.BoolAttribute{MarkdownDescription: "Send certificate request. Defaults to `true`.", Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
			"send_certificate":         schema.StringAttribute{MarkdownDescription: "Send certificate mode.", Optional: true, Computed: true},
			"keying_tries":             schema.StringAttribute{MarkdownDescription: "Number of keying tries.", Optional: true},
			"description":              schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func IPsecConnectionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an IPsec connection.",
		Attributes: map[string]dschema.Attribute{
			"id":                       dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"enabled":                  dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether enabled."},
			"proposals":                dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "IKE proposals."},
			"unique":                   dschema.StringAttribute{Computed: true, MarkdownDescription: "Uniqueness policy."},
			"aggressive":               dschema.BoolAttribute{Computed: true, MarkdownDescription: "Aggressive mode."},
			"version":                  dschema.StringAttribute{Computed: true, MarkdownDescription: "IKE version."},
			"mobike":                   dschema.BoolAttribute{Computed: true, MarkdownDescription: "MOBIKE enabled."},
			"local_addresses":          dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Local addresses."},
			"remote_addresses":         dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Remote addresses."},
			"local_port":               dschema.StringAttribute{Computed: true, MarkdownDescription: "Local port."},
			"remote_port":              dschema.StringAttribute{Computed: true, MarkdownDescription: "Remote port."},
			"udp_encapsulation":        dschema.BoolAttribute{Computed: true, MarkdownDescription: "UDP encapsulation."},
			"reauthentication_time":    dschema.StringAttribute{Computed: true, MarkdownDescription: "Reauthentication time."},
			"rekey_time":               dschema.StringAttribute{Computed: true, MarkdownDescription: "Rekey time."},
			"ike_lifetime":             dschema.StringAttribute{Computed: true, MarkdownDescription: "IKE lifetime."},
			"dpd_delay":                dschema.StringAttribute{Computed: true, MarkdownDescription: "DPD delay."},
			"dpd_timeout":              dschema.StringAttribute{Computed: true, MarkdownDescription: "DPD timeout."},
			"ip_pools":                 dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "IP pools."},
			"send_certificate_request": dschema.BoolAttribute{Computed: true, MarkdownDescription: "Send cert request."},
			"send_certificate":         dschema.StringAttribute{Computed: true, MarkdownDescription: "Send certificate."},
			"keying_tries":             dschema.StringAttribute{Computed: true, MarkdownDescription: "Keying tries."},
			"description":              dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func sliceToSet(s []string) types.Set {
	var list []attr.Value
	for _, i := range s {
		if i == "" { continue }
		list = append(list, basetypes.NewStringValue(i))
	}
	v, _ := types.SetValue(types.StringType, list)
	return v
}

func setToSlice(s types.Set) []string {
	var list []string
	s.ElementsAs(context.Background(), &list, false)
	return list
}

func convertIPsecConnectionSchemaToStruct(d *IPsecConnectionResourceModel) (*ipsec.IPsecConnection, error) {
	return &ipsec.IPsecConnection{
		Enabled:                tools.BoolToString(d.Enabled.ValueBool()),
		Proposals:              setToSlice(d.Proposals),
		Unique:                 api.SelectedMap(d.Unique.ValueString()),
		Aggressive:             tools.BoolToString(d.Aggressive.ValueBool()),
		Version:                api.SelectedMap(d.Version.ValueString()),
		Mobike:                 tools.BoolToString(d.Mobike.ValueBool()),
		LocalAddresses:         setToSlice(d.LocalAddresses),
		RemoteAddresses:        setToSlice(d.RemoteAddresses),
		LocalPort:              api.SelectedMap(d.LocalPort.ValueString()),
		RemotePort:             api.SelectedMap(d.RemotePort.ValueString()),
		UDPEncapsulation:       tools.BoolToString(d.UDPEncapsulation.ValueBool()),
		ReauthenticationTime:   d.ReauthenticationTime.ValueString(),
		RekeyTime:              d.RekeyTime.ValueString(),
		IKELifetime:            d.IKELifetime.ValueString(),
		DPDDelay:               d.DPDDelay.ValueString(),
		DPDTimeout:             d.DPDTimeout.ValueString(),
		IPPools:                setToSlice(d.IPPools),
		SendCertificateRequest: tools.BoolToString(d.SendCertificateRequest.ValueBool()),
		SendCertificate:        api.SelectedMap(d.SendCertificate.ValueString()),
		KeyingTries:            d.KeyingTries.ValueString(),
		Description:            d.Description.ValueString(),
	}, nil
}

func convertIPsecConnectionStructToSchema(d *ipsec.IPsecConnection) (*IPsecConnectionResourceModel, error) {
	return &IPsecConnectionResourceModel{
		Enabled:                types.BoolValue(tools.StringToBool(d.Enabled)),
		Proposals:              sliceToSet(d.Proposals),
		Unique:                 types.StringValue(d.Unique.String()),
		Aggressive:             types.BoolValue(tools.StringToBool(d.Aggressive)),
		Version:                types.StringValue(d.Version.String()),
		Mobike:                 types.BoolValue(tools.StringToBool(d.Mobike)),
		LocalAddresses:         sliceToSet(d.LocalAddresses),
		RemoteAddresses:        sliceToSet(d.RemoteAddresses),
		LocalPort:              types.StringValue(d.LocalPort.String()),
		RemotePort:             types.StringValue(d.RemotePort.String()),
		UDPEncapsulation:       types.BoolValue(tools.StringToBool(d.UDPEncapsulation)),
		ReauthenticationTime:   tools.StringOrNull(d.ReauthenticationTime),
		RekeyTime:              tools.StringOrNull(d.RekeyTime),
		IKELifetime:            tools.StringOrNull(d.IKELifetime),
		DPDDelay:               tools.StringOrNull(d.DPDDelay),
		DPDTimeout:             tools.StringOrNull(d.DPDTimeout),
		IPPools:                sliceToSet(d.IPPools),
		SendCertificateRequest: types.BoolValue(tools.StringToBool(d.SendCertificateRequest)),
		SendCertificate:        types.StringValue(d.SendCertificate.String()),
		KeyingTries:            tools.StringOrNull(d.KeyingTries),
		Description:            tools.StringOrNull(d.Description),
	}, nil
}
