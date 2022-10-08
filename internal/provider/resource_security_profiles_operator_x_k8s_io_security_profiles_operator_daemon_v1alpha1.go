/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource)(nil)
)

type SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AllowedSeccompActions *[]string `tfsdk:"allowed_seccomp_actions" yaml:"allowedSeccompActions,omitempty"`

		AllowedSyscalls *[]string `tfsdk:"allowed_syscalls" yaml:"allowedSyscalls,omitempty"`

		EnableAppArmor *bool `tfsdk:"enable_app_armor" yaml:"enableAppArmor,omitempty"`

		EnableBpfRecorder *bool `tfsdk:"enable_bpf_recorder" yaml:"enableBpfRecorder,omitempty"`

		EnableLogEnricher *bool `tfsdk:"enable_log_enricher" yaml:"enableLogEnricher,omitempty"`

		EnableProfiling *bool `tfsdk:"enable_profiling" yaml:"enableProfiling,omitempty"`

		EnableSelinux *bool `tfsdk:"enable_selinux" yaml:"enableSelinux,omitempty"`

		HostProcVolumePath *string `tfsdk:"host_proc_volume_path" yaml:"hostProcVolumePath,omitempty"`

		SelinuxOptions *struct {
			AllowedSystemProfiles *[]string `tfsdk:"allowed_system_profiles" yaml:"allowedSystemProfiles,omitempty"`
		} `tfsdk:"selinux_options" yaml:"selinuxOptions,omitempty"`

		SelinuxTypeTag *string `tfsdk:"selinux_type_tag" yaml:"selinuxTypeTag,omitempty"`

		StaticWebhookConfig *bool `tfsdk:"static_webhook_config" yaml:"staticWebhookConfig,omitempty"`

		Tolerations *[]struct {
			Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

			TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

		Verbosity *int64 `tfsdk:"verbosity" yaml:"verbosity,omitempty"`

		WebhookOptions *[]struct {
			FailurePolicy *string `tfsdk:"failure_policy" yaml:"failurePolicy,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			NamespaceSelector *struct {
				MatchExpressions *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`
		} `tfsdk:"webhook_options" yaml:"webhookOptions,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewSecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource() resource.Resource {
	return &SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource{}
}

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1"
}

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "SecurityProfilesOperatorDaemon is the Schema to configure the spod deployment.",
		MarkdownDescription: "SecurityProfilesOperatorDaemon is the Schema to configure the spod deployment.",
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
						PlanModifiers: []tfsdk.AttributePlanModifier{
							resource.RequiresReplace(),
						},
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
				Description:         "SPODStatus defines the desired state of SPOD.",
				MarkdownDescription: "SPODStatus defines the desired state of SPOD.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"allowed_seccomp_actions": {
						Description:         "AllowedSeccompActions if specified, a list of allowed seccomp actions.",
						MarkdownDescription: "AllowedSeccompActions if specified, a list of allowed seccomp actions.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"allowed_syscalls": {
						Description:         "AllowedSyscalls if specified, a list of system calls which are allowed in seccomp profiles.",
						MarkdownDescription: "AllowedSyscalls if specified, a list of system calls which are allowed in seccomp profiles.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_app_armor": {
						Description:         "tells the operator whether or not to enable AppArmor support for this SPOD instance.",
						MarkdownDescription: "tells the operator whether or not to enable AppArmor support for this SPOD instance.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_bpf_recorder": {
						Description:         "tells the operator whether or not to enable bpf recorder support for this SPOD instance.",
						MarkdownDescription: "tells the operator whether or not to enable bpf recorder support for this SPOD instance.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_log_enricher": {
						Description:         "tells the operator whether or not to enable log enrichment support for this SPOD instance.",
						MarkdownDescription: "tells the operator whether or not to enable log enrichment support for this SPOD instance.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_profiling": {
						Description:         "EnableProfiling tells the operator whether or not to enable profiling support for this SPOD instance.",
						MarkdownDescription: "EnableProfiling tells the operator whether or not to enable profiling support for this SPOD instance.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_selinux": {
						Description:         "tells the operator whether or not to enable SELinux support for this SPOD instance.",
						MarkdownDescription: "tells the operator whether or not to enable SELinux support for this SPOD instance.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host_proc_volume_path": {
						Description:         "HostProcVolumePath is the path for specifying a custom host /proc volume, which is required for the log-enricher as well as bpf-recorder to retrieve the container ID for a process ID. This can be helpful for nested environments, for example when using 'kind'.",
						MarkdownDescription: "HostProcVolumePath is the path for specifying a custom host /proc volume, which is required for the log-enricher as well as bpf-recorder to retrieve the container ID for a process ID. This can be helpful for nested environments, for example when using 'kind'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selinux_options": {
						Description:         "Defines options specific to the SELinux functionality of the SecurityProfilesOperator",
						MarkdownDescription: "Defines options specific to the SELinux functionality of the SecurityProfilesOperator",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allowed_system_profiles": {
								Description:         "Lists the profiles coming from the system itself that are allowed to be inherited by workloads. Use this with care, as this might provide a lot of permissions depending on the policy.",
								MarkdownDescription: "Lists the profiles coming from the system itself that are allowed to be inherited by workloads. Use this with care, as this might provide a lot of permissions depending on the policy.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selinux_type_tag": {
						Description:         "If specified, the SELinux type tag applied to the security context of SPOD.",
						MarkdownDescription: "If specified, the SELinux type tag applied to the security context of SPOD.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"static_webhook_config": {
						Description:         "StaticWebhookConfig indicates whether the webhook configuration and its related resources are statically deployed. In this case, the operator will not create or update the webhook configuration and its related resources.",
						MarkdownDescription: "StaticWebhookConfig indicates whether the webhook configuration and its related resources are statically deployed. In this case, the operator will not create or update the webhook configuration and its related resources.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tolerations": {
						Description:         "If specified, the SPOD's tolerations.",
						MarkdownDescription: "If specified, the SPOD's tolerations.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"effect": {
								Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
								MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"key": {
								Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
								MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"operator": {
								Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
								MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"toleration_seconds": {
								Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
								MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
								MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

					"verbosity": {
						Description:         "Verbosity specifies the logging verbosity of the daemon.",
						MarkdownDescription: "Verbosity specifies the logging verbosity of the daemon.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"webhook_options": {
						Description:         "WebhookOpts set custom namespace selectors and failure mode for SPO's webhooks",
						MarkdownDescription: "WebhookOpts set custom namespace selectors and failure mode for SPO's webhooks",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"failure_policy": {
								Description:         "FailurePolicy sets the webhook failure policy",
								MarkdownDescription: "FailurePolicy sets the webhook failure policy",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name specifies which webhook do we configure",
								MarkdownDescription: "Name specifies which webhook do we configure",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace_selector": {
								Description:         "NamespaceSelector sets webhook's namespace selector",
								MarkdownDescription: "NamespaceSelector sets webhook's namespace selector",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"match_expressions": {
										Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
										MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the label key that the selector applies to.",
												MarkdownDescription: "key is the label key that the selector applies to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"operator": {
												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"values": {
												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match_labels": {
										Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
										MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",

										Type: types.MapType{ElemType: types.StringType},

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

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1")

	var state SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("security-profiles-operator.x-k8s.io/v1alpha1")
	goModel.Kind = utilities.Ptr("SecurityProfilesOperatorDaemon")

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

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1")

	var state SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("security-profiles-operator.x-k8s.io/v1alpha1")
	goModel.Kind = utilities.Ptr("SecurityProfilesOperatorDaemon")

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

func (r *SecurityProfilesOperatorXK8SIoSecurityProfilesOperatorDaemonV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
