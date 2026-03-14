package dnsmasq_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqHostResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsmasqHostResourceConfig("testhost", "example.com", "Test host"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_host.test", "hostname", "testhost"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_host.test", "domain", "example.com"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_host.test", "description", "Test host"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_host.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_host.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsmasqHostResourceConfig("testhost", "example.com", "Updated host description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_host.test", "description", "Updated host description"),
				),
			},
		},
	})
}

func testAccDnsmasqHostResourceConfig(hostname, domain, description string) string {
	return fmt.Sprintf(`
resource "opnsense_dnsmasq_host" "test" {
  hostname    = %[1]q
  domain      = %[2]q
  description = %[3]q
}
`, hostname, domain, description)
}
