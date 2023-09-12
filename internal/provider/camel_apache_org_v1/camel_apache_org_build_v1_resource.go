/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"strings"
	"time"
)

var (
	_ resource.Resource                = &CamelApacheOrgBuildV1Resource{}
	_ resource.ResourceWithConfigure   = &CamelApacheOrgBuildV1Resource{}
	_ resource.ResourceWithImportState = &CamelApacheOrgBuildV1Resource{}
)

func NewCamelApacheOrgBuildV1Resource() resource.Resource {
	return &CamelApacheOrgBuildV1Resource{}
}

type CamelApacheOrgBuildV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type CamelApacheOrgBuildV1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *CamelApacheOrgBuildV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_build_v1"
}

func (r *CamelApacheOrgBuildV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
			},

			"wait_for_upsert": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
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
								Optional:            true,
								Computed:            false,
							},

							"limit_memory": schema.StringAttribute{
								Description:         "The maximum amount of memory required. Only used for 'pod' strategy",
								MarkdownDescription: "The maximum amount of memory required. Only used for 'pod' strategy",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"operator_namespace": schema.StringAttribute{
								Description:         "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
								MarkdownDescription: "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"order_strategy": schema.StringAttribute{
								Description:         "the build order strategy to adopt",
								MarkdownDescription: "the build order strategy to adopt",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("dependencies", "fifo", "sequential"),
								},
							},

							"request_cpu": schema.StringAttribute{
								Description:         "The minimum amount of CPU required. Only used for 'pod' strategy",
								MarkdownDescription: "The minimum amount of CPU required. Only used for 'pod' strategy",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"request_memory": schema.StringAttribute{
								Description:         "The minimum amount of memory required. Only used for 'pod' strategy",
								MarkdownDescription: "The minimum amount of memory required. Only used for 'pod' strategy",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"strategy": schema.StringAttribute{
								Description:         "the strategy to adopt",
								MarkdownDescription: "the strategy to adopt",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("routine", "pod"),
								},
							},

							"tool_image": schema.StringAttribute{
								Description:         "The container image to be used to run the build.",
								MarkdownDescription: "The container image to be used to run the build.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_running_builds": schema.Int64Attribute{
						Description:         "the maximum amount of parallel running builds started by this operator instance Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						MarkdownDescription: "the maximum amount of parallel running builds started by this operator instance Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operator_namespace": schema.StringAttribute{
						Description:         "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation). Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						MarkdownDescription: "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation). Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						Required:            false,
						Optional:            true,
						Computed:            false,
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
											Optional:            true,
											Computed:            false,
										},

										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"executor_image": schema.StringAttribute{
											Description:         "docker image to use",
											MarkdownDescription: "docker image to use",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "final image name",
											MarkdownDescription: "final image name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"platform": schema.StringAttribute{
											Description:         "The platform of build image",
											MarkdownDescription: "The platform of build image",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registry": schema.SingleNestedAttribute{
											Description:         "where to publish the final image",
											MarkdownDescription: "where to publish the final image",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "the URI to access",
													MarkdownDescription: "the URI to access",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca": schema.StringAttribute{
													Description:         "the configmap which stores the Certificate Authority",
													MarkdownDescription: "the configmap which stores the Certificate Authority",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure": schema.BoolAttribute{
													Description:         "if the container registry is insecure (ie, http only)",
													MarkdownDescription: "if the container registry is insecure (ie, http only)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"organization": schema.StringAttribute{
													Description:         "the registry organization",
													MarkdownDescription: "the registry organization",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret": schema.StringAttribute{
													Description:         "the secret where credentials are stored",
													MarkdownDescription: "the secret where credentials are stored",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"verbose": schema.BoolAttribute{
											Description:         "log more information",
											MarkdownDescription: "log more information",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"builder": schema.SingleNestedAttribute{
									Description:         "a BuilderTask, used to generate and package the project",
									MarkdownDescription: "a BuilderTask, used to generate and package the project",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "the base image layer",
											MarkdownDescription: "the base image layer",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"build_dir": schema.StringAttribute{
											Description:         "workspace directory to use",
											MarkdownDescription: "workspace directory to use",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"configuration": schema.SingleNestedAttribute{
											Description:         "The configuration that should be used to perform the Build.",
											MarkdownDescription: "The configuration that should be used to perform the Build.",
											Attributes: map[string]schema.Attribute{
												"limit_cpu": schema.StringAttribute{
													Description:         "The maximum amount of CPU required. Only used for 'pod' strategy",
													MarkdownDescription: "The maximum amount of CPU required. Only used for 'pod' strategy",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"limit_memory": schema.StringAttribute{
													Description:         "The maximum amount of memory required. Only used for 'pod' strategy",
													MarkdownDescription: "The maximum amount of memory required. Only used for 'pod' strategy",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator_namespace": schema.StringAttribute{
													Description:         "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
													MarkdownDescription: "The namespace where to run the builder Pod (must be the same of the operator in charge of this Build reconciliation).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"order_strategy": schema.StringAttribute{
													Description:         "the build order strategy to adopt",
													MarkdownDescription: "the build order strategy to adopt",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("dependencies", "fifo", "sequential"),
													},
												},

												"request_cpu": schema.StringAttribute{
													Description:         "The minimum amount of CPU required. Only used for 'pod' strategy",
													MarkdownDescription: "The minimum amount of CPU required. Only used for 'pod' strategy",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"request_memory": schema.StringAttribute{
													Description:         "The minimum amount of memory required. Only used for 'pod' strategy",
													MarkdownDescription: "The minimum amount of memory required. Only used for 'pod' strategy",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"strategy": schema.StringAttribute{
													Description:         "the strategy to adopt",
													MarkdownDescription: "the strategy to adopt",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("routine", "pod"),
													},
												},

												"tool_image": schema.StringAttribute{
													Description:         "The container image to be used to run the build.",
													MarkdownDescription: "The container image to be used to run the build.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"dependencies": schema.ListAttribute{
											Description:         "the list of dependencies to use for this build",
											MarkdownDescription: "the list of dependencies to use for this build",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"cli_options": schema.ListAttribute{
													Description:         "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",
													MarkdownDescription: "The CLI options that are appended to the list of arguments for Maven commands, e.g., '-V,--no-transfer-progress,-Dstyle.color=never'. See https://maven.apache.org/ref/3.8.4/maven-embedder/cli.html.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"extension": schema.ListNestedAttribute{
													Description:         "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",
													MarkdownDescription: "The Maven build extensions. See https://maven.apache.org/guides/mini/guide-using-extensions.html.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"artifact_id": schema.StringAttribute{
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"group_id": schema.StringAttribute{
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"version": schema.StringAttribute{
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",
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

												"local_repository": schema.StringAttribute{
													Description:         "The path of the local Maven repository.",
													MarkdownDescription: "The path of the local Maven repository.",
													Required:            false,
													Optional:            true,
													Computed:            false,
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"secret_key_ref": schema.SingleNestedAttribute{
																Description:         "Selects a key of a secret.",
																MarkdownDescription: "Selects a key of a secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"properties": schema.MapAttribute{
													Description:         "The Maven properties.",
													MarkdownDescription: "The Maven properties.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"repositories": schema.ListNestedAttribute{
													Description:         "additional repositories",
													MarkdownDescription: "additional repositories",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"id": schema.StringAttribute{
																Description:         "identifies the repository",
																MarkdownDescription: "identifies the repository",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "name of the repository",
																MarkdownDescription: "name of the repository",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"releases": schema.SingleNestedAttribute{
																Description:         "can use stable releases",
																MarkdownDescription: "can use stable releases",
																Attributes: map[string]schema.Attribute{
																	"checksum_policy": schema.StringAttribute{
																		Description:         "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																		MarkdownDescription: "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"enabled": schema.BoolAttribute{
																		Description:         "is the policy activated or not",
																		MarkdownDescription: "is the policy activated or not",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"update_policy": schema.StringAttribute{
																		Description:         "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																		MarkdownDescription: "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"snapshots": schema.SingleNestedAttribute{
																Description:         "can use snapshot",
																MarkdownDescription: "can use snapshot",
																Attributes: map[string]schema.Attribute{
																	"checksum_policy": schema.StringAttribute{
																		Description:         "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																		MarkdownDescription: "When Maven deploys files to the repository, it also deploys corresponding checksum files. Your options are to 'ignore', 'fail', or 'warn' on missing or incorrect checksums.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"enabled": schema.BoolAttribute{
																		Description:         "is the policy activated or not",
																		MarkdownDescription: "is the policy activated or not",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"update_policy": schema.StringAttribute{
																		Description:         "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																		MarkdownDescription: "This element specifies how often updates should attempt to occur. Maven will compare the local POM's timestamp (stored in a repository's maven-metadata file) to the remote. The choices are: 'always', 'daily' (default), 'interval:X' (where X is an integer in minutes) or 'never'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"url": schema.StringAttribute{
																Description:         "location of the repository",
																MarkdownDescription: "location of the repository",
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
																Optional:            true,
																Computed:            false,
															},

															"id": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"password": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"username": schema.StringAttribute{
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
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

														"secret_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a secret.",
															MarkdownDescription: "Selects a key of a secret.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from.  Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

														"secret_key_ref": schema.SingleNestedAttribute{
															Description:         "Selects a key of a secret.",
															MarkdownDescription: "Selects a key of a secret.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "The key of the secret to select from.  Must be a valid secret key.",
																	MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"runtime": schema.SingleNestedAttribute{
											Description:         "the configuration required for the runtime application",
											MarkdownDescription: "the configuration required for the runtime application",
											Attributes: map[string]schema.Attribute{
												"application_class": schema.StringAttribute{
													Description:         "application entry point (main) to be executed",
													MarkdownDescription: "application entry point (main) to be executed",
													Required:            true,
													Optional:            false,
													Computed:            false,
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"group_id": schema.StringAttribute{
																		Description:         "Maven Group",
																		MarkdownDescription: "Maven Group",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"version": schema.StringAttribute{
																		Description:         "Maven Version",
																		MarkdownDescription: "Maven Version",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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

												"dependencies": schema.ListNestedAttribute{
													Description:         "list of dependencies needed to run the application",
													MarkdownDescription: "list of dependencies needed to run the application",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"artifact_id": schema.StringAttribute{
																Description:         "Maven Artifact",
																MarkdownDescription: "Maven Artifact",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"group_id": schema.StringAttribute{
																Description:         "Maven Group",
																MarkdownDescription: "Maven Group",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"version": schema.StringAttribute{
																Description:         "Maven Version",
																MarkdownDescription: "Maven Version",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"metadata": schema.MapAttribute{
													Description:         "set of metadata",
													MarkdownDescription: "set of metadata",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"provider": schema.StringAttribute{
													Description:         "Camel main application provider, ie, Camel Quarkus",
													MarkdownDescription: "Camel main application provider, ie, Camel Quarkus",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"version": schema.StringAttribute{
													Description:         "Camel K Runtime version",
													MarkdownDescription: "Camel K Runtime version",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
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
														Optional:            true,
														Computed:            false,
													},

													"content": schema.StringAttribute{
														Description:         "the source code (plain text)",
														MarkdownDescription: "the source code (plain text)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"content_key": schema.StringAttribute{
														Description:         "the confimap key holding the source content",
														MarkdownDescription: "the confimap key holding the source content",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"content_ref": schema.StringAttribute{
														Description:         "the confimap reference holding the source content",
														MarkdownDescription: "the confimap reference holding the source content",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"content_type": schema.StringAttribute{
														Description:         "the content type (tipically text or binary)",
														MarkdownDescription: "the content type (tipically text or binary)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"interceptors": schema.ListAttribute{
														Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
														MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"language": schema.StringAttribute{
														Description:         "specify which is the language (Camel DSL) used to interpret this source code",
														MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"loader": schema.StringAttribute{
														Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
														MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "the name of the specification",
														MarkdownDescription: "the name of the specification",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "the path where the file is stored",
														MarkdownDescription: "the path where the file is stored",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"property_names": schema.ListAttribute{
														Description:         "List of property names defined in the source (e.g. if type is 'template')",
														MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"raw_content": schema.StringAttribute{
														Description:         "the source code (binary)",
														MarkdownDescription: "the source code (binary)",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															validators.Base64Validator(),
														},
													},

													"type": schema.StringAttribute{
														Description:         "Type defines the kind of source described by this object",
														MarkdownDescription: "Type defines the kind of source described by this object",
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

										"steps": schema.ListAttribute{
											Description:         "the list of steps to execute (see pkg/builder/)",
											MarkdownDescription: "the list of steps to execute (see pkg/builder/)",
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

								"custom": schema.SingleNestedAttribute{
									Description:         "UserTask is used to execute any generic custom operation.",
									MarkdownDescription: "UserTask is used to execute any generic custom operation.",
									Attributes: map[string]schema.Attribute{
										"command": schema.StringAttribute{
											Description:         "the command to execute",
											MarkdownDescription: "the command to execute",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "the container image to use",
											MarkdownDescription: "the container image to use",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"jib": schema.SingleNestedAttribute{
									Description:         "a JibTask, for Jib strategy",
									MarkdownDescription: "a JibTask, for Jib strategy",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "base image layer",
											MarkdownDescription: "base image layer",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "final image name",
											MarkdownDescription: "final image name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registry": schema.SingleNestedAttribute{
											Description:         "where to publish the final image",
											MarkdownDescription: "where to publish the final image",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "the URI to access",
													MarkdownDescription: "the URI to access",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca": schema.StringAttribute{
													Description:         "the configmap which stores the Certificate Authority",
													MarkdownDescription: "the configmap which stores the Certificate Authority",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure": schema.BoolAttribute{
													Description:         "if the container registry is insecure (ie, http only)",
													MarkdownDescription: "if the container registry is insecure (ie, http only)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"organization": schema.StringAttribute{
													Description:         "the registry organization",
													MarkdownDescription: "the registry organization",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret": schema.StringAttribute{
													Description:         "the secret where credentials are stored",
													MarkdownDescription: "the secret where credentials are stored",
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

								"kaniko": schema.SingleNestedAttribute{
									Description:         "a KanikoTask, for Kaniko strategy",
									MarkdownDescription: "a KanikoTask, for Kaniko strategy",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "base image layer",
											MarkdownDescription: "base image layer",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"cache": schema.SingleNestedAttribute{
											Description:         "use a cache",
											MarkdownDescription: "use a cache",
											Attributes: map[string]schema.Attribute{
												"enabled": schema.BoolAttribute{
													Description:         "true if a cache is enabled",
													MarkdownDescription: "true if a cache is enabled",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"persistent_volume_claim": schema.StringAttribute{
													Description:         "the PVC used to store the cache",
													MarkdownDescription: "the PVC used to store the cache",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"executor_image": schema.StringAttribute{
											Description:         "docker image to use",
											MarkdownDescription: "docker image to use",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "final image name",
											MarkdownDescription: "final image name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registry": schema.SingleNestedAttribute{
											Description:         "where to publish the final image",
											MarkdownDescription: "where to publish the final image",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "the URI to access",
													MarkdownDescription: "the URI to access",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca": schema.StringAttribute{
													Description:         "the configmap which stores the Certificate Authority",
													MarkdownDescription: "the configmap which stores the Certificate Authority",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure": schema.BoolAttribute{
													Description:         "if the container registry is insecure (ie, http only)",
													MarkdownDescription: "if the container registry is insecure (ie, http only)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"organization": schema.StringAttribute{
													Description:         "the registry organization",
													MarkdownDescription: "the registry organization",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret": schema.StringAttribute{
													Description:         "the secret where credentials are stored",
													MarkdownDescription: "the secret where credentials are stored",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"verbose": schema.BoolAttribute{
											Description:         "log more information",
											MarkdownDescription: "log more information",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"s2i": schema.SingleNestedAttribute{
									Description:         "a S2iTask, for S2I strategy",
									MarkdownDescription: "a S2iTask, for S2I strategy",
									Attributes: map[string]schema.Attribute{
										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tag": schema.StringAttribute{
											Description:         "used by the ImageStream",
											MarkdownDescription: "used by the ImageStream",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"spectrum": schema.SingleNestedAttribute{
									Description:         "a SpectrumTask, for Spectrum strategy",
									MarkdownDescription: "a SpectrumTask, for Spectrum strategy",
									Attributes: map[string]schema.Attribute{
										"base_image": schema.StringAttribute{
											Description:         "base image layer",
											MarkdownDescription: "base image layer",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"context_dir": schema.StringAttribute{
											Description:         "can be useful to share info with other tasks",
											MarkdownDescription: "can be useful to share info with other tasks",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "final image name",
											MarkdownDescription: "final image name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "name of the task",
											MarkdownDescription: "name of the task",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"registry": schema.SingleNestedAttribute{
											Description:         "where to publish the final image",
											MarkdownDescription: "where to publish the final image",
											Attributes: map[string]schema.Attribute{
												"address": schema.StringAttribute{
													Description:         "the URI to access",
													MarkdownDescription: "the URI to access",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca": schema.StringAttribute{
													Description:         "the configmap which stores the Certificate Authority",
													MarkdownDescription: "the configmap which stores the Certificate Authority",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure": schema.BoolAttribute{
													Description:         "if the container registry is insecure (ie, http only)",
													MarkdownDescription: "if the container registry is insecure (ie, http only)",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"organization": schema.StringAttribute{
													Description:         "the registry organization",
													MarkdownDescription: "the registry organization",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret": schema.StringAttribute{
													Description:         "the secret where credentials are stored",
													MarkdownDescription: "the secret where credentials are stored",
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

					"timeout": schema.StringAttribute{
						Description:         "Timeout defines the Build maximum execution duration. The Build deadline is set to the Build start time plus the Timeout duration. If the Build deadline is exceeded, the Build context is canceled, and its phase set to BuildPhaseFailed.",
						MarkdownDescription: "Timeout defines the Build maximum execution duration. The Build deadline is set to the Build start time plus the Timeout duration. If the Build deadline is exceeded, the Build context is canceled, and its phase set to BuildPhaseFailed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tool_image": schema.StringAttribute{
						Description:         "The container image to be used to run the build. Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
						MarkdownDescription: "The container image to be used to run the build. Deprecated: no longer in use in Camel K 2 - maintained for backward compatibility",
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

func (r *CamelApacheOrgBuildV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *CamelApacheOrgBuildV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_camel_apache_org_build_v1")

	var model CamelApacheOrgBuildV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("camel.apache.org/v1")
	model.Kind = pointer.String("Build")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "builds"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CamelApacheOrgBuildV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CamelApacheOrgBuildV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_build_v1")

	var data CamelApacheOrgBuildV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "builds"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CamelApacheOrgBuildV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *CamelApacheOrgBuildV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_camel_apache_org_build_v1")

	var model CamelApacheOrgBuildV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("camel.apache.org/v1")
	model.Kind = pointer.String("Build")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "builds"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse CamelApacheOrgBuildV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *CamelApacheOrgBuildV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_camel_apache_org_build_v1")

	var data CamelApacheOrgBuildV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "builds"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "builds"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *CamelApacheOrgBuildV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
