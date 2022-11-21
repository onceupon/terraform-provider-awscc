// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediaconnect

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
	registry.AddResourceFactory("awscc_mediaconnect_flow_entitlement", flowEntitlementResource)
}

// flowEntitlementResource returns the Terraform awscc_mediaconnect_flow_entitlement resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaConnect::FlowEntitlement resource.
func flowEntitlementResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"data_transfer_subscriber_fee_percent": {
			// Property: DataTransferSubscriberFeePercent
			// CloudFormation resource type schema:
			//
			//	{
			//	  "default": 0,
			//	  "description": "Percentage from 0-100 of the data transfer cost to be billed to the subscriber.",
			//	  "type": "integer"
			//	}
			Description: "Percentage from 0-100 of the data transfer cost to be billed to the subscriber.",
			Type:        types.Int64Type,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.Int64Value(0)),
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "A description of the entitlement.",
			//	  "type": "string"
			//	}
			Description: "A description of the entitlement.",
			Type:        types.StringType,
			Required:    true,
		},
		"encryption": {
			// Property: Encryption
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The type of encryption that will be used on the output that is associated with this entitlement.",
			//	  "properties": {
			//	    "Algorithm": {
			//	      "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
			//	      "enum": [
			//	        "aes128",
			//	        "aes192",
			//	        "aes256"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "ConstantInitializationVector": {
			//	      "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
			//	      "type": "string"
			//	    },
			//	    "DeviceId": {
			//	      "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//	      "type": "string"
			//	    },
			//	    "KeyType": {
			//	      "default": "static-key",
			//	      "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
			//	      "enum": [
			//	        "speke",
			//	        "static-key"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "Region": {
			//	      "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//	      "type": "string"
			//	    },
			//	    "ResourceId": {
			//	      "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//	      "type": "string"
			//	    },
			//	    "RoleArn": {
			//	      "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
			//	      "type": "string"
			//	    },
			//	    "SecretArn": {
			//	      "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
			//	      "type": "string"
			//	    },
			//	    "Url": {
			//	      "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "required": [
			//	    "Algorithm",
			//	    "RoleArn"
			//	  ],
			//	  "type": "object"
			//	}
			Description: "The type of encryption that will be used on the output that is associated with this entitlement.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"algorithm": {
						// Property: Algorithm
						Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"aes128",
								"aes192",
								"aes256",
							}),
						},
					},
					"constant_initialization_vector": {
						// Property: ConstantInitializationVector
						Description: "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"device_id": {
						// Property: DeviceId
						Description: "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"key_type": {
						// Property: KeyType
						Description: "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"speke",
								"static-key",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.StringValue("static-key")),
							resource.UseStateForUnknown(),
						},
					},
					"region": {
						// Property: Region
						Description: "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"resource_id": {
						// Property: ResourceId
						Description: "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"role_arn": {
						// Property: RoleArn
						Description: "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
						Type:        types.StringType,
						Required:    true,
					},
					"secret_arn": {
						// Property: SecretArn
						Description: " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"url": {
						// Property: Url
						Description: "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
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
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"entitlement_arn": {
			// Property: EntitlementArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ARN of the entitlement.",
			//	  "type": "string"
			//	}
			Description: "The ARN of the entitlement.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"entitlement_status": {
			// Property: EntitlementStatus
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": " An indication of whether the entitlement is enabled.",
			//	  "enum": [
			//	    "ENABLED",
			//	    "DISABLED"
			//	  ],
			//	  "type": "string"
			//	}
			Description: " An indication of whether the entitlement is enabled.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"ENABLED",
					"DISABLED",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The ARN of the flow.",
			//	  "type": "string"
			//	}
			Description: "The ARN of the flow.",
			Type:        types.StringType,
			Required:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the entitlement.",
			//	  "type": "string"
			//	}
			Description: "The name of the entitlement.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"subscribers": {
			// Property: Subscribers
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.",
			//	  "items": {
			//	    "type": "string"
			//	  },
			//	  "type": "array"
			//	}
			Description: "The AWS account IDs that you want to share your content with. The receiving accounts (subscribers) will be allowed to create their own flow using your content as the source.",
			Type:        types.ListType{ElemType: types.StringType},
			Required:    true,
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
		Description: "Resource schema for AWS::MediaConnect::FlowEntitlement",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::FlowEntitlement").WithTerraformTypeName("awscc_mediaconnect_flow_entitlement")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm":                            "Algorithm",
		"constant_initialization_vector":       "ConstantInitializationVector",
		"data_transfer_subscriber_fee_percent": "DataTransferSubscriberFeePercent",
		"description":                          "Description",
		"device_id":                            "DeviceId",
		"encryption":                           "Encryption",
		"entitlement_arn":                      "EntitlementArn",
		"entitlement_status":                   "EntitlementStatus",
		"flow_arn":                             "FlowArn",
		"key_type":                             "KeyType",
		"name":                                 "Name",
		"region":                               "Region",
		"resource_id":                          "ResourceId",
		"role_arn":                             "RoleArn",
		"secret_arn":                           "SecretArn",
		"subscribers":                          "Subscribers",
		"url":                                  "Url",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
