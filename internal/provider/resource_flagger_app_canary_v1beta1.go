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

type FlaggerAppCanaryV1Beta1Resource struct{}

var (
	_ resource.Resource = (*FlaggerAppCanaryV1Beta1Resource)(nil)
)

type FlaggerAppCanaryV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type FlaggerAppCanaryV1Beta1GoModel struct {
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
		Analysis *struct {
			Alerts *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ProviderRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"provider_ref" yaml:"providerRef,omitempty"`

				Severity *string `tfsdk:"severity" yaml:"severity,omitempty"`
			} `tfsdk:"alerts" yaml:"alerts,omitempty"`

			CanaryReadyThreshold utilities.DynamicNumber `tfsdk:"canary_ready_threshold" yaml:"canaryReadyThreshold,omitempty"`

			Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

			Iterations utilities.DynamicNumber `tfsdk:"iterations" yaml:"iterations,omitempty"`

			Match *[]struct {
				Headers *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

					Suffix *string `tfsdk:"suffix" yaml:"suffix,omitempty"`
				} `tfsdk:"headers" yaml:"headers,omitempty"`

				SourceLabels *map[string]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			MaxWeight utilities.DynamicNumber `tfsdk:"max_weight" yaml:"maxWeight,omitempty"`

			Metrics *[]struct {
				Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Query *string `tfsdk:"query" yaml:"query,omitempty"`

				TemplateRef *struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
				} `tfsdk:"template_ref" yaml:"templateRef,omitempty"`

				Threshold utilities.DynamicNumber `tfsdk:"threshold" yaml:"threshold,omitempty"`

				ThresholdRange *struct {
					Max utilities.DynamicNumber `tfsdk:"max" yaml:"max,omitempty"`

					Min utilities.DynamicNumber `tfsdk:"min" yaml:"min,omitempty"`
				} `tfsdk:"threshold_range" yaml:"thresholdRange,omitempty"`
			} `tfsdk:"metrics" yaml:"metrics,omitempty"`

			Mirror *bool `tfsdk:"mirror" yaml:"mirror,omitempty"`

			MirrorWeight utilities.DynamicNumber `tfsdk:"mirror_weight" yaml:"mirrorWeight,omitempty"`

			PrimaryReadyThreshold utilities.DynamicNumber `tfsdk:"primary_ready_threshold" yaml:"primaryReadyThreshold,omitempty"`

			StepWeight utilities.DynamicNumber `tfsdk:"step_weight" yaml:"stepWeight,omitempty"`

			StepWeightPromotion utilities.DynamicNumber `tfsdk:"step_weight_promotion" yaml:"stepWeightPromotion,omitempty"`

			StepWeights *[]string `tfsdk:"step_weights" yaml:"stepWeights,omitempty"`

			Threshold utilities.DynamicNumber `tfsdk:"threshold" yaml:"threshold,omitempty"`

			Webhooks *[]struct {
				Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

				MuteAlert *bool `tfsdk:"mute_alert" yaml:"muteAlert,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"webhooks" yaml:"webhooks,omitempty"`
		} `tfsdk:"analysis" yaml:"analysis,omitempty"`

		AutoscalerRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			PrimaryScalerQueries *map[string]string `tfsdk:"primary_scaler_queries" yaml:"primaryScalerQueries,omitempty"`
		} `tfsdk:"autoscaler_ref" yaml:"autoscalerRef,omitempty"`

		IngressRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"ingress_ref" yaml:"ingressRef,omitempty"`

		MetricsServer *string `tfsdk:"metrics_server" yaml:"metricsServer,omitempty"`

		ProgressDeadlineSeconds utilities.DynamicNumber `tfsdk:"progress_deadline_seconds" yaml:"progressDeadlineSeconds,omitempty"`

		Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`

		RevertOnDeletion *bool `tfsdk:"revert_on_deletion" yaml:"revertOnDeletion,omitempty"`

		Service *struct {
			Apex *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"apex" yaml:"apex,omitempty"`

			AppProtocol *string `tfsdk:"app_protocol" yaml:"appProtocol,omitempty"`

			Backends *[]string `tfsdk:"backends" yaml:"backends,omitempty"`

			Canary *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"canary" yaml:"canary,omitempty"`

			CorsPolicy *struct {
				AllowCredentials *bool `tfsdk:"allow_credentials" yaml:"allowCredentials,omitempty"`

				AllowHeaders *[]string `tfsdk:"allow_headers" yaml:"allowHeaders,omitempty"`

				AllowMethods *[]string `tfsdk:"allow_methods" yaml:"allowMethods,omitempty"`

				AllowOrigin *[]string `tfsdk:"allow_origin" yaml:"allowOrigin,omitempty"`

				AllowOrigins *[]struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"allow_origins" yaml:"allowOrigins,omitempty"`

				ExposeHeaders *[]string `tfsdk:"expose_headers" yaml:"exposeHeaders,omitempty"`

				MaxAge *string `tfsdk:"max_age" yaml:"maxAge,omitempty"`
			} `tfsdk:"cors_policy" yaml:"corsPolicy,omitempty"`

			Delegation *bool `tfsdk:"delegation" yaml:"delegation,omitempty"`

			GatewayRefs *[]struct {
				Group *string `tfsdk:"group" yaml:"group,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				SectionName *string `tfsdk:"section_name" yaml:"sectionName,omitempty"`
			} `tfsdk:"gateway_refs" yaml:"gatewayRefs,omitempty"`

			Gateways *[]string `tfsdk:"gateways" yaml:"gateways,omitempty"`

			Headers *struct {
				Request *struct {
					Add *map[string]string `tfsdk:"add" yaml:"add,omitempty"`

					Remove *[]string `tfsdk:"remove" yaml:"remove,omitempty"`

					Set *map[string]string `tfsdk:"set" yaml:"set,omitempty"`
				} `tfsdk:"request" yaml:"request,omitempty"`

				Response *struct {
					Add *map[string]string `tfsdk:"add" yaml:"add,omitempty"`

					Remove *[]string `tfsdk:"remove" yaml:"remove,omitempty"`

					Set *map[string]string `tfsdk:"set" yaml:"set,omitempty"`
				} `tfsdk:"response" yaml:"response,omitempty"`
			} `tfsdk:"headers" yaml:"headers,omitempty"`

			Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

			Match *[]struct {
				Authority *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"authority" yaml:"authority,omitempty"`

				Gateways *[]string `tfsdk:"gateways" yaml:"gateways,omitempty"`

				Headers *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"headers" yaml:"headers,omitempty"`

				IgnoreUriCase *bool `tfsdk:"ignore_uri_case" yaml:"ignoreUriCase,omitempty"`

				Method *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"method" yaml:"method,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				QueryParams *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"query_params" yaml:"queryParams,omitempty"`

				Scheme *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"scheme" yaml:"scheme,omitempty"`

				SourceLabels *map[string]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

				SourceNamespace *string `tfsdk:"source_namespace" yaml:"sourceNamespace,omitempty"`

				Uri *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"uri" yaml:"uri,omitempty"`

				WithoutHeaders *struct {
					Exact *string `tfsdk:"exact" yaml:"exact,omitempty"`

					Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

					Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
				} `tfsdk:"without_headers" yaml:"withoutHeaders,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			MeshName *string `tfsdk:"mesh_name" yaml:"meshName,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Port utilities.DynamicNumber `tfsdk:"port" yaml:"port,omitempty"`

			PortDiscovery *bool `tfsdk:"port_discovery" yaml:"portDiscovery,omitempty"`

			PortName *string `tfsdk:"port_name" yaml:"portName,omitempty"`

			Primary *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"primary" yaml:"primary,omitempty"`

			Retries *struct {
				Attempts *int64 `tfsdk:"attempts" yaml:"attempts,omitempty"`

				PerTryTimeout *string `tfsdk:"per_try_timeout" yaml:"perTryTimeout,omitempty"`

				RetryOn *string `tfsdk:"retry_on" yaml:"retryOn,omitempty"`
			} `tfsdk:"retries" yaml:"retries,omitempty"`

			Rewrite *struct {
				Uri *string `tfsdk:"uri" yaml:"uri,omitempty"`
			} `tfsdk:"rewrite" yaml:"rewrite,omitempty"`

			TargetPort utilities.IntOrString `tfsdk:"target_port" yaml:"targetPort,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

			TrafficPolicy *struct {
				ConnectionPool *struct {
					Http *struct {
						H2UpgradePolicy *string `tfsdk:"h2_upgrade_policy" yaml:"h2UpgradePolicy,omitempty"`

						Http1MaxPendingRequests *int64 `tfsdk:"http1_max_pending_requests" yaml:"http1MaxPendingRequests,omitempty"`

						Http2MaxRequests *int64 `tfsdk:"http2_max_requests" yaml:"http2MaxRequests,omitempty"`

						IdleTimeout *string `tfsdk:"idle_timeout" yaml:"idleTimeout,omitempty"`

						MaxRequestsPerConnection *int64 `tfsdk:"max_requests_per_connection" yaml:"maxRequestsPerConnection,omitempty"`

						MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`
					} `tfsdk:"http" yaml:"http,omitempty"`
				} `tfsdk:"connection_pool" yaml:"connectionPool,omitempty"`

				LoadBalancer *struct {
					ConsistentHash *struct {
						HttpCookie *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
						} `tfsdk:"http_cookie" yaml:"httpCookie,omitempty"`

						HttpHeaderName *string `tfsdk:"http_header_name" yaml:"httpHeaderName,omitempty"`

						HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" yaml:"httpQueryParameterName,omitempty"`

						MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`

						UseSourceIp *bool `tfsdk:"use_source_ip" yaml:"useSourceIp,omitempty"`
					} `tfsdk:"consistent_hash" yaml:"consistentHash,omitempty"`

					LocalityLbSetting *struct {
						Distribute *[]struct {
							From *string `tfsdk:"from" yaml:"from,omitempty"`

							To *map[string]string `tfsdk:"to" yaml:"to,omitempty"`
						} `tfsdk:"distribute" yaml:"distribute,omitempty"`

						Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

						Failover *[]struct {
							From *string `tfsdk:"from" yaml:"from,omitempty"`

							To *string `tfsdk:"to" yaml:"to,omitempty"`
						} `tfsdk:"failover" yaml:"failover,omitempty"`
					} `tfsdk:"locality_lb_setting" yaml:"localityLbSetting,omitempty"`

					Simple *string `tfsdk:"simple" yaml:"simple,omitempty"`
				} `tfsdk:"load_balancer" yaml:"loadBalancer,omitempty"`

				OutlierDetection *struct {
					BaseEjectionTime *string `tfsdk:"base_ejection_time" yaml:"baseEjectionTime,omitempty"`

					Consecutive5xxErrors *int64 `tfsdk:"consecutive5xx_errors" yaml:"consecutive5xxErrors,omitempty"`

					ConsecutiveErrors *int64 `tfsdk:"consecutive_errors" yaml:"consecutiveErrors,omitempty"`

					ConsecutiveGatewayErrors *int64 `tfsdk:"consecutive_gateway_errors" yaml:"consecutiveGatewayErrors,omitempty"`

					Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

					MaxEjectionPercent *int64 `tfsdk:"max_ejection_percent" yaml:"maxEjectionPercent,omitempty"`

					MinHealthPercent *int64 `tfsdk:"min_health_percent" yaml:"minHealthPercent,omitempty"`
				} `tfsdk:"outlier_detection" yaml:"outlierDetection,omitempty"`

				Tls *struct {
					CaCertificates *string `tfsdk:"ca_certificates" yaml:"caCertificates,omitempty"`

					ClientCertificate *string `tfsdk:"client_certificate" yaml:"clientCertificate,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

					SubjectAltNames *[]string `tfsdk:"subject_alt_names" yaml:"subjectAltNames,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`
			} `tfsdk:"traffic_policy" yaml:"trafficPolicy,omitempty"`
		} `tfsdk:"service" yaml:"service,omitempty"`

		SkipAnalysis *bool `tfsdk:"skip_analysis" yaml:"skipAnalysis,omitempty"`

		TargetRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"target_ref" yaml:"targetRef,omitempty"`

		UpstreamRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"upstream_ref" yaml:"upstreamRef,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewFlaggerAppCanaryV1Beta1Resource() resource.Resource {
	return &FlaggerAppCanaryV1Beta1Resource{}
}

func (r *FlaggerAppCanaryV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flagger_app_canary_v1beta1"
}

func (r *FlaggerAppCanaryV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Canary is the Schema for the Canary API.",
		MarkdownDescription: "Canary is the Schema for the Canary API.",
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
				Description:         "CanarySpec defines the desired state of a Canary.",
				MarkdownDescription: "CanarySpec defines the desired state of a Canary.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"analysis": {
						Description:         "Canary analysis for this canary",
						MarkdownDescription: "Canary analysis for this canary",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"alerts": {
								Description:         "Alert list for this canary analysis",
								MarkdownDescription: "Alert list for this canary analysis",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the this alert",
										MarkdownDescription: "Name of the this alert",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"provider_ref": {
										Description:         "Alert provider reference",
										MarkdownDescription: "Alert provider reference",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the alert provider",
												MarkdownDescription: "Name of the alert provider",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace of the alert provider",
												MarkdownDescription: "Namespace of the alert provider",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: true,
										Optional: false,
										Computed: false,
									},

									"severity": {
										Description:         "Severity level can be info, warn, error (default info)",
										MarkdownDescription: "Severity level can be info, warn, error (default info)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("", "info", "warn", "error"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"canary_ready_threshold": {
								Description:         "Percentage of pods that need to be available to consider canary as ready",
								MarkdownDescription: "Percentage of pods that need to be available to consider canary as ready",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interval": {
								Description:         "Schedule interval for this canary",
								MarkdownDescription: "Schedule interval for this canary",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+(m|s)`), ""),
								},
							},

							"iterations": {
								Description:         "Number of checks to run for A/B Testing and Blue/Green",
								MarkdownDescription: "Number of checks to run for A/B Testing and Blue/Green",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match": {
								Description:         "A/B testing match conditions",
								MarkdownDescription: "A/B testing match conditions",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"headers": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax)",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax)",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"suffix": {
												Description:         "",
												MarkdownDescription: "",

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

									"source_labels": {
										Description:         "Applicable only when the 'mesh' gateway is included in the service.gateways list",
										MarkdownDescription: "Applicable only when the 'mesh' gateway is included in the service.gateways list",

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

							"max_weight": {
								Description:         "Max traffic weight routed to canary",
								MarkdownDescription: "Max traffic weight routed to canary",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics": {
								Description:         "Metric check list for this canary",
								MarkdownDescription: "Metric check list for this canary",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"interval": {
										Description:         "Interval of the query",
										MarkdownDescription: "Interval of the query",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+(m|s)`), ""),
										},
									},

									"name": {
										Description:         "Name of the metric",
										MarkdownDescription: "Name of the metric",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"query": {
										Description:         "Prometheus query",
										MarkdownDescription: "Prometheus query",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"template_ref": {
										Description:         "Metric template reference",
										MarkdownDescription: "Metric template reference",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of this metric template",
												MarkdownDescription: "Name of this metric template",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"namespace": {
												Description:         "Namespace of this metric template",
												MarkdownDescription: "Namespace of this metric template",

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

									"threshold": {
										Description:         "Max value accepted for this metric",
										MarkdownDescription: "Max value accepted for this metric",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"threshold_range": {
										Description:         "Range accepted for this metric",
										MarkdownDescription: "Range accepted for this metric",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max": {
												Description:         "Max value accepted for this metric",
												MarkdownDescription: "Max value accepted for this metric",

												Type: utilities.DynamicNumberType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min": {
												Description:         "Min value accepted for this metric",
												MarkdownDescription: "Min value accepted for this metric",

												Type: utilities.DynamicNumberType{},

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

							"mirror": {
								Description:         "Mirror traffic to canary",
								MarkdownDescription: "Mirror traffic to canary",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mirror_weight": {
								Description:         "Weight of traffic to be mirrored",
								MarkdownDescription: "Weight of traffic to be mirrored",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"primary_ready_threshold": {
								Description:         "Percentage of pods that need to be available to consider primary as ready",
								MarkdownDescription: "Percentage of pods that need to be available to consider primary as ready",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"step_weight": {
								Description:         "Incremental traffic step weight for the analysis phase",
								MarkdownDescription: "Incremental traffic step weight for the analysis phase",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"step_weight_promotion": {
								Description:         "Incremental traffic step weight for the promotion phase",
								MarkdownDescription: "Incremental traffic step weight for the promotion phase",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"step_weights": {
								Description:         "Incremental traffic step weights for the analysis phase",
								MarkdownDescription: "Incremental traffic step weights for the analysis phase",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"threshold": {
								Description:         "Max number of failed checks before rollback",
								MarkdownDescription: "Max number of failed checks before rollback",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"webhooks": {
								Description:         "Webhook list for this canary",
								MarkdownDescription: "Webhook list for this canary",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Metadata (key-value pairs) for this webhook",
										MarkdownDescription: "Metadata (key-value pairs) for this webhook",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mute_alert": {
										Description:         "Mute all alerts for the webhook",
										MarkdownDescription: "Mute all alerts for the webhook",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the webhook",
										MarkdownDescription: "Name of the webhook",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"timeout": {
										Description:         "Request timeout for this webhook",
										MarkdownDescription: "Request timeout for this webhook",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+(m|s)`), ""),
										},
									},

									"type": {
										Description:         "Type of the webhook pre, post or during rollout",
										MarkdownDescription: "Type of the webhook pre, post or during rollout",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("", "confirm-rollout", "pre-rollout", "rollout", "confirm-promotion", "post-rollout", "event", "rollback", "confirm-traffic-increase"),
										},
									},

									"url": {
										Description:         "URL address of this webhook",
										MarkdownDescription: "URL address of this webhook",

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
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"autoscaler_ref": {
						Description:         "Scaler selector",
						MarkdownDescription: "Scaler selector",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("HorizontalPodAutoscaler", "ScaledObject"),
								},
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"primary_scaler_queries": {
								Description:         "",
								MarkdownDescription: "",

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

					"ingress_ref": {
						Description:         "Ingress selector",
						MarkdownDescription: "Ingress selector",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Ingress"),
								},
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

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

					"metrics_server": {
						Description:         "Prometheus URL",
						MarkdownDescription: "Prometheus URL",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"progress_deadline_seconds": {
						Description:         "Deployment progress deadline",
						MarkdownDescription: "Deployment progress deadline",

						Type: utilities.DynamicNumberType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"provider": {
						Description:         "Traffic managent provider",
						MarkdownDescription: "Traffic managent provider",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"revert_on_deletion": {
						Description:         "Revert mutated resources to original spec on deletion",
						MarkdownDescription: "Revert mutated resources to original spec on deletion",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": {
						Description:         "Kubernetes Service spec",
						MarkdownDescription: "Kubernetes Service spec",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"apex": {
								Description:         "Metadata to add to the apex service",
								MarkdownDescription: "Metadata to add to the apex service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "",
										MarkdownDescription: "",

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

							"app_protocol": {
								Description:         "Application protocol of the port",
								MarkdownDescription: "Application protocol of the port",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"backends": {
								Description:         "AppMesh backend array",
								MarkdownDescription: "AppMesh backend array",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"canary": {
								Description:         "Metadata to add to the canary service",
								MarkdownDescription: "Metadata to add to the canary service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "",
										MarkdownDescription: "",

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

							"cors_policy": {
								Description:         "Istio Cross-Origin Resource Sharing policy (CORS)",
								MarkdownDescription: "Istio Cross-Origin Resource Sharing policy (CORS)",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allow_credentials": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_headers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_methods": {
										Description:         "List of HTTP methods allowed to access the resource",
										MarkdownDescription: "List of HTTP methods allowed to access the resource",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_origin": {
										Description:         "The list of origins that are allowed to perform CORS requests.",
										MarkdownDescription: "The list of origins that are allowed to perform CORS requests.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allow_origins": {
										Description:         "String patterns that match allowed origins",
										MarkdownDescription: "String patterns that match allowed origins",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "",
												MarkdownDescription: "",

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

									"expose_headers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_age": {
										Description:         "",
										MarkdownDescription: "",

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

							"delegation": {
								Description:         "enable behaving as a delegate VirtualService",
								MarkdownDescription: "enable behaving as a delegate VirtualService",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway_refs": {
								Description:         "The list of parent Gateways for a HTTPRoute",
								MarkdownDescription: "The list of parent Gateways for a HTTPRoute",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"group": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtMost(253),

											stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
										},
									},

									"kind": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),

											stringvalidator.LengthAtMost(63),

											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
										},
									},

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),

											stringvalidator.LengthAtMost(253),
										},
									},

									"namespace": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),

											stringvalidator.LengthAtMost(63),

											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
										},
									},

									"section_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),

											stringvalidator.LengthAtMost(253),

											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateways": {
								Description:         "The list of Istio gateway for this virtual service",
								MarkdownDescription: "The list of Istio gateway for this virtual service",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"headers": {
								Description:         "Headers operations",
								MarkdownDescription: "Headers operations",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"request": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"add": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"remove": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set": {
												Description:         "",
												MarkdownDescription: "",

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

									"response": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"add": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"remove": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"set": {
												Description:         "",
												MarkdownDescription: "",

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

							"hosts": {
								Description:         "The list of host names for this service",
								MarkdownDescription: "The list of host names for this service",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match": {
								Description:         "URI match conditions",
								MarkdownDescription: "URI match conditions",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"authority": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

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

									"gateways": {
										Description:         "Names of gateways where the rule should be applied.",
										MarkdownDescription: "Names of gateways where the rule should be applied.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"headers": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

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

									"ignore_uri_case": {
										Description:         "Flag to specify whether the URI matching should be case-insensitive.",
										MarkdownDescription: "Flag to specify whether the URI matching should be case-insensitive.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"method": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

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

									"name": {
										Description:         "The name assigned to a match.",
										MarkdownDescription: "The name assigned to a match.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Specifies the ports on the host that is being addressed.",
										MarkdownDescription: "Specifies the ports on the host that is being addressed.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_params": {
										Description:         "Query parameters for matching.",
										MarkdownDescription: "Query parameters for matching.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

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

									"scheme": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

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

									"source_labels": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_namespace": {
										Description:         "Source namespace constraining the applicability of a rule to workloads in that namespace.",
										MarkdownDescription: "Source namespace constraining the applicability of a rule to workloads in that namespace.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uri": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

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

									"without_headers": {
										Description:         "withoutHeader has the same syntax with the header, but has opposite meaning.",
										MarkdownDescription: "withoutHeader has the same syntax with the header, but has opposite meaning.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exact": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prefix": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"regex": {
												Description:         "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",
												MarkdownDescription: "RE2 style regex-based match (https://github.com/google/re2/wiki/Syntax).",

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

							"mesh_name": {
								Description:         "AppMesh mesh name",
								MarkdownDescription: "AppMesh mesh name",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Kubernetes service name",
								MarkdownDescription: "Kubernetes service name",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "Container port number",
								MarkdownDescription: "Container port number",

								Type: utilities.DynamicNumberType{},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"port_discovery": {
								Description:         "Enable port dicovery",
								MarkdownDescription: "Enable port dicovery",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port_name": {
								Description:         "Container port name",
								MarkdownDescription: "Container port name",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"primary": {
								Description:         "Metadata to add to the primary service",
								MarkdownDescription: "Metadata to add to the primary service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "",
										MarkdownDescription: "",

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

							"retries": {
								Description:         "Retry policy for HTTP requests",
								MarkdownDescription: "Retry policy for HTTP requests",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"attempts": {
										Description:         "Number of retries for a given request",
										MarkdownDescription: "Number of retries for a given request",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"per_try_timeout": {
										Description:         "Timeout per retry attempt for a given request",
										MarkdownDescription: "Timeout per retry attempt for a given request",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retry_on": {
										Description:         "Specifies the conditions under which retry takes place",
										MarkdownDescription: "Specifies the conditions under which retry takes place",

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

							"rewrite": {
								Description:         "Rewrite HTTP URIs",
								MarkdownDescription: "Rewrite HTTP URIs",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"uri": {
										Description:         "",
										MarkdownDescription: "",

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

							"target_port": {
								Description:         "Container target port name",
								MarkdownDescription: "Container target port name",

								Type: utilities.IntOrStringType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": {
								Description:         "HTTP or gRPC request timeout",
								MarkdownDescription: "HTTP or gRPC request timeout",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"traffic_policy": {
								Description:         "Istio traffic policy",
								MarkdownDescription: "Istio traffic policy",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"connection_pool": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"http": {
												Description:         "HTTP connection pool settings.",
												MarkdownDescription: "HTTP connection pool settings.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"h2_upgrade_policy": {
														Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
														MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
														},
													},

													"http1_max_pending_requests": {
														Description:         "Maximum number of pending HTTP requests to a destination.",
														MarkdownDescription: "Maximum number of pending HTTP requests to a destination.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http2_max_requests": {
														Description:         "Maximum number of requests to a backend.",
														MarkdownDescription: "Maximum number of requests to a backend.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"idle_timeout": {
														Description:         "The idle timeout for upstream connection pool connections.",
														MarkdownDescription: "The idle timeout for upstream connection pool connections.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_requests_per_connection": {
														Description:         "Maximum number of requests per connection to a backend.",
														MarkdownDescription: "Maximum number of requests per connection to a backend.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_retries": {
														Description:         "",
														MarkdownDescription: "",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"load_balancer": {
										Description:         "Settings controlling the load balancer algorithms.",
										MarkdownDescription: "Settings controlling the load balancer algorithms.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"consistent_hash": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"http_cookie": {
														Description:         "Hash based on HTTP cookie.",
														MarkdownDescription: "Hash based on HTTP cookie.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the cookie.",
																MarkdownDescription: "Name of the cookie.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "Path to set for the cookie.",
																MarkdownDescription: "Path to set for the cookie.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ttl": {
																Description:         "Lifetime of the cookie.",
																MarkdownDescription: "Lifetime of the cookie.",

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

													"http_header_name": {
														Description:         "Hash based on a specific HTTP header.",
														MarkdownDescription: "Hash based on a specific HTTP header.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http_query_parameter_name": {
														Description:         "Hash based on a specific HTTP query parameter.",
														MarkdownDescription: "Hash based on a specific HTTP query parameter.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"minimum_ring_size": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"use_source_ip": {
														Description:         "Hash based on the source IP address.",
														MarkdownDescription: "Hash based on the source IP address.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"locality_lb_setting": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"distribute": {
														Description:         "Optional: only one of distribute or failover can be set.",
														MarkdownDescription: "Optional: only one of distribute or failover can be set.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"from": {
																Description:         "Originating locality, '/' separated, e.g.",
																MarkdownDescription: "Originating locality, '/' separated, e.g.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"to": {
																Description:         "Map of upstream localities to traffic distribution weights.",
																MarkdownDescription: "Map of upstream localities to traffic distribution weights.",

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

													"enabled": {
														Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
														MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"failover": {
														Description:         "Optional: only failover or distribute can be set.",
														MarkdownDescription: "Optional: only failover or distribute can be set.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"from": {
																Description:         "Originating region.",
																MarkdownDescription: "Originating region.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"to": {
																Description:         "",
																MarkdownDescription: "",

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

											"simple": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ROUND_ROBIN", "LEAST_CONN", "RANDOM", "PASSTHROUGH"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"outlier_detection": {
										Description:         "Settings controlling eviction of unhealthy hosts from the load balancing pool.",
										MarkdownDescription: "Settings controlling eviction of unhealthy hosts from the load balancing pool.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"base_ejection_time": {
												Description:         "Minimum ejection duration.",
												MarkdownDescription: "Minimum ejection duration.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive5xx_errors": {
												Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
												MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive_errors": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive_gateway_errors": {
												Description:         "Number of gateway errors before a host is ejected from the connection pool.",
												MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"interval": {
												Description:         "Time interval between ejection sweep analysis.",
												MarkdownDescription: "Time interval between ejection sweep analysis.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_ejection_percent": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_health_percent": {
												Description:         "",
												MarkdownDescription: "",

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

									"tls": {
										Description:         "Istio TLS related settings for connections to the upstream service",
										MarkdownDescription: "Istio TLS related settings for connections to the upstream service",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_certificates": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_certificate": {
												Description:         "REQUIRED if mode is 'MUTUAL'.",
												MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
												},
											},

											"private_key": {
												Description:         "REQUIRED if mode is 'MUTUAL'.",
												MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sni": {
												Description:         "SNI string to present to the server during TLS handshake.",
												MarkdownDescription: "SNI string to present to the server during TLS handshake.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subject_alt_names": {
												Description:         "",
												MarkdownDescription: "",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"skip_analysis": {
						Description:         "Skip analysis and promote canary",
						MarkdownDescription: "Skip analysis and promote canary",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_ref": {
						Description:         "Target selector",
						MarkdownDescription: "Target selector",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("DaemonSet", "Deployment", "Service"),
								},
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},

					"upstream_ref": {
						Description:         "Gloo Upstream selector",
						MarkdownDescription: "Gloo Upstream selector",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"kind": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Upstream"),
								},
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "",
								MarkdownDescription: "",

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
		},
	}, nil
}

func (r *FlaggerAppCanaryV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_flagger_app_canary_v1beta1")

	var state FlaggerAppCanaryV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel FlaggerAppCanaryV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("flagger.app/v1beta1")
	goModel.Kind = utilities.Ptr("Canary")

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

func (r *FlaggerAppCanaryV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_flagger_app_canary_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *FlaggerAppCanaryV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_flagger_app_canary_v1beta1")

	var state FlaggerAppCanaryV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel FlaggerAppCanaryV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("flagger.app/v1beta1")
	goModel.Kind = utilities.Ptr("Canary")

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

func (r *FlaggerAppCanaryV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_flagger_app_canary_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
