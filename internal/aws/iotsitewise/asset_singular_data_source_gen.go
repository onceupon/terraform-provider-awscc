// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iotsitewise

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iotsitewise_asset", assetDataSourceType)
}

// assetDataSourceType returns the Terraform awscc_iotsitewise_asset data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoTSiteWise::Asset resource type.
func assetDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"asset_arn": {
			// Property: AssetArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN of the asset",
			//   "type": "string"
			// }
			Description: "The ARN of the asset",
			Type:        types.StringType,
			Computed:    true,
		},
		"asset_description": {
			// Property: AssetDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description for the asset",
			//   "type": "string"
			// }
			Description: "A description for the asset",
			Type:        types.StringType,
			Computed:    true,
		},
		"asset_hierarchies": {
			// Property: AssetHierarchies
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A hierarchy specifies allowed parent/child asset relationships.",
			//     "properties": {
			//       "ChildAssetId": {
			//         "description": "The ID of the child asset to be associated.",
			//         "type": "string"
			//       },
			//       "LogicalId": {
			//         "description": "The LogicalID of a hierarchy in the parent asset's model.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "LogicalId",
			//       "ChildAssetId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"child_asset_id": {
						// Property: ChildAssetId
						Description: "The ID of the child asset to be associated.",
						Type:        types.StringType,
						Computed:    true,
					},
					"logical_id": {
						// Property: LogicalId
						Description: "The LogicalID of a hierarchy in the parent asset's model.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"asset_id": {
			// Property: AssetId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the asset",
			//   "type": "string"
			// }
			Description: "The ID of the asset",
			Type:        types.StringType,
			Computed:    true,
		},
		"asset_model_id": {
			// Property: AssetModelId
			// CloudFormation resource type schema:
			// {
			//   "description": "The ID of the asset model from which to create the asset.",
			//   "type": "string"
			// }
			Description: "The ID of the asset model from which to create the asset.",
			Type:        types.StringType,
			Computed:    true,
		},
		"asset_name": {
			// Property: AssetName
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique, friendly name for the asset.",
			//   "type": "string"
			// }
			Description: "A unique, friendly name for the asset.",
			Type:        types.StringType,
			Computed:    true,
		},
		"asset_properties": {
			// Property: AssetProperties
			// CloudFormation resource type schema:
			// {
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The asset property's definition, alias, and notification state.",
			//     "properties": {
			//       "Alias": {
			//         "description": "The property alias that identifies the property.",
			//         "type": "string"
			//       },
			//       "LogicalId": {
			//         "description": "Customer provided ID for property.",
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "NotificationState": {
			//         "description": "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
			//         "enum": [
			//           "ENABLED",
			//           "DISABLED"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "LogicalId"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"alias": {
						// Property: Alias
						Description: "The property alias that identifies the property.",
						Type:        types.StringType,
						Computed:    true,
					},
					"logical_id": {
						// Property: LogicalId
						Description: "Customer provided ID for property.",
						Type:        types.StringType,
						Computed:    true,
					},
					"notification_state": {
						// Property: NotificationState
						Description: "The MQTT notification state (ENABLED or DISABLED) for this asset property.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of key-value pairs that contain metadata for the asset.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "type": "string"
			//       },
			//       "Value": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of key-value pairs that contain metadata for the asset.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoTSiteWise::Asset",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTSiteWise::Asset").WithTerraformTypeName("awscc_iotsitewise_asset")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alias":              "Alias",
		"asset_arn":          "AssetArn",
		"asset_description":  "AssetDescription",
		"asset_hierarchies":  "AssetHierarchies",
		"asset_id":           "AssetId",
		"asset_model_id":     "AssetModelId",
		"asset_name":         "AssetName",
		"asset_properties":   "AssetProperties",
		"child_asset_id":     "ChildAssetId",
		"key":                "Key",
		"logical_id":         "LogicalId",
		"notification_state": "NotificationState",
		"tags":               "Tags",
		"value":              "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
