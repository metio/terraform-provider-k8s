/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apiextensions_crossplane_io_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
)

var (
	_ resource.Resource                = &ApiextensionsCrossplaneIoCompositionRevisionV1Resource{}
	_ resource.ResourceWithConfigure   = &ApiextensionsCrossplaneIoCompositionRevisionV1Resource{}
	_ resource.ResourceWithImportState = &ApiextensionsCrossplaneIoCompositionRevisionV1Resource{}
)

func NewApiextensionsCrossplaneIoCompositionRevisionV1Resource() resource.Resource {
	return &ApiextensionsCrossplaneIoCompositionRevisionV1Resource{}
}

type ApiextensionsCrossplaneIoCompositionRevisionV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type ApiextensionsCrossplaneIoCompositionRevisionV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CompositeTypeRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
		} `tfsdk:"composite_type_ref" json:"compositeTypeRef,omitempty"`
		Environment *struct {
			DefaultData        *map[string]string `tfsdk:"default_data" json:"defaultData,omitempty"`
			EnvironmentConfigs *[]struct {
				Ref *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
				Selector *struct {
					MatchLabels *[]struct {
						FromFieldPathPolicy *string `tfsdk:"from_field_path_policy" json:"fromFieldPathPolicy,omitempty"`
						Key                 *string `tfsdk:"key" json:"key,omitempty"`
						Type                *string `tfsdk:"type" json:"type,omitempty"`
						Value               *string `tfsdk:"value" json:"value,omitempty"`
						ValueFromFieldPath  *string `tfsdk:"value_from_field_path" json:"valueFromFieldPath,omitempty"`
					} `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					MaxMatch        *int64  `tfsdk:"max_match" json:"maxMatch,omitempty"`
					Mode            *string `tfsdk:"mode" json:"mode,omitempty"`
					SortByFieldPath *string `tfsdk:"sort_by_field_path" json:"sortByFieldPath,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"environment_configs" json:"environmentConfigs,omitempty"`
			Patches *[]struct {
				Combine *struct {
					Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
					String   *struct {
						Fmt *string `tfsdk:"fmt" json:"fmt,omitempty"`
					} `tfsdk:"string" json:"string,omitempty"`
					Variables *[]struct {
						FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"combine" json:"combine,omitempty"`
				FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
				Policy        *struct {
					FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
					MergeOptions  *struct {
						AppendSlice   *bool `tfsdk:"append_slice" json:"appendSlice,omitempty"`
						KeepMapValues *bool `tfsdk:"keep_map_values" json:"keepMapValues,omitempty"`
					} `tfsdk:"merge_options" json:"mergeOptions,omitempty"`
				} `tfsdk:"policy" json:"policy,omitempty"`
				ToFieldPath *string `tfsdk:"to_field_path" json:"toFieldPath,omitempty"`
				Transforms  *[]struct {
					Convert *struct {
						Format *string `tfsdk:"format" json:"format,omitempty"`
						ToType *string `tfsdk:"to_type" json:"toType,omitempty"`
					} `tfsdk:"convert" json:"convert,omitempty"`
					Map   *map[string]string `tfsdk:"map" json:"map,omitempty"`
					Match *struct {
						FallbackTo    *string            `tfsdk:"fallback_to" json:"fallbackTo,omitempty"`
						FallbackValue *map[string]string `tfsdk:"fallback_value" json:"fallbackValue,omitempty"`
						Patterns      *[]struct {
							Literal *string            `tfsdk:"literal" json:"literal,omitempty"`
							Regexp  *string            `tfsdk:"regexp" json:"regexp,omitempty"`
							Result  *map[string]string `tfsdk:"result" json:"result,omitempty"`
							Type    *string            `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"patterns" json:"patterns,omitempty"`
					} `tfsdk:"match" json:"match,omitempty"`
					Math *struct {
						ClampMax *int64  `tfsdk:"clamp_max" json:"clampMax,omitempty"`
						ClampMin *int64  `tfsdk:"clamp_min" json:"clampMin,omitempty"`
						Multiply *int64  `tfsdk:"multiply" json:"multiply,omitempty"`
						Type     *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"math" json:"math,omitempty"`
					String *struct {
						Convert *string `tfsdk:"convert" json:"convert,omitempty"`
						Fmt     *string `tfsdk:"fmt" json:"fmt,omitempty"`
						Regexp  *struct {
							Group *int64  `tfsdk:"group" json:"group,omitempty"`
							Match *string `tfsdk:"match" json:"match,omitempty"`
						} `tfsdk:"regexp" json:"regexp,omitempty"`
						Trim *string `tfsdk:"trim" json:"trim,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"string" json:"string,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"transforms" json:"transforms,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"patches" json:"patches,omitempty"`
			Policy *struct {
				Resolution *string `tfsdk:"resolution" json:"resolution,omitempty"`
				Resolve    *string `tfsdk:"resolve" json:"resolve,omitempty"`
			} `tfsdk:"policy" json:"policy,omitempty"`
		} `tfsdk:"environment" json:"environment,omitempty"`
		Functions *[]struct {
			Config    *map[string]string `tfsdk:"config" json:"config,omitempty"`
			Container *struct {
				Image            *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				ImagePullSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
				Network *struct {
					Policy *string `tfsdk:"policy" json:"policy,omitempty"`
				} `tfsdk:"network" json:"network,omitempty"`
				Resources *struct {
					Limits *struct {
						Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
						Memory *string `tfsdk:"memory" json:"memory,omitempty"`
					} `tfsdk:"limits" json:"limits,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Runner *struct {
					Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"runner" json:"runner,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"container" json:"container,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"functions" json:"functions,omitempty"`
		PatchSets *[]struct {
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Patches *[]struct {
				Combine *struct {
					Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
					String   *struct {
						Fmt *string `tfsdk:"fmt" json:"fmt,omitempty"`
					} `tfsdk:"string" json:"string,omitempty"`
					Variables *[]struct {
						FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"combine" json:"combine,omitempty"`
				FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
				PatchSetName  *string `tfsdk:"patch_set_name" json:"patchSetName,omitempty"`
				Policy        *struct {
					FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
					MergeOptions  *struct {
						AppendSlice   *bool `tfsdk:"append_slice" json:"appendSlice,omitempty"`
						KeepMapValues *bool `tfsdk:"keep_map_values" json:"keepMapValues,omitempty"`
					} `tfsdk:"merge_options" json:"mergeOptions,omitempty"`
				} `tfsdk:"policy" json:"policy,omitempty"`
				ToFieldPath *string `tfsdk:"to_field_path" json:"toFieldPath,omitempty"`
				Transforms  *[]struct {
					Convert *struct {
						Format *string `tfsdk:"format" json:"format,omitempty"`
						ToType *string `tfsdk:"to_type" json:"toType,omitempty"`
					} `tfsdk:"convert" json:"convert,omitempty"`
					Map   *map[string]string `tfsdk:"map" json:"map,omitempty"`
					Match *struct {
						FallbackTo    *string            `tfsdk:"fallback_to" json:"fallbackTo,omitempty"`
						FallbackValue *map[string]string `tfsdk:"fallback_value" json:"fallbackValue,omitempty"`
						Patterns      *[]struct {
							Literal *string            `tfsdk:"literal" json:"literal,omitempty"`
							Regexp  *string            `tfsdk:"regexp" json:"regexp,omitempty"`
							Result  *map[string]string `tfsdk:"result" json:"result,omitempty"`
							Type    *string            `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"patterns" json:"patterns,omitempty"`
					} `tfsdk:"match" json:"match,omitempty"`
					Math *struct {
						ClampMax *int64  `tfsdk:"clamp_max" json:"clampMax,omitempty"`
						ClampMin *int64  `tfsdk:"clamp_min" json:"clampMin,omitempty"`
						Multiply *int64  `tfsdk:"multiply" json:"multiply,omitempty"`
						Type     *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"math" json:"math,omitempty"`
					String *struct {
						Convert *string `tfsdk:"convert" json:"convert,omitempty"`
						Fmt     *string `tfsdk:"fmt" json:"fmt,omitempty"`
						Regexp  *struct {
							Group *int64  `tfsdk:"group" json:"group,omitempty"`
							Match *string `tfsdk:"match" json:"match,omitempty"`
						} `tfsdk:"regexp" json:"regexp,omitempty"`
						Trim *string `tfsdk:"trim" json:"trim,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"string" json:"string,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"transforms" json:"transforms,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"patches" json:"patches,omitempty"`
		} `tfsdk:"patch_sets" json:"patchSets,omitempty"`
		PublishConnectionDetailsWithStoreConfigRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"publish_connection_details_with_store_config_ref" json:"publishConnectionDetailsWithStoreConfigRef,omitempty"`
		Resources *[]struct {
			Base              *map[string]string `tfsdk:"base" json:"base,omitempty"`
			ConnectionDetails *[]struct {
				FromConnectionSecretKey *string `tfsdk:"from_connection_secret_key" json:"fromConnectionSecretKey,omitempty"`
				FromFieldPath           *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
				Name                    *string `tfsdk:"name" json:"name,omitempty"`
				Type                    *string `tfsdk:"type" json:"type,omitempty"`
				Value                   *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"connection_details" json:"connectionDetails,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
			Patches *[]struct {
				Combine *struct {
					Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
					String   *struct {
						Fmt *string `tfsdk:"fmt" json:"fmt,omitempty"`
					} `tfsdk:"string" json:"string,omitempty"`
					Variables *[]struct {
						FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"combine" json:"combine,omitempty"`
				FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
				PatchSetName  *string `tfsdk:"patch_set_name" json:"patchSetName,omitempty"`
				Policy        *struct {
					FromFieldPath *string `tfsdk:"from_field_path" json:"fromFieldPath,omitempty"`
					MergeOptions  *struct {
						AppendSlice   *bool `tfsdk:"append_slice" json:"appendSlice,omitempty"`
						KeepMapValues *bool `tfsdk:"keep_map_values" json:"keepMapValues,omitempty"`
					} `tfsdk:"merge_options" json:"mergeOptions,omitempty"`
				} `tfsdk:"policy" json:"policy,omitempty"`
				ToFieldPath *string `tfsdk:"to_field_path" json:"toFieldPath,omitempty"`
				Transforms  *[]struct {
					Convert *struct {
						Format *string `tfsdk:"format" json:"format,omitempty"`
						ToType *string `tfsdk:"to_type" json:"toType,omitempty"`
					} `tfsdk:"convert" json:"convert,omitempty"`
					Map   *map[string]string `tfsdk:"map" json:"map,omitempty"`
					Match *struct {
						FallbackTo    *string            `tfsdk:"fallback_to" json:"fallbackTo,omitempty"`
						FallbackValue *map[string]string `tfsdk:"fallback_value" json:"fallbackValue,omitempty"`
						Patterns      *[]struct {
							Literal *string            `tfsdk:"literal" json:"literal,omitempty"`
							Regexp  *string            `tfsdk:"regexp" json:"regexp,omitempty"`
							Result  *map[string]string `tfsdk:"result" json:"result,omitempty"`
							Type    *string            `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"patterns" json:"patterns,omitempty"`
					} `tfsdk:"match" json:"match,omitempty"`
					Math *struct {
						ClampMax *int64  `tfsdk:"clamp_max" json:"clampMax,omitempty"`
						ClampMin *int64  `tfsdk:"clamp_min" json:"clampMin,omitempty"`
						Multiply *int64  `tfsdk:"multiply" json:"multiply,omitempty"`
						Type     *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"math" json:"math,omitempty"`
					String *struct {
						Convert *string `tfsdk:"convert" json:"convert,omitempty"`
						Fmt     *string `tfsdk:"fmt" json:"fmt,omitempty"`
						Regexp  *struct {
							Group *int64  `tfsdk:"group" json:"group,omitempty"`
							Match *string `tfsdk:"match" json:"match,omitempty"`
						} `tfsdk:"regexp" json:"regexp,omitempty"`
						Trim *string `tfsdk:"trim" json:"trim,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"string" json:"string,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"transforms" json:"transforms,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"patches" json:"patches,omitempty"`
			ReadinessChecks *[]struct {
				FieldPath      *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				MatchCondition *struct {
					Status *string `tfsdk:"status" json:"status,omitempty"`
					Type   *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"match_condition" json:"matchCondition,omitempty"`
				MatchInteger *int64  `tfsdk:"match_integer" json:"matchInteger,omitempty"`
				MatchString  *string `tfsdk:"match_string" json:"matchString,omitempty"`
				Type         *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"readiness_checks" json:"readinessChecks,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Revision                          *int64  `tfsdk:"revision" json:"revision,omitempty"`
		WriteConnectionSecretsToNamespace *string `tfsdk:"write_connection_secrets_to_namespace" json:"writeConnectionSecretsToNamespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apiextensions_crossplane_io_composition_revision_v1"
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A CompositionRevision represents a revision in time of a Composition. Revisions are created by Crossplane; they should be treated as immutable.",
		MarkdownDescription: "A CompositionRevision represents a revision in time of a Composition. Revisions are created by Crossplane; they should be treated as immutable.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
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
				Description:         "CompositionRevisionSpec specifies the desired state of the composition revision.",
				MarkdownDescription: "CompositionRevisionSpec specifies the desired state of the composition revision.",
				Attributes: map[string]schema.Attribute{
					"composite_type_ref": schema.SingleNestedAttribute{
						Description:         "CompositeTypeRef specifies the type of composite resource that this composition is compatible with.",
						MarkdownDescription: "CompositeTypeRef specifies the type of composite resource that this composition is compatible with.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "APIVersion of the type.",
								MarkdownDescription: "APIVersion of the type.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the type.",
								MarkdownDescription: "Kind of the type.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"environment": schema.SingleNestedAttribute{
						Description:         "Environment configures the environment in which resources are rendered.",
						MarkdownDescription: "Environment configures the environment in which resources are rendered.",
						Attributes: map[string]schema.Attribute{
							"default_data": schema.MapAttribute{
								Description:         "DefaultData statically defines the initial state of the environment. It has the same schema-less structure as the data field in environment configs. It is overwritten by the selected environment configs.",
								MarkdownDescription: "DefaultData statically defines the initial state of the environment. It has the same schema-less structure as the data field in environment configs. It is overwritten by the selected environment configs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"environment_configs": schema.ListNestedAttribute{
								Description:         "EnvironmentConfigs selects a list of 'EnvironmentConfig's. The resolved resources are stored in the composite resource at 'spec.environmentConfigRefs' and is only updated if it is null.  The list of references is used to compute an in-memory environment at compose time. The data of all object is merged in the order they are listed, meaning the values of EnvironmentConfigs with a larger index take priority over ones with smaller indices.  The computed environment can be accessed in a composition using 'FromEnvironmentFieldPath' and 'CombineFromEnvironment' patches.",
								MarkdownDescription: "EnvironmentConfigs selects a list of 'EnvironmentConfig's. The resolved resources are stored in the composite resource at 'spec.environmentConfigRefs' and is only updated if it is null.  The list of references is used to compute an in-memory environment at compose time. The data of all object is merged in the order they are listed, meaning the values of EnvironmentConfigs with a larger index take priority over ones with smaller indices.  The computed environment can be accessed in a composition using 'FromEnvironmentFieldPath' and 'CombineFromEnvironment' patches.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ref": schema.SingleNestedAttribute{
											Description:         "Ref is a named reference to a single EnvironmentConfig. Either Ref or Selector is required.",
											MarkdownDescription: "Ref is a named reference to a single EnvironmentConfig. Either Ref or Selector is required.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The name of the object.",
													MarkdownDescription: "The name of the object.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"selector": schema.SingleNestedAttribute{
											Description:         "Selector selects EnvironmentConfig(s) via labels.",
											MarkdownDescription: "Selector selects EnvironmentConfig(s) via labels.",
											Attributes: map[string]schema.Attribute{
												"match_labels": schema.ListNestedAttribute{
													Description:         "MatchLabels ensures an object with matching labels is selected.",
													MarkdownDescription: "MatchLabels ensures an object with matching labels is selected.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from_field_path_policy": schema.StringAttribute{
																Description:         "FromFieldPathPolicy specifies the policy for the valueFromFieldPath. The default is Required, meaning that an error will be returned if the field is not found in the composite resource. Optional means that if the field is not found in the composite resource, that label pair will just be skipped. N.B. other specified label matchers will still be used to retrieve the desired environment config, if any.",
																MarkdownDescription: "FromFieldPathPolicy specifies the policy for the valueFromFieldPath. The default is Required, meaning that an error will be returned if the field is not found in the composite resource. Optional means that if the field is not found in the composite resource, that label pair will just be skipped. N.B. other specified label matchers will still be used to retrieve the desired environment config, if any.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Optional", "Required"),
																},
															},

															"key": schema.StringAttribute{
																Description:         "Key of the label to match.",
																MarkdownDescription: "Key of the label to match.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type specifies where the value for a label comes from.",
																MarkdownDescription: "Type specifies where the value for a label comes from.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("FromCompositeFieldPath", "Value"),
																},
															},

															"value": schema.StringAttribute{
																Description:         "Value specifies a literal label value.",
																MarkdownDescription: "Value specifies a literal label value.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value_from_field_path": schema.StringAttribute{
																Description:         "ValueFromFieldPath specifies the field path to look for the label value.",
																MarkdownDescription: "ValueFromFieldPath specifies the field path to look for the label value.",
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

												"max_match": schema.Int64Attribute{
													Description:         "MaxMatch specifies the number of extracted EnvironmentConfigs in Multiple mode, extracts all if nil.",
													MarkdownDescription: "MaxMatch specifies the number of extracted EnvironmentConfigs in Multiple mode, extracts all if nil.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Mode specifies retrieval strategy: 'Single' or 'Multiple'.",
													MarkdownDescription: "Mode specifies retrieval strategy: 'Single' or 'Multiple'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Single", "Multiple"),
													},
												},

												"sort_by_field_path": schema.StringAttribute{
													Description:         "SortByFieldPath is the path to the field based on which list of EnvironmentConfigs is alphabetically sorted.",
													MarkdownDescription: "SortByFieldPath is the path to the field based on which list of EnvironmentConfigs is alphabetically sorted.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"type": schema.StringAttribute{
											Description:         "Type specifies the way the EnvironmentConfig is selected. Default is 'Reference'",
											MarkdownDescription: "Type specifies the way the EnvironmentConfig is selected. Default is 'Reference'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Reference", "Selector"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"patches": schema.ListNestedAttribute{
								Description:         "Patches is a list of environment patches that are executed before a composition's resources are composed.",
								MarkdownDescription: "Patches is a list of environment patches that are executed before a composition's resources are composed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"combine": schema.SingleNestedAttribute{
											Description:         "Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.",
											MarkdownDescription: "Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.",
											Attributes: map[string]schema.Attribute{
												"strategy": schema.StringAttribute{
													Description:         "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",
													MarkdownDescription: "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("string"),
													},
												},

												"string": schema.SingleNestedAttribute{
													Description:         "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",
													MarkdownDescription: "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",
													Attributes: map[string]schema.Attribute{
														"fmt": schema.StringAttribute{
															Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
															MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"variables": schema.ListNestedAttribute{
													Description:         "Variables are the list of variables whose values will be retrieved and combined.",
													MarkdownDescription: "Variables are the list of variables whose values will be retrieved and combined.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"from_field_path": schema.StringAttribute{
																Description:         "FromFieldPath is the path of the field on the source whose value is to be used as input.",
																MarkdownDescription: "FromFieldPath is the path of the field on the source whose value is to be used as input.",
																Required:            true,
																Optional:            false,
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

										"from_field_path": schema.StringAttribute{
											Description:         "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath or ToCompositeFieldPath.",
											MarkdownDescription: "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath or ToCompositeFieldPath.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"policy": schema.SingleNestedAttribute{
											Description:         "Policy configures the specifics of patching behaviour.",
											MarkdownDescription: "Policy configures the specifics of patching behaviour.",
											Attributes: map[string]schema.Attribute{
												"from_field_path": schema.StringAttribute{
													Description:         "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",
													MarkdownDescription: "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Optional", "Required"),
													},
												},

												"merge_options": schema.SingleNestedAttribute{
													Description:         "MergeOptions Specifies merge options on a field path",
													MarkdownDescription: "MergeOptions Specifies merge options on a field path",
													Attributes: map[string]schema.Attribute{
														"append_slice": schema.BoolAttribute{
															Description:         "Specifies that already existing elements in a merged slice should be preserved",
															MarkdownDescription: "Specifies that already existing elements in a merged slice should be preserved",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"keep_map_values": schema.BoolAttribute{
															Description:         "Specifies that already existing values in a merged map should be preserved",
															MarkdownDescription: "Specifies that already existing values in a merged map should be preserved",
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

										"to_field_path": schema.StringAttribute{
											Description:         "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",
											MarkdownDescription: "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"transforms": schema.ListNestedAttribute{
											Description:         "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",
											MarkdownDescription: "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"convert": schema.SingleNestedAttribute{
														Description:         "Convert is used to cast the input into the given output type.",
														MarkdownDescription: "Convert is used to cast the input into the given output type.",
														Attributes: map[string]schema.Attribute{
															"format": schema.StringAttribute{
																Description:         "The expected input format.  * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions.  If this property is null, the default conversion is applied.",
																MarkdownDescription: "The expected input format.  * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions.  If this property is null, the default conversion is applied.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "quantity", "json"),
																},
															},

															"to_type": schema.StringAttribute{
																Description:         "ToType is the type of the output of this transform.",
																MarkdownDescription: "ToType is the type of the output of this transform.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("string", "int", "int64", "bool", "float64", "object", "list"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"map": schema.MapAttribute{
														Description:         "Map uses the input as a key in the given map and returns the value.",
														MarkdownDescription: "Map uses the input as a key in the given map and returns the value.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"match": schema.SingleNestedAttribute{
														Description:         "Match is a more complex version of Map that matches a list of patterns.",
														MarkdownDescription: "Match is a more complex version of Map that matches a list of patterns.",
														Attributes: map[string]schema.Attribute{
															"fallback_to": schema.StringAttribute{
																Description:         "Determines to what value the transform should fallback if no pattern matches.",
																MarkdownDescription: "Determines to what value the transform should fallback if no pattern matches.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Value", "Input"),
																},
															},

															"fallback_value": schema.MapAttribute{
																Description:         "The fallback value that should be returned by the transform if now pattern matches.",
																MarkdownDescription: "The fallback value that should be returned by the transform if now pattern matches.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"patterns": schema.ListNestedAttribute{
																Description:         "The patterns that should be tested against the input string. Patterns are tested in order. The value of the first match is used as result of this transform.",
																MarkdownDescription: "The patterns that should be tested against the input string. Patterns are tested in order. The value of the first match is used as result of this transform.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"literal": schema.StringAttribute{
																			Description:         "Literal exactly matches the input string (case sensitive). Is required if 'type' is 'literal'.",
																			MarkdownDescription: "Literal exactly matches the input string (case sensitive). Is required if 'type' is 'literal'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"regexp": schema.StringAttribute{
																			Description:         "Regexp to match against the input string. Is required if 'type' is 'regexp'.",
																			MarkdownDescription: "Regexp to match against the input string. Is required if 'type' is 'regexp'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"result": schema.MapAttribute{
																			Description:         "The value that is used as result of the transform if the pattern matches.",
																			MarkdownDescription: "The value that is used as result of the transform if the pattern matches.",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how the pattern matches the input.  * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default.  * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
																			MarkdownDescription: "Type specifies how the pattern matches the input.  * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default.  * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("literal", "regexp"),
																			},
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

													"math": schema.SingleNestedAttribute{
														Description:         "Math is used to transform the input via mathematical operations such as multiplication.",
														MarkdownDescription: "Math is used to transform the input via mathematical operations such as multiplication.",
														Attributes: map[string]schema.Attribute{
															"clamp_max": schema.Int64Attribute{
																Description:         "ClampMax makes sure that the value is not bigger than the given value.",
																MarkdownDescription: "ClampMax makes sure that the value is not bigger than the given value.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"clamp_min": schema.Int64Attribute{
																Description:         "ClampMin makes sure that the value is not smaller than the given value.",
																MarkdownDescription: "ClampMin makes sure that the value is not smaller than the given value.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"multiply": schema.Int64Attribute{
																Description:         "Multiply the value.",
																MarkdownDescription: "Multiply the value.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type of the math transform to be run.",
																MarkdownDescription: "Type of the math transform to be run.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Multiply", "ClampMin", "ClampMax"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"string": schema.SingleNestedAttribute{
														Description:         "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",
														MarkdownDescription: "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",
														Attributes: map[string]schema.Attribute{
															"convert": schema.StringAttribute{
																Description:         "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON.",
																MarkdownDescription: "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("ToUpper", "ToLower", "ToBase64", "FromBase64", "ToJson", "ToSha1", "ToSha256", "ToSha512"),
																},
															},

															"fmt": schema.StringAttribute{
																Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"regexp": schema.SingleNestedAttribute{
																Description:         "Extract a match from the input using a regular expression.",
																MarkdownDescription: "Extract a match from the input using a regular expression.",
																Attributes: map[string]schema.Attribute{
																	"group": schema.Int64Attribute{
																		Description:         "Group number to match. 0 (the default) matches the entire expression.",
																		MarkdownDescription: "Group number to match. 0 (the default) matches the entire expression.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"match": schema.StringAttribute{
																		Description:         "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",
																		MarkdownDescription: "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"trim": schema.StringAttribute{
																Description:         "Trim the prefix or suffix from the input",
																MarkdownDescription: "Trim the prefix or suffix from the input",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "Type of the string transform to be run.",
																MarkdownDescription: "Type of the string transform to be run.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Format", "Convert", "TrimPrefix", "TrimSuffix", "Regexp"),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of the transform to be run.",
														MarkdownDescription: "Type of the transform to be run.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("map", "match", "math", "string", "convert"),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"type": schema.StringAttribute{
											Description:         "Type sets the patching behaviour to be used. Each patch type may require its own fields to be set on the Patch object.",
											MarkdownDescription: "Type sets the patching behaviour to be used. Each patch type may require its own fields to be set on the Patch object.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("FromCompositeFieldPath", "ToCompositeFieldPath", "CombineFromComposite", "CombineToComposite"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy": schema.SingleNestedAttribute{
								Description:         "Policy represents the Resolve and Resolution policies which apply to all EnvironmentSourceReferences in EnvironmentConfigs list.",
								MarkdownDescription: "Policy represents the Resolve and Resolution policies which apply to all EnvironmentSourceReferences in EnvironmentConfigs list.",
								Attributes: map[string]schema.Attribute{
									"resolution": schema.StringAttribute{
										Description:         "Resolution specifies whether resolution of this reference is required. The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.",
										MarkdownDescription: "Resolution specifies whether resolution of this reference is required. The default is 'Required', which means the reconcile will fail if the reference cannot be resolved. 'Optional' means this reference will be a no-op if it cannot be resolved.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Required", "Optional"),
										},
									},

									"resolve": schema.StringAttribute{
										Description:         "Resolve specifies when this reference should be resolved. The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.",
										MarkdownDescription: "Resolve specifies when this reference should be resolved. The default is 'IfNotPresent', which will attempt to resolve the reference only when the corresponding field is not present. Use 'Always' to resolve the reference on every reconcile.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "IfNotPresent"),
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

					"functions": schema.ListNestedAttribute{
						Description:         "Functions is list of Composition Functions that will be used when a composite resource referring to this composition is created. At least one of resources and functions must be specified. If both are specified the resources will be rendered first, then passed to the functions for further processing.",
						MarkdownDescription: "Functions is list of Composition Functions that will be used when a composite resource referring to this composition is created. At least one of resources and functions must be specified. If both are specified the resources will be rendered first, then passed to the functions for further processing.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config": schema.MapAttribute{
									Description:         "Config is an optional, arbitrary Kubernetes resource (i.e. a resource with an apiVersion and kind) that will be passed to the Composition Function as the 'config' block of its FunctionIO.",
									MarkdownDescription: "Config is an optional, arbitrary Kubernetes resource (i.e. a resource with an apiVersion and kind) that will be passed to the Composition Function as the 'config' block of its FunctionIO.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"container": schema.SingleNestedAttribute{
									Description:         "Container configuration of this function.",
									MarkdownDescription: "Container configuration of this function.",
									Attributes: map[string]schema.Attribute{
										"image": schema.StringAttribute{
											Description:         "Image specifies the OCI image in which the function is packaged. The image should include an entrypoint that reads a FunctionIO from stdin and emits it, optionally mutated, to stdout.",
											MarkdownDescription: "Image specifies the OCI image in which the function is packaged. The image should include an entrypoint that reads a FunctionIO from stdin and emits it, optionally mutated, to stdout.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"image_pull_policy": schema.StringAttribute{
											Description:         "ImagePullPolicy defines the pull policy for the function image.",
											MarkdownDescription: "ImagePullPolicy defines the pull policy for the function image.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("IfNotPresent", "Always", "Never"),
											},
										},

										"image_pull_secrets": schema.ListNestedAttribute{
											Description:         "ImagePullSecrets are used to pull images from private OCI registries.",
											MarkdownDescription: "ImagePullSecrets are used to pull images from private OCI registries.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"network": schema.SingleNestedAttribute{
											Description:         "Network configuration for the Composition Function.",
											MarkdownDescription: "Network configuration for the Composition Function.",
											Attributes: map[string]schema.Attribute{
												"policy": schema.StringAttribute{
													Description:         "Policy specifies the network policy under which the Composition Function will run. Defaults to 'Isolated' - i.e. no network access. Specify 'Runner' to allow the function the same network access as its runner.",
													MarkdownDescription: "Policy specifies the network policy under which the Composition Function will run. Defaults to 'Isolated' - i.e. no network access. Specify 'Runner' to allow the function the same network access as its runner.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Isolated", "Runner"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "Resources that may be used by the Composition Function.",
											MarkdownDescription: "Resources that may be used by the Composition Function.",
											Attributes: map[string]schema.Attribute{
												"limits": schema.SingleNestedAttribute{
													Description:         "Limits specify the maximum compute resources that may be used by the Composition Function.",
													MarkdownDescription: "Limits specify the maximum compute resources that may be used by the Composition Function.",
													Attributes: map[string]schema.Attribute{
														"cpu": schema.StringAttribute{
															Description:         "CPU, in cores. (500m = .5 cores)",
															MarkdownDescription: "CPU, in cores. (500m = .5 cores)",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"memory": schema.StringAttribute{
															Description:         "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
															MarkdownDescription: "Memory, in bytes. (500Gi = 500GiB = 500 * 1024 * 1024 * 1024)",
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

										"runner": schema.SingleNestedAttribute{
											Description:         "Runner configuration for the Composition Function.",
											MarkdownDescription: "Runner configuration for the Composition Function.",
											Attributes: map[string]schema.Attribute{
												"endpoint": schema.StringAttribute{
													Description:         "Endpoint specifies how and where Crossplane should reach the runner it uses to invoke containerized Composition Functions.",
													MarkdownDescription: "Endpoint specifies how and where Crossplane should reach the runner it uses to invoke containerized Composition Functions.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"timeout": schema.StringAttribute{
											Description:         "Timeout after which the Composition Function will be killed.",
											MarkdownDescription: "Timeout after which the Composition Function will be killed.",
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
									Description:         "Name of this function. Must be unique within its Composition.",
									MarkdownDescription: "Name of this function. Must be unique within its Composition.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "Type of this function.",
									MarkdownDescription: "Type of this function.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Container"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"patch_sets": schema.ListNestedAttribute{
						Description:         "PatchSets define a named set of patches that may be included by any resource in this Composition. PatchSets cannot themselves refer to other PatchSets.",
						MarkdownDescription: "PatchSets define a named set of patches that may be included by any resource in this Composition. PatchSets cannot themselves refer to other PatchSets.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of this PatchSet.",
									MarkdownDescription: "Name of this PatchSet.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"patches": schema.ListNestedAttribute{
									Description:         "Patches will be applied as an overlay to the base resource.",
									MarkdownDescription: "Patches will be applied as an overlay to the base resource.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"combine": schema.SingleNestedAttribute{
												Description:         "Combine is the patch configuration for a CombineFromComposite, CombineFromEnvironment, CombineToComposite or CombineToEnvironment patch.",
												MarkdownDescription: "Combine is the patch configuration for a CombineFromComposite, CombineFromEnvironment, CombineToComposite or CombineToEnvironment patch.",
												Attributes: map[string]schema.Attribute{
													"strategy": schema.StringAttribute{
														Description:         "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",
														MarkdownDescription: "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("string"),
														},
													},

													"string": schema.SingleNestedAttribute{
														Description:         "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",
														MarkdownDescription: "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",
														Attributes: map[string]schema.Attribute{
															"fmt": schema.StringAttribute{
																Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"variables": schema.ListNestedAttribute{
														Description:         "Variables are the list of variables whose values will be retrieved and combined.",
														MarkdownDescription: "Variables are the list of variables whose values will be retrieved and combined.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"from_field_path": schema.StringAttribute{
																	Description:         "FromFieldPath is the path of the field on the source whose value is to be used as input.",
																	MarkdownDescription: "FromFieldPath is the path of the field on the source whose value is to be used as input.",
																	Required:            true,
																	Optional:            false,
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

											"from_field_path": schema.StringAttribute{
												Description:         "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath, FromEnvironmentFieldPath, ToCompositeFieldPath, ToEnvironmentFieldPath.",
												MarkdownDescription: "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath, FromEnvironmentFieldPath, ToCompositeFieldPath, ToEnvironmentFieldPath.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"patch_set_name": schema.StringAttribute{
												Description:         "PatchSetName to include patches from. Required when type is PatchSet.",
												MarkdownDescription: "PatchSetName to include patches from. Required when type is PatchSet.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"policy": schema.SingleNestedAttribute{
												Description:         "Policy configures the specifics of patching behaviour.",
												MarkdownDescription: "Policy configures the specifics of patching behaviour.",
												Attributes: map[string]schema.Attribute{
													"from_field_path": schema.StringAttribute{
														Description:         "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",
														MarkdownDescription: "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Optional", "Required"),
														},
													},

													"merge_options": schema.SingleNestedAttribute{
														Description:         "MergeOptions Specifies merge options on a field path",
														MarkdownDescription: "MergeOptions Specifies merge options on a field path",
														Attributes: map[string]schema.Attribute{
															"append_slice": schema.BoolAttribute{
																Description:         "Specifies that already existing elements in a merged slice should be preserved",
																MarkdownDescription: "Specifies that already existing elements in a merged slice should be preserved",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"keep_map_values": schema.BoolAttribute{
																Description:         "Specifies that already existing values in a merged map should be preserved",
																MarkdownDescription: "Specifies that already existing values in a merged map should be preserved",
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

											"to_field_path": schema.StringAttribute{
												Description:         "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",
												MarkdownDescription: "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"transforms": schema.ListNestedAttribute{
												Description:         "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",
												MarkdownDescription: "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"convert": schema.SingleNestedAttribute{
															Description:         "Convert is used to cast the input into the given output type.",
															MarkdownDescription: "Convert is used to cast the input into the given output type.",
															Attributes: map[string]schema.Attribute{
																"format": schema.StringAttribute{
																	Description:         "The expected input format.  * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions.  If this property is null, the default conversion is applied.",
																	MarkdownDescription: "The expected input format.  * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions.  If this property is null, the default conversion is applied.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("none", "quantity", "json"),
																	},
																},

																"to_type": schema.StringAttribute{
																	Description:         "ToType is the type of the output of this transform.",
																	MarkdownDescription: "ToType is the type of the output of this transform.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("string", "int", "int64", "bool", "float64", "object", "list"),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"map": schema.MapAttribute{
															Description:         "Map uses the input as a key in the given map and returns the value.",
															MarkdownDescription: "Map uses the input as a key in the given map and returns the value.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"match": schema.SingleNestedAttribute{
															Description:         "Match is a more complex version of Map that matches a list of patterns.",
															MarkdownDescription: "Match is a more complex version of Map that matches a list of patterns.",
															Attributes: map[string]schema.Attribute{
																"fallback_to": schema.StringAttribute{
																	Description:         "Determines to what value the transform should fallback if no pattern matches.",
																	MarkdownDescription: "Determines to what value the transform should fallback if no pattern matches.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Value", "Input"),
																	},
																},

																"fallback_value": schema.MapAttribute{
																	Description:         "The fallback value that should be returned by the transform if now pattern matches.",
																	MarkdownDescription: "The fallback value that should be returned by the transform if now pattern matches.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"patterns": schema.ListNestedAttribute{
																	Description:         "The patterns that should be tested against the input string. Patterns are tested in order. The value of the first match is used as result of this transform.",
																	MarkdownDescription: "The patterns that should be tested against the input string. Patterns are tested in order. The value of the first match is used as result of this transform.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"literal": schema.StringAttribute{
																				Description:         "Literal exactly matches the input string (case sensitive). Is required if 'type' is 'literal'.",
																				MarkdownDescription: "Literal exactly matches the input string (case sensitive). Is required if 'type' is 'literal'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"regexp": schema.StringAttribute{
																				Description:         "Regexp to match against the input string. Is required if 'type' is 'regexp'.",
																				MarkdownDescription: "Regexp to match against the input string. Is required if 'type' is 'regexp'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"result": schema.MapAttribute{
																				Description:         "The value that is used as result of the transform if the pattern matches.",
																				MarkdownDescription: "The value that is used as result of the transform if the pattern matches.",
																				ElementType:         types.StringType,
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"type": schema.StringAttribute{
																				Description:         "Type specifies how the pattern matches the input.  * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default.  * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
																				MarkdownDescription: "Type specifies how the pattern matches the input.  * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default.  * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("literal", "regexp"),
																				},
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

														"math": schema.SingleNestedAttribute{
															Description:         "Math is used to transform the input via mathematical operations such as multiplication.",
															MarkdownDescription: "Math is used to transform the input via mathematical operations such as multiplication.",
															Attributes: map[string]schema.Attribute{
																"clamp_max": schema.Int64Attribute{
																	Description:         "ClampMax makes sure that the value is not bigger than the given value.",
																	MarkdownDescription: "ClampMax makes sure that the value is not bigger than the given value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"clamp_min": schema.Int64Attribute{
																	Description:         "ClampMin makes sure that the value is not smaller than the given value.",
																	MarkdownDescription: "ClampMin makes sure that the value is not smaller than the given value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"multiply": schema.Int64Attribute{
																	Description:         "Multiply the value.",
																	MarkdownDescription: "Multiply the value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "Type of the math transform to be run.",
																	MarkdownDescription: "Type of the math transform to be run.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Multiply", "ClampMin", "ClampMax"),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"string": schema.SingleNestedAttribute{
															Description:         "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",
															MarkdownDescription: "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",
															Attributes: map[string]schema.Attribute{
																"convert": schema.StringAttribute{
																	Description:         "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON.",
																	MarkdownDescription: "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("ToUpper", "ToLower", "ToBase64", "FromBase64", "ToJson", "ToSha1", "ToSha256", "ToSha512"),
																	},
																},

																"fmt": schema.StringAttribute{
																	Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																	MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"regexp": schema.SingleNestedAttribute{
																	Description:         "Extract a match from the input using a regular expression.",
																	MarkdownDescription: "Extract a match from the input using a regular expression.",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.Int64Attribute{
																			Description:         "Group number to match. 0 (the default) matches the entire expression.",
																			MarkdownDescription: "Group number to match. 0 (the default) matches the entire expression.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"match": schema.StringAttribute{
																			Description:         "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",
																			MarkdownDescription: "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"trim": schema.StringAttribute{
																	Description:         "Trim the prefix or suffix from the input",
																	MarkdownDescription: "Trim the prefix or suffix from the input",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "Type of the string transform to be run.",
																	MarkdownDescription: "Type of the string transform to be run.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Format", "Convert", "TrimPrefix", "TrimSuffix", "Regexp"),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"type": schema.StringAttribute{
															Description:         "Type of the transform to be run.",
															MarkdownDescription: "Type of the transform to be run.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("map", "match", "math", "string", "convert"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "Type sets the patching behaviour to be used. Each patch type may require its own fields to be set on the Patch object.",
												MarkdownDescription: "Type sets the patching behaviour to be used. Each patch type may require its own fields to be set on the Patch object.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("FromCompositeFieldPath", "FromEnvironmentFieldPath", "PatchSet", "ToCompositeFieldPath", "ToEnvironmentFieldPath", "CombineFromEnvironment", "CombineFromComposite", "CombineToComposite", "CombineToEnvironment"),
												},
											},
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

					"publish_connection_details_with_store_config_ref": schema.SingleNestedAttribute{
						Description:         "PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection details of composite resources dynamically provisioned using this composition will be published.",
						MarkdownDescription: "PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection details of composite resources dynamically provisioned using this composition will be published.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referenced StoreConfig.",
								MarkdownDescription: "Name of the referenced StoreConfig.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": schema.ListNestedAttribute{
						Description:         "Resources is the list of resource templates that will be used when a composite resource referring to this composition is created.",
						MarkdownDescription: "Resources is the list of resource templates that will be used when a composite resource referring to this composition is created.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"base": schema.MapAttribute{
									Description:         "Base is the target resource that the patches will be applied on.",
									MarkdownDescription: "Base is the target resource that the patches will be applied on.",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"connection_details": schema.ListNestedAttribute{
									Description:         "ConnectionDetails lists the propagation secret keys from this target resource to the composition instance connection secret.",
									MarkdownDescription: "ConnectionDetails lists the propagation secret keys from this target resource to the composition instance connection secret.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"from_connection_secret_key": schema.StringAttribute{
												Description:         "FromConnectionSecretKey is the key that will be used to fetch the value from the composed resource's connection secret.",
												MarkdownDescription: "FromConnectionSecretKey is the key that will be used to fetch the value from the composed resource's connection secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"from_field_path": schema.StringAttribute{
												Description:         "FromFieldPath is the path of the field on the composed resource whose value to be used as input. Name must be specified if the type is FromFieldPath.",
												MarkdownDescription: "FromFieldPath is the path of the field on the composed resource whose value to be used as input. Name must be specified if the type is FromFieldPath.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the connection secret key that will be propagated to the connection secret of the composition instance. Leave empty if you'd like to use the same key name.",
												MarkdownDescription: "Name of the connection secret key that will be propagated to the connection secret of the composition instance. Leave empty if you'd like to use the same key name.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type sets the connection detail fetching behaviour to be used. Each connection detail type may require its own fields to be set on the ConnectionDetail object. If the type is omitted Crossplane will attempt to infer it based on which other fields were specified. If multiple fields are specified the order of precedence is: 1. FromValue 2. FromConnectionSecretKey 3. FromFieldPath",
												MarkdownDescription: "Type sets the connection detail fetching behaviour to be used. Each connection detail type may require its own fields to be set on the ConnectionDetail object. If the type is omitted Crossplane will attempt to infer it based on which other fields were specified. If multiple fields are specified the order of precedence is: 1. FromValue 2. FromConnectionSecretKey 3. FromFieldPath",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("FromConnectionSecretKey", "FromFieldPath", "FromValue"),
												},
											},

											"value": schema.StringAttribute{
												Description:         "Value that will be propagated to the connection secret of the composite resource. May be set to inject a fixed, non-sensitive connection secret value, for example a well-known port.",
												MarkdownDescription: "Value that will be propagated to the connection secret of the composite resource. May be set to inject a fixed, non-sensitive connection secret value, for example a well-known port.",
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

								"name": schema.StringAttribute{
									Description:         "A Name uniquely identifies this entry within its Composition's resources array. Names are optional but *strongly* recommended. When all entries in the resources array are named entries may added, deleted, and reordered as long as their names do not change. When entries are not named the length and order of the resources array should be treated as immutable. Either all or no entries must be named.",
									MarkdownDescription: "A Name uniquely identifies this entry within its Composition's resources array. Names are optional but *strongly* recommended. When all entries in the resources array are named entries may added, deleted, and reordered as long as their names do not change. When entries are not named the length and order of the resources array should be treated as immutable. Either all or no entries must be named.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"patches": schema.ListNestedAttribute{
									Description:         "Patches will be applied as overlay to the base resource.",
									MarkdownDescription: "Patches will be applied as overlay to the base resource.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"combine": schema.SingleNestedAttribute{
												Description:         "Combine is the patch configuration for a CombineFromComposite, CombineFromEnvironment, CombineToComposite or CombineToEnvironment patch.",
												MarkdownDescription: "Combine is the patch configuration for a CombineFromComposite, CombineFromEnvironment, CombineToComposite or CombineToEnvironment patch.",
												Attributes: map[string]schema.Attribute{
													"strategy": schema.StringAttribute{
														Description:         "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",
														MarkdownDescription: "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("string"),
														},
													},

													"string": schema.SingleNestedAttribute{
														Description:         "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",
														MarkdownDescription: "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",
														Attributes: map[string]schema.Attribute{
															"fmt": schema.StringAttribute{
																Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"variables": schema.ListNestedAttribute{
														Description:         "Variables are the list of variables whose values will be retrieved and combined.",
														MarkdownDescription: "Variables are the list of variables whose values will be retrieved and combined.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"from_field_path": schema.StringAttribute{
																	Description:         "FromFieldPath is the path of the field on the source whose value is to be used as input.",
																	MarkdownDescription: "FromFieldPath is the path of the field on the source whose value is to be used as input.",
																	Required:            true,
																	Optional:            false,
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

											"from_field_path": schema.StringAttribute{
												Description:         "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath, FromEnvironmentFieldPath, ToCompositeFieldPath, ToEnvironmentFieldPath.",
												MarkdownDescription: "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath, FromEnvironmentFieldPath, ToCompositeFieldPath, ToEnvironmentFieldPath.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"patch_set_name": schema.StringAttribute{
												Description:         "PatchSetName to include patches from. Required when type is PatchSet.",
												MarkdownDescription: "PatchSetName to include patches from. Required when type is PatchSet.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"policy": schema.SingleNestedAttribute{
												Description:         "Policy configures the specifics of patching behaviour.",
												MarkdownDescription: "Policy configures the specifics of patching behaviour.",
												Attributes: map[string]schema.Attribute{
													"from_field_path": schema.StringAttribute{
														Description:         "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",
														MarkdownDescription: "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Optional", "Required"),
														},
													},

													"merge_options": schema.SingleNestedAttribute{
														Description:         "MergeOptions Specifies merge options on a field path",
														MarkdownDescription: "MergeOptions Specifies merge options on a field path",
														Attributes: map[string]schema.Attribute{
															"append_slice": schema.BoolAttribute{
																Description:         "Specifies that already existing elements in a merged slice should be preserved",
																MarkdownDescription: "Specifies that already existing elements in a merged slice should be preserved",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"keep_map_values": schema.BoolAttribute{
																Description:         "Specifies that already existing values in a merged map should be preserved",
																MarkdownDescription: "Specifies that already existing values in a merged map should be preserved",
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

											"to_field_path": schema.StringAttribute{
												Description:         "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",
												MarkdownDescription: "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"transforms": schema.ListNestedAttribute{
												Description:         "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",
												MarkdownDescription: "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"convert": schema.SingleNestedAttribute{
															Description:         "Convert is used to cast the input into the given output type.",
															MarkdownDescription: "Convert is used to cast the input into the given output type.",
															Attributes: map[string]schema.Attribute{
																"format": schema.StringAttribute{
																	Description:         "The expected input format.  * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions.  If this property is null, the default conversion is applied.",
																	MarkdownDescription: "The expected input format.  * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions.  If this property is null, the default conversion is applied.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("none", "quantity", "json"),
																	},
																},

																"to_type": schema.StringAttribute{
																	Description:         "ToType is the type of the output of this transform.",
																	MarkdownDescription: "ToType is the type of the output of this transform.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("string", "int", "int64", "bool", "float64", "object", "list"),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"map": schema.MapAttribute{
															Description:         "Map uses the input as a key in the given map and returns the value.",
															MarkdownDescription: "Map uses the input as a key in the given map and returns the value.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"match": schema.SingleNestedAttribute{
															Description:         "Match is a more complex version of Map that matches a list of patterns.",
															MarkdownDescription: "Match is a more complex version of Map that matches a list of patterns.",
															Attributes: map[string]schema.Attribute{
																"fallback_to": schema.StringAttribute{
																	Description:         "Determines to what value the transform should fallback if no pattern matches.",
																	MarkdownDescription: "Determines to what value the transform should fallback if no pattern matches.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Value", "Input"),
																	},
																},

																"fallback_value": schema.MapAttribute{
																	Description:         "The fallback value that should be returned by the transform if now pattern matches.",
																	MarkdownDescription: "The fallback value that should be returned by the transform if now pattern matches.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"patterns": schema.ListNestedAttribute{
																	Description:         "The patterns that should be tested against the input string. Patterns are tested in order. The value of the first match is used as result of this transform.",
																	MarkdownDescription: "The patterns that should be tested against the input string. Patterns are tested in order. The value of the first match is used as result of this transform.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"literal": schema.StringAttribute{
																				Description:         "Literal exactly matches the input string (case sensitive). Is required if 'type' is 'literal'.",
																				MarkdownDescription: "Literal exactly matches the input string (case sensitive). Is required if 'type' is 'literal'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"regexp": schema.StringAttribute{
																				Description:         "Regexp to match against the input string. Is required if 'type' is 'regexp'.",
																				MarkdownDescription: "Regexp to match against the input string. Is required if 'type' is 'regexp'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"result": schema.MapAttribute{
																				Description:         "The value that is used as result of the transform if the pattern matches.",
																				MarkdownDescription: "The value that is used as result of the transform if the pattern matches.",
																				ElementType:         types.StringType,
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"type": schema.StringAttribute{
																				Description:         "Type specifies how the pattern matches the input.  * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default.  * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
																				MarkdownDescription: "Type specifies how the pattern matches the input.  * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default.  * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("literal", "regexp"),
																				},
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

														"math": schema.SingleNestedAttribute{
															Description:         "Math is used to transform the input via mathematical operations such as multiplication.",
															MarkdownDescription: "Math is used to transform the input via mathematical operations such as multiplication.",
															Attributes: map[string]schema.Attribute{
																"clamp_max": schema.Int64Attribute{
																	Description:         "ClampMax makes sure that the value is not bigger than the given value.",
																	MarkdownDescription: "ClampMax makes sure that the value is not bigger than the given value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"clamp_min": schema.Int64Attribute{
																	Description:         "ClampMin makes sure that the value is not smaller than the given value.",
																	MarkdownDescription: "ClampMin makes sure that the value is not smaller than the given value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"multiply": schema.Int64Attribute{
																	Description:         "Multiply the value.",
																	MarkdownDescription: "Multiply the value.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "Type of the math transform to be run.",
																	MarkdownDescription: "Type of the math transform to be run.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Multiply", "ClampMin", "ClampMax"),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"string": schema.SingleNestedAttribute{
															Description:         "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",
															MarkdownDescription: "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",
															Attributes: map[string]schema.Attribute{
																"convert": schema.StringAttribute{
																	Description:         "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON.",
																	MarkdownDescription: "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("ToUpper", "ToLower", "ToBase64", "FromBase64", "ToJson", "ToSha1", "ToSha256", "ToSha512"),
																	},
																},

																"fmt": schema.StringAttribute{
																	Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																	MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"regexp": schema.SingleNestedAttribute{
																	Description:         "Extract a match from the input using a regular expression.",
																	MarkdownDescription: "Extract a match from the input using a regular expression.",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.Int64Attribute{
																			Description:         "Group number to match. 0 (the default) matches the entire expression.",
																			MarkdownDescription: "Group number to match. 0 (the default) matches the entire expression.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"match": schema.StringAttribute{
																			Description:         "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",
																			MarkdownDescription: "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"trim": schema.StringAttribute{
																	Description:         "Trim the prefix or suffix from the input",
																	MarkdownDescription: "Trim the prefix or suffix from the input",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "Type of the string transform to be run.",
																	MarkdownDescription: "Type of the string transform to be run.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Format", "Convert", "TrimPrefix", "TrimSuffix", "Regexp"),
																	},
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"type": schema.StringAttribute{
															Description:         "Type of the transform to be run.",
															MarkdownDescription: "Type of the transform to be run.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("map", "match", "math", "string", "convert"),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "Type sets the patching behaviour to be used. Each patch type may require its own fields to be set on the Patch object.",
												MarkdownDescription: "Type sets the patching behaviour to be used. Each patch type may require its own fields to be set on the Patch object.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("FromCompositeFieldPath", "FromEnvironmentFieldPath", "PatchSet", "ToCompositeFieldPath", "ToEnvironmentFieldPath", "CombineFromEnvironment", "CombineFromComposite", "CombineToComposite", "CombineToEnvironment"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"readiness_checks": schema.ListNestedAttribute{
									Description:         "ReadinessChecks allows users to define custom readiness checks. All checks have to return true in order for resource to be considered ready. The default readiness check is to have the 'Ready' condition to be 'True'.",
									MarkdownDescription: "ReadinessChecks allows users to define custom readiness checks. All checks have to return true in order for resource to be considered ready. The default readiness check is to have the 'Ready' condition to be 'True'.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"field_path": schema.StringAttribute{
												Description:         "FieldPath shows the path of the field whose value will be used.",
												MarkdownDescription: "FieldPath shows the path of the field whose value will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match_condition": schema.SingleNestedAttribute{
												Description:         "MatchCondition specifies the condition you'd like to match if you're using 'MatchCondition' type.",
												MarkdownDescription: "MatchCondition specifies the condition you'd like to match if you're using 'MatchCondition' type.",
												Attributes: map[string]schema.Attribute{
													"status": schema.StringAttribute{
														Description:         "Status is the status of the condition you'd like to match.",
														MarkdownDescription: "Status is the status of the condition you'd like to match.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type indicates the type of condition you'd like to use.",
														MarkdownDescription: "Type indicates the type of condition you'd like to use.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"match_integer": schema.Int64Attribute{
												Description:         "MatchInt is the value you'd like to match if you're using 'MatchInt' type.",
												MarkdownDescription: "MatchInt is the value you'd like to match if you're using 'MatchInt' type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match_string": schema.StringAttribute{
												Description:         "MatchString is the value you'd like to match if you're using 'MatchString' type.",
												MarkdownDescription: "MatchString is the value you'd like to match if you're using 'MatchString' type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type indicates the type of probe you'd like to use.",
												MarkdownDescription: "Type indicates the type of probe you'd like to use.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("MatchString", "MatchInteger", "NonEmpty", "MatchCondition", "MatchTrue", "MatchFalse", "None"),
												},
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

					"revision": schema.Int64Attribute{
						Description:         "Revision number. Newer revisions have larger numbers.",
						MarkdownDescription: "Revision number. Newer revisions have larger numbers.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"write_connection_secrets_to_namespace": schema.StringAttribute{
						Description:         "WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created. This field is planned to be removed in a future release in favor of PublishConnectionDetailsWithStoreConfigRef. Currently, both could be set independently and connection details would be published to both without affecting each other as long as related fields at MR level specified.",
						MarkdownDescription: "WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created. This field is planned to be removed in a future release in favor of PublishConnectionDetailsWithStoreConfigRef. Currently, both could be set independently and connection details would be published to both without affecting each other as long as related fields at MR level specified.",
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

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apiextensions_crossplane_io_composition_revision_v1")

	var model ApiextensionsCrossplaneIoCompositionRevisionV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("apiextensions.crossplane.io/v1")
	model.Kind = pointer.String("CompositionRevision")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "apiextensions.crossplane.io", Version: "v1", Resource: "compositionrevisions"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ApiextensionsCrossplaneIoCompositionRevisionV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apiextensions_crossplane_io_composition_revision_v1")

	var data ApiextensionsCrossplaneIoCompositionRevisionV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apiextensions.crossplane.io", Version: "v1", Resource: "compositionrevisions"}).
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

	var readResponse ApiextensionsCrossplaneIoCompositionRevisionV1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apiextensions_crossplane_io_composition_revision_v1")

	var model ApiextensionsCrossplaneIoCompositionRevisionV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apiextensions.crossplane.io/v1")
	model.Kind = pointer.String("CompositionRevision")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "apiextensions.crossplane.io", Version: "v1", Resource: "compositionrevisions"}).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ApiextensionsCrossplaneIoCompositionRevisionV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apiextensions_crossplane_io_composition_revision_v1")

	var data ApiextensionsCrossplaneIoCompositionRevisionV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "apiextensions.crossplane.io", Version: "v1", Resource: "compositionrevisions"}).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
