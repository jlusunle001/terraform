package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/acceptance/common"
)

func TestAccDcsMaintainWindowV1DataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { common.TestAccPreCheck(t) },
		Providers: common.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDcsMaintainWindowV1DataSource_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcsMaintainWindowV1DataSourceID("data.opentelekomcloud_dcs_maintainwindow_v1.maintainwindow1"),
					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_dcs_maintainwindow_v1.maintainwindow1", "seq", "1"),
					resource.TestCheckResourceAttr(
						"data.opentelekomcloud_dcs_maintainwindow_v1.maintainwindow1", "begin", "22"),
				),
			},
		},
	})
}

func testAccCheckDcsMaintainWindowV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find Dcs maintainwindow data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Dcs maintainwindow data source ID not set")
		}

		return nil
	}
}

var testAccDcsMaintainWindowV1DataSource_basic = fmt.Sprintf(`
data "opentelekomcloud_dcs_maintainwindow_v1" "maintainwindow1" {
  seq = 1
}
`)