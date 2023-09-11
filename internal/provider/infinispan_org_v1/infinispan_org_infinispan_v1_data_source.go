/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infinispan_org_v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &InfinispanOrgInfinispanV1DataSource{}
	_ datasource.DataSourceWithConfigure = &InfinispanOrgInfinispanV1DataSource{}
)

func NewInfinispanOrgInfinispanV1DataSource() datasource.DataSource {
	return &InfinispanOrgInfinispanV1DataSource{}
}

type InfinispanOrgInfinispanV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type InfinispanOrgInfinispanV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Affinity *struct {
			NodeAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					Preference *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchFields *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_fields" json:"matchFields,omitempty"`
					} `tfsdk:"preference" json:"preference,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *struct {
					NodeSelectorTerms *[]struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchFields *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_fields" json:"matchFields,omitempty"`
					} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
			PodAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
			PodAntiAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		Autoscale *struct {
			Disabled           *bool  `tfsdk:"disabled" json:"disabled,omitempty"`
			MaxMemUsagePercent *int64 `tfsdk:"max_mem_usage_percent" json:"maxMemUsagePercent,omitempty"`
			MaxReplicas        *int64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
			MinMemUsagePercent *int64 `tfsdk:"min_mem_usage_percent" json:"minMemUsagePercent,omitempty"`
			MinReplicas        *int64 `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
		} `tfsdk:"autoscale" json:"autoscale,omitempty"`
		CloudEvents *struct {
			Acks              *string `tfsdk:"acks" json:"acks,omitempty"`
			BootstrapServers  *string `tfsdk:"bootstrap_servers" json:"bootstrapServers,omitempty"`
			CacheEntriesTopic *string `tfsdk:"cache_entries_topic" json:"cacheEntriesTopic,omitempty"`
		} `tfsdk:"cloud_events" json:"cloudEvents,omitempty"`
		ConfigListener *struct {
			Cpu     *string `tfsdk:"cpu" json:"cpu,omitempty"`
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Logging *struct {
				Level *string `tfsdk:"level" json:"level,omitempty"`
			} `tfsdk:"logging" json:"logging,omitempty"`
			Memory *string `tfsdk:"memory" json:"memory,omitempty"`
		} `tfsdk:"config_listener" json:"configListener,omitempty"`
		ConfigMapName *string `tfsdk:"config_map_name" json:"configMapName,omitempty"`
		Container     *struct {
			CliExtraJvmOpts    *string `tfsdk:"cli_extra_jvm_opts" json:"cliExtraJvmOpts,omitempty"`
			Cpu                *string `tfsdk:"cpu" json:"cpu,omitempty"`
			ExtraJvmOpts       *string `tfsdk:"extra_jvm_opts" json:"extraJvmOpts,omitempty"`
			Memory             *string `tfsdk:"memory" json:"memory,omitempty"`
			RouterExtraJvmOpts *string `tfsdk:"router_extra_jvm_opts" json:"routerExtraJvmOpts,omitempty"`
		} `tfsdk:"container" json:"container,omitempty"`
		Dependencies *struct {
			Artifacts *[]struct {
				Hash  *string `tfsdk:"hash" json:"hash,omitempty"`
				Maven *string `tfsdk:"maven" json:"maven,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Url   *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"artifacts" json:"artifacts,omitempty"`
			VolumeClaimName *string `tfsdk:"volume_claim_name" json:"volumeClaimName,omitempty"`
		} `tfsdk:"dependencies" json:"dependencies,omitempty"`
		Expose *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Host        *string            `tfsdk:"host" json:"host,omitempty"`
			NodePort    *int64             `tfsdk:"node_port" json:"nodePort,omitempty"`
			Port        *int64             `tfsdk:"port" json:"port,omitempty"`
			Type        *string            `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"expose" json:"expose,omitempty"`
		Image *string `tfsdk:"image" json:"image,omitempty"`
		Jmx   *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"jmx" json:"jmx,omitempty"`
		Logging *struct {
			Categories *map[string]string `tfsdk:"categories" json:"categories,omitempty"`
		} `tfsdk:"logging" json:"logging,omitempty"`
		Replicas   *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Scheduling *struct {
			PriorityClassName *string `tfsdk:"priority_class_name" json:"PriorityClassName,omitempty"`
			Affinity          *struct {
				NodeAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						Preference *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchFields *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_fields" json:"matchFields,omitempty"`
						} `tfsdk:"preference" json:"preference,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
					RequiredDuringSchedulingIgnoredDuringExecution *struct {
						NodeSelectorTerms *[]struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchFields *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_fields" json:"matchFields,omitempty"`
						} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
				PodAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						PodAffinityTerm *struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
							TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
					RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
				PodAntiAffinity *struct {
					PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
						PodAffinityTerm *struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
							Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
							TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
					RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
				} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
			} `tfsdk:"affinity" json:"affinity,omitempty"`
		} `tfsdk:"scheduling" json:"scheduling,omitempty"`
		Security *struct {
			Authorization *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				Roles   *[]struct {
					Name        *string   `tfsdk:"name" json:"name,omitempty"`
					Permissions *[]string `tfsdk:"permissions" json:"permissions,omitempty"`
				} `tfsdk:"roles" json:"roles,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			CredentialStoreSecretName *string `tfsdk:"credential_store_secret_name" json:"credentialStoreSecretName,omitempty"`
			EndpointAuthentication    *bool   `tfsdk:"endpoint_authentication" json:"endpointAuthentication,omitempty"`
			EndpointEncryption        *struct {
				CertSecretName       *string `tfsdk:"cert_secret_name" json:"certSecretName,omitempty"`
				CertServiceName      *string `tfsdk:"cert_service_name" json:"certServiceName,omitempty"`
				ClientCert           *string `tfsdk:"client_cert" json:"clientCert,omitempty"`
				ClientCertSecretName *string `tfsdk:"client_cert_secret_name" json:"clientCertSecretName,omitempty"`
				Type                 *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"endpoint_encryption" json:"endpointEncryption,omitempty"`
			EndpointSecretName *string `tfsdk:"endpoint_secret_name" json:"endpointSecretName,omitempty"`
		} `tfsdk:"security" json:"security,omitempty"`
		Service *struct {
			Container *struct {
				EphemeralStorage *bool   `tfsdk:"ephemeral_storage" json:"ephemeralStorage,omitempty"`
				Storage          *string `tfsdk:"storage" json:"storage,omitempty"`
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			} `tfsdk:"container" json:"container,omitempty"`
			ReplicationFactor *int64 `tfsdk:"replication_factor" json:"replicationFactor,omitempty"`
			Sites             *struct {
				Local *struct {
					Discovery *struct {
						Cpu        *string `tfsdk:"cpu" json:"cpu,omitempty"`
						Heartbeats *struct {
							Enabled  *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
							Interval *int64 `tfsdk:"interval" json:"interval,omitempty"`
							Timeout  *int64 `tfsdk:"timeout" json:"timeout,omitempty"`
						} `tfsdk:"heartbeats" json:"heartbeats,omitempty"`
						LaunchGossipRouter *bool   `tfsdk:"launch_gossip_router" json:"launchGossipRouter,omitempty"`
						Memory             *string `tfsdk:"memory" json:"memory,omitempty"`
						SuspectEvents      *bool   `tfsdk:"suspect_events" json:"suspectEvents,omitempty"`
						Type               *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"discovery" json:"discovery,omitempty"`
					Encryption *struct {
						Protocol       *string `tfsdk:"protocol" json:"protocol,omitempty"`
						RouterKeyStore *struct {
							Alias      *string `tfsdk:"alias" json:"alias,omitempty"`
							Filename   *string `tfsdk:"filename" json:"filename,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"router_key_store" json:"routerKeyStore,omitempty"`
						TransportKeyStore *struct {
							Alias      *string `tfsdk:"alias" json:"alias,omitempty"`
							Filename   *string `tfsdk:"filename" json:"filename,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"transport_key_store" json:"transportKeyStore,omitempty"`
						TrustStore *struct {
							Filename   *string `tfsdk:"filename" json:"filename,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"trust_store" json:"trustStore,omitempty"`
					} `tfsdk:"encryption" json:"encryption,omitempty"`
					Expose *struct {
						Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						NodePort      *int64             `tfsdk:"node_port" json:"nodePort,omitempty"`
						Port          *int64             `tfsdk:"port" json:"port,omitempty"`
						RouteHostName *string            `tfsdk:"route_host_name" json:"routeHostName,omitempty"`
						Type          *string            `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"expose" json:"expose,omitempty"`
					MaxRelayNodes *int64  `tfsdk:"max_relay_nodes" json:"maxRelayNodes,omitempty"`
					Name          *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"local" json:"local,omitempty"`
				Locations *[]struct {
					ClusterName *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
					Host        *string `tfsdk:"host" json:"host,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					Url         *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"locations" json:"locations,omitempty"`
			} `tfsdk:"sites" json:"sites,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		Upgrades *struct {
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"upgrades" json:"upgrades,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfinispanOrgInfinispanV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infinispan_org_infinispan_v1"
}

func (r *InfinispanOrgInfinispanV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Infinispan is the Schema for the infinispans API",
		MarkdownDescription: "Infinispan is the Schema for the infinispans API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "InfinispanSpec defines the desired state of Infinispan",
				MarkdownDescription: "InfinispanSpec defines the desired state of Infinispan",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Deprecated. Use scheduling.affinity instead",
						MarkdownDescription: "Deprecated. Use scheduling.affinity instead",
						Attributes: map[string]schema.Attribute{
							"node_affinity": schema.SingleNestedAttribute{
								Description:         "Describes node affinity scheduling rules for the pod.",
								MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"preference": schema.SingleNestedAttribute{
													Description:         "A node selector term, associated with the corresponding weight.",
													MarkdownDescription: "A node selector term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's labels.",
															MarkdownDescription: "A list of node selector requirements by node's labels.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_fields": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"weight": schema.Int64Attribute{
													Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
													MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
										Attributes: map[string]schema.Attribute{
											"node_selector_terms": schema.ListNestedAttribute{
												Description:         "Required. A list of node selector terms. The terms are ORed.",
												MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's labels.",
															MarkdownDescription: "A list of node selector requirements by node's labels.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_fields": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"pod_affinity": schema.SingleNestedAttribute{
								Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Required. A pod affinity term, associated with the corresponding weight.",
													MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
															MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"weight": schema.Int64Attribute{
													Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"namespaces": schema.ListAttribute{
													Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
													MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"pod_anti_affinity": schema.SingleNestedAttribute{
								Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Required. A pod affinity term, associated with the corresponding weight.",
													MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
															MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"weight": schema.Int64Attribute{
													Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"namespaces": schema.ListAttribute{
													Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
													MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"autoscale": schema.SingleNestedAttribute{
						Description:         "Autoscale describe autoscaling configuration for the cluster",
						MarkdownDescription: "Autoscale describe autoscaling configuration for the cluster",
						Attributes: map[string]schema.Attribute{
							"disabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_mem_usage_percent": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_replicas": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_mem_usage_percent": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"min_replicas": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"cloud_events": schema.SingleNestedAttribute{
						Description:         "InfinispanCloudEvents describes how Infinispan is connected with Cloud Event, see Kafka docs for more info",
						MarkdownDescription: "InfinispanCloudEvents describes how Infinispan is connected with Cloud Event, see Kafka docs for more info",
						Attributes: map[string]schema.Attribute{
							"acks": schema.StringAttribute{
								Description:         "Acks configuration for the producer ack-value",
								MarkdownDescription: "Acks configuration for the producer ack-value",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"bootstrap_servers": schema.StringAttribute{
								Description:         "BootstrapServers is comma separated list of boostrap server:port addresses",
								MarkdownDescription: "BootstrapServers is comma separated list of boostrap server:port addresses",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cache_entries_topic": schema.StringAttribute{
								Description:         "CacheEntriesTopic is the name of the topic on which events will be published",
								MarkdownDescription: "CacheEntriesTopic is the name of the topic on which events will be published",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"config_listener": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cpu": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enabled": schema.BoolAttribute{
								Description:         "If true, a dedicated pod is used to ensure that all config resources created on the Infinispan server have a matching CR resource",
								MarkdownDescription: "If true, a dedicated pod is used to ensure that all config resources created on the Infinispan server have a matching CR resource",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"logging": schema.SingleNestedAttribute{
								Description:         "ConfigListener logging configuration",
								MarkdownDescription: "ConfigListener logging configuration",
								Attributes: map[string]schema.Attribute{
									"level": schema.StringAttribute{
										Description:         "The logging level to be used by the ConfigListener pod",
										MarkdownDescription: "The logging level to be used by the ConfigListener pod",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"memory": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"config_map_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"container": schema.SingleNestedAttribute{
						Description:         "InfinispanContainerSpec specify resource requirements per container",
						MarkdownDescription: "InfinispanContainerSpec specify resource requirements per container",
						Attributes: map[string]schema.Attribute{
							"cli_extra_jvm_opts": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"cpu": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"extra_jvm_opts": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"memory": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"router_extra_jvm_opts": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"dependencies": schema.SingleNestedAttribute{
						Description:         "External dependencies needed by the Infinispan cluster",
						MarkdownDescription: "External dependencies needed by the Infinispan cluster",
						Attributes: map[string]schema.Attribute{
							"artifacts": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hash": schema.StringAttribute{
											Description:         "Checksum that you can use to verify downloaded files.",
											MarkdownDescription: "Checksum that you can use to verify downloaded files.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"maven": schema.StringAttribute{
											Description:         "Coordinates of a maven artifact in the 'groupId:artifactId:version' format, for example 'org.postgresql:postgresql:42.3.1'.",
											MarkdownDescription: "Coordinates of a maven artifact in the 'groupId:artifactId:version' format, for example 'org.postgresql:postgresql:42.3.1'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"type": schema.StringAttribute{
											Description:         "Deprecated, no longer has any effect. Specifies the type of file you want to download. If not specified, the file type is automatically determined from the extension.",
											MarkdownDescription: "Deprecated, no longer has any effect. Specifies the type of file you want to download. If not specified, the file type is automatically determined from the extension.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"url": schema.StringAttribute{
											Description:         "URL of the file you want to download.",
											MarkdownDescription: "URL of the file you want to download.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"volume_claim_name": schema.StringAttribute{
								Description:         "The Persistent Volume Claim that holds custom libraries",
								MarkdownDescription: "The Persistent Volume Claim that holds custom libraries",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"expose": schema.SingleNestedAttribute{
						Description:         "ExposeSpec describe how Infinispan will be exposed externally",
						MarkdownDescription: "ExposeSpec describe how Infinispan will be exposed externally",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"host": schema.StringAttribute{
								Description:         "The network hostname for your Infinispan cluster",
								MarkdownDescription: "The network hostname for your Infinispan cluster",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Type specifies different exposition methods for data grid",
								MarkdownDescription: "Type specifies different exposition methods for data grid",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"jmx": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "If true, a JMX endpoint is exposed on the admin service",
								MarkdownDescription: "If true, a JMX endpoint is exposed on the admin service",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"logging": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"categories": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"replicas": schema.Int64Attribute{
						Description:         "The number of nodes in the Infinispan cluster.",
						MarkdownDescription: "The number of nodes in the Infinispan cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scheduling": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"priority_class_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"affinity": schema.SingleNestedAttribute{
								Description:         "Affinity is a group of affinity scheduling rules.",
								MarkdownDescription: "Affinity is a group of affinity scheduling rules.",
								Attributes: map[string]schema.Attribute{
									"node_affinity": schema.SingleNestedAttribute{
										Description:         "Describes node affinity scheduling rules for the pod.",
										MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"preference": schema.SingleNestedAttribute{
															Description:         "A node selector term, associated with the corresponding weight.",
															MarkdownDescription: "A node selector term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's labels.",
																	MarkdownDescription: "A list of node selector requirements by node's labels.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_fields": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's fields.",
																	MarkdownDescription: "A list of node selector requirements by node's fields.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"weight": schema.Int64Attribute{
															Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
															MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
												Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
												MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
												Attributes: map[string]schema.Attribute{
													"node_selector_terms": schema.ListNestedAttribute{
														Description:         "Required. A list of node selector terms. The terms are ORed.",
														MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's labels.",
																	MarkdownDescription: "A list of node selector requirements by node's labels.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_fields": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's fields.",
																	MarkdownDescription: "A list of node selector requirements by node's fields.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"pod_affinity": schema.SingleNestedAttribute{
										Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
										MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "Required. A pod affinity term, associated with the corresponding weight.",
															MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"match_labels": schema.MapAttribute{
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"namespaces": schema.ListAttribute{
																	Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
																	MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"weight": schema.Int64Attribute{
															Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
															MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
															MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"pod_anti_affinity": schema.SingleNestedAttribute{
										Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
										MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "Required. A pod affinity term, associated with the corresponding weight.",
															MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
																	Attributes: map[string]schema.Attribute{
																		"match_expressions": schema.ListNestedAttribute{
																			Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "key is the label key that the selector applies to.",
																						MarkdownDescription: "key is the label key that the selector applies to.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"match_labels": schema.MapAttribute{
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"namespaces": schema.ListAttribute{
																	Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
																	MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"weight": schema.Int64Attribute{
															Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
															MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
															MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"security": schema.SingleNestedAttribute{
						Description:         "InfinispanSecurity info for the user application connection",
						MarkdownDescription: "InfinispanSecurity info for the user application connection",
						Attributes: map[string]schema.Attribute{
							"authorization": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"roles": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"permissions": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"credential_store_secret_name": schema.StringAttribute{
								Description:         "A secret that contains CredentialStore alias and password combinations",
								MarkdownDescription: "A secret that contains CredentialStore alias and password combinations",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"endpoint_authentication": schema.BoolAttribute{
								Description:         "Enable or disable user authentication",
								MarkdownDescription: "Enable or disable user authentication",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"endpoint_encryption": schema.SingleNestedAttribute{
								Description:         "EndpointEncryption configuration",
								MarkdownDescription: "EndpointEncryption configuration",
								Attributes: map[string]schema.Attribute{
									"cert_secret_name": schema.StringAttribute{
										Description:         "The secret that contains TLS certificates",
										MarkdownDescription: "The secret that contains TLS certificates",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"cert_service_name": schema.StringAttribute{
										Description:         "A service that provides TLS certificates",
										MarkdownDescription: "A service that provides TLS certificates",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"client_cert": schema.StringAttribute{
										Description:         "ClientCertType specifies a client certificate validation mechanism.",
										MarkdownDescription: "ClientCertType specifies a client certificate validation mechanism.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"client_cert_secret_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "Disable or modify endpoint encryption.",
										MarkdownDescription: "Disable or modify endpoint encryption.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"endpoint_secret_name": schema.StringAttribute{
								Description:         "The secret that contains user credentials.",
								MarkdownDescription: "The secret that contains user credentials.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "InfinispanServiceSpec specify configuration for specific service",
						MarkdownDescription: "InfinispanServiceSpec specify configuration for specific service",
						Attributes: map[string]schema.Attribute{
							"container": schema.SingleNestedAttribute{
								Description:         "InfinispanServiceContainerSpec resource requirements specific for service",
								MarkdownDescription: "InfinispanServiceContainerSpec resource requirements specific for service",
								Attributes: map[string]schema.Attribute{
									"ephemeral_storage": schema.BoolAttribute{
										Description:         "Enable/disable container ephemeral storage",
										MarkdownDescription: "Enable/disable container ephemeral storage",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"storage": schema.StringAttribute{
										Description:         "The amount of storage for the persistent volume claim.",
										MarkdownDescription: "The amount of storage for the persistent volume claim.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"storage_class_name": schema.StringAttribute{
										Description:         "The storage class object for persistent volume claims",
										MarkdownDescription: "The storage class object for persistent volume claims",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"replication_factor": schema.Int64Attribute{
								Description:         "Cache replication factor, or number of copies for each entry.",
								MarkdownDescription: "Cache replication factor, or number of copies for each entry.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"sites": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"local": schema.SingleNestedAttribute{
										Description:         "InfinispanSitesLocalSpec enables cross-site replication",
										MarkdownDescription: "InfinispanSitesLocalSpec enables cross-site replication",
										Attributes: map[string]schema.Attribute{
											"discovery": schema.SingleNestedAttribute{
												Description:         "DiscoverySiteSpec configures the corss-site replication discovery",
												MarkdownDescription: "DiscoverySiteSpec configures the corss-site replication discovery",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.StringAttribute{
														Description:         "CPU resource request for Gossip Router if enabled",
														MarkdownDescription: "CPU resource request for Gossip Router if enabled",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"heartbeats": schema.SingleNestedAttribute{
														Description:         "Configures the Gossip Router heartbeats to keep the connection open",
														MarkdownDescription: "Configures the Gossip Router heartbeats to keep the connection open",
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Description:         "Enables the Gossip Router heartbeats",
																MarkdownDescription: "Enables the Gossip Router heartbeats",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"interval": schema.Int64Attribute{
																Description:         "Sends a heartbeat to the GossipRouter every interval milliseconds",
																MarkdownDescription: "Sends a heartbeat to the GossipRouter every interval milliseconds",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"timeout": schema.Int64Attribute{
																Description:         "Max time (millsecoonds) with no received message or heartbeat after which the connection to a GossipRouter is closed",
																MarkdownDescription: "Max time (millsecoonds) with no received message or heartbeat after which the connection to a GossipRouter is closed",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"launch_gossip_router": schema.BoolAttribute{
														Description:         "Enables (default) or disables the Gossip Router pod and cross-site services",
														MarkdownDescription: "Enables (default) or disables the Gossip Router pod and cross-site services",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"memory": schema.StringAttribute{
														Description:         "Memory resource request for Gossip Router if enabled",
														MarkdownDescription: "Memory resource request for Gossip Router if enabled",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"suspect_events": schema.BoolAttribute{
														Description:         "Enables the JGroups suspect events if the Gossip Router detects a connection closed",
														MarkdownDescription: "Enables the JGroups suspect events if the Gossip Router detects a connection closed",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "Configures the discovery mode for cross-site replication",
														MarkdownDescription: "Configures the discovery mode for cross-site replication",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"encryption": schema.SingleNestedAttribute{
												Description:         "EncryptionSiteSpec enables TLS for cross-site replication",
												MarkdownDescription: "EncryptionSiteSpec enables TLS for cross-site replication",
												Attributes: map[string]schema.Attribute{
													"protocol": schema.StringAttribute{
														Description:         "TLSProtocol specifies the TLS protocol",
														MarkdownDescription: "TLSProtocol specifies the TLS protocol",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"router_key_store": schema.SingleNestedAttribute{
														Description:         "CrossSiteKeyStore keystore configuration for cross-site replication with TLS",
														MarkdownDescription: "CrossSiteKeyStore keystore configuration for cross-site replication with TLS",
														Attributes: map[string]schema.Attribute{
															"alias": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"filename": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"transport_key_store": schema.SingleNestedAttribute{
														Description:         "CrossSiteKeyStore keystore configuration for cross-site replication with TLS",
														MarkdownDescription: "CrossSiteKeyStore keystore configuration for cross-site replication with TLS",
														Attributes: map[string]schema.Attribute{
															"alias": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"filename": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"trust_store": schema.SingleNestedAttribute{
														Description:         "CrossSiteTrustStore truststore configuration for cross-site replication with TLS",
														MarkdownDescription: "CrossSiteTrustStore truststore configuration for cross-site replication with TLS",
														Attributes: map[string]schema.Attribute{
															"filename": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"secret_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"expose": schema.SingleNestedAttribute{
												Description:         "CrossSiteExposeSpec describe how Infinispan Cross-Site service will be exposed externally",
												MarkdownDescription: "CrossSiteExposeSpec describe how Infinispan Cross-Site service will be exposed externally",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"node_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"route_host_name": schema.StringAttribute{
														Description:         "RouteHostName optionally, specifies a custom hostname to be used by Openshift Route.",
														MarkdownDescription: "RouteHostName optionally, specifies a custom hostname to be used by Openshift Route.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "Type specifies different exposition methods for data grid",
														MarkdownDescription: "Type specifies different exposition methods for data grid",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"max_relay_nodes": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"locations": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"cluster_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"host": schema.StringAttribute{
													Description:         "Deprecated and to be removed on subsequent release. Use .URL with infinispan+xsite schema instead.",
													MarkdownDescription: "Deprecated and to be removed on subsequent release. Use .URL with infinispan+xsite schema instead.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"port": schema.Int64Attribute{
													Description:         "Deprecated and to be removed on subsequent release. Use .URL with infinispan+xsite schema instead.",
													MarkdownDescription: "Deprecated and to be removed on subsequent release. Use .URL with infinispan+xsite schema instead.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret_name": schema.StringAttribute{
													Description:         "The access secret that allows backups to a remote site",
													MarkdownDescription: "The access secret that allows backups to a remote site",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"url": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"type": schema.StringAttribute{
								Description:         "The service type",
								MarkdownDescription: "The service type",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"upgrades": schema.SingleNestedAttribute{
						Description:         "Strategy to use when doing upgrades",
						MarkdownDescription: "Strategy to use when doing upgrades",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"version": schema.StringAttribute{
						Description:         "The semantic version of the Infinispan cluster.",
						MarkdownDescription: "The semantic version of the Infinispan cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *InfinispanOrgInfinispanV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *InfinispanOrgInfinispanV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_infinispan_org_infinispan_v1")

	var data InfinispanOrgInfinispanV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infinispan.org", Version: "v1", Resource: "infinispans"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name and namespace configured.\n\n"+
						"Namespace: %s\n"+
						"Name: %s", data.Metadata.Namespace, data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse InfinispanOrgInfinispanV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("infinispan.org/v1")
	data.Kind = pointer.String("Infinispan")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
