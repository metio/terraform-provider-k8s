/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package autoscaling_k8s_elastic_co_v1alpha1

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
	_ datasource.DataSource = &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Manifest{}
)

func NewAutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Manifest() datasource.DataSource {
	return &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Manifest{}
}

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Manifest struct{}

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ManifestData struct {
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
		ElasticsearchRef *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"elasticsearch_ref" json:"elasticsearchRef,omitempty"`
		Policies *[]struct {
			Deciders  *map[string]map[string]string `tfsdk:"deciders" json:"deciders,omitempty"`
			Name      *string                       `tfsdk:"name" json:"name,omitempty"`
			Resources *struct {
				Cpu *struct {
					Max                   *string `tfsdk:"max" json:"max,omitempty"`
					Min                   *string `tfsdk:"min" json:"min,omitempty"`
					RequestsToLimitsRatio *string `tfsdk:"requests_to_limits_ratio" json:"requestsToLimitsRatio,omitempty"`
				} `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *struct {
					Max                   *string `tfsdk:"max" json:"max,omitempty"`
					Min                   *string `tfsdk:"min" json:"min,omitempty"`
					RequestsToLimitsRatio *string `tfsdk:"requests_to_limits_ratio" json:"requestsToLimitsRatio,omitempty"`
				} `tfsdk:"memory" json:"memory,omitempty"`
				NodeCount *struct {
					Max *int64 `tfsdk:"max" json:"max,omitempty"`
					Min *int64 `tfsdk:"min" json:"min,omitempty"`
				} `tfsdk:"node_count" json:"nodeCount,omitempty"`
				Storage *struct {
					Max                   *string `tfsdk:"max" json:"max,omitempty"`
					Min                   *string `tfsdk:"min" json:"min,omitempty"`
					RequestsToLimitsRatio *string `tfsdk:"requests_to_limits_ratio" json:"requestsToLimitsRatio,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Roles *[]string `tfsdk:"roles" json:"roles,omitempty"`
		} `tfsdk:"policies" json:"policies,omitempty"`
		PollingPeriod *string `tfsdk:"polling_period" json:"pollingPeriod,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1_manifest"
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ElasticsearchAutoscaler represents an ElasticsearchAutoscaler resource in a Kubernetes cluster.",
		MarkdownDescription: "ElasticsearchAutoscaler represents an ElasticsearchAutoscaler resource in a Kubernetes cluster.",
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
				Description:         "ElasticsearchAutoscalerSpec holds the specification of an Elasticsearch autoscaler resource.",
				MarkdownDescription: "ElasticsearchAutoscalerSpec holds the specification of an Elasticsearch autoscaler resource.",
				Attributes: map[string]schema.Attribute{
					"elasticsearch_ref": schema.SingleNestedAttribute{
						Description:         "ElasticsearchRef is a reference to an Elasticsearch cluster that exists in the same namespace.",
						MarkdownDescription: "ElasticsearchRef is a reference to an Elasticsearch cluster that exists in the same namespace.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name is the name of the Elasticsearch resource to scale automatically.",
								MarkdownDescription: "Name is the name of the Elasticsearch resource to scale automatically.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"policies": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"deciders": schema.MapAttribute{
									Description:         "Deciders allow the user to override default settings for autoscaling deciders.",
									MarkdownDescription: "Deciders allow the user to override default settings for autoscaling deciders.",
									ElementType:         types.MapType{ElemType: types.StringType},
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name identifies the autoscaling policy in the autoscaling specification.",
									MarkdownDescription: "Name identifies the autoscaling policy in the autoscaling specification.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "AutoscalingResources model the limits, submitted by the user, for the supported resources in an autoscaling policy. Only the node count range is mandatory. For other resources, a limit range is required only if the Elasticsearch autoscaling capacity API returns a requirement for a given resource. For example, the memory limit range is only required if the autoscaling API response contains a memory requirement. If there is no limit range for a resource, and if that resource is not mandatory, then the resources in the NodeSets managed by the autoscaling policy are left untouched.",
									MarkdownDescription: "AutoscalingResources model the limits, submitted by the user, for the supported resources in an autoscaling policy. Only the node count range is mandatory. For other resources, a limit range is required only if the Elasticsearch autoscaling capacity API returns a requirement for a given resource. For example, the memory limit range is only required if the autoscaling API response contains a memory requirement. If there is no limit range for a resource, and if that resource is not mandatory, then the resources in the NodeSets managed by the autoscaling policy are left untouched.",
									Attributes: map[string]schema.Attribute{
										"cpu": schema.SingleNestedAttribute{
											Description:         "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											MarkdownDescription: "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											Attributes: map[string]schema.Attribute{
												"max": schema.StringAttribute{
													Description:         "Max represents the upper limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Max represents the upper limit for the resources managed by the autoscaler.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"min": schema.StringAttribute{
													Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"requests_to_limits_ratio": schema.StringAttribute{
													Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"memory": schema.SingleNestedAttribute{
											Description:         "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											MarkdownDescription: "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											Attributes: map[string]schema.Attribute{
												"max": schema.StringAttribute{
													Description:         "Max represents the upper limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Max represents the upper limit for the resources managed by the autoscaler.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"min": schema.StringAttribute{
													Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"requests_to_limits_ratio": schema.StringAttribute{
													Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"node_count": schema.SingleNestedAttribute{
											Description:         "NodeCountRange is used to model the minimum and the maximum number of nodes over all the NodeSets managed by the same autoscaling policy.",
											MarkdownDescription: "NodeCountRange is used to model the minimum and the maximum number of nodes over all the NodeSets managed by the same autoscaling policy.",
											Attributes: map[string]schema.Attribute{
												"max": schema.Int64Attribute{
													Description:         "Max represents the maximum number of nodes in a tier.",
													MarkdownDescription: "Max represents the maximum number of nodes in a tier.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"min": schema.Int64Attribute{
													Description:         "Min represents the minimum number of nodes in a tier.",
													MarkdownDescription: "Min represents the minimum number of nodes in a tier.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"storage": schema.SingleNestedAttribute{
											Description:         "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											MarkdownDescription: "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											Attributes: map[string]schema.Attribute{
												"max": schema.StringAttribute{
													Description:         "Max represents the upper limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Max represents the upper limit for the resources managed by the autoscaler.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"min": schema.StringAttribute{
													Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"requests_to_limits_ratio": schema.StringAttribute{
													Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
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
									Required: true,
									Optional: false,
									Computed: false,
								},

								"roles": schema.ListAttribute{
									Description:         "An autoscaling policy must target a unique set of roles.",
									MarkdownDescription: "An autoscaling policy must target a unique set of roles.",
									ElementType:         types.StringType,
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

					"polling_period": schema.StringAttribute{
						Description:         "PollingPeriod is the period at which to synchronize with the Elasticsearch autoscaling API.",
						MarkdownDescription: "PollingPeriod is the period at which to synchronize with the Elasticsearch autoscaling API.",
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

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1_manifest")

	var model AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("autoscaling.k8s.elastic.co/v1alpha1")
	model.Kind = pointer.String("ElasticsearchAutoscaler")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
