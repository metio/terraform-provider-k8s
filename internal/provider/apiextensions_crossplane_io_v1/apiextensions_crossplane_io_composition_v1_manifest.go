/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apiextensions_crossplane_io_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &ApiextensionsCrossplaneIoCompositionV1Manifest{}
)

func NewApiextensionsCrossplaneIoCompositionV1Manifest() datasource.DataSource {
	return &ApiextensionsCrossplaneIoCompositionV1Manifest{}
}

type ApiextensionsCrossplaneIoCompositionV1Manifest struct{}

type ApiextensionsCrossplaneIoCompositionV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

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
		Mode      *string `tfsdk:"mode" json:"mode,omitempty"`
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
						Join    *struct {
							Separator *string `tfsdk:"separator" json:"separator,omitempty"`
						} `tfsdk:"join" json:"join,omitempty"`
						Regexp *struct {
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
		Pipeline *[]struct {
			Credentials *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				SecretRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"credentials" json:"credentials,omitempty"`
			FunctionRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"function_ref" json:"functionRef,omitempty"`
			Input *map[string]string `tfsdk:"input" json:"input,omitempty"`
			Step  *string            `tfsdk:"step" json:"step,omitempty"`
		} `tfsdk:"pipeline" json:"pipeline,omitempty"`
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
						Join    *struct {
							Separator *string `tfsdk:"separator" json:"separator,omitempty"`
						} `tfsdk:"join" json:"join,omitempty"`
						Regexp *struct {
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
		WriteConnectionSecretsToNamespace *string `tfsdk:"write_connection_secrets_to_namespace" json:"writeConnectionSecretsToNamespace,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ApiextensionsCrossplaneIoCompositionV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apiextensions_crossplane_io_composition_v1_manifest"
}

func (r *ApiextensionsCrossplaneIoCompositionV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "A Composition defines a collection of managed resources or functions that Crossplane uses to create and manage new composite resources. Read the Crossplane documentation for [more information about Compositions](https://docs.crossplane.io/latest/concepts/compositions).",
		MarkdownDescription: "A Composition defines a collection of managed resources or functions that Crossplane uses to create and manage new composite resources. Read the Crossplane documentation for [more information about Compositions](https://docs.crossplane.io/latest/concepts/compositions).",
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
				Description:         "CompositionSpec specifies desired state of a composition.",
				MarkdownDescription: "CompositionSpec specifies desired state of a composition.",
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

					"mode": schema.StringAttribute{
						Description:         "Mode controls what type or 'mode' of Composition will be used. 'Pipeline' indicates that a Composition specifies a pipeline of Composition Functions, each of which is responsible for producing composed resources that Crossplane should create or update. 'Resources' indicates that a Composition uses what is commonly referred to as 'Patch & Transform' or P&T composition. This mode of Composition uses an array of resources, each a template for a composed resource. All Compositions should use Pipeline mode. Resources mode is deprecated. Resources mode won't be removed in Crossplane 1.x, and will remain the default to avoid breaking legacy Compositions. However, it's no longer accepting new features, and only accepting security related bug fixes.",
						MarkdownDescription: "Mode controls what type or 'mode' of Composition will be used. 'Pipeline' indicates that a Composition specifies a pipeline of Composition Functions, each of which is responsible for producing composed resources that Crossplane should create or update. 'Resources' indicates that a Composition uses what is commonly referred to as 'Patch & Transform' or P&T composition. This mode of Composition uses an array of resources, each a template for a composed resource. All Compositions should use Pipeline mode. Resources mode is deprecated. Resources mode won't be removed in Crossplane 1.x, and will remain the default to avoid breaking legacy Compositions. However, it's no longer accepting new features, and only accepting security related bug fixes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Resources", "Pipeline"),
						},
					},

					"patch_sets": schema.ListNestedAttribute{
						Description:         "PatchSets define a named set of patches that may be included by any resource in this Composition. PatchSets cannot themselves refer to other PatchSets. PatchSets are only used by the 'Resources' mode of Composition. They are ignored by other modes. Deprecated: Use Composition Functions instead.",
						MarkdownDescription: "PatchSets define a named set of patches that may be included by any resource in this Composition. PatchSets cannot themselves refer to other PatchSets. PatchSets are only used by the 'Resources' mode of Composition. They are ignored by other modes. Deprecated: Use Composition Functions instead.",
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
														Description:         "MergeOptions Specifies merge options on a field path.",
														MarkdownDescription: "MergeOptions Specifies merge options on a field path.",
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
																	Description:         "The expected input format. * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions. If this property is null, the default conversion is applied.",
																	MarkdownDescription: "The expected input format. * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions. If this property is null, the default conversion is applied.",
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
																		stringvalidator.OneOf("string", "int", "int64", "bool", "float64", "object", "array"),
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
																				Description:         "Type specifies how the pattern matches the input. * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default. * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
																				MarkdownDescription: "Type specifies how the pattern matches the input. * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default. * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
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
																	Description:         "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON. 'ToAdler32' generate a addler32 hash based on the input string.",
																	MarkdownDescription: "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON. 'ToAdler32' generate a addler32 hash based on the input string.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("ToUpper", "ToLower", "ToBase64", "FromBase64", "ToJson", "ToSha1", "ToSha256", "ToSha512", "ToAdler32"),
																	},
																},

																"fmt": schema.StringAttribute{
																	Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																	MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"join": schema.SingleNestedAttribute{
																	Description:         "Join defines parameters to join a slice of values to a string.",
																	MarkdownDescription: "Join defines parameters to join a slice of values to a string.",
																	Attributes: map[string]schema.Attribute{
																		"separator": schema.StringAttribute{
																			Description:         "Separator defines the character that should separate the values from each other in the joined string.",
																			MarkdownDescription: "Separator defines the character that should separate the values from each other in the joined string.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
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
																		stringvalidator.OneOf("Format", "Convert", "TrimPrefix", "TrimSuffix", "Regexp", "Join"),
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
													stringvalidator.OneOf("FromCompositeFieldPath", "PatchSet", "ToCompositeFieldPath", "CombineFromComposite", "CombineToComposite"),
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

					"pipeline": schema.ListNestedAttribute{
						Description:         "Pipeline is a list of composition function steps that will be used when a composite resource referring to this composition is created. One of resources and pipeline must be specified - you cannot specify both. The Pipeline is only used by the 'Pipeline' mode of Composition. It is ignored by other modes.",
						MarkdownDescription: "Pipeline is a list of composition function steps that will be used when a composite resource referring to this composition is created. One of resources and pipeline must be specified - you cannot specify both. The Pipeline is only used by the 'Pipeline' mode of Composition. It is ignored by other modes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"credentials": schema.ListNestedAttribute{
									Description:         "Credentials are optional credentials that the Composition Function needs.",
									MarkdownDescription: "Credentials are optional credentials that the Composition Function needs.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of this set of credentials.",
												MarkdownDescription: "Name of this set of credentials.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "A SecretRef is a reference to a secret containing credentials that should be supplied to the function.",
												MarkdownDescription: "A SecretRef is a reference to a secret containing credentials that should be supplied to the function.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the secret.",
														MarkdownDescription: "Name of the secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the secret.",
														MarkdownDescription: "Namespace of the secret.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"source": schema.StringAttribute{
												Description:         "Source of the function credentials.",
												MarkdownDescription: "Source of the function credentials.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("None", "Secret"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"function_ref": schema.SingleNestedAttribute{
									Description:         "FunctionRef is a reference to the Composition Function this step should execute.",
									MarkdownDescription: "FunctionRef is a reference to the Composition Function this step should execute.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referenced Function.",
											MarkdownDescription: "Name of the referenced Function.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"input": schema.MapAttribute{
									Description:         "Input is an optional, arbitrary Kubernetes resource (i.e. a resource with an apiVersion and kind) that will be passed to the Composition Function as the 'input' of its RunFunctionRequest.",
									MarkdownDescription: "Input is an optional, arbitrary Kubernetes resource (i.e. a resource with an apiVersion and kind) that will be passed to the Composition Function as the 'input' of its RunFunctionRequest.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"step": schema.StringAttribute{
									Description:         "Step name. Must be unique within its Pipeline.",
									MarkdownDescription: "Step name. Must be unique within its Pipeline.",
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

					"publish_connection_details_with_store_config_ref": schema.SingleNestedAttribute{
						Description:         "PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection details of composite resources dynamically provisioned using this composition will be published. THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored unless the relevant Crossplane feature flag is enabled, and may be changed or removed without notice.",
						MarkdownDescription: "PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection details of composite resources dynamically provisioned using this composition will be published. THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored unless the relevant Crossplane feature flag is enabled, and may be changed or removed without notice.",
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
						Description:         "Resources is a list of resource templates that will be used when a composite resource referring to this composition is created. Resources are only used by the 'Resources' mode of Composition. They are ignored by other modes. Deprecated: Use Composition Functions instead.",
						MarkdownDescription: "Resources is a list of resource templates that will be used when a composite resource referring to this composition is created. Resources are only used by the 'Resources' mode of Composition. They are ignored by other modes. Deprecated: Use Composition Functions instead.",
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
														Description:         "MergeOptions Specifies merge options on a field path.",
														MarkdownDescription: "MergeOptions Specifies merge options on a field path.",
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
																	Description:         "The expected input format. * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions. If this property is null, the default conversion is applied.",
																	MarkdownDescription: "The expected input format. * 'quantity' - parses the input as a K8s ['resource.Quantity'](https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Quantity). Only used during 'string -> float64' conversions. * 'json' - parses the input as a JSON string. Only used during 'string -> object' or 'string -> list' conversions. If this property is null, the default conversion is applied.",
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
																		stringvalidator.OneOf("string", "int", "int64", "bool", "float64", "object", "array"),
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
																				Description:         "Type specifies how the pattern matches the input. * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default. * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
																				MarkdownDescription: "Type specifies how the pattern matches the input. * 'literal' - the pattern value has to exactly match (case sensitive) the input string. This is the default. * 'regexp' - the pattern treated as a regular expression against which the input string is tested. Crossplane will throw an error if the key is not a valid regexp.",
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
																	Description:         "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON. 'ToAdler32' generate a addler32 hash based on the input string.",
																	MarkdownDescription: "Optional conversion method to be specified. 'ToUpper' and 'ToLower' change the letter case of the input string. 'ToBase64' and 'FromBase64' perform a base64 conversion based on the input string. 'ToJson' converts any input value into its raw JSON representation. 'ToSha1', 'ToSha256' and 'ToSha512' generate a hash value based on the input converted to JSON. 'ToAdler32' generate a addler32 hash based on the input string.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("ToUpper", "ToLower", "ToBase64", "FromBase64", "ToJson", "ToSha1", "ToSha256", "ToSha512", "ToAdler32"),
																	},
																},

																"fmt": schema.StringAttribute{
																	Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																	MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"join": schema.SingleNestedAttribute{
																	Description:         "Join defines parameters to join a slice of values to a string.",
																	MarkdownDescription: "Join defines parameters to join a slice of values to a string.",
																	Attributes: map[string]schema.Attribute{
																		"separator": schema.StringAttribute{
																			Description:         "Separator defines the character that should separate the values from each other in the joined string.",
																			MarkdownDescription: "Separator defines the character that should separate the values from each other in the joined string.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
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
																		stringvalidator.OneOf("Format", "Convert", "TrimPrefix", "TrimSuffix", "Regexp", "Join"),
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
													stringvalidator.OneOf("FromCompositeFieldPath", "PatchSet", "ToCompositeFieldPath", "CombineFromComposite", "CombineToComposite"),
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

					"write_connection_secrets_to_namespace": schema.StringAttribute{
						Description:         "WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created. This field is planned to be replaced in a future release in favor of PublishConnectionDetailsWithStoreConfigRef. Currently, both could be set independently and connection details would be published to both without affecting each other as long as related fields at MR level specified.",
						MarkdownDescription: "WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created. This field is planned to be replaced in a future release in favor of PublishConnectionDetailsWithStoreConfigRef. Currently, both could be set independently and connection details would be published to both without affecting each other as long as related fields at MR level specified.",
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

func (r *ApiextensionsCrossplaneIoCompositionV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apiextensions_crossplane_io_composition_v1_manifest")

	var model ApiextensionsCrossplaneIoCompositionV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apiextensions.crossplane.io/v1")
	model.Kind = pointer.String("Composition")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
