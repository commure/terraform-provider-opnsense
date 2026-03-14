package service

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/ipsec"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

type IPsecPSKResourceModel struct {
	IdentityLocal  types.String `tfsdk:"identity_local"`
	IdentityRemote types.String `tfsdk:"identity_remote"`
	PreSharedKey   types.String `tfsdk:"pre_shared_key"`
	Type           types.String `tfsdk:"type"`
	Description    types.String `tfsdk:"description"`
	Id             types.String `tfsdk:"id"`
}

func IPsecPSKResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an IPsec pre-shared key in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"identity_local":  schema.StringAttribute{MarkdownDescription: "Local identity.", Required: true},
			"identity_remote": schema.StringAttribute{MarkdownDescription: "Remote identity.", Optional: true},
			"pre_shared_key":  schema.StringAttribute{MarkdownDescription: "Pre-shared key.", Required: true, Sensitive: true},
			"type":            schema.StringAttribute{MarkdownDescription: "Key type.", Optional: true, Computed: true},
			"description":     schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{Computed: true, MarkdownDescription: "UUID of the resource.", PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()}},
		},
	}
}

func IPsecPSKDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about an IPsec pre-shared key.",
		Attributes: map[string]dschema.Attribute{
			"id":              dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"identity_local":  dschema.StringAttribute{Computed: true, MarkdownDescription: "Local identity."},
			"identity_remote": dschema.StringAttribute{Computed: true, MarkdownDescription: "Remote identity."},
			"pre_shared_key":  dschema.StringAttribute{Computed: true, Sensitive: true, MarkdownDescription: "Pre-shared key."},
			"type":            dschema.StringAttribute{Computed: true, MarkdownDescription: "Key type."},
			"description":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertIPsecPSKSchemaToStruct(d *IPsecPSKResourceModel) (*ipsec.IPsecPSK, error) {
	return &ipsec.IPsecPSK{
		IdentityLocal:  d.IdentityLocal.ValueString(),
		IdentityRemote: d.IdentityRemote.ValueString(),
		PreSharedKey:   d.PreSharedKey.ValueString(),
		Type:           api.SelectedMap(d.Type.ValueString()),
		Description:    d.Description.ValueString(),
	}, nil
}

func convertIPsecPSKStructToSchema(d *ipsec.IPsecPSK) (*IPsecPSKResourceModel, error) {
	return &IPsecPSKResourceModel{
		IdentityLocal:  types.StringValue(d.IdentityLocal),
		IdentityRemote: tools.StringOrNull(d.IdentityRemote),
		PreSharedKey:   types.StringValue(d.PreSharedKey),
		Type:           types.StringValue(d.Type.String()),
		Description:    tools.StringOrNull(d.Description),
	}, nil
}
