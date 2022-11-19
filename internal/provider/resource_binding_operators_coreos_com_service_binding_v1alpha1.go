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

type BindingOperatorsCoreosComServiceBindingV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*BindingOperatorsCoreosComServiceBindingV1Alpha1Resource)(nil)
)

type BindingOperatorsCoreosComServiceBindingV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type BindingOperatorsCoreosComServiceBindingV1Alpha1GoModel struct {
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
		Application *struct {
			BindingPath *struct {
				ContainersPath *string `tfsdk:"containers_path" yaml:"containersPath,omitempty"`

				SecretPath *string `tfsdk:"secret_path" yaml:"secretPath,omitempty"`
			} `tfsdk:"binding_path" yaml:"bindingPath,omitempty"`

			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"application" yaml:"application,omitempty"`

		BindAsFiles *bool `tfsdk:"bind_as_files" yaml:"bindAsFiles,omitempty"`

		DetectBindingResources *bool `tfsdk:"detect_binding_resources" yaml:"detectBindingResources,omitempty"`

		Mappings *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"mappings" yaml:"mappings,omitempty"`

		Name *string `tfsdk:"name" yaml:"name,omitempty"`

		NamingStrategy *string `tfsdk:"naming_strategy" yaml:"namingStrategy,omitempty"`

		Services *[]struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Id *string `tfsdk:"id" yaml:"id,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"services" yaml:"services,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewBindingOperatorsCoreosComServiceBindingV1Alpha1Resource() resource.Resource {
	return &BindingOperatorsCoreosComServiceBindingV1Alpha1Resource{}
}

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_binding_operators_coreos_com_service_binding_v1alpha1"
}

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "ServiceBindingSpec defines the desired state of ServiceBinding.",
				MarkdownDescription: "ServiceBindingSpec defines the desired state of ServiceBinding.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"application": {
						Description:         "Application identifies the application connecting to the backing service.",
						MarkdownDescription: "Application identifies the application connecting to the backing service.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"binding_path": {
								Description:         "BindingPath refers to the paths in the application workload's schema where the binding workload would be referenced.  If BindingPath is not specified, then the default path locations are used.  The default location for ContainersPath is 'spec.template.spec.containers'.  If SecretPath is not specified, then the name of the secret object does not need to be specified.",
								MarkdownDescription: "BindingPath refers to the paths in the application workload's schema where the binding workload would be referenced.  If BindingPath is not specified, then the default path locations are used.  The default location for ContainersPath is 'spec.template.spec.containers'.  If SecretPath is not specified, then the name of the secret object does not need to be specified.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"containers_path": {
										Description:         "ContainersPath defines the path to the corev1.Containers reference. If BindingPath is not specified, the default location is 'spec.template.spec.containers'.",
										MarkdownDescription: "ContainersPath defines the path to the corev1.Containers reference. If BindingPath is not specified, the default location is 'spec.template.spec.containers'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_path": {
										Description:         "SecretPath defines the path to a string field where the name of the secret object is going to be assigned.  Note: The name of the secret object is same as that of the name of service binding custom resource (metadata.name).",
										MarkdownDescription: "SecretPath defines the path to a string field where the name of the secret object is going to be assigned.  Note: The name of the secret object is same as that of the name of service binding custom resource (metadata.name).",

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
								Description:         "Group of the referent.",
								MarkdownDescription: "Group of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"label_selector": {
								Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
								MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

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

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

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

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource": {
								Description:         "Resource of the referent.",
								MarkdownDescription: "Resource of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Version of the referent.",
								MarkdownDescription: "Version of the referent.",

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

					"bind_as_files": {
						Description:         "BindAsFiles makes the binding values available as files in the application's container.  By default, values are mounted under the path '/bindings'; this can be changed by setting the SERVICE_BINDING_ROOT environment variable.",
						MarkdownDescription: "BindAsFiles makes the binding values available as files in the application's container.  By default, values are mounted under the path '/bindings'; this can be changed by setting the SERVICE_BINDING_ROOT environment variable.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"detect_binding_resources": {
						Description:         "DetectBindingResources is a flag that, when set to true, will cause SBO to search for binding information in the owned resources of the specified services.  If this binding information exists, then the application is bound to these subresources.",
						MarkdownDescription: "DetectBindingResources is a flag that, when set to true, will cause SBO to search for binding information in the owned resources of the specified services.  If this binding information exists, then the application is bound to these subresources.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mappings": {
						Description:         "Mappings specifies custom mappings.",
						MarkdownDescription: "Mappings specifies custom mappings.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name is the name of new binding.",
								MarkdownDescription: "Name is the name of new binding.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"value": {
								Description:         "Value specificies a go template that will be rendered and injected into the application.",
								MarkdownDescription: "Value specificies a go template that will be rendered and injected into the application.",

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

					"name": {
						Description:         "Name is the name of the service as projected into the workload container.  Defaults to .metadata.name.",
						MarkdownDescription: "Name is the name of the service as projected into the workload container.  Defaults to .metadata.name.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.LengthAtMost(253),

							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9\-\.]*$`), ""),
						},
					},

					"naming_strategy": {
						Description:         "NamingStrategy defines custom string template for preparing binding names.  It can be set to pre-defined strategies: 'none', 'lowercase', or 'uppercase'.  Otherwise, it is treated as a custom go template, and it is handled accordingly.",
						MarkdownDescription: "NamingStrategy defines custom string template for preparing binding names.  It can be set to pre-defined strategies: 'none', 'lowercase', or 'uppercase'.  Otherwise, it is treated as a custom go template, and it is handled accordingly.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"services": {
						Description:         "Services indicates the backing services to be connected to by an application.  At least one service must be specified.",
						MarkdownDescription: "Services indicates the backing services to be connected to by an application.  At least one service must be specified.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "Group of the referent.",
								MarkdownDescription: "Group of the referent.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent.  If unspecified, assumes the same namespace as ServiceBinding.",
								MarkdownDescription: "Namespace of the referent.  If unspecified, assumes the same namespace as ServiceBinding.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource": {
								Description:         "Resource of the referent.",
								MarkdownDescription: "Resource of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Version of the referent.",
								MarkdownDescription: "Version of the referent.",

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

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_binding_operators_coreos_com_service_binding_v1alpha1")

	var state BindingOperatorsCoreosComServiceBindingV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel BindingOperatorsCoreosComServiceBindingV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("binding.operators.coreos.com/v1alpha1")
	goModel.Kind = utilities.Ptr("ServiceBinding")

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

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_binding_operators_coreos_com_service_binding_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_binding_operators_coreos_com_service_binding_v1alpha1")

	var state BindingOperatorsCoreosComServiceBindingV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel BindingOperatorsCoreosComServiceBindingV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("binding.operators.coreos.com/v1alpha1")
	goModel.Kind = utilities.Ptr("ServiceBinding")

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

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_binding_operators_coreos_com_service_binding_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
