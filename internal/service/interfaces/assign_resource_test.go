package interfaces_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccInterfacesAssignResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccInterfacesAssignResourceConfig("192.168.1.1", "24", "Test interface assign"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_assign.test", "ip", "192.168.1.1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_assign.test", "subnet", "24"),
					resource.TestCheckResourceAttr("opnsense_interfaces_assign.test", "description", "Test interface assign"),
					resource.TestCheckResourceAttrSet("opnsense_interfaces_assign.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_interfaces_assign.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccInterfacesAssignResourceConfig("192.168.2.1", "24", "Updated interface assign"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_interfaces_assign.test", "ip", "192.168.2.1"),
					resource.TestCheckResourceAttr("opnsense_interfaces_assign.test", "description", "Updated interface assign"),
				),
			},
		},
	})
}

func testAccInterfacesAssignResourceConfig(ip, subnet, description string) string {
	return fmt.Sprintf(`
resource "opnsense_interfaces_assign" "test" {
  ip          = %[1]q
  subnet      = %[2]q
  description = %[3]q
}
`, ip, subnet, description)
}
