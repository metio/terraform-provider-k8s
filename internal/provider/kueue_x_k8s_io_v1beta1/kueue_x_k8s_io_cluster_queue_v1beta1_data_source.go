/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kueue_x_k8s_io_v1beta1

import (
	"context"
	"encoding/json"
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
	_ datasource.DataSource              = &KueueXK8SIoClusterQueueV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &KueueXK8SIoClusterQueueV1Beta1DataSource{}
)

func NewKueueXK8SIoClusterQueueV1Beta1DataSource() datasource.DataSource {
	return &KueueXK8SIoClusterQueueV1Beta1DataSource{}
}

type KueueXK8SIoClusterQueueV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KueueXK8SIoClusterQueueV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AdmissionChecks   *[]string `tfsdk:"admission_checks" json:"admissionChecks,omitempty"`
		Cohort            *string   `tfsdk:"cohort" json:"cohort,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		Preemption *struct {
			ReclaimWithinCohort *string `tfsdk:"reclaim_within_cohort" json:"reclaimWithinCohort,omitempty"`
			WithinClusterQueue  *string `tfsdk:"within_cluster_queue" json:"withinClusterQueue,omitempty"`
		} `tfsdk:"preemption" json:"preemption,omitempty"`
		QueueingStrategy *string `tfsdk:"queueing_strategy" json:"queueingStrategy,omitempty"`
		ResourceGroups   *[]struct {
			CoveredResources *[]string `tfsdk:"covered_resources" json:"coveredResources,omitempty"`
			Flavors          *[]struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Resources *[]struct {
					BorrowingLimit *string `tfsdk:"borrowing_limit" json:"borrowingLimit,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					NominalQuota   *string `tfsdk:"nominal_quota" json:"nominalQuota,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"flavors" json:"flavors,omitempty"`
		} `tfsdk:"resource_groups" json:"resourceGroups,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KueueXK8SIoClusterQueueV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kueue_x_k8s_io_cluster_queue_v1beta1"
}

