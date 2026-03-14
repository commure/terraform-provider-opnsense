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

type DnsmasqOptionResourceModel struct {
	Type         types.String `tfsdk:"type"`
	OptionV4     types.String `tfsdk:"option_v4"`
	OptionV6     types.String `tfsdk:"option_v6"`
	Interface    types.String `tfsdk:"interface"`
	TypeSetTags  types.Set    `tfsdk:"type_set_tags"`
	TypeMatchTag types.String `tfsdk:"type_match_tag"`
	Value        types.String `tfsdk:"value"`
	Force        types.Bool   `tfsdk:"force"`
	Description  types.String `tfsdk:"description"`
	Id           types.String `tfsdk:"id"`
}

func DnsmasqOptionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a Dnsmasq DHCP option in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"type":           schema.StringAttribute{MarkdownDescription: "Option type.", Optional: true, Computed: true},
			"option_v4":      schema.StringAttribute{MarkdownDescription: "DHCP option (IPv4).", Optional: true, Computed: true},
			"option_v6":      schema.StringAttribute{MarkdownDescription: "DHCP option (IPv6).", Optional: true, Computed: true},
			"interface":      schema.StringAttribute{MarkdownDescription: "Interface.", Optional: true, Computed: true},
			"type_set_tags": schema.SetAttribute{
				MarkdownDescription: "Tags to set.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"type_match_tag": schema.StringAttribute{MarkdownDescription: "Tag to match.", Optional: true, Computed: true},
			"value":          schema.StringAttribute{MarkdownDescription: "Option value.", Optional: true},
			"force": schema.BoolAttribute{
				MarkdownDescription: "Force this option. Defaults to `false`.",
				Optional: true, Computed: true,
			},
			"description": schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func DnsmasqOptionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a Dnsmasq DHCP option.",
		Attributes: map[string]dschema.Attribute{
			"id":             dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"type":           dschema.StringAttribute{Computed: true, MarkdownDescription: "Option type."},
			"option_v4":      dschema.StringAttribute{Computed: true, MarkdownDescription: "DHCP option (IPv4)."},
			"option_v6":      dschema.StringAttribute{Computed: true, MarkdownDescription: "DHCP option (IPv6)."},
			"interface":      dschema.StringAttribute{Computed: true, MarkdownDescription: "Interface."},
			"type_set_tags":  dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Tags to set."},
			"type_match_tag": dschema.StringAttribute{Computed: true, MarkdownDescription: "Tag to match."},
			"value":          dschema.StringAttribute{Computed: true, MarkdownDescription: "Option value."},
			"force":          dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether forced."},
			"description":    dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertDnsmasqOptionSchemaToStruct(d *DnsmasqOptionResourceModel) (*dnsmasq.Option, error) {
	var tagList []string
	d.TypeSetTags.ElementsAs(context.Background(), &tagList, false)
	return &dnsmasq.Option{
		Type:         api.SelectedMap(d.Type.ValueString()),
		OptionV4:     api.SelectedMap(d.OptionV4.ValueString()),
		OptionV6:     api.SelectedMap(d.OptionV6.ValueString()),
		Interface:    api.SelectedMap(d.Interface.ValueString()),
		TypeSetTags:  tagList,
		TypeMatchTag: api.SelectedMap(d.TypeMatchTag.ValueString()),
		Value:        d.Value.ValueString(),
		Force:        tools.BoolToString(d.Force.ValueBool()),
		Description:  d.Description.ValueString(),
	}, nil
}

func convertDnsmasqOptionStructToSchema(d *dnsmasq.Option) (*DnsmasqOptionResourceModel, error) {
	model := &DnsmasqOptionResourceModel{
		Type:         types.StringValue(d.Type.String()),
		OptionV4:     types.StringValue(d.OptionV4.String()),
		OptionV6:     types.StringValue(d.OptionV6.String()),
		Interface:    types.StringValue(d.Interface.String()),
		TypeMatchTag: types.StringValue(d.TypeMatchTag.String()),
		Value:        tools.StringOrNull(d.Value),
		Force:        types.BoolValue(tools.StringToBool(d.Force)),
		Description:  tools.StringOrNull(d.Description),
	}
	var tagList []attr.Value
	for _, i := range d.TypeSetTags {
		if i == "" { continue }
		tagList = append(tagList, basetypes.NewStringValue(i))
	}
	tagSet, _ := types.SetValue(types.StringType, tagList)
	model.TypeSetTags = tagSet
	return model, nil
}
