package dnsmasq

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
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

type hostResourceModel struct {
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

	Id types.String `tfsdk:"id"`
}

func hostResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense Dnsmasq host entry.",

		Attributes: map[string]schema.Attribute{
			"hostname": schema.StringAttribute{
				MarkdownDescription: "The hostname.",
				Optional:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The domain.",
				Optional:            true,
			},
			"is_local_domain": schema.BoolAttribute{
				MarkdownDescription: "Whether this is a local domain. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"ip_addresses": schema.SetAttribute{
				MarkdownDescription: "Set of IP addresses. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"alias_records": schema.SetAttribute{
				MarkdownDescription: "Set of alias records. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"cname_records": schema.SetAttribute{
				MarkdownDescription: "Set of CNAME records. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"client_id": schema.StringAttribute{
				MarkdownDescription: "The client ID.",
				Optional:            true,
			},
			"hardware_addresses": schema.SetAttribute{
				MarkdownDescription: "Set of hardware (MAC) addresses. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"tag": schema.StringAttribute{
				MarkdownDescription: "The tag for this host.",
				Optional:            true,
			},
			"is_ignored": schema.BoolAttribute{
				MarkdownDescription: "Whether this host is ignored. Defaults to `false`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this host entry.",
				Optional:            true,
			},
			"comments": schema.StringAttribute{
				MarkdownDescription: "Comments for this host entry.",
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

func hostDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense Dnsmasq host entry.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"hostname": dschema.StringAttribute{
				MarkdownDescription: "The hostname.",
				Computed:            true,
			},
			"domain": dschema.StringAttribute{
				MarkdownDescription: "The domain.",
				Computed:            true,
			},
			"is_local_domain": dschema.BoolAttribute{
				MarkdownDescription: "Whether this is a local domain.",
				Computed:            true,
			},
			"ip_addresses": dschema.SetAttribute{
				MarkdownDescription: "Set of IP addresses.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"alias_records": dschema.SetAttribute{
				MarkdownDescription: "Set of alias records.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"cname_records": dschema.SetAttribute{
				MarkdownDescription: "Set of CNAME records.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"client_id": dschema.StringAttribute{
				MarkdownDescription: "The client ID.",
				Computed:            true,
			},
			"hardware_addresses": dschema.SetAttribute{
				MarkdownDescription: "Set of hardware (MAC) addresses.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"tag": dschema.StringAttribute{
				MarkdownDescription: "The tag for this host.",
				Computed:            true,
			},
			"is_ignored": dschema.BoolAttribute{
				MarkdownDescription: "Whether this host is ignored.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this host entry.",
				Computed:            true,
			},
			"comments": dschema.StringAttribute{
				MarkdownDescription: "Comments for this host entry.",
				Computed:            true,
			},
		},
	}
}

func convertHostSchemaToStruct(d *hostResourceModel) (*dnsmasq.Host, error) {
	var ipAddressesList []string
	d.IpAddresses.ElementsAs(context.Background(), &ipAddressesList, false)

	var aliasRecordsList []string
	d.AliasRecords.ElementsAs(context.Background(), &aliasRecordsList, false)

	var cnameRecordsList []string
	d.CnameRecords.ElementsAs(context.Background(), &cnameRecordsList, false)

	var hardwareAddressesList []string
	d.HardwareAddresses.ElementsAs(context.Background(), &hardwareAddressesList, false)

	return &dnsmasq.Host{
		Hostname:          d.Hostname.ValueString(),
		Domain:            d.Domain.ValueString(),
		IsLocalDomain:     tools.BoolToString(d.IsLocalDomain.ValueBool()),
		IpAddresses:       ipAddressesList,
		AliasRecords:      aliasRecordsList,
		CnameRecords:      cnameRecordsList,
		ClientId:          d.ClientId.ValueString(),
		HardwareAddresses: hardwareAddressesList,
		Tag:               api.SelectedMap(d.Tag.ValueString()),
		IsIgnored:         tools.BoolToString(d.IsIgnored.ValueBool()),
		Description:       d.Description.ValueString(),
		Comments:          d.Comments.ValueString(),
	}, nil
}

func convertHostStructToSchema(d *dnsmasq.Host) (*hostResourceModel, error) {
	model := &hostResourceModel{
		Hostname:      tools.StringOrNull(d.Hostname),
		Domain:        tools.StringOrNull(d.Domain),
		IsLocalDomain: types.BoolValue(tools.StringToBool(d.IsLocalDomain)),
		ClientId:      tools.StringOrNull(d.ClientId),
		Tag:           tools.StringOrNull(d.Tag.String()),
		IsIgnored:     types.BoolValue(tools.StringToBool(d.IsIgnored)),
		Description:   tools.StringOrNull(d.Description),
		Comments:      tools.StringOrNull(d.Comments),
	}

	// Parse 'IpAddresses'
	var ipAddressesList []attr.Value
	for _, i := range d.IpAddresses {
		if i == "" {
			continue
		}
		ipAddressesList = append(ipAddressesList, basetypes.NewStringValue(i))
	}
	ipAddressesTypeList, _ := types.SetValue(types.StringType, ipAddressesList)
	model.IpAddresses = ipAddressesTypeList

	// Parse 'AliasRecords'
	var aliasRecordsList []attr.Value
	for _, i := range d.AliasRecords {
		if i == "" {
			continue
		}
		aliasRecordsList = append(aliasRecordsList, basetypes.NewStringValue(i))
	}
	aliasRecordsTypeList, _ := types.SetValue(types.StringType, aliasRecordsList)
	model.AliasRecords = aliasRecordsTypeList

	// Parse 'CnameRecords'
	var cnameRecordsList []attr.Value
	for _, i := range d.CnameRecords {
		if i == "" {
			continue
		}
		cnameRecordsList = append(cnameRecordsList, basetypes.NewStringValue(i))
	}
	cnameRecordsTypeList, _ := types.SetValue(types.StringType, cnameRecordsList)
	model.CnameRecords = cnameRecordsTypeList

	// Parse 'HardwareAddresses'
	var hardwareAddressesList []attr.Value
	for _, i := range d.HardwareAddresses {
		if i == "" {
			continue
		}
		hardwareAddressesList = append(hardwareAddressesList, basetypes.NewStringValue(i))
	}
	hardwareAddressesTypeList, _ := types.SetValue(types.StringType, hardwareAddressesList)
	model.HardwareAddresses = hardwareAddressesTypeList

	return model, nil
}
