---
subcategory: "Dedicated Host (DEH)"
---

# opentelekomcloud_deh_server_v1

Use this data source to get details about the server on a specified Dedicated Host.

## Example Usage

```hcl
variable "deh_id" {}
variable "server_id" {}

data "opentelekomcloud_deh_server_v1" "deh_server" {
  id        = var.deh_id
  server_id = var.server_id
}
```

## Argument Reference

The arguments of this data source act as filters for querying the server on specified dedicated host.

* `dedicated_host_id` - (Optional) The Dedicated Host ID.

* `server_id` - (Optional) The Server ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `user_id` - The ID of the user to which the server belongs.

* `name` - The server name.

* `flavor` - The ID of server specifications.

* `metadata` - The metadata of the server.

* `status` - The status of the server.

* `tenant_id` - The ID of the tenant to which the server belongs.

* `addresses` - The network addresses of the server.
