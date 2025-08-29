/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package fluxcd_controlplane_io_v1

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
	_ datasource.DataSource = &FluxcdControlplaneIoResourceSetV1Manifest{}
)

func NewFluxcdControlplaneIoResourceSetV1Manifest() datasource.DataSource {
	return &FluxcdControlplaneIoResourceSetV1Manifest{}
}

type FluxcdControlplaneIoResourceSetV1Manifest struct{}

type FluxcdControlplaneIoResourceSetV1ManifestData struct {
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
		CommonMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"common_metadata" json:"commonMetadata,omitempty"`
		DependsOn *[]struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
			Ready      *bool   `tfsdk:"ready" json:"ready,omitempty"`
			ReadyExpr  *string `tfsdk:"ready_expr" json:"readyExpr,omitempty"`
		} `tfsdk:"depends_on" json:"dependsOn,omitempty"`
		Inputs *[]struct {
		} `tfsdk:"inputs" json:"inputs,omitempty"`
		InputsFrom *[]struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Selector   *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"inputs_from" json:"inputsFrom,omitempty"`
		Resources          *[]string `tfsdk:"resources" json:"resources,omitempty"`
		ResourcesTemplate  *string   `tfsdk:"resources_template" json:"resourcesTemplate,omitempty"`
		ServiceAccountName *string   `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		Wait               *bool     `tfsdk:"wait" json:"wait,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FluxcdControlplaneIoResourceSetV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_fluxcd_controlplane_io_resource_set_v1_manifest"
}

func (r *FluxcdControlplaneIoResourceSetV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResourceSet is the Schema for the ResourceSets API.",
		MarkdownDescription: "ResourceSet is the Schema for the ResourceSets API.",
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
				Description:         "ResourceSetSpec defines the desired state of ResourceSet",
				MarkdownDescription: "ResourceSetSpec defines the desired state of ResourceSet",
				Attributes: map[string]schema.Attribute{
					"common_metadata": schema.SingleNestedAttribute{
						Description:         "CommonMetadata specifies the common labels and annotations that are applied to all resources. Any existing label or annotation will be overridden if its key matches a common one.",
						MarkdownDescription: "CommonMetadata specifies the common labels and annotations that are applied to all resources. Any existing label or annotation will be overridden if its key matches a common one.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to be added to the object's metadata.",
								MarkdownDescription: "Annotations to be added to the object's metadata.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels to be added to the object's metadata.",
								MarkdownDescription: "Labels to be added to the object's metadata.",
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

					"depends_on": schema.ListNestedAttribute{
						Description:         "DependsOn specifies the list of Kubernetes resources that must exist on the cluster before the reconciliation process starts.",
						MarkdownDescription: "DependsOn specifies the list of Kubernetes resources that must exist on the cluster before the reconciliation process starts.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion of the resource to depend on.",
									MarkdownDescription: "APIVersion of the resource to depend on.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the resource to depend on.",
									MarkdownDescription: "Kind of the resource to depend on.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the resource to depend on.",
									MarkdownDescription: "Name of the resource to depend on.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespace of the resource to depend on.",
									MarkdownDescription: "Namespace of the resource to depend on.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ready": schema.BoolAttribute{
									Description:         "Ready checks if the resource Ready status condition is true.",
									MarkdownDescription: "Ready checks if the resource Ready status condition is true.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ready_expr": schema.StringAttribute{
									Description:         "ReadyExpr checks if the resource satisfies the given CEL expression. The expression replaces the default readiness check and is only evaluated if Ready is set to 'true'.",
									MarkdownDescription: "ReadyExpr checks if the resource satisfies the given CEL expression. The expression replaces the default readiness check and is only evaluated if Ready is set to 'true'.",
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

					"inputs": schema.ListNestedAttribute{
						Description:         "Inputs contains the list of ResourceSet inputs.",
						MarkdownDescription: "Inputs contains the list of ResourceSet inputs.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"inputs_from": schema.ListNestedAttribute{
						Description:         "InputsFrom contains the list of references to input providers. When set, the inputs are fetched from the providers and concatenated with the in-line inputs defined in the ResourceSet.",
						MarkdownDescription: "InputsFrom contains the list of references to input providers. When set, the inputs are fetched from the providers and concatenated with the in-line inputs defined in the ResourceSet.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"api_version": schema.StringAttribute{
									Description:         "APIVersion of the input provider resource. When not set, the APIVersion of the ResourceSet is used.",
									MarkdownDescription: "APIVersion of the input provider resource. When not set, the APIVersion of the ResourceSet is used.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "Kind of the input provider resource.",
									MarkdownDescription: "Kind of the input provider resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("ResourceSetInputProvider"),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name of the input provider resource. Cannot be set when the Selector field is set.",
									MarkdownDescription: "Name of the input provider resource. Cannot be set when the Selector field is set.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"selector": schema.SingleNestedAttribute{
									Description:         "Selector is a label selector to filter the input provider resources as an alternative to the Name field.",
									MarkdownDescription: "Selector is a label selector to filter the input provider resources as an alternative to the Name field.",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": schema.ListAttribute{
						Description:         "Resources contains the list of Kubernetes resources to reconcile.",
						MarkdownDescription: "Resources contains the list of Kubernetes resources to reconcile.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources_template": schema.StringAttribute{
						Description:         "ResourcesTemplate is a Go template that generates the list of Kubernetes resources to reconcile. The template is rendered as multi-document YAML, the resources should be separated by '---'. When both Resources and ResourcesTemplate are set, the resulting objects are merged and deduplicated, with the ones from Resources taking precedence.",
						MarkdownDescription: "ResourcesTemplate is a Go template that generates the list of Kubernetes resources to reconcile. The template is rendered as multi-document YAML, the resources should be separated by '---'. When both Resources and ResourcesTemplate are set, the resulting objects are merged and deduplicated, with the ones from Resources taking precedence.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "The name of the Kubernetes service account to impersonate when reconciling the generated resources.",
						MarkdownDescription: "The name of the Kubernetes service account to impersonate when reconciling the generated resources.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"wait": schema.BoolAttribute{
						Description:         "Wait instructs the controller to check the health of all the reconciled resources.",
						MarkdownDescription: "Wait instructs the controller to check the health of all the reconciled resources.",
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

func (r *FluxcdControlplaneIoResourceSetV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_fluxcd_controlplane_io_resource_set_v1_manifest")

	var model FluxcdControlplaneIoResourceSetV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("fluxcd.controlplane.io/v1")
	model.Kind = pointer.String("ResourceSet")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
