package auth_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAuthGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccAuthGroupResourceConfig("testgroup_tf", "Test group"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "name", "testgroup_tf"),
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "description", "Test group"),
					resource.TestCheckResourceAttrSet("opnsense_auth_group.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "opnsense_auth_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccAuthGroupResourceConfig("testgroup_tf", "Updated group description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_group.test", "description", "Updated group description"),
				),
			},
		},
	})
}

func testAccAuthGroupResourceConfig(name, description string) string {
	return fmt.Sprintf(`
resource "opnsense_auth_group" "test" {
  name        = %[1]q
  description = %[2]q
}
`, name, description)
}
