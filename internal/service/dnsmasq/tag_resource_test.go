package dnsmasq_test

import (
	"fmt"
	"testing"

	"github.com/browningluke/terraform-provider-opnsense/internal/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDnsmasqTagResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDnsmasqTagResourceConfig("test_tag"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_tag.test", "tag", "test_tag"),
					resource.TestCheckResourceAttrSet("opnsense_dnsmasq_tag.test", "id"),
				),
			},
			{
				ResourceName:      "opnsense_dnsmasq_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccDnsmasqTagResourceConfig("updated_tag"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("opnsense_dnsmasq_tag.test", "tag", "updated_tag"),
				),
			},
		},
	})
}

func testAccDnsmasqTagResourceConfig(tag string) string {
	return fmt.Sprintf(`
resource "opnsense_dnsmasq_tag" "test" {
  tag = %[1]q
}
`, tag)
}
