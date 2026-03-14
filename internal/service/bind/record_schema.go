package bind

import (
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/bind"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type recordResourceModel struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Name    types.String `tfsdk:"name"`
	Type    types.String `tfsdk:"type"`
	Value   types.String `tfsdk:"value"`
	Domain  types.String `tfsdk:"domain"`

	Id types.String `tfsdk:"id"`
}

func recordResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manages an OPNsense BIND record.",

		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable this record. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The record name.",
				Required:            true,
			},
			"type": schema.StringAttribute{
				MarkdownDescription: "The record type (e.g. A, AAAA, CNAME, MX, etc.).",
				Required:            true,
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "The record value.",
				Required:            true,
			},
			"domain": schema.StringAttribute{
				MarkdownDescription: "The domain this record belongs to.",
				Required:            true,
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

func recordDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Reads an OPNsense BIND record.",

		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the resource.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether this record is enabled.",
				Computed:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "The record name.",
				Computed:            true,
			},
			"type": dschema.StringAttribute{
				MarkdownDescription: "The record type.",
				Computed:            true,
			},
			"value": dschema.StringAttribute{
				MarkdownDescription: "The record value.",
				Computed:            true,
			},
			"domain": dschema.StringAttribute{
				MarkdownDescription: "The domain this record belongs to.",
				Computed:            true,
			},
		},
	}
}

func convertRecordSchemaToStruct(d *recordResourceModel) (*bind.Record, error) {
	return &bind.Record{
		Enabled: tools.BoolToString(d.Enabled.ValueBool()),
		Name:    d.Name.ValueString(),
		Type:    api.SelectedMap(d.Type.ValueString()),
		Value:   d.Value.ValueString(),
		Domain:  api.SelectedMap(d.Domain.ValueString()),
	}, nil
}

func convertRecordStructToSchema(d *bind.Record) (*recordResourceModel, error) {
	return &recordResourceModel{
		Enabled: types.BoolValue(tools.StringToBool(d.Enabled)),
		Name:    types.StringValue(d.Name),
		Type:    types.StringValue(d.Type.String()),
		Value:   types.StringValue(d.Value),
		Domain:  types.StringValue(d.Domain.String()),
	}, nil
}
