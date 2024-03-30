/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cloudfront_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &CloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1Manifest{}
)

func NewCloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1Manifest() datasource.DataSource {
	return &CloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1Manifest{}
}

type CloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1Manifest struct{}

type CloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1ManifestData struct {
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
		ResponseHeadersPolicyConfig *struct {
			Comment    *string `tfsdk:"comment" json:"comment,omitempty"`
			CorsConfig *struct {
				AccessControlAllowCredentials *bool `tfsdk:"access_control_allow_credentials" json:"accessControlAllowCredentials,omitempty"`
				AccessControlAllowHeaders     *struct {
					Items *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"access_control_allow_headers" json:"accessControlAllowHeaders,omitempty"`
				AccessControlAllowMethods *struct {
					Items *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"access_control_allow_methods" json:"accessControlAllowMethods,omitempty"`
				AccessControlAllowOrigins *struct {
					Items *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"access_control_allow_origins" json:"accessControlAllowOrigins,omitempty"`
				AccessControlExposeHeaders *struct {
					Items *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"access_control_expose_headers" json:"accessControlExposeHeaders,omitempty"`
				AccessControlMaxAgeSec *int64 `tfsdk:"access_control_max_age_sec" json:"accessControlMaxAgeSec,omitempty"`
				OriginOverride         *bool  `tfsdk:"origin_override" json:"originOverride,omitempty"`
			} `tfsdk:"cors_config" json:"corsConfig,omitempty"`
			CustomHeadersConfig *struct {
				Items *[]struct {
					Header   *string `tfsdk:"header" json:"header,omitempty"`
					Override *bool   `tfsdk:"override" json:"override,omitempty"`
					Value    *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"items" json:"items,omitempty"`
			} `tfsdk:"custom_headers_config" json:"customHeadersConfig,omitempty"`
			Name                *string `tfsdk:"name" json:"name,omitempty"`
			RemoveHeadersConfig *struct {
				Items *[]struct {
					Header *string `tfsdk:"header" json:"header,omitempty"`
				} `tfsdk:"items" json:"items,omitempty"`
			} `tfsdk:"remove_headers_config" json:"removeHeadersConfig,omitempty"`
			SecurityHeadersConfig *struct {
				ContentSecurityPolicy *struct {
					ContentSecurityPolicy *string `tfsdk:"content_security_policy" json:"contentSecurityPolicy,omitempty"`
					Override              *bool   `tfsdk:"override" json:"override,omitempty"`
				} `tfsdk:"content_security_policy" json:"contentSecurityPolicy,omitempty"`
				ContentTypeOptions *struct {
					Override *bool `tfsdk:"override" json:"override,omitempty"`
				} `tfsdk:"content_type_options" json:"contentTypeOptions,omitempty"`
				FrameOptions *struct {
					FrameOption *string `tfsdk:"frame_option" json:"frameOption,omitempty"`
					Override    *bool   `tfsdk:"override" json:"override,omitempty"`
				} `tfsdk:"frame_options" json:"frameOptions,omitempty"`
				ReferrerPolicy *struct {
					Override       *bool   `tfsdk:"override" json:"override,omitempty"`
					ReferrerPolicy *string `tfsdk:"referrer_policy" json:"referrerPolicy,omitempty"`
				} `tfsdk:"referrer_policy" json:"referrerPolicy,omitempty"`
				StrictTransportSecurity *struct {
					AccessControlMaxAgeSec *int64 `tfsdk:"access_control_max_age_sec" json:"accessControlMaxAgeSec,omitempty"`
					IncludeSubdomains      *bool  `tfsdk:"include_subdomains" json:"includeSubdomains,omitempty"`
					Override               *bool  `tfsdk:"override" json:"override,omitempty"`
					Preload                *bool  `tfsdk:"preload" json:"preload,omitempty"`
				} `tfsdk:"strict_transport_security" json:"strictTransportSecurity,omitempty"`
				XSSProtection *struct {
					ModeBlock  *bool   `tfsdk:"mode_block" json:"modeBlock,omitempty"`
					Override   *bool   `tfsdk:"override" json:"override,omitempty"`
					Protection *bool   `tfsdk:"protection" json:"protection,omitempty"`
					ReportURI  *string `tfsdk:"report_uri" json:"reportURI,omitempty"`
				} `tfsdk:"x_ss_protection" json:"xSSProtection,omitempty"`
			} `tfsdk:"security_headers_config" json:"securityHeadersConfig,omitempty"`
			ServerTimingHeadersConfig *struct {
				Enabled      *bool    `tfsdk:"enabled" json:"enabled,omitempty"`
				SamplingRate *float64 `tfsdk:"sampling_rate" json:"samplingRate,omitempty"`
			} `tfsdk:"server_timing_headers_config" json:"serverTimingHeadersConfig,omitempty"`
		} `tfsdk:"response_headers_policy_config" json:"responseHeadersPolicyConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest"
}

func (r *CloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ResponseHeadersPolicy is the Schema for the ResponseHeadersPolicies API",
		MarkdownDescription: "ResponseHeadersPolicy is the Schema for the ResponseHeadersPolicies API",
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
				Description:         "ResponseHeadersPolicySpec defines the desired state of ResponseHeadersPolicy.A response headers policy.A response headers policy contains information about a set of HTTP responseheaders.After you create a response headers policy, you can use its ID to attachit to one or more cache behaviors in a CloudFront distribution. When it'sattached to a cache behavior, the response headers policy affects the HTTPheaders that CloudFront includes in HTTP responses to requests that matchthe cache behavior. CloudFront adds or removes response headers accordingto the configuration of the response headers policy.For more information, see Adding or removing HTTP headers in CloudFront responses(https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html)in the Amazon CloudFront Developer Guide.",
				MarkdownDescription: "ResponseHeadersPolicySpec defines the desired state of ResponseHeadersPolicy.A response headers policy.A response headers policy contains information about a set of HTTP responseheaders.After you create a response headers policy, you can use its ID to attachit to one or more cache behaviors in a CloudFront distribution. When it'sattached to a cache behavior, the response headers policy affects the HTTPheaders that CloudFront includes in HTTP responses to requests that matchthe cache behavior. CloudFront adds or removes response headers accordingto the configuration of the response headers policy.For more information, see Adding or removing HTTP headers in CloudFront responses(https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/modifying-response-headers.html)in the Amazon CloudFront Developer Guide.",
				Attributes: map[string]schema.Attribute{
					"response_headers_policy_config": schema.SingleNestedAttribute{
						Description:         "Contains metadata about the response headers policy, and a set of configurationsthat specify the HTTP headers.",
						MarkdownDescription: "Contains metadata about the response headers policy, and a set of configurationsthat specify the HTTP headers.",
						Attributes: map[string]schema.Attribute{
							"comment": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cors_config": schema.SingleNestedAttribute{
								Description:         "A configuration for a set of HTTP response headers that are used for cross-originresource sharing (CORS). CloudFront adds these headers to HTTP responsesthat it sends for CORS requests that match a cache behavior associated withthis response headers policy.For more information about CORS, see Cross-Origin Resource Sharing (CORS)(https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) in the MDN Web Docs.",
								MarkdownDescription: "A configuration for a set of HTTP response headers that are used for cross-originresource sharing (CORS). CloudFront adds these headers to HTTP responsesthat it sends for CORS requests that match a cache behavior associated withthis response headers policy.For more information about CORS, see Cross-Origin Resource Sharing (CORS)(https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) in the MDN Web Docs.",
								Attributes: map[string]schema.Attribute{
									"access_control_allow_credentials": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"access_control_allow_headers": schema.SingleNestedAttribute{
										Description:         "A list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-HeadersHTTP response header.For more information about the Access-Control-Allow-Headers HTTP responseheader, see Access-Control-Allow-Headers (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers)in the MDN Web Docs.",
										MarkdownDescription: "A list of HTTP header names that CloudFront includes as values for the Access-Control-Allow-HeadersHTTP response header.For more information about the Access-Control-Allow-Headers HTTP responseheader, see Access-Control-Allow-Headers (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Headers)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"items": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"access_control_allow_methods": schema.SingleNestedAttribute{
										Description:         "A list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-MethodsHTTP response header.For more information about the Access-Control-Allow-Methods HTTP responseheader, see Access-Control-Allow-Methods (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods)in the MDN Web Docs.",
										MarkdownDescription: "A list of HTTP methods that CloudFront includes as values for the Access-Control-Allow-MethodsHTTP response header.For more information about the Access-Control-Allow-Methods HTTP responseheader, see Access-Control-Allow-Methods (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Methods)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"items": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"access_control_allow_origins": schema.SingleNestedAttribute{
										Description:         "A list of origins (domain names) that CloudFront can use as the value forthe Access-Control-Allow-Origin HTTP response header.For more information about the Access-Control-Allow-Origin HTTP responseheader, see Access-Control-Allow-Origin (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin)in the MDN Web Docs.",
										MarkdownDescription: "A list of origins (domain names) that CloudFront can use as the value forthe Access-Control-Allow-Origin HTTP response header.For more information about the Access-Control-Allow-Origin HTTP responseheader, see Access-Control-Allow-Origin (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Allow-Origin)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"items": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"access_control_expose_headers": schema.SingleNestedAttribute{
										Description:         "A list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-HeadersHTTP response header.For more information about the Access-Control-Expose-Headers HTTP responseheader, see Access-Control-Expose-Headers (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers)in the MDN Web Docs.",
										MarkdownDescription: "A list of HTTP headers that CloudFront includes as values for the Access-Control-Expose-HeadersHTTP response header.For more information about the Access-Control-Expose-Headers HTTP responseheader, see Access-Control-Expose-Headers (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Access-Control-Expose-Headers)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"items": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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

									"access_control_max_age_sec": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"origin_override": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_headers_config": schema.SingleNestedAttribute{
								Description:         "A list of HTTP response header names and their values. CloudFront includesthese headers in HTTP responses that it sends for requests that match a cachebehavior that's associated with this response headers policy.",
								MarkdownDescription: "A list of HTTP response header names and their values. CloudFront includesthese headers in HTTP responses that it sends for requests that match a cachebehavior that's associated with this response headers policy.",
								Attributes: map[string]schema.Attribute{
									"items": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"header": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"override": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"remove_headers_config": schema.SingleNestedAttribute{
								Description:         "A list of HTTP header names that CloudFront removes from HTTP responses torequests that match the cache behavior that this response headers policyis attached to.",
								MarkdownDescription: "A list of HTTP header names that CloudFront removes from HTTP responses torequests that match the cache behavior that this response headers policyis attached to.",
								Attributes: map[string]schema.Attribute{
									"items": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"header": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

							"security_headers_config": schema.SingleNestedAttribute{
								Description:         "A configuration for a set of security-related HTTP response headers. CloudFrontadds these headers to HTTP responses that it sends for requests that matcha cache behavior associated with this response headers policy.",
								MarkdownDescription: "A configuration for a set of security-related HTTP response headers. CloudFrontadds these headers to HTTP responses that it sends for requests that matcha cache behavior associated with this response headers policy.",
								Attributes: map[string]schema.Attribute{
									"content_security_policy": schema.SingleNestedAttribute{
										Description:         "The policy directives and their values that CloudFront includes as valuesfor the Content-Security-Policy HTTP response header.For more information about the Content-Security-Policy HTTP response header,see Content-Security-Policy (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy)in the MDN Web Docs.",
										MarkdownDescription: "The policy directives and their values that CloudFront includes as valuesfor the Content-Security-Policy HTTP response header.For more information about the Content-Security-Policy HTTP response header,see Content-Security-Policy (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"content_security_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"override": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"content_type_options": schema.SingleNestedAttribute{
										Description:         "Determines whether CloudFront includes the X-Content-Type-Options HTTP responseheader with its value set to nosniff.For more information about the X-Content-Type-Options HTTP response header,see X-Content-Type-Options (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options)in the MDN Web Docs.",
										MarkdownDescription: "Determines whether CloudFront includes the X-Content-Type-Options HTTP responseheader with its value set to nosniff.For more information about the X-Content-Type-Options HTTP response header,see X-Content-Type-Options (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"override": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"frame_options": schema.SingleNestedAttribute{
										Description:         "Determines whether CloudFront includes the X-Frame-Options HTTP responseheader and the header's value.For more information about the X-Frame-Options HTTP response header, seeX-Frame-Options (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options)in the MDN Web Docs.",
										MarkdownDescription: "Determines whether CloudFront includes the X-Frame-Options HTTP responseheader and the header's value.For more information about the X-Frame-Options HTTP response header, seeX-Frame-Options (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"frame_option": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"override": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"referrer_policy": schema.SingleNestedAttribute{
										Description:         "Determines whether CloudFront includes the Referrer-Policy HTTP responseheader and the header's value.For more information about the Referrer-Policy HTTP response header, seeReferrer-Policy (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy)in the MDN Web Docs.",
										MarkdownDescription: "Determines whether CloudFront includes the Referrer-Policy HTTP responseheader and the header's value.For more information about the Referrer-Policy HTTP response header, seeReferrer-Policy (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"override": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"referrer_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"strict_transport_security": schema.SingleNestedAttribute{
										Description:         "Determines whether CloudFront includes the Strict-Transport-Security HTTPresponse header and the header's value.For more information about the Strict-Transport-Security HTTP response header,see Strict-Transport-Security (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security)in the MDN Web Docs.",
										MarkdownDescription: "Determines whether CloudFront includes the Strict-Transport-Security HTTPresponse header and the header's value.For more information about the Strict-Transport-Security HTTP response header,see Strict-Transport-Security (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"access_control_max_age_sec": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"include_subdomains": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"override": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"preload": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"x_ss_protection": schema.SingleNestedAttribute{
										Description:         "Determines whether CloudFront includes the X-XSS-Protection HTTP responseheader and the header's value.For more information about the X-XSS-Protection HTTP response header, seeX-XSS-Protection (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection)in the MDN Web Docs.",
										MarkdownDescription: "Determines whether CloudFront includes the X-XSS-Protection HTTP responseheader and the header's value.For more information about the X-XSS-Protection HTTP response header, seeX-XSS-Protection (https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection)in the MDN Web Docs.",
										Attributes: map[string]schema.Attribute{
											"mode_block": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"override": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"protection": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"report_uri": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

							"server_timing_headers_config": schema.SingleNestedAttribute{
								Description:         "A configuration for enabling the Server-Timing header in HTTP responses sentfrom CloudFront. CloudFront adds this header to HTTP responses that it sendsin response to requests that match a cache behavior that's associated withthis response headers policy.You can use the Server-Timing header to view metrics that can help you gaininsights about the behavior and performance of CloudFront. For example, youcan see which cache layer served a cache hit, or the first byte latency fromthe origin when there was a cache miss. You can use the metrics in the Server-Timingheader to troubleshoot issues or test the efficiency of your CloudFront configuration.For more information, see Server-Timing header (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/understanding-response-headers-policies.html#server-timing-header)in the Amazon CloudFront Developer Guide.",
								MarkdownDescription: "A configuration for enabling the Server-Timing header in HTTP responses sentfrom CloudFront. CloudFront adds this header to HTTP responses that it sendsin response to requests that match a cache behavior that's associated withthis response headers policy.You can use the Server-Timing header to view metrics that can help you gaininsights about the behavior and performance of CloudFront. For example, youcan see which cache layer served a cache hit, or the first byte latency fromthe origin when there was a cache miss. You can use the metrics in the Server-Timingheader to troubleshoot issues or test the efficiency of your CloudFront configuration.For more information, see Server-Timing header (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/understanding-response-headers-policies.html#server-timing-header)in the Amazon CloudFront Developer Guide.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sampling_rate": schema.Float64Attribute{
										Description:         "",
										MarkdownDescription: "",
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
	}
}

func (r *CloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloudfront_services_k8s_aws_response_headers_policy_v1alpha1_manifest")

	var model CloudfrontServicesK8SAwsResponseHeadersPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cloudfront.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("ResponseHeadersPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
