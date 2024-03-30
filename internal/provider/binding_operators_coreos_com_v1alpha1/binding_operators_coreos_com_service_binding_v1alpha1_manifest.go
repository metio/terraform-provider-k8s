/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package binding_operators_coreos_com_v1alpha1

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
	_ datasource.DataSource = &BindingOperatorsCoreosComServiceBindingV1Alpha1Manifest{}
)

func NewBindingOperatorsCoreosComServiceBindingV1Alpha1Manifest() datasource.DataSource {
	return &BindingOperatorsCoreosComServiceBindingV1Alpha1Manifest{}
}

type BindingOperatorsCoreosComServiceBindingV1Alpha1Manifest struct{}

type BindingOperatorsCoreosComServiceBindingV1Alpha1ManifestData struct {
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
		Application *struct {
			BindingPath *struct {
				ContainersPath *string `tfsdk:"containers_path" json:"containersPath,omitempty"`
				SecretPath     *string `tfsdk:"secret_path" json:"secretPath,omitempty"`
			} `tfsdk:"binding_path" json:"bindingPath,omitempty"`
			Group         *string `tfsdk:"group" json:"group,omitempty"`
			Kind          *string `tfsdk:"kind" json:"kind,omitempty"`
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Resource *string `tfsdk:"resource" json:"resource,omitempty"`
			Version  *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"application" json:"application,omitempty"`
		BindAsFiles            *bool `tfsdk:"bind_as_files" json:"bindAsFiles,omitempty"`
		DetectBindingResources *bool `tfsdk:"detect_binding_resources" json:"detectBindingResources,omitempty"`
		Mappings               *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"mappings" json:"mappings,omitempty"`
		Name           *string `tfsdk:"name" json:"name,omitempty"`
		NamingStrategy *string `tfsdk:"naming_strategy" json:"namingStrategy,omitempty"`
		Services       *[]struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Id        *string `tfsdk:"id" json:"id,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Resource  *string `tfsdk:"resource" json:"resource,omitempty"`
			Version   *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_binding_operators_coreos_com_service_binding_v1alpha1_manifest"
}

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "ServiceBindingSpec defines the desired state of ServiceBinding.",
				MarkdownDescription: "ServiceBindingSpec defines the desired state of ServiceBinding.",
				Attributes: map[string]schema.Attribute{
					"application": schema.SingleNestedAttribute{
						Description:         "Application identifies the application connecting to the backing service.",
						MarkdownDescription: "Application identifies the application connecting to the backing service.",
						Attributes: map[string]schema.Attribute{
							"binding_path": schema.SingleNestedAttribute{
								Description:         "BindingPath refers to the paths in the application workload's schema where the binding workload would be referenced.  If BindingPath is not specified, then the default path locations are used.  The default location for ContainersPath is 'spec.template.spec.containers'.  If SecretPath is not specified, then the name of the secret object does not need to be specified.",
								MarkdownDescription: "BindingPath refers to the paths in the application workload's schema where the binding workload would be referenced.  If BindingPath is not specified, then the default path locations are used.  The default location for ContainersPath is 'spec.template.spec.containers'.  If SecretPath is not specified, then the name of the secret object does not need to be specified.",
								Attributes: map[string]schema.Attribute{
									"containers_path": schema.StringAttribute{
										Description:         "ContainersPath defines the path to the corev1.Containers reference. If BindingPath is not specified, the default location is 'spec.template.spec.containers'.",
										MarkdownDescription: "ContainersPath defines the path to the corev1.Containers reference. If BindingPath is not specified, the default location is 'spec.template.spec.containers'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_path": schema.StringAttribute{
										Description:         "SecretPath defines the path to a string field where the name of the secret object is going to be assigned.  Note: The name of the secret object is same as that of the name of service binding custom resource (metadata.name).",
										MarkdownDescription: "SecretPath defines the path to a string field where the name of the secret object is going to be assigned.  Note: The name of the secret object is same as that of the name of service binding custom resource (metadata.name).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"group": schema.StringAttribute{
								Description:         "Group of the referent.",
								MarkdownDescription: "Group of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.",
								MarkdownDescription: "Kind of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label_selector": schema.SingleNestedAttribute{
								Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
								MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
								Attributes: map[string]schema.Attribute{
									"match_expressions": schema.ListNestedAttribute{
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "key is the label key that the selector applies to.",
													MarkdownDescription: "key is the label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"values": schema.ListAttribute{
													Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

									"match_labels": schema.MapAttribute{
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource": schema.StringAttribute{
								Description:         "Resource of the referent.",
								MarkdownDescription: "Resource of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version of the referent.",
								MarkdownDescription: "Version of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"bind_as_files": schema.BoolAttribute{
						Description:         "BindAsFiles makes the binding values available as files in the application's container.  By default, values are mounted under the path '/bindings'; this can be changed by setting the SERVICE_BINDING_ROOT environment variable.",
						MarkdownDescription: "BindAsFiles makes the binding values available as files in the application's container.  By default, values are mounted under the path '/bindings'; this can be changed by setting the SERVICE_BINDING_ROOT environment variable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"detect_binding_resources": schema.BoolAttribute{
						Description:         "DetectBindingResources is a flag that, when set to true, will cause SBO to search for binding information in the owned resources of the specified services.  If this binding information exists, then the application is bound to these subresources.",
						MarkdownDescription: "DetectBindingResources is a flag that, when set to true, will cause SBO to search for binding information in the owned resources of the specified services.  If this binding information exists, then the application is bound to these subresources.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"mappings": schema.ListNestedAttribute{
						Description:         "Mappings specifies custom mappings.",
						MarkdownDescription: "Mappings specifies custom mappings.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name is the name of new binding.",
									MarkdownDescription: "Name is the name of new binding.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value specificies a go template that will be rendered and injected into the application.",
									MarkdownDescription: "Value specificies a go template that will be rendered and injected into the application.",
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

					"name": schema.StringAttribute{
						Description:         "Name is the name of the service as projected into the workload container.  Defaults to .metadata.name.",
						MarkdownDescription: "Name is the name of the service as projected into the workload container.  Defaults to .metadata.name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9\-\.]*$`), ""),
						},
					},

					"naming_strategy": schema.StringAttribute{
						Description:         "NamingStrategy defines custom string template for preparing binding names.  It can be set to pre-defined strategies: 'none', 'lowercase', or 'uppercase'.  Otherwise, it is treated as a custom go template, and it is handled accordingly.",
						MarkdownDescription: "NamingStrategy defines custom string template for preparing binding names.  It can be set to pre-defined strategies: 'none', 'lowercase', or 'uppercase'.  Otherwise, it is treated as a custom go template, and it is handled accordingly.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"services": schema.ListNestedAttribute{
						Description:         "Services indicates the backing services to be connected to by an application.  At least one service must be specified.",
						MarkdownDescription: "Services indicates the backing services to be connected to by an application.  At least one service must be specified.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group of the referent.",
									MarkdownDescription: "Group of the referent.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the referent.",
									MarkdownDescription: "Kind of the referent.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the referent.",
									MarkdownDescription: "Name of the referent.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the referent.  If unspecified, assumes the same namespace as ServiceBinding.",
									MarkdownDescription: "Namespace of the referent.  If unspecified, assumes the same namespace as ServiceBinding.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resource": schema.StringAttribute{
									Description:         "Resource of the referent.",
									MarkdownDescription: "Resource of the referent.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"version": schema.StringAttribute{
									Description:         "Version of the referent.",
									MarkdownDescription: "Version of the referent.",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *BindingOperatorsCoreosComServiceBindingV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_binding_operators_coreos_com_service_binding_v1alpha1_manifest")

	var model BindingOperatorsCoreosComServiceBindingV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("binding.operators.coreos.com/v1alpha1")
	model.Kind = pointer.String("ServiceBinding")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
