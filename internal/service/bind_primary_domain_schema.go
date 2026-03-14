package service

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/bind"
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

type BindPrimaryDomainResourceModel struct {
	DomainName    types.String `tfsdk:"domain_name"`
	Enabled       types.Bool   `tfsdk:"enabled"`
	AllowTransfer types.Set    `tfsdk:"allow_transfer"`
	AllowQuery    types.Set    `tfsdk:"allow_query"`
	TimeToLive    types.String `tfsdk:"ttl"`
	Refresh       types.String `tfsdk:"refresh"`
	Retry         types.String `tfsdk:"retry"`
	Expire        types.String `tfsdk:"expire"`
	Negative      types.String `tfsdk:"negative"`
	MailAdmin     types.String `tfsdk:"mail_admin"`
	DnsServer     types.String `tfsdk:"dns_server"`
	Id            types.String `tfsdk:"id"`
}

func BindPrimaryDomainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a BIND primary (master) DNS domain in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"domain_name": schema.StringAttribute{
				MarkdownDescription: "The domain name.",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this domain. Defaults to `true`.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(true),
			},
			"allow_transfer": schema.SetAttribute{
				MarkdownDescription: "ACLs allowed to transfer this zone.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"allow_query": schema.SetAttribute{
				MarkdownDescription: "ACLs allowed to query this zone.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ttl": schema.StringAttribute{
				MarkdownDescription: "Default TTL.", Optional: true,
			},
			"refresh": schema.StringAttribute{
				MarkdownDescription: "SOA refresh interval.", Optional: true,
			},
			"retry": schema.StringAttribute{
				MarkdownDescription: "SOA retry interval.", Optional: true,
			},
			"expire": schema.StringAttribute{
				MarkdownDescription: "SOA expire time.", Optional: true,
			},
			"negative": schema.StringAttribute{
				MarkdownDescription: "SOA negative cache TTL.", Optional: true,
			},
			"mail_admin": schema.StringAttribute{
				MarkdownDescription: "Mail admin for the SOA record.", Optional: true,
			},
			"dns_server": schema.StringAttribute{
				MarkdownDescription: "Primary DNS server for the SOA record.", Optional: true,
			},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func BindPrimaryDomainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a BIND primary domain.",
		Attributes: map[string]dschema.Attribute{
			"id":             dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"domain_name":    dschema.StringAttribute{Computed: true, MarkdownDescription: "The domain name."},
			"enabled":        dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether this domain is enabled."},
			"allow_transfer": dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "ACLs allowed to transfer."},
			"allow_query":    dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "ACLs allowed to query."},
			"ttl":            dschema.StringAttribute{Computed: true, MarkdownDescription: "Default TTL."},
			"refresh":        dschema.StringAttribute{Computed: true, MarkdownDescription: "SOA refresh."},
			"retry":          dschema.StringAttribute{Computed: true, MarkdownDescription: "SOA retry."},
			"expire":         dschema.StringAttribute{Computed: true, MarkdownDescription: "SOA expire."},
			"negative":       dschema.StringAttribute{Computed: true, MarkdownDescription: "SOA negative cache TTL."},
			"mail_admin":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Mail admin."},
			"dns_server":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Primary DNS server."},
		},
	}
}

func convertBindPrimaryDomainSchemaToStruct(d *BindPrimaryDomainResourceModel) (*bind.PrimaryDomain, error) {
	var allowTransferList []string
	d.AllowTransfer.ElementsAs(context.Background(), &allowTransferList, false)
	var allowQueryList []string
	d.AllowQuery.ElementsAs(context.Background(), &allowQueryList, false)

	return &bind.PrimaryDomain{
		DomainName:    d.DomainName.ValueString(),
		Enabled:       tools.BoolToString(d.Enabled.ValueBool()),
		AllowTransfer: allowTransferList,
		AllowQuery:    allowQueryList,
		TimeToLive:    d.TimeToLive.ValueString(),
		Refresh:       d.Refresh.ValueString(),
		Retry:         d.Retry.ValueString(),
		Expire:        d.Expire.ValueString(),
		Negative:      d.Negative.ValueString(),
		MailAdmin:     d.MailAdmin.ValueString(),
		DnsServer:     d.DnsServer.ValueString(),
	}, nil
}

func convertBindPrimaryDomainStructToSchema(d *bind.PrimaryDomain) (*BindPrimaryDomainResourceModel, error) {
	model := &BindPrimaryDomainResourceModel{
		DomainName: types.StringValue(d.DomainName),
		Enabled:    types.BoolValue(tools.StringToBool(d.Enabled)),
		TimeToLive: tools.StringOrNull(d.TimeToLive),
		Refresh:    tools.StringOrNull(d.Refresh),
		Retry:      tools.StringOrNull(d.Retry),
		Expire:     tools.StringOrNull(d.Expire),
		Negative:   tools.StringOrNull(d.Negative),
		MailAdmin:  tools.StringOrNull(d.MailAdmin),
		DnsServer:  tools.StringOrNull(d.DnsServer),
	}

	var atList []attr.Value
	for _, i := range d.AllowTransfer {
		if i == "" { continue }
		atList = append(atList, basetypes.NewStringValue(i))
	}
	atSet, _ := types.SetValue(types.StringType, atList)
	model.AllowTransfer = atSet

	var aqList []attr.Value
	for _, i := range d.AllowQuery {
		if i == "" { continue }
		aqList = append(aqList, basetypes.NewStringValue(i))
	}
	aqSet, _ := types.SetValue(types.StringType, aqList)
	model.AllowQuery = aqSet

	return model, nil
}
