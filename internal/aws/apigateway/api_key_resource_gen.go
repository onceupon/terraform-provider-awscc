// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_apigateway_api_key", apiKeyResource)
}

// apiKeyResource returns the Terraform awscc_apigateway_api_key resource.
// This Terraform resource corresponds to the CloudFormation AWS::ApiGateway::ApiKey resource.
func apiKeyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"api_key_id": {
			// Property: APIKeyId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs ",
			//	  "type": "string"
			//	}
			Description: "A Unique Key ID which identifies the API Key. Generated by the Create API and returned by the Read and List APIs ",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"customer_id": {
			// Property: CustomerId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.",
			//	  "type": "string"
			//	}
			Description: "An AWS Marketplace customer identifier to use when integrating with the AWS SaaS Marketplace.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A description of the purpose of the API key.",
			//	  "type": "string"
			//	}
			Description: "A description of the purpose of the API key.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"enabled": {
			// Property: Enabled
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": false,
			//	  "description": "Indicates whether the API key can be used by clients.",
			//	  "type": "boolean"
			//	}
			Description: "Indicates whether the API key can be used by clients.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.BoolValue(false)),
				resource.UseStateForUnknown(),
			},
		},
		"generate_distinct_id": {
			// Property: GenerateDistinctId
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "Specifies whether the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.",
			//	  "type": "boolean"
			//	}
			Description: "Specifies whether the key identifier is distinct from the created API key value. This parameter is deprecated and should not be used.",
			Type:        types.BoolType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
			// GenerateDistinctId is a write-only property.
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A name for the API key. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.",
			//	  "type": "string"
			//	}
			Description: "A name for the API key. If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the API key name.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"stage_keys": {
			// Property: StageKeys
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A list of stages to associate with this API key.",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "RestApiId": {
			//	        "description": "The ID of a RestApi resource that includes the stage with which you want to associate the API key.",
			//	        "type": "string"
			//	      },
			//	      "StageName": {
			//	        "description": "The name of the stage with which to associate the API key. The stage must be included in the RestApi resource that you specified in the RestApiId property. ",
			//	        "type": "string"
			//	      }
			//	    },
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Description: "A list of stages to associate with this API key.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"rest_api_id": {
						// Property: RestApiId
						Description: "The ID of a RestApi resource that includes the stage with which you want to associate the API key.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"stage_name": {
						// Property: StageName
						Description: "The name of the stage with which to associate the API key. The stage must be included in the RestApi resource that you specified in the RestApiId property. ",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "An array of arbitrary tags (key-value pairs) to associate with the API key.",
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
			//	        "maxLength": 256,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Value",
			//	      "Key"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "type": "array",
			//	  "uniqueItems": false
			//	}
			Description: "An array of arbitrary tags (key-value pairs) to associate with the API key.",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenAtMost(256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"value": {
			// Property: Value
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The value of the API key. Must be at least 20 characters long.",
			//	  "type": "string"
			//	}
			Description: "The value of the API key. Must be at least 20 characters long.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			resource.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::ApiGateway::ApiKey",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ApiGateway::ApiKey").WithTerraformTypeName("awscc_apigateway_api_key")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"api_key_id":           "APIKeyId",
		"customer_id":          "CustomerId",
		"description":          "Description",
		"enabled":              "Enabled",
		"generate_distinct_id": "GenerateDistinctId",
		"key":                  "Key",
		"name":                 "Name",
		"rest_api_id":          "RestApiId",
		"stage_keys":           "StageKeys",
		"stage_name":           "StageName",
		"tags":                 "Tags",
		"value":                "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/GenerateDistinctId",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
