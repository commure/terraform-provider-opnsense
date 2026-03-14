package service

import (
	"context"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/firewall"
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

type FirewallNatOneToOneResourceModel struct {
	Enabled           types.Bool   `tfsdk:"enabled"`
	Log               types.Bool   `tfsdk:"log"`
	Sequence          types.String `tfsdk:"sequence"`
	Interface         types.String `tfsdk:"interface"`
	Type              types.String `tfsdk:"type"`
	SourceNet         types.String `tfsdk:"source_net"`
	SourceInvert      types.Bool   `tfsdk:"source_invert"`
	DestinationNet    types.String `tfsdk:"destination_net"`
	DestinationInvert types.Bool   `tfsdk:"destination_invert"`
	ExternalNet       types.String `tfsdk:"external_net"`
	NatReflection     types.String `tfsdk:"nat_reflection"`
	Categories        types.Set    `tfsdk:"categories"`
	Description       types.String `tfsdk:"description"`
	Id                types.String `tfsdk:"id"`
}

func FirewallNatOneToOneResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a one-to-one NAT rule in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this rule. Defaults to `true`.",
				Optional: true, Computed: true, Default: booldefault.StaticBool(true),
			},
			"log": schema.BoolAttribute{
				MarkdownDescription: "Log packets. Defaults to `false`.",
				Optional: true, Computed: true, Default: booldefault.StaticBool(false),
			},
			"sequence":      schema.StringAttribute{MarkdownDescription: "Sequence order.", Optional: true},
			"interface":     schema.StringAttribute{MarkdownDescription: "Interface.", Required: true},
			"type":          schema.StringAttribute{MarkdownDescription: "NAT type.", Optional: true, Computed: true},
			"source_net":    schema.StringAttribute{MarkdownDescription: "Source network.", Optional: true},
			"source_invert": schema.BoolAttribute{
				MarkdownDescription: "Invert source match. Defaults to `false`.",
				Optional: true, Computed: true, Default: booldefault.StaticBool(false),
			},
			"destination_net":    schema.StringAttribute{MarkdownDescription: "Destination network.", Optional: true},
			"destination_invert": schema.BoolAttribute{
				MarkdownDescription: "Invert destination match. Defaults to `false`.",
				Optional: true, Computed: true, Default: booldefault.StaticBool(false),
			},
			"external_net":  schema.StringAttribute{MarkdownDescription: "External (translation) network.", Optional: true},
			"nat_reflection": schema.StringAttribute{MarkdownDescription: "NAT reflection mode.", Optional: true, Computed: true},
			"categories": schema.SetAttribute{
				MarkdownDescription: "Categories.",
				Optional: true, Computed: true, ElementType: types.StringType,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"description": schema.StringAttribute{MarkdownDescription: "Description.", Optional: true},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func FirewallNatOneToOneDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a one-to-one NAT rule.",
		Attributes: map[string]dschema.Attribute{
			"id":                 dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"enabled":            dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether enabled."},
			"log":                dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether logging."},
			"sequence":           dschema.StringAttribute{Computed: true, MarkdownDescription: "Sequence."},
			"interface":          dschema.StringAttribute{Computed: true, MarkdownDescription: "Interface."},
			"type":               dschema.StringAttribute{Computed: true, MarkdownDescription: "NAT type."},
			"source_net":         dschema.StringAttribute{Computed: true, MarkdownDescription: "Source network."},
			"source_invert":      dschema.BoolAttribute{Computed: true, MarkdownDescription: "Source inverted."},
			"destination_net":    dschema.StringAttribute{Computed: true, MarkdownDescription: "Destination network."},
			"destination_invert": dschema.BoolAttribute{Computed: true, MarkdownDescription: "Destination inverted."},
			"external_net":       dschema.StringAttribute{Computed: true, MarkdownDescription: "External network."},
			"nat_reflection":     dschema.StringAttribute{Computed: true, MarkdownDescription: "NAT reflection."},
			"categories":         dschema.SetAttribute{Computed: true, ElementType: types.StringType, MarkdownDescription: "Categories."},
			"description":        dschema.StringAttribute{Computed: true, MarkdownDescription: "Description."},
		},
	}
}

func convertFirewallNatOneToOneSchemaToStruct(d *FirewallNatOneToOneResourceModel) (*firewall.NatOneToOne, error) {
	var catList []string
	d.Categories.ElementsAs(context.Background(), &catList, false)
	return &firewall.NatOneToOne{
		Enabled:           tools.BoolToString(d.Enabled.ValueBool()),
		Log:               tools.BoolToString(d.Log.ValueBool()),
		Sequence:          d.Sequence.ValueString(),
		Interface:         api.SelectedMap(d.Interface.ValueString()),
		Type:              api.SelectedMap(d.Type.ValueString()),
		SourceNet:         d.SourceNet.ValueString(),
		SourceInvert:      tools.BoolToString(d.SourceInvert.ValueBool()),
		DestinationNet:    d.DestinationNet.ValueString(),
		DestinationInvert: tools.BoolToString(d.DestinationInvert.ValueBool()),
		ExternalNet:       d.ExternalNet.ValueString(),
		NatReflection:     api.SelectedMap(d.NatReflection.ValueString()),
		Categories:        catList,
		Description:       d.Description.ValueString(),
	}, nil
}

func convertFirewallNatOneToOneStructToSchema(d *firewall.NatOneToOne) (*FirewallNatOneToOneResourceModel, error) {
	model := &FirewallNatOneToOneResourceModel{
		Enabled:           types.BoolValue(tools.StringToBool(d.Enabled)),
		Log:               types.BoolValue(tools.StringToBool(d.Log)),
		Sequence:          tools.StringOrNull(d.Sequence),
		Interface:         types.StringValue(d.Interface.String()),
		Type:              types.StringValue(d.Type.String()),
		SourceNet:         tools.StringOrNull(d.SourceNet),
		SourceInvert:      types.BoolValue(tools.StringToBool(d.SourceInvert)),
		DestinationNet:    tools.StringOrNull(d.DestinationNet),
		DestinationInvert: types.BoolValue(tools.StringToBool(d.DestinationInvert)),
		ExternalNet:       tools.StringOrNull(d.ExternalNet),
		NatReflection:     types.StringValue(d.NatReflection.String()),
		Description:       tools.StringOrNull(d.Description),
	}
	var catList []attr.Value
	for _, i := range d.Categories {
		if i == "" { continue }
		catList = append(catList, basetypes.NewStringValue(i))
	}
	catSet, _ := types.SetValue(types.StringType, catList)
	model.Categories = catSet
	return model, nil
}
