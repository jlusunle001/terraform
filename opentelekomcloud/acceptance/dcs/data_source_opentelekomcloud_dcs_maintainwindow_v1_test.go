package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/acceptance/common"
)

const dataMaintainWindowName = "data.opentelekomcloud_dcs_maintainwindow_v1.maintainwindow1"

func TestAccDcsMaintainWindowV1DataSource_basic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { common.TestAccPreCheck(t) },
		ProviderFactories: common.TestAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDcsMaintainWindowV1DataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDcsMaintainWindowV1DataSourceID(dataMaintainWindowName),
					resource.TestCheckResourceAttr(dataMaintainWindowName, "seq", "1"),
					resource.TestCheckResourceAttr(dataMaintainWindowName, "begin", "22"),
				),
			},
		},
	})
}

func testAccCheckDcsMaintainWindowV1DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("can't find DCS maintainwindow data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("dcs maintainwindow data source ID not set")
		}

		return nil
	}
}

const testAccDcsMaintainWindowV1DataSourceBasic = `
data "opentelekomcloud_dcs_maintainwindow_v1" "maintainwindow1" {
  seq = 1
}
`
