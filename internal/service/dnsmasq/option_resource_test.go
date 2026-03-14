package dnsmasq_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqOptionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsmasqOptionResourceConfig("192.168.1.1,192.168.1.2", "Test DNS option"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "value", "192.168.1.1,192.168.1.2"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "description", "Test DNS option"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_option.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_option.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsmasqOptionResourceConfig("10.0.0.1", "Updated DNS option"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "value", "10.0.0.1"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_option.test", "description", "Updated DNS option"),
				),
			},
		},
	})
}

func testAccDnsmasqOptionResourceConfig(value, description string) string {
	return fmt.Sprintf(`
resource "opnsense_dnsmasq_option" "test" {
  value       = %[1]q
  description = %[2]q
}
`, value, description)
}
