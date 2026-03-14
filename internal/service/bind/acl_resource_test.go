package bind_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBindAclResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccBindAclResourceConfig("test_acl", "1"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_bind_acl.test", "name", "test_acl"),
					resource.TestCheckResourceAttr("opnsense_bind_acl.test", "enabled", "1"),
					resource.TestCheckResourceAttrSet("opnsense_bind_acl.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_bind_acl.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccBindAclResourceConfig("test_acl_upd", "0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_bind_acl.test", "name", "test_acl_upd"),
					resource.TestCheckResourceAttr("opnsense_bind_acl.test", "enabled", "0"),
				),
			},
		},
	})
}

func testAccBindAclResourceConfig(name, enabled string) string {
	return fmt.Sprintf(`
resource "opnsense_bind_acl" "test" {
  name    = %[1]q
  enabled = %[2]q
}
`, name, enabled)
}