func (r *KueueXK8SIoClusterQueueV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterQueue is the Schema for the clusterQueue API.",
		MarkdownDescription: "ClusterQueue is the Schema for the clusterQueue API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "ClusterQueueSpec defines the desired state of ClusterQueue",
				MarkdownDescription: "ClusterQueueSpec defines the desired state of ClusterQueue",
				Attributes: map[string]schema.Attribute{
					"admission_checks": schema.ListAttribute{
						Description:         "admissionChecks lists the AdmissionChecks required by this ClusterQueue",
						MarkdownDescription: "admissionChecks lists the AdmissionChecks required by this ClusterQueue",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cohort": schema.StringAttribute{
						Description:         "cohort that this ClusterQueue belongs to. CQs that belong to the same cohort can borrow unused resources from each other.  A CQ can be a member of a single borrowing cohort. A workload submitted to a queue referencing this CQ can borrow quota from any CQ in the cohort. Only quota for the [resource, flavor] pairs listed in the CQ can be borrowed. If empty, this ClusterQueue cannot borrow from any other ClusterQueue and vice versa.  A cohort is a name that links CQs together, but it doesn't reference any object.  Validation of a cohort name is equivalent to that of object names: subdomain in DNS (RFC 1123).",
						MarkdownDescription: "cohort that this ClusterQueue belongs to. CQs that belong to the same cohort can borrow unused resources from each other.  A CQ can be a member of a single borrowing cohort. A workload submitted to a queue referencing this CQ can borrow quota from any CQ in the cohort. Only quota for the [resource, flavor] pairs listed in the CQ can be borrowed. If empty, this ClusterQueue cannot borrow from any other ClusterQueue and vice versa.  A cohort is a name that links CQs together, but it doesn't reference any object.  Validation of a cohort name is equivalent to that of object names: subdomain in DNS (RFC 1123).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "namespaceSelector defines which namespaces are allowed to submit workloads to this clusterQueue. Beyond this basic support for policy, an policy agent like Gatekeeper should be used to enforce more advanced policies. Defaults to null which is a nothing selector (no namespaces eligible). If set to an empty selector '{}', then all namespaces are eligible.",
						MarkdownDescription: "namespaceSelector defines which namespaces are allowed to submit workloads to this clusterQueue. Beyond this basic support for policy, an policy agent like Gatekeeper should be used to enforce more advanced policies. Defaults to null which is a nothing selector (no namespaces eligible). If set to an empty selector '{}', then all namespaces are eligible.",
						Attributes: map[string]schema.Attribute{
							"match_expressions": schema.ListNestedAttribute{
								Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "key is the label key that the selector applies to.",
											MarkdownDescription: "key is the label key that the selector applies to.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"operator": schema.StringAttribute{
											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"values": schema.ListAttribute{
											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

							"match_labels": schema.MapAttribute{
								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"preemption": schema.SingleNestedAttribute{
						Description:         "preemption describes policies to preempt Workloads from this ClusterQueue or the ClusterQueue's cohort.  Preemption can happen in two scenarios:  - When a Workload fits within the nominal quota of the ClusterQueue, but the quota is currently borrowed by other ClusterQueues in the cohort. Preempting Workloads in other ClusterQueues allows this ClusterQueue to reclaim its nominal quota. - When a Workload doesn't fit within the nominal quota of the ClusterQueue and there are admitted Workloads in the ClusterQueue with lower priority.  The preemption algorithm tries to find a minimal set of Workloads to preempt to accomomdate the pending Workload, preempting Workloads with lower priority first.",
						MarkdownDescription: "preemption describes policies to preempt Workloads from this ClusterQueue or the ClusterQueue's cohort.  Preemption can happen in two scenarios:  - When a Workload fits within the nominal quota of the ClusterQueue, but the quota is currently borrowed by other ClusterQueues in the cohort. Preempting Workloads in other ClusterQueues allows this ClusterQueue to reclaim its nominal quota. - When a Workload doesn't fit within the nominal quota of the ClusterQueue and there are admitted Workloads in the ClusterQueue with lower priority.  The preemption algorithm tries to find a minimal set of Workloads to preempt to accomomdate the pending Workload, preempting Workloads with lower priority first.",
						Attributes: map[string]schema.Attribute{
							"reclaim_within_cohort": schema.StringAttribute{
								Description:         "reclaimWithinCohort determines whether a pending Workload can preempt Workloads from other ClusterQueues in the cohort that are using more than their nominal quota. The possible values are:  - 'Never' (default): do not preempt Workloads in the cohort. - 'LowerPriority': if the pending Workload fits within the nominal quota of its ClusterQueue, only preempt Workloads in the cohort that have lower priority than the pending Workload. - 'Any': if the pending Workload fits within the nominal quota of its ClusterQueue, preempt any Workload in the cohort, irrespective of priority.",
								MarkdownDescription: "reclaimWithinCohort determines whether a pending Workload can preempt Workloads from other ClusterQueues in the cohort that are using more than their nominal quota. The possible values are:  - 'Never' (default): do not preempt Workloads in the cohort. - 'LowerPriority': if the pending Workload fits within the nominal quota of its ClusterQueue, only preempt Workloads in the cohort that have lower priority than the pending Workload. - 'Any': if the pending Workload fits within the nominal quota of its ClusterQueue, preempt any Workload in the cohort, irrespective of priority.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"within_cluster_queue": schema.StringAttribute{
								Description:         "withinClusterQueue determines whether a pending Workload that doesn't fit within the nominal quota for its ClusterQueue, can preempt active Workloads in the ClusterQueue. The possible values are:  - 'Never' (default): do not preempt Workloads in the ClusterQueue. - 'LowerPriority': only preempt Workloads in the ClusterQueue that have lower priority than the pending Workload. - 'LowerOrNewerEqualPriority': only preempt Workloads in the ClusterQueue that either have a lower priority than the pending workload or equal priority and are newer than the pending workload.",
								MarkdownDescription: "withinClusterQueue determines whether a pending Workload that doesn't fit within the nominal quota for its ClusterQueue, can preempt active Workloads in the ClusterQueue. The possible values are:  - 'Never' (default): do not preempt Workloads in the ClusterQueue. - 'LowerPriority': only preempt Workloads in the ClusterQueue that have lower priority than the pending Workload. - 'LowerOrNewerEqualPriority': only preempt Workloads in the ClusterQueue that either have a lower priority than the pending workload or equal priority and are newer than the pending workload.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"queueing_strategy": schema.StringAttribute{
						Description:         "QueueingStrategy indicates the queueing strategy of the workloads across the queues in this ClusterQueue. This field is immutable. Current Supported Strategies:  - StrictFIFO: workloads are ordered strictly by creation time. Older workloads that can't be admitted will block admitting newer workloads even if they fit available quota. - BestEffortFIFO: workloads are ordered by creation time, however older workloads that can't be admitted will not block admitting newer workloads that fit existing quota.",
						MarkdownDescription: "QueueingStrategy indicates the queueing strategy of the workloads across the queues in this ClusterQueue. This field is immutable. Current Supported Strategies:  - StrictFIFO: workloads are ordered strictly by creation time. Older workloads that can't be admitted will block admitting newer workloads even if they fit available quota. - BestEffortFIFO: workloads are ordered by creation time, however older workloads that can't be admitted will not block admitting newer workloads that fit existing quota.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"resource_groups": schema.ListNestedAttribute{
						Description:         "resourceGroups describes groups of resources. Each resource group defines the list of resources and a list of flavors that provide quotas for these resources. Each resource and each flavor can only form part of one resource group. resourceGroups can be up to 16.",
						MarkdownDescription: "resourceGroups describes groups of resources. Each resource group defines the list of resources and a list of flavors that provide quotas for these resources. Each resource and each flavor can only form part of one resource group. resourceGroups can be up to 16.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"covered_resources": schema.ListAttribute{
									Description:         "coveredResources is the list of resources covered by the flavors in this group. Examples: cpu, memory, vendor.com/gpu. The list cannot be empty and it can contain up to 16 resources.",
									MarkdownDescription: "coveredResources is the list of resources covered by the flavors in this group. Examples: cpu, memory, vendor.com/gpu. The list cannot be empty and it can contain up to 16 resources.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"flavors": schema.ListNestedAttribute{
									Description:         "flavors is the list of flavors that provide the resources of this group. Typically, different flavors represent different hardware models (e.g., gpu models, cpu architectures) or pricing models (on-demand vs spot cpus). Each flavor MUST list all the resources listed for this group in the same order as the .resources field. The list cannot be empty and it can contain up to 16 flavors.",
									MarkdownDescription: "flavors is the list of flavors that provide the resources of this group. Typically, different flavors represent different hardware models (e.g., gpu models, cpu architectures) or pricing models (on-demand vs spot cpus). Each flavor MUST list all the resources listed for this group in the same order as the .resources field. The list cannot be empty and it can contain up to 16 flavors.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name of this flavor. The name should match the .metadata.name of a ResourceFlavor. If a matching ResourceFlavor does not exist, the ClusterQueue will have an Active condition set to False.",
												MarkdownDescription: "name of this flavor. The name should match the .metadata.name of a ResourceFlavor. If a matching ResourceFlavor does not exist, the ClusterQueue will have an Active condition set to False.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"resources": schema.ListNestedAttribute{
												Description:         "resources is the list of quotas for this flavor per resource. There could be up to 16 resources.",
												MarkdownDescription: "resources is the list of quotas for this flavor per resource. There could be up to 16 resources.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"borrowing_limit": schema.StringAttribute{
															Description:         "borrowingLimit is the maximum amount of quota for the [flavor, resource] combination that this ClusterQueue is allowed to borrow from the unused quota of other ClusterQueues in the same cohort. In total, at a given time, Workloads in a ClusterQueue can consume a quantity of quota equal to nominalQuota+borrowingLimit, assuming the other ClusterQueues in the cohort have enough unused quota. If null, it means that there is no borrowing limit. If not null, it must be non-negative. borrowingLimit must be null if spec.cohort is empty.",
															MarkdownDescription: "borrowingLimit is the maximum amount of quota for the [flavor, resource] combination that this ClusterQueue is allowed to borrow from the unused quota of other ClusterQueues in the same cohort. In total, at a given time, Workloads in a ClusterQueue can consume a quantity of quota equal to nominalQuota+borrowingLimit, assuming the other ClusterQueues in the cohort have enough unused quota. If null, it means that there is no borrowing limit. If not null, it must be non-negative. borrowingLimit must be null if spec.cohort is empty.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "name of this resource.",
															MarkdownDescription: "name of this resource.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"nominal_quota": schema.StringAttribute{
															Description:         "nominalQuota is the quantity of this resource that is available for Workloads admitted by this ClusterQueue at a point in time. The nominalQuota must be non-negative. nominalQuota should represent the resources in the cluster available for running jobs (after discounting resources consumed by system components and pods not managed by kueue). In an autoscaled cluster, nominalQuota should account for resources that can be provided by a component such as Kubernetes cluster-autoscaler.  If the ClusterQueue belongs to a cohort, the sum of the quotas for each (flavor, resource) combination defines the maximum quantity that can be allocated by a ClusterQueue in the cohort.",
															MarkdownDescription: "nominalQuota is the quantity of this resource that is available for Workloads admitted by this ClusterQueue at a point in time. The nominalQuota must be non-negative. nominalQuota should represent the resources in the cluster available for running jobs (after discounting resources consumed by system components and pods not managed by kueue). In an autoscaled cluster, nominalQuota should account for resources that can be provided by a component such as Kubernetes cluster-autoscaler.  If the ClusterQueue belongs to a cohort, the sum of the quotas for each (flavor, resource) combination defines the maximum quantity that can be allocated by a ClusterQueue in the cohort.",
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
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},
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
		},
	}
}

func (r *KueueXK8SIoClusterQueueV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *KueueXK8SIoClusterQueueV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kueue_x_k8s_io_cluster_queue_v1beta1")

	var data KueueXK8SIoClusterQueueV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kueue.x-k8s.io", Version: "v1beta1", Resource: "clusterqueues"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse KueueXK8SIoClusterQueueV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("kueue.x-k8s.io/v1beta1")
	data.Kind = pointer.String("ClusterQueue")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
