---
subcategory: "Virtual Private Cloud (VPC)"
---

# opentelekomcloud_vpc_route_ids_v2

Use this data source to get a list of route ids for a vpc_id.

This resource can be useful for getting back a list of route ids for a vpc.

## Example Usage

```hcl
variable "vpc_id" {}

data "opentelekomcloud_vpc_route_ids_v2" "example" {
  vpc_id = var.vpc_id
}

data "opentelekomcloud_vpc_route_v2" "vpc_route" {
  count = length(data.opentelekomcloud_vpc_route_ids_v2.example.ids)
  id    = data.opentelekomcloud_vpc_route_ids_v2.example.ids[count.index]
}

output "route_nexthop" {
  value = [data.opentelekomcloud_vpc_route_v2.vpc_route.*.nexthop]
}
```

## Argument Reference

* `vpc_id` - (Required) The VPC ID that you want to filter from.

## Attributes Reference

* `ids` - A list of all the route ids found. This data source will fail if none are found.
