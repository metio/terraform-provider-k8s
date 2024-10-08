/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kueue_x_k8s_io_v1beta1

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
	_ datasource.DataSource = &KueueXK8SIoClusterQueueV1Beta1Manifest{}
)

func NewKueueXK8SIoClusterQueueV1Beta1Manifest() datasource.DataSource {
	return &KueueXK8SIoClusterQueueV1Beta1Manifest{}
}

type KueueXK8SIoClusterQueueV1Beta1Manifest struct{}

type KueueXK8SIoClusterQueueV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AdmissionChecks         *[]string `tfsdk:"admission_checks" json:"admissionChecks,omitempty"`
		AdmissionChecksStrategy *struct {
			AdmissionChecks *[]struct {
				Name      *string   `tfsdk:"name" json:"name,omitempty"`
				OnFlavors *[]string `tfsdk:"on_flavors" json:"onFlavors,omitempty"`
			} `tfsdk:"admission_checks" json:"admissionChecks,omitempty"`
		} `tfsdk:"admission_checks_strategy" json:"admissionChecksStrategy,omitempty"`
		Cohort      *string `tfsdk:"cohort" json:"cohort,omitempty"`
		FairSharing *struct {
			Weight *string `tfsdk:"weight" json:"weight,omitempty"`
		} `tfsdk:"fair_sharing" json:"fairSharing,omitempty"`
		FlavorFungibility *struct {
			WhenCanBorrow  *string `tfsdk:"when_can_borrow" json:"whenCanBorrow,omitempty"`
			WhenCanPreempt *string `tfsdk:"when_can_preempt" json:"whenCanPreempt,omitempty"`
		} `tfsdk:"flavor_fungibility" json:"flavorFungibility,omitempty"`
		NamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key      *string   `tfsdk:"key" json:"key,omitempty"`
				Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
				Values   *[]string `tfsdk:"values" json:"values,omitempty"`
			} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
		Preemption *struct {
			BorrowWithinCohort *struct {
				MaxPriorityThreshold *int64  `tfsdk:"max_priority_threshold" json:"maxPriorityThreshold,omitempty"`
				Policy               *string `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"borrow_within_cohort" json:"borrowWithinCohort,omitempty"`
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
					LendingLimit   *string `tfsdk:"lending_limit" json:"lendingLimit,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					NominalQuota   *string `tfsdk:"nominal_quota" json:"nominalQuota,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"flavors" json:"flavors,omitempty"`
		} `tfsdk:"resource_groups" json:"resourceGroups,omitempty"`
		StopPolicy *string `tfsdk:"stop_policy" json:"stopPolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KueueXK8SIoClusterQueueV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kueue_x_k8s_io_cluster_queue_v1beta1_manifest"
}

func (r *KueueXK8SIoClusterQueueV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterQueue is the Schema for the clusterQueue API.",
		MarkdownDescription: "ClusterQueue is the Schema for the clusterQueue API.",
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
				Description:         "ClusterQueueSpec defines the desired state of ClusterQueue",
				MarkdownDescription: "ClusterQueueSpec defines the desired state of ClusterQueue",
				Attributes: map[string]schema.Attribute{
					"admission_checks": schema.ListAttribute{
						Description:         "admissionChecks lists the AdmissionChecks required by this ClusterQueue. Cannot be used along with AdmissionCheckStrategy.",
						MarkdownDescription: "admissionChecks lists the AdmissionChecks required by this ClusterQueue. Cannot be used along with AdmissionCheckStrategy.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"admission_checks_strategy": schema.SingleNestedAttribute{
						Description:         "admissionCheckStrategy defines a list of strategies to determine which ResourceFlavors require AdmissionChecks. This property cannot be used in conjunction with the 'admissionChecks' property.",
						MarkdownDescription: "admissionCheckStrategy defines a list of strategies to determine which ResourceFlavors require AdmissionChecks. This property cannot be used in conjunction with the 'admissionChecks' property.",
						Attributes: map[string]schema.Attribute{
							"admission_checks": schema.ListNestedAttribute{
								Description:         "admissionChecks is a list of strategies for AdmissionChecks",
								MarkdownDescription: "admissionChecks is a list of strategies for AdmissionChecks",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is an AdmissionCheck's name.",
											MarkdownDescription: "name is an AdmissionCheck's name.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"on_flavors": schema.ListAttribute{
											Description:         "onFlavors is a list of ResourceFlavors' names that this AdmissionCheck should run for. If empty, the AdmissionCheck will run for all workloads submitted to the ClusterQueue.",
											MarkdownDescription: "onFlavors is a list of ResourceFlavors' names that this AdmissionCheck should run for. If empty, the AdmissionCheck will run for all workloads submitted to the ClusterQueue.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cohort": schema.StringAttribute{
						Description:         "cohort that this ClusterQueue belongs to. CQs that belong to the same cohort can borrow unused resources from each other. A CQ can be a member of a single borrowing cohort. A workload submitted to a queue referencing this CQ can borrow quota from any CQ in the cohort. Only quota for the [resource, flavor] pairs listed in the CQ can be borrowed. If empty, this ClusterQueue cannot borrow from any other ClusterQueue and vice versa. A cohort is a name that links CQs together, but it doesn't reference any object. Validation of a cohort name is equivalent to that of object names: subdomain in DNS (RFC 1123).",
						MarkdownDescription: "cohort that this ClusterQueue belongs to. CQs that belong to the same cohort can borrow unused resources from each other. A CQ can be a member of a single borrowing cohort. A workload submitted to a queue referencing this CQ can borrow quota from any CQ in the cohort. Only quota for the [resource, flavor] pairs listed in the CQ can be borrowed. If empty, this ClusterQueue cannot borrow from any other ClusterQueue and vice versa. A cohort is a name that links CQs together, but it doesn't reference any object. Validation of a cohort name is equivalent to that of object names: subdomain in DNS (RFC 1123).",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(253),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
						},
					},

					"fair_sharing": schema.SingleNestedAttribute{
						Description:         "fairSharing defines the properties of the ClusterQueue when participating in fair sharing. The values are only relevant if fair sharing is enabled in the Kueue configuration.",
						MarkdownDescription: "fairSharing defines the properties of the ClusterQueue when participating in fair sharing. The values are only relevant if fair sharing is enabled in the Kueue configuration.",
						Attributes: map[string]schema.Attribute{
							"weight": schema.StringAttribute{
								Description:         "weight gives a comparative advantage to this ClusterQueue when competing for unused resources in the cohort against other ClusterQueues. The share of a ClusterQueue is based on the dominant resource usage above nominal quotas for each resource, divided by the weight. Admission prioritizes scheduling workloads from ClusterQueues with the lowest share and preempting workloads from the ClusterQueues with the highest share. A zero weight implies infinite share value, meaning that this ClusterQueue will always be at disadvantage against other ClusterQueues.",
								MarkdownDescription: "weight gives a comparative advantage to this ClusterQueue when competing for unused resources in the cohort against other ClusterQueues. The share of a ClusterQueue is based on the dominant resource usage above nominal quotas for each resource, divided by the weight. Admission prioritizes scheduling workloads from ClusterQueues with the lowest share and preempting workloads from the ClusterQueues with the highest share. A zero weight implies infinite share value, meaning that this ClusterQueue will always be at disadvantage against other ClusterQueues.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"flavor_fungibility": schema.SingleNestedAttribute{
						Description:         "flavorFungibility defines whether a workload should try the next flavor before borrowing or preempting in the flavor being evaluated.",
						MarkdownDescription: "flavorFungibility defines whether a workload should try the next flavor before borrowing or preempting in the flavor being evaluated.",
						Attributes: map[string]schema.Attribute{
							"when_can_borrow": schema.StringAttribute{
								Description:         "whenCanBorrow determines whether a workload should try the next flavor before borrowing in current flavor. The possible values are: - 'Borrow' (default): allocate in current flavor if borrowing is possible. - 'TryNextFlavor': try next flavor even if the current flavor has enough resources to borrow.",
								MarkdownDescription: "whenCanBorrow determines whether a workload should try the next flavor before borrowing in current flavor. The possible values are: - 'Borrow' (default): allocate in current flavor if borrowing is possible. - 'TryNextFlavor': try next flavor even if the current flavor has enough resources to borrow.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Borrow", "TryNextFlavor"),
								},
							},

							"when_can_preempt": schema.StringAttribute{
								Description:         "whenCanPreempt determines whether a workload should try the next flavor before borrowing in current flavor. The possible values are: - 'Preempt': allocate in current flavor if it's possible to preempt some workloads. - 'TryNextFlavor' (default): try next flavor even if there are enough candidates for preemption in the current flavor.",
								MarkdownDescription: "whenCanPreempt determines whether a workload should try the next flavor before borrowing in current flavor. The possible values are: - 'Preempt': allocate in current flavor if it's possible to preempt some workloads. - 'TryNextFlavor' (default): try next flavor even if there are enough candidates for preemption in the current flavor.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Preempt", "TryNextFlavor"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"namespace_selector": schema.SingleNestedAttribute{
						Description:         "namespaceSelector defines which namespaces are allowed to submit workloads to this clusterQueue. Beyond this basic support for policy, a policy agent like Gatekeeper should be used to enforce more advanced policies. Defaults to null which is a nothing selector (no namespaces eligible). If set to an empty selector '{}', then all namespaces are eligible.",
						MarkdownDescription: "namespaceSelector defines which namespaces are allowed to submit workloads to this clusterQueue. Beyond this basic support for policy, a policy agent like Gatekeeper should be used to enforce more advanced policies. Defaults to null which is a nothing selector (no namespaces eligible). If set to an empty selector '{}', then all namespaces are eligible.",
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

					"preemption": schema.SingleNestedAttribute{
						Description:         "preemption describes policies to preempt Workloads from this ClusterQueue or the ClusterQueue's cohort. Preemption can happen in two scenarios: - When a Workload fits within the nominal quota of the ClusterQueue, but the quota is currently borrowed by other ClusterQueues in the cohort. Preempting Workloads in other ClusterQueues allows this ClusterQueue to reclaim its nominal quota. - When a Workload doesn't fit within the nominal quota of the ClusterQueue and there are admitted Workloads in the ClusterQueue with lower priority. The preemption algorithm tries to find a minimal set of Workloads to preempt to accomomdate the pending Workload, preempting Workloads with lower priority first.",
						MarkdownDescription: "preemption describes policies to preempt Workloads from this ClusterQueue or the ClusterQueue's cohort. Preemption can happen in two scenarios: - When a Workload fits within the nominal quota of the ClusterQueue, but the quota is currently borrowed by other ClusterQueues in the cohort. Preempting Workloads in other ClusterQueues allows this ClusterQueue to reclaim its nominal quota. - When a Workload doesn't fit within the nominal quota of the ClusterQueue and there are admitted Workloads in the ClusterQueue with lower priority. The preemption algorithm tries to find a minimal set of Workloads to preempt to accomomdate the pending Workload, preempting Workloads with lower priority first.",
						Attributes: map[string]schema.Attribute{
							"borrow_within_cohort": schema.SingleNestedAttribute{
								Description:         "borrowWithinCohort provides configuration to allow preemption within cohort while borrowing.",
								MarkdownDescription: "borrowWithinCohort provides configuration to allow preemption within cohort while borrowing.",
								Attributes: map[string]schema.Attribute{
									"max_priority_threshold": schema.Int64Attribute{
										Description:         "maxPriorityThreshold allows to restrict the set of workloads which might be preempted by a borrowing workload, to only workloads with priority less than or equal to the specified threshold priority. When the threshold is not specified, then any workload satisfying the policy can be preempted by the borrowing workload.",
										MarkdownDescription: "maxPriorityThreshold allows to restrict the set of workloads which might be preempted by a borrowing workload, to only workloads with priority less than or equal to the specified threshold priority. When the threshold is not specified, then any workload satisfying the policy can be preempted by the borrowing workload.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"policy": schema.StringAttribute{
										Description:         "policy determines the policy for preemption to reclaim quota within cohort while borrowing. Possible values are: - 'Never' (default): do not allow for preemption, in other ClusterQueues within the cohort, for a borrowing workload. - 'LowerPriority': allow preemption, in other ClusterQueues within the cohort, for a borrowing workload, but only if the preempted workloads are of lower priority.",
										MarkdownDescription: "policy determines the policy for preemption to reclaim quota within cohort while borrowing. Possible values are: - 'Never' (default): do not allow for preemption, in other ClusterQueues within the cohort, for a borrowing workload. - 'LowerPriority': allow preemption, in other ClusterQueues within the cohort, for a borrowing workload, but only if the preempted workloads are of lower priority.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Never", "LowerPriority"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"reclaim_within_cohort": schema.StringAttribute{
								Description:         "reclaimWithinCohort determines whether a pending Workload can preempt Workloads from other ClusterQueues in the cohort that are using more than their nominal quota. The possible values are: - 'Never' (default): do not preempt Workloads in the cohort. - 'LowerPriority': if the pending Workload fits within the nominal quota of its ClusterQueue, only preempt Workloads in the cohort that have lower priority than the pending Workload. - 'Any': if the pending Workload fits within the nominal quota of its ClusterQueue, preempt any Workload in the cohort, irrespective of priority.",
								MarkdownDescription: "reclaimWithinCohort determines whether a pending Workload can preempt Workloads from other ClusterQueues in the cohort that are using more than their nominal quota. The possible values are: - 'Never' (default): do not preempt Workloads in the cohort. - 'LowerPriority': if the pending Workload fits within the nominal quota of its ClusterQueue, only preempt Workloads in the cohort that have lower priority than the pending Workload. - 'Any': if the pending Workload fits within the nominal quota of its ClusterQueue, preempt any Workload in the cohort, irrespective of priority.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Never", "LowerPriority", "Any"),
								},
							},

							"within_cluster_queue": schema.StringAttribute{
								Description:         "withinClusterQueue determines whether a pending Workload that doesn't fit within the nominal quota for its ClusterQueue, can preempt active Workloads in the ClusterQueue. The possible values are: - 'Never' (default): do not preempt Workloads in the ClusterQueue. - 'LowerPriority': only preempt Workloads in the ClusterQueue that have lower priority than the pending Workload. - 'LowerOrNewerEqualPriority': only preempt Workloads in the ClusterQueue that either have a lower priority than the pending workload or equal priority and are newer than the pending workload.",
								MarkdownDescription: "withinClusterQueue determines whether a pending Workload that doesn't fit within the nominal quota for its ClusterQueue, can preempt active Workloads in the ClusterQueue. The possible values are: - 'Never' (default): do not preempt Workloads in the ClusterQueue. - 'LowerPriority': only preempt Workloads in the ClusterQueue that have lower priority than the pending Workload. - 'LowerOrNewerEqualPriority': only preempt Workloads in the ClusterQueue that either have a lower priority than the pending workload or equal priority and are newer than the pending workload.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Never", "LowerPriority", "LowerOrNewerEqualPriority"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"queueing_strategy": schema.StringAttribute{
						Description:         "QueueingStrategy indicates the queueing strategy of the workloads across the queues in this ClusterQueue. Current Supported Strategies: - StrictFIFO: workloads are ordered strictly by creation time. Older workloads that can't be admitted will block admitting newer workloads even if they fit available quota. - BestEffortFIFO: workloads are ordered by creation time, however older workloads that can't be admitted will not block admitting newer workloads that fit existing quota.",
						MarkdownDescription: "QueueingStrategy indicates the queueing strategy of the workloads across the queues in this ClusterQueue. Current Supported Strategies: - StrictFIFO: workloads are ordered strictly by creation time. Older workloads that can't be admitted will block admitting newer workloads even if they fit available quota. - BestEffortFIFO: workloads are ordered by creation time, however older workloads that can't be admitted will not block admitting newer workloads that fit existing quota.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("StrictFIFO", "BestEffortFIFO"),
						},
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
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"flavors": schema.ListNestedAttribute{
									Description:         "flavors is the list of flavors that provide the resources of this group. Typically, different flavors represent different hardware models (e.g., gpu models, cpu architectures) or pricing models (on-demand vs spot cpus). Each flavor MUST list all the resources listed for this group in the same order as the .resources field. The list cannot be empty and it can contain up to 16 flavors.",
									MarkdownDescription: "flavors is the list of flavors that provide the resources of this group. Typically, different flavors represent different hardware models (e.g., gpu models, cpu architectures) or pricing models (on-demand vs spot cpus). Each flavor MUST list all the resources listed for this group in the same order as the .resources field. The list cannot be empty and it can contain up to 16 flavors.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name of this flavor. The name should match the .metadata.name of a ResourceFlavor. If a matching ResourceFlavor does not exist, the ClusterQueue will have an Active condition set to False.",
												MarkdownDescription: "name of this flavor. The name should match the .metadata.name of a ResourceFlavor. If a matching ResourceFlavor does not exist, the ClusterQueue will have an Active condition set to False.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(253),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
												},
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
															Optional:            true,
															Computed:            false,
														},

														"lending_limit": schema.StringAttribute{
															Description:         "lendingLimit is the maximum amount of unused quota for the [flavor, resource] combination that this ClusterQueue can lend to other ClusterQueues in the same cohort. In total, at a given time, ClusterQueue reserves for its exclusive use a quantity of quota equals to nominalQuota - lendingLimit. If null, it means that there is no lending limit, meaning that all the nominalQuota can be borrowed by other clusterQueues in the cohort. If not null, it must be non-negative. lendingLimit must be null if spec.cohort is empty. This field is in beta stage and is enabled by default.",
															MarkdownDescription: "lendingLimit is the maximum amount of unused quota for the [flavor, resource] combination that this ClusterQueue can lend to other ClusterQueues in the same cohort. In total, at a given time, ClusterQueue reserves for its exclusive use a quantity of quota equals to nominalQuota - lendingLimit. If null, it means that there is no lending limit, meaning that all the nominalQuota can be borrowed by other clusterQueues in the cohort. If not null, it must be non-negative. lendingLimit must be null if spec.cohort is empty. This field is in beta stage and is enabled by default.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "name of this resource.",
															MarkdownDescription: "name of this resource.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"nominal_quota": schema.StringAttribute{
															Description:         "nominalQuota is the quantity of this resource that is available for Workloads admitted by this ClusterQueue at a point in time. The nominalQuota must be non-negative. nominalQuota should represent the resources in the cluster available for running jobs (after discounting resources consumed by system components and pods not managed by kueue). In an autoscaled cluster, nominalQuota should account for resources that can be provided by a component such as Kubernetes cluster-autoscaler. If the ClusterQueue belongs to a cohort, the sum of the quotas for each (flavor, resource) combination defines the maximum quantity that can be allocated by a ClusterQueue in the cohort.",
															MarkdownDescription: "nominalQuota is the quantity of this resource that is available for Workloads admitted by this ClusterQueue at a point in time. The nominalQuota must be non-negative. nominalQuota should represent the resources in the cluster available for running jobs (after discounting resources consumed by system components and pods not managed by kueue). In an autoscaled cluster, nominalQuota should account for resources that can be provided by a component such as Kubernetes cluster-autoscaler. If the ClusterQueue belongs to a cohort, the sum of the quotas for each (flavor, resource) combination defines the maximum quantity that can be allocated by a ClusterQueue in the cohort.",
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
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"stop_policy": schema.StringAttribute{
						Description:         "stopPolicy - if set to a value different from None, the ClusterQueue is considered Inactive, no new reservation being made. Depending on its value, its associated workloads will: - None - Workloads are admitted - HoldAndDrain - Admitted workloads are evicted and Reserving workloads will cancel the reservation. - Hold - Admitted workloads will run to completion and Reserving workloads will cancel the reservation.",
						MarkdownDescription: "stopPolicy - if set to a value different from None, the ClusterQueue is considered Inactive, no new reservation being made. Depending on its value, its associated workloads will: - None - Workloads are admitted - HoldAndDrain - Admitted workloads are evicted and Reserving workloads will cancel the reservation. - Hold - Admitted workloads will run to completion and Reserving workloads will cancel the reservation.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("None", "Hold", "HoldAndDrain"),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *KueueXK8SIoClusterQueueV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kueue_x_k8s_io_cluster_queue_v1beta1_manifest")

	var model KueueXK8SIoClusterQueueV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kueue.x-k8s.io/v1beta1")
	model.Kind = pointer.String("ClusterQueue")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
