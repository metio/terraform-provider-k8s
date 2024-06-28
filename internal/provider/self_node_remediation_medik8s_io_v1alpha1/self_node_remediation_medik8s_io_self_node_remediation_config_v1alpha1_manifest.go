/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package self_node_remediation_medik8s_io_v1alpha1

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
	_ datasource.DataSource = &SelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1Manifest{}
)

func NewSelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1Manifest() datasource.DataSource {
	return &SelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1Manifest{}
}

type SelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1Manifest struct{}

type SelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1ManifestData struct {
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
		ApiCheckInterval    *string `tfsdk:"api_check_interval" json:"apiCheckInterval,omitempty"`
		ApiServerTimeout    *string `tfsdk:"api_server_timeout" json:"apiServerTimeout,omitempty"`
		CustomDsTolerations *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"custom_ds_tolerations" json:"customDsTolerations,omitempty"`
		EndpointHealthCheckUrl              *string `tfsdk:"endpoint_health_check_url" json:"endpointHealthCheckUrl,omitempty"`
		HostPort                            *int64  `tfsdk:"host_port" json:"hostPort,omitempty"`
		IsSoftwareRebootEnabled             *bool   `tfsdk:"is_software_reboot_enabled" json:"isSoftwareRebootEnabled,omitempty"`
		MaxApiErrorThreshold                *int64  `tfsdk:"max_api_error_threshold" json:"maxApiErrorThreshold,omitempty"`
		PeerApiServerTimeout                *string `tfsdk:"peer_api_server_timeout" json:"peerApiServerTimeout,omitempty"`
		PeerDialTimeout                     *string `tfsdk:"peer_dial_timeout" json:"peerDialTimeout,omitempty"`
		PeerRequestTimeout                  *string `tfsdk:"peer_request_timeout" json:"peerRequestTimeout,omitempty"`
		PeerUpdateInterval                  *string `tfsdk:"peer_update_interval" json:"peerUpdateInterval,omitempty"`
		SafeTimeToAssumeNodeRebootedSeconds *int64  `tfsdk:"safe_time_to_assume_node_rebooted_seconds" json:"safeTimeToAssumeNodeRebootedSeconds,omitempty"`
		WatchdogFilePath                    *string `tfsdk:"watchdog_file_path" json:"watchdogFilePath,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_self_node_remediation_medik8s_io_self_node_remediation_config_v1alpha1_manifest"
}

