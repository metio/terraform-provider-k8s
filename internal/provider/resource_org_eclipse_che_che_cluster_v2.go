/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type OrgEclipseCheCheClusterV2Resource struct{}

var (
	_ resource.Resource = (*OrgEclipseCheCheClusterV2Resource)(nil)
)

type OrgEclipseCheCheClusterV2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type OrgEclipseCheCheClusterV2GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Components *struct {
			CheServer *struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" yaml:"clusterRoles,omitempty"`

				Debug *bool `tfsdk:"debug" yaml:"debug,omitempty"`

				Deployment *struct {
					Containers *[]struct {
						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

								SecretKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Resources *struct {
							Limits *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"limits" yaml:"limits,omitempty"`

							Request *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"request" yaml:"request,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`
					} `tfsdk:"containers" yaml:"containers,omitempty"`

					SecurityContext *struct {
						FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

						RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				ExtraProperties *map[string]string `tfsdk:"extra_properties" yaml:"extraProperties,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				Proxy *struct {
					CredentialsSecretName *string `tfsdk:"credentials_secret_name" yaml:"credentialsSecretName,omitempty"`

					NonProxyHosts *[]string `tfsdk:"non_proxy_hosts" yaml:"nonProxyHosts,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"proxy" yaml:"proxy,omitempty"`
			} `tfsdk:"che_server" yaml:"cheServer,omitempty"`

			Dashboard *struct {
				Deployment *struct {
					Containers *[]struct {
						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

								SecretKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Resources *struct {
							Limits *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"limits" yaml:"limits,omitempty"`

							Request *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"request" yaml:"request,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`
					} `tfsdk:"containers" yaml:"containers,omitempty"`

					SecurityContext *struct {
						FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

						RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				HeaderMessage *struct {
					Show *bool `tfsdk:"show" yaml:"show,omitempty"`

					Text *string `tfsdk:"text" yaml:"text,omitempty"`
				} `tfsdk:"header_message" yaml:"headerMessage,omitempty"`
			} `tfsdk:"dashboard" yaml:"dashboard,omitempty"`

			Database *struct {
				CredentialsSecretName *string `tfsdk:"credentials_secret_name" yaml:"credentialsSecretName,omitempty"`

				Deployment *struct {
					Containers *[]struct {
						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

								SecretKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Resources *struct {
							Limits *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"limits" yaml:"limits,omitempty"`

							Request *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"request" yaml:"request,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`
					} `tfsdk:"containers" yaml:"containers,omitempty"`

					SecurityContext *struct {
						FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

						RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				ExternalDb *bool `tfsdk:"external_db" yaml:"externalDb,omitempty"`

				PostgresDb *string `tfsdk:"postgres_db" yaml:"postgresDb,omitempty"`

				PostgresHostName *string `tfsdk:"postgres_host_name" yaml:"postgresHostName,omitempty"`

				PostgresPort *string `tfsdk:"postgres_port" yaml:"postgresPort,omitempty"`

				Pvc *struct {
					ClaimSize *string `tfsdk:"claim_size" yaml:"claimSize,omitempty"`

					StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`
				} `tfsdk:"pvc" yaml:"pvc,omitempty"`
			} `tfsdk:"database" yaml:"database,omitempty"`

			DevWorkspace *struct {
				RunningLimit *string `tfsdk:"running_limit" yaml:"runningLimit,omitempty"`
			} `tfsdk:"dev_workspace" yaml:"devWorkspace,omitempty"`

			DevfileRegistry *struct {
				Deployment *struct {
					Containers *[]struct {
						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

								SecretKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Resources *struct {
							Limits *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"limits" yaml:"limits,omitempty"`

							Request *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"request" yaml:"request,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`
					} `tfsdk:"containers" yaml:"containers,omitempty"`

					SecurityContext *struct {
						FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

						RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				DisableInternalRegistry *bool `tfsdk:"disable_internal_registry" yaml:"disableInternalRegistry,omitempty"`

				ExternalDevfileRegistries *[]struct {
					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"external_devfile_registries" yaml:"externalDevfileRegistries,omitempty"`
			} `tfsdk:"devfile_registry" yaml:"devfileRegistry,omitempty"`

			ImagePuller *struct {
				Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

				Spec *struct {
					Affinity *string `tfsdk:"affinity" yaml:"affinity,omitempty"`

					CachingCPULimit *string `tfsdk:"caching_cpu_limit" yaml:"cachingCPULimit,omitempty"`

					CachingCPURequest *string `tfsdk:"caching_cpu_request" yaml:"cachingCPURequest,omitempty"`

					CachingIntervalHours *string `tfsdk:"caching_interval_hours" yaml:"cachingIntervalHours,omitempty"`

					CachingMemoryLimit *string `tfsdk:"caching_memory_limit" yaml:"cachingMemoryLimit,omitempty"`

					CachingMemoryRequest *string `tfsdk:"caching_memory_request" yaml:"cachingMemoryRequest,omitempty"`

					ConfigMapName *string `tfsdk:"config_map_name" yaml:"configMapName,omitempty"`

					DaemonsetName *string `tfsdk:"daemonset_name" yaml:"daemonsetName,omitempty"`

					DeploymentName *string `tfsdk:"deployment_name" yaml:"deploymentName,omitempty"`

					ImagePullSecrets *string `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					ImagePullerImage *string `tfsdk:"image_puller_image" yaml:"imagePullerImage,omitempty"`

					Images *string `tfsdk:"images" yaml:"images,omitempty"`

					NodeSelector *string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`
				} `tfsdk:"spec" yaml:"spec,omitempty"`
			} `tfsdk:"image_puller" yaml:"imagePuller,omitempty"`

			Metrics *struct {
				Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`
			} `tfsdk:"metrics" yaml:"metrics,omitempty"`

			PluginRegistry *struct {
				Deployment *struct {
					Containers *[]struct {
						Env *[]struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`

							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

									FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

									Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

								SecretKeyRef *struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
						} `tfsdk:"env" yaml:"env,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`

						ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Resources *struct {
							Limits *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"limits" yaml:"limits,omitempty"`

							Request *struct {
								Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

								Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
							} `tfsdk:"request" yaml:"request,omitempty"`
						} `tfsdk:"resources" yaml:"resources,omitempty"`
					} `tfsdk:"containers" yaml:"containers,omitempty"`

					SecurityContext *struct {
						FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

						RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				DisableInternalRegistry *bool `tfsdk:"disable_internal_registry" yaml:"disableInternalRegistry,omitempty"`

				ExternalPluginRegistries *[]struct {
					Url *string `tfsdk:"url" yaml:"url,omitempty"`
				} `tfsdk:"external_plugin_registries" yaml:"externalPluginRegistries,omitempty"`

				OpenVSXURL *string `tfsdk:"open_vsx_url" yaml:"openVSXURL,omitempty"`
			} `tfsdk:"plugin_registry" yaml:"pluginRegistry,omitempty"`
		} `tfsdk:"components" yaml:"components,omitempty"`

		ContainerRegistry *struct {
			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Organization *string `tfsdk:"organization" yaml:"organization,omitempty"`
		} `tfsdk:"container_registry" yaml:"containerRegistry,omitempty"`

		DevEnvironments *struct {
			ContainerBuildConfiguration *struct {
				OpenShiftSecurityContextConstraint *string `tfsdk:"open_shift_security_context_constraint" yaml:"openShiftSecurityContextConstraint,omitempty"`
			} `tfsdk:"container_build_configuration" yaml:"containerBuildConfiguration,omitempty"`

			DefaultComponents *[]struct {
				Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

				ComponentType *string `tfsdk:"component_type" yaml:"componentType,omitempty"`

				Container *struct {
					Annotation *struct {
						Deployment *map[string]string `tfsdk:"deployment" yaml:"deployment,omitempty"`

						Service *map[string]string `tfsdk:"service" yaml:"service,omitempty"`
					} `tfsdk:"annotation" yaml:"annotation,omitempty"`

					Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

					CpuLimit *string `tfsdk:"cpu_limit" yaml:"cpuLimit,omitempty"`

					CpuRequest *string `tfsdk:"cpu_request" yaml:"cpuRequest,omitempty"`

					DedicatedPod *bool `tfsdk:"dedicated_pod" yaml:"dedicatedPod,omitempty"`

					Endpoints *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

						TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
					} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

					Image *string `tfsdk:"image" yaml:"image,omitempty"`

					MemoryLimit *string `tfsdk:"memory_limit" yaml:"memoryLimit,omitempty"`

					MemoryRequest *string `tfsdk:"memory_request" yaml:"memoryRequest,omitempty"`

					MountSources *bool `tfsdk:"mount_sources" yaml:"mountSources,omitempty"`

					SourceMapping *string `tfsdk:"source_mapping" yaml:"sourceMapping,omitempty"`

					VolumeMounts *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`
				} `tfsdk:"container" yaml:"container,omitempty"`

				Custom *struct {
					ComponentClass *string `tfsdk:"component_class" yaml:"componentClass,omitempty"`

					EmbeddedResource utilities.Dynamic `tfsdk:"embedded_resource" yaml:"embeddedResource,omitempty"`
				} `tfsdk:"custom" yaml:"custom,omitempty"`

				Image *struct {
					AutoBuild *bool `tfsdk:"auto_build" yaml:"autoBuild,omitempty"`

					Dockerfile *struct {
						Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

						BuildContext *string `tfsdk:"build_context" yaml:"buildContext,omitempty"`

						DevfileRegistry *struct {
							Id *string `tfsdk:"id" yaml:"id,omitempty"`

							RegistryUrl *string `tfsdk:"registry_url" yaml:"registryUrl,omitempty"`
						} `tfsdk:"devfile_registry" yaml:"devfileRegistry,omitempty"`

						Git *struct {
							CheckoutFrom *struct {
								Remote *string `tfsdk:"remote" yaml:"remote,omitempty"`

								Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`
							} `tfsdk:"checkout_from" yaml:"checkoutFrom,omitempty"`

							FileLocation *string `tfsdk:"file_location" yaml:"fileLocation,omitempty"`

							Remotes *map[string]string `tfsdk:"remotes" yaml:"remotes,omitempty"`
						} `tfsdk:"git" yaml:"git,omitempty"`

						RootRequired *bool `tfsdk:"root_required" yaml:"rootRequired,omitempty"`

						SrcType *string `tfsdk:"src_type" yaml:"srcType,omitempty"`

						Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
					} `tfsdk:"dockerfile" yaml:"dockerfile,omitempty"`

					ImageName *string `tfsdk:"image_name" yaml:"imageName,omitempty"`

					ImageType *string `tfsdk:"image_type" yaml:"imageType,omitempty"`
				} `tfsdk:"image" yaml:"image,omitempty"`

				Kubernetes *struct {
					DeployByDefault *bool `tfsdk:"deploy_by_default" yaml:"deployByDefault,omitempty"`

					Endpoints *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

						TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
					} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

					Inlined *string `tfsdk:"inlined" yaml:"inlined,omitempty"`

					LocationType *string `tfsdk:"location_type" yaml:"locationType,omitempty"`

					Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
				} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Openshift *struct {
					DeployByDefault *bool `tfsdk:"deploy_by_default" yaml:"deployByDefault,omitempty"`

					Endpoints *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

						TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
					} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

					Inlined *string `tfsdk:"inlined" yaml:"inlined,omitempty"`

					LocationType *string `tfsdk:"location_type" yaml:"locationType,omitempty"`

					Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
				} `tfsdk:"openshift" yaml:"openshift,omitempty"`

				Plugin *struct {
					Commands *[]struct {
						Apply *struct {
							Component *string `tfsdk:"component" yaml:"component,omitempty"`

							Group *struct {
								IsDefault *bool `tfsdk:"is_default" yaml:"isDefault,omitempty"`

								Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
							} `tfsdk:"group" yaml:"group,omitempty"`

							Label *string `tfsdk:"label" yaml:"label,omitempty"`
						} `tfsdk:"apply" yaml:"apply,omitempty"`

						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						CommandType *string `tfsdk:"command_type" yaml:"commandType,omitempty"`

						Composite *struct {
							Commands *[]string `tfsdk:"commands" yaml:"commands,omitempty"`

							Group *struct {
								IsDefault *bool `tfsdk:"is_default" yaml:"isDefault,omitempty"`

								Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
							} `tfsdk:"group" yaml:"group,omitempty"`

							Label *string `tfsdk:"label" yaml:"label,omitempty"`

							Parallel *bool `tfsdk:"parallel" yaml:"parallel,omitempty"`
						} `tfsdk:"composite" yaml:"composite,omitempty"`

						Exec *struct {
							CommandLine *string `tfsdk:"command_line" yaml:"commandLine,omitempty"`

							Component *string `tfsdk:"component" yaml:"component,omitempty"`

							Env *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"env" yaml:"env,omitempty"`

							Group *struct {
								IsDefault *bool `tfsdk:"is_default" yaml:"isDefault,omitempty"`

								Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
							} `tfsdk:"group" yaml:"group,omitempty"`

							HotReloadCapable *bool `tfsdk:"hot_reload_capable" yaml:"hotReloadCapable,omitempty"`

							Label *string `tfsdk:"label" yaml:"label,omitempty"`

							WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						Id *string `tfsdk:"id" yaml:"id,omitempty"`
					} `tfsdk:"commands" yaml:"commands,omitempty"`

					Components *[]struct {
						Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

						ComponentType *string `tfsdk:"component_type" yaml:"componentType,omitempty"`

						Container *struct {
							Annotation *struct {
								Deployment *map[string]string `tfsdk:"deployment" yaml:"deployment,omitempty"`

								Service *map[string]string `tfsdk:"service" yaml:"service,omitempty"`
							} `tfsdk:"annotation" yaml:"annotation,omitempty"`

							Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

							CpuLimit *string `tfsdk:"cpu_limit" yaml:"cpuLimit,omitempty"`

							CpuRequest *string `tfsdk:"cpu_request" yaml:"cpuRequest,omitempty"`

							DedicatedPod *bool `tfsdk:"dedicated_pod" yaml:"dedicatedPod,omitempty"`

							Endpoints *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

								Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

								Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

								Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

								TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
							} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

							Env *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"env" yaml:"env,omitempty"`

							Image *string `tfsdk:"image" yaml:"image,omitempty"`

							MemoryLimit *string `tfsdk:"memory_limit" yaml:"memoryLimit,omitempty"`

							MemoryRequest *string `tfsdk:"memory_request" yaml:"memoryRequest,omitempty"`

							MountSources *bool `tfsdk:"mount_sources" yaml:"mountSources,omitempty"`

							SourceMapping *string `tfsdk:"source_mapping" yaml:"sourceMapping,omitempty"`

							VolumeMounts *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`
						} `tfsdk:"container" yaml:"container,omitempty"`

						Image *struct {
							AutoBuild *bool `tfsdk:"auto_build" yaml:"autoBuild,omitempty"`

							Dockerfile *struct {
								Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

								BuildContext *string `tfsdk:"build_context" yaml:"buildContext,omitempty"`

								DevfileRegistry *struct {
									Id *string `tfsdk:"id" yaml:"id,omitempty"`

									RegistryUrl *string `tfsdk:"registry_url" yaml:"registryUrl,omitempty"`
								} `tfsdk:"devfile_registry" yaml:"devfileRegistry,omitempty"`

								Git *struct {
									CheckoutFrom *struct {
										Remote *string `tfsdk:"remote" yaml:"remote,omitempty"`

										Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`
									} `tfsdk:"checkout_from" yaml:"checkoutFrom,omitempty"`

									FileLocation *string `tfsdk:"file_location" yaml:"fileLocation,omitempty"`

									Remotes *map[string]string `tfsdk:"remotes" yaml:"remotes,omitempty"`
								} `tfsdk:"git" yaml:"git,omitempty"`

								RootRequired *bool `tfsdk:"root_required" yaml:"rootRequired,omitempty"`

								SrcType *string `tfsdk:"src_type" yaml:"srcType,omitempty"`

								Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
							} `tfsdk:"dockerfile" yaml:"dockerfile,omitempty"`

							ImageName *string `tfsdk:"image_name" yaml:"imageName,omitempty"`

							ImageType *string `tfsdk:"image_type" yaml:"imageType,omitempty"`
						} `tfsdk:"image" yaml:"image,omitempty"`

						Kubernetes *struct {
							DeployByDefault *bool `tfsdk:"deploy_by_default" yaml:"deployByDefault,omitempty"`

							Endpoints *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

								Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

								Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

								Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

								TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
							} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

							Inlined *string `tfsdk:"inlined" yaml:"inlined,omitempty"`

							LocationType *string `tfsdk:"location_type" yaml:"locationType,omitempty"`

							Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
						} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Openshift *struct {
							DeployByDefault *bool `tfsdk:"deploy_by_default" yaml:"deployByDefault,omitempty"`

							Endpoints *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" yaml:"annotation,omitempty"`

								Attributes utilities.Dynamic `tfsdk:"attributes" yaml:"attributes,omitempty"`

								Exposure *string `tfsdk:"exposure" yaml:"exposure,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

								Secure *bool `tfsdk:"secure" yaml:"secure,omitempty"`

								TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
							} `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

							Inlined *string `tfsdk:"inlined" yaml:"inlined,omitempty"`

							LocationType *string `tfsdk:"location_type" yaml:"locationType,omitempty"`

							Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
						} `tfsdk:"openshift" yaml:"openshift,omitempty"`

						Volume *struct {
							Ephemeral *bool `tfsdk:"ephemeral" yaml:"ephemeral,omitempty"`

							Size *string `tfsdk:"size" yaml:"size,omitempty"`
						} `tfsdk:"volume" yaml:"volume,omitempty"`
					} `tfsdk:"components" yaml:"components,omitempty"`

					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					ImportReferenceType *string `tfsdk:"import_reference_type" yaml:"importReferenceType,omitempty"`

					Kubernetes *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

					RegistryUrl *string `tfsdk:"registry_url" yaml:"registryUrl,omitempty"`

					Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`

					Version *string `tfsdk:"version" yaml:"version,omitempty"`
				} `tfsdk:"plugin" yaml:"plugin,omitempty"`

				Volume *struct {
					Ephemeral *bool `tfsdk:"ephemeral" yaml:"ephemeral,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`
				} `tfsdk:"volume" yaml:"volume,omitempty"`
			} `tfsdk:"default_components" yaml:"defaultComponents,omitempty"`

			DefaultEditor *string `tfsdk:"default_editor" yaml:"defaultEditor,omitempty"`

			DefaultNamespace *struct {
				AutoProvision *bool `tfsdk:"auto_provision" yaml:"autoProvision,omitempty"`

				Template *string `tfsdk:"template" yaml:"template,omitempty"`
			} `tfsdk:"default_namespace" yaml:"defaultNamespace,omitempty"`

			DefaultPlugins *[]struct {
				Editor *string `tfsdk:"editor" yaml:"editor,omitempty"`

				Plugins *[]string `tfsdk:"plugins" yaml:"plugins,omitempty"`
			} `tfsdk:"default_plugins" yaml:"defaultPlugins,omitempty"`

			DisableContainerBuildCapabilities *bool `tfsdk:"disable_container_build_capabilities" yaml:"disableContainerBuildCapabilities,omitempty"`

			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

			SecondsOfInactivityBeforeIdling *int64 `tfsdk:"seconds_of_inactivity_before_idling" yaml:"secondsOfInactivityBeforeIdling,omitempty"`

			SecondsOfRunBeforeIdling *int64 `tfsdk:"seconds_of_run_before_idling" yaml:"secondsOfRunBeforeIdling,omitempty"`

			Storage *struct {
				PerUserStrategyPvcConfig *struct {
					ClaimSize *string `tfsdk:"claim_size" yaml:"claimSize,omitempty"`

					StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`
				} `tfsdk:"per_user_strategy_pvc_config" yaml:"perUserStrategyPvcConfig,omitempty"`

				PerWorkspaceStrategyPvcConfig *struct {
					ClaimSize *string `tfsdk:"claim_size" yaml:"claimSize,omitempty"`

					StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`
				} `tfsdk:"per_workspace_strategy_pvc_config" yaml:"perWorkspaceStrategyPvcConfig,omitempty"`

				PvcStrategy *string `tfsdk:"pvc_strategy" yaml:"pvcStrategy,omitempty"`
			} `tfsdk:"storage" yaml:"storage,omitempty"`

			Tolerations *[]struct {
				Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

			TrustedCerts *struct {
				GitTrustedCertsConfigMapName *string `tfsdk:"git_trusted_certs_config_map_name" yaml:"gitTrustedCertsConfigMapName,omitempty"`
			} `tfsdk:"trusted_certs" yaml:"trustedCerts,omitempty"`
		} `tfsdk:"dev_environments" yaml:"devEnvironments,omitempty"`

		GitServices *struct {
			Bitbucket *[]struct {
				Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"bitbucket" yaml:"bitbucket,omitempty"`

			Github *[]struct {
				DisableSubdomainIsolation *bool `tfsdk:"disable_subdomain_isolation" yaml:"disableSubdomainIsolation,omitempty"`

				Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"github" yaml:"github,omitempty"`

			Gitlab *[]struct {
				Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"gitlab" yaml:"gitlab,omitempty"`
		} `tfsdk:"git_services" yaml:"gitServices,omitempty"`

		Networking *struct {
			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Auth *struct {
				Gateway *struct {
					ConfigLabels *map[string]string `tfsdk:"config_labels" yaml:"configLabels,omitempty"`

					Deployment *struct {
						Containers *[]struct {
							Env *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									ConfigMapKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
									} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

									FieldRef *struct {
										ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

										FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
									} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

									ResourceFieldRef *struct {
										ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

										Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

										Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
									} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"env" yaml:"env,omitempty"`

							Image *string `tfsdk:"image" yaml:"image,omitempty"`

							ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Resources *struct {
								Limits *struct {
									Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

									Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
								} `tfsdk:"limits" yaml:"limits,omitempty"`

								Request *struct {
									Cpu utilities.IntOrString `tfsdk:"cpu" yaml:"cpu,omitempty"`

									Memory utilities.IntOrString `tfsdk:"memory" yaml:"memory,omitempty"`
								} `tfsdk:"request" yaml:"request,omitempty"`
							} `tfsdk:"resources" yaml:"resources,omitempty"`
						} `tfsdk:"containers" yaml:"containers,omitempty"`

						SecurityContext *struct {
							FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

							RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`
						} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
					} `tfsdk:"deployment" yaml:"deployment,omitempty"`
				} `tfsdk:"gateway" yaml:"gateway,omitempty"`

				IdentityProviderURL *string `tfsdk:"identity_provider_url" yaml:"identityProviderURL,omitempty"`

				IdentityToken *string `tfsdk:"identity_token" yaml:"identityToken,omitempty"`

				OAuthAccessTokenInactivityTimeoutSeconds *int64 `tfsdk:"o_auth_access_token_inactivity_timeout_seconds" yaml:"oAuthAccessTokenInactivityTimeoutSeconds,omitempty"`

				OAuthAccessTokenMaxAgeSeconds *int64 `tfsdk:"o_auth_access_token_max_age_seconds" yaml:"oAuthAccessTokenMaxAgeSeconds,omitempty"`

				OAuthClientName *string `tfsdk:"o_auth_client_name" yaml:"oAuthClientName,omitempty"`

				OAuthScope *string `tfsdk:"o_auth_scope" yaml:"oAuthScope,omitempty"`

				OAuthSecret *string `tfsdk:"o_auth_secret" yaml:"oAuthSecret,omitempty"`
			} `tfsdk:"auth" yaml:"auth,omitempty"`

			Domain *string `tfsdk:"domain" yaml:"domain,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

			TlsSecretName *string `tfsdk:"tls_secret_name" yaml:"tlsSecretName,omitempty"`
		} `tfsdk:"networking" yaml:"networking,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewOrgEclipseCheCheClusterV2Resource() resource.Resource {
	return &OrgEclipseCheCheClusterV2Resource{}
}

func (r *OrgEclipseCheCheClusterV2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_org_eclipse_che_che_cluster_v2"
}

func (r *OrgEclipseCheCheClusterV2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "The 'CheCluster' custom resource allows defining and managing Eclipse Che server installation. Based on these settings, the  Operator automatically creates and maintains several ConfigMaps: 'che', 'plugin-registry', 'devfile-registry' that will contain the appropriate environment variables of the various components of the installation. These generated ConfigMaps must NOT be updated manually.",
		MarkdownDescription: "The 'CheCluster' custom resource allows defining and managing Eclipse Che server installation. Based on these settings, the  Operator automatically creates and maintains several ConfigMaps: 'che', 'plugin-registry', 'devfile-registry' that will contain the appropriate environment variables of the various components of the installation. These generated ConfigMaps must NOT be updated manually.",
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

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
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
				Description:         "Desired configuration of Eclipse Che installation.",
				MarkdownDescription: "Desired configuration of Eclipse Che installation.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"components": {
						Description:         "Che components configuration.",
						MarkdownDescription: "Che components configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"che_server": {
								Description:         "General configuration settings related to the Che server.",
								MarkdownDescription: "General configuration settings related to the Che server.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cluster_roles": {
										Description:         "ClusterRoles assigned to Che ServiceAccount. The defaults roles are: - '<che-namespace>-cheworkspaces-namespaces-clusterrole' - '<che-namespace>-cheworkspaces-clusterrole' - '<che-namespace>-cheworkspaces-devworkspace-clusterrole' where the <che-namespace> is the namespace where the CheCluster CRD is created. Each role must have a 'app.kubernetes.io/part-of=che.eclipse.org' label. The Che Operator must already have all permissions in these ClusterRoles to grant them.",
										MarkdownDescription: "ClusterRoles assigned to Che ServiceAccount. The defaults roles are: - '<che-namespace>-cheworkspaces-namespaces-clusterrole' - '<che-namespace>-cheworkspaces-clusterrole' - '<che-namespace>-cheworkspaces-devworkspace-clusterrole' where the <che-namespace> is the namespace where the CheCluster CRD is created. Each role must have a 'app.kubernetes.io/part-of=che.eclipse.org' label. The Che Operator must already have all permissions in these ClusterRoles to grant them.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"debug": {
										Description:         "Enables the debug mode for Che server.",
										MarkdownDescription: "Enables the debug mode for Che server.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": {
										Description:         "Deployment override options.",
										MarkdownDescription: "Deployment override options.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"containers": {
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"env": {
														Description:         "List of environment variables to set in the container.",
														MarkdownDescription: "List of environment variables to set in the container.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_from": {
																Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map_key_ref": {
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": {
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": {
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": {
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",

																				Type: types.BoolType,

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

													"image": {
														Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
														MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image_pull_policy": {
														Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
														MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
														},
													},

													"name": {
														Description:         "Container name.",
														MarkdownDescription: "Container name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "Compute resources required by this container.",
														MarkdownDescription: "Compute resources required by this container.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "Describes the maximum amount of compute resources allowed.",
																MarkdownDescription: "Describes the maximum amount of compute resources allowed.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request": {
																Description:         "Describes the minimum amount of compute resources required.",
																MarkdownDescription: "Describes the minimum amount of compute resources required.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

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

											"security_context": {
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_group": {
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",

														Type: types.Int64Type,

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

									"extra_properties": {
										Description:         "A map of additional environment variables applied in the generated 'che' ConfigMap to be used by the Che server in addition to the values already generated from other fields of the 'CheCluster' custom resource (CR). If the 'extraProperties' field contains a property normally generated in 'che' ConfigMap from other CR fields, the value defined in the 'extraProperties' is used instead.",
										MarkdownDescription: "A map of additional environment variables applied in the generated 'che' ConfigMap to be used by the Che server in addition to the values already generated from other fields of the 'CheCluster' custom resource (CR). If the 'extraProperties' field contains a property normally generated in 'che' ConfigMap from other CR fields, the value defined in the 'extraProperties' is used instead.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_level": {
										Description:         "The log level for the Che server: 'INFO' or 'DEBUG'.",
										MarkdownDescription: "The log level for the Che server: 'INFO' or 'DEBUG'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"proxy": {
										Description:         "Proxy server settings for Kubernetes cluster. No additional configuration is required for OpenShift cluster. By specifying these settings for the OpenShift cluster, you override the OpenShift proxy configuration.",
										MarkdownDescription: "Proxy server settings for Kubernetes cluster. No additional configuration is required for OpenShift cluster. By specifying these settings for the OpenShift cluster, you override the OpenShift proxy configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"credentials_secret_name": {
												Description:         "The secret name that contains 'user' and 'password' for a proxy server. The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
												MarkdownDescription: "The secret name that contains 'user' and 'password' for a proxy server. The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"non_proxy_hosts": {
												Description:         "A list of hosts that can be reached directly, bypassing the proxy. Specify wild card domain use the following form '.<DOMAIN>', for example:    - localhost    - my.host.com    - 123.42.12.32 Use only when a proxy configuration is required. The Operator respects OpenShift cluster-wide proxy configuration, defining 'nonProxyHosts' in a custom resource leads to merging non-proxy hosts lists from the cluster proxy configuration, and the ones defined in the custom resources. See the following page: https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html.",
												MarkdownDescription: "A list of hosts that can be reached directly, bypassing the proxy. Specify wild card domain use the following form '.<DOMAIN>', for example:    - localhost    - my.host.com    - 123.42.12.32 Use only when a proxy configuration is required. The Operator respects OpenShift cluster-wide proxy configuration, defining 'nonProxyHosts' in a custom resource leads to merging non-proxy hosts lists from the cluster proxy configuration, and the ones defined in the custom resources. See the following page: https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Proxy server port.",
												MarkdownDescription: "Proxy server port.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": {
												Description:         "URL (protocol+hostname) of the proxy server. Use only when a proxy configuration is required. The Operator respects OpenShift cluster-wide proxy configuration, defining 'url' in a custom resource leads to overriding the cluster proxy configuration. See the following page: https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html.",
												MarkdownDescription: "URL (protocol+hostname) of the proxy server. Use only when a proxy configuration is required. The Operator respects OpenShift cluster-wide proxy configuration, defining 'url' in a custom resource leads to overriding the cluster proxy configuration. See the following page: https://docs.openshift.com/container-platform/4.4/networking/enable-cluster-wide-proxy.html.",

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

							"dashboard": {
								Description:         "Configuration settings related to the dashboard used by the Che installation.",
								MarkdownDescription: "Configuration settings related to the dashboard used by the Che installation.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"deployment": {
										Description:         "Deployment override options.",
										MarkdownDescription: "Deployment override options.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"containers": {
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"env": {
														Description:         "List of environment variables to set in the container.",
														MarkdownDescription: "List of environment variables to set in the container.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_from": {
																Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map_key_ref": {
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": {
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": {
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": {
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",

																				Type: types.BoolType,

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

													"image": {
														Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
														MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image_pull_policy": {
														Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
														MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
														},
													},

													"name": {
														Description:         "Container name.",
														MarkdownDescription: "Container name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "Compute resources required by this container.",
														MarkdownDescription: "Compute resources required by this container.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "Describes the maximum amount of compute resources allowed.",
																MarkdownDescription: "Describes the maximum amount of compute resources allowed.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request": {
																Description:         "Describes the minimum amount of compute resources required.",
																MarkdownDescription: "Describes the minimum amount of compute resources required.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

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

											"security_context": {
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_group": {
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",

														Type: types.Int64Type,

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

									"header_message": {
										Description:         "Dashboard header message.",
										MarkdownDescription: "Dashboard header message.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"show": {
												Description:         "Instructs dashboard to show the message.",
												MarkdownDescription: "Instructs dashboard to show the message.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"text": {
												Description:         "Warning message displayed on the user dashboard.",
												MarkdownDescription: "Warning message displayed on the user dashboard.",

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

							"database": {
								Description:         "Configuration settings related to the database used by the Che installation.",
								MarkdownDescription: "Configuration settings related to the database used by the Che installation.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials_secret_name": {
										Description:         "The secret that contains PostgreSQL 'user' and 'password' that the Che server uses to connect to the database. The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
										MarkdownDescription: "The secret that contains PostgreSQL 'user' and 'password' that the Che server uses to connect to the database. The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": {
										Description:         "Deployment override options.",
										MarkdownDescription: "Deployment override options.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"containers": {
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"env": {
														Description:         "List of environment variables to set in the container.",
														MarkdownDescription: "List of environment variables to set in the container.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_from": {
																Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map_key_ref": {
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": {
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": {
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": {
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",

																				Type: types.BoolType,

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

													"image": {
														Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
														MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image_pull_policy": {
														Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
														MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
														},
													},

													"name": {
														Description:         "Container name.",
														MarkdownDescription: "Container name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "Compute resources required by this container.",
														MarkdownDescription: "Compute resources required by this container.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "Describes the maximum amount of compute resources allowed.",
																MarkdownDescription: "Describes the maximum amount of compute resources allowed.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request": {
																Description:         "Describes the minimum amount of compute resources required.",
																MarkdownDescription: "Describes the minimum amount of compute resources required.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

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

											"security_context": {
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_group": {
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",

														Type: types.Int64Type,

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

									"external_db": {
										Description:         "Instructs the Operator to deploy a dedicated database. By default, a dedicated PostgreSQL database is deployed as part of the Che installation. When 'externalDb' is set as 'true', no dedicated database is deployed by the Operator and you need to provide connection details about the external database you want to use.",
										MarkdownDescription: "Instructs the Operator to deploy a dedicated database. By default, a dedicated PostgreSQL database is deployed as part of the Che installation. When 'externalDb' is set as 'true', no dedicated database is deployed by the Operator and you need to provide connection details about the external database you want to use.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"postgres_db": {
										Description:         "PostgreSQL database name that the Che server uses to connect to the database.",
										MarkdownDescription: "PostgreSQL database name that the Che server uses to connect to the database.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"postgres_host_name": {
										Description:         "PostgreSQL database hostname that the Che server connects to. Override this value only when using an external database. See field 'externalDb'.",
										MarkdownDescription: "PostgreSQL database hostname that the Che server connects to. Override this value only when using an external database. See field 'externalDb'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"postgres_port": {
										Description:         "PostgreSQL Database port the Che server connects to. Override this value only when using an external database. See field 'externalDb'.",
										MarkdownDescription: "PostgreSQL Database port the Che server connects to. Override this value only when using an external database. See field 'externalDb'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pvc": {
										Description:         "PVC settings for PostgreSQL database.",
										MarkdownDescription: "PVC settings for PostgreSQL database.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"claim_size": {
												Description:         "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",
												MarkdownDescription: "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_class": {
												Description:         "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",
												MarkdownDescription: "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",

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

							"dev_workspace": {
								Description:         "DevWorkspace Operator configuration.",
								MarkdownDescription: "DevWorkspace Operator configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"running_limit": {
										Description:         "The maximum number of running workspaces per user.",
										MarkdownDescription: "The maximum number of running workspaces per user.",

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

							"devfile_registry": {
								Description:         "Configuration settings related to the devfile registry used by the Che installation.",
								MarkdownDescription: "Configuration settings related to the devfile registry used by the Che installation.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"deployment": {
										Description:         "Deployment override options.",
										MarkdownDescription: "Deployment override options.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"containers": {
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"env": {
														Description:         "List of environment variables to set in the container.",
														MarkdownDescription: "List of environment variables to set in the container.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_from": {
																Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map_key_ref": {
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": {
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": {
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": {
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",

																				Type: types.BoolType,

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

													"image": {
														Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
														MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image_pull_policy": {
														Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
														MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
														},
													},

													"name": {
														Description:         "Container name.",
														MarkdownDescription: "Container name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "Compute resources required by this container.",
														MarkdownDescription: "Compute resources required by this container.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "Describes the maximum amount of compute resources allowed.",
																MarkdownDescription: "Describes the maximum amount of compute resources allowed.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request": {
																Description:         "Describes the minimum amount of compute resources required.",
																MarkdownDescription: "Describes the minimum amount of compute resources required.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

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

											"security_context": {
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_group": {
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",

														Type: types.Int64Type,

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

									"disable_internal_registry": {
										Description:         "Disables internal devfile registry.",
										MarkdownDescription: "Disables internal devfile registry.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"external_devfile_registries": {
										Description:         "External devfile registries serving sample ready-to-use devfiles.",
										MarkdownDescription: "External devfile registries serving sample ready-to-use devfiles.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"url": {
												Description:         "The public UR of the devfile registry that serves sample ready-to-use devfiles.",
												MarkdownDescription: "The public UR of the devfile registry that serves sample ready-to-use devfiles.",

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

							"image_puller": {
								Description:         "Kubernetes Image Puller configuration.",
								MarkdownDescription: "Kubernetes Image Puller configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enable": {
										Description:         "Install and configure the community supported Kubernetes Image Puller Operator. When you set the value to 'true' without providing any specs, it creates a default Kubernetes Image Puller object managed by the Operator. When you set the value to 'false', the Kubernetes Image Puller object is deleted, and the Operator uninstalled, regardless of whether a spec is provided. If you leave the 'spec.images' field empty, a set of recommended workspace-related images is automatically detected and pre-pulled after installation. Note that while this Operator and its behavior is community-supported, its payload may be commercially-supported for pulling commercially-supported images.",
										MarkdownDescription: "Install and configure the community supported Kubernetes Image Puller Operator. When you set the value to 'true' without providing any specs, it creates a default Kubernetes Image Puller object managed by the Operator. When you set the value to 'false', the Kubernetes Image Puller object is deleted, and the Operator uninstalled, regardless of whether a spec is provided. If you leave the 'spec.images' field empty, a set of recommended workspace-related images is automatically detected and pre-pulled after installation. Note that while this Operator and its behavior is community-supported, its payload may be commercially-supported for pulling commercially-supported images.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"spec": {
										Description:         "A Kubernetes Image Puller spec to configure the image puller in the CheCluster.",
										MarkdownDescription: "A Kubernetes Image Puller spec to configure the image puller in the CheCluster.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"caching_cpu_limit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"caching_cpu_request": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"caching_interval_hours": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"caching_memory_limit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"caching_memory_request": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"config_map_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"daemonset_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"deployment_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_pull_secrets": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image_puller_image": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"images": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selector": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics": {
								Description:         "Che server metrics configuration.",
								MarkdownDescription: "Che server metrics configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enable": {
										Description:         "Enables 'metrics' for the Che server endpoint.",
										MarkdownDescription: "Enables 'metrics' for the Che server endpoint.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry": {
								Description:         "Configuration settings related to the plug-in registry used by the Che installation.",
								MarkdownDescription: "Configuration settings related to the plug-in registry used by the Che installation.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"deployment": {
										Description:         "Deployment override options.",
										MarkdownDescription: "Deployment override options.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"containers": {
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"env": {
														Description:         "List of environment variables to set in the container.",
														MarkdownDescription: "List of environment variables to set in the container.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"value": {
																Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value_from": {
																Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"config_map_key_ref": {
																		Description:         "Selects a key of a ConfigMap.",
																		MarkdownDescription: "Selects a key of a ConfigMap.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key to select.",
																				MarkdownDescription: "The key to select.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the ConfigMap or its key must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"field_ref": {
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"api_version": {
																				Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																				MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_path": {
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"resource_field_ref": {
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"container_name": {
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"divisor": {
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource": {
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret_key_ref": {
																		Description:         "Selects a key of a secret in the pod's namespace",
																		MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"optional": {
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",

																				Type: types.BoolType,

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

													"image": {
														Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
														MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image_pull_policy": {
														Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
														MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
														},
													},

													"name": {
														Description:         "Container name.",
														MarkdownDescription: "Container name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"resources": {
														Description:         "Compute resources required by this container.",
														MarkdownDescription: "Compute resources required by this container.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"limits": {
																Description:         "Describes the maximum amount of compute resources allowed.",
																MarkdownDescription: "Describes the maximum amount of compute resources allowed.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"request": {
																Description:         "Describes the minimum amount of compute resources required.",
																MarkdownDescription: "Describes the minimum amount of compute resources required.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cpu": {
																		Description:         "CPU, in cores. (500m = .5 cores)",
																		MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"memory": {
																		Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																		MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																		Type: utilities.IntOrStringType{},

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

											"security_context": {
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_group": {
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",

														Type: types.Int64Type,

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

									"disable_internal_registry": {
										Description:         "Disables internal plug-in registry.",
										MarkdownDescription: "Disables internal plug-in registry.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"external_plugin_registries": {
										Description:         "External plugin registries.",
										MarkdownDescription: "External plugin registries.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"url": {
												Description:         "Public URL of the plug-in registry.",
												MarkdownDescription: "Public URL of the plug-in registry.",

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

									"open_vsx_url": {
										Description:         "Open VSX registry URL. If omitted an embedded instance will be used.",
										MarkdownDescription: "Open VSX registry URL. If omitted an embedded instance will be used.",

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

					"container_registry": {
						Description:         "Configuration of an alternative registry that stores Che images.",
						MarkdownDescription: "Configuration of an alternative registry that stores Che images.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"hostname": {
								Description:         "An optional hostname or URL of an alternative container registry to pull images from. This value overrides the container registry hostname defined in all the default container images involved in a Che deployment. This is particularly useful for installing Che in a restricted environment.",
								MarkdownDescription: "An optional hostname or URL of an alternative container registry to pull images from. This value overrides the container registry hostname defined in all the default container images involved in a Che deployment. This is particularly useful for installing Che in a restricted environment.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"organization": {
								Description:         "An optional repository name of an alternative registry to pull images from. This value overrides the container registry organization defined in all the default container images involved in a Che deployment. This is particularly useful for installing Eclipse Che in a restricted environment.",
								MarkdownDescription: "An optional repository name of an alternative registry to pull images from. This value overrides the container registry organization defined in all the default container images involved in a Che deployment. This is particularly useful for installing Eclipse Che in a restricted environment.",

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

					"dev_environments": {
						Description:         "Development environment default configuration options.",
						MarkdownDescription: "Development environment default configuration options.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"container_build_configuration": {
								Description:         "Container build configuration.",
								MarkdownDescription: "Container build configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"open_shift_security_context_constraint": {
										Description:         "OpenShift security context constraint to build containers.",
										MarkdownDescription: "OpenShift security context constraint to build containers.",

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

							"default_components": {
								Description:         "Default components applied to DevWorkspaces. These default components are meant to be used when a Devfile, that does not contain any components.",
								MarkdownDescription: "Default components applied to DevWorkspaces. These default components are meant to be used when a Devfile, that does not contain any components.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"attributes": {
										Description:         "Map of implementation-dependant free-form YAML attributes.",
										MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"component_type": {
										Description:         "Type of component",
										MarkdownDescription: "Type of component",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image", "Plugin", "Custom"),
										},
									},

									"container": {
										Description:         "Allows adding and configuring devworkspace-related containers",
										MarkdownDescription: "Allows adding and configuring devworkspace-related containers",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotation": {
												Description:         "Annotations that should be added to specific resources for this container",
												MarkdownDescription: "Annotations that should be added to specific resources for this container",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"deployment": {
														Description:         "Annotations to be added to deployment",
														MarkdownDescription: "Annotations to be added to deployment",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service": {
														Description:         "Annotations to be added to service",
														MarkdownDescription: "Annotations to be added to service",

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

											"args": {
												Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",
												MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": {
												Description:         "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",
												MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cpu_limit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cpu_request": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dedicated_pod": {
												Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",
												MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"endpoints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"annotation": {
														Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
														MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attributes": {
														Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exposure": {
														Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("public", "internal", "none"),
														},
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": {
														Description:         "Path of the endpoint URL",
														MarkdownDescription: "Path of the endpoint URL",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": {
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_port": {
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"env": {
												Description:         "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",
												MarkdownDescription: "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"memory_limit": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"memory_request": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mount_sources": {
												Description:         "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
												MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_mapping": {
												Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
												MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_mounts": {
												Description:         "List of volumes mounts that should be mounted is this container.",
												MarkdownDescription: "List of volumes mounts that should be mounted is this container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
														MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": {
														Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
														MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",

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

									"custom": {
										Description:         "Custom component whose logic is implementation-dependant and should be provided by the user possibly through some dedicated controller",
										MarkdownDescription: "Custom component whose logic is implementation-dependant and should be provided by the user possibly through some dedicated controller",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"component_class": {
												Description:         "Class of component that the associated implementation controller should use to process this command with the appropriate logic",
												MarkdownDescription: "Class of component that the associated implementation controller should use to process this command with the appropriate logic",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"embedded_resource": {
												Description:         "Additional free-form configuration for this custom component that the implementation controller will know how to use",
												MarkdownDescription: "Additional free-form configuration for this custom component that the implementation controller will know how to use",

												Type: utilities.DynamicType{},

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "Allows specifying the definition of an image for outer loop builds",
										MarkdownDescription: "Allows specifying the definition of an image for outer loop builds",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"auto_build": {
												Description:         "Defines if the image should be built during startup.  Default value is 'false'",
												MarkdownDescription: "Defines if the image should be built during startup.  Default value is 'false'",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dockerfile": {
												Description:         "Allows specifying dockerfile type build",
												MarkdownDescription: "Allows specifying dockerfile type build",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"args": {
														Description:         "The arguments to supply to the dockerfile build.",
														MarkdownDescription: "The arguments to supply to the dockerfile build.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"build_context": {
														Description:         "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
														MarkdownDescription: "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"devfile_registry": {
														Description:         "Dockerfile's Devfile Registry source",
														MarkdownDescription: "Dockerfile's Devfile Registry source",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"id": {
																Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"registry_url": {
																Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",

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

													"git": {
														Description:         "Dockerfile's Git source",
														MarkdownDescription: "Dockerfile's Git source",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"checkout_from": {
																Description:         "Defines from what the project should be checked out. Required if there are more than one remote configured",
																MarkdownDescription: "Defines from what the project should be checked out. Required if there are more than one remote configured",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"remote": {
																		Description:         "The remote name should be used as init. Required if there are more than one remote configured",
																		MarkdownDescription: "The remote name should be used as init. Required if there are more than one remote configured",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"revision": {
																		Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																		MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",

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

															"file_location": {
																Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"remotes": {
																Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",

																Type: types.MapType{ElemType: types.StringType},

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"root_required": {
														Description:         "Specify if a privileged builder pod is required.  Default value is 'false'",
														MarkdownDescription: "Specify if a privileged builder pod is required.  Default value is 'false'",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"src_type": {
														Description:         "Type of Dockerfile src",
														MarkdownDescription: "Type of Dockerfile src",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Uri", "DevfileRegistry", "Git"),
														},
													},

													"uri": {
														Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
														MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",

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

											"image_name": {
												Description:         "Name of the image for the resulting outerloop build",
												MarkdownDescription: "Name of the image for the resulting outerloop build",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"image_type": {
												Description:         "Type of image",
												MarkdownDescription: "Type of image",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Dockerfile"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kubernetes": {
										Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
										MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"deploy_by_default": {
												Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
												MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"endpoints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"annotation": {
														Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
														MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attributes": {
														Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exposure": {
														Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("public", "internal", "none"),
														},
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": {
														Description:         "Path of the endpoint URL",
														MarkdownDescription: "Path of the endpoint URL",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": {
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_port": {
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"inlined": {
												Description:         "Inlined manifest",
												MarkdownDescription: "Inlined manifest",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"location_type": {
												Description:         "Type of Kubernetes-like location",
												MarkdownDescription: "Type of Kubernetes-like location",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Uri", "Inlined"),
												},
											},

											"uri": {
												Description:         "Location in a file fetched from a uri.",
												MarkdownDescription: "Location in a file fetched from a uri.",

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

									"name": {
										Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
										MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtMost(63),

											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
										},
									},

									"openshift": {
										Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
										MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"deploy_by_default": {
												Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
												MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"endpoints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"annotation": {
														Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
														MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"attributes": {
														Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exposure": {
														Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("public", "internal", "none"),
														},
													},

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": {
														Description:         "Path of the endpoint URL",
														MarkdownDescription: "Path of the endpoint URL",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"protocol": {
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": {
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_port": {
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"inlined": {
												Description:         "Inlined manifest",
												MarkdownDescription: "Inlined manifest",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"location_type": {
												Description:         "Type of Kubernetes-like location",
												MarkdownDescription: "Type of Kubernetes-like location",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Uri", "Inlined"),
												},
											},

											"uri": {
												Description:         "Location in a file fetched from a uri.",
												MarkdownDescription: "Location in a file fetched from a uri.",

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

									"plugin": {
										Description:         "Allows importing a plugin.  Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
										MarkdownDescription: "Allows importing a plugin.  Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"commands": {
												Description:         "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
												MarkdownDescription: "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"apply": {
														Description:         "Command that consists in applying a given component definition, typically bound to a devworkspace event.  For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'.  When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
														MarkdownDescription: "Command that consists in applying a given component definition, typically bound to a devworkspace event.  For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'.  When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"component": {
																Description:         "Describes component that will be applied",
																MarkdownDescription: "Describes component that will be applied",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"group": {
																Description:         "Defines the group this command is part of",
																MarkdownDescription: "Defines the group this command is part of",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"is_default": {
																		Description:         "Identifies the default command for a given group kind",
																		MarkdownDescription: "Identifies the default command for a given group kind",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"kind": {
																		Description:         "Kind of group the command is part of",
																		MarkdownDescription: "Kind of group the command is part of",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label": {
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",

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

													"attributes": {
														Description:         "Map of implementation-dependant free-form YAML attributes.",
														MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"command_type": {
														Description:         "Type of devworkspace command",
														MarkdownDescription: "Type of devworkspace command",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Exec", "Apply", "Composite"),
														},
													},

													"composite": {
														Description:         "Composite command that allows executing several sub-commands either sequentially or concurrently",
														MarkdownDescription: "Composite command that allows executing several sub-commands either sequentially or concurrently",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"commands": {
																Description:         "The commands that comprise this composite command",
																MarkdownDescription: "The commands that comprise this composite command",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"group": {
																Description:         "Defines the group this command is part of",
																MarkdownDescription: "Defines the group this command is part of",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"is_default": {
																		Description:         "Identifies the default command for a given group kind",
																		MarkdownDescription: "Identifies the default command for a given group kind",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"kind": {
																		Description:         "Kind of group the command is part of",
																		MarkdownDescription: "Kind of group the command is part of",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label": {
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"parallel": {
																Description:         "Indicates if the sub-commands should be executed concurrently",
																MarkdownDescription: "Indicates if the sub-commands should be executed concurrently",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exec": {
														Description:         "CLI Command executed in an existing component container",
														MarkdownDescription: "CLI Command executed in an existing component container",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command_line": {
																Description:         "The actual command-line string  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																MarkdownDescription: "The actual command-line string  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"component": {
																Description:         "Describes component to which given action relates",
																MarkdownDescription: "Describes component to which given action relates",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"env": {
																Description:         "Optional list of environment variables that have to be set before running the command",
																MarkdownDescription: "Optional list of environment variables that have to be set before running the command",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
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

															"group": {
																Description:         "Defines the group this command is part of",
																MarkdownDescription: "Defines the group this command is part of",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"is_default": {
																		Description:         "Identifies the default command for a given group kind",
																		MarkdownDescription: "Identifies the default command for a given group kind",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"kind": {
																		Description:         "Kind of group the command is part of",
																		MarkdownDescription: "Kind of group the command is part of",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"hot_reload_capable": {
																Description:         "Whether the command is capable to reload itself when source code changes. If set to 'true' the command won't be restarted and it is expected to handle file changes on its own.  Default value is 'false'",
																MarkdownDescription: "Whether the command is capable to reload itself when source code changes. If set to 'true' the command won't be restarted and it is expected to handle file changes on its own.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"label": {
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"working_dir": {
																Description:         "Working directory where the command should be executed  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																MarkdownDescription: "Working directory where the command should be executed  Special variables that can be used:   - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping.   - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",

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

													"id": {
														Description:         "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
														MarkdownDescription: "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"components": {
												Description:         "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
												MarkdownDescription: "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"attributes": {
														Description:         "Map of implementation-dependant free-form YAML attributes.",
														MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"component_type": {
														Description:         "Type of component",
														MarkdownDescription: "Type of component",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image"),
														},
													},

													"container": {
														Description:         "Allows adding and configuring devworkspace-related containers",
														MarkdownDescription: "Allows adding and configuring devworkspace-related containers",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"annotation": {
																Description:         "Annotations that should be added to specific resources for this container",
																MarkdownDescription: "Annotations that should be added to specific resources for this container",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"deployment": {
																		Description:         "Annotations to be added to deployment",
																		MarkdownDescription: "Annotations to be added to deployment",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"service": {
																		Description:         "Annotations to be added to service",
																		MarkdownDescription: "Annotations to be added to service",

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

															"args": {
																Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",
																MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.  Defaults to an empty array, meaning use whatever is defined in the image.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"command": {
																Description:         "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",
																MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image.  Defaults to an empty array, meaning use whatever is defined in the image.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"cpu_limit": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"cpu_request": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"dedicated_pod": {
																Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",
																MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoints": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"annotation": {
																		Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																		MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"attributes": {
																		Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																		MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"exposure": {
																		Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																		MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("public", "internal", "none"),
																		},
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtMost(63),

																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"path": {
																		Description:         "Path of the endpoint URL",
																		MarkdownDescription: "Path of the endpoint URL",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"protocol": {
																		Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																		MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																		},
																	},

																	"secure": {
																		Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																		MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"target_port": {
																		Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																		MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"env": {
																Description:         "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",
																MarkdownDescription: "Environment variables used in this container.  The following variables are reserved and cannot be overridden via env:   - '$PROJECTS_ROOT'   - '$PROJECT_SOURCE'",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
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

															"image": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"memory_limit": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"memory_request": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"mount_sources": {
																Description:         "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component.  Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"source_mapping": {
																Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
																MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_mounts": {
																Description:         "List of volumes mounts that should be mounted is this container.",
																MarkdownDescription: "List of volumes mounts that should be mounted is this container.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																		MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtMost(63),

																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"path": {
																		Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
																		MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",

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

													"image": {
														Description:         "Allows specifying the definition of an image for outer loop builds",
														MarkdownDescription: "Allows specifying the definition of an image for outer loop builds",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"auto_build": {
																Description:         "Defines if the image should be built during startup.  Default value is 'false'",
																MarkdownDescription: "Defines if the image should be built during startup.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"dockerfile": {
																Description:         "Allows specifying dockerfile type build",
																MarkdownDescription: "Allows specifying dockerfile type build",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"args": {
																		Description:         "The arguments to supply to the dockerfile build.",
																		MarkdownDescription: "The arguments to supply to the dockerfile build.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"build_context": {
																		Description:         "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
																		MarkdownDescription: "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"devfile_registry": {
																		Description:         "Dockerfile's Devfile Registry source",
																		MarkdownDescription: "Dockerfile's Devfile Registry source",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"id": {
																				Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																				MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"registry_url": {
																				Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																				MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",

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

																	"git": {
																		Description:         "Dockerfile's Git source",
																		MarkdownDescription: "Dockerfile's Git source",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"checkout_from": {
																				Description:         "Defines from what the project should be checked out. Required if there are more than one remote configured",
																				MarkdownDescription: "Defines from what the project should be checked out. Required if there are more than one remote configured",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"remote": {
																						Description:         "The remote name should be used as init. Required if there are more than one remote configured",
																						MarkdownDescription: "The remote name should be used as init. Required if there are more than one remote configured",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"revision": {
																						Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																						MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",

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

																			"file_location": {
																				Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																				MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"remotes": {
																				Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																				MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",

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

																	"root_required": {
																		Description:         "Specify if a privileged builder pod is required.  Default value is 'false'",
																		MarkdownDescription: "Specify if a privileged builder pod is required.  Default value is 'false'",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"src_type": {
																		Description:         "Type of Dockerfile src",
																		MarkdownDescription: "Type of Dockerfile src",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("Uri", "DevfileRegistry", "Git"),
																		},
																	},

																	"uri": {
																		Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
																		MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",

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

															"image_name": {
																Description:         "Name of the image for the resulting outerloop build",
																MarkdownDescription: "Name of the image for the resulting outerloop build",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image_type": {
																Description:         "Type of image",
																MarkdownDescription: "Type of image",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Dockerfile", "AutoBuild"),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kubernetes": {
														Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
														MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"deploy_by_default": {
																Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
																MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoints": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"annotation": {
																		Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																		MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"attributes": {
																		Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																		MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"exposure": {
																		Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																		MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("public", "internal", "none"),
																		},
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtMost(63),

																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"path": {
																		Description:         "Path of the endpoint URL",
																		MarkdownDescription: "Path of the endpoint URL",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"protocol": {
																		Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																		MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																		},
																	},

																	"secure": {
																		Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																		MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"target_port": {
																		Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																		MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"inlined": {
																Description:         "Inlined manifest",
																MarkdownDescription: "Inlined manifest",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"location_type": {
																Description:         "Type of Kubernetes-like location",
																MarkdownDescription: "Type of Kubernetes-like location",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Uri", "Inlined"),
																},
															},

															"uri": {
																Description:         "Location in a file fetched from a uri.",
																MarkdownDescription: "Location in a file fetched from a uri.",

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

													"name": {
														Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
														MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtMost(63),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"openshift": {
														Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
														MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"deploy_by_default": {
																Description:         "Defines if the component should be deployed during startup.  Default value is 'false'",
																MarkdownDescription: "Defines if the component should be deployed during startup.  Default value is 'false'",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"endpoints": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"annotation": {
																		Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																		MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"attributes": {
																		Description:         "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",
																		MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.  Examples of Che-specific attributes:  - cookiesAuthEnabled: 'true' / 'false',  - type: 'terminal' / 'ide' / 'ide-dev',",

																		Type: utilities.DynamicType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"exposure": {
																		Description:         "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",
																		MarkdownDescription: "Describes how the endpoint should be exposed on the network.  - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route.  - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network.  - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address.  Default value is 'public'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("public", "internal", "none"),
																		},
																	},

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtMost(63),

																			stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																		},
																	},

																	"path": {
																		Description:         "Path of the endpoint URL",
																		MarkdownDescription: "Path of the endpoint URL",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"protocol": {
																		Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",
																		MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.  - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.  - 'https': Endpoint will have 'https' traffic, typically on a TCP connection.  - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.  - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.  - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.  - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.  Default value is 'http'",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																		},
																	},

																	"secure": {
																		Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																		MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"target_port": {
																		Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																		MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"inlined": {
																Description:         "Inlined manifest",
																MarkdownDescription: "Inlined manifest",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"location_type": {
																Description:         "Type of Kubernetes-like location",
																MarkdownDescription: "Type of Kubernetes-like location",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Uri", "Inlined"),
																},
															},

															"uri": {
																Description:         "Location in a file fetched from a uri.",
																MarkdownDescription: "Location in a file fetched from a uri.",

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

													"volume": {
														Description:         "Allows specifying the definition of a volume shared by several other components",
														MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"ephemeral": {
																Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
																MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"size": {
																Description:         "Size of the volume",
																MarkdownDescription: "Size of the volume",

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

											"id": {
												Description:         "Id in a registry that contains a Devfile yaml file",
												MarkdownDescription: "Id in a registry that contains a Devfile yaml file",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"import_reference_type": {
												Description:         "type of location from where the referenced template structure should be retrieved",
												MarkdownDescription: "type of location from where the referenced template structure should be retrieved",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Uri", "Id", "Kubernetes"),
												},
											},

											"kubernetes": {
												Description:         "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",
												MarkdownDescription: "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"namespace": {
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

											"registry_url": {
												Description:         "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
												MarkdownDescription: "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"uri": {
												Description:         "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
												MarkdownDescription: "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"version": {
												Description:         "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
												MarkdownDescription: "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^(latest)|(([1-9])\.([0-9]+)\.([0-9]+)(\-[0-9a-z-]+(\.[0-9a-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?)$`), ""),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume": {
										Description:         "Allows specifying the definition of a volume shared by several other components",
										MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ephemeral": {
												Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
												MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "Size of the volume",
												MarkdownDescription: "Size of the volume",

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

							"default_editor": {
								Description:         "The default editor to workspace create with. It could be a plugin ID or a URI. The plugin ID must have 'publisher/plugin/version' format. The URI must start from 'http://' or 'https://'.",
								MarkdownDescription: "The default editor to workspace create with. It could be a plugin ID or a URI. The plugin ID must have 'publisher/plugin/version' format. The URI must start from 'http://' or 'https://'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"default_namespace": {
								Description:         "User's default namespace.",
								MarkdownDescription: "User's default namespace.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auto_provision": {
										Description:         "Indicates if is allowed to automatically create a user namespace. If it set to false, then user namespace must be pre-created by a cluster administrator.",
										MarkdownDescription: "Indicates if is allowed to automatically create a user namespace. If it set to false, then user namespace must be pre-created by a cluster administrator.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"template": {
										Description:         "If you don't create the user namespaces in advance, this field defines the Kubernetes namespace created when you start your first workspace. You can use '<username>' and '<userid>' placeholders, such as che-workspace-<username>.",
										MarkdownDescription: "If you don't create the user namespaces in advance, this field defines the Kubernetes namespace created when you start your first workspace. You can use '<username>' and '<userid>' placeholders, such as che-workspace-<username>.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`<username>|<userid>`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"default_plugins": {
								Description:         "Default plug-ins applied to DevWorkspaces.",
								MarkdownDescription: "Default plug-ins applied to DevWorkspaces.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"editor": {
										Description:         "The editor ID to specify default plug-ins for.",
										MarkdownDescription: "The editor ID to specify default plug-ins for.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"plugins": {
										Description:         "Default plug-in URIs for the specified editor.",
										MarkdownDescription: "Default plug-in URIs for the specified editor.",

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

							"disable_container_build_capabilities": {
								Description:         "Disables the container build capabilities.",
								MarkdownDescription: "Disables the container build capabilities.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_selector": {
								Description:         "The node selector limits the nodes that can run the workspace pods.",
								MarkdownDescription: "The node selector limits the nodes that can run the workspace pods.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"seconds_of_inactivity_before_idling": {
								Description:         "Idle timeout for workspaces in seconds. This timeout is the duration after which a workspace will be idled if there is no activity. To disable workspace idling due to inactivity, set this value to -1.",
								MarkdownDescription: "Idle timeout for workspaces in seconds. This timeout is the duration after which a workspace will be idled if there is no activity. To disable workspace idling due to inactivity, set this value to -1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"seconds_of_run_before_idling": {
								Description:         "Run timeout for workspaces in seconds. This timeout is the maximum duration a workspace runs. To disable workspace run timeout, set this value to -1.",
								MarkdownDescription: "Run timeout for workspaces in seconds. This timeout is the maximum duration a workspace runs. To disable workspace run timeout, set this value to -1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage": {
								Description:         "Workspaces persistent storage.",
								MarkdownDescription: "Workspaces persistent storage.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"per_user_strategy_pvc_config": {
										Description:         "PVC settings when using the 'per-user' PVC strategy.",
										MarkdownDescription: "PVC settings when using the 'per-user' PVC strategy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"claim_size": {
												Description:         "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",
												MarkdownDescription: "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_class": {
												Description:         "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",
												MarkdownDescription: "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",

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

									"per_workspace_strategy_pvc_config": {
										Description:         "PVC settings when using the 'per-workspace' PVC strategy.",
										MarkdownDescription: "PVC settings when using the 'per-workspace' PVC strategy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"claim_size": {
												Description:         "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",
												MarkdownDescription: "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_class": {
												Description:         "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",
												MarkdownDescription: "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",

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

									"pvc_strategy": {
										Description:         "Persistent volume claim strategy for the Che server. The supported strategies are: 'per-user' (all workspaces PVCs in one volume) and 'per-workspace' (each workspace is given its own individual PVC). For details, see https://github.com/eclipse/che/issues/21185.",
										MarkdownDescription: "Persistent volume claim strategy for the Che server. The supported strategies are: 'per-user' (all workspaces PVCs in one volume) and 'per-workspace' (each workspace is given its own individual PVC). For details, see https://github.com/eclipse/che/issues/21185.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("common", "per-user", "per-workspace"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tolerations": {
								Description:         "The pod tolerations of the workspace pods limit where the workspace pods can run.",
								MarkdownDescription: "The pod tolerations of the workspace pods limit where the workspace pods can run.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"effect": {
										Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
										MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key": {
										Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
										MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"operator": {
										Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
										MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"toleration_seconds": {
										Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
										MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
										MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

							"trusted_certs": {
								Description:         "Trusted certificate settings.",
								MarkdownDescription: "Trusted certificate settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"git_trusted_certs_config_map_name": {
										Description:         "The ConfigMap contains certificates to propagate to the Che components and to provide a particular configuration for Git. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/deploying-che-with-support-for-git-repositories-with-self-signed-certificates/ The ConfigMap must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
										MarkdownDescription: "The ConfigMap contains certificates to propagate to the Che components and to provide a particular configuration for Git. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/deploying-che-with-support-for-git-repositories-with-self-signed-certificates/ The ConfigMap must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",

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

					"git_services": {
						Description:         "A configuration that allows users to work with remote Git repositories.",
						MarkdownDescription: "A configuration that allows users to work with remote Git repositories.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bitbucket": {
								Description:         "Enables users to work with repositories hosted on Bitbucket (bitbucket.org or self-hosted).",
								MarkdownDescription: "Enables users to work with repositories hosted on Bitbucket (bitbucket.org or self-hosted).",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"endpoint": {
										Description:         "Bitbucket server endpoint URL.",
										MarkdownDescription: "Bitbucket server endpoint URL.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "Kubernetes secret, that contains Base64-encoded Bitbucket OAuth 1.0 or OAuth 2.0 data. For OAuth 1.0: private key, Bitbucket Application link consumer key and Bitbucket Application link shared secret must be stored in 'private.key', 'consumer.key' and 'shared_secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/. For OAuth 2.0: Bitbucket OAuth consumer key and Bitbucket OAuth consumer secret must be stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-the-bitbucket-cloud/.",
										MarkdownDescription: "Kubernetes secret, that contains Base64-encoded Bitbucket OAuth 1.0 or OAuth 2.0 data. For OAuth 1.0: private key, Bitbucket Application link consumer key and Bitbucket Application link shared secret must be stored in 'private.key', 'consumer.key' and 'shared_secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/. For OAuth 2.0: Bitbucket OAuth consumer key and Bitbucket OAuth consumer secret must be stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-the-bitbucket-cloud/.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"github": {
								Description:         "Enables users to work with repositories hosted on GitHub (github.com or GitHub Enterprise).",
								MarkdownDescription: "Enables users to work with repositories hosted on GitHub (github.com or GitHub Enterprise).",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"disable_subdomain_isolation": {
										Description:         "Disables subdomain isolation.",
										MarkdownDescription: "Disables subdomain isolation.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"endpoint": {
										Description:         "GitHub server endpoint URL.",
										MarkdownDescription: "GitHub server endpoint URL.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_name": {
										Description:         "Kubernetes secret, that contains Base64-encoded GitHub OAuth Client id and GitHub OAuth Client secret, that stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
										MarkdownDescription: "Kubernetes secret, that contains Base64-encoded GitHub OAuth Client id and GitHub OAuth Client secret, that stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gitlab": {
								Description:         "Enables users to work with repositories hosted on GitLab (gitlab.com or self-hosted).",
								MarkdownDescription: "Enables users to work with repositories hosted on GitLab (gitlab.com or self-hosted).",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"endpoint": {
										Description:         "GitLab server endpoint URL.",
										MarkdownDescription: "GitLab server endpoint URL.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_name": {
										Description:         "Kubernetes secret, that contains Base64-encoded GitHub Application id and GitLab Application Client secret, that stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
										MarkdownDescription: "Kubernetes secret, that contains Base64-encoded GitHub Application id and GitLab Application Client secret, that stored in 'id' and 'secret' keys respectively. See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",

										Type: types.StringType,

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

					"networking": {
						Description:         "Networking, Che authentication, and TLS configuration.",
						MarkdownDescription: "Networking, Che authentication, and TLS configuration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotations": {
								Description:         "Defines annotations which will be set for an Ingress (a route for OpenShift platform). The defaults for kubernetes platforms are:     kubernetes.io/ingress.class:                       'nginx'     nginx.ingress.kubernetes.io/proxy-read-timeout:    '3600',     nginx.ingress.kubernetes.io/proxy-connect-timeout: '3600',     nginx.ingress.kubernetes.io/ssl-redirect:          'true'",
								MarkdownDescription: "Defines annotations which will be set for an Ingress (a route for OpenShift platform). The defaults for kubernetes platforms are:     kubernetes.io/ingress.class:                       'nginx'     nginx.ingress.kubernetes.io/proxy-read-timeout:    '3600',     nginx.ingress.kubernetes.io/proxy-connect-timeout: '3600',     nginx.ingress.kubernetes.io/ssl-redirect:          'true'",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"auth": {
								Description:         "Authentication settings.",
								MarkdownDescription: "Authentication settings.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"gateway": {
										Description:         "Gateway settings.",
										MarkdownDescription: "Gateway settings.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_labels": {
												Description:         "Gateway configuration labels.",
												MarkdownDescription: "Gateway configuration labels.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"deployment": {
												Description:         "Deployment override options. Since gateway deployment consists of several containers, they must be distinguished in the configuration by their names: - 'gateway' - 'configbump' - 'oauth-proxy' - 'kube-rbac-proxy'",
												MarkdownDescription: "Deployment override options. Since gateway deployment consists of several containers, they must be distinguished in the configuration by their names: - 'gateway' - 'configbump' - 'oauth-proxy' - 'kube-rbac-proxy'",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"containers": {
														Description:         "List of containers belonging to the pod.",
														MarkdownDescription: "List of containers belonging to the pod.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"env": {
																Description:         "List of environment variables to set in the container.",
																MarkdownDescription: "List of environment variables to set in the container.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																		MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																		MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"config_map_key_ref": {
																				Description:         "Selects a key of a ConfigMap.",
																				MarkdownDescription: "Selects a key of a ConfigMap.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "The key to select.",
																						MarkdownDescription: "The key to select.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"optional": {
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

																						Type: types.BoolType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_ref": {
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"api_version": {
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"field_path": {
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource_field_ref": {
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"container_name": {
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"divisor": {
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",

																						Type: utilities.IntOrStringType{},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"resource": {
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},
																				}),

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_key_ref": {
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "The key of the secret to select from.  Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"name": {
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"optional": {
																						Description:         "Specify whether the Secret or its key must be defined",
																						MarkdownDescription: "Specify whether the Secret or its key must be defined",

																						Type: types.BoolType,

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

															"image": {
																Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
																MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image_pull_policy": {
																Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
																MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
																},
															},

															"name": {
																Description:         "Container name.",
																MarkdownDescription: "Container name.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"resources": {
																Description:         "Compute resources required by this container.",
																MarkdownDescription: "Compute resources required by this container.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"limits": {
																		Description:         "Describes the maximum amount of compute resources allowed.",
																		MarkdownDescription: "Describes the maximum amount of compute resources allowed.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"cpu": {
																				Description:         "CPU, in cores. (500m = .5 cores)",
																				MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"memory": {
																				Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																				MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"request": {
																		Description:         "Describes the minimum amount of compute resources required.",
																		MarkdownDescription: "Describes the minimum amount of compute resources required.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"cpu": {
																				Description:         "CPU, in cores. (500m = .5 cores)",
																				MarkdownDescription: "CPU, in cores. (500m = .5 cores)",

																				Type: utilities.IntOrStringType{},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"memory": {
																				Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
																				MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",

																				Type: utilities.IntOrStringType{},

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

													"security_context": {
														Description:         "Security options the pod should run with.",
														MarkdownDescription: "Security options the pod should run with.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"fs_group": {
																Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
																MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"run_as_user": {
																Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
																MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",

																Type: types.Int64Type,

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

									"identity_provider_url": {
										Description:         "Public URL of the Identity Provider server.",
										MarkdownDescription: "Public URL of the Identity Provider server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"identity_token": {
										Description:         "Identity token to be passed to upstream. There are two types of tokens supported: 'id_token' and 'access_token'. Default value is 'id_token'. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
										MarkdownDescription: "Identity token to be passed to upstream. There are two types of tokens supported: 'id_token' and 'access_token'. Default value is 'id_token'. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("id_token", "access_token"),
										},
									},

									"o_auth_access_token_inactivity_timeout_seconds": {
										Description:         "Inactivity timeout for tokens to set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side. 0 means tokens for this client never time out.",
										MarkdownDescription: "Inactivity timeout for tokens to set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side. 0 means tokens for this client never time out.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"o_auth_access_token_max_age_seconds": {
										Description:         "Access token max age for tokens to set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side. 0 means no expiration.",
										MarkdownDescription: "Access token max age for tokens to set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side. 0 means no expiration.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"o_auth_client_name": {
										Description:         "Name of the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.",
										MarkdownDescription: "Name of the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"o_auth_scope": {
										Description:         "Access Token Scope. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
										MarkdownDescription: "Access Token Scope. This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"o_auth_secret": {
										Description:         "Name of the secret set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.",
										MarkdownDescription: "Name of the secret set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.",

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

							"domain": {
								Description:         "For an OpenShift cluster, the Operator uses the domain to generate a hostname for the route. The generated hostname follows this pattern: che-<che-namespace>.<domain>. The <che-namespace> is the namespace where the CheCluster CRD is created. In conjunction with labels, it creates a route served by a non-default Ingress controller. For a Kubernetes cluster, it contains a global ingress domain. There are no default values: you must specify them.",
								MarkdownDescription: "For an OpenShift cluster, the Operator uses the domain to generate a hostname for the route. The generated hostname follows this pattern: che-<che-namespace>.<domain>. The <che-namespace> is the namespace where the CheCluster CRD is created. In conjunction with labels, it creates a route served by a non-default Ingress controller. For a Kubernetes cluster, it contains a global ingress domain. There are no default values: you must specify them.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": {
								Description:         "The public hostname of the installed Che server.",
								MarkdownDescription: "The public hostname of the installed Che server.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"labels": {
								Description:         "Defines labels which will be set for an Ingress (a route for OpenShift platform).",
								MarkdownDescription: "Defines labels which will be set for an Ingress (a route for OpenShift platform).",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_secret_name": {
								Description:         "The name of the secret used to set up Ingress TLS termination. If the field is an empty string, the default cluster certificate is used. The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "The name of the secret used to set up Ingress TLS termination. If the field is an empty string, the default cluster certificate is used. The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",

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
		},
	}, nil
}

func (r *OrgEclipseCheCheClusterV2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_org_eclipse_che_che_cluster_v2")

	var state OrgEclipseCheCheClusterV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OrgEclipseCheCheClusterV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("org.eclipse.che/v2")
	goModel.Kind = utilities.Ptr("CheCluster")

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

func (r *OrgEclipseCheCheClusterV2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_org_eclipse_che_che_cluster_v2")
	// NO-OP: All data is already in Terraform state
}

func (r *OrgEclipseCheCheClusterV2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_org_eclipse_che_che_cluster_v2")

	var state OrgEclipseCheCheClusterV2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OrgEclipseCheCheClusterV2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("org.eclipse.che/v2")
	goModel.Kind = utilities.Ptr("CheCluster")

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

func (r *OrgEclipseCheCheClusterV2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_org_eclipse_che_che_cluster_v2")
	// NO-OP: Terraform removes the state automatically for us
}
