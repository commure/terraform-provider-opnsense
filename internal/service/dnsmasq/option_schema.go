package dnsmasq

import (
	"context"

	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/dnsmasq"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type optionResourceModel struct {
	Type         types.String `tfsdk:"type"`
	OptionV4     types.String `tfsdk:"option_v4"`
	OptionV6     types.String `tfsdk:"option_v6"`
	Interface    types.String `tfsdk:"interface"`
	TypeSetTags  types.Set    `tfsdk:"type_set_tags"`
	TypeMatchTag types.String `tfsdk:"type_match_tag"`
	Value        types.String `tfsdk:"value"`
	Force        types.String `tfsdk:"force"`
	Description  types.String `tfsdk:"description"`

	Id types.String `tfsdk:"id"`
}

func optionResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense Dnsmasq DHCP option.",

		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				MarkdownDescription: "The option type.",
				Optional:            true,
			},
			"option_v4": schema.StringAttribute{
				MarkdownDescription: "The IPv4 DHCP option.",
				Optional:            true,
			},
			"option_v6": schema.StringAttribute{
				MarkdownDescription: "The IPv6 DHCP option.",
				Optional:            true,
			},
			"interface": schema.StringAttribute{
				MarkdownDescription: "The interface.",
				Optional:            true,
			},
			"type_set_tags": schema.SetAttribute{
				MarkdownDescription: "Tags to set. Defaults to `[]`.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.StringType,
				Default:             setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"type_match_tag": schema.StringAttribute{
				MarkdownDescription: "Tag to match.",
				Optional:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The option value.",
				Optional:            true,
			},
			"force": schema.StringAttribute{
				MarkdownDescription: "Force this option.",
				Optional:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Description for this option.",
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

func optionDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense Dnsmasq DHCP option.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"type": dschema.StringAttribute{
				MarkdownDescription: "The option type.",
				Computed:            true,
			},
			"option_v4": dschema.StringAttribute{
				MarkdownDescription: "The IPv4 DHCP option.",
				Computed:            true,
			},
			"option_v6": dschema.StringAttribute{
				MarkdownDescription: "The IPv6 DHCP option.",
				Computed:            true,
			},
			"interface": dschema.StringAttribute{
				MarkdownDescription: "The interface.",
				Computed:            true,
			},
			"type_set_tags": dschema.SetAttribute{
				MarkdownDescription: "Tags to set.",
				Computed:            true,
				ElementType:         types.StringType,
			},
			"type_match_tag": dschema.StringAttribute{
				MarkdownDescription: "Tag to match.",
				Computed:            true,
			},
			"value": dschema.StringAttribute{
				MarkdownDescription: "The option value.",
				Computed:            true,
			},
			"force": dschema.StringAttribute{
				MarkdownDescription: "Force this option.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description for this option.",
				Computed:            true,
			},
		},
	}
}

func convertOptionSchemaToStruct(d *optionResourceModel) (*dnsmasq.Option, error) {
	var typeSetTagsList []string
	d.TypeSetTags.ElementsAs(context.Background(), &typeSetTagsList, false)

	return &dnsmasq.Option{
		Type:         api.SelectedMap(d.Type.ValueString()),
		OptionV4:     api.SelectedMap(d.OptionV4.ValueString()),
		OptionV6:     api.SelectedMap(d.OptionV6.ValueString()),
		Interface:    api.SelectedMap(d.Interface.ValueString()),
		TypeSetTags:  typeSetTagsList,
		TypeMatchTag: api.SelectedMap(d.TypeMatchTag.ValueString()),
		Value:        d.Value.ValueString(),
		Force:        d.Force.ValueString(),
		Description:  d.Description.ValueString(),
	}, nil
}

func convertOptionStructToSchema(d *dnsmasq.Option) (*optionResourceModel, error) {
	model := &optionResourceModel{
		Type:         tools.StringOrNull(d.Type.String()),
		OptionV4:     tools.StringOrNull(d.OptionV4.String()),
		OptionV6:     tools.StringOrNull(d.OptionV6.String()),
		Interface:    tools.StringOrNull(d.Interface.String()),
		TypeMatchTag: tools.StringOrNull(d.TypeMatchTag.String()),
		Value:        tools.StringOrNull(d.Value),
		Force:        tools.StringOrNull(d.Force),
		Description:  tools.StringOrNull(d.Description),
	}

	var typeSetTagsList []attr.Value
	for _, i := range d.TypeSetTags {
		if i == "" {
			continue
		}
		typeSetTagsList = append(typeSetTagsList, basetypes.NewStringValue(i))
	}
	typeSetTagsTypeList, _ := types.SetValue(types.StringType, typeSetTagsList)
	model.TypeSetTags = typeSetTagsTypeList

	return model, nil
}
