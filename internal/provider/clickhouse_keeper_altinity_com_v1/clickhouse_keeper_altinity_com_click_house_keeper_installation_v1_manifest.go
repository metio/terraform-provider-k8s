/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package clickhouse_keeper_altinity_com_v1

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
	_ datasource.DataSource = &ClickhouseKeeperAltinityComClickHouseKeeperInstallationV1Manifest{}
)

func NewClickhouseKeeperAltinityComClickHouseKeeperInstallationV1Manifest() datasource.DataSource {
	return &ClickhouseKeeperAltinityComClickHouseKeeperInstallationV1Manifest{}
}

type ClickhouseKeeperAltinityComClickHouseKeeperInstallationV1Manifest struct{}

type ClickhouseKeeperAltinityComClickHouseKeeperInstallationV1ManifestData struct {
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
				Files  *map[string]string `tfsdk:"files" json:"files,omitempty"`
				Layout *struct {
					Replicas *[]struct {
						Files    *map[string]string `tfsdk:"files" json:"files,omitempty"`
						Name     *string            `tfsdk:"name" json:"name,omitempty"`
						Settings *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
						Shards   *[]struct {
							Files     *map[string]string `tfsdk:"files" json:"files,omitempty"`
							Name      *string            `tfsdk:"name" json:"name,omitempty"`
							RaftPort  *int64             `tfsdk:"raft_port" json:"raftPort,omitempty"`
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
							ZkPort *int64 `tfsdk:"zk_port" json:"zkPort,omitempty"`
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
				} `tfsdk:"layout" json:"layout,omitempty"`
				Name              *string            `tfsdk:"name" json:"name,omitempty"`
				PdbManaged        *string            `tfsdk:"pdb_managed" json:"pdbManaged,omitempty"`
				PdbMaxUnavailable *int64             `tfsdk:"pdb_max_unavailable" json:"pdbMaxUnavailable,omitempty"`
				Settings          *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
				Templates         *struct {
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
			} `tfsdk:"clusters" json:"clusters,omitempty"`
			Files    *map[string]string `tfsdk:"files" json:"files,omitempty"`
			Settings *map[string]string `tfsdk:"settings" json:"settings,omitempty"`
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
		} `tfsdk:"reconciling" json:"reconciling,omitempty"`
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
					Files     *map[string]string `tfsdk:"files" json:"files,omitempty"`
					Name      *string            `tfsdk:"name" json:"name,omitempty"`
					RaftPort  *int64             `tfsdk:"raft_port" json:"raftPort,omitempty"`
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
					ZkPort *int64 `tfsdk:"zk_port" json:"zkPort,omitempty"`
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
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClickhouseKeeperAltinityComClickHouseKeeperInstallationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_clickhouse_keeper_altinity_com_click_house_keeper_installation_v1_manifest"
}

func (r *ClickhouseKeeperAltinityComClickHouseKeeperInstallationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
								Description:         "describes clusters layout and allows change settings on cluster-level and replica-level ",
								MarkdownDescription: "describes clusters layout and allows change settings on cluster-level and replica-level ",
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

																		"raft_port": schema.Int64Attribute{
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

																		"settings": schema.MapAttribute{
																			Description:         "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in 'Pod' only in one shard related to current replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' override top-level 'chi.spec.configuration.settings', cluster-level 'chi.spec.configuration.clusters.settings' and replica-level 'chi.spec.configuration.clusters.layout.replicas.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																			MarkdownDescription: "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in 'Pod' only in one shard related to current replica during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' override top-level 'chi.spec.configuration.settings', cluster-level 'chi.spec.configuration.clusters.settings' and replica-level 'chi.spec.configuration.clusters.layout.replicas.settings' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
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

																		"zk_port": schema.Int64Attribute{
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

										"pdb_managed": schema.StringAttribute{
											Description:         "Specifies whether the Pod Disruption Budget (PDB) should be managed. During the next installation, if PDB management is enabled, the operator will attempt to retrieve any existing PDB. If none is found, it will create a new one and initiate a reconciliation loop. If PDB management is disabled, the existing PDB will remain intact, and the reconciliation loop will not be executed. By default, PDB management is enabled. ",
											MarkdownDescription: "Specifies whether the Pod Disruption Budget (PDB) should be managed. During the next installation, if PDB management is enabled, the operator will attempt to retrieve any existing PDB. If none is found, it will create a new one and initiate a reconciliation loop. If PDB management is disabled, the existing PDB will remain intact, and the reconciliation loop will not be executed. By default, PDB management is enabled. ",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("", "0", "1", "False", "false", "True", "true", "No", "no", "Yes", "yes", "Off", "off", "On", "on", "Disable", "disable", "Enable", "enable", "Disabled", "disabled", "Enabled", "enabled"),
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"files": schema.MapAttribute{
								Description:         "allows define content of any setting ",
								MarkdownDescription: "allows define content of any setting ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"settings": schema.MapAttribute{
								Description:         "allows configure multiple aspects and behavior for 'clickhouse-keeper' instance ",
								MarkdownDescription: "allows configure multiple aspects and behavior for 'clickhouse-keeper' instance ",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
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
						Description:         "Suspend reconciliation of resources managed by a ClickHouse Keeper. Works as the following: - When 'suspend' is 'true' operator stops reconciling all resources. - When 'suspend' is 'false' or not set, operator reconciles all resources. ",
						MarkdownDescription: "Suspend reconciliation of resources managed by a ClickHouse Keeper. Works as the following: - When 'suspend' is 'true' operator stops reconciling all resources. - When 'suspend' is 'false' or not set, operator reconciles all resources. ",
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

												"raft_port": schema.Int64Attribute{
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

												"settings": schema.MapAttribute{
													Description:         "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' where this template will apply during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
													MarkdownDescription: "optional, allows configure 'clickhouse-server' settings inside <yandex>...</yandex> tag in each 'Pod' where this template will apply during generate 'ConfigMap' which will mount in '/etc/clickhouse-server/conf.d/' More details: https://clickhouse.tech/docs/en/operations/settings/settings/ ",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
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

												"zk_port": schema.Int64Attribute{
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
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ClickhouseKeeperAltinityComClickHouseKeeperInstallationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_clickhouse_keeper_altinity_com_click_house_keeper_installation_v1_manifest")

	var model ClickhouseKeeperAltinityComClickHouseKeeperInstallationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("clickhouse-keeper.altinity.com/v1")
	model.Kind = pointer.String("ClickHouseKeeperInstallation")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
