package service

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ipsec"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

type IPsecChildResourceModel struct {
	Enabled         types.Bool   `tfsdk:"enabled"`
	Connection      types.String `tfsdk:"connection"`
	Proposals       types.Set    `tfsdk:"proposals"`
	SHA256_96       types.Bool   `tfsdk:"sha256_96"`
	StartAction     types.String `tfsdk:"start_action"`
	CloseAction     types.String `tfsdk:"close_action"`
	DPDAction       types.String `tfsdk:"dpd_action"`
	Mode            types.String `tfsdk:"mode"`
	InstallPolicies types.Bool   `tfsdk:"install_policies"`
	LocalNetworks   types.Set    `tfsdk:"local_networks"`
	RemoteNetworks  types.Set    `tfsdk:"remote_networks"`
	RequestID       types.String `tfsdk:"request_id"`
	RekeyTime       types.String `tfsdk:"rekey_time"`
	Description     types.String `tfsdk:"description"`
	Id              types.String `tfsdk:"id"`
}

func IPsecChildResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an IPsec child SA in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"enabled":    schema.BoolAttribute{MarkdownDescription: "Enable. Defaults to `true`.", Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
			"connection": schema.StringAttribute{MarkdownDescription: "Connection reference.", Required: true},
			"proposals":  schema.SetAttribute{MarkdownDescription: "ESP proposals.", Optional: true, Computed: true, ElementType: types.StringType, Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType))},
			"sha256_96":  schema.BoolAttribute{MarkdownDescription: "Use truncated SHA256 (96 bit). Defaults to `false`.", Optional: true, Computed: true, Default: booldefault.StaticBool(false)},
			"start_action":     schema.StringAttribute{MarkdownDescription: "Start action.", Optional: true, Computed: true},
			"close_action":     schema.StringAttribute{MarkdownDescription: "Close action.", Optional: true, Computed: true},
			"dpd_action":       schema.StringAttribute{MarkdownDescription: "DPD action.", Optional: true, Computed: true},
			"mode":             schema.StringAttribute{MarkdownDescription: "Mode (tunnel, transport, etc.).", Optional: true, Computed: true},
			"install_policies": schema.BoolAttribute{MarkdownDescription: "Install policies. Defaults to `true`.", Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
			"local_networks":   schema.SetAttribute{MarkdownDescription: "Local traffic selectors.", Optional: true, Computed: true, ElementType: types.StringType, Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType))},
			"remote_networks":  schema.SetAttribute{MarkdownDescription: "Remote traffic selectors.", Optional: true, Computed: true, ElementType: types.StringType, Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType))},
			"request_id":       schema.StringAttribute{MarkdownDescription: "Request ID.", Optional: true},
			"rekey_time":       schema.StringAttribute{MarkdownDescription: "Rekey time.", Optional: true},
			"description":      schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{Computed: true, MarkdownDescription: "UUID of the resource.", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
		},
	}
}

func IPsecChildDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an IPsec child SA.",
		Attributes: map[string]dschema.Attribute{
			"id":               dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"enabled":          dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether enabled."},
			"connection":       dschema.StringAttribute{Computed: true, MarkdownDescription: "Connection."},
			"proposals":        dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "ESP proposals."},
			"sha256_96":        dschema.BoolAttribute{Computed: true, MarkdownDescription: "Truncated SHA256."},
			"start_action":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Start action."},
			"close_action":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Close action."},
			"dpd_action":       dschema.StringAttribute{Computed: true, MarkdownDescription: "DPD action."},
			"mode":             dschema.StringAttribute{Computed: true, MarkdownDescription: "Mode."},
			"install_policies": dschema.BoolAttribute{Computed: true, MarkdownDescription: "Install policies."},
			"local_networks":   dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Local networks."},
			"remote_networks":  dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Remote networks."},
			"request_id":       dschema.StringAttribute{Computed: true, MarkdownDescription: "Request ID."},
			"rekey_time":       dschema.StringAttribute{Computed: true, MarkdownDescription: "Rekey time."},
			"description":      dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertIPsecChildSchemaToStruct(d *IPsecChildResourceModel) (*ipsec.IPsecChild, error) {
	return &ipsec.IPsecChild{
		Enabled:         tools.BoolToString(d.Enabled.ValueBool()),
		Connection:      api.SelectedMap(d.Connection.ValueString()),
		Proposals:       setToSlice(d.Proposals),
		SHA256_96:       tools.BoolToString(d.SHA256_96.ValueBool()),
		StartAction:     api.SelectedMap(d.StartAction.ValueString()),
		CloseAction:     api.SelectedMap(d.CloseAction.ValueString()),
		DPDAction:       api.SelectedMap(d.DPDAction.ValueString()),
		Mode:            api.SelectedMap(d.Mode.ValueString()),
		InstallPolicies: tools.BoolToString(d.InstallPolicies.ValueBool()),
		LocalNetworks:   setToSlice(d.LocalNetworks),
		RemoteNetworks:  setToSlice(d.RemoteNetworks),
		RequestID:       d.RequestID.ValueString(),
		RekeyTime:       d.RekeyTime.ValueString(),
		Description:     d.Description.ValueString(),
	}, nil
}

func convertIPsecChildStructToSchema(d *ipsec.IPsecChild) (*IPsecChildResourceModel, error) {
	return &IPsecChildResourceModel{
		Enabled:         types.BoolValue(tools.StringToBool(d.Enabled)),
		Connection:      types.StringValue(d.Connection.String()),
		Proposals:       sliceToSet(d.Proposals),
		SHA256_96:       types.BoolValue(tools.StringToBool(d.SHA256_96)),
		StartAction:     types.StringValue(d.StartAction.String()),
		CloseAction:     types.StringValue(d.CloseAction.String()),
		DPDAction:       types.StringValue(d.DPDAction.String()),
		Mode:            types.StringValue(d.Mode.String()),
		InstallPolicies: types.BoolValue(tools.StringToBool(d.InstallPolicies)),
		LocalNetworks:   sliceToSet(d.LocalNetworks),
		RemoteNetworks:  sliceToSet(d.RemoteNetworks),
		RequestID:       tools.StringOrNull(d.RequestID),
		RekeyTime:       tools.StringOrNull(d.RekeyTime),
		Description:     tools.StringOrNull(d.Description),
	}, nil
}
