// Code generated by generators/resource/main.go; DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceTypeFactory("awscc_lightsail_load_balancer_tls_certificate", loadBalancerTlsCertificateResourceType)
}

// loadBalancerTlsCertificateResourceType returns the Terraform awscc_lightsail_load_balancer_tls_certificate resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Lightsail::LoadBalancerTlsCertificate resource type.
func loadBalancerTlsCertificateResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"certificate_alternative_names": {
			// Property: CertificateAlternativeNames
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.",
			//   "insertionOrder": false,
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of strings listing alternative domains and subdomains for your SSL/TLS certificate.",
			Type:        types.SetType{ElemType: types.StringType},
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"certificate_domain_name": {
			// Property: CertificateDomainName
			// CloudFormation resource type schema:
			// {
			//   "description": "The domain name (e.g., example.com ) for your SSL/TLS certificate.",
			//   "type": "string"
			// }
			Description: "The domain name (e.g., example.com ) for your SSL/TLS certificate.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"certificate_name": {
			// Property: CertificateName
			// CloudFormation resource type schema:
			// {
			//   "description": "The SSL/TLS certificate name.",
			//   "type": "string"
			// }
			Description: "The SSL/TLS certificate name.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"is_attached": {
			// Property: IsAttached
			// CloudFormation resource type schema:
			// {
			//   "description": "When true, the SSL/TLS certificate is attached to the Lightsail load balancer.",
			//   "type": "boolean"
			// }
			Description: "When true, the SSL/TLS certificate is attached to the Lightsail load balancer.",
			Type:        types.BoolType,
			Optional:    true,
		},
		"load_balancer_name": {
			// Property: LoadBalancerName
			// CloudFormation resource type schema:
			// {
			//   "description": "The name of your load balancer.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The name of your load balancer.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"load_balancer_tls_certificate_arn": {
			// Property: LoadBalancerTlsCertificateArn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"status": {
			// Property: Status
			// CloudFormation resource type schema:
			// {
			//   "description": "The validation status of the SSL/TLS certificate.",
			//   "type": "string"
			// }
			Description: "The validation status of the SSL/TLS certificate.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource Type definition for AWS::Lightsail::LoadBalancerTlsCertificate",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lightsail::LoadBalancerTlsCertificate").WithTerraformTypeName("awscc_lightsail_load_balancer_tls_certificate")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_alternative_names":     "CertificateAlternativeNames",
		"certificate_domain_name":           "CertificateDomainName",
		"certificate_name":                  "CertificateName",
		"is_attached":                       "IsAttached",
		"load_balancer_name":                "LoadBalancerName",
		"load_balancer_tls_certificate_arn": "LoadBalancerTlsCertificateArn",
		"status":                            "Status",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
