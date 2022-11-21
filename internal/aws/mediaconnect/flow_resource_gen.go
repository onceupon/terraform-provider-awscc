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
	registry.AddResourceFactory("awscc_mediaconnect_flow", flowResource)
}

// flowResource returns the Terraform awscc_mediaconnect_flow resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaConnect::Flow resource.
func flowResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]tfsdk.Attribute{
		"availability_zone": {
			// Property: AvailabilityZone
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.",
			//	  "type": "string"
			//	}
			Description: "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
				resource.RequiresReplace(),
			},
		},
		"flow_arn": {
			// Property: FlowArn
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			//	  "type": "string"
			//	}
			Description: "The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"flow_availability_zone": {
			// Property: FlowAvailabilityZone
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.(ReadOnly)",
			//	  "type": "string"
			//	}
			Description: "The Availability Zone that you want to create the flow in. These options are limited to the Availability Zones within the current AWS.(ReadOnly)",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			//
			//	{
			//	  "description": "The name of the flow.",
			//	  "type": "string"
			//	}
			Description: "The name of the flow.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				resource.RequiresReplace(),
			},
		},
		"source": {
			// Property: Source
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The source of the flow.",
			//	  "properties": {
			//	    "Decryption": {
			//	      "additionalProperties": false,
			//	      "description": "The type of decryption that is used on the content ingested from this source.",
			//	      "properties": {
			//	        "Algorithm": {
			//	          "description": "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
			//	          "enum": [
			//	            "aes128",
			//	            "aes192",
			//	            "aes256"
			//	          ],
			//	          "type": "string"
			//	        },
			//	        "ConstantInitializationVector": {
			//	          "description": "A 128-bit, 16-byte hex value represented by a 32-character string, to be used with the key for encrypting content. This parameter is not valid for static key encryption.",
			//	          "type": "string"
			//	        },
			//	        "DeviceId": {
			//	          "description": "The value of one of the devices that you configured with your digital rights management (DRM) platform key provider. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//	          "type": "string"
			//	        },
			//	        "KeyType": {
			//	          "default": "static-key",
			//	          "description": "The type of key that is used for the encryption. If no keyType is provided, the service will use the default setting (static-key).",
			//	          "enum": [
			//	            "speke",
			//	            "static-key",
			//	            "srt-password"
			//	          ],
			//	          "type": "string"
			//	        },
			//	        "Region": {
			//	          "description": "The AWS Region that the API Gateway proxy endpoint was created in. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//	          "type": "string"
			//	        },
			//	        "ResourceId": {
			//	          "description": "An identifier for the content. The service sends this value to the key server to identify the current endpoint. The resource ID is also known as the content ID. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//	          "type": "string"
			//	        },
			//	        "RoleArn": {
			//	          "description": "The ARN of the role that you created during setup (when you set up AWS Elemental MediaConnect as a trusted entity).",
			//	          "type": "string"
			//	        },
			//	        "SecretArn": {
			//	          "description": " The ARN of the secret that you created in AWS Secrets Manager to store the encryption key. This parameter is required for static key encryption and is not valid for SPEKE encryption.",
			//	          "type": "string"
			//	        },
			//	        "Url": {
			//	          "description": "The URL from the API Gateway proxy that you set up to talk to your key server. This parameter is required for SPEKE encryption and is not valid for static key encryption.",
			//	          "type": "string"
			//	        }
			//	      },
			//	      "required": [
			//	        "RoleArn"
			//	      ],
			//	      "type": "object"
			//	    },
			//	    "Description": {
			//	      "description": "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
			//	      "type": "string"
			//	    },
			//	    "EntitlementArn": {
			//	      "description": "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
			//	      "type": "string"
			//	    },
			//	    "IngestIp": {
			//	      "description": "The IP address that the flow will be listening on for incoming content.",
			//	      "type": "string"
			//	    },
			//	    "IngestPort": {
			//	      "description": "The port that the flow will be listening on for incoming content.",
			//	      "type": "integer"
			//	    },
			//	    "MaxBitrate": {
			//	      "description": "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
			//	      "type": "integer"
			//	    },
			//	    "MaxLatency": {
			//	      "default": 2000,
			//	      "description": "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
			//	      "type": "integer"
			//	    },
			//	    "MinLatency": {
			//	      "default": 2000,
			//	      "description": "The minimum latency in milliseconds.",
			//	      "type": "integer"
			//	    },
			//	    "Name": {
			//	      "description": "The name of the source.",
			//	      "type": "string"
			//	    },
			//	    "Protocol": {
			//	      "description": "The protocol that is used by the source or output.",
			//	      "enum": [
			//	        "zixi-push",
			//	        "rtp-fec",
			//	        "rtp",
			//	        "rist",
			//	        "srt-listener"
			//	      ],
			//	      "type": "string"
			//	    },
			//	    "SourceArn": {
			//	      "description": "The ARN of the source.",
			//	      "type": "string"
			//	    },
			//	    "SourceIngestPort": {
			//	      "description": "The port that the flow will be listening on for incoming content.(ReadOnly)",
			//	      "type": "string"
			//	    },
			//	    "StreamId": {
			//	      "description": "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
			//	      "type": "string"
			//	    },
			//	    "VpcInterfaceName": {
			//	      "description": "The name of the VPC Interface this Source is configured with.",
			//	      "type": "string"
			//	    },
			//	    "WhitelistCidr": {
			//	      "description": "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The source of the flow.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"decryption": {
						// Property: Decryption
						Description: "The type of decryption that is used on the content ingested from this source.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"algorithm": {
									// Property: Algorithm
									Description: "The type of algorithm that is used for the encryption (such as aes128, aes192, or aes256).",
									Type:        types.StringType,
									Optional:    true,
									Computed:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringInSlice([]string{
											"aes128",
											"aes192",
											"aes256",
										}),
									},
									PlanModifiers: []tfsdk.AttributePlanModifier{
										resource.UseStateForUnknown(),
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
											"srt-password",
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
					"description": {
						// Property: Description
						Description: "A description for the source. This value is not used or seen outside of the current AWS Elemental MediaConnect account.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"entitlement_arn": {
						// Property: EntitlementArn
						Description: "The ARN of the entitlement that allows you to subscribe to content that comes from another AWS account. The entitlement is set by the content originator and the ARN is generated as part of the originator's flow.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"ingest_ip": {
						// Property: IngestIp
						Description: "The IP address that the flow will be listening on for incoming content.",
						Type:        types.StringType,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"ingest_port": {
						// Property: IngestPort
						Description: "The port that the flow will be listening on for incoming content.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"max_bitrate": {
						// Property: MaxBitrate
						Description: "The smoothing max bitrate for RIST, RTP, and RTP-FEC streams.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"max_latency": {
						// Property: MaxLatency
						Description: "The maximum latency in milliseconds. This parameter applies only to RIST-based and Zixi-based streams.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.Int64Value(2000)),
							resource.UseStateForUnknown(),
						},
					},
					"min_latency": {
						// Property: MinLatency
						Description: "The minimum latency in milliseconds.",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.Int64Value(2000)),
							resource.UseStateForUnknown(),
						},
					},
					"name": {
						// Property: Name
						Description: "The name of the source.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
							resource.RequiresReplace(),
						},
					},
					"protocol": {
						// Property: Protocol
						Description: "The protocol that is used by the source or output.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"zixi-push",
								"rtp-fec",
								"rtp",
								"rist",
								"srt-listener",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"source_arn": {
						// Property: SourceArn
						Description: "The ARN of the source.",
						Type:        types.StringType,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"source_ingest_port": {
						// Property: SourceIngestPort
						Description: "The port that the flow will be listening on for incoming content.(ReadOnly)",
						Type:        types.StringType,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"stream_id": {
						// Property: StreamId
						Description: "The stream ID that you want to use for this transport. This parameter applies only to Zixi-based streams.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"vpc_interface_name": {
						// Property: VpcInterfaceName
						Description: "The name of the VPC Interface this Source is configured with.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"whitelist_cidr": {
						// Property: WhitelistCidr
						Description: "The range of IP addresses that should be allowed to contribute content to your source. These IP addresses should be in the form of a Classless Inter-Domain Routing (CIDR) block; for example, 10.0.0.0/16.",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
				},
			),
			Required: true,
		},
		"source_failover_config": {
			// Property: SourceFailoverConfig
			// CloudFormation resource type schema:
			//
			//	{
			//	  "additionalProperties": false,
			//	  "description": "The source failover config of the flow.",
			//	  "properties": {
			//	    "RecoveryWindow": {
			//	      "description": "Search window time to look for dash-7 packets",
			//	      "type": "integer"
			//	    },
			//	    "State": {
			//	      "enum": [
			//	        "ENABLED",
			//	        "DISABLED"
			//	      ],
			//	      "type": "string"
			//	    }
			//	  },
			//	  "type": "object"
			//	}
			Description: "The source failover config of the flow.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"recovery_window": {
						// Property: RecoveryWindow
						Description: "Search window time to look for dash-7 packets",
						Type:        types.Int64Type,
						Optional:    true,
						Computed:    true,
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.UseStateForUnknown(),
						},
					},
					"state": {
						// Property: State
						Type:     types.StringType,
						Optional: true,
						Computed: true,
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
				},
			),
			Optional: true,
			Computed: true,
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
		Description: "Resource schema for AWS::MediaConnect::Flow",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaConnect::Flow").WithTerraformTypeName("awscc_mediaconnect_flow")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"algorithm":                      "Algorithm",
		"availability_zone":              "AvailabilityZone",
		"constant_initialization_vector": "ConstantInitializationVector",
		"decryption":                     "Decryption",
		"description":                    "Description",
		"device_id":                      "DeviceId",
		"entitlement_arn":                "EntitlementArn",
		"flow_arn":                       "FlowArn",
		"flow_availability_zone":         "FlowAvailabilityZone",
		"ingest_ip":                      "IngestIp",
		"ingest_port":                    "IngestPort",
		"key_type":                       "KeyType",
		"max_bitrate":                    "MaxBitrate",
		"max_latency":                    "MaxLatency",
		"min_latency":                    "MinLatency",
		"name":                           "Name",
		"protocol":                       "Protocol",
		"recovery_window":                "RecoveryWindow",
		"region":                         "Region",
		"resource_id":                    "ResourceId",
		"role_arn":                       "RoleArn",
		"secret_arn":                     "SecretArn",
		"source":                         "Source",
		"source_arn":                     "SourceArn",
		"source_failover_config":         "SourceFailoverConfig",
		"source_ingest_port":             "SourceIngestPort",
		"state":                          "State",
		"stream_id":                      "StreamId",
		"url":                            "Url",
		"vpc_interface_name":             "VpcInterfaceName",
		"whitelist_cidr":                 "WhitelistCidr",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
