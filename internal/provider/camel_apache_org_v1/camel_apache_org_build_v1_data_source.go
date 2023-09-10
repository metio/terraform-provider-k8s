/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &CamelApacheOrgBuildV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CamelApacheOrgBuildV1DataSource{}
)

func NewCamelApacheOrgBuildV1DataSource() datasource.DataSource {
	return &CamelApacheOrgBuildV1DataSource{}
}

type CamelApacheOrgBuildV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CamelApacheOrgBuildV1DataSourceData struct {
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
		Configuration *struct {
			LimitCPU          *string `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
			LimitMemory       *string `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
			OperatorNamespace *string `tfsdk:"operator_namespace" json:"operatorNamespace,omitempty"`
			OrderStrategy     *string `tfsdk:"order_strategy" json:"orderStrategy,omitempty"`
			RequestCPU        *string `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
			RequestMemory     *string `tfsdk:"request_memory" json:"requestMemory,omitempty"`
			Strategy          *string `tfsdk:"strategy" json:"strategy,omitempty"`
			ToolImage         *string `tfsdk:"tool_image" json:"toolImage,omitempty"`
		} `tfsdk:"configuration" json:"configuration,omitempty"`
		MaxRunningBuilds  *int64  `tfsdk:"max_running_builds" json:"maxRunningBuilds,omitempty"`
		OperatorNamespace *string `tfsdk:"operator_namespace" json:"operatorNamespace,omitempty"`
		Tasks             *[]struct {
			Buildah *struct {
				BaseImage     *string `tfsdk:"base_image" json:"baseImage,omitempty"`
				ContextDir    *string `tfsdk:"context_dir" json:"contextDir,omitempty"`
				ExecutorImage *string `tfsdk:"executor_image" json:"executorImage,omitempty"`
				Image         *string `tfsdk:"image" json:"image,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				Platform      *string `tfsdk:"platform" json:"platform,omitempty"`
				Registry      *struct {
					Address      *string `tfsdk:"address" json:"address,omitempty"`
					Ca           *string `tfsdk:"ca" json:"ca,omitempty"`
					Insecure     *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					Organization *string `tfsdk:"organization" json:"organization,omitempty"`
					Secret       *string `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"registry" json:"registry,omitempty"`
				Verbose *bool `tfsdk:"verbose" json:"verbose,omitempty"`
			} `tfsdk:"buildah" json:"buildah,omitempty"`
			Builder *struct {
				BaseImage     *string `tfsdk:"base_image" json:"baseImage,omitempty"`
				BuildDir      *string `tfsdk:"build_dir" json:"buildDir,omitempty"`
				Configuration *struct {
					LimitCPU          *string `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
					LimitMemory       *string `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
					OperatorNamespace *string `tfsdk:"operator_namespace" json:"operatorNamespace,omitempty"`
					OrderStrategy     *string `tfsdk:"order_strategy" json:"orderStrategy,omitempty"`
					RequestCPU        *string `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
					RequestMemory     *string `tfsdk:"request_memory" json:"requestMemory,omitempty"`
					Strategy          *string `tfsdk:"strategy" json:"strategy,omitempty"`
					ToolImage         *string `tfsdk:"tool_image" json:"toolImage,omitempty"`
				} `tfsdk:"configuration" json:"configuration,omitempty"`
				Dependencies *[]string `tfsdk:"dependencies" json:"dependencies,omitempty"`
				Maven        *struct {
					CaSecrets *[]struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"ca_secrets" json:"caSecrets,omitempty"`
					CliOptions *[]string `tfsdk:"cli_options" json:"cliOptions,omitempty"`
					Extension  *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
						GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
						Version    *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"extension" json:"extension,omitempty"`
					LocalRepository *string `tfsdk:"local_repository" json:"localRepository,omitempty"`
					Profiles        *[]struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"profiles" json:"profiles,omitempty"`
					Properties   *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
					Repositories *[]struct {
						Id       *string `tfsdk:"id" json:"id,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Releases *struct {
							ChecksumPolicy *string `tfsdk:"checksum_policy" json:"checksumPolicy,omitempty"`
							Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
							UpdatePolicy   *string `tfsdk:"update_policy" json:"updatePolicy,omitempty"`
						} `tfsdk:"releases" json:"releases,omitempty"`
						Snapshots *struct {
							ChecksumPolicy *string `tfsdk:"checksum_policy" json:"checksumPolicy,omitempty"`
							Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
							UpdatePolicy   *string `tfsdk:"update_policy" json:"updatePolicy,omitempty"`
						} `tfsdk:"snapshots" json:"snapshots,omitempty"`
						Url *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"repositories" json:"repositories,omitempty"`
					Servers *[]struct {
						Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
						Id            *string            `tfsdk:"id" json:"id,omitempty"`
						Password      *string            `tfsdk:"password" json:"password,omitempty"`
						Username      *string            `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"servers" json:"servers,omitempty"`
					Settings *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"settings" json:"settings,omitempty"`
					SettingsSecurity *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						SecretKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"settings_security" json:"settingsSecurity,omitempty"`
				} `tfsdk:"maven" json:"maven,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Runtime *struct {
					ApplicationClass *string `tfsdk:"application_class" json:"applicationClass,omitempty"`
					Capabilities     *struct {
						Dependencies *[]struct {
							ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
							GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
							Version    *string `tfsdk:"version" json:"version,omitempty"`
						} `tfsdk:"dependencies" json:"dependencies,omitempty"`
					} `tfsdk:"capabilities" json:"capabilities,omitempty"`
					Dependencies *[]struct {
						ArtifactId *string `tfsdk:"artifact_id" json:"artifactId,omitempty"`
						GroupId    *string `tfsdk:"group_id" json:"groupId,omitempty"`
						Version    *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"dependencies" json:"dependencies,omitempty"`
					Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
					Provider *string            `tfsdk:"provider" json:"provider,omitempty"`
					Version  *string            `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"runtime" json:"runtime,omitempty"`
				Sources *[]struct {
					Compression    *bool     `tfsdk:"compression" json:"compression,omitempty"`
					Content        *string   `tfsdk:"content" json:"content,omitempty"`
					ContentKey     *string   `tfsdk:"content_key" json:"contentKey,omitempty"`
					ContentRef     *string   `tfsdk:"content_ref" json:"contentRef,omitempty"`
					ContentType    *string   `tfsdk:"content_type" json:"contentType,omitempty"`
					Interceptors   *[]string `tfsdk:"interceptors" json:"interceptors,omitempty"`
					Language       *string   `tfsdk:"language" json:"language,omitempty"`
					Loader         *string   `tfsdk:"loader" json:"loader,omitempty"`
					Name           *string   `tfsdk:"name" json:"name,omitempty"`
					Path           *string   `tfsdk:"path" json:"path,omitempty"`
					Property_names *[]string `tfsdk:"property_names" json:"property-names,omitempty"`
					RawContent     *string   `tfsdk:"raw_content" json:"rawContent,omitempty"`
					Type           *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"sources" json:"sources,omitempty"`
				Steps *[]string `tfsdk:"steps" json:"steps,omitempty"`
			} `tfsdk:"builder" json:"builder,omitempty"`
			Custom *struct {
				Command *string `tfsdk:"command" json:"command,omitempty"`
				Image   *string `tfsdk:"image" json:"image,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"custom" json:"custom,omitempty"`
			Jib *struct {
				BaseImage  *string `tfsdk:"base_image" json:"baseImage,omitempty"`
				ContextDir *string `tfsdk:"context_dir" json:"contextDir,omitempty"`
				Image      *string `tfsdk:"image" json:"image,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Registry   *struct {
					Address      *string `tfsdk:"address" json:"address,omitempty"`
					Ca           *string `tfsdk:"ca" json:"ca,omitempty"`
					Insecure     *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					Organization *string `tfsdk:"organization" json:"organization,omitempty"`
					Secret       *string `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"registry" json:"registry,omitempty"`
			} `tfsdk:"jib" json:"jib,omitempty"`
			Kaniko *struct {
				BaseImage *string `tfsdk:"base_image" json:"baseImage,omitempty"`
				Cache     *struct {
					Enabled               *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					PersistentVolumeClaim *string `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
				} `tfsdk:"cache" json:"cache,omitempty"`
				ContextDir    *string `tfsdk:"context_dir" json:"contextDir,omitempty"`
				ExecutorImage *string `tfsdk:"executor_image" json:"executorImage,omitempty"`
				Image         *string `tfsdk:"image" json:"image,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				Registry      *struct {
					Address      *string `tfsdk:"address" json:"address,omitempty"`
					Ca           *string `tfsdk:"ca" json:"ca,omitempty"`
					Insecure     *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					Organization *string `tfsdk:"organization" json:"organization,omitempty"`
					Secret       *string `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"registry" json:"registry,omitempty"`
				Verbose *bool `tfsdk:"verbose" json:"verbose,omitempty"`
			} `tfsdk:"kaniko" json:"kaniko,omitempty"`
			S2i *struct {
				ContextDir *string `tfsdk:"context_dir" json:"contextDir,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Tag        *string `tfsdk:"tag" json:"tag,omitempty"`
			} `tfsdk:"s2i" json:"s2i,omitempty"`
			Spectrum *struct {
				BaseImage  *string `tfsdk:"base_image" json:"baseImage,omitempty"`
				ContextDir *string `tfsdk:"context_dir" json:"contextDir,omitempty"`
				Image      *string `tfsdk:"image" json:"image,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Registry   *struct {
					Address      *string `tfsdk:"address" json:"address,omitempty"`
					Ca           *string `tfsdk:"ca" json:"ca,omitempty"`
					Insecure     *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					Organization *string `tfsdk:"organization" json:"organization,omitempty"`
					Secret       *string `tfsdk:"secret" json:"secret,omitempty"`
				} `tfsdk:"registry" json:"registry,omitempty"`
			} `tfsdk:"spectrum" json:"spectrum,omitempty"`
		} `tfsdk:"tasks" json:"tasks,omitempty"`
		Timeout   *string `tfsdk:"timeout" json:"timeout,omitempty"`
		ToolImage *string `tfsdk:"tool_image" json:"toolImage,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgBuildV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_build_v1"
}

func (r *CamelApacheOrgBuildV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Build is the Schema for the builds API.",
		MarkdownDescription: "Build is the Schema for the builds API.",
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
				Description:         "BuildSpec defines the list of tasks to be execute for a Build. From Camel K version 2, it would be more appropriate to think it as pipeline.",
				MarkdownDescription: "BuildSpec defines the list of tasks to be execute for a Build. From Camel K version 2, it would be more appropriate to think it as pipeline.",
				Attributes: map[string]schema.Attribute{
					"configuration": schema.SingleNestedAttribute{
						Description:         "The configuration that should be used to perform the Build. Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						MarkdownDescription: "The configuration that should be used to perform the Build. Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						Attributes: map[string]schema.Attribute{
							"limit_cpu": schema.StringAttribute{
								Description:         "The maximum amount of CPU required. Only used for 'pod' strategy",
								MarkdownDescription: "The maximum amount of CPU required. Only used for 'pod' strategy",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"limit_memory": schema.StringAttribute{
								Description:         "The maximum amount of memory required. Only used for 'pod' strategy",
								MarkdownDescription: "The maximum amount of memory required. Only used for 'pod' strategy",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"operator_namespace": schema.StringAttribute{
								Description:         "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
								MarkdownDescription: "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"order_strategy": schema.StringAttribute{
								Description:         "the build order strategy to adopt",
								MarkdownDescription: "the build order strategy to adopt",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"request_cpu": schema.StringAttribute{
								Description:         "The minimum amount of CPU required. Only used for 'pod' strategy",
								MarkdownDescription: "The minimum amount of CPU required. Only used for 'pod' strategy",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"request_memory": schema.StringAttribute{
								Description:         "The minimum amount of memory required. Only used for 'pod' strategy",
								MarkdownDescription: "The minimum amount of memory required. Only used for 'pod' strategy",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"strategy": schema.StringAttribute{
								Description:         "the strategy to adopt",
								MarkdownDescription: "the strategy to adopt",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tool_image": schema.StringAttribute{
								Description:         "The container image to be used to run the build.",
								MarkdownDescription: "The container image to be used to run the build.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"max_running_builds": schema.Int64Attribute{
						Description:         "the maximum amount of parallel running builds started by this operator instance Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						MarkdownDescription: "the maximum amount of parallel running builds started by this operator instance Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"operator_namespace": schema.StringAttribute{
						Description:         "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation). Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						MarkdownDescription: "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation). Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tasks": schema.ListNestedAttribute{
						Description:         "The sequence of tasks (pipeline) to be performed.",
						MarkdownDescription: "The sequence of tasks (pipeline) to be performed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"buildah": schema.SingleNestedAttribute{
									Description:         "a BuildahTask, for Buildah strategy",
									MarkdownDescription: "a BuildahTask, for Buildah strategy",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "base image layer",
											MarkdownDescription: "base image layer",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"executor_image": schema.StringAttribute{
											Description:         "docker image to use",
											MarkdownDescription: "docker image to use",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image": schema.StringAttribute{
											Description:         "final image name",
											MarkdownDescription: "final image name",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"platform": schema.StringAttribute{
											Description:         "The platform of build image",
											MarkdownDescription: "The platform of build image",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"registry": schema.SingleNestedAttribute{
											Description:         "where to publish the final image",
											MarkdownDescription: "where to publish the final image",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "the URI to access",
													MarkdownDescription: "the URI to access",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ca": schema.StringAttribute{
													Description:         "the configmap which stores the Certificate Authority",
													MarkdownDescription: "the configmap which stores the Certificate Authority",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"insecure": schema.BoolAttribute{
													Description:         "if the container registry is insecure (ie, http only)",
													MarkdownDescription: "if the container registry is insecure (ie, http only)",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"organization": schema.StringAttribute{
													Description:         "the registry organization",
													MarkdownDescription: "the registry organization",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret": schema.StringAttribute{
													Description:         "the secret where credentials are stored",
													MarkdownDescription: "the secret where credentials are stored",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"verbose": schema.BoolAttribute{
											Description:         "log more information",
											MarkdownDescription: "log more information",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"builder": schema.SingleNestedAttribute{
									Description:         "a BuilderTask, used to generate and package the project",
									MarkdownDescription: "a BuilderTask, used to generate and package the project",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "the base image layer",
											MarkdownDescription: "the base image layer",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"build_dir": schema.StringAttribute{
											Description:         "workspace directory to use",
											MarkdownDescription: "workspace directory to use",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"configuration": schema.SingleNestedAttribute{
											Description:         "The configuration that should be used to perform the Build.",
											MarkdownDescription: "The configuration that should be used to perform the Build.",
											Attributes: map[string]schema.Attribute{
												"limit_cpu": schema.StringAttribute{
													Description:         "The maximum amount of CPU required. Only used for 'pod' strategy",
													MarkdownDescription: "The maximum amount of CPU required. Only used for 'pod' strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"limit_memory": schema.StringAttribute{
													Description:         "The maximum amount of memory required. Only used for 'pod' strategy",
													MarkdownDescription: "The maximum amount of memory required. Only used for 'pod' strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator_namespace": schema.StringAttribute{
													Description:         "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
													MarkdownDescription: "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"order_strategy": schema.StringAttribute{
													Description:         "the build order strategy to adopt",
													MarkdownDescription: "the build order strategy to adopt",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"request_cpu": schema.StringAttribute{
													Description:         "The minimum amount of CPU required. Only used for 'pod' strategy",
													MarkdownDescription: "The minimum amount of CPU required. Only used for 'pod' strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"request_memory": schema.StringAttribute{
													Description:         "The minimum amount of memory required. Only used for 'pod' strategy",
													MarkdownDescription: "The minimum amount of memory required. Only used for 'pod' strategy",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"strategy": schema.StringAttribute{
													Description:         "the strategy to adopt",
													MarkdownDescription: "the strategy to adopt",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tool_image": schema.StringAttribute{
													Description:         "The container image to be used to run the build.",
													MarkdownDescription: "The container image to be used to run the build.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"dependencies": schema.ListAttribute{
											Description:         "the list of dependencies to use for this build",
											MarkdownDescription: "the list of dependencies to use for this build",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"maven": schema.SingleNestedAttribute{
											Description:         "the configuration required by Maven for the application build phase",
											MarkdownDescription: "the configuration required by Maven for the application build phase",
											Attributes: map[string]schema.Attribute{
												"ca_secrets": schema.ListNestedAttribute{
													Description:         "The Secrets name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",
													MarkdownDescription: "The Secrets name and key, containing the CA certificate(s) used to connect to remote Maven repositories. It can contain X.509 certificates, and PKCS#7 formatted certificate chains. A JKS formatted keystore is automatically created to store the CA certificate(s), and configured to be used as a trusted certificate(s) by the Maven commands. Note that the root CA certificates are also imported into the created keystore.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the Secret or its key must be defined",
																MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

												"cli_options": schema.ListAttribute{
													Description:         "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",
													MarkdownDescription: "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"extension": schema.ListNestedAttribute{
													Description:         "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",
													MarkdownDescription: "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"artifact_id": schema.StringAttribute{
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"group_id": schema.StringAttribute{
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"version": schema.StringAttribute{
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",
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

												"local_repository": schema.StringAttribute{
													Description:         "The path of the local Maven repository.",
													MarkdownDescription: "The path of the local Maven repository.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"profiles": schema.ListNestedAttribute{
													Description:         "A reference to the ConfigMap or Secret key that contains the Maven profile.",
													MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the Maven profile.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config_map_key_ref": schema.SingleNestedAttribute{
																Description:         "Selects a key of a ConfigMap.",
																MarkdownDescription: "Selects a key of a ConfigMap.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"secret_key_ref": schema.SingleNestedAttribute{
																Description:         "Selects a key of a secret.",
																MarkdownDescription: "Selects a key of a secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"properties": schema.MapAttribute{
													Description:         "The Maven properties.",
													MarkdownDescription: "The Maven properties.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"repositories": schema.ListNestedAttribute{
													Description:         "additional repositories",
													MarkdownDescription: "additional repositories",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"id": schema.StringAttribute{
																Description:         "identifies the repository",
																MarkdownDescription: "identifies the repository",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"name": schema.StringAttribute{
																Description:         "name of the repository",
																MarkdownDescription: "name of the repository",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"releases": schema.SingleNestedAttribute{
																Description:         "can use stable releases",
																MarkdownDescription: "can use stable releases",
																Attributes: map[string]schema.Attribute{
																	"checksum_policy": schema.StringAttribute{
																		Description:         "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																		MarkdownDescription: "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"enabled": schema.BoolAttribute{
																		Description:         "is the policy activated or not",
																		MarkdownDescription: "is the policy activated or not",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"update_policy": schema.StringAttribute{
																		Description:         "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																		MarkdownDescription: "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"snapshots": schema.SingleNestedAttribute{
																Description:         "can use snapshot",
																MarkdownDescription: "can use snapshot",
																Attributes: map[string]schema.Attribute{
																	"checksum_policy": schema.StringAttribute{
																		Description:         "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																		MarkdownDescription: "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"enabled": schema.BoolAttribute{
																		Description:         "is the policy activated or not",
																		MarkdownDescription: "is the policy activated or not",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"update_policy": schema.StringAttribute{
																		Description:         "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																		MarkdownDescription: "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"url": schema.StringAttribute{
																Description:         "location of the repository",
																MarkdownDescription: "location of the repository",
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

												"servers": schema.ListNestedAttribute{
													Description:         "Servers (auth)",
													MarkdownDescription: "Servers (auth)",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"configuration": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"password": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"username": schema.StringAttribute{
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

												"settings": schema.SingleNestedAttribute{
													Description:         "A reference to the ConfigMap or Secret key that contains the Maven settings.",
													MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the Maven settings.",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a ConfigMap.",
															MarkdownDescription: "Selects a key of a ConfigMap.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key to select.",
																	MarkdownDescription: "The key to select.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the ConfigMap or its key must be defined",
																	MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"secret_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a secret.",
															MarkdownDescription: "Selects a key of a secret.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from.  Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the Secret or its key must be defined",
																	MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

												"settings_security": schema.SingleNestedAttribute{
													Description:         "A reference to the ConfigMap or Secret key that contains the security of the Maven settings.",
													MarkdownDescription: "A reference to the ConfigMap or Secret key that contains the security of the Maven settings.",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a ConfigMap.",
															MarkdownDescription: "Selects a key of a ConfigMap.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key to select.",
																	MarkdownDescription: "The key to select.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the ConfigMap or its key must be defined",
																	MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"secret_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a secret.",
															MarkdownDescription: "Selects a key of a secret.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from.  Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"optional": schema.BoolAttribute{
																	Description:         "Specify whether the Secret or its key must be defined",
																	MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"runtime": schema.SingleNestedAttribute{
											Description:         "the configuration required for the runtime application",
											MarkdownDescription: "the configuration required for the runtime application",
											Attributes: map[string]schema.Attribute{
												"application_class": schema.StringAttribute{
													Description:         "application entry point (main) to be executed",
													MarkdownDescription: "application entry point (main) to be executed",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"capabilities": schema.SingleNestedAttribute{
													Description:         "features offered by this runtime",
													MarkdownDescription: "features offered by this runtime",
													Attributes: map[string]schema.Attribute{
														"dependencies": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"artifact_id": schema.StringAttribute{
																		Description:         "Maven Artifact",
																		MarkdownDescription: "Maven Artifact",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"group_id": schema.StringAttribute{
																		Description:         "Maven Group",
																		MarkdownDescription: "Maven Group",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"version": schema.StringAttribute{
																		Description:         "Maven Version",
																		MarkdownDescription: "Maven Version",
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

												"dependencies": schema.ListNestedAttribute{
													Description:         "list of dependencies needed to run the application",
													MarkdownDescription: "list of dependencies needed to run the application",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"artifact_id": schema.StringAttribute{
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"group_id": schema.StringAttribute{
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"version": schema.StringAttribute{
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",
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

												"metadata": schema.MapAttribute{
													Description:         "set of metadata",
													MarkdownDescription: "set of metadata",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"provider": schema.StringAttribute{
													Description:         "Camel main application provider, ie, Camel Quarkus",
													MarkdownDescription: "Camel main application provider, ie, Camel Quarkus",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"version": schema.StringAttribute{
													Description:         "Camel K Runtime version",
													MarkdownDescription: "Camel K Runtime version",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"sources": schema.ListNestedAttribute{
											Description:         "the sources to add at build time",
											MarkdownDescription: "the sources to add at build time",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"compression": schema.BoolAttribute{
														Description:         "if the content is compressed (base64 encrypted)",
														MarkdownDescription: "if the content is compressed (base64 encrypted)",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"content": schema.StringAttribute{
														Description:         "the source code (plain text)",
														MarkdownDescription: "the source code (plain text)",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"content_key": schema.StringAttribute{
														Description:         "the confimap key holding the source content",
														MarkdownDescription: "the confimap key holding the source content",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"content_ref": schema.StringAttribute{
														Description:         "the confimap reference holding the source content",
														MarkdownDescription: "the confimap reference holding the source content",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"content_type": schema.StringAttribute{
														Description:         "the content type (tipically text or binary)",
														MarkdownDescription: "the content type (tipically text or binary)",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"interceptors": schema.ListAttribute{
														Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
														MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"language": schema.StringAttribute{
														Description:         "specify which is the language (Camel DSL) used to interpret this source code",
														MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"loader": schema.StringAttribute{
														Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
														MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "the name of the specification",
														MarkdownDescription: "the name of the specification",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"path": schema.StringAttribute{
														Description:         "the path where the file is stored",
														MarkdownDescription: "the path where the file is stored",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"property_names": schema.ListAttribute{
														Description:         "List of property names defined in the source (e.g. if type is 'template')",
														MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"raw_content": schema.StringAttribute{
														Description:         "the source code (binary)",
														MarkdownDescription: "the source code (binary)",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "Type defines the kind of source described by this object",
														MarkdownDescription: "Type defines the kind of source described by this object",
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

										"steps": schema.ListAttribute{
											Description:         "the list of steps to execute (see pkg/builder/)",
											MarkdownDescription: "the list of steps to execute (see pkg/builder/)",
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

								"custom": schema.SingleNestedAttribute{
									Description:         "UserTask is used to execute any generic custom operation.",
									MarkdownDescription: "UserTask is used to execute any generic custom operation.",
									Attributes: map[string]schema.Attribute{
										"command": schema.StringAttribute{
											Description:         "the command to execute",
											MarkdownDescription: "the command to execute",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image": schema.StringAttribute{
											Description:         "the container image to use",
											MarkdownDescription: "the container image to use",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"jib": schema.SingleNestedAttribute{
									Description:         "a JibTask, for Jib strategy",
									MarkdownDescription: "a JibTask, for Jib strategy",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "base image layer",
											MarkdownDescription: "base image layer",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image": schema.StringAttribute{
											Description:         "final image name",
											MarkdownDescription: "final image name",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"registry": schema.SingleNestedAttribute{
											Description:         "where to publish the final image",
											MarkdownDescription: "where to publish the final image",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "the URI to access",
													MarkdownDescription: "the URI to access",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ca": schema.StringAttribute{
													Description:         "the configmap which stores the Certificate Authority",
													MarkdownDescription: "the configmap which stores the Certificate Authority",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"insecure": schema.BoolAttribute{
													Description:         "if the container registry is insecure (ie, http only)",
													MarkdownDescription: "if the container registry is insecure (ie, http only)",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"organization": schema.StringAttribute{
													Description:         "the registry organization",
													MarkdownDescription: "the registry organization",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret": schema.StringAttribute{
													Description:         "the secret where credentials are stored",
													MarkdownDescription: "the secret where credentials are stored",
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

								"kaniko": schema.SingleNestedAttribute{
									Description:         "a KanikoTask, for Kaniko strategy",
									MarkdownDescription: "a KanikoTask, for Kaniko strategy",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "base image layer",
											MarkdownDescription: "base image layer",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"cache": schema.SingleNestedAttribute{
											Description:         "use a cache",
											MarkdownDescription: "use a cache",
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
													Description:         "true if a cache is enabled",
													MarkdownDescription: "true if a cache is enabled",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"persistent_volume_claim": schema.StringAttribute{
													Description:         "the PVC used to store the cache",
													MarkdownDescription: "the PVC used to store the cache",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"executor_image": schema.StringAttribute{
											Description:         "docker image to use",
											MarkdownDescription: "docker image to use",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image": schema.StringAttribute{
											Description:         "final image name",
											MarkdownDescription: "final image name",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"registry": schema.SingleNestedAttribute{
											Description:         "where to publish the final image",
											MarkdownDescription: "where to publish the final image",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "the URI to access",
													MarkdownDescription: "the URI to access",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ca": schema.StringAttribute{
													Description:         "the configmap which stores the Certificate Authority",
													MarkdownDescription: "the configmap which stores the Certificate Authority",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"insecure": schema.BoolAttribute{
													Description:         "if the container registry is insecure (ie, http only)",
													MarkdownDescription: "if the container registry is insecure (ie, http only)",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"organization": schema.StringAttribute{
													Description:         "the registry organization",
													MarkdownDescription: "the registry organization",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret": schema.StringAttribute{
													Description:         "the secret where credentials are stored",
													MarkdownDescription: "the secret where credentials are stored",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"verbose": schema.BoolAttribute{
											Description:         "log more information",
											MarkdownDescription: "log more information",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"s2i": schema.SingleNestedAttribute{
									Description:         "a S2iTask, for S2I strategy",
									MarkdownDescription: "a S2iTask, for S2I strategy",
									Attributes: map[string]schema.Attribute{
										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tag": schema.StringAttribute{
											Description:         "used by the ImageStream",
											MarkdownDescription: "used by the ImageStream",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"spectrum": schema.SingleNestedAttribute{
									Description:         "a SpectrumTask, for Spectrum strategy",
									MarkdownDescription: "a SpectrumTask, for Spectrum strategy",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "base image layer",
											MarkdownDescription: "base image layer",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image": schema.StringAttribute{
											Description:         "final image name",
											MarkdownDescription: "final image name",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"registry": schema.SingleNestedAttribute{
											Description:         "where to publish the final image",
											MarkdownDescription: "where to publish the final image",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "the URI to access",
													MarkdownDescription: "the URI to access",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ca": schema.StringAttribute{
													Description:         "the configmap which stores the Certificate Authority",
													MarkdownDescription: "the configmap which stores the Certificate Authority",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"insecure": schema.BoolAttribute{
													Description:         "if the container registry is insecure (ie, http only)",
													MarkdownDescription: "if the container registry is insecure (ie, http only)",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"organization": schema.StringAttribute{
													Description:         "the registry organization",
													MarkdownDescription: "the registry organization",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"secret": schema.StringAttribute{
													Description:         "the secret where credentials are stored",
													MarkdownDescription: "the secret where credentials are stored",
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
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"timeout": schema.StringAttribute{
						Description:         "Timeout defines the Build maximum execution duration. The Build deadline is set to the Build start time plus the Timeout duration. If the Build deadline is exceeded, the Build context is canceled, and its phase set to BuildPhaseFailed.",
						MarkdownDescription: "Timeout defines the Build maximum execution duration. The Build deadline is set to the Build start time plus the Timeout duration. If the Build deadline is exceeded, the Build context is canceled, and its phase set to BuildPhaseFailed.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tool_image": schema.StringAttribute{
						Description:         "The container image to be used to run the build. Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						MarkdownDescription: "The container image to be used to run the build. Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
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

func (r *CamelApacheOrgBuildV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CamelApacheOrgBuildV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_camel_apache_org_build_v1")

	var data CamelApacheOrgBuildV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "Build"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse CamelApacheOrgBuildV1DataSourceData
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
	data.ApiVersion = pointer.String("camel.apache.org/v1")
	data.Kind = pointer.String("Build")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
