/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package multicluster_crd_antrea_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &MulticlusterCrdAntreaIoResourceImportV1Alpha1Manifest{}
)

func NewMulticlusterCrdAntreaIoResourceImportV1Alpha1Manifest() datasource.DataSource {
	return &MulticlusterCrdAntreaIoResourceImportV1Alpha1Manifest{}
}

type MulticlusterCrdAntreaIoResourceImportV1Alpha1Manifest struct{}

type MulticlusterCrdAntreaIoResourceImportV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClusterID   *[]string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
		Clusterinfo *struct {
			ClusterID    *string `tfsdk:"cluster_id" json:"clusterID,omitempty"`
			GatewayInfos *[]struct {
				GatewayIP *string `tfsdk:"gateway_ip" json:"gatewayIP,omitempty"`
			} `tfsdk:"gateway_infos" json:"gatewayInfos,omitempty"`
			PodCIDRs    *[]string `tfsdk:"pod_cid_rs" json:"podCIDRs,omitempty"`
			ServiceCIDR *string   `tfsdk:"service_cidr" json:"serviceCIDR,omitempty"`
			WireGuard   *struct {
				PublicKey *string `tfsdk:"public_key" json:"publicKey,omitempty"`
			} `tfsdk:"wire_guard" json:"wireGuard,omitempty"`
		} `tfsdk:"clusterinfo" json:"clusterinfo,omitempty"`
		Clusternetworkpolicy *struct {
			AppliedTo *[]struct {
				ExternalEntitySelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"external_entity_selector" json:"externalEntitySelector,omitempty"`
				Group             *string `tfsdk:"group" json:"group,omitempty"`
				NamespaceSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
				NodeSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				PodSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
				Service *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service" json:"service,omitempty"`
				ServiceAccount *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			} `tfsdk:"applied_to" json:"appliedTo,omitempty"`
			Egress *[]struct {
				Action    *string `tfsdk:"action" json:"action,omitempty"`
				AppliedTo *[]struct {
					ExternalEntitySelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"external_entity_selector" json:"externalEntitySelector,omitempty"`
					Group             *string `tfsdk:"group" json:"group,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					NodeSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					Service *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					ServiceAccount *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"applied_to" json:"appliedTo,omitempty"`
				EnableLogging *bool `tfsdk:"enable_logging" json:"enableLogging,omitempty"`
				From          *[]struct {
					ExternalEntitySelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"external_entity_selector" json:"externalEntitySelector,omitempty"`
					Fqdn    *string `tfsdk:"fqdn" json:"fqdn,omitempty"`
					Group   *string `tfsdk:"group" json:"group,omitempty"`
					IpBlock *struct {
						Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
					} `tfsdk:"ip_block" json:"ipBlock,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *struct {
						Match *string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					Scope          *string `tfsdk:"scope" json:"scope,omitempty"`
					ServiceAccount *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
				L7Protocols *[]struct {
					Http *struct {
						Host   *string `tfsdk:"host" json:"host,omitempty"`
						Method *string `tfsdk:"method" json:"method,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Tls *struct {
						Sni *string `tfsdk:"sni" json:"sni,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"l7_protocols" json:"l7Protocols,omitempty"`
				LogLabel *string `tfsdk:"log_label" json:"logLabel,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Ports    *[]struct {
					EndPort       *int64  `tfsdk:"end_port" json:"endPort,omitempty"`
					Port          *string `tfsdk:"port" json:"port,omitempty"`
					Protocol      *string `tfsdk:"protocol" json:"protocol,omitempty"`
					SourceEndPort *int64  `tfsdk:"source_end_port" json:"sourceEndPort,omitempty"`
					SourcePort    *int64  `tfsdk:"source_port" json:"sourcePort,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				Protocols *[]struct {
					Icmp *struct {
						IcmpCode *int64 `tfsdk:"icmp_code" json:"icmpCode,omitempty"`
						IcmpType *int64 `tfsdk:"icmp_type" json:"icmpType,omitempty"`
					} `tfsdk:"icmp" json:"icmp,omitempty"`
					Igmp *struct {
						GroupAddress *string `tfsdk:"group_address" json:"groupAddress,omitempty"`
						IgmpType     *int64  `tfsdk:"igmp_type" json:"igmpType,omitempty"`
					} `tfsdk:"igmp" json:"igmp,omitempty"`
				} `tfsdk:"protocols" json:"protocols,omitempty"`
				To *[]struct {
					ExternalEntitySelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"external_entity_selector" json:"externalEntitySelector,omitempty"`
					Fqdn    *string `tfsdk:"fqdn" json:"fqdn,omitempty"`
					Group   *string `tfsdk:"group" json:"group,omitempty"`
					IpBlock *struct {
						Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
					} `tfsdk:"ip_block" json:"ipBlock,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *struct {
						Match *string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					Scope          *string `tfsdk:"scope" json:"scope,omitempty"`
					ServiceAccount *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"to" json:"to,omitempty"`
				ToServices *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Scope     *string `tfsdk:"scope" json:"scope,omitempty"`
				} `tfsdk:"to_services" json:"toServices,omitempty"`
			} `tfsdk:"egress" json:"egress,omitempty"`
			Ingress *[]struct {
				Action    *string `tfsdk:"action" json:"action,omitempty"`
				AppliedTo *[]struct {
					ExternalEntitySelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"external_entity_selector" json:"externalEntitySelector,omitempty"`
					Group             *string `tfsdk:"group" json:"group,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					NodeSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					Service *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
					ServiceAccount *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"applied_to" json:"appliedTo,omitempty"`
				EnableLogging *bool `tfsdk:"enable_logging" json:"enableLogging,omitempty"`
				From          *[]struct {
					ExternalEntitySelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"external_entity_selector" json:"externalEntitySelector,omitempty"`
					Fqdn    *string `tfsdk:"fqdn" json:"fqdn,omitempty"`
					Group   *string `tfsdk:"group" json:"group,omitempty"`
					IpBlock *struct {
						Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
					} `tfsdk:"ip_block" json:"ipBlock,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *struct {
						Match *string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					Scope          *string `tfsdk:"scope" json:"scope,omitempty"`
					ServiceAccount *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"from" json:"from,omitempty"`
				L7Protocols *[]struct {
					Http *struct {
						Host   *string `tfsdk:"host" json:"host,omitempty"`
						Method *string `tfsdk:"method" json:"method,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Tls *struct {
						Sni *string `tfsdk:"sni" json:"sni,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"l7_protocols" json:"l7Protocols,omitempty"`
				LogLabel *string `tfsdk:"log_label" json:"logLabel,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Ports    *[]struct {
					EndPort       *int64  `tfsdk:"end_port" json:"endPort,omitempty"`
					Port          *string `tfsdk:"port" json:"port,omitempty"`
					Protocol      *string `tfsdk:"protocol" json:"protocol,omitempty"`
					SourceEndPort *int64  `tfsdk:"source_end_port" json:"sourceEndPort,omitempty"`
					SourcePort    *int64  `tfsdk:"source_port" json:"sourcePort,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				Protocols *[]struct {
					Icmp *struct {
						IcmpCode *int64 `tfsdk:"icmp_code" json:"icmpCode,omitempty"`
						IcmpType *int64 `tfsdk:"icmp_type" json:"icmpType,omitempty"`
					} `tfsdk:"icmp" json:"icmp,omitempty"`
					Igmp *struct {
						GroupAddress *string `tfsdk:"group_address" json:"groupAddress,omitempty"`
						IgmpType     *int64  `tfsdk:"igmp_type" json:"igmpType,omitempty"`
					} `tfsdk:"igmp" json:"igmp,omitempty"`
				} `tfsdk:"protocols" json:"protocols,omitempty"`
				To *[]struct {
					ExternalEntitySelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"external_entity_selector" json:"externalEntitySelector,omitempty"`
					Fqdn    *string `tfsdk:"fqdn" json:"fqdn,omitempty"`
					Group   *string `tfsdk:"group" json:"group,omitempty"`
					IpBlock *struct {
						Cidr *string `tfsdk:"cidr" json:"cidr,omitempty"`
					} `tfsdk:"ip_block" json:"ipBlock,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces *struct {
						Match *string `tfsdk:"match" json:"match,omitempty"`
					} `tfsdk:"namespaces" json:"namespaces,omitempty"`
					NodeSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
					Scope          *string `tfsdk:"scope" json:"scope,omitempty"`
					ServiceAccount *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				} `tfsdk:"to" json:"to,omitempty"`
				ToServices *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Scope     *string `tfsdk:"scope" json:"scope,omitempty"`
				} `tfsdk:"to_services" json:"toServices,omitempty"`
			} `tfsdk:"ingress" json:"ingress,omitempty"`
			Priority *float64 `tfsdk:"priority" json:"priority,omitempty"`
			Tier     *string  `tfsdk:"tier" json:"tier,omitempty"`
		} `tfsdk:"clusternetworkpolicy" json:"clusternetworkpolicy,omitempty"`
		Endpoints *struct {
			Subsets *[]struct {
				Addresses *[]struct {
					Hostname  *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Ip        *string `tfsdk:"ip" json:"ip,omitempty"`
					NodeName  *string `tfsdk:"node_name" json:"nodeName,omitempty"`
					TargetRef *struct {
						ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
						ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
						Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
					} `tfsdk:"target_ref" json:"targetRef,omitempty"`
				} `tfsdk:"addresses" json:"addresses,omitempty"`
				NotReadyAddresses *[]struct {
					Hostname  *string `tfsdk:"hostname" json:"hostname,omitempty"`
					Ip        *string `tfsdk:"ip" json:"ip,omitempty"`
					NodeName  *string `tfsdk:"node_name" json:"nodeName,omitempty"`
					TargetRef *struct {
						ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
						ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
						Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
					} `tfsdk:"target_ref" json:"targetRef,omitempty"`
				} `tfsdk:"not_ready_addresses" json:"notReadyAddresses,omitempty"`
				Ports *[]struct {
					AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"subsets" json:"subsets,omitempty"`
		} `tfsdk:"endpoints" json:"endpoints,omitempty"`
		Externalentity *struct {
			Externalentityspec *struct {
				Endpoints *[]struct {
					Ip   *string `tfsdk:"ip" json:"ip,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"endpoints" json:"endpoints,omitempty"`
				ExternalNode *string `tfsdk:"external_node" json:"externalNode,omitempty"`
				Ports        *[]struct {
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Port     *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
			} `tfsdk:"externalentityspec" json:"externalentityspec,omitempty"`
		} `tfsdk:"externalentity" json:"externalentity,omitempty"`
		Kind          *string `tfsdk:"kind" json:"kind,omitempty"`
		LabelIdentity *struct {
			Id    *int64  `tfsdk:"id" json:"id,omitempty"`
			Label *string `tfsdk:"label" json:"label,omitempty"`
		} `tfsdk:"label_identity" json:"labelIdentity,omitempty"`
		Name      *string `tfsdk:"name" json:"name,omitempty"`
		Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		Raw       *struct {
			Data *string `tfsdk:"data" json:"data,omitempty"`
		} `tfsdk:"raw" json:"raw,omitempty"`
		ServiceImport *struct {
			ApiVersion *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string            `tfsdk:"kind" json:"kind,omitempty"`
			Metadata   *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec       *struct {
				Ips   *[]string `tfsdk:"ips" json:"ips,omitempty"`
				Ports *[]struct {
					AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				SessionAffinity       *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
				SessionAffinityConfig *struct {
					ClientIP *struct {
						TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
					} `tfsdk:"client_ip" json:"clientIP,omitempty"`
				} `tfsdk:"session_affinity_config" json:"sessionAffinityConfig,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
			Status *struct {
				Clusters *[]struct {
					Cluster *string `tfsdk:"cluster" json:"cluster,omitempty"`
				} `tfsdk:"clusters" json:"clusters,omitempty"`
			} `tfsdk:"status" json:"status,omitempty"`
		} `tfsdk:"service_import" json:"serviceImport,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MulticlusterCrdAntreaIoResourceImportV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_multicluster_crd_antrea_io_resource_import_v1alpha1_manifest"
}

func (r *MulticlusterCrdAntreaIoResourceImportV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResourceImport is the Schema for the resourceimports API.",
		MarkdownDescription: "ResourceImport is the Schema for the resourceimports API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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

					"namespace": schema.StringAttribute{
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
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
				Description:         "ResourceImportSpec defines the desired state of ResourceImport.",
				MarkdownDescription: "ResourceImportSpec defines the desired state of ResourceImport.",
				Attributes: map[string]schema.Attribute{
					"cluster_id": schema.ListAttribute{
						Description:         "ClusterIDs specifies the member clusters this resource to import to. When not specified, import to all member clusters.",
						MarkdownDescription: "ClusterIDs specifies the member clusters this resource to import to. When not specified, import to all member clusters.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clusterinfo": schema.SingleNestedAttribute{
						Description:         "If imported resource is ClusterInfo.",
						MarkdownDescription: "If imported resource is ClusterInfo.",
						Attributes: map[string]schema.Attribute{
							"cluster_id": schema.StringAttribute{
								Description:         "ClusterID of the member cluster.",
								MarkdownDescription: "ClusterID of the member cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gateway_infos": schema.ListNestedAttribute{
								Description:         "GatewayInfos has information of Gateways",
								MarkdownDescription: "GatewayInfos has information of Gateways",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"gateway_ip": schema.StringAttribute{
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

							"pod_cid_rs": schema.ListAttribute{
								Description:         "PodCIDRs is the Pod IP address CIDRs.",
								MarkdownDescription: "PodCIDRs is the Pod IP address CIDRs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_cidr": schema.StringAttribute{
								Description:         "ServiceCIDR is the IP ranges used by Service ClusterIP.",
								MarkdownDescription: "ServiceCIDR is the IP ranges used by Service ClusterIP.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wire_guard": schema.SingleNestedAttribute{
								Description:         "WireGuardInfo includes information of a WireGuard tunnel.",
								MarkdownDescription: "WireGuardInfo includes information of a WireGuard tunnel.",
								Attributes: map[string]schema.Attribute{
									"public_key": schema.StringAttribute{
										Description:         "Public key of the WireGuard tunnel.",
										MarkdownDescription: "Public key of the WireGuard tunnel.",
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

					"clusternetworkpolicy": schema.SingleNestedAttribute{
						Description:         "If imported resource is AntreaClusterNetworkPolicy.",
						MarkdownDescription: "If imported resource is AntreaClusterNetworkPolicy.",
						Attributes: map[string]schema.Attribute{
							"applied_to": schema.ListNestedAttribute{
								Description:         "Select workloads on which the rules will be applied to. Cannot be set in conjunction with AppliedTo in each rule.",
								MarkdownDescription: "Select workloads on which the rules will be applied to. Cannot be set in conjunction with AppliedTo in each rule.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"external_entity_selector": schema.SingleNestedAttribute{
											Description:         "Select ExternalEntities from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
											MarkdownDescription: "Select ExternalEntities from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

										"group": schema.StringAttribute{
											Description:         "Group is the name of the ClusterGroup which can be set as an AppliedTo in place of a stand-alone selector. A Group cannot be set with any other selector.",
											MarkdownDescription: "Group is the name of the ClusterGroup which can be set as an AppliedTo in place of a stand-alone selector. A Group cannot be set with any other selector.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace_selector": schema.SingleNestedAttribute{
											Description:         "Select all Pods from Namespaces matched by this selector, as workloads in AppliedTo fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
											MarkdownDescription: "Select all Pods from Namespaces matched by this selector, as workloads in AppliedTo fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
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

										"node_selector": schema.SingleNestedAttribute{
											Description:         "Select Nodes in cluster as workloads in AppliedTo fields. Cannot be set with any other selector.",
											MarkdownDescription: "Select Nodes in cluster as workloads in AppliedTo fields. Cannot be set with any other selector.",
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

										"pod_selector": schema.SingleNestedAttribute{
											Description:         "Select Pods from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
											MarkdownDescription: "Select Pods from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

										"service": schema.SingleNestedAttribute{
											Description:         "Select a certain Service which matches the NamespacedName. A Service can only be set in either policy level AppliedTo field in a policy that only has ingress rules or rule level AppliedTo field in an ingress rule. Only a NodePort Service can be referred by this field. Cannot be set with any other selector.",
											MarkdownDescription: "Select a certain Service which matches the NamespacedName. A Service can only be set in either policy level AppliedTo field in a policy that only has ingress rules or rule level AppliedTo field in an ingress rule. Only a NodePort Service can be referred by this field. Cannot be set with any other selector.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
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

										"service_account": schema.SingleNestedAttribute{
											Description:         "Select all Pods with the ServiceAccount matched by this field, as workloads in AppliedTo fields. Cannot be set with any other selector.",
											MarkdownDescription: "Select all Pods with the ServiceAccount matched by this field, as workloads in AppliedTo fields. Cannot be set with any other selector.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"egress": schema.ListNestedAttribute{
								Description:         "Set of egress rules evaluated based on the order in which they are set. Currently Egress rule supports setting the 'To' field but not the 'From' field within a Rule.",
								MarkdownDescription: "Set of egress rules evaluated based on the order in which they are set. Currently Egress rule supports setting the 'To' field but not the 'From' field within a Rule.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action specifies the action to be applied on the rule.",
											MarkdownDescription: "Action specifies the action to be applied on the rule.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"applied_to": schema.ListNestedAttribute{
											Description:         "Select workloads on which this rule will be applied to. Cannot be set in conjunction with NetworkPolicySpec/ClusterNetworkPolicySpec.AppliedTo.",
											MarkdownDescription: "Select workloads on which this rule will be applied to. Cannot be set in conjunction with NetworkPolicySpec/ClusterNetworkPolicySpec.AppliedTo.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"external_entity_selector": schema.SingleNestedAttribute{
														Description:         "Select ExternalEntities from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select ExternalEntities from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"group": schema.StringAttribute{
														Description:         "Group is the name of the ClusterGroup which can be set as an AppliedTo in place of a stand-alone selector. A Group cannot be set with any other selector.",
														MarkdownDescription: "Group is the name of the ClusterGroup which can be set as an AppliedTo in place of a stand-alone selector. A Group cannot be set with any other selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "Select all Pods from Namespaces matched by this selector, as workloads in AppliedTo fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
														MarkdownDescription: "Select all Pods from Namespaces matched by this selector, as workloads in AppliedTo fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
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

													"node_selector": schema.SingleNestedAttribute{
														Description:         "Select Nodes in cluster as workloads in AppliedTo fields. Cannot be set with any other selector.",
														MarkdownDescription: "Select Nodes in cluster as workloads in AppliedTo fields. Cannot be set with any other selector.",
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

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "Select Pods from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select Pods from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"service": schema.SingleNestedAttribute{
														Description:         "Select a certain Service which matches the NamespacedName. A Service can only be set in either policy level AppliedTo field in a policy that only has ingress rules or rule level AppliedTo field in an ingress rule. Only a NodePort Service can be referred by this field. Cannot be set with any other selector.",
														MarkdownDescription: "Select a certain Service which matches the NamespacedName. A Service can only be set in either policy level AppliedTo field in a policy that only has ingress rules or rule level AppliedTo field in an ingress rule. Only a NodePort Service can be referred by this field. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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

													"service_account": schema.SingleNestedAttribute{
														Description:         "Select all Pods with the ServiceAccount matched by this field, as workloads in AppliedTo fields. Cannot be set with any other selector.",
														MarkdownDescription: "Select all Pods with the ServiceAccount matched by this field, as workloads in AppliedTo fields. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"enable_logging": schema.BoolAttribute{
											Description:         "EnableLogging is used to indicate if agent should generate logs when rules are matched. Should be default to false.",
											MarkdownDescription: "EnableLogging is used to indicate if agent should generate logs when rules are matched. Should be default to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"from": schema.ListNestedAttribute{
											Description:         "Rule is matched if traffic originates from workloads selected by this field. If this field is empty, this rule matches all sources.",
											MarkdownDescription: "Rule is matched if traffic originates from workloads selected by this field. If this field is empty, this rule matches all sources.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"external_entity_selector": schema.SingleNestedAttribute{
														Description:         "Select ExternalEntities from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select ExternalEntities from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"fqdn": schema.StringAttribute{
														Description:         "Restrict egress access to the Fully Qualified Domain Names prescribed by name or by wildcard match patterns. This field can only be set for NetworkPolicyPeer of egress rules. Supported formats are: Exact FQDNs such as 'google.com'. Wildcard expressions such as '*wayfair.com'.",
														MarkdownDescription: "Restrict egress access to the Fully Qualified Domain Names prescribed by name or by wildcard match patterns. This field can only be set for NetworkPolicyPeer of egress rules. Supported formats are: Exact FQDNs such as 'google.com'. Wildcard expressions such as '*wayfair.com'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"group": schema.StringAttribute{
														Description:         "Group is the name of the ClusterGroup which can be set within an Ingress or Egress rule in place of a stand-alone selector. A Group cannot be set with any other selector.",
														MarkdownDescription: "Group is the name of the ClusterGroup which can be set within an Ingress or Egress rule in place of a stand-alone selector. A Group cannot be set with any other selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip_block": schema.SingleNestedAttribute{
														Description:         "IPBlock describes the IPAddresses/IPBlocks that is matched in to/from. IPBlock cannot be set as part of the AppliedTo field. Cannot be set with any other selector.",
														MarkdownDescription: "IPBlock describes the IPAddresses/IPBlocks that is matched in to/from. IPBlock cannot be set as part of the AppliedTo field. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"cidr": schema.StringAttribute{
																Description:         "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24'.",
																MarkdownDescription: "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24'.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "Select all Pods from Namespaces matched by this selector, as workloads in To/From fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
														MarkdownDescription: "Select all Pods from Namespaces matched by this selector, as workloads in To/From fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
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

													"namespaces": schema.SingleNestedAttribute{
														Description:         "Select Pod/ExternalEntity from Namespaces matched by specific criteria. Current supported criteria is match: Self, which selects from the same Namespace of the appliedTo workloads. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. This field can only be set when NetworkPolicyPeer is created for ClusterNetworkPolicy ingress/egress rules. Cannot be set with NamespaceSelector.",
														MarkdownDescription: "Select Pod/ExternalEntity from Namespaces matched by specific criteria. Current supported criteria is match: Self, which selects from the same Namespace of the appliedTo workloads. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. This field can only be set when NetworkPolicyPeer is created for ClusterNetworkPolicy ingress/egress rules. Cannot be set with NamespaceSelector.",
														Attributes: map[string]schema.Attribute{
															"match": schema.StringAttribute{
																Description:         "NamespaceMatchType describes Namespace matching strategy.",
																MarkdownDescription: "NamespaceMatchType describes Namespace matching strategy.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selector": schema.SingleNestedAttribute{
														Description:         "Select certain Nodes which match the label selector. A NodeSelector cannot be set with any other selector.",
														MarkdownDescription: "Select certain Nodes which match the label selector. A NodeSelector cannot be set with any other selector.",
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

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "Select Pods from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select Pods from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"scope": schema.StringAttribute{
														Description:         "Define scope of the Pod/NamespaceSelector(s) of this peer. Can only be used in ingress NetworkPolicyPeers. Defaults to 'Cluster'.",
														MarkdownDescription: "Define scope of the Pod/NamespaceSelector(s) of this peer. Can only be used in ingress NetworkPolicyPeers. Defaults to 'Cluster'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_account": schema.SingleNestedAttribute{
														Description:         "Select all Pods with the ServiceAccount matched by this field, as workloads in To/From fields. Cannot be set with any other selector.",
														MarkdownDescription: "Select all Pods with the ServiceAccount matched by this field, as workloads in To/From fields. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"l7_protocols": schema.ListNestedAttribute{
											Description:         "Set of layer 7 protocols matched by the rule. If this field is set, action can only be Allow. When this field is used in a rule, any traffic matching the other layer 3/4 criteria of the rule (typically the 5-tuple) will be forwarded to an application-aware engine for protocol detection and rule enforcement, and the traffic will be allowed if the layer 7 criteria is also matched, otherwise it will be dropped. Therefore, any rules after a layer 7 rule will not be enforced for the traffic.",
											MarkdownDescription: "Set of layer 7 protocols matched by the rule. If this field is set, action can only be Allow. When this field is used in a rule, any traffic matching the other layer 3/4 criteria of the rule (typically the 5-tuple) will be forwarded to an application-aware engine for protocol detection and rule enforcement, and the traffic will be allowed if the layer 7 criteria is also matched, otherwise it will be dropped. Therefore, any rules after a layer 7 rule will not be enforced for the traffic.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"http": schema.SingleNestedAttribute{
														Description:         "HTTPProtocol matches HTTP requests with specific host, method, and path. All fields could be used alone or together. If all fields are not provided, it matches all HTTP requests.",
														MarkdownDescription: "HTTPProtocol matches HTTP requests with specific host, method, and path. All fields could be used alone or together. If all fields are not provided, it matches all HTTP requests.",
														Attributes: map[string]schema.Attribute{
															"host": schema.StringAttribute{
																Description:         "Host represents the hostname present in the URI or the HTTP Host header to match. It does not contain the port associated with the host.",
																MarkdownDescription: "Host represents the hostname present in the URI or the HTTP Host header to match. It does not contain the port associated with the host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"method": schema.StringAttribute{
																Description:         "Method represents the HTTP method to match. It could be GET, POST, PUT, HEAD, DELETE, TRACE, OPTIONS, CONNECT and PATCH.",
																MarkdownDescription: "Method represents the HTTP method to match. It could be GET, POST, PUT, HEAD, DELETE, TRACE, OPTIONS, CONNECT and PATCH.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "Path represents the URI path to match (Ex. '/index.html', '/admin').",
																MarkdownDescription: "Path represents the URI path to match (Ex. '/index.html', '/admin').",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"tls": schema.SingleNestedAttribute{
														Description:         "TLSProtocol matches TLS handshake packets with specific SNI. If the field is not provided, this matches all TLS handshake packets.",
														MarkdownDescription: "TLSProtocol matches TLS handshake packets with specific SNI. If the field is not provided, this matches all TLS handshake packets.",
														Attributes: map[string]schema.Attribute{
															"sni": schema.StringAttribute{
																Description:         "SNI (Server Name Indication) indicates the server domain name in the TLS/SSL hello message.",
																MarkdownDescription: "SNI (Server Name Indication) indicates the server domain name in the TLS/SSL hello message.",
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

										"log_label": schema.StringAttribute{
											Description:         "LogLabel is a user-defined arbitrary string which will be printed in the NetworkPolicy logs.",
											MarkdownDescription: "LogLabel is a user-defined arbitrary string which will be printed in the NetworkPolicy logs.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name describes the intention of this rule. Name should be unique within the policy.",
											MarkdownDescription: "Name describes the intention of this rule. Name should be unique within the policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "Set of ports and protocols matched by the rule. If this field and Protocols are unset or empty, this rule matches all ports.",
											MarkdownDescription: "Set of ports and protocols matched by the rule. If this field and Protocols are unset or empty, this rule matches all ports.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"end_port": schema.Int64Attribute{
														Description:         "EndPort defines the end of the port range, inclusive. It can only be specified when a numerical 'port' is specified.",
														MarkdownDescription: "EndPort defines the end of the port range, inclusive. It can only be specified when a numerical 'port' is specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.StringAttribute{
														Description:         "The port on the given protocol. This can be either a numerical or named port on a Pod. If this field is not provided, this matches all port names and numbers.",
														MarkdownDescription: "The port on the given protocol. This can be either a numerical or named port on a Pod. If this field is not provided, this matches all port names and numbers.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"source_end_port": schema.Int64Attribute{
														Description:         "SourceEndPort defines the end of the source port range, inclusive. It can only be specified when 'sourcePort' is specified.",
														MarkdownDescription: "SourceEndPort defines the end of the source port range, inclusive. It can only be specified when 'sourcePort' is specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"source_port": schema.Int64Attribute{
														Description:         "The source port on the given protocol. This can only be a numerical port. If this field is not provided, rule matches all source ports.",
														MarkdownDescription: "The source port on the given protocol. This can only be a numerical port. If this field is not provided, rule matches all source ports.",
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

										"protocols": schema.ListNestedAttribute{
											Description:         "Set of protocols matched by the rule. If this field and Ports are unset or empty, this rule matches all protocols supported.",
											MarkdownDescription: "Set of protocols matched by the rule. If this field and Ports are unset or empty, this rule matches all protocols supported.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"icmp": schema.SingleNestedAttribute{
														Description:         "ICMPProtocol matches ICMP traffic with specific ICMPType and/or ICMPCode. All fields could be used alone or together. If all fields are not provided, this matches all ICMP traffic.",
														MarkdownDescription: "ICMPProtocol matches ICMP traffic with specific ICMPType and/or ICMPCode. All fields could be used alone or together. If all fields are not provided, this matches all ICMP traffic.",
														Attributes: map[string]schema.Attribute{
															"icmp_code": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icmp_type": schema.Int64Attribute{
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

													"igmp": schema.SingleNestedAttribute{
														Description:         "IGMPProtocol matches IGMP traffic with IGMPType and GroupAddress. IGMPType must be filled with: IGMPQuery    int32 = 0x11 IGMPReportV1 int32 = 0x12 IGMPReportV2 int32 = 0x16 IGMPReportV3 int32 = 0x22 If groupAddress is empty, all groupAddresses will be matched.",
														MarkdownDescription: "IGMPProtocol matches IGMP traffic with IGMPType and GroupAddress. IGMPType must be filled with: IGMPQuery    int32 = 0x11 IGMPReportV1 int32 = 0x12 IGMPReportV2 int32 = 0x16 IGMPReportV3 int32 = 0x22 If groupAddress is empty, all groupAddresses will be matched.",
														Attributes: map[string]schema.Attribute{
															"group_address": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"igmp_type": schema.Int64Attribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"to": schema.ListNestedAttribute{
											Description:         "Rule is matched if traffic is intended for workloads selected by this field. This field can't be used with ToServices. If this field and ToServices are both empty or missing this rule matches all destinations.",
											MarkdownDescription: "Rule is matched if traffic is intended for workloads selected by this field. This field can't be used with ToServices. If this field and ToServices are both empty or missing this rule matches all destinations.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"external_entity_selector": schema.SingleNestedAttribute{
														Description:         "Select ExternalEntities from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select ExternalEntities from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"fqdn": schema.StringAttribute{
														Description:         "Restrict egress access to the Fully Qualified Domain Names prescribed by name or by wildcard match patterns. This field can only be set for NetworkPolicyPeer of egress rules. Supported formats are: Exact FQDNs such as 'google.com'. Wildcard expressions such as '*wayfair.com'.",
														MarkdownDescription: "Restrict egress access to the Fully Qualified Domain Names prescribed by name or by wildcard match patterns. This field can only be set for NetworkPolicyPeer of egress rules. Supported formats are: Exact FQDNs such as 'google.com'. Wildcard expressions such as '*wayfair.com'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"group": schema.StringAttribute{
														Description:         "Group is the name of the ClusterGroup which can be set within an Ingress or Egress rule in place of a stand-alone selector. A Group cannot be set with any other selector.",
														MarkdownDescription: "Group is the name of the ClusterGroup which can be set within an Ingress or Egress rule in place of a stand-alone selector. A Group cannot be set with any other selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip_block": schema.SingleNestedAttribute{
														Description:         "IPBlock describes the IPAddresses/IPBlocks that is matched in to/from. IPBlock cannot be set as part of the AppliedTo field. Cannot be set with any other selector.",
														MarkdownDescription: "IPBlock describes the IPAddresses/IPBlocks that is matched in to/from. IPBlock cannot be set as part of the AppliedTo field. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"cidr": schema.StringAttribute{
																Description:         "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24'.",
																MarkdownDescription: "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24'.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "Select all Pods from Namespaces matched by this selector, as workloads in To/From fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
														MarkdownDescription: "Select all Pods from Namespaces matched by this selector, as workloads in To/From fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
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

													"namespaces": schema.SingleNestedAttribute{
														Description:         "Select Pod/ExternalEntity from Namespaces matched by specific criteria. Current supported criteria is match: Self, which selects from the same Namespace of the appliedTo workloads. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. This field can only be set when NetworkPolicyPeer is created for ClusterNetworkPolicy ingress/egress rules. Cannot be set with NamespaceSelector.",
														MarkdownDescription: "Select Pod/ExternalEntity from Namespaces matched by specific criteria. Current supported criteria is match: Self, which selects from the same Namespace of the appliedTo workloads. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. This field can only be set when NetworkPolicyPeer is created for ClusterNetworkPolicy ingress/egress rules. Cannot be set with NamespaceSelector.",
														Attributes: map[string]schema.Attribute{
															"match": schema.StringAttribute{
																Description:         "NamespaceMatchType describes Namespace matching strategy.",
																MarkdownDescription: "NamespaceMatchType describes Namespace matching strategy.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selector": schema.SingleNestedAttribute{
														Description:         "Select certain Nodes which match the label selector. A NodeSelector cannot be set with any other selector.",
														MarkdownDescription: "Select certain Nodes which match the label selector. A NodeSelector cannot be set with any other selector.",
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

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "Select Pods from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select Pods from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"scope": schema.StringAttribute{
														Description:         "Define scope of the Pod/NamespaceSelector(s) of this peer. Can only be used in ingress NetworkPolicyPeers. Defaults to 'Cluster'.",
														MarkdownDescription: "Define scope of the Pod/NamespaceSelector(s) of this peer. Can only be used in ingress NetworkPolicyPeers. Defaults to 'Cluster'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_account": schema.SingleNestedAttribute{
														Description:         "Select all Pods with the ServiceAccount matched by this field, as workloads in To/From fields. Cannot be set with any other selector.",
														MarkdownDescription: "Select all Pods with the ServiceAccount matched by this field, as workloads in To/From fields. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"to_services": schema.ListNestedAttribute{
											Description:         "Rule is matched if traffic is intended for a Service listed in this field. Currently, only ClusterIP types Services are supported in this field. When scope is set to ClusterSet, it matches traffic intended for a multi-cluster Service listed in this field. Service name and Namespace provided should match the original exported Service. This field can only be used when AntreaProxy is enabled. This field can't be used with To or Ports. If this field and To are both empty or missing, this rule matches all destinations.",
											MarkdownDescription: "Rule is matched if traffic is intended for a Service listed in this field. Currently, only ClusterIP types Services are supported in this field. When scope is set to ClusterSet, it matches traffic intended for a multi-cluster Service listed in this field. Service name and Namespace provided should match the original exported Service. This field can only be used when AntreaProxy is enabled. This field can't be used with To or Ports. If this field and To are both empty or missing, this rule matches all destinations.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scope": schema.StringAttribute{
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress": schema.ListNestedAttribute{
								Description:         "Set of ingress rules evaluated based on the order in which they are set. Currently Ingress rule supports setting the 'From' field but not the 'To' field within a Rule.",
								MarkdownDescription: "Set of ingress rules evaluated based on the order in which they are set. Currently Ingress rule supports setting the 'From' field but not the 'To' field within a Rule.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action specifies the action to be applied on the rule.",
											MarkdownDescription: "Action specifies the action to be applied on the rule.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"applied_to": schema.ListNestedAttribute{
											Description:         "Select workloads on which this rule will be applied to. Cannot be set in conjunction with NetworkPolicySpec/ClusterNetworkPolicySpec.AppliedTo.",
											MarkdownDescription: "Select workloads on which this rule will be applied to. Cannot be set in conjunction with NetworkPolicySpec/ClusterNetworkPolicySpec.AppliedTo.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"external_entity_selector": schema.SingleNestedAttribute{
														Description:         "Select ExternalEntities from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select ExternalEntities from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"group": schema.StringAttribute{
														Description:         "Group is the name of the ClusterGroup which can be set as an AppliedTo in place of a stand-alone selector. A Group cannot be set with any other selector.",
														MarkdownDescription: "Group is the name of the ClusterGroup which can be set as an AppliedTo in place of a stand-alone selector. A Group cannot be set with any other selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "Select all Pods from Namespaces matched by this selector, as workloads in AppliedTo fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
														MarkdownDescription: "Select all Pods from Namespaces matched by this selector, as workloads in AppliedTo fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
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

													"node_selector": schema.SingleNestedAttribute{
														Description:         "Select Nodes in cluster as workloads in AppliedTo fields. Cannot be set with any other selector.",
														MarkdownDescription: "Select Nodes in cluster as workloads in AppliedTo fields. Cannot be set with any other selector.",
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

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "Select Pods from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select Pods from NetworkPolicy's Namespace as workloads in AppliedTo fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"service": schema.SingleNestedAttribute{
														Description:         "Select a certain Service which matches the NamespacedName. A Service can only be set in either policy level AppliedTo field in a policy that only has ingress rules or rule level AppliedTo field in an ingress rule. Only a NodePort Service can be referred by this field. Cannot be set with any other selector.",
														MarkdownDescription: "Select a certain Service which matches the NamespacedName. A Service can only be set in either policy level AppliedTo field in a policy that only has ingress rules or rule level AppliedTo field in an ingress rule. Only a NodePort Service can be referred by this field. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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

													"service_account": schema.SingleNestedAttribute{
														Description:         "Select all Pods with the ServiceAccount matched by this field, as workloads in AppliedTo fields. Cannot be set with any other selector.",
														MarkdownDescription: "Select all Pods with the ServiceAccount matched by this field, as workloads in AppliedTo fields. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"enable_logging": schema.BoolAttribute{
											Description:         "EnableLogging is used to indicate if agent should generate logs when rules are matched. Should be default to false.",
											MarkdownDescription: "EnableLogging is used to indicate if agent should generate logs when rules are matched. Should be default to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"from": schema.ListNestedAttribute{
											Description:         "Rule is matched if traffic originates from workloads selected by this field. If this field is empty, this rule matches all sources.",
											MarkdownDescription: "Rule is matched if traffic originates from workloads selected by this field. If this field is empty, this rule matches all sources.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"external_entity_selector": schema.SingleNestedAttribute{
														Description:         "Select ExternalEntities from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select ExternalEntities from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"fqdn": schema.StringAttribute{
														Description:         "Restrict egress access to the Fully Qualified Domain Names prescribed by name or by wildcard match patterns. This field can only be set for NetworkPolicyPeer of egress rules. Supported formats are: Exact FQDNs such as 'google.com'. Wildcard expressions such as '*wayfair.com'.",
														MarkdownDescription: "Restrict egress access to the Fully Qualified Domain Names prescribed by name or by wildcard match patterns. This field can only be set for NetworkPolicyPeer of egress rules. Supported formats are: Exact FQDNs such as 'google.com'. Wildcard expressions such as '*wayfair.com'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"group": schema.StringAttribute{
														Description:         "Group is the name of the ClusterGroup which can be set within an Ingress or Egress rule in place of a stand-alone selector. A Group cannot be set with any other selector.",
														MarkdownDescription: "Group is the name of the ClusterGroup which can be set within an Ingress or Egress rule in place of a stand-alone selector. A Group cannot be set with any other selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip_block": schema.SingleNestedAttribute{
														Description:         "IPBlock describes the IPAddresses/IPBlocks that is matched in to/from. IPBlock cannot be set as part of the AppliedTo field. Cannot be set with any other selector.",
														MarkdownDescription: "IPBlock describes the IPAddresses/IPBlocks that is matched in to/from. IPBlock cannot be set as part of the AppliedTo field. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"cidr": schema.StringAttribute{
																Description:         "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24'.",
																MarkdownDescription: "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24'.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "Select all Pods from Namespaces matched by this selector, as workloads in To/From fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
														MarkdownDescription: "Select all Pods from Namespaces matched by this selector, as workloads in To/From fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
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

													"namespaces": schema.SingleNestedAttribute{
														Description:         "Select Pod/ExternalEntity from Namespaces matched by specific criteria. Current supported criteria is match: Self, which selects from the same Namespace of the appliedTo workloads. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. This field can only be set when NetworkPolicyPeer is created for ClusterNetworkPolicy ingress/egress rules. Cannot be set with NamespaceSelector.",
														MarkdownDescription: "Select Pod/ExternalEntity from Namespaces matched by specific criteria. Current supported criteria is match: Self, which selects from the same Namespace of the appliedTo workloads. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. This field can only be set when NetworkPolicyPeer is created for ClusterNetworkPolicy ingress/egress rules. Cannot be set with NamespaceSelector.",
														Attributes: map[string]schema.Attribute{
															"match": schema.StringAttribute{
																Description:         "NamespaceMatchType describes Namespace matching strategy.",
																MarkdownDescription: "NamespaceMatchType describes Namespace matching strategy.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selector": schema.SingleNestedAttribute{
														Description:         "Select certain Nodes which match the label selector. A NodeSelector cannot be set with any other selector.",
														MarkdownDescription: "Select certain Nodes which match the label selector. A NodeSelector cannot be set with any other selector.",
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

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "Select Pods from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select Pods from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"scope": schema.StringAttribute{
														Description:         "Define scope of the Pod/NamespaceSelector(s) of this peer. Can only be used in ingress NetworkPolicyPeers. Defaults to 'Cluster'.",
														MarkdownDescription: "Define scope of the Pod/NamespaceSelector(s) of this peer. Can only be used in ingress NetworkPolicyPeers. Defaults to 'Cluster'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_account": schema.SingleNestedAttribute{
														Description:         "Select all Pods with the ServiceAccount matched by this field, as workloads in To/From fields. Cannot be set with any other selector.",
														MarkdownDescription: "Select all Pods with the ServiceAccount matched by this field, as workloads in To/From fields. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"l7_protocols": schema.ListNestedAttribute{
											Description:         "Set of layer 7 protocols matched by the rule. If this field is set, action can only be Allow. When this field is used in a rule, any traffic matching the other layer 3/4 criteria of the rule (typically the 5-tuple) will be forwarded to an application-aware engine for protocol detection and rule enforcement, and the traffic will be allowed if the layer 7 criteria is also matched, otherwise it will be dropped. Therefore, any rules after a layer 7 rule will not be enforced for the traffic.",
											MarkdownDescription: "Set of layer 7 protocols matched by the rule. If this field is set, action can only be Allow. When this field is used in a rule, any traffic matching the other layer 3/4 criteria of the rule (typically the 5-tuple) will be forwarded to an application-aware engine for protocol detection and rule enforcement, and the traffic will be allowed if the layer 7 criteria is also matched, otherwise it will be dropped. Therefore, any rules after a layer 7 rule will not be enforced for the traffic.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"http": schema.SingleNestedAttribute{
														Description:         "HTTPProtocol matches HTTP requests with specific host, method, and path. All fields could be used alone or together. If all fields are not provided, it matches all HTTP requests.",
														MarkdownDescription: "HTTPProtocol matches HTTP requests with specific host, method, and path. All fields could be used alone or together. If all fields are not provided, it matches all HTTP requests.",
														Attributes: map[string]schema.Attribute{
															"host": schema.StringAttribute{
																Description:         "Host represents the hostname present in the URI or the HTTP Host header to match. It does not contain the port associated with the host.",
																MarkdownDescription: "Host represents the hostname present in the URI or the HTTP Host header to match. It does not contain the port associated with the host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"method": schema.StringAttribute{
																Description:         "Method represents the HTTP method to match. It could be GET, POST, PUT, HEAD, DELETE, TRACE, OPTIONS, CONNECT and PATCH.",
																MarkdownDescription: "Method represents the HTTP method to match. It could be GET, POST, PUT, HEAD, DELETE, TRACE, OPTIONS, CONNECT and PATCH.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "Path represents the URI path to match (Ex. '/index.html', '/admin').",
																MarkdownDescription: "Path represents the URI path to match (Ex. '/index.html', '/admin').",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"tls": schema.SingleNestedAttribute{
														Description:         "TLSProtocol matches TLS handshake packets with specific SNI. If the field is not provided, this matches all TLS handshake packets.",
														MarkdownDescription: "TLSProtocol matches TLS handshake packets with specific SNI. If the field is not provided, this matches all TLS handshake packets.",
														Attributes: map[string]schema.Attribute{
															"sni": schema.StringAttribute{
																Description:         "SNI (Server Name Indication) indicates the server domain name in the TLS/SSL hello message.",
																MarkdownDescription: "SNI (Server Name Indication) indicates the server domain name in the TLS/SSL hello message.",
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

										"log_label": schema.StringAttribute{
											Description:         "LogLabel is a user-defined arbitrary string which will be printed in the NetworkPolicy logs.",
											MarkdownDescription: "LogLabel is a user-defined arbitrary string which will be printed in the NetworkPolicy logs.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name describes the intention of this rule. Name should be unique within the policy.",
											MarkdownDescription: "Name describes the intention of this rule. Name should be unique within the policy.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "Set of ports and protocols matched by the rule. If this field and Protocols are unset or empty, this rule matches all ports.",
											MarkdownDescription: "Set of ports and protocols matched by the rule. If this field and Protocols are unset or empty, this rule matches all ports.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"end_port": schema.Int64Attribute{
														Description:         "EndPort defines the end of the port range, inclusive. It can only be specified when a numerical 'port' is specified.",
														MarkdownDescription: "EndPort defines the end of the port range, inclusive. It can only be specified when a numerical 'port' is specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.StringAttribute{
														Description:         "The port on the given protocol. This can be either a numerical or named port on a Pod. If this field is not provided, this matches all port names and numbers.",
														MarkdownDescription: "The port on the given protocol. This can be either a numerical or named port on a Pod. If this field is not provided, this matches all port names and numbers.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"source_end_port": schema.Int64Attribute{
														Description:         "SourceEndPort defines the end of the source port range, inclusive. It can only be specified when 'sourcePort' is specified.",
														MarkdownDescription: "SourceEndPort defines the end of the source port range, inclusive. It can only be specified when 'sourcePort' is specified.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"source_port": schema.Int64Attribute{
														Description:         "The source port on the given protocol. This can only be a numerical port. If this field is not provided, rule matches all source ports.",
														MarkdownDescription: "The source port on the given protocol. This can only be a numerical port. If this field is not provided, rule matches all source ports.",
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

										"protocols": schema.ListNestedAttribute{
											Description:         "Set of protocols matched by the rule. If this field and Ports are unset or empty, this rule matches all protocols supported.",
											MarkdownDescription: "Set of protocols matched by the rule. If this field and Ports are unset or empty, this rule matches all protocols supported.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"icmp": schema.SingleNestedAttribute{
														Description:         "ICMPProtocol matches ICMP traffic with specific ICMPType and/or ICMPCode. All fields could be used alone or together. If all fields are not provided, this matches all ICMP traffic.",
														MarkdownDescription: "ICMPProtocol matches ICMP traffic with specific ICMPType and/or ICMPCode. All fields could be used alone or together. If all fields are not provided, this matches all ICMP traffic.",
														Attributes: map[string]schema.Attribute{
															"icmp_code": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"icmp_type": schema.Int64Attribute{
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

													"igmp": schema.SingleNestedAttribute{
														Description:         "IGMPProtocol matches IGMP traffic with IGMPType and GroupAddress. IGMPType must be filled with: IGMPQuery    int32 = 0x11 IGMPReportV1 int32 = 0x12 IGMPReportV2 int32 = 0x16 IGMPReportV3 int32 = 0x22 If groupAddress is empty, all groupAddresses will be matched.",
														MarkdownDescription: "IGMPProtocol matches IGMP traffic with IGMPType and GroupAddress. IGMPType must be filled with: IGMPQuery    int32 = 0x11 IGMPReportV1 int32 = 0x12 IGMPReportV2 int32 = 0x16 IGMPReportV3 int32 = 0x22 If groupAddress is empty, all groupAddresses will be matched.",
														Attributes: map[string]schema.Attribute{
															"group_address": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"igmp_type": schema.Int64Attribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"to": schema.ListNestedAttribute{
											Description:         "Rule is matched if traffic is intended for workloads selected by this field. This field can't be used with ToServices. If this field and ToServices are both empty or missing this rule matches all destinations.",
											MarkdownDescription: "Rule is matched if traffic is intended for workloads selected by this field. This field can't be used with ToServices. If this field and ToServices are both empty or missing this rule matches all destinations.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"external_entity_selector": schema.SingleNestedAttribute{
														Description:         "Select ExternalEntities from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select ExternalEntities from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, ExternalEntities are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"fqdn": schema.StringAttribute{
														Description:         "Restrict egress access to the Fully Qualified Domain Names prescribed by name or by wildcard match patterns. This field can only be set for NetworkPolicyPeer of egress rules. Supported formats are: Exact FQDNs such as 'google.com'. Wildcard expressions such as '*wayfair.com'.",
														MarkdownDescription: "Restrict egress access to the Fully Qualified Domain Names prescribed by name or by wildcard match patterns. This field can only be set for NetworkPolicyPeer of egress rules. Supported formats are: Exact FQDNs such as 'google.com'. Wildcard expressions such as '*wayfair.com'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"group": schema.StringAttribute{
														Description:         "Group is the name of the ClusterGroup which can be set within an Ingress or Egress rule in place of a stand-alone selector. A Group cannot be set with any other selector.",
														MarkdownDescription: "Group is the name of the ClusterGroup which can be set within an Ingress or Egress rule in place of a stand-alone selector. A Group cannot be set with any other selector.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip_block": schema.SingleNestedAttribute{
														Description:         "IPBlock describes the IPAddresses/IPBlocks that is matched in to/from. IPBlock cannot be set as part of the AppliedTo field. Cannot be set with any other selector.",
														MarkdownDescription: "IPBlock describes the IPAddresses/IPBlocks that is matched in to/from. IPBlock cannot be set as part of the AppliedTo field. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"cidr": schema.StringAttribute{
																Description:         "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24'.",
																MarkdownDescription: "CIDR is a string representing the IP Block Valid examples are '192.168.1.1/24'.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "Select all Pods from Namespaces matched by this selector, as workloads in To/From fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
														MarkdownDescription: "Select all Pods from Namespaces matched by this selector, as workloads in To/From fields. If set with PodSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. Cannot be set with Namespaces.",
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

													"namespaces": schema.SingleNestedAttribute{
														Description:         "Select Pod/ExternalEntity from Namespaces matched by specific criteria. Current supported criteria is match: Self, which selects from the same Namespace of the appliedTo workloads. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. This field can only be set when NetworkPolicyPeer is created for ClusterNetworkPolicy ingress/egress rules. Cannot be set with NamespaceSelector.",
														MarkdownDescription: "Select Pod/ExternalEntity from Namespaces matched by specific criteria. Current supported criteria is match: Self, which selects from the same Namespace of the appliedTo workloads. Cannot be set with any other selector except PodSelector or ExternalEntitySelector. This field can only be set when NetworkPolicyPeer is created for ClusterNetworkPolicy ingress/egress rules. Cannot be set with NamespaceSelector.",
														Attributes: map[string]schema.Attribute{
															"match": schema.StringAttribute{
																Description:         "NamespaceMatchType describes Namespace matching strategy.",
																MarkdownDescription: "NamespaceMatchType describes Namespace matching strategy.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_selector": schema.SingleNestedAttribute{
														Description:         "Select certain Nodes which match the label selector. A NodeSelector cannot be set with any other selector.",
														MarkdownDescription: "Select certain Nodes which match the label selector. A NodeSelector cannot be set with any other selector.",
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

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "Select Pods from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
														MarkdownDescription: "Select Pods from NetworkPolicy's Namespace as workloads in To/From fields. If set with NamespaceSelector, Pods are matched from Namespaces matched by the NamespaceSelector. Cannot be set with any other selector except NamespaceSelector.",
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

													"scope": schema.StringAttribute{
														Description:         "Define scope of the Pod/NamespaceSelector(s) of this peer. Can only be used in ingress NetworkPolicyPeers. Defaults to 'Cluster'.",
														MarkdownDescription: "Define scope of the Pod/NamespaceSelector(s) of this peer. Can only be used in ingress NetworkPolicyPeers. Defaults to 'Cluster'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_account": schema.SingleNestedAttribute{
														Description:         "Select all Pods with the ServiceAccount matched by this field, as workloads in To/From fields. Cannot be set with any other selector.",
														MarkdownDescription: "Select all Pods with the ServiceAccount matched by this field, as workloads in To/From fields. Cannot be set with any other selector.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"to_services": schema.ListNestedAttribute{
											Description:         "Rule is matched if traffic is intended for a Service listed in this field. Currently, only ClusterIP types Services are supported in this field. When scope is set to ClusterSet, it matches traffic intended for a multi-cluster Service listed in this field. Service name and Namespace provided should match the original exported Service. This field can only be used when AntreaProxy is enabled. This field can't be used with To or Ports. If this field and To are both empty or missing, this rule matches all destinations.",
											MarkdownDescription: "Rule is matched if traffic is intended for a Service listed in this field. Currently, only ClusterIP types Services are supported in this field. When scope is set to ClusterSet, it matches traffic intended for a multi-cluster Service listed in this field. Service name and Namespace provided should match the original exported Service. This field can only be used when AntreaProxy is enabled. This field can't be used with To or Ports. If this field and To are both empty or missing, this rule matches all destinations.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scope": schema.StringAttribute{
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"priority": schema.Float64Attribute{
								Description:         "Priority specfies the order of the ClusterNetworkPolicy relative to other AntreaClusterNetworkPolicies.",
								MarkdownDescription: "Priority specfies the order of the ClusterNetworkPolicy relative to other AntreaClusterNetworkPolicies.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"tier": schema.StringAttribute{
								Description:         "Tier specifies the tier to which this ClusterNetworkPolicy belongs to. The ClusterNetworkPolicy order will be determined based on the combination of the Tier's Priority and the ClusterNetworkPolicy's own Priority. If not specified, this policy will be created in the Application Tier right above the K8s NetworkPolicy which resides at the bottom.",
								MarkdownDescription: "Tier specifies the tier to which this ClusterNetworkPolicy belongs to. The ClusterNetworkPolicy order will be determined based on the combination of the Tier's Priority and the ClusterNetworkPolicy's own Priority. If not specified, this policy will be created in the Application Tier right above the K8s NetworkPolicy which resides at the bottom.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"endpoints": schema.SingleNestedAttribute{
						Description:         "If imported resource is EndPoints.",
						MarkdownDescription: "If imported resource is EndPoints.",
						Attributes: map[string]schema.Attribute{
							"subsets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"addresses": schema.ListNestedAttribute{
											Description:         "IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.",
											MarkdownDescription: "IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"hostname": schema.StringAttribute{
														Description:         "The Hostname of this endpoint",
														MarkdownDescription: "The Hostname of this endpoint",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip": schema.StringAttribute{
														Description:         "The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready. TODO: This should allow hostname or IP, See #4447.",
														MarkdownDescription: "The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready. TODO: This should allow hostname or IP, See #4447.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"node_name": schema.StringAttribute{
														Description:         "Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.",
														MarkdownDescription: "Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_ref": schema.SingleNestedAttribute{
														Description:         "Reference to object providing the endpoint.",
														MarkdownDescription: "Reference to object providing the endpoint.",
														Attributes: map[string]schema.Attribute{
															"api_version": schema.StringAttribute{
																Description:         "API version of the referent.",
																MarkdownDescription: "API version of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"field_path": schema.StringAttribute{
																Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
																MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resource_version": schema.StringAttribute{
																Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uid": schema.StringAttribute{
																Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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

										"not_ready_addresses": schema.ListNestedAttribute{
											Description:         "IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.",
											MarkdownDescription: "IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"hostname": schema.StringAttribute{
														Description:         "The Hostname of this endpoint",
														MarkdownDescription: "The Hostname of this endpoint",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ip": schema.StringAttribute{
														Description:         "The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready. TODO: This should allow hostname or IP, See #4447.",
														MarkdownDescription: "The IP of this endpoint. May not be loopback (127.0.0.0/8), link-local (169.254.0.0/16), or link-local multicast ((224.0.0.0/24). IPv6 is also accepted but not fully supported on all platforms. Also, certain kubernetes components, like kube-proxy, are not IPv6 ready. TODO: This should allow hostname or IP, See #4447.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"node_name": schema.StringAttribute{
														Description:         "Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.",
														MarkdownDescription: "Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_ref": schema.SingleNestedAttribute{
														Description:         "Reference to object providing the endpoint.",
														MarkdownDescription: "Reference to object providing the endpoint.",
														Attributes: map[string]schema.Attribute{
															"api_version": schema.StringAttribute{
																Description:         "API version of the referent.",
																MarkdownDescription: "API version of the referent.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"field_path": schema.StringAttribute{
																Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
																MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"kind": schema.StringAttribute{
																Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"namespace": schema.StringAttribute{
																Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resource_version": schema.StringAttribute{
																Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"uid": schema.StringAttribute{
																Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
																MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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

										"ports": schema.ListNestedAttribute{
											Description:         "Port numbers available on the related IP addresses.",
											MarkdownDescription: "Port numbers available on the related IP addresses.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"app_protocol": schema.StringAttribute{
														Description:         "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.",
														MarkdownDescription: "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.",
														MarkdownDescription: "The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"port": schema.Int64Attribute{
														Description:         "The port number of the endpoint.",
														MarkdownDescription: "The port number of the endpoint.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
														MarkdownDescription: "The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.",
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

					"externalentity": schema.SingleNestedAttribute{
						Description:         "If imported resource is ExternalEntity.",
						MarkdownDescription: "If imported resource is ExternalEntity.",
						Attributes: map[string]schema.Attribute{
							"externalentityspec": schema.SingleNestedAttribute{
								Description:         "ExternalEntitySpec defines the desired state for ExternalEntity.",
								MarkdownDescription: "ExternalEntitySpec defines the desired state for ExternalEntity.",
								Attributes: map[string]schema.Attribute{
									"endpoints": schema.ListNestedAttribute{
										Description:         "Endpoints is a list of external endpoints associated with this entity.",
										MarkdownDescription: "Endpoints is a list of external endpoints associated with this entity.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"ip": schema.StringAttribute{
													Description:         "IP associated with this endpoint.",
													MarkdownDescription: "IP associated with this endpoint.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name identifies this endpoint. Could be the network interface name in case of VMs.",
													MarkdownDescription: "Name identifies this endpoint. Could be the network interface name in case of VMs.",
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

									"external_node": schema.StringAttribute{
										Description:         "ExternalNode is the opaque identifier of the agent/controller responsible for additional processing or handling of this external entity.",
										MarkdownDescription: "ExternalNode is the opaque identifier of the agent/controller responsible for additional processing or handling of this external entity.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ports": schema.ListNestedAttribute{
										Description:         "Ports maintain the list of named ports.",
										MarkdownDescription: "Ports maintain the list of named ports.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name associated with the Port.",
													MarkdownDescription: "Name associated with the Port.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "The port on the given protocol.",
													MarkdownDescription: "The port on the given protocol.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
													MarkdownDescription: "The protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
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

					"kind": schema.StringAttribute{
						Description:         "Kind of imported resource.",
						MarkdownDescription: "Kind of imported resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"label_identity": schema.SingleNestedAttribute{
						Description:         "If imported resource kind is LabelIdentity.",
						MarkdownDescription: "If imported resource kind is LabelIdentity.",
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Description:         "ID is the ID allocated for the label identity by the leader cluster.",
								MarkdownDescription: "ID is the ID allocated for the label identity by the leader cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label": schema.StringAttribute{
								Description:         "Label is the normalized string of a label identity. The format of normalized label identity is 'ns:(?P<nslabels>(.)*)&pod:(?P<podlabels>(.)*)' E.g., 'ns:kubernetes.io/metadata.name=kube-system&pod:app=db'",
								MarkdownDescription: "Label is the normalized string of a label identity. The format of normalized label identity is 'ns:(?P<nslabels>(.)*)&pod:(?P<podlabels>(.)*)' E.g., 'ns:kubernetes.io/metadata.name=kube-system&pod:app=db'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of imported resource.",
						MarkdownDescription: "Name of imported resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "Namespace of imported resource.",
						MarkdownDescription: "Namespace of imported resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"raw": schema.SingleNestedAttribute{
						Description:         "If imported resource kind is unknown.",
						MarkdownDescription: "If imported resource kind is unknown.",
						Attributes: map[string]schema.Attribute{
							"data": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.Base64Validator(),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_import": schema.SingleNestedAttribute{
						Description:         "If imported resource is ServiceImport.",
						MarkdownDescription: "If imported resource is ServiceImport.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spec": schema.SingleNestedAttribute{
								Description:         "spec defines the behavior of a ServiceImport.",
								MarkdownDescription: "spec defines the behavior of a ServiceImport.",
								Attributes: map[string]schema.Attribute{
									"ips": schema.ListAttribute{
										Description:         "ip will be used as the VIP for this service when type is ClusterSetIP.",
										MarkdownDescription: "ip will be used as the VIP for this service when type is ClusterSetIP.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ports": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"app_protocol": schema.StringAttribute{
													Description:         "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. Field can be enabled with ServiceAppProtocol feature gate.",
													MarkdownDescription: "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. Field can be enabled with ServiceAppProtocol feature gate.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
													MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "The port that will be exposed by this service.",
													MarkdownDescription: "The port that will be exposed by this service.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
													MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"session_affinity": schema.StringAttribute{
										Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. Ignored when type is Headless More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. Ignored when type is Headless More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"session_affinity_config": schema.SingleNestedAttribute{
										Description:         "sessionAffinityConfig contains session affinity configuration.",
										MarkdownDescription: "sessionAffinityConfig contains session affinity configuration.",
										Attributes: map[string]schema.Attribute{
											"client_ip": schema.SingleNestedAttribute{
												Description:         "clientIP contains the configurations of Client IP based session affinity.",
												MarkdownDescription: "clientIP contains the configurations of Client IP based session affinity.",
												Attributes: map[string]schema.Attribute{
													"timeout_seconds": schema.Int64Attribute{
														Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
														MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'. Default value is 10800(for 3 hours).",
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

									"type": schema.StringAttribute{
										Description:         "type defines the type of this service. Must be ClusterSetIP or Headless.",
										MarkdownDescription: "type defines the type of this service. Must be ClusterSetIP or Headless.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("ClusterSetIP", "Headless"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"status": schema.SingleNestedAttribute{
								Description:         "status contains information about the exported services that form the multi-cluster service referenced by this ServiceImport.",
								MarkdownDescription: "status contains information about the exported services that form the multi-cluster service referenced by this ServiceImport.",
								Attributes: map[string]schema.Attribute{
									"clusters": schema.ListNestedAttribute{
										Description:         "clusters is the list of exporting clusters from which this service was derived.",
										MarkdownDescription: "clusters is the list of exporting clusters from which this service was derived.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cluster": schema.StringAttribute{
													Description:         "cluster is the name of the exporting cluster. Must be a valid RFC-1123 DNS label.",
													MarkdownDescription: "cluster is the name of the exporting cluster. Must be a valid RFC-1123 DNS label.",
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
								Required: false,
								Optional: true,
								Computed: false,
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
		},
	}
}

func (r *MulticlusterCrdAntreaIoResourceImportV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_multicluster_crd_antrea_io_resource_import_v1alpha1_manifest")

	var model MulticlusterCrdAntreaIoResourceImportV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("multicluster.crd.antrea.io/v1alpha1")
	model.Kind = pointer.String("ResourceImport")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
