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

type IPsecAuthRemoteResourceModel struct {
	Enabled        types.Bool   `tfsdk:"enabled"`
	Connection     types.String `tfsdk:"connection"`
	Round          types.String `tfsdk:"round"`
	Authentication types.String `tfsdk:"authentication"`
	AuthId         types.String `tfsdk:"auth_id"`
	EAPId          types.String `tfsdk:"eap_id"`
	Certificates   types.Set    `tfsdk:"certificates"`
	PublicKeys     types.Set    `tfsdk:"public_keys"`
	Description    types.String `tfsdk:"description"`
	Id             types.String `tfsdk:"id"`
}

func IPsecAuthRemoteResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an IPsec remote authentication in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"enabled":        schema.BoolAttribute{MarkdownDescription: "Enable. Defaults to `true`.", Optional: true, Computed: true, Default: booldefault.StaticBool(true)},
			"connection":     schema.StringAttribute{MarkdownDescription: "Connection reference.", Required: true},
			"round":          schema.StringAttribute{MarkdownDescription: "Authentication round.", Optional: true},
			"authentication": schema.StringAttribute{MarkdownDescription: "Authentication method.", Optional: true, Computed: true},
			"auth_id":        schema.StringAttribute{MarkdownDescription: "Identity.", Optional: true},
			"eap_id":         schema.StringAttribute{MarkdownDescription: "EAP identity.", Optional: true},
			"certificates":   schema.SetAttribute{MarkdownDescription: "Certificates.", Optional: true, Computed: true, ElementType: types.StringType, Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType))},
			"public_keys":    schema.SetAttribute{MarkdownDescription: "Public keys.", Optional: true, Computed: true, ElementType: types.StringType, Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType))},
			"description":    schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{Computed: true, MarkdownDescription: "UUID of the resource.", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
		},
	}
}

func IPsecAuthRemoteDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an IPsec remote authentication.",
		Attributes: map[string]dschema.Attribute{
			"id":             dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"enabled":        dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether enabled."},
			"connection":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Connection."},
			"round":          dschema.StringAttribute{Computed: true, MarkdownDescription: "Round."},
			"authentication": dschema.StringAttribute{Computed: true, MarkdownDescription: "Authentication method."},
			"auth_id":        dschema.StringAttribute{Computed: true, MarkdownDescription: "Identity."},
			"eap_id":         dschema.StringAttribute{Computed: true, MarkdownDescription: "EAP identity."},
			"certificates":   dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Certificates."},
			"public_keys":    dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Public keys."},
			"description":    dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertIPsecAuthRemoteSchemaToStruct(d *IPsecAuthRemoteResourceModel) (*ipsec.IPsecAuthRemote, error) {
	return &ipsec.IPsecAuthRemote{
		Enabled:        tools.BoolToString(d.Enabled.ValueBool()),
		Connection:     api.SelectedMap(d.Connection.ValueString()),
		Round:          d.Round.ValueString(),
		Authentication: api.SelectedMap(d.Authentication.ValueString()),
		Id:             d.AuthId.ValueString(),
		EAPId:          d.EAPId.ValueString(),
		Certificates:   setToSlice(d.Certificates),
		PublicKeys:     setToSlice(d.PublicKeys),
		Description:    d.Description.ValueString(),
	}, nil
}

func convertIPsecAuthRemoteStructToSchema(d *ipsec.IPsecAuthRemote) (*IPsecAuthRemoteResourceModel, error) {
	return &IPsecAuthRemoteResourceModel{
		Enabled:        types.BoolValue(tools.StringToBool(d.Enabled)),
		Connection:     types.StringValue(d.Connection.String()),
		Round:          tools.StringOrNull(d.Round),
		Authentication: types.StringValue(d.Authentication.String()),
		AuthId:         tools.StringOrNull(d.Id),
		EAPId:          tools.StringOrNull(d.EAPId),
		Certificates:   sliceToSet(d.Certificates),
		PublicKeys:     sliceToSet(d.PublicKeys),
		Description:    tools.StringOrNull(d.Description),
	}, nil
}
