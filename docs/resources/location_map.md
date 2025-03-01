---
page_title: "awscc_location_map Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Definition of AWS::Location::Map Resource Type
---

# awscc_location_map (Resource)

Definition of AWS::Location::Map Resource Type

## Example Usage

### Basic

```terraform
resource "awscc_location_map" "example" {
  map_name = "example"

  configuration = {
    style = "VectorHereExplore"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `map_name` (String)

### Optional

- `description` (String)
- `pricing_plan` (String)

### Read-Only

- `arn` (String)
- `create_time` (String) The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)
- `data_source` (String)
- `id` (String) Uniquely identifies the resource.
- `map_arn` (String)
- `update_time` (String) The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `style` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_location_map.example <resource ID>
```