/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package airflow_stackable_tech_v1alpha1

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
	_ datasource.DataSource = &AirflowStackableTechAirflowClusterV1Alpha1Manifest{}
)

func NewAirflowStackableTechAirflowClusterV1Alpha1Manifest() datasource.DataSource {
	return &AirflowStackableTechAirflowClusterV1Alpha1Manifest{}
}

type AirflowStackableTechAirflowClusterV1Alpha1Manifest struct{}

type AirflowStackableTechAirflowClusterV1Alpha1ManifestData struct {
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
		CeleryExecutors *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
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
				Resources *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
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
					Resources *struct {
						Cpu *struct {
							Max *string `tfsdk:"max" json:"max,omitempty"`
							Min *string `tfsdk:"min" json:"min,omitempty"`
						} `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *struct {
							Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
							RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
						} `tfsdk:"memory" json:"memory,omitempty"`
						Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
				EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
				PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
				Replicas        *int64                        `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"role_groups" json:"roleGroups,omitempty"`
		} `tfsdk:"celery_executors" json:"celeryExecutors,omitempty"`
		ClusterConfig *struct {
			Authentication *[]struct {
				AuthenticationClass  *string `tfsdk:"authentication_class" json:"authenticationClass,omitempty"`
				SyncRolesAt          *string `tfsdk:"sync_roles_at" json:"syncRolesAt,omitempty"`
				UserRegistration     *bool   `tfsdk:"user_registration" json:"userRegistration,omitempty"`
				UserRegistrationRole *string `tfsdk:"user_registration_role" json:"userRegistrationRole,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			CredentialsSecret *string `tfsdk:"credentials_secret" json:"credentialsSecret,omitempty"`
			DagsGitSync       *[]struct {
				Branch            *string            `tfsdk:"branch" json:"branch,omitempty"`
				CredentialsSecret *string            `tfsdk:"credentials_secret" json:"credentialsSecret,omitempty"`
				Depth             *int64             `tfsdk:"depth" json:"depth,omitempty"`
				GitFolder         *string            `tfsdk:"git_folder" json:"gitFolder,omitempty"`
				GitSyncConf       *map[string]string `tfsdk:"git_sync_conf" json:"gitSyncConf,omitempty"`
				Repo              *string            `tfsdk:"repo" json:"repo,omitempty"`
				Wait              *int64             `tfsdk:"wait" json:"wait,omitempty"`
			} `tfsdk:"dags_git_sync" json:"dagsGitSync,omitempty"`
			ExposeConfig                  *bool                `tfsdk:"expose_config" json:"exposeConfig,omitempty"`
			ListenerClass                 *string              `tfsdk:"listener_class" json:"listenerClass,omitempty"`
			LoadExamples                  *bool                `tfsdk:"load_examples" json:"loadExamples,omitempty"`
			VectorAggregatorConfigMapName *string              `tfsdk:"vector_aggregator_config_map_name" json:"vectorAggregatorConfigMapName,omitempty"`
			VolumeMounts                  *[]map[string]string `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			Volumes                       *[]map[string]string `tfsdk:"volumes" json:"volumes,omitempty"`
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
		KubernetesExecutors *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
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
				Resources *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
			EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
			PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
		} `tfsdk:"kubernetes_executors" json:"kubernetesExecutors,omitempty"`
		Schedulers *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
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
				Resources *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
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
					Resources *struct {
						Cpu *struct {
							Max *string `tfsdk:"max" json:"max,omitempty"`
							Min *string `tfsdk:"min" json:"min,omitempty"`
						} `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *struct {
							Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
							RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
						} `tfsdk:"memory" json:"memory,omitempty"`
						Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
				EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
				PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
				Replicas        *int64                        `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"role_groups" json:"roleGroups,omitempty"`
		} `tfsdk:"schedulers" json:"schedulers,omitempty"`
		Webservers *struct {
			CliOverrides *map[string]string `tfsdk:"cli_overrides" json:"cliOverrides,omitempty"`
			Config       *struct {
				Affinity *struct {
					NodeAffinity    *map[string]string `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodAffinity     *map[string]string `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAntiAffinity *map[string]string `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
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
				Resources *struct {
					Cpu *struct {
						Max *string `tfsdk:"max" json:"max,omitempty"`
						Min *string `tfsdk:"min" json:"min,omitempty"`
					} `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *struct {
						Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
						RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
					} `tfsdk:"memory" json:"memory,omitempty"`
					Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
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
					Resources *struct {
						Cpu *struct {
							Max *string `tfsdk:"max" json:"max,omitempty"`
							Min *string `tfsdk:"min" json:"min,omitempty"`
						} `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *struct {
							Limit         *string            `tfsdk:"limit" json:"limit,omitempty"`
							RuntimeLimits *map[string]string `tfsdk:"runtime_limits" json:"runtimeLimits,omitempty"`
						} `tfsdk:"memory" json:"memory,omitempty"`
						Storage *map[string]string `tfsdk:"storage" json:"storage,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"config" json:"config,omitempty"`
				ConfigOverrides *map[string]map[string]string `tfsdk:"config_overrides" json:"configOverrides,omitempty"`
				EnvOverrides    *map[string]string            `tfsdk:"env_overrides" json:"envOverrides,omitempty"`
				PodOverrides    *map[string]string            `tfsdk:"pod_overrides" json:"podOverrides,omitempty"`
				Replicas        *int64                        `tfsdk:"replicas" json:"replicas,omitempty"`
			} `tfsdk:"role_groups" json:"roleGroups,omitempty"`
		} `tfsdk:"webservers" json:"webservers,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AirflowStackableTechAirflowClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_airflow_stackable_tech_airflow_cluster_v1alpha1_manifest"
}

func (r *AirflowStackableTechAirflowClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for AirflowClusterSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for AirflowClusterSpec via 'CustomResource'",
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
				Description:         "An Airflow cluster stacklet. This resource is managed by the Stackable operator for Apache Airflow. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/airflow/). The CRD contains three roles: webserver, scheduler and worker/celeryExecutor. You can use either the celeryExecutor or the kubernetesExecutor.",
				MarkdownDescription: "An Airflow cluster stacklet. This resource is managed by the Stackable operator for Apache Airflow. Find more information on how to use it and the resources that the operator generates in the [operator documentation](https://docs.stackable.tech/home/nightly/airflow/). The CRD contains three roles: webserver, scheduler and worker/celeryExecutor. You can use either the celeryExecutor or the kubernetesExecutor.",
				Attributes: map[string]schema.Attribute{
					"celery_executors": schema.SingleNestedAttribute{
						Description:         "The celery executor. Deployed with an explicit number of replicas.",
						MarkdownDescription: "The celery executor. Deployed with an explicit number of replicas.",
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

											"storage": schema.MapAttribute{
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

													"storage": schema.MapAttribute{
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

					"cluster_config": schema.SingleNestedAttribute{
						Description:         "Configuration that applies to all roles and role groups. This includes settings for authentication, git sync, service exposition and volumes, among other things.",
						MarkdownDescription: "Configuration that applies to all roles and role groups. This includes settings for authentication, git sync, service exposition and volumes, among other things.",
						Attributes: map[string]schema.Attribute{
							"authentication": schema.ListNestedAttribute{
								Description:         "The Airflow [authentication](https://docs.stackable.tech/home/nightly/airflow/usage-guide/security.html) settings. Currently the underlying Flask App Builder only supports one authentication mechanism at a time. This means the operator will error out if multiple references to an AuthenticationClass are provided.",
								MarkdownDescription: "The Airflow [authentication](https://docs.stackable.tech/home/nightly/airflow/usage-guide/security.html) settings. Currently the underlying Flask App Builder only supports one authentication mechanism at a time. This means the operator will error out if multiple references to an AuthenticationClass are provided.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authentication_class": schema.StringAttribute{
											Description:         "Name of the [AuthenticationClass](https://docs.stackable.tech/home/nightly/concepts/authentication.html#authenticationclass) used to authenticate the users. At the moment only LDAP is supported. If not specified the default authentication (AUTH_DB) will be used.",
											MarkdownDescription: "Name of the [AuthenticationClass](https://docs.stackable.tech/home/nightly/concepts/authentication.html#authenticationclass) used to authenticate the users. At the moment only LDAP is supported. If not specified the default authentication (AUTH_DB) will be used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sync_roles_at": schema.StringAttribute{
											Description:         "If we should replace ALL the user's roles each login, or only on registration. Gets mapped to 'AUTH_ROLES_SYNC_AT_LOGIN'",
											MarkdownDescription: "If we should replace ALL the user's roles each login, or only on registration. Gets mapped to 'AUTH_ROLES_SYNC_AT_LOGIN'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Registration", "Login"),
											},
										},

										"user_registration": schema.BoolAttribute{
											Description:         "Allow users who are not already in the FAB DB. Gets mapped to 'AUTH_USER_REGISTRATION'",
											MarkdownDescription: "Allow users who are not already in the FAB DB. Gets mapped to 'AUTH_USER_REGISTRATION'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"user_registration_role": schema.StringAttribute{
											Description:         "This role will be given in addition to any AUTH_ROLES_MAPPING. Gets mapped to 'AUTH_USER_REGISTRATION_ROLE'",
											MarkdownDescription: "This role will be given in addition to any AUTH_ROLES_MAPPING. Gets mapped to 'AUTH_USER_REGISTRATION_ROLE'",
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

							"credentials_secret": schema.StringAttribute{
								Description:         "The name of the Secret object containing the admin user credentials and database connection details. Read the [getting started guide first steps](https://docs.stackable.tech/home/nightly/airflow/getting_started/first_steps) to find out more.",
								MarkdownDescription: "The name of the Secret object containing the admin user credentials and database connection details. Read the [getting started guide first steps](https://docs.stackable.tech/home/nightly/airflow/getting_started/first_steps) to find out more.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"dags_git_sync": schema.ListNestedAttribute{
								Description:         "The 'gitSync' settings allow configuring DAGs to mount via 'git-sync'. Learn more in the [mounting DAGs documentation](https://docs.stackable.tech/home/nightly/airflow/usage-guide/mounting-dags#_via_git_sync).",
								MarkdownDescription: "The 'gitSync' settings allow configuring DAGs to mount via 'git-sync'. Learn more in the [mounting DAGs documentation](https://docs.stackable.tech/home/nightly/airflow/usage-guide/mounting-dags#_via_git_sync).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"branch": schema.StringAttribute{
											Description:         "The branch to clone. Defaults to 'main'. Since git-sync v4.x.x this field is mapped to the flag '--ref'.",
											MarkdownDescription: "The branch to clone. Defaults to 'main'. Since git-sync v4.x.x this field is mapped to the flag '--ref'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"credentials_secret": schema.StringAttribute{
											Description:         "The name of the Secret used to access the repository if it is not public. This should include two fields: 'user' and 'password'. The 'password' field can either be an actual password (not recommended) or a GitHub token, as described [here](https://github.com/kubernetes/git-sync/tree/v4.2.4?tab=readme-ov-file#manual).",
											MarkdownDescription: "The name of the Secret used to access the repository if it is not public. This should include two fields: 'user' and 'password'. The 'password' field can either be an actual password (not recommended) or a GitHub token, as described [here](https://github.com/kubernetes/git-sync/tree/v4.2.4?tab=readme-ov-file#manual).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"depth": schema.Int64Attribute{
											Description:         "The depth of syncing i.e. the number of commits to clone; defaults to 1.",
											MarkdownDescription: "The depth of syncing i.e. the number of commits to clone; defaults to 1.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"git_folder": schema.StringAttribute{
											Description:         "The location of the DAG folder, relative to the synced repository root.",
											MarkdownDescription: "The location of the DAG folder, relative to the synced repository root.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"git_sync_conf": schema.MapAttribute{
											Description:         "A map of optional configuration settings that are listed in the [git-sync documentation](https://github.com/kubernetes/git-sync/tree/v4.2.4?tab=readme-ov-file#manual). Read the [git sync example](https://docs.stackable.tech/home/nightly/airflow/usage-guide/mounting-dags#_example).",
											MarkdownDescription: "A map of optional configuration settings that are listed in the [git-sync documentation](https://github.com/kubernetes/git-sync/tree/v4.2.4?tab=readme-ov-file#manual). Read the [git sync example](https://docs.stackable.tech/home/nightly/airflow/usage-guide/mounting-dags#_example).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"repo": schema.StringAttribute{
											Description:         "The git repository URL that will be cloned, for example: 'https://github.com/stackabletech/airflow-operator'.",
											MarkdownDescription: "The git repository URL that will be cloned, for example: 'https://github.com/stackabletech/airflow-operator'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"wait": schema.Int64Attribute{
											Description:         "The synchronization interval in seconds; defaults to 20 seconds. Since git-sync v4.x.x this field is mapped to the flag '--period'.",
											MarkdownDescription: "The synchronization interval in seconds; defaults to 20 seconds. Since git-sync v4.x.x this field is mapped to the flag '--period'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"expose_config": schema.BoolAttribute{
								Description:         "for internal use only - not for production use.",
								MarkdownDescription: "for internal use only - not for production use.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"listener_class": schema.StringAttribute{
								Description:         "This field controls which type of Service the Operator creates for this AirflowCluster: * cluster-internal: Use a ClusterIP service * external-unstable: Use a NodePort service * external-stable: Use a LoadBalancer service This is a temporary solution with the goal to keep yaml manifests forward compatible. In the future, this setting will control which [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass.html) will be used to expose the service, and ListenerClass names will stay the same, allowing for a non-breaking change.",
								MarkdownDescription: "This field controls which type of Service the Operator creates for this AirflowCluster: * cluster-internal: Use a ClusterIP service * external-unstable: Use a NodePort service * external-stable: Use a LoadBalancer service This is a temporary solution with the goal to keep yaml manifests forward compatible. In the future, this setting will control which [ListenerClass](https://docs.stackable.tech/home/nightly/listener-operator/listenerclass.html) will be used to expose the service, and ListenerClass names will stay the same, allowing for a non-breaking change.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("cluster-internal", "external-unstable", "external-stable"),
								},
							},

							"load_examples": schema.BoolAttribute{
								Description:         "Whether to load example DAGs or not; defaults to false. The examples are used in the [getting started guide](https://docs.stackable.tech/home/nightly/airflow/getting_started/).",
								MarkdownDescription: "Whether to load example DAGs or not; defaults to false. The examples are used in the [getting started guide](https://docs.stackable.tech/home/nightly/airflow/getting_started/).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vector_aggregator_config_map_name": schema.StringAttribute{
								Description:         "Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.",
								MarkdownDescription: "Name of the Vector aggregator [discovery ConfigMap](https://docs.stackable.tech/home/nightly/concepts/service_discovery). It must contain the key 'ADDRESS' with the address of the Vector aggregator. Follow the [logging tutorial](https://docs.stackable.tech/home/nightly/tutorials/logging-vector-aggregator) to learn how to configure log aggregation with Vector.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_mounts": schema.ListAttribute{
								Description:         "Additional volumes to mount. Use together with 'volumes' to define volumes.",
								MarkdownDescription: "Additional volumes to mount. Use together with 'volumes' to define volumes.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volumes": schema.ListAttribute{
								Description:         "Additional volumes to define. Use together with 'volumeMounts' to mount the volumes.",
								MarkdownDescription: "Additional volumes to define. Use together with 'volumeMounts' to mount the volumes.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
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

					"kubernetes_executors": schema.SingleNestedAttribute{
						Description:         "With the Kuberentes executor, executor Pods are created on demand.",
						MarkdownDescription: "With the Kuberentes executor, executor Pods are created on demand.",
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

											"storage": schema.MapAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"schedulers": schema.SingleNestedAttribute{
						Description:         "The 'scheduler' is responsible for triggering jobs and persisting their metadata to the backend database. Jobs are scheduled on the workers/executors.",
						MarkdownDescription: "The 'scheduler' is responsible for triggering jobs and persisting their metadata to the backend database. Jobs are scheduled on the workers/executors.",
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

											"storage": schema.MapAttribute{
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

													"storage": schema.MapAttribute{
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

					"webservers": schema.SingleNestedAttribute{
						Description:         "The 'webserver' role provides the main UI for user interaction.",
						MarkdownDescription: "The 'webserver' role provides the main UI for user interaction.",
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

											"storage": schema.MapAttribute{
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

													"storage": schema.MapAttribute{
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

func (r *AirflowStackableTechAirflowClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_airflow_stackable_tech_airflow_cluster_v1alpha1_manifest")

	var model AirflowStackableTechAirflowClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("airflow.stackable.tech/v1alpha1")
	model.Kind = pointer.String("AirflowCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
