/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package zookeeper_stackable_tech_v1alpha1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ZookeeperStackableTechZookeeperClusterV1Alpha1Manifest{}
)

func NewZookeeperStackableTechZookeeperClusterV1Alpha1Manifest() datasource.DataSource {
	return &ZookeeperStackableTechZookeeperClusterV1Alpha1Manifest{}
}

type ZookeeperStackableTechZookeeperClusterV1Alpha1Manifest struct{}

type ZookeeperStackableTechZookeeperClusterV1Alpha1ManifestData struct {
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
		ClusterConfig *struct {
			Authentication *[]struct {
				AuthenticationClass *string `tfsdk:"authentication_class" json:"authenticationClass,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			ListenerClass *string `tfsdk:"listener_class" json:"listenerClass,omitempty"`
			Tls           *struct {
				QuorumSecretClass *string `tfsdk:"quorum_secret_class" json:"quorumSecretClass,omitempty"`
				ServerSecretClass *string `tfsdk:"server_secret_class" json:"serverSecretClass,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			VectorAggregatorConfigMapName *string `tfsdk:"vector_aggregator_config_map_name" json:"vectorAggregatorConfigMapName,omitempty"`
		} `tfsdk:"cluster_config" json:"clusterConfig,omitempty"`
		ClusterOperation *struct {
			ReconciliationPaused *bool `tfsdk:"reconciliation_paused" json:"reconciliationPaused,omitempty"`
			Stopped              *bool `tfsdk:"stopped" json:"stopped,omitempty"`
		} `tfsdk:"cluster_operation" json:"clusterOperation,omitempty"`
		Image *struct {
			Custom         *string `tfsdk:"custom" json:"custom,omitempty"`
			ProductVersion *string `tfsdk:"product_version" json:"productVersion,omitempty"`
			PullPolicy     *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
			PullSecrets    *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
			Repo             *string `tfsdk:"repo" json:"repo,omitempty"`
			StackableVersion *string `tfsdk:"stackable_version" json:"stackableVersion,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		Servers *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
				InitLimit               *int64  `tfsdk:"init_limit" json:"initLimit,omitempty"`
				Logging                 *struct {
					Containers *struct {
						Console *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"console" json:"console,omitempty"`
						Custom *struct {
							ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
						} `tfsdk:"custom" json:"custom,omitempty"`
						File *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"file" json:"file,omitempty"`
						Loggers *struct {
							Level *string `tfsdk:"level" json:"level,omitempty"`
						} `tfsdk:"loggers" json:"loggers,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					EnableVectorAgent *bool `tfsdk:"enable_vector_agent" json:"enableVectorAgent,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				MyidOffset *int64 `tfsdk:"myid_offset" json:"myidOffset,omitempty"`
				Resources  *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *struct {
						Data *struct {
							Capacity  *string `tfsdk:"capacity" json:"capacity,omitempty"`
							Selectors *struct {
								MatchExpressions *[]struct {
									Key      *string   `tfsdk:"key" json:"key,omitempty"`
									Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
									Values   *[]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
								MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
							} `tfsdk:"selectors" json:"selectors,omitempty"`
							StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
						} `tfsdk:"data" json:"data,omitempty"`
					} `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				SyncLimit *int64 `tfsdk:"sync_limit" json:"syncLimit,omitempty"`
				TickTime  *int64 `tfsdk:"tick_time" json:"tickTime,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
			EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
			PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
			RoleConfig      *struct {
				PodDisruptionBudget *struct {
					Enabled        *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					MaxUnavailable *int64 `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			} `tfsdk:"role_config" json:"roleConfig,omitempty"`
			RoleGroups *struct {
				CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
				Config       *struct {
					Affinity *struct {
						NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
						NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
						PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
						PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" json:"affinity,omitempty"`
					GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
					InitLimit               *int64  `tfsdk:"init_limit" json:"initLimit,omitempty"`
					Logging                 *struct {
						Containers *struct {
							Console *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"console" json:"console,omitempty"`
							Custom *struct {
								ConfigMap *string `tfsdk:"config_map" json:"configMap,omitempty"`
							} `tfsdk:"custom" json:"custom,omitempty"`
							File *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"file" json:"file,omitempty"`
							Loggers *struct {
								Level *string `tfsdk:"level" json:"level,omitempty"`
							} `tfsdk:"loggers" json:"loggers,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						EnableVectorAgent *bool `tfsdk:"enable_vector_agent" json:"enableVectorAgent,omitempty"`
					} `tfsdk:"logging" json:"logging,omitempty"`
					MyidOffset *int64 `tfsdk:"myid_offset" json:"myidOffset,omitempty"`
					Resources  *struct {
						Cpu *struct {
							Max *string `tfsdk:"max" json:"max,omitempty"`
							Min *string `tfsdk:"min" json:"min,omitempty"`
						} `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *struct {
							Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
							RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
						} `tfsdk:"memory" json:"memory,omitempty"`
						Storage *struct {
							Data *struct {
								Capacity  *string `tfsdk:"capacity" json:"capacity,omitempty"`
								Selectors *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"selectors" json:"selectors,omitempty"`
								StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
							} `tfsdk:"data" json:"data,omitempty"`
						} `tfsdk:"storage" json:"storage,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					SyncLimit *int64 `tfsdk:"sync_limit" json:"syncLimit,omitempty"`
					TickTime  *int64 `tfsdk:"tick_time" json:"tickTime,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
				EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
				PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
				Replicas        *int64                        `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"role_groups" json:"roleGroups,omitempty"`
		} `tfsdk:"servers" json:"servers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ZookeeperStackableTechZookeeperClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_zookeeper_stackable_tech_zookeeper_cluster_v1alpha1_manifest"
}

func (r *ZookeeperStackableTechZookeeperClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for ZookeeperClusterSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for ZookeeperClusterSpec via 'CustomResource'",
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
				Description:         "A ZooKeeper cluster stacklet. This resource is managed by the Stackable operator for Apache ZooKeeper. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/zookeeper/).",
				MarkdownDescription: "A ZooKeeper cluster stacklet. This resource is managed by the Stackable operator for Apache ZooKeeper. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/zookeeper/).",
				Attributes: map[string]schema.Attribute{
					"cluster_config": schema.SingleNestedAttribute{
						Description:         "Settings that affect all roles and role groups. The settings in the 'clusterConfig' are cluster wide settings that do not need to be configurable at role or role group level.",
						MarkdownDescription: "Settings that affect all roles and role groups. The settings in the 'clusterConfig' are cluster wide settings that do not need to be configurable at role or role group level.",
						Attributes: map[string]schema.Attribute{
							"authentication": schema.ListNestedAttribute{
								Description:         "Authentication settings for ZooKeeper like mTLS authentication. Read more in the [authentication usage guide](https://docs.stackable.tech/home/nightly/zookeeper/usage_guide/authentication).",
								MarkdownDescription: "Authentication settings for ZooKeeper like mTLS authentication. Read more in the [authentication usage guide](https://docs.stackable.tech/home/nightly/zookeeper/usage_guide/authentication).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authentication_class": schema.StringAttribute{
											Description:         "The [AuthenticationClass](https://docs.stackable.tech/home/stable/concepts/authentication) to use. ## mTLS Only affects client connections. This setting controls: - If clients need to authenticate themselves against the server via TLS - Which ca.crt to use when validating the provided client certs This will override the server TLS settings (if set) in 'spec.clusterConfig.tls.serverSecretClass'.",
											MarkdownDescription: "The [AuthenticationClass](https://docs.stackable.tech/home/stable/concepts/authentication) to use. ## mTLS Only affects client connections. This setting controls: - If clients need to authenticate themselves against the server via TLS - Which ca.crt to use when validating the provided client certs This will override the server TLS settings (if set) in 'spec.clusterConfig.tls.serverSecretClass'.",
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

							"listener_class": schema.StringAttribute{
								Description:         "This field controls which type of Service the Operator creates for this ZookeeperCluster: * cluster-internal: Use a ClusterIP service * external-unstable: Use a NodePort service This is a temporary solution with the goal to keep yaml manifests forward compatible. In the future, this setting will control which [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass.html) will be used to expose the service, and ListenerClass names will stay the same, allowing for a non-breaking change.",
								MarkdownDescription: "This field controls which type of Service the Operator creates for this ZookeeperCluster: * cluster-internal: Use a ClusterIP service * external-unstable: Use a NodePort service This is a temporary solution with the goal to keep yaml manifests forward compatible. In the future, this setting will control which [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass.html) will be used to expose the service, and ListenerClass names will stay the same, allowing for a non-breaking change.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("cluster-internal", "external-unstable"),
								},
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS encryption settings for ZooKeeper (server, quorum). Read more in the [encryption usage guide](https://docs.stackable.tech/home/nightly/zookeeper/usage_guide/encryption).",
								MarkdownDescription: "TLS encryption settings for ZooKeeper (server, quorum). Read more in the [encryption usage guide](https://docs.stackable.tech/home/nightly/zookeeper/usage_guide/encryption).",
								Attributes: map[string]schema.Attribute{
									"quorum_secret_class": schema.StringAttribute{
										Description:         "The [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) to use for internal quorum communication. Use mutual verification between Zookeeper Nodes (mandatory). This setting controls: - Which cert the servers should use to authenticate themselves against other servers - Which ca.crt to use when validating the other server Defaults to 'tls'",
										MarkdownDescription: "The [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) to use for internal quorum communication. Use mutual verification between Zookeeper Nodes (mandatory). This setting controls: - Which cert the servers should use to authenticate themselves against other servers - Which ca.crt to use when validating the other server Defaults to 'tls'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server_secret_class": schema.StringAttribute{
										Description:         "The [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) to use for client connections. This setting controls: - If TLS encryption is used at all - Which cert the servers should use to authenticate themselves against the client Defaults to 'tls'.",
										MarkdownDescription: "The [SecretClass](https://docs.stackable.tech/home/nightly/secret-operator/secretclass) to use for client connections. This setting controls: - If TLS encryption is used at all - Which cert the servers should use to authenticate themselves against the client Defaults to 'tls'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"vector_aggregator_config_map_name": schema.StringAttribute{
								Description:         "Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.",
								MarkdownDescription: "Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_operation": schema.SingleNestedAttribute{
						Description:         "[Cluster operations](https://docs.stackable.tech/home/nightly/concepts/operations/cluster_operations) properties, allow stopping the product instance as well as pausing reconciliation.",
						MarkdownDescription: "[Cluster operations](https://docs.stackable.tech/home/nightly/concepts/operations/cluster_operations) properties, allow stopping the product instance as well as pausing reconciliation.",
						Attributes: map[string]schema.Attribute{
							"reconciliation_paused": schema.BoolAttribute{
								Description:         "Flag to stop cluster reconciliation by the operator. This means that all changes in the custom resource spec are ignored until this flag is set to false or removed. The operator will however still watch the deployed resources at the time and update the custom resource status field. If applied at the same time with 'stopped', 'reconciliationPaused' will take precedence over 'stopped' and stop the reconciliation immediately.",
								MarkdownDescription: "Flag to stop cluster reconciliation by the operator. This means that all changes in the custom resource spec are ignored until this flag is set to false or removed. The operator will however still watch the deployed resources at the time and update the custom resource status field. If applied at the same time with 'stopped', 'reconciliationPaused' will take precedence over 'stopped' and stop the reconciliation immediately.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stopped": schema.BoolAttribute{
								Description:         "Flag to stop the cluster. This means all deployed resources (e.g. Services, StatefulSets, ConfigMaps) are kept but all deployed Pods (e.g. replicas from a StatefulSet) are scaled to 0 and therefore stopped and removed. If applied at the same time with 'reconciliationPaused', the latter will pause reconciliation and 'stopped' will take no effect until 'reconciliationPaused' is set to false or removed.",
								MarkdownDescription: "Flag to stop the cluster. This means all deployed resources (e.g. Services, StatefulSets, ConfigMaps) are kept but all deployed Pods (e.g. replicas from a StatefulSet) are scaled to 0 and therefore stopped and removed. If applied at the same time with 'reconciliationPaused', the latter will pause reconciliation and 'stopped' will take no effect until 'reconciliationPaused' is set to false or removed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "Specify which image to use, the easiest way is to only configure the 'productVersion'. You can also configure a custom image registry to pull from, as well as completely custom images. Consult the [Product image selection documentation](https://docs.stackable.tech/home/nightly/concepts/product_image_selection) for details.",
						MarkdownDescription: "Specify which image to use, the easiest way is to only configure the 'productVersion'. You can also configure a custom image registry to pull from, as well as completely custom images. Consult the [Product image selection documentation](https://docs.stackable.tech/home/nightly/concepts/product_image_selection) for details.",
						Attributes: map[string]schema.Attribute{
							"custom": schema.StringAttribute{
								Description:         "Overwrite the docker image. Specify the full docker image name, e.g. 'docker.stackable.tech/stackable/superset:1.4.1-stackable2.1.0'",
								MarkdownDescription: "Overwrite the docker image. Specify the full docker image name, e.g. 'docker.stackable.tech/stackable/superset:1.4.1-stackable2.1.0'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"product_version": schema.StringAttribute{
								Description:         "Version of the product, e.g. '1.4.1'.",
								MarkdownDescription: "Version of the product, e.g. '1.4.1'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pull_policy": schema.StringAttribute{
								Description:         "[Pull policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) used when pulling the image.",
								MarkdownDescription: "[Pull policy](https://kubernetes.io/docs/concepts/containers/images/#image-pull-policy) used when pulling the image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
								},
							},

							"pull_secrets": schema.ListNestedAttribute{
								Description:         "[Image pull secrets](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod) to pull images from a private registry.",
								MarkdownDescription: "[Image pull secrets](https://kubernetes.io/docs/concepts/containers/images/#specifying-imagepullsecrets-on-a-pod) to pull images from a private registry.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

							"repo": schema.StringAttribute{
								Description:         "Name of the docker repo, e.g. 'docker.stackable.tech/stackable'",
								MarkdownDescription: "Name of the docker repo, e.g. 'docker.stackable.tech/stackable'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"stackable_version": schema.StringAttribute{
								Description:         "Stackable version of the product, e.g. '23.4', '23.4.1' or '0.0.0-dev'. If not specified, the operator will use its own version, e.g. '23.4.1'. When using a nightly operator or a pr version, it will use the nightly '0.0.0-dev' image.",
								MarkdownDescription: "Stackable version of the product, e.g. '23.4', '23.4.1' or '0.0.0-dev'. If not specified, the operator will use its own version, e.g. '23.4.1'. When using a nightly operator or a pr version, it will use the nightly '0.0.0-dev' image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"servers": schema.SingleNestedAttribute{
						Description:         "This struct represents a role - e.g. HDFS datanodes or Trino workers. It has a key-value-map containing all the roleGroups that are part of this role. Additionally, there is a 'config', which is configurable at the role *and* roleGroup level. Everything at roleGroup level is merged on top of what is configured on role level. There is also a second form of config, which can only be configured at role level, the 'roleConfig'. You can learn more about this in the [Roles and role group concept documentation](https://docs.stackable.tech/home/nightly/concepts/roles-and-role-groups).",
						MarkdownDescription: "This struct represents a role - e.g. HDFS datanodes or Trino workers. It has a key-value-map containing all the roleGroups that are part of this role. Additionally, there is a 'config', which is configurable at the role *and* roleGroup level. Everything at roleGroup level is merged on top of what is configured on role level. There is also a second form of config, which can only be configured at role level, the 'roleConfig'. You can learn more about this in the [Roles and role group concept documentation](https://docs.stackable.tech/home/nightly/concepts/roles-and-role-groups).",
						Attributes: map[string]schema.Attribute{
							"cli_overrides": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										MarkdownDescription: "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
										Attributes: map[string]schema.Attribute{
											"node_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_selector": schema.MapAttribute{
												Description:         "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity": schema.MapAttribute{
												Description:         "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
												MarkdownDescription: "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
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

									"graceful_shutdown_timeout": schema.StringAttribute{
										Description:         "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
										MarkdownDescription: "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"init_limit": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"logging": schema.SingleNestedAttribute{
										Description:         "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										MarkdownDescription: "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
										Attributes: map[string]schema.Attribute{
											"containers": schema.SingleNestedAttribute{
												Description:         "Log configuration per container.",
												MarkdownDescription: "Log configuration per container.",
												Attributes: map[string]schema.Attribute{
													"console": schema.SingleNestedAttribute{
														Description:         "Configuration for the console appender",
														MarkdownDescription: "Configuration for the console appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"custom": schema.SingleNestedAttribute{
														Description:         "Custom log configuration provided in a ConfigMap",
														MarkdownDescription: "Custom log configuration provided in a ConfigMap",
														Attributes: map[string]schema.Attribute{
															"config_map": schema.StringAttribute{
																Description:         "ConfigMap containing the log configuration files",
																MarkdownDescription: "ConfigMap containing the log configuration files",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"file": schema.SingleNestedAttribute{
														Description:         "Configuration for the file appender",
														MarkdownDescription: "Configuration for the file appender",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"loggers": schema.SingleNestedAttribute{
														Description:         "Configuration per logger",
														MarkdownDescription: "Configuration per logger",
														Attributes: map[string]schema.Attribute{
															"level": schema.StringAttribute{
																Description:         "The log level threshold. Log events with a lower log level are discarded.",
																MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
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

											"enable_vector_agent": schema.BoolAttribute{
												Description:         "Wether or not to deploy a container with the Vector log agent.",
												MarkdownDescription: "Wether or not to deploy a container with the Vector log agent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"myid_offset": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
										Attributes: map[string]schema.Attribute{
											"cpu": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max": schema.StringAttribute{
														Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min": schema.StringAttribute{
														Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"limit": schema.StringAttribute{
														Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"runtime_limits": schema.MapAttribute{
														Description:         "Additional options that can be specified.",
														MarkdownDescription: "Additional options that can be specified.",
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

											"storage": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"data": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"capacity": schema.StringAttribute{
																Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ''' <quantity> ::= <signedNumber><suffix> (Note that <suffix> may be empty, from the '' case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= '+' | '-' <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | '' | k | M | G | T | P | E (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> ''' No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi' Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
																MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ''' <quantity> ::= <signedNumber><suffix> (Note that <suffix> may be empty, from the '' case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= '+' | '-' <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | '' | k | M | G | T | P | E (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> ''' No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi' Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"selectors": schema.SingleNestedAttribute{
																Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
																MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

															"storage_class": schema.StringAttribute{
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
												Required: false,
												Optional: true,
												Computed: false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"sync_limit": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"tick_time": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"config_overrides": schema.MapAttribute{
								Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"env_overrides": schema.MapAttribute{
								Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_overrides": schema.MapAttribute{
								Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_config": schema.SingleNestedAttribute{
								Description:         "This is a product-agnostic RoleConfig, which is sufficient for most of the products.",
								MarkdownDescription: "This is a product-agnostic RoleConfig, which is sufficient for most of the products.",
								Attributes: map[string]schema.Attribute{
									"pod_disruption_budget": schema.SingleNestedAttribute{
										Description:         "This struct is used to configure: 1. If PodDisruptionBudgets are created by the operator 2. The allowed number of Pods to be unavailable ('maxUnavailable') Learn more in the [allowed Pod disruptions documentation](https://docs.stackable.tech/home/nightly/concepts/operations/pod_disruptions).",
										MarkdownDescription: "This struct is used to configure: 1. If PodDisruptionBudgets are created by the operator 2. The allowed number of Pods to be unavailable ('maxUnavailable') Learn more in the [allowed Pod disruptions documentation](https://docs.stackable.tech/home/nightly/concepts/operations/pod_disruptions).",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether a PodDisruptionBudget should be written out for this role. Disabling this enables you to specify your own - custom - one. Defaults to true.",
												MarkdownDescription: "Whether a PodDisruptionBudget should be written out for this role. Disabling this enables you to specify your own - custom - one. Defaults to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_unavailable": schema.Int64Attribute{
												Description:         "The number of Pods that are allowed to be down because of voluntary disruptions. If you don't explicitly set this, the operator will use a sane default based upon knowledge about the individual product.",
												MarkdownDescription: "The number of Pods that are allowed to be down because of voluntary disruptions. If you don't explicitly set this, the operator will use a sane default based upon knowledge about the individual product.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
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

							"role_groups": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cli_overrides": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"affinity": schema.SingleNestedAttribute{
												Description:         "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
												MarkdownDescription: "These configuration settings control [Pod placement](https://docs.stackable.tech/home/nightly/concepts/operations/pod_placement).",
												Attributes: map[string]schema.Attribute{
													"node_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.nodeAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_selector": schema.MapAttribute{
														Description:         "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Simple key-value pairs forming a nodeSelector, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.podAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_anti_affinity": schema.MapAttribute{
														Description:         "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
														MarkdownDescription: "Same as the 'spec.affinity.podAntiAffinity' field on the Pod, see the [Kubernetes docs](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node)",
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

											"graceful_shutdown_timeout": schema.StringAttribute{
												Description:         "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
												MarkdownDescription: "Time period Pods have to gracefully shut down, e.g. '30m', '1h' or '2d'. Consult the operator documentation for details.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"init_limit": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"logging": schema.SingleNestedAttribute{
												Description:         "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
												MarkdownDescription: "Logging configuration, learn more in the [logging concept documentation](https://docs.stackable.tech/home/nightly/concepts/logging).",
												Attributes: map[string]schema.Attribute{
													"containers": schema.SingleNestedAttribute{
														Description:         "Log configuration per container.",
														MarkdownDescription: "Log configuration per container.",
														Attributes: map[string]schema.Attribute{
															"console": schema.SingleNestedAttribute{
																Description:         "Configuration for the console appender",
																MarkdownDescription: "Configuration for the console appender",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"custom": schema.SingleNestedAttribute{
																Description:         "Custom log configuration provided in a ConfigMap",
																MarkdownDescription: "Custom log configuration provided in a ConfigMap",
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.StringAttribute{
																		Description:         "ConfigMap containing the log configuration files",
																		MarkdownDescription: "ConfigMap containing the log configuration files",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"file": schema.SingleNestedAttribute{
																Description:         "Configuration for the file appender",
																MarkdownDescription: "Configuration for the file appender",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"loggers": schema.SingleNestedAttribute{
																Description:         "Configuration per logger",
																MarkdownDescription: "Configuration per logger",
																Attributes: map[string]schema.Attribute{
																	"level": schema.StringAttribute{
																		Description:         "The log level threshold. Log events with a lower log level are discarded.",
																		MarkdownDescription: "The log level threshold. Log events with a lower log level are discarded.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL", "NONE"),
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

													"enable_vector_agent": schema.BoolAttribute{
														Description:         "Wether or not to deploy a container with the Vector log agent.",
														MarkdownDescription: "Wether or not to deploy a container with the Vector log agent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"myid_offset": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
												MarkdownDescription: "Resource usage is configured here, this includes CPU usage, memory usage and disk storage usage, if this role needs any.",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"max": schema.StringAttribute{
																Description:         "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																MarkdownDescription: "The maximum amount of CPU cores that can be requested by Pods. Equivalent to the 'limit' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"min": schema.StringAttribute{
																Description:         "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																MarkdownDescription: "The minimal amount of CPU cores that Pods need to run. Equivalent to the 'request' for Pod resource configuration. Cores are specified either as a decimal point number or as milli units. For example:'1.5' will be 1.5 cores, also written as '1500m'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"memory": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"limit": schema.StringAttribute{
																Description:         "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
																MarkdownDescription: "The maximum amount of memory that should be available to the Pod. Specified as a byte [Quantity](https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/), which means these suffixes are supported: E, P, T, G, M, k. You can also use the power-of-two equivalents: Ei, Pi, Ti, Gi, Mi, Ki. For example, the following represent roughly the same value: '128974848, 129e6, 129M, 128974848000m, 123Mi'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"runtime_limits": schema.MapAttribute{
																Description:         "Additional options that can be specified.",
																MarkdownDescription: "Additional options that can be specified.",
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

													"storage": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"data": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"capacity": schema.StringAttribute{
																		Description:         "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ''' <quantity> ::= <signedNumber><suffix> (Note that <suffix> may be empty, from the '' case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= '+' | '-' <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | '' | k | M | G | T | P | E (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> ''' No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi' Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
																		MarkdownDescription: "Quantity is a fixed-point representation of a number. It provides convenient marshaling/unmarshaling in JSON and YAML, in addition to String() and AsInt64() accessors. The serialization format is: ''' <quantity> ::= <signedNumber><suffix> (Note that <suffix> may be empty, from the '' case in <decimalSI>.) <digit> ::= 0 | 1 | ... | 9 <digits> ::= <digit> | <digit><digits> <number> ::= <digits> | <digits>.<digits> | <digits>. | .<digits> <sign> ::= '+' | '-' <signedNumber> ::= <number> | <sign><number> <suffix> ::= <binarySI> | <decimalExponent> | <decimalSI> <binarySI> ::= Ki | Mi | Gi | Ti | Pi | Ei (International System of units; See: http://physics.nist.gov/cuu/Units/binary.html) <decimalSI> ::= m | '' | k | M | G | T | P | E (Note that 1024 = 1Ki but 1000 = 1k; I didn't choose the capitalization.) <decimalExponent> ::= 'e' <signedNumber> | 'E' <signedNumber> ''' No matter which of the three exponent forms is used, no quantity may represent a number greater than 2^63-1 in magnitude, nor may it have more than 3 decimal places. Numbers larger or more precise will be capped or rounded up. (E.g.: 0.1m will rounded up to 1m.) This may be extended in the future if we require larger or smaller quantities. When a Quantity is parsed from a string, it will remember the type of suffix it had, and will use the same type again when it is serialized. Before serializing, Quantity will be put in 'canonical form'. This means that Exponent/suffix will be adjusted up or down (with a corresponding increase or decrease in Mantissa) such that: - No precision is lost - No fractional digits will be emitted - The exponent (or suffix) is as large as possible. The sign will be omitted unless the number is negative. Examples: - 1.5 will be serialized as '1500m' - 1.5Gi will be serialized as '1536Mi' Note that the quantity will NEVER be internally represented by a floating point number. That is the whole point of this exercise. Non-canonical values will still parse as long as they are well formed, but will be re-emitted in their canonical form. (So always use canonical form, or don't diff.) This format is intended to make it difficult to use these numbers without writing some sort of special handling code in the hopes that that will cause implementors to also use a fixed point implementation.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"selectors": schema.SingleNestedAttribute{
																		Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
																		MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

																	"storage_class": schema.StringAttribute{
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
														Required: false,
														Optional: true,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"sync_limit": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"tick_time": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"config_overrides": schema.MapAttribute{
										Description:         "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
										MarkdownDescription: "The 'configOverrides' can be used to configure properties in product config files that are not exposed in the CRD. Read the [config overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#config-overrides) and consult the operator specific usage guide documentation for details on the available config files and settings for the specific product.",
										ElementType:         types.MapType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env_overrides": schema.MapAttribute{
										Description:         "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
										MarkdownDescription: "'envOverrides' configure environment variables to be set in the Pods. It is a map from strings to strings - environment variables and the value to set. Read the [environment variable overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#env-overrides) for more information and consult the operator specific usage guide to find out about the product specific environment variables that are available.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_overrides": schema.MapAttribute{
										Description:         "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
										MarkdownDescription: "In the 'podOverrides' property you can define a [PodTemplateSpec](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#podtemplatespec-v1-core) to override any property that can be set on a Kubernetes Pod. Read the [Pod overrides documentation](https://docs.stackable.tech/home/nightly/concepts/overrides#pod-overrides) for more information.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"replicas": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *ZookeeperStackableTechZookeeperClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_zookeeper_stackable_tech_zookeeper_cluster_v1alpha1_manifest")

	var model ZookeeperStackableTechZookeeperClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("zookeeper.stackable.tech/v1alpha1")
	model.Kind = pointer.String("ZookeeperCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
