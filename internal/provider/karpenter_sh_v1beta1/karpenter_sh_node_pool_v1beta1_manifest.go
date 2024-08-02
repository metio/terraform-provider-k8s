/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package karpenter_sh_v1beta1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &KarpenterShNodePoolV1Beta1Manifest{}
)

func NewKarpenterShNodePoolV1Beta1Manifest() datasource.DataSource {
	return &KarpenterShNodePoolV1Beta1Manifest{}
}

type KarpenterShNodePoolV1Beta1Manifest struct{}

type KarpenterShNodePoolV1Beta1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Disruption *struct {
			Budgets *[]struct {
				Duration *string   `tfsdk:"duration" json:"duration,omitempty"`
				Nodes    *string   `tfsdk:"nodes" json:"nodes,omitempty"`
				Reasons  *[]string `tfsdk:"reasons" json:"reasons,omitempty"`
				Schedule *string   `tfsdk:"schedule" json:"schedule,omitempty"`
			} `tfsdk:"budgets" json:"budgets,omitempty"`
			ConsolidateAfter    *string `tfsdk:"consolidate_after" json:"consolidateAfter,omitempty"`
			ConsolidationPolicy *string `tfsdk:"consolidation_policy" json:"consolidationPolicy,omitempty"`
			ExpireAfter         *string `tfsdk:"expire_after" json:"expireAfter,omitempty"`
		} `tfsdk:"disruption" json:"disruption,omitempty"`
		Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
		Template *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				Kubelet *struct {
					ClusterDNS                  *[]string          `tfsdk:"cluster_dns" json:"clusterDNS,omitempty"`
					CpuCFSQuota                 *bool              `tfsdk:"cpu_cfs_quota" json:"cpuCFSQuota,omitempty"`
					EvictionHard                *map[string]string `tfsdk:"eviction_hard" json:"evictionHard,omitempty"`
					EvictionMaxPodGracePeriod   *int64             `tfsdk:"eviction_max_pod_grace_period" json:"evictionMaxPodGracePeriod,omitempty"`
					EvictionSoft                *map[string]string `tfsdk:"eviction_soft" json:"evictionSoft,omitempty"`
					EvictionSoftGracePeriod     *map[string]string `tfsdk:"eviction_soft_grace_period" json:"evictionSoftGracePeriod,omitempty"`
					ImageGCHighThresholdPercent *int64             `tfsdk:"image_gc_high_threshold_percent" json:"imageGCHighThresholdPercent,omitempty"`
					ImageGCLowThresholdPercent  *int64             `tfsdk:"image_gc_low_threshold_percent" json:"imageGCLowThresholdPercent,omitempty"`
					KubeReserved                *map[string]string `tfsdk:"kube_reserved" json:"kubeReserved,omitempty"`
					MaxPods                     *int64             `tfsdk:"max_pods" json:"maxPods,omitempty"`
					PodsPerCore                 *int64             `tfsdk:"pods_per_core" json:"podsPerCore,omitempty"`
					SystemReserved              *map[string]string `tfsdk:"system_reserved" json:"systemReserved,omitempty"`
				} `tfsdk:"kubelet" json:"kubelet,omitempty"`
				NodeClassRef *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"node_class_ref" json:"nodeClassRef,omitempty"`
				Requirements *[]struct {
					Key       *string   `tfsdk:"key" json:"key,omitempty"`
					MinValues *int64    `tfsdk:"min_values" json:"minValues,omitempty"`
					Operator  *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values    *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"requirements" json:"requirements,omitempty"`
				Resources *struct {
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				StartupTaints *[]struct {
					Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"startup_taints" json:"startupTaints,omitempty"`
				Taints *[]struct {
					Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
					Key       *string `tfsdk:"key" json:"key,omitempty"`
					TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"taints" json:"taints,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KarpenterShNodePoolV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_karpenter_sh_node_pool_v1beta1_manifest"
}

func (r *KarpenterShNodePoolV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NodePool is the Schema for the NodePools API",
		MarkdownDescription: "NodePool is the Schema for the NodePools API",
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
				Description:         "NodePoolSpec is the top level nodepool specification. Nodepoolslaunch nodes in response to pods that are unschedulable. A single nodepoolis capable of managing a diverse set of nodes. Node properties are determinedfrom a combination of nodepool and pod scheduling constraints.",
				MarkdownDescription: "NodePoolSpec is the top level nodepool specification. Nodepoolslaunch nodes in response to pods that are unschedulable. A single nodepoolis capable of managing a diverse set of nodes. Node properties are determinedfrom a combination of nodepool and pod scheduling constraints.",
				Attributes: map[string]schema.Attribute{
					"disruption": schema.SingleNestedAttribute{
						Description:         "Disruption contains the parameters that relate to Karpenter's disruption logic",
						MarkdownDescription: "Disruption contains the parameters that relate to Karpenter's disruption logic",
						Attributes: map[string]schema.Attribute{
							"budgets": schema.ListNestedAttribute{
								Description:         "Budgets is a list of Budgets.If there are multiple active budgets, Karpenter usesthe most restrictive value. If left undefined,this will default to one budget with a value to 10%.",
								MarkdownDescription: "Budgets is a list of Budgets.If there are multiple active budgets, Karpenter usesthe most restrictive value. If left undefined,this will default to one budget with a value to 10%.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"duration": schema.StringAttribute{
											Description:         "Duration determines how long a Budget is active since each Schedule hit.Only minutes and hours are accepted, as cron does not work in seconds.If omitted, the budget is always active.This is required if Schedule is set.This regex has an optional 0s at the end since the duration.String() always addsa 0s at the end.",
											MarkdownDescription: "Duration determines how long a Budget is active since each Schedule hit.Only minutes and hours are accepted, as cron does not work in seconds.If omitted, the budget is always active.This is required if Schedule is set.This regex has an optional 0s at the end since the duration.String() always addsa 0s at the end.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^((([0-9]+(h|m))|([0-9]+h[0-9]+m))(0s)?)$`), ""),
											},
										},

										"nodes": schema.StringAttribute{
											Description:         "Nodes dictates the maximum number of NodeClaims owned by this NodePoolthat can be terminating at once. This is calculated by counting nodes thathave a deletion timestamp set, or are actively being deleted by Karpenter.This field is required when specifying a budget.This cannot be of type intstr.IntOrString since kubebuilder doesn't support patternchecking for int nodes for IntOrString nodes.Ref: https://github.com/kubernetes-sigs/controller-tools/blob/55efe4be40394a288216dab63156b0a64fb82929/pkg/crd/markers/validation.go#L379-L388",
											MarkdownDescription: "Nodes dictates the maximum number of NodeClaims owned by this NodePoolthat can be terminating at once. This is calculated by counting nodes thathave a deletion timestamp set, or are actively being deleted by Karpenter.This field is required when specifying a budget.This cannot be of type intstr.IntOrString since kubebuilder doesn't support patternchecking for int nodes for IntOrString nodes.Ref: https://github.com/kubernetes-sigs/controller-tools/blob/55efe4be40394a288216dab63156b0a64fb82929/pkg/crd/markers/validation.go#L379-L388",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^((100|[0-9]{1,2})%|[0-9]+)$`), ""),
											},
										},

										"reasons": schema.ListAttribute{
											Description:         "Reasons is a list of disruption methods that this budget applies to. If Reasons is not set, this budget applies to all methods.Otherwise, this will apply to each reason defined.allowed reasons are Underutilized, Empty, and Drifted.",
											MarkdownDescription: "Reasons is a list of disruption methods that this budget applies to. If Reasons is not set, this budget applies to all methods.Otherwise, this will apply to each reason defined.allowed reasons are Underutilized, Empty, and Drifted.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"schedule": schema.StringAttribute{
											Description:         "Schedule specifies when a budget begins being active, followingthe upstream cronjob syntax. If omitted, the budget is always active.Timezones are not supported.This field is required if Duration is set.",
											MarkdownDescription: "Schedule specifies when a budget begins being active, followingthe upstream cronjob syntax. If omitted, the budget is always active.Timezones are not supported.This field is required if Duration is set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(@(annually|yearly|monthly|weekly|daily|midnight|hourly))|((.+)\s(.+)\s(.+)\s(.+)\s(.+))$`), ""),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"consolidate_after": schema.StringAttribute{
								Description:         "ConsolidateAfter is the duration the controller will waitbefore attempting to terminate nodes that are underutilized.Refer to ConsolidationPolicy for how underutilization is considered.",
								MarkdownDescription: "ConsolidateAfter is the duration the controller will waitbefore attempting to terminate nodes that are underutilized.Refer to ConsolidationPolicy for how underutilization is considered.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+(s|m|h))+)|(Never)$`), ""),
								},
							},

							"consolidation_policy": schema.StringAttribute{
								Description:         "ConsolidationPolicy describes which nodes Karpenter can disrupt through its consolidationalgorithm. This policy defaults to 'WhenUnderutilized' if not specified",
								MarkdownDescription: "ConsolidationPolicy describes which nodes Karpenter can disrupt through its consolidationalgorithm. This policy defaults to 'WhenUnderutilized' if not specified",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("WhenEmpty", "WhenUnderutilized"),
								},
							},

							"expire_after": schema.StringAttribute{
								Description:         "ExpireAfter is the duration the controller will waitbefore terminating a node, measured from when the node is created. Thisis useful to implement features like eventually consistent node upgrade,memory leak protection, and disruption testing.",
								MarkdownDescription: "ExpireAfter is the duration the controller will waitbefore terminating a node, measured from when the node is created. Thisis useful to implement features like eventually consistent node upgrade,memory leak protection, and disruption testing.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+(s|m|h))+)|(Never)$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"limits": schema.MapAttribute{
						Description:         "Limits define a set of bounds for provisioning capacity.",
						MarkdownDescription: "Limits define a set of bounds for provisioning capacity.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template contains the template of possibilities for the provisioning logic to launch a NodeClaim with.NodeClaims launched from this NodePool will often be further constrained than the template specifies.",
						MarkdownDescription: "Template contains the template of possibilities for the provisioning logic to launch a NodeClaim with.NodeClaims launched from this NodePool will often be further constrained than the template specifies.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
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

							"spec": schema.SingleNestedAttribute{
								Description:         "NodeClaimSpec describes the desired state of the NodeClaim",
								MarkdownDescription: "NodeClaimSpec describes the desired state of the NodeClaim",
								Attributes: map[string]schema.Attribute{
									"kubelet": schema.SingleNestedAttribute{
										Description:         "Kubelet defines args to be used when configuring kubelet on provisioned nodes.They are a subset of the upstream types, recognizing not all options may be supported.Wherever possible, the types and names should reflect the upstream kubelet types.",
										MarkdownDescription: "Kubelet defines args to be used when configuring kubelet on provisioned nodes.They are a subset of the upstream types, recognizing not all options may be supported.Wherever possible, the types and names should reflect the upstream kubelet types.",
										Attributes: map[string]schema.Attribute{
											"cluster_dns": schema.ListAttribute{
												Description:         "clusterDNS is a list of IP addresses for the cluster DNS server.Note that not all providers may use all addresses.",
												MarkdownDescription: "clusterDNS is a list of IP addresses for the cluster DNS server.Note that not all providers may use all addresses.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cpu_cfs_quota": schema.BoolAttribute{
												Description:         "CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.",
												MarkdownDescription: "CPUCFSQuota enables CPU CFS quota enforcement for containers that specify CPU limits.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"eviction_hard": schema.MapAttribute{
												Description:         "EvictionHard is the map of signal names to quantities that define hard eviction thresholds",
												MarkdownDescription: "EvictionHard is the map of signal names to quantities that define hard eviction thresholds",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"eviction_max_pod_grace_period": schema.Int64Attribute{
												Description:         "EvictionMaxPodGracePeriod is the maximum allowed grace period (in seconds) to use when terminating pods inresponse to soft eviction thresholds being met.",
												MarkdownDescription: "EvictionMaxPodGracePeriod is the maximum allowed grace period (in seconds) to use when terminating pods inresponse to soft eviction thresholds being met.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"eviction_soft": schema.MapAttribute{
												Description:         "EvictionSoft is the map of signal names to quantities that define soft eviction thresholds",
												MarkdownDescription: "EvictionSoft is the map of signal names to quantities that define soft eviction thresholds",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"eviction_soft_grace_period": schema.MapAttribute{
												Description:         "EvictionSoftGracePeriod is the map of signal names to quantities that define grace periods for each eviction signal",
												MarkdownDescription: "EvictionSoftGracePeriod is the map of signal names to quantities that define grace periods for each eviction signal",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_gc_high_threshold_percent": schema.Int64Attribute{
												Description:         "ImageGCHighThresholdPercent is the percent of disk usage after which imagegarbage collection is always run. The percent is calculated by dividing thisfield value by 100, so this field must be between 0 and 100, inclusive.When specified, the value must be greater than ImageGCLowThresholdPercent.",
												MarkdownDescription: "ImageGCHighThresholdPercent is the percent of disk usage after which imagegarbage collection is always run. The percent is calculated by dividing thisfield value by 100, so this field must be between 0 and 100, inclusive.When specified, the value must be greater than ImageGCLowThresholdPercent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"image_gc_low_threshold_percent": schema.Int64Attribute{
												Description:         "ImageGCLowThresholdPercent is the percent of disk usage before which imagegarbage collection is never run. Lowest disk usage to garbage collect to.The percent is calculated by dividing this field value by 100,so the field value must be between 0 and 100, inclusive.When specified, the value must be less than imageGCHighThresholdPercent",
												MarkdownDescription: "ImageGCLowThresholdPercent is the percent of disk usage before which imagegarbage collection is never run. Lowest disk usage to garbage collect to.The percent is calculated by dividing this field value by 100,so the field value must be between 0 and 100, inclusive.When specified, the value must be less than imageGCHighThresholdPercent",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
													int64validator.AtMost(100),
												},
											},

											"kube_reserved": schema.MapAttribute{
												Description:         "KubeReserved contains resources reserved for Kubernetes system components.",
												MarkdownDescription: "KubeReserved contains resources reserved for Kubernetes system components.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_pods": schema.Int64Attribute{
												Description:         "MaxPods is an override for the maximum number of pods that can run ona worker node instance.",
												MarkdownDescription: "MaxPods is an override for the maximum number of pods that can run ona worker node instance.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"pods_per_core": schema.Int64Attribute{
												Description:         "PodsPerCore is an override for the number of pods that can run on a worker nodeinstance based on the number of cpu cores. This value cannot exceed MaxPods, so, ifMaxPods is a lower value, that value will be used.",
												MarkdownDescription: "PodsPerCore is an override for the number of pods that can run on a worker nodeinstance based on the number of cpu cores. This value cannot exceed MaxPods, so, ifMaxPods is a lower value, that value will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"system_reserved": schema.MapAttribute{
												Description:         "SystemReserved contains resources reserved for OS system daemons and kernel memory.",
												MarkdownDescription: "SystemReserved contains resources reserved for OS system daemons and kernel memory.",
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

									"node_class_ref": schema.SingleNestedAttribute{
										Description:         "NodeClassRef is a reference to an object that defines provider specific configuration",
										MarkdownDescription: "NodeClassRef is a reference to an object that defines provider specific configuration",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent",
												MarkdownDescription: "API version of the referent",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
												MarkdownDescription: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
												MarkdownDescription: "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"requirements": schema.ListNestedAttribute{
										Description:         "Requirements are layered with GetLabels and applied to every node.",
										MarkdownDescription: "Requirements are layered with GetLabels and applied to every node.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The label key that the selector applies to.",
													MarkdownDescription: "The label key that the selector applies to.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtMost(316),
														stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
													},
												},

												"min_values": schema.Int64Attribute{
													Description:         "This field is ALPHA and can be dropped or replaced at any timeMinValues is the minimum number of unique values required to define the flexibility of the specific requirement.",
													MarkdownDescription: "This field is ALPHA and can be dropped or replaced at any timeMinValues is the minimum number of unique values required to define the flexibility of the specific requirement.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(50),
													},
												},

												"operator": schema.StringAttribute{
													Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
													MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist", "Gt", "Lt"),
													},
												},

												"values": schema.ListAttribute{
													Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
													MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

									"resources": schema.SingleNestedAttribute{
										Description:         "Resources models the resource requirements for the NodeClaim to launch",
										MarkdownDescription: "Resources models the resource requirements for the NodeClaim to launch",
										Attributes: map[string]schema.Attribute{
											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum required resources for the NodeClaim to launch",
												MarkdownDescription: "Requests describes the minimum required resources for the NodeClaim to launch",
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

									"startup_taints": schema.ListNestedAttribute{
										Description:         "StartupTaints are taints that are applied to nodes upon startup which are expected to be removed automaticallywithin a short period of time, typically by a DaemonSet that tolerates the taint. These are commonly used bydaemonsets to allow initialization and enforce startup ordering.  StartupTaints are ignored for provisioningpurposes in that pods are not required to tolerate a StartupTaint in order to have nodes provisioned for them.",
										MarkdownDescription: "StartupTaints are taints that are applied to nodes upon startup which are expected to be removed automaticallywithin a short period of time, typically by a DaemonSet that tolerates the taint. These are commonly used bydaemonsets to allow initialization and enforce startup ordering.  StartupTaints are ignored for provisioningpurposes in that pods are not required to tolerate a StartupTaint in order to have nodes provisioned for them.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Required. The effect of the taint on podsthat do not tolerate the taint.Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Required. The effect of the taint on podsthat do not tolerate the taint.Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("NoSchedule", "PreferNoSchedule", "NoExecute"),
													},
												},

												"key": schema.StringAttribute{
													Description:         "Required. The taint key to be applied to a node.",
													MarkdownDescription: "Required. The taint key to be applied to a node.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
													},
												},

												"time_added": schema.StringAttribute{
													Description:         "TimeAdded represents the time at which the taint was added.It is only written for NoExecute taints.",
													MarkdownDescription: "TimeAdded represents the time at which the taint was added.It is only written for NoExecute taints.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														validators.DateTime64Validator(),
													},
												},

												"value": schema.StringAttribute{
													Description:         "The taint value corresponding to the taint key.",
													MarkdownDescription: "The taint value corresponding to the taint key.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"taints": schema.ListNestedAttribute{
										Description:         "Taints will be applied to the NodeClaim's node.",
										MarkdownDescription: "Taints will be applied to the NodeClaim's node.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Required. The effect of the taint on podsthat do not tolerate the taint.Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Required. The effect of the taint on podsthat do not tolerate the taint.Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("NoSchedule", "PreferNoSchedule", "NoExecute"),
													},
												},

												"key": schema.StringAttribute{
													Description:         "Required. The taint key to be applied to a node.",
													MarkdownDescription: "Required. The taint key to be applied to a node.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
														stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
													},
												},

												"time_added": schema.StringAttribute{
													Description:         "TimeAdded represents the time at which the taint was added.It is only written for NoExecute taints.",
													MarkdownDescription: "TimeAdded represents the time at which the taint was added.It is only written for NoExecute taints.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														validators.DateTime64Validator(),
													},
												},

												"value": schema.StringAttribute{
													Description:         "The taint value corresponding to the taint key.",
													MarkdownDescription: "The taint value corresponding to the taint key.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*(\/))?([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$`), ""),
													},
												},
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
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"weight": schema.Int64Attribute{
						Description:         "Weight is the priority given to the nodepool during scheduling. A highernumerical weight indicates that this nodepool will be orderedahead of other nodepools with lower weights. A nodepool with no weightwill be treated as if it is a nodepool with a weight of 0.",
						MarkdownDescription: "Weight is the priority given to the nodepool during scheduling. A highernumerical weight indicates that this nodepool will be orderedahead of other nodepools with lower weights. A nodepool with no weightwill be treated as if it is a nodepool with a weight of 0.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtMost(100),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *KarpenterShNodePoolV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_karpenter_sh_node_pool_v1beta1_manifest")

	var model KarpenterShNodePoolV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("karpenter.sh/v1beta1")
	model.Kind = pointer.String("NodePool")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
