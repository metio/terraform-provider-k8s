/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package workspace_devfile_io_v1alpha2

import (
	"context"
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
	_ datasource.DataSource = &WorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2Manifest{}
)

func NewWorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2Manifest() datasource.DataSource {
	return &WorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2Manifest{}
}

type WorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2Manifest struct{}

type WorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2ManifestData struct {
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
		Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
		Commands   *[]struct {
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
			Custom *struct {
				CommandClass     *string            `tfsdk:"command_class" json:"commandClass,omitempty"`
				EmbeddedResource *map[string]string `tfsdk:"embedded_resource" json:"embeddedResource,omitempty"`
				Group            *struct {
					IsDefault *bool   `tfsdk:"is_default" json:"isDefault,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Label *string `tfsdk:"label" json:"label,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
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
		} `tfsdk:"components" json:"components,omitempty"`
		DependentProjects *[]struct {
			Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
			ClonePath  *string            `tfsdk:"clone_path" json:"clonePath,omitempty"`
			Custom     *struct {
				EmbeddedResource   *map[string]string `tfsdk:"embedded_resource" json:"embeddedResource,omitempty"`
				ProjectSourceClass *string            `tfsdk:"project_source_class" json:"projectSourceClass,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Git *struct {
				CheckoutFrom *struct {
					Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
					Revision *string `tfsdk:"revision" json:"revision,omitempty"`
				} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
				Remotes *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
			} `tfsdk:"git" json:"git,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			SourceType *string `tfsdk:"source_type" json:"sourceType,omitempty"`
			Zip        *struct {
				Location *string `tfsdk:"location" json:"location,omitempty"`
			} `tfsdk:"zip" json:"zip,omitempty"`
		} `tfsdk:"dependent_projects" json:"dependentProjects,omitempty"`
		Events *struct {
			PostStart *[]string `tfsdk:"post_start" json:"postStart,omitempty"`
			PostStop  *[]string `tfsdk:"post_stop" json:"postStop,omitempty"`
			PreStart  *[]string `tfsdk:"pre_start" json:"preStart,omitempty"`
			PreStop   *[]string `tfsdk:"pre_stop" json:"preStop,omitempty"`
		} `tfsdk:"events" json:"events,omitempty"`
		Parent *struct {
			Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
			Commands   *[]struct {
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
			} `tfsdk:"components" json:"components,omitempty"`
			DependentProjects *[]struct {
				Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ClonePath  *string            `tfsdk:"clone_path" json:"clonePath,omitempty"`
				Git        *struct {
					CheckoutFrom *struct {
						Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
						Revision *string `tfsdk:"revision" json:"revision,omitempty"`
					} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
					Remotes *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
				} `tfsdk:"git" json:"git,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				SourceType *string `tfsdk:"source_type" json:"sourceType,omitempty"`
				Zip        *struct {
					Location *string `tfsdk:"location" json:"location,omitempty"`
				} `tfsdk:"zip" json:"zip,omitempty"`
			} `tfsdk:"dependent_projects" json:"dependentProjects,omitempty"`
			Id                  *string `tfsdk:"id" json:"id,omitempty"`
			ImportReferenceType *string `tfsdk:"import_reference_type" json:"importReferenceType,omitempty"`
			Kubernetes          *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
			Projects *[]struct {
				Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				ClonePath  *string            `tfsdk:"clone_path" json:"clonePath,omitempty"`
				Git        *struct {
					CheckoutFrom *struct {
						Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
						Revision *string `tfsdk:"revision" json:"revision,omitempty"`
					} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
					Remotes *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
				} `tfsdk:"git" json:"git,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				SourceType *string `tfsdk:"source_type" json:"sourceType,omitempty"`
				Zip        *struct {
					Location *string `tfsdk:"location" json:"location,omitempty"`
				} `tfsdk:"zip" json:"zip,omitempty"`
			} `tfsdk:"projects" json:"projects,omitempty"`
			RegistryUrl     *string `tfsdk:"registry_url" json:"registryUrl,omitempty"`
			StarterProjects *[]struct {
				Attributes  *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				Description *string            `tfsdk:"description" json:"description,omitempty"`
				Git         *struct {
					CheckoutFrom *struct {
						Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
						Revision *string `tfsdk:"revision" json:"revision,omitempty"`
					} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
					Remotes *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
				} `tfsdk:"git" json:"git,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				SourceType *string `tfsdk:"source_type" json:"sourceType,omitempty"`
				SubDir     *string `tfsdk:"sub_dir" json:"subDir,omitempty"`
				Zip        *struct {
					Location *string `tfsdk:"location" json:"location,omitempty"`
				} `tfsdk:"zip" json:"zip,omitempty"`
			} `tfsdk:"starter_projects" json:"starterProjects,omitempty"`
			Uri       *string            `tfsdk:"uri" json:"uri,omitempty"`
			Variables *map[string]string `tfsdk:"variables" json:"variables,omitempty"`
			Version   *string            `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"parent" json:"parent,omitempty"`
		Projects *[]struct {
			Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
			ClonePath  *string            `tfsdk:"clone_path" json:"clonePath,omitempty"`
			Custom     *struct {
				EmbeddedResource   *map[string]string `tfsdk:"embedded_resource" json:"embeddedResource,omitempty"`
				ProjectSourceClass *string            `tfsdk:"project_source_class" json:"projectSourceClass,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Git *struct {
				CheckoutFrom *struct {
					Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
					Revision *string `tfsdk:"revision" json:"revision,omitempty"`
				} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
				Remotes *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
			} `tfsdk:"git" json:"git,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			SourceType *string `tfsdk:"source_type" json:"sourceType,omitempty"`
			Zip        *struct {
				Location *string `tfsdk:"location" json:"location,omitempty"`
			} `tfsdk:"zip" json:"zip,omitempty"`
		} `tfsdk:"projects" json:"projects,omitempty"`
		StarterProjects *[]struct {
			Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
			Custom     *struct {
				EmbeddedResource   *map[string]string `tfsdk:"embedded_resource" json:"embeddedResource,omitempty"`
				ProjectSourceClass *string            `tfsdk:"project_source_class" json:"projectSourceClass,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Description *string `tfsdk:"description" json:"description,omitempty"`
			Git         *struct {
				CheckoutFrom *struct {
					Remote   *string `tfsdk:"remote" json:"remote,omitempty"`
					Revision *string `tfsdk:"revision" json:"revision,omitempty"`
				} `tfsdk:"checkout_from" json:"checkoutFrom,omitempty"`
				Remotes *map[string]string `tfsdk:"remotes" json:"remotes,omitempty"`
			} `tfsdk:"git" json:"git,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			SourceType *string `tfsdk:"source_type" json:"sourceType,omitempty"`
			SubDir     *string `tfsdk:"sub_dir" json:"subDir,omitempty"`
			Zip        *struct {
				Location *string `tfsdk:"location" json:"location,omitempty"`
			} `tfsdk:"zip" json:"zip,omitempty"`
		} `tfsdk:"starter_projects" json:"starterProjects,omitempty"`
		Variables *map[string]string `tfsdk:"variables" json:"variables,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *WorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_workspace_devfile_io_dev_workspace_template_v1alpha2_manifest"
}

func (r *WorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DevWorkspaceTemplate is the Schema for the devworkspacetemplates API",
		MarkdownDescription: "DevWorkspaceTemplate is the Schema for the devworkspacetemplates API",
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
				Description:         "Structure of the devworkspace. This is also the specification of a devworkspace template.",
				MarkdownDescription: "Structure of the devworkspace. This is also the specification of a devworkspace template.",
				Attributes: map[string]schema.Attribute{
					"attributes": schema.MapAttribute{
						Description:         "Map of implementation-dependant free-form YAML attributes.",
						MarkdownDescription: "Map of implementation-dependant free-form YAML attributes.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"commands": schema.ListNestedAttribute{
						Description:         "Predefined, ready-to-use, devworkspace-related commands",
						MarkdownDescription: "Predefined, ready-to-use, devworkspace-related commands",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"apply": schema.SingleNestedAttribute{
									Description:         "Command that consists in applying a given component definition, typically bound to a devworkspace event. For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'. When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
									MarkdownDescription: "Command that consists in applying a given component definition, typically bound to a devworkspace event. For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'. When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
									Attributes: map[string]schema.Attribute{
										"component": schema.StringAttribute{
											Description:         "Describes component that will be applied",
											MarkdownDescription: "Describes component that will be applied",
											Required:            true,
											Optional:            false,
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
													Required:            true,
													Optional:            false,
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
											Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
											MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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
										stringvalidator.OneOf("Exec", "Apply", "Composite", "Custom"),
									},
								},

								"composite": schema.SingleNestedAttribute{
									Description:         "Composite command that allows executing several sub-commands either sequentially or concurrently",
									MarkdownDescription: "Composite command that allows executing several sub-commands either sequentially or concurrently",
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
													Required:            true,
													Optional:            false,
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
											Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
											MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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

								"custom": schema.SingleNestedAttribute{
									Description:         "Custom command whose logic is implementation-dependant and should be provided by the user possibly through some dedicated plugin",
									MarkdownDescription: "Custom command whose logic is implementation-dependant and should be provided by the user possibly through some dedicated plugin",
									Attributes: map[string]schema.Attribute{
										"command_class": schema.StringAttribute{
											Description:         "Class of command that the associated implementation component should use to process this command with the appropriate logic",
											MarkdownDescription: "Class of command that the associated implementation component should use to process this command with the appropriate logic",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"embedded_resource": schema.MapAttribute{
											Description:         "Additional free-form configuration for this custom command that the implementation component will know how to use",
											MarkdownDescription: "Additional free-form configuration for this custom command that the implementation component will know how to use",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
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
													Required:            true,
													Optional:            false,
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
											Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
											MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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
											Description:         "The actual command-line string Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
											MarkdownDescription: "The actual command-line string Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"component": schema.StringAttribute{
											Description:         "Describes component to which given action relates",
											MarkdownDescription: "Describes component to which given action relates",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"env": schema.ListNestedAttribute{
											Description:         "Optional list of environment variables that have to be set before running the command",
											MarkdownDescription: "Optional list of environment variables that have to be set before running the command",
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
													Required:            true,
													Optional:            false,
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
											Description:         "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'. Default value is 'false'",
											MarkdownDescription: "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'. Default value is 'false'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"label": schema.StringAttribute{
											Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
											MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"working_dir": schema.StringAttribute{
											Description:         "Working directory where the command should be executed Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
											MarkdownDescription: "Working directory where the command should be executed Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
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
									Description:         "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
									MarkdownDescription: "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
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
						Description:         "List of the devworkspace components, such as editor and plugins, user-provided containers, or other types of components",
						MarkdownDescription: "List of the devworkspace components, such as editor and plugins, user-provided containers, or other types of components",
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
											Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.",
											MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"command": schema.ListAttribute{
											Description:         "The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.",
											MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.",
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
											Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod. Default value is 'false'",
											MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod. Default value is 'false'",
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
														Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exposure": schema.StringAttribute{
														Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": schema.BoolAttribute{
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_port": schema.Int64Attribute{
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
											Description:         "Environment variables used in this container. The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
											MarkdownDescription: "Environment variables used in this container. The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
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
											Description:         "Toggles whether or not the project source code should be mounted in the component. Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
											MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component. Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source_mapping": schema.StringAttribute{
											Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
											MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
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
														Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
														MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"path": schema.StringAttribute{
														Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
														MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
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
									Description:         "Custom component whose logic is implementation-dependant and should be provided by the user possibly through some dedicated controller",
									MarkdownDescription: "Custom component whose logic is implementation-dependant and should be provided by the user possibly through some dedicated controller",
									Attributes: map[string]schema.Attribute{
										"component_class": schema.StringAttribute{
											Description:         "Class of component that the associated implementation controller should use to process this command with the appropriate logic",
											MarkdownDescription: "Class of component that the associated implementation controller should use to process this command with the appropriate logic",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"embedded_resource": schema.MapAttribute{
											Description:         "Additional free-form configuration for this custom component that the implementation controller will know how to use",
											MarkdownDescription: "Additional free-form configuration for this custom component that the implementation controller will know how to use",
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
											Description:         "Defines if the image should be built during startup. Default value is 'false'",
											MarkdownDescription: "Defines if the image should be built during startup. Default value is 'false'",
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
															Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
															MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"registry_url": schema.StringAttribute{
															Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
															MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
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
																	Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																	MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
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
															Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
															MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"remotes": schema.MapAttribute{
															Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
															MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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
													Description:         "Specify if a privileged builder pod is required. Default value is 'false'",
													MarkdownDescription: "Specify if a privileged builder pod is required. Default value is 'false'",
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
													Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
													MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
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
									Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
									MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
									Attributes: map[string]schema.Attribute{
										"deploy_by_default": schema.BoolAttribute{
											Description:         "Defines if the component should be deployed during startup. Default value is 'false'",
											MarkdownDescription: "Defines if the component should be deployed during startup. Default value is 'false'",
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
														Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exposure": schema.StringAttribute{
														Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": schema.BoolAttribute{
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_port": schema.Int64Attribute{
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
									Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
									MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
									},
								},

								"openshift": schema.SingleNestedAttribute{
									Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
									MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
									Attributes: map[string]schema.Attribute{
										"deploy_by_default": schema.BoolAttribute{
											Description:         "Defines if the component should be deployed during startup. Default value is 'false'",
											MarkdownDescription: "Defines if the component should be deployed during startup. Default value is 'false'",
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
														Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
														MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exposure": schema.StringAttribute{
														Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
														MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
														Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
														MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
														},
													},

													"secure": schema.BoolAttribute{
														Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"target_port": schema.Int64Attribute{
														Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
														MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
									Description:         "Allows importing a plugin. Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
									MarkdownDescription: "Allows importing a plugin. Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
									Attributes: map[string]schema.Attribute{
										"commands": schema.ListNestedAttribute{
											Description:         "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
											MarkdownDescription: "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"apply": schema.SingleNestedAttribute{
														Description:         "Command that consists in applying a given component definition, typically bound to a devworkspace event. For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'. When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
														MarkdownDescription: "Command that consists in applying a given component definition, typically bound to a devworkspace event. For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'. When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
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
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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
														Description:         "Composite command that allows executing several sub-commands either sequentially or concurrently",
														MarkdownDescription: "Composite command that allows executing several sub-commands either sequentially or concurrently",
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
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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
																Description:         "The actual command-line string Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																MarkdownDescription: "The actual command-line string Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
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
																Description:         "Optional list of environment variables that have to be set before running the command",
																MarkdownDescription: "Optional list of environment variables that have to be set before running the command",
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
																Description:         "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'. Default value is 'false'",
																MarkdownDescription: "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'. Default value is 'false'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label": schema.StringAttribute{
																Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"working_dir": schema.StringAttribute{
																Description:         "Working directory where the command should be executed Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																MarkdownDescription: "Working directory where the command should be executed Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
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
														Description:         "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
														MarkdownDescription: "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
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
											Description:         "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
											MarkdownDescription: "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
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
																Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.",
																MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"command": schema.ListAttribute{
																Description:         "The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.",
																MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.",
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
																Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod. Default value is 'false'",
																MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod. Default value is 'false'",
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
																			Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																			MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"exposure": schema.StringAttribute{
																			Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																			MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																			Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																			MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																			},
																		},

																		"secure": schema.BoolAttribute{
																			Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																			MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"target_port": schema.Int64Attribute{
																			Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																			MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
																Description:         "Environment variables used in this container. The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
																MarkdownDescription: "Environment variables used in this container. The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
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
																Description:         "Toggles whether or not the project source code should be mounted in the component. Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component. Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"source_mapping": schema.StringAttribute{
																Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
																MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
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
																			Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																			MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtMost(63),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																			},
																		},

																		"path": schema.StringAttribute{
																			Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
																			MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
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
																Description:         "Defines if the image should be built during startup. Default value is 'false'",
																MarkdownDescription: "Defines if the image should be built during startup. Default value is 'false'",
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
																				Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																				MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"registry_url": schema.StringAttribute{
																				Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																				MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
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
																						Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																						MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
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
																				Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																				MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"remotes": schema.MapAttribute{
																				Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																				MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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
																		Description:         "Specify if a privileged builder pod is required. Default value is 'false'",
																		MarkdownDescription: "Specify if a privileged builder pod is required. Default value is 'false'",
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
																		Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
																		MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
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
														Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
														MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
														Attributes: map[string]schema.Attribute{
															"deploy_by_default": schema.BoolAttribute{
																Description:         "Defines if the component should be deployed during startup. Default value is 'false'",
																MarkdownDescription: "Defines if the component should be deployed during startup. Default value is 'false'",
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
																			Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																			MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"exposure": schema.StringAttribute{
																			Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																			MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																			Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																			MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																			},
																		},

																		"secure": schema.BoolAttribute{
																			Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																			MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"target_port": schema.Int64Attribute{
																			Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																			MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
														Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
														MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
														},
													},

													"openshift": schema.SingleNestedAttribute{
														Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
														MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
														Attributes: map[string]schema.Attribute{
															"deploy_by_default": schema.BoolAttribute{
																Description:         "Defines if the component should be deployed during startup. Default value is 'false'",
																MarkdownDescription: "Defines if the component should be deployed during startup. Default value is 'false'",
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
																			Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																			MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"exposure": schema.StringAttribute{
																			Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																			MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																			Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																			MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																			},
																		},

																		"secure": schema.BoolAttribute{
																			Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																			MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"target_port": schema.Int64Attribute{
																			Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																			MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
														Description:         "Allows specifying the definition of a volume shared by several other components",
														MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",
														Attributes: map[string]schema.Attribute{
															"ephemeral": schema.BoolAttribute{
																Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
																MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
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
											Description:         "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
											MarkdownDescription: "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"uri": schema.StringAttribute{
											Description:         "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
											MarkdownDescription: "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
											MarkdownDescription: "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
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
									Description:         "Allows specifying the definition of a volume shared by several other components",
									MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",
									Attributes: map[string]schema.Attribute{
										"ephemeral": schema.BoolAttribute{
											Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
											MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
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

					"dependent_projects": schema.ListNestedAttribute{
						Description:         "Additional projects related to the main project in the devfile, contianing names and sources locations",
						MarkdownDescription: "Additional projects related to the main project in the devfile, contianing names and sources locations",
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

								"clone_path": schema.StringAttribute{
									Description:         "Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.",
									MarkdownDescription: "Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"custom": schema.SingleNestedAttribute{
									Description:         "Project's Custom source",
									MarkdownDescription: "Project's Custom source",
									Attributes: map[string]schema.Attribute{
										"embedded_resource": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"project_source_class": schema.StringAttribute{
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

								"git": schema.SingleNestedAttribute{
									Description:         "Project's Git source",
									MarkdownDescription: "Project's Git source",
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
													Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
													MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"remotes": schema.MapAttribute{
											Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
											MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

								"name": schema.StringAttribute{
									Description:         "Project name",
									MarkdownDescription: "Project name",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
									},
								},

								"source_type": schema.StringAttribute{
									Description:         "Type of project source",
									MarkdownDescription: "Type of project source",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Git", "Zip", "Custom"),
									},
								},

								"zip": schema.SingleNestedAttribute{
									Description:         "Project's Zip source",
									MarkdownDescription: "Project's Zip source",
									Attributes: map[string]schema.Attribute{
										"location": schema.StringAttribute{
											Description:         "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
											MarkdownDescription: "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
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

					"events": schema.SingleNestedAttribute{
						Description:         "Bindings of commands to events. Each command is referred-to by its name.",
						MarkdownDescription: "Bindings of commands to events. Each command is referred-to by its name.",
						Attributes: map[string]schema.Attribute{
							"post_start": schema.ListAttribute{
								Description:         "IDs of commands that should be executed after the devworkspace is completely started. In the case of Che-Theia, these commands should be executed after all plugins and extensions have started, including project cloning. This means that those commands are not triggered until the user opens the IDE in his browser.",
								MarkdownDescription: "IDs of commands that should be executed after the devworkspace is completely started. In the case of Che-Theia, these commands should be executed after all plugins and extensions have started, including project cloning. This means that those commands are not triggered until the user opens the IDE in his browser.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"post_stop": schema.ListAttribute{
								Description:         "IDs of commands that should be executed after stopping the devworkspace.",
								MarkdownDescription: "IDs of commands that should be executed after stopping the devworkspace.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pre_start": schema.ListAttribute{
								Description:         "IDs of commands that should be executed before the devworkspace start. Kubernetes-wise, these commands would typically be executed in init containers of the devworkspace POD.",
								MarkdownDescription: "IDs of commands that should be executed before the devworkspace start. Kubernetes-wise, these commands would typically be executed in init containers of the devworkspace POD.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pre_stop": schema.ListAttribute{
								Description:         "IDs of commands that should be executed before stopping the devworkspace.",
								MarkdownDescription: "IDs of commands that should be executed before stopping the devworkspace.",
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

					"parent": schema.SingleNestedAttribute{
						Description:         "Parent devworkspace template",
						MarkdownDescription: "Parent devworkspace template",
						Attributes: map[string]schema.Attribute{
							"attributes": schema.MapAttribute{
								Description:         "Overrides of attributes encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
								MarkdownDescription: "Overrides of attributes encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"commands": schema.ListNestedAttribute{
								Description:         "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
								MarkdownDescription: "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"apply": schema.SingleNestedAttribute{
											Description:         "Command that consists in applying a given component definition, typically bound to a devworkspace event. For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'. When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
											MarkdownDescription: "Command that consists in applying a given component definition, typically bound to a devworkspace event. For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'. When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
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
													Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
													MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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
											Description:         "Composite command that allows executing several sub-commands either sequentially or concurrently",
											MarkdownDescription: "Composite command that allows executing several sub-commands either sequentially or concurrently",
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
													Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
													MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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
													Description:         "The actual command-line string Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
													MarkdownDescription: "The actual command-line string Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
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
													Description:         "Optional list of environment variables that have to be set before running the command",
													MarkdownDescription: "Optional list of environment variables that have to be set before running the command",
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
													Description:         "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'. Default value is 'false'",
													MarkdownDescription: "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'. Default value is 'false'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"label": schema.StringAttribute{
													Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
													MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"working_dir": schema.StringAttribute{
													Description:         "Working directory where the command should be executed Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
													MarkdownDescription: "Working directory where the command should be executed Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
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
											Description:         "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
											MarkdownDescription: "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
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
								Description:         "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
								MarkdownDescription: "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
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
												stringvalidator.OneOf("Container", "Kubernetes", "Openshift", "Volume", "Image", "Plugin"),
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
													Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.",
													MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"command": schema.ListAttribute{
													Description:         "The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.",
													MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.",
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
													Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod. Default value is 'false'",
													MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod. Default value is 'false'",
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
																Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
													Description:         "Environment variables used in this container. The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
													MarkdownDescription: "Environment variables used in this container. The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
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
													Description:         "Toggles whether or not the project source code should be mounted in the component. Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
													MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component. Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_mapping": schema.StringAttribute{
													Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
													MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
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
																Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"path": schema.StringAttribute{
																Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
																MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
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
													Description:         "Defines if the image should be built during startup. Default value is 'false'",
													MarkdownDescription: "Defines if the image should be built during startup. Default value is 'false'",
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
																	Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																	MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"registry_url": schema.StringAttribute{
																	Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																	MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
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
																			Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																			MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
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
																	Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																	MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"remotes": schema.MapAttribute{
																	Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																	MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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
															Description:         "Specify if a privileged builder pod is required. Default value is 'false'",
															MarkdownDescription: "Specify if a privileged builder pod is required. Default value is 'false'",
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
															Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
															MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
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
											Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
											MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
											Attributes: map[string]schema.Attribute{
												"deploy_by_default": schema.BoolAttribute{
													Description:         "Defines if the component should be deployed during startup. Default value is 'false'",
													MarkdownDescription: "Defines if the component should be deployed during startup. Default value is 'false'",
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
																Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
											Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
											MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
											},
										},

										"openshift": schema.SingleNestedAttribute{
											Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
											MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
											Attributes: map[string]schema.Attribute{
												"deploy_by_default": schema.BoolAttribute{
													Description:         "Defines if the component should be deployed during startup. Default value is 'false'",
													MarkdownDescription: "Defines if the component should be deployed during startup. Default value is 'false'",
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
																Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"exposure": schema.StringAttribute{
																Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																},
															},

															"secure": schema.BoolAttribute{
																Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"target_port": schema.Int64Attribute{
																Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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

										"plugin": schema.SingleNestedAttribute{
											Description:         "Allows importing a plugin. Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
											MarkdownDescription: "Allows importing a plugin. Plugins are mainly imported devfiles that contribute components, commands and events as a consistent single unit. They are defined in either YAML files following the devfile syntax, or as 'DevWorkspaceTemplate' Kubernetes Custom Resources",
											Attributes: map[string]schema.Attribute{
												"commands": schema.ListNestedAttribute{
													Description:         "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
													MarkdownDescription: "Overrides of commands encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"apply": schema.SingleNestedAttribute{
																Description:         "Command that consists in applying a given component definition, typically bound to a devworkspace event. For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'. When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
																MarkdownDescription: "Command that consists in applying a given component definition, typically bound to a devworkspace event. For example, when an 'apply' command is bound to a 'preStart' event, and references a 'container' component, it will start the container as a K8S initContainer in the devworkspace POD, unless the component has its 'dedicatedPod' field set to 'true'. When no 'apply' command exist for a given component, it is assumed the component will be applied at devworkspace start by default, unless 'deployByDefault' for that component is set to false.",
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
																		Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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
																Description:         "Composite command that allows executing several sub-commands either sequentially or concurrently",
																MarkdownDescription: "Composite command that allows executing several sub-commands either sequentially or concurrently",
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
																		Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
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
																		Description:         "The actual command-line string Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		MarkdownDescription: "The actual command-line string Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
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
																		Description:         "Optional list of environment variables that have to be set before running the command",
																		MarkdownDescription: "Optional list of environment variables that have to be set before running the command",
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
																		Description:         "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'. Default value is 'false'",
																		MarkdownDescription: "Specify whether the command is restarted or not when the source code changes. If set to 'true' the command won't be restarted. A *hotReloadCapable* 'run' or 'debug' command is expected to handle file changes on its own and won't be restarted. A *hotReloadCapable* 'build' command is expected to be executed only once and won't be executed again. This field is taken into account only for commands 'build', 'run' and 'debug' with 'isDefault' set to 'true'. Default value is 'false'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"label": schema.StringAttribute{
																		Description:         "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		MarkdownDescription: "Optional label that provides a label for this command to be used in Editor UI menus for example",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"working_dir": schema.StringAttribute{
																		Description:         "Working directory where the command should be executed Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
																		MarkdownDescription: "Working directory where the command should be executed Special variables that can be used: - '$PROJECTS_ROOT': A path where projects sources are mounted as defined by container component's sourceMapping. - '$PROJECT_SOURCE': A path to a project source ($PROJECTS_ROOT/<project-name>). If there are multiple projects, this will point to the directory of the first one.",
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
																Description:         "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
																MarkdownDescription: "Mandatory identifier that allows referencing this command in composite commands, from a parent, or in events.",
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
													Description:         "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
													MarkdownDescription: "Overrides of components encapsulated in a parent devfile or a plugin. Overriding is done according to K8S strategic merge patch standard rules.",
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
																		Description:         "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.",
																		MarkdownDescription: "The arguments to supply to the command running the dockerimage component. The arguments are supplied either to the default command provided in the image or to the overridden command. Defaults to an empty array, meaning use whatever is defined in the image.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"command": schema.ListAttribute{
																		Description:         "The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.",
																		MarkdownDescription: "The command to run in the dockerimage component instead of the default one provided in the image. Defaults to an empty array, meaning use whatever is defined in the image.",
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
																		Description:         "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod. Default value is 'false'",
																		MarkdownDescription: "Specify if a container should run in its own separated pod, instead of running as part of the main development environment pod. Default value is 'false'",
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
																					Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
																		Description:         "Environment variables used in this container. The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
																		MarkdownDescription: "Environment variables used in this container. The following variables are reserved and cannot be overridden via env: - '$PROJECTS_ROOT' - '$PROJECT_SOURCE'",
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
																		Description:         "Toggles whether or not the project source code should be mounted in the component. Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																		MarkdownDescription: "Toggles whether or not the project source code should be mounted in the component. Defaults to true for all component types except plugins and components that set 'dedicatedPod' to true.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"source_mapping": schema.StringAttribute{
																		Description:         "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
																		MarkdownDescription: "Optional specification of the path in the container where project sources should be transferred/mounted when 'mountSources' is 'true'. When omitted, the default value of /projects is used.",
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
																					Description:         "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																					MarkdownDescription: "The volume mount name is the name of an existing 'Volume' component. If several containers mount the same volume name then they will reuse the same volume and will be able to access to the same files.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.LengthAtMost(63),
																						stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																					},
																				},

																				"path": schema.StringAttribute{
																					Description:         "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
																					MarkdownDescription: "The path in the component container where the volume should be mounted. If not path is mentioned, default path is the is '/<name>'.",
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
																		Description:         "Defines if the image should be built during startup. Default value is 'false'",
																		MarkdownDescription: "Defines if the image should be built during startup. Default value is 'false'",
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
																						Description:         "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																						MarkdownDescription: "Id in a devfile registry that contains a Dockerfile. The src in the OCI registry required for the Dockerfile build will be downloaded for building the image.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"registry_url": schema.StringAttribute{
																						Description:         "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
																						MarkdownDescription: "Devfile Registry URL to pull the Dockerfile from when using the Devfile Registry as Dockerfile src. To ensure the Dockerfile gets resolved consistently in different environments, it is recommended to always specify the 'devfileRegistryUrl' when 'Id' is used.",
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
																								Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
																								MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
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
																						Description:         "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																						MarkdownDescription: "Location of the Dockerfile in the Git repository when using git as Dockerfile src. Defaults to Dockerfile.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"remotes": schema.MapAttribute{
																						Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
																						MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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
																				Description:         "Specify if a privileged builder pod is required. Default value is 'false'",
																				MarkdownDescription: "Specify if a privileged builder pod is required. Default value is 'false'",
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
																				Description:         "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
																				MarkdownDescription: "URI Reference of a Dockerfile. It can be a full URL or a relative URI from the current devfile as the base URI.",
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
																Description:         "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
																MarkdownDescription: "Allows importing into the devworkspace the Kubernetes resources defined in a given manifest. For example this allows reusing the Kubernetes definitions used to deploy some runtime components in production.",
																Attributes: map[string]schema.Attribute{
																	"deploy_by_default": schema.BoolAttribute{
																		Description:         "Defines if the component should be deployed during startup. Default value is 'false'",
																		MarkdownDescription: "Defines if the component should be deployed during startup. Default value is 'false'",
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
																					Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
																Description:         "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
																MarkdownDescription: "Mandatory name that allows referencing the component from other elements (such as commands) or from an external devfile that may reference this component through a parent or a plugin.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
																},
															},

															"openshift": schema.SingleNestedAttribute{
																Description:         "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
																MarkdownDescription: "Allows importing into the devworkspace the OpenShift resources defined in a given manifest. For example this allows reusing the OpenShift definitions used to deploy some runtime components in production.",
																Attributes: map[string]schema.Attribute{
																	"deploy_by_default": schema.BoolAttribute{
																		Description:         "Defines if the component should be deployed during startup. Default value is 'false'",
																		MarkdownDescription: "Defines if the component should be deployed during startup. Default value is 'false'",
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
																					Description:         "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																					MarkdownDescription: "Map of implementation-dependant string-based free-form attributes. Examples of Che-specific attributes: - cookiesAuthEnabled: 'true' / 'false', - type: 'terminal' / 'ide' / 'ide-dev',",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"exposure": schema.StringAttribute{
																					Description:         "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
																					MarkdownDescription: "Describes how the endpoint should be exposed on the network. - 'public' means that the endpoint will be exposed on the public network, typically through a K8S ingress or an OpenShift route. - 'internal' means that the endpoint will be exposed internally outside of the main devworkspace POD, typically by K8S services, to be consumed by other elements running on the same cloud internal network. - 'none' means that the endpoint will not be exposed and will only be accessible inside the main devworkspace POD, on a local address. Default value is 'public'",
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
																					Description:         "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																					MarkdownDescription: "Describes the application and transport protocols of the traffic that will go through this endpoint. - 'http': Endpoint will have 'http' traffic, typically on a TCP connection. It will be automaticaly promoted to 'https' when the 'secure' field is set to 'true'. - 'https': Endpoint will have 'https' traffic, typically on a TCP connection. - 'ws': Endpoint will have 'ws' traffic, typically on a TCP connection. It will be automaticaly promoted to 'wss' when the 'secure' field is set to 'true'. - 'wss': Endpoint will have 'wss' traffic, typically on a TCP connection. - 'tcp': Endpoint will have traffic on a TCP connection, without specifying an application protocol. - 'udp': Endpoint will have traffic on an UDP connection, without specifying an application protocol. Default value is 'http'",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																					Validators: []validator.String{
																						stringvalidator.OneOf("http", "https", "ws", "wss", "tcp", "udp"),
																					},
																				},

																				"secure": schema.BoolAttribute{
																					Description:         "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					MarkdownDescription: "Describes whether the endpoint should be secured and protected by some authentication process. This requires a protocol of 'https' or 'wss'.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"target_port": schema.Int64Attribute{
																					Description:         "Port number to be used within the container component. The same port cannot be used by two different container components.",
																					MarkdownDescription: "Port number to be used within the container component. The same port cannot be used by two different container components.",
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
																Description:         "Allows specifying the definition of a volume shared by several other components",
																MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",
																Attributes: map[string]schema.Attribute{
																	"ephemeral": schema.BoolAttribute{
																		Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
																		MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
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

												"registry_url": schema.StringAttribute{
													Description:         "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
													MarkdownDescription: "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"uri": schema.StringAttribute{
													Description:         "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
													MarkdownDescription: "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
													MarkdownDescription: "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
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
											Description:         "Allows specifying the definition of a volume shared by several other components",
											MarkdownDescription: "Allows specifying the definition of a volume shared by several other components",
											Attributes: map[string]schema.Attribute{
												"ephemeral": schema.BoolAttribute{
													Description:         "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
													MarkdownDescription: "Ephemeral volumes are not stored persistently across restarts. Defaults to false",
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

							"dependent_projects": schema.ListNestedAttribute{
								Description:         "Overrides of dependentProjects encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
								MarkdownDescription: "Overrides of dependentProjects encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
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

										"clone_path": schema.StringAttribute{
											Description:         "Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.",
											MarkdownDescription: "Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"git": schema.SingleNestedAttribute{
											Description:         "Project's Git source",
											MarkdownDescription: "Project's Git source",
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
															Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
															MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"remotes": schema.MapAttribute{
													Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
													MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

										"name": schema.StringAttribute{
											Description:         "Project name",
											MarkdownDescription: "Project name",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
											},
										},

										"source_type": schema.StringAttribute{
											Description:         "Type of project source",
											MarkdownDescription: "Type of project source",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Git", "Zip"),
											},
										},

										"zip": schema.SingleNestedAttribute{
											Description:         "Project's Zip source",
											MarkdownDescription: "Project's Zip source",
											Attributes: map[string]schema.Attribute{
												"location": schema.StringAttribute{
													Description:         "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
													MarkdownDescription: "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
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

							"projects": schema.ListNestedAttribute{
								Description:         "Overrides of projects encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
								MarkdownDescription: "Overrides of projects encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
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

										"clone_path": schema.StringAttribute{
											Description:         "Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.",
											MarkdownDescription: "Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"git": schema.SingleNestedAttribute{
											Description:         "Project's Git source",
											MarkdownDescription: "Project's Git source",
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
															Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
															MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"remotes": schema.MapAttribute{
													Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
													MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

										"name": schema.StringAttribute{
											Description:         "Project name",
											MarkdownDescription: "Project name",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
											},
										},

										"source_type": schema.StringAttribute{
											Description:         "Type of project source",
											MarkdownDescription: "Type of project source",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Git", "Zip"),
											},
										},

										"zip": schema.SingleNestedAttribute{
											Description:         "Project's Zip source",
											MarkdownDescription: "Project's Zip source",
											Attributes: map[string]schema.Attribute{
												"location": schema.StringAttribute{
													Description:         "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
													MarkdownDescription: "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
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

							"registry_url": schema.StringAttribute{
								Description:         "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
								MarkdownDescription: "Registry URL to pull the parent devfile from when using id in the parent reference. To ensure the parent devfile gets resolved consistently in different environments, it is recommended to always specify the 'registryUrl' when 'id' is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"starter_projects": schema.ListNestedAttribute{
								Description:         "Overrides of starterProjects encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
								MarkdownDescription: "Overrides of starterProjects encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
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

										"description": schema.StringAttribute{
											Description:         "Description of a starter project",
											MarkdownDescription: "Description of a starter project",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"git": schema.SingleNestedAttribute{
											Description:         "Project's Git source",
											MarkdownDescription: "Project's Git source",
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
															Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
															MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"remotes": schema.MapAttribute{
													Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
													MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

										"name": schema.StringAttribute{
											Description:         "Project name",
											MarkdownDescription: "Project name",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
											},
										},

										"source_type": schema.StringAttribute{
											Description:         "Type of project source",
											MarkdownDescription: "Type of project source",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Git", "Zip"),
											},
										},

										"sub_dir": schema.StringAttribute{
											Description:         "Sub-directory from a starter project to be used as root for starter project.",
											MarkdownDescription: "Sub-directory from a starter project to be used as root for starter project.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"zip": schema.SingleNestedAttribute{
											Description:         "Project's Zip source",
											MarkdownDescription: "Project's Zip source",
											Attributes: map[string]schema.Attribute{
												"location": schema.StringAttribute{
													Description:         "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
													MarkdownDescription: "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
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

							"uri": schema.StringAttribute{
								Description:         "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
								MarkdownDescription: "URI Reference of a parent devfile YAML file. It can be a full URL or a relative URI with the current devfile as the base URI.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"variables": schema.MapAttribute{
								Description:         "Overrides of variables encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
								MarkdownDescription: "Overrides of variables encapsulated in a parent devfile. Overriding is done according to K8S strategic merge patch standard rules.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
								MarkdownDescription: "Specific stack/sample version to pull the parent devfile from, when using id in the parent reference. To specify 'version', 'id' must be defined and used as the import reference source. 'version' can be either a specific stack version, or 'latest'. If no 'version' specified, default version will be used.",
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

					"projects": schema.ListNestedAttribute{
						Description:         "Projects worked on in the devworkspace, containing names and sources locations",
						MarkdownDescription: "Projects worked on in the devworkspace, containing names and sources locations",
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

								"clone_path": schema.StringAttribute{
									Description:         "Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.",
									MarkdownDescription: "Path relative to the root of the projects to which this project should be cloned into. This is a unix-style relative path (i.e. uses forward slashes). The path is invalid if it is absolute or tries to escape the project root through the usage of '..'. If not specified, defaults to the project name.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"custom": schema.SingleNestedAttribute{
									Description:         "Project's Custom source",
									MarkdownDescription: "Project's Custom source",
									Attributes: map[string]schema.Attribute{
										"embedded_resource": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"project_source_class": schema.StringAttribute{
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

								"git": schema.SingleNestedAttribute{
									Description:         "Project's Git source",
									MarkdownDescription: "Project's Git source",
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
													Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
													MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"remotes": schema.MapAttribute{
											Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
											MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

								"name": schema.StringAttribute{
									Description:         "Project name",
									MarkdownDescription: "Project name",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
									},
								},

								"source_type": schema.StringAttribute{
									Description:         "Type of project source",
									MarkdownDescription: "Type of project source",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Git", "Zip", "Custom"),
									},
								},

								"zip": schema.SingleNestedAttribute{
									Description:         "Project's Zip source",
									MarkdownDescription: "Project's Zip source",
									Attributes: map[string]schema.Attribute{
										"location": schema.StringAttribute{
											Description:         "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
											MarkdownDescription: "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
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

					"starter_projects": schema.ListNestedAttribute{
						Description:         "StarterProjects is a project that can be used as a starting point when bootstrapping new projects",
						MarkdownDescription: "StarterProjects is a project that can be used as a starting point when bootstrapping new projects",
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

								"custom": schema.SingleNestedAttribute{
									Description:         "Project's Custom source",
									MarkdownDescription: "Project's Custom source",
									Attributes: map[string]schema.Attribute{
										"embedded_resource": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"project_source_class": schema.StringAttribute{
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

								"description": schema.StringAttribute{
									Description:         "Description of a starter project",
									MarkdownDescription: "Description of a starter project",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"git": schema.SingleNestedAttribute{
									Description:         "Project's Git source",
									MarkdownDescription: "Project's Git source",
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
													Description:         "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
													MarkdownDescription: "The revision to checkout from. Should be branch name, tag or commit id. Default branch is used if missing or specified revision is not found.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"remotes": schema.MapAttribute{
											Description:         "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
											MarkdownDescription: "The remotes map which should be initialized in the git project. Projects must have at least one remote configured while StarterProjects & Image Component's Git source can only have at most one remote configured.",
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

								"name": schema.StringAttribute{
									Description:         "Project name",
									MarkdownDescription: "Project name",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
									},
								},

								"source_type": schema.StringAttribute{
									Description:         "Type of project source",
									MarkdownDescription: "Type of project source",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Git", "Zip", "Custom"),
									},
								},

								"sub_dir": schema.StringAttribute{
									Description:         "Sub-directory from a starter project to be used as root for starter project.",
									MarkdownDescription: "Sub-directory from a starter project to be used as root for starter project.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"zip": schema.SingleNestedAttribute{
									Description:         "Project's Zip source",
									MarkdownDescription: "Project's Zip source",
									Attributes: map[string]schema.Attribute{
										"location": schema.StringAttribute{
											Description:         "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
											MarkdownDescription: "Zip project's source location address. Should be file path of the archive, e.g. file://$FILE_PATH",
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

					"variables": schema.MapAttribute{
						Description:         "Map of key-value variables used for string replacement in the devfile. Values can be referenced via {{variable-key}} to replace the corresponding value in string fields in the devfile. Replacement cannot be used for - schemaVersion, metadata, parent source - element identifiers, e.g. command id, component name, endpoint name, project name - references to identifiers, e.g. in events, a command's component, container's volume mount name - string enums, e.g. command group kind, endpoint exposure",
						MarkdownDescription: "Map of key-value variables used for string replacement in the devfile. Values can be referenced via {{variable-key}} to replace the corresponding value in string fields in the devfile. Replacement cannot be used for - schemaVersion, metadata, parent source - element identifiers, e.g. command id, component name, endpoint name, project name - references to identifiers, e.g. in events, a command's component, container's volume mount name - string enums, e.g. command group kind, endpoint exposure",
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

func (r *WorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_workspace_devfile_io_dev_workspace_template_v1alpha2_manifest")

	var model WorkspaceDevfileIoDevWorkspaceTemplateV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("workspace.devfile.io/v1alpha2")
	model.Kind = pointer.String("DevWorkspaceTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
