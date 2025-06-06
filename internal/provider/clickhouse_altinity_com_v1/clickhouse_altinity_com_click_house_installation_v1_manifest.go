/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package clickhouse_altinity_com_v1

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
	_ datasource.DataSource = &ClickhouseAltinityComClickHouseInstallationV1Manifest{}
)

func NewClickhouseAltinityComClickHouseInstallationV1Manifest() datasource.DataSource {
	return &ClickhouseAltinityComClickHouseInstallationV1Manifest{}
}

type ClickhouseAltinityComClickHouseInstallationV1Manifest struct{}

type ClickhouseAltinityComClickHouseInstallationV1ManifestData struct {
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
		Configuration *struct {
			Clusters *[]struct {
				Files    *map[string]string `tfsdk:"files" json:"files,omitempty"`
				Insecure *string            `tfsdk:"insecure" json:"insecure,omitempty"`
				Layout   *struct {
					Replicas *[]struct {
						Files    *map[string]string `tfsdk:"files" json:"files,omitempty"`
						Name     *string            `tfsdk:"name" json:"name,omitempty"`
						Settings *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
						Shards   *[]struct {
							Files               *map[string]string `tfsdk:"files" json:"files,omitempty"`
							HttpPort            *int64             `tfsdk:"http_port" json:"httpPort,omitempty"`
							HttpsPort           *int64             `tfsdk:"https_port" json:"httpsPort,omitempty"`
							Insecure            *string            `tfsdk:"insecure" json:"insecure,omitempty"`
							InterserverHTTPPort *int64             `tfsdk:"interserver_http_port" json:"interserverHTTPPort,omitempty"`
							Name                *string            `tfsdk:"name" json:"name,omitempty"`
							Secure              *string            `tfsdk:"secure" json:"secure,omitempty"`
							Settings            *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
							TcpPort             *int64             `tfsdk:"tcp_port" json:"tcpPort,omitempty"`
							Templates           *struct {
								ClusterServiceTemplate  *string   `tfsdk:"cluster_service_template" json:"clusterServiceTemplate,omitempty"`
								DataVolumeClaimTemplate *string   `tfsdk:"data_volume_claim_template" json:"dataVolumeClaimTemplate,omitempty"`
								HostTemplate            *string   `tfsdk:"host_template" json:"hostTemplate,omitempty"`
								LogVolumeClaimTemplate  *string   `tfsdk:"log_volume_claim_template" json:"logVolumeClaimTemplate,omitempty"`
								PodTemplate             *string   `tfsdk:"pod_template" json:"podTemplate,omitempty"`
								ReplicaServiceTemplate  *string   `tfsdk:"replica_service_template" json:"replicaServiceTemplate,omitempty"`
								ServiceTemplate         *string   `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
								ServiceTemplates        *[]string `tfsdk:"service_templates" json:"serviceTemplates,omitempty"`
								ShardServiceTemplate    *string   `tfsdk:"shard_service_template" json:"shardServiceTemplate,omitempty"`
								VolumeClaimTemplate     *string   `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
							} `tfsdk:"templates" json:"templates,omitempty"`
							TlsPort *int64 `tfsdk:"tls_port" json:"tlsPort,omitempty"`
						} `tfsdk:"shards" json:"shards,omitempty"`
						ShardsCount *int64 `tfsdk:"shards_count" json:"shardsCount,omitempty"`
						Templates   *struct {
							ClusterServiceTemplate  *string   `tfsdk:"cluster_service_template" json:"clusterServiceTemplate,omitempty"`
							DataVolumeClaimTemplate *string   `tfsdk:"data_volume_claim_template" json:"dataVolumeClaimTemplate,omitempty"`
							HostTemplate            *string   `tfsdk:"host_template" json:"hostTemplate,omitempty"`
							LogVolumeClaimTemplate  *string   `tfsdk:"log_volume_claim_template" json:"logVolumeClaimTemplate,omitempty"`
							PodTemplate             *string   `tfsdk:"pod_template" json:"podTemplate,omitempty"`
							ReplicaServiceTemplate  *string   `tfsdk:"replica_service_template" json:"replicaServiceTemplate,omitempty"`
							ServiceTemplate         *string   `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
							ServiceTemplates        *[]string `tfsdk:"service_templates" json:"serviceTemplates,omitempty"`
							ShardServiceTemplate    *string   `tfsdk:"shard_service_template" json:"shardServiceTemplate,omitempty"`
							VolumeClaimTemplate     *string   `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
						} `tfsdk:"templates" json:"templates,omitempty"`
					} `tfsdk:"replicas" json:"replicas,omitempty"`
					ReplicasCount *int64 `tfsdk:"replicas_count" json:"replicasCount,omitempty"`
					Shards        *[]struct {
						DefinitionType      *string            `tfsdk:"definition_type" json:"definitionType,omitempty"`
						Files               *map[string]string `tfsdk:"files" json:"files,omitempty"`
						InternalReplication *string            `tfsdk:"internal_replication" json:"internalReplication,omitempty"`
						Name                *string            `tfsdk:"name" json:"name,omitempty"`
						Replicas            *[]struct {
							Files               *map[string]string `tfsdk:"files" json:"files,omitempty"`
							HttpPort            *int64             `tfsdk:"http_port" json:"httpPort,omitempty"`
							HttpsPort           *int64             `tfsdk:"https_port" json:"httpsPort,omitempty"`
							Insecure            *string            `tfsdk:"insecure" json:"insecure,omitempty"`
							InterserverHTTPPort *int64             `tfsdk:"interserver_http_port" json:"interserverHTTPPort,omitempty"`
							Name                *string            `tfsdk:"name" json:"name,omitempty"`
							Secure              *string            `tfsdk:"secure" json:"secure,omitempty"`
							Settings            *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
							TcpPort             *int64             `tfsdk:"tcp_port" json:"tcpPort,omitempty"`
							Templates           *struct {
								ClusterServiceTemplate  *string   `tfsdk:"cluster_service_template" json:"clusterServiceTemplate,omitempty"`
								DataVolumeClaimTemplate *string   `tfsdk:"data_volume_claim_template" json:"dataVolumeClaimTemplate,omitempty"`
								HostTemplate            *string   `tfsdk:"host_template" json:"hostTemplate,omitempty"`
								LogVolumeClaimTemplate  *string   `tfsdk:"log_volume_claim_template" json:"logVolumeClaimTemplate,omitempty"`
								PodTemplate             *string   `tfsdk:"pod_template" json:"podTemplate,omitempty"`
								ReplicaServiceTemplate  *string   `tfsdk:"replica_service_template" json:"replicaServiceTemplate,omitempty"`
								ServiceTemplate         *string   `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
								ServiceTemplates        *[]string `tfsdk:"service_templates" json:"serviceTemplates,omitempty"`
								ShardServiceTemplate    *string   `tfsdk:"shard_service_template" json:"shardServiceTemplate,omitempty"`
								VolumeClaimTemplate     *string   `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
							} `tfsdk:"templates" json:"templates,omitempty"`
							TlsPort *int64 `tfsdk:"tls_port" json:"tlsPort,omitempty"`
						} `tfsdk:"replicas" json:"replicas,omitempty"`
						ReplicasCount *int64             `tfsdk:"replicas_count" json:"replicasCount,omitempty"`
						Settings      *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
						Templates     *struct {
							ClusterServiceTemplate  *string   `tfsdk:"cluster_service_template" json:"clusterServiceTemplate,omitempty"`
							DataVolumeClaimTemplate *string   `tfsdk:"data_volume_claim_template" json:"dataVolumeClaimTemplate,omitempty"`
							HostTemplate            *string   `tfsdk:"host_template" json:"hostTemplate,omitempty"`
							LogVolumeClaimTemplate  *string   `tfsdk:"log_volume_claim_template" json:"logVolumeClaimTemplate,omitempty"`
							PodTemplate             *string   `tfsdk:"pod_template" json:"podTemplate,omitempty"`
							ReplicaServiceTemplate  *string   `tfsdk:"replica_service_template" json:"replicaServiceTemplate,omitempty"`
							ServiceTemplate         *string   `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
							ServiceTemplates        *[]string `tfsdk:"service_templates" json:"serviceTemplates,omitempty"`
							ShardServiceTemplate    *string   `tfsdk:"shard_service_template" json:"shardServiceTemplate,omitempty"`
							VolumeClaimTemplate     *string   `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
						} `tfsdk:"templates" json:"templates,omitempty"`
						Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
					} `tfsdk:"shards" json:"shards,omitempty"`
					ShardsCount *int64 `tfsdk:"shards_count" json:"shardsCount,omitempty"`
				} `tfsdk:"layout" json:"layout,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				PdbMaxUnavailable *int64  `tfsdk:"pdb_max_unavailable" json:"pdbMaxUnavailable,omitempty"`
				Reconcile         *struct {
					Runtime *struct {
						ReconcileShardsMaxConcurrencyPercent *int64 `tfsdk:"reconcile_shards_max_concurrency_percent" json:"reconcileShardsMaxConcurrencyPercent,omitempty"`
						ReconcileShardsThreadsNumber         *int64 `tfsdk:"reconcile_shards_threads_number" json:"reconcileShardsThreadsNumber,omitempty"`
					} `tfsdk:"runtime" json:"runtime,omitempty"`
				} `tfsdk:"reconcile" json:"reconcile,omitempty"`
				SchemaPolicy *struct {
					Replica *string `tfsdk:"replica" json:"replica,omitempty"`
					Shard   *string `tfsdk:"shard" json:"shard,omitempty"`
				} `tfsdk:"schema_policy" json:"schemaPolicy,omitempty"`
				Secret *struct {
					Auto      *string `tfsdk:"auto" json:"auto,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Secure    *string            `tfsdk:"secure" json:"secure,omitempty"`
				Settings  *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
				Templates *struct {
					ClusterServiceTemplate  *string   `tfsdk:"cluster_service_template" json:"clusterServiceTemplate,omitempty"`
					DataVolumeClaimTemplate *string   `tfsdk:"data_volume_claim_template" json:"dataVolumeClaimTemplate,omitempty"`
					HostTemplate            *string   `tfsdk:"host_template" json:"hostTemplate,omitempty"`
					LogVolumeClaimTemplate  *string   `tfsdk:"log_volume_claim_template" json:"logVolumeClaimTemplate,omitempty"`
					PodTemplate             *string   `tfsdk:"pod_template" json:"podTemplate,omitempty"`
					ReplicaServiceTemplate  *string   `tfsdk:"replica_service_template" json:"replicaServiceTemplate,omitempty"`
					ServiceTemplate         *string   `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
					ServiceTemplates        *[]string `tfsdk:"service_templates" json:"serviceTemplates,omitempty"`
					ShardServiceTemplate    *string   `tfsdk:"shard_service_template" json:"shardServiceTemplate,omitempty"`
					VolumeClaimTemplate     *string   `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
				} `tfsdk:"templates" json:"templates,omitempty"`
				Zookeeper *struct {
					Identity *string `tfsdk:"identity" json:"identity,omitempty"`
					Nodes    *[]struct {
						AvailabilityZone *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
						Host             *string `tfsdk:"host" json:"host,omitempty"`
						Port             *int64  `tfsdk:"port" json:"port,omitempty"`
						Secure           *string `tfsdk:"secure" json:"secure,omitempty"`
					} `tfsdk:"nodes" json:"nodes,omitempty"`
					Operation_timeout_ms *int64  `tfsdk:"operation_timeout_ms" json:"operation_timeout_ms,omitempty"`
					Root                 *string `tfsdk:"root" json:"root,omitempty"`
					Session_timeout_ms   *int64  `tfsdk:"session_timeout_ms" json:"session_timeout_ms,omitempty"`
				} `tfsdk:"zookeeper" json:"zookeeper,omitempty"`
			} `tfsdk:"clusters" json:"clusters,omitempty"`
			Files     *map[string]string `tfsdk:"files" json:"files,omitempty"`
			Profiles  *map[string]string `tfsdk:"profiles" json:"profiles,omitempty"`
			Quotas    *map[string]string `tfsdk:"quotas" json:"quotas,omitempty"`
			Settings  *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
			Users     *map[string]string `tfsdk:"users" json:"users,omitempty"`
			Zookeeper *struct {
				Identity *string `tfsdk:"identity" json:"identity,omitempty"`
				Nodes    *[]struct {
					AvailabilityZone *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
					Host             *string `tfsdk:"host" json:"host,omitempty"`
					Port             *int64  `tfsdk:"port" json:"port,omitempty"`
					Secure           *string `tfsdk:"secure" json:"secure,omitempty"`
				} `tfsdk:"nodes" json:"nodes,omitempty"`
				Operation_timeout_ms *int64  `tfsdk:"operation_timeout_ms" json:"operation_timeout_ms,omitempty"`
				Root                 *string `tfsdk:"root" json:"root,omitempty"`
				Session_timeout_ms   *int64  `tfsdk:"session_timeout_ms" json:"session_timeout_ms,omitempty"`
			} `tfsdk:"zookeeper" json:"zookeeper,omitempty"`
		} `tfsdk:"configuration" json:"configuration,omitempty"`
		Defaults *struct {
			DistributedDDL *struct {
				Profile *string `tfsdk:"profile" json:"profile,omitempty"`
			} `tfsdk:"distributed_ddl" json:"distributedDDL,omitempty"`
			ReplicasUseFQDN   *string `tfsdk:"replicas_use_fqdn" json:"replicasUseFQDN,omitempty"`
			StorageManagement *struct {
				Provisioner   *string `tfsdk:"provisioner" json:"provisioner,omitempty"`
				ReclaimPolicy *string `tfsdk:"reclaim_policy" json:"reclaimPolicy,omitempty"`
			} `tfsdk:"storage_management" json:"storageManagement,omitempty"`
			Templates *struct {
				ClusterServiceTemplate  *string   `tfsdk:"cluster_service_template" json:"clusterServiceTemplate,omitempty"`
				DataVolumeClaimTemplate *string   `tfsdk:"data_volume_claim_template" json:"dataVolumeClaimTemplate,omitempty"`
				HostTemplate            *string   `tfsdk:"host_template" json:"hostTemplate,omitempty"`
				LogVolumeClaimTemplate  *string   `tfsdk:"log_volume_claim_template" json:"logVolumeClaimTemplate,omitempty"`
				PodTemplate             *string   `tfsdk:"pod_template" json:"podTemplate,omitempty"`
				ReplicaServiceTemplate  *string   `tfsdk:"replica_service_template" json:"replicaServiceTemplate,omitempty"`
				ServiceTemplate         *string   `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
				ServiceTemplates        *[]string `tfsdk:"service_templates" json:"serviceTemplates,omitempty"`
				ShardServiceTemplate    *string   `tfsdk:"shard_service_template" json:"shardServiceTemplate,omitempty"`
				VolumeClaimTemplate     *string   `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
			} `tfsdk:"templates" json:"templates,omitempty"`
		} `tfsdk:"defaults" json:"defaults,omitempty"`
		NamespaceDomainPattern *string `tfsdk:"namespace_domain_pattern" json:"namespaceDomainPattern,omitempty"`
		Reconciling            *struct {
			Cleanup *struct {
				ReconcileFailedObjects *struct {
					ConfigMap   *string `tfsdk:"config_map" json:"configMap,omitempty"`
					Pvc         *string `tfsdk:"pvc" json:"pvc,omitempty"`
					Service     *string `tfsdk:"service" json:"service,omitempty"`
					StatefulSet *string `tfsdk:"stateful_set" json:"statefulSet,omitempty"`
				} `tfsdk:"reconcile_failed_objects" json:"reconcileFailedObjects,omitempty"`
				UnknownObjects *struct {
					ConfigMap   *string `tfsdk:"config_map" json:"configMap,omitempty"`
					Pvc         *string `tfsdk:"pvc" json:"pvc,omitempty"`
					Service     *string `tfsdk:"service" json:"service,omitempty"`
					StatefulSet *string `tfsdk:"stateful_set" json:"statefulSet,omitempty"`
				} `tfsdk:"unknown_objects" json:"unknownObjects,omitempty"`
			} `tfsdk:"cleanup" json:"cleanup,omitempty"`
			ConfigMapPropagationTimeout *int64  `tfsdk:"config_map_propagation_timeout" json:"configMapPropagationTimeout,omitempty"`
			Policy                      *string `tfsdk:"policy" json:"policy,omitempty"`
			Runtime                     *struct {
				ReconcileShardsMaxConcurrencyPercent *int64 `tfsdk:"reconcile_shards_max_concurrency_percent" json:"reconcileShardsMaxConcurrencyPercent,omitempty"`
				ReconcileShardsThreadsNumber         *int64 `tfsdk:"reconcile_shards_threads_number" json:"reconcileShardsThreadsNumber,omitempty"`
			} `tfsdk:"runtime" json:"runtime,omitempty"`
		} `tfsdk:"reconciling" json:"reconciling,omitempty"`
		Restart   *string `tfsdk:"restart" json:"restart,omitempty"`
		Stop      *string `tfsdk:"stop" json:"stop,omitempty"`
		Suspend   *string `tfsdk:"suspend" json:"suspend,omitempty"`
		TaskID    *string `tfsdk:"task_id" json:"taskID,omitempty"`
		Templates *struct {
			HostTemplates *[]struct {
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				PortDistribution *[]struct {
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"port_distribution" json:"portDistribution,omitempty"`
				Spec *struct {
					Files               *map[string]string `tfsdk:"files" json:"files,omitempty"`
					HttpPort            *int64             `tfsdk:"http_port" json:"httpPort,omitempty"`
					HttpsPort           *int64             `tfsdk:"https_port" json:"httpsPort,omitempty"`
					Insecure            *string            `tfsdk:"insecure" json:"insecure,omitempty"`
					InterserverHTTPPort *int64             `tfsdk:"interserver_http_port" json:"interserverHTTPPort,omitempty"`
					Name                *string            `tfsdk:"name" json:"name,omitempty"`
					Secure              *string            `tfsdk:"secure" json:"secure,omitempty"`
					Settings            *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
					TcpPort             *int64             `tfsdk:"tcp_port" json:"tcpPort,omitempty"`
					Templates           *struct {
						ClusterServiceTemplate  *string   `tfsdk:"cluster_service_template" json:"clusterServiceTemplate,omitempty"`
						DataVolumeClaimTemplate *string   `tfsdk:"data_volume_claim_template" json:"dataVolumeClaimTemplate,omitempty"`
						HostTemplate            *string   `tfsdk:"host_template" json:"hostTemplate,omitempty"`
						LogVolumeClaimTemplate  *string   `tfsdk:"log_volume_claim_template" json:"logVolumeClaimTemplate,omitempty"`
						PodTemplate             *string   `tfsdk:"pod_template" json:"podTemplate,omitempty"`
						ReplicaServiceTemplate  *string   `tfsdk:"replica_service_template" json:"replicaServiceTemplate,omitempty"`
						ServiceTemplate         *string   `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
						ServiceTemplates        *[]string `tfsdk:"service_templates" json:"serviceTemplates,omitempty"`
						ShardServiceTemplate    *string   `tfsdk:"shard_service_template" json:"shardServiceTemplate,omitempty"`
						VolumeClaimTemplate     *string   `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
					} `tfsdk:"templates" json:"templates,omitempty"`
					TlsPort *int64 `tfsdk:"tls_port" json:"tlsPort,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"host_templates" json:"hostTemplates,omitempty"`
			PodTemplates *[]struct {
				Distribution    *string            `tfsdk:"distribution" json:"distribution,omitempty"`
				GenerateName    *string            `tfsdk:"generate_name" json:"generateName,omitempty"`
				Metadata        *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				Name            *string            `tfsdk:"name" json:"name,omitempty"`
				PodDistribution *[]struct {
					Number      *int64  `tfsdk:"number" json:"number,omitempty"`
					Scope       *string `tfsdk:"scope" json:"scope,omitempty"`
					TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					Type        *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"pod_distribution" json:"podDistribution,omitempty"`
				Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				Zone *struct {
					Key    *string   `tfsdk:"key" json:"key,omitempty"`
					Values *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"pod_templates" json:"podTemplates,omitempty"`
			ServiceTemplates *[]struct {
				GenerateName *string            `tfsdk:"generate_name" json:"generateName,omitempty"`
				Metadata     *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				Name         *string            `tfsdk:"name" json:"name,omitempty"`
				Spec         *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"service_templates" json:"serviceTemplates,omitempty"`
			VolumeClaimTemplates *[]struct {
				Metadata      *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				Name          *string            `tfsdk:"name" json:"name,omitempty"`
				Provisioner   *string            `tfsdk:"provisioner" json:"provisioner,omitempty"`
				ReclaimPolicy *string            `tfsdk:"reclaim_policy" json:"reclaimPolicy,omitempty"`
				Spec          *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		} `tfsdk:"templates" json:"templates,omitempty"`
		Templating *struct {
			ChiSelector *map[string]string `tfsdk:"chi_selector" json:"chiSelector,omitempty"`
			Policy      *string            `tfsdk:"policy" json:"policy,omitempty"`
		} `tfsdk:"templating" json:"templating,omitempty"`
		Troubleshoot *string `tfsdk:"troubleshoot" json:"troubleshoot,omitempty"`
		UseTemplates *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			UseType   *string `tfsdk:"use_type" json:"useType,omitempty"`
		} `tfsdk:"use_templates" json:"useTemplates,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClickhouseAltinityComClickHouseInstallationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_clickhouse_altinity_com_click_house_installation_v1_manifest"
}

func (r *ClickhouseAltinityComClickHouseInstallationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "define a set of Kubernetes resources (StatefulSet, PVC, Service, ConfigMap) which describe behavior one or more clusters",
		MarkdownDescription: "define a set of Kubernetes resources (StatefulSet, PVC, Service, ConfigMap) which describe behavior one or more clusters",
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
				Description:         "Specification of the desired behavior of one or more ClickHouse clusters More info: https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md ",
				MarkdownDescription: "Specification of the desired behavior of one or more ClickHouse clusters More info: https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md ",
				Attributes: map[string]schema.Attribute{
					"configuration": schema.SingleNestedAttribute{
						Description:         "allows configure multiple aspects and behavior for 'clickhouse-server' instance and also allows describe multiple 'clickhouse-server' clusters inside one 'chi' resource",
						MarkdownDescription: "allows configure multiple aspects and behavior for 'clickhouse-server' instance and also allows describe multiple 'clickhouse-server' clusters inside one 'chi' resource",
						Attributes: map[string]schema.Attribute{
							"clusters": schema.ListNestedAttribute{
								Description:         "describes clusters layout and allows change settings on cluster-level, shard-level and replica-level every cluster is a set of StatefulSet, one StatefulSet contains only one Pod with 'clickhouse-server' all Pods will rendered in <remote_server> part of ClickHouse configs, mounted from ConfigMap as '/etc/clickhouse-server/config.d/chop-generated-remote_servers.xml' Clusters will use for Distributed table engine, more details: https://clickhouse.tech/docs/en/engines/table-engines/special/distributed/ If 'cluster' contains zookeeper settings (could be inherited from top 'chi' level), when you can create *ReplicatedMergeTree tables ",
								MarkdownDescription: "describes clusters layout and allows change settings on cluster-level, shard-level and replica-level every cluster is a set of StatefulSet, one StatefulSet contains only one Pod with 'clickhouse-server' all Pods will rendered in <remote_server> part of ClickHouse configs, mounted from ConfigMap as '/etc/clickhouse-server/config.d/chop-generated-remote_servers.xml' Clusters will use for Distributed table engine, more details: https://clickhouse.tech/docs/en/engines/table-engines/special/distributed/ If 'cluster' contains zookeeper settings (could be inherited from top 'chi' level), when you can create *ReplicatedMergeTree tables ",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"files": schema.MapAttribute{
											Description:         "optional, allows define content of any setting file inside each 'Pod' on current cluster during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files' ",
											MarkdownDescription: "optional, allows define content of any setting file inside each 'Pod' on current cluster during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files' ",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"insecure": schema.StringAttribute{
											Description:         "optional, open insecure ports for cluster, defaults to 'yes'",
											MarkdownDescription: "optional, open insecure ports for cluster, defaults to 'yes'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
											},
										},

										"layout": schema.SingleNestedAttribute{
											Description:         "describe current cluster layout, how much shards in cluster, how much replica in shard allows override settings on each shard and replica separatelly ",
											MarkdownDescription: "describe current cluster layout, how much shards in cluster, how much replica in shard allows override settings on each shard and replica separatelly ",
											Attributes: map[string]schema.Attribute{
												"replicas": schema.ListNestedAttribute{
													Description:         "optional, allows override top-level 'chi.spec.configuration' and cluster-level 'chi.spec.configuration.clusters' configuration for each replica and each shard relates to selected replica, use it only if you fully understand what you do",
													MarkdownDescription: "optional, allows override top-level 'chi.spec.configuration' and cluster-level 'chi.spec.configuration.clusters' configuration for each replica and each shard relates to selected replica, use it only if you fully understand what you do",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"files": schema.MapAttribute{
																Description:         "optional, allows define content of any setting file inside each 'Pod' only in one replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files' and cluster-level 'chi.spec.configuration.clusters.files', will ignore if 'chi.spec.configuration.clusters.layout.shards' presents ",
																MarkdownDescription: "optional, allows define content of any setting file inside each 'Pod' only in one replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files' and cluster-level 'chi.spec.configuration.clusters.files', will ignore if 'chi.spec.configuration.clusters.layout.shards' presents ",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "optional, by default replica name is generated, but you can override it and setup custom name",
																MarkdownDescription: "optional, by default replica name is generated, but you can override it and setup custom name",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(15),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-]{0,15}$`), ""),
																},
															},

															"settings": schema.MapAttribute{
																Description:         "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in 'Pod' only in one replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' override top-level 'chi.spec.configuration.settings', cluster-level 'chi.spec.configuration.clusters.settings' and will ignore if shard-level 'chi.spec.configuration.clusters.layout.shards' present More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																MarkdownDescription: "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in 'Pod' only in one replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' override top-level 'chi.spec.configuration.settings', cluster-level 'chi.spec.configuration.clusters.settings' and will ignore if shard-level 'chi.spec.configuration.clusters.layout.shards' present More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"shards": schema.ListNestedAttribute{
																Description:         "optional, list of shards related to current replica, will ignore if 'chi.spec.configuration.clusters.layout.shards' presents",
																MarkdownDescription: "optional, list of shards related to current replica, will ignore if 'chi.spec.configuration.clusters.layout.shards' presents",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"files": schema.MapAttribute{
																			Description:         "optional, allows define content of any setting file inside each 'Pod' only in one shard related to current replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files' and cluster-level 'chi.spec.configuration.clusters.files', will ignore if 'chi.spec.configuration.clusters.layout.shards' presents ",
																			MarkdownDescription: "optional, allows define content of any setting file inside each 'Pod' only in one shard related to current replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files' and cluster-level 'chi.spec.configuration.clusters.files', will ignore if 'chi.spec.configuration.clusters.layout.shards' presents ",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_port": schema.Int64Attribute{
																			Description:         "optional, setup 'Pod.spec.containers.ports' with name 'http' for selected shard, override 'chi.spec.templates.hostTemplates.spec.httpPort' allows connect to 'clickhouse-server' via HTTP protocol via kubernetes 'Service' ",
																			MarkdownDescription: "optional, setup 'Pod.spec.containers.ports' with name 'http' for selected shard, override 'chi.spec.templates.hostTemplates.spec.httpPort' allows connect to 'clickhouse-server' via HTTP protocol via kubernetes 'Service' ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"https_port": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"insecure": schema.StringAttribute{
																			Description:         "optional, open insecure ports for cluster, defaults to 'yes' ",
																			MarkdownDescription: "optional, open insecure ports for cluster, defaults to 'yes' ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
																			},
																		},

																		"interserver_http_port": schema.Int64Attribute{
																			Description:         "optional, setup 'Pod.spec.containers.ports' with name 'interserver' for selected shard, override 'chi.spec.templates.hostTemplates.spec.interserverHTTPPort' allows connect between replicas inside same shard during fetch replicated data parts HTTP protocol ",
																			MarkdownDescription: "optional, setup 'Pod.spec.containers.ports' with name 'interserver' for selected shard, override 'chi.spec.templates.hostTemplates.spec.interserverHTTPPort' allows connect between replicas inside same shard during fetch replicated data parts HTTP protocol ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"name": schema.StringAttribute{
																			Description:         "optional, by default shard name is generated, but you can override it and setup custom name",
																			MarkdownDescription: "optional, by default shard name is generated, but you can override it and setup custom name",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(15),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-]{0,15}$`), ""),
																			},
																		},

																		"secure": schema.StringAttribute{
																			Description:         "optional, open secure ports ",
																			MarkdownDescription: "optional, open secure ports ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
																			},
																		},

																		"settings": schema.MapAttribute{
																			Description:         "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in 'Pod' only in one shard related to current replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' override top-level 'chi.spec.configuration.settings', cluster-level 'chi.spec.configuration.clusters.settings' and replica-level 'chi.spec.configuration.clusters.layout.replicas.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																			MarkdownDescription: "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in 'Pod' only in one shard related to current replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' override top-level 'chi.spec.configuration.settings', cluster-level 'chi.spec.configuration.clusters.settings' and replica-level 'chi.spec.configuration.clusters.layout.replicas.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"tcp_port": schema.Int64Attribute{
																			Description:         "optional, setup 'Pod.spec.containers.ports' with name 'tcp' for selected shard, override 'chi.spec.templates.hostTemplates.spec.tcpPort' allows connect to 'clickhouse-server' via TCP Native protocol via kubernetes 'Service' ",
																			MarkdownDescription: "optional, setup 'Pod.spec.containers.ports' with name 'tcp' for selected shard, override 'chi.spec.templates.hostTemplates.spec.tcpPort' allows connect to 'clickhouse-server' via TCP Native protocol via kubernetes 'Service' ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"templates": schema.SingleNestedAttribute{
																			Description:         "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected replica override top-level 'chi.spec.configuration.templates', cluster-level 'chi.spec.configuration.clusters.templates', replica-level 'chi.spec.configuration.clusters.layout.replicas.templates' ",
																			MarkdownDescription: "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected replica override top-level 'chi.spec.configuration.templates', cluster-level 'chi.spec.configuration.clusters.templates', replica-level 'chi.spec.configuration.clusters.layout.replicas.templates' ",
																			Attributes: map[string]schema.Attribute{
																				"cluster_service_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"data_volume_claim_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"host_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"log_volume_claim_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"pod_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replica_service_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"service_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																					MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"service_templates": schema.ListAttribute{
																					Description:         "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																					MarkdownDescription: "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"shard_service_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_claim_template": schema.StringAttribute{
																					Description:         "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"tls_port": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"shards_count": schema.Int64Attribute{
																Description:         "optional, count of shards related to current replica, you can override each shard behavior on low-level 'chi.spec.configuration.clusters.layout.replicas.shards'",
																MarkdownDescription: "optional, count of shards related to current replica, you can override each shard behavior on low-level 'chi.spec.configuration.clusters.layout.replicas.shards'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																},
															},

															"templates": schema.SingleNestedAttribute{
																Description:         "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected replica override top-level 'chi.spec.configuration.templates', cluster-level 'chi.spec.configuration.clusters.templates' ",
																MarkdownDescription: "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected replica override top-level 'chi.spec.configuration.templates', cluster-level 'chi.spec.configuration.clusters.templates' ",
																Attributes: map[string]schema.Attribute{
																	"cluster_service_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"data_volume_claim_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"host_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"log_volume_claim_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pod_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"replica_service_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"service_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																		MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"service_templates": schema.ListAttribute{
																		Description:         "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																		MarkdownDescription: "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"shard_service_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"volume_claim_template": schema.StringAttribute{
																		Description:         "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
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

												"replicas_count": schema.Int64Attribute{
													Description:         "how much replicas in each shards for current cluster will run in Kubernetes, each replica is a separate 'StatefulSet' which contains only one 'Pod' with 'clickhouse-server' instance, every shard contains 1 replica by default' ",
													MarkdownDescription: "how much replicas in each shards for current cluster will run in Kubernetes, each replica is a separate 'StatefulSet' which contains only one 'Pod' with 'clickhouse-server' instance, every shard contains 1 replica by default' ",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"shards": schema.ListNestedAttribute{
													Description:         "optional, allows override top-level 'chi.spec.configuration', cluster-level 'chi.spec.configuration.clusters' settings for each shard separately, use it only if you fully understand what you do' ",
													MarkdownDescription: "optional, allows override top-level 'chi.spec.configuration', cluster-level 'chi.spec.configuration.clusters' settings for each shard separately, use it only if you fully understand what you do' ",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"definition_type": schema.StringAttribute{
																Description:         "DEPRECATED - to be removed soon",
																MarkdownDescription: "DEPRECATED - to be removed soon",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"files": schema.MapAttribute{
																Description:         "optional, allows define content of any setting file inside each 'Pod' only in one shard during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files' and cluster-level 'chi.spec.configuration.clusters.files' ",
																MarkdownDescription: "optional, allows define content of any setting file inside each 'Pod' only in one shard during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files' and cluster-level 'chi.spec.configuration.clusters.files' ",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"internal_replication": schema.StringAttribute{
																Description:         "optional, 'true' by default when 'chi.spec.configuration.clusters[].layout.ReplicaCount' > 1 and 0 otherwise allows setup <internal_replication> setting which will use during insert into tables with 'Distributed' engine for insert only in one live replica and other replicas will download inserted data during replication, will apply in <remote_servers> inside ConfigMap which will mount in /etc/clickhouse-server/config.d/chop-generated-remote_servers.xml More details: https://clickhouse.tech/docs/en/engines/table-engines/special/distributed/ ",
																MarkdownDescription: "optional, 'true' by default when 'chi.spec.configuration.clusters[].layout.ReplicaCount' > 1 and 0 otherwise allows setup <internal_replication> setting which will use during insert into tables with 'Distributed' engine for insert only in one live replica and other replicas will download inserted data during replication, will apply in <remote_servers> inside ConfigMap which will mount in /etc/clickhouse-server/config.d/chop-generated-remote_servers.xml More details: https://clickhouse.tech/docs/en/engines/table-engines/special/distributed/ ",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "optional, by default shard name is generated, but you can override it and setup custom name",
																MarkdownDescription: "optional, by default shard name is generated, but you can override it and setup custom name",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.LengthAtMost(15),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-]{0,15}$`), ""),
																},
															},

															"replicas": schema.ListNestedAttribute{
																Description:         "optional, allows override behavior for selected replicas from cluster-level 'chi.spec.configuration.clusters' and shard-level 'chi.spec.configuration.clusters.layout.shards' ",
																MarkdownDescription: "optional, allows override behavior for selected replicas from cluster-level 'chi.spec.configuration.clusters' and shard-level 'chi.spec.configuration.clusters.layout.shards' ",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"files": schema.MapAttribute{
																			Description:         "optional, allows define content of any setting file inside 'Pod' only in one replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files', cluster-level 'chi.spec.configuration.clusters.files' and shard-level 'chi.spec.configuration.clusters.layout.shards.files' ",
																			MarkdownDescription: "optional, allows define content of any setting file inside 'Pod' only in one replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' override top-level 'chi.spec.configuration.files', cluster-level 'chi.spec.configuration.clusters.files' and shard-level 'chi.spec.configuration.clusters.layout.shards.files' ",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_port": schema.Int64Attribute{
																			Description:         "optional, setup 'Pod.spec.containers.ports' with name 'http' for selected replica, override 'chi.spec.templates.hostTemplates.spec.httpPort' allows connect to 'clickhouse-server' via HTTP protocol via kubernetes 'Service' ",
																			MarkdownDescription: "optional, setup 'Pod.spec.containers.ports' with name 'http' for selected replica, override 'chi.spec.templates.hostTemplates.spec.httpPort' allows connect to 'clickhouse-server' via HTTP protocol via kubernetes 'Service' ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"https_port": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"insecure": schema.StringAttribute{
																			Description:         "optional, open insecure ports for cluster, defaults to 'yes' ",
																			MarkdownDescription: "optional, open insecure ports for cluster, defaults to 'yes' ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
																			},
																		},

																		"interserver_http_port": schema.Int64Attribute{
																			Description:         "optional, setup 'Pod.spec.containers.ports' with name 'interserver' for selected replica, override 'chi.spec.templates.hostTemplates.spec.interserverHTTPPort' allows connect between replicas inside same shard during fetch replicated data parts HTTP protocol ",
																			MarkdownDescription: "optional, setup 'Pod.spec.containers.ports' with name 'interserver' for selected replica, override 'chi.spec.templates.hostTemplates.spec.interserverHTTPPort' allows connect between replicas inside same shard during fetch replicated data parts HTTP protocol ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"name": schema.StringAttribute{
																			Description:         "optional, by default replica name is generated, but you can override it and setup custom name",
																			MarkdownDescription: "optional, by default replica name is generated, but you can override it and setup custom name",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(15),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-]{0,15}$`), ""),
																			},
																		},

																		"secure": schema.StringAttribute{
																			Description:         "optional, open secure ports ",
																			MarkdownDescription: "optional, open secure ports ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
																			},
																		},

																		"settings": schema.MapAttribute{
																			Description:         "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in 'Pod' only in one replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' override top-level 'chi.spec.configuration.settings', cluster-level 'chi.spec.configuration.clusters.settings' and shard-level 'chi.spec.configuration.clusters.layout.shards.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																			MarkdownDescription: "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in 'Pod' only in one replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' override top-level 'chi.spec.configuration.settings', cluster-level 'chi.spec.configuration.clusters.settings' and shard-level 'chi.spec.configuration.clusters.layout.shards.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"tcp_port": schema.Int64Attribute{
																			Description:         "optional, setup 'Pod.spec.containers.ports' with name 'tcp' for selected replica, override 'chi.spec.templates.hostTemplates.spec.tcpPort' allows connect to 'clickhouse-server' via TCP Native protocol via kubernetes 'Service' ",
																			MarkdownDescription: "optional, setup 'Pod.spec.containers.ports' with name 'tcp' for selected replica, override 'chi.spec.templates.hostTemplates.spec.tcpPort' allows connect to 'clickhouse-server' via TCP Native protocol via kubernetes 'Service' ",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},

																		"templates": schema.SingleNestedAttribute{
																			Description:         "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected replica override top-level 'chi.spec.configuration.templates', cluster-level 'chi.spec.configuration.clusters.templates' and shard-level 'chi.spec.configuration.clusters.layout.shards.templates' ",
																			MarkdownDescription: "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected replica override top-level 'chi.spec.configuration.templates', cluster-level 'chi.spec.configuration.clusters.templates' and shard-level 'chi.spec.configuration.clusters.layout.shards.templates' ",
																			Attributes: map[string]schema.Attribute{
																				"cluster_service_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"data_volume_claim_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"host_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"log_volume_claim_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"pod_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"replica_service_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"service_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																					MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"service_templates": schema.ListAttribute{
																					Description:         "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																					MarkdownDescription: "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"shard_service_template": schema.StringAttribute{
																					Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_claim_template": schema.StringAttribute{
																					Description:         "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					MarkdownDescription: "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"tls_port": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.Int64{
																				int64validator.AtLeast(1),
																				int64validator.AtMost(65535),
																			},
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"replicas_count": schema.Int64Attribute{
																Description:         "optional, how much replicas in selected shard for selected ClickHouse cluster will run in Kubernetes, each replica is a separate 'StatefulSet' which contains only one 'Pod' with 'clickhouse-server' instance, shard contains 1 replica by default override cluster-level 'chi.spec.configuration.clusters.layout.replicasCount' ",
																MarkdownDescription: "optional, how much replicas in selected shard for selected ClickHouse cluster will run in Kubernetes, each replica is a separate 'StatefulSet' which contains only one 'Pod' with 'clickhouse-server' instance, shard contains 1 replica by default override cluster-level 'chi.spec.configuration.clusters.layout.replicasCount' ",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(1),
																},
															},

															"settings": schema.MapAttribute{
																Description:         "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' only in one shard during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' override top-level 'chi.spec.configuration.settings' and cluster-level 'chi.spec.configuration.clusters.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																MarkdownDescription: "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' only in one shard during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' override top-level 'chi.spec.configuration.settings' and cluster-level 'chi.spec.configuration.clusters.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"templates": schema.SingleNestedAttribute{
																Description:         "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected shard override top-level 'chi.spec.configuration.templates' and cluster-level 'chi.spec.configuration.clusters.templates' ",
																MarkdownDescription: "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected shard override top-level 'chi.spec.configuration.templates' and cluster-level 'chi.spec.configuration.clusters.templates' ",
																Attributes: map[string]schema.Attribute{
																	"cluster_service_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"data_volume_claim_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"host_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"log_volume_claim_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"pod_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"replica_service_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"service_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																		MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"service_templates": schema.ListAttribute{
																		Description:         "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																		MarkdownDescription: "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"shard_service_template": schema.StringAttribute{
																		Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"volume_claim_template": schema.StringAttribute{
																		Description:         "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		MarkdownDescription: "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"weight": schema.Int64Attribute{
																Description:         "optional, 1 by default, allows setup shard <weight> setting which will use during insert into tables with 'Distributed' engine, will apply in <remote_servers> inside ConfigMap which will mount in /etc/clickhouse-server/config.d/chop-generated-remote_servers.xml More details: https://clickhouse.tech/docs/en/engines/table-engines/special/distributed/ ",
																MarkdownDescription: "optional, 1 by default, allows setup shard <weight> setting which will use during insert into tables with 'Distributed' engine, will apply in <remote_servers> inside ConfigMap which will mount in /etc/clickhouse-server/config.d/chop-generated-remote_servers.xml More details: https://clickhouse.tech/docs/en/engines/table-engines/special/distributed/ ",
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

												"shards_count": schema.Int64Attribute{
													Description:         "how much shards for current ClickHouse cluster will run in Kubernetes, each shard contains shared-nothing part of data and contains set of replicas, cluster contains 1 shard by default' ",
													MarkdownDescription: "how much shards for current ClickHouse cluster will run in Kubernetes, each shard contains shared-nothing part of data and contains set of replicas, cluster contains 1 shard by default' ",
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
											Description:         "cluster name, used to identify set of servers and wide used during generate names of related Kubernetes resources",
											MarkdownDescription: "cluster name, used to identify set of servers and wide used during generate names of related Kubernetes resources",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(15),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-]{0,15}$`), ""),
											},
										},

										"pdb_max_unavailable": schema.Int64Attribute{
											Description:         "Pod eviction is allowed if at most 'pdbMaxUnavailable' pods are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'. ",
											MarkdownDescription: "Pod eviction is allowed if at most 'pdbMaxUnavailable' pods are unavailable after the eviction, i.e. even in absence of the evicted pod. For example, one can prevent all voluntary evictions by specifying 0. This is a mutually exclusive setting with 'minAvailable'. ",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
												int64validator.AtMost(65535),
											},
										},

										"reconcile": schema.SingleNestedAttribute{
											Description:         "allow tuning reconciling process",
											MarkdownDescription: "allow tuning reconciling process",
											Attributes: map[string]schema.Attribute{
												"runtime": schema.SingleNestedAttribute{
													Description:         "runtime parameters for clickhouse-operator process which are used during reconcile cycle",
													MarkdownDescription: "runtime parameters for clickhouse-operator process which are used during reconcile cycle",
													Attributes: map[string]schema.Attribute{
														"reconcile_shards_max_concurrency_percent": schema.Int64Attribute{
															Description:         "The maximum percentage of cluster shards that may be reconciled in parallel, 50 percent by default.",
															MarkdownDescription: "The maximum percentage of cluster shards that may be reconciled in parallel, 50 percent by default.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(0),
																int64validator.AtMost(100),
															},
														},

														"reconcile_shards_threads_number": schema.Int64Attribute{
															Description:         "How many goroutines will be used to reconcile shards of a cluster in parallel, 1 by default",
															MarkdownDescription: "How many goroutines will be used to reconcile shards of a cluster in parallel, 1 by default",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(1),
																int64validator.AtMost(65535),
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

										"schema_policy": schema.SingleNestedAttribute{
											Description:         "describes how schema is propagated within replicas and shards ",
											MarkdownDescription: "describes how schema is propagated within replicas and shards ",
											Attributes: map[string]schema.Attribute{
												"replica": schema.StringAttribute{
													Description:         "how schema is propagated within a replica",
													MarkdownDescription: "how schema is propagated within a replica",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("", "None", "All"),
													},
												},

												"shard": schema.StringAttribute{
													Description:         "how schema is propagated between shards",
													MarkdownDescription: "how schema is propagated between shards",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("", "None", "All", "DistributedTablesOnly"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secret": schema.SingleNestedAttribute{
											Description:         "optional, shared secret value to secure cluster communications",
											MarkdownDescription: "optional, shared secret value to secure cluster communications",
											Attributes: map[string]schema.Attribute{
												"auto": schema.StringAttribute{
													Description:         "Auto-generate shared secret value to secure cluster communications",
													MarkdownDescription: "Auto-generate shared secret value to secure cluster communications",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
													},
												},

												"value": schema.StringAttribute{
													Description:         "Cluster shared secret value in plain text",
													MarkdownDescription: "Cluster shared secret value in plain text",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value_from": schema.SingleNestedAttribute{
													Description:         "Cluster shared secret source",
													MarkdownDescription: "Cluster shared secret source",
													Attributes: map[string]schema.Attribute{
														"secret_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a secret in the clickhouse installation namespace. Should not be used if value is not empty. ",
															MarkdownDescription: "Selects a key of a secret in the clickhouse installation namespace. Should not be used if value is not empty. ",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from. Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names ",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names ",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the Secret or its key must be defined",
																	MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secure": schema.StringAttribute{
											Description:         "optional, open secure ports for cluster",
											MarkdownDescription: "optional, open secure ports for cluster",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
											},
										},

										"settings": schema.MapAttribute{
											Description:         "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' only in one cluster during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' override top-level 'chi.spec.configuration.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
											MarkdownDescription: "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' only in one cluster during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' override top-level 'chi.spec.configuration.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"templates": schema.SingleNestedAttribute{
											Description:         "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected cluster override top-level 'chi.spec.configuration.templates' ",
											MarkdownDescription: "optional, configuration of the templates names which will use for generate Kubernetes resources according to selected cluster override top-level 'chi.spec.configuration.templates' ",
											Attributes: map[string]schema.Attribute{
												"cluster_service_template": schema.StringAttribute{
													Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
													MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"data_volume_claim_template": schema.StringAttribute{
													Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
													MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"host_template": schema.StringAttribute{
													Description:         "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
													MarkdownDescription: "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"log_volume_claim_template": schema.StringAttribute{
													Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
													MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_template": schema.StringAttribute{
													Description:         "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
													MarkdownDescription: "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replica_service_template": schema.StringAttribute{
													Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
													MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service_template": schema.StringAttribute{
													Description:         "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
													MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"service_templates": schema.ListAttribute{
													Description:         "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
													MarkdownDescription: "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"shard_service_template": schema.StringAttribute{
													Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
													MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_claim_template": schema.StringAttribute{
													Description:         "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
													MarkdownDescription: "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"zookeeper": schema.SingleNestedAttribute{
											Description:         "optional, allows configure <yandex><zookeeper>..</zookeeper></yandex> section in each 'Pod' only in current ClickHouse cluster, during generate 'ConfigMap' which will mounted in '/etc/clickhouse-server/config.d/' override top-level 'chi.spec.configuration.zookeeper' settings ",
											MarkdownDescription: "optional, allows configure <yandex><zookeeper>..</zookeeper></yandex> section in each 'Pod' only in current ClickHouse cluster, during generate 'ConfigMap' which will mounted in '/etc/clickhouse-server/config.d/' override top-level 'chi.spec.configuration.zookeeper' settings ",
											Attributes: map[string]schema.Attribute{
												"identity": schema.StringAttribute{
													Description:         "optional access credentials string with 'user:password' format used when use digest authorization in Zookeeper",
													MarkdownDescription: "optional access credentials string with 'user:password' format used when use digest authorization in Zookeeper",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"nodes": schema.ListNestedAttribute{
													Description:         "describe every available zookeeper cluster node for interaction",
													MarkdownDescription: "describe every available zookeeper cluster node for interaction",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"availability_zone": schema.StringAttribute{
																Description:         "availability zone for Zookeeper node",
																MarkdownDescription: "availability zone for Zookeeper node",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"host": schema.StringAttribute{
																Description:         "dns name or ip address for Zookeeper node",
																MarkdownDescription: "dns name or ip address for Zookeeper node",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"port": schema.Int64Attribute{
																Description:         "TCP port which used to connect to Zookeeper node",
																MarkdownDescription: "TCP port which used to connect to Zookeeper node",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.Int64{
																	int64validator.AtLeast(0),
																	int64validator.AtMost(65535),
																},
															},

															"secure": schema.StringAttribute{
																Description:         "if a secure connection to Zookeeper is required",
																MarkdownDescription: "if a secure connection to Zookeeper is required",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"operation_timeout_ms": schema.Int64Attribute{
													Description:         "one operation timeout during Zookeeper transactions",
													MarkdownDescription: "one operation timeout during Zookeeper transactions",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"root": schema.StringAttribute{
													Description:         "optional root znode path inside zookeeper to store ClickHouse related data (replication queue or distributed DDL)",
													MarkdownDescription: "optional root znode path inside zookeeper to store ClickHouse related data (replication queue or distributed DDL)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"session_timeout_ms": schema.Int64Attribute{
													Description:         "session timeout during connect to Zookeeper",
													MarkdownDescription: "session timeout during connect to Zookeeper",
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

							"files": schema.MapAttribute{
								Description:         "allows define content of any setting file inside each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' every key in this object is the file name every value in this object is the file content you can use '!!binary |' and base64 for binary files, see details here https://yaml.org/type/binary.html each key could contains prefix like {users}, {common}, {host} or config.d, users.d, cond.d, wrong prefixes will ignored, subfolders also will ignored More details: https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-05-files-nested.yaml any key could contains 'valueFrom' with 'secretKeyRef' which allow pass values from kubernetes secrets secrets will mounted into pod as separate volume in /etc/clickhouse-server/secrets.d/ and will automatically update when update secret it useful for pass SSL certificates from cert-manager or similar tool look into https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-01-overview.yaml for examples ",
								MarkdownDescription: "allows define content of any setting file inside each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' every key in this object is the file name every value in this object is the file content you can use '!!binary |' and base64 for binary files, see details here https://yaml.org/type/binary.html each key could contains prefix like {users}, {common}, {host} or config.d, users.d, cond.d, wrong prefixes will ignored, subfolders also will ignored More details: https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-05-files-nested.yaml any key could contains 'valueFrom' with 'secretKeyRef' which allow pass values from kubernetes secrets secrets will mounted into pod as separate volume in /etc/clickhouse-server/secrets.d/ and will automatically update when update secret it useful for pass SSL certificates from cert-manager or similar tool look into https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-01-overview.yaml for examples ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"profiles": schema.MapAttribute{
								Description:         "allows configure <yandex><profiles>..</profiles></yandex> section in each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/users.d/' you can configure any aspect of settings profile More details: https://clickhouse.tech/docs/en/operations/settings/settings-profiles/ Your yaml code will convert to XML, see examples https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specconfigurationprofiles ",
								MarkdownDescription: "allows configure <yandex><profiles>..</profiles></yandex> section in each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/users.d/' you can configure any aspect of settings profile More details: https://clickhouse.tech/docs/en/operations/settings/settings-profiles/ Your yaml code will convert to XML, see examples https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specconfigurationprofiles ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"quotas": schema.MapAttribute{
								Description:         "allows configure <yandex><quotas>..</quotas></yandex> section in each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/users.d/' you can configure any aspect of resource quotas More details: https://clickhouse.tech/docs/en/operations/quotas/ Your yaml code will convert to XML, see examples https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specconfigurationquotas ",
								MarkdownDescription: "allows configure <yandex><quotas>..</quotas></yandex> section in each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/users.d/' you can configure any aspect of resource quotas More details: https://clickhouse.tech/docs/en/operations/quotas/ Your yaml code will convert to XML, see examples https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specconfigurationquotas ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"settings": schema.MapAttribute{
								Description:         "allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ Your yaml code will convert to XML, see examples https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specconfigurationsettings any key could contains 'valueFrom' with 'secretKeyRef' which allow pass password from kubernetes secrets look into https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-01-overview.yaml for examples secret value will pass in 'pod.spec.env', and generate with from_env=XXX in XML in /etc/clickhouse-server/config.d/chop-generated-settings.xml it not allow automatically updates when updates 'secret', change spec.taskID for manually trigger reconcile cycle ",
								MarkdownDescription: "allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ Your yaml code will convert to XML, see examples https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specconfigurationsettings any key could contains 'valueFrom' with 'secretKeyRef' which allow pass password from kubernetes secrets look into https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-01-overview.yaml for examples secret value will pass in 'pod.spec.env', and generate with from_env=XXX in XML in /etc/clickhouse-server/config.d/chop-generated-settings.xml it not allow automatically updates when updates 'secret', change spec.taskID for manually trigger reconcile cycle ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"users": schema.MapAttribute{
								Description:         "allows configure <yandex><users>..</users></yandex> section in each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/users.d/' you can configure password hashed, authorization restrictions, database level security row filters etc. More details: https://clickhouse.tech/docs/en/operations/settings/settings-users/ Your yaml code will convert to XML, see examples https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specconfigurationusers any key could contains 'valueFrom' with 'secretKeyRef' which allow pass password from kubernetes secrets secret value will pass in 'pod.spec.containers.evn', and generate with from_env=XXX in XML in /etc/clickhouse-server/users.d/chop-generated-users.xml it not allow automatically updates when updates 'secret', change spec.taskID for manually trigger reconcile cycle look into https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-01-overview.yaml for examples any key with prefix 'k8s_secret_' shall has value with format namespace/secret/key or secret/key in this case value from secret will write directly into XML tag during render *-usersd ConfigMap any key with prefix 'k8s_secret_env' shall has value with format namespace/secret/key or secret/key in this case value from secret will write into environment variable and write to XML tag via from_env=XXX look into https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-01-overview.yaml for examples ",
								MarkdownDescription: "allows configure <yandex><users>..</users></yandex> section in each 'Pod' during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/users.d/' you can configure password hashed, authorization restrictions, database level security row filters etc. More details: https://clickhouse.tech/docs/en/operations/settings/settings-users/ Your yaml code will convert to XML, see examples https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specconfigurationusers any key could contains 'valueFrom' with 'secretKeyRef' which allow pass password from kubernetes secrets secret value will pass in 'pod.spec.containers.evn', and generate with from_env=XXX in XML in /etc/clickhouse-server/users.d/chop-generated-users.xml it not allow automatically updates when updates 'secret', change spec.taskID for manually trigger reconcile cycle look into https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-01-overview.yaml for examples any key with prefix 'k8s_secret_' shall has value with format namespace/secret/key or secret/key in this case value from secret will write directly into XML tag during render *-usersd ConfigMap any key with prefix 'k8s_secret_env' shall has value with format namespace/secret/key or secret/key in this case value from secret will write into environment variable and write to XML tag via from_env=XXX look into https://github.com/Altinity/clickhouse-operator/blob/master/docs/chi-examples/05-settings-01-overview.yaml for examples ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"zookeeper": schema.SingleNestedAttribute{
								Description:         "allows configure <yandex><zookeeper>..</zookeeper></yandex> section in each 'Pod' during generate 'ConfigMap' which will mounted in '/etc/clickhouse-server/config.d/' 'clickhouse-operator' itself doesn't manage Zookeeper, please install Zookeeper separatelly look examples on https://github.com/Altinity/clickhouse-operator/tree/master/deploy/zookeeper/ currently, zookeeper (or clickhouse-keeper replacement) used for *ReplicatedMergeTree table engines and for 'distributed_ddl' More details: https://clickhouse.tech/docs/en/operations/server-configuration-parameters/settings/#server-settings_zookeeper ",
								MarkdownDescription: "allows configure <yandex><zookeeper>..</zookeeper></yandex> section in each 'Pod' during generate 'ConfigMap' which will mounted in '/etc/clickhouse-server/config.d/' 'clickhouse-operator' itself doesn't manage Zookeeper, please install Zookeeper separatelly look examples on https://github.com/Altinity/clickhouse-operator/tree/master/deploy/zookeeper/ currently, zookeeper (or clickhouse-keeper replacement) used for *ReplicatedMergeTree table engines and for 'distributed_ddl' More details: https://clickhouse.tech/docs/en/operations/server-configuration-parameters/settings/#server-settings_zookeeper ",
								Attributes: map[string]schema.Attribute{
									"identity": schema.StringAttribute{
										Description:         "optional access credentials string with 'user:password' format used when use digest authorization in Zookeeper",
										MarkdownDescription: "optional access credentials string with 'user:password' format used when use digest authorization in Zookeeper",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"nodes": schema.ListNestedAttribute{
										Description:         "describe every available zookeeper cluster node for interaction",
										MarkdownDescription: "describe every available zookeeper cluster node for interaction",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"availability_zone": schema.StringAttribute{
													Description:         "availability zone for Zookeeper node",
													MarkdownDescription: "availability zone for Zookeeper node",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"host": schema.StringAttribute{
													Description:         "dns name or ip address for Zookeeper node",
													MarkdownDescription: "dns name or ip address for Zookeeper node",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "TCP port which used to connect to Zookeeper node",
													MarkdownDescription: "TCP port which used to connect to Zookeeper node",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
														int64validator.AtMost(65535),
													},
												},

												"secure": schema.StringAttribute{
													Description:         "if a secure connection to Zookeeper is required",
													MarkdownDescription: "if a secure connection to Zookeeper is required",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"operation_timeout_ms": schema.Int64Attribute{
										Description:         "one operation timeout during Zookeeper transactions",
										MarkdownDescription: "one operation timeout during Zookeeper transactions",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"root": schema.StringAttribute{
										Description:         "optional root znode path inside zookeeper to store ClickHouse related data (replication queue or distributed DDL)",
										MarkdownDescription: "optional root znode path inside zookeeper to store ClickHouse related data (replication queue or distributed DDL)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"session_timeout_ms": schema.Int64Attribute{
										Description:         "session timeout during connect to Zookeeper",
										MarkdownDescription: "session timeout during connect to Zookeeper",
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

					"defaults": schema.SingleNestedAttribute{
						Description:         "define default behavior for whole ClickHouseInstallation, some behavior can be re-define on cluster, shard and replica level More info: https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specdefaults ",
						MarkdownDescription: "define default behavior for whole ClickHouseInstallation, some behavior can be re-define on cluster, shard and replica level More info: https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#specdefaults ",
						Attributes: map[string]schema.Attribute{
							"distributed_ddl": schema.SingleNestedAttribute{
								Description:         "allows change '<yandex><distributed_ddl></distributed_ddl></yandex>' settings More info: https://clickhouse.tech/docs/en/operations/server-configuration-parameters/settings/#server-settings-distributed_ddl ",
								MarkdownDescription: "allows change '<yandex><distributed_ddl></distributed_ddl></yandex>' settings More info: https://clickhouse.tech/docs/en/operations/server-configuration-parameters/settings/#server-settings-distributed_ddl ",
								Attributes: map[string]schema.Attribute{
									"profile": schema.StringAttribute{
										Description:         "Settings from this profile will be used to execute DDL queries",
										MarkdownDescription: "Settings from this profile will be used to execute DDL queries",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas_use_fqdn": schema.StringAttribute{
								Description:         "define should replicas be specified by FQDN in '<host></host>'. In case of 'no' will use short hostname and clickhouse-server will use kubernetes default suffixes for DNS lookup 'no' by default ",
								MarkdownDescription: "define should replicas be specified by FQDN in '<host></host>'. In case of 'no' will use short hostname and clickhouse-server will use kubernetes default suffixes for DNS lookup 'no' by default ",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
								},
							},

							"storage_management": schema.SingleNestedAttribute{
								Description:         "default storage management options",
								MarkdownDescription: "default storage management options",
								Attributes: map[string]schema.Attribute{
									"provisioner": schema.StringAttribute{
										Description:         "defines 'PVC' provisioner - be it StatefulSet or the Operator",
										MarkdownDescription: "defines 'PVC' provisioner - be it StatefulSet or the Operator",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "StatefulSet", "Operator"),
										},
									},

									"reclaim_policy": schema.StringAttribute{
										Description:         "defines behavior of 'PVC' deletion. 'Delete' by default, if 'Retain' specified then 'PVC' will be kept when deleting StatefulSet ",
										MarkdownDescription: "defines behavior of 'PVC' deletion. 'Delete' by default, if 'Retain' specified then 'PVC' will be kept when deleting StatefulSet ",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("", "Retain", "Delete"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"templates": schema.SingleNestedAttribute{
								Description:         "optional, configuration of the templates names which will use for generate Kubernetes resources according to one or more ClickHouse clusters described in current ClickHouseInstallation (chi) resource",
								MarkdownDescription: "optional, configuration of the templates names which will use for generate Kubernetes resources according to one or more ClickHouse clusters described in current ClickHouseInstallation (chi) resource",
								Attributes: map[string]schema.Attribute{
									"cluster_service_template": schema.StringAttribute{
										Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
										MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"data_volume_claim_template": schema.StringAttribute{
										Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
										MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_template": schema.StringAttribute{
										Description:         "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
										MarkdownDescription: "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_volume_claim_template": schema.StringAttribute{
										Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
										MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_template": schema.StringAttribute{
										Description:         "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
										MarkdownDescription: "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replica_service_template": schema.StringAttribute{
										Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
										MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_template": schema.StringAttribute{
										Description:         "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
										MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_templates": schema.ListAttribute{
										Description:         "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
										MarkdownDescription: "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"shard_service_template": schema.StringAttribute{
										Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
										MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_claim_template": schema.StringAttribute{
										Description:         "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
										MarkdownDescription: "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
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

					"namespace_domain_pattern": schema.StringAttribute{
						Description:         "Custom domain pattern which will be used for DNS names of 'Service' or 'Pod'. Typical use scenario - custom cluster domain in Kubernetes cluster Example: %s.svc.my.test ",
						MarkdownDescription: "Custom domain pattern which will be used for DNS names of 'Service' or 'Pod'. Typical use scenario - custom cluster domain in Kubernetes cluster Example: %s.svc.my.test ",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"reconciling": schema.SingleNestedAttribute{
						Description:         "Optional, allows tuning reconciling cycle for ClickhouseInstallation from clickhouse-operator side",
						MarkdownDescription: "Optional, allows tuning reconciling cycle for ClickhouseInstallation from clickhouse-operator side",
						Attributes: map[string]schema.Attribute{
							"cleanup": schema.SingleNestedAttribute{
								Description:         "Optional, defines behavior for cleanup Kubernetes resources during reconcile cycle",
								MarkdownDescription: "Optional, defines behavior for cleanup Kubernetes resources during reconcile cycle",
								Attributes: map[string]schema.Attribute{
									"reconcile_failed_objects": schema.SingleNestedAttribute{
										Description:         "Describes what clickhouse-operator should do with Kubernetes resources which are failed during reconcile. Default behavior is 'Retain'' ",
										MarkdownDescription: "Describes what clickhouse-operator should do with Kubernetes resources which are failed during reconcile. Default behavior is 'Retain'' ",
										Attributes: map[string]schema.Attribute{
											"config_map": schema.StringAttribute{
												Description:         "Behavior policy for failed ConfigMap, 'Retain' by default",
												MarkdownDescription: "Behavior policy for failed ConfigMap, 'Retain' by default",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "Retain", "Delete"),
												},
											},

											"pvc": schema.StringAttribute{
												Description:         "Behavior policy for failed PVC, 'Retain' by default",
												MarkdownDescription: "Behavior policy for failed PVC, 'Retain' by default",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "Retain", "Delete"),
												},
											},

											"service": schema.StringAttribute{
												Description:         "Behavior policy for failed Service, 'Retain' by default",
												MarkdownDescription: "Behavior policy for failed Service, 'Retain' by default",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "Retain", "Delete"),
												},
											},

											"stateful_set": schema.StringAttribute{
												Description:         "Behavior policy for failed StatefulSet, 'Retain' by default",
												MarkdownDescription: "Behavior policy for failed StatefulSet, 'Retain' by default",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "Retain", "Delete"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"unknown_objects": schema.SingleNestedAttribute{
										Description:         "Describes what clickhouse-operator should do with found Kubernetes resources which should be managed by clickhouse-operator, but do not have 'ownerReference' to any currently managed 'ClickHouseInstallation' resource. Default behavior is 'Delete'' ",
										MarkdownDescription: "Describes what clickhouse-operator should do with found Kubernetes resources which should be managed by clickhouse-operator, but do not have 'ownerReference' to any currently managed 'ClickHouseInstallation' resource. Default behavior is 'Delete'' ",
										Attributes: map[string]schema.Attribute{
											"config_map": schema.StringAttribute{
												Description:         "Behavior policy for unknown ConfigMap, 'Delete' by default",
												MarkdownDescription: "Behavior policy for unknown ConfigMap, 'Delete' by default",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "Retain", "Delete"),
												},
											},

											"pvc": schema.StringAttribute{
												Description:         "Behavior policy for unknown PVC, 'Delete' by default",
												MarkdownDescription: "Behavior policy for unknown PVC, 'Delete' by default",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "Retain", "Delete"),
												},
											},

											"service": schema.StringAttribute{
												Description:         "Behavior policy for unknown Service, 'Delete' by default",
												MarkdownDescription: "Behavior policy for unknown Service, 'Delete' by default",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "Retain", "Delete"),
												},
											},

											"stateful_set": schema.StringAttribute{
												Description:         "Behavior policy for unknown StatefulSet, 'Delete' by default",
												MarkdownDescription: "Behavior policy for unknown StatefulSet, 'Delete' by default",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("", "Retain", "Delete"),
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

							"config_map_propagation_timeout": schema.Int64Attribute{
								Description:         "Timeout in seconds for 'clickhouse-operator' to wait for modified 'ConfigMap' to propagate into the 'Pod' More details: https://kubernetes.io/docs/concepts/configuration/configmap/#mounted-configmaps-are-updated-automatically ",
								MarkdownDescription: "Timeout in seconds for 'clickhouse-operator' to wait for modified 'ConfigMap' to propagate into the 'Pod' More details: https://kubernetes.io/docs/concepts/configuration/configmap/#mounted-configmaps-are-updated-automatically ",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(3600),
								},
							},

							"policy": schema.StringAttribute{
								Description:         "DISCUSSED TO BE DEPRECATED Syntax sugar Overrides all three 'reconcile.host.wait.{exclude, queries, include}' values from the operator's config Possible values: - wait - should wait to exclude host, complete queries and include host back into the cluster - nowait - should NOT wait to exclude host, complete queries and include host back into the cluster ",
								MarkdownDescription: "DISCUSSED TO BE DEPRECATED Syntax sugar Overrides all three 'reconcile.host.wait.{exclude, queries, include}' values from the operator's config Possible values: - wait - should wait to exclude host, complete queries and include host back into the cluster - nowait - should NOT wait to exclude host, complete queries and include host back into the cluster ",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "wait", "nowait"),
								},
							},

							"runtime": schema.SingleNestedAttribute{
								Description:         "runtime parameters for clickhouse-operator process which are used during reconcile cycle",
								MarkdownDescription: "runtime parameters for clickhouse-operator process which are used during reconcile cycle",
								Attributes: map[string]schema.Attribute{
									"reconcile_shards_max_concurrency_percent": schema.Int64Attribute{
										Description:         "The maximum percentage of cluster shards that may be reconciled in parallel, 50 percent by default.",
										MarkdownDescription: "The maximum percentage of cluster shards that may be reconciled in parallel, 50 percent by default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(100),
										},
									},

									"reconcile_shards_threads_number": schema.Int64Attribute{
										Description:         "How many goroutines will be used to reconcile shards of a cluster in parallel, 1 by default",
										MarkdownDescription: "How many goroutines will be used to reconcile shards of a cluster in parallel, 1 by default",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(65535),
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

					"restart": schema.StringAttribute{
						Description:         "In case 'RollingUpdate' specified, the operator will always restart ClickHouse pods during reconcile. This options is used in rare cases when force restart is required and is typically removed after the use in order to avoid unneeded restarts. ",
						MarkdownDescription: "In case 'RollingUpdate' specified, the operator will always restart ClickHouse pods during reconcile. This options is used in rare cases when force restart is required and is typically removed after the use in order to avoid unneeded restarts. ",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "RollingUpdate"),
						},
					},

					"stop": schema.StringAttribute{
						Description:         "Allows to stop all ClickHouse clusters defined in a CHI. Works as the following: - When 'stop' is '1' operator sets 'Replicas: 0' in each StatefulSet. Thie leads to having all 'Pods' and 'Service' deleted. All PVCs are kept intact. - When 'stop' is '0' operator sets 'Replicas: 1' and 'Pod's and 'Service's will created again and all retained PVCs will be attached to 'Pod's. ",
						MarkdownDescription: "Allows to stop all ClickHouse clusters defined in a CHI. Works as the following: - When 'stop' is '1' operator sets 'Replicas: 0' in each StatefulSet. Thie leads to having all 'Pods' and 'Service' deleted. All PVCs are kept intact. - When 'stop' is '0' operator sets 'Replicas: 1' and 'Pod's and 'Service's will created again and all retained PVCs will be attached to 'Pod's. ",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
						},
					},

					"suspend": schema.StringAttribute{
						Description:         "Suspend reconciliation of resources managed by a ClickHouse Installation. Works as the following: - When 'suspend' is 'true' operator stops reconciling all resources. - When 'suspend' is 'false' or not set, operator reconciles all resources. ",
						MarkdownDescription: "Suspend reconciliation of resources managed by a ClickHouse Installation. Works as the following: - When 'suspend' is 'true' operator stops reconciling all resources. - When 'suspend' is 'false' or not set, operator reconciles all resources. ",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
						},
					},

					"task_id": schema.StringAttribute{
						Description:         "Allows to define custom taskID for CHI update and watch status of this update execution. Displayed in all .status.taskID* fields. By default (if not filled) every update of CHI manifest will generate random taskID ",
						MarkdownDescription: "Allows to define custom taskID for CHI update and watch status of this update execution. Displayed in all .status.taskID* fields. By default (if not filled) every update of CHI manifest will generate random taskID ",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"templates": schema.SingleNestedAttribute{
						Description:         "allows define templates which will use for render Kubernetes resources like StatefulSet, ConfigMap, Service, PVC, by default, clickhouse-operator have own templates, but you can override it",
						MarkdownDescription: "allows define templates which will use for render Kubernetes resources like StatefulSet, ConfigMap, Service, PVC, by default, clickhouse-operator have own templates, but you can override it",
						Attributes: map[string]schema.Attribute{
							"host_templates": schema.ListNestedAttribute{
								Description:         "hostTemplate will use during apply to generate 'clickhose-server' config files",
								MarkdownDescription: "hostTemplate will use during apply to generate 'clickhose-server' config files",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "template name, could use to link inside top-level 'chi.spec.defaults.templates.hostTemplate', cluster-level 'chi.spec.configuration.clusters.templates.hostTemplate', shard-level 'chi.spec.configuration.clusters.layout.shards.temlates.hostTemplate', replica-level 'chi.spec.configuration.clusters.layout.replicas.templates.hostTemplate'",
											MarkdownDescription: "template name, could use to link inside top-level 'chi.spec.defaults.templates.hostTemplate', cluster-level 'chi.spec.configuration.clusters.templates.hostTemplate', shard-level 'chi.spec.configuration.clusters.layout.shards.temlates.hostTemplate', replica-level 'chi.spec.configuration.clusters.layout.replicas.templates.hostTemplate'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"port_distribution": schema.ListNestedAttribute{
											Description:         "define how will distribute numeric values of named ports in 'Pod.spec.containers.ports' and clickhouse-server configs",
											MarkdownDescription: "define how will distribute numeric values of named ports in 'Pod.spec.containers.ports' and clickhouse-server configs",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														Description:         "type of distribution, when 'Unspecified' (default value) then all listen ports on clickhouse-server configuration in all Pods will have the same value, when 'ClusterScopeIndex' then ports will increment to offset from base value depends on shard and replica index inside cluster with combination of 'chi.spec.templates.podTemlates.spec.HostNetwork' it allows setup ClickHouse cluster inside Kubernetes and provide access via external network bypass Kubernetes internal network",
														MarkdownDescription: "type of distribution, when 'Unspecified' (default value) then all listen ports on clickhouse-server configuration in all Pods will have the same value, when 'ClusterScopeIndex' then ports will increment to offset from base value depends on shard and replica index inside cluster with combination of 'chi.spec.templates.podTemlates.spec.HostNetwork' it allows setup ClickHouse cluster inside Kubernetes and provide access via external network bypass Kubernetes internal network",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("", "Unspecified", "ClusterScopeIndex"),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"spec": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"files": schema.MapAttribute{
													Description:         "optional, allows define content of any setting file inside each 'Pod' where this template will apply during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' ",
													MarkdownDescription: "optional, allows define content of any setting file inside each 'Pod' where this template will apply during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/config.d/' or '/etc/clickhouse-server/conf.d/' or '/etc/clickhouse-server/users.d/' ",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_port": schema.Int64Attribute{
													Description:         "optional, setup 'http_port' inside 'clickhouse-server' settings for each Pod where current template will apply if specified, should have equal value with 'chi.spec.templates.podTemplates.spec.containers.ports[name=http]' More info: https://clickhouse.tech/docs/en/interfaces/http/ ",
													MarkdownDescription: "optional, setup 'http_port' inside 'clickhouse-server' settings for each Pod where current template will apply if specified, should have equal value with 'chi.spec.templates.podTemplates.spec.containers.ports[name=http]' More info: https://clickhouse.tech/docs/en/interfaces/http/ ",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(65535),
													},
												},

												"https_port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(65535),
													},
												},

												"insecure": schema.StringAttribute{
													Description:         "optional, open insecure ports for cluster, defaults to 'yes' ",
													MarkdownDescription: "optional, open insecure ports for cluster, defaults to 'yes' ",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
													},
												},

												"interserver_http_port": schema.Int64Attribute{
													Description:         "optional, setup 'interserver_http_port' inside 'clickhouse-server' settings for each Pod where current template will apply if specified, should have equal value with 'chi.spec.templates.podTemplates.spec.containers.ports[name=interserver]' More info: https://clickhouse.tech/docs/en/operations/server-configuration-parameters/settings/#interserver-http-port ",
													MarkdownDescription: "optional, setup 'interserver_http_port' inside 'clickhouse-server' settings for each Pod where current template will apply if specified, should have equal value with 'chi.spec.templates.podTemplates.spec.containers.ports[name=interserver]' More info: https://clickhouse.tech/docs/en/operations/server-configuration-parameters/settings/#interserver-http-port ",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(65535),
													},
												},

												"name": schema.StringAttribute{
													Description:         "by default, hostname will generate, but this allows define custom name for each 'clickhuse-server'",
													MarkdownDescription: "by default, hostname will generate, but this allows define custom name for each 'clickhuse-server'",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.LengthAtMost(15),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9-]{0,15}$`), ""),
													},
												},

												"secure": schema.StringAttribute{
													Description:         "optional, open secure ports ",
													MarkdownDescription: "optional, open secure ports ",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
													},
												},

												"settings": schema.MapAttribute{
													Description:         "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' where this template will apply during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
													MarkdownDescription: "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' where this template will apply during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tcp_port": schema.Int64Attribute{
													Description:         "optional, setup 'tcp_port' inside 'clickhouse-server' settings for each Pod where current template will apply if specified, should have equal value with 'chi.spec.templates.podTemplates.spec.containers.ports[name=tcp]' More info: https://clickhouse.tech/docs/en/interfaces/tcp/ ",
													MarkdownDescription: "optional, setup 'tcp_port' inside 'clickhouse-server' settings for each Pod where current template will apply if specified, should have equal value with 'chi.spec.templates.podTemplates.spec.containers.ports[name=tcp]' More info: https://clickhouse.tech/docs/en/interfaces/tcp/ ",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(65535),
													},
												},

												"templates": schema.SingleNestedAttribute{
													Description:         "be careful, this part of CRD allows override template inside template, don't use it if you don't understand what you do",
													MarkdownDescription: "be careful, this part of CRD allows override template inside template, don't use it if you don't understand what you do",
													Attributes: map[string]schema.Attribute{
														"cluster_service_template": schema.StringAttribute{
															Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
															MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each clickhouse cluster described in 'chi.spec.configuration.clusters'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"data_volume_claim_template": schema.StringAttribute{
															Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
															MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"host_template": schema.StringAttribute{
															Description:         "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
															MarkdownDescription: "optional, template name from chi.spec.templates.hostTemplates, which will apply to configure every 'clickhouse-server' instance during render ConfigMap resources which will mount into 'Pod'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"log_volume_claim_template": schema.StringAttribute{
															Description:         "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
															MarkdownDescription: "optional, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse log directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pod_template": schema.StringAttribute{
															Description:         "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
															MarkdownDescription: "optional, template name from chi.spec.templates.podTemplates, allows customization each 'Pod' resource during render and reconcile each StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"replica_service_template": schema.StringAttribute{
															Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
															MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each replica inside each shard inside each clickhouse cluster described in 'chi.spec.configuration.clusters'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_template": schema.StringAttribute{
															Description:         "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
															MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resource, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_templates": schema.ListAttribute{
															Description:         "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
															MarkdownDescription: "optional, template names from chi.spec.templates.serviceTemplates. used for customization of the 'Service' resources, created by 'clickhouse-operator' to cover all clusters in whole 'chi' resource",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"shard_service_template": schema.StringAttribute{
															Description:         "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
															MarkdownDescription: "optional, template name from chi.spec.templates.serviceTemplates, allows customization for each 'Service' resource which will created by 'clickhouse-operator' which cover each shard inside clickhouse cluster described in 'chi.spec.configuration.clusters'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_claim_template": schema.StringAttribute{
															Description:         "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
															MarkdownDescription: "optional, alias for dataVolumeClaimTemplate, template name from chi.spec.templates.volumeClaimTemplates, allows customization each 'PVC' which will mount for clickhouse data directory in each 'Pod' during render and reconcile every StatefulSet.spec resource described in 'chi.spec.configuration.clusters'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"tls_port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(65535),
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

							"pod_templates": schema.ListNestedAttribute{
								Description:         "podTemplate will use during render 'Pod' inside 'StatefulSet.spec' and allows define rendered 'Pod.spec', pod scheduling distribution and pod zone More information: https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#spectemplatespodtemplates ",
								MarkdownDescription: "podTemplate will use during render 'Pod' inside 'StatefulSet.spec' and allows define rendered 'Pod.spec', pod scheduling distribution and pod zone More information: https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#spectemplatespodtemplates ",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"distribution": schema.StringAttribute{
											Description:         "DEPRECATED, shortcut for 'chi.spec.templates.podTemplates.spec.affinity.podAntiAffinity'",
											MarkdownDescription: "DEPRECATED, shortcut for 'chi.spec.templates.podTemplates.spec.affinity.podAntiAffinity'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("", "Unspecified", "OnePerHost"),
											},
										},

										"generate_name": schema.StringAttribute{
											Description:         "allows define format for generated 'Pod' name, look to https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#spectemplatesservicetemplates for details about available template variables",
											MarkdownDescription: "allows define format for generated 'Pod' name, look to https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#spectemplatesservicetemplates for details about available template variables",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metadata": schema.MapAttribute{
											Description:         "allows pass standard object's metadata from template to Pod More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata ",
											MarkdownDescription: "allows pass standard object's metadata from template to Pod More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata ",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "template name, could use to link inside top-level 'chi.spec.defaults.templates.podTemplate', cluster-level 'chi.spec.configuration.clusters.templates.podTemplate', shard-level 'chi.spec.configuration.clusters.layout.shards.temlates.podTemplate', replica-level 'chi.spec.configuration.clusters.layout.replicas.templates.podTemplate'",
											MarkdownDescription: "template name, could use to link inside top-level 'chi.spec.defaults.templates.podTemplate', cluster-level 'chi.spec.configuration.clusters.templates.podTemplate', shard-level 'chi.spec.configuration.clusters.layout.shards.temlates.podTemplate', replica-level 'chi.spec.configuration.clusters.layout.replicas.templates.podTemplate'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pod_distribution": schema.ListNestedAttribute{
											Description:         "define ClickHouse Pod distribution policy between Kubernetes Nodes inside Shard, Replica, Namespace, CHI, another ClickHouse cluster",
											MarkdownDescription: "define ClickHouse Pod distribution policy between Kubernetes Nodes inside Shard, Replica, Namespace, CHI, another ClickHouse cluster",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"number": schema.Int64Attribute{
														Description:         "define, how much ClickHouse Pods could be inside selected scope with selected distribution type",
														MarkdownDescription: "define, how much ClickHouse Pods could be inside selected scope with selected distribution type",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(65535),
														},
													},

													"scope": schema.StringAttribute{
														Description:         "scope for apply each podDistribution",
														MarkdownDescription: "scope for apply each podDistribution",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("", "Unspecified", "Shard", "Replica", "Cluster", "ClickHouseInstallation", "Namespace"),
														},
													},

													"topology_key": schema.StringAttribute{
														Description:         "use for inter-pod affinity look to 'pod.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.topologyKey', more info: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity' ",
														MarkdownDescription: "use for inter-pod affinity look to 'pod.spec.affinity.podAntiAffinity.preferredDuringSchedulingIgnoredDuringExecution.podAffinityTerm.topologyKey', more info: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity' ",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "you can define multiple affinity policy types",
														MarkdownDescription: "you can define multiple affinity policy types",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("", "Unspecified", "ClickHouseAntiAffinity", "ShardAntiAffinity", "ReplicaAntiAffinity", "AnotherNamespaceAntiAffinity", "AnotherClickHouseInstallationAntiAffinity", "AnotherClusterAntiAffinity", "MaxNumberPerNode", "NamespaceAffinity", "ClickHouseInstallationAffinity", "ClusterAffinity", "ShardAffinity", "ReplicaAffinity", "PreviousTailAffinity", "CircularReplication"),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"spec": schema.MapAttribute{
											Description:         "allows define whole Pod.spec inside StaefulSet.spec, look to https://kubernetes.io/docs/concepts/workloads/pods/#pod-templates for details",
											MarkdownDescription: "allows define whole Pod.spec inside StaefulSet.spec, look to https://kubernetes.io/docs/concepts/workloads/pods/#pod-templates for details",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"zone": schema.SingleNestedAttribute{
											Description:         "allows define custom zone name and will separate ClickHouse 'Pods' between nodes, shortcut for 'chi.spec.templates.podTemplates.spec.affinity.podAntiAffinity'",
											MarkdownDescription: "allows define custom zone name and will separate ClickHouse 'Pods' between nodes, shortcut for 'chi.spec.templates.podTemplates.spec.affinity.podAntiAffinity'",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "optional, if defined, allows select kubernetes nodes by label with 'name' equal 'key'",
													MarkdownDescription: "optional, if defined, allows select kubernetes nodes by label with 'name' equal 'key'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "optional, if defined, allows select kubernetes nodes by label with 'value' in 'values'",
													MarkdownDescription: "optional, if defined, allows select kubernetes nodes by label with 'value' in 'values'",
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

							"service_templates": schema.ListNestedAttribute{
								Description:         "allows define template for rendering 'Service' which would get endpoint from Pods which scoped chi-wide, cluster-wide, shard-wide, replica-wide level ",
								MarkdownDescription: "allows define template for rendering 'Service' which would get endpoint from Pods which scoped chi-wide, cluster-wide, shard-wide, replica-wide level ",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"generate_name": schema.StringAttribute{
											Description:         "allows define format for generated 'Service' name, look to https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#spectemplatesservicetemplates for details about available template variables' ",
											MarkdownDescription: "allows define format for generated 'Service' name, look to https://github.com/Altinity/clickhouse-operator/blob/master/docs/custom_resource_explained.md#spectemplatesservicetemplates for details about available template variables' ",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"metadata": schema.MapAttribute{
											Description:         "allows pass standard object's metadata from template to Service Could be use for define specificly for Cloud Provider metadata which impact to behavior of service More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata ",
											MarkdownDescription: "allows pass standard object's metadata from template to Service Could be use for define specificly for Cloud Provider metadata which impact to behavior of service More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata ",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "template name, could use to link inside chi-level 'chi.spec.defaults.templates.serviceTemplate' cluster-level 'chi.spec.configuration.clusters.templates.clusterServiceTemplate' shard-level 'chi.spec.configuration.clusters.layout.shards.temlates.shardServiceTemplate' replica-level 'chi.spec.configuration.clusters.layout.replicas.templates.replicaServiceTemplate' or 'chi.spec.configuration.clusters.layout.shards.replicas.replicaServiceTemplate' ",
											MarkdownDescription: "template name, could use to link inside chi-level 'chi.spec.defaults.templates.serviceTemplate' cluster-level 'chi.spec.configuration.clusters.templates.clusterServiceTemplate' shard-level 'chi.spec.configuration.clusters.layout.shards.temlates.shardServiceTemplate' replica-level 'chi.spec.configuration.clusters.layout.replicas.templates.replicaServiceTemplate' or 'chi.spec.configuration.clusters.layout.shards.replicas.replicaServiceTemplate' ",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"spec": schema.MapAttribute{
											Description:         "describe behavior of generated Service More info: https://kubernetes.io/docs/concepts/services-networking/service/ ",
											MarkdownDescription: "describe behavior of generated Service More info: https://kubernetes.io/docs/concepts/services-networking/service/ ",
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

							"volume_claim_templates": schema.ListNestedAttribute{
								Description:         "allows define template for rendering 'PVC' kubernetes resource, which would use inside 'Pod' for mount clickhouse 'data', clickhouse 'logs' or something else ",
								MarkdownDescription: "allows define template for rendering 'PVC' kubernetes resource, which would use inside 'Pod' for mount clickhouse 'data', clickhouse 'logs' or something else ",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.MapAttribute{
											Description:         "allows to pass standard object's metadata from template to PVC More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata ",
											MarkdownDescription: "allows to pass standard object's metadata from template to PVC More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata ",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "template name, could use to link inside top-level 'chi.spec.defaults.templates.dataVolumeClaimTemplate' or 'chi.spec.defaults.templates.logVolumeClaimTemplate', cluster-level 'chi.spec.configuration.clusters.templates.dataVolumeClaimTemplate' or 'chi.spec.configuration.clusters.templates.logVolumeClaimTemplate', shard-level 'chi.spec.configuration.clusters.layout.shards.temlates.dataVolumeClaimTemplate' or 'chi.spec.configuration.clusters.layout.shards.temlates.logVolumeClaimTemplate' replica-level 'chi.spec.configuration.clusters.layout.replicas.templates.dataVolumeClaimTemplate' or 'chi.spec.configuration.clusters.layout.replicas.templates.logVolumeClaimTemplate' ",
											MarkdownDescription: "template name, could use to link inside top-level 'chi.spec.defaults.templates.dataVolumeClaimTemplate' or 'chi.spec.defaults.templates.logVolumeClaimTemplate', cluster-level 'chi.spec.configuration.clusters.templates.dataVolumeClaimTemplate' or 'chi.spec.configuration.clusters.templates.logVolumeClaimTemplate', shard-level 'chi.spec.configuration.clusters.layout.shards.temlates.dataVolumeClaimTemplate' or 'chi.spec.configuration.clusters.layout.shards.temlates.logVolumeClaimTemplate' replica-level 'chi.spec.configuration.clusters.layout.replicas.templates.dataVolumeClaimTemplate' or 'chi.spec.configuration.clusters.layout.replicas.templates.logVolumeClaimTemplate' ",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"provisioner": schema.StringAttribute{
											Description:         "defines 'PVC' provisioner - be it StatefulSet or the Operator",
											MarkdownDescription: "defines 'PVC' provisioner - be it StatefulSet or the Operator",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("", "StatefulSet", "Operator"),
											},
										},

										"reclaim_policy": schema.StringAttribute{
											Description:         "defines behavior of 'PVC' deletion. 'Delete' by default, if 'Retain' specified then 'PVC' will be kept when deleting StatefulSet ",
											MarkdownDescription: "defines behavior of 'PVC' deletion. 'Delete' by default, if 'Retain' specified then 'PVC' will be kept when deleting StatefulSet ",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("", "Retain", "Delete"),
											},
										},

										"spec": schema.MapAttribute{
											Description:         "allows define all aspects of 'PVC' resource More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims ",
											MarkdownDescription: "allows define all aspects of 'PVC' resource More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#persistentvolumeclaims ",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"templating": schema.SingleNestedAttribute{
						Description:         "Optional, applicable inside ClickHouseInstallationTemplate only. Defines current ClickHouseInstallationTemplate application options to target ClickHouseInstallation(s).' ",
						MarkdownDescription: "Optional, applicable inside ClickHouseInstallationTemplate only. Defines current ClickHouseInstallationTemplate application options to target ClickHouseInstallation(s).' ",
						Attributes: map[string]schema.Attribute{
							"chi_selector": schema.MapAttribute{
								Description:         "Optional, defines selector for ClickHouseInstallation(s) to be templated with ClickhouseInstallationTemplate",
								MarkdownDescription: "Optional, defines selector for ClickHouseInstallation(s) to be templated with ClickhouseInstallationTemplate",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"policy": schema.StringAttribute{
								Description:         "When defined as 'auto' inside ClickhouseInstallationTemplate, this ClickhouseInstallationTemplate will be auto-added into ClickHouseInstallation, selectable by 'chiSelector'. Default value is 'manual', meaning ClickHouseInstallation should request this ClickhouseInstallationTemplate explicitly. ",
								MarkdownDescription: "When defined as 'auto' inside ClickhouseInstallationTemplate, this ClickhouseInstallationTemplate will be auto-added into ClickHouseInstallation, selectable by 'chiSelector'. Default value is 'manual', meaning ClickHouseInstallation should request this ClickhouseInstallationTemplate explicitly. ",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "auto", "manual"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"troubleshoot": schema.StringAttribute{
						Description:         "Allows to troubleshoot Pods during CrashLoopBack state. This may happen when wrong configuration applied, in this case 'clickhouse-server' wouldn't start. Command within ClickHouse container is modified with 'sleep' in order to avoid quick restarts and give time to troubleshoot via CLI. Liveness and Readiness probes are disabled as well. ",
						MarkdownDescription: "Allows to troubleshoot Pods during CrashLoopBack state. This may happen when wrong configuration applied, in this case 'clickhouse-server' wouldn't start. Command within ClickHouse container is modified with 'sleep' in order to avoid quick restarts and give time to troubleshoot via CLI. Liveness and Readiness probes are disabled as well. ",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
						},
					},

					"use_templates": schema.ListNestedAttribute{
						Description:         "list of 'ClickHouseInstallationTemplate' (chit) resource names which will merge with current 'CHI' manifest during render Kubernetes resources to create related ClickHouse clusters' ",
						MarkdownDescription: "list of 'ClickHouseInstallationTemplate' (chit) resource names which will merge with current 'CHI' manifest during render Kubernetes resources to create related ClickHouse clusters' ",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "name of 'ClickHouseInstallationTemplate' (chit) resource",
									MarkdownDescription: "name of 'ClickHouseInstallationTemplate' (chit) resource",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Kubernetes namespace where need search 'chit' resource, depending on 'watchNamespaces' settings in 'clichouse-operator'",
									MarkdownDescription: "Kubernetes namespace where need search 'chit' resource, depending on 'watchNamespaces' settings in 'clichouse-operator'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"use_type": schema.StringAttribute{
									Description:         "optional, current strategy is only merge, and current 'chi' settings have more priority than merged template 'chit'",
									MarkdownDescription: "optional, current strategy is only merge, and current 'chi' settings have more priority than merged template 'chit'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("", "merge"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ClickhouseAltinityComClickHouseInstallationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_clickhouse_altinity_com_click_house_installation_v1_manifest")

	var model ClickhouseAltinityComClickHouseInstallationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("clickhouse.altinity.com/v1")
	model.Kind = pointer.String("ClickHouseInstallation")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
