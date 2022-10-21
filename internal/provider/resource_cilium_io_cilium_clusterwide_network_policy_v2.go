/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type CiliumIoCiliumClusterwideNetworkPolicyV2Resource struct{}

var (
	_ resource.Resource = (*CiliumIoCiliumClusterwideNetworkPolicyV2Resource)(nil)
)

type CiliumIoCiliumClusterwideNetworkPolicyV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
	Specs      types.List   `tfsdk:"specs"`
}

type CiliumIoCiliumClusterwideNetworkPolicyV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		Egress *[]struct {
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" yaml:"family,omitempty"`

					Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`
			} `tfsdk:"icmps" yaml:"icmps,omitempty"`

			ToCIDR *[]string `tfsdk:"to_cidr" yaml:"toCIDR,omitempty"`

			ToCIDRSet *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
			} `tfsdk:"to_cidr_set" yaml:"toCIDRSet,omitempty"`

			ToEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"to_endpoints" yaml:"toEndpoints,omitempty"`

			ToEntities *[]string `tfsdk:"to_entities" yaml:"toEntities,omitempty"`

			ToFQDNs *[]struct {
				MatchName *string `tfsdk:"match_name" yaml:"matchName,omitempty"`

				MatchPattern *string `tfsdk:"match_pattern" yaml:"matchPattern,omitempty"`
			} `tfsdk:"to_fqd_ns" yaml:"toFQDNs,omitempty"`

			ToGroups *[]struct {
				Aws *struct {
					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Region *string `tfsdk:"region" yaml:"region,omitempty"`

					SecurityGroupsIds *[]string `tfsdk:"security_groups_ids" yaml:"securityGroupsIds,omitempty"`

					SecurityGroupsNames *[]string `tfsdk:"security_groups_names" yaml:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" yaml:"aws,omitempty"`
			} `tfsdk:"to_groups" yaml:"toGroups,omitempty"`

			ToPorts *[]struct {
				OriginatingTLS *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					TrustedCA *string `tfsdk:"trusted_ca" yaml:"trustedCA,omitempty"`
				} `tfsdk:"originating_tls" yaml:"originatingTLS,omitempty"`

				Ports *[]struct {
					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				Rules *struct {
					Dns *[]struct {
						MatchName *string `tfsdk:"match_name" yaml:"matchName,omitempty"`

						MatchPattern *string `tfsdk:"match_pattern" yaml:"matchPattern,omitempty"`
					} `tfsdk:"dns" yaml:"dns,omitempty"`

					Http *[]struct {
						HeaderMatches *[]struct {
							Mismatch *string `tfsdk:"mismatch" yaml:"mismatch,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Secret *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"secret" yaml:"secret,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"header_matches" yaml:"headerMatches,omitempty"`

						Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"http" yaml:"http,omitempty"`

					Kafka *[]struct {
						ApiKey *string `tfsdk:"api_key" yaml:"apiKey,omitempty"`

						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`
					} `tfsdk:"kafka" yaml:"kafka,omitempty"`

					L7 *[]map[string]string `tfsdk:"l7" yaml:"l7,omitempty"`

					L7proto *string `tfsdk:"l7proto" yaml:"l7proto,omitempty"`
				} `tfsdk:"rules" yaml:"rules,omitempty"`

				TerminatingTLS *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					TrustedCA *string `tfsdk:"trusted_ca" yaml:"trustedCA,omitempty"`
				} `tfsdk:"terminating_tls" yaml:"terminatingTLS,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`

			ToRequires *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"to_requires" yaml:"toRequires,omitempty"`

			ToServices *[]struct {
				K8sService *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
				} `tfsdk:"k8s_service" yaml:"k8sService,omitempty"`

				K8sServiceSelector *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"k8s_service_selector" yaml:"k8sServiceSelector,omitempty"`
			} `tfsdk:"to_services" yaml:"toServices,omitempty"`
		} `tfsdk:"egress" yaml:"egress,omitempty"`

		EgressDeny *[]struct {
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" yaml:"family,omitempty"`

					Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`
			} `tfsdk:"icmps" yaml:"icmps,omitempty"`

			ToCIDR *[]string `tfsdk:"to_cidr" yaml:"toCIDR,omitempty"`

			ToCIDRSet *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
			} `tfsdk:"to_cidr_set" yaml:"toCIDRSet,omitempty"`

			ToEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"to_endpoints" yaml:"toEndpoints,omitempty"`

			ToEntities *[]string `tfsdk:"to_entities" yaml:"toEntities,omitempty"`

			ToGroups *[]struct {
				Aws *struct {
					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Region *string `tfsdk:"region" yaml:"region,omitempty"`

					SecurityGroupsIds *[]string `tfsdk:"security_groups_ids" yaml:"securityGroupsIds,omitempty"`

					SecurityGroupsNames *[]string `tfsdk:"security_groups_names" yaml:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" yaml:"aws,omitempty"`
			} `tfsdk:"to_groups" yaml:"toGroups,omitempty"`

			ToPorts *[]struct {
				Ports *[]struct {
					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`

			ToRequires *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"to_requires" yaml:"toRequires,omitempty"`

			ToServices *[]struct {
				K8sService *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
				} `tfsdk:"k8s_service" yaml:"k8sService,omitempty"`

				K8sServiceSelector *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"k8s_service_selector" yaml:"k8sServiceSelector,omitempty"`
			} `tfsdk:"to_services" yaml:"toServices,omitempty"`
		} `tfsdk:"egress_deny" yaml:"egressDeny,omitempty"`

		EndpointSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"endpoint_selector" yaml:"endpointSelector,omitempty"`

		Ingress *[]struct {
			FromCIDR *[]string `tfsdk:"from_cidr" yaml:"fromCIDR,omitempty"`

			FromCIDRSet *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
			} `tfsdk:"from_cidr_set" yaml:"fromCIDRSet,omitempty"`

			FromEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"from_endpoints" yaml:"fromEndpoints,omitempty"`

			FromEntities *[]string `tfsdk:"from_entities" yaml:"fromEntities,omitempty"`

			FromRequires *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"from_requires" yaml:"fromRequires,omitempty"`

			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" yaml:"family,omitempty"`

					Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`
			} `tfsdk:"icmps" yaml:"icmps,omitempty"`

			ToPorts *[]struct {
				OriginatingTLS *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					TrustedCA *string `tfsdk:"trusted_ca" yaml:"trustedCA,omitempty"`
				} `tfsdk:"originating_tls" yaml:"originatingTLS,omitempty"`

				Ports *[]struct {
					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				Rules *struct {
					Dns *[]struct {
						MatchName *string `tfsdk:"match_name" yaml:"matchName,omitempty"`

						MatchPattern *string `tfsdk:"match_pattern" yaml:"matchPattern,omitempty"`
					} `tfsdk:"dns" yaml:"dns,omitempty"`

					Http *[]struct {
						HeaderMatches *[]struct {
							Mismatch *string `tfsdk:"mismatch" yaml:"mismatch,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Secret *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"secret" yaml:"secret,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"header_matches" yaml:"headerMatches,omitempty"`

						Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"http" yaml:"http,omitempty"`

					Kafka *[]struct {
						ApiKey *string `tfsdk:"api_key" yaml:"apiKey,omitempty"`

						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`
					} `tfsdk:"kafka" yaml:"kafka,omitempty"`

					L7 *[]map[string]string `tfsdk:"l7" yaml:"l7,omitempty"`

					L7proto *string `tfsdk:"l7proto" yaml:"l7proto,omitempty"`
				} `tfsdk:"rules" yaml:"rules,omitempty"`

				TerminatingTLS *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					TrustedCA *string `tfsdk:"trusted_ca" yaml:"trustedCA,omitempty"`
				} `tfsdk:"terminating_tls" yaml:"terminatingTLS,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`
		} `tfsdk:"ingress" yaml:"ingress,omitempty"`

		IngressDeny *[]struct {
			FromCIDR *[]string `tfsdk:"from_cidr" yaml:"fromCIDR,omitempty"`

			FromCIDRSet *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
			} `tfsdk:"from_cidr_set" yaml:"fromCIDRSet,omitempty"`

			FromEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"from_endpoints" yaml:"fromEndpoints,omitempty"`

			FromEntities *[]string `tfsdk:"from_entities" yaml:"fromEntities,omitempty"`

			FromRequires *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"from_requires" yaml:"fromRequires,omitempty"`

			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" yaml:"family,omitempty"`

					Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`
			} `tfsdk:"icmps" yaml:"icmps,omitempty"`

			ToPorts *[]struct {
				Ports *[]struct {
					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`
		} `tfsdk:"ingress_deny" yaml:"ingressDeny,omitempty"`

		Labels *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Source *string `tfsdk:"source" yaml:"source,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"labels" yaml:"labels,omitempty"`

		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`

	Specs *[]struct {
		Description *string `tfsdk:"description" yaml:"description,omitempty"`

		Egress *[]struct {
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" yaml:"family,omitempty"`

					Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`
			} `tfsdk:"icmps" yaml:"icmps,omitempty"`

			ToCIDR *[]string `tfsdk:"to_cidr" yaml:"toCIDR,omitempty"`

			ToCIDRSet *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
			} `tfsdk:"to_cidr_set" yaml:"toCIDRSet,omitempty"`

			ToEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"to_endpoints" yaml:"toEndpoints,omitempty"`

			ToEntities *[]string `tfsdk:"to_entities" yaml:"toEntities,omitempty"`

			ToFQDNs *[]struct {
				MatchName *string `tfsdk:"match_name" yaml:"matchName,omitempty"`

				MatchPattern *string `tfsdk:"match_pattern" yaml:"matchPattern,omitempty"`
			} `tfsdk:"to_fqd_ns" yaml:"toFQDNs,omitempty"`

			ToGroups *[]struct {
				Aws *struct {
					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Region *string `tfsdk:"region" yaml:"region,omitempty"`

					SecurityGroupsIds *[]string `tfsdk:"security_groups_ids" yaml:"securityGroupsIds,omitempty"`

					SecurityGroupsNames *[]string `tfsdk:"security_groups_names" yaml:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" yaml:"aws,omitempty"`
			} `tfsdk:"to_groups" yaml:"toGroups,omitempty"`

			ToPorts *[]struct {
				OriginatingTLS *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					TrustedCA *string `tfsdk:"trusted_ca" yaml:"trustedCA,omitempty"`
				} `tfsdk:"originating_tls" yaml:"originatingTLS,omitempty"`

				Ports *[]struct {
					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				Rules *struct {
					Dns *[]struct {
						MatchName *string `tfsdk:"match_name" yaml:"matchName,omitempty"`

						MatchPattern *string `tfsdk:"match_pattern" yaml:"matchPattern,omitempty"`
					} `tfsdk:"dns" yaml:"dns,omitempty"`

					Http *[]struct {
						HeaderMatches *[]struct {
							Mismatch *string `tfsdk:"mismatch" yaml:"mismatch,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Secret *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"secret" yaml:"secret,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"header_matches" yaml:"headerMatches,omitempty"`

						Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"http" yaml:"http,omitempty"`

					Kafka *[]struct {
						ApiKey *string `tfsdk:"api_key" yaml:"apiKey,omitempty"`

						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`
					} `tfsdk:"kafka" yaml:"kafka,omitempty"`

					L7 *[]map[string]string `tfsdk:"l7" yaml:"l7,omitempty"`

					L7proto *string `tfsdk:"l7proto" yaml:"l7proto,omitempty"`
				} `tfsdk:"rules" yaml:"rules,omitempty"`

				TerminatingTLS *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					TrustedCA *string `tfsdk:"trusted_ca" yaml:"trustedCA,omitempty"`
				} `tfsdk:"terminating_tls" yaml:"terminatingTLS,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`

			ToRequires *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"to_requires" yaml:"toRequires,omitempty"`

			ToServices *[]struct {
				K8sService *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
				} `tfsdk:"k8s_service" yaml:"k8sService,omitempty"`

				K8sServiceSelector *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"k8s_service_selector" yaml:"k8sServiceSelector,omitempty"`
			} `tfsdk:"to_services" yaml:"toServices,omitempty"`
		} `tfsdk:"egress" yaml:"egress,omitempty"`

		EgressDeny *[]struct {
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" yaml:"family,omitempty"`

					Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`
			} `tfsdk:"icmps" yaml:"icmps,omitempty"`

			ToCIDR *[]string `tfsdk:"to_cidr" yaml:"toCIDR,omitempty"`

			ToCIDRSet *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
			} `tfsdk:"to_cidr_set" yaml:"toCIDRSet,omitempty"`

			ToEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"to_endpoints" yaml:"toEndpoints,omitempty"`

			ToEntities *[]string `tfsdk:"to_entities" yaml:"toEntities,omitempty"`

			ToGroups *[]struct {
				Aws *struct {
					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Region *string `tfsdk:"region" yaml:"region,omitempty"`

					SecurityGroupsIds *[]string `tfsdk:"security_groups_ids" yaml:"securityGroupsIds,omitempty"`

					SecurityGroupsNames *[]string `tfsdk:"security_groups_names" yaml:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" yaml:"aws,omitempty"`
			} `tfsdk:"to_groups" yaml:"toGroups,omitempty"`

			ToPorts *[]struct {
				Ports *[]struct {
					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`

			ToRequires *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"to_requires" yaml:"toRequires,omitempty"`

			ToServices *[]struct {
				K8sService *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
				} `tfsdk:"k8s_service" yaml:"k8sService,omitempty"`

				K8sServiceSelector *struct {
					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

					Selector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"selector" yaml:"selector,omitempty"`
				} `tfsdk:"k8s_service_selector" yaml:"k8sServiceSelector,omitempty"`
			} `tfsdk:"to_services" yaml:"toServices,omitempty"`
		} `tfsdk:"egress_deny" yaml:"egressDeny,omitempty"`

		EndpointSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"endpoint_selector" yaml:"endpointSelector,omitempty"`

		Ingress *[]struct {
			FromCIDR *[]string `tfsdk:"from_cidr" yaml:"fromCIDR,omitempty"`

			FromCIDRSet *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
			} `tfsdk:"from_cidr_set" yaml:"fromCIDRSet,omitempty"`

			FromEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"from_endpoints" yaml:"fromEndpoints,omitempty"`

			FromEntities *[]string `tfsdk:"from_entities" yaml:"fromEntities,omitempty"`

			FromRequires *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"from_requires" yaml:"fromRequires,omitempty"`

			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" yaml:"family,omitempty"`

					Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`
			} `tfsdk:"icmps" yaml:"icmps,omitempty"`

			ToPorts *[]struct {
				OriginatingTLS *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					TrustedCA *string `tfsdk:"trusted_ca" yaml:"trustedCA,omitempty"`
				} `tfsdk:"originating_tls" yaml:"originatingTLS,omitempty"`

				Ports *[]struct {
					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				Rules *struct {
					Dns *[]struct {
						MatchName *string `tfsdk:"match_name" yaml:"matchName,omitempty"`

						MatchPattern *string `tfsdk:"match_pattern" yaml:"matchPattern,omitempty"`
					} `tfsdk:"dns" yaml:"dns,omitempty"`

					Http *[]struct {
						HeaderMatches *[]struct {
							Mismatch *string `tfsdk:"mismatch" yaml:"mismatch,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Secret *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"secret" yaml:"secret,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"header_matches" yaml:"headerMatches,omitempty"`

						Headers *[]string `tfsdk:"headers" yaml:"headers,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Method *string `tfsdk:"method" yaml:"method,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"http" yaml:"http,omitempty"`

					Kafka *[]struct {
						ApiKey *string `tfsdk:"api_key" yaml:"apiKey,omitempty"`

						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						ClientID *string `tfsdk:"client_id" yaml:"clientID,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Topic *string `tfsdk:"topic" yaml:"topic,omitempty"`
					} `tfsdk:"kafka" yaml:"kafka,omitempty"`

					L7 *[]map[string]string `tfsdk:"l7" yaml:"l7,omitempty"`

					L7proto *string `tfsdk:"l7proto" yaml:"l7proto,omitempty"`
				} `tfsdk:"rules" yaml:"rules,omitempty"`

				TerminatingTLS *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Secret *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`

					TrustedCA *string `tfsdk:"trusted_ca" yaml:"trustedCA,omitempty"`
				} `tfsdk:"terminating_tls" yaml:"terminatingTLS,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`
		} `tfsdk:"ingress" yaml:"ingress,omitempty"`

		IngressDeny *[]struct {
			FromCIDR *[]string `tfsdk:"from_cidr" yaml:"fromCIDR,omitempty"`

			FromCIDRSet *[]struct {
				Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

				Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
			} `tfsdk:"from_cidr_set" yaml:"fromCIDRSet,omitempty"`

			FromEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"from_endpoints" yaml:"fromEndpoints,omitempty"`

			FromEntities *[]string `tfsdk:"from_entities" yaml:"fromEntities,omitempty"`

			FromRequires *[]struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"from_requires" yaml:"fromRequires,omitempty"`

			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" yaml:"family,omitempty"`

					Type *int64 `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`
			} `tfsdk:"icmps" yaml:"icmps,omitempty"`

			ToPorts *[]struct {
				Ports *[]struct {
					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`
			} `tfsdk:"to_ports" yaml:"toPorts,omitempty"`
		} `tfsdk:"ingress_deny" yaml:"ingressDeny,omitempty"`

		Labels *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Source *string `tfsdk:"source" yaml:"source,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"labels" yaml:"labels,omitempty"`

		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`
	} `tfsdk:"specs" yaml:"specs,omitempty"`
}

func NewCiliumIoCiliumClusterwideNetworkPolicyV2Resource() resource.Resource {
	return &CiliumIoCiliumClusterwideNetworkPolicyV2Resource{}
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cilium_io_cilium_clusterwide_network_policy_v2"
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CiliumClusterwideNetworkPolicy is a Kubernetes third-party resource with an modified version of CiliumNetworkPolicy which is cluster scoped rather than namespace scoped.",
		MarkdownDescription: "CiliumClusterwideNetworkPolicy is a Kubernetes third-party resource with an modified version of CiliumNetworkPolicy which is cluster scoped rather than namespace scoped.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Spec is the desired Cilium specific rule specification.",
				MarkdownDescription: "Spec is the desired Cilium specific rule specification.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"description": {
						Description:         "Description is a free form string, it can be used by the creator of the rule to store human readable explanation of the purpose of this rule. Rules cannot be identified by comment.",
						MarkdownDescription: "Description is a free form string, it can be used by the creator of the rule to store human readable explanation of the purpose of this rule. Rules cannot be identified by comment.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"egress": {
						Description:         "Egress is a list of EgressRule which are enforced at egress. If omitted or empty, this rule does not apply at egress.",
						MarkdownDescription: "Egress is a list of EgressRule which are enforced at egress. If omitted or empty, this rule does not apply at egress.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"icmps": {
								Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is allowed to initiate type 8 ICMP connections.",
								MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is allowed to initiate type 8 ICMP connections.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"fields": {
										Description:         "Fields is a list of ICMP fields.",
										MarkdownDescription: "Fields is a list of ICMP fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"family": {
												Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
												MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("IPv4", "IPv6"),
												},
											},

											"type": {
												Description:         "Type is a ICMP-type. It should be 0-255 (8bit).",
												MarkdownDescription: "Type is a ICMP-type. It should be 0-255 (8bit).",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(255),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_cidr": {
								Description:         "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
								MarkdownDescription: "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_cidr_set": {
								Description:         "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
								MarkdownDescription: "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "CIDR is a CIDR prefix / IP Block.",
										MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
										},
									},

									"except": {
										Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
										MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_endpoints": {
								Description:         "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
								MarkdownDescription: "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_entities": {
								Description:         "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster' and 'host'",
								MarkdownDescription: "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster' and 'host'",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_fqd_ns": {
								Description:         "ToFQDN allows whitelisting DNS names in place of IPs. The IPs that result from DNS resolution of 'ToFQDN.MatchName's are added to the same EgressRule object as ToCIDRSet entries, and behave accordingly. Any L4 and L7 rules within this EgressRule will also apply to these IPs. The DNS -> IP mapping is re-resolved periodically from within the cilium-agent, and the IPs in the DNS response are effected in the policy for selected pods as-is (i.e. the list of IPs is not modified in any way). Note: An explicit rule to allow for DNS traffic is needed for the pods, as ToFQDN counts as an egress rule and will enforce egress policy when PolicyEnforcment=default. Note: If the resolved IPs are IPs within the kubernetes cluster, the ToFQDN rule will not apply to that IP. Note: ToFQDN cannot occur in the same policy as other To* rules.  The current implementation has a number of limitations: - The DNS resolution originates from cilium-agent, and not from the pods. Differences between the responses seen by cilium agent and a particular pod will whitelist the incorrect IP. - DNS TTLs are ignored, and cilium-agent will repoll on a short interval (5 seconds). Each change to the DNS data will trigger a policy regeneration. This may result in delayed updates to the policy for an endpoint when the data changes often or the system is under load.",
								MarkdownDescription: "ToFQDN allows whitelisting DNS names in place of IPs. The IPs that result from DNS resolution of 'ToFQDN.MatchName's are added to the same EgressRule object as ToCIDRSet entries, and behave accordingly. Any L4 and L7 rules within this EgressRule will also apply to these IPs. The DNS -> IP mapping is re-resolved periodically from within the cilium-agent, and the IPs in the DNS response are effected in the policy for selected pods as-is (i.e. the list of IPs is not modified in any way). Note: An explicit rule to allow for DNS traffic is needed for the pods, as ToFQDN counts as an egress rule and will enforce egress policy when PolicyEnforcment=default. Note: If the resolved IPs are IPs within the kubernetes cluster, the ToFQDN rule will not apply to that IP. Note: ToFQDN cannot occur in the same policy as other To* rules.  The current implementation has a number of limitations: - The DNS resolution originates from cilium-agent, and not from the pods. Differences between the responses seen by cilium agent and a particular pod will whitelist the incorrect IP. - DNS TTLs are ignored, and cilium-agent will repoll on a short interval (5 seconds). Each change to the DNS data will trigger a policy regeneration. This may result in delayed updates to the policy for an endpoint when the data changes often or the system is under load.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_name": {
										Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
										MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
										},
									},

									"match_pattern": {
										Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",
										MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_groups": {
								Description:         "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws:     securityGroupsIds:     - 'sg-XXXXXXXXXXXXX'",
								MarkdownDescription: "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws:     securityGroupsIds:     - 'sg-XXXXXXXXXXXXX'",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"aws": {
										Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
										MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"labels": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"region": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_groups_ids": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_groups_names": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is allowed to initiate connections to destination port 8080/tcp",
								MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is allowed to initiate connections to destination port 8080/tcp",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"originating_tls": {
										Description:         "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
										MarkdownDescription: "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",
												MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
														MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"trusted_ca": {
												Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "Ports is a list of L4 port/protocol",
										MarkdownDescription: "Ports is a list of L4 port/protocol",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
												MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rules": {
										Description:         "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
										MarkdownDescription: "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dns": {
												Description:         "DNS-specific rules.",
												MarkdownDescription: "DNS-specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"match_name": {
														Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
														MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
														},
													},

													"match_pattern": {
														Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",
														MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http": {
												Description:         "HTTP specific rules.",
												MarkdownDescription: "HTTP specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"header_matches": {
														Description:         "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
														MarkdownDescription: "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"mismatch": {
																Description:         "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																MarkdownDescription: "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("LOG", "ADD", "DELETE", "REPLACE"),
																},
															},

															"name": {
																Description:         "Name identifies the header.",
																MarkdownDescription: "Name identifies the header.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																MarkdownDescription: "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name is the name of the secret.",
																		MarkdownDescription: "Name is the name of the secret.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"namespace": {
																		Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																		MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																MarkdownDescription: "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": {
														Description:         "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
														MarkdownDescription: "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host": {
														Description:         "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
														MarkdownDescription: "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
														MarkdownDescription: "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
														MarkdownDescription: "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kafka": {
												Description:         "Kafka-specific rules.",
												MarkdownDescription: "Kafka-specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_key": {
														Description:         "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
														MarkdownDescription: "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"api_version": {
														Description:         "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
														MarkdownDescription: "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"client_id": {
														Description:         "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
														MarkdownDescription: "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"role": {
														Description:         "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported:  - 'produce': Allow producing to the topics specified in the rule  - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
														MarkdownDescription: "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported:  - 'produce': Allow producing to the topics specified in the rule  - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("produce", "consume"),
														},
													},

													"topic": {
														Description:         "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
														MarkdownDescription: "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(255),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"l7": {
												Description:         "Key-value pair rules.",
												MarkdownDescription: "Key-value pair rules.",

												Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"l7proto": {
												Description:         "Name of the L7 protocol for which the Key-value pair rules apply.",
												MarkdownDescription: "Name of the L7 protocol for which the Key-value pair rules apply.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"terminating_tls": {
										Description:         "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
										MarkdownDescription: "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",
												MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
														MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"trusted_ca": {
												Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_requires": {
								Description:         "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
								MarkdownDescription: "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_services": {
								Description:         "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
								MarkdownDescription: "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"k8s_service": {
										Description:         "K8sService selects service by name and namespace pair",
										MarkdownDescription: "K8sService selects service by name and namespace pair",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"k8s_service_selector": {
										Description:         "K8sServiceSelector selects services by k8s labels and namespace",
										MarkdownDescription: "K8sServiceSelector selects services by k8s labels and namespace",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "ServiceSelector is a label selector for k8s services",
												MarkdownDescription: "ServiceSelector is a label selector for k8s services",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"egress_deny": {
						Description:         "EgressDeny is a list of EgressDenyRule which are enforced at egress. Any rule inserted here will by denied regardless of the allowed egress rules in the 'egress' field. If omitted or empty, this rule does not apply at egress.",
						MarkdownDescription: "EgressDeny is a list of EgressDenyRule which are enforced at egress. Any rule inserted here will by denied regardless of the allowed egress rules in the 'egress' field. If omitted or empty, this rule does not apply at egress.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"icmps": {
								Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is not allowed to initiate type 8 ICMP connections.",
								MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is not allowed to initiate type 8 ICMP connections.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"fields": {
										Description:         "Fields is a list of ICMP fields.",
										MarkdownDescription: "Fields is a list of ICMP fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"family": {
												Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
												MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("IPv4", "IPv6"),
												},
											},

											"type": {
												Description:         "Type is a ICMP-type. It should be 0-255 (8bit).",
												MarkdownDescription: "Type is a ICMP-type. It should be 0-255 (8bit).",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(255),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_cidr": {
								Description:         "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
								MarkdownDescription: "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_cidr_set": {
								Description:         "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
								MarkdownDescription: "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "CIDR is a CIDR prefix / IP Block.",
										MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
										},
									},

									"except": {
										Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
										MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_endpoints": {
								Description:         "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
								MarkdownDescription: "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_entities": {
								Description:         "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster' and 'host'",
								MarkdownDescription: "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster' and 'host'",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_groups": {
								Description:         "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws:     securityGroupsIds:     - 'sg-XXXXXXXXXXXXX'",
								MarkdownDescription: "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws:     securityGroupsIds:     - 'sg-XXXXXXXXXXXXX'",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"aws": {
										Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
										MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"labels": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"region": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_groups_ids": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_groups_names": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is not allowed to initiate connections to destination port 8080/tcp",
								MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is not allowed to initiate connections to destination port 8080/tcp",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"ports": {
										Description:         "Ports is a list of L4 port/protocol",
										MarkdownDescription: "Ports is a list of L4 port/protocol",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
												MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_requires": {
								Description:         "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
								MarkdownDescription: "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_services": {
								Description:         "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
								MarkdownDescription: "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"k8s_service": {
										Description:         "K8sService selects service by name and namespace pair",
										MarkdownDescription: "K8sService selects service by name and namespace pair",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"k8s_service_selector": {
										Description:         "K8sServiceSelector selects services by k8s labels and namespace",
										MarkdownDescription: "K8sServiceSelector selects services by k8s labels and namespace",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "ServiceSelector is a label selector for k8s services",
												MarkdownDescription: "ServiceSelector is a label selector for k8s services",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"endpoint_selector": {
						Description:         "EndpointSelector selects all endpoints which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive.",
						MarkdownDescription: "EndpointSelector selects all endpoints which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
										},
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": {
						Description:         "Ingress is a list of IngressRule which are enforced at ingress. If omitted or empty, this rule does not apply at ingress.",
						MarkdownDescription: "Ingress is a list of IngressRule which are enforced at ingress. If omitted or empty, this rule does not apply at ingress.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from_cidr": {
								Description:         "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
								MarkdownDescription: "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_cidr_set": {
								Description:         "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
								MarkdownDescription: "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "CIDR is a CIDR prefix / IP Block.",
										MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
										},
									},

									"except": {
										Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
										MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_endpoints": {
								Description:         "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
								MarkdownDescription: "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_entities": {
								Description:         "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
								MarkdownDescription: "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_requires": {
								Description:         "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
								MarkdownDescription: "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"icmps": {
								Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming type 8 ICMP connections.",
								MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming type 8 ICMP connections.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"fields": {
										Description:         "Fields is a list of ICMP fields.",
										MarkdownDescription: "Fields is a list of ICMP fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"family": {
												Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
												MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("IPv4", "IPv6"),
												},
											},

											"type": {
												Description:         "Type is a ICMP-type. It should be 0-255 (8bit).",
												MarkdownDescription: "Type is a ICMP-type. It should be 0-255 (8bit).",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(255),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming connections on port 80/tcp.",
								MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming connections on port 80/tcp.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"originating_tls": {
										Description:         "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
										MarkdownDescription: "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",
												MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
														MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"trusted_ca": {
												Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "Ports is a list of L4 port/protocol",
										MarkdownDescription: "Ports is a list of L4 port/protocol",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
												MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rules": {
										Description:         "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
										MarkdownDescription: "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dns": {
												Description:         "DNS-specific rules.",
												MarkdownDescription: "DNS-specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"match_name": {
														Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
														MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
														},
													},

													"match_pattern": {
														Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",
														MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http": {
												Description:         "HTTP specific rules.",
												MarkdownDescription: "HTTP specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"header_matches": {
														Description:         "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
														MarkdownDescription: "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"mismatch": {
																Description:         "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																MarkdownDescription: "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("LOG", "ADD", "DELETE", "REPLACE"),
																},
															},

															"name": {
																Description:         "Name identifies the header.",
																MarkdownDescription: "Name identifies the header.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																MarkdownDescription: "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name is the name of the secret.",
																		MarkdownDescription: "Name is the name of the secret.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"namespace": {
																		Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																		MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																MarkdownDescription: "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": {
														Description:         "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
														MarkdownDescription: "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host": {
														Description:         "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
														MarkdownDescription: "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
														MarkdownDescription: "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
														MarkdownDescription: "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kafka": {
												Description:         "Kafka-specific rules.",
												MarkdownDescription: "Kafka-specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_key": {
														Description:         "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
														MarkdownDescription: "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"api_version": {
														Description:         "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
														MarkdownDescription: "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"client_id": {
														Description:         "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
														MarkdownDescription: "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"role": {
														Description:         "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported:  - 'produce': Allow producing to the topics specified in the rule  - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
														MarkdownDescription: "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported:  - 'produce': Allow producing to the topics specified in the rule  - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("produce", "consume"),
														},
													},

													"topic": {
														Description:         "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
														MarkdownDescription: "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(255),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"l7": {
												Description:         "Key-value pair rules.",
												MarkdownDescription: "Key-value pair rules.",

												Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"l7proto": {
												Description:         "Name of the L7 protocol for which the Key-value pair rules apply.",
												MarkdownDescription: "Name of the L7 protocol for which the Key-value pair rules apply.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"terminating_tls": {
										Description:         "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
										MarkdownDescription: "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",
												MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
														MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"trusted_ca": {
												Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_deny": {
						Description:         "IngressDeny is a list of IngressDenyRule which are enforced at ingress. Any rule inserted here will by denied regardless of the allowed ingress rules in the 'ingress' field. If omitted or empty, this rule does not apply at ingress.",
						MarkdownDescription: "IngressDeny is a list of IngressDenyRule which are enforced at ingress. Any rule inserted here will by denied regardless of the allowed ingress rules in the 'ingress' field. If omitted or empty, this rule does not apply at ingress.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from_cidr": {
								Description:         "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
								MarkdownDescription: "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_cidr_set": {
								Description:         "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
								MarkdownDescription: "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "CIDR is a CIDR prefix / IP Block.",
										MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
										},
									},

									"except": {
										Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
										MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_endpoints": {
								Description:         "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
								MarkdownDescription: "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_entities": {
								Description:         "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
								MarkdownDescription: "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_requires": {
								Description:         "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
								MarkdownDescription: "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"icmps": {
								Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming type 8 ICMP connections.",
								MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming type 8 ICMP connections.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"fields": {
										Description:         "Fields is a list of ICMP fields.",
										MarkdownDescription: "Fields is a list of ICMP fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"family": {
												Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
												MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("IPv4", "IPv6"),
												},
											},

											"type": {
												Description:         "Type is a ICMP-type. It should be 0-255 (8bit).",
												MarkdownDescription: "Type is a ICMP-type. It should be 0-255 (8bit).",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(255),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming connections on port 80/tcp.",
								MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming connections on port 80/tcp.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"ports": {
										Description:         "Ports is a list of L4 port/protocol",
										MarkdownDescription: "Ports is a list of L4 port/protocol",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
												MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"labels": {
						Description:         "Labels is a list of optional strings which can be used to re-identify the rule or to store metadata. It is possible to lookup or delete strings based on labels. Labels are not required to be unique, multiple rules can have overlapping or identical labels.",
						MarkdownDescription: "Labels is a list of optional strings which can be used to re-identify the rule or to store metadata. It is possible to lookup or delete strings based on labels. Labels are not required to be unique, multiple rules can have overlapping or identical labels.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"source": {
								Description:         "Source can be one of the above values (e.g.: LabelSourceContainer).",
								MarkdownDescription: "Source can be one of the above values (e.g.: LabelSourceContainer).",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "NodeSelector selects all nodes which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive. Can only be used in CiliumClusterwideNetworkPolicies.",
						MarkdownDescription: "NodeSelector selects all nodes which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive. Can only be used in CiliumClusterwideNetworkPolicies.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
										},
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},

			"specs": {
				Description:         "Specs is a list of desired Cilium specific rule specification.",
				MarkdownDescription: "Specs is a list of desired Cilium specific rule specification.",

				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

					"description": {
						Description:         "Description is a free form string, it can be used by the creator of the rule to store human readable explanation of the purpose of this rule. Rules cannot be identified by comment.",
						MarkdownDescription: "Description is a free form string, it can be used by the creator of the rule to store human readable explanation of the purpose of this rule. Rules cannot be identified by comment.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"egress": {
						Description:         "Egress is a list of EgressRule which are enforced at egress. If omitted or empty, this rule does not apply at egress.",
						MarkdownDescription: "Egress is a list of EgressRule which are enforced at egress. If omitted or empty, this rule does not apply at egress.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"icmps": {
								Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is allowed to initiate type 8 ICMP connections.",
								MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is allowed to initiate type 8 ICMP connections.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"fields": {
										Description:         "Fields is a list of ICMP fields.",
										MarkdownDescription: "Fields is a list of ICMP fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"family": {
												Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
												MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("IPv4", "IPv6"),
												},
											},

											"type": {
												Description:         "Type is a ICMP-type. It should be 0-255 (8bit).",
												MarkdownDescription: "Type is a ICMP-type. It should be 0-255 (8bit).",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(255),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_cidr": {
								Description:         "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
								MarkdownDescription: "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_cidr_set": {
								Description:         "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
								MarkdownDescription: "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "CIDR is a CIDR prefix / IP Block.",
										MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
										},
									},

									"except": {
										Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
										MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_endpoints": {
								Description:         "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
								MarkdownDescription: "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_entities": {
								Description:         "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster' and 'host'",
								MarkdownDescription: "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster' and 'host'",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_fqd_ns": {
								Description:         "ToFQDN allows whitelisting DNS names in place of IPs. The IPs that result from DNS resolution of 'ToFQDN.MatchName's are added to the same EgressRule object as ToCIDRSet entries, and behave accordingly. Any L4 and L7 rules within this EgressRule will also apply to these IPs. The DNS -> IP mapping is re-resolved periodically from within the cilium-agent, and the IPs in the DNS response are effected in the policy for selected pods as-is (i.e. the list of IPs is not modified in any way). Note: An explicit rule to allow for DNS traffic is needed for the pods, as ToFQDN counts as an egress rule and will enforce egress policy when PolicyEnforcment=default. Note: If the resolved IPs are IPs within the kubernetes cluster, the ToFQDN rule will not apply to that IP. Note: ToFQDN cannot occur in the same policy as other To* rules.  The current implementation has a number of limitations: - The DNS resolution originates from cilium-agent, and not from the pods. Differences between the responses seen by cilium agent and a particular pod will whitelist the incorrect IP. - DNS TTLs are ignored, and cilium-agent will repoll on a short interval (5 seconds). Each change to the DNS data will trigger a policy regeneration. This may result in delayed updates to the policy for an endpoint when the data changes often or the system is under load.",
								MarkdownDescription: "ToFQDN allows whitelisting DNS names in place of IPs. The IPs that result from DNS resolution of 'ToFQDN.MatchName's are added to the same EgressRule object as ToCIDRSet entries, and behave accordingly. Any L4 and L7 rules within this EgressRule will also apply to these IPs. The DNS -> IP mapping is re-resolved periodically from within the cilium-agent, and the IPs in the DNS response are effected in the policy for selected pods as-is (i.e. the list of IPs is not modified in any way). Note: An explicit rule to allow for DNS traffic is needed for the pods, as ToFQDN counts as an egress rule and will enforce egress policy when PolicyEnforcment=default. Note: If the resolved IPs are IPs within the kubernetes cluster, the ToFQDN rule will not apply to that IP. Note: ToFQDN cannot occur in the same policy as other To* rules.  The current implementation has a number of limitations: - The DNS resolution originates from cilium-agent, and not from the pods. Differences between the responses seen by cilium agent and a particular pod will whitelist the incorrect IP. - DNS TTLs are ignored, and cilium-agent will repoll on a short interval (5 seconds). Each change to the DNS data will trigger a policy regeneration. This may result in delayed updates to the policy for an endpoint when the data changes often or the system is under load.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_name": {
										Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
										MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
										},
									},

									"match_pattern": {
										Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",
										MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_groups": {
								Description:         "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws:     securityGroupsIds:     - 'sg-XXXXXXXXXXXXX'",
								MarkdownDescription: "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws:     securityGroupsIds:     - 'sg-XXXXXXXXXXXXX'",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"aws": {
										Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
										MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"labels": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"region": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_groups_ids": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_groups_names": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is allowed to initiate connections to destination port 8080/tcp",
								MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is allowed to initiate connections to destination port 8080/tcp",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"originating_tls": {
										Description:         "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
										MarkdownDescription: "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",
												MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
														MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"trusted_ca": {
												Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "Ports is a list of L4 port/protocol",
										MarkdownDescription: "Ports is a list of L4 port/protocol",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
												MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rules": {
										Description:         "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
										MarkdownDescription: "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dns": {
												Description:         "DNS-specific rules.",
												MarkdownDescription: "DNS-specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"match_name": {
														Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
														MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
														},
													},

													"match_pattern": {
														Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",
														MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http": {
												Description:         "HTTP specific rules.",
												MarkdownDescription: "HTTP specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"header_matches": {
														Description:         "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
														MarkdownDescription: "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"mismatch": {
																Description:         "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																MarkdownDescription: "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("LOG", "ADD", "DELETE", "REPLACE"),
																},
															},

															"name": {
																Description:         "Name identifies the header.",
																MarkdownDescription: "Name identifies the header.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																MarkdownDescription: "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name is the name of the secret.",
																		MarkdownDescription: "Name is the name of the secret.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"namespace": {
																		Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																		MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																MarkdownDescription: "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": {
														Description:         "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
														MarkdownDescription: "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host": {
														Description:         "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
														MarkdownDescription: "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
														MarkdownDescription: "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
														MarkdownDescription: "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kafka": {
												Description:         "Kafka-specific rules.",
												MarkdownDescription: "Kafka-specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_key": {
														Description:         "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
														MarkdownDescription: "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"api_version": {
														Description:         "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
														MarkdownDescription: "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"client_id": {
														Description:         "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
														MarkdownDescription: "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"role": {
														Description:         "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported:  - 'produce': Allow producing to the topics specified in the rule  - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
														MarkdownDescription: "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported:  - 'produce': Allow producing to the topics specified in the rule  - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("produce", "consume"),
														},
													},

													"topic": {
														Description:         "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
														MarkdownDescription: "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(255),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"l7": {
												Description:         "Key-value pair rules.",
												MarkdownDescription: "Key-value pair rules.",

												Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"l7proto": {
												Description:         "Name of the L7 protocol for which the Key-value pair rules apply.",
												MarkdownDescription: "Name of the L7 protocol for which the Key-value pair rules apply.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"terminating_tls": {
										Description:         "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
										MarkdownDescription: "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",
												MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
														MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"trusted_ca": {
												Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_requires": {
								Description:         "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
								MarkdownDescription: "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_services": {
								Description:         "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
								MarkdownDescription: "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"k8s_service": {
										Description:         "K8sService selects service by name and namespace pair",
										MarkdownDescription: "K8sService selects service by name and namespace pair",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"k8s_service_selector": {
										Description:         "K8sServiceSelector selects services by k8s labels and namespace",
										MarkdownDescription: "K8sServiceSelector selects services by k8s labels and namespace",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "ServiceSelector is a label selector for k8s services",
												MarkdownDescription: "ServiceSelector is a label selector for k8s services",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"egress_deny": {
						Description:         "EgressDeny is a list of EgressDenyRule which are enforced at egress. Any rule inserted here will by denied regardless of the allowed egress rules in the 'egress' field. If omitted or empty, this rule does not apply at egress.",
						MarkdownDescription: "EgressDeny is a list of EgressDenyRule which are enforced at egress. Any rule inserted here will by denied regardless of the allowed egress rules in the 'egress' field. If omitted or empty, this rule does not apply at egress.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"icmps": {
								Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is not allowed to initiate type 8 ICMP connections.",
								MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is not allowed to initiate type 8 ICMP connections.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"fields": {
										Description:         "Fields is a list of ICMP fields.",
										MarkdownDescription: "Fields is a list of ICMP fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"family": {
												Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
												MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("IPv4", "IPv6"),
												},
											},

											"type": {
												Description:         "Type is a ICMP-type. It should be 0-255 (8bit).",
												MarkdownDescription: "Type is a ICMP-type. It should be 0-255 (8bit).",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(255),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_cidr": {
								Description:         "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
								MarkdownDescription: "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_cidr_set": {
								Description:         "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
								MarkdownDescription: "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "CIDR is a CIDR prefix / IP Block.",
										MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
										},
									},

									"except": {
										Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
										MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_endpoints": {
								Description:         "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
								MarkdownDescription: "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_entities": {
								Description:         "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster' and 'host'",
								MarkdownDescription: "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster' and 'host'",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_groups": {
								Description:         "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws:     securityGroupsIds:     - 'sg-XXXXXXXXXXXXX'",
								MarkdownDescription: "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws:     securityGroupsIds:     - 'sg-XXXXXXXXXXXXX'",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"aws": {
										Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
										MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"labels": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"region": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_groups_ids": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_groups_names": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is not allowed to initiate connections to destination port 8080/tcp",
								MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is not allowed to initiate connections to destination port 8080/tcp",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"ports": {
										Description:         "Ports is a list of L4 port/protocol",
										MarkdownDescription: "Ports is a list of L4 port/protocol",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
												MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_requires": {
								Description:         "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
								MarkdownDescription: "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_services": {
								Description:         "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
								MarkdownDescription: "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"k8s_service": {
										Description:         "K8sService selects service by name and namespace pair",
										MarkdownDescription: "K8sService selects service by name and namespace pair",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"service_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"k8s_service_selector": {
										Description:         "K8sServiceSelector selects services by k8s labels and namespace",
										MarkdownDescription: "K8sServiceSelector selects services by k8s labels and namespace",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"namespace": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"selector": {
												Description:         "ServiceSelector is a label selector for k8s services",
												MarkdownDescription: "ServiceSelector is a label selector for k8s services",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"operator": {
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": {
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_labels": {
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"endpoint_selector": {
						Description:         "EndpointSelector selects all endpoints which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive.",
						MarkdownDescription: "EndpointSelector selects all endpoints which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
										},
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": {
						Description:         "Ingress is a list of IngressRule which are enforced at ingress. If omitted or empty, this rule does not apply at ingress.",
						MarkdownDescription: "Ingress is a list of IngressRule which are enforced at ingress. If omitted or empty, this rule does not apply at ingress.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from_cidr": {
								Description:         "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
								MarkdownDescription: "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_cidr_set": {
								Description:         "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
								MarkdownDescription: "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "CIDR is a CIDR prefix / IP Block.",
										MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
										},
									},

									"except": {
										Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
										MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_endpoints": {
								Description:         "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
								MarkdownDescription: "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_entities": {
								Description:         "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
								MarkdownDescription: "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_requires": {
								Description:         "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
								MarkdownDescription: "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"icmps": {
								Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming type 8 ICMP connections.",
								MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming type 8 ICMP connections.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"fields": {
										Description:         "Fields is a list of ICMP fields.",
										MarkdownDescription: "Fields is a list of ICMP fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"family": {
												Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
												MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("IPv4", "IPv6"),
												},
											},

											"type": {
												Description:         "Type is a ICMP-type. It should be 0-255 (8bit).",
												MarkdownDescription: "Type is a ICMP-type. It should be 0-255 (8bit).",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(255),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming connections on port 80/tcp.",
								MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming connections on port 80/tcp.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"originating_tls": {
										Description:         "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
										MarkdownDescription: "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",
												MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
														MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"trusted_ca": {
												Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ports": {
										Description:         "Ports is a list of L4 port/protocol",
										MarkdownDescription: "Ports is a list of L4 port/protocol",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
												MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"rules": {
										Description:         "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
										MarkdownDescription: "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dns": {
												Description:         "DNS-specific rules.",
												MarkdownDescription: "DNS-specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"match_name": {
														Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
														MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
														},
													},

													"match_pattern": {
														Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",
														MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level   www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io'   except those containing '.' separator, subcilium.io and sub-cilium.io match,   www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub'   sub.cilium.io and subdomain.cilium.io match, www.cilium.io,   blog.cilium.io, cilium.io and google.com do not",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http": {
												Description:         "HTTP specific rules.",
												MarkdownDescription: "HTTP specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"header_matches": {
														Description:         "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
														MarkdownDescription: "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"mismatch": {
																Description:         "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																MarkdownDescription: "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("LOG", "ADD", "DELETE", "REPLACE"),
																},
															},

															"name": {
																Description:         "Name identifies the header.",
																MarkdownDescription: "Name identifies the header.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"secret": {
																Description:         "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																MarkdownDescription: "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name is the name of the secret.",
																		MarkdownDescription: "Name is the name of the secret.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"namespace": {
																		Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																		MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
																Description:         "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																MarkdownDescription: "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": {
														Description:         "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
														MarkdownDescription: "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host": {
														Description:         "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
														MarkdownDescription: "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
														MarkdownDescription: "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
														MarkdownDescription: "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kafka": {
												Description:         "Kafka-specific rules.",
												MarkdownDescription: "Kafka-specific rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"api_key": {
														Description:         "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
														MarkdownDescription: "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"api_version": {
														Description:         "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
														MarkdownDescription: "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"client_id": {
														Description:         "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
														MarkdownDescription: "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"role": {
														Description:         "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported:  - 'produce': Allow producing to the topics specified in the rule  - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
														MarkdownDescription: "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported:  - 'produce': Allow producing to the topics specified in the rule  - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("produce", "consume"),
														},
													},

													"topic": {
														Description:         "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
														MarkdownDescription: "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(255),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"l7": {
												Description:         "Key-value pair rules.",
												MarkdownDescription: "Key-value pair rules.",

												Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"l7proto": {
												Description:         "Name of the L7 protocol for which the Key-value pair rules apply.",
												MarkdownDescription: "Name of the L7 protocol for which the Key-value pair rules apply.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"terminating_tls": {
										Description:         "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
										MarkdownDescription: "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret": {
												Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",
												MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items:  - 'ca.crt'  - Which represents the trusted CA to verify remote source.  - 'tls.crt' - Which represents the public key certificate.  - 'tls.key' - Which represents the private key matching the public key                certificate.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name is the name of the secret.",
														MarkdownDescription: "Name is the name of the secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
														Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
														MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"trusted_ca": {
												Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
												MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_deny": {
						Description:         "IngressDeny is a list of IngressDenyRule which are enforced at ingress. Any rule inserted here will by denied regardless of the allowed ingress rules in the 'ingress' field. If omitted or empty, this rule does not apply at ingress.",
						MarkdownDescription: "IngressDeny is a list of IngressDenyRule which are enforced at ingress. Any rule inserted here will by denied regardless of the allowed ingress rules in the 'ingress' field. If omitted or empty, this rule does not apply at ingress.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"from_cidr": {
								Description:         "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
								MarkdownDescription: "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_cidr_set": {
								Description:         "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
								MarkdownDescription: "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"cidr": {
										Description:         "CIDR is a CIDR prefix / IP Block.",
										MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
										},
									},

									"except": {
										Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
										MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_endpoints": {
								Description:         "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
								MarkdownDescription: "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_entities": {
								Description:         "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
								MarkdownDescription: "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"from_requires": {
								Description:         "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
								MarkdownDescription: "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"icmps": {
								Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming type 8 ICMP connections.",
								MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming type 8 ICMP connections.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"fields": {
										Description:         "Fields is a list of ICMP fields.",
										MarkdownDescription: "Fields is a list of ICMP fields.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"family": {
												Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
												MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("IPv4", "IPv6"),
												},
											},

											"type": {
												Description:         "Type is a ICMP-type. It should be 0-255 (8bit).",
												MarkdownDescription: "Type is a ICMP-type. It should be 0-255 (8bit).",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(255),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"to_ports": {
								Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming connections on port 80/tcp.",
								MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming connections on port 80/tcp.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"ports": {
										Description:         "Ports is a list of L4 port/protocol",
										MarkdownDescription: "Ports is a list of L4 port/protocol",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
												MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
												},
											},

											"protocol": {
												Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
												MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"labels": {
						Description:         "Labels is a list of optional strings which can be used to re-identify the rule or to store metadata. It is possible to lookup or delete strings based on labels. Labels are not required to be unique, multiple rules can have overlapping or identical labels.",
						MarkdownDescription: "Labels is a list of optional strings which can be used to re-identify the rule or to store metadata. It is possible to lookup or delete strings based on labels. Labels are not required to be unique, multiple rules can have overlapping or identical labels.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"source": {
								Description:         "Source can be one of the above values (e.g.: LabelSourceContainer).",
								MarkdownDescription: "Source can be one of the above values (e.g.: LabelSourceContainer).",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "NodeSelector selects all nodes which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive. Can only be used in CiliumClusterwideNetworkPolicies.",
						MarkdownDescription: "NodeSelector selects all nodes which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive. Can only be used in CiliumClusterwideNetworkPolicies.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_expressions": {
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "key is the label key that the selector applies to.",
										MarkdownDescription: "key is the label key that the selector applies to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"operator": {
										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
										},
									},

									"values": {
										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": {
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_cilium_io_cilium_clusterwide_network_policy_v2")

	var state CiliumIoCiliumClusterwideNetworkPolicyV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumClusterwideNetworkPolicyV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2")
	goModel.Kind = utilities.Ptr("CiliumClusterwideNetworkPolicy")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_clusterwide_network_policy_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_cilium_io_cilium_clusterwide_network_policy_v2")

	var state CiliumIoCiliumClusterwideNetworkPolicyV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CiliumIoCiliumClusterwideNetworkPolicyV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("cilium.io/v2")
	goModel.Kind = utilities.Ptr("CiliumClusterwideNetworkPolicy")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_cilium_io_cilium_clusterwide_network_policy_v2")
	// NO-OP: Terraform removes the state automatically for us
}
