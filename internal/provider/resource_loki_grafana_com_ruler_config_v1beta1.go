/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type LokiGrafanaComRulerConfigV1Beta1Resource struct{}

var (
	_ resource.Resource = (*LokiGrafanaComRulerConfigV1Beta1Resource)(nil)
)

type LokiGrafanaComRulerConfigV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LokiGrafanaComRulerConfigV1Beta1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Alertmanager *struct {
			Discovery *struct {
				EnableSRV *bool `tfsdk:"enable_srv" yaml:"enableSRV,omitempty"`

				RefreshInterval *string `tfsdk:"refresh_interval" yaml:"refreshInterval,omitempty"`
			} `tfsdk:"discovery" yaml:"discovery,omitempty"`

			EnableV2 *bool `tfsdk:"enable_v2" yaml:"enableV2,omitempty"`

			Endpoints *[]string `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

			ExternalLabels *map[string]string `tfsdk:"external_labels" yaml:"externalLabels,omitempty"`

			ExternalUrl *string `tfsdk:"external_url" yaml:"externalUrl,omitempty"`

			NotificationQueue *struct {
				Capacity *int64 `tfsdk:"capacity" yaml:"capacity,omitempty"`

				ForGracePeriod *string `tfsdk:"for_grace_period" yaml:"forGracePeriod,omitempty"`

				ForOutageTolerance *string `tfsdk:"for_outage_tolerance" yaml:"forOutageTolerance,omitempty"`

				ResendDelay *string `tfsdk:"resend_delay" yaml:"resendDelay,omitempty"`

				Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
			} `tfsdk:"notification_queue" yaml:"notificationQueue,omitempty"`

			RelabelConfigs *[]struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				Modulus *int64 `tfsdk:"modulus" yaml:"modulus,omitempty"`

				Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

				Replacement *string `tfsdk:"replacement" yaml:"replacement,omitempty"`

				Separator *string `tfsdk:"separator" yaml:"separator,omitempty"`

				SourceLabels *[]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

				TargetLabel *string `tfsdk:"target_label" yaml:"targetLabel,omitempty"`
			} `tfsdk:"relabel_configs" yaml:"relabelConfigs,omitempty"`
		} `tfsdk:"alertmanager" yaml:"alertmanager,omitempty"`

		EvaluationInterval *string `tfsdk:"evaluation_interval" yaml:"evaluationInterval,omitempty"`

		PollInterval *string `tfsdk:"poll_interval" yaml:"pollInterval,omitempty"`

		RemoteWrite *struct {
			Client *struct {
				AdditionalHeaders *map[string]string `tfsdk:"additional_headers" yaml:"additionalHeaders,omitempty"`

				Authorization *string `tfsdk:"authorization" yaml:"authorization,omitempty"`

				AuthorizationSecretName *string `tfsdk:"authorization_secret_name" yaml:"authorizationSecretName,omitempty"`

				FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ProxyUrl *string `tfsdk:"proxy_url" yaml:"proxyUrl,omitempty"`

				RelabelConfigs *[]struct {
					Action *string `tfsdk:"action" yaml:"action,omitempty"`

					Modulus *int64 `tfsdk:"modulus" yaml:"modulus,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

					Replacement *string `tfsdk:"replacement" yaml:"replacement,omitempty"`

					Separator *string `tfsdk:"separator" yaml:"separator,omitempty"`

					SourceLabels *[]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

					TargetLabel *string `tfsdk:"target_label" yaml:"targetLabel,omitempty"`
				} `tfsdk:"relabel_configs" yaml:"relabelConfigs,omitempty"`

				Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"client" yaml:"client,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Queue *struct {
				BatchSendDeadline *string `tfsdk:"batch_send_deadline" yaml:"batchSendDeadline,omitempty"`

				Capacity *int64 `tfsdk:"capacity" yaml:"capacity,omitempty"`

				MaxBackOffPeriod *string `tfsdk:"max_back_off_period" yaml:"maxBackOffPeriod,omitempty"`

				MaxSamplesPerSend *int64 `tfsdk:"max_samples_per_send" yaml:"maxSamplesPerSend,omitempty"`

				MaxShards *int64 `tfsdk:"max_shards" yaml:"maxShards,omitempty"`

				MinBackOffPeriod *string `tfsdk:"min_back_off_period" yaml:"minBackOffPeriod,omitempty"`

				MinShards *int64 `tfsdk:"min_shards" yaml:"minShards,omitempty"`
			} `tfsdk:"queue" yaml:"queue,omitempty"`

			RefreshPeriod *string `tfsdk:"refresh_period" yaml:"refreshPeriod,omitempty"`
		} `tfsdk:"remote_write" yaml:"remoteWrite,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLokiGrafanaComRulerConfigV1Beta1Resource() resource.Resource {
	return &LokiGrafanaComRulerConfigV1Beta1Resource{}
}

func (r *LokiGrafanaComRulerConfigV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_loki_grafana_com_ruler_config_v1beta1"
}

func (r *LokiGrafanaComRulerConfigV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "RulerConfig is the Schema for the rulerconfigs API",
		MarkdownDescription: "RulerConfig is the Schema for the rulerconfigs API",
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
				Description:         "RulerConfigSpec defines the desired state of Ruler",
				MarkdownDescription: "RulerConfigSpec defines the desired state of Ruler",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"alertmanager": {
						Description:         "Defines alert manager configuration to notify on firing alerts.",
						MarkdownDescription: "Defines alert manager configuration to notify on firing alerts.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"discovery": {
								Description:         "Defines the configuration for DNS-based discovery of AlertManager hosts.",
								MarkdownDescription: "Defines the configuration for DNS-based discovery of AlertManager hosts.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enable_srv": {
										Description:         "Use DNS SRV records to discover Alertmanager hosts.",
										MarkdownDescription: "Use DNS SRV records to discover Alertmanager hosts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"refresh_interval": {
										Description:         "How long to wait between refreshing DNS resolutions of Alertmanager hosts.",
										MarkdownDescription: "How long to wait between refreshing DNS resolutions of Alertmanager hosts.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_v2": {
								Description:         "If enabled, then requests to Alertmanager use the v2 API.",
								MarkdownDescription: "If enabled, then requests to Alertmanager use the v2 API.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"endpoints": {
								Description:         "List of AlertManager URLs to send notifications to. Each Alertmanager URL is treated as a separate group in the configuration. Multiple Alertmanagers in HA per group can be supported by using DNS resolution (See EnableDNSDiscovery).",
								MarkdownDescription: "List of AlertManager URLs to send notifications to. Each Alertmanager URL is treated as a separate group in the configuration. Multiple Alertmanagers in HA per group can be supported by using DNS resolution (See EnableDNSDiscovery).",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"external_labels": {
								Description:         "Additional labels to add to all alerts.",
								MarkdownDescription: "Additional labels to add to all alerts.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_url": {
								Description:         "URL for alerts return path.",
								MarkdownDescription: "URL for alerts return path.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"notification_queue": {
								Description:         "Defines the configuration for the notification queue to AlertManager hosts.",
								MarkdownDescription: "Defines the configuration for the notification queue to AlertManager hosts.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"capacity": {
										Description:         "Capacity of the queue for notifications to be sent to the Alertmanager.",
										MarkdownDescription: "Capacity of the queue for notifications to be sent to the Alertmanager.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"for_grace_period": {
										Description:         "Minimum duration between alert and restored 'for' state. This is maintained only for alerts with configured 'for' time greater than the grace period.",
										MarkdownDescription: "Minimum duration between alert and restored 'for' state. This is maintained only for alerts with configured 'for' time greater than the grace period.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"for_outage_tolerance": {
										Description:         "Max time to tolerate outage for restoring 'for' state of alert.",
										MarkdownDescription: "Max time to tolerate outage for restoring 'for' state of alert.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"resend_delay": {
										Description:         "Minimum amount of time to wait before resending an alert to Alertmanager.",
										MarkdownDescription: "Minimum amount of time to wait before resending an alert to Alertmanager.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"timeout": {
										Description:         "HTTP timeout duration when sending notifications to the Alertmanager.",
										MarkdownDescription: "HTTP timeout duration when sending notifications to the Alertmanager.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"relabel_configs": {
								Description:         "List of alert relabel configurations.",
								MarkdownDescription: "List of alert relabel configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action to perform based on regex matching. Default is 'replace'",
										MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("drop", "hashmod", "keep", "labeldrop", "labelkeep", "labelmap", "replace"),
										},
									},

									"modulus": {
										Description:         "Modulus to take of the hash of the source label values.",
										MarkdownDescription: "Modulus to take of the hash of the source label values.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"regex": {
										Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
										MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replacement": {
										Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
										MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"separator": {
										Description:         "Separator placed between concatenated source label values. default is ';'.",
										MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_labels": {
										Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
										MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"target_label": {
										Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
										MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"evaluation_interval": {
						Description:         "Interval on how frequently to evaluate rules.",
						MarkdownDescription: "Interval on how frequently to evaluate rules.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
						},
					},

					"poll_interval": {
						Description:         "Interval on how frequently to poll for new rule definitions.",
						MarkdownDescription: "Interval on how frequently to poll for new rule definitions.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
						},
					},

					"remote_write": {
						Description:         "Defines a remote write endpoint to write recording rule metrics.",
						MarkdownDescription: "Defines a remote write endpoint to write recording rule metrics.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"client": {
								Description:         "Defines the configuration for remote write client.",
								MarkdownDescription: "Defines the configuration for remote write client.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_headers": {
										Description:         "Additional HTTP headers to be sent along with each remote write request.",
										MarkdownDescription: "Additional HTTP headers to be sent along with each remote write request.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorization": {
										Description:         "Type of authorzation to use to access the remote write endpoint",
										MarkdownDescription: "Type of authorzation to use to access the remote write endpoint",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("basic", "header"),
										},
									},

									"authorization_secret_name": {
										Description:         "Name of a secret in the namespace configured for authorization secrets.",
										MarkdownDescription: "Name of a secret in the namespace configured for authorization secrets.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"follow_redirects": {
										Description:         "Configure whether HTTP requests follow HTTP 3xx redirects.",
										MarkdownDescription: "Configure whether HTTP requests follow HTTP 3xx redirects.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the remote write config, which if specified must be unique among remote write configs.",
										MarkdownDescription: "Name of the remote write config, which if specified must be unique among remote write configs.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"proxy_url": {
										Description:         "Optional proxy URL.",
										MarkdownDescription: "Optional proxy URL.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"relabel_configs": {
										Description:         "List of remote write relabel configurations.",
										MarkdownDescription: "List of remote write relabel configurations.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"action": {
												Description:         "Action to perform based on regex matching. Default is 'replace'",
												MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("drop", "hashmod", "keep", "labeldrop", "labelkeep", "labelmap", "replace"),
												},
											},

											"modulus": {
												Description:         "Modulus to take of the hash of the source label values.",
												MarkdownDescription: "Modulus to take of the hash of the source label values.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
												MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"replacement": {
												Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
												MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"separator": {
												Description:         "Separator placed between concatenated source label values. default is ';'.",
												MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"source_labels": {
												Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
												MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"target_label": {
												Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
												MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",

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

									"timeout": {
										Description:         "Timeout for requests to the remote write endpoint.",
										MarkdownDescription: "Timeout for requests to the remote write endpoint.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"url": {
										Description:         "The URL of the endpoint to send samples to.",
										MarkdownDescription: "The URL of the endpoint to send samples to.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "Enable remote-write functionality.",
								MarkdownDescription: "Enable remote-write functionality.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"queue": {
								Description:         "Defines the configuration for remote write client queue.",
								MarkdownDescription: "Defines the configuration for remote write client queue.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"batch_send_deadline": {
										Description:         "Maximum time a sample will wait in buffer.",
										MarkdownDescription: "Maximum time a sample will wait in buffer.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"capacity": {
										Description:         "Number of samples to buffer per shard before we block reading of more",
										MarkdownDescription: "Number of samples to buffer per shard before we block reading of more",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_back_off_period": {
										Description:         "Maximum retry delay.",
										MarkdownDescription: "Maximum retry delay.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"max_samples_per_send": {
										Description:         "Maximum number of samples per send.",
										MarkdownDescription: "Maximum number of samples per send.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_shards": {
										Description:         "Maximum number of shards, i.e. amount of concurrency.",
										MarkdownDescription: "Maximum number of shards, i.e. amount of concurrency.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_back_off_period": {
										Description:         "Initial retry delay. Gets doubled for every retry.",
										MarkdownDescription: "Initial retry delay. Gets doubled for every retry.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"min_shards": {
										Description:         "Minimum number of shards, i.e. amount of concurrency.",
										MarkdownDescription: "Minimum number of shards, i.e. amount of concurrency.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"refresh_period": {
								Description:         "Minimum period to wait between refreshing remote-write reconfigurations.",
								MarkdownDescription: "Minimum period to wait between refreshing remote-write reconfigurations.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
								},
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

func (r *LokiGrafanaComRulerConfigV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_loki_grafana_com_ruler_config_v1beta1")

	var state LokiGrafanaComRulerConfigV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LokiGrafanaComRulerConfigV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("loki.grafana.com/v1beta1")
	goModel.Kind = utilities.Ptr("RulerConfig")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LokiGrafanaComRulerConfigV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_loki_grafana_com_ruler_config_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *LokiGrafanaComRulerConfigV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_loki_grafana_com_ruler_config_v1beta1")

	var state LokiGrafanaComRulerConfigV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LokiGrafanaComRulerConfigV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("loki.grafana.com/v1beta1")
	goModel.Kind = utilities.Ptr("RulerConfig")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *LokiGrafanaComRulerConfigV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_loki_grafana_com_ruler_config_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
