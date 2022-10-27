/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

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

type ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource)(nil)
)

type ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		CompositeTypeRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`
		} `tfsdk:"composite_type_ref" yaml:"compositeTypeRef,omitempty"`

		PatchSets *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Patches *[]struct {
				Combine *struct {
					Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

					String *struct {
						Fmt *string `tfsdk:"fmt" yaml:"fmt,omitempty"`
					} `tfsdk:"string" yaml:"string,omitempty"`

					Variables *[]struct {
						FromFieldPath *string `tfsdk:"from_field_path" yaml:"fromFieldPath,omitempty"`
					} `tfsdk:"variables" yaml:"variables,omitempty"`
				} `tfsdk:"combine" yaml:"combine,omitempty"`

				FromFieldPath *string `tfsdk:"from_field_path" yaml:"fromFieldPath,omitempty"`

				PatchSetName *string `tfsdk:"patch_set_name" yaml:"patchSetName,omitempty"`

				Policy *struct {
					FromFieldPath *string `tfsdk:"from_field_path" yaml:"fromFieldPath,omitempty"`
				} `tfsdk:"policy" yaml:"policy,omitempty"`

				ToFieldPath *string `tfsdk:"to_field_path" yaml:"toFieldPath,omitempty"`

				Transforms *[]struct {
					Convert *struct {
						ToType *string `tfsdk:"to_type" yaml:"toType,omitempty"`
					} `tfsdk:"convert" yaml:"convert,omitempty"`

					Map *map[string]string `tfsdk:"map" yaml:"map,omitempty"`

					Math *struct {
						Multiply *int64 `tfsdk:"multiply" yaml:"multiply,omitempty"`
					} `tfsdk:"math" yaml:"math,omitempty"`

					String *struct {
						Convert *string `tfsdk:"convert" yaml:"convert,omitempty"`

						Fmt *string `tfsdk:"fmt" yaml:"fmt,omitempty"`

						Regexp *struct {
							Group *int64 `tfsdk:"group" yaml:"group,omitempty"`

							Match *string `tfsdk:"match" yaml:"match,omitempty"`
						} `tfsdk:"regexp" yaml:"regexp,omitempty"`

						Trim *string `tfsdk:"trim" yaml:"trim,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"string" yaml:"string,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"transforms" yaml:"transforms,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"patches" yaml:"patches,omitempty"`
		} `tfsdk:"patch_sets" yaml:"patchSets,omitempty"`

		PublishConnectionDetailsWithStoreConfigRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"publish_connection_details_with_store_config_ref" yaml:"publishConnectionDetailsWithStoreConfigRef,omitempty"`

		Resources *[]struct {
			Base utilities.Dynamic `tfsdk:"base" yaml:"base,omitempty"`

			ConnectionDetails *[]struct {
				FromConnectionSecretKey *string `tfsdk:"from_connection_secret_key" yaml:"fromConnectionSecretKey,omitempty"`

				FromFieldPath *string `tfsdk:"from_field_path" yaml:"fromFieldPath,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"connection_details" yaml:"connectionDetails,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Patches *[]struct {
				Combine *struct {
					Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

					String *struct {
						Fmt *string `tfsdk:"fmt" yaml:"fmt,omitempty"`
					} `tfsdk:"string" yaml:"string,omitempty"`

					Variables *[]struct {
						FromFieldPath *string `tfsdk:"from_field_path" yaml:"fromFieldPath,omitempty"`
					} `tfsdk:"variables" yaml:"variables,omitempty"`
				} `tfsdk:"combine" yaml:"combine,omitempty"`

				FromFieldPath *string `tfsdk:"from_field_path" yaml:"fromFieldPath,omitempty"`

				PatchSetName *string `tfsdk:"patch_set_name" yaml:"patchSetName,omitempty"`

				Policy *struct {
					FromFieldPath *string `tfsdk:"from_field_path" yaml:"fromFieldPath,omitempty"`
				} `tfsdk:"policy" yaml:"policy,omitempty"`

				ToFieldPath *string `tfsdk:"to_field_path" yaml:"toFieldPath,omitempty"`

				Transforms *[]struct {
					Convert *struct {
						ToType *string `tfsdk:"to_type" yaml:"toType,omitempty"`
					} `tfsdk:"convert" yaml:"convert,omitempty"`

					Map *map[string]string `tfsdk:"map" yaml:"map,omitempty"`

					Math *struct {
						Multiply *int64 `tfsdk:"multiply" yaml:"multiply,omitempty"`
					} `tfsdk:"math" yaml:"math,omitempty"`

					String *struct {
						Convert *string `tfsdk:"convert" yaml:"convert,omitempty"`

						Fmt *string `tfsdk:"fmt" yaml:"fmt,omitempty"`

						Regexp *struct {
							Group *int64 `tfsdk:"group" yaml:"group,omitempty"`

							Match *string `tfsdk:"match" yaml:"match,omitempty"`
						} `tfsdk:"regexp" yaml:"regexp,omitempty"`

						Trim *string `tfsdk:"trim" yaml:"trim,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"string" yaml:"string,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"transforms" yaml:"transforms,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"patches" yaml:"patches,omitempty"`

			ReadinessChecks *[]struct {
				FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

				MatchInteger *int64 `tfsdk:"match_integer" yaml:"matchInteger,omitempty"`

				MatchString *string `tfsdk:"match_string" yaml:"matchString,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"readiness_checks" yaml:"readinessChecks,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		Revision *int64 `tfsdk:"revision" yaml:"revision,omitempty"`

		WriteConnectionSecretsToNamespace *string `tfsdk:"write_connection_secrets_to_namespace" yaml:"writeConnectionSecretsToNamespace,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource() resource.Resource {
	return &ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource{}
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiextensions_crossplane_io_composition_revision_v1alpha1"
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "A CompositionRevision represents a revision in time of a Composition. Revisions are created by Crossplane; they should be treated as immutable.",
		MarkdownDescription: "A CompositionRevision represents a revision in time of a Composition. Revisions are created by Crossplane; they should be treated as immutable.",
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
				Description:         "CompositionRevisionSpec specifies the desired state of the composition revision.",
				MarkdownDescription: "CompositionRevisionSpec specifies the desired state of the composition revision.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"composite_type_ref": {
						Description:         "CompositeTypeRef specifies the type of composite resource that this composition is compatible with.",
						MarkdownDescription: "CompositeTypeRef specifies the type of composite resource that this composition is compatible with.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "APIVersion of the type.",
								MarkdownDescription: "APIVersion of the type.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the type.",
								MarkdownDescription: "Kind of the type.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"patch_sets": {
						Description:         "PatchSets define a named set of patches that may be included by any resource in this Composition. PatchSets cannot themselves refer to other PatchSets.",
						MarkdownDescription: "PatchSets define a named set of patches that may be included by any resource in this Composition. PatchSets cannot themselves refer to other PatchSets.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of this PatchSet.",
								MarkdownDescription: "Name of this PatchSet.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"patches": {
								Description:         "Patches will be applied as an overlay to the base resource.",
								MarkdownDescription: "Patches will be applied as an overlay to the base resource.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"combine": {
										Description:         "Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.",
										MarkdownDescription: "Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"strategy": {
												Description:         "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",
												MarkdownDescription: "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("string"),
												},
											},

											"string": {
												Description:         "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",
												MarkdownDescription: "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fmt": {
														Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
														MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",

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

											"variables": {
												Description:         "Variables are the list of variables whose values will be retrieved and combined.",
												MarkdownDescription: "Variables are the list of variables whose values will be retrieved and combined.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"from_field_path": {
														Description:         "FromFieldPath is the path of the field on the source whose value is to be used as input.",
														MarkdownDescription: "FromFieldPath is the path of the field on the source whose value is to be used as input.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_field_path": {
										Description:         "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath or ToCompositeFieldPath.",
										MarkdownDescription: "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath or ToCompositeFieldPath.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"patch_set_name": {
										Description:         "PatchSetName to include patches from. Required when type is PatchSet.",
										MarkdownDescription: "PatchSetName to include patches from. Required when type is PatchSet.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"policy": {
										Description:         "Policy configures the specifics of patching behaviour.",
										MarkdownDescription: "Policy configures the specifics of patching behaviour.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"from_field_path": {
												Description:         "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",
												MarkdownDescription: "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Optional", "Required"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_field_path": {
										Description:         "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",
										MarkdownDescription: "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"transforms": {
										Description:         "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",
										MarkdownDescription: "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"convert": {
												Description:         "Convert is used to cast the input into the given output type.",
												MarkdownDescription: "Convert is used to cast the input into the given output type.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"to_type": {
														Description:         "ToType is the type of the output of this transform.",
														MarkdownDescription: "ToType is the type of the output of this transform.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("string", "int", "int64", "bool", "float64"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"map": {
												Description:         "Map uses the input as a key in the given map and returns the value.",
												MarkdownDescription: "Map uses the input as a key in the given map and returns the value.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"math": {
												Description:         "Math is used to transform the input via mathematical operations such as multiplication.",
												MarkdownDescription: "Math is used to transform the input via mathematical operations such as multiplication.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"multiply": {
														Description:         "Multiply the value.",
														MarkdownDescription: "Multiply the value.",

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

											"string": {
												Description:         "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",
												MarkdownDescription: "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"convert": {
														Description:         "Convert the type of conversion to Upper/Lower case.",
														MarkdownDescription: "Convert the type of conversion to Upper/Lower case.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("ToUpper", "ToLower", "ToBase64", "FromBase64"),
														},
													},

													"fmt": {
														Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
														MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"regexp": {
														Description:         "Extract a match from the input using a regular expression.",
														MarkdownDescription: "Extract a match from the input using a regular expression.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"group": {
																Description:         "Group number to match. 0 (the default) matches the entire expression.",
																MarkdownDescription: "Group number to match. 0 (the default) matches the entire expression.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match": {
																Description:         "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",
																MarkdownDescription: "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",

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

													"trim": {
														Description:         "Trim the prefix or suffix from the input",
														MarkdownDescription: "Trim the prefix or suffix from the input",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "Type of the string transform to be run.",
														MarkdownDescription: "Type of the string transform to be run.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Format", "Convert", "TrimPrefix", "TrimSuffix", "Regexp"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Type of the transform to be run.",
												MarkdownDescription: "Type of the transform to be run.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("map", "math", "string", "convert"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Type sets the patching behaviour to be used. Each patch type may require its' own fields to be set on the Patch object.",
										MarkdownDescription: "Type sets the patching behaviour to be used. Each patch type may require its' own fields to be set on the Patch object.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("FromCompositeFieldPath", "PatchSet", "ToCompositeFieldPath", "CombineFromComposite", "CombineToComposite"),
										},
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"publish_connection_details_with_store_config_ref": {
						Description:         "PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection details of composite resources dynamically provisioned using this composition will be published.",
						MarkdownDescription: "PublishConnectionDetailsWithStoreConfig specifies the secret store config with which the connection details of composite resources dynamically provisioned using this composition will be published.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the referenced StoreConfig.",
								MarkdownDescription: "Name of the referenced StoreConfig.",

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

					"resources": {
						Description:         "Resources is the list of resource templates that will be used when a composite resource referring to this composition is created.",
						MarkdownDescription: "Resources is the list of resource templates that will be used when a composite resource referring to this composition is created.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"base": {
								Description:         "Base is the target resource that the patches will be applied on.",
								MarkdownDescription: "Base is the target resource that the patches will be applied on.",

								Type: utilities.DynamicType{},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"connection_details": {
								Description:         "ConnectionDetails lists the propagation secret keys from this target resource to the composition instance connection secret.",
								MarkdownDescription: "ConnectionDetails lists the propagation secret keys from this target resource to the composition instance connection secret.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"from_connection_secret_key": {
										Description:         "FromConnectionSecretKey is the key that will be used to fetch the value from the given target resource's secret.",
										MarkdownDescription: "FromConnectionSecretKey is the key that will be used to fetch the value from the given target resource's secret.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_field_path": {
										Description:         "FromFieldPath is the path of the field on the composed resource whose value to be used as input. Name must be specified if the type is FromFieldPath is specified.",
										MarkdownDescription: "FromFieldPath is the path of the field on the composed resource whose value to be used as input. Name must be specified if the type is FromFieldPath is specified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the connection secret key that will be propagated to the connection secret of the composition instance. Leave empty if you'd like to use the same key name.",
										MarkdownDescription: "Name of the connection secret key that will be propagated to the connection secret of the composition instance. Leave empty if you'd like to use the same key name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Type sets the connection detail fetching behaviour to be used. Each connection detail type may require its own fields to be set on the ConnectionDetail object. If the type is omitted Crossplane will attempt to infer it based on which other fields were specified.",
										MarkdownDescription: "Type sets the connection detail fetching behaviour to be used. Each connection detail type may require its own fields to be set on the ConnectionDetail object. If the type is omitted Crossplane will attempt to infer it based on which other fields were specified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("FromConnectionSecretKey", "FromFieldPath", "FromValue"),
										},
									},

									"value": {
										Description:         "Value that will be propagated to the connection secret of the composition instance. Typically you should use FromConnectionSecretKey instead, but an explicit value may be set to inject a fixed, non-sensitive connection secret values, for example a well-known port. Supercedes FromConnectionSecretKey when set.",
										MarkdownDescription: "Value that will be propagated to the connection secret of the composition instance. Typically you should use FromConnectionSecretKey instead, but an explicit value may be set to inject a fixed, non-sensitive connection secret values, for example a well-known port. Supercedes FromConnectionSecretKey when set.",

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
								Description:         "A Name uniquely identifies this entry within its Composition's resources array. Names are optional but *strongly* recommended. When all entries in the resources array are named entries may added, deleted, and reordered as long as their names do not change. When entries are not named the length and order of the resources array should be treated as immutable. Either all or no entries must be named.",
								MarkdownDescription: "A Name uniquely identifies this entry within its Composition's resources array. Names are optional but *strongly* recommended. When all entries in the resources array are named entries may added, deleted, and reordered as long as their names do not change. When entries are not named the length and order of the resources array should be treated as immutable. Either all or no entries must be named.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"patches": {
								Description:         "Patches will be applied as overlay to the base resource.",
								MarkdownDescription: "Patches will be applied as overlay to the base resource.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"combine": {
										Description:         "Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.",
										MarkdownDescription: "Combine is the patch configuration for a CombineFromComposite or CombineToComposite patch.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"strategy": {
												Description:         "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",
												MarkdownDescription: "Strategy defines the strategy to use to combine the input variable values. Currently only string is supported.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("string"),
												},
											},

											"string": {
												Description:         "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",
												MarkdownDescription: "String declares that input variables should be combined into a single string, using the relevant settings for formatting purposes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fmt": {
														Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
														MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",

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

											"variables": {
												Description:         "Variables are the list of variables whose values will be retrieved and combined.",
												MarkdownDescription: "Variables are the list of variables whose values will be retrieved and combined.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"from_field_path": {
														Description:         "FromFieldPath is the path of the field on the source whose value is to be used as input.",
														MarkdownDescription: "FromFieldPath is the path of the field on the source whose value is to be used as input.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"from_field_path": {
										Description:         "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath or ToCompositeFieldPath.",
										MarkdownDescription: "FromFieldPath is the path of the field on the resource whose value is to be used as input. Required when type is FromCompositeFieldPath or ToCompositeFieldPath.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"patch_set_name": {
										Description:         "PatchSetName to include patches from. Required when type is PatchSet.",
										MarkdownDescription: "PatchSetName to include patches from. Required when type is PatchSet.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"policy": {
										Description:         "Policy configures the specifics of patching behaviour.",
										MarkdownDescription: "Policy configures the specifics of patching behaviour.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"from_field_path": {
												Description:         "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",
												MarkdownDescription: "FromFieldPath specifies how to patch from a field path. The default is 'Optional', which means the patch will be a no-op if the specified fromFieldPath does not exist. Use 'Required' if the patch should fail if the specified path does not exist.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Optional", "Required"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_field_path": {
										Description:         "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",
										MarkdownDescription: "ToFieldPath is the path of the field on the resource whose value will be changed with the result of transforms. Leave empty if you'd like to propagate to the same path as fromFieldPath.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"transforms": {
										Description:         "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",
										MarkdownDescription: "Transforms are the list of functions that are used as a FIFO pipe for the input to be transformed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"convert": {
												Description:         "Convert is used to cast the input into the given output type.",
												MarkdownDescription: "Convert is used to cast the input into the given output type.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"to_type": {
														Description:         "ToType is the type of the output of this transform.",
														MarkdownDescription: "ToType is the type of the output of this transform.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("string", "int", "int64", "bool", "float64"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"map": {
												Description:         "Map uses the input as a key in the given map and returns the value.",
												MarkdownDescription: "Map uses the input as a key in the given map and returns the value.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"math": {
												Description:         "Math is used to transform the input via mathematical operations such as multiplication.",
												MarkdownDescription: "Math is used to transform the input via mathematical operations such as multiplication.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"multiply": {
														Description:         "Multiply the value.",
														MarkdownDescription: "Multiply the value.",

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

											"string": {
												Description:         "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",
												MarkdownDescription: "String is used to transform the input into a string or a different kind of string. Note that the input does not necessarily need to be a string.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"convert": {
														Description:         "Convert the type of conversion to Upper/Lower case.",
														MarkdownDescription: "Convert the type of conversion to Upper/Lower case.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("ToUpper", "ToLower", "ToBase64", "FromBase64"),
														},
													},

													"fmt": {
														Description:         "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",
														MarkdownDescription: "Format the input using a Go format string. See https://golang.org/pkg/fmt/ for details.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"regexp": {
														Description:         "Extract a match from the input using a regular expression.",
														MarkdownDescription: "Extract a match from the input using a regular expression.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"group": {
																Description:         "Group number to match. 0 (the default) matches the entire expression.",
																MarkdownDescription: "Group number to match. 0 (the default) matches the entire expression.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"match": {
																Description:         "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",
																MarkdownDescription: "Match string. May optionally include submatches, aka capture groups. See https://pkg.go.dev/regexp/ for details.",

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

													"trim": {
														Description:         "Trim the prefix or suffix from the input",
														MarkdownDescription: "Trim the prefix or suffix from the input",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "Type of the string transform to be run.",
														MarkdownDescription: "Type of the string transform to be run.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("Format", "Convert", "TrimPrefix", "TrimSuffix", "Regexp"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Type of the transform to be run.",
												MarkdownDescription: "Type of the transform to be run.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("map", "math", "string", "convert"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Type sets the patching behaviour to be used. Each patch type may require its' own fields to be set on the Patch object.",
										MarkdownDescription: "Type sets the patching behaviour to be used. Each patch type may require its' own fields to be set on the Patch object.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("FromCompositeFieldPath", "PatchSet", "ToCompositeFieldPath", "CombineFromComposite", "CombineToComposite"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"readiness_checks": {
								Description:         "ReadinessChecks allows users to define custom readiness checks. All checks have to return true in order for resource to be considered ready. The default readiness check is to have the 'Ready' condition to be 'True'.",
								MarkdownDescription: "ReadinessChecks allows users to define custom readiness checks. All checks have to return true in order for resource to be considered ready. The default readiness check is to have the 'Ready' condition to be 'True'.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"field_path": {
										Description:         "FieldPath shows the path of the field whose value will be used.",
										MarkdownDescription: "FieldPath shows the path of the field whose value will be used.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_integer": {
										Description:         "MatchInt is the value you'd like to match if you're using 'MatchInt' type.",
										MarkdownDescription: "MatchInt is the value you'd like to match if you're using 'MatchInt' type.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_string": {
										Description:         "MatchString is the value you'd like to match if you're using 'MatchString' type.",
										MarkdownDescription: "MatchString is the value you'd like to match if you're using 'MatchString' type.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Type indicates the type of probe you'd like to use.",
										MarkdownDescription: "Type indicates the type of probe you'd like to use.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("MatchString", "MatchInteger", "NonEmpty", "None"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"revision": {
						Description:         "Revision number. Newer revisions have larger numbers.",
						MarkdownDescription: "Revision number. Newer revisions have larger numbers.",

						Type: types.Int64Type,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"write_connection_secrets_to_namespace": {
						Description:         "WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created. This field is planned to be removed in a future release in favor of PublishConnectionDetailsWithStoreConfigRef. Currently, both could be set independently and connection details would be published to both without affecting each other as long as related fields at MR level specified.",
						MarkdownDescription: "WriteConnectionSecretsToNamespace specifies the namespace in which the connection secrets of composite resource dynamically provisioned using this composition will be created. This field is planned to be removed in a future release in favor of PublishConnectionDetailsWithStoreConfigRef. Currently, both could be set independently and connection details would be published to both without affecting each other as long as related fields at MR level specified.",

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
		},
	}, nil
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apiextensions_crossplane_io_composition_revision_v1alpha1")

	var state ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apiextensions.crossplane.io/v1alpha1")
	goModel.Kind = utilities.Ptr("CompositionRevision")

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

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apiextensions_crossplane_io_composition_revision_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apiextensions_crossplane_io_composition_revision_v1alpha1")

	var state ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apiextensions.crossplane.io/v1alpha1")
	goModel.Kind = utilities.Ptr("CompositionRevision")

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

func (r *ApiextensionsCrossplaneIoCompositionRevisionV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apiextensions_crossplane_io_composition_revision_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
