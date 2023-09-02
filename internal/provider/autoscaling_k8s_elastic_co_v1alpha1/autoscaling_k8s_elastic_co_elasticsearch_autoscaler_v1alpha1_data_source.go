/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package autoscaling_k8s_elastic_co_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource{}
)

func NewAutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource() datasource.DataSource {
	return &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource{}
}

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1"
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ElasticsearchAutoscaler represents an ElasticsearchAutoscaler resource in a Kubernetes cluster.",
		MarkdownDescription: "ElasticsearchAutoscaler represents an ElasticsearchAutoscaler resource in a Kubernetes cluster.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "Name identifies the autoscaling policy in the autoscaling specification.",
									MarkdownDescription: "Name identifies the autoscaling policy in the autoscaling specification.",
									Required:            false,
									Optional:            false,
									Computed:            true,
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
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min": schema.StringAttribute{
													Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"requests_to_limits_ratio": schema.StringAttribute{
													Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"memory": schema.SingleNestedAttribute{
											Description:         "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											MarkdownDescription: "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											Attributes: map[string]schema.Attribute{
												"max": schema.StringAttribute{
													Description:         "Max represents the upper limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Max represents the upper limit for the resources managed by the autoscaler.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min": schema.StringAttribute{
													Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"requests_to_limits_ratio": schema.StringAttribute{
													Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"node_count": schema.SingleNestedAttribute{
											Description:         "NodeCountRange is used to model the minimum and the maximum number of nodes over all the NodeSets managed by the same autoscaling policy.",
											MarkdownDescription: "NodeCountRange is used to model the minimum and the maximum number of nodes over all the NodeSets managed by the same autoscaling policy.",
											Attributes: map[string]schema.Attribute{
												"max": schema.Int64Attribute{
													Description:         "Max represents the maximum number of nodes in a tier.",
													MarkdownDescription: "Max represents the maximum number of nodes in a tier.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min": schema.Int64Attribute{
													Description:         "Min represents the minimum number of nodes in a tier.",
													MarkdownDescription: "Min represents the minimum number of nodes in a tier.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"storage": schema.SingleNestedAttribute{
											Description:         "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											MarkdownDescription: "QuantityRange models a resource limit range for resources which can be expressed with resource.Quantity.",
											Attributes: map[string]schema.Attribute{
												"max": schema.StringAttribute{
													Description:         "Max represents the upper limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Max represents the upper limit for the resources managed by the autoscaler.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min": schema.StringAttribute{
													Description:         "Min represents the lower limit for the resources managed by the autoscaler.",
													MarkdownDescription: "Min represents the lower limit for the resources managed by the autoscaler.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"requests_to_limits_ratio": schema.StringAttribute{
													Description:         "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													MarkdownDescription: "RequestsToLimitsRatio allows to customize Kubernetes resource Limit based on the Request.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"roles": schema.ListAttribute{
									Description:         "An autoscaling policy must target a unique set of roles.",
									MarkdownDescription: "An autoscaling policy must target a unique set of roles.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"polling_period": schema.StringAttribute{
						Description:         "PollingPeriod is the period at which to synchronize with the Elasticsearch autoscaling API.",
						MarkdownDescription: "PollingPeriod is the period at which to synchronize with the Elasticsearch autoscaling API.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")

	var data AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.k8s.elastic.co", Version: "v1alpha1", Resource: "ElasticsearchAutoscaler"}).
		Namespace(data.Metadata.Namespace).
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

	var readResponse AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1DataSourceData
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

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("autoscaling.k8s.elastic.co/v1alpha1")
	data.Kind = pointer.String("ElasticsearchAutoscaler")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
