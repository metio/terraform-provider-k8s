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
	"strings"
	"time"
)

var (
	_ resource.Resource                = &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource{}
)

func NewAutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource() resource.Resource {
	return &AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource{}
}

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitForUpsert  types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete  types.Object `tfsdk:"wait_for_delete" json:"-"`

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

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1"
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"wait_for_upsert": schema.ListNestedAttribute{
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
						"poll_interval": schema.StringAttribute{
							Description:         "The length of time to wait before checking again.",
							MarkdownDescription: "The length of time to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("5s"),
						},
					},
				},
			},

			"wait_for_delete": schema.SingleNestedAttribute{
				Description:         "Wait for deletion of resources.",
				MarkdownDescription: "Wait for deletion of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"timeout": schema.StringAttribute{
						Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("30s"),
					},
					"poll_interval": schema.StringAttribute{
						Description:         "The length of time to wait before checking again.",
						MarkdownDescription: "The length of time to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             stringdefault.StaticString("5s"),
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
						Required: false,
						Optional: true,
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

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedResourceDataError(request.ProviderData))
	}
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")

	var model AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("autoscaling.k8s.elastic.co/v1alpha1")
	model.Kind = pointer.String("ElasticsearchAutoscaler")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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
		Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.k8s.elastic.co", Version: "v1alpha1", Resource: "elasticsearchautoscalers"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")

	var data AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.k8s.elastic.co", Version: "v1alpha1", Resource: "elasticsearchautoscalers"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")

	var model AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("autoscaling.k8s.elastic.co/v1alpha1")
	model.Kind = pointer.String("ElasticsearchAutoscaler")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonMarshalError(err))
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
		Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.k8s.elastic.co", Version: "v1alpha1", Resource: "elasticsearchautoscalers"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.Append(utilities.PatchError(err))
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_autoscaling_k8s_elastic_co_elasticsearch_autoscaler_v1alpha1")

	var data AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.k8s.elastic.co", Version: "v1alpha1", Resource: "elasticsearchautoscalers"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "autoscaling.k8s.elastic.co", Version: "v1alpha1", Resource: "elasticsearchautoscalers"}).
				Namespace(data.Metadata.Namespace).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout == time.Second*0 {
				break
			}
			if time.Now().After(startTime.Add(timeout)) {
				response.Diagnostics.Append(utilities.WaitTimeoutExceeded())
				return
			}
			time.Sleep(pollInterval)
		}
	}
}

func (r *AutoscalingK8SElasticCoElasticsearchAutoscalerV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
