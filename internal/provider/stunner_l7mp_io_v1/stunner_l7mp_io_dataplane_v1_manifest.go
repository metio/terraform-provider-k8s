/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package stunner_l7mp_io_v1

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
	_ datasource.DataSource = &StunnerL7MpIoDataplaneV1Manifest{}
)

func NewStunnerL7MpIoDataplaneV1Manifest() datasource.DataSource {
	return &StunnerL7MpIoDataplaneV1Manifest{}
}

type StunnerL7MpIoDataplaneV1Manifest struct{}

type StunnerL7MpIoDataplaneV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
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
		Args                     *[]string `tfsdk:"args" json:"args,omitempty"`
		Command                  *[]string `tfsdk:"command" json:"command,omitempty"`
		ContainerSecurityContext *struct {
			AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
			Capabilities             *struct {
				Add  *[]string `tfsdk:"add" json:"add,omitempty"`
				Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
			} `tfsdk:"capabilities" json:"capabilities,omitempty"`
			Privileged             *bool   `tfsdk:"privileged" json:"privileged,omitempty"`
			ProcMount              *string `tfsdk:"proc_mount" json:"procMount,omitempty"`
			ReadOnlyRootFilesystem *bool   `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
			RunAsGroup             *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
			RunAsNonRoot           *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
			RunAsUser              *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
			SeLinuxOptions         *struct {
				Level *string `tfsdk:"level" json:"level,omitempty"`
				Role  *string `tfsdk:"role" json:"role,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				User  *string `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
			SeccompProfile *struct {
				LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
				Type             *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
			WindowsOptions *struct {
				GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
				GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
				HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
				RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
			} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
		} `tfsdk:"container_security_context" json:"containerSecurityContext,omitempty"`
		DisableHealthCheck    *bool `tfsdk:"disable_health_check" json:"disableHealthCheck,omitempty"`
		EnableMetricsEndpoint *bool `tfsdk:"enable_metrics_endpoint" json:"enableMetricsEndpoint,omitempty"`
		Env                   *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				FieldRef *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
				ResourceFieldRef *struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
					Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
				} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
				SecretKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"env" json:"env,omitempty"`
		EnvFrom *[]struct {
			ConfigMapRef *struct {
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
			Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
			SecretRef *struct {
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"env_from" json:"envFrom,omitempty"`
		HostNetwork      *bool   `tfsdk:"host_network" json:"hostNetwork,omitempty"`
		Image            *string `tfsdk:"image" json:"image,omitempty"`
		ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		SecurityContext *struct {
			FsGroup             *int64  `tfsdk:"fs_group" json:"fsGroup,omitempty"`
			FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" json:"fsGroupChangePolicy,omitempty"`
			RunAsGroup          *int64  `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
			RunAsNonRoot        *bool   `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
			RunAsUser           *int64  `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
			SeLinuxOptions      *struct {
				Level *string `tfsdk:"level" json:"level,omitempty"`
				Role  *string `tfsdk:"role" json:"role,omitempty"`
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				User  *string `tfsdk:"user" json:"user,omitempty"`
			} `tfsdk:"se_linux_options" json:"seLinuxOptions,omitempty"`
			SeccompProfile *struct {
				LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
				Type             *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
			SupplementalGroups *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
			Sysctls            *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"sysctls" json:"sysctls,omitempty"`
			WindowsOptions *struct {
				GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
				GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
				HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
				RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
			} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
		} `tfsdk:"security_context" json:"securityContext,omitempty"`
		TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
		Tolerations                   *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		TopologySpreadConstraints *[]struct {
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			MatchLabelKeys     *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
			MaxSkew            *int64    `tfsdk:"max_skew" json:"maxSkew,omitempty"`
			MinDomains         *int64    `tfsdk:"min_domains" json:"minDomains,omitempty"`
			NodeAffinityPolicy *string   `tfsdk:"node_affinity_policy" json:"nodeAffinityPolicy,omitempty"`
			NodeTaintsPolicy   *string   `tfsdk:"node_taints_policy" json:"nodeTaintsPolicy,omitempty"`
			TopologyKey        *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
			WhenUnsatisfiable  *string   `tfsdk:"when_unsatisfiable" json:"whenUnsatisfiable,omitempty"`
		} `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StunnerL7MpIoDataplaneV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_stunner_l7mp_io_dataplane_v1_manifest"
}

func (r *StunnerL7MpIoDataplaneV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Dataplane is a collection of configuration parameters that can be used for spawning a 'stunnerd'instance for a Gateway. Labels and annotations on the Dataplane object will be copied verbatiminto the target Deployment.",
		MarkdownDescription: "Dataplane is a collection of configuration parameters that can be used for spawning a 'stunnerd'instance for a Gateway. Labels and annotations on the Dataplane object will be copied verbatiminto the target Deployment.",
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
				Description:         "Spec defines the behavior of a Dataplane resource.",
				MarkdownDescription: "Spec defines the behavior of a Dataplane resource.",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Scheduling constraints.",
						MarkdownDescription: "Scheduling constraints.",
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

					"args": schema.ListAttribute{
						Description:         "Arguments to the entrypoint.",
						MarkdownDescription: "Arguments to the entrypoint.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"command": schema.ListAttribute{
						Description:         "Entrypoint array. Defaults: 'stunnerd'.",
						MarkdownDescription: "Entrypoint array. Defaults: 'stunnerd'.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"container_security_context": schema.SingleNestedAttribute{
						Description:         "ContainerSecurityContext holds container-level security attributes specifically for thestunnerd container.",
						MarkdownDescription: "ContainerSecurityContext holds container-level security attributes specifically for thestunnerd container.",
						Attributes: map[string]schema.Attribute{
							"allow_privilege_escalation": schema.BoolAttribute{
								Description:         "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"capabilities": schema.SingleNestedAttribute{
								Description:         "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"add": schema.ListAttribute{
										Description:         "Added capabilities",
										MarkdownDescription: "Added capabilities",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"drop": schema.ListAttribute{
										Description:         "Removed capabilities",
										MarkdownDescription: "Removed capabilities",
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

							"privileged": schema.BoolAttribute{
								Description:         "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proc_mount": schema.StringAttribute{
								Description:         "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"read_only_root_filesystem": schema.BoolAttribute{
								Description:         "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_group": schema.Int64Attribute{
								Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_non_root": schema.BoolAttribute{
								Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
								MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_user": schema.Int64Attribute{
								Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"se_linux_options": schema.SingleNestedAttribute{
								Description:         "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"level": schema.StringAttribute{
										Description:         "Level is SELinux level label that applies to the container.",
										MarkdownDescription: "Level is SELinux level label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"role": schema.StringAttribute{
										Description:         "Role is a SELinux role label that applies to the container.",
										MarkdownDescription: "Role is a SELinux role label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is a SELinux type label that applies to the container.",
										MarkdownDescription: "Type is a SELinux type label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user": schema.StringAttribute{
										Description:         "User is a SELinux user label that applies to the container.",
										MarkdownDescription: "User is a SELinux user label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"seccomp_profile": schema.SingleNestedAttribute{
								Description:         "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"localhost_profile": schema.StringAttribute{
										Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
										MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
										MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"windows_options": schema.SingleNestedAttribute{
								Description:         "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
								MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
								Attributes: map[string]schema.Attribute{
									"gmsa_credential_spec": schema.StringAttribute{
										Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
										MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"gmsa_credential_spec_name": schema.StringAttribute{
										Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
										MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_process": schema.BoolAttribute{
										Description:         "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
										MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_user_name": schema.StringAttribute{
										Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_health_check": schema.BoolAttribute{
						Description:         "Disable health-checking. Default is to enable HTTP health-checks on port 8086: aliveness probe responder will be exposed on path '/live' and readiness probe on path'/ready'.",
						MarkdownDescription: "Disable health-checking. Default is to enable HTTP health-checks on port 8086: aliveness probe responder will be exposed on path '/live' and readiness probe on path'/ready'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_metrics_endpoint": schema.BoolAttribute{
						Description:         "EnableMetricsEnpoint can be used to enable metrics scraping (Prometheus). If enabled, ametrics endpoint will be available at http://0.0.0.0:8080 at all dataplane pods. Defaultis no metrics collection.",
						MarkdownDescription: "EnableMetricsEnpoint can be used to enable metrics scraping (Prometheus). If enabled, ametrics endpoint will be available at http://0.0.0.0:8080 at all dataplane pods. Defaultis no metrics collection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env": schema.ListNestedAttribute{
						Description:         "List of environment variables to set in the stunnerd container.",
						MarkdownDescription: "List of environment variables to set in the stunnerd container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
									MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
									MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value_from": schema.SingleNestedAttribute{
									Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
									MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a ConfigMap.",
											MarkdownDescription: "Selects a key of a ConfigMap.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key to select.",
													MarkdownDescription: "The key to select.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap or its key must be defined",
													MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"field_ref": schema.SingleNestedAttribute{
											Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
													MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
													Description:         "Path of the field to select in the specified API version.",
													MarkdownDescription: "Path of the field to select in the specified API version.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resource_field_ref": schema.SingleNestedAttribute{
											Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											Attributes: map[string]schema.Attribute{
												"container_name": schema.StringAttribute{
													Description:         "Container name: required for volumes, optional for env vars",
													MarkdownDescription: "Container name: required for volumes, optional for env vars",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"divisor": schema.StringAttribute{
													Description:         "Specifies the output format of the exposed resources, defaults to '1'",
													MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource": schema.StringAttribute{
													Description:         "Required: resource to select",
													MarkdownDescription: "Required: resource to select",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a secret in the pod's namespace",
											MarkdownDescription: "Selects a key of a secret in the pod's namespace",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

					"env_from": schema.ListNestedAttribute{
						Description:         "List of sources to populate environment variables in the stunnerd container.",
						MarkdownDescription: "List of sources to populate environment variables in the stunnerd container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_map_ref": schema.SingleNestedAttribute{
									Description:         "The ConfigMap to select from",
									MarkdownDescription: "The ConfigMap to select from",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the ConfigMap must be defined",
											MarkdownDescription: "Specify whether the ConfigMap must be defined",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"prefix": schema.StringAttribute{
									Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
									MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "The Secret to select from",
									MarkdownDescription: "The Secret to select from",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the Secret must be defined",
											MarkdownDescription: "Specify whether the Secret must be defined",
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

					"host_network": schema.BoolAttribute{
						Description:         "Host networking requested for the stunnerd pod to use the host's network namespace.Can be used to implement public TURN servers with Kubernetes.  Defaults to false.",
						MarkdownDescription: "Host networking requested for the stunnerd pod to use the host's network namespace.Can be used to implement public TURN servers with Kubernetes.  Defaults to false.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Container image name.",
						MarkdownDescription: "Container image name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "Image pull policy. One of Always, Never, IfNotPresent.",
						MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "ImagePullSecrets is an optional list of references to secrets to use for pulling thestunnerd image. Note that the referenced secrets are not watched by the operator, somodifications will in effect only for newly created pods. Also note that the Secret isalways searched in the same namespace as the Gateway, which allows to use separate pullsecrets per each namespace.",
						MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets to use for pulling thestunnerd image. Note that the referenced secrets are not watched by the operator, somodifications will in effect only for newly created pods. Also note that the Secret isalways searched in the same namespace as the Gateway, which allows to use separate pullsecrets per each namespace.",
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

					"replicas": schema.Int64Attribute{
						Description:         "Number of desired pods. If empty or set to 1, use whatever is in the target Deployment.Otherwise, enforce this setting, overwiting whatever is set in the Deployment (this mayblock autoscaling the dataplane though). Defaults to 1.",
						MarkdownDescription: "Number of desired pods. If empty or set to 1, use whatever is in the target Deployment.Otherwise, enforce this setting, overwiting whatever is set in the Deployment (this mayblock autoscaling the dataplane though). Defaults to 1.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources required by stunnerd.",
						MarkdownDescription: "Resources required by stunnerd.",
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

					"security_context": schema.SingleNestedAttribute{
						Description:         "SecurityContext holds pod-level security attributes and common container settings.",
						MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.",
						Attributes: map[string]schema.Attribute{
							"fs_group": schema.Int64Attribute{
								Description:         "A special supplemental group that applies to all containers in a pod.Some volume types allow the Kubelet to change the ownership of that volumeto be owned by the pod:1. The owning GID will be the FSGroup2. The setgid bit is set (new files created in the volume will be owned by FSGroup)3. The permission bits are OR'd with rw-rw----If unset, the Kubelet will not modify the ownership and permissions of any volume.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "A special supplemental group that applies to all containers in a pod.Some volume types allow the Kubelet to change the ownership of that volumeto be owned by the pod:1. The owning GID will be the FSGroup2. The setgid bit is set (new files created in the volume will be owned by FSGroup)3. The permission bits are OR'd with rw-rw----If unset, the Kubelet will not modify the ownership and permissions of any volume.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fs_group_change_policy": schema.StringAttribute{
								Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volumebefore being exposed inside Pod. This field will only apply tovolume types which support fsGroup based ownership(and permissions).It will have no effect on ephemeral volume types such as: secret, configmapsand emptydir.Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volumebefore being exposed inside Pod. This field will only apply tovolume types which support fsGroup based ownership(and permissions).It will have no effect on ephemeral volume types such as: secret, configmapsand emptydir.Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_group": schema.Int64Attribute{
								Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_non_root": schema.BoolAttribute{
								Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
								MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"run_as_user": schema.Int64Attribute{
								Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"se_linux_options": schema.SingleNestedAttribute{
								Description:         "The SELinux context to be applied to all containers.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in SecurityContext.  If set inboth SecurityContext and PodSecurityContext, the value specified in SecurityContexttakes precedence for that container.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The SELinux context to be applied to all containers.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in SecurityContext.  If set inboth SecurityContext and PodSecurityContext, the value specified in SecurityContexttakes precedence for that container.Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"level": schema.StringAttribute{
										Description:         "Level is SELinux level label that applies to the container.",
										MarkdownDescription: "Level is SELinux level label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"role": schema.StringAttribute{
										Description:         "Role is a SELinux role label that applies to the container.",
										MarkdownDescription: "Role is a SELinux role label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is a SELinux type label that applies to the container.",
										MarkdownDescription: "Type is a SELinux type label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user": schema.StringAttribute{
										Description:         "User is a SELinux user label that applies to the container.",
										MarkdownDescription: "User is a SELinux user label that applies to the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"seccomp_profile": schema.SingleNestedAttribute{
								Description:         "The seccomp options to use by the containers in this pod.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "The seccomp options to use by the containers in this pod.Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"localhost_profile": schema.StringAttribute{
										Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
										MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
										MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"supplemental_groups": schema.ListAttribute{
								Description:         "A list of groups applied to the first process run in each container, in additionto the container's primary GID, the fsGroup (if specified), and group membershipsdefined in the container image for the uid of the container process. If unspecified,no additional groups are added to any container. Note that group membershipsdefined in the container image for the uid of the container process are still effective,even if they are not included in this list.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "A list of groups applied to the first process run in each container, in additionto the container's primary GID, the fsGroup (if specified), and group membershipsdefined in the container image for the uid of the container process. If unspecified,no additional groups are added to any container. Note that group membershipsdefined in the container image for the uid of the container process are still effective,even if they are not included in this list.Note that this field cannot be set when spec.os.name is windows.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sysctls": schema.ListNestedAttribute{
								Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupportedsysctls (by the container runtime) might fail to launch.Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupportedsysctls (by the container runtime) might fail to launch.Note that this field cannot be set when spec.os.name is windows.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of a property to set",
											MarkdownDescription: "Name of a property to set",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value of a property to set",
											MarkdownDescription: "Value of a property to set",
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

							"windows_options": schema.SingleNestedAttribute{
								Description:         "The Windows specific settings applied to all containers.If unspecified, the options within a container's SecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
								MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options within a container's SecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
								Attributes: map[string]schema.Attribute{
									"gmsa_credential_spec": schema.StringAttribute{
										Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
										MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"gmsa_credential_spec_name": schema.StringAttribute{
										Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
										MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_process": schema.BoolAttribute{
										Description:         "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
										MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_user_name": schema.StringAttribute{
										Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"termination_grace_period_seconds": schema.Int64Attribute{
						Description:         "Optional duration in seconds the stunnerd needs to terminate gracefully. Defaults to 3600 seconds.",
						MarkdownDescription: "Optional duration in seconds the stunnerd needs to terminate gracefully. Defaults to 3600 seconds.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "If specified, the pod's tolerations.",
						MarkdownDescription: "If specified, the pod's tolerations.",
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

					"topology_spread_constraints": schema.ListNestedAttribute{
						Description:         "TopologySpreadConstraints describes how stunnerd pods ought to spread across topologydomains.",
						MarkdownDescription: "TopologySpreadConstraints describes how stunnerd pods ought to spread across topologydomains.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"label_selector": schema.SingleNestedAttribute{
									Description:         "LabelSelector is used to find matching pods.Pods that match this label selector are counted to determine the number of podsin their corresponding topology domain.",
									MarkdownDescription: "LabelSelector is used to find matching pods.Pods that match this label selector are counted to determine the number of podsin their corresponding topology domain.",
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
									Description:         "MatchLabelKeys is a set of pod label keys to select the pods over whichspreading will be calculated. The keys are used to lookup values from theincoming pod labels, those key-value labels are ANDed with labelSelectorto select the group of existing pods over which spreading will be calculatedfor the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.MatchLabelKeys cannot be set when LabelSelector isn't set.Keys that don't exist in the incoming pod labels willbe ignored. A null or empty list means only match against labelSelector.This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
									MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over whichspreading will be calculated. The keys are used to lookup values from theincoming pod labels, those key-value labels are ANDed with labelSelectorto select the group of existing pods over which spreading will be calculatedfor the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.MatchLabelKeys cannot be set when LabelSelector isn't set.Keys that don't exist in the incoming pod labels willbe ignored. A null or empty list means only match against labelSelector.This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_skew": schema.Int64Attribute{
									Description:         "MaxSkew describes the degree to which pods may be unevenly distributed.When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted differencebetween the number of matching pods in the target topology and the global minimum.The global minimum is the minimum number of matching pods in an eligible domainor zero if the number of eligible domains is less than MinDomains.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 2/2/1:In this case, the global minimum is 1.| zone1 | zone2 | zone3 ||  P P  |  P P  |   P   |- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)violate MaxSkew(1).- if MaxSkew is 2, incoming pod can be scheduled onto any zone.When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedenceto topologies that satisfy it.It's a required field. Default value is 1 and 0 is not allowed.",
									MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed.When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted differencebetween the number of matching pods in the target topology and the global minimum.The global minimum is the minimum number of matching pods in an eligible domainor zero if the number of eligible domains is less than MinDomains.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 2/2/1:In this case, the global minimum is 1.| zone1 | zone2 | zone3 ||  P P  |  P P  |   P   |- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)violate MaxSkew(1).- if MaxSkew is 2, incoming pod can be scheduled onto any zone.When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedenceto topologies that satisfy it.It's a required field. Default value is 1 and 0 is not allowed.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"min_domains": schema.Int64Attribute{
									Description:         "MinDomains indicates a minimum number of eligible domains.When the number of eligible domains with matching topology keys is less than minDomains,Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed.And when the number of eligible domains with matching topology keys equals or greater than minDomains,this value has no effect on scheduling.As a result, when the number of eligible domains is less than minDomains,scheduler won't schedule more than maxSkew Pods to those domains.If value is nil, the constraint behaves as if MinDomains is equal to 1.Valid values are integers greater than 0.When value is not nil, WhenUnsatisfiable must be DoNotSchedule.For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the samelabelSelector spread as 2/2/2:| zone1 | zone2 | zone3 ||  P P  |  P P  |  P P  |The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0.In this situation, new pod with the same labelSelector cannot be scheduled,because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,it will violate MaxSkew.This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
									MarkdownDescription: "MinDomains indicates a minimum number of eligible domains.When the number of eligible domains with matching topology keys is less than minDomains,Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed.And when the number of eligible domains with matching topology keys equals or greater than minDomains,this value has no effect on scheduling.As a result, when the number of eligible domains is less than minDomains,scheduler won't schedule more than maxSkew Pods to those domains.If value is nil, the constraint behaves as if MinDomains is equal to 1.Valid values are integers greater than 0.When value is not nil, WhenUnsatisfiable must be DoNotSchedule.For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the samelabelSelector spread as 2/2/2:| zone1 | zone2 | zone3 ||  P P  |  P P  |  P P  |The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0.In this situation, new pod with the same labelSelector cannot be scheduled,because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,it will violate MaxSkew.This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"node_affinity_policy": schema.StringAttribute{
									Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelectorwhen calculating pod topology spread skew. Options are:- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.If this value is nil, the behavior is equivalent to the Honor policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
									MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelectorwhen calculating pod topology spread skew. Options are:- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.If this value is nil, the behavior is equivalent to the Honor policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"node_taints_policy": schema.StringAttribute{
									Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculatingpod topology spread skew. Options are:- Honor: nodes without taints, along with tainted nodes for which the incoming podhas a toleration, are included.- Ignore: node taints are ignored. All nodes are included.If this value is nil, the behavior is equivalent to the Ignore policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
									MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculatingpod topology spread skew. Options are:- Honor: nodes without taints, along with tainted nodes for which the incoming podhas a toleration, are included.- Ignore: node taints are ignored. All nodes are included.If this value is nil, the behavior is equivalent to the Ignore policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"topology_key": schema.StringAttribute{
									Description:         "TopologyKey is the key of node labels. Nodes that have a label with this keyand identical values are considered to be in the same topology.We consider each <key, value> as a 'bucket', and try to put balanced numberof pods into each bucket.We define a domain as a particular instance of a topology.Also, we define an eligible domain as a domain whose nodes meet the requirements ofnodeAffinityPolicy and nodeTaintsPolicy.e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology.And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology.It's a required field.",
									MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this keyand identical values are considered to be in the same topology.We consider each <key, value> as a 'bucket', and try to put balanced numberof pods into each bucket.We define a domain as a particular instance of a topology.Also, we define an eligible domain as a domain whose nodes meet the requirements ofnodeAffinityPolicy and nodeTaintsPolicy.e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology.And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology.It's a required field.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"when_unsatisfiable": schema.StringAttribute{
									Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfythe spread constraint.- DoNotSchedule (default) tells the scheduler not to schedule it.- ScheduleAnyway tells the scheduler to schedule the pod in any location,  but giving higher precedence to topologies that would help reduce the  skew.A constraint is considered 'Unsatisfiable' for an incoming podif and only if every possible node assignment for that pod would violate'MaxSkew' on some topology.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 3/1/1:| zone1 | zone2 | zone3 || P P P |   P   |   P   |If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduledto zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfiesMaxSkew(1). In other words, the cluster can still be imbalanced, but schedulerwon't make it *more* imbalanced.It's a required field.",
									MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfythe spread constraint.- DoNotSchedule (default) tells the scheduler not to schedule it.- ScheduleAnyway tells the scheduler to schedule the pod in any location,  but giving higher precedence to topologies that would help reduce the  skew.A constraint is considered 'Unsatisfiable' for an incoming podif and only if every possible node assignment for that pod would violate'MaxSkew' on some topology.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 3/1/1:| zone1 | zone2 | zone3 || P P P |   P   |   P   |If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduledto zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfiesMaxSkew(1). In other words, the cluster can still be imbalanced, but schedulerwon't make it *more* imbalanced.It's a required field.",
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
	}
}

func (r *StunnerL7MpIoDataplaneV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_stunner_l7mp_io_dataplane_v1_manifest")

	var model StunnerL7MpIoDataplaneV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("stunner.l7mp.io/v1")
	model.Kind = pointer.String("Dataplane")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
