package dnsmasq_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqBootResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsmasqBootResourceConfig("pxelinux.0", "tftp-server", "192.168.1.1", "Test PXE boot"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "filename", "pxelinux.0"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "servername", "tftp-server"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "server_address", "192.168.1.1"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "description", "Test PXE boot"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_boot.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_boot.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsmasqBootResourceConfig("pxelinux.0", "tftp-server-updated", "192.168.1.2", "Updated PXE boot"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "servername", "tftp-server-updated"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "server_address", "192.168.1.2"),
					resource.TestCheckResourceAttr("opnsense_dnsmasq_boot.test", "description", "Updated PXE boot"),
				),
			},
		},
	})
}

func testAccDnsmasqBootResourceConfig(filename, servername, serverAddress, description string) string {
	return fmt.Sprintf(`
resource "opnsense_dnsmasq_boot" "test" {
  filename       = %[1]q
  servername     = %[2]q
  server_address = %[3]q
  description    = %[4]q
}
`, filename, servername, serverAddress, description)
}
