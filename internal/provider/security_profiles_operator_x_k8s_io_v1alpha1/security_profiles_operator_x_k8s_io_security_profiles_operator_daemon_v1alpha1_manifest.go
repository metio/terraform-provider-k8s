/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package security_profiles_operator_x_k8s_io_v1alpha1

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
	_ datasource.DataSource = &SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Manifest{}
)

func NewSecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Manifest() datasource.DataSource {
	return &SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Manifest{}
}

type SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Manifest struct{}

type SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1ManifestData struct {
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
		Affinity *struct {
			NodeAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					Preference *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchFields *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_fields" json:"matchFields,omitempty"`
					} `tfsdk:"preference" json:"preference,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *struct {
					NodeSelectorTerms *[]struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchFields *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_fields" json:"matchFields,omitempty"`
					} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
			PodAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
						MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
					MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
			PodAntiAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
						MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
					MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		AllowedSeccompActions      *[]string `tfsdk:"allowed_seccomp_actions" json:"allowedSeccompActions,omitempty"`
		AllowedSyscalls            *[]string `tfsdk:"allowed_syscalls" json:"allowedSyscalls,omitempty"`
		DaemonResourceRequirements *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"daemon_resource_requirements" json:"daemonResourceRequirements,omitempty"`
		DisableOciArtifactSignatureVerification *bool   `tfsdk:"disable_oci_artifact_signature_verification" json:"disableOciArtifactSignatureVerification,omitempty"`
		EnableAppArmor                          *bool   `tfsdk:"enable_app_armor" json:"enableAppArmor,omitempty"`
		EnableBpfRecorder                       *bool   `tfsdk:"enable_bpf_recorder" json:"enableBpfRecorder,omitempty"`
		EnableLogEnricher                       *bool   `tfsdk:"enable_log_enricher" json:"enableLogEnricher,omitempty"`
		EnableMemoryOptimization                *bool   `tfsdk:"enable_memory_optimization" json:"enableMemoryOptimization,omitempty"`
		EnableProfiling                         *bool   `tfsdk:"enable_profiling" json:"enableProfiling,omitempty"`
		EnableSelinux                           *bool   `tfsdk:"enable_selinux" json:"enableSelinux,omitempty"`
		HostProcVolumePath                      *string `tfsdk:"host_proc_volume_path" json:"hostProcVolumePath,omitempty"`
		ImagePullSecrets                        *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		SelinuxOptions    *struct {
			AllowedSystemProfiles *[]string `tfsdk:"allowed_system_profiles" json:"allowedSystemProfiles,omitempty"`
		} `tfsdk:"selinux_options" json:"selinuxOptions,omitempty"`
		SelinuxTypeTag      *string `tfsdk:"selinux_type_tag" json:"selinuxTypeTag,omitempty"`
		StaticWebhookConfig *bool   `tfsdk:"static_webhook_config" json:"staticWebhookConfig,omitempty"`
		Tolerations         *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		Verbosity      *int64 `tfsdk:"verbosity" json:"verbosity,omitempty"`
		WebhookOptions *[]struct {
			FailurePolicy     *string `tfsdk:"failure_policy" json:"failurePolicy,omitempty"`
			Name              *string `tfsdk:"name" json:"name,omitempty"`
			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
			ObjectSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"object_selector" json:"objectSelector,omitempty"`
		} `tfsdk:"webhook_options" json:"webhookOptions,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1_manifest"
}

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SecurityProfilesOperatorDaemon is the Schema to configure the spod deployment.",
		MarkdownDescription: "SecurityProfilesOperatorDaemon is the Schema to configure the spod deployment.",
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
				Description:         "SPODStatus defines the desired state of SPOD.",
				MarkdownDescription: "SPODStatus defines the desired state of SPOD.",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Affinity if specified, the SPOD's affinity.",
						MarkdownDescription: "Affinity if specified, the SPOD's affinity.",
						Attributes: map[string]schema.Attribute{
							"node_affinity": schema.SingleNestedAttribute{
								Description:         "Describes node affinity scheduling rules for the pod.",
								MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"preference": schema.SingleNestedAttribute{
													Description:         "A node selector term, associated with the corresponding weight.",
													MarkdownDescription: "A node selector term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's labels.",
															MarkdownDescription: "A list of node selector requirements by node's labels.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
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
															Required: false,
															Optional: true,
															Computed: false,
														},

														"match_fields": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
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
															Required: false,
															Optional: true,
															Computed: false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"weight": schema.Int64Attribute{
													Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
													MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
										Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
										Attributes: map[string]schema.Attribute{
											"node_selector_terms": schema.ListNestedAttribute{
												Description:         "Required. A list of node selector terms. The terms are ORed.",
												MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's labels.",
															MarkdownDescription: "A list of node selector requirements by node's labels.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
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
															Required: false,
															Optional: true,
															Computed: false,
														},

														"match_fields": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
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
															Required: false,
															Optional: true,
															Computed: false,
														},
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
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_affinity": schema.SingleNestedAttribute{
								Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Required. A pod affinity term, associated with the corresponding weight.",
													MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"match_label_keys": schema.ListAttribute{
															Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"weight": schema.Int64Attribute{
													Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"match_label_keys": schema.ListAttribute{
													Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mismatch_label_keys": schema.ListAttribute{
													Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"namespaces": schema.ListAttribute{
													Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
													Required:            true,
													Optional:            false,
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

							"pod_anti_affinity": schema.SingleNestedAttribute{
								Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Required. A pod affinity term, associated with the corresponding weight.",
													MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"match_label_keys": schema.ListAttribute{
															Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"weight": schema.Int64Attribute{
													Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"match_label_keys": schema.ListAttribute{
													Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mismatch_label_keys": schema.ListAttribute{
													Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"namespaces": schema.ListAttribute{
													Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
													Required:            true,
													Optional:            false,
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"allowed_seccomp_actions": schema.ListAttribute{
						Description:         "AllowedSeccompActions if specified, a list of allowed seccomp actions.",
						MarkdownDescription: "AllowedSeccompActions if specified, a list of allowed seccomp actions.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"allowed_syscalls": schema.ListAttribute{
						Description:         "AllowedSyscalls if specified, a list of system calls which are allowedin seccomp profiles.",
						MarkdownDescription: "AllowedSyscalls if specified, a list of system calls which are allowedin seccomp profiles.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"daemon_resource_requirements": schema.SingleNestedAttribute{
						Description:         "DaemonResourceRequirements if defined, overwrites the default resource requirementsof SPOD daemon.",
						MarkdownDescription: "DaemonResourceRequirements if defined, overwrites the default resource requirementsof SPOD daemon.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"disable_oci_artifact_signature_verification": schema.BoolAttribute{
						Description:         "DisableOCIArtifactSignatureVerification can be used to disable OCIartifact signature verification.",
						MarkdownDescription: "DisableOCIArtifactSignatureVerification can be used to disable OCIartifact signature verification.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_app_armor": schema.BoolAttribute{
						Description:         "tells the operator whether or not to enable AppArmor support for thisSPOD instance.",
						MarkdownDescription: "tells the operator whether or not to enable AppArmor support for thisSPOD instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_bpf_recorder": schema.BoolAttribute{
						Description:         "tells the operator whether or not to enable bpf recorder support for thisSPOD instance.",
						MarkdownDescription: "tells the operator whether or not to enable bpf recorder support for thisSPOD instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_log_enricher": schema.BoolAttribute{
						Description:         "tells the operator whether or not to enable log enrichment support for thisSPOD instance.",
						MarkdownDescription: "tells the operator whether or not to enable log enrichment support for thisSPOD instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_memory_optimization": schema.BoolAttribute{
						Description:         "EnableMemoryOptimization enables memory optimization in the controllerrunning inside of SPOD instance and watching for pods in the cluster.This will make the controller loading in the cache memory only the podslabelled explicitly for profile recording with 'spo.x-k8s.io/enable-recording=true'.",
						MarkdownDescription: "EnableMemoryOptimization enables memory optimization in the controllerrunning inside of SPOD instance and watching for pods in the cluster.This will make the controller loading in the cache memory only the podslabelled explicitly for profile recording with 'spo.x-k8s.io/enable-recording=true'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_profiling": schema.BoolAttribute{
						Description:         "EnableProfiling tells the operator whether or not to enable profilingsupport for this SPOD instance.",
						MarkdownDescription: "EnableProfiling tells the operator whether or not to enable profilingsupport for this SPOD instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_selinux": schema.BoolAttribute{
						Description:         "tells the operator whether or not to enable SELinux support for thisSPOD instance.",
						MarkdownDescription: "tells the operator whether or not to enable SELinux support for thisSPOD instance.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_proc_volume_path": schema.StringAttribute{
						Description:         "HostProcVolumePath is the path for specifying a custom host /procvolume, which is required for the log-enricher as well as bpf-recorderto retrieve the container ID for a process ID. This can be helpful fornested environments, for example when using 'kind'.",
						MarkdownDescription: "HostProcVolumePath is the path for specifying a custom host /procvolume, which is required for the log-enricher as well as bpf-recorderto retrieve the container ID for a process ID. This can be helpful fornested environments, for example when using 'kind'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "ImagePullSecrets if defined, list of references to secrets in the security-profiles-operator'snamespace to use for pulling the images from SPOD pod from a private registry.",
						MarkdownDescription: "ImagePullSecrets if defined, list of references to secrets in the security-profiles-operator'snamespace to use for pulling the images from SPOD pod from a private registry.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
									MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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

					"priority_class_name": schema.StringAttribute{
						Description:         "PriorityClassName if defined, indicates the spod pod priority class.",
						MarkdownDescription: "PriorityClassName if defined, indicates the spod pod priority class.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"selinux_options": schema.SingleNestedAttribute{
						Description:         "Defines options specific to the SELinuxfunctionality of the SecurityProfilesOperator",
						MarkdownDescription: "Defines options specific to the SELinuxfunctionality of the SecurityProfilesOperator",
						Attributes: map[string]schema.Attribute{
							"allowed_system_profiles": schema.ListAttribute{
								Description:         "Lists the profiles coming from the system itself that areallowed to be inherited by workloads. Use this with care,as this might provide a lot of permissions depending on thepolicy.",
								MarkdownDescription: "Lists the profiles coming from the system itself that areallowed to be inherited by workloads. Use this with care,as this might provide a lot of permissions depending on thepolicy.",
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

					"selinux_type_tag": schema.StringAttribute{
						Description:         "If specified, the SELinux type tag applied to the security context of SPOD.",
						MarkdownDescription: "If specified, the SELinux type tag applied to the security context of SPOD.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"static_webhook_config": schema.BoolAttribute{
						Description:         "StaticWebhookConfig indicates whether the webhook configuration and itsrelated resources are statically deployed. In this case, the operator willnot create or update the webhook configuration and its related resources.",
						MarkdownDescription: "StaticWebhookConfig indicates whether the webhook configuration and itsrelated resources are statically deployed. In this case, the operator willnot create or update the webhook configuration and its related resources.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "If specified, the SPOD's tolerations.",
						MarkdownDescription: "If specified, the SPOD's tolerations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

					"verbosity": schema.Int64Attribute{
						Description:         "Verbosity specifies the logging verbosity of the daemon.",
						MarkdownDescription: "Verbosity specifies the logging verbosity of the daemon.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"webhook_options": schema.ListNestedAttribute{
						Description:         "WebhookOpts set custom namespace selectors and failure mode forSPO's webhooks",
						MarkdownDescription: "WebhookOpts set custom namespace selectors and failure mode forSPO's webhooks",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"failure_policy": schema.StringAttribute{
									Description:         "FailurePolicy sets the webhook failure policy",
									MarkdownDescription: "FailurePolicy sets the webhook failure policy",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name specifies which webhook do we configure",
									MarkdownDescription: "Name specifies which webhook do we configure",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace_selector": schema.SingleNestedAttribute{
									Description:         "NamespaceSelector sets webhook's namespace selector",
									MarkdownDescription: "NamespaceSelector sets webhook's namespace selector",
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
														Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"object_selector": schema.SingleNestedAttribute{
									Description:         "ObjectSelector sets webhook's object selector",
									MarkdownDescription: "ObjectSelector sets webhook's object selector",
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
														Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1_manifest")

	var model SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("security-profiles-operator.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("SecurityProfilesOperatorDaemon")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