func (r *SelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "SelfNodeRemediationConfig is the Schema for the selfnoderemediationconfigs API in which a user can configure the self node remediation agents",
		MarkdownDescription: "SelfNodeRemediationConfig is the Schema for the selfnoderemediationconfigs API in which a user can configure the self node remediation agents",
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
				Description:         "SelfNodeRemediationConfigSpec defines the desired state of SelfNodeRemediationConfig",
				MarkdownDescription: "SelfNodeRemediationConfigSpec defines the desired state of SelfNodeRemediationConfig",
				Attributes: map[string]schema.Attribute{
					"api_check_interval": schema.StringAttribute{
						Description:         "The frequency for api-server connectivity check.Valid time units are 'ms', 's', 'm', 'h'.the frequency for api-server connectivity check",
						MarkdownDescription: "The frequency for api-server connectivity check.Valid time units are 'ms', 's', 'm', 'h'.the frequency for api-server connectivity check",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"api_server_timeout": schema.StringAttribute{
						Description:         "Timeout for each api-connectivity check.Valid time units are 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Timeout for each api-connectivity check.Valid time units are 'ms', 's', 'm', 'h'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"custom_ds_tolerations": schema.ListNestedAttribute{
						Description:         "CustomDsTolerations allows to add custom tolerations snr agents that are running on the ds in order to support remediation for different types of nodes.",
						MarkdownDescription: "CustomDsTolerations allows to add custom tolerations snr agents that are running on the ds in order to support remediation for different types of nodes.",
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

					"endpoint_health_check_url": schema.StringAttribute{
						Description:         "EndpointHealthCheckUrl is an url that self node remediation agents which run on control-plane node will try to access when they can't contact their peers.This is a part of self diagnostics which will decide whether the node should be remediated or not.It will be ignored when empty (which is the default).",
						MarkdownDescription: "EndpointHealthCheckUrl is an url that self node remediation agents which run on control-plane node will try to access when they can't contact their peers.This is a part of self diagnostics which will decide whether the node should be remediated or not.It will be ignored when empty (which is the default).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_port": schema.Int64Attribute{
						Description:         "HostPort is used for internal communication between SNR agents.",
						MarkdownDescription: "HostPort is used for internal communication between SNR agents.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"is_software_reboot_enabled": schema.BoolAttribute{
						Description:         "IsSoftwareRebootEnabled indicates whether self node remediation agent will do software reboot,if the watchdog device can not be used or will use watchdog only,without a fallback to software reboot.",
						MarkdownDescription: "IsSoftwareRebootEnabled indicates whether self node remediation agent will do software reboot,if the watchdog device can not be used or will use watchdog only,without a fallback to software reboot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_api_error_threshold": schema.Int64Attribute{
						Description:         "After this threshold, the node will start contacting its peers.",
						MarkdownDescription: "After this threshold, the node will start contacting its peers.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"peer_api_server_timeout": schema.StringAttribute{
						Description:         "The timeout for api-server connectivity check.Valid time units are 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "The timeout for api-server connectivity check.Valid time units are 'ms', 's', 'm', 'h'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"peer_dial_timeout": schema.StringAttribute{
						Description:         "Timeout for establishing connection to peer.Valid time units are 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Timeout for establishing connection to peer.Valid time units are 'ms', 's', 'm', 'h'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"peer_request_timeout": schema.StringAttribute{
						Description:         "Timeout for each peer request.Valid time units are 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "Timeout for each peer request.Valid time units are 'ms', 's', 'm', 'h'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"peer_update_interval": schema.StringAttribute{
						Description:         "The frequency for updating peers.Valid time units are 'ms', 's', 'm', 'h'.",
						MarkdownDescription: "The frequency for updating peers.Valid time units are 'ms', 's', 'm', 'h'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+(\.[0-9]+)?(ns|us|µs|ms|s|m|h))+$`), ""),
						},
					},

					"safe_time_to_assume_node_rebooted_seconds": schema.Int64Attribute{
						Description:         "SafeTimeToAssumeNodeRebootedSeconds is the time after which the healthy self node remediationagents will assume the unhealthy node has been rebooted, and it is safe to recover affected workloads.This is extremely important as starting replacement Pods while they are still running on the failednode will likely lead to data corruption and violation of run-once semantics.In an effort to prevent this, the operator ignores values lower than a minimum calculated from theApiCheckInterval, ApiServerTimeout, MaxApiErrorThreshold, PeerDialTimeout, and PeerRequestTimeout fields,and the unhealthy node's individual watchdog timeout.",
						MarkdownDescription: "SafeTimeToAssumeNodeRebootedSeconds is the time after which the healthy self node remediationagents will assume the unhealthy node has been rebooted, and it is safe to recover affected workloads.This is extremely important as starting replacement Pods while they are still running on the failednode will likely lead to data corruption and violation of run-once semantics.In an effort to prevent this, the operator ignores values lower than a minimum calculated from theApiCheckInterval, ApiServerTimeout, MaxApiErrorThreshold, PeerDialTimeout, and PeerRequestTimeout fields,and the unhealthy node's individual watchdog timeout.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"watchdog_file_path": schema.StringAttribute{
						Description:         "WatchdogFilePath is the watchdog file path that should be available on each node, e.g. /dev/watchdog.",
						MarkdownDescription: "WatchdogFilePath is the watchdog file path that should be available on each node, e.g. /dev/watchdog.",
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

func (r *SelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_self_node_remediation_medik8s_io_self_node_remediation_config_v1alpha1_manifest")

	var model SelfNodeRemediationMedik8SIoSelfNodeRemediationConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("self-node-remediation.medik8s.io/v1alpha1")
	model.Kind = pointer.String("SelfNodeRemediationConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
