package service

import (
	"context"
	"fmt"
	"github.com/browningluke/opnsense-go/pkg/api"
	"github.com/browningluke/opnsense-go/pkg/opnsense"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = &IPsecVTIDataSource{}

func NewIPsecVTIDataSource() datasource.DataSource {
	return &IPsecVTIDataSource{}
}

type IPsecVTIDataSource struct {
	client opnsense.Client
}

func (d *IPsecVTIDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsec_vti"
}

func (d *IPsecVTIDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = IPsecVTIDataSourceSchema()
}

func (d *IPsecVTIDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	apiClient, ok := req.ProviderData.(*api.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *opnsense.Client, got: %T.", req.ProviderData))
		return
	}
	d.client = opnsense.NewClient(apiClient)
}

func (d *IPsecVTIDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *IPsecVTIResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	obj, err := d.client.Ipsec().GetIPsecVTI(ctx, data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read IPsec VTI, got error: %s", err))
		return
	}
	model, err := convertIPsecVTIStructToSchema(obj)
	if err != nil {
		resp.Diagnostics.AddError("Client Error",
			fmt.Sprintf("Unable to read IPsec VTI, got error: %s", err))
		return
	}
	model.Id = data.Id
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
