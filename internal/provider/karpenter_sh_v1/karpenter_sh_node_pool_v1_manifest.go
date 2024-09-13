/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package karpenter_sh_v1

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
	_ datasource.DataSource = &KarpenterShNodePoolV1Manifest{}
)

func NewKarpenterShNodePoolV1Manifest() datasource.DataSource {
	return &KarpenterShNodePoolV1Manifest{}
}

type KarpenterShNodePoolV1Manifest struct{}

type KarpenterShNodePoolV1ManifestData struct {
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
		} `tfsdk:"disruption" json:"disruption,omitempty"`
		Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
		Template *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Spec *struct {
				ExpireAfter  *string `tfsdk:"expire_after" json:"expireAfter,omitempty"`
				NodeClassRef *struct {
					Group *string `tfsdk:"group" json:"group,omitempty"`
					Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"node_class_ref" json:"nodeClassRef,omitempty"`
				Requirements *[]struct {
					Key       *string   `tfsdk:"key" json:"key,omitempty"`
					MinValues *int64    `tfsdk:"min_values" json:"minValues,omitempty"`
					Operator  *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values    *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"requirements" json:"requirements,omitempty"`
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
				TerminationGracePeriod *string `tfsdk:"termination_grace_period" json:"terminationGracePeriod,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KarpenterShNodePoolV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_karpenter_sh_node_pool_v1_manifest"
}

func (r *KarpenterShNodePoolV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "NodePoolSpec is the top level nodepool specification. Nodepools launch nodes in response to pods that are unschedulable. A single nodepool is capable of managing a diverse set of nodes. Node properties are determined from a combination of nodepool and pod scheduling constraints.",
				MarkdownDescription: "NodePoolSpec is the top level nodepool specification. Nodepools launch nodes in response to pods that are unschedulable. A single nodepool is capable of managing a diverse set of nodes. Node properties are determined from a combination of nodepool and pod scheduling constraints.",
				Attributes: map[string]schema.Attribute{
					"disruption": schema.SingleNestedAttribute{
						Description:         "Disruption contains the parameters that relate to Karpenter's disruption logic",
						MarkdownDescription: "Disruption contains the parameters that relate to Karpenter's disruption logic",
						Attributes: map[string]schema.Attribute{
							"budgets": schema.ListNestedAttribute{
								Description:         "Budgets is a list of Budgets. If there are multiple active budgets, Karpenter uses the most restrictive value. If left undefined, this will default to one budget with a value to 10%.",
								MarkdownDescription: "Budgets is a list of Budgets. If there are multiple active budgets, Karpenter uses the most restrictive value. If left undefined, this will default to one budget with a value to 10%.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"duration": schema.StringAttribute{
											Description:         "Duration determines how long a Budget is active since each Schedule hit. Only minutes and hours are accepted, as cron does not work in seconds. If omitted, the budget is always active. This is required if Schedule is set. This regex has an optional 0s at the end since the duration.String() always adds a 0s at the end.",
											MarkdownDescription: "Duration determines how long a Budget is active since each Schedule hit. Only minutes and hours are accepted, as cron does not work in seconds. If omitted, the budget is always active. This is required if Schedule is set. This regex has an optional 0s at the end since the duration.String() always adds a 0s at the end.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^((([0-9]+(h|m))|([0-9]+h[0-9]+m))(0s)?)$`), ""),
											},
										},

										"nodes": schema.StringAttribute{
											Description:         "Nodes dictates the maximum number of NodeClaims owned by this NodePool that can be terminating at once. This is calculated by counting nodes that have a deletion timestamp set, or are actively being deleted by Karpenter. This field is required when specifying a budget. This cannot be of type intstr.IntOrString since kubebuilder doesn't support pattern checking for int nodes for IntOrString nodes. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/55efe4be40394a288216dab63156b0a64fb82929/pkg/crd/markers/validation.go#L379-L388",
											MarkdownDescription: "Nodes dictates the maximum number of NodeClaims owned by this NodePool that can be terminating at once. This is calculated by counting nodes that have a deletion timestamp set, or are actively being deleted by Karpenter. This field is required when specifying a budget. This cannot be of type intstr.IntOrString since kubebuilder doesn't support pattern checking for int nodes for IntOrString nodes. Ref: https://github.com/kubernetes-sigs/controller-tools/blob/55efe4be40394a288216dab63156b0a64fb82929/pkg/crd/markers/validation.go#L379-L388",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^((100|[0-9]{1,2})%|[0-9]+)$`), ""),
											},
										},

										"reasons": schema.ListAttribute{
											Description:         "Reasons is a list of disruption methods that this budget applies to. If Reasons is not set, this budget applies to all methods. Otherwise, this will apply to each reason defined. allowed reasons are Underutilized, Empty, and Drifted and additional CloudProvider-specific reasons.",
											MarkdownDescription: "Reasons is a list of disruption methods that this budget applies to. If Reasons is not set, this budget applies to all methods. Otherwise, this will apply to each reason defined. allowed reasons are Underutilized, Empty, and Drifted and additional CloudProvider-specific reasons.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"schedule": schema.StringAttribute{
											Description:         "Schedule specifies when a budget begins being active, following the upstream cronjob syntax. If omitted, the budget is always active. Timezones are not supported. This field is required if Duration is set.",
											MarkdownDescription: "Schedule specifies when a budget begins being active, following the upstream cronjob syntax. If omitted, the budget is always active. Timezones are not supported. This field is required if Duration is set.",
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
								Description:         "ConsolidateAfter is the duration the controller will wait before attempting to terminate nodes that are underutilized. Refer to ConsolidationPolicy for how underutilization is considered.",
								MarkdownDescription: "ConsolidateAfter is the duration the controller will wait before attempting to terminate nodes that are underutilized. Refer to ConsolidationPolicy for how underutilization is considered.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+(s|m|h))+)|(Never)$`), ""),
								},
							},

							"consolidation_policy": schema.StringAttribute{
								Description:         "ConsolidationPolicy describes which nodes Karpenter can disrupt through its consolidation algorithm. This policy defaults to 'WhenEmptyOrUnderutilized' if not specified",
								MarkdownDescription: "ConsolidationPolicy describes which nodes Karpenter can disrupt through its consolidation algorithm. This policy defaults to 'WhenEmptyOrUnderutilized' if not specified",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("WhenEmpty", "WhenEmptyOrUnderutilized"),
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
						Description:         "Template contains the template of possibilities for the provisioning logic to launch a NodeClaim with. NodeClaims launched from this NodePool will often be further constrained than the template specifies.",
						MarkdownDescription: "Template contains the template of possibilities for the provisioning logic to launch a NodeClaim with. NodeClaims launched from this NodePool will often be further constrained than the template specifies.",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels",
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
								Description:         "NodeClaimTemplateSpec describes the desired state of the NodeClaim in the Nodepool NodeClaimTemplateSpec is used in the NodePool's NodeClaimTemplate, with the resource requests omitted since users are not able to set resource requests in the NodePool.",
								MarkdownDescription: "NodeClaimTemplateSpec describes the desired state of the NodeClaim in the Nodepool NodeClaimTemplateSpec is used in the NodePool's NodeClaimTemplate, with the resource requests omitted since users are not able to set resource requests in the NodePool.",
								Attributes: map[string]schema.Attribute{
									"expire_after": schema.StringAttribute{
										Description:         "ExpireAfter is the duration the controller will wait before terminating a node, measured from when the node is created. This is useful to implement features like eventually consistent node upgrade, memory leak protection, and disruption testing.",
										MarkdownDescription: "ExpireAfter is the duration the controller will wait before terminating a node, measured from when the node is created. This is useful to implement features like eventually consistent node upgrade, memory leak protection, and disruption testing.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+(s|m|h))+)|(Never)$`), ""),
										},
									},

									"node_class_ref": schema.SingleNestedAttribute{
										Description:         "NodeClassRef is a reference to an object that defines provider specific configuration",
										MarkdownDescription: "NodeClassRef is a reference to an object that defines provider specific configuration",
										Attributes: map[string]schema.Attribute{
											"group": schema.StringAttribute{
												Description:         "API version of the referent",
												MarkdownDescription: "API version of the referent",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^[^/]*$`), ""),
												},
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
												MarkdownDescription: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
												Required:            true,
												Optional:            false,
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
													Description:         "This field is ALPHA and can be dropped or replaced at any time MinValues is the minimum number of unique values required to define the flexibility of the specific requirement.",
													MarkdownDescription: "This field is ALPHA and can be dropped or replaced at any time MinValues is the minimum number of unique values required to define the flexibility of the specific requirement.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
														int64validator.AtMost(50),
													},
												},

												"operator": schema.StringAttribute{
													Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
													MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist", "Gt", "Lt"),
													},
												},

												"values": schema.ListAttribute{
													Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
													MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

									"startup_taints": schema.ListNestedAttribute{
										Description:         "StartupTaints are taints that are applied to nodes upon startup which are expected to be removed automatically within a short period of time, typically by a DaemonSet that tolerates the taint. These are commonly used by daemonsets to allow initialization and enforce startup ordering. StartupTaints are ignored for provisioning purposes in that pods are not required to tolerate a StartupTaint in order to have nodes provisioned for them.",
										MarkdownDescription: "StartupTaints are taints that are applied to nodes upon startup which are expected to be removed automatically within a short period of time, typically by a DaemonSet that tolerates the taint. These are commonly used by daemonsets to allow initialization and enforce startup ordering. StartupTaints are ignored for provisioning purposes in that pods are not required to tolerate a StartupTaint in order to have nodes provisioned for them.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
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
													Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
													MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
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
													Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
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
													Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
													MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
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

									"termination_grace_period": schema.StringAttribute{
										Description:         "TerminationGracePeriod is the maximum duration the controller will wait before forcefully deleting the pods on a node, measured from when deletion is first initiated. Warning: this feature takes precedence over a Pod's terminationGracePeriodSeconds value, and bypasses any blocked PDBs or the karpenter.sh/do-not-disrupt annotation. This field is intended to be used by cluster administrators to enforce that nodes can be cycled within a given time period. When set, drifted nodes will begin draining even if there are pods blocking eviction. Draining will respect PDBs and the do-not-disrupt annotation until the TGP is reached. Karpenter will preemptively delete pods so their terminationGracePeriodSeconds align with the node's terminationGracePeriod. If a pod would be terminated without being granted its full terminationGracePeriodSeconds prior to the node timeout, that pod will be deleted at T = node timeout - pod terminationGracePeriodSeconds. The feature can also be used to allow maximum time limits for long-running jobs which can delay node termination with preStop hooks. If left undefined, the controller will wait indefinitely for pods to be drained.",
										MarkdownDescription: "TerminationGracePeriod is the maximum duration the controller will wait before forcefully deleting the pods on a node, measured from when deletion is first initiated. Warning: this feature takes precedence over a Pod's terminationGracePeriodSeconds value, and bypasses any blocked PDBs or the karpenter.sh/do-not-disrupt annotation. This field is intended to be used by cluster administrators to enforce that nodes can be cycled within a given time period. When set, drifted nodes will begin draining even if there are pods blocking eviction. Draining will respect PDBs and the do-not-disrupt annotation until the TGP is reached. Karpenter will preemptively delete pods so their terminationGracePeriodSeconds align with the node's terminationGracePeriod. If a pod would be terminated without being granted its full terminationGracePeriodSeconds prior to the node timeout, that pod will be deleted at T = node timeout - pod terminationGracePeriodSeconds. The feature can also be used to allow maximum time limits for long-running jobs which can delay node termination with preStop hooks. If left undefined, the controller will wait indefinitely for pods to be drained.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(s|m|h))+$`), ""),
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

					"weight": schema.Int64Attribute{
						Description:         "Weight is the priority given to the nodepool during scheduling. A higher numerical weight indicates that this nodepool will be ordered ahead of other nodepools with lower weights. A nodepool with no weight will be treated as if it is a nodepool with a weight of 0.",
						MarkdownDescription: "Weight is the priority given to the nodepool during scheduling. A higher numerical weight indicates that this nodepool will be ordered ahead of other nodepools with lower weights. A nodepool with no weight will be treated as if it is a nodepool with a weight of 0.",
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

func (r *KarpenterShNodePoolV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_karpenter_sh_node_pool_v1_manifest")

	var model KarpenterShNodePoolV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("karpenter.sh/v1")
	model.Kind = pointer.String("NodePool")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
