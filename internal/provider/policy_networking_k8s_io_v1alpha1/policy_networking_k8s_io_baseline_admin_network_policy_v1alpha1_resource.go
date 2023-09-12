/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package policy_networking_k8s_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
	"time"
)

var (
	_ resource.Resource                = &PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource{}
)

func NewPolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource() resource.Resource {
	return &PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource{}
}

type PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1ResourceData struct {
	ID                  types.String `tfsdk:"id" json:"-"`
	ForceConflicts      types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager        types.String `tfsdk:"field_manager" json:"-"`
	DeletionPropagation types.String `tfsdk:"deletion_propagation" json:"-"`
	WaitForUpsert       types.List   `tfsdk:"wait_for_upsert" json:"-"`
	WaitForDelete       types.Object `tfsdk:"wait_for_delete" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Egress *[]struct {
			Action *string `tfsdk:"action" json:"action,omitempty"`
			Name   *string `tfsdk:"name" json:"name,omitempty"`
			Ports  *[]struct {
				NamedPort  *string `tfsdk:"named_port" json:"namedPort,omitempty"`
				PortNumber *struct {
					Port     *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"port_number" json:"portNumber,omitempty"`
				PortRange *struct {
					End      *int64  `tfsdk:"end" json:"end,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Start    *int64  `tfsdk:"start" json:"start,omitempty"`
				} `tfsdk:"port_range" json:"portRange,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
			To *[]struct {
				Namespaces *struct {
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					NotSameLabels *[]string `tfsdk:"not_same_labels" json:"notSameLabels,omitempty"`
					SameLabels    *[]string `tfsdk:"same_labels" json:"sameLabels,omitempty"`
				} `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Pods *struct {
					Namespaces *struct {
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						NotSameLabels *[]string `tfsdk:"not_same_labels" json:"notSameLabels,omitempty"`
						SameLabels    *[]string `tfsdk:"same_labels" json:"sameLabels,omitempty"`
					} `tfsdk:"namespaces" json:"namespaces,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
				} `tfsdk:"pods" json:"pods,omitempty"`
			} `tfsdk:"to" json:"to,omitempty"`
		} `tfsdk:"egress" json:"egress,omitempty"`
		Ingress *[]struct {
			Action *string `tfsdk:"action" json:"action,omitempty"`
			From   *[]struct {
				Namespaces *struct {
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					NotSameLabels *[]string `tfsdk:"not_same_labels" json:"notSameLabels,omitempty"`
					SameLabels    *[]string `tfsdk:"same_labels" json:"sameLabels,omitempty"`
				} `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Pods *struct {
					Namespaces *struct {
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						NotSameLabels *[]string `tfsdk:"not_same_labels" json:"notSameLabels,omitempty"`
						SameLabels    *[]string `tfsdk:"same_labels" json:"sameLabels,omitempty"`
					} `tfsdk:"namespaces" json:"namespaces,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
				} `tfsdk:"pods" json:"pods,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Ports *[]struct {
				NamedPort  *string `tfsdk:"named_port" json:"namedPort,omitempty"`
				PortNumber *struct {
					Port     *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"port_number" json:"portNumber,omitempty"`
				PortRange *struct {
					End      *int64  `tfsdk:"end" json:"end,omitempty"`
					Protocol *string `tfsdk:"protocol" json:"protocol,omitempty"`
					Start    *int64  `tfsdk:"start" json:"start,omitempty"`
				} `tfsdk:"port_range" json:"portRange,omitempty"`
			} `tfsdk:"ports" json:"ports,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Subject *struct {
			Namespaces *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespaces" json:"namespaces,omitempty"`
			Pods *struct {
				NamespaceSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
				PodSelector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
			} `tfsdk:"pods" json:"pods,omitempty"`
		} `tfsdk:"subject" json:"subject,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1"
}

func (r *PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BaselineAdminNetworkPolicy is a cluster level resource that is part of the AdminNetworkPolicy API.",
		MarkdownDescription: "BaselineAdminNetworkPolicy is a cluster level resource that is part of the AdminNetworkPolicy API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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

			"field_manager": schema.StringAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},

			"deletion_propagation": schema.StringAttribute{
				Description:         "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				MarkdownDescription: "Decides if a deletion will propagate to the dependents of the object, and how the garbage collector will handle the propagation.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("Orphan", "Background", "Foreground"),
				},
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
						"timeout": schema.Int64Attribute{
							Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(30),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
						},
						"poll_interval": schema.Int64Attribute{
							Description:         "The number of seconds to wait before checking again.",
							MarkdownDescription: "The number of seconds to wait before checking again.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             int64default.StaticInt64(5),
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
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
					"timeout": schema.Int64Attribute{
						Description:         "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						MarkdownDescription: "The number of seconds to wait before giving up. Zero means check once and don't wait.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(30),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},
					"poll_interval": schema.Int64Attribute{
						Description:         "The number of seconds to wait before checking again.",
						MarkdownDescription: "The number of seconds to wait before checking again.",
						Required:            false,
						Optional:            true,
						Computed:            true,
						Default:             int64default.StaticInt64(5),
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
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
				Description:         "Specification of the desired behavior of BaselineAdminNetworkPolicy.",
				MarkdownDescription: "Specification of the desired behavior of BaselineAdminNetworkPolicy.",
				Attributes: map[string]schema.Attribute{
					"egress": schema.ListNestedAttribute{
						Description:         "Egress is the list of Egress rules to be applied to the selected pods if they are not matched by any AdminNetworkPolicy or NetworkPolicy rules. A total of 100 Egress rules will be allowed in each BANP instance. The relative precedence of egress rules within a single BANP object will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the egress rules would take the highest precedence. BANPs with no egress rules do not affect egress traffic.",
						MarkdownDescription: "Egress is the list of Egress rules to be applied to the selected pods if they are not matched by any AdminNetworkPolicy or NetworkPolicy rules. A total of 100 Egress rules will be allowed in each BANP instance. The relative precedence of egress rules within a single BANP object will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the egress rules would take the highest precedence. BANPs with no egress rules do not affect egress traffic.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic Deny: denies the selected traffic",
									MarkdownDescription: "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic Deny: denies the selected traffic",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied BaselineAdminNetworkPolicies.",
									MarkdownDescription: "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied BaselineAdminNetworkPolicies.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(100),
									},
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Ports allows for matching traffic based on port and protocols. This field is a list of destination ports for the outging egress traffic. If Ports is not set then the rule does not filter traffic via port.",
									MarkdownDescription: "Ports allows for matching traffic based on port and protocols. This field is a list of destination ports for the outging egress traffic. If Ports is not set then the rule does not filter traffic via port.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"named_port": schema.StringAttribute{
												Description:         "NamedPort selects a port on a pod(s) based on name.",
												MarkdownDescription: "NamedPort selects a port on a pod(s) based on name.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("port_number"), path.MatchRelative().AtParent().AtName("port_range")),
												},
											},

											"port_number": schema.SingleNestedAttribute{
												Description:         "Port selects a port on a pod(s) based on number.",
												MarkdownDescription: "Port selects a port on a pod(s) based on number.",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Number defines a network port value.",
														MarkdownDescription: "Number defines a network port value.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("named_port"), path.MatchRelative().AtParent().AtName("port_range")),
												},
											},

											"port_range": schema.SingleNestedAttribute{
												Description:         "PortRange selects a port range on a pod(s) based on provided start and end values.",
												MarkdownDescription: "PortRange selects a port range on a pod(s) based on provided start and end values.",
												Attributes: map[string]schema.Attribute{
													"end": schema.Int64Attribute{
														Description:         "End defines a network port that is the end of a port range, the End value must be greater than Start.",
														MarkdownDescription: "End defines a network port that is the end of a port range, the End value must be greater than Start.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"start": schema.Int64Attribute{
														Description:         "Start defines a network port that is the start of a port range, the Start value must be less than End.",
														MarkdownDescription: "Start defines a network port that is the start of a port range, the Start value must be less than End.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("named_port"), path.MatchRelative().AtParent().AtName("port_number")),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"to": schema.ListNestedAttribute{
									Description:         "To is the list of destinations whose traffic this rule applies to. If any AdminNetworkPolicyPeer matches the destination of outgoing traffic then the specified action is applied. This field must be defined and contain at least one item.",
									MarkdownDescription: "To is the list of destinations whose traffic this rule applies to. If any AdminNetworkPolicyPeer matches the destination of outgoing traffic then the specified action is applied. This field must be defined and contain at least one item.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"namespaces": schema.SingleNestedAttribute{
												Description:         "Namespaces defines a way to select a set of Namespaces.",
												MarkdownDescription: "Namespaces defines a way to select a set of Namespaces.",
												Attributes: map[string]schema.Attribute{
													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
														MarkdownDescription: "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
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
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("not_same_labels"), path.MatchRelative().AtParent().AtName("same_labels")),
														},
													},

													"not_same_labels": schema.ListAttribute{
														Description:         "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
														MarkdownDescription: "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.List{
															listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespace_selector"), path.MatchRelative().AtParent().AtName("same_labels")),
														},
													},

													"same_labels": schema.ListAttribute{
														Description:         "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
														MarkdownDescription: "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.List{
															listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespace_selector"), path.MatchRelative().AtParent().AtName("not_same_labels")),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("pods")),
												},
											},

											"pods": schema.SingleNestedAttribute{
												Description:         "Pods defines a way to select a set of pods in in a set of namespaces.",
												MarkdownDescription: "Pods defines a way to select a set of pods in in a set of namespaces.",
												Attributes: map[string]schema.Attribute{
													"namespaces": schema.SingleNestedAttribute{
														Description:         "Namespaces is used to select a set of Namespaces.",
														MarkdownDescription: "Namespaces is used to select a set of Namespaces.",
														Attributes: map[string]schema.Attribute{
															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
																MarkdownDescription: "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
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
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("not_same_labels"), path.MatchRelative().AtParent().AtName("same_labels")),
																},
															},

															"not_same_labels": schema.ListAttribute{
																Description:         "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
																MarkdownDescription: "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.List{
																	listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespace_selector"), path.MatchRelative().AtParent().AtName("same_labels")),
																},
															},

															"same_labels": schema.ListAttribute{
																Description:         "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
																MarkdownDescription: "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.List{
																	listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespace_selector"), path.MatchRelative().AtParent().AtName("not_same_labels")),
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "PodSelector is a labelSelector used to select Pods, This field is NOT optional, follows standard label selector semantics and if present but empty, it selects all Pods.",
														MarkdownDescription: "PodSelector is a labelSelector used to select Pods, This field is NOT optional, follows standard label selector semantics and if present but empty, it selects all Pods.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespaces")),
												},
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

					"ingress": schema.ListNestedAttribute{
						Description:         "Ingress is the list of Ingress rules to be applied to the selected pods if they are not matched by any AdminNetworkPolicy or NetworkPolicy rules. A total of 100 Ingress rules will be allowed in each BANP instance. The relative precedence of ingress rules within a single BANP object will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the ingress rules would take the highest precedence. BANPs with no ingress rules do not affect ingress traffic.",
						MarkdownDescription: "Ingress is the list of Ingress rules to be applied to the selected pods if they are not matched by any AdminNetworkPolicy or NetworkPolicy rules. A total of 100 Ingress rules will be allowed in each BANP instance. The relative precedence of ingress rules within a single BANP object will be determined by the order in which the rule is written. Thus, a rule that appears at the top of the ingress rules would take the highest precedence. BANPs with no ingress rules do not affect ingress traffic.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic Deny: denies the selected traffic",
									MarkdownDescription: "Action specifies the effect this rule will have on matching traffic. Currently the following actions are supported: Allow: allows the selected traffic Deny: denies the selected traffic",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"from": schema.ListNestedAttribute{
									Description:         "From is the list of sources whose traffic this rule applies to. If any AdminNetworkPolicyPeer matches the source of incoming traffic then the specified action is applied. This field must be defined and contain at least one item.",
									MarkdownDescription: "From is the list of sources whose traffic this rule applies to. If any AdminNetworkPolicyPeer matches the source of incoming traffic then the specified action is applied. This field must be defined and contain at least one item.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"namespaces": schema.SingleNestedAttribute{
												Description:         "Namespaces defines a way to select a set of Namespaces.",
												MarkdownDescription: "Namespaces defines a way to select a set of Namespaces.",
												Attributes: map[string]schema.Attribute{
													"namespace_selector": schema.SingleNestedAttribute{
														Description:         "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
														MarkdownDescription: "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
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
														Validators: []validator.Object{
															objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("not_same_labels"), path.MatchRelative().AtParent().AtName("same_labels")),
														},
													},

													"not_same_labels": schema.ListAttribute{
														Description:         "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
														MarkdownDescription: "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.List{
															listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespace_selector"), path.MatchRelative().AtParent().AtName("same_labels")),
														},
													},

													"same_labels": schema.ListAttribute{
														Description:         "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
														MarkdownDescription: "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.List{
															listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespace_selector"), path.MatchRelative().AtParent().AtName("not_same_labels")),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("pods")),
												},
											},

											"pods": schema.SingleNestedAttribute{
												Description:         "Pods defines a way to select a set of pods in in a set of namespaces.",
												MarkdownDescription: "Pods defines a way to select a set of pods in in a set of namespaces.",
												Attributes: map[string]schema.Attribute{
													"namespaces": schema.SingleNestedAttribute{
														Description:         "Namespaces is used to select a set of Namespaces.",
														MarkdownDescription: "Namespaces is used to select a set of Namespaces.",
														Attributes: map[string]schema.Attribute{
															"namespace_selector": schema.SingleNestedAttribute{
																Description:         "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
																MarkdownDescription: "NamespaceSelector is a labelSelector used to select Namespaces, This field follows standard label selector semantics; if present but empty, it selects all Namespaces.",
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
																Validators: []validator.Object{
																	objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("not_same_labels"), path.MatchRelative().AtParent().AtName("same_labels")),
																},
															},

															"not_same_labels": schema.ListAttribute{
																Description:         "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
																MarkdownDescription: "NotSameLabels is used to select a set of Namespaces that do not have certain values for a set of label(s). To be selected a Namespace must have all of the labels defined in NotSameLabels, AND at least one of them must have different values than the subject of this policy. If NotSameLabels is empty then nothing is selected.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.List{
																	listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespace_selector"), path.MatchRelative().AtParent().AtName("same_labels")),
																},
															},

															"same_labels": schema.ListAttribute{
																Description:         "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
																MarkdownDescription: "SameLabels is used to select a set of Namespaces that share the same values for a set of labels. To be selected a Namespace must have all of the labels defined in SameLabels, AND they must all have the same value as the subject of this policy. If Samelabels is Empty then nothing is selected.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.List{
																	listvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespace_selector"), path.MatchRelative().AtParent().AtName("not_same_labels")),
																},
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"pod_selector": schema.SingleNestedAttribute{
														Description:         "PodSelector is a labelSelector used to select Pods, This field is NOT optional, follows standard label selector semantics and if present but empty, it selects all Pods.",
														MarkdownDescription: "PodSelector is a labelSelector used to select Pods, This field is NOT optional, follows standard label selector semantics and if present but empty, it selects all Pods.",
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
														Required: true,
														Optional: false,
														Computed: false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespaces")),
												},
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied BaselineAdminNetworkPolicies.",
									MarkdownDescription: "Name is an identifier for this rule, that may be no more than 100 characters in length. This field should be used by the implementation to help improve observability, readability and error-reporting for any applied BaselineAdminNetworkPolicies.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(100),
									},
								},

								"ports": schema.ListNestedAttribute{
									Description:         "Ports allows for matching traffic based on port and protocols. This field is a list of ports which should be matched on the pods selected for this policy i.e the subject of the policy. So it matches on the destination port for the ingress traffic. If Ports is not set then the rule does not filter traffic via port.",
									MarkdownDescription: "Ports allows for matching traffic based on port and protocols. This field is a list of ports which should be matched on the pods selected for this policy i.e the subject of the policy. So it matches on the destination port for the ingress traffic. If Ports is not set then the rule does not filter traffic via port.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"named_port": schema.StringAttribute{
												Description:         "NamedPort selects a port on a pod(s) based on name.",
												MarkdownDescription: "NamedPort selects a port on a pod(s) based on name.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("port_number"), path.MatchRelative().AtParent().AtName("port_range")),
												},
											},

											"port_number": schema.SingleNestedAttribute{
												Description:         "Port selects a port on a pod(s) based on number.",
												MarkdownDescription: "Port selects a port on a pod(s) based on number.",
												Attributes: map[string]schema.Attribute{
													"port": schema.Int64Attribute{
														Description:         "Number defines a network port value.",
														MarkdownDescription: "Number defines a network port value.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("named_port"), path.MatchRelative().AtParent().AtName("port_range")),
												},
											},

											"port_range": schema.SingleNestedAttribute{
												Description:         "PortRange selects a port range on a pod(s) based on provided start and end values.",
												MarkdownDescription: "PortRange selects a port range on a pod(s) based on provided start and end values.",
												Attributes: map[string]schema.Attribute{
													"end": schema.Int64Attribute{
														Description:         "End defines a network port that is the end of a port range, the End value must be greater than Start.",
														MarkdownDescription: "End defines a network port that is the end of a port range, the End value must be greater than Start.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														MarkdownDescription: "Protocol is the network protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"start": schema.Int64Attribute{
														Description:         "Start defines a network port that is the start of a port range, the Start value must be less than End.",
														MarkdownDescription: "Start defines a network port that is the start of a port range, the Start value must be less than End.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(1),
															int64validator.AtMost(65535),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
												Validators: []validator.Object{
													objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("named_port"), path.MatchRelative().AtParent().AtName("port_number")),
												},
											},
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

					"subject": schema.SingleNestedAttribute{
						Description:         "Subject defines the pods to which this BaselineAdminNetworkPolicy applies.",
						MarkdownDescription: "Subject defines the pods to which this BaselineAdminNetworkPolicy applies.",
						Attributes: map[string]schema.Attribute{
							"namespaces": schema.SingleNestedAttribute{
								Description:         "Namespaces is used to select pods via namespace selectors.",
								MarkdownDescription: "Namespaces is used to select pods via namespace selectors.",
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
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("pods")),
								},
							},

							"pods": schema.SingleNestedAttribute{
								Description:         "Pods is used to select pods via namespace AND pod selectors.",
								MarkdownDescription: "Pods is used to select pods via namespace AND pod selectors.",
								Attributes: map[string]schema.Attribute{
									"namespace_selector": schema.SingleNestedAttribute{
										Description:         "NamespaceSelector follows standard label selector semantics; if empty, it selects all Namespaces.",
										MarkdownDescription: "NamespaceSelector follows standard label selector semantics; if empty, it selects all Namespaces.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},

									"pod_selector": schema.SingleNestedAttribute{
										Description:         "PodSelector is used to explicitly select pods within a namespace; if empty, it selects all Pods.",
										MarkdownDescription: "PodSelector is used to explicitly select pods within a namespace; if empty, it selects all Pods.",
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
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
								Validators: []validator.Object{
									objectvalidator.ExactlyOneOf(path.MatchRelative().AtParent().AtName("namespaces")),
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

func (r *PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1")

	var model PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("policy.networking.k8s.io/v1alpha1")
	model.Kind = pointer.String("BaselineAdminNetworkPolicy")

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
		Resource(k8sSchema.GroupVersionResource{Group: "policy.networking.k8s.io", Version: "v1alpha1", Resource: "baselineadminnetworkpolicies"}).
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

	var readResponse PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec
	if model.ForceConflicts.IsUnknown() {
		model.ForceConflicts = types.BoolNull()
	}
	if model.FieldManager.IsUnknown() {
		model.FieldManager = types.StringNull()
	}
	if model.DeletionPropagation.IsUnknown() {
		model.DeletionPropagation = types.StringNull()
	}
	if model.WaitForUpsert.IsUnknown() {
		model.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if model.WaitForDelete.IsUnknown() {
		model.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1")

	var data PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "policy.networking.k8s.io", Version: "v1alpha1", Resource: "baselineadminnetworkpolicies"}).
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

	var readResponse PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec
	if data.ForceConflicts.IsUnknown() {
		data.ForceConflicts = types.BoolNull()
	}
	if data.FieldManager.IsUnknown() {
		data.FieldManager = types.StringNull()
	}
	if data.DeletionPropagation.IsUnknown() {
		data.DeletionPropagation = types.StringNull()
	}
	if data.WaitForUpsert.IsUnknown() {
		data.WaitForUpsert = types.ListNull(types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"jsonpath":      types.StringType,
				"value":         types.StringType,
				"timeout":       types.Int64Type,
				"poll_interval": types.Int64Type,
			},
		})
	}
	if data.WaitForDelete.IsUnknown() {
		data.WaitForDelete = types.ObjectNull(map[string]attr.Type{
			"timeout":       types.Int64Type,
			"poll_interval": types.Int64Type,
		})
	}

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1")

	var model PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("policy.networking.k8s.io/v1alpha1")
	model.Kind = pointer.String("BaselineAdminNetworkPolicy")

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
		Resource(k8sSchema.GroupVersionResource{Group: "policy.networking.k8s.io", Version: "v1alpha1", Resource: "baselineadminnetworkpolicies"}).
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

	var readResponse PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_policy_networking_k8s_io_baseline_admin_network_policy_v1alpha1")

	var data PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	deleteOptions := meta.DeleteOptions{}
	if !data.DeletionPropagation.IsNull() && !data.DeletionPropagation.IsUnknown() {
		deleteOptions.PropagationPolicy = utilities.MapDeletionPropagation(data.DeletionPropagation.ValueString())
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "policy.networking.k8s.io", Version: "v1alpha1", Resource: "baselineadminnetworkpolicies"}).
		Delete(ctx, data.Metadata.Name, deleteOptions)
	if utilities.IsDeletionError(err) {
		response.Diagnostics.Append(utilities.DeleteError(err))
		return
	}

	if !data.WaitForDelete.IsNull() && !data.WaitForDelete.IsUnknown() {
		timeout := utilities.DetermineTimeout(data.WaitForDelete.Attributes())
		pollInterval := utilities.DeterminePollInterval(data.WaitForDelete.Attributes())

		startTime := time.Now()
		for {
			_, err := r.kubernetesClient.
				Resource(k8sSchema.GroupVersionResource{Group: "policy.networking.k8s.io", Version: "v1alpha1", Resource: "baselineadminnetworkpolicies"}).
				Get(ctx, data.Metadata.Name, meta.GetOptions{})
			if utilities.IsNotFound(err) || timeout.Milliseconds() == 0 {
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

func (r *PolicyNetworkingK8SIoBaselineAdminNetworkPolicyV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	if request.ID == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'name' Got: '%q'", request.ID),
		)
		return
	}
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	resource.ImportStatePassthroughID(ctx, path.Root("metadata").AtName("name"), request, response)
}
