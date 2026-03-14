package service

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
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

type DnsmasqHostResourceModel struct {
	Hostname          types.String `tfsdk:"hostname"`
	Domain            types.String `tfsdk:"domain"`
	IsLocalDomain     types.Bool   `tfsdk:"is_local_domain"`
	IpAddresses       types.Set    `tfsdk:"ip_addresses"`
	AliasRecords      types.Set    `tfsdk:"alias_records"`
	CnameRecords      types.Set    `tfsdk:"cname_records"`
	ClientId          types.String `tfsdk:"client_id"`
	HardwareAddresses types.Set    `tfsdk:"hardware_addresses"`
	Tag               types.String `tfsdk:"tag"`
	IsIgnored         types.Bool   `tfsdk:"is_ignored"`
	Description       types.String `tfsdk:"description"`
	Comments          types.String `tfsdk:"comments"`
	Id                types.String `tfsdk:"id"`
}

func DnsmasqHostResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a Dnsmasq host entry in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"hostname":    schema.StringAttribute{MarkdownDescription: "Hostname.", Required: true},
			"domain":      schema.StringAttribute{MarkdownDescription: "Domain.", Optional: true},
			"is_local_domain": schema.BoolAttribute{
				MarkdownDescription: "Whether this is a local domain. Defaults to `false`.",
				Optional: true, Computed: true, Default: booldefault.StaticBool(false),
			},
			"ip_addresses": schema.SetAttribute{
				MarkdownDescription: "IP addresses.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"alias_records": schema.SetAttribute{
				MarkdownDescription: "Alias records.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"cname_records": schema.SetAttribute{
				MarkdownDescription: "CNAME records.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"client_id":          schema.StringAttribute{MarkdownDescription: "Client ID.", Optional: true},
			"hardware_addresses": schema.SetAttribute{
				MarkdownDescription: "Hardware (MAC) addresses.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tag":         schema.StringAttribute{MarkdownDescription: "Tag.", Optional: true, Computed: true},
			"is_ignored":  schema.BoolAttribute{
				MarkdownDescription: "Whether this host is ignored. Defaults to `false`.",
				Optional: true, Computed: true, Default: booldefault.StaticBool(false),
			},
			"description": schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"comments":    schema.StringAttribute{MarkdownDescription: "Comments.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func DnsmasqHostDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a Dnsmasq host entry.",
		Attributes: map[string]dschema.Attribute{
			"id":                 dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"hostname":           dschema.StringAttribute{Computed: true, MarkdownDescription: "Hostname."},
			"domain":             dschema.StringAttribute{Computed: true, MarkdownDescription: "Domain."},
			"is_local_domain":    dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether local domain."},
			"ip_addresses":       dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "IP addresses."},
			"alias_records":      dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Alias records."},
			"cname_records":      dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "CNAME records."},
			"client_id":          dschema.StringAttribute{Computed: true, MarkdownDescription: "Client ID."},
			"hardware_addresses": dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Hardware addresses."},
			"tag":                dschema.StringAttribute{Computed: true, MarkdownDescription: "Tag."},
			"is_ignored":         dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether ignored."},
			"description":        dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
			"comments":           dschema.StringAttribute{Computed: true, MarkdownDescription: "Comments."},
		},
	}
}

func convertDnsmasqHostSchemaToStruct(d *DnsmasqHostResourceModel) (*dnsmasq.Host, error) {
	var ipList []string
	d.IpAddresses.ElementsAs(context.Background(), &ipList, false)
	var aliasList []string
	d.AliasRecords.ElementsAs(context.Background(), &aliasList, false)
	var cnameList []string
	d.CnameRecords.ElementsAs(context.Background(), &cnameList, false)
	var hwList []string
	d.HardwareAddresses.ElementsAs(context.Background(), &hwList, false)

	return &dnsmasq.Host{
		Hostname:          d.Hostname.ValueString(),
		Domain:            d.Domain.ValueString(),
		IsLocalDomain:     tools.BoolToString(d.IsLocalDomain.ValueBool()),
		IpAddresses:       ipList,
		AliasRecords:      aliasList,
		CnameRecords:      cnameList,
		ClientId:          d.ClientId.ValueString(),
		HardwareAddresses: hwList,
		Tag:               api.SelectedMap(d.Tag.ValueString()),
		IsIgnored:         tools.BoolToString(d.IsIgnored.ValueBool()),
		Description:       d.Description.ValueString(),
		Comments:          d.Comments.ValueString(),
	}, nil
}

func convertDnsmasqHostStructToSchema(d *dnsmasq.Host) (*DnsmasqHostResourceModel, error) {
	model := &DnsmasqHostResourceModel{
		Hostname:      types.StringValue(d.Hostname),
		Domain:        tools.StringOrNull(d.Domain),
		IsLocalDomain: types.BoolValue(tools.StringToBool(d.IsLocalDomain)),
		ClientId:      tools.StringOrNull(d.ClientId),
		Tag:           types.StringValue(d.Tag.String()),
		IsIgnored:     types.BoolValue(tools.StringToBool(d.IsIgnored)),
		Description:   tools.StringOrNull(d.Description),
		Comments:      tools.StringOrNull(d.Comments),
	}

	setFromSlice := func(s []string) types.Set {
		var list []attr.Value
		for _, i := range s {
			if i == "" { continue }
			list = append(list, basetypes.NewStringValue(i))
		}
		v, _ := types.SetValue(types.StringType, list)
		return v
	}

	model.IpAddresses = setFromSlice(d.IpAddresses)
	model.AliasRecords = setFromSlice(d.AliasRecords)
	model.CnameRecords = setFromSlice(d.CnameRecords)
	model.HardwareAddresses = setFromSlice(d.HardwareAddresses)

	return model, nil
}
