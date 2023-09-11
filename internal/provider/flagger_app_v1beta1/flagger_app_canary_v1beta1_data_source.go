/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package flagger_app_v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &FlaggerAppCanaryV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &FlaggerAppCanaryV1Beta1DataSource{}
)

func NewFlaggerAppCanaryV1Beta1DataSource() datasource.DataSource {
	return &FlaggerAppCanaryV1Beta1DataSource{}
}

type FlaggerAppCanaryV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type FlaggerAppCanaryV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Analysis *struct {
			Alerts *[]struct {
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				ProviderRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"provider_ref" json:"providerRef,omitempty"`
				Severity *string `tfsdk:"severity" json:"severity,omitempty"`
			} `tfsdk:"alerts" json:"alerts,omitempty"`
			CanaryReadyThreshold *float64 `tfsdk:"canary_ready_threshold" json:"canaryReadyThreshold,omitempty"`
			Interval             *string  `tfsdk:"interval" json:"interval,omitempty"`
			Iterations           *float64 `tfsdk:"iterations" json:"iterations,omitempty"`
			Match                *[]struct {
				Headers *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
					Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				QueryParams *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"query_params" json:"queryParams,omitempty"`
				SourceLabels *map[string]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			MaxWeight *float64 `tfsdk:"max_weight" json:"maxWeight,omitempty"`
			Metrics   *[]struct {
				Interval    *string `tfsdk:"interval" json:"interval,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Query       *string `tfsdk:"query" json:"query,omitempty"`
				TemplateRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"template_ref" json:"templateRef,omitempty"`
				TemplateVariables *map[string]string `tfsdk:"template_variables" json:"templateVariables,omitempty"`
				Threshold         *float64           `tfsdk:"threshold" json:"threshold,omitempty"`
				ThresholdRange    *struct {
					Max *float64 `tfsdk:"max" json:"max,omitempty"`
					Min *float64 `tfsdk:"min" json:"min,omitempty"`
				} `tfsdk:"threshold_range" json:"thresholdRange,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Mirror                *bool    `tfsdk:"mirror" json:"mirror,omitempty"`
			MirrorWeight          *float64 `tfsdk:"mirror_weight" json:"mirrorWeight,omitempty"`
			PrimaryReadyThreshold *float64 `tfsdk:"primary_ready_threshold" json:"primaryReadyThreshold,omitempty"`
			SessionAffinity       *struct {
				CookieName *string  `tfsdk:"cookie_name" json:"cookieName,omitempty"`
				MaxAge     *float64 `tfsdk:"max_age" json:"maxAge,omitempty"`
			} `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
			StepWeight          *float64  `tfsdk:"step_weight" json:"stepWeight,omitempty"`
			StepWeightPromotion *float64  `tfsdk:"step_weight_promotion" json:"stepWeightPromotion,omitempty"`
			StepWeights         *[]string `tfsdk:"step_weights" json:"stepWeights,omitempty"`
			Threshold           *float64  `tfsdk:"threshold" json:"threshold,omitempty"`
			Webhooks            *[]struct {
				Metadata  *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				MuteAlert *bool              `tfsdk:"mute_alert" json:"muteAlert,omitempty"`
				Name      *string            `tfsdk:"name" json:"name,omitempty"`
				Timeout   *string            `tfsdk:"timeout" json:"timeout,omitempty"`
				Type      *string            `tfsdk:"type" json:"type,omitempty"`
				Url       *string            `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"webhooks" json:"webhooks,omitempty"`
		} `tfsdk:"analysis" json:"analysis,omitempty"`
		AutoscalerRef *struct {
			ApiVersion            *string            `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind                  *string            `tfsdk:"kind" json:"kind,omitempty"`
			Name                  *string            `tfsdk:"name" json:"name,omitempty"`
			PrimaryScalerQueries  *map[string]string `tfsdk:"primary_scaler_queries" json:"primaryScalerQueries,omitempty"`
			PrimaryScalerReplicas *struct {
				MaxReplicas *float64 `tfsdk:"max_replicas" json:"maxReplicas,omitempty"`
				MinReplicas *float64 `tfsdk:"min_replicas" json:"minReplicas,omitempty"`
			} `tfsdk:"primary_scaler_replicas" json:"primaryScalerReplicas,omitempty"`
		} `tfsdk:"autoscaler_ref" json:"autoscalerRef,omitempty"`
		IngressRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"ingress_ref" json:"ingressRef,omitempty"`
		MetricsServer           *string  `tfsdk:"metrics_server" json:"metricsServer,omitempty"`
		ProgressDeadlineSeconds *float64 `tfsdk:"progress_deadline_seconds" json:"progressDeadlineSeconds,omitempty"`
		Provider                *string  `tfsdk:"provider" json:"provider,omitempty"`
		RevertOnDeletion        *bool    `tfsdk:"revert_on_deletion" json:"revertOnDeletion,omitempty"`
		RouteRef                *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"route_ref" json:"routeRef,omitempty"`
		Service *struct {
			Apex *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"apex" json:"apex,omitempty"`
			AppProtocol *string   `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
			Backends    *[]string `tfsdk:"backends" json:"backends,omitempty"`
			Canary      *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"canary" json:"canary,omitempty"`
			CorsPolicy *struct {
				AllowCredentials *bool     `tfsdk:"allow_credentials" json:"allowCredentials,omitempty"`
				AllowHeaders     *[]string `tfsdk:"allow_headers" json:"allowHeaders,omitempty"`
				AllowMethods     *[]string `tfsdk:"allow_methods" json:"allowMethods,omitempty"`
				AllowOrigin      *[]string `tfsdk:"allow_origin" json:"allowOrigin,omitempty"`
				AllowOrigins     *[]struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"allow_origins" json:"allowOrigins,omitempty"`
				ExposeHeaders *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
				MaxAge        *string   `tfsdk:"max_age" json:"maxAge,omitempty"`
			} `tfsdk:"cors_policy" json:"corsPolicy,omitempty"`
			Delegation  *bool `tfsdk:"delegation" json:"delegation,omitempty"`
			GatewayRefs *[]struct {
				Group       *string `tfsdk:"group" json:"group,omitempty"`
				Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Namespace   *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Port        *int64  `tfsdk:"port" json:"port,omitempty"`
				SectionName *string `tfsdk:"section_name" json:"sectionName,omitempty"`
			} `tfsdk:"gateway_refs" json:"gatewayRefs,omitempty"`
			Gateways *[]string `tfsdk:"gateways" json:"gateways,omitempty"`
			Headers  *struct {
				Request *struct {
					Add    *map[string]string `tfsdk:"add" json:"add,omitempty"`
					Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
					Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"request" json:"request,omitempty"`
				Response *struct {
					Add    *map[string]string `tfsdk:"add" json:"add,omitempty"`
					Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
					Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"response" json:"response,omitempty"`
			} `tfsdk:"headers" json:"headers,omitempty"`
			Hosts *[]string `tfsdk:"hosts" json:"hosts,omitempty"`
			Match *[]struct {
				Authority *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"authority" json:"authority,omitempty"`
				Gateways *[]string `tfsdk:"gateways" json:"gateways,omitempty"`
				Headers  *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				IgnoreUriCase *bool `tfsdk:"ignore_uri_case" json:"ignoreUriCase,omitempty"`
				Method        *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"method" json:"method,omitempty"`
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Port        *int64  `tfsdk:"port" json:"port,omitempty"`
				QueryParams *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"query_params" json:"queryParams,omitempty"`
				Scheme *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"scheme" json:"scheme,omitempty"`
				SourceLabels    *map[string]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				SourceNamespace *string            `tfsdk:"source_namespace" json:"sourceNamespace,omitempty"`
				Uri             *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"uri" json:"uri,omitempty"`
				WithoutHeaders *struct {
					Exact  *string `tfsdk:"exact" json:"exact,omitempty"`
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Regex  *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"without_headers" json:"withoutHeaders,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			MeshName      *string  `tfsdk:"mesh_name" json:"meshName,omitempty"`
			Name          *string  `tfsdk:"name" json:"name,omitempty"`
			Port          *float64 `tfsdk:"port" json:"port,omitempty"`
			PortDiscovery *bool    `tfsdk:"port_discovery" json:"portDiscovery,omitempty"`
			PortName      *string  `tfsdk:"port_name" json:"portName,omitempty"`
			Primary       *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"primary" json:"primary,omitempty"`
			Retries *struct {
				Attempts      *int64  `tfsdk:"attempts" json:"attempts,omitempty"`
				PerTryTimeout *string `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
				RetryOn       *string `tfsdk:"retry_on" json:"retryOn,omitempty"`
			} `tfsdk:"retries" json:"retries,omitempty"`
			Rewrite *struct {
				Uri *string `tfsdk:"uri" json:"uri,omitempty"`
			} `tfsdk:"rewrite" json:"rewrite,omitempty"`
			TargetPort    *string `tfsdk:"target_port" json:"targetPort,omitempty"`
			Timeout       *string `tfsdk:"timeout" json:"timeout,omitempty"`
			TrafficPolicy *struct {
				ConnectionPool *struct {
					Http *struct {
						H2UpgradePolicy          *string `tfsdk:"h2_upgrade_policy" json:"h2UpgradePolicy,omitempty"`
						Http1MaxPendingRequests  *int64  `tfsdk:"http1_max_pending_requests" json:"http1MaxPendingRequests,omitempty"`
						Http2MaxRequests         *int64  `tfsdk:"http2_max_requests" json:"http2MaxRequests,omitempty"`
						IdleTimeout              *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
						MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
						MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
				} `tfsdk:"connection_pool" json:"connectionPool,omitempty"`
				LoadBalancer *struct {
					ConsistentHash *struct {
						HttpCookie *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
							Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
						} `tfsdk:"http_cookie" json:"httpCookie,omitempty"`
						HttpHeaderName         *string `tfsdk:"http_header_name" json:"httpHeaderName,omitempty"`
						HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" json:"httpQueryParameterName,omitempty"`
						MinimumRingSize        *int64  `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
						UseSourceIp            *bool   `tfsdk:"use_source_ip" json:"useSourceIp,omitempty"`
					} `tfsdk:"consistent_hash" json:"consistentHash,omitempty"`
					LocalityLbSetting *struct {
						Distribute *[]struct {
							From *string            `tfsdk:"from" json:"from,omitempty"`
							To   *map[string]string `tfsdk:"to" json:"to,omitempty"`
						} `tfsdk:"distribute" json:"distribute,omitempty"`
						Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
						Failover *[]struct {
							From *string `tfsdk:"from" json:"from,omitempty"`
							To   *string `tfsdk:"to" json:"to,omitempty"`
						} `tfsdk:"failover" json:"failover,omitempty"`
					} `tfsdk:"locality_lb_setting" json:"localityLbSetting,omitempty"`
					Simple *string `tfsdk:"simple" json:"simple,omitempty"`
				} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
				OutlierDetection *struct {
					BaseEjectionTime         *string `tfsdk:"base_ejection_time" json:"baseEjectionTime,omitempty"`
					Consecutive5xxErrors     *int64  `tfsdk:"consecutive5xx_errors" json:"consecutive5xxErrors,omitempty"`
					ConsecutiveErrors        *int64  `tfsdk:"consecutive_errors" json:"consecutiveErrors,omitempty"`
					ConsecutiveGatewayErrors *int64  `tfsdk:"consecutive_gateway_errors" json:"consecutiveGatewayErrors,omitempty"`
					Interval                 *string `tfsdk:"interval" json:"interval,omitempty"`
					MaxEjectionPercent       *int64  `tfsdk:"max_ejection_percent" json:"maxEjectionPercent,omitempty"`
					MinHealthPercent         *int64  `tfsdk:"min_health_percent" json:"minHealthPercent,omitempty"`
				} `tfsdk:"outlier_detection" json:"outlierDetection,omitempty"`
				Tls *struct {
					CaCertificates    *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
					ClientCertificate *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
					Mode              *string   `tfsdk:"mode" json:"mode,omitempty"`
					PrivateKey        *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
					Sni               *string   `tfsdk:"sni" json:"sni,omitempty"`
					SubjectAltNames   *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"traffic_policy" json:"trafficPolicy,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		SkipAnalysis *bool `tfsdk:"skip_analysis" json:"skipAnalysis,omitempty"`
		Suspend      *bool `tfsdk:"suspend" json:"suspend,omitempty"`
		TargetRef    *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		UpstreamRef *struct {
			ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			Kind       *string `tfsdk:"kind" json:"kind,omitempty"`
			Name       *string `tfsdk:"name" json:"name,omitempty"`
			Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"upstream_ref" json:"upstreamRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *FlaggerAppCanaryV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_flagger_app_canary_v1beta1"
}

func (r *FlaggerAppCanaryV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Canary is the Schema for the Canary API.",
		MarkdownDescription: "Canary is the Schema for the Canary API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "CanarySpec defines the desired state of a Canary.",
				MarkdownDescription: "CanarySpec defines the desired state of a Canary.",
				Attributes: map[string]schema.Attribute{
					"analysis": schema.SingleNestedAttribute{
						Description:         "Canary analysis for this canary",
						MarkdownDescription: "Canary analysis for this canary",
						Attributes: map[string]schema.Attribute{
							"alerts": schema.ListNestedAttribute{
								Description:         "Alert list for this canary analysis",
								MarkdownDescription: "Alert list for this canary analysis",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the this alert",
											MarkdownDescription: "Name of the this alert",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"provider_ref": schema.SingleNestedAttribute{
											Description:         "Alert provider reference",
											MarkdownDescription: "Alert provider reference",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the alert provider",
													MarkdownDescription: "Name of the alert provider",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of the alert provider",
													MarkdownDescription: "Namespace of the alert provider",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"severity": schema.StringAttribute{
											Description:         "Severity level can be info, warn, error (default info)",
											MarkdownDescription: "Severity level can be info, warn, error (default info)",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"canary_ready_threshold": schema.Float64Attribute{
								Description:         "Percentage of pods that need to be available to consider canary as ready",
								MarkdownDescription: "Percentage of pods that need to be available to consider canary as ready",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"interval": schema.StringAttribute{
								Description:         "Schedule interval for this canary",
								MarkdownDescription: "Schedule interval for this canary",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"iterations": schema.Float64Attribute{
								Description:         "Number of checks to run for A/B Testing and Blue/Green",
								MarkdownDescription: "Number of checks to run for A/B Testing and Blue/Green",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"match": schema.ListNestedAttribute{
								Description:         "A/B testing match conditions",
								MarkdownDescription: "A/B testing match conditions",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"headers": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax)",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax)",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"suffix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"query_params": schema.SingleNestedAttribute{
											Description:         "Query parameters for matching.",
											MarkdownDescription: "Query parameters for matching.",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"source_labels": schema.MapAttribute{
											Description:         "Applicable only when the 'mesh' gateway is included in the service.gateways list",
											MarkdownDescription: "Applicable only when the 'mesh' gateway is included in the service.gateways list",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"max_weight": schema.Float64Attribute{
								Description:         "Max traffic weight routed to canary",
								MarkdownDescription: "Max traffic weight routed to canary",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"metrics": schema.ListNestedAttribute{
								Description:         "Metric check list for this canary",
								MarkdownDescription: "Metric check list for this canary",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"interval": schema.StringAttribute{
											Description:         "Interval of the query",
											MarkdownDescription: "Interval of the query",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the metric",
											MarkdownDescription: "Name of the metric",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"query": schema.StringAttribute{
											Description:         "Prometheus query",
											MarkdownDescription: "Prometheus query",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"template_ref": schema.SingleNestedAttribute{
											Description:         "Metric template reference",
											MarkdownDescription: "Metric template reference",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of this metric template",
													MarkdownDescription: "Name of this metric template",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace": schema.StringAttribute{
													Description:         "Namespace of this metric template",
													MarkdownDescription: "Namespace of this metric template",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"template_variables": schema.MapAttribute{
											Description:         "Additional variables to be used in the metrics query (key-value pairs)",
											MarkdownDescription: "Additional variables to be used in the metrics query (key-value pairs)",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"threshold": schema.Float64Attribute{
											Description:         "Max value accepted for this metric",
											MarkdownDescription: "Max value accepted for this metric",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"threshold_range": schema.SingleNestedAttribute{
											Description:         "Range accepted for this metric",
											MarkdownDescription: "Range accepted for this metric",
											Attributes: map[string]schema.Attribute{
												"max": schema.Float64Attribute{
													Description:         "Max value accepted for this metric",
													MarkdownDescription: "Max value accepted for this metric",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min": schema.Float64Attribute{
													Description:         "Min value accepted for this metric",
													MarkdownDescription: "Min value accepted for this metric",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"mirror": schema.BoolAttribute{
								Description:         "Mirror traffic to canary",
								MarkdownDescription: "Mirror traffic to canary",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"mirror_weight": schema.Float64Attribute{
								Description:         "Weight of traffic to be mirrored",
								MarkdownDescription: "Weight of traffic to be mirrored",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"primary_ready_threshold": schema.Float64Attribute{
								Description:         "Percentage of pods that need to be available to consider primary as ready",
								MarkdownDescription: "Percentage of pods that need to be available to consider primary as ready",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"session_affinity": schema.SingleNestedAttribute{
								Description:         "SessionAffinity represents the session affinity settings for a canary run.",
								MarkdownDescription: "SessionAffinity represents the session affinity settings for a canary run.",
								Attributes: map[string]schema.Attribute{
									"cookie_name": schema.StringAttribute{
										Description:         "CookieName is the key that will be used for the session affinity cookie.",
										MarkdownDescription: "CookieName is the key that will be used for the session affinity cookie.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_age": schema.Float64Attribute{
										Description:         "MaxAge indicates the number of seconds until the session affinity cookie will expire.",
										MarkdownDescription: "MaxAge indicates the number of seconds until the session affinity cookie will expire.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"step_weight": schema.Float64Attribute{
								Description:         "Incremental traffic step weight for the analysis phase",
								MarkdownDescription: "Incremental traffic step weight for the analysis phase",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"step_weight_promotion": schema.Float64Attribute{
								Description:         "Incremental traffic step weight for the promotion phase",
								MarkdownDescription: "Incremental traffic step weight for the promotion phase",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"step_weights": schema.ListAttribute{
								Description:         "Incremental traffic step weights for the analysis phase",
								MarkdownDescription: "Incremental traffic step weights for the analysis phase",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"threshold": schema.Float64Attribute{
								Description:         "Max number of failed checks before rollback",
								MarkdownDescription: "Max number of failed checks before rollback",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"webhooks": schema.ListNestedAttribute{
								Description:         "Webhook list for this canary",
								MarkdownDescription: "Webhook list for this canary",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"metadata": schema.MapAttribute{
											Description:         "Metadata (key-value pairs) for this webhook",
											MarkdownDescription: "Metadata (key-value pairs) for this webhook",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mute_alert": schema.BoolAttribute{
											Description:         "Mute all alerts for the webhook",
											MarkdownDescription: "Mute all alerts for the webhook",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the webhook",
											MarkdownDescription: "Name of the webhook",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"timeout": schema.StringAttribute{
											Description:         "Request timeout for this webhook",
											MarkdownDescription: "Request timeout for this webhook",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"type": schema.StringAttribute{
											Description:         "Type of the webhook pre, post or during rollout",
											MarkdownDescription: "Type of the webhook pre, post or during rollout",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"url": schema.StringAttribute{
											Description:         "URL address of this webhook",
											MarkdownDescription: "URL address of this webhook",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"autoscaler_ref": schema.SingleNestedAttribute{
						Description:         "Scaler selector",
						MarkdownDescription: "Scaler selector",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"primary_scaler_queries": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"primary_scaler_replicas": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"max_replicas": schema.Float64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"min_replicas": schema.Float64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ingress_ref": schema.SingleNestedAttribute{
						Description:         "Ingress selector",
						MarkdownDescription: "Ingress selector",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"metrics_server": schema.StringAttribute{
						Description:         "Prometheus URL",
						MarkdownDescription: "Prometheus URL",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"progress_deadline_seconds": schema.Float64Attribute{
						Description:         "Deployment progress deadline",
						MarkdownDescription: "Deployment progress deadline",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"provider": schema.StringAttribute{
						Description:         "Traffic managent provider",
						MarkdownDescription: "Traffic managent provider",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"revert_on_deletion": schema.BoolAttribute{
						Description:         "Revert mutated resources to original spec on deletion",
						MarkdownDescription: "Revert mutated resources to original spec on deletion",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"route_ref": schema.SingleNestedAttribute{
						Description:         "APISIX route selector",
						MarkdownDescription: "APISIX route selector",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "Kubernetes Service spec",
						MarkdownDescription: "Kubernetes Service spec",
						Attributes: map[string]schema.Attribute{
							"apex": schema.SingleNestedAttribute{
								Description:         "Metadata to add to the apex service",
								MarkdownDescription: "Metadata to add to the apex service",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"app_protocol": schema.StringAttribute{
								Description:         "Application protocol of the port",
								MarkdownDescription: "Application protocol of the port",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"backends": schema.ListAttribute{
								Description:         "AppMesh backend array",
								MarkdownDescription: "AppMesh backend array",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"canary": schema.SingleNestedAttribute{
								Description:         "Metadata to add to the canary service",
								MarkdownDescription: "Metadata to add to the canary service",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"cors_policy": schema.SingleNestedAttribute{
								Description:         "Istio Cross-Origin Resource Sharing policy (CORS)",
								MarkdownDescription: "Istio Cross-Origin Resource Sharing policy (CORS)",
								Attributes: map[string]schema.Attribute{
									"allow_credentials": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"allow_headers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"allow_methods": schema.ListAttribute{
										Description:         "List of HTTP methods allowed to access the resource",
										MarkdownDescription: "List of HTTP methods allowed to access the resource",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"allow_origin": schema.ListAttribute{
										Description:         "The list of origins that are allowed to perform CORS requests.",
										MarkdownDescription: "The list of origins that are allowed to perform CORS requests.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"allow_origins": schema.ListNestedAttribute{
										Description:         "String patterns that match allowed origins",
										MarkdownDescription: "String patterns that match allowed origins",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"expose_headers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_age": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"delegation": schema.BoolAttribute{
								Description:         "enable behaving as a delegate VirtualService",
								MarkdownDescription: "enable behaving as a delegate VirtualService",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"gateway_refs": schema.ListNestedAttribute{
								Description:         "The list of parent Gateways for a HTTPRoute",
								MarkdownDescription: "The list of parent Gateways for a HTTPRoute",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"group": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"kind": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"namespace": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"section_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"gateways": schema.ListAttribute{
								Description:         "The list of Istio gateway for this virtual service",
								MarkdownDescription: "The list of Istio gateway for this virtual service",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"headers": schema.SingleNestedAttribute{
								Description:         "Headers operations",
								MarkdownDescription: "Headers operations",
								Attributes: map[string]schema.Attribute{
									"request": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"add": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"remove": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"set": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"response": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"add": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"remove": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"set": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"hosts": schema.ListAttribute{
								Description:         "The list of host names for this service",
								MarkdownDescription: "The list of host names for this service",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"match": schema.ListNestedAttribute{
								Description:         "URI match conditions",
								MarkdownDescription: "URI match conditions",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authority": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"gateways": schema.ListAttribute{
											Description:         "Names of gateways where the rule should be applied.",
											MarkdownDescription: "Names of gateways where the rule should be applied.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"headers": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"ignore_uri_case": schema.BoolAttribute{
											Description:         "Flag to specify whether the URI matching should be case-insensitive.",
											MarkdownDescription: "Flag to specify whether the URI matching should be case-insensitive.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"method": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"name": schema.StringAttribute{
											Description:         "The name assigned to a match.",
											MarkdownDescription: "The name assigned to a match.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"port": schema.Int64Attribute{
											Description:         "Specifies the ports on the host that is being addressed.",
											MarkdownDescription: "Specifies the ports on the host that is being addressed.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"query_params": schema.SingleNestedAttribute{
											Description:         "Query parameters for matching.",
											MarkdownDescription: "Query parameters for matching.",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"scheme": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"source_labels": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"source_namespace": schema.StringAttribute{
											Description:         "Source namespace constraining the applicability of a rule to workloads in that namespace.",
											MarkdownDescription: "Source namespace constraining the applicability of a rule to workloads in that namespace.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"uri": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"without_headers": schema.SingleNestedAttribute{
											Description:         "withoutHeader has the same syntax with the header, but has opposite meaning.",
											MarkdownDescription: "withoutHeader has the same syntax with the header, but has opposite meaning.",
											Attributes: map[string]schema.Attribute{
												"exact": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex": schema.StringAttribute{
													Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"mesh_name": schema.StringAttribute{
								Description:         "AppMesh mesh name",
								MarkdownDescription: "AppMesh mesh name",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Kubernetes service name",
								MarkdownDescription: "Kubernetes service name",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port": schema.Float64Attribute{
								Description:         "Container port number",
								MarkdownDescription: "Container port number",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port_discovery": schema.BoolAttribute{
								Description:         "Enable port dicovery",
								MarkdownDescription: "Enable port dicovery",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"port_name": schema.StringAttribute{
								Description:         "Container port name",
								MarkdownDescription: "Container port name",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"primary": schema.SingleNestedAttribute{
								Description:         "Metadata to add to the primary service",
								MarkdownDescription: "Metadata to add to the primary service",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"retries": schema.SingleNestedAttribute{
								Description:         "Retry policy for HTTP requests",
								MarkdownDescription: "Retry policy for HTTP requests",
								Attributes: map[string]schema.Attribute{
									"attempts": schema.Int64Attribute{
										Description:         "Number of retries for a given request",
										MarkdownDescription: "Number of retries for a given request",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"per_try_timeout": schema.StringAttribute{
										Description:         "Timeout per retry attempt for a given request",
										MarkdownDescription: "Timeout per retry attempt for a given request",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"retry_on": schema.StringAttribute{
										Description:         "Specifies the conditions under which retry takes place",
										MarkdownDescription: "Specifies the conditions under which retry takes place",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"rewrite": schema.SingleNestedAttribute{
								Description:         "Rewrite HTTP URIs",
								MarkdownDescription: "Rewrite HTTP URIs",
								Attributes: map[string]schema.Attribute{
									"uri": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"target_port": schema.StringAttribute{
								Description:         "Container target port name",
								MarkdownDescription: "Container target port name",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"timeout": schema.StringAttribute{
								Description:         "HTTP or gRPC request timeout",
								MarkdownDescription: "HTTP or gRPC request timeout",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"traffic_policy": schema.SingleNestedAttribute{
								Description:         "Istio traffic policy",
								MarkdownDescription: "Istio traffic policy",
								Attributes: map[string]schema.Attribute{
									"connection_pool": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"http": schema.SingleNestedAttribute{
												Description:         "HTTP connection pool settings.",
												MarkdownDescription: "HTTP connection pool settings.",
												Attributes: map[string]schema.Attribute{
													"h2_upgrade_policy": schema.StringAttribute{
														Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
														MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"http1_max_pending_requests": schema.Int64Attribute{
														Description:         "Maximum number of pending HTTP requests to a destination.",
														MarkdownDescription: "Maximum number of pending HTTP requests to a destination.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"http2_max_requests": schema.Int64Attribute{
														Description:         "Maximum number of requests to a backend.",
														MarkdownDescription: "Maximum number of requests to a backend.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"idle_timeout": schema.StringAttribute{
														Description:         "The idle timeout for upstream connection pool connections.",
														MarkdownDescription: "The idle timeout for upstream connection pool connections.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_requests_per_connection": schema.Int64Attribute{
														Description:         "Maximum number of requests per connection to a backend.",
														MarkdownDescription: "Maximum number of requests per connection to a backend.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_retries": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"load_balancer": schema.SingleNestedAttribute{
										Description:         "Settings controlling the load balancer algorithms.",
										MarkdownDescription: "Settings controlling the load balancer algorithms.",
										Attributes: map[string]schema.Attribute{
											"consistent_hash": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"http_cookie": schema.SingleNestedAttribute{
														Description:         "Hash based on HTTP cookie.",
														MarkdownDescription: "Hash based on HTTP cookie.",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the cookie.",
																MarkdownDescription: "Name of the cookie.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "Path to set for the cookie.",
																MarkdownDescription: "Path to set for the cookie.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"ttl": schema.StringAttribute{
																Description:         "Lifetime of the cookie.",
																MarkdownDescription: "Lifetime of the cookie.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"http_header_name": schema.StringAttribute{
														Description:         "Hash based on a specific HTTP header.",
														MarkdownDescription: "Hash based on a specific HTTP header.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"http_query_parameter_name": schema.StringAttribute{
														Description:         "Hash based on a specific HTTP query parameter.",
														MarkdownDescription: "Hash based on a specific HTTP query parameter.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"minimum_ring_size": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"use_source_ip": schema.BoolAttribute{
														Description:         "Hash based on the source IP address.",
														MarkdownDescription: "Hash based on the source IP address.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"locality_lb_setting": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"distribute": schema.ListNestedAttribute{
														Description:         "Optional: only one of distribute or failover can be set.",
														MarkdownDescription: "Optional: only one of distribute or failover can be set.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"from": schema.StringAttribute{
																	Description:         "Originating locality, '/' separated, e.g.",
																	MarkdownDescription: "Originating locality, '/' separated, e.g.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"to": schema.MapAttribute{
																	Description:         "Map of upstream localities to traffic distribution weights.",
																	MarkdownDescription: "Map of upstream localities to traffic distribution weights.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"enabled": schema.BoolAttribute{
														Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
														MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"failover": schema.ListNestedAttribute{
														Description:         "Optional: only failover or distribute can be set.",
														MarkdownDescription: "Optional: only failover or distribute can be set.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"from": schema.StringAttribute{
																	Description:         "Originating region.",
																	MarkdownDescription: "Originating region.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"to": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"simple": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"outlier_detection": schema.SingleNestedAttribute{
										Description:         "Settings controlling eviction of unhealthy hosts from the load balancing pool.",
										MarkdownDescription: "Settings controlling eviction of unhealthy hosts from the load balancing pool.",
										Attributes: map[string]schema.Attribute{
											"base_ejection_time": schema.StringAttribute{
												Description:         "Minimum ejection duration.",
												MarkdownDescription: "Minimum ejection duration.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"consecutive5xx_errors": schema.Int64Attribute{
												Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
												MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"consecutive_errors": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"consecutive_gateway_errors": schema.Int64Attribute{
												Description:         "Number of gateway errors before a host is ejected from the connection pool.",
												MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"interval": schema.StringAttribute{
												Description:         "Time interval between ejection sweep analysis.",
												MarkdownDescription: "Time interval between ejection sweep analysis.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_ejection_percent": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"min_health_percent": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "Istio TLS related settings for connections to the upstream service",
										MarkdownDescription: "Istio TLS related settings for connections to the upstream service",
										Attributes: map[string]schema.Attribute{
											"ca_certificates": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"client_certificate": schema.StringAttribute{
												Description:         "REQUIRED if mode is 'MUTUAL'.",
												MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"mode": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"private_key": schema.StringAttribute{
												Description:         "REQUIRED if mode is 'MUTUAL'.",
												MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sni": schema.StringAttribute{
												Description:         "SNI string to present to the server during TLS handshake.",
												MarkdownDescription: "SNI string to present to the server during TLS handshake.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"subject_alt_names": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"skip_analysis": schema.BoolAttribute{
						Description:         "Skip analysis and promote canary",
						MarkdownDescription: "Skip analysis and promote canary",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend Canary disabling/pausing all canary runs",
						MarkdownDescription: "Suspend Canary disabling/pausing all canary runs",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "Target selector",
						MarkdownDescription: "Target selector",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"upstream_ref": schema.SingleNestedAttribute{
						Description:         "Gloo Upstream selector",
						MarkdownDescription: "Gloo Upstream selector",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *FlaggerAppCanaryV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *FlaggerAppCanaryV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_flagger_app_canary_v1beta1")

	var data FlaggerAppCanaryV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "flagger.app", Version: "v1beta1", Resource: "canaries"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse FlaggerAppCanaryV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("flagger.app/v1beta1")
	data.Kind = pointer.String("Canary")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
