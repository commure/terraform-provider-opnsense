package dnsmasq

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newBootResource,
		newDomainResource,
		newHostResource,
		newOptionResource,
		newRangeResource,
		newTagResource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newBootDataSource,
		newDomainDataSource,
		newHostDataSource,
		newOptionDataSource,
		newRangeDataSource,
		newTagDataSource,
	}
}
