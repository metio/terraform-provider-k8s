/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package org_eclipse_che_v2

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
	_ datasource.DataSource = &OrgEclipseCheCheClusterV2Manifest{}
)

func NewOrgEclipseCheCheClusterV2Manifest() datasource.DataSource {
	return &OrgEclipseCheCheClusterV2Manifest{}
}

type OrgEclipseCheCheClusterV2Manifest struct{}

type OrgEclipseCheCheClusterV2ManifestData struct {
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
		Components *struct {
			CheServer *struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
				Debug        *bool     `tfsdk:"debug" json:"debug,omitempty"`
				Deployment   *struct {
					Containers *[]struct {
						Env *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Value     *string `tfsdk:"value" json:"value,omitempty"`
							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
									FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
									Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
									Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"env" json:"env,omitempty"`
						Image           *string `tfsdk:"image" json:"image,omitempty"`
						ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Resources       *struct {
							Limits *struct {
								Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
								Memory *string `tfsdk:"memory" json:"memory,omitempty"`
							} `tfsdk:"limits" json:"limits,omitempty"`
							Request *struct {
								Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
								Memory *string `tfsdk:"memory" json:"memory,omitempty"`
							} `tfsdk:"request" json:"request,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					SecurityContext *struct {
						FsGroup   *int64 `tfsdk:"fs_group" json:"fsGroup,omitempty"`
						RunAsUser *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					} `tfsdk:"security_context" json:"securityContext,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
				ExtraProperties *map[string]string `tfsdk:"extra_properties" json:"extraProperties,omitempty"`
				LogLevel        *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
				Proxy           *struct {
					CredentialsSecretName *string   `tfsdk:"credentials_secret_name" json:"credentialsSecretName,omitempty"`
					NonProxyHosts         *[]string `tfsdk:"non_proxy_hosts" json:"nonProxyHosts,omitempty"`
					Port                  *string   `tfsdk:"port" json:"port,omitempty"`
					Url                   *string   `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
			} `tfsdk:"che_server" json:"cheServer,omitempty"`
			Dashboard *struct {
				Branding *struct {
					Logo *struct {
						Base64data *string `tfsdk:"base64data" json:"base64data,omitempty"`
						Mediatype  *string `tfsdk:"mediatype" json:"mediatype,omitempty"`
					} `tfsdk:"logo" json:"logo,omitempty"`
				} `tfsdk:"branding" json:"branding,omitempty"`
				Deployment *struct {
					Containers *[]struct {
						Env *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Value     *string `tfsdk:"value" json:"value,omitempty"`
							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
									FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
									Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
									Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"env" json:"env,omitempty"`
						Image           *string `tfsdk:"image" json:"image,omitempty"`
						ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Resources       *struct {
							Limits *struct {
								Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
								Memory *string `tfsdk:"memory" json:"memory,omitempty"`
							} `tfsdk:"limits" json:"limits,omitempty"`
							Request *struct {
								Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
								Memory *string `tfsdk:"memory" json:"memory,omitempty"`
							} `tfsdk:"request" json:"request,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					SecurityContext *struct {
						FsGroup   *int64 `tfsdk:"fs_group" json:"fsGroup,omitempty"`
						RunAsUser *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					} `tfsdk:"security_context" json:"securityContext,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
				HeaderMessage *struct {
					Show *bool   `tfsdk:"show" json:"show,omitempty"`
					Text *string `tfsdk:"text" json:"text,omitempty"`
				} `tfsdk:"header_message" json:"headerMessage,omitempty"`
				LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			} `tfsdk:"dashboard" json:"dashboard,omitempty"`
			DevWorkspace *struct {
				RunningLimit *string `tfsdk:"running_limit" json:"runningLimit,omitempty"`
			} `tfsdk:"dev_workspace" json:"devWorkspace,omitempty"`
			DevfileRegistry *struct {
				Deployment *struct {
					Containers *[]struct {
						Env *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Value     *string `tfsdk:"value" json:"value,omitempty"`
							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
									FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
									Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
									Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"env" json:"env,omitempty"`
						Image           *string `tfsdk:"image" json:"image,omitempty"`
						ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Resources       *struct {
							Limits *struct {
								Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
								Memory *string `tfsdk:"memory" json:"memory,omitempty"`
							} `tfsdk:"limits" json:"limits,omitempty"`
							Request *struct {
								Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
								Memory *string `tfsdk:"memory" json:"memory,omitempty"`
							} `tfsdk:"request" json:"request,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					SecurityContext *struct {
						FsGroup   *int64 `tfsdk:"fs_group" json:"fsGroup,omitempty"`
						RunAsUser *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					} `tfsdk:"security_context" json:"securityContext,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
				DisableInternalRegistry   *bool `tfsdk:"disable_internal_registry" json:"disableInternalRegistry,omitempty"`
				ExternalDevfileRegistries *[]struct {
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"external_devfile_registries" json:"externalDevfileRegistries,omitempty"`
			} `tfsdk:"devfile_registry" json:"devfileRegistry,omitempty"`
			ImagePuller *struct {
				Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				Spec   *struct {
					Affinity             *string `tfsdk:"affinity" json:"affinity,omitempty"`
					CachingCPULimit      *string `tfsdk:"caching_cpu_limit" json:"cachingCPULimit,omitempty"`
					CachingCPURequest    *string `tfsdk:"caching_cpu_request" json:"cachingCPURequest,omitempty"`
					CachingIntervalHours *string `tfsdk:"caching_interval_hours" json:"cachingIntervalHours,omitempty"`
					CachingMemoryLimit   *string `tfsdk:"caching_memory_limit" json:"cachingMemoryLimit,omitempty"`
					CachingMemoryRequest *string `tfsdk:"caching_memory_request" json:"cachingMemoryRequest,omitempty"`
					ConfigMapName        *string `tfsdk:"config_map_name" json:"configMapName,omitempty"`
					DaemonsetName        *string `tfsdk:"daemonset_name" json:"daemonsetName,omitempty"`
					DeploymentName       *string `tfsdk:"deployment_name" json:"deploymentName,omitempty"`
					ImagePullSecrets     *string `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
					ImagePullerImage     *string `tfsdk:"image_puller_image" json:"imagePullerImage,omitempty"`
					Images               *string `tfsdk:"images" json:"images,omitempty"`
					NodeSelector         *string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"image_puller" json:"imagePuller,omitempty"`
			Metrics *struct {
				Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			PluginRegistry *struct {
				Deployment *struct {
					Containers *[]struct {
						Env *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Value     *string `tfsdk:"value" json:"value,omitempty"`
							ValueFrom *struct {
								ConfigMapKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
								FieldRef *struct {
									ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
									FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
								} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
								ResourceFieldRef *struct {
									ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
									Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
									Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
								} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
								SecretKeyRef *struct {
									Key      *string `tfsdk:"key" json:"key,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
									Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
								} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
							} `tfsdk:"value_from" json:"valueFrom,omitempty"`
						} `tfsdk:"env" json:"env,omitempty"`
						Image           *string `tfsdk:"image" json:"image,omitempty"`
						ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Resources       *struct {
							Limits *struct {
								Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
								Memory *string `tfsdk:"memory" json:"memory,omitempty"`
							} `tfsdk:"limits" json:"limits,omitempty"`
							Request *struct {
								Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
								Memory *string `tfsdk:"memory" json:"memory,omitempty"`
							} `tfsdk:"request" json:"request,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
					} `tfsdk:"containers" json:"containers,omitempty"`
					SecurityContext *struct {
						FsGroup   *int64 `tfsdk:"fs_group" json:"fsGroup,omitempty"`
						RunAsUser *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					} `tfsdk:"security_context" json:"securityContext,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
				DisableInternalRegistry  *bool `tfsdk:"disable_internal_registry" json:"disableInternalRegistry,omitempty"`
				ExternalPluginRegistries *[]struct {
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"external_plugin_registries" json:"externalPluginRegistries,omitempty"`
				OpenVSXURL *string `tfsdk:"open_vsx_url" json:"openVSXURL,omitempty"`
			} `tfsdk:"plugin_registry" json:"pluginRegistry,omitempty"`
		} `tfsdk:"components" json:"components,omitempty"`
		ContainerRegistry *struct {
			Hostname     *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Organization *string `tfsdk:"organization" json:"organization,omitempty"`
		} `tfsdk:"container_registry" json:"containerRegistry,omitempty"`
		DevEnvironments *struct {
			ContainerBuildConfiguration *struct {
				OpenShiftSecurityContextConstraint *string `tfsdk:"open_shift_security_context_constraint" json:"openShiftSecurityContextConstraint,omitempty"`
			} `tfsdk:"container_build_configuration" json:"containerBuildConfiguration,omitempty"`
			DefaultComponents *[]struct {
				Attributes    *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ComponentType *string            `tfsdk:"component_type" json:"componentType,omitempty"`
				Container     *struct {
					Annotation *struct {
						Deployment *map[string]string `tfsdk:"deployment" json:"deployment,omitempty"`
						Service    *map[string]string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"annotation" json:"annotation,omitempty"`
					Args         *[]string `tfsdk:"args" json:"args,omitempty"`
					Command      *[]string `tfsdk:"command" json:"command,omitempty"`
					CpuLimit     *string   `tfsdk:"cpu_limit" json:"cpuLimit,omitempty"`
					CpuRequest   *string   `tfsdk:"cpu_request" json:"cpuRequest,omitempty"`
					DedicatedPod *bool     `tfsdk:"dedicated_pod" json:"dedicatedPod,omitempty"`
					Endpoints    *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
						Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Path       *string            `tfsdk:"path" json:"path,omitempty"`
						Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
						Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
						TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
					} `tfsdk:"endpoints" json:"endpoints,omitempty"`
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
					Image         *string `tfsdk:"image" json:"image,omitempty"`
					MemoryLimit   *string `tfsdk:"memory_limit" json:"memoryLimit,omitempty"`
					MemoryRequest *string `tfsdk:"memory_request" json:"memoryRequest,omitempty"`
					MountSources  *bool   `tfsdk:"mount_sources" json:"mountSources,omitempty"`
					SourceMapping *string `tfsdk:"source_mapping" json:"sourceMapping,omitempty"`
					VolumeMounts  *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
				} `tfsdk:"container" json:"container,omitempty"`
				Custom *struct {
					ComponentClass   *string            `tfsdk:"component_class" json:"componentClass,omitempty"`
					EmbeddedResource *map[string]string `tfsdk:"embedded_resource" json:"embeddedResource,omitempty"`
				} `tfsdk:"custom" json:"custom,omitempty"`
				Image *struct {
					AutoBuild  *bool `tfsdk:"auto_build" json:"autoBuild,omitempty"`
					Dockerfile *struct {
						Args            *[]string `tfsdk:"args" json:"args,omitempty"`
						BuildContext    *string   `tfsdk:"build_context" json:"buildContext,omitempty"`
						DevfileRegistry *struct {
							Id          *string `tfsdk:"id" json:"id,omitempty"`
							RegistryUrl *string `tfsdk:"registry_url" json:"registryUrl,omitempty"`
						} `tfsdk:"devfile_registry" json:"devfileRegistry,omitempty"`
						Git *struct {
							CheckoutFrom *struct {
								Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
								Revision *string `tfsdk:"revision" json:"revision,omitempty"`
							} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
							FileLocation *string            `tfsdk:"file_location" json:"fileLocation,omitempty"`
							Remotes      *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
						} `tfsdk:"git" json:"git,omitempty"`
						RootRequired *bool   `tfsdk:"root_required" json:"rootRequired,omitempty"`
						SrcType      *string `tfsdk:"src_type" json:"srcType,omitempty"`
						Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
					} `tfsdk:"dockerfile" json:"dockerfile,omitempty"`
					ImageName *string `tfsdk:"image_name" json:"imageName,omitempty"`
					ImageType *string `tfsdk:"image_type" json:"imageType,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
				Kubernetes *struct {
					DeployByDefault *bool `tfsdk:"deploy_by_default" json:"deployByDefault,omitempty"`
					Endpoints       *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
						Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Path       *string            `tfsdk:"path" json:"path,omitempty"`
						Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
						Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
						TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
					} `tfsdk:"endpoints" json:"endpoints,omitempty"`
					Inlined      *string `tfsdk:"inlined" json:"inlined,omitempty"`
					LocationType *string `tfsdk:"location_type" json:"locationType,omitempty"`
					Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Openshift *struct {
					DeployByDefault *bool `tfsdk:"deploy_by_default" json:"deployByDefault,omitempty"`
					Endpoints       *[]struct {
						Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
						Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
						Name       *string            `tfsdk:"name" json:"name,omitempty"`
						Path       *string            `tfsdk:"path" json:"path,omitempty"`
						Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
						Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
						TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
					} `tfsdk:"endpoints" json:"endpoints,omitempty"`
					Inlined      *string `tfsdk:"inlined" json:"inlined,omitempty"`
					LocationType *string `tfsdk:"location_type" json:"locationType,omitempty"`
					Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
				} `tfsdk:"openshift" json:"openshift,omitempty"`
				Plugin *struct {
					Commands *[]struct {
						Apply *struct {
							Component *string `tfsdk:"component" json:"component,omitempty"`
							Group     *struct {
								IsDefault *bool   `tfsdk:"is_default" json:"isDefault,omitempty"`
								Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
							} `tfsdk:"group" json:"group,omitempty"`
							Label *string `tfsdk:"label" json:"label,omitempty"`
						} `tfsdk:"apply" json:"apply,omitempty"`
						Attributes  *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						CommandType *string            `tfsdk:"command_type" json:"commandType,omitempty"`
						Composite   *struct {
							Commands *[]string `tfsdk:"commands" json:"commands,omitempty"`
							Group    *struct {
								IsDefault *bool   `tfsdk:"is_default" json:"isDefault,omitempty"`
								Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
							} `tfsdk:"group" json:"group,omitempty"`
							Label    *string `tfsdk:"label" json:"label,omitempty"`
							Parallel *bool   `tfsdk:"parallel" json:"parallel,omitempty"`
						} `tfsdk:"composite" json:"composite,omitempty"`
						Exec *struct {
							CommandLine *string `tfsdk:"command_line" json:"commandLine,omitempty"`
							Component   *string `tfsdk:"component" json:"component,omitempty"`
							Env         *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"env" json:"env,omitempty"`
							Group *struct {
								IsDefault *bool   `tfsdk:"is_default" json:"isDefault,omitempty"`
								Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
							} `tfsdk:"group" json:"group,omitempty"`
							HotReloadCapable *bool   `tfsdk:"hot_reload_capable" json:"hotReloadCapable,omitempty"`
							Label            *string `tfsdk:"label" json:"label,omitempty"`
							WorkingDir       *string `tfsdk:"working_dir" json:"workingDir,omitempty"`
						} `tfsdk:"exec" json:"exec,omitempty"`
						Id *string `tfsdk:"id" json:"id,omitempty"`
					} `tfsdk:"commands" json:"commands,omitempty"`
					Components *[]struct {
						Attributes    *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
						ComponentType *string            `tfsdk:"component_type" json:"componentType,omitempty"`
						Container     *struct {
							Annotation *struct {
								Deployment *map[string]string `tfsdk:"deployment" json:"deployment,omitempty"`
								Service    *map[string]string `tfsdk:"service" json:"service,omitempty"`
							} `tfsdk:"annotation" json:"annotation,omitempty"`
							Args         *[]string `tfsdk:"args" json:"args,omitempty"`
							Command      *[]string `tfsdk:"command" json:"command,omitempty"`
							CpuLimit     *string   `tfsdk:"cpu_limit" json:"cpuLimit,omitempty"`
							CpuRequest   *string   `tfsdk:"cpu_request" json:"cpuRequest,omitempty"`
							DedicatedPod *bool     `tfsdk:"dedicated_pod" json:"dedicatedPod,omitempty"`
							Endpoints    *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
								Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
								Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
								Name       *string            `tfsdk:"name" json:"name,omitempty"`
								Path       *string            `tfsdk:"path" json:"path,omitempty"`
								Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
								Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
								TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
							} `tfsdk:"endpoints" json:"endpoints,omitempty"`
							Env *[]struct {
								Name  *string `tfsdk:"name" json:"name,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"env" json:"env,omitempty"`
							Image         *string `tfsdk:"image" json:"image,omitempty"`
							MemoryLimit   *string `tfsdk:"memory_limit" json:"memoryLimit,omitempty"`
							MemoryRequest *string `tfsdk:"memory_request" json:"memoryRequest,omitempty"`
							MountSources  *bool   `tfsdk:"mount_sources" json:"mountSources,omitempty"`
							SourceMapping *string `tfsdk:"source_mapping" json:"sourceMapping,omitempty"`
							VolumeMounts  *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
						} `tfsdk:"container" json:"container,omitempty"`
						Image *struct {
							AutoBuild  *bool `tfsdk:"auto_build" json:"autoBuild,omitempty"`
							Dockerfile *struct {
								Args            *[]string `tfsdk:"args" json:"args,omitempty"`
								BuildContext    *string   `tfsdk:"build_context" json:"buildContext,omitempty"`
								DevfileRegistry *struct {
									Id          *string `tfsdk:"id" json:"id,omitempty"`
									RegistryUrl *string `tfsdk:"registry_url" json:"registryUrl,omitempty"`
								} `tfsdk:"devfile_registry" json:"devfileRegistry,omitempty"`
								Git *struct {
									CheckoutFrom *struct {
										Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
										Revision *string `tfsdk:"revision" json:"revision,omitempty"`
									} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
									FileLocation *string            `tfsdk:"file_location" json:"fileLocation,omitempty"`
									Remotes      *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
								} `tfsdk:"git" json:"git,omitempty"`
								RootRequired *bool   `tfsdk:"root_required" json:"rootRequired,omitempty"`
								SrcType      *string `tfsdk:"src_type" json:"srcType,omitempty"`
								Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
							} `tfsdk:"dockerfile" json:"dockerfile,omitempty"`
							ImageName *string `tfsdk:"image_name" json:"imageName,omitempty"`
							ImageType *string `tfsdk:"image_type" json:"imageType,omitempty"`
						} `tfsdk:"image" json:"image,omitempty"`
						Kubernetes *struct {
							DeployByDefault *bool `tfsdk:"deploy_by_default" json:"deployByDefault,omitempty"`
							Endpoints       *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
								Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
								Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
								Name       *string            `tfsdk:"name" json:"name,omitempty"`
								Path       *string            `tfsdk:"path" json:"path,omitempty"`
								Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
								Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
								TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
							} `tfsdk:"endpoints" json:"endpoints,omitempty"`
							Inlined      *string `tfsdk:"inlined" json:"inlined,omitempty"`
							LocationType *string `tfsdk:"location_type" json:"locationType,omitempty"`
							Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Openshift *struct {
							DeployByDefault *bool `tfsdk:"deploy_by_default" json:"deployByDefault,omitempty"`
							Endpoints       *[]struct {
								Annotation *map[string]string `tfsdk:"annotation" json:"annotation,omitempty"`
								Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
								Exposure   *string            `tfsdk:"exposure" json:"exposure,omitempty"`
								Name       *string            `tfsdk:"name" json:"name,omitempty"`
								Path       *string            `tfsdk:"path" json:"path,omitempty"`
								Protocol   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
								Secure     *bool              `tfsdk:"secure" json:"secure,omitempty"`
								TargetPort *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
							} `tfsdk:"endpoints" json:"endpoints,omitempty"`
							Inlined      *string `tfsdk:"inlined" json:"inlined,omitempty"`
							LocationType *string `tfsdk:"location_type" json:"locationType,omitempty"`
							Uri          *string `tfsdk:"uri" json:"uri,omitempty"`
						} `tfsdk:"openshift" json:"openshift,omitempty"`
						Volume *struct {
							Ephemeral *bool   `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
							Size      *string `tfsdk:"size" json:"size,omitempty"`
						} `tfsdk:"volume" json:"volume,omitempty"`
					} `tfsdk:"components" json:"components,omitempty"`
					Id                  *string `tfsdk:"id" json:"id,omitempty"`
					ImportReferenceType *string `tfsdk:"import_reference_type" json:"importReferenceType,omitempty"`
					Kubernetes          *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
					RegistryUrl *string `tfsdk:"registry_url" json:"registryUrl,omitempty"`
					Uri         *string `tfsdk:"uri" json:"uri,omitempty"`
					Version     *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"plugin" json:"plugin,omitempty"`
				Volume *struct {
					Ephemeral *bool   `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
					Size      *string `tfsdk:"size" json:"size,omitempty"`
				} `tfsdk:"volume" json:"volume,omitempty"`
			} `tfsdk:"default_components" json:"defaultComponents,omitempty"`
			DefaultEditor    *string `tfsdk:"default_editor" json:"defaultEditor,omitempty"`
			DefaultNamespace *struct {
				AutoProvision *bool   `tfsdk:"auto_provision" json:"autoProvision,omitempty"`
				Template      *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"default_namespace" json:"defaultNamespace,omitempty"`
			DefaultPlugins *[]struct {
				Editor  *string   `tfsdk:"editor" json:"editor,omitempty"`
				Plugins *[]string `tfsdk:"plugins" json:"plugins,omitempty"`
			} `tfsdk:"default_plugins" json:"defaultPlugins,omitempty"`
			DeploymentStrategy                *string `tfsdk:"deployment_strategy" json:"deploymentStrategy,omitempty"`
			DisableContainerBuildCapabilities *bool   `tfsdk:"disable_container_build_capabilities" json:"disableContainerBuildCapabilities,omitempty"`
			GatewayContainer                  *struct {
				Env *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						ResourceFieldRef *struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
							Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				Image           *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Resources       *struct {
					Limits *struct {
						Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *string `tfsdk:"memory" json:"memory,omitempty"`
					} `tfsdk:"limits" json:"limits,omitempty"`
					Request *struct {
						Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *string `tfsdk:"memory" json:"memory,omitempty"`
					} `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"gateway_container" json:"gatewayContainer,omitempty"`
			IgnoredUnrecoverableEvents             *[]string          `tfsdk:"ignored_unrecoverable_events" json:"ignoredUnrecoverableEvents,omitempty"`
			ImagePullPolicy                        *string            `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			MaxNumberOfRunningWorkspacesPerCluster *int64             `tfsdk:"max_number_of_running_workspaces_per_cluster" json:"maxNumberOfRunningWorkspacesPerCluster,omitempty"`
			MaxNumberOfRunningWorkspacesPerUser    *int64             `tfsdk:"max_number_of_running_workspaces_per_user" json:"maxNumberOfRunningWorkspacesPerUser,omitempty"`
			MaxNumberOfWorkspacesPerUser           *int64             `tfsdk:"max_number_of_workspaces_per_user" json:"maxNumberOfWorkspacesPerUser,omitempty"`
			NodeSelector                           *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PersistUserHome                        *struct {
				DisableInitContainer *bool `tfsdk:"disable_init_container" json:"disableInitContainer,omitempty"`
				Enabled              *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"persist_user_home" json:"persistUserHome,omitempty"`
			PodSchedulerName      *string `tfsdk:"pod_scheduler_name" json:"podSchedulerName,omitempty"`
			ProjectCloneContainer *struct {
				Env *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						ResourceFieldRef *struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
							Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				Image           *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Resources       *struct {
					Limits *struct {
						Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *string `tfsdk:"memory" json:"memory,omitempty"`
					} `tfsdk:"limits" json:"limits,omitempty"`
					Request *struct {
						Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *string `tfsdk:"memory" json:"memory,omitempty"`
					} `tfsdk:"request" json:"request,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"project_clone_container" json:"projectCloneContainer,omitempty"`
			SecondsOfInactivityBeforeIdling *int64 `tfsdk:"seconds_of_inactivity_before_idling" json:"secondsOfInactivityBeforeIdling,omitempty"`
			SecondsOfRunBeforeIdling        *int64 `tfsdk:"seconds_of_run_before_idling" json:"secondsOfRunBeforeIdling,omitempty"`
			Security                        *struct {
				ContainerSecurityContext *struct {
					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
					Capabilities             *struct {
						Add  *[]string `tfsdk:"add" json:"add,omitempty"`
						Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
					} `tfsdk:"capabilities" json:"capabilities,omitempty"`
					Privileged             *bool   `tfsdk:"privileged" json:"privileged,omitempty"`
					ProcMount              *string `tfsdk:"proc_mount" json:"procMount,omitempty"`
					ReadOnlyRootFilesystem *bool   `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
					RunAsGroup             *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
					RunAsNonRoot           *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser              *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					SeLinuxOptions         *struct {
						Level *string `tfsdk:"level" json:"level,omitempty"`
						Role  *string `tfsdk:"role" json:"role,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						User  *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
					SeccompProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
						Type             *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
					WindowsOptions *struct {
						GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
						HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
						RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
				} `tfsdk:"container_security_context" json:"containerSecurityContext,omitempty"`
				PodSecurityContext *struct {
					FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
					FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
					RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
					RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					SeLinuxOptions      *struct {
						Level *string `tfsdk:"level" json:"level,omitempty"`
						Role  *string `tfsdk:"role" json:"role,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						User  *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
					SeccompProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
						Type             *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
					SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
					Sysctls            *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"sysctls" json:"sysctls,omitempty"`
					WindowsOptions *struct {
						GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
						HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
						RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
				} `tfsdk:"pod_security_context" json:"podSecurityContext,omitempty"`
			} `tfsdk:"security" json:"security,omitempty"`
			ServiceAccount       *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			ServiceAccountTokens *[]struct {
				Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
				ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
				MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				Path              *string `tfsdk:"path" json:"path,omitempty"`
			} `tfsdk:"service_account_tokens" json:"serviceAccountTokens,omitempty"`
			StartTimeoutSeconds *int64 `tfsdk:"start_timeout_seconds" json:"startTimeoutSeconds,omitempty"`
			Storage             *struct {
				PerUserStrategyPvcConfig *struct {
					ClaimSize    *string `tfsdk:"claim_size" json:"claimSize,omitempty"`
					StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				} `tfsdk:"per_user_strategy_pvc_config" json:"perUserStrategyPvcConfig,omitempty"`
				PerWorkspaceStrategyPvcConfig *struct {
					ClaimSize    *string `tfsdk:"claim_size" json:"claimSize,omitempty"`
					StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
				} `tfsdk:"per_workspace_strategy_pvc_config" json:"perWorkspaceStrategyPvcConfig,omitempty"`
				PvcStrategy *string `tfsdk:"pvc_strategy" json:"pvcStrategy,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			Tolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TrustedCerts *struct {
				GitTrustedCertsConfigMapName *string `tfsdk:"git_trusted_certs_config_map_name" json:"gitTrustedCertsConfigMapName,omitempty"`
			} `tfsdk:"trusted_certs" json:"trustedCerts,omitempty"`
			User *struct {
				ClusterRoles *[]string `tfsdk:"cluster_roles" json:"clusterRoles,omitempty"`
			} `tfsdk:"user" json:"user,omitempty"`
			WorkspacesPodAnnotations *map[string]string `tfsdk:"workspaces_pod_annotations" json:"workspacesPodAnnotations,omitempty"`
		} `tfsdk:"dev_environments" json:"devEnvironments,omitempty"`
		GitServices *struct {
			Azure *[]struct {
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"azure" json:"azure,omitempty"`
			Bitbucket *[]struct {
				Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"bitbucket" json:"bitbucket,omitempty"`
			Github *[]struct {
				DisableSubdomainIsolation *bool   `tfsdk:"disable_subdomain_isolation" json:"disableSubdomainIsolation,omitempty"`
				Endpoint                  *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				SecretName                *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"github" json:"github,omitempty"`
			Gitlab *[]struct {
				Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"gitlab" json:"gitlab,omitempty"`
		} `tfsdk:"git_services" json:"gitServices,omitempty"`
		Networking *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Auth        *struct {
				AdvancedAuthorization *struct {
					AllowGroups *[]string `tfsdk:"allow_groups" json:"allowGroups,omitempty"`
					AllowUsers  *[]string `tfsdk:"allow_users" json:"allowUsers,omitempty"`
					DenyGroups  *[]string `tfsdk:"deny_groups" json:"denyGroups,omitempty"`
					DenyUsers   *[]string `tfsdk:"deny_users" json:"denyUsers,omitempty"`
				} `tfsdk:"advanced_authorization" json:"advancedAuthorization,omitempty"`
				Gateway *struct {
					ConfigLabels *map[string]string `tfsdk:"config_labels" json:"configLabels,omitempty"`
					Deployment   *struct {
						Containers *[]struct {
							Env *[]struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Value     *string `tfsdk:"value" json:"value,omitempty"`
								ValueFrom *struct {
									ConfigMapKeyRef *struct {
										Key      *string `tfsdk:"key" json:"key,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
									FieldRef *struct {
										ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
										FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
									} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
									ResourceFieldRef *struct {
										ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
										Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
										Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
									} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
									SecretKeyRef *struct {
										Key      *string `tfsdk:"key" json:"key,omitempty"`
										Name     *string `tfsdk:"name" json:"name,omitempty"`
										Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
									} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" json:"valueFrom,omitempty"`
							} `tfsdk:"env" json:"env,omitempty"`
							Image           *string `tfsdk:"image" json:"image,omitempty"`
							ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
							Name            *string `tfsdk:"name" json:"name,omitempty"`
							Resources       *struct {
								Limits *struct {
									Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
									Memory *string `tfsdk:"memory" json:"memory,omitempty"`
								} `tfsdk:"limits" json:"limits,omitempty"`
								Request *struct {
									Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
									Memory *string `tfsdk:"memory" json:"memory,omitempty"`
								} `tfsdk:"request" json:"request,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
						} `tfsdk:"containers" json:"containers,omitempty"`
						SecurityContext *struct {
							FsGroup   *int64 `tfsdk:"fs_group" json:"fsGroup,omitempty"`
							RunAsUser *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
						} `tfsdk:"security_context" json:"securityContext,omitempty"`
					} `tfsdk:"deployment" json:"deployment,omitempty"`
					KubeRbacProxy *struct {
						LogLevel *int64 `tfsdk:"log_level" json:"logLevel,omitempty"`
					} `tfsdk:"kube_rbac_proxy" json:"kubeRbacProxy,omitempty"`
					OAuthProxy *struct {
						CookieExpireSeconds *int64 `tfsdk:"cookie_expire_seconds" json:"cookieExpireSeconds,omitempty"`
					} `tfsdk:"o_auth_proxy" json:"oAuthProxy,omitempty"`
					Traefik *struct {
						LogLevel *string `tfsdk:"log_level" json:"logLevel,omitempty"`
					} `tfsdk:"traefik" json:"traefik,omitempty"`
				} `tfsdk:"gateway" json:"gateway,omitempty"`
				IdentityProviderURL                      *string `tfsdk:"identity_provider_url" json:"identityProviderURL,omitempty"`
				IdentityToken                            *string `tfsdk:"identity_token" json:"identityToken,omitempty"`
				OAuthAccessTokenInactivityTimeoutSeconds *int64  `tfsdk:"o_auth_access_token_inactivity_timeout_seconds" json:"oAuthAccessTokenInactivityTimeoutSeconds,omitempty"`
				OAuthAccessTokenMaxAgeSeconds            *int64  `tfsdk:"o_auth_access_token_max_age_seconds" json:"oAuthAccessTokenMaxAgeSeconds,omitempty"`
				OAuthClientName                          *string `tfsdk:"o_auth_client_name" json:"oAuthClientName,omitempty"`
				OAuthScope                               *string `tfsdk:"o_auth_scope" json:"oAuthScope,omitempty"`
				OAuthSecret                              *string `tfsdk:"o_auth_secret" json:"oAuthSecret,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
			Domain           *string            `tfsdk:"domain" json:"domain,omitempty"`
			Hostname         *string            `tfsdk:"hostname" json:"hostname,omitempty"`
			IngressClassName *string            `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
			Labels           *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			TlsSecretName    *string            `tfsdk:"tls_secret_name" json:"tlsSecretName,omitempty"`
		} `tfsdk:"networking" json:"networking,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OrgEclipseCheCheClusterV2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_org_eclipse_che_che_cluster_v2_manifest"
}

func (r *OrgEclipseCheCheClusterV2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "The 'CheCluster' custom resource allows defining and managing Eclipse Che server installation.Based on these settings, the  Operator automatically creates and maintains several ConfigMaps:'che', 'plugin-registry' that will contain the appropriate environment variablesof the various components of the installation. These generated ConfigMaps must NOT be updated manually.",
		MarkdownDescription: "The 'CheCluster' custom resource allows defining and managing Eclipse Che server installation.Based on these settings, the  Operator automatically creates and maintains several ConfigMaps:'che', 'plugin-registry' that will contain the appropriate environment variablesof the various components of the installation. These generated ConfigMaps must NOT be updated manually.",
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
				Description:         "Desired configuration of Eclipse Che installation.",
				MarkdownDescription: "Desired configuration of Eclipse Che installation.",
				Attributes: map[string]schema.Attribute{
					"components": schema.SingleNestedAttribute{
						Description:         "Che components configuration.",
						MarkdownDescription: "Che components configuration.",
						Attributes: map[string]schema.Attribute{
							"che_server": schema.SingleNestedAttribute{
								Description:         "General configuration settings related to the Che server.",
								MarkdownDescription: "General configuration settings related to the Che server.",
								Attributes: map[string]schema.Attribute{
									"cluster_roles": schema.ListAttribute{
										Description:         "Additional ClusterRoles assigned to Che ServiceAccount.Each role must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.The defaults roles are:- '<che-namespace>-cheworkspaces-clusterrole'- '<che-namespace>-cheworkspaces-namespaces-clusterrole'- '<che-namespace>-cheworkspaces-devworkspace-clusterrole'where the <che-namespace> is the namespace where the CheCluster CR is created.The Che Operator must already have all permissions in these ClusterRoles to grant them.",
										MarkdownDescription: "Additional ClusterRoles assigned to Che ServiceAccount.Each role must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.The defaults roles are:- '<che-namespace>-cheworkspaces-clusterrole'- '<che-namespace>-cheworkspaces-namespaces-clusterrole'- '<che-namespace>-cheworkspaces-devworkspace-clusterrole'where the <che-namespace> is the namespace where the CheCluster CR is created.The Che Operator must already have all permissions in these ClusterRoles to grant them.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"debug": schema.BoolAttribute{
										Description:         "Enables the debug mode for Che server.",
										MarkdownDescription: "Enables the debug mode for Che server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"deployment": schema.SingleNestedAttribute{
										Description:         "Deployment override options.",
										MarkdownDescription: "Deployment override options.",
										Attributes: map[string]schema.Attribute{
											"containers": schema.ListNestedAttribute{
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"env": schema.ListNestedAttribute{
															Description:         "List of environment variables to set in the container.",
															MarkdownDescription: "List of environment variables to set in the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																		MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																		MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
																		Attributes: map[string]schema.Attribute{
																			"config_map_key_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a key of a ConfigMap.",
																				MarkdownDescription: "Selects a key of a ConfigMap.",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key to select.",
																						MarkdownDescription: "The key to select.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				Attributes: map[string]schema.Attribute{
																					"api_version": schema.StringAttribute{
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"field_path": schema.StringAttribute{
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource_field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				Attributes: map[string]schema.Attribute{
																					"container_name": schema.StringAttribute{
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"divisor": schema.StringAttribute{
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"resource": schema.StringAttribute{
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key of the secret to select from.  Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"image": schema.StringAttribute{
															Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
															MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_policy": schema.StringAttribute{
															Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
															MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
															},
														},

														"name": schema.StringAttribute{
															Description:         "Container name.",
															MarkdownDescription: "Container name.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Compute resources required by this container.",
															MarkdownDescription: "Compute resources required by this container.",
															Attributes: map[string]schema.Attribute{
																"limits": schema.SingleNestedAttribute{
																	Description:         "Describes the maximum amount of compute resources allowed.",
																	MarkdownDescription: "Describes the maximum amount of compute resources allowed.",
																	Attributes: map[string]schema.Attribute{
																		"cpu": schema.StringAttribute{
																			Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"memory": schema.StringAttribute{
																			Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"request": schema.SingleNestedAttribute{
																	Description:         "Describes the minimum amount of compute resources required.",
																	MarkdownDescription: "Describes the minimum amount of compute resources required.",
																	Attributes: map[string]schema.Attribute{
																		"cpu": schema.StringAttribute{
																			Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"memory": schema.StringAttribute{
																			Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": schema.SingleNestedAttribute{
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",
												Attributes: map[string]schema.Attribute{
													"fs_group": schema.Int64Attribute{
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user": schema.Int64Attribute{
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",
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

									"extra_properties": schema.MapAttribute{
										Description:         "A map of additional environment variables applied in the generated 'che' ConfigMap to be used by the Che serverin addition to the values already generated from other fields of the 'CheCluster' custom resource (CR).If the 'extraProperties' field contains a property normally generated in 'che' ConfigMap from other CR fields,the value defined in the 'extraProperties' is used instead.",
										MarkdownDescription: "A map of additional environment variables applied in the generated 'che' ConfigMap to be used by the Che serverin addition to the values already generated from other fields of the 'CheCluster' custom resource (CR).If the 'extraProperties' field contains a property normally generated in 'che' ConfigMap from other CR fields,the value defined in the 'extraProperties' is used instead.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"log_level": schema.StringAttribute{
										Description:         "The log level for the Che server: 'INFO' or 'DEBUG'.",
										MarkdownDescription: "The log level for the Che server: 'INFO' or 'DEBUG'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy server settings for Kubernetes cluster. No additional configuration is required for OpenShift cluster.By specifying these settings for the OpenShift cluster, you override the OpenShift proxy configuration.",
										MarkdownDescription: "Proxy server settings for Kubernetes cluster. No additional configuration is required for OpenShift cluster.By specifying these settings for the OpenShift cluster, you override the OpenShift proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"credentials_secret_name": schema.StringAttribute{
												Description:         "The secret name that contains 'user' and 'password' for a proxy server.The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
												MarkdownDescription: "The secret name that contains 'user' and 'password' for a proxy server.The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"non_proxy_hosts": schema.ListAttribute{
												Description:         "A list of hosts that can be reached directly, bypassing the proxy.Specify wild card domain use the following form '.<DOMAIN>', for example:   - localhost   - my.host.com   - 123.42.12.32Use only when a proxy configuration is required. The Operator respects OpenShift cluster-wide proxy configuration,defining 'nonProxyHosts' in a custom resource leads to merging non-proxy hosts lists from the cluster proxy configuration, and the ones defined in the custom resources.See the following page: https://docs.openshift.com/container-platform/latest/networking/enable-cluster-wide-proxy.html.",
												MarkdownDescription: "A list of hosts that can be reached directly, bypassing the proxy.Specify wild card domain use the following form '.<DOMAIN>', for example:   - localhost   - my.host.com   - 123.42.12.32Use only when a proxy configuration is required. The Operator respects OpenShift cluster-wide proxy configuration,defining 'nonProxyHosts' in a custom resource leads to merging non-proxy hosts lists from the cluster proxy configuration, and the ones defined in the custom resources.See the following page: https://docs.openshift.com/container-platform/latest/networking/enable-cluster-wide-proxy.html.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.StringAttribute{
												Description:         "Proxy server port.",
												MarkdownDescription: "Proxy server port.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url": schema.StringAttribute{
												Description:         "URL (protocol+hostname) of the proxy server.Use only when a proxy configuration is required. The Operator respects OpenShift cluster-wide proxy configuration,defining 'url' in a custom resource leads to overriding the cluster proxy configuration.See the following page: https://docs.openshift.com/container-platform/latest/networking/enable-cluster-wide-proxy.html.",
												MarkdownDescription: "URL (protocol+hostname) of the proxy server.Use only when a proxy configuration is required. The Operator respects OpenShift cluster-wide proxy configuration,defining 'url' in a custom resource leads to overriding the cluster proxy configuration.See the following page: https://docs.openshift.com/container-platform/latest/networking/enable-cluster-wide-proxy.html.",
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

							"dashboard": schema.SingleNestedAttribute{
								Description:         "Configuration settings related to the dashboard used by the Che installation.",
								MarkdownDescription: "Configuration settings related to the dashboard used by the Che installation.",
								Attributes: map[string]schema.Attribute{
									"branding": schema.SingleNestedAttribute{
										Description:         "Dashboard branding resources.",
										MarkdownDescription: "Dashboard branding resources.",
										Attributes: map[string]schema.Attribute{
											"logo": schema.SingleNestedAttribute{
												Description:         "Dashboard logo.",
												MarkdownDescription: "Dashboard logo.",
												Attributes: map[string]schema.Attribute{
													"base64data": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"mediatype": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

									"deployment": schema.SingleNestedAttribute{
										Description:         "Deployment override options.",
										MarkdownDescription: "Deployment override options.",
										Attributes: map[string]schema.Attribute{
											"containers": schema.ListNestedAttribute{
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"env": schema.ListNestedAttribute{
															Description:         "List of environment variables to set in the container.",
															MarkdownDescription: "List of environment variables to set in the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																		MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																		MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
																		Attributes: map[string]schema.Attribute{
																			"config_map_key_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a key of a ConfigMap.",
																				MarkdownDescription: "Selects a key of a ConfigMap.",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key to select.",
																						MarkdownDescription: "The key to select.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				Attributes: map[string]schema.Attribute{
																					"api_version": schema.StringAttribute{
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"field_path": schema.StringAttribute{
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource_field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				Attributes: map[string]schema.Attribute{
																					"container_name": schema.StringAttribute{
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"divisor": schema.StringAttribute{
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"resource": schema.StringAttribute{
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key of the secret to select from.  Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"image": schema.StringAttribute{
															Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
															MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_policy": schema.StringAttribute{
															Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
															MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
															},
														},

														"name": schema.StringAttribute{
															Description:         "Container name.",
															MarkdownDescription: "Container name.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Compute resources required by this container.",
															MarkdownDescription: "Compute resources required by this container.",
															Attributes: map[string]schema.Attribute{
																"limits": schema.SingleNestedAttribute{
																	Description:         "Describes the maximum amount of compute resources allowed.",
																	MarkdownDescription: "Describes the maximum amount of compute resources allowed.",
																	Attributes: map[string]schema.Attribute{
																		"cpu": schema.StringAttribute{
																			Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"memory": schema.StringAttribute{
																			Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"request": schema.SingleNestedAttribute{
																	Description:         "Describes the minimum amount of compute resources required.",
																	MarkdownDescription: "Describes the minimum amount of compute resources required.",
																	Attributes: map[string]schema.Attribute{
																		"cpu": schema.StringAttribute{
																			Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"memory": schema.StringAttribute{
																			Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": schema.SingleNestedAttribute{
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",
												Attributes: map[string]schema.Attribute{
													"fs_group": schema.Int64Attribute{
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user": schema.Int64Attribute{
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",
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

									"header_message": schema.SingleNestedAttribute{
										Description:         "Dashboard header message.",
										MarkdownDescription: "Dashboard header message.",
										Attributes: map[string]schema.Attribute{
											"show": schema.BoolAttribute{
												Description:         "Instructs dashboard to show the message.",
												MarkdownDescription: "Instructs dashboard to show the message.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"text": schema.StringAttribute{
												Description:         "Warning message displayed on the user dashboard.",
												MarkdownDescription: "Warning message displayed on the user dashboard.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_level": schema.StringAttribute{
										Description:         "The log level for the Dashboard.",
										MarkdownDescription: "The log level for the Dashboard.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("DEBUG", "INFO", "WARN", "ERROR", "FATAL", "TRACE", "SILENT"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"dev_workspace": schema.SingleNestedAttribute{
								Description:         "DevWorkspace Operator configuration.",
								MarkdownDescription: "DevWorkspace Operator configuration.",
								Attributes: map[string]schema.Attribute{
									"running_limit": schema.StringAttribute{
										Description:         "Deprecated in favor of 'MaxNumberOfRunningWorkspacesPerUser'The maximum number of running workspaces per user.",
										MarkdownDescription: "Deprecated in favor of 'MaxNumberOfRunningWorkspacesPerUser'The maximum number of running workspaces per user.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"devfile_registry": schema.SingleNestedAttribute{
								Description:         "Configuration settings related to the devfile registry used by the Che installation.",
								MarkdownDescription: "Configuration settings related to the devfile registry used by the Che installation.",
								Attributes: map[string]schema.Attribute{
									"deployment": schema.SingleNestedAttribute{
										Description:         "Deprecated deployment override options.",
										MarkdownDescription: "Deprecated deployment override options.",
										Attributes: map[string]schema.Attribute{
											"containers": schema.ListNestedAttribute{
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"env": schema.ListNestedAttribute{
															Description:         "List of environment variables to set in the container.",
															MarkdownDescription: "List of environment variables to set in the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																		MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																		MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
																		Attributes: map[string]schema.Attribute{
																			"config_map_key_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a key of a ConfigMap.",
																				MarkdownDescription: "Selects a key of a ConfigMap.",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key to select.",
																						MarkdownDescription: "The key to select.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				Attributes: map[string]schema.Attribute{
																					"api_version": schema.StringAttribute{
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"field_path": schema.StringAttribute{
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource_field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				Attributes: map[string]schema.Attribute{
																					"container_name": schema.StringAttribute{
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"divisor": schema.StringAttribute{
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"resource": schema.StringAttribute{
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key of the secret to select from.  Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"image": schema.StringAttribute{
															Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
															MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_policy": schema.StringAttribute{
															Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
															MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
															},
														},

														"name": schema.StringAttribute{
															Description:         "Container name.",
															MarkdownDescription: "Container name.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Compute resources required by this container.",
															MarkdownDescription: "Compute resources required by this container.",
															Attributes: map[string]schema.Attribute{
																"limits": schema.SingleNestedAttribute{
																	Description:         "Describes the maximum amount of compute resources allowed.",
																	MarkdownDescription: "Describes the maximum amount of compute resources allowed.",
																	Attributes: map[string]schema.Attribute{
																		"cpu": schema.StringAttribute{
																			Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"memory": schema.StringAttribute{
																			Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"request": schema.SingleNestedAttribute{
																	Description:         "Describes the minimum amount of compute resources required.",
																	MarkdownDescription: "Describes the minimum amount of compute resources required.",
																	Attributes: map[string]schema.Attribute{
																		"cpu": schema.StringAttribute{
																			Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"memory": schema.StringAttribute{
																			Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": schema.SingleNestedAttribute{
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",
												Attributes: map[string]schema.Attribute{
													"fs_group": schema.Int64Attribute{
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user": schema.Int64Attribute{
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",
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

									"disable_internal_registry": schema.BoolAttribute{
										Description:         "Disables internal devfile registry.",
										MarkdownDescription: "Disables internal devfile registry.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_devfile_registries": schema.ListNestedAttribute{
										Description:         "External devfile registries serving sample ready-to-use devfiles.",
										MarkdownDescription: "External devfile registries serving sample ready-to-use devfiles.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"url": schema.StringAttribute{
													Description:         "The public UR of the devfile registry that serves sample ready-to-use devfiles.",
													MarkdownDescription: "The public UR of the devfile registry that serves sample ready-to-use devfiles.",
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

							"image_puller": schema.SingleNestedAttribute{
								Description:         "Kubernetes Image Puller configuration.",
								MarkdownDescription: "Kubernetes Image Puller configuration.",
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										Description:         "Install and configure the community supported Kubernetes Image Puller Operator. When you set the value to 'true' without providing any specs,it creates a default Kubernetes Image Puller object managed by the Operator.When you set the value to 'false', the Kubernetes Image Puller object is deleted, and the Operator uninstalled,regardless of whether a spec is provided.If you leave the 'spec.images' field empty, a set of recommended workspace-related images is automatically detected andpre-pulled after installation.Note that while this Operator and its behavior is community-supported, its payload may be commercially-supportedfor pulling commercially-supported images.",
										MarkdownDescription: "Install and configure the community supported Kubernetes Image Puller Operator. When you set the value to 'true' without providing any specs,it creates a default Kubernetes Image Puller object managed by the Operator.When you set the value to 'false', the Kubernetes Image Puller object is deleted, and the Operator uninstalled,regardless of whether a spec is provided.If you leave the 'spec.images' field empty, a set of recommended workspace-related images is automatically detected andpre-pulled after installation.Note that while this Operator and its behavior is community-supported, its payload may be commercially-supportedfor pulling commercially-supported images.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "A Kubernetes Image Puller spec to configure the image puller in the CheCluster.",
										MarkdownDescription: "A Kubernetes Image Puller spec to configure the image puller in the CheCluster.",
										Attributes: map[string]schema.Attribute{
											"affinity": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"caching_cpu_limit": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"caching_cpu_request": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"caching_interval_hours": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"caching_memory_limit": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"caching_memory_request": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"config_map_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"daemonset_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"deployment_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_pull_secrets": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_puller_image": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"images": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_selector": schema.StringAttribute{
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

							"metrics": schema.SingleNestedAttribute{
								Description:         "Che server metrics configuration.",
								MarkdownDescription: "Che server metrics configuration.",
								Attributes: map[string]schema.Attribute{
									"enable": schema.BoolAttribute{
										Description:         "Enables 'metrics' for the Che server endpoint.",
										MarkdownDescription: "Enables 'metrics' for the Che server endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"plugin_registry": schema.SingleNestedAttribute{
								Description:         "Configuration settings related to the plug-in registry used by the Che installation.",
								MarkdownDescription: "Configuration settings related to the plug-in registry used by the Che installation.",
								Attributes: map[string]schema.Attribute{
									"deployment": schema.SingleNestedAttribute{
										Description:         "Deployment override options.",
										MarkdownDescription: "Deployment override options.",
										Attributes: map[string]schema.Attribute{
											"containers": schema.ListNestedAttribute{
												Description:         "List of containers belonging to the pod.",
												MarkdownDescription: "List of containers belonging to the pod.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"env": schema.ListNestedAttribute{
															Description:         "List of environment variables to set in the container.",
															MarkdownDescription: "List of environment variables to set in the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																		MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value_from": schema.SingleNestedAttribute{
																		Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																		MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
																		Attributes: map[string]schema.Attribute{
																			"config_map_key_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a key of a ConfigMap.",
																				MarkdownDescription: "Selects a key of a ConfigMap.",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key to select.",
																						MarkdownDescription: "The key to select.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				Attributes: map[string]schema.Attribute{
																					"api_version": schema.StringAttribute{
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"field_path": schema.StringAttribute{
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"resource_field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				Attributes: map[string]schema.Attribute{
																					"container_name": schema.StringAttribute{
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"divisor": schema.StringAttribute{
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"resource": schema.StringAttribute{
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_key_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key of the secret to select from.  Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"image": schema.StringAttribute{
															Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
															MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_policy": schema.StringAttribute{
															Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
															MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
															},
														},

														"name": schema.StringAttribute{
															Description:         "Container name.",
															MarkdownDescription: "Container name.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Compute resources required by this container.",
															MarkdownDescription: "Compute resources required by this container.",
															Attributes: map[string]schema.Attribute{
																"limits": schema.SingleNestedAttribute{
																	Description:         "Describes the maximum amount of compute resources allowed.",
																	MarkdownDescription: "Describes the maximum amount of compute resources allowed.",
																	Attributes: map[string]schema.Attribute{
																		"cpu": schema.StringAttribute{
																			Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"memory": schema.StringAttribute{
																			Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"request": schema.SingleNestedAttribute{
																	Description:         "Describes the minimum amount of compute resources required.",
																	MarkdownDescription: "Describes the minimum amount of compute resources required.",
																	Attributes: map[string]schema.Attribute{
																		"cpu": schema.StringAttribute{
																			Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"memory": schema.StringAttribute{
																			Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																			MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": schema.SingleNestedAttribute{
												Description:         "Security options the pod should run with.",
												MarkdownDescription: "Security options the pod should run with.",
												Attributes: map[string]schema.Attribute{
													"fs_group": schema.Int64Attribute{
														Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user": schema.Int64Attribute{
														Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",
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

									"disable_internal_registry": schema.BoolAttribute{
										Description:         "Disables internal plug-in registry.",
										MarkdownDescription: "Disables internal plug-in registry.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_plugin_registries": schema.ListNestedAttribute{
										Description:         "External plugin registries.",
										MarkdownDescription: "External plugin registries.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"url": schema.StringAttribute{
													Description:         "Public URL of the plug-in registry.",
													MarkdownDescription: "Public URL of the plug-in registry.",
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

									"open_vsx_url": schema.StringAttribute{
										Description:         "Open VSX registry URL. If omitted an embedded instance will be used.",
										MarkdownDescription: "Open VSX registry URL. If omitted an embedded instance will be used.",
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

					"container_registry": schema.SingleNestedAttribute{
						Description:         "Configuration of an alternative registry that stores Che images.",
						MarkdownDescription: "Configuration of an alternative registry that stores Che images.",
						Attributes: map[string]schema.Attribute{
							"hostname": schema.StringAttribute{
								Description:         "An optional hostname or URL of an alternative container registry to pull images from.This value overrides the container registry hostname defined in all the default container images involved in a Che deployment.This is particularly useful for installing Che in a restricted environment.",
								MarkdownDescription: "An optional hostname or URL of an alternative container registry to pull images from.This value overrides the container registry hostname defined in all the default container images involved in a Che deployment.This is particularly useful for installing Che in a restricted environment.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"organization": schema.StringAttribute{
								Description:         "An optional repository name of an alternative registry to pull images from.This value overrides the container registry organization defined in all the default container images involved in a Che deployment.This is particularly useful for installing Eclipse Che in a restricted environment.",
								MarkdownDescription: "An optional repository name of an alternative registry to pull images from.This value overrides the container registry organization defined in all the default container images involved in a Che deployment.This is particularly useful for installing Eclipse Che in a restricted environment.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"dev_environments": schema.SingleNestedAttribute{
						Description:         "Development environment default configuration options.",
						MarkdownDescription: "Development environment default configuration options.",
						Attributes: map[string]schema.Attribute{
							"container_build_configuration": schema.SingleNestedAttribute{
								Description:         "Container build configuration.",
								MarkdownDescription: "Container build configuration.",
								Attributes: map[string]schema.Attribute{
									"open_shift_security_context_constraint": schema.StringAttribute{
										Description:         "OpenShift security context constraint to build containers.",
										MarkdownDescription: "OpenShift security context constraint to build containers.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"default_components": schema.ListNestedAttribute{
								Description:         "Default components applied to DevWorkspaces.These default components are meant to be used when a Devfile, that does not contain any components.",
								MarkdownDescription: "Default components applied to DevWorkspaces.These default components are meant to be used when a Devfile, that does not contain any components.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"attributes": schema.MapAttribute{
											Description:         "Map of implementation-dependant free-form YAML attributes.",
											MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"component_type": schema.StringAttribute{
											Description:         "Type of component",
											MarkdownDescription: "Type of component",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image", "Plugin", "Custom"),
											},
										},

										"container": schema.SingleNestedAttribute{
											Description:         "Allows adding and configuring devworkspace-related containers",
											MarkdownDescription: "Allows adding and configuring devworkspace-related containers",
											Attributes: map[string]schema.Attribute{
												"annotation": schema.SingleNestedAttribute{
													Description:         "Annotations that should be added to specific resources for this container",
													MarkdownDescription: "Annotations that should be added to specific resources for this container",
													Attributes: map[string]schema.Attribute{
														"deployment": schema.MapAttribute{
															Description:         "Annotations to be added to deployment",
															MarkdownDescription: "Annotations to be added to deployment",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service": schema.MapAttribute{
															Description:         "Annotations to be added to service",
															MarkdownDescription: "Annotations to be added to service",
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

												"args": schema.ListAttribute{
													Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.Defaults to an empty array, meaning use whatever is defined in the image.",
													MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.Defaults to an empty array, meaning use whatever is defined in the image.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"command": schema.ListAttribute{
													Description:         "The command to run in the dockerimage component instead of the default one provided in the image.Defaults to an empty array, meaning use whatever is defined in the image.",
													MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image.Defaults to an empty array, meaning use whatever is defined in the image.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cpu_limit": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"cpu_request": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dedicated_pod": schema.BoolAttribute{
													Description:         "Specify if a container should run in its own separated pod,instead of running as part of the main development environment pod.Default value is 'false'",
													MarkdownDescription: "Specify if a container should run in its own separated pod,instead of running as part of the main development environment pod.Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"endpoints": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"annotation": schema.MapAttribute{
																Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("public", "internal", "none"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path of the endpoint URL",
																MarkdownDescription: "Path of the endpoint URL",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"protocol": schema.StringAttribute{
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannotbe used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannotbe used by two different container components.",
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

												"env": schema.ListNestedAttribute{
													Description:         "Environment variables used in this container.The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
													MarkdownDescription: "Environment variables used in this container.The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

												"image": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"memory_limit": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"memory_request": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mount_sources": schema.BoolAttribute{
													Description:         "Toggles whether or not the project source code shouldbe mounted in the component.Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
													MarkdownDescription: "Toggles whether or not the project source code shouldbe mounted in the component.Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_mapping": schema.StringAttribute{
													Description:         "Optional specification of the path in the container whereproject sources should be transferred/mounted when 'mountSources' is 'true'.When omitted, the default value of /projects is used.",
													MarkdownDescription: "Optional specification of the path in the container whereproject sources should be transferred/mounted when 'mountSources' is 'true'.When omitted, the default value of /projects is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_mounts": schema.ListNestedAttribute{
													Description:         "List of volumes mounts that should be mounted is this container.",
													MarkdownDescription: "List of volumes mounts that should be mounted is this container.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "The volume mount name is the name of an existing 'Volume' component.If several containers mount the same volume namethen they will reuse the same volume and will be able to access to the same files.",
																MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component.If several containers mount the same volume namethen they will reuse the same volume and will be able to access to the same files.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "The path in the component container where the volume should be mounted.If not path is mentioned, default path is the is '/<name>'.",
																MarkdownDescription: "The path in the component container where the volume should be mounted.If not path is mentioned, default path is the is '/<name>'.",
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

										"custom": schema.SingleNestedAttribute{
											Description:         "Custom component whose logic is implementation-dependantand should be provided by the userpossibly through some dedicated controller",
											MarkdownDescription: "Custom component whose logic is implementation-dependantand should be provided by the userpossibly through some dedicated controller",
											Attributes: map[string]schema.Attribute{
												"component_class": schema.StringAttribute{
													Description:         "Class of component that the associated implementation controllershould use to process this command with the appropriate logic",
													MarkdownDescription: "Class of component that the associated implementation controllershould use to process this command with the appropriate logic",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"embedded_resource": schema.MapAttribute{
													Description:         "Additional free-form configuration for this custom componentthat the implementation controller will know how to use",
													MarkdownDescription: "Additional free-form configuration for this custom componentthat the implementation controller will know how to use",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"image": schema.SingleNestedAttribute{
											Description:         "Allows specifying the definition of an image for outer loop builds",
											MarkdownDescription: "Allows specifying the definition of an image for outer loop builds",
											Attributes: map[string]schema.Attribute{
												"auto_build": schema.BoolAttribute{
													Description:         "Defines if the image should be built during startup.Default value is 'false'",
													MarkdownDescription: "Defines if the image should be built during startup.Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dockerfile": schema.SingleNestedAttribute{
													Description:         "Allows specifying dockerfile type build",
													MarkdownDescription: "Allows specifying dockerfile type build",
													Attributes: map[string]schema.Attribute{
														"args": schema.ListAttribute{
															Description:         "The arguments to supply to the dockerfile build.",
															MarkdownDescription: "The arguments to supply to the dockerfile build.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"build_context": schema.StringAttribute{
															Description:         "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
															MarkdownDescription: "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"devfile_registry": schema.SingleNestedAttribute{
															Description:         "Dockerfile's Devfile Registry source",
															MarkdownDescription: "Dockerfile's Devfile Registry source",
															Attributes: map[string]schema.Attribute{
																"id": schema.StringAttribute{
																	Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registryrequired for the Dockerfile build will be downloaded for building the image.",
																	MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registryrequired for the Dockerfile build will be downloaded for building the image.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"registry_url": schema.StringAttribute{
																	Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src.To ensure the Dockerfile gets resolved consistently in different environments,it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																	MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src.To ensure the Dockerfile gets resolved consistently in different environments,it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"git": schema.SingleNestedAttribute{
															Description:         "Dockerfile's Git source",
															MarkdownDescription: "Dockerfile's Git source",
															Attributes: map[string]schema.Attribute{
																"checkout_from": schema.SingleNestedAttribute{
																	Description:         "Defines from what the project should be checked out. Required if there are more than one remote configured",
																	MarkdownDescription: "Defines from what the project should be checked out. Required if there are more than one remote configured",
																	Attributes: map[string]schema.Attribute{
																		"remote": schema.StringAttribute{
																			Description:         "The remote name should be used as init. Required if there are more than one remote configured",
																			MarkdownDescription: "The remote name should be used as init. Required if there are more than one remote configured",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"revision": schema.StringAttribute{
																			Description:         "The revision to checkout from. Should be branch name, tag or commit id.Default branch is used if missing or specified revision is not found.",
																			MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id.Default branch is used if missing or specified revision is not found.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"file_location": schema.StringAttribute{
																	Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src.Defaults to Dockerfile.",
																	MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src.Defaults to Dockerfile.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"remotes": schema.MapAttribute{
																	Description:         "The remotes map which should be initialized in the git project.Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																	MarkdownDescription: "The remotes map which should be initialized in the git project.Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"root_required": schema.BoolAttribute{
															Description:         "Specify if a privileged builder pod is required.Default value is 'false'",
															MarkdownDescription: "Specify if a privileged builder pod is required.Default value is 'false'",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"src_type": schema.StringAttribute{
															Description:         "Type of Dockerfile src",
															MarkdownDescription: "Type of Dockerfile src",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Uri", "DevfileRegistry", "Git"),
															},
														},

														"uri": schema.StringAttribute{
															Description:         "URI Reference of a Dockerfile.It can be a full URL or a relative URI from the current devfile as the base URI.",
															MarkdownDescription: "URI Reference of a Dockerfile.It can be a full URL or a relative URI from the current devfile as the base URI.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"image_name": schema.StringAttribute{
													Description:         "Name of the image for the resulting outerloop build",
													MarkdownDescription: "Name of the image for the resulting outerloop build",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"image_type": schema.StringAttribute{
													Description:         "Type of image",
													MarkdownDescription: "Type of image",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Dockerfile"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kubernetes": schema.SingleNestedAttribute{
											Description:         "Allows importing into the devworkspace the Kubernetes resourcesdefined in a given manifest. For example this allows reusing the Kubernetesdefinitions used to deploy some runtime components in production.",
											MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resourcesdefined in a given manifest. For example this allows reusing the Kubernetesdefinitions used to deploy some runtime components in production.",
											Attributes: map[string]schema.Attribute{
												"deploy_by_default": schema.BoolAttribute{
													Description:         "Defines if the component should be deployed during startup.Default value is 'false'",
													MarkdownDescription: "Defines if the component should be deployed during startup.Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"endpoints": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"annotation": schema.MapAttribute{
																Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("public", "internal", "none"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path of the endpoint URL",
																MarkdownDescription: "Path of the endpoint URL",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"protocol": schema.StringAttribute{
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannotbe used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannotbe used by two different container components.",
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

												"inlined": schema.StringAttribute{
													Description:         "Inlined manifest",
													MarkdownDescription: "Inlined manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"location_type": schema.StringAttribute{
													Description:         "Type of Kubernetes-like location",
													MarkdownDescription: "Type of Kubernetes-like location",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Uri", "Inlined"),
													},
												},

												"uri": schema.StringAttribute{
													Description:         "Location in a file fetched from a uri.",
													MarkdownDescription: "Location in a file fetched from a uri.",
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
											Description:         "Mandatory name that allows referencing the componentfrom other elements (such as commands) or from an externaldevfile that may reference this component through a parent or a plugin.",
											MarkdownDescription: "Mandatory name that allows referencing the componentfrom other elements (such as commands) or from an externaldevfile that may reference this component through a parent or a plugin.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
											},
										},

										"openshift": schema.SingleNestedAttribute{
											Description:         "Allows importing into the devworkspace the OpenShift resourcesdefined in a given manifest. For example this allows reusing the OpenShiftdefinitions used to deploy some runtime components in production.",
											MarkdownDescription: "Allows importing into the devworkspace the OpenShift resourcesdefined in a given manifest. For example this allows reusing the OpenShiftdefinitions used to deploy some runtime components in production.",
											Attributes: map[string]schema.Attribute{
												"deploy_by_default": schema.BoolAttribute{
													Description:         "Defines if the component should be deployed during startup.Default value is 'false'",
													MarkdownDescription: "Defines if the component should be deployed during startup.Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"endpoints": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"annotation": schema.MapAttribute{
																Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("public", "internal", "none"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "Path of the endpoint URL",
																MarkdownDescription: "Path of the endpoint URL",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"protocol": schema.StringAttribute{
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannotbe used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannotbe used by two different container components.",
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

												"inlined": schema.StringAttribute{
													Description:         "Inlined manifest",
													MarkdownDescription: "Inlined manifest",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"location_type": schema.StringAttribute{
													Description:         "Type of Kubernetes-like location",
													MarkdownDescription: "Type of Kubernetes-like location",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Uri", "Inlined"),
													},
												},

												"uri": schema.StringAttribute{
													Description:         "Location in a file fetched from a uri.",
													MarkdownDescription: "Location in a file fetched from a uri.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"plugin": schema.SingleNestedAttribute{
											Description:         "Allows importing a plugin.Plugins are mainly imported devfiles that contribute components, commandsand events as a consistent single unit. They are defined in either YAML filesfollowing the devfile syntax,or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
											MarkdownDescription: "Allows importing a plugin.Plugins are mainly imported devfiles that contribute components, commandsand events as a consistent single unit. They are defined in either YAML filesfollowing the devfile syntax,or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
											Attributes: map[string]schema.Attribute{
												"commands": schema.ListNestedAttribute{
													Description:         "Overrides of commands encapsulated in a parent devfile or a plugin.Overriding is done according to K8S strategic merge patch standard rules.",
													MarkdownDescription: "Overrides of commands encapsulated in a parent devfile or a plugin.Overriding is done according to K8S strategic merge patch standard rules.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"apply": schema.SingleNestedAttribute{
																Description:         "Command that consists in applying a given component definition,typically bound to a devworkspace event.For example, when an 'apply' command is bound to a 'preStart' event,and references a 'container' component, it will start the container as aK8S initContainer in the devworkspace POD, unless the component has its'dedicatedPod' field set to 'true'.When no 'apply' command exist for a given component,it is assumed the component will be applied at devworkspace startby default, unless 'deployByDefault' for that component is set to false.",
																MarkdownDescription: "Command that consists in applying a given component definition,typically bound to a devworkspace event.For example, when an 'apply' command is bound to a 'preStart' event,and references a 'container' component, it will start the container as aK8S initContainer in the devworkspace POD, unless the component has its'dedicatedPod' field set to 'true'.When no 'apply' command exist for a given component,it is assumed the component will be applied at devworkspace startby default, unless 'deployByDefault' for that component is set to false.",
																Attributes: map[string]schema.Attribute{
																	"component": schema.StringAttribute{
																		Description:         "Describes component that will be applied",
																		MarkdownDescription: "Describes component that will be applied",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"group": schema.SingleNestedAttribute{
																		Description:         "Defines the group this command is part of",
																		MarkdownDescription: "Defines the group this command is part of",
																		Attributes: map[string]schema.Attribute{
																			"is_default": schema.BoolAttribute{
																				Description:         "Identifies the default command for a given group kind",
																				MarkdownDescription: "Identifies the default command for a given group kind",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"kind": schema.StringAttribute{
																				Description:         "Kind of group the command is part of",
																				MarkdownDescription: "Kind of group the command is part of",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"label": schema.StringAttribute{
																		Description:         "Optional label that provides a label for this commandto be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this commandto be used in Editor UI menus for example",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant free-form YAML attributes.",
																MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"command_type": schema.StringAttribute{
																Description:         "Type of devworkspace command",
																MarkdownDescription: "Type of devworkspace command",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Exec", "Apply", "Composite"),
																},
															},

															"composite": schema.SingleNestedAttribute{
																Description:         "Composite command that allows executing several sub-commandseither sequentially or concurrently",
																MarkdownDescription: "Composite command that allows executing several sub-commandseither sequentially or concurrently",
																Attributes: map[string]schema.Attribute{
																	"commands": schema.ListAttribute{
																		Description:         "The commands that comprise this composite command",
																		MarkdownDescription: "The commands that comprise this composite command",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"group": schema.SingleNestedAttribute{
																		Description:         "Defines the group this command is part of",
																		MarkdownDescription: "Defines the group this command is part of",
																		Attributes: map[string]schema.Attribute{
																			"is_default": schema.BoolAttribute{
																				Description:         "Identifies the default command for a given group kind",
																				MarkdownDescription: "Identifies the default command for a given group kind",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"kind": schema.StringAttribute{
																				Description:         "Kind of group the command is part of",
																				MarkdownDescription: "Kind of group the command is part of",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"label": schema.StringAttribute{
																		Description:         "Optional label that provides a label for this commandto be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this commandto be used in Editor UI menus for example",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"parallel": schema.BoolAttribute{
																		Description:         "Indicates if the sub-commands should be executed concurrently",
																		MarkdownDescription: "Indicates if the sub-commands should be executed concurrently",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"exec": schema.SingleNestedAttribute{
																Description:         "CLI Command executed in an existing component container",
																MarkdownDescription: "CLI Command executed in an existing component container",
																Attributes: map[string]schema.Attribute{
																	"command_line": schema.StringAttribute{
																		Description:         "The actual command-line stringSpecial variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		MarkdownDescription: "The actual command-line stringSpecial variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"component": schema.StringAttribute{
																		Description:         "Describes component to which given action relates",
																		MarkdownDescription: "Describes component to which given action relates",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"env": schema.ListNestedAttribute{
																		Description:         "Optional list of environment variables that have to be setbefore running the command",
																		MarkdownDescription: "Optional list of environment variables that have to be setbefore running the command",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
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

																	"group": schema.SingleNestedAttribute{
																		Description:         "Defines the group this command is part of",
																		MarkdownDescription: "Defines the group this command is part of",
																		Attributes: map[string]schema.Attribute{
																			"is_default": schema.BoolAttribute{
																				Description:         "Identifies the default command for a given group kind",
																				MarkdownDescription: "Identifies the default command for a given group kind",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"kind": schema.StringAttribute{
																				Description:         "Kind of group the command is part of",
																				MarkdownDescription: "Kind of group the command is part of",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("build", "run", "test", "debug", "deploy"),
																				},
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"hot_reload_capable": schema.BoolAttribute{
																		Description:         "Specify whether the command is restarted or not when the source code changes.If set to 'true' the command won't be restarted.A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted.A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again.This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'.Default value is 'false'",
																		MarkdownDescription: "Specify whether the command is restarted or not when the source code changes.If set to 'true' the command won't be restarted.A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted.A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again.This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'.Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"label": schema.StringAttribute{
																		Description:         "Optional label that provides a label for this commandto be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this commandto be used in Editor UI menus for example",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"working_dir": schema.StringAttribute{
																		Description:         "Working directory where the command should be executedSpecial variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		MarkdownDescription: "Working directory where the command should be executedSpecial variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"id": schema.StringAttribute{
																Description:         "Mandatory identifier that allows referencingthis command in composite commands, froma parent, or in events.",
																MarkdownDescription: "Mandatory identifier that allows referencingthis command in composite commands, froma parent, or in events.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"components": schema.ListNestedAttribute{
													Description:         "Overrides of components encapsulated in a parent devfile or a plugin.Overriding is done according to K8S strategic merge patch standard rules.",
													MarkdownDescription: "Overrides of components encapsulated in a parent devfile or a plugin.Overriding is done according to K8S strategic merge patch standard rules.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"attributes": schema.MapAttribute{
																Description:         "Map of implementation-dependant free-form YAML attributes.",
																MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"component_type": schema.StringAttribute{
																Description:         "Type of component",
																MarkdownDescription: "Type of component",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image"),
																},
															},

															"container": schema.SingleNestedAttribute{
																Description:         "Allows adding and configuring devworkspace-related containers",
																MarkdownDescription: "Allows adding and configuring devworkspace-related containers",
																Attributes: map[string]schema.Attribute{
																	"annotation": schema.SingleNestedAttribute{
																		Description:         "Annotations that should be added to specific resources for this container",
																		MarkdownDescription: "Annotations that should be added to specific resources for this container",
																		Attributes: map[string]schema.Attribute{
																			"deployment": schema.MapAttribute{
																				Description:         "Annotations to be added to deployment",
																				MarkdownDescription: "Annotations to be added to deployment",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"service": schema.MapAttribute{
																				Description:         "Annotations to be added to service",
																				MarkdownDescription: "Annotations to be added to service",
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

																	"args": schema.ListAttribute{
																		Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.Defaults to an empty array, meaning use whatever is defined in the image.",
																		MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command.Defaults to an empty array, meaning use whatever is defined in the image.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"command": schema.ListAttribute{
																		Description:         "The command to run in the dockerimage component instead of the default one provided in the image.Defaults to an empty array, meaning use whatever is defined in the image.",
																		MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image.Defaults to an empty array, meaning use whatever is defined in the image.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"cpu_limit": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"cpu_request": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"dedicated_pod": schema.BoolAttribute{
																		Description:         "Specify if a container should run in its own separated pod,instead of running as part of the main development environment pod.Default value is 'false'",
																		MarkdownDescription: "Specify if a container should run in its own separated pod,instead of running as part of the main development environment pod.Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"endpoints": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"annotation": schema.MapAttribute{
																					Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"attributes": schema.MapAttribute{
																					Description:         "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("public", "internal", "none"),
																					},
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "Path of the endpoint URL",
																					MarkdownDescription: "Path of the endpoint URL",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"protocol": schema.StringAttribute{
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannotbe used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannotbe used by two different container components.",
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

																	"env": schema.ListNestedAttribute{
																		Description:         "Environment variables used in this container.The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
																		MarkdownDescription: "Environment variables used in this container.The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
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

																	"image": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"memory_limit": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"memory_request": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"mount_sources": schema.BoolAttribute{
																		Description:         "Toggles whether or not the project source code shouldbe mounted in the component.Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																		MarkdownDescription: "Toggles whether or not the project source code shouldbe mounted in the component.Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"source_mapping": schema.StringAttribute{
																		Description:         "Optional specification of the path in the container whereproject sources should be transferred/mounted when 'mountSources' is 'true'.When omitted, the default value of /projects is used.",
																		MarkdownDescription: "Optional specification of the path in the container whereproject sources should be transferred/mounted when 'mountSources' is 'true'.When omitted, the default value of /projects is used.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"volume_mounts": schema.ListNestedAttribute{
																		Description:         "List of volumes mounts that should be mounted is this container.",
																		MarkdownDescription: "List of volumes mounts that should be mounted is this container.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "The volume mount name is the name of an existing 'Volume' component.If several containers mount the same volume namethen they will reuse the same volume and will be able to access to the same files.",
																					MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component.If several containers mount the same volume namethen they will reuse the same volume and will be able to access to the same files.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "The path in the component container where the volume should be mounted.If not path is mentioned, default path is the is '/<name>'.",
																					MarkdownDescription: "The path in the component container where the volume should be mounted.If not path is mentioned, default path is the is '/<name>'.",
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

															"image": schema.SingleNestedAttribute{
																Description:         "Allows specifying the definition of an image for outer loop builds",
																MarkdownDescription: "Allows specifying the definition of an image for outer loop builds",
																Attributes: map[string]schema.Attribute{
																	"auto_build": schema.BoolAttribute{
																		Description:         "Defines if the image should be built during startup.Default value is 'false'",
																		MarkdownDescription: "Defines if the image should be built during startup.Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"dockerfile": schema.SingleNestedAttribute{
																		Description:         "Allows specifying dockerfile type build",
																		MarkdownDescription: "Allows specifying dockerfile type build",
																		Attributes: map[string]schema.Attribute{
																			"args": schema.ListAttribute{
																				Description:         "The arguments to supply to the dockerfile build.",
																				MarkdownDescription: "The arguments to supply to the dockerfile build.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"build_context": schema.StringAttribute{
																				Description:         "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
																				MarkdownDescription: "Path of source directory to establish build context. Defaults to ${PROJECT_SOURCE} in the container",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"devfile_registry": schema.SingleNestedAttribute{
																				Description:         "Dockerfile's Devfile Registry source",
																				MarkdownDescription: "Dockerfile's Devfile Registry source",
																				Attributes: map[string]schema.Attribute{
																					"id": schema.StringAttribute{
																						Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registryrequired for the Dockerfile build will be downloaded for building the image.",
																						MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registryrequired for the Dockerfile build will be downloaded for building the image.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"registry_url": schema.StringAttribute{
																						Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src.To ensure the Dockerfile gets resolved consistently in different environments,it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																						MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src.To ensure the Dockerfile gets resolved consistently in different environments,it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"git": schema.SingleNestedAttribute{
																				Description:         "Dockerfile's Git source",
																				MarkdownDescription: "Dockerfile's Git source",
																				Attributes: map[string]schema.Attribute{
																					"checkout_from": schema.SingleNestedAttribute{
																						Description:         "Defines from what the project should be checked out. Required if there are more than one remote configured",
																						MarkdownDescription: "Defines from what the project should be checked out. Required if there are more than one remote configured",
																						Attributes: map[string]schema.Attribute{
																							"remote": schema.StringAttribute{
																								Description:         "The remote name should be used as init. Required if there are more than one remote configured",
																								MarkdownDescription: "The remote name should be used as init. Required if there are more than one remote configured",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"revision": schema.StringAttribute{
																								Description:         "The revision to checkout from. Should be branch name, tag or commit id.Default branch is used if missing or specified revision is not found.",
																								MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id.Default branch is used if missing or specified revision is not found.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"file_location": schema.StringAttribute{
																						Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src.Defaults to Dockerfile.",
																						MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src.Defaults to Dockerfile.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"remotes": schema.MapAttribute{
																						Description:         "The remotes map which should be initialized in the git project.Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																						MarkdownDescription: "The remotes map which should be initialized in the git project.Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

																			"root_required": schema.BoolAttribute{
																				Description:         "Specify if a privileged builder pod is required.Default value is 'false'",
																				MarkdownDescription: "Specify if a privileged builder pod is required.Default value is 'false'",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"src_type": schema.StringAttribute{
																				Description:         "Type of Dockerfile src",
																				MarkdownDescription: "Type of Dockerfile src",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("Uri", "DevfileRegistry", "Git"),
																				},
																			},

																			"uri": schema.StringAttribute{
																				Description:         "URI Reference of a Dockerfile.It can be a full URL or a relative URI from the current devfile as the base URI.",
																				MarkdownDescription: "URI Reference of a Dockerfile.It can be a full URL or a relative URI from the current devfile as the base URI.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"image_name": schema.StringAttribute{
																		Description:         "Name of the image for the resulting outerloop build",
																		MarkdownDescription: "Name of the image for the resulting outerloop build",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"image_type": schema.StringAttribute{
																		Description:         "Type of image",
																		MarkdownDescription: "Type of image",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Dockerfile", "AutoBuild"),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"kubernetes": schema.SingleNestedAttribute{
																Description:         "Allows importing into the devworkspace the Kubernetes resourcesdefined in a given manifest. For example this allows reusing the Kubernetesdefinitions used to deploy some runtime components in production.",
																MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resourcesdefined in a given manifest. For example this allows reusing the Kubernetesdefinitions used to deploy some runtime components in production.",
																Attributes: map[string]schema.Attribute{
																	"deploy_by_default": schema.BoolAttribute{
																		Description:         "Defines if the component should be deployed during startup.Default value is 'false'",
																		MarkdownDescription: "Defines if the component should be deployed during startup.Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"endpoints": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"annotation": schema.MapAttribute{
																					Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"attributes": schema.MapAttribute{
																					Description:         "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("public", "internal", "none"),
																					},
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "Path of the endpoint URL",
																					MarkdownDescription: "Path of the endpoint URL",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"protocol": schema.StringAttribute{
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannotbe used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannotbe used by two different container components.",
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

																	"inlined": schema.StringAttribute{
																		Description:         "Inlined manifest",
																		MarkdownDescription: "Inlined manifest",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"location_type": schema.StringAttribute{
																		Description:         "Type of Kubernetes-like location",
																		MarkdownDescription: "Type of Kubernetes-like location",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Uri", "Inlined"),
																		},
																	},

																	"uri": schema.StringAttribute{
																		Description:         "Location in a file fetched from a uri.",
																		MarkdownDescription: "Location in a file fetched from a uri.",
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
																Description:         "Mandatory name that allows referencing the componentfrom other elements (such as commands) or from an externaldevfile that may reference this component through a parent or a plugin.",
																MarkdownDescription: "Mandatory name that allows referencing the componentfrom other elements (such as commands) or from an externaldevfile that may reference this component through a parent or a plugin.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"openshift": schema.SingleNestedAttribute{
																Description:         "Allows importing into the devworkspace the OpenShift resourcesdefined in a given manifest. For example this allows reusing the OpenShiftdefinitions used to deploy some runtime components in production.",
																MarkdownDescription: "Allows importing into the devworkspace the OpenShift resourcesdefined in a given manifest. For example this allows reusing the OpenShiftdefinitions used to deploy some runtime components in production.",
																Attributes: map[string]schema.Attribute{
																	"deploy_by_default": schema.BoolAttribute{
																		Description:         "Defines if the component should be deployed during startup.Default value is 'false'",
																		MarkdownDescription: "Defines if the component should be deployed during startup.Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"endpoints": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"annotation": schema.MapAttribute{
																					Description:         "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					MarkdownDescription: "Annotations to be added to Kubernetes Ingress or Openshift Route",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"attributes": schema.MapAttribute{
																					Description:         "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes.Examples of Che-specific attributes:- cookiesAuthEnabled: 'true' / 'false',- type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network.- 'public' means that the endpoint will be exposed on the public network, typically througha K8S ingress or an OpenShift route.- 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD,typically by K8S services, to be consumed by other elements runningon the same cloud internal network.- 'none' means that the endpoint will not be exposed and will only be accessibleinside the main devworkspace POD, on a local address.Default value is 'public'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("public", "internal", "none"),
																					},
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "Path of the endpoint URL",
																					MarkdownDescription: "Path of the endpoint URL",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"protocol": schema.StringAttribute{
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint.- 'http': Endpoint will have 'http' traffic, typically on a TCP connection.It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'.- 'https': Endpoint will have 'https' traffic, typically on a TCP connection.- 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection.It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'.- 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection.- 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol.- 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol.Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by someauthentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannotbe used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannotbe used by two different container components.",
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

																	"inlined": schema.StringAttribute{
																		Description:         "Inlined manifest",
																		MarkdownDescription: "Inlined manifest",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"location_type": schema.StringAttribute{
																		Description:         "Type of Kubernetes-like location",
																		MarkdownDescription: "Type of Kubernetes-like location",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Uri", "Inlined"),
																		},
																	},

																	"uri": schema.StringAttribute{
																		Description:         "Location in a file fetched from a uri.",
																		MarkdownDescription: "Location in a file fetched from a uri.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume": schema.SingleNestedAttribute{
																Description:         "Allows specifying the definition of a volumeshared by several other components",
																MarkdownDescription: "Allows specifying the definition of a volumeshared by several other components",
																Attributes: map[string]schema.Attribute{
																	"ephemeral": schema.BoolAttribute{
																		Description:         "Ephemeral volumes are not stored persistently across restarts. Defaultsto false",
																		MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaultsto false",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"size": schema.StringAttribute{
																		Description:         "Size of the volume",
																		MarkdownDescription: "Size of the volume",
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

												"id": schema.StringAttribute{
													Description:         "Id in a registry that contains a Devfile yaml file",
													MarkdownDescription: "Id in a registry that contains a Devfile yaml file",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"import_reference_type": schema.StringAttribute{
													Description:         "type of location from where the referenced template structure should be retrieved",
													MarkdownDescription: "type of location from where the referenced template structure should be retrieved",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Uri", "Id", "Kubernetes"),
													},
												},

												"kubernetes": schema.SingleNestedAttribute{
													Description:         "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",
													MarkdownDescription: "Reference to a Kubernetes CRD of type DevWorkspaceTemplate",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
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

												"registry_url": schema.StringAttribute{
													Description:         "Registry URL to pull the parent devfile from when using id in the parent reference.To ensure the parent devfile gets resolved consistently in different environments,it is recommended to always specify the 'registryUrl' when 'id' is used.",
													MarkdownDescription: "Registry URL to pull the parent devfile from when using id in the parent reference.To ensure the parent devfile gets resolved consistently in different environments,it is recommended to always specify the 'registryUrl' when 'id' is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "URI Reference of a parent devfile YAML file.It can be a full URL or a relative URI with the current devfile as the base URI.",
													MarkdownDescription: "URI Reference of a parent devfile YAML file.It can be a full URL or a relative URI with the current devfile as the base URI.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference.To specify 'version', 'id' must be defined and used as the import reference source.'version' can be either a specific stack version, or 'latest'.If no 'version' specified, default version will be used.",
													MarkdownDescription: "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference.To specify 'version', 'id' must be defined and used as the import reference source.'version' can be either a specific stack version, or 'latest'.If no 'version' specified, default version will be used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^(latest)|(([1-9])\.([0-9]+)\.([0-9]+)(\-[0-9a-z-]+(\.[0-9a-z-]+)*)?(\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?)$`), ""),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"volume": schema.SingleNestedAttribute{
											Description:         "Allows specifying the definition of a volumeshared by several other components",
											MarkdownDescription: "Allows specifying the definition of a volumeshared by several other components",
											Attributes: map[string]schema.Attribute{
												"ephemeral": schema.BoolAttribute{
													Description:         "Ephemeral volumes are not stored persistently across restarts. Defaultsto false",
													MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaultsto false",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size": schema.StringAttribute{
													Description:         "Size of the volume",
													MarkdownDescription: "Size of the volume",
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

							"default_editor": schema.StringAttribute{
								Description:         "The default editor to workspace create with. It could be a plugin ID or a URI.The plugin ID must have 'publisher/name/version' format.The URI must start from 'http://' or 'https://'.",
								MarkdownDescription: "The default editor to workspace create with. It could be a plugin ID or a URI.The plugin ID must have 'publisher/name/version' format.The URI must start from 'http://' or 'https://'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_namespace": schema.SingleNestedAttribute{
								Description:         "User's default namespace.",
								MarkdownDescription: "User's default namespace.",
								Attributes: map[string]schema.Attribute{
									"auto_provision": schema.BoolAttribute{
										Description:         "Indicates if is allowed to automatically create a user namespace.If it set to false, then user namespace must be pre-created by a cluster administrator.",
										MarkdownDescription: "Indicates if is allowed to automatically create a user namespace.If it set to false, then user namespace must be pre-created by a cluster administrator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"template": schema.StringAttribute{
										Description:         "If you don't create the user namespaces in advance, this field defines the Kubernetes namespace created when you start your first workspace.You can use '<username>' and '<userid>' placeholders, such as che-workspace-<username>.",
										MarkdownDescription: "If you don't create the user namespaces in advance, this field defines the Kubernetes namespace created when you start your first workspace.You can use '<username>' and '<userid>' placeholders, such as che-workspace-<username>.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`<username>|<userid>`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"default_plugins": schema.ListNestedAttribute{
								Description:         "Default plug-ins applied to DevWorkspaces.",
								MarkdownDescription: "Default plug-ins applied to DevWorkspaces.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"editor": schema.StringAttribute{
											Description:         "The editor ID to specify default plug-ins for.The plugin ID must have 'publisher/name/version' format.",
											MarkdownDescription: "The editor ID to specify default plug-ins for.The plugin ID must have 'publisher/name/version' format.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"plugins": schema.ListAttribute{
											Description:         "Default plug-in URIs for the specified editor.",
											MarkdownDescription: "Default plug-in URIs for the specified editor.",
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

							"deployment_strategy": schema.StringAttribute{
								Description:         "DeploymentStrategy defines the deployment strategy to use to replace existing workspace podswith new ones. The available deployment stragies are 'Recreate' and 'RollingUpdate'.With the 'Recreate' deployment strategy, the existing workspace pod is killed before the new one is created.With the 'RollingUpdate' deployment strategy, a new workspace pod is created and the existing workspace pod is deletedonly when the new workspace pod is in a ready state.If not specified, the default 'Recreate' deployment strategy is used.",
								MarkdownDescription: "DeploymentStrategy defines the deployment strategy to use to replace existing workspace podswith new ones. The available deployment stragies are 'Recreate' and 'RollingUpdate'.With the 'Recreate' deployment strategy, the existing workspace pod is killed before the new one is created.With the 'RollingUpdate' deployment strategy, a new workspace pod is created and the existing workspace pod is deletedonly when the new workspace pod is in a ready state.If not specified, the default 'Recreate' deployment strategy is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Recreate", "RollingUpdate"),
								},
							},

							"disable_container_build_capabilities": schema.BoolAttribute{
								Description:         "Disables the container build capabilities.When set to 'false' (the default value), the devEnvironments.security.containerSecurityContextfield is ignored, and the following container SecurityContext is applied: containerSecurityContext:   allowPrivilegeEscalation: true   capabilities:     add:     - SETGID     - SETUID",
								MarkdownDescription: "Disables the container build capabilities.When set to 'false' (the default value), the devEnvironments.security.containerSecurityContextfield is ignored, and the following container SecurityContext is applied: containerSecurityContext:   allowPrivilegeEscalation: true   capabilities:     add:     - SETGID     - SETUID",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gateway_container": schema.SingleNestedAttribute{
								Description:         "GatewayContainer configuration.",
								MarkdownDescription: "GatewayContainer configuration.",
								Attributes: map[string]schema.Attribute{
									"env": schema.ListNestedAttribute{
										Description:         "List of environment variables to set in the container.",
										MarkdownDescription: "List of environment variables to set in the container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
													MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
													MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value_from": schema.SingleNestedAttribute{
													Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
													MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a ConfigMap.",
															MarkdownDescription: "Selects a key of a ConfigMap.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key to select.",
																	MarkdownDescription: "The key to select.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the ConfigMap or its key must be defined",
																	MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"field_ref": schema.SingleNestedAttribute{
															Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
															MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																	MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "Path of the field to select in the specified API version.",
																	MarkdownDescription: "Path of the field to select in the specified API version.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"resource_field_ref": schema.SingleNestedAttribute{
															Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
															MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
															Attributes: map[string]schema.Attribute{
																"container_name": schema.StringAttribute{
																	Description:         "Container name: required for volumes, optional for env vars",
																	MarkdownDescription: "Container name: required for volumes, optional for env vars",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"divisor": schema.StringAttribute{
																	Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																	MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resource": schema.StringAttribute{
																	Description:         "Required: resource to select",
																	MarkdownDescription: "Required: resource to select",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"secret_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a secret in the pod's namespace",
															MarkdownDescription: "Selects a key of a secret in the pod's namespace",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from.  Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            true,
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": schema.StringAttribute{
										Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
										MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
										MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
										},
									},

									"name": schema.StringAttribute{
										Description:         "Container name.",
										MarkdownDescription: "Container name.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Compute resources required by this container.",
										MarkdownDescription: "Compute resources required by this container.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.SingleNestedAttribute{
												Description:         "Describes the maximum amount of compute resources allowed.",
												MarkdownDescription: "Describes the maximum amount of compute resources allowed.",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.StringAttribute{
														Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"memory": schema.StringAttribute{
														Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"request": schema.SingleNestedAttribute{
												Description:         "Describes the minimum amount of compute resources required.",
												MarkdownDescription: "Describes the minimum amount of compute resources required.",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.StringAttribute{
														Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"memory": schema.StringAttribute{
														Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
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

							"ignored_unrecoverable_events": schema.ListAttribute{
								Description:         "IgnoredUnrecoverableEvents defines a list of Kubernetes event names that shouldbe ignored when deciding to fail a workspace that is starting. This option should be usedif a transient cluster issue is triggering false-positives (for example, ifthe cluster occasionally encounters FailedScheduling events). Events listedhere will not trigger workspace failures.",
								MarkdownDescription: "IgnoredUnrecoverableEvents defines a list of Kubernetes event names that shouldbe ignored when deciding to fail a workspace that is starting. This option should be usedif a transient cluster issue is triggering false-positives (for example, ifthe cluster occasionally encounters FailedScheduling events). Events listedhere will not trigger workspace failures.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "ImagePullPolicy defines the imagePullPolicy used for containers in a DevWorkspace.",
								MarkdownDescription: "ImagePullPolicy defines the imagePullPolicy used for containers in a DevWorkspace.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
								},
							},

							"max_number_of_running_workspaces_per_cluster": schema.Int64Attribute{
								Description:         "The maximum number of concurrently running workspaces across the entire Kubernetes cluster.This applies to all users in the system. If the value is set to -1, it means there isno limit on the number of running workspaces.",
								MarkdownDescription: "The maximum number of concurrently running workspaces across the entire Kubernetes cluster.This applies to all users in the system. If the value is set to -1, it means there isno limit on the number of running workspaces.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(-1),
								},
							},

							"max_number_of_running_workspaces_per_user": schema.Int64Attribute{
								Description:         "The maximum number of running workspaces per user.The value, -1, allows users to run an unlimited number of workspaces.",
								MarkdownDescription: "The maximum number of running workspaces per user.The value, -1, allows users to run an unlimited number of workspaces.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(-1),
								},
							},

							"max_number_of_workspaces_per_user": schema.Int64Attribute{
								Description:         "Total number of workspaces, both stopped and running, that a user can keep.The value, -1, allows users to keep an unlimited number of workspaces.",
								MarkdownDescription: "Total number of workspaces, both stopped and running, that a user can keep.The value, -1, allows users to keep an unlimited number of workspaces.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(-1),
								},
							},

							"node_selector": schema.MapAttribute{
								Description:         "The node selector limits the nodes that can run the workspace pods.",
								MarkdownDescription: "The node selector limits the nodes that can run the workspace pods.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"persist_user_home": schema.SingleNestedAttribute{
								Description:         "PersistUserHome defines configuration options for persisting theuser home directory in workspaces.",
								MarkdownDescription: "PersistUserHome defines configuration options for persisting theuser home directory in workspaces.",
								Attributes: map[string]schema.Attribute{
									"disable_init_container": schema.BoolAttribute{
										Description:         "Determines whether the init container that initializes the persistent home directory should be disabled.When the '/home/user' directory is persisted, the init container is used to initialize the directory beforethe workspace starts. If set to true, the init container will not be created.Disabling the init container allows home persistence to be initialized by the entrypoint present in the workspace's first container component.This field is not used if the 'devEnvironments.persistUserHome.enabled' field is set to false.The init container is enabled by default.",
										MarkdownDescription: "Determines whether the init container that initializes the persistent home directory should be disabled.When the '/home/user' directory is persisted, the init container is used to initialize the directory beforethe workspace starts. If set to true, the init container will not be created.Disabling the init container allows home persistence to be initialized by the entrypoint present in the workspace's first container component.This field is not used if the 'devEnvironments.persistUserHome.enabled' field is set to false.The init container is enabled by default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Determines whether the user home directory in workspaces should persist betweenworkspace shutdown and startup.Must be used with the 'per-user' or 'per-workspace' PVC strategy in order to take effect.Disabled by default.",
										MarkdownDescription: "Determines whether the user home directory in workspaces should persist betweenworkspace shutdown and startup.Must be used with the 'per-user' or 'per-workspace' PVC strategy in order to take effect.Disabled by default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_scheduler_name": schema.StringAttribute{
								Description:         "Pod scheduler for the workspace pods.If not specified, the pod scheduler is set to the default scheduler on the cluster.",
								MarkdownDescription: "Pod scheduler for the workspace pods.If not specified, the pod scheduler is set to the default scheduler on the cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"project_clone_container": schema.SingleNestedAttribute{
								Description:         "Project clone container configuration.",
								MarkdownDescription: "Project clone container configuration.",
								Attributes: map[string]schema.Attribute{
									"env": schema.ListNestedAttribute{
										Description:         "List of environment variables to set in the container.",
										MarkdownDescription: "List of environment variables to set in the container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
													MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
													MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value_from": schema.SingleNestedAttribute{
													Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
													MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a ConfigMap.",
															MarkdownDescription: "Selects a key of a ConfigMap.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key to select.",
																	MarkdownDescription: "The key to select.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the ConfigMap or its key must be defined",
																	MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"field_ref": schema.SingleNestedAttribute{
															Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
															MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																	MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "Path of the field to select in the specified API version.",
																	MarkdownDescription: "Path of the field to select in the specified API version.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"resource_field_ref": schema.SingleNestedAttribute{
															Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
															MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
															Attributes: map[string]schema.Attribute{
																"container_name": schema.StringAttribute{
																	Description:         "Container name: required for volumes, optional for env vars",
																	MarkdownDescription: "Container name: required for volumes, optional for env vars",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"divisor": schema.StringAttribute{
																	Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																	MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resource": schema.StringAttribute{
																	Description:         "Required: resource to select",
																	MarkdownDescription: "Required: resource to select",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"secret_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a secret in the pod's namespace",
															MarkdownDescription: "Selects a key of a secret in the pod's namespace",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from.  Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            true,
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": schema.StringAttribute{
										Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
										MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
										MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
										},
									},

									"name": schema.StringAttribute{
										Description:         "Container name.",
										MarkdownDescription: "Container name.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Compute resources required by this container.",
										MarkdownDescription: "Compute resources required by this container.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.SingleNestedAttribute{
												Description:         "Describes the maximum amount of compute resources allowed.",
												MarkdownDescription: "Describes the maximum amount of compute resources allowed.",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.StringAttribute{
														Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"memory": schema.StringAttribute{
														Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"request": schema.SingleNestedAttribute{
												Description:         "Describes the minimum amount of compute resources required.",
												MarkdownDescription: "Describes the minimum amount of compute resources required.",
												Attributes: map[string]schema.Attribute{
													"cpu": schema.StringAttribute{
														Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"memory": schema.StringAttribute{
														Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
														MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
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

							"seconds_of_inactivity_before_idling": schema.Int64Attribute{
								Description:         "Idle timeout for workspaces in seconds.This timeout is the duration after which a workspace will be idled if there is no activity.To disable workspace idling due to inactivity, set this value to -1.",
								MarkdownDescription: "Idle timeout for workspaces in seconds.This timeout is the duration after which a workspace will be idled if there is no activity.To disable workspace idling due to inactivity, set this value to -1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"seconds_of_run_before_idling": schema.Int64Attribute{
								Description:         "Run timeout for workspaces in seconds.This timeout is the maximum duration a workspace runs.To disable workspace run timeout, set this value to -1.",
								MarkdownDescription: "Run timeout for workspaces in seconds.This timeout is the maximum duration a workspace runs.To disable workspace run timeout, set this value to -1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security": schema.SingleNestedAttribute{
								Description:         "Workspace security configuration.",
								MarkdownDescription: "Workspace security configuration.",
								Attributes: map[string]schema.Attribute{
									"container_security_context": schema.SingleNestedAttribute{
										Description:         "Container SecurityContext used by all workspace-related containers.If set, defined values are merged into the default Container SecurityContext configuration.Requires devEnvironments.disableContainerBuildCapabilities to be set to 'true' in order to take effect.",
										MarkdownDescription: "Container SecurityContext used by all workspace-related containers.If set, defined values are merged into the default Container SecurityContext configuration.Requires devEnvironments.disableContainerBuildCapabilities to be set to 'true' in order to take effect.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capabilities": schema.SingleNestedAttribute{
												Description:         "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
												Attributes: map[string]schema.Attribute{
													"add": schema.ListAttribute{
														Description:         "Added capabilities",
														MarkdownDescription: "Added capabilities",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"drop": schema.ListAttribute{
														Description:         "Removed capabilities",
														MarkdownDescription: "Removed capabilities",
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

											"privileged": schema.BoolAttribute{
												Description:         "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"proc_mount": schema.StringAttribute{
												Description:         "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_only_root_filesystem": schema.BoolAttribute{
												Description:         "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												Attributes: map[string]schema.Attribute{
													"level": schema.StringAttribute{
														Description:         "Level is SELinux level label that applies to the container.",
														MarkdownDescription: "Level is SELinux level label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"role": schema.StringAttribute{
														Description:         "Role is a SELinux role label that applies to the container.",
														MarkdownDescription: "Role is a SELinux role label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type is a SELinux type label that applies to the container.",
														MarkdownDescription: "Type is a SELinux type label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user": schema.StringAttribute{
														Description:         "User is a SELinux user label that applies to the container.",
														MarkdownDescription: "User is a SELinux user label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"seccomp_profile": schema.SingleNestedAttribute{
												Description:         "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must only be set if type is 'Localhost'.",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must only be set if type is 'Localhost'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"windows_options": schema.SingleNestedAttribute{
												Description:         "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
												MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
												Attributes: map[string]schema.Attribute{
													"gmsa_credential_spec": schema.StringAttribute{
														Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"gmsa_credential_spec_name": schema.StringAttribute{
														Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
														MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_process": schema.BoolAttribute{
														Description:         "HostProcess determines if a container should be run as a 'Host Process' container.This field is alpha-level and will only be honored by components that enable theWindowsHostProcessContainers feature flag. Setting this field without the featureflag will result in errors when validating the Pod. All of a Pod's containers musthave the same effective HostProcess value (it is not allowed to have a mix of HostProcesscontainers and non-HostProcess containers).  In addition, if HostProcess is truethen HostNetwork must also be set to true.",
														MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.This field is alpha-level and will only be honored by components that enable theWindowsHostProcessContainers feature flag. Setting this field without the featureflag will result in errors when validating the Pod. All of a Pod's containers musthave the same effective HostProcess value (it is not allowed to have a mix of HostProcesscontainers and non-HostProcess containers).  In addition, if HostProcess is truethen HostNetwork must also be set to true.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user_name": schema.StringAttribute{
														Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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

									"pod_security_context": schema.SingleNestedAttribute{
										Description:         "PodSecurityContext used by all workspace-related pods.If set, defined values are merged into the default PodSecurityContext configuration.",
										MarkdownDescription: "PodSecurityContext used by all workspace-related pods.If set, defined values are merged into the default PodSecurityContext configuration.",
										Attributes: map[string]schema.Attribute{
											"fs_group": schema.Int64Attribute{
												Description:         "A special supplemental group that applies to all containers in a pod.Some volume types allow the Kubelet to change the ownership of that volumeto be owned by the pod:1. The owning GID will be the FSGroup2. The setgid bit is set (new files created in the volume will be owned by FSGroup)If unset, the Kubelet will not modify the ownership and permissions of any volume.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "A special supplemental group that applies to all containers in a pod.Some volume types allow the Kubelet to change the ownership of that volumeto be owned by the pod:1. The owning GID will be the FSGroup2. The setgid bit is set (new files created in the volume will be owned by FSGroup)If unset, the Kubelet will not modify the ownership and permissions of any volume.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"fs_group_change_policy": schema.StringAttribute{
												Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volumebefore being exposed inside Pod. This field will only apply tovolume types which support fsGroup based ownership(and permissions).It will have no effect on ephemeral volume types such as: secret, configmapsand emptydir.Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volumebefore being exposed inside Pod. This field will only apply tovolume types which support fsGroup based ownership(and permissions).It will have no effect on ephemeral volume types such as: secret, configmapsand emptydir.Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "The SELinux context to be applied to all containers.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in SecurityContext.  If set inboth SecurityContext and PodSecurityContext, the value specified in SecurityContexttakes precedence for that container.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The SELinux context to be applied to all containers.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in SecurityContext.  If set inboth SecurityContext and PodSecurityContext, the value specified in SecurityContexttakes precedence for that container.Note that this field cannot be set when spec.os.name is windows.",
												Attributes: map[string]schema.Attribute{
													"level": schema.StringAttribute{
														Description:         "Level is SELinux level label that applies to the container.",
														MarkdownDescription: "Level is SELinux level label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"role": schema.StringAttribute{
														Description:         "Role is a SELinux role label that applies to the container.",
														MarkdownDescription: "Role is a SELinux role label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type is a SELinux type label that applies to the container.",
														MarkdownDescription: "Type is a SELinux type label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user": schema.StringAttribute{
														Description:         "User is a SELinux user label that applies to the container.",
														MarkdownDescription: "User is a SELinux user label that applies to the container.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"seccomp_profile": schema.SingleNestedAttribute{
												Description:         "The seccomp options to use by the containers in this pod.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The seccomp options to use by the containers in this pod.Note that this field cannot be set when spec.os.name is windows.",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must only be set if type is 'Localhost'.",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must only be set if type is 'Localhost'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"supplemental_groups": schema.ListAttribute{
												Description:         "A list of groups applied to the first process run in each container, in additionto the container's primary GID, the fsGroup (if specified), and group membershipsdefined in the container image for the uid of the container process. If unspecified,no additional groups are added to any container. Note that group membershipsdefined in the container image for the uid of the container process are still effective,even if they are not included in this list.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "A list of groups applied to the first process run in each container, in additionto the container's primary GID, the fsGroup (if specified), and group membershipsdefined in the container image for the uid of the container process. If unspecified,no additional groups are added to any container. Note that group membershipsdefined in the container image for the uid of the container process are still effective,even if they are not included in this list.Note that this field cannot be set when spec.os.name is windows.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sysctls": schema.ListNestedAttribute{
												Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupportedsysctls (by the container runtime) might fail to launch.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupportedsysctls (by the container runtime) might fail to launch.Note that this field cannot be set when spec.os.name is windows.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of a property to set",
															MarkdownDescription: "Name of a property to set",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value of a property to set",
															MarkdownDescription: "Value of a property to set",
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

											"windows_options": schema.SingleNestedAttribute{
												Description:         "The Windows specific settings applied to all containers.If unspecified, the options within a container's SecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
												MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options within a container's SecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
												Attributes: map[string]schema.Attribute{
													"gmsa_credential_spec": schema.StringAttribute{
														Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"gmsa_credential_spec_name": schema.StringAttribute{
														Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
														MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_process": schema.BoolAttribute{
														Description:         "HostProcess determines if a container should be run as a 'Host Process' container.This field is alpha-level and will only be honored by components that enable theWindowsHostProcessContainers feature flag. Setting this field without the featureflag will result in errors when validating the Pod. All of a Pod's containers musthave the same effective HostProcess value (it is not allowed to have a mix of HostProcesscontainers and non-HostProcess containers).  In addition, if HostProcess is truethen HostNetwork must also be set to true.",
														MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.This field is alpha-level and will only be honored by components that enable theWindowsHostProcessContainers feature flag. Setting this field without the featureflag will result in errors when validating the Pod. All of a Pod's containers musthave the same effective HostProcess value (it is not allowed to have a mix of HostProcesscontainers and non-HostProcess containers).  In addition, if HostProcess is truethen HostNetwork must also be set to true.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user_name": schema.StringAttribute{
														Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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

							"service_account": schema.StringAttribute{
								Description:         "ServiceAccount to use by the DevWorkspace operator when starting the workspaces.",
								MarkdownDescription: "ServiceAccount to use by the DevWorkspace operator when starting the workspaces.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},

							"service_account_tokens": schema.ListNestedAttribute{
								Description:         "List of ServiceAccount tokens that will be mounted into workspace pods as projected volumes.",
								MarkdownDescription: "List of ServiceAccount tokens that will be mounted into workspace pods as projected volumes.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"audience": schema.StringAttribute{
											Description:         "Audience is the intended audience of the token. A recipient of a tokenmust identify itself with an identifier specified in the audience of thetoken, and otherwise should reject the token. The audience defaults to theidentifier of the apiserver.",
											MarkdownDescription: "Audience is the intended audience of the token. A recipient of a tokenmust identify itself with an identifier specified in the audience of thetoken, and otherwise should reject the token. The audience defaults to theidentifier of the apiserver.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"expiration_seconds": schema.Int64Attribute{
											Description:         "ExpirationSeconds is the requested duration of validity of the serviceaccount token. As the token approaches expiration, the kubelet volumeplugin will proactively rotate the service account token. The kubelet willstart trying to rotate the token if the token is older than 80 percent ofits time to live or if the token is older than 24 hours. Defaults to 1 hourand must be at least 10 minutes.",
											MarkdownDescription: "ExpirationSeconds is the requested duration of validity of the serviceaccount token. As the token approaches expiration, the kubelet volumeplugin will proactively rotate the service account token. The kubelet willstart trying to rotate the token if the token is older than 80 percent ofits time to live or if the token is older than 24 hours. Defaults to 1 hourand must be at least 10 minutes.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(600),
											},
										},

										"mount_path": schema.StringAttribute{
											Description:         "Path within the workspace container at which the token should be mounted.  Mustnot contain ':'.",
											MarkdownDescription: "Path within the workspace container at which the token should be mounted.  Mustnot contain ':'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Identifiable name of the ServiceAccount token.If multiple ServiceAccount tokens use the same mount path, a generic name will be usedfor the projected volume instead.",
											MarkdownDescription: "Identifiable name of the ServiceAccount token.If multiple ServiceAccount tokens use the same mount path, a generic name will be usedfor the projected volume instead.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "Path is the path relative to the mount point of the file to project thetoken into.",
											MarkdownDescription: "Path is the path relative to the mount point of the file to project thetoken into.",
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

							"start_timeout_seconds": schema.Int64Attribute{
								Description:         "StartTimeoutSeconds determines the maximum duration (in seconds) that a workspace can take to startbefore it is automatically failed.If not specified, the default value of 300 seconds (5 minutes) is used.",
								MarkdownDescription: "StartTimeoutSeconds determines the maximum duration (in seconds) that a workspace can take to startbefore it is automatically failed.If not specified, the default value of 300 seconds (5 minutes) is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"storage": schema.SingleNestedAttribute{
								Description:         "Workspaces persistent storage.",
								MarkdownDescription: "Workspaces persistent storage.",
								Attributes: map[string]schema.Attribute{
									"per_user_strategy_pvc_config": schema.SingleNestedAttribute{
										Description:         "PVC settings when using the 'per-user' PVC strategy.",
										MarkdownDescription: "PVC settings when using the 'per-user' PVC strategy.",
										Attributes: map[string]schema.Attribute{
											"claim_size": schema.StringAttribute{
												Description:         "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",
												MarkdownDescription: "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"storage_class": schema.StringAttribute{
												Description:         "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",
												MarkdownDescription: "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"per_workspace_strategy_pvc_config": schema.SingleNestedAttribute{
										Description:         "PVC settings when using the 'per-workspace' PVC strategy.",
										MarkdownDescription: "PVC settings when using the 'per-workspace' PVC strategy.",
										Attributes: map[string]schema.Attribute{
											"claim_size": schema.StringAttribute{
												Description:         "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",
												MarkdownDescription: "Persistent Volume Claim size. To update the claim size, the storage class that provisions it must support resizing.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"storage_class": schema.StringAttribute{
												Description:         "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",
												MarkdownDescription: "Storage class for the Persistent Volume Claim. When omitted or left blank, a default storage class is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pvc_strategy": schema.StringAttribute{
										Description:         "Persistent volume claim strategy for the Che server.The supported strategies are: 'per-user' (all workspaces PVCs in one volume),'per-workspace' (each workspace is given its own individual PVC)and 'ephemeral' (non-persistent storage where local changes will be lost whenthe workspace is stopped.)",
										MarkdownDescription: "Persistent volume claim strategy for the Che server.The supported strategies are: 'per-user' (all workspaces PVCs in one volume),'per-workspace' (each workspace is given its own individual PVC)and 'ephemeral' (non-persistent storage where local changes will be lost whenthe workspace is stopped.)",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("common", "per-user", "per-workspace", "ephemeral"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "The pod tolerations of the workspace pods limit where the workspace pods can run.",
								MarkdownDescription: "The pod tolerations of the workspace pods limit where the workspace pods can run.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"trusted_certs": schema.SingleNestedAttribute{
								Description:         "Trusted certificate settings.",
								MarkdownDescription: "Trusted certificate settings.",
								Attributes: map[string]schema.Attribute{
									"git_trusted_certs_config_map_name": schema.StringAttribute{
										Description:         "The ConfigMap contains certificates to propagate to the Che components and to provide a particular configuration for Git.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/deploying-che-with-support-for-git-repositories-with-self-signed-certificates/The ConfigMap must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
										MarkdownDescription: "The ConfigMap contains certificates to propagate to the Che components and to provide a particular configuration for Git.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/deploying-che-with-support-for-git-repositories-with-self-signed-certificates/The ConfigMap must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"user": schema.SingleNestedAttribute{
								Description:         "User configuration.",
								MarkdownDescription: "User configuration.",
								Attributes: map[string]schema.Attribute{
									"cluster_roles": schema.ListAttribute{
										Description:         "Additional ClusterRoles assigned to the user.The role must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
										MarkdownDescription: "Additional ClusterRoles assigned to the user.The role must have 'app.kubernetes.io/part-of=che.eclipse.org' label.",
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

							"workspaces_pod_annotations": schema.MapAttribute{
								Description:         "WorkspacesPodAnnotations defines additional annotations for workspace pods.",
								MarkdownDescription: "WorkspacesPodAnnotations defines additional annotations for workspace pods.",
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

					"git_services": schema.SingleNestedAttribute{
						Description:         "A configuration that allows users to work with remote Git repositories.",
						MarkdownDescription: "A configuration that allows users to work with remote Git repositories.",
						Attributes: map[string]schema.Attribute{
							"azure": schema.ListNestedAttribute{
								Description:         "Enables users to work with repositories hosted on Azure DevOps Service (dev.azure.com).",
								MarkdownDescription: "Enables users to work with repositories hosted on Azure DevOps Service (dev.azure.com).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"secret_name": schema.StringAttribute{
											Description:         "Kubernetes secret, that contains Base64-encoded Azure DevOps Service Application ID and Client Secret.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-microsoft-azure-devops-services",
											MarkdownDescription: "Kubernetes secret, that contains Base64-encoded Azure DevOps Service Application ID and Client Secret.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-microsoft-azure-devops-services",
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

							"bitbucket": schema.ListNestedAttribute{
								Description:         "Enables users to work with repositories hosted on Bitbucket (bitbucket.org or self-hosted).",
								MarkdownDescription: "Enables users to work with repositories hosted on Bitbucket (bitbucket.org or self-hosted).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"endpoint": schema.StringAttribute{
											Description:         "Bitbucket server endpoint URL.Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/.",
											MarkdownDescription: "Bitbucket server endpoint URL.Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Kubernetes secret, that contains Base64-encoded Bitbucket OAuth 1.0 or OAuth 2.0 data.See the following pages for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/and https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-the-bitbucket-cloud/.",
											MarkdownDescription: "Kubernetes secret, that contains Base64-encoded Bitbucket OAuth 1.0 or OAuth 2.0 data.See the following pages for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-1-for-a-bitbucket-server/and https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-the-bitbucket-cloud/.",
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

							"github": schema.ListNestedAttribute{
								Description:         "Enables users to work with repositories hosted on GitHub (github.com or GitHub Enterprise).",
								MarkdownDescription: "Enables users to work with repositories hosted on GitHub (github.com or GitHub Enterprise).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"disable_subdomain_isolation": schema.BoolAttribute{
											Description:         "Disables subdomain isolation.Deprecated in favor of 'che.eclipse.org/scm-github-disable-subdomain-isolation' annotation.See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
											MarkdownDescription: "Disables subdomain isolation.Deprecated in favor of 'che.eclipse.org/scm-github-disable-subdomain-isolation' annotation.See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"endpoint": schema.StringAttribute{
											Description:         "GitHub server endpoint URL.Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation.See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
											MarkdownDescription: "GitHub server endpoint URL.Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation.See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Kubernetes secret, that contains Base64-encoded GitHub OAuth Client id and GitHub OAuth Client secret.See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
											MarkdownDescription: "Kubernetes secret, that contains Base64-encoded GitHub OAuth Client id and GitHub OAuth Client secret.See the following page for details: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-github/.",
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

							"gitlab": schema.ListNestedAttribute{
								Description:         "Enables users to work with repositories hosted on GitLab (gitlab.com or self-hosted).",
								MarkdownDescription: "Enables users to work with repositories hosted on GitLab (gitlab.com or self-hosted).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"endpoint": schema.StringAttribute{
											Description:         "GitLab server endpoint URL.Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
											MarkdownDescription: "GitLab server endpoint URL.Deprecated in favor of 'che.eclipse.org/scm-server-endpoint' annotation.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Kubernetes secret, that contains Base64-encoded GitHub Application id and GitLab Application Client secret.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
											MarkdownDescription: "Kubernetes secret, that contains Base64-encoded GitHub Application id and GitLab Application Client secret.See the following page: https://www.eclipse.org/che/docs/stable/administration-guide/configuring-oauth-2-for-gitlab/.",
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

					"networking": schema.SingleNestedAttribute{
						Description:         "Networking, Che authentication, and TLS configuration.",
						MarkdownDescription: "Networking, Che authentication, and TLS configuration.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Defines annotations which will be set for an Ingress (a route for OpenShift platform).The defaults for kubernetes platforms are:    kubernetes.io/ingress.class:                       'nginx'    nginx.ingress.kubernetes.io/proxy-read-timeout:    '3600',    nginx.ingress.kubernetes.io/proxy-connect-timeout: '3600',    nginx.ingress.kubernetes.io/ssl-redirect:          'true'",
								MarkdownDescription: "Defines annotations which will be set for an Ingress (a route for OpenShift platform).The defaults for kubernetes platforms are:    kubernetes.io/ingress.class:                       'nginx'    nginx.ingress.kubernetes.io/proxy-read-timeout:    '3600',    nginx.ingress.kubernetes.io/proxy-connect-timeout: '3600',    nginx.ingress.kubernetes.io/ssl-redirect:          'true'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"auth": schema.SingleNestedAttribute{
								Description:         "Authentication settings.",
								MarkdownDescription: "Authentication settings.",
								Attributes: map[string]schema.Attribute{
									"advanced_authorization": schema.SingleNestedAttribute{
										Description:         "Advance authorization settings. Determines which users and groups are allowed to access Che.User is allowed to access Che if he/she is either in the 'allowUsers' list or is member of group from 'allowGroups' listand not in neither the 'denyUsers' list nor is member of group from 'denyGroups' list.If 'allowUsers' and 'allowGroups' are empty, then all users are allowed to access Che.if 'denyUsers' and 'denyGroups' are empty, then no users are denied to access Che.",
										MarkdownDescription: "Advance authorization settings. Determines which users and groups are allowed to access Che.User is allowed to access Che if he/she is either in the 'allowUsers' list or is member of group from 'allowGroups' listand not in neither the 'denyUsers' list nor is member of group from 'denyGroups' list.If 'allowUsers' and 'allowGroups' are empty, then all users are allowed to access Che.if 'denyUsers' and 'denyGroups' are empty, then no users are denied to access Che.",
										Attributes: map[string]schema.Attribute{
											"allow_groups": schema.ListAttribute{
												Description:         "List of groups allowed to access Che (currently supported in OpenShift only).",
												MarkdownDescription: "List of groups allowed to access Che (currently supported in OpenShift only).",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"allow_users": schema.ListAttribute{
												Description:         "List of users allowed to access Che.",
												MarkdownDescription: "List of users allowed to access Che.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"deny_groups": schema.ListAttribute{
												Description:         "List of groups denied to access Che (currently supported in OpenShift only).",
												MarkdownDescription: "List of groups denied to access Che (currently supported in OpenShift only).",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"deny_users": schema.ListAttribute{
												Description:         "List of users denied to access Che.",
												MarkdownDescription: "List of users denied to access Che.",
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

									"gateway": schema.SingleNestedAttribute{
										Description:         "Gateway settings.",
										MarkdownDescription: "Gateway settings.",
										Attributes: map[string]schema.Attribute{
											"config_labels": schema.MapAttribute{
												Description:         "Gateway configuration labels.",
												MarkdownDescription: "Gateway configuration labels.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"deployment": schema.SingleNestedAttribute{
												Description:         "Deployment override options.Since gateway deployment consists of several containers, they must be distinguished in the configuration by their names:- 'gateway'- 'configbump'- 'oauth-proxy'- 'kube-rbac-proxy'",
												MarkdownDescription: "Deployment override options.Since gateway deployment consists of several containers, they must be distinguished in the configuration by their names:- 'gateway'- 'configbump'- 'oauth-proxy'- 'kube-rbac-proxy'",
												Attributes: map[string]schema.Attribute{
													"containers": schema.ListNestedAttribute{
														Description:         "List of containers belonging to the pod.",
														MarkdownDescription: "List of containers belonging to the pod.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"env": schema.ListNestedAttribute{
																	Description:         "List of environment variables to set in the container.",
																	MarkdownDescription: "List of environment variables to set in the container.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																				MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																				MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"value_from": schema.SingleNestedAttribute{
																				Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																				MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
																				Attributes: map[string]schema.Attribute{
																					"config_map_key_ref": schema.SingleNestedAttribute{
																						Description:         "Selects a key of a ConfigMap.",
																						MarkdownDescription: "Selects a key of a ConfigMap.",
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The key to select.",
																								MarkdownDescription: "The key to select.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"name": schema.StringAttribute{
																								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"optional": schema.BoolAttribute{
																								Description:         "Specify whether the ConfigMap or its key must be defined",
																								MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"field_ref": schema.SingleNestedAttribute{
																						Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																						MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																						Attributes: map[string]schema.Attribute{
																							"api_version": schema.StringAttribute{
																								Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																								MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"field_path": schema.StringAttribute{
																								Description:         "Path of the field to select in the specified API version.",
																								MarkdownDescription: "Path of the field to select in the specified API version.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"resource_field_ref": schema.SingleNestedAttribute{
																						Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																						MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																						Attributes: map[string]schema.Attribute{
																							"container_name": schema.StringAttribute{
																								Description:         "Container name: required for volumes, optional for env vars",
																								MarkdownDescription: "Container name: required for volumes, optional for env vars",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"divisor": schema.StringAttribute{
																								Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																								MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"resource": schema.StringAttribute{
																								Description:         "Required: resource to select",
																								MarkdownDescription: "Required: resource to select",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},
																						},
																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"secret_key_ref": schema.SingleNestedAttribute{
																						Description:         "Selects a key of a secret in the pod's namespace",
																						MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "The key of the secret to select from.  Must be a valid secret key.",
																								MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"name": schema.StringAttribute{
																								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																								Required:            false,
																								Optional:            true,
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
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"image": schema.StringAttribute{
																	Description:         "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
																	MarkdownDescription: "Container image. Omit it or leave it empty to use the default container image provided by the Operator.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image_pull_policy": schema.StringAttribute{
																	Description:         "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
																	MarkdownDescription: "Image pull policy. Default value is 'Always' for 'nightly', 'next' or 'latest' images, and 'IfNotPresent' in other cases.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Always", "IfNotPresent", "Never"),
																	},
																},

																"name": schema.StringAttribute{
																	Description:         "Container name.",
																	MarkdownDescription: "Container name.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Compute resources required by this container.",
																	MarkdownDescription: "Compute resources required by this container.",
																	Attributes: map[string]schema.Attribute{
																		"limits": schema.SingleNestedAttribute{
																			Description:         "Describes the maximum amount of compute resources allowed.",
																			MarkdownDescription: "Describes the maximum amount of compute resources allowed.",
																			Attributes: map[string]schema.Attribute{
																				"cpu": schema.StringAttribute{
																					Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																					MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"memory": schema.StringAttribute{
																					Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																					MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"request": schema.SingleNestedAttribute{
																			Description:         "Describes the minimum amount of compute resources required.",
																			MarkdownDescription: "Describes the minimum amount of compute resources required.",
																			Attributes: map[string]schema.Attribute{
																				"cpu": schema.StringAttribute{
																					Description:         "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																					MarkdownDescription: "CPU, in cores. (500m = .5 cores)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"memory": schema.StringAttribute{
																					Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
																					MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)If the value is not specified, then the default value is set depending on the component.If value is '0', then no value is set for the component.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"security_context": schema.SingleNestedAttribute{
														Description:         "Security options the pod should run with.",
														MarkdownDescription: "Security options the pod should run with.",
														Attributes: map[string]schema.Attribute{
															"fs_group": schema.Int64Attribute{
																Description:         "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
																MarkdownDescription: "A special supplemental group that applies to all containers in a pod. The default value is '1724'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"run_as_user": schema.Int64Attribute{
																Description:         "The UID to run the entrypoint of the container process. The default value is '1724'.",
																MarkdownDescription: "The UID to run the entrypoint of the container process. The default value is '1724'.",
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

											"kube_rbac_proxy": schema.SingleNestedAttribute{
												Description:         "Configuration for kube-rbac-proxy within the Che gateway pod.",
												MarkdownDescription: "Configuration for kube-rbac-proxy within the Che gateway pod.",
												Attributes: map[string]schema.Attribute{
													"log_level": schema.Int64Attribute{
														Description:         "The glog log level for the kube-rbac-proxy container within the gateway pod. Larger values represent a higher verbosity. The default value is '0'.",
														MarkdownDescription: "The glog log level for the kube-rbac-proxy container within the gateway pod. Larger values represent a higher verbosity. The default value is '0'.",
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

											"o_auth_proxy": schema.SingleNestedAttribute{
												Description:         "Configuration for oauth-proxy within the Che gateway pod.",
												MarkdownDescription: "Configuration for oauth-proxy within the Che gateway pod.",
												Attributes: map[string]schema.Attribute{
													"cookie_expire_seconds": schema.Int64Attribute{
														Description:         "Expire timeframe for cookie. If set to 0, cookie becomes a session-cookie which will expire when the browser is closed.",
														MarkdownDescription: "Expire timeframe for cookie. If set to 0, cookie becomes a session-cookie which will expire when the browser is closed.",
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

											"traefik": schema.SingleNestedAttribute{
												Description:         "Configuration for Traefik within the Che gateway pod.",
												MarkdownDescription: "Configuration for Traefik within the Che gateway pod.",
												Attributes: map[string]schema.Attribute{
													"log_level": schema.StringAttribute{
														Description:         "The log level for the Traefik container within the gateway pod: 'DEBUG', 'INFO', 'WARN', 'ERROR', 'FATAL', or 'PANIC'. The default value is 'INFO'",
														MarkdownDescription: "The log level for the Traefik container within the gateway pod: 'DEBUG', 'INFO', 'WARN', 'ERROR', 'FATAL', or 'PANIC'. The default value is 'INFO'",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("DEBUG", "INFO", "WARN", "ERROR", "FATAL", "PANIC"),
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

									"identity_provider_url": schema.StringAttribute{
										Description:         "Public URL of the Identity Provider server.",
										MarkdownDescription: "Public URL of the Identity Provider server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"identity_token": schema.StringAttribute{
										Description:         "Identity token to be passed to upstream. There are two types of tokens supported: 'id_token' and 'access_token'.Default value is 'id_token'.This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
										MarkdownDescription: "Identity token to be passed to upstream. There are two types of tokens supported: 'id_token' and 'access_token'.Default value is 'id_token'.This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("id_token", "access_token"),
										},
									},

									"o_auth_access_token_inactivity_timeout_seconds": schema.Int64Attribute{
										Description:         "Inactivity timeout for tokens to set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.0 means tokens for this client never time out.",
										MarkdownDescription: "Inactivity timeout for tokens to set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.0 means tokens for this client never time out.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"o_auth_access_token_max_age_seconds": schema.Int64Attribute{
										Description:         "Access token max age for tokens to set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.0 means no expiration.",
										MarkdownDescription: "Access token max age for tokens to set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.0 means no expiration.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"o_auth_client_name": schema.StringAttribute{
										Description:         "Name of the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.",
										MarkdownDescription: "Name of the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"o_auth_scope": schema.StringAttribute{
										Description:         "Access Token Scope.This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
										MarkdownDescription: "Access Token Scope.This field is specific to Che installations made for Kubernetes only and ignored for OpenShift.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"o_auth_secret": schema.StringAttribute{
										Description:         "Name of the secret set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.For Kubernetes, this can either be the plain text oAuthSecret value, or the name of a kubernetes secret which contains akey 'oAuthSecret' and the value is the secret. NOTE: this secret must exist in the same namespace as the 'CheCluster'resource and contain the label 'app.kubernetes.io/part-of=che.eclipse.org'.",
										MarkdownDescription: "Name of the secret set in the OpenShift 'OAuthClient' resource used to set up identity federation on the OpenShift side.For Kubernetes, this can either be the plain text oAuthSecret value, or the name of a kubernetes secret which contains akey 'oAuthSecret' and the value is the secret. NOTE: this secret must exist in the same namespace as the 'CheCluster'resource and contain the label 'app.kubernetes.io/part-of=che.eclipse.org'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"domain": schema.StringAttribute{
								Description:         "For an OpenShift cluster, the Operator uses the domain to generate a hostname for the route.The generated hostname follows this pattern: che-<che-namespace>.<domain>. The <che-namespace> is the namespace where the CheCluster CRD is created.In conjunction with labels, it creates a route served by a non-default Ingress controller.For a Kubernetes cluster, it contains a global ingress domain. There are no default values: you must specify them.",
								MarkdownDescription: "For an OpenShift cluster, the Operator uses the domain to generate a hostname for the route.The generated hostname follows this pattern: che-<che-namespace>.<domain>. The <che-namespace> is the namespace where the CheCluster CRD is created.In conjunction with labels, it creates a route served by a non-default Ingress controller.For a Kubernetes cluster, it contains a global ingress domain. There are no default values: you must specify them.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "The public hostname of the installed Che server.",
								MarkdownDescription: "The public hostname of the installed Che server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress_class_name": schema.StringAttribute{
								Description:         "IngressClassName is the name of an IngressClass cluster resource.If a class name is defined in both the 'IngressClassName' field and the 'kubernetes.io/ingress.class' annotation,'IngressClassName' field takes precedence.",
								MarkdownDescription: "IngressClassName is the name of an IngressClass cluster resource.If a class name is defined in both the 'IngressClassName' field and the 'kubernetes.io/ingress.class' annotation,'IngressClassName' field takes precedence.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Defines labels which will be set for an Ingress (a route for OpenShift platform).",
								MarkdownDescription: "Defines labels which will be set for an Ingress (a route for OpenShift platform).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tls_secret_name": schema.StringAttribute{
								Description:         "The name of the secret used to set up Ingress TLS termination.If the field is an empty string, the default cluster certificate is used.The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
								MarkdownDescription: "The name of the secret used to set up Ingress TLS termination.If the field is an empty string, the default cluster certificate is used.The secret must have a 'app.kubernetes.io/part-of=che.eclipse.org' label.",
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
	}
}

func (r *OrgEclipseCheCheClusterV2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_org_eclipse_che_che_cluster_v2_manifest")

	var model OrgEclipseCheCheClusterV2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("org.eclipse.che/v2")
	model.Kind = pointer.String("CheCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
