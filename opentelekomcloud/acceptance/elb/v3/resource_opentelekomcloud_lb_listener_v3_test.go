package acceptance

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/opentelekomcloud/gophertelekomcloud/openstack/elb/v3/listeners"
	th "github.com/opentelekomcloud/gophertelekomcloud/testhelper"
	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/acceptance/common"
	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/acceptance/common/quotas"
	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/acceptance/env"
	"github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/common/cfg"
	elbv3 "github.com/opentelekomcloud/terraform-provider-opentelekomcloud/opentelekomcloud/services/elb/v3"
)

const resourceListenerName = "opentelekomcloud_lb_listener_v3.listener_1"

func TestAccLBV3Listener_basic(t *testing.T) {
	var listener listeners.Listener

	t.Parallel()
	th.AssertNoErr(t, quotas.LbCertificate.Acquire())
	th.AssertNoErr(t, quotas.LoadBalancer.Acquire())
	th.AssertNoErr(t, quotas.LbListener.Acquire())
	defer func() {
		quotas.LbListener.Release()
		quotas.LoadBalancer.Release()
		quotas.LbCertificate.Release()
	}()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { common.TestAccPreCheck(t) },
		ProviderFactories: common.TestAccProviderFactories,
		CheckDestroy:      testAccCheckLBV3ListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLBV3ListenerConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLBV3ListenerExists(resourceListenerName, &listener),
					resource.TestCheckResourceAttr(resourceListenerName, "name", "listener_1"),
					resource.TestCheckResourceAttr(resourceListenerName, "description", "some interesting description"),
				),
			},
			{
				Config: testAccLBV3ListenerConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceListenerName, "name", "listener_1_updated"),
					resource.TestCheckResourceAttr(resourceListenerName, "description", ""),
				),
			},
		},
	})
}

func TestAccLBV3Listener_import(t *testing.T) {
	t.Parallel()
	th.AssertNoErr(t, quotas.LbCertificate.Acquire())
	th.AssertNoErr(t, quotas.LoadBalancer.Acquire())
	th.AssertNoErr(t, quotas.LbListener.Acquire())
	defer func() {
		quotas.LbListener.Release()
		quotas.LoadBalancer.Release()
		quotas.LbCertificate.Release()
	}()

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { common.TestAccPreCheck(t) },
		ProviderFactories: common.TestAccProviderFactories,
		CheckDestroy:      testAccCheckLBV3ListenerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccLBV3ListenerConfigBasic,
			},
			{
				ResourceName:      resourceListenerName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckLBV3ListenerDestroy(s *terraform.State) error {
	config := common.TestAccProvider.Meta().(*cfg.Config)
	client, err := config.ElbV3Client(env.OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf(elbv3.ErrCreateClient, err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "opentelekomcloud_lb_listener_v3" {
			continue
		}

		_, err := listeners.Get(client, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("listener still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckLBV3ListenerExists(n string, listener *listeners.Listener) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := common.TestAccProvider.Meta().(*cfg.Config)
		client, err := config.ElbV3Client(env.OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf(elbv3.ErrCreateClient, err)
		}

		found, err := listeners.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("listener not found")
		}

		*listener = *found

		return nil
	}
}

var testAccLBV3ListenerConfigBasic = fmt.Sprintf(`
%s

resource "opentelekomcloud_lb_loadbalancer_v3" "loadbalancer_1" {
  name        = "loadbalancer_1"
  router_id   = data.opentelekomcloud_vpc_subnet_v1.shared_subnet.vpc_id
  network_ids = [data.opentelekomcloud_vpc_subnet_v1.shared_subnet.network_id]

  availability_zones = ["%s"]
}

resource "opentelekomcloud_lb_certificate_v3" "certificate_1" {
  name        = "certificate_1"
  type        = "server"
  private_key = %s
  certificate = %s
}

resource "opentelekomcloud_lb_listener_v3" "listener_1" {
  name                      = "listener_1"
  description               = "some interesting description"
  loadbalancer_id           = opentelekomcloud_lb_loadbalancer_v3.loadbalancer_1.id
  protocol                  = "HTTPS"
  protocol_port             = 443
  default_tls_container_ref = opentelekomcloud_lb_certificate_v3.certificate_1.id

  insert_headers {
    forwarded_host = true
  }
}
`, common.DataSourceSubnet, env.OS_AVAILABILITY_ZONE, privateKey, certificate)

var testAccLBV3ListenerConfigUpdate = fmt.Sprintf(`
%s

resource "opentelekomcloud_lb_loadbalancer_v3" "loadbalancer_1" {
  name        = "loadbalancer_1_updated"
  router_id   = data.opentelekomcloud_vpc_subnet_v1.shared_subnet.vpc_id
  network_ids = [data.opentelekomcloud_vpc_subnet_v1.shared_subnet.network_id]

  availability_zones = ["%s"]
}

resource "opentelekomcloud_lb_certificate_v3" "certificate_1" {
  name        = "certificate_1"
  type        = "server"
  private_key = %s
  certificate = %s
}

resource "opentelekomcloud_lb_listener_v3" "listener_1" {
  name                      = "listener_1_updated"
  loadbalancer_id           = opentelekomcloud_lb_loadbalancer_v3.loadbalancer_1.id
  protocol                  = "HTTPS"
  protocol_port             = 443
  default_tls_container_ref = opentelekomcloud_lb_certificate_v3.certificate_1.id
}
`, common.DataSourceSubnet, env.OS_AVAILABILITY_ZONE, privateKey, certificate)