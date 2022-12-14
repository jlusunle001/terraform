---
subcategory: "Dedicated Host (DEH)"
---

# opentelekomcloud_deh_host_v1

Use this data source to get details about the allocated dedicated hosts from OpenTelekomCloud.

## Example Usage

```hcl
variable "deh_id" {}

data "opentelekomcloud_deh_host_v1" "deh_host" {
  id = var.deh_id
}
```

## Argument Reference

The arguments of this data source act as filters for querying the allocated dedicated host.

* `id` - (Optional) The Dedicated Host ID.

* `name` - (Optional) The Dedicated Host name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `host_type` - The Dedicated Host type.

* `host_type_name` - The Dedicated Host name of type.

* `status` - The Dedicated Host status.

* `availability_zone` - The Availability Zone to which the Dedicated Host belongs.

* `tenant_id` -  The UUID of the tenant in a multi-tenancy cloud.

* `auto_placement` - Allows a instance to be automatically placed onto the available Dedicated Hosts.

* `available_vcpus` - Thenumber of available vCPUs for the Dedicated Host.

* `available_memory` - The size of available memory for the Dedicated Host.

* `sockets` - The number of host physical sockets.

* `instance_total` - The number of the placed VMs.

* `memory` - The size of host physical memory (MB).

* `vcpus` - The number of host vCPUs.

* `available_instance_capacities` - The VM flavors placed on the Dedicated Host.

* `cores` - The number of hosts physical cores.

* `instance_uuids` - The VMs started on the Dedicated Host.
