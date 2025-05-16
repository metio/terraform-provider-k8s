/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_networking_x_k8s_io_v1alpha1

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
	_ datasource.DataSource = &GatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1Manifest{}
)

func NewGatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1Manifest() datasource.DataSource {
	return &GatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1Manifest{}
}

type GatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1Manifest struct{}

type GatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1ManifestData struct {
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
		RetryConstraint *struct {
			Budget *struct {
				Interval *string `tfsdk:"interval" json:"interval,omitempty"`
				Percent  *int64  `tfsdk:"percent" json:"percent,omitempty"`
			} `tfsdk:"budget" json:"budget,omitempty"`
			MinRetryRate *struct {
				Count    *int64  `tfsdk:"count" json:"count,omitempty"`
				Interval *string `tfsdk:"interval" json:"interval,omitempty"`
			} `tfsdk:"min_retry_rate" json:"minRetryRate,omitempty"`
		} `tfsdk:"retry_constraint" json:"retryConstraint,omitempty"`
		SessionPersistence *struct {
			AbsoluteTimeout *string `tfsdk:"absolute_timeout" json:"absoluteTimeout,omitempty"`
			CookieConfig    *struct {
				LifetimeType *string `tfsdk:"lifetime_type" json:"lifetimeType,omitempty"`
			} `tfsdk:"cookie_config" json:"cookieConfig,omitempty"`
			IdleTimeout *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
			SessionName *string `tfsdk:"session_name" json:"sessionName,omitempty"`
			Type        *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"session_persistence" json:"sessionPersistence,omitempty"`
		TargetRefs *[]struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"target_refs" json:"targetRefs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_networking_x_k8s_io_x_backend_traffic_policy_v1alpha1_manifest"
}

func (r *GatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "XBackendTrafficPolicy defines the configuration for how traffic to a target backend should be handled.",
		MarkdownDescription: "XBackendTrafficPolicy defines the configuration for how traffic to a target backend should be handled.",
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
				Description:         "Spec defines the desired state of BackendTrafficPolicy.",
				MarkdownDescription: "Spec defines the desired state of BackendTrafficPolicy.",
				Attributes: map[string]schema.Attribute{
					"retry_constraint": schema.SingleNestedAttribute{
						Description:         "RetryConstraint defines the configuration for when to allow or prevent further retries to a target backend, by dynamically calculating a 'retry budget'. This budget is calculated based on the percentage of incoming traffic composed of retries over a given time interval. Once the budget is exceeded, additional retries will be rejected. For example, if the retry budget interval is 10 seconds, there have been 1000 active requests in the past 10 seconds, and the allowed percentage of requests that can be retried is 20% (the default), then 200 of those requests may be composed of retries. Active requests will only be considered for the duration of the interval when calculating the retry budget. Retrying the same original request multiple times within the retry budget interval will lead to each retry being counted towards calculating the budget. Configuring a RetryConstraint in BackendTrafficPolicy is compatible with HTTPRoute Retry settings for each HTTPRouteRule that targets the same backend. While the HTTPRouteRule Retry stanza can specify whether a request will be retried, and the number of retry attempts each client may perform, RetryConstraint helps prevent cascading failures such as retry storms during periods of consistent failures. After the retry budget has been exceeded, additional retries to the backend MUST return a 503 response to the client. Additional configurations for defining a constraint on retries MAY be defined in the future. Support: Extended",
						MarkdownDescription: "RetryConstraint defines the configuration for when to allow or prevent further retries to a target backend, by dynamically calculating a 'retry budget'. This budget is calculated based on the percentage of incoming traffic composed of retries over a given time interval. Once the budget is exceeded, additional retries will be rejected. For example, if the retry budget interval is 10 seconds, there have been 1000 active requests in the past 10 seconds, and the allowed percentage of requests that can be retried is 20% (the default), then 200 of those requests may be composed of retries. Active requests will only be considered for the duration of the interval when calculating the retry budget. Retrying the same original request multiple times within the retry budget interval will lead to each retry being counted towards calculating the budget. Configuring a RetryConstraint in BackendTrafficPolicy is compatible with HTTPRoute Retry settings for each HTTPRouteRule that targets the same backend. While the HTTPRouteRule Retry stanza can specify whether a request will be retried, and the number of retry attempts each client may perform, RetryConstraint helps prevent cascading failures such as retry storms during periods of consistent failures. After the retry budget has been exceeded, additional retries to the backend MUST return a 503 response to the client. Additional configurations for defining a constraint on retries MAY be defined in the future. Support: Extended",
						Attributes: map[string]schema.Attribute{
							"budget": schema.SingleNestedAttribute{
								Description:         "Budget holds the details of the retry budget configuration.",
								MarkdownDescription: "Budget holds the details of the retry budget configuration.",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Description:         "Interval defines the duration in which requests will be considered for calculating the budget for retries. Support: Extended",
										MarkdownDescription: "Interval defines the duration in which requests will be considered for calculating the budget for retries. Support: Extended",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
										},
									},

									"percent": schema.Int64Attribute{
										Description:         "Percent defines the maximum percentage of active requests that may be made up of retries. Support: Extended",
										MarkdownDescription: "Percent defines the maximum percentage of active requests that may be made up of retries. Support: Extended",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(100),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"min_retry_rate": schema.SingleNestedAttribute{
								Description:         "MinRetryRate defines the minimum rate of retries that will be allowable over a specified duration of time. The effective overall minimum rate of retries targeting the backend service may be much higher, as there can be any number of clients which are applying this setting locally. This ensures that requests can still be retried during periods of low traffic, where the budget for retries may be calculated as a very low value. Support: Extended",
								MarkdownDescription: "MinRetryRate defines the minimum rate of retries that will be allowable over a specified duration of time. The effective overall minimum rate of retries targeting the backend service may be much higher, as there can be any number of clients which are applying this setting locally. This ensures that requests can still be retried during periods of low traffic, where the budget for retries may be calculated as a very low value. Support: Extended",
								Attributes: map[string]schema.Attribute{
									"count": schema.Int64Attribute{
										Description:         "Count specifies the number of requests per time interval. Support: Extended",
										MarkdownDescription: "Count specifies the number of requests per time interval. Support: Extended",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(1e+06),
										},
									},

									"interval": schema.StringAttribute{
										Description:         "Interval specifies the divisor of the rate of requests, the amount of time during which the given count of requests occur. Support: Extended",
										MarkdownDescription: "Interval specifies the divisor of the rate of requests, the amount of time during which the given count of requests occur. Support: Extended",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
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

					"session_persistence": schema.SingleNestedAttribute{
						Description:         "SessionPersistence defines and configures session persistence for the backend. Support: Extended",
						MarkdownDescription: "SessionPersistence defines and configures session persistence for the backend. Support: Extended",
						Attributes: map[string]schema.Attribute{
							"absolute_timeout": schema.StringAttribute{
								Description:         "AbsoluteTimeout defines the absolute timeout of the persistent session. Once the AbsoluteTimeout duration has elapsed, the session becomes invalid. Support: Extended",
								MarkdownDescription: "AbsoluteTimeout defines the absolute timeout of the persistent session. Once the AbsoluteTimeout duration has elapsed, the session becomes invalid. Support: Extended",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
								},
							},

							"cookie_config": schema.SingleNestedAttribute{
								Description:         "CookieConfig provides configuration settings that are specific to cookie-based session persistence. Support: Core",
								MarkdownDescription: "CookieConfig provides configuration settings that are specific to cookie-based session persistence. Support: Core",
								Attributes: map[string]schema.Attribute{
									"lifetime_type": schema.StringAttribute{
										Description:         "LifetimeType specifies whether the cookie has a permanent or session-based lifetime. A permanent cookie persists until its specified expiry time, defined by the Expires or Max-Age cookie attributes, while a session cookie is deleted when the current session ends. When set to 'Permanent', AbsoluteTimeout indicates the cookie's lifetime via the Expires or Max-Age cookie attributes and is required. When set to 'Session', AbsoluteTimeout indicates the absolute lifetime of the cookie tracked by the gateway and is optional. Defaults to 'Session'. Support: Core for 'Session' type Support: Extended for 'Permanent' type",
										MarkdownDescription: "LifetimeType specifies whether the cookie has a permanent or session-based lifetime. A permanent cookie persists until its specified expiry time, defined by the Expires or Max-Age cookie attributes, while a session cookie is deleted when the current session ends. When set to 'Permanent', AbsoluteTimeout indicates the cookie's lifetime via the Expires or Max-Age cookie attributes and is required. When set to 'Session', AbsoluteTimeout indicates the absolute lifetime of the cookie tracked by the gateway and is optional. Defaults to 'Session'. Support: Core for 'Session' type Support: Extended for 'Permanent' type",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Permanent", "Session"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"idle_timeout": schema.StringAttribute{
								Description:         "IdleTimeout defines the idle timeout of the persistent session. Once the session has been idle for more than the specified IdleTimeout duration, the session becomes invalid. Support: Extended",
								MarkdownDescription: "IdleTimeout defines the idle timeout of the persistent session. Once the session has been idle for more than the specified IdleTimeout duration, the session becomes invalid. Support: Extended",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
								},
							},

							"session_name": schema.StringAttribute{
								Description:         "SessionName defines the name of the persistent session token which may be reflected in the cookie or the header. Users should avoid reusing session names to prevent unintended consequences, such as rejection or unpredictable behavior. Support: Implementation-specific",
								MarkdownDescription: "SessionName defines the name of the persistent session token which may be reflected in the cookie or the header. Users should avoid reusing session names to prevent unintended consequences, such as rejection or unpredictable behavior. Support: Implementation-specific",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(128),
								},
							},

							"type": schema.StringAttribute{
								Description:         "Type defines the type of session persistence such as through the use a header or cookie. Defaults to cookie based session persistence. Support: Core for 'Cookie' type Support: Extended for 'Header' type",
								MarkdownDescription: "Type defines the type of session persistence such as through the use a header or cookie. Defaults to cookie based session persistence. Support: Core for 'Cookie' type Support: Extended for 'Header' type",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Cookie", "Header"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_refs": schema.ListNestedAttribute{
						Description:         "TargetRefs identifies API object(s) to apply this policy to. Currently, Backends (A grouping of like endpoints such as Service, ServiceImport, or any implementation-specific backendRef) are the only valid API target references. Currently, a TargetRef can not be scoped to a specific port on a Service.",
						MarkdownDescription: "TargetRefs identifies API object(s) to apply this policy to. Currently, Backends (A grouping of like endpoints such as Service, ServiceImport, or any implementation-specific backendRef) are the only valid API target references. Currently, a TargetRef can not be scoped to a specific port on a Service.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "Group is the group of the target resource.",
									MarkdownDescription: "Group is the group of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(253),
										stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is kind of the target resource.",
									MarkdownDescription: "Kind is kind of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(63),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the target resource.",
									MarkdownDescription: "Name is the name of the target resource.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(253),
									},
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

func (r *GatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_networking_x_k8s_io_x_backend_traffic_policy_v1alpha1_manifest")

	var model GatewayNetworkingXK8SIoXbackendTrafficPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.networking.x-k8s.io/v1alpha1")
	model.Kind = pointer.String("XBackendTrafficPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
