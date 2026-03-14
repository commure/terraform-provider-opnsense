package bind_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBindPrimaryDomainResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBindPrimaryDomainResourceConfig("test.example.com", "1", "admin@test.example.com", "ns1.test.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_bind_primary_domain.test", "domain_name", "test.example.com"),
					resource.TestCheckResourceAttr("opnsense_bind_primary_domain.test", "enabled", "1"),
					resource.TestCheckResourceAttr("opnsense_bind_primary_domain.test", "mail_admin", "admin@test.example.com"),
					resource.TestCheckResourceAttr("opnsense_bind_primary_domain.test", "dns_server", "ns1.test.example.com"),
					resource.TestCheckResourceAttrSet("opnsense_bind_primary_domain.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_bind_primary_domain.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBindPrimaryDomainResourceConfig("test.example.com", "1", "postmaster@test.example.com", "ns2.test.example.com"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_bind_primary_domain.test", "mail_admin", "postmaster@test.example.com"),
					resource.TestCheckResourceAttr("opnsense_bind_primary_domain.test", "dns_server", "ns2.test.example.com"),
				),
			},
		},
	})
}

func testAccBindPrimaryDomainResourceConfig(domainName, enabled, mailAdmin, dnsServer string) string {
	return fmt.Sprintf(`
resource "opnsense_bind_primary_domain" "test" {
  domain_name = %[1]q
  enabled     = %[2]q
  mail_admin  = %[3]q
  dns_server  = %[4]q
}
`, domainName, enabled, mailAdmin, dnsServer)
}
