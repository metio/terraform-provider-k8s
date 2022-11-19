/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource struct{}

var (
	_ resource.Resource = (*ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource)(nil)
)

type ApiextensionsCrossplaneIoCompositeResourceDefinitionV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ApiextensionsCrossplaneIoCompositeResourceDefinitionV1GoModel struct {
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
		ClaimNames *struct {
			Categories *[]string `tfsdk:"categories" yaml:"categories,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			ListKind *string `tfsdk:"list_kind" yaml:"listKind,omitempty"`

			Plural *string `tfsdk:"plural" yaml:"plural,omitempty"`

			ShortNames *[]string `tfsdk:"short_names" yaml:"shortNames,omitempty"`

			Singular *string `tfsdk:"singular" yaml:"singular,omitempty"`
		} `tfsdk:"claim_names" yaml:"claimNames,omitempty"`

		ConnectionSecretKeys *[]string `tfsdk:"connection_secret_keys" yaml:"connectionSecretKeys,omitempty"`

		DefaultCompositionRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"default_composition_ref" yaml:"defaultCompositionRef,omitempty"`

		EnforcedCompositionRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"enforced_composition_ref" yaml:"enforcedCompositionRef,omitempty"`

		Group *string `tfsdk:"group" yaml:"group,omitempty"`

		Names *struct {
			Categories *[]string `tfsdk:"categories" yaml:"categories,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			ListKind *string `tfsdk:"list_kind" yaml:"listKind,omitempty"`

			Plural *string `tfsdk:"plural" yaml:"plural,omitempty"`

			ShortNames *[]string `tfsdk:"short_names" yaml:"shortNames,omitempty"`

			Singular *string `tfsdk:"singular" yaml:"singular,omitempty"`
		} `tfsdk:"names" yaml:"names,omitempty"`

		Versions *[]struct {
			AdditionalPrinterColumns *[]struct {
				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Format *string `tfsdk:"format" yaml:"format,omitempty"`

				JsonPath *string `tfsdk:"json_path" yaml:"jsonPath,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"additional_printer_columns" yaml:"additionalPrinterColumns,omitempty"`

			Deprecated *bool `tfsdk:"deprecated" yaml:"deprecated,omitempty"`

			DeprecationWarning *string `tfsdk:"deprecation_warning" yaml:"deprecationWarning,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Referenceable *bool `tfsdk:"referenceable" yaml:"referenceable,omitempty"`

			Schema *struct {
				OpenAPIV3Schema utilities.Dynamic `tfsdk:"open_apiv3_schema" yaml:"openAPIV3Schema,omitempty"`
			} `tfsdk:"schema" yaml:"schema,omitempty"`

			Served *bool `tfsdk:"served" yaml:"served,omitempty"`
		} `tfsdk:"versions" yaml:"versions,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource() resource.Resource {
	return &ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource{}
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apiextensions_crossplane_io_composite_resource_definition_v1"
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "A CompositeResourceDefinition defines a new kind of composite infrastructure resource. The new resource is composed of other composite or managed infrastructure resources.",
		MarkdownDescription: "A CompositeResourceDefinition defines a new kind of composite infrastructure resource. The new resource is composed of other composite or managed infrastructure resources.",
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
				Description:         "CompositeResourceDefinitionSpec specifies the desired state of the definition.",
				MarkdownDescription: "CompositeResourceDefinitionSpec specifies the desired state of the definition.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"claim_names": {
						Description:         "ClaimNames specifies the names of an optional composite resource claim. When claim names are specified Crossplane will create a namespaced 'composite resource claim' CRD that corresponds to the defined composite resource. This composite resource claim acts as a namespaced proxy for the composite resource; creating, updating, or deleting the claim will create, update, or delete a corresponding composite resource. You may add claim names to an existing CompositeResourceDefinition, but they cannot be changed or removed once they have been set.",
						MarkdownDescription: "ClaimNames specifies the names of an optional composite resource claim. When claim names are specified Crossplane will create a namespaced 'composite resource claim' CRD that corresponds to the defined composite resource. This composite resource claim acts as a namespaced proxy for the composite resource; creating, updating, or deleting the claim will create, update, or delete a corresponding composite resource. You may add claim names to an existing CompositeResourceDefinition, but they cannot be changed or removed once they have been set.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"categories": {
								Description:         "categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like 'kubectl get all'.",
								MarkdownDescription: "categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like 'kubectl get all'.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the 'kind' attribute in API calls.",
								MarkdownDescription: "kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the 'kind' attribute in API calls.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"list_kind": {
								Description:         "listKind is the serialized kind of the list for this resource. Defaults to ''kind'List'.",
								MarkdownDescription: "listKind is the serialized kind of the list for this resource. Defaults to ''kind'List'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plural": {
								Description:         "plural is the plural name of the resource to serve. The custom resources are served under '/apis/<group>/<version>/.../<plural>'. Must match the name of the CustomResourceDefinition (in the form '<names.plural>.<group>'). Must be all lowercase.",
								MarkdownDescription: "plural is the plural name of the resource to serve. The custom resources are served under '/apis/<group>/<version>/.../<plural>'. Must match the name of the CustomResourceDefinition (in the form '<names.plural>.<group>'). Must be all lowercase.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"short_names": {
								Description:         "shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like 'kubectl get <shortname>'. It must be all lowercase.",
								MarkdownDescription: "shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like 'kubectl get <shortname>'. It must be all lowercase.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"singular": {
								Description:         "singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased 'kind'.",
								MarkdownDescription: "singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased 'kind'.",

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

					"connection_secret_keys": {
						Description:         "ConnectionSecretKeys is the list of keys that will be exposed to the end user of the defined kind. If the list is empty, all keys will be published.",
						MarkdownDescription: "ConnectionSecretKeys is the list of keys that will be exposed to the end user of the defined kind. If the list is empty, all keys will be published.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"default_composition_ref": {
						Description:         "DefaultCompositionRef refers to the Composition resource that will be used in case no composition selector is given.",
						MarkdownDescription: "DefaultCompositionRef refers to the Composition resource that will be used in case no composition selector is given.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the Composition.",
								MarkdownDescription: "Name of the Composition.",

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

					"enforced_composition_ref": {
						Description:         "EnforcedCompositionRef refers to the Composition resource that will be used by all composite instances whose schema is defined by this definition.",
						MarkdownDescription: "EnforcedCompositionRef refers to the Composition resource that will be used by all composite instances whose schema is defined by this definition.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the Composition.",
								MarkdownDescription: "Name of the Composition.",

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

					"group": {
						Description:         "Group specifies the API group of the defined composite resource. Composite resources are served under '/apis/<group>/...'. Must match the name of the XRD (in the form '<names.plural>.<group>').",
						MarkdownDescription: "Group specifies the API group of the defined composite resource. Composite resources are served under '/apis/<group>/...'. Must match the name of the XRD (in the form '<names.plural>.<group>').",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"names": {
						Description:         "Names specifies the resource and kind names of the defined composite resource.",
						MarkdownDescription: "Names specifies the resource and kind names of the defined composite resource.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"categories": {
								Description:         "categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like 'kubectl get all'.",
								MarkdownDescription: "categories is a list of grouped resources this custom resource belongs to (e.g. 'all'). This is published in API discovery documents, and used by clients to support invocations like 'kubectl get all'.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the 'kind' attribute in API calls.",
								MarkdownDescription: "kind is the serialized kind of the resource. It is normally CamelCase and singular. Custom resource instances will use this value as the 'kind' attribute in API calls.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"list_kind": {
								Description:         "listKind is the serialized kind of the list for this resource. Defaults to ''kind'List'.",
								MarkdownDescription: "listKind is the serialized kind of the list for this resource. Defaults to ''kind'List'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"plural": {
								Description:         "plural is the plural name of the resource to serve. The custom resources are served under '/apis/<group>/<version>/.../<plural>'. Must match the name of the CustomResourceDefinition (in the form '<names.plural>.<group>'). Must be all lowercase.",
								MarkdownDescription: "plural is the plural name of the resource to serve. The custom resources are served under '/apis/<group>/<version>/.../<plural>'. Must match the name of the CustomResourceDefinition (in the form '<names.plural>.<group>'). Must be all lowercase.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"short_names": {
								Description:         "shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like 'kubectl get <shortname>'. It must be all lowercase.",
								MarkdownDescription: "shortNames are short names for the resource, exposed in API discovery documents, and used by clients to support invocations like 'kubectl get <shortname>'. It must be all lowercase.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"singular": {
								Description:         "singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased 'kind'.",
								MarkdownDescription: "singular is the singular name of the resource. It must be all lowercase. Defaults to lowercased 'kind'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"versions": {
						Description:         "Versions is the list of all API versions of the defined composite resource. Version names are used to compute the order in which served versions are listed in API discovery. If the version string is 'kube-like', it will sort above non 'kube-like' version strings, which are ordered lexicographically. 'Kube-like' versions start with a 'v', then are followed by a number (the major version), then optionally the string 'alpha' or 'beta' and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10. Note that all versions must have identical schemas; Crossplane does not currently support conversion between different version schemas.",
						MarkdownDescription: "Versions is the list of all API versions of the defined composite resource. Version names are used to compute the order in which served versions are listed in API discovery. If the version string is 'kube-like', it will sort above non 'kube-like' version strings, which are ordered lexicographically. 'Kube-like' versions start with a 'v', then are followed by a number (the major version), then optionally the string 'alpha' or 'beta' and another number (the minor version). These are sorted first by GA > beta > alpha (where GA is a version with no suffix such as beta or alpha), and then by comparing major version, then minor version. An example sorted list of versions: v10, v2, v1, v11beta2, v10beta3, v3beta1, v12alpha1, v11alpha2, foo1, foo10. Note that all versions must have identical schemas; Crossplane does not currently support conversion between different version schemas.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"additional_printer_columns": {
								Description:         "AdditionalPrinterColumns specifies additional columns returned in Table output. If no columns are specified, a single column displaying the age of the custom resource is used. See the following link for details: https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables",
								MarkdownDescription: "AdditionalPrinterColumns specifies additional columns returned in Table output. If no columns are specified, a single column displaying the age of the custom resource is used. See the following link for details: https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"description": {
										Description:         "description is a human readable description of this column.",
										MarkdownDescription: "description is a human readable description of this column.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"format": {
										Description:         "format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.",
										MarkdownDescription: "format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"json_path": {
										Description:         "jsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column.",
										MarkdownDescription: "jsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "name is a human readable name for the column.",
										MarkdownDescription: "name is a human readable name for the column.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"priority": {
										Description:         "priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.",
										MarkdownDescription: "priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.",
										MarkdownDescription: "type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.",

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

							"deprecated": {
								Description:         "The deprecated field specifies that this version is deprecated and should not be used.",
								MarkdownDescription: "The deprecated field specifies that this version is deprecated and should not be used.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"deprecation_warning": {
								Description:         "DeprecationWarning specifies the message that should be shown to the user when using this version.",
								MarkdownDescription: "DeprecationWarning specifies the message that should be shown to the user when using this version.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of this version, e.g. “v1”, “v2beta1”, etc. Composite resources are served under this version at '/apis/<group>/<version>/...' if 'served' is true.",
								MarkdownDescription: "Name of this version, e.g. “v1”, “v2beta1”, etc. Composite resources are served under this version at '/apis/<group>/<version>/...' if 'served' is true.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"referenceable": {
								Description:         "Referenceable specifies that this version may be referenced by a Composition in order to configure which resources an XR may be composed of. Exactly one version must be marked as referenceable; all Compositions must target only the referenceable version. The referenceable version must be served.",
								MarkdownDescription: "Referenceable specifies that this version may be referenced by a Composition in order to configure which resources an XR may be composed of. Exactly one version must be marked as referenceable; all Compositions must target only the referenceable version. The referenceable version must be served.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"schema": {
								Description:         "Schema describes the schema used for validation, pruning, and defaulting of this version of the defined composite resource. Fields required by all composite resources will be injected into this schema automatically, and will override equivalently named fields in this schema. Omitting this schema results in a schema that contains only the fields required by all composite resources.",
								MarkdownDescription: "Schema describes the schema used for validation, pruning, and defaulting of this version of the defined composite resource. Fields required by all composite resources will be injected into this schema automatically, and will override equivalently named fields in this schema. Omitting this schema results in a schema that contains only the fields required by all composite resources.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"open_apiv3_schema": {
										Description:         "OpenAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.",
										MarkdownDescription: "OpenAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"served": {
								Description:         "Served specifies that this version should be served via REST APIs.",
								MarkdownDescription: "Served specifies that this version should be served via REST APIs.",

								Type: types.BoolType,

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
		},
	}, nil
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apiextensions_crossplane_io_composite_resource_definition_v1")

	var state ApiextensionsCrossplaneIoCompositeResourceDefinitionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ApiextensionsCrossplaneIoCompositeResourceDefinitionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apiextensions.crossplane.io/v1")
	goModel.Kind = utilities.Ptr("CompositeResourceDefinition")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apiextensions_crossplane_io_composite_resource_definition_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apiextensions_crossplane_io_composite_resource_definition_v1")

	var state ApiextensionsCrossplaneIoCompositeResourceDefinitionV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ApiextensionsCrossplaneIoCompositeResourceDefinitionV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apiextensions.crossplane.io/v1")
	goModel.Kind = utilities.Ptr("CompositeResourceDefinition")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *ApiextensionsCrossplaneIoCompositeResourceDefinitionV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apiextensions_crossplane_io_composite_resource_definition_v1")
	// NO-OP: Terraform removes the state automatically for us
}
