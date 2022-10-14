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

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource)(nil)
)

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1GoModel struct {
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
		ElasticsearchRef *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"elasticsearch_ref" yaml:"elasticsearchRef,omitempty"`

		Policies *[]struct {
			Deciders *map[string]map[string]string `tfsdk:"deciders" yaml:"deciders,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Resources *struct {
				Cpu *struct {
					Max *string `tfsdk:"max" yaml:"max,omitempty"`

					Min *string `tfsdk:"min" yaml:"min,omitempty"`

					RequestsToLimitsRatio *string `tfsdk:"requests_to_limits_ratio" yaml:"requestsToLimitsRatio,omitempty"`
				} `tfsdk:"cpu" yaml:"cpu,omitempty"`

				Memory *struct {
					Max *string `tfsdk:"max" yaml:"max,omitempty"`

					Min *string `tfsdk:"min" yaml:"min,omitempty"`

					RequestsToLimitsRatio *string `tfsdk:"requests_to_limits_ratio" yaml:"requestsToLimitsRatio,omitempty"`
				} `tfsdk:"memory" yaml:"memory,omitempty"`

				NodeCount *struct {
					Max *int64 `tfsdk:"max" yaml:"max,omitempty"`

					Min *int64 `tfsdk:"min" yaml:"min,omitempty"`
				} `tfsdk:"node_count" yaml:"nodeCount,omitempty"`

				Storage *struct {
					Max *string `tfsdk:"max" yaml:"max,omitempty"`

					Min *string `tfsdk:"min" yaml:"min,omitempty"`

					RequestsToLimitsRatio *string `tfsdk:"requests_to_limits_ratio" yaml:"requestsToLimitsRatio,omitempty"`
				} `tfsdk:"storage" yaml:"storage,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`
		} `tfsdk:"policies" yaml:"policies,omitempty"`

		PollingPeriod *string `tfsdk:"polling_period" yaml:"pollingPeriod,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource() resource.Resource {
	return &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource{}
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1"
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ElasticsearchAutoscaler represents an ElasticsearchAutoscaler resource in a Kubernetes cluster.",
		MarkdownDescription: "ElasticsearchAutoscaler represents an ElasticsearchAutoscaler resource in a Kubernetes cluster.",
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
				Description:         "ElasticsearchAutoscalerSpec holds the specification of an Elasticsearch autoscaler resource.",
				MarkdownDescription: "ElasticsearchAutoscalerSpec holds the specification of an Elasticsearch autoscaler resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"elasticsearch_ref": {
						Description:         "ElasticsearchRef is a reference to an Elasticsearch cluster that exists in the same namespace.",
						MarkdownDescription: "ElasticsearchRef is a reference to an Elasticsearch cluster that exists in the same namespace.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name is the name of the Elasticsearch resource to scale automatically.",
								MarkdownDescription: "Name is the name of the Elasticsearch resource to scale automatically.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"policies": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"deciders": {
								Description:         "Deciders allow the user to override default settings for autoscaling deciders.",
								MarkdownDescription: "Deciders allow the user to override default settings for autoscaling deciders.",

								Type: types.MapType{ElemType: types.MapType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name identifies the autoscaling policy in the autoscaling specification.",
								MarkdownDescription: "Name identifies the autoscaling policy in the autoscaling specification.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "AutoscalingResources model the limits, submitted by the user, for the supported resources in an autoscaling policy. Only the node count range is mandatory. For other resources, a limit range is required only if the Elasticsearch autoscaling capacity API returns a requirement for a given resource. For example, the memory limit range is only required if the autoscaling API response contains a memory requirement. If there is no limit range for a resource, and if that resource is not mandatory, then the resources in the NodeSets managed by the autoscaling policy are left untouched.",
								MarkdownDescription: "AutoscalingResources model the limits, submitted by the user, for the supported resources in an autoscaling policy. Only the node count range is mandatory. For other resources, a limit range is required only if the Elasticsearch autoscaling capacity API returns a requirement for a given resource. For example, the memory limit range is only required if the autoscaling API response contains a memory requirement. If there is no limit range for a resource, and if that resource is not mandatory, then the resources in the NodeSets managed by the autoscaling policy are left untouched.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cpu": {
										Description:         "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
										MarkdownDescription: "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max": {
												Description:         "Max represents the upper limit for the resources managed by the autoscaler.",
												MarkdownDescription: "Max represents the upper limit for the resources managed by the autoscaler.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"min": {
												Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
												MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"requests_to_limits_ratio": {
												Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
												MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",

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

									"memory": {
										Description:         "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
										MarkdownDescription: "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max": {
												Description:         "Max represents the upper limit for the resources managed by the autoscaler.",
												MarkdownDescription: "Max represents the upper limit for the resources managed by the autoscaler.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"min": {
												Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
												MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"requests_to_limits_ratio": {
												Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
												MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",

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

									"node_count": {
										Description:         "NodeCountRange is used to model the minimum and the maximum number of nodes over all the NodeSets managed by the same autoscaling policy.",
										MarkdownDescription: "NodeCountRange is used to model the minimum and the maximum number of nodes over all the NodeSets managed by the same autoscaling policy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max": {
												Description:         "Max represents the maximum number of nodes in a tier.",
												MarkdownDescription: "Max represents the maximum number of nodes in a tier.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"min": {
												Description:         "Min represents the minimum number of nodes in a tier.",
												MarkdownDescription: "Min represents the minimum number of nodes in a tier.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"storage": {
										Description:         "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
										MarkdownDescription: "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max": {
												Description:         "Max represents the upper limit for the resources managed by the autoscaler.",
												MarkdownDescription: "Max represents the upper limit for the resources managed by the autoscaler.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"min": {
												Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
												MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"requests_to_limits_ratio": {
												Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
												MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",

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
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"roles": {
								Description:         "An autoscaling policy must target a unique set of roles.",
								MarkdownDescription: "An autoscaling policy must target a unique set of roles.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"polling_period": {
						Description:         "PollingPeriod is the period at which to synchronize with the Elasticsearch autoscaling API.",
						MarkdownDescription: "PollingPeriod is the period at which to synchronize with the Elasticsearch autoscaling API.",

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

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")

	var state AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("autoscaling.k8s.elastic.co/v1alpha1")
	goModel.Kind = utilities.Ptr("ElasticsearchAutoscaler")

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

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")

	var state AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("autoscaling.k8s.elastic.co/v1alpha1")
	goModel.Kind = utilities.Ptr("ElasticsearchAutoscaler")

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

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
