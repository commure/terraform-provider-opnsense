package bind

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/bind"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type primaryDomainResourceModel struct {
	Enabled       types.Bool   `tfsdk:"enabled"`
	DomainName    types.String `tfsdk:"domain_name"`
	AllowTransfer types.Set    `tfsdk:"allow_transfer"`
	AllowQuery    types.Set    `tfsdk:"allow_query"`
	TTL           types.String `tfsdk:"ttl"`
	Refresh       types.String `tfsdk:"refresh"`
	Retry         types.String `tfsdk:"retry"`
	Expire        types.String `tfsdk:"expire"`
	Negative      types.String `tfsdk:"negative"`
	MailAdmin     types.String `tfsdk:"mail_admin"`
	DnsServer     types.String `tfsdk:"dns_server"`

	Id types.String `tfsdk:"id"`
}

func primaryDomainResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense BIND primary domain.",

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this primary domain. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"domain_name": schema.StringAttribute{
				MarkdownDescription: "The domain name.",
				Required:            true,
			},
			"allow_transfer": schema.SetAttribute{
				MarkdownDescription: "ACLs allowed to transfer. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"allow_query": schema.SetAttribute{
				MarkdownDescription: "ACLs allowed to query. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"ttl": schema.StringAttribute{
				MarkdownDescription: "Time to live.",
				Optional:            true,
			},
			"refresh": schema.StringAttribute{
				MarkdownDescription: "Refresh interval.",
				Optional:            true,
			},
			"retry": schema.StringAttribute{
				MarkdownDescription: "Retry interval.",
				Optional:            true,
			},
			"expire": schema.StringAttribute{
				MarkdownDescription: "Expire time.",
				Optional:            true,
			},
			"negative": schema.StringAttribute{
				MarkdownDescription: "Negative cache TTL.",
				Optional:            true,
			},
			"mail_admin": schema.StringAttribute{
				MarkdownDescription: "Mail admin address.",
				Optional:            true,
			},
			"dns_server": schema.StringAttribute{
				MarkdownDescription: "DNS server hostname.",
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

func primaryDomainDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense BIND primary domain.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether this primary domain is enabled.",
				Computed:            true,
			},
			"domain_name": dschema.StringAttribute{
				MarkdownDescription: "The domain name.",
				Computed:            true,
			},
			"allow_transfer": dschema.SetAttribute{
				MarkdownDescription: "ACLs allowed to transfer.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"allow_query": dschema.SetAttribute{
				MarkdownDescription: "ACLs allowed to query.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"ttl": dschema.StringAttribute{
				MarkdownDescription: "Time to live.",
				Computed:            true,
			},
			"refresh": dschema.StringAttribute{
				MarkdownDescription: "Refresh interval.",
				Computed:            true,
			},
			"retry": dschema.StringAttribute{
				MarkdownDescription: "Retry interval.",
				Computed:            true,
			},
			"expire": dschema.StringAttribute{
				MarkdownDescription: "Expire time.",
				Computed:            true,
			},
			"negative": dschema.StringAttribute{
				MarkdownDescription: "Negative cache TTL.",
				Computed:            true,
			},
			"mail_admin": dschema.StringAttribute{
				MarkdownDescription: "Mail admin address.",
				Computed:            true,
			},
			"dns_server": dschema.StringAttribute{
				MarkdownDescription: "DNS server hostname.",
				Computed:            true,
			},
		},
	}
}

func convertPrimaryDomainSchemaToStruct(d *primaryDomainResourceModel) (*bind.PrimaryDomain, error) {
	var allowTransferList []string
	d.AllowTransfer.ElementsAs(context.Background(), &allowTransferList, false)

	var allowQueryList []string
	d.AllowQuery.ElementsAs(context.Background(), &allowQueryList, false)

	return &bind.PrimaryDomain{
		Enabled:       tools.BoolToString(d.Enabled.ValueBool()),
		DomainName:    d.DomainName.ValueString(),
		AllowTransfer: allowTransferList,
		AllowQuery:    allowQueryList,
		TimeToLive:    d.TTL.ValueString(),
		Refresh:       d.Refresh.ValueString(),
		Retry:         d.Retry.ValueString(),
		Expire:        d.Expire.ValueString(),
		Negative:      d.Negative.ValueString(),
		MailAdmin:     d.MailAdmin.ValueString(),
		DnsServer:     d.DnsServer.ValueString(),
	}, nil
}

func convertPrimaryDomainStructToSchema(d *bind.PrimaryDomain) (*primaryDomainResourceModel, error) {
	model := &primaryDomainResourceModel{
		Enabled:    types.BoolValue(tools.StringToBool(d.Enabled)),
		DomainName: types.StringValue(d.DomainName),
		TTL:        tools.StringOrNull(d.TimeToLive),
		Refresh:    tools.StringOrNull(d.Refresh),
		Retry:      tools.StringOrNull(d.Retry),
		Expire:     tools.StringOrNull(d.Expire),
		Negative:   tools.StringOrNull(d.Negative),
		MailAdmin:  tools.StringOrNull(d.MailAdmin),
		DnsServer:  tools.StringOrNull(d.DnsServer),
	}

	var allowTransferList []attr.Value
	for _, i := range d.AllowTransfer {
		if i == "" {
			continue
		}
		allowTransferList = append(allowTransferList, basetypes.NewStringValue(i))
	}
	allowTransferTypeList, _ := types.SetValue(types.StringType, allowTransferList)
	model.AllowTransfer = allowTransferTypeList

	var allowQueryList []attr.Value
	for _, i := range d.AllowQuery {
		if i == "" {
			continue
		}
		allowQueryList = append(allowQueryList, basetypes.NewStringValue(i))
	}
	allowQueryTypeList, _ := types.SetValue(types.StringType, allowQueryList)
	model.AllowQuery = allowQueryTypeList

	return model, nil
}
