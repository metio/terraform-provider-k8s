/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package loki_grafana_com_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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
	"regexp"
	"strings"
)

var (
	_ resource.Resource                = &LokiGrafanaComRulerConfigV1Resource{}
	_ resource.ResourceWithConfigure   = &LokiGrafanaComRulerConfigV1Resource{}
	_ resource.ResourceWithImportState = &LokiGrafanaComRulerConfigV1Resource{}
)

func NewLokiGrafanaComRulerConfigV1Resource() resource.Resource {
	return &LokiGrafanaComRulerConfigV1Resource{}
}

type LokiGrafanaComRulerConfigV1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type LokiGrafanaComRulerConfigV1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Alertmanager *struct {
			Client *struct {
				BasicAuth *struct {
					Password *string `tfsdk:"password" json:"password,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
				HeaderAuth *struct {
					Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
					CredentialsFile *string `tfsdk:"credentials_file" json:"credentialsFile,omitempty"`
					Type            *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"header_auth" json:"headerAuth,omitempty"`
				Tls *struct {
					CaPath     *string `tfsdk:"ca_path" json:"caPath,omitempty"`
					CertPath   *string `tfsdk:"cert_path" json:"certPath,omitempty"`
					KeyPath    *string `tfsdk:"key_path" json:"keyPath,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"client" json:"client,omitempty"`
			Discovery *struct {
				EnableSRV       *bool   `tfsdk:"enable_srv" json:"enableSRV,omitempty"`
				RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
			} `tfsdk:"discovery" json:"discovery,omitempty"`
			EnableV2          *bool              `tfsdk:"enable_v2" json:"enableV2,omitempty"`
			Endpoints         *[]string          `tfsdk:"endpoints" json:"endpoints,omitempty"`
			ExternalLabels    *map[string]string `tfsdk:"external_labels" json:"externalLabels,omitempty"`
			ExternalUrl       *string            `tfsdk:"external_url" json:"externalUrl,omitempty"`
			NotificationQueue *struct {
				Capacity           *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
				ForGracePeriod     *string `tfsdk:"for_grace_period" json:"forGracePeriod,omitempty"`
				ForOutageTolerance *string `tfsdk:"for_outage_tolerance" json:"forOutageTolerance,omitempty"`
				ResendDelay        *string `tfsdk:"resend_delay" json:"resendDelay,omitempty"`
				Timeout            *string `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"notification_queue" json:"notificationQueue,omitempty"`
			RelabelConfigs *[]struct {
				Action       *string   `tfsdk:"action" json:"action,omitempty"`
				Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
				Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
				Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
				Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
				SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
			} `tfsdk:"relabel_configs" json:"relabelConfigs,omitempty"`
		} `tfsdk:"alertmanager" json:"alertmanager,omitempty"`
		EvaluationInterval *string `tfsdk:"evaluation_interval" json:"evaluationInterval,omitempty"`
		Overrides          *struct {
			Alertmanager *struct {
				Client *struct {
					BasicAuth *struct {
						Password *string `tfsdk:"password" json:"password,omitempty"`
						Username *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
					HeaderAuth *struct {
						Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
						CredentialsFile *string `tfsdk:"credentials_file" json:"credentialsFile,omitempty"`
						Type            *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"header_auth" json:"headerAuth,omitempty"`
					Tls *struct {
						CaPath     *string `tfsdk:"ca_path" json:"caPath,omitempty"`
						CertPath   *string `tfsdk:"cert_path" json:"certPath,omitempty"`
						KeyPath    *string `tfsdk:"key_path" json:"keyPath,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"client" json:"client,omitempty"`
				Discovery *struct {
					EnableSRV       *bool   `tfsdk:"enable_srv" json:"enableSRV,omitempty"`
					RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
				} `tfsdk:"discovery" json:"discovery,omitempty"`
				EnableV2          *bool              `tfsdk:"enable_v2" json:"enableV2,omitempty"`
				Endpoints         *[]string          `tfsdk:"endpoints" json:"endpoints,omitempty"`
				ExternalLabels    *map[string]string `tfsdk:"external_labels" json:"externalLabels,omitempty"`
				ExternalUrl       *string            `tfsdk:"external_url" json:"externalUrl,omitempty"`
				NotificationQueue *struct {
					Capacity           *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
					ForGracePeriod     *string `tfsdk:"for_grace_period" json:"forGracePeriod,omitempty"`
					ForOutageTolerance *string `tfsdk:"for_outage_tolerance" json:"forOutageTolerance,omitempty"`
					ResendDelay        *string `tfsdk:"resend_delay" json:"resendDelay,omitempty"`
					Timeout            *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"notification_queue" json:"notificationQueue,omitempty"`
				RelabelConfigs *[]struct {
					Action       *string   `tfsdk:"action" json:"action,omitempty"`
					Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
					Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
					Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
					Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
					SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
					TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
				} `tfsdk:"relabel_configs" json:"relabelConfigs,omitempty"`
			} `tfsdk:"alertmanager" json:"alertmanager,omitempty"`
		} `tfsdk:"overrides" json:"overrides,omitempty"`
		PollInterval *string `tfsdk:"poll_interval" json:"pollInterval,omitempty"`
		RemoteWrite  *struct {
			Client *struct {
				AdditionalHeaders       *map[string]string `tfsdk:"additional_headers" json:"additionalHeaders,omitempty"`
				Authorization           *string            `tfsdk:"authorization" json:"authorization,omitempty"`
				AuthorizationSecretName *string            `tfsdk:"authorization_secret_name" json:"authorizationSecretName,omitempty"`
				FollowRedirects         *bool              `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
				Name                    *string            `tfsdk:"name" json:"name,omitempty"`
				ProxyUrl                *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
				RelabelConfigs          *[]struct {
					Action       *string   `tfsdk:"action" json:"action,omitempty"`
					Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
					Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
					Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
					Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
					SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
					TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
				} `tfsdk:"relabel_configs" json:"relabelConfigs,omitempty"`
				Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
				Url     *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"client" json:"client,omitempty"`
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Queue   *struct {
				BatchSendDeadline *string `tfsdk:"batch_send_deadline" json:"batchSendDeadline,omitempty"`
				Capacity          *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
				MaxBackOffPeriod  *string `tfsdk:"max_back_off_period" json:"maxBackOffPeriod,omitempty"`
				MaxSamplesPerSend *int64  `tfsdk:"max_samples_per_send" json:"maxSamplesPerSend,omitempty"`
				MaxShards         *int64  `tfsdk:"max_shards" json:"maxShards,omitempty"`
				MinBackOffPeriod  *string `tfsdk:"min_back_off_period" json:"minBackOffPeriod,omitempty"`
				MinShards         *int64  `tfsdk:"min_shards" json:"minShards,omitempty"`
			} `tfsdk:"queue" json:"queue,omitempty"`
			RefreshPeriod *string `tfsdk:"refresh_period" json:"refreshPeriod,omitempty"`
		} `tfsdk:"remote_write" json:"remoteWrite,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LokiGrafanaComRulerConfigV1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_loki_grafana_com_ruler_config_v1"
}

func (r *LokiGrafanaComRulerConfigV1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RulerConfig is the Schema for the rulerconfigs API",
		MarkdownDescription: "RulerConfig is the Schema for the rulerconfigs API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
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
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
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
				Description:         "RulerConfigSpec defines the desired state of Ruler",
				MarkdownDescription: "RulerConfigSpec defines the desired state of Ruler",
				Attributes: map[string]schema.Attribute{
					"alertmanager": schema.SingleNestedAttribute{
						Description:         "Defines alert manager configuration to notify on firing alerts.",
						MarkdownDescription: "Defines alert manager configuration to notify on firing alerts.",
						Attributes: map[string]schema.Attribute{
							"client": schema.SingleNestedAttribute{
								Description:         "Client configuration for reaching the alertmanager endpoint.",
								MarkdownDescription: "Client configuration for reaching the alertmanager endpoint.",
								Attributes: map[string]schema.Attribute{
									"basic_auth": schema.SingleNestedAttribute{
										Description:         "Basic authentication configuration for reaching the alertmanager endpoints.",
										MarkdownDescription: "Basic authentication configuration for reaching the alertmanager endpoints.",
										Attributes: map[string]schema.Attribute{
											"password": schema.StringAttribute{
												Description:         "The subject's password for the basic authentication configuration.",
												MarkdownDescription: "The subject's password for the basic authentication configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"username": schema.StringAttribute{
												Description:         "The subject's username for the basic authentication configuration.",
												MarkdownDescription: "The subject's username for the basic authentication configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"header_auth": schema.SingleNestedAttribute{
										Description:         "Header authentication configuration for reaching the alertmanager endpoints.",
										MarkdownDescription: "Header authentication configuration for reaching the alertmanager endpoints.",
										Attributes: map[string]schema.Attribute{
											"credentials": schema.StringAttribute{
												Description:         "The credentials for the header authentication configuration.",
												MarkdownDescription: "The credentials for the header authentication configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"credentials_file": schema.StringAttribute{
												Description:         "The credentials file for the Header authentication configuration. It is mutually exclusive with 'credentials'.",
												MarkdownDescription: "The credentials file for the Header authentication configuration. It is mutually exclusive with 'credentials'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "The authentication type for the header authentication configuration.",
												MarkdownDescription: "The authentication type for the header authentication configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS configuration for reaching the alertmanager endpoints.",
										MarkdownDescription: "TLS configuration for reaching the alertmanager endpoints.",
										Attributes: map[string]schema.Attribute{
											"ca_path": schema.StringAttribute{
												Description:         "The CA certificate file path for the TLS configuration.",
												MarkdownDescription: "The CA certificate file path for the TLS configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_path": schema.StringAttribute{
												Description:         "The client-side certificate file path for the TLS configuration.",
												MarkdownDescription: "The client-side certificate file path for the TLS configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_path": schema.StringAttribute{
												Description:         "The client-side key file path for the TLS configuration.",
												MarkdownDescription: "The client-side key file path for the TLS configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"server_name": schema.StringAttribute{
												Description:         "The server name to validate in the alertmanager server certificates.",
												MarkdownDescription: "The server name to validate in the alertmanager server certificates.",
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

							"discovery": schema.SingleNestedAttribute{
								Description:         "Defines the configuration for DNS-based discovery of AlertManager hosts.",
								MarkdownDescription: "Defines the configuration for DNS-based discovery of AlertManager hosts.",
								Attributes: map[string]schema.Attribute{
									"enable_srv": schema.BoolAttribute{
										Description:         "Use DNS SRV records to discover Alertmanager hosts.",
										MarkdownDescription: "Use DNS SRV records to discover Alertmanager hosts.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"refresh_interval": schema.StringAttribute{
										Description:         "How long to wait between refreshing DNS resolutions of Alertmanager hosts.",
										MarkdownDescription: "How long to wait between refreshing DNS resolutions of Alertmanager hosts.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_v2": schema.BoolAttribute{
								Description:         "If enabled, then requests to Alertmanager use the v2 API.",
								MarkdownDescription: "If enabled, then requests to Alertmanager use the v2 API.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"endpoints": schema.ListAttribute{
								Description:         "List of AlertManager URLs to send notifications to. Each Alertmanager URL is treated as a separate group in the configuration. Multiple Alertmanagers in HA per group can be supported by using DNS resolution (See EnableDNSDiscovery).",
								MarkdownDescription: "List of AlertManager URLs to send notifications to. Each Alertmanager URL is treated as a separate group in the configuration. Multiple Alertmanagers in HA per group can be supported by using DNS resolution (See EnableDNSDiscovery).",
								ElementType:         types.StringType,
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"external_labels": schema.MapAttribute{
								Description:         "Additional labels to add to all alerts.",
								MarkdownDescription: "Additional labels to add to all alerts.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_url": schema.StringAttribute{
								Description:         "URL for alerts return path.",
								MarkdownDescription: "URL for alerts return path.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"notification_queue": schema.SingleNestedAttribute{
								Description:         "Defines the configuration for the notification queue to AlertManager hosts.",
								MarkdownDescription: "Defines the configuration for the notification queue to AlertManager hosts.",
								Attributes: map[string]schema.Attribute{
									"capacity": schema.Int64Attribute{
										Description:         "Capacity of the queue for notifications to be sent to the Alertmanager.",
										MarkdownDescription: "Capacity of the queue for notifications to be sent to the Alertmanager.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"for_grace_period": schema.StringAttribute{
										Description:         "Minimum duration between alert and restored 'for' state. This is maintained only for alerts with configured 'for' time greater than the grace period.",
										MarkdownDescription: "Minimum duration between alert and restored 'for' state. This is maintained only for alerts with configured 'for' time greater than the grace period.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"for_outage_tolerance": schema.StringAttribute{
										Description:         "Max time to tolerate outage for restoring 'for' state of alert.",
										MarkdownDescription: "Max time to tolerate outage for restoring 'for' state of alert.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"resend_delay": schema.StringAttribute{
										Description:         "Minimum amount of time to wait before resending an alert to Alertmanager.",
										MarkdownDescription: "Minimum amount of time to wait before resending an alert to Alertmanager.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"timeout": schema.StringAttribute{
										Description:         "HTTP timeout duration when sending notifications to the Alertmanager.",
										MarkdownDescription: "HTTP timeout duration when sending notifications to the Alertmanager.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"relabel_configs": schema.ListNestedAttribute{
								Description:         "List of alert relabel configurations.",
								MarkdownDescription: "List of alert relabel configurations.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action to perform based on regex matching. Default is 'replace'",
											MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("drop", "hashmod", "keep", "labeldrop", "labelkeep", "labelmap", "replace"),
											},
										},

										"modulus": schema.Int64Attribute{
											Description:         "Modulus to take of the hash of the source label values.",
											MarkdownDescription: "Modulus to take of the hash of the source label values.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex": schema.StringAttribute{
											Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
											MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replacement": schema.StringAttribute{
											Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
											MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"separator": schema.StringAttribute{
											Description:         "Separator placed between concatenated source label values. default is ';'.",
											MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source_labels": schema.ListAttribute{
											Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
											MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target_label": schema.StringAttribute{
											Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
											MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
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

					"evaluation_interval": schema.StringAttribute{
						Description:         "Interval on how frequently to evaluate rules.",
						MarkdownDescription: "Interval on how frequently to evaluate rules.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
						},
					},

					"overrides": schema.SingleNestedAttribute{
						Description:         "Overrides defines the config overrides to be applied per-tenant.",
						MarkdownDescription: "Overrides defines the config overrides to be applied per-tenant.",
						Attributes: map[string]schema.Attribute{
							"alertmanager": schema.SingleNestedAttribute{
								Description:         "AlertManagerOverrides defines the overrides to apply to the alertmanager config.",
								MarkdownDescription: "AlertManagerOverrides defines the overrides to apply to the alertmanager config.",
								Attributes: map[string]schema.Attribute{
									"client": schema.SingleNestedAttribute{
										Description:         "Client configuration for reaching the alertmanager endpoint.",
										MarkdownDescription: "Client configuration for reaching the alertmanager endpoint.",
										Attributes: map[string]schema.Attribute{
											"basic_auth": schema.SingleNestedAttribute{
												Description:         "Basic authentication configuration for reaching the alertmanager endpoints.",
												MarkdownDescription: "Basic authentication configuration for reaching the alertmanager endpoints.",
												Attributes: map[string]schema.Attribute{
													"password": schema.StringAttribute{
														Description:         "The subject's password for the basic authentication configuration.",
														MarkdownDescription: "The subject's password for the basic authentication configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"username": schema.StringAttribute{
														Description:         "The subject's username for the basic authentication configuration.",
														MarkdownDescription: "The subject's username for the basic authentication configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"header_auth": schema.SingleNestedAttribute{
												Description:         "Header authentication configuration for reaching the alertmanager endpoints.",
												MarkdownDescription: "Header authentication configuration for reaching the alertmanager endpoints.",
												Attributes: map[string]schema.Attribute{
													"credentials": schema.StringAttribute{
														Description:         "The credentials for the header authentication configuration.",
														MarkdownDescription: "The credentials for the header authentication configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"credentials_file": schema.StringAttribute{
														Description:         "The credentials file for the Header authentication configuration. It is mutually exclusive with 'credentials'.",
														MarkdownDescription: "The credentials file for the Header authentication configuration. It is mutually exclusive with 'credentials'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "The authentication type for the header authentication configuration.",
														MarkdownDescription: "The authentication type for the header authentication configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls": schema.SingleNestedAttribute{
												Description:         "TLS configuration for reaching the alertmanager endpoints.",
												MarkdownDescription: "TLS configuration for reaching the alertmanager endpoints.",
												Attributes: map[string]schema.Attribute{
													"ca_path": schema.StringAttribute{
														Description:         "The CA certificate file path for the TLS configuration.",
														MarkdownDescription: "The CA certificate file path for the TLS configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_path": schema.StringAttribute{
														Description:         "The client-side certificate file path for the TLS configuration.",
														MarkdownDescription: "The client-side certificate file path for the TLS configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_path": schema.StringAttribute{
														Description:         "The client-side key file path for the TLS configuration.",
														MarkdownDescription: "The client-side key file path for the TLS configuration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"server_name": schema.StringAttribute{
														Description:         "The server name to validate in the alertmanager server certificates.",
														MarkdownDescription: "The server name to validate in the alertmanager server certificates.",
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

									"discovery": schema.SingleNestedAttribute{
										Description:         "Defines the configuration for DNS-based discovery of AlertManager hosts.",
										MarkdownDescription: "Defines the configuration for DNS-based discovery of AlertManager hosts.",
										Attributes: map[string]schema.Attribute{
											"enable_srv": schema.BoolAttribute{
												Description:         "Use DNS SRV records to discover Alertmanager hosts.",
												MarkdownDescription: "Use DNS SRV records to discover Alertmanager hosts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"refresh_interval": schema.StringAttribute{
												Description:         "How long to wait between refreshing DNS resolutions of Alertmanager hosts.",
												MarkdownDescription: "How long to wait between refreshing DNS resolutions of Alertmanager hosts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable_v2": schema.BoolAttribute{
										Description:         "If enabled, then requests to Alertmanager use the v2 API.",
										MarkdownDescription: "If enabled, then requests to Alertmanager use the v2 API.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"endpoints": schema.ListAttribute{
										Description:         "List of AlertManager URLs to send notifications to. Each Alertmanager URL is treated as a separate group in the configuration. Multiple Alertmanagers in HA per group can be supported by using DNS resolution (See EnableDNSDiscovery).",
										MarkdownDescription: "List of AlertManager URLs to send notifications to. Each Alertmanager URL is treated as a separate group in the configuration. Multiple Alertmanagers in HA per group can be supported by using DNS resolution (See EnableDNSDiscovery).",
										ElementType:         types.StringType,
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"external_labels": schema.MapAttribute{
										Description:         "Additional labels to add to all alerts.",
										MarkdownDescription: "Additional labels to add to all alerts.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_url": schema.StringAttribute{
										Description:         "URL for alerts return path.",
										MarkdownDescription: "URL for alerts return path.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"notification_queue": schema.SingleNestedAttribute{
										Description:         "Defines the configuration for the notification queue to AlertManager hosts.",
										MarkdownDescription: "Defines the configuration for the notification queue to AlertManager hosts.",
										Attributes: map[string]schema.Attribute{
											"capacity": schema.Int64Attribute{
												Description:         "Capacity of the queue for notifications to be sent to the Alertmanager.",
												MarkdownDescription: "Capacity of the queue for notifications to be sent to the Alertmanager.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"for_grace_period": schema.StringAttribute{
												Description:         "Minimum duration between alert and restored 'for' state. This is maintained only for alerts with configured 'for' time greater than the grace period.",
												MarkdownDescription: "Minimum duration between alert and restored 'for' state. This is maintained only for alerts with configured 'for' time greater than the grace period.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
												},
											},

											"for_outage_tolerance": schema.StringAttribute{
												Description:         "Max time to tolerate outage for restoring 'for' state of alert.",
												MarkdownDescription: "Max time to tolerate outage for restoring 'for' state of alert.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
												},
											},

											"resend_delay": schema.StringAttribute{
												Description:         "Minimum amount of time to wait before resending an alert to Alertmanager.",
												MarkdownDescription: "Minimum amount of time to wait before resending an alert to Alertmanager.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
												},
											},

											"timeout": schema.StringAttribute{
												Description:         "HTTP timeout duration when sending notifications to the Alertmanager.",
												MarkdownDescription: "HTTP timeout duration when sending notifications to the Alertmanager.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"relabel_configs": schema.ListNestedAttribute{
										Description:         "List of alert relabel configurations.",
										MarkdownDescription: "List of alert relabel configurations.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action to perform based on regex matching. Default is 'replace'",
													MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("drop", "hashmod", "keep", "labeldrop", "labelkeep", "labelmap", "replace"),
													},
												},

												"modulus": schema.Int64Attribute{
													Description:         "Modulus to take of the hash of the source label values.",
													MarkdownDescription: "Modulus to take of the hash of the source label values.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"regex": schema.StringAttribute{
													Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
													MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replacement": schema.StringAttribute{
													Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
													MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"separator": schema.StringAttribute{
													Description:         "Separator placed between concatenated source label values. default is ';'.",
													MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_labels": schema.ListAttribute{
													Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
													MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"target_label": schema.StringAttribute{
													Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
													MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"poll_interval": schema.StringAttribute{
						Description:         "Interval on how frequently to poll for new rule definitions.",
						MarkdownDescription: "Interval on how frequently to poll for new rule definitions.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
						},
					},

					"remote_write": schema.SingleNestedAttribute{
						Description:         "Defines a remote write endpoint to write recording rule metrics.",
						MarkdownDescription: "Defines a remote write endpoint to write recording rule metrics.",
						Attributes: map[string]schema.Attribute{
							"client": schema.SingleNestedAttribute{
								Description:         "Defines the configuration for remote write client.",
								MarkdownDescription: "Defines the configuration for remote write client.",
								Attributes: map[string]schema.Attribute{
									"additional_headers": schema.MapAttribute{
										Description:         "Additional HTTP headers to be sent along with each remote write request.",
										MarkdownDescription: "Additional HTTP headers to be sent along with each remote write request.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"authorization": schema.StringAttribute{
										Description:         "Type of authorzation to use to access the remote write endpoint",
										MarkdownDescription: "Type of authorzation to use to access the remote write endpoint",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("basic", "header"),
										},
									},

									"authorization_secret_name": schema.StringAttribute{
										Description:         "Name of a secret in the namespace configured for authorization secrets.",
										MarkdownDescription: "Name of a secret in the namespace configured for authorization secrets.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"follow_redirects": schema.BoolAttribute{
										Description:         "Configure whether HTTP requests follow HTTP 3xx redirects.",
										MarkdownDescription: "Configure whether HTTP requests follow HTTP 3xx redirects.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the remote write config, which if specified must be unique among remote write configs.",
										MarkdownDescription: "Name of the remote write config, which if specified must be unique among remote write configs.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"proxy_url": schema.StringAttribute{
										Description:         "Optional proxy URL.",
										MarkdownDescription: "Optional proxy URL.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"relabel_configs": schema.ListNestedAttribute{
										Description:         "List of remote write relabel configurations.",
										MarkdownDescription: "List of remote write relabel configurations.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"action": schema.StringAttribute{
													Description:         "Action to perform based on regex matching. Default is 'replace'",
													MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("drop", "hashmod", "keep", "labeldrop", "labelkeep", "labelmap", "replace"),
													},
												},

												"modulus": schema.Int64Attribute{
													Description:         "Modulus to take of the hash of the source label values.",
													MarkdownDescription: "Modulus to take of the hash of the source label values.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"regex": schema.StringAttribute{
													Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
													MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"replacement": schema.StringAttribute{
													Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
													MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"separator": schema.StringAttribute{
													Description:         "Separator placed between concatenated source label values. default is ';'.",
													MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"source_labels": schema.ListAttribute{
													Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
													MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"target_label": schema.StringAttribute{
													Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
													MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
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

									"timeout": schema.StringAttribute{
										Description:         "Timeout for requests to the remote write endpoint.",
										MarkdownDescription: "Timeout for requests to the remote write endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"url": schema.StringAttribute{
										Description:         "The URL of the endpoint to send samples to.",
										MarkdownDescription: "The URL of the endpoint to send samples to.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enable remote-write functionality.",
								MarkdownDescription: "Enable remote-write functionality.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"queue": schema.SingleNestedAttribute{
								Description:         "Defines the configuration for remote write client queue.",
								MarkdownDescription: "Defines the configuration for remote write client queue.",
								Attributes: map[string]schema.Attribute{
									"batch_send_deadline": schema.StringAttribute{
										Description:         "Maximum time a sample will wait in buffer.",
										MarkdownDescription: "Maximum time a sample will wait in buffer.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"capacity": schema.Int64Attribute{
										Description:         "Number of samples to buffer per shard before we block reading of more",
										MarkdownDescription: "Number of samples to buffer per shard before we block reading of more",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_back_off_period": schema.StringAttribute{
										Description:         "Maximum retry delay.",
										MarkdownDescription: "Maximum retry delay.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"max_samples_per_send": schema.Int64Attribute{
										Description:         "Maximum number of samples per send.",
										MarkdownDescription: "Maximum number of samples per send.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_shards": schema.Int64Attribute{
										Description:         "Maximum number of shards, i.e. amount of concurrency.",
										MarkdownDescription: "Maximum number of shards, i.e. amount of concurrency.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_back_off_period": schema.StringAttribute{
										Description:         "Initial retry delay. Gets doubled for every retry.",
										MarkdownDescription: "Initial retry delay. Gets doubled for every retry.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
										},
									},

									"min_shards": schema.Int64Attribute{
										Description:         "Minimum number of shards, i.e. amount of concurrency.",
										MarkdownDescription: "Minimum number of shards, i.e. amount of concurrency.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"refresh_period": schema.StringAttribute{
								Description:         "Minimum period to wait between refreshing remote-write reconfigurations.",
								MarkdownDescription: "Minimum period to wait between refreshing remote-write reconfigurations.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)`), ""),
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

func (r *LokiGrafanaComRulerConfigV1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *LokiGrafanaComRulerConfigV1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_loki_grafana_com_ruler_config_v1")

	var model LokiGrafanaComRulerConfigV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("loki.grafana.com/v1")
	model.Kind = pointer.String("RulerConfig")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "loki.grafana.com", Version: "v1", Resource: "rulerconfigs"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LokiGrafanaComRulerConfigV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *LokiGrafanaComRulerConfigV1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_loki_grafana_com_ruler_config_v1")

	var data LokiGrafanaComRulerConfigV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "loki.grafana.com", Version: "v1", Resource: "rulerconfigs"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LokiGrafanaComRulerConfigV1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *LokiGrafanaComRulerConfigV1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_loki_grafana_com_ruler_config_v1")

	var model LokiGrafanaComRulerConfigV1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("loki.grafana.com/v1")
	model.Kind = pointer.String("RulerConfig")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
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
		Resource(k8sSchema.GroupVersionResource{Group: "loki.grafana.com", Version: "v1", Resource: "rulerconfigs"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LokiGrafanaComRulerConfigV1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *LokiGrafanaComRulerConfigV1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_loki_grafana_com_ruler_config_v1")

	var data LokiGrafanaComRulerConfigV1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "loki.grafana.com", Version: "v1", Resource: "rulerconfigs"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *LokiGrafanaComRulerConfigV1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
