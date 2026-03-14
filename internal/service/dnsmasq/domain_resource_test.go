package dnsmasq_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqDomainResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsmasqDomainResourceConfig("example.com", "8.8.8.8", "Test domain override"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "domain", "example.com"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "ip", "8.8.8.8"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "description", "Test domain override"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_domain.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_domain.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsmasqDomainResourceConfig("example.com", "8.8.4.4", "Updated domain override"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "ip", "8.8.4.4"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_domain.test", "description", "Updated domain override"),
				),
			},
		},
	})
}

func testAccDnsmasqDomainResourceConfig(domain, ip, description string) string {
	return fmt.Sprintf(`
resource "opnsense_dnsmasq_domain" "test" {
  domain      = %[1]q
  ip          = %[2]q
  description = %[3]q
}
`, domain, ip, description)
}
