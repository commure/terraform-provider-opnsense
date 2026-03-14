package service

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"terraform-provider-opnsense/internal/tools"
)

type DnsmasqBootResourceModel struct {
	Interface     types.String `tfsdk:"interface"`
	Tag           types.Set    `tfsdk:"tag"`
	Filename      types.String `tfsdk:"filename"`
	Servername    types.String `tfsdk:"servername"`
	ServerAddress types.String `tfsdk:"server_address"`
	Description   types.String `tfsdk:"description"`
	Id            types.String `tfsdk:"id"`
}

func DnsmasqBootResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a Dnsmasq PXE boot entry in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"interface": schema.StringAttribute{
				MarkdownDescription: "Interface for this boot entry.",
				Optional:            true,
				Computed:            true,
			},
			"tag": schema.SetAttribute{
				MarkdownDescription: "Tag(s) for this boot entry.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"filename": schema.StringAttribute{
				MarkdownDescription: "Boot filename.",
				Optional:            true,
			},
			"servername": schema.StringAttribute{
				MarkdownDescription: "Server name.",
				Optional:            true,
			},
			"server_address": schema.StringAttribute{
				MarkdownDescription: "Server address.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description.",
				Optional:            true,
			},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func DnsmasqBootDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a Dnsmasq PXE boot entry.",
		Attributes: map[string]dschema.Attribute{
			"id":             dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"interface":      dschema.StringAttribute{Computed: true, MarkdownDescription: "Interface."},
			"tag":            dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Tag(s)."},
			"filename":       dschema.StringAttribute{Computed: true, MarkdownDescription: "Boot filename."},
			"servername":     dschema.StringAttribute{Computed: true, MarkdownDescription: "Server name."},
			"server_address": dschema.StringAttribute{Computed: true, MarkdownDescription: "Server address."},
			"description":    dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertDnsmasqBootSchemaToStruct(d *DnsmasqBootResourceModel) (*dnsmasq.Boot, error) {
	var tagList []string
	d.Tag.ElementsAs(context.Background(), &tagList, false)
	return &dnsmasq.Boot{
		Interface:     api.SelectedMap(d.Interface.ValueString()),
		Tag:           tagList,
		Filename:      d.Filename.ValueString(),
		Servername:    d.Servername.ValueString(),
		ServerAddress: d.ServerAddress.ValueString(),
		Description:   d.Description.ValueString(),
	}, nil
}

func convertDnsmasqBootStructToSchema(d *dnsmasq.Boot) (*DnsmasqBootResourceModel, error) {
	model := &DnsmasqBootResourceModel{
		Interface:     types.StringValue(d.Interface.String()),
		Filename:      tools.StringOrNull(d.Filename),
		Servername:    tools.StringOrNull(d.Servername),
		ServerAddress: tools.StringOrNull(d.ServerAddress),
		Description:   tools.StringOrNull(d.Description),
	}
	var tagList []attr.Value
	for _, i := range d.Tag {
		if i == "" { continue }
		tagList = append(tagList, basetypes.NewStringValue(i))
	}
	tagSet, _ := types.SetValue(types.StringType, tagList)
	model.Tag = tagSet
	return model, nil
}
