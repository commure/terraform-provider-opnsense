package dnsmasq_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqRangeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsmasqRangeResourceConfig("192.168.1.100", "192.168.1.200", "255.255.255.0", "Test DHCP range"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "start_address", "192.168.1.100"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "end_address", "192.168.1.200"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "subnet_mask", "255.255.255.0"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "description", "Test DHCP range"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_range.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_range.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsmasqRangeResourceConfig("192.168.1.50", "192.168.1.150", "255.255.255.0", "Updated DHCP range"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "start_address", "192.168.1.50"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "end_address", "192.168.1.150"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_range.test", "description", "Updated DHCP range"),
				),
			},
		},
	})
}

func testAccDnsmasqRangeResourceConfig(startAddr, endAddr, subnetMask, description string) string {
	return fmt.Sprintf(`
resource "opnsense_dnsmasq_range" "test" {
  start_address = %[1]q
  end_address   = %[2]q
  subnet_mask   = %[3]q
  description   = %[4]q
}
`, startAddr, endAddr, subnetMask, description)
}
