package auth_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAuthUserResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccAuthUserResourceConfig("testuser_tf", "test@example.com", "Test User"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "name", "testuser_tf"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "email", "test@example.com"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "fullname", "Test User"),
					resource.TestCheckResourceAttrSet("opnsense_auth_user.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "opnsense_auth_user.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			// Update and Read testing
			{
				Config: testAccAuthUserResourceConfig("testuser_tf", "updated@example.com", "Updated User"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "email", "updated@example.com"),
					resource.TestCheckResourceAttr("opnsense_auth_user.test", "fullname", "Updated User"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccAuthUserResourceConfig(name, email, fullname string) string {
	return fmt.Sprintf(`
resource "opnsense_auth_user" "test" {
  name     = %[1]q
  email    = %[2]q
  fullname = %[3]q
  password = "TestPass123!"
}
`, name, email, fullname)
}
