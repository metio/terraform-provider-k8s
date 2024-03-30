/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &AppsKubeblocksIoClusterV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoClusterV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoClusterV1Alpha1Manifest{}
}

type AppsKubeblocksIoClusterV1Alpha1Manifest struct{}

type AppsKubeblocksIoClusterV1Alpha1ManifestData struct {
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
		Affinity *struct {
			NodeLabels      *map[string]string `tfsdk:"node_labels" json:"nodeLabels,omitempty"`
			PodAntiAffinity *string            `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
			Tenancy         *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
			TopologyKeys    *[]string          `tfsdk:"topology_keys" json:"topologyKeys,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		AvailabilityPolicy *string `tfsdk:"availability_policy" json:"availabilityPolicy,omitempty"`
		Backup             *struct {
			CronExpression          *string `tfsdk:"cron_expression" json:"cronExpression,omitempty"`
			Enabled                 *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Method                  *string `tfsdk:"method" json:"method,omitempty"`
			PitrEnabled             *bool   `tfsdk:"pitr_enabled" json:"pitrEnabled,omitempty"`
			RepoName                *string `tfsdk:"repo_name" json:"repoName,omitempty"`
			RetentionPeriod         *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
			StartingDeadlineMinutes *int64  `tfsdk:"starting_deadline_minutes" json:"startingDeadlineMinutes,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		ClusterDefinitionRef *string `tfsdk:"cluster_definition_ref" json:"clusterDefinitionRef,omitempty"`
		ClusterVersionRef    *string `tfsdk:"cluster_version_ref" json:"clusterVersionRef,omitempty"`
		ComponentSpecs       *[]struct {
			Affinity *struct {
				NodeLabels      *map[string]string `tfsdk:"node_labels" json:"nodeLabels,omitempty"`
				PodAntiAffinity *string            `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				Tenancy         *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
				TopologyKeys    *[]string          `tfsdk:"topology_keys" json:"topologyKeys,omitempty"`
			} `tfsdk:"affinity" json:"affinity,omitempty"`
			ClassDefRef *struct {
				Class *string `tfsdk:"class" json:"class,omitempty"`
				Name  *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"class_def_ref" json:"classDefRef,omitempty"`
			ComponentDef    *string   `tfsdk:"component_def" json:"componentDef,omitempty"`
			ComponentDefRef *string   `tfsdk:"component_def_ref" json:"componentDefRef,omitempty"`
			EnabledLogs     *[]string `tfsdk:"enabled_logs" json:"enabledLogs,omitempty"`
			Instances       *[]string `tfsdk:"instances" json:"instances,omitempty"`
			Issuer          *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				SecretRef *struct {
					Ca   *string `tfsdk:"ca" json:"ca,omitempty"`
					Cert *string `tfsdk:"cert" json:"cert,omitempty"`
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"issuer" json:"issuer,omitempty"`
			Monitor   *bool     `tfsdk:"monitor" json:"monitor,omitempty"`
			Name      *string   `tfsdk:"name" json:"name,omitempty"`
			Nodes     *[]string `tfsdk:"nodes" json:"nodes,omitempty"`
			Replicas  *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			RsmTransformPolicy *string `tfsdk:"rsm_transform_policy" json:"rsmTransformPolicy,omitempty"`
			ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
			ServiceRefs        *[]struct {
				Cluster           *string `tfsdk:"cluster" json:"cluster,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				Namespace         *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ServiceDescriptor *string `tfsdk:"service_descriptor" json:"serviceDescriptor,omitempty"`
			} `tfsdk:"service_refs" json:"serviceRefs,omitempty"`
			Services *[]struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				ServiceType *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
			SwitchPolicy *struct {
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"switch_policy" json:"switchPolicy,omitempty"`
			Tls              *bool              `tfsdk:"tls" json:"tls,omitempty"`
			Tolerations      *map[string]string `tfsdk:"tolerations" json:"tolerations,omitempty"`
			UpdateStrategy   *string            `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
			UserResourceRefs *struct {
				ConfigMapRefs *[]struct {
					AsVolumeFrom *[]string `tfsdk:"as_volume_from" json:"asVolumeFrom,omitempty"`
					ConfigMap    *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					MountPoint *string `tfsdk:"mount_point" json:"mountPoint,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					SubPath    *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				} `tfsdk:"config_map_refs" json:"configMapRefs,omitempty"`
				SecretRefs *[]struct {
					AsVolumeFrom *[]string `tfsdk:"as_volume_from" json:"asVolumeFrom,omitempty"`
					MountPoint   *string   `tfsdk:"mount_point" json:"mountPoint,omitempty"`
					Name         *string   `tfsdk:"name" json:"name,omitempty"`
					Secret       *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					SubPath *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
			} `tfsdk:"user_resource_refs" json:"userResourceRefs,omitempty"`
			VolumeClaimTemplates *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Spec *struct {
					AccessModes *map[string]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
					Resources   *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
					VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		} `tfsdk:"component_specs" json:"componentSpecs,omitempty"`
		Monitor *struct {
			MonitoringInterval *string `tfsdk:"monitoring_interval" json:"monitoringInterval,omitempty"`
		} `tfsdk:"monitor" json:"monitor,omitempty"`
		Network *struct {
			HostNetworkAccessible *bool `tfsdk:"host_network_accessible" json:"hostNetworkAccessible,omitempty"`
			PubliclyAccessible    *bool `tfsdk:"publicly_accessible" json:"publiclyAccessible,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources *struct {
			Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
			Memory *string `tfsdk:"memory" json:"memory,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Services      *map[string]string `tfsdk:"services" json:"services,omitempty"`
		ShardingSpecs *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Shards   *int64  `tfsdk:"shards" json:"shards,omitempty"`
			Template *struct {
				Affinity *struct {
					NodeLabels      *map[string]string `tfsdk:"node_labels" json:"nodeLabels,omitempty"`
					PodAntiAffinity *string            `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					Tenancy         *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
					TopologyKeys    *[]string          `tfsdk:"topology_keys" json:"topologyKeys,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				ClassDefRef *struct {
					Class *string `tfsdk:"class" json:"class,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"class_def_ref" json:"classDefRef,omitempty"`
				ComponentDef    *string   `tfsdk:"component_def" json:"componentDef,omitempty"`
				ComponentDefRef *string   `tfsdk:"component_def_ref" json:"componentDefRef,omitempty"`
				EnabledLogs     *[]string `tfsdk:"enabled_logs" json:"enabledLogs,omitempty"`
				Instances       *[]string `tfsdk:"instances" json:"instances,omitempty"`
				Issuer          *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					SecretRef *struct {
						Ca   *string `tfsdk:"ca" json:"ca,omitempty"`
						Cert *string `tfsdk:"cert" json:"cert,omitempty"`
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"issuer" json:"issuer,omitempty"`
				Monitor   *bool     `tfsdk:"monitor" json:"monitor,omitempty"`
				Name      *string   `tfsdk:"name" json:"name,omitempty"`
				Nodes     *[]string `tfsdk:"nodes" json:"nodes,omitempty"`
				Replicas  *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				RsmTransformPolicy *string `tfsdk:"rsm_transform_policy" json:"rsmTransformPolicy,omitempty"`
				ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
				ServiceRefs        *[]struct {
					Cluster           *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Name              *string `tfsdk:"name" json:"name,omitempty"`
					Namespace         *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ServiceDescriptor *string `tfsdk:"service_descriptor" json:"serviceDescriptor,omitempty"`
				} `tfsdk:"service_refs" json:"serviceRefs,omitempty"`
				Services *[]struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
					ServiceType *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
				} `tfsdk:"services" json:"services,omitempty"`
				SwitchPolicy *struct {
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"switch_policy" json:"switchPolicy,omitempty"`
				Tls              *bool              `tfsdk:"tls" json:"tls,omitempty"`
				Tolerations      *map[string]string `tfsdk:"tolerations" json:"tolerations,omitempty"`
				UpdateStrategy   *string            `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
				UserResourceRefs *struct {
					ConfigMapRefs *[]struct {
						AsVolumeFrom *[]string `tfsdk:"as_volume_from" json:"asVolumeFrom,omitempty"`
						ConfigMap    *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						MountPoint *string `tfsdk:"mount_point" json:"mountPoint,omitempty"`
						Name       *string `tfsdk:"name" json:"name,omitempty"`
						SubPath    *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					} `tfsdk:"config_map_refs" json:"configMapRefs,omitempty"`
					SecretRefs *[]struct {
						AsVolumeFrom *[]string `tfsdk:"as_volume_from" json:"asVolumeFrom,omitempty"`
						MountPoint   *string   `tfsdk:"mount_point" json:"mountPoint,omitempty"`
						Name         *string   `tfsdk:"name" json:"name,omitempty"`
						Secret       *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						SubPath *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
				} `tfsdk:"user_resource_refs" json:"userResourceRefs,omitempty"`
				VolumeClaimTemplates *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Spec *struct {
						AccessModes *map[string]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
						Resources   *struct {
							Claims *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"claims" json:"claims,omitempty"`
							Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
							Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
						StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
						VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
		} `tfsdk:"sharding_specs" json:"shardingSpecs,omitempty"`
		Storage *struct {
			Size *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Tenancy           *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
		TerminationPolicy *string            `tfsdk:"termination_policy" json:"terminationPolicy,omitempty"`
		Tolerations       *map[string]string `tfsdk:"tolerations" json:"tolerations,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_cluster_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cluster is the Schema for the clusters API.",
		MarkdownDescription: "Cluster is the Schema for the clusters API.",
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
				Description:         "ClusterSpec defines the desired state of Cluster.",
				MarkdownDescription: "ClusterSpec defines the desired state of Cluster.",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "A group of affinity scheduling rules.",
						MarkdownDescription: "A group of affinity scheduling rules.",
						Attributes: map[string]schema.Attribute{
							"node_labels": schema.MapAttribute{
								Description:         "Indicates that pods must be scheduled to the nodes with the specified node labels.",
								MarkdownDescription: "Indicates that pods must be scheduled to the nodes with the specified node labels.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_anti_affinity": schema.StringAttribute{
								Description:         "Specifies the anti-affinity level of pods within a component.",
								MarkdownDescription: "Specifies the anti-affinity level of pods within a component.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Preferred", "Required"),
								},
							},

							"tenancy": schema.StringAttribute{
								Description:         "Defines how pods are distributed across nodes.",
								MarkdownDescription: "Defines how pods are distributed across nodes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("SharedNode", "DedicatedNode"),
								},
							},

							"topology_keys": schema.ListAttribute{
								Description:         "Represents the key of node labels.  Nodes with a label containing this key and identical values are considered to be in the same topology. This is used as the topology domain for pod anti-affinity and pod spread constraint. Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey, along with any other custom label key.",
								MarkdownDescription: "Represents the key of node labels.  Nodes with a label containing this key and identical values are considered to be in the same topology. This is used as the topology domain for pod anti-affinity and pod spread constraint. Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey, along with any other custom label key.",
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

					"availability_policy": schema.StringAttribute{
						Description:         "Describes the availability policy, including zone, node, and none.",
						MarkdownDescription: "Describes the availability policy, including zone, node, and none.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("zone", "node", "none"),
						},
					},

					"backup": schema.SingleNestedAttribute{
						Description:         "Cluster backup configuration.",
						MarkdownDescription: "Cluster backup configuration.",
						Attributes: map[string]schema.Attribute{
							"cron_expression": schema.StringAttribute{
								Description:         "The cron expression for the schedule. The timezone is in UTC. See https://en.wikipedia.org/wiki/Cron.",
								MarkdownDescription: "The cron expression for the schedule. The timezone is in UTC. See https://en.wikipedia.org/wiki/Cron.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Specifies whether automated backup is enabled.",
								MarkdownDescription: "Specifies whether automated backup is enabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "Specifies the backup method to use, as defined in backupPolicy.",
								MarkdownDescription: "Specifies the backup method to use, as defined in backupPolicy.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"pitr_enabled": schema.BoolAttribute{
								Description:         "Specifies whether to enable point-in-time recovery.",
								MarkdownDescription: "Specifies whether to enable point-in-time recovery.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repo_name": schema.StringAttribute{
								Description:         "Specifies the name of the backupRepo. If not set, the default backupRepo will be used.",
								MarkdownDescription: "Specifies the name of the backupRepo. If not set, the default backupRepo will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retention_period": schema.StringAttribute{
								Description:         "Determines the duration for which the backup should be retained. All backups older than this period will be removed by the controller.  For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 	2y - months: 	6mo - days: 		30d - hours: 	12h - minutes: 	30m  You can also combine the above durations. For example: 30d12h30m",
								MarkdownDescription: "Determines the duration for which the backup should be retained. All backups older than this period will be removed by the controller.  For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 	2y - months: 	6mo - days: 		30d - hours: 	12h - minutes: 	30m  You can also combine the above durations. For example: 30d12h30m",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"starting_deadline_minutes": schema.Int64Attribute{
								Description:         "Defines the deadline in minutes for starting the backup job if it misses its scheduled time for any reason.",
								MarkdownDescription: "Defines the deadline in minutes for starting the backup job if it misses its scheduled time for any reason.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(1440),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_definition_ref": schema.StringAttribute{
						Description:         "Refers to the ClusterDefinition name. If not specified, ComponentDef must be specified for each Component in ComponentSpecs.",
						MarkdownDescription: "Refers to the ClusterDefinition name. If not specified, ComponentDef must be specified for each Component in ComponentSpecs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
						},
					},

					"cluster_version_ref": schema.StringAttribute{
						Description:         "Refers to the ClusterVersion name.",
						MarkdownDescription: "Refers to the ClusterVersion name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
						},
					},

					"component_specs": schema.ListNestedAttribute{
						Description:         "List of componentSpec used to define the components that make up a cluster. ComponentSpecs and ShardingSpecs cannot both be empty at the same time.",
						MarkdownDescription: "List of componentSpec used to define the components that make up a cluster. ComponentSpecs and ShardingSpecs cannot both be empty at the same time.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"affinity": schema.SingleNestedAttribute{
									Description:         "A group of affinity scheduling rules.",
									MarkdownDescription: "A group of affinity scheduling rules.",
									Attributes: map[string]schema.Attribute{
										"node_labels": schema.MapAttribute{
											Description:         "Indicates that pods must be scheduled to the nodes with the specified node labels.",
											MarkdownDescription: "Indicates that pods must be scheduled to the nodes with the specified node labels.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pod_anti_affinity": schema.StringAttribute{
											Description:         "Specifies the anti-affinity level of pods within a component.",
											MarkdownDescription: "Specifies the anti-affinity level of pods within a component.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Preferred", "Required"),
											},
										},

										"tenancy": schema.StringAttribute{
											Description:         "Defines how pods are distributed across nodes.",
											MarkdownDescription: "Defines how pods are distributed across nodes.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("SharedNode", "DedicatedNode"),
											},
										},

										"topology_keys": schema.ListAttribute{
											Description:         "Represents the key of node labels.  Nodes with a label containing this key and identical values are considered to be in the same topology. This is used as the topology domain for pod anti-affinity and pod spread constraint. Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey, along with any other custom label key.",
											MarkdownDescription: "Represents the key of node labels.  Nodes with a label containing this key and identical values are considered to be in the same topology. This is used as the topology domain for pod anti-affinity and pod spread constraint. Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey, along with any other custom label key.",
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

								"class_def_ref": schema.SingleNestedAttribute{
									Description:         "References the class defined in ComponentClassDefinition.",
									MarkdownDescription: "References the class defined in ComponentClassDefinition.",
									Attributes: map[string]schema.Attribute{
										"class": schema.StringAttribute{
											Description:         "Defines the name of the class that is defined in the ComponentClassDefinition.",
											MarkdownDescription: "Defines the name of the class that is defined in the ComponentClassDefinition.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the ComponentClassDefinition.",
											MarkdownDescription: "Specifies the name of the ComponentClassDefinition.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"component_def": schema.StringAttribute{
									Description:         "References the name of the ComponentDefinition. If both componentDefRef and componentDef are provided, the componentDef will take precedence over componentDefRef.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDef is immutable'",
									MarkdownDescription: "References the name of the ComponentDefinition. If both componentDefRef and componentDef are provided, the componentDef will take precedence over componentDefRef.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDef is immutable'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(22),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"component_def_ref": schema.StringAttribute{
									Description:         "References the componentDef defined in the ClusterDefinition spec. Must comply with the IANA Service Naming rule.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDefRef is immutable'",
									MarkdownDescription: "References the componentDef defined in the ClusterDefinition spec. Must comply with the IANA Service Naming rule.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDefRef is immutable'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(22),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
									},
								},

								"enabled_logs": schema.ListAttribute{
									Description:         "Indicates which log file takes effect in the database cluster.",
									MarkdownDescription: "Indicates which log file takes effect in the database cluster.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"instances": schema.ListAttribute{
									Description:         "Defines the list of instances to be deleted priorly. If the RsmTransformPolicy is specified as ToPod, the list of instances will be used.",
									MarkdownDescription: "Defines the list of instances to be deleted priorly. If the RsmTransformPolicy is specified as ToPod, the list of instances will be used.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"issuer": schema.SingleNestedAttribute{
									Description:         "Defines provider context for TLS certs. Required when TLS is enabled.",
									MarkdownDescription: "Defines provider context for TLS certs. Required when TLS is enabled.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "The issuer for TLS certificates.",
											MarkdownDescription: "The issuer for TLS certificates.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "SecretRef is the reference to the TLS certificates secret. It is required when the issuer is set to UserProvided.",
											MarkdownDescription: "SecretRef is the reference to the TLS certificates secret. It is required when the issuer is set to UserProvided.",
											Attributes: map[string]schema.Attribute{
												"ca": schema.StringAttribute{
													Description:         "CA cert key in Secret",
													MarkdownDescription: "CA cert key in Secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"cert": schema.StringAttribute{
													Description:         "Cert key in Secret",
													MarkdownDescription: "Cert key in Secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key of TLS private key in Secret",
													MarkdownDescription: "Key of TLS private key in Secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the Secret",
													MarkdownDescription: "Name of the Secret",
													Required:            true,
													Optional:            false,
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

								"monitor": schema.BoolAttribute{
									Description:         "To enable monitoring.",
									MarkdownDescription: "To enable monitoring.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Specifies the name of the cluster's component. This name is also part of the Service DNS name and must comply with the IANA Service Naming rule. When ClusterComponentSpec is referenced as a template, the name is optional. Otherwise, it is required.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='name is immutable'",
									MarkdownDescription: "Specifies the name of the cluster's component. This name is also part of the Service DNS name and must comply with the IANA Service Naming rule. When ClusterComponentSpec is referenced as a template, the name is optional. Otherwise, it is required.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='name is immutable'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(22),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
									},
								},

								"nodes": schema.ListAttribute{
									Description:         "Defines the list of nodes that pods can schedule. If the RsmTransformPolicy is specified as ToPod, the list of nodes will be used. If the list of nodes is empty, no specific node will be assigned. However, if the list of nodes is filled, all pods will be evenly scheduled across the nodes in the list.",
									MarkdownDescription: "Defines the list of nodes that pods can schedule. If the RsmTransformPolicy is specified as ToPod, the list of nodes will be used. If the list of nodes is empty, no specific node will be assigned. However, if the list of nodes is filled, all pods will be evenly scheduled across the nodes in the list.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "Specifies the number of component replicas.",
									MarkdownDescription: "Specifies the number of component replicas.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Specifies the resources requests and limits of the workload.",
									MarkdownDescription: "Specifies the resources requests and limits of the workload.",
									Attributes: map[string]schema.Attribute{
										"claims": schema.ListNestedAttribute{
											Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
											MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
														MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

										"limits": schema.MapAttribute{
											Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

								"rsm_transform_policy": schema.StringAttribute{
									Description:         "Defines the policy to generate sts using rsm.",
									MarkdownDescription: "Defines the policy to generate sts using rsm.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("ToPod", "ToSts"),
									},
								},

								"service_account_name": schema.StringAttribute{
									Description:         "Specifies the name of the ServiceAccount that the running component depends on.",
									MarkdownDescription: "Specifies the name of the ServiceAccount that the running component depends on.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"service_refs": schema.ListNestedAttribute{
									Description:         "Defines service references for the current component.  Based on the referenced services, they can be categorized into two types:  - Service provided by external sources: These services are provided by external sources and are not managed by KubeBlocks. They can be Kubernetes-based or non-Kubernetes services. For external services, an additional ServiceDescriptor object is needed to establish the service binding. - Service provided by other KubeBlocks clusters: These services are provided by other KubeBlocks clusters. Binding to these services is done by specifying the name of the hosting cluster.  Each type of service reference requires specific configurations and bindings to establish the connection and interaction with the respective services. Note that the ServiceRef has cluster-level semantic consistency, meaning that within the same Cluster, service references with the same ServiceRef.Name are considered to be the same service. It is only allowed to bind to the same Cluster or ServiceDescriptor.",
									MarkdownDescription: "Defines service references for the current component.  Based on the referenced services, they can be categorized into two types:  - Service provided by external sources: These services are provided by external sources and are not managed by KubeBlocks. They can be Kubernetes-based or non-Kubernetes services. For external services, an additional ServiceDescriptor object is needed to establish the service binding. - Service provided by other KubeBlocks clusters: These services are provided by other KubeBlocks clusters. Binding to these services is done by specifying the name of the hosting cluster.  Each type of service reference requires specific configurations and bindings to establish the connection and interaction with the respective services. Note that the ServiceRef has cluster-level semantic consistency, meaning that within the same Cluster, service references with the same ServiceRef.Name are considered to be the same service. It is only allowed to bind to the same Cluster or ServiceDescriptor.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cluster": schema.StringAttribute{
												Description:         "The name of the KubeBlocks cluster being referenced when a service provided by another KubeBlocks cluster is being referenced.  By default, the clusterDefinition.spec.connectionCredential secret corresponding to the referenced Cluster will be used to bind to the current component. The connection credential secret should include and correspond to the following fields: endpoint, port, username, and password when a KubeBlocks cluster is being referenced.  Under this referencing approach, the ServiceKind and ServiceVersion of service reference declaration defined in the ClusterDefinition will not be validated. If both Cluster and ServiceDescriptor are specified, the Cluster takes precedence.",
												MarkdownDescription: "The name of the KubeBlocks cluster being referenced when a service provided by another KubeBlocks cluster is being referenced.  By default, the clusterDefinition.spec.connectionCredential secret corresponding to the referenced Cluster will be used to bind to the current component. The connection credential secret should include and correspond to the following fields: endpoint, port, username, and password when a KubeBlocks cluster is being referenced.  Under this referencing approach, the ServiceKind and ServiceVersion of service reference declaration defined in the ClusterDefinition will not be validated. If both Cluster and ServiceDescriptor are specified, the Cluster takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Specifies the identifier of the service reference declaration. It corresponds to the serviceRefDeclaration name defined in the clusterDefinition.componentDefs[*].serviceRefDeclarations[*].name.",
												MarkdownDescription: "Specifies the identifier of the service reference declaration. It corresponds to the serviceRefDeclaration name defined in the clusterDefinition.componentDefs[*].serviceRefDeclarations[*].name.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object. If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current cluster by default.",
												MarkdownDescription: "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object. If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current cluster by default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_descriptor": schema.StringAttribute{
												Description:         "The service descriptor of the service provided by external sources.  When referencing a service provided by external sources, the ServiceDescriptor object name is required to establish the service binding. The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKind and serviceVersion defined in the service reference declaration in the ClusterDefinition.  If both Cluster and ServiceDescriptor are specified, the Cluster takes precedence.",
												MarkdownDescription: "The service descriptor of the service provided by external sources.  When referencing a service provided by external sources, the ServiceDescriptor object name is required to establish the service binding. The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKind and serviceVersion defined in the service reference declaration in the ClusterDefinition.  If both Cluster and ServiceDescriptor are specified, the Cluster takes precedence.",
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

								"services": schema.ListNestedAttribute{
									Description:         "Services expose endpoints that can be accessed by clients.",
									MarkdownDescription: "Services expose endpoints that can be accessed by clients.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "If ServiceType is LoadBalancer, cloud provider related parameters can be put here. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												MarkdownDescription: "If ServiceType is LoadBalancer, cloud provider related parameters can be put here. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "The name of the service.",
												MarkdownDescription: "The name of the service.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(15),
												},
											},

											"service_type": schema.StringAttribute{
												Description:         "Determines how the Service is exposed. Valid options are ClusterIP, NodePort, and LoadBalancer.  - 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP.  More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
												MarkdownDescription: "Determines how the Service is exposed. Valid options are ClusterIP, NodePort, and LoadBalancer.  - 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP.  More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"switch_policy": schema.SingleNestedAttribute{
									Description:         "Defines the strategy for switchover and failover when workloadType is Replication.",
									MarkdownDescription: "Defines the strategy for switchover and failover when workloadType is Replication.",
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											Description:         "Type specifies the type of switch policy to be applied.",
											MarkdownDescription: "Type specifies the type of switch policy to be applied.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Noop"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tls": schema.BoolAttribute{
									Description:         "Enables or disables TLS certs.",
									MarkdownDescription: "Enables or disables TLS certs.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tolerations": schema.MapAttribute{
									Description:         "Attached to tolerate any taint that matches the triple 'key,value,effect' using the matching operator 'operator'.",
									MarkdownDescription: "Attached to tolerate any taint that matches the triple 'key,value,effect' using the matching operator 'operator'.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"update_strategy": schema.StringAttribute{
									Description:         "Defines the update strategy for the component. Not supported.",
									MarkdownDescription: "Defines the update strategy for the component. Not supported.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Serial", "BestEffortParallel", "Parallel"),
									},
								},

								"user_resource_refs": schema.SingleNestedAttribute{
									Description:         "Defines the user-defined volumes.",
									MarkdownDescription: "Defines the user-defined volumes.",
									Attributes: map[string]schema.Attribute{
										"config_map_refs": schema.ListNestedAttribute{
											Description:         "ConfigMapRefs defines the user-defined config maps.",
											MarkdownDescription: "ConfigMapRefs defines the user-defined config maps.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"as_volume_from": schema.ListAttribute{
														Description:         "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
														MarkdownDescription: "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"config_map": schema.SingleNestedAttribute{
														Description:         "ConfigMap specifies the ConfigMap to be mounted as a volume.",
														MarkdownDescription: "ConfigMap specifies the ConfigMap to be mounted as a volume.",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "key is the key to project.",
																			MarkdownDescription: "key is the key to project.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
																Description:         "optional specify whether the ConfigMap or its keys must be defined",
																MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"mount_point": schema.StringAttribute{
														Description:         "MountPoint is the filesystem path where the volume will be mounted.",
														MarkdownDescription: "MountPoint is the filesystem path where the volume will be mounted.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(256),
															stringvalidator.RegexMatches(regexp.MustCompile(`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
														MarkdownDescription: "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
														},
													},

													"sub_path": schema.StringAttribute{
														Description:         "SubPath specifies a path within the volume from which to mount.",
														MarkdownDescription: "SubPath specifies a path within the volume from which to mount.",
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

										"secret_refs": schema.ListNestedAttribute{
											Description:         "SecretRefs defines the user-defined secrets.",
											MarkdownDescription: "SecretRefs defines the user-defined secrets.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"as_volume_from": schema.ListAttribute{
														Description:         "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
														MarkdownDescription: "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"mount_point": schema.StringAttribute{
														Description:         "MountPoint is the filesystem path where the volume will be mounted.",
														MarkdownDescription: "MountPoint is the filesystem path where the volume will be mounted.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(256),
															stringvalidator.RegexMatches(regexp.MustCompile(`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
														MarkdownDescription: "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
														},
													},

													"secret": schema.SingleNestedAttribute{
														Description:         "Secret specifies the secret to be mounted as a volume.",
														MarkdownDescription: "Secret specifies the secret to be mounted as a volume.",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "key is the key to project.",
																			MarkdownDescription: "key is the key to project.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

															"optional": schema.BoolAttribute{
																Description:         "optional field specify whether the Secret or its keys must be defined",
																MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
																Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"sub_path": schema.StringAttribute{
														Description:         "SubPath specifies a path within the volume from which to mount.",
														MarkdownDescription: "SubPath specifies a path within the volume from which to mount.",
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

								"volume_claim_templates": schema.ListNestedAttribute{
									Description:         "Provides information for statefulset.spec.volumeClaimTemplates.",
									MarkdownDescription: "Provides information for statefulset.spec.volumeClaimTemplates.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Refers to 'clusterDefinition.spec.componentDefs.containers.volumeMounts.name'.",
												MarkdownDescription: "Refers to 'clusterDefinition.spec.componentDefs.containers.volumeMounts.name'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"spec": schema.SingleNestedAttribute{
												Description:         "Defines the desired characteristics of a volume requested by a pod author.",
												MarkdownDescription: "Defines the desired characteristics of a volume requested by a pod author.",
												Attributes: map[string]schema.Attribute{
													"access_modes": schema.MapAttribute{
														Description:         "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
														MarkdownDescription: "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
														MarkdownDescription: "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
														Attributes: map[string]schema.Attribute{
															"claims": schema.ListNestedAttribute{
																Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																			MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

															"limits": schema.MapAttribute{
																Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"requests": schema.MapAttribute{
																Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

													"storage_class_name": schema.StringAttribute{
														Description:         "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
														MarkdownDescription: "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"volume_mode": schema.StringAttribute{
														Description:         "Defines what type of volume is required by the claim.",
														MarkdownDescription: "Defines what type of volume is required by the claim.",
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

					"monitor": schema.SingleNestedAttribute{
						Description:         "The configuration of monitor.",
						MarkdownDescription: "The configuration of monitor.",
						Attributes: map[string]schema.Attribute{
							"monitoring_interval": schema.StringAttribute{
								Description:         "Defines the frequency at which monitoring occurs. If set to 0, monitoring is disabled.",
								MarkdownDescription: "Defines the frequency at which monitoring occurs. If set to 0, monitoring is disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "The configuration of network.",
						MarkdownDescription: "The configuration of network.",
						Attributes: map[string]schema.Attribute{
							"host_network_accessible": schema.BoolAttribute{
								Description:         "Indicates whether the host network can be accessed. By default, this is set to false.",
								MarkdownDescription: "Indicates whether the host network can be accessed. By default, this is set to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"publicly_accessible": schema.BoolAttribute{
								Description:         "Indicates whether the network is accessible to the public. By default, this is set to false.",
								MarkdownDescription: "Indicates whether the network is accessible to the public. By default, this is set to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Specifies the replicas of the first componentSpec, if the replicas of the first componentSpec is specified, this value will be ignored.",
						MarkdownDescription: "Specifies the replicas of the first componentSpec, if the replicas of the first componentSpec is specified, this value will be ignored.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Specifies the resources of the first componentSpec, if the resources of the first componentSpec is specified, this value will be ignored.",
						MarkdownDescription: "Specifies the resources of the first componentSpec, if the resources of the first componentSpec is specified, this value will be ignored.",
						Attributes: map[string]schema.Attribute{
							"cpu": schema.StringAttribute{
								Description:         "Specifies the amount of processing power the cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Specifies the amount of processing power the cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"memory": schema.StringAttribute{
								Description:         "Specifies the amount of memory the cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Specifies the amount of memory the cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"services": schema.MapAttribute{
						Description:         "Defines the services to access a cluster.",
						MarkdownDescription: "Defines the services to access a cluster.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sharding_specs": schema.ListNestedAttribute{
						Description:         "List of ShardingSpec used to define components with a sharding topology structure that make up a cluster. ShardingSpecs and ComponentSpecs cannot both be empty at the same time.",
						MarkdownDescription: "List of ShardingSpec used to define components with a sharding topology structure that make up a cluster. ShardingSpecs and ComponentSpecs cannot both be empty at the same time.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Specifies the identifier for the sharding configuration. This identifier is included as part of the Service DNS name and must comply with IANA Service Naming rules. It is used to generate the names of underlying components following the pattern '$(ShardingSpec.Name)-$(ShardID)'. Note that the name of the component template defined in ShardingSpec.Template.Name will be disregarded.",
									MarkdownDescription: "Specifies the identifier for the sharding configuration. This identifier is included as part of the Service DNS name and must comply with IANA Service Naming rules. It is used to generate the names of underlying components following the pattern '$(ShardingSpec.Name)-$(ShardID)'. Note that the name of the component template defined in ShardingSpec.Template.Name will be disregarded.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(15),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"shards": schema.Int64Attribute{
									Description:         "Specifies the number of components, all of which will have identical specifications and definitions.  The number of replicas for each component should be defined by template.replicas. The logical relationship between these components should be maintained by the components themselves. KubeBlocks only provides lifecycle management for sharding, including:  1. Executing the postProvision Action defined in the ComponentDefinition when the number of shards increases, provided the conditions are met. 2. Executing the preTerminate Action defined in the ComponentDefinition when the number of shards decreases, provided the conditions are met. Resources and data associated with the corresponding Component will also be deleted.",
									MarkdownDescription: "Specifies the number of components, all of which will have identical specifications and definitions.  The number of replicas for each component should be defined by template.replicas. The logical relationship between these components should be maintained by the components themselves. KubeBlocks only provides lifecycle management for sharding, including:  1. Executing the postProvision Action defined in the ComponentDefinition when the number of shards increases, provided the conditions are met. 2. Executing the preTerminate Action defined in the ComponentDefinition when the number of shards decreases, provided the conditions are met. Resources and data associated with the corresponding Component will also be deleted.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(2048),
									},
								},

								"template": schema.SingleNestedAttribute{
									Description:         "The blueprint for the components. Generates a set of components (also referred to as shards) based on this template. All components or shards generated will have identical specifications and definitions.",
									MarkdownDescription: "The blueprint for the components. Generates a set of components (also referred to as shards) based on this template. All components or shards generated will have identical specifications and definitions.",
									Attributes: map[string]schema.Attribute{
										"affinity": schema.SingleNestedAttribute{
											Description:         "A group of affinity scheduling rules.",
											MarkdownDescription: "A group of affinity scheduling rules.",
											Attributes: map[string]schema.Attribute{
												"node_labels": schema.MapAttribute{
													Description:         "Indicates that pods must be scheduled to the nodes with the specified node labels.",
													MarkdownDescription: "Indicates that pods must be scheduled to the nodes with the specified node labels.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_anti_affinity": schema.StringAttribute{
													Description:         "Specifies the anti-affinity level of pods within a component.",
													MarkdownDescription: "Specifies the anti-affinity level of pods within a component.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Preferred", "Required"),
													},
												},

												"tenancy": schema.StringAttribute{
													Description:         "Defines how pods are distributed across nodes.",
													MarkdownDescription: "Defines how pods are distributed across nodes.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("SharedNode", "DedicatedNode"),
													},
												},

												"topology_keys": schema.ListAttribute{
													Description:         "Represents the key of node labels.  Nodes with a label containing this key and identical values are considered to be in the same topology. This is used as the topology domain for pod anti-affinity and pod spread constraint. Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey, along with any other custom label key.",
													MarkdownDescription: "Represents the key of node labels.  Nodes with a label containing this key and identical values are considered to be in the same topology. This is used as the topology domain for pod anti-affinity and pod spread constraint. Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey, along with any other custom label key.",
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

										"class_def_ref": schema.SingleNestedAttribute{
											Description:         "References the class defined in ComponentClassDefinition.",
											MarkdownDescription: "References the class defined in ComponentClassDefinition.",
											Attributes: map[string]schema.Attribute{
												"class": schema.StringAttribute{
													Description:         "Defines the name of the class that is defined in the ComponentClassDefinition.",
													MarkdownDescription: "Defines the name of the class that is defined in the ComponentClassDefinition.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Specifies the name of the ComponentClassDefinition.",
													MarkdownDescription: "Specifies the name of the ComponentClassDefinition.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(63),
														stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"component_def": schema.StringAttribute{
											Description:         "References the name of the ComponentDefinition. If both componentDefRef and componentDef are provided, the componentDef will take precedence over componentDefRef.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDef is immutable'",
											MarkdownDescription: "References the name of the ComponentDefinition. If both componentDefRef and componentDef are provided, the componentDef will take precedence over componentDefRef.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDef is immutable'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(22),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"component_def_ref": schema.StringAttribute{
											Description:         "References the componentDef defined in the ClusterDefinition spec. Must comply with the IANA Service Naming rule.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDefRef is immutable'",
											MarkdownDescription: "References the componentDef defined in the ClusterDefinition spec. Must comply with the IANA Service Naming rule.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDefRef is immutable'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(22),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},

										"enabled_logs": schema.ListAttribute{
											Description:         "Indicates which log file takes effect in the database cluster.",
											MarkdownDescription: "Indicates which log file takes effect in the database cluster.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"instances": schema.ListAttribute{
											Description:         "Defines the list of instances to be deleted priorly. If the RsmTransformPolicy is specified as ToPod, the list of instances will be used.",
											MarkdownDescription: "Defines the list of instances to be deleted priorly. If the RsmTransformPolicy is specified as ToPod, the list of instances will be used.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"issuer": schema.SingleNestedAttribute{
											Description:         "Defines provider context for TLS certs. Required when TLS is enabled.",
											MarkdownDescription: "Defines provider context for TLS certs. Required when TLS is enabled.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The issuer for TLS certificates.",
													MarkdownDescription: "The issuer for TLS certificates.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "SecretRef is the reference to the TLS certificates secret. It is required when the issuer is set to UserProvided.",
													MarkdownDescription: "SecretRef is the reference to the TLS certificates secret. It is required when the issuer is set to UserProvided.",
													Attributes: map[string]schema.Attribute{
														"ca": schema.StringAttribute{
															Description:         "CA cert key in Secret",
															MarkdownDescription: "CA cert key in Secret",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"cert": schema.StringAttribute{
															Description:         "Cert key in Secret",
															MarkdownDescription: "Cert key in Secret",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Key of TLS private key in Secret",
															MarkdownDescription: "Key of TLS private key in Secret",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the Secret",
															MarkdownDescription: "Name of the Secret",
															Required:            true,
															Optional:            false,
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

										"monitor": schema.BoolAttribute{
											Description:         "To enable monitoring.",
											MarkdownDescription: "To enable monitoring.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the cluster's component. This name is also part of the Service DNS name and must comply with the IANA Service Naming rule. When ClusterComponentSpec is referenced as a template, the name is optional. Otherwise, it is required.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='name is immutable'",
											MarkdownDescription: "Specifies the name of the cluster's component. This name is also part of the Service DNS name and must comply with the IANA Service Naming rule. When ClusterComponentSpec is referenced as a template, the name is optional. Otherwise, it is required.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='name is immutable'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(22),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},

										"nodes": schema.ListAttribute{
											Description:         "Defines the list of nodes that pods can schedule. If the RsmTransformPolicy is specified as ToPod, the list of nodes will be used. If the list of nodes is empty, no specific node will be assigned. However, if the list of nodes is filled, all pods will be evenly scheduled across the nodes in the list.",
											MarkdownDescription: "Defines the list of nodes that pods can schedule. If the RsmTransformPolicy is specified as ToPod, the list of nodes will be used. If the list of nodes is empty, no specific node will be assigned. However, if the list of nodes is filled, all pods will be evenly scheduled across the nodes in the list.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replicas": schema.Int64Attribute{
											Description:         "Specifies the number of component replicas.",
											MarkdownDescription: "Specifies the number of component replicas.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "Specifies the resources requests and limits of the workload.",
											MarkdownDescription: "Specifies the resources requests and limits of the workload.",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
													MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

												"limits": schema.MapAttribute{
													Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

										"rsm_transform_policy": schema.StringAttribute{
											Description:         "Defines the policy to generate sts using rsm.",
											MarkdownDescription: "Defines the policy to generate sts using rsm.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("ToPod", "ToSts"),
											},
										},

										"service_account_name": schema.StringAttribute{
											Description:         "Specifies the name of the ServiceAccount that the running component depends on.",
											MarkdownDescription: "Specifies the name of the ServiceAccount that the running component depends on.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_refs": schema.ListNestedAttribute{
											Description:         "Defines service references for the current component.  Based on the referenced services, they can be categorized into two types:  - Service provided by external sources: These services are provided by external sources and are not managed by KubeBlocks. They can be Kubernetes-based or non-Kubernetes services. For external services, an additional ServiceDescriptor object is needed to establish the service binding. - Service provided by other KubeBlocks clusters: These services are provided by other KubeBlocks clusters. Binding to these services is done by specifying the name of the hosting cluster.  Each type of service reference requires specific configurations and bindings to establish the connection and interaction with the respective services. Note that the ServiceRef has cluster-level semantic consistency, meaning that within the same Cluster, service references with the same ServiceRef.Name are considered to be the same service. It is only allowed to bind to the same Cluster or ServiceDescriptor.",
											MarkdownDescription: "Defines service references for the current component.  Based on the referenced services, they can be categorized into two types:  - Service provided by external sources: These services are provided by external sources and are not managed by KubeBlocks. They can be Kubernetes-based or non-Kubernetes services. For external services, an additional ServiceDescriptor object is needed to establish the service binding. - Service provided by other KubeBlocks clusters: These services are provided by other KubeBlocks clusters. Binding to these services is done by specifying the name of the hosting cluster.  Each type of service reference requires specific configurations and bindings to establish the connection and interaction with the respective services. Note that the ServiceRef has cluster-level semantic consistency, meaning that within the same Cluster, service references with the same ServiceRef.Name are considered to be the same service. It is only allowed to bind to the same Cluster or ServiceDescriptor.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "The name of the KubeBlocks cluster being referenced when a service provided by another KubeBlocks cluster is being referenced.  By default, the clusterDefinition.spec.connectionCredential secret corresponding to the referenced Cluster will be used to bind to the current component. The connection credential secret should include and correspond to the following fields: endpoint, port, username, and password when a KubeBlocks cluster is being referenced.  Under this referencing approach, the ServiceKind and ServiceVersion of service reference declaration defined in the ClusterDefinition will not be validated. If both Cluster and ServiceDescriptor are specified, the Cluster takes precedence.",
														MarkdownDescription: "The name of the KubeBlocks cluster being referenced when a service provided by another KubeBlocks cluster is being referenced.  By default, the clusterDefinition.spec.connectionCredential secret corresponding to the referenced Cluster will be used to bind to the current component. The connection credential secret should include and correspond to the following fields: endpoint, port, username, and password when a KubeBlocks cluster is being referenced.  Under this referencing approach, the ServiceKind and ServiceVersion of service reference declaration defined in the ClusterDefinition will not be validated. If both Cluster and ServiceDescriptor are specified, the Cluster takes precedence.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Specifies the identifier of the service reference declaration. It corresponds to the serviceRefDeclaration name defined in the clusterDefinition.componentDefs[*].serviceRefDeclarations[*].name.",
														MarkdownDescription: "Specifies the identifier of the service reference declaration. It corresponds to the serviceRefDeclaration name defined in the clusterDefinition.componentDefs[*].serviceRefDeclarations[*].name.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object. If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current cluster by default.",
														MarkdownDescription: "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object. If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current cluster by default.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_descriptor": schema.StringAttribute{
														Description:         "The service descriptor of the service provided by external sources.  When referencing a service provided by external sources, the ServiceDescriptor object name is required to establish the service binding. The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKind and serviceVersion defined in the service reference declaration in the ClusterDefinition.  If both Cluster and ServiceDescriptor are specified, the Cluster takes precedence.",
														MarkdownDescription: "The service descriptor of the service provided by external sources.  When referencing a service provided by external sources, the ServiceDescriptor object name is required to establish the service binding. The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKind and serviceVersion defined in the service reference declaration in the ClusterDefinition.  If both Cluster and ServiceDescriptor are specified, the Cluster takes precedence.",
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

										"services": schema.ListNestedAttribute{
											Description:         "Services expose endpoints that can be accessed by clients.",
											MarkdownDescription: "Services expose endpoints that can be accessed by clients.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "If ServiceType is LoadBalancer, cloud provider related parameters can be put here. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
														MarkdownDescription: "If ServiceType is LoadBalancer, cloud provider related parameters can be put here. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the service.",
														MarkdownDescription: "The name of the service.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(15),
														},
													},

													"service_type": schema.StringAttribute{
														Description:         "Determines how the Service is exposed. Valid options are ClusterIP, NodePort, and LoadBalancer.  - 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP.  More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
														MarkdownDescription: "Determines how the Service is exposed. Valid options are ClusterIP, NodePort, and LoadBalancer.  - 'ClusterIP' allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is 'None', no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP.  More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer"),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"switch_policy": schema.SingleNestedAttribute{
											Description:         "Defines the strategy for switchover and failover when workloadType is Replication.",
											MarkdownDescription: "Defines the strategy for switchover and failover when workloadType is Replication.",
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													Description:         "Type specifies the type of switch policy to be applied.",
													MarkdownDescription: "Type specifies the type of switch policy to be applied.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Noop"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls": schema.BoolAttribute{
											Description:         "Enables or disables TLS certs.",
											MarkdownDescription: "Enables or disables TLS certs.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tolerations": schema.MapAttribute{
											Description:         "Attached to tolerate any taint that matches the triple 'key,value,effect' using the matching operator 'operator'.",
											MarkdownDescription: "Attached to tolerate any taint that matches the triple 'key,value,effect' using the matching operator 'operator'.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"update_strategy": schema.StringAttribute{
											Description:         "Defines the update strategy for the component. Not supported.",
											MarkdownDescription: "Defines the update strategy for the component. Not supported.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Serial", "BestEffortParallel", "Parallel"),
											},
										},

										"user_resource_refs": schema.SingleNestedAttribute{
											Description:         "Defines the user-defined volumes.",
											MarkdownDescription: "Defines the user-defined volumes.",
											Attributes: map[string]schema.Attribute{
												"config_map_refs": schema.ListNestedAttribute{
													Description:         "ConfigMapRefs defines the user-defined config maps.",
													MarkdownDescription: "ConfigMapRefs defines the user-defined config maps.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"as_volume_from": schema.ListAttribute{
																Description:         "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
																MarkdownDescription: "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"config_map": schema.SingleNestedAttribute{
																Description:         "ConfigMap specifies the ConfigMap to be mounted as a volume.",
																MarkdownDescription: "ConfigMap specifies the ConfigMap to be mounted as a volume.",
																Attributes: map[string]schema.Attribute{
																	"default_mode": schema.Int64Attribute{
																		Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"items": schema.ListNestedAttribute{
																		Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the key to project.",
																					MarkdownDescription: "key is the key to project.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"mode": schema.Int64Attribute{
																					Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																					MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "optional specify whether the ConfigMap or its keys must be defined",
																		MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_point": schema.StringAttribute{
																Description:         "MountPoint is the filesystem path where the volume will be mounted.",
																MarkdownDescription: "MountPoint is the filesystem path where the volume will be mounted.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(256),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
																MarkdownDescription: "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
																},
															},

															"sub_path": schema.StringAttribute{
																Description:         "SubPath specifies a path within the volume from which to mount.",
																MarkdownDescription: "SubPath specifies a path within the volume from which to mount.",
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

												"secret_refs": schema.ListNestedAttribute{
													Description:         "SecretRefs defines the user-defined secrets.",
													MarkdownDescription: "SecretRefs defines the user-defined secrets.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"as_volume_from": schema.ListAttribute{
																Description:         "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
																MarkdownDescription: "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"mount_point": schema.StringAttribute{
																Description:         "MountPoint is the filesystem path where the volume will be mounted.",
																MarkdownDescription: "MountPoint is the filesystem path where the volume will be mounted.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(256),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
																MarkdownDescription: "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
																},
															},

															"secret": schema.SingleNestedAttribute{
																Description:         "Secret specifies the secret to be mounted as a volume.",
																MarkdownDescription: "Secret specifies the secret to be mounted as a volume.",
																Attributes: map[string]schema.Attribute{
																	"default_mode": schema.Int64Attribute{
																		Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"items": schema.ListNestedAttribute{
																		Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "key is the key to project.",
																					MarkdownDescription: "key is the key to project.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"mode": schema.Int64Attribute{
																					Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																					MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																	"optional": schema.BoolAttribute{
																		Description:         "optional field specify whether the Secret or its keys must be defined",
																		MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"secret_name": schema.StringAttribute{
																		Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																		MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"sub_path": schema.StringAttribute{
																Description:         "SubPath specifies a path within the volume from which to mount.",
																MarkdownDescription: "SubPath specifies a path within the volume from which to mount.",
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

										"volume_claim_templates": schema.ListNestedAttribute{
											Description:         "Provides information for statefulset.spec.volumeClaimTemplates.",
											MarkdownDescription: "Provides information for statefulset.spec.volumeClaimTemplates.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Refers to 'clusterDefinition.spec.componentDefs.containers.volumeMounts.name'.",
														MarkdownDescription: "Refers to 'clusterDefinition.spec.componentDefs.containers.volumeMounts.name'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"spec": schema.SingleNestedAttribute{
														Description:         "Defines the desired characteristics of a volume requested by a pod author.",
														MarkdownDescription: "Defines the desired characteristics of a volume requested by a pod author.",
														Attributes: map[string]schema.Attribute{
															"access_modes": schema.MapAttribute{
																Description:         "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																MarkdownDescription: "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources": schema.SingleNestedAttribute{
																Description:         "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																MarkdownDescription: "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																Attributes: map[string]schema.Attribute{
																	"claims": schema.ListNestedAttribute{
																		Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																		MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																					MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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

																	"limits": schema.MapAttribute{
																		Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"requests": schema.MapAttribute{
																		Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

															"storage_class_name": schema.StringAttribute{
																Description:         "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																MarkdownDescription: "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"volume_mode": schema.StringAttribute{
																Description:         "Defines what type of volume is required by the claim.",
																MarkdownDescription: "Defines what type of volume is required by the claim.",
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
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Specifies the storage of the first componentSpec, if the storage of the first componentSpec is specified, this value will be ignored.",
						MarkdownDescription: "Specifies the storage of the first componentSpec, if the storage of the first componentSpec is specified, this value will be ignored.",
						Attributes: map[string]schema.Attribute{
							"size": schema.StringAttribute{
								Description:         "Specifies the amount of storage the cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Specifies the amount of storage the cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tenancy": schema.StringAttribute{
						Description:         "Describes how pods are distributed across node.",
						MarkdownDescription: "Describes how pods are distributed across node.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("SharedNode", "DedicatedNode"),
						},
					},

					"termination_policy": schema.StringAttribute{
						Description:         "Specifies the cluster termination policy.  - DoNotTerminate will block delete operation. - Halt will delete workload resources such as statefulset, deployment workloads but keep PVCs. - Delete is based on Halt and deletes PVCs. - WipeOut is based on Delete and wipe out all volume snapshots and snapshot data from backup storage location.",
						MarkdownDescription: "Specifies the cluster termination policy.  - DoNotTerminate will block delete operation. - Halt will delete workload resources such as statefulset, deployment workloads but keep PVCs. - Delete is based on Halt and deletes PVCs. - WipeOut is based on Delete and wipe out all volume snapshots and snapshot data from backup storage location.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("DoNotTerminate", "Halt", "Delete", "WipeOut"),
						},
					},

					"tolerations": schema.MapAttribute{
						Description:         "Attached to tolerate any taint that matches the triple 'key,value,effect' using the matching operator 'operator'.",
						MarkdownDescription: "Attached to tolerate any taint that matches the triple 'key,value,effect' using the matching operator 'operator'.",
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
	}
}

func (r *AppsKubeblocksIoClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_cluster_v1alpha1_manifest")

	var model AppsKubeblocksIoClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
