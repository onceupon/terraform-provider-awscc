// Code generated by generators/resource/main.go; DO NOT EDIT.

package iotfleetwise

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceFactory("awscc_iotfleetwise_model_manifest", modelManifestResource)
}

// modelManifestResource returns the Terraform awscc_iotfleetwise_model_manifest resource.
// This Terraform resource corresponds to the CloudFormation AWS::IoTFleetWise::ModelManifest resource.
func modelManifestResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"creation_time": {
			// Property: CreationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 2048,
			//	  "minLength": 1,
			//	  "pattern": "",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 2048),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"last_modification_time": {
			// Property: LastModificationTime
			// CloudFormation resource type schema:
			//
			//	{
			//	  "format": "date-time",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "maxLength": 100,
			//	  "minLength": 1,
			//	  "pattern": "^[a-zA-Z\\d\\-_:]+$",
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 100),
				validate.StringMatch(regexp.MustCompile("^[a-zA-Z\\d\\-_:]+$"), ""),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"nodes": {
			// Property: Nodes
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "minItems": 1,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Type:     types.SetType{ElemType: types.StringType},
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"signal_catalog_arn": {
			// Property: SignalCatalogArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Required: true,
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": "DRAFT",
			//	  "enum": [
			//	    "ACTIVE",
			//	    "DRAFT"
			//	  ],
			//	  "type": "string"
			//	}
			Type:     types.StringType,
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ACTIVE",
					"DRAFT",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.StringValue("DRAFT")),
				resource.UseStateForUnknown(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			//
			//	{
			//	  "insertionOrder": false,
			//	  "items": {
			//	    "additionalProperties": false,
			//	    "properties": {
			//	      "Key": {
			//	        "maxLength": 128,
			//	        "minLength": 1,
			//	        "type": "string"
			//	      },
			//	      "Value": {
			//	        "maxLength": 256,
			//	        "minLength": 0,
			//	        "type": "string"
			//	      }
			//	    },
			//	    "required": [
			//	      "Key",
			//	      "Value"
			//	    ],
			//	    "type": "object"
			//	  },
			//	  "maxItems": 50,
			//	  "minItems": 0,
			//	  "type": "array",
			//	  "uniqueItems": true
			//	}
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
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
		Description: "Definition of AWS::IoTFleetWise::ModelManifest Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoTFleetWise::ModelManifest").WithTerraformTypeName("awscc_iotfleetwise_model_manifest")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"creation_time":          "CreationTime",
		"description":            "Description",
		"key":                    "Key",
		"last_modification_time": "LastModificationTime",
		"name":                   "Name",
		"nodes":                  "Nodes",
		"signal_catalog_arn":     "SignalCatalogArn",
		"status":                 "Status",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
