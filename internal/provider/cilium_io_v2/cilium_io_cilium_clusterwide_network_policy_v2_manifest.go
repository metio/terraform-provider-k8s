/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cilium_io_v2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &CiliumIoCiliumClusterwideNetworkPolicyV2Manifest{}
)

func NewCiliumIoCiliumClusterwideNetworkPolicyV2Manifest() datasource.DataSource {
	return &CiliumIoCiliumClusterwideNetworkPolicyV2Manifest{}
}

type CiliumIoCiliumClusterwideNetworkPolicyV2Manifest struct{}

type CiliumIoCiliumClusterwideNetworkPolicyV2ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Description *string `tfsdk:"description" json:"description,omitempty"`
		Egress      *[]struct {
			Authentication *struct {
				Mode *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" json:"family,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
			} `tfsdk:"icmps" json:"icmps,omitempty"`
			ToCIDR    *[]string `tfsdk:"to_cidr" json:"toCIDR,omitempty"`
			ToCIDRSet *[]struct {
				Cidr         *string   `tfsdk:"cidr" json:"cidr,omitempty"`
				CidrGroupRef *string   `tfsdk:"cidr_group_ref" json:"cidrGroupRef,omitempty"`
				Except       *[]string `tfsdk:"except" json:"except,omitempty"`
			} `tfsdk:"to_cidr_set" json:"toCIDRSet,omitempty"`
			ToEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_endpoints" json:"toEndpoints,omitempty"`
			ToEntities *[]string `tfsdk:"to_entities" json:"toEntities,omitempty"`
			ToFQDNs    *[]struct {
				MatchName    *string `tfsdk:"match_name" json:"matchName,omitempty"`
				MatchPattern *string `tfsdk:"match_pattern" json:"matchPattern,omitempty"`
			} `tfsdk:"to_fqd_ns" json:"toFQDNs,omitempty"`
			ToGroups *[]struct {
				Aws *struct {
					Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Region              *string            `tfsdk:"region" json:"region,omitempty"`
					SecurityGroupsIds   *[]string          `tfsdk:"security_groups_ids" json:"securityGroupsIds,omitempty"`
					SecurityGroupsNames *[]string          `tfsdk:"security_groups_names" json:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
			} `tfsdk:"to_groups" json:"toGroups,omitempty"`
			ToNodes *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_nodes" json:"toNodes,omitempty"`
			ToPorts *[]struct {
				Listener *struct {
					EnvoyConfig *struct {
						Kind *string `tfsdk:"kind" json:"kind,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"envoy_config" json:"envoyConfig,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Priority *int64  `tfsdk:"priority" json:"priority,omitempty"`
				} `tfsdk:"listener" json:"listener,omitempty"`
				OriginatingTLS *struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					PrivateKey  *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					Secret      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					TrustedCA *string `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
				} `tfsdk:"originating_tls" json:"originatingTLS,omitempty"`
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				Rules *struct {
					Dns *[]struct {
						MatchName    *string `tfsdk:"match_name" json:"matchName,omitempty"`
						MatchPattern *string `tfsdk:"match_pattern" json:"matchPattern,omitempty"`
					} `tfsdk:"dns" json:"dns,omitempty"`
					Http *[]struct {
						HeaderMatches *[]struct {
							Mismatch *string `tfsdk:"mismatch" json:"mismatch,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Secret   *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"header_matches" json:"headerMatches,omitempty"`
						Headers *[]string `tfsdk:"headers" json:"headers,omitempty"`
						Host    *string   `tfsdk:"host" json:"host,omitempty"`
						Method  *string   `tfsdk:"method" json:"method,omitempty"`
						Path    *string   `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Kafka *[]struct {
						ApiKey     *string `tfsdk:"api_key" json:"apiKey,omitempty"`
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						ClientID   *string `tfsdk:"client_id" json:"clientID,omitempty"`
						Role       *string `tfsdk:"role" json:"role,omitempty"`
						Topic      *string `tfsdk:"topic" json:"topic,omitempty"`
					} `tfsdk:"kafka" json:"kafka,omitempty"`
					L7      *[]map[string]string `tfsdk:"l7" json:"l7,omitempty"`
					L7proto *string              `tfsdk:"l7proto" json:"l7proto,omitempty"`
				} `tfsdk:"rules" json:"rules,omitempty"`
				ServerNames    *[]string `tfsdk:"server_names" json:"serverNames,omitempty"`
				TerminatingTLS *struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					PrivateKey  *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					Secret      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					TrustedCA *string `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
				} `tfsdk:"terminating_tls" json:"terminatingTLS,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
			ToRequires *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_requires" json:"toRequires,omitempty"`
			ToServices *[]struct {
				K8sService *struct {
					Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"k8s_service" json:"k8sService,omitempty"`
				K8sServiceSelector *struct {
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"k8s_service_selector" json:"k8sServiceSelector,omitempty"`
			} `tfsdk:"to_services" json:"toServices,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		EgressDeny *[]struct {
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" json:"family,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
			} `tfsdk:"icmps" json:"icmps,omitempty"`
			ToCIDR    *[]string `tfsdk:"to_cidr" json:"toCIDR,omitempty"`
			ToCIDRSet *[]struct {
				Cidr         *string   `tfsdk:"cidr" json:"cidr,omitempty"`
				CidrGroupRef *string   `tfsdk:"cidr_group_ref" json:"cidrGroupRef,omitempty"`
				Except       *[]string `tfsdk:"except" json:"except,omitempty"`
			} `tfsdk:"to_cidr_set" json:"toCIDRSet,omitempty"`
			ToEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_endpoints" json:"toEndpoints,omitempty"`
			ToEntities *[]string `tfsdk:"to_entities" json:"toEntities,omitempty"`
			ToGroups   *[]struct {
				Aws *struct {
					Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Region              *string            `tfsdk:"region" json:"region,omitempty"`
					SecurityGroupsIds   *[]string          `tfsdk:"security_groups_ids" json:"securityGroupsIds,omitempty"`
					SecurityGroupsNames *[]string          `tfsdk:"security_groups_names" json:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
			} `tfsdk:"to_groups" json:"toGroups,omitempty"`
			ToNodes *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_nodes" json:"toNodes,omitempty"`
			ToPorts *[]struct {
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
			ToRequires *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_requires" json:"toRequires,omitempty"`
			ToServices *[]struct {
				K8sService *struct {
					Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"k8s_service" json:"k8sService,omitempty"`
				K8sServiceSelector *struct {
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"k8s_service_selector" json:"k8sServiceSelector,omitempty"`
			} `tfsdk:"to_services" json:"toServices,omitempty"`
		} `tfsdk:"egress_deny" json:"egressDeny,omitempty"`
		EnableDefaultDeny *struct {
			Egress  *bool `tfsdk:"egress" json:"egress,omitempty"`
			Ingress *bool `tfsdk:"ingress" json:"ingress,omitempty"`
		} `tfsdk:"enable_default_deny" json:"enableDefaultDeny,omitempty"`
		EndpointSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"endpoint_selector" json:"endpointSelector,omitempty"`
		Ingress *[]struct {
			Authentication *struct {
				Mode *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			FromCIDR    *[]string `tfsdk:"from_cidr" json:"fromCIDR,omitempty"`
			FromCIDRSet *[]struct {
				Cidr         *string   `tfsdk:"cidr" json:"cidr,omitempty"`
				CidrGroupRef *string   `tfsdk:"cidr_group_ref" json:"cidrGroupRef,omitempty"`
				Except       *[]string `tfsdk:"except" json:"except,omitempty"`
			} `tfsdk:"from_cidr_set" json:"fromCIDRSet,omitempty"`
			FromEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_endpoints" json:"fromEndpoints,omitempty"`
			FromEntities *[]string `tfsdk:"from_entities" json:"fromEntities,omitempty"`
			FromGroups   *[]struct {
				Aws *struct {
					Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Region              *string            `tfsdk:"region" json:"region,omitempty"`
					SecurityGroupsIds   *[]string          `tfsdk:"security_groups_ids" json:"securityGroupsIds,omitempty"`
					SecurityGroupsNames *[]string          `tfsdk:"security_groups_names" json:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
			} `tfsdk:"from_groups" json:"fromGroups,omitempty"`
			FromNodes *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_nodes" json:"fromNodes,omitempty"`
			FromRequires *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_requires" json:"fromRequires,omitempty"`
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" json:"family,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
			} `tfsdk:"icmps" json:"icmps,omitempty"`
			ToPorts *[]struct {
				Listener *struct {
					EnvoyConfig *struct {
						Kind *string `tfsdk:"kind" json:"kind,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"envoy_config" json:"envoyConfig,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Priority *int64  `tfsdk:"priority" json:"priority,omitempty"`
				} `tfsdk:"listener" json:"listener,omitempty"`
				OriginatingTLS *struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					PrivateKey  *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					Secret      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					TrustedCA *string `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
				} `tfsdk:"originating_tls" json:"originatingTLS,omitempty"`
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				Rules *struct {
					Dns *[]struct {
						MatchName    *string `tfsdk:"match_name" json:"matchName,omitempty"`
						MatchPattern *string `tfsdk:"match_pattern" json:"matchPattern,omitempty"`
					} `tfsdk:"dns" json:"dns,omitempty"`
					Http *[]struct {
						HeaderMatches *[]struct {
							Mismatch *string `tfsdk:"mismatch" json:"mismatch,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Secret   *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"header_matches" json:"headerMatches,omitempty"`
						Headers *[]string `tfsdk:"headers" json:"headers,omitempty"`
						Host    *string   `tfsdk:"host" json:"host,omitempty"`
						Method  *string   `tfsdk:"method" json:"method,omitempty"`
						Path    *string   `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Kafka *[]struct {
						ApiKey     *string `tfsdk:"api_key" json:"apiKey,omitempty"`
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						ClientID   *string `tfsdk:"client_id" json:"clientID,omitempty"`
						Role       *string `tfsdk:"role" json:"role,omitempty"`
						Topic      *string `tfsdk:"topic" json:"topic,omitempty"`
					} `tfsdk:"kafka" json:"kafka,omitempty"`
					L7      *[]map[string]string `tfsdk:"l7" json:"l7,omitempty"`
					L7proto *string              `tfsdk:"l7proto" json:"l7proto,omitempty"`
				} `tfsdk:"rules" json:"rules,omitempty"`
				ServerNames    *[]string `tfsdk:"server_names" json:"serverNames,omitempty"`
				TerminatingTLS *struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					PrivateKey  *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					Secret      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					TrustedCA *string `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
				} `tfsdk:"terminating_tls" json:"terminatingTLS,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		IngressDeny *[]struct {
			FromCIDR    *[]string `tfsdk:"from_cidr" json:"fromCIDR,omitempty"`
			FromCIDRSet *[]struct {
				Cidr         *string   `tfsdk:"cidr" json:"cidr,omitempty"`
				CidrGroupRef *string   `tfsdk:"cidr_group_ref" json:"cidrGroupRef,omitempty"`
				Except       *[]string `tfsdk:"except" json:"except,omitempty"`
			} `tfsdk:"from_cidr_set" json:"fromCIDRSet,omitempty"`
			FromEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_endpoints" json:"fromEndpoints,omitempty"`
			FromEntities *[]string `tfsdk:"from_entities" json:"fromEntities,omitempty"`
			FromGroups   *[]struct {
				Aws *struct {
					Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Region              *string            `tfsdk:"region" json:"region,omitempty"`
					SecurityGroupsIds   *[]string          `tfsdk:"security_groups_ids" json:"securityGroupsIds,omitempty"`
					SecurityGroupsNames *[]string          `tfsdk:"security_groups_names" json:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
			} `tfsdk:"from_groups" json:"fromGroups,omitempty"`
			FromNodes *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_nodes" json:"fromNodes,omitempty"`
			FromRequires *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_requires" json:"fromRequires,omitempty"`
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" json:"family,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
			} `tfsdk:"icmps" json:"icmps,omitempty"`
			ToPorts *[]struct {
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
		} `tfsdk:"ingress_deny" json:"ingressDeny,omitempty"`
		Labels *[]struct {
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
			Value  *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"labels" json:"labels,omitempty"`
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
	Specs *[]struct {
		Description *string `tfsdk:"description" json:"description,omitempty"`
		Egress      *[]struct {
			Authentication *struct {
				Mode *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" json:"family,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
			} `tfsdk:"icmps" json:"icmps,omitempty"`
			ToCIDR    *[]string `tfsdk:"to_cidr" json:"toCIDR,omitempty"`
			ToCIDRSet *[]struct {
				Cidr         *string   `tfsdk:"cidr" json:"cidr,omitempty"`
				CidrGroupRef *string   `tfsdk:"cidr_group_ref" json:"cidrGroupRef,omitempty"`
				Except       *[]string `tfsdk:"except" json:"except,omitempty"`
			} `tfsdk:"to_cidr_set" json:"toCIDRSet,omitempty"`
			ToEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_endpoints" json:"toEndpoints,omitempty"`
			ToEntities *[]string `tfsdk:"to_entities" json:"toEntities,omitempty"`
			ToFQDNs    *[]struct {
				MatchName    *string `tfsdk:"match_name" json:"matchName,omitempty"`
				MatchPattern *string `tfsdk:"match_pattern" json:"matchPattern,omitempty"`
			} `tfsdk:"to_fqd_ns" json:"toFQDNs,omitempty"`
			ToGroups *[]struct {
				Aws *struct {
					Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Region              *string            `tfsdk:"region" json:"region,omitempty"`
					SecurityGroupsIds   *[]string          `tfsdk:"security_groups_ids" json:"securityGroupsIds,omitempty"`
					SecurityGroupsNames *[]string          `tfsdk:"security_groups_names" json:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
			} `tfsdk:"to_groups" json:"toGroups,omitempty"`
			ToNodes *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_nodes" json:"toNodes,omitempty"`
			ToPorts *[]struct {
				Listener *struct {
					EnvoyConfig *struct {
						Kind *string `tfsdk:"kind" json:"kind,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"envoy_config" json:"envoyConfig,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Priority *int64  `tfsdk:"priority" json:"priority,omitempty"`
				} `tfsdk:"listener" json:"listener,omitempty"`
				OriginatingTLS *struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					PrivateKey  *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					Secret      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					TrustedCA *string `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
				} `tfsdk:"originating_tls" json:"originatingTLS,omitempty"`
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				Rules *struct {
					Dns *[]struct {
						MatchName    *string `tfsdk:"match_name" json:"matchName,omitempty"`
						MatchPattern *string `tfsdk:"match_pattern" json:"matchPattern,omitempty"`
					} `tfsdk:"dns" json:"dns,omitempty"`
					Http *[]struct {
						HeaderMatches *[]struct {
							Mismatch *string `tfsdk:"mismatch" json:"mismatch,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Secret   *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"header_matches" json:"headerMatches,omitempty"`
						Headers *[]string `tfsdk:"headers" json:"headers,omitempty"`
						Host    *string   `tfsdk:"host" json:"host,omitempty"`
						Method  *string   `tfsdk:"method" json:"method,omitempty"`
						Path    *string   `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Kafka *[]struct {
						ApiKey     *string `tfsdk:"api_key" json:"apiKey,omitempty"`
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						ClientID   *string `tfsdk:"client_id" json:"clientID,omitempty"`
						Role       *string `tfsdk:"role" json:"role,omitempty"`
						Topic      *string `tfsdk:"topic" json:"topic,omitempty"`
					} `tfsdk:"kafka" json:"kafka,omitempty"`
					L7      *[]map[string]string `tfsdk:"l7" json:"l7,omitempty"`
					L7proto *string              `tfsdk:"l7proto" json:"l7proto,omitempty"`
				} `tfsdk:"rules" json:"rules,omitempty"`
				ServerNames    *[]string `tfsdk:"server_names" json:"serverNames,omitempty"`
				TerminatingTLS *struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					PrivateKey  *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					Secret      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					TrustedCA *string `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
				} `tfsdk:"terminating_tls" json:"terminatingTLS,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
			ToRequires *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_requires" json:"toRequires,omitempty"`
			ToServices *[]struct {
				K8sService *struct {
					Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"k8s_service" json:"k8sService,omitempty"`
				K8sServiceSelector *struct {
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"k8s_service_selector" json:"k8sServiceSelector,omitempty"`
			} `tfsdk:"to_services" json:"toServices,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		EgressDeny *[]struct {
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" json:"family,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
			} `tfsdk:"icmps" json:"icmps,omitempty"`
			ToCIDR    *[]string `tfsdk:"to_cidr" json:"toCIDR,omitempty"`
			ToCIDRSet *[]struct {
				Cidr         *string   `tfsdk:"cidr" json:"cidr,omitempty"`
				CidrGroupRef *string   `tfsdk:"cidr_group_ref" json:"cidrGroupRef,omitempty"`
				Except       *[]string `tfsdk:"except" json:"except,omitempty"`
			} `tfsdk:"to_cidr_set" json:"toCIDRSet,omitempty"`
			ToEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_endpoints" json:"toEndpoints,omitempty"`
			ToEntities *[]string `tfsdk:"to_entities" json:"toEntities,omitempty"`
			ToGroups   *[]struct {
				Aws *struct {
					Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Region              *string            `tfsdk:"region" json:"region,omitempty"`
					SecurityGroupsIds   *[]string          `tfsdk:"security_groups_ids" json:"securityGroupsIds,omitempty"`
					SecurityGroupsNames *[]string          `tfsdk:"security_groups_names" json:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
			} `tfsdk:"to_groups" json:"toGroups,omitempty"`
			ToNodes *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_nodes" json:"toNodes,omitempty"`
			ToPorts *[]struct {
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
			ToRequires *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"to_requires" json:"toRequires,omitempty"`
			ToServices *[]struct {
				K8sService *struct {
					Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"k8s_service" json:"k8sService,omitempty"`
				K8sServiceSelector *struct {
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Selector  *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"k8s_service_selector" json:"k8sServiceSelector,omitempty"`
			} `tfsdk:"to_services" json:"toServices,omitempty"`
		} `tfsdk:"egress_deny" json:"egressDeny,omitempty"`
		EnableDefaultDeny *struct {
			Egress  *bool `tfsdk:"egress" json:"egress,omitempty"`
			Ingress *bool `tfsdk:"ingress" json:"ingress,omitempty"`
		} `tfsdk:"enable_default_deny" json:"enableDefaultDeny,omitempty"`
		EndpointSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"endpoint_selector" json:"endpointSelector,omitempty"`
		Ingress *[]struct {
			Authentication *struct {
				Mode *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			FromCIDR    *[]string `tfsdk:"from_cidr" json:"fromCIDR,omitempty"`
			FromCIDRSet *[]struct {
				Cidr         *string   `tfsdk:"cidr" json:"cidr,omitempty"`
				CidrGroupRef *string   `tfsdk:"cidr_group_ref" json:"cidrGroupRef,omitempty"`
				Except       *[]string `tfsdk:"except" json:"except,omitempty"`
			} `tfsdk:"from_cidr_set" json:"fromCIDRSet,omitempty"`
			FromEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_endpoints" json:"fromEndpoints,omitempty"`
			FromEntities *[]string `tfsdk:"from_entities" json:"fromEntities,omitempty"`
			FromGroups   *[]struct {
				Aws *struct {
					Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Region              *string            `tfsdk:"region" json:"region,omitempty"`
					SecurityGroupsIds   *[]string          `tfsdk:"security_groups_ids" json:"securityGroupsIds,omitempty"`
					SecurityGroupsNames *[]string          `tfsdk:"security_groups_names" json:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
			} `tfsdk:"from_groups" json:"fromGroups,omitempty"`
			FromNodes *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_nodes" json:"fromNodes,omitempty"`
			FromRequires *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_requires" json:"fromRequires,omitempty"`
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" json:"family,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
			} `tfsdk:"icmps" json:"icmps,omitempty"`
			ToPorts *[]struct {
				Listener *struct {
					EnvoyConfig *struct {
						Kind *string `tfsdk:"kind" json:"kind,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"envoy_config" json:"envoyConfig,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Priority *int64  `tfsdk:"priority" json:"priority,omitempty"`
				} `tfsdk:"listener" json:"listener,omitempty"`
				OriginatingTLS *struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					PrivateKey  *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					Secret      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					TrustedCA *string `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
				} `tfsdk:"originating_tls" json:"originatingTLS,omitempty"`
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				Rules *struct {
					Dns *[]struct {
						MatchName    *string `tfsdk:"match_name" json:"matchName,omitempty"`
						MatchPattern *string `tfsdk:"match_pattern" json:"matchPattern,omitempty"`
					} `tfsdk:"dns" json:"dns,omitempty"`
					Http *[]struct {
						HeaderMatches *[]struct {
							Mismatch *string `tfsdk:"mismatch" json:"mismatch,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Secret   *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"secret" json:"secret,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"header_matches" json:"headerMatches,omitempty"`
						Headers *[]string `tfsdk:"headers" json:"headers,omitempty"`
						Host    *string   `tfsdk:"host" json:"host,omitempty"`
						Method  *string   `tfsdk:"method" json:"method,omitempty"`
						Path    *string   `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Kafka *[]struct {
						ApiKey     *string `tfsdk:"api_key" json:"apiKey,omitempty"`
						ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						ClientID   *string `tfsdk:"client_id" json:"clientID,omitempty"`
						Role       *string `tfsdk:"role" json:"role,omitempty"`
						Topic      *string `tfsdk:"topic" json:"topic,omitempty"`
					} `tfsdk:"kafka" json:"kafka,omitempty"`
					L7      *[]map[string]string `tfsdk:"l7" json:"l7,omitempty"`
					L7proto *string              `tfsdk:"l7proto" json:"l7proto,omitempty"`
				} `tfsdk:"rules" json:"rules,omitempty"`
				ServerNames    *[]string `tfsdk:"server_names" json:"serverNames,omitempty"`
				TerminatingTLS *struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					PrivateKey  *string `tfsdk:"private_key" json:"privateKey,omitempty"`
					Secret      *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					TrustedCA *string `tfsdk:"trusted_ca" json:"trustedCA,omitempty"`
				} `tfsdk:"terminating_tls" json:"terminatingTLS,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		IngressDeny *[]struct {
			FromCIDR    *[]string `tfsdk:"from_cidr" json:"fromCIDR,omitempty"`
			FromCIDRSet *[]struct {
				Cidr         *string   `tfsdk:"cidr" json:"cidr,omitempty"`
				CidrGroupRef *string   `tfsdk:"cidr_group_ref" json:"cidrGroupRef,omitempty"`
				Except       *[]string `tfsdk:"except" json:"except,omitempty"`
			} `tfsdk:"from_cidr_set" json:"fromCIDRSet,omitempty"`
			FromEndpoints *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_endpoints" json:"fromEndpoints,omitempty"`
			FromEntities *[]string `tfsdk:"from_entities" json:"fromEntities,omitempty"`
			FromGroups   *[]struct {
				Aws *struct {
					Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Region              *string            `tfsdk:"region" json:"region,omitempty"`
					SecurityGroupsIds   *[]string          `tfsdk:"security_groups_ids" json:"securityGroupsIds,omitempty"`
					SecurityGroupsNames *[]string          `tfsdk:"security_groups_names" json:"securityGroupsNames,omitempty"`
				} `tfsdk:"aws" json:"aws,omitempty"`
			} `tfsdk:"from_groups" json:"fromGroups,omitempty"`
			FromNodes *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_nodes" json:"fromNodes,omitempty"`
			FromRequires *[]struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"from_requires" json:"fromRequires,omitempty"`
			Icmps *[]struct {
				Fields *[]struct {
					Family *string `tfsdk:"family" json:"family,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
			} `tfsdk:"icmps" json:"icmps,omitempty"`
			ToPorts *[]struct {
				Ports *[]struct {
					Port     *string `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"to_ports" json:"toPorts,omitempty"`
		} `tfsdk:"ingress_deny" json:"ingressDeny,omitempty"`
		Labels *[]struct {
			Key    *string `tfsdk:"key" json:"key,omitempty"`
			Source *string `tfsdk:"source" json:"source,omitempty"`
			Value  *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"labels" json:"labels,omitempty"`
		NodeSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
	} `tfsdk:"specs" json:"specs,omitempty"`
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cilium_io_cilium_clusterwide_network_policy_v2_manifest"
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CiliumClusterwideNetworkPolicy is a Kubernetes third-party resource with an modified version of CiliumNetworkPolicy which is cluster scoped rather than namespace scoped.",
		MarkdownDescription: "CiliumClusterwideNetworkPolicy is a Kubernetes third-party resource with an modified version of CiliumNetworkPolicy which is cluster scoped rather than namespace scoped.",
		Attributes: map[string]schema.Attribute{
			"yaml": schema.StringAttribute{
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"metadata": schema.SingleNestedAttribute{
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Optional:            false,
				Computed:            false,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							validators.NameValidator(),
							stringvalidator.LengthAtLeast(1),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.LabelValidator(),
						},
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "Spec is the desired Cilium specific rule specification.",
				MarkdownDescription: "Spec is the desired Cilium specific rule specification.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "Description is a free form string, it can be used by the creator of the rule to store human readable explanation of the purpose of this rule. Rules cannot be identified by comment.",
						MarkdownDescription: "Description is a free form string, it can be used by the creator of the rule to store human readable explanation of the purpose of this rule. Rules cannot be identified by comment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"egress": schema.ListNestedAttribute{
						Description:         "Egress is a list of EgressRule which are enforced at egress. If omitted or empty, this rule does not apply at egress.",
						MarkdownDescription: "Egress is a list of EgressRule which are enforced at egress. If omitted or empty, this rule does not apply at egress.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"authentication": schema.SingleNestedAttribute{
									Description:         "Authentication is the required authentication type for the allowed traffic, if any.",
									MarkdownDescription: "Authentication is the required authentication type for the allowed traffic, if any.",
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description:         "Mode is the required authentication mode for the allowed traffic, if any.",
											MarkdownDescription: "Mode is the required authentication mode for the allowed traffic, if any.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("disabled", "required", "test-always-fail"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"icmps": schema.ListNestedAttribute{
									Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is allowed to initiate type 8 ICMP connections.",
									MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is allowed to initiate type 8 ICMP connections.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"fields": schema.ListNestedAttribute{
												Description:         "Fields is a list of ICMP fields.",
												MarkdownDescription: "Fields is a list of ICMP fields.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"family": schema.StringAttribute{
															Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
															MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("IPv4", "IPv6"),
															},
														},

														"type": schema.StringAttribute{
															Description:         "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
															MarkdownDescription: "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_cidr": schema.ListAttribute{
									Description:         "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
									MarkdownDescription: "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"to_cidr_set": schema.ListNestedAttribute{
									Description:         "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
									MarkdownDescription: "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cidr": schema.StringAttribute{
												Description:         "CIDR is a CIDR prefix / IP Block.",
												MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
												},
											},

											"cidr_group_ref": schema.StringAttribute{
												Description:         "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
												MarkdownDescription: "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"except": schema.ListAttribute{
												Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
												MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_endpoints": schema.ListNestedAttribute{
									Description:         "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
									MarkdownDescription: "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_entities": schema.ListAttribute{
									Description:         "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster','host','remote-node','kube-apiserver', 'init', 'health','unmanaged' and 'all'.",
									MarkdownDescription: "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster','host','remote-node','kube-apiserver', 'init', 'health','unmanaged' and 'all'.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"to_fqd_ns": schema.ListNestedAttribute{
									Description:         "ToFQDN allows whitelisting DNS names in place of IPs. The IPs that result from DNS resolution of 'ToFQDN.MatchName's are added to the same EgressRule object as ToCIDRSet entries, and behave accordingly. Any L4 and L7 rules within this EgressRule will also apply to these IPs. The DNS -> IP mapping is re-resolved periodically from within the cilium-agent, and the IPs in the DNS response are effected in the policy for selected pods as-is (i.e. the list of IPs is not modified in any way). Note: An explicit rule to allow for DNS traffic is needed for the pods, as ToFQDN counts as an egress rule and will enforce egress policy when PolicyEnforcment=default. Note: If the resolved IPs are IPs within the kubernetes cluster, the ToFQDN rule will not apply to that IP. Note: ToFQDN cannot occur in the same policy as other To* rules.",
									MarkdownDescription: "ToFQDN allows whitelisting DNS names in place of IPs. The IPs that result from DNS resolution of 'ToFQDN.MatchName's are added to the same EgressRule object as ToCIDRSet entries, and behave accordingly. Any L4 and L7 rules within this EgressRule will also apply to these IPs. The DNS -> IP mapping is re-resolved periodically from within the cilium-agent, and the IPs in the DNS response are effected in the policy for selected pods as-is (i.e. the list of IPs is not modified in any way). Note: An explicit rule to allow for DNS traffic is needed for the pods, as ToFQDN counts as an egress rule and will enforce egress policy when PolicyEnforcment=default. Note: If the resolved IPs are IPs within the kubernetes cluster, the ToFQDN rule will not apply to that IP. Note: ToFQDN cannot occur in the same policy as other To* rules.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_name": schema.StringAttribute{
												Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
												MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
												},
											},

											"match_pattern": schema.StringAttribute{
												Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
												MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_groups": schema.ListNestedAttribute{
									Description:         "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
									MarkdownDescription: "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"aws": schema.SingleNestedAttribute{
												Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
												MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
												Attributes: map[string]schema.Attribute{
													"labels": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups_ids": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups_names": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_nodes": schema.ListNestedAttribute{
									Description:         "ToNodes is a list of nodes identified by an EndpointSelector to which endpoints subject to the rule is allowed to communicate.",
									MarkdownDescription: "ToNodes is a list of nodes identified by an EndpointSelector to which endpoints subject to the rule is allowed to communicate.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_ports": schema.ListNestedAttribute{
									Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is allowed to initiate connections to destination port 8080/tcp",
									MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is allowed to initiate connections to destination port 8080/tcp",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"listener": schema.SingleNestedAttribute{
												Description:         "listener specifies the name of a custom Envoy listener to which this traffic should be redirected to.",
												MarkdownDescription: "listener specifies the name of a custom Envoy listener to which this traffic should be redirected to.",
												Attributes: map[string]schema.Attribute{
													"envoy_config": schema.SingleNestedAttribute{
														Description:         "EnvoyConfig is a reference to the CEC or CCEC resource in which the listener is defined.",
														MarkdownDescription: "EnvoyConfig is a reference to the CEC or CCEC resource in which the listener is defined.",
														Attributes: map[string]schema.Attribute{
															"kind": schema.StringAttribute{
																Description:         "Kind is the resource type being referred to. Defaults to CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig for CiliumNetworkPolicy and CiliumClusterwideNetworkPolicy, respectively. The only case this is currently explicitly needed is when referring to a CiliumClusterwideEnvoyConfig from CiliumNetworkPolicy, as using a namespaced listener from a cluster scoped policy is not allowed.",
																MarkdownDescription: "Kind is the resource type being referred to. Defaults to CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig for CiliumNetworkPolicy and CiliumClusterwideNetworkPolicy, respectively. The only case this is currently explicitly needed is when referring to a CiliumClusterwideEnvoyConfig from CiliumNetworkPolicy, as using a namespaced listener from a cluster scoped policy is not allowed.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("CiliumEnvoyConfig", "CiliumClusterwideEnvoyConfig"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name is the resource name of the CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig where the listener is defined in.",
																MarkdownDescription: "Name is the resource name of the CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig where the listener is defined in.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the listener.",
														MarkdownDescription: "Name is the name of the listener.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"priority": schema.Int64Attribute{
														Description:         "Priority for this Listener that is used when multiple rules would apply different listeners to a policy map entry. Behavior of this is implementation dependent.",
														MarkdownDescription: "Priority for this Listener that is used when multiple rules would apply different listeners to a policy map entry. Behavior of this is implementation dependent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(100),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"originating_tls": schema.SingleNestedAttribute{
												Description:         "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
												MarkdownDescription: "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
												Attributes: map[string]schema.Attribute{
													"certificate": schema.StringAttribute{
														Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"private_key": schema.StringAttribute{
														Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret": schema.SingleNestedAttribute{
														Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
														MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of the secret.",
																MarkdownDescription: "Name is the name of the secret.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"trusted_ca": schema.StringAttribute{
														Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"ports": schema.ListNestedAttribute{
												Description:         "Ports is a list of L4 port/protocol",
												MarkdownDescription: "Ports is a list of L4 port/protocol",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"port": schema.StringAttribute{
															Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
															MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
															},
														},

														"protocol": schema.StringAttribute{
															Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
															MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"rules": schema.SingleNestedAttribute{
												Description:         "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
												MarkdownDescription: "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
												Attributes: map[string]schema.Attribute{
													"dns": schema.ListNestedAttribute{
														Description:         "DNS-specific rules.",
														MarkdownDescription: "DNS-specific rules.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match_name": schema.StringAttribute{
																	Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
																	MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
																	},
																},

																"match_pattern": schema.StringAttribute{
																	Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
																	MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"http": schema.ListNestedAttribute{
														Description:         "HTTP specific rules.",
														MarkdownDescription: "HTTP specific rules.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"header_matches": schema.ListNestedAttribute{
																	Description:         "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
																	MarkdownDescription: "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mismatch": schema.StringAttribute{
																				Description:         "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																				MarkdownDescription: "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("LOG", "ADD", "DELETE", "REPLACE"),
																				},
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name identifies the header.",
																				MarkdownDescription: "Name identifies the header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"secret": schema.SingleNestedAttribute{
																				Description:         "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																				MarkdownDescription: "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name is the name of the secret.",
																						MarkdownDescription: "Name is the name of the secret.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"namespace": schema.StringAttribute{
																						Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																						MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																				MarkdownDescription: "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"headers": schema.ListAttribute{
																	Description:         "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
																	MarkdownDescription: "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"host": schema.StringAttribute{
																	Description:         "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
																	MarkdownDescription: "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"method": schema.StringAttribute{
																	Description:         "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
																	MarkdownDescription: "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
																	MarkdownDescription: "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"kafka": schema.ListNestedAttribute{
														Description:         "Kafka-specific rules.",
														MarkdownDescription: "Kafka-specific rules.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_key": schema.StringAttribute{
																	Description:         "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
																	MarkdownDescription: "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"api_version": schema.StringAttribute{
																	Description:         "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
																	MarkdownDescription: "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"client_id": schema.StringAttribute{
																	Description:         "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
																	MarkdownDescription: "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"role": schema.StringAttribute{
																	Description:         "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported: - 'produce': Allow producing to the topics specified in the rule - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
																	MarkdownDescription: "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported: - 'produce': Allow producing to the topics specified in the rule - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("produce", "consume"),
																	},
																},

																"topic": schema.StringAttribute{
																	Description:         "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
																	MarkdownDescription: "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtMost(255),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"l7": schema.ListAttribute{
														Description:         "Key-value pair rules.",
														MarkdownDescription: "Key-value pair rules.",
														ElementType:         types.MapType{ElemType: types.StringType},
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"l7proto": schema.StringAttribute{
														Description:         "Name of the L7 protocol for which the Key-value pair rules apply.",
														MarkdownDescription: "Name of the L7 protocol for which the Key-value pair rules apply.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_names": schema.ListAttribute{
												Description:         "ServerNames is a list of allowed TLS SNI values. If not empty, then TLS must be present and one of the provided SNIs must be indicated in the TLS handshake.",
												MarkdownDescription: "ServerNames is a list of allowed TLS SNI values. If not empty, then TLS must be present and one of the provided SNIs must be indicated in the TLS handshake.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"terminating_tls": schema.SingleNestedAttribute{
												Description:         "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
												MarkdownDescription: "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
												Attributes: map[string]schema.Attribute{
													"certificate": schema.StringAttribute{
														Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"private_key": schema.StringAttribute{
														Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret": schema.SingleNestedAttribute{
														Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
														MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of the secret.",
																MarkdownDescription: "Name is the name of the secret.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"trusted_ca": schema.StringAttribute{
														Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_requires": schema.ListNestedAttribute{
									Description:         "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
									MarkdownDescription: "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_services": schema.ListNestedAttribute{
									Description:         "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
									MarkdownDescription: "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"k8s_service": schema.SingleNestedAttribute{
												Description:         "K8sService selects service by name and namespace pair",
												MarkdownDescription: "K8sService selects service by name and namespace pair",
												Attributes: map[string]schema.Attribute{
													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"k8s_service_selector": schema.SingleNestedAttribute{
												Description:         "K8sServiceSelector selects services by k8s labels and namespace",
												MarkdownDescription: "K8sServiceSelector selects services by k8s labels and namespace",
												Attributes: map[string]schema.Attribute{
													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"selector": schema.SingleNestedAttribute{
														Description:         "ServiceSelector is a label selector for k8s services",
														MarkdownDescription: "ServiceSelector is a label selector for k8s services",
														Attributes: map[string]schema.Attribute{
															"match_expressions": schema.ListNestedAttribute{
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "key is the label key that the selector applies to.",
																			MarkdownDescription: "key is the label key that the selector applies to.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"operator": schema.StringAttribute{
																			Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																			MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																			},
																		},

																		"values": schema.ListAttribute{
																			Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																			MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": schema.MapAttribute{
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"egress_deny": schema.ListNestedAttribute{
						Description:         "EgressDeny is a list of EgressDenyRule which are enforced at egress. Any rule inserted here will be denied regardless of the allowed egress rules in the 'egress' field. If omitted or empty, this rule does not apply at egress.",
						MarkdownDescription: "EgressDeny is a list of EgressDenyRule which are enforced at egress. Any rule inserted here will be denied regardless of the allowed egress rules in the 'egress' field. If omitted or empty, this rule does not apply at egress.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"icmps": schema.ListNestedAttribute{
									Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is not allowed to initiate type 8 ICMP connections.",
									MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is not allowed to initiate type 8 ICMP connections.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"fields": schema.ListNestedAttribute{
												Description:         "Fields is a list of ICMP fields.",
												MarkdownDescription: "Fields is a list of ICMP fields.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"family": schema.StringAttribute{
															Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
															MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("IPv4", "IPv6"),
															},
														},

														"type": schema.StringAttribute{
															Description:         "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
															MarkdownDescription: "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_cidr": schema.ListAttribute{
									Description:         "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
									MarkdownDescription: "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"to_cidr_set": schema.ListNestedAttribute{
									Description:         "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
									MarkdownDescription: "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cidr": schema.StringAttribute{
												Description:         "CIDR is a CIDR prefix / IP Block.",
												MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
												},
											},

											"cidr_group_ref": schema.StringAttribute{
												Description:         "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
												MarkdownDescription: "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"except": schema.ListAttribute{
												Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
												MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_endpoints": schema.ListNestedAttribute{
									Description:         "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
									MarkdownDescription: "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_entities": schema.ListAttribute{
									Description:         "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster','host','remote-node','kube-apiserver', 'init', 'health','unmanaged' and 'all'.",
									MarkdownDescription: "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster','host','remote-node','kube-apiserver', 'init', 'health','unmanaged' and 'all'.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"to_groups": schema.ListNestedAttribute{
									Description:         "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
									MarkdownDescription: "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"aws": schema.SingleNestedAttribute{
												Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
												MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
												Attributes: map[string]schema.Attribute{
													"labels": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups_ids": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups_names": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_nodes": schema.ListNestedAttribute{
									Description:         "ToNodes is a list of nodes identified by an EndpointSelector to which endpoints subject to the rule is allowed to communicate.",
									MarkdownDescription: "ToNodes is a list of nodes identified by an EndpointSelector to which endpoints subject to the rule is allowed to communicate.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_ports": schema.ListNestedAttribute{
									Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is not allowed to initiate connections to destination port 8080/tcp",
									MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is not allowed to initiate connections to destination port 8080/tcp",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ports": schema.ListNestedAttribute{
												Description:         "Ports is a list of L4 port/protocol",
												MarkdownDescription: "Ports is a list of L4 port/protocol",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"port": schema.StringAttribute{
															Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
															MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
															},
														},

														"protocol": schema.StringAttribute{
															Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
															MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_requires": schema.ListNestedAttribute{
									Description:         "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
									MarkdownDescription: "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_services": schema.ListNestedAttribute{
									Description:         "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
									MarkdownDescription: "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"k8s_service": schema.SingleNestedAttribute{
												Description:         "K8sService selects service by name and namespace pair",
												MarkdownDescription: "K8sService selects service by name and namespace pair",
												Attributes: map[string]schema.Attribute{
													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"k8s_service_selector": schema.SingleNestedAttribute{
												Description:         "K8sServiceSelector selects services by k8s labels and namespace",
												MarkdownDescription: "K8sServiceSelector selects services by k8s labels and namespace",
												Attributes: map[string]schema.Attribute{
													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"selector": schema.SingleNestedAttribute{
														Description:         "ServiceSelector is a label selector for k8s services",
														MarkdownDescription: "ServiceSelector is a label selector for k8s services",
														Attributes: map[string]schema.Attribute{
															"match_expressions": schema.ListNestedAttribute{
																Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "key is the label key that the selector applies to.",
																			MarkdownDescription: "key is the label key that the selector applies to.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"operator": schema.StringAttribute{
																			Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																			MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																			},
																		},

																		"values": schema.ListAttribute{
																			Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																			MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"match_labels": schema.MapAttribute{
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_default_deny": schema.SingleNestedAttribute{
						Description:         "EnableDefaultDeny determines whether this policy configures the subject endpoint(s) to have a default deny mode. If enabled, this causes all traffic not explicitly allowed by a network policy to be dropped.  If not specified, the default is true for each traffic direction that has rules, and false otherwise. For example, if a policy only has Ingress or IngressDeny rules, then the default for ingress is true and egress is false.  If multiple policies apply to an endpoint, that endpoint's default deny will be enabled if any policy requests it.  This is useful for creating broad-based network policies that will not cause endpoints to enter default-deny mode.",
						MarkdownDescription: "EnableDefaultDeny determines whether this policy configures the subject endpoint(s) to have a default deny mode. If enabled, this causes all traffic not explicitly allowed by a network policy to be dropped.  If not specified, the default is true for each traffic direction that has rules, and false otherwise. For example, if a policy only has Ingress or IngressDeny rules, then the default for ingress is true and egress is false.  If multiple policies apply to an endpoint, that endpoint's default deny will be enabled if any policy requests it.  This is useful for creating broad-based network policies that will not cause endpoints to enter default-deny mode.",
						Attributes: map[string]schema.Attribute{
							"egress": schema.BoolAttribute{
								Description:         "Whether or not the endpoint should have a default-deny rule applied to egress traffic.",
								MarkdownDescription: "Whether or not the endpoint should have a default-deny rule applied to egress traffic.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress": schema.BoolAttribute{
								Description:         "Whether or not the endpoint should have a default-deny rule applied to ingress traffic.",
								MarkdownDescription: "Whether or not the endpoint should have a default-deny rule applied to ingress traffic.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"endpoint_selector": schema.SingleNestedAttribute{
						Description:         "EndpointSelector selects all endpoints which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive.",
						MarkdownDescription: "EndpointSelector selects all endpoints which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
											},
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": schema.ListNestedAttribute{
						Description:         "Ingress is a list of IngressRule which are enforced at ingress. If omitted or empty, this rule does not apply at ingress.",
						MarkdownDescription: "Ingress is a list of IngressRule which are enforced at ingress. If omitted or empty, this rule does not apply at ingress.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"authentication": schema.SingleNestedAttribute{
									Description:         "Authentication is the required authentication type for the allowed traffic, if any.",
									MarkdownDescription: "Authentication is the required authentication type for the allowed traffic, if any.",
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description:         "Mode is the required authentication mode for the allowed traffic, if any.",
											MarkdownDescription: "Mode is the required authentication mode for the allowed traffic, if any.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("disabled", "required", "test-always-fail"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_cidr": schema.ListAttribute{
									Description:         "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
									MarkdownDescription: "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"from_cidr_set": schema.ListNestedAttribute{
									Description:         "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
									MarkdownDescription: "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cidr": schema.StringAttribute{
												Description:         "CIDR is a CIDR prefix / IP Block.",
												MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
												},
											},

											"cidr_group_ref": schema.StringAttribute{
												Description:         "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
												MarkdownDescription: "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"except": schema.ListAttribute{
												Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
												MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_endpoints": schema.ListNestedAttribute{
									Description:         "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
									MarkdownDescription: "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_entities": schema.ListAttribute{
									Description:         "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
									MarkdownDescription: "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"from_groups": schema.ListNestedAttribute{
									Description:         "FromGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: FromGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
									MarkdownDescription: "FromGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: FromGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"aws": schema.SingleNestedAttribute{
												Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
												MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
												Attributes: map[string]schema.Attribute{
													"labels": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups_ids": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups_names": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_nodes": schema.ListNestedAttribute{
									Description:         "FromNodes is a list of nodes identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.",
									MarkdownDescription: "FromNodes is a list of nodes identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_requires": schema.ListNestedAttribute{
									Description:         "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
									MarkdownDescription: "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"icmps": schema.ListNestedAttribute{
									Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming type 8 ICMP connections.",
									MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming type 8 ICMP connections.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"fields": schema.ListNestedAttribute{
												Description:         "Fields is a list of ICMP fields.",
												MarkdownDescription: "Fields is a list of ICMP fields.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"family": schema.StringAttribute{
															Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
															MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("IPv4", "IPv6"),
															},
														},

														"type": schema.StringAttribute{
															Description:         "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
															MarkdownDescription: "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_ports": schema.ListNestedAttribute{
									Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming connections on port 80/tcp.",
									MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming connections on port 80/tcp.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"listener": schema.SingleNestedAttribute{
												Description:         "listener specifies the name of a custom Envoy listener to which this traffic should be redirected to.",
												MarkdownDescription: "listener specifies the name of a custom Envoy listener to which this traffic should be redirected to.",
												Attributes: map[string]schema.Attribute{
													"envoy_config": schema.SingleNestedAttribute{
														Description:         "EnvoyConfig is a reference to the CEC or CCEC resource in which the listener is defined.",
														MarkdownDescription: "EnvoyConfig is a reference to the CEC or CCEC resource in which the listener is defined.",
														Attributes: map[string]schema.Attribute{
															"kind": schema.StringAttribute{
																Description:         "Kind is the resource type being referred to. Defaults to CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig for CiliumNetworkPolicy and CiliumClusterwideNetworkPolicy, respectively. The only case this is currently explicitly needed is when referring to a CiliumClusterwideEnvoyConfig from CiliumNetworkPolicy, as using a namespaced listener from a cluster scoped policy is not allowed.",
																MarkdownDescription: "Kind is the resource type being referred to. Defaults to CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig for CiliumNetworkPolicy and CiliumClusterwideNetworkPolicy, respectively. The only case this is currently explicitly needed is when referring to a CiliumClusterwideEnvoyConfig from CiliumNetworkPolicy, as using a namespaced listener from a cluster scoped policy is not allowed.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("CiliumEnvoyConfig", "CiliumClusterwideEnvoyConfig"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name is the resource name of the CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig where the listener is defined in.",
																MarkdownDescription: "Name is the resource name of the CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig where the listener is defined in.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the listener.",
														MarkdownDescription: "Name is the name of the listener.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"priority": schema.Int64Attribute{
														Description:         "Priority for this Listener that is used when multiple rules would apply different listeners to a policy map entry. Behavior of this is implementation dependent.",
														MarkdownDescription: "Priority for this Listener that is used when multiple rules would apply different listeners to a policy map entry. Behavior of this is implementation dependent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(100),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"originating_tls": schema.SingleNestedAttribute{
												Description:         "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
												MarkdownDescription: "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
												Attributes: map[string]schema.Attribute{
													"certificate": schema.StringAttribute{
														Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"private_key": schema.StringAttribute{
														Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret": schema.SingleNestedAttribute{
														Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
														MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of the secret.",
																MarkdownDescription: "Name is the name of the secret.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"trusted_ca": schema.StringAttribute{
														Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"ports": schema.ListNestedAttribute{
												Description:         "Ports is a list of L4 port/protocol",
												MarkdownDescription: "Ports is a list of L4 port/protocol",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"port": schema.StringAttribute{
															Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
															MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
															},
														},

														"protocol": schema.StringAttribute{
															Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
															MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"rules": schema.SingleNestedAttribute{
												Description:         "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
												MarkdownDescription: "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
												Attributes: map[string]schema.Attribute{
													"dns": schema.ListNestedAttribute{
														Description:         "DNS-specific rules.",
														MarkdownDescription: "DNS-specific rules.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match_name": schema.StringAttribute{
																	Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
																	MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
																	},
																},

																"match_pattern": schema.StringAttribute{
																	Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
																	MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"http": schema.ListNestedAttribute{
														Description:         "HTTP specific rules.",
														MarkdownDescription: "HTTP specific rules.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"header_matches": schema.ListNestedAttribute{
																	Description:         "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
																	MarkdownDescription: "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"mismatch": schema.StringAttribute{
																				Description:         "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																				MarkdownDescription: "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("LOG", "ADD", "DELETE", "REPLACE"),
																				},
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name identifies the header.",
																				MarkdownDescription: "Name identifies the header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"secret": schema.SingleNestedAttribute{
																				Description:         "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																				MarkdownDescription: "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name is the name of the secret.",
																						MarkdownDescription: "Name is the name of the secret.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"namespace": schema.StringAttribute{
																						Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																						MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																				MarkdownDescription: "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"headers": schema.ListAttribute{
																	Description:         "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
																	MarkdownDescription: "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"host": schema.StringAttribute{
																	Description:         "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
																	MarkdownDescription: "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"method": schema.StringAttribute{
																	Description:         "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
																	MarkdownDescription: "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
																	MarkdownDescription: "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"kafka": schema.ListNestedAttribute{
														Description:         "Kafka-specific rules.",
														MarkdownDescription: "Kafka-specific rules.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"api_key": schema.StringAttribute{
																	Description:         "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
																	MarkdownDescription: "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"api_version": schema.StringAttribute{
																	Description:         "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
																	MarkdownDescription: "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"client_id": schema.StringAttribute{
																	Description:         "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
																	MarkdownDescription: "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"role": schema.StringAttribute{
																	Description:         "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported: - 'produce': Allow producing to the topics specified in the rule - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
																	MarkdownDescription: "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported: - 'produce': Allow producing to the topics specified in the rule - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("produce", "consume"),
																	},
																},

																"topic": schema.StringAttribute{
																	Description:         "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
																	MarkdownDescription: "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtMost(255),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"l7": schema.ListAttribute{
														Description:         "Key-value pair rules.",
														MarkdownDescription: "Key-value pair rules.",
														ElementType:         types.MapType{ElemType: types.StringType},
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"l7proto": schema.StringAttribute{
														Description:         "Name of the L7 protocol for which the Key-value pair rules apply.",
														MarkdownDescription: "Name of the L7 protocol for which the Key-value pair rules apply.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"server_names": schema.ListAttribute{
												Description:         "ServerNames is a list of allowed TLS SNI values. If not empty, then TLS must be present and one of the provided SNIs must be indicated in the TLS handshake.",
												MarkdownDescription: "ServerNames is a list of allowed TLS SNI values. If not empty, then TLS must be present and one of the provided SNIs must be indicated in the TLS handshake.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"terminating_tls": schema.SingleNestedAttribute{
												Description:         "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
												MarkdownDescription: "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
												Attributes: map[string]schema.Attribute{
													"certificate": schema.StringAttribute{
														Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"private_key": schema.StringAttribute{
														Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret": schema.SingleNestedAttribute{
														Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
														MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name is the name of the secret.",
																MarkdownDescription: "Name is the name of the secret.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"trusted_ca": schema.StringAttribute{
														Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
														MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress_deny": schema.ListNestedAttribute{
						Description:         "IngressDeny is a list of IngressDenyRule which are enforced at ingress. Any rule inserted here will be denied regardless of the allowed ingress rules in the 'ingress' field. If omitted or empty, this rule does not apply at ingress.",
						MarkdownDescription: "IngressDeny is a list of IngressDenyRule which are enforced at ingress. Any rule inserted here will be denied regardless of the allowed ingress rules in the 'ingress' field. If omitted or empty, this rule does not apply at ingress.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"from_cidr": schema.ListAttribute{
									Description:         "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
									MarkdownDescription: "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"from_cidr_set": schema.ListNestedAttribute{
									Description:         "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
									MarkdownDescription: "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cidr": schema.StringAttribute{
												Description:         "CIDR is a CIDR prefix / IP Block.",
												MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
												},
											},

											"cidr_group_ref": schema.StringAttribute{
												Description:         "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
												MarkdownDescription: "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
											},

											"except": schema.ListAttribute{
												Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
												MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_endpoints": schema.ListNestedAttribute{
									Description:         "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
									MarkdownDescription: "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_entities": schema.ListAttribute{
									Description:         "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
									MarkdownDescription: "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"from_groups": schema.ListNestedAttribute{
									Description:         "FromGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: FromGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
									MarkdownDescription: "FromGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: FromGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"aws": schema.SingleNestedAttribute{
												Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
												MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
												Attributes: map[string]schema.Attribute{
													"labels": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups_ids": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"security_groups_names": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_nodes": schema.ListNestedAttribute{
									Description:         "FromNodes is a list of nodes identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.",
									MarkdownDescription: "FromNodes is a list of nodes identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"from_requires": schema.ListNestedAttribute{
									Description:         "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
									MarkdownDescription: "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
															},
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_labels": schema.MapAttribute{
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"icmps": schema.ListNestedAttribute{
									Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming type 8 ICMP connections.",
									MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming type 8 ICMP connections.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"fields": schema.ListNestedAttribute{
												Description:         "Fields is a list of ICMP fields.",
												MarkdownDescription: "Fields is a list of ICMP fields.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"family": schema.StringAttribute{
															Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
															MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("IPv4", "IPv6"),
															},
														},

														"type": schema.StringAttribute{
															Description:         "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
															MarkdownDescription: "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to_ports": schema.ListNestedAttribute{
									Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming connections on port 80/tcp.",
									MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming connections on port 80/tcp.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ports": schema.ListNestedAttribute{
												Description:         "Ports is a list of L4 port/protocol",
												MarkdownDescription: "Ports is a list of L4 port/protocol",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"port": schema.StringAttribute{
															Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
															MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
															},
														},

														"protocol": schema.StringAttribute{
															Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
															MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"labels": schema.ListNestedAttribute{
						Description:         "Labels is a list of optional strings which can be used to re-identify the rule or to store metadata. It is possible to lookup or delete strings based on labels. Labels are not required to be unique, multiple rules can have overlapping or identical labels.",
						MarkdownDescription: "Labels is a list of optional strings which can be used to re-identify the rule or to store metadata. It is possible to lookup or delete strings based on labels. Labels are not required to be unique, multiple rules can have overlapping or identical labels.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"source": schema.StringAttribute{
									Description:         "Source can be one of the above values (e.g.: LabelSourceContainer).",
									MarkdownDescription: "Source can be one of the above values (e.g.: LabelSourceContainer).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": schema.SingleNestedAttribute{
						Description:         "NodeSelector selects all nodes which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive. Can only be used in CiliumClusterwideNetworkPolicies.",
						MarkdownDescription: "NodeSelector selects all nodes which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive. Can only be used in CiliumClusterwideNetworkPolicies.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
											},
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},

			"specs": schema.ListNestedAttribute{
				Description:         "Specs is a list of desired Cilium specific rule specification.",
				MarkdownDescription: "Specs is a list of desired Cilium specific rule specification.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							Description:         "Description is a free form string, it can be used by the creator of the rule to store human readable explanation of the purpose of this rule. Rules cannot be identified by comment.",
							MarkdownDescription: "Description is a free form string, it can be used by the creator of the rule to store human readable explanation of the purpose of this rule. Rules cannot be identified by comment.",
							Required:            false,
							Optional:            true,
							Computed:            false,
						},

						"egress": schema.ListNestedAttribute{
							Description:         "Egress is a list of EgressRule which are enforced at egress. If omitted or empty, this rule does not apply at egress.",
							MarkdownDescription: "Egress is a list of EgressRule which are enforced at egress. If omitted or empty, this rule does not apply at egress.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"authentication": schema.SingleNestedAttribute{
										Description:         "Authentication is the required authentication type for the allowed traffic, if any.",
										MarkdownDescription: "Authentication is the required authentication type for the allowed traffic, if any.",
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Description:         "Mode is the required authentication mode for the allowed traffic, if any.",
												MarkdownDescription: "Mode is the required authentication mode for the allowed traffic, if any.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("disabled", "required", "test-always-fail"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"icmps": schema.ListNestedAttribute{
										Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is allowed to initiate type 8 ICMP connections.",
										MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is allowed to initiate type 8 ICMP connections.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"fields": schema.ListNestedAttribute{
													Description:         "Fields is a list of ICMP fields.",
													MarkdownDescription: "Fields is a list of ICMP fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"family": schema.StringAttribute{
																Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
																MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("IPv4", "IPv6"),
																},
															},

															"type": schema.StringAttribute{
																Description:         "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
																MarkdownDescription: "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_cidr": schema.ListAttribute{
										Description:         "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
										MarkdownDescription: "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"to_cidr_set": schema.ListNestedAttribute{
										Description:         "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
										MarkdownDescription: "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Description:         "CIDR is a CIDR prefix / IP Block.",
													MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
													},
												},

												"cidr_group_ref": schema.StringAttribute{
													Description:         "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
													MarkdownDescription: "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(253),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
													},
												},

												"except": schema.ListAttribute{
													Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
													MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_endpoints": schema.ListNestedAttribute{
										Description:         "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
										MarkdownDescription: "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_entities": schema.ListAttribute{
										Description:         "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster','host','remote-node','kube-apiserver', 'init', 'health','unmanaged' and 'all'.",
										MarkdownDescription: "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster','host','remote-node','kube-apiserver', 'init', 'health','unmanaged' and 'all'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"to_fqd_ns": schema.ListNestedAttribute{
										Description:         "ToFQDN allows whitelisting DNS names in place of IPs. The IPs that result from DNS resolution of 'ToFQDN.MatchName's are added to the same EgressRule object as ToCIDRSet entries, and behave accordingly. Any L4 and L7 rules within this EgressRule will also apply to these IPs. The DNS -> IP mapping is re-resolved periodically from within the cilium-agent, and the IPs in the DNS response are effected in the policy for selected pods as-is (i.e. the list of IPs is not modified in any way). Note: An explicit rule to allow for DNS traffic is needed for the pods, as ToFQDN counts as an egress rule and will enforce egress policy when PolicyEnforcment=default. Note: If the resolved IPs are IPs within the kubernetes cluster, the ToFQDN rule will not apply to that IP. Note: ToFQDN cannot occur in the same policy as other To* rules.",
										MarkdownDescription: "ToFQDN allows whitelisting DNS names in place of IPs. The IPs that result from DNS resolution of 'ToFQDN.MatchName's are added to the same EgressRule object as ToCIDRSet entries, and behave accordingly. Any L4 and L7 rules within this EgressRule will also apply to these IPs. The DNS -> IP mapping is re-resolved periodically from within the cilium-agent, and the IPs in the DNS response are effected in the policy for selected pods as-is (i.e. the list of IPs is not modified in any way). Note: An explicit rule to allow for DNS traffic is needed for the pods, as ToFQDN counts as an egress rule and will enforce egress policy when PolicyEnforcment=default. Note: If the resolved IPs are IPs within the kubernetes cluster, the ToFQDN rule will not apply to that IP. Note: ToFQDN cannot occur in the same policy as other To* rules.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_name": schema.StringAttribute{
													Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
													MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
													},
												},

												"match_pattern": schema.StringAttribute{
													Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
													MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_groups": schema.ListNestedAttribute{
										Description:         "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
										MarkdownDescription: "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"aws": schema.SingleNestedAttribute{
													Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
													MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
													Attributes: map[string]schema.Attribute{
														"labels": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"region": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_groups_ids": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_groups_names": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_nodes": schema.ListNestedAttribute{
										Description:         "ToNodes is a list of nodes identified by an EndpointSelector to which endpoints subject to the rule is allowed to communicate.",
										MarkdownDescription: "ToNodes is a list of nodes identified by an EndpointSelector to which endpoints subject to the rule is allowed to communicate.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_ports": schema.ListNestedAttribute{
										Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is allowed to initiate connections to destination port 8080/tcp",
										MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is allowed to initiate connections to destination port 8080/tcp",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"listener": schema.SingleNestedAttribute{
													Description:         "listener specifies the name of a custom Envoy listener to which this traffic should be redirected to.",
													MarkdownDescription: "listener specifies the name of a custom Envoy listener to which this traffic should be redirected to.",
													Attributes: map[string]schema.Attribute{
														"envoy_config": schema.SingleNestedAttribute{
															Description:         "EnvoyConfig is a reference to the CEC or CCEC resource in which the listener is defined.",
															MarkdownDescription: "EnvoyConfig is a reference to the CEC or CCEC resource in which the listener is defined.",
															Attributes: map[string]schema.Attribute{
																"kind": schema.StringAttribute{
																	Description:         "Kind is the resource type being referred to. Defaults to CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig for CiliumNetworkPolicy and CiliumClusterwideNetworkPolicy, respectively. The only case this is currently explicitly needed is when referring to a CiliumClusterwideEnvoyConfig from CiliumNetworkPolicy, as using a namespaced listener from a cluster scoped policy is not allowed.",
																	MarkdownDescription: "Kind is the resource type being referred to. Defaults to CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig for CiliumNetworkPolicy and CiliumClusterwideNetworkPolicy, respectively. The only case this is currently explicitly needed is when referring to a CiliumClusterwideEnvoyConfig from CiliumNetworkPolicy, as using a namespaced listener from a cluster scoped policy is not allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("CiliumEnvoyConfig", "CiliumClusterwideEnvoyConfig"),
																	},
																},

																"name": schema.StringAttribute{
																	Description:         "Name is the resource name of the CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig where the listener is defined in.",
																	MarkdownDescription: "Name is the resource name of the CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig where the listener is defined in.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of the listener.",
															MarkdownDescription: "Name is the name of the listener.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"priority": schema.Int64Attribute{
															Description:         "Priority for this Listener that is used when multiple rules would apply different listeners to a policy map entry. Behavior of this is implementation dependent.",
															MarkdownDescription: "Priority for this Listener that is used when multiple rules would apply different listeners to a policy map entry. Behavior of this is implementation dependent.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(100),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"originating_tls": schema.SingleNestedAttribute{
													Description:         "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
													MarkdownDescription: "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
													Attributes: map[string]schema.Attribute{
														"certificate": schema.StringAttribute{
															Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"private_key": schema.StringAttribute{
															Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
															MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the secret.",
																	MarkdownDescription: "Name is the name of the secret.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																	MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"trusted_ca": schema.StringAttribute{
															Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"ports": schema.ListNestedAttribute{
													Description:         "Ports is a list of L4 port/protocol",
													MarkdownDescription: "Ports is a list of L4 port/protocol",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"port": schema.StringAttribute{
																Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
																MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
																},
															},

															"protocol": schema.StringAttribute{
																Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
																MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"rules": schema.SingleNestedAttribute{
													Description:         "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
													MarkdownDescription: "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
													Attributes: map[string]schema.Attribute{
														"dns": schema.ListNestedAttribute{
															Description:         "DNS-specific rules.",
															MarkdownDescription: "DNS-specific rules.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"match_name": schema.StringAttribute{
																		Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
																		MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
																		},
																	},

																	"match_pattern": schema.StringAttribute{
																		Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
																		MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
																		},
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"http": schema.ListNestedAttribute{
															Description:         "HTTP specific rules.",
															MarkdownDescription: "HTTP specific rules.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"header_matches": schema.ListNestedAttribute{
																		Description:         "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
																		MarkdownDescription: "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"mismatch": schema.StringAttribute{
																					Description:         "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																					MarkdownDescription: "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("LOG", "ADD", "DELETE", "REPLACE"),
																					},
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name identifies the header.",
																					MarkdownDescription: "Name identifies the header.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"secret": schema.SingleNestedAttribute{
																					Description:         "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																					MarkdownDescription: "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Name is the name of the secret.",
																							MarkdownDescription: "Name is the name of the secret.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"namespace": schema.StringAttribute{
																							Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																							MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																					MarkdownDescription: "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": schema.ListAttribute{
																		Description:         "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
																		MarkdownDescription: "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"host": schema.StringAttribute{
																		Description:         "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
																		MarkdownDescription: "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"method": schema.StringAttribute{
																		Description:         "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
																		MarkdownDescription: "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"path": schema.StringAttribute{
																		Description:         "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
																		MarkdownDescription: "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"kafka": schema.ListNestedAttribute{
															Description:         "Kafka-specific rules.",
															MarkdownDescription: "Kafka-specific rules.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"api_key": schema.StringAttribute{
																		Description:         "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
																		MarkdownDescription: "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"api_version": schema.StringAttribute{
																		Description:         "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
																		MarkdownDescription: "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"client_id": schema.StringAttribute{
																		Description:         "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
																		MarkdownDescription: "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"role": schema.StringAttribute{
																		Description:         "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported: - 'produce': Allow producing to the topics specified in the rule - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
																		MarkdownDescription: "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported: - 'produce': Allow producing to the topics specified in the rule - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("produce", "consume"),
																		},
																	},

																	"topic": schema.StringAttribute{
																		Description:         "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
																		MarkdownDescription: "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(255),
																		},
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"l7": schema.ListAttribute{
															Description:         "Key-value pair rules.",
															MarkdownDescription: "Key-value pair rules.",
															ElementType:         types.MapType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"l7proto": schema.StringAttribute{
															Description:         "Name of the L7 protocol for which the Key-value pair rules apply.",
															MarkdownDescription: "Name of the L7 protocol for which the Key-value pair rules apply.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"server_names": schema.ListAttribute{
													Description:         "ServerNames is a list of allowed TLS SNI values. If not empty, then TLS must be present and one of the provided SNIs must be indicated in the TLS handshake.",
													MarkdownDescription: "ServerNames is a list of allowed TLS SNI values. If not empty, then TLS must be present and one of the provided SNIs must be indicated in the TLS handshake.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"terminating_tls": schema.SingleNestedAttribute{
													Description:         "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
													MarkdownDescription: "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
													Attributes: map[string]schema.Attribute{
														"certificate": schema.StringAttribute{
															Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"private_key": schema.StringAttribute{
															Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
															MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the secret.",
																	MarkdownDescription: "Name is the name of the secret.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																	MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"trusted_ca": schema.StringAttribute{
															Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_requires": schema.ListNestedAttribute{
										Description:         "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
										MarkdownDescription: "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_services": schema.ListNestedAttribute{
										Description:         "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
										MarkdownDescription: "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"k8s_service": schema.SingleNestedAttribute{
													Description:         "K8sService selects service by name and namespace pair",
													MarkdownDescription: "K8sService selects service by name and namespace pair",
													Attributes: map[string]schema.Attribute{
														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"k8s_service_selector": schema.SingleNestedAttribute{
													Description:         "K8sServiceSelector selects services by k8s labels and namespace",
													MarkdownDescription: "K8sServiceSelector selects services by k8s labels and namespace",
													Attributes: map[string]schema.Attribute{
														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"selector": schema.SingleNestedAttribute{
															Description:         "ServiceSelector is a label selector for k8s services",
															MarkdownDescription: "ServiceSelector is a label selector for k8s services",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																				},
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"egress_deny": schema.ListNestedAttribute{
							Description:         "EgressDeny is a list of EgressDenyRule which are enforced at egress. Any rule inserted here will be denied regardless of the allowed egress rules in the 'egress' field. If omitted or empty, this rule does not apply at egress.",
							MarkdownDescription: "EgressDeny is a list of EgressDenyRule which are enforced at egress. Any rule inserted here will be denied regardless of the allowed egress rules in the 'egress' field. If omitted or empty, this rule does not apply at egress.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"icmps": schema.ListNestedAttribute{
										Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is not allowed to initiate type 8 ICMP connections.",
										MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'app=httpd' is not allowed to initiate type 8 ICMP connections.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"fields": schema.ListNestedAttribute{
													Description:         "Fields is a list of ICMP fields.",
													MarkdownDescription: "Fields is a list of ICMP fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"family": schema.StringAttribute{
																Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
																MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("IPv4", "IPv6"),
																},
															},

															"type": schema.StringAttribute{
																Description:         "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
																MarkdownDescription: "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_cidr": schema.ListAttribute{
										Description:         "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
										MarkdownDescription: "ToCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections. Only connections destined for outside of the cluster and not targeting the host will be subject to CIDR rules.  This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"to_cidr_set": schema.ListNestedAttribute{
										Description:         "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
										MarkdownDescription: "ToCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to initiate connections to in addition to connections which are allowed via ToEndpoints, along with a list of subnets contained within their corresponding IP block to which traffic should not be allowed. This will match on the destination IP address of outgoing connections. Adding a prefix into ToCIDR or into ToCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between ToCIDR and ToCIDRSet.  Example: Any endpoint with the label 'app=database-proxy' is allowed to initiate connections to 10.2.3.0/24 except from IPs in subnet 10.2.3.0/28.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Description:         "CIDR is a CIDR prefix / IP Block.",
													MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
													},
												},

												"cidr_group_ref": schema.StringAttribute{
													Description:         "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
													MarkdownDescription: "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(253),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
													},
												},

												"except": schema.ListAttribute{
													Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
													MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_endpoints": schema.ListNestedAttribute{
										Description:         "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
										MarkdownDescription: "ToEndpoints is a list of endpoints identified by an EndpointSelector to which the endpoints subject to the rule are allowed to communicate.  Example: Any endpoint with the label 'role=frontend' can communicate with any endpoint carrying the label 'role=backend'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_entities": schema.ListAttribute{
										Description:         "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster','host','remote-node','kube-apiserver', 'init', 'health','unmanaged' and 'all'.",
										MarkdownDescription: "ToEntities is a list of special entities to which the endpoint subject to the rule is allowed to initiate connections. Supported entities are 'world', 'cluster','host','remote-node','kube-apiserver', 'init', 'health','unmanaged' and 'all'.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"to_groups": schema.ListNestedAttribute{
										Description:         "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
										MarkdownDescription: "ToGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: toGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"aws": schema.SingleNestedAttribute{
													Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
													MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
													Attributes: map[string]schema.Attribute{
														"labels": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"region": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_groups_ids": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_groups_names": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_nodes": schema.ListNestedAttribute{
										Description:         "ToNodes is a list of nodes identified by an EndpointSelector to which endpoints subject to the rule is allowed to communicate.",
										MarkdownDescription: "ToNodes is a list of nodes identified by an EndpointSelector to which endpoints subject to the rule is allowed to communicate.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_ports": schema.ListNestedAttribute{
										Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is not allowed to initiate connections to destination port 8080/tcp",
										MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to connect to.  Example: Any endpoint with the label 'role=frontend' is not allowed to initiate connections to destination port 8080/tcp",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"ports": schema.ListNestedAttribute{
													Description:         "Ports is a list of L4 port/protocol",
													MarkdownDescription: "Ports is a list of L4 port/protocol",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"port": schema.StringAttribute{
																Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
																MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
																},
															},

															"protocol": schema.StringAttribute{
																Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
																MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_requires": schema.ListNestedAttribute{
										Description:         "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
										MarkdownDescription: "ToRequires is a list of additional constraints which must be met in order for the selected endpoints to be able to connect to other endpoints. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching ToEndpoints.  Example: Any Endpoint with the label 'team=A' requires any endpoint to which it communicates to also carry the label 'team=A'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_services": schema.ListNestedAttribute{
										Description:         "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
										MarkdownDescription: "ToServices is a list of services to which the endpoint subject to the rule is allowed to initiate connections. Currently Cilium only supports toServices for K8s services without selectors.  Example: Any endpoint with the label 'app=backend-app' is allowed to initiate connections to all cidrs backing the 'external-service' service",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"k8s_service": schema.SingleNestedAttribute{
													Description:         "K8sService selects service by name and namespace pair",
													MarkdownDescription: "K8sService selects service by name and namespace pair",
													Attributes: map[string]schema.Attribute{
														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"k8s_service_selector": schema.SingleNestedAttribute{
													Description:         "K8sServiceSelector selects services by k8s labels and namespace",
													MarkdownDescription: "K8sServiceSelector selects services by k8s labels and namespace",
													Attributes: map[string]schema.Attribute{
														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"selector": schema.SingleNestedAttribute{
															Description:         "ServiceSelector is a label selector for k8s services",
															MarkdownDescription: "ServiceSelector is a label selector for k8s services",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																				},
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"enable_default_deny": schema.SingleNestedAttribute{
							Description:         "EnableDefaultDeny determines whether this policy configures the subject endpoint(s) to have a default deny mode. If enabled, this causes all traffic not explicitly allowed by a network policy to be dropped.  If not specified, the default is true for each traffic direction that has rules, and false otherwise. For example, if a policy only has Ingress or IngressDeny rules, then the default for ingress is true and egress is false.  If multiple policies apply to an endpoint, that endpoint's default deny will be enabled if any policy requests it.  This is useful for creating broad-based network policies that will not cause endpoints to enter default-deny mode.",
							MarkdownDescription: "EnableDefaultDeny determines whether this policy configures the subject endpoint(s) to have a default deny mode. If enabled, this causes all traffic not explicitly allowed by a network policy to be dropped.  If not specified, the default is true for each traffic direction that has rules, and false otherwise. For example, if a policy only has Ingress or IngressDeny rules, then the default for ingress is true and egress is false.  If multiple policies apply to an endpoint, that endpoint's default deny will be enabled if any policy requests it.  This is useful for creating broad-based network policies that will not cause endpoints to enter default-deny mode.",
							Attributes: map[string]schema.Attribute{
								"egress": schema.BoolAttribute{
									Description:         "Whether or not the endpoint should have a default-deny rule applied to egress traffic.",
									MarkdownDescription: "Whether or not the endpoint should have a default-deny rule applied to egress traffic.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ingress": schema.BoolAttribute{
									Description:         "Whether or not the endpoint should have a default-deny rule applied to ingress traffic.",
									MarkdownDescription: "Whether or not the endpoint should have a default-deny rule applied to ingress traffic.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"endpoint_selector": schema.SingleNestedAttribute{
							Description:         "EndpointSelector selects all endpoints which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive.",
							MarkdownDescription: "EndpointSelector selects all endpoints which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive.",
							Attributes: map[string]schema.Attribute{
								"match_expressions": schema.ListNestedAttribute{
									Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
									MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"operator": schema.StringAttribute{
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": schema.ListAttribute{
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"match_labels": schema.MapAttribute{
									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"ingress": schema.ListNestedAttribute{
							Description:         "Ingress is a list of IngressRule which are enforced at ingress. If omitted or empty, this rule does not apply at ingress.",
							MarkdownDescription: "Ingress is a list of IngressRule which are enforced at ingress. If omitted or empty, this rule does not apply at ingress.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"authentication": schema.SingleNestedAttribute{
										Description:         "Authentication is the required authentication type for the allowed traffic, if any.",
										MarkdownDescription: "Authentication is the required authentication type for the allowed traffic, if any.",
										Attributes: map[string]schema.Attribute{
											"mode": schema.StringAttribute{
												Description:         "Mode is the required authentication mode for the allowed traffic, if any.",
												MarkdownDescription: "Mode is the required authentication mode for the allowed traffic, if any.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("disabled", "required", "test-always-fail"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_cidr": schema.ListAttribute{
										Description:         "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
										MarkdownDescription: "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"from_cidr_set": schema.ListNestedAttribute{
										Description:         "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
										MarkdownDescription: "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Description:         "CIDR is a CIDR prefix / IP Block.",
													MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
													},
												},

												"cidr_group_ref": schema.StringAttribute{
													Description:         "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
													MarkdownDescription: "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(253),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
													},
												},

												"except": schema.ListAttribute{
													Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
													MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_endpoints": schema.ListNestedAttribute{
										Description:         "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
										MarkdownDescription: "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_entities": schema.ListAttribute{
										Description:         "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
										MarkdownDescription: "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"from_groups": schema.ListNestedAttribute{
										Description:         "FromGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: FromGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
										MarkdownDescription: "FromGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: FromGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"aws": schema.SingleNestedAttribute{
													Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
													MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
													Attributes: map[string]schema.Attribute{
														"labels": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"region": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_groups_ids": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_groups_names": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_nodes": schema.ListNestedAttribute{
										Description:         "FromNodes is a list of nodes identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.",
										MarkdownDescription: "FromNodes is a list of nodes identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_requires": schema.ListNestedAttribute{
										Description:         "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
										MarkdownDescription: "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"icmps": schema.ListNestedAttribute{
										Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming type 8 ICMP connections.",
										MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming type 8 ICMP connections.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"fields": schema.ListNestedAttribute{
													Description:         "Fields is a list of ICMP fields.",
													MarkdownDescription: "Fields is a list of ICMP fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"family": schema.StringAttribute{
																Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
																MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("IPv4", "IPv6"),
																},
															},

															"type": schema.StringAttribute{
																Description:         "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
																MarkdownDescription: "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_ports": schema.ListNestedAttribute{
										Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming connections on port 80/tcp.",
										MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can only accept incoming connections on port 80/tcp.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"listener": schema.SingleNestedAttribute{
													Description:         "listener specifies the name of a custom Envoy listener to which this traffic should be redirected to.",
													MarkdownDescription: "listener specifies the name of a custom Envoy listener to which this traffic should be redirected to.",
													Attributes: map[string]schema.Attribute{
														"envoy_config": schema.SingleNestedAttribute{
															Description:         "EnvoyConfig is a reference to the CEC or CCEC resource in which the listener is defined.",
															MarkdownDescription: "EnvoyConfig is a reference to the CEC or CCEC resource in which the listener is defined.",
															Attributes: map[string]schema.Attribute{
																"kind": schema.StringAttribute{
																	Description:         "Kind is the resource type being referred to. Defaults to CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig for CiliumNetworkPolicy and CiliumClusterwideNetworkPolicy, respectively. The only case this is currently explicitly needed is when referring to a CiliumClusterwideEnvoyConfig from CiliumNetworkPolicy, as using a namespaced listener from a cluster scoped policy is not allowed.",
																	MarkdownDescription: "Kind is the resource type being referred to. Defaults to CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig for CiliumNetworkPolicy and CiliumClusterwideNetworkPolicy, respectively. The only case this is currently explicitly needed is when referring to a CiliumClusterwideEnvoyConfig from CiliumNetworkPolicy, as using a namespaced listener from a cluster scoped policy is not allowed.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("CiliumEnvoyConfig", "CiliumClusterwideEnvoyConfig"),
																	},
																},

																"name": schema.StringAttribute{
																	Description:         "Name is the resource name of the CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig where the listener is defined in.",
																	MarkdownDescription: "Name is the resource name of the CiliumEnvoyConfig or CiliumClusterwideEnvoyConfig where the listener is defined in.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of the listener.",
															MarkdownDescription: "Name is the name of the listener.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"priority": schema.Int64Attribute{
															Description:         "Priority for this Listener that is used when multiple rules would apply different listeners to a policy map entry. Behavior of this is implementation dependent.",
															MarkdownDescription: "Priority for this Listener that is used when multiple rules would apply different listeners to a policy map entry. Behavior of this is implementation dependent.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(100),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"originating_tls": schema.SingleNestedAttribute{
													Description:         "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
													MarkdownDescription: "OriginatingTLS is the TLS context for the connections originated by the L7 proxy.  For egress policy this specifies the client-side TLS parameters for the upstream connection originating from the L7 proxy to the remote destination. For ingress policy this specifies the client-side TLS parameters for the connection from the L7 proxy to the local endpoint.",
													Attributes: map[string]schema.Attribute{
														"certificate": schema.StringAttribute{
															Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"private_key": schema.StringAttribute{
															Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
															MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the secret.",
																	MarkdownDescription: "Name is the name of the secret.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																	MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"trusted_ca": schema.StringAttribute{
															Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"ports": schema.ListNestedAttribute{
													Description:         "Ports is a list of L4 port/protocol",
													MarkdownDescription: "Ports is a list of L4 port/protocol",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"port": schema.StringAttribute{
																Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
																MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
																},
															},

															"protocol": schema.StringAttribute{
																Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
																MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"rules": schema.SingleNestedAttribute{
													Description:         "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
													MarkdownDescription: "Rules is a list of additional port level rules which must be met in order for the PortRule to allow the traffic. If omitted or empty, no layer 7 rules are enforced.",
													Attributes: map[string]schema.Attribute{
														"dns": schema.ListNestedAttribute{
															Description:         "DNS-specific rules.",
															MarkdownDescription: "DNS-specific rules.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"match_name": schema.StringAttribute{
																		Description:         "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
																		MarkdownDescription: "MatchName matches literal DNS names. A trailing '.' is automatically added when missing.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_]+[.]?)+$`), ""),
																		},
																	},

																	"match_pattern": schema.StringAttribute{
																		Description:         "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
																		MarkdownDescription: "MatchPattern allows using wildcards to match DNS names. All wildcards are case insensitive. The wildcards are: - '*' matches 0 or more DNS valid characters, and may occur anywhere in the pattern. As a special case a '*' as the leftmost character, without a following '.' matches all subdomains as well as the name to the right. A trailing '.' is automatically added when missing.  Examples: '*.cilium.io' matches subomains of cilium at that level www.cilium.io and blog.cilium.io match, cilium.io and google.com do not '*cilium.io' matches cilium.io and all subdomains ends with 'cilium.io' except those containing '.' separator, subcilium.io and sub-cilium.io match, www.cilium.io and blog.cilium.io does not sub*.cilium.io matches subdomains of cilium where the subdomain component begins with 'sub' sub.cilium.io and subdomain.cilium.io match, www.cilium.io, blog.cilium.io, cilium.io and google.com do not",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.RegexMatches(regexp.MustCompile(`^([-a-zA-Z0-9_*]+[.]?)+$`), ""),
																		},
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"http": schema.ListNestedAttribute{
															Description:         "HTTP specific rules.",
															MarkdownDescription: "HTTP specific rules.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"header_matches": schema.ListNestedAttribute{
																		Description:         "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
																		MarkdownDescription: "HeaderMatches is a list of HTTP headers which must be present and match against the given values. Mismatch field can be used to specify what to do when there is no match.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"mismatch": schema.StringAttribute{
																					Description:         "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																					MarkdownDescription: "Mismatch identifies what to do in case there is no match. The default is to drop the request. Otherwise the overall rule is still considered as matching, but the mismatches are logged in the access log.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("LOG", "ADD", "DELETE", "REPLACE"),
																					},
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name identifies the header.",
																					MarkdownDescription: "Name identifies the header.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"secret": schema.SingleNestedAttribute{
																					Description:         "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																					MarkdownDescription: "Secret refers to a secret that contains the value to be matched against. The secret must only contain one entry. If the referred secret does not exist, and there is no 'Value' specified, the match will fail.",
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Name is the name of the secret.",
																							MarkdownDescription: "Name is the name of the secret.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"namespace": schema.StringAttribute{
																							Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																							MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"value": schema.StringAttribute{
																					Description:         "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																					MarkdownDescription: "Value matches the exact value of the header. Can be specified either alone or together with 'Secret'; will be used as the header value if the secret can not be found in the latter case.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"headers": schema.ListAttribute{
																		Description:         "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
																		MarkdownDescription: "Headers is a list of HTTP headers which must be present in the request. If omitted or empty, requests are allowed regardless of headers present.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"host": schema.StringAttribute{
																		Description:         "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
																		MarkdownDescription: "Host is an extended POSIX regex matched against the host header of a request, e.g. 'foo.com'  If omitted or empty, the value of the host header is ignored.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"method": schema.StringAttribute{
																		Description:         "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
																		MarkdownDescription: "Method is an extended POSIX regex matched against the method of a request, e.g. 'GET', 'POST', 'PUT', 'PATCH', 'DELETE', ...  If omitted or empty, all methods are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"path": schema.StringAttribute{
																		Description:         "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
																		MarkdownDescription: "Path is an extended POSIX regex matched against the path of a request. Currently it can contain characters disallowed from the conventional 'path' part of a URL as defined by RFC 3986.  If omitted or empty, all paths are all allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"kafka": schema.ListNestedAttribute{
															Description:         "Kafka-specific rules.",
															MarkdownDescription: "Kafka-specific rules.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"api_key": schema.StringAttribute{
																		Description:         "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
																		MarkdownDescription: "APIKey is a case-insensitive string matched against the key of a request, e.g. 'produce', 'fetch', 'createtopic', 'deletetopic', et al Reference: https://kafka.apache.org/protocol#protocol_api_keys  If omitted or empty, and if Role is not specified, then all keys are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"api_version": schema.StringAttribute{
																		Description:         "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
																		MarkdownDescription: "APIVersion is the version matched against the api version of the Kafka message. If set, it has to be a string representing a positive integer.  If omitted or empty, all versions are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"client_id": schema.StringAttribute{
																		Description:         "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
																		MarkdownDescription: "ClientID is the client identifier as provided in the request.  From Kafka protocol documentation: This is a user supplied identifier for the client application. The user can use any identifier they like and it will be used when logging errors, monitoring aggregates, etc. For example, one might want to monitor not just the requests per second overall, but the number coming from each client application (each of which could reside on multiple servers). This id acts as a logical grouping across all requests from a particular client.  If omitted or empty, all client identifiers are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"role": schema.StringAttribute{
																		Description:         "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported: - 'produce': Allow producing to the topics specified in the rule - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
																		MarkdownDescription: "Role is a case-insensitive string and describes a group of API keys necessary to perform certain higher-level Kafka operations such as 'produce' or 'consume'. A Role automatically expands into all APIKeys required to perform the specified higher-level operation.  The following values are supported: - 'produce': Allow producing to the topics specified in the rule - 'consume': Allow consuming from the topics specified in the rule  This field is incompatible with the APIKey field, i.e APIKey and Role cannot both be specified in the same rule.  If omitted or empty, and if APIKey is not specified, then all keys are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("produce", "consume"),
																		},
																	},

																	"topic": schema.StringAttribute{
																		Description:         "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
																		MarkdownDescription: "Topic is the topic name contained in the message. If a Kafka request contains multiple topics, then all topics must be allowed or the message will be rejected.  This constraint is ignored if the matched request message type doesn't contain any topic. Maximum size of Topic can be 249 characters as per recent Kafka spec and allowed characters are a-z, A-Z, 0-9, -, . and _.  Older Kafka versions had longer topic lengths of 255, but in Kafka 0.10 version the length was changed from 255 to 249. For compatibility reasons we are using 255.  If omitted or empty, all topics are allowed.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(255),
																		},
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"l7": schema.ListAttribute{
															Description:         "Key-value pair rules.",
															MarkdownDescription: "Key-value pair rules.",
															ElementType:         types.MapType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"l7proto": schema.StringAttribute{
															Description:         "Name of the L7 protocol for which the Key-value pair rules apply.",
															MarkdownDescription: "Name of the L7 protocol for which the Key-value pair rules apply.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"server_names": schema.ListAttribute{
													Description:         "ServerNames is a list of allowed TLS SNI values. If not empty, then TLS must be present and one of the provided SNIs must be indicated in the TLS handshake.",
													MarkdownDescription: "ServerNames is a list of allowed TLS SNI values. If not empty, then TLS must be present and one of the provided SNIs must be indicated in the TLS handshake.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"terminating_tls": schema.SingleNestedAttribute{
													Description:         "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
													MarkdownDescription: "TerminatingTLS is the TLS context for the connection terminated by the L7 proxy.  For egress policy this specifies the server-side TLS parameters to be applied on the connections originated from the local endpoint and terminated by the L7 proxy. For ingress policy this specifies the server-side TLS parameters to be applied on the connections originated from a remote source and terminated by the L7 proxy.",
													Attributes: map[string]schema.Attribute{
														"certificate": schema.StringAttribute{
															Description:         "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "Certificate is the file name or k8s secret item name for the certificate chain. If omitted, 'tls.crt' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"private_key": schema.StringAttribute{
															Description:         "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "PrivateKey is the file name or k8s secret item name for the private key matching the certificate chain. If omitted, 'tls.key' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret": schema.SingleNestedAttribute{
															Description:         "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
															MarkdownDescription: "Secret is the secret that contains the certificates and private key for the TLS context. By default, Cilium will search in this secret for the following items: - 'ca.crt'  - Which represents the trusted CA to verify remote source. - 'tls.crt' - Which represents the public key certificate. - 'tls.key' - Which represents the private key matching the public key certificate.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the secret.",
																	MarkdownDescription: "Name is the name of the secret.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"namespace": schema.StringAttribute{
																	Description:         "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																	MarkdownDescription: "Namespace is the namespace in which the secret exists. Context of use determines the default value if left out (e.g., 'default').",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: true,
															Optional: false,
															Computed: false,
														},

														"trusted_ca": schema.StringAttribute{
															Description:         "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
															MarkdownDescription: "TrustedCA is the file name or k8s secret item name for the trusted CA. If omitted, 'ca.crt' is assumed, if it exists. If given, the item must exist.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"ingress_deny": schema.ListNestedAttribute{
							Description:         "IngressDeny is a list of IngressDenyRule which are enforced at ingress. Any rule inserted here will be denied regardless of the allowed ingress rules in the 'ingress' field. If omitted or empty, this rule does not apply at ingress.",
							MarkdownDescription: "IngressDeny is a list of IngressDenyRule which are enforced at ingress. Any rule inserted here will be denied regardless of the allowed ingress rules in the 'ingress' field. If omitted or empty, this rule does not apply at ingress.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"from_cidr": schema.ListAttribute{
										Description:         "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
										MarkdownDescription: "FromCIDR is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from. Only connections which do *not* originate from the cluster or from the local host are subject to CIDR rules. In order to allow in-cluster connectivity, use the FromEndpoints field.  This will match on the source IP address of incoming connections. Adding  a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is  equivalent.  Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.3.9.1",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"from_cidr_set": schema.ListNestedAttribute{
										Description:         "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
										MarkdownDescription: "FromCIDRSet is a list of IP blocks which the endpoint subject to the rule is allowed to receive connections from in addition to FromEndpoints, along with a list of subnets contained within their corresponding IP block from which traffic should not be allowed. This will match on the source IP address of incoming connections. Adding a prefix into FromCIDR or into FromCIDRSet with no ExcludeCIDRs is equivalent. Overlaps are allowed between FromCIDR and FromCIDRSet.  Example: Any endpoint with the label 'app=my-legacy-pet' is allowed to receive connections from 10.0.0.0/8 except from IPs in subnet 10.96.0.0/12.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cidr": schema.StringAttribute{
													Description:         "CIDR is a CIDR prefix / IP Block.",
													MarkdownDescription: "CIDR is a CIDR prefix / IP Block.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\/([0-9]|[1-2][0-9]|3[0-2])$|^s*((([0-9A-Fa-f]{1,4}:){7}(:|([0-9A-Fa-f]{1,4})))|(([0-9A-Fa-f]{1,4}:){6}:([0-9A-Fa-f]{1,4})?)|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){0,1}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){0,2}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){0,3}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){0,4}):([0-9A-Fa-f]{1,4})?))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){0,5}):([0-9A-Fa-f]{1,4})?))|(:(:|((:[0-9A-Fa-f]{1,4}){1,7}))))(%.+)?s*/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8])$`), ""),
													},
												},

												"cidr_group_ref": schema.StringAttribute{
													Description:         "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
													MarkdownDescription: "CIDRGroupRef is a reference to a CiliumCIDRGroup object. A CiliumCIDRGroup contains a list of CIDRs that the endpoint, subject to the rule, can (Ingress/Egress) or cannot (IngressDeny/EgressDeny) receive connections from.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(253),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
													},
												},

												"except": schema.ListAttribute{
													Description:         "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
													MarkdownDescription: "ExceptCIDRs is a list of IP blocks which the endpoint subject to the rule is not allowed to initiate connections to. These CIDR prefixes should be contained within Cidr, using ExceptCIDRs together with CIDRGroupRef is not supported yet. These exceptions are only applied to the Cidr in this CIDRRule, and do not apply to any other CIDR prefixes in any other CIDRRules.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_endpoints": schema.ListNestedAttribute{
										Description:         "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
										MarkdownDescription: "FromEndpoints is a list of endpoints identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.  Example: Any endpoint with the label 'role=backend' can be consumed by any endpoint carrying the label 'role=frontend'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_entities": schema.ListAttribute{
										Description:         "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
										MarkdownDescription: "FromEntities is a list of special entities which the endpoint subject to the rule is allowed to receive connections from. Supported entities are 'world', 'cluster' and 'host'",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"from_groups": schema.ListNestedAttribute{
										Description:         "FromGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: FromGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
										MarkdownDescription: "FromGroups is a directive that allows the integration with multiple outside providers. Currently, only AWS is supported, and the rule can select by multiple sub directives:  Example: FromGroups: - aws: securityGroupsIds: - 'sg-XXXXXXXXXXXXX'",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"aws": schema.SingleNestedAttribute{
													Description:         "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
													MarkdownDescription: "AWSGroup is an structure that can be used to whitelisting information from AWS integration",
													Attributes: map[string]schema.Attribute{
														"labels": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"region": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_groups_ids": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_groups_names": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_nodes": schema.ListNestedAttribute{
										Description:         "FromNodes is a list of nodes identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.",
										MarkdownDescription: "FromNodes is a list of nodes identified by an EndpointSelector which are allowed to communicate with the endpoint subject to the rule.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_requires": schema.ListNestedAttribute{
										Description:         "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
										MarkdownDescription: "FromRequires is a list of additional constraints which must be met in order for the selected endpoints to be reachable. These additional constraints do no by itself grant access privileges and must always be accompanied with at least one matching FromEndpoints.  Example: Any Endpoint with the label 'team=A' requires consuming endpoint to also carry the label 'team=A'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"match_expressions": schema.ListNestedAttribute{
													Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the label key that the selector applies to.",
																MarkdownDescription: "key is the label key that the selector applies to.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"match_labels": schema.MapAttribute{
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"icmps": schema.ListNestedAttribute{
										Description:         "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming type 8 ICMP connections.",
										MarkdownDescription: "ICMPs is a list of ICMP rule identified by type number which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming type 8 ICMP connections.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"fields": schema.ListNestedAttribute{
													Description:         "Fields is a list of ICMP fields.",
													MarkdownDescription: "Fields is a list of ICMP fields.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"family": schema.StringAttribute{
																Description:         "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
																MarkdownDescription: "Family is a IP address version. Currently, we support 'IPv4' and 'IPv6'. 'IPv4' is set as default.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("IPv4", "IPv6"),
																},
															},

															"type": schema.StringAttribute{
																Description:         "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
																MarkdownDescription: "Type is a ICMP-type. It should be an 8bit code (0-255), or it's CamelCase name (for example, 'EchoReply'). Allowed ICMP types are: Ipv4: EchoReply | DestinationUnreachable | Redirect | Echo | EchoRequest | RouterAdvertisement | RouterSelection | TimeExceeded | ParameterProblem | Timestamp | TimestampReply | Photuris | ExtendedEcho Request | ExtendedEcho Reply Ipv6: DestinationUnreachable | PacketTooBig | TimeExceeded | ParameterProblem | EchoRequest | EchoReply | MulticastListenerQuery| MulticastListenerReport | MulticastListenerDone | RouterSolicitation | RouterAdvertisement | NeighborSolicitation | NeighborAdvertisement | RedirectMessage | RouterRenumbering | ICMPNodeInformationQuery | ICMPNodeInformationResponse | InverseNeighborDiscoverySolicitation | InverseNeighborDiscoveryAdvertisement | HomeAgentAddressDiscoveryRequest | HomeAgentAddressDiscoveryReply | MobilePrefixSolicitation | MobilePrefixAdvertisement | DuplicateAddressRequestCodeSuffix | DuplicateAddressConfirmationCodeSuffix | ExtendedEchoRequest | ExtendedEchoReply",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_ports": schema.ListNestedAttribute{
										Description:         "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming connections on port 80/tcp.",
										MarkdownDescription: "ToPorts is a list of destination ports identified by port number and protocol which the endpoint subject to the rule is not allowed to receive connections on.  Example: Any endpoint with the label 'app=httpd' can not accept incoming connections on port 80/tcp.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"ports": schema.ListNestedAttribute{
													Description:         "Ports is a list of L4 port/protocol",
													MarkdownDescription: "Ports is a list of L4 port/protocol",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"port": schema.StringAttribute{
																Description:         "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
																MarkdownDescription: "Port is an L4 port number. For now the string will be strictly parsed as a single uint16. In the future, this field may support ranges in the form '1024-2048 Port can also be a port name, which must contain at least one [a-z], and may also contain [0-9] and '-' anywhere except adjacent to another '-' or in the beginning or the end.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^(6553[0-5]|655[0-2][0-9]|65[0-4][0-9]{2}|6[0-4][0-9]{3}|[1-5][0-9]{4}|[0-9]{1,4})|([a-zA-Z0-9]-?)*[a-zA-Z](-?[a-zA-Z0-9])*$`), ""),
																},
															},

															"protocol": schema.StringAttribute{
																Description:         "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
																MarkdownDescription: "Protocol is the L4 protocol. If omitted or empty, any protocol matches. Accepted values: 'TCP', 'UDP', 'SCTP', 'ANY'  Matching on ICMP is not supported.  Named port specified for a container may narrow this down, but may not contradict this.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TCP", "UDP", "SCTP", "ANY"),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"labels": schema.ListNestedAttribute{
							Description:         "Labels is a list of optional strings which can be used to re-identify the rule or to store metadata. It is possible to lookup or delete strings based on labels. Labels are not required to be unique, multiple rules can have overlapping or identical labels.",
							MarkdownDescription: "Labels is a list of optional strings which can be used to re-identify the rule or to store metadata. It is possible to lookup or delete strings based on labels. Labels are not required to be unique, multiple rules can have overlapping or identical labels.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"source": schema.StringAttribute{
										Description:         "Source can be one of the above values (e.g.: LabelSourceContainer).",
										MarkdownDescription: "Source can be one of the above values (e.g.: LabelSourceContainer).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},

						"node_selector": schema.SingleNestedAttribute{
							Description:         "NodeSelector selects all nodes which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive. Can only be used in CiliumClusterwideNetworkPolicies.",
							MarkdownDescription: "NodeSelector selects all nodes which should be subject to this rule. EndpointSelector and NodeSelector cannot be both empty and are mutually exclusive. Can only be used in CiliumClusterwideNetworkPolicies.",
							Attributes: map[string]schema.Attribute{
								"match_expressions": schema.ListNestedAttribute{
									Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
									MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"operator": schema.StringAttribute{
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
												},
											},

											"values": schema.ListAttribute{
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"match_labels": schema.MapAttribute{
									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
							Required: false,
							Optional: true,
							Computed: false,
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *CiliumIoCiliumClusterwideNetworkPolicyV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cilium_io_cilium_clusterwide_network_policy_v2_manifest")

	var model CiliumIoCiliumClusterwideNetworkPolicyV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cilium.io/v2")
	model.Kind = pointer.String("CiliumClusterwideNetworkPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
