package service

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-opnsense/internal/tools"
)

type BindRecordResourceModel struct {
	Domain  types.String `tfsdk:"domain"`
	Enabled types.Bool   `tfsdk:"enabled"`
	Name    types.String `tfsdk:"name"`
	Type    types.String `tfsdk:"type"`
	Value   types.String `tfsdk:"value"`
	Id      types.String `tfsdk:"id"`
}

func BindRecordResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages a BIND DNS record in OPNsense.",
		Attributes: map[string]schema.Attribute{
			"domain": schema.StringAttribute{
				MarkdownDescription: "The domain this record belongs to.",
				Required:            true,
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this record. Defaults to `true`.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Record name.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "DNS record type (A, AAAA, CNAME, MX, etc.).",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "Record value.",
				Required:            true,
			},
			"id": schema.StringAttribute{
				Computed: true, MarkdownDescription: "UUID of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
		},
	}
}

func BindRecordDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Gets information about a BIND DNS record.",
		Attributes: map[string]dschema.Attribute{
			"id":      dschema.StringAttribute{Required: true, MarkdownDescription: "UUID of the resource."},
			"domain":  dschema.StringAttribute{Computed: true, MarkdownDescription: "The domain."},
			"enabled": dschema.BoolAttribute{Computed: true, MarkdownDescription: "Whether enabled."},
			"name":    dschema.StringAttribute{Computed: true, MarkdownDescription: "Record name."},
			"type":    dschema.StringAttribute{Computed: true, MarkdownDescription: "Record type."},
			"value":   dschema.StringAttribute{Computed: true, MarkdownDescription: "Record value."},
		},
	}
}

func convertBindRecordSchemaToStruct(d *BindRecordResourceModel) (*bind.Record, error) {
	return &bind.Record{
		Domain:  api.SelectedMap(d.Domain.ValueString()),
		Enabled: tools.BoolToString(d.Enabled.ValueBool()),
		Name:    d.Name.ValueString(),
		Type:    api.SelectedMap(d.Type.ValueString()),
		Value:   d.Value.ValueString(),
	}, nil
}

func convertBindRecordStructToSchema(d *bind.Record) (*BindRecordResourceModel, error) {
	return &BindRecordResourceModel{
		Domain:  types.StringValue(d.Domain.String()),
		Enabled: types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:    types.StringValue(d.Name),
		Type:    types.StringValue(d.Type.String()),
		Value:   types.StringValue(d.Value),
	}, nil
}
