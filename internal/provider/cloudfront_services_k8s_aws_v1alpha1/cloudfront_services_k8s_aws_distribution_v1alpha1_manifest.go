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
	_ datasource.DataSource = &CloudfrontServicesK8SAwsDistributionV1Alpha1Manifest{}
)

func NewCloudfrontServicesK8SAwsDistributionV1Alpha1Manifest() datasource.DataSource {
	return &CloudfrontServicesK8SAwsDistributionV1Alpha1Manifest{}
}

type CloudfrontServicesK8SAwsDistributionV1Alpha1Manifest struct{}

type CloudfrontServicesK8SAwsDistributionV1Alpha1ManifestData struct {
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
		DistributionConfig *struct {
			Aliases *struct {
				Items *[]string `tfsdk:"items" json:"items,omitempty"`
			} `tfsdk:"aliases" json:"aliases,omitempty"`
			CacheBehaviors *struct {
				Items *[]struct {
					AllowedMethods *struct {
						CachedMethods *struct {
							Items *[]string `tfsdk:"items" json:"items,omitempty"`
						} `tfsdk:"cached_methods" json:"cachedMethods,omitempty"`
						Items *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"allowed_methods" json:"allowedMethods,omitempty"`
					CachePolicyID          *string `tfsdk:"cache_policy_id" json:"cachePolicyID,omitempty"`
					Compress               *bool   `tfsdk:"compress" json:"compress,omitempty"`
					DefaultTTL             *int64  `tfsdk:"default_ttl" json:"defaultTTL,omitempty"`
					FieldLevelEncryptionID *string `tfsdk:"field_level_encryption_id" json:"fieldLevelEncryptionID,omitempty"`
					ForwardedValues        *struct {
						Cookies *struct {
							Forward          *string `tfsdk:"forward" json:"forward,omitempty"`
							WhitelistedNames *struct {
								Items *[]string `tfsdk:"items" json:"items,omitempty"`
							} `tfsdk:"whitelisted_names" json:"whitelistedNames,omitempty"`
						} `tfsdk:"cookies" json:"cookies,omitempty"`
						Headers *struct {
							Items *[]string `tfsdk:"items" json:"items,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						QueryString          *bool `tfsdk:"query_string" json:"queryString,omitempty"`
						QueryStringCacheKeys *struct {
							Items *[]string `tfsdk:"items" json:"items,omitempty"`
						} `tfsdk:"query_string_cache_keys" json:"queryStringCacheKeys,omitempty"`
					} `tfsdk:"forwarded_values" json:"forwardedValues,omitempty"`
					FunctionAssociations *struct {
						Items *[]struct {
							EventType   *string `tfsdk:"event_type" json:"eventType,omitempty"`
							FunctionARN *string `tfsdk:"function_arn" json:"functionARN,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"function_associations" json:"functionAssociations,omitempty"`
					LambdaFunctionAssociations *struct {
						Items *[]struct {
							EventType         *string `tfsdk:"event_type" json:"eventType,omitempty"`
							IncludeBody       *bool   `tfsdk:"include_body" json:"includeBody,omitempty"`
							LambdaFunctionARN *string `tfsdk:"lambda_function_arn" json:"lambdaFunctionARN,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"lambda_function_associations" json:"lambdaFunctionAssociations,omitempty"`
					MaxTTL                  *int64  `tfsdk:"max_ttl" json:"maxTTL,omitempty"`
					MinTTL                  *int64  `tfsdk:"min_ttl" json:"minTTL,omitempty"`
					OriginRequestPolicyID   *string `tfsdk:"origin_request_policy_id" json:"originRequestPolicyID,omitempty"`
					PathPattern             *string `tfsdk:"path_pattern" json:"pathPattern,omitempty"`
					RealtimeLogConfigARN    *string `tfsdk:"realtime_log_config_arn" json:"realtimeLogConfigARN,omitempty"`
					ResponseHeadersPolicyID *string `tfsdk:"response_headers_policy_id" json:"responseHeadersPolicyID,omitempty"`
					SmoothStreaming         *bool   `tfsdk:"smooth_streaming" json:"smoothStreaming,omitempty"`
					TargetOriginID          *string `tfsdk:"target_origin_id" json:"targetOriginID,omitempty"`
					TrustedKeyGroups        *struct {
						Enabled *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
						Items   *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"trusted_key_groups" json:"trustedKeyGroups,omitempty"`
					TrustedSigners *struct {
						Enabled *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
						Items   *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"trusted_signers" json:"trustedSigners,omitempty"`
					ViewerProtocolPolicy *string `tfsdk:"viewer_protocol_policy" json:"viewerProtocolPolicy,omitempty"`
				} `tfsdk:"items" json:"items,omitempty"`
			} `tfsdk:"cache_behaviors" json:"cacheBehaviors,omitempty"`
			Comment                      *string `tfsdk:"comment" json:"comment,omitempty"`
			ContinuousDeploymentPolicyID *string `tfsdk:"continuous_deployment_policy_id" json:"continuousDeploymentPolicyID,omitempty"`
			CustomErrorResponses         *struct {
				Items *[]struct {
					ErrorCachingMinTTL *int64  `tfsdk:"error_caching_min_ttl" json:"errorCachingMinTTL,omitempty"`
					ErrorCode          *int64  `tfsdk:"error_code" json:"errorCode,omitempty"`
					ResponseCode       *string `tfsdk:"response_code" json:"responseCode,omitempty"`
					ResponsePagePath   *string `tfsdk:"response_page_path" json:"responsePagePath,omitempty"`
				} `tfsdk:"items" json:"items,omitempty"`
			} `tfsdk:"custom_error_responses" json:"customErrorResponses,omitempty"`
			DefaultCacheBehavior *struct {
				AllowedMethods *struct {
					CachedMethods *struct {
						Items *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"cached_methods" json:"cachedMethods,omitempty"`
					Items *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"allowed_methods" json:"allowedMethods,omitempty"`
				CachePolicyID          *string `tfsdk:"cache_policy_id" json:"cachePolicyID,omitempty"`
				Compress               *bool   `tfsdk:"compress" json:"compress,omitempty"`
				DefaultTTL             *int64  `tfsdk:"default_ttl" json:"defaultTTL,omitempty"`
				FieldLevelEncryptionID *string `tfsdk:"field_level_encryption_id" json:"fieldLevelEncryptionID,omitempty"`
				ForwardedValues        *struct {
					Cookies *struct {
						Forward          *string `tfsdk:"forward" json:"forward,omitempty"`
						WhitelistedNames *struct {
							Items *[]string `tfsdk:"items" json:"items,omitempty"`
						} `tfsdk:"whitelisted_names" json:"whitelistedNames,omitempty"`
					} `tfsdk:"cookies" json:"cookies,omitempty"`
					Headers *struct {
						Items *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					QueryString          *bool `tfsdk:"query_string" json:"queryString,omitempty"`
					QueryStringCacheKeys *struct {
						Items *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"query_string_cache_keys" json:"queryStringCacheKeys,omitempty"`
				} `tfsdk:"forwarded_values" json:"forwardedValues,omitempty"`
				FunctionAssociations *struct {
					Items *[]struct {
						EventType   *string `tfsdk:"event_type" json:"eventType,omitempty"`
						FunctionARN *string `tfsdk:"function_arn" json:"functionARN,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"function_associations" json:"functionAssociations,omitempty"`
				LambdaFunctionAssociations *struct {
					Items *[]struct {
						EventType         *string `tfsdk:"event_type" json:"eventType,omitempty"`
						IncludeBody       *bool   `tfsdk:"include_body" json:"includeBody,omitempty"`
						LambdaFunctionARN *string `tfsdk:"lambda_function_arn" json:"lambdaFunctionARN,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"lambda_function_associations" json:"lambdaFunctionAssociations,omitempty"`
				MaxTTL                  *int64  `tfsdk:"max_ttl" json:"maxTTL,omitempty"`
				MinTTL                  *int64  `tfsdk:"min_ttl" json:"minTTL,omitempty"`
				OriginRequestPolicyID   *string `tfsdk:"origin_request_policy_id" json:"originRequestPolicyID,omitempty"`
				RealtimeLogConfigARN    *string `tfsdk:"realtime_log_config_arn" json:"realtimeLogConfigARN,omitempty"`
				ResponseHeadersPolicyID *string `tfsdk:"response_headers_policy_id" json:"responseHeadersPolicyID,omitempty"`
				SmoothStreaming         *bool   `tfsdk:"smooth_streaming" json:"smoothStreaming,omitempty"`
				TargetOriginID          *string `tfsdk:"target_origin_id" json:"targetOriginID,omitempty"`
				TrustedKeyGroups        *struct {
					Enabled *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
					Items   *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"trusted_key_groups" json:"trustedKeyGroups,omitempty"`
				TrustedSigners *struct {
					Enabled *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
					Items   *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"trusted_signers" json:"trustedSigners,omitempty"`
				ViewerProtocolPolicy *string `tfsdk:"viewer_protocol_policy" json:"viewerProtocolPolicy,omitempty"`
			} `tfsdk:"default_cache_behavior" json:"defaultCacheBehavior,omitempty"`
			DefaultRootObject *string `tfsdk:"default_root_object" json:"defaultRootObject,omitempty"`
			Enabled           *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			HttpVersion       *string `tfsdk:"http_version" json:"httpVersion,omitempty"`
			IsIPV6Enabled     *bool   `tfsdk:"is_ipv6_enabled" json:"isIPV6Enabled,omitempty"`
			Logging           *struct {
				Bucket         *string `tfsdk:"bucket" json:"bucket,omitempty"`
				Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				IncludeCookies *bool   `tfsdk:"include_cookies" json:"includeCookies,omitempty"`
				Prefix         *string `tfsdk:"prefix" json:"prefix,omitempty"`
			} `tfsdk:"logging" json:"logging,omitempty"`
			OriginGroups *struct {
				Items *[]struct {
					FailoverCriteria *struct {
						StatusCodes *struct {
							Items *[]string `tfsdk:"items" json:"items,omitempty"`
						} `tfsdk:"status_codes" json:"statusCodes,omitempty"`
					} `tfsdk:"failover_criteria" json:"failoverCriteria,omitempty"`
					Id      *string `tfsdk:"id" json:"id,omitempty"`
					Members *struct {
						Items *[]struct {
							OriginID *string `tfsdk:"origin_id" json:"originID,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"members" json:"members,omitempty"`
				} `tfsdk:"items" json:"items,omitempty"`
			} `tfsdk:"origin_groups" json:"originGroups,omitempty"`
			Origins *struct {
				Items *[]struct {
					ConnectionAttempts *int64 `tfsdk:"connection_attempts" json:"connectionAttempts,omitempty"`
					ConnectionTimeout  *int64 `tfsdk:"connection_timeout" json:"connectionTimeout,omitempty"`
					CustomHeaders      *struct {
						Items *[]struct {
							HeaderName  *string `tfsdk:"header_name" json:"headerName,omitempty"`
							HeaderValue *string `tfsdk:"header_value" json:"headerValue,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"custom_headers" json:"customHeaders,omitempty"`
					CustomOriginConfig *struct {
						HttpPort               *int64  `tfsdk:"http_port" json:"httpPort,omitempty"`
						HttpSPort              *int64  `tfsdk:"http_s_port" json:"httpSPort,omitempty"`
						OriginKeepaliveTimeout *int64  `tfsdk:"origin_keepalive_timeout" json:"originKeepaliveTimeout,omitempty"`
						OriginProtocolPolicy   *string `tfsdk:"origin_protocol_policy" json:"originProtocolPolicy,omitempty"`
						OriginReadTimeout      *int64  `tfsdk:"origin_read_timeout" json:"originReadTimeout,omitempty"`
						OriginSSLProtocols     *struct {
							Items *[]string `tfsdk:"items" json:"items,omitempty"`
						} `tfsdk:"origin_ssl_protocols" json:"originSSLProtocols,omitempty"`
					} `tfsdk:"custom_origin_config" json:"customOriginConfig,omitempty"`
					DomainName            *string `tfsdk:"domain_name" json:"domainName,omitempty"`
					Id                    *string `tfsdk:"id" json:"id,omitempty"`
					OriginAccessControlID *string `tfsdk:"origin_access_control_id" json:"originAccessControlID,omitempty"`
					OriginPath            *string `tfsdk:"origin_path" json:"originPath,omitempty"`
					OriginShield          *struct {
						Enabled            *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
						OriginShieldRegion *string `tfsdk:"origin_shield_region" json:"originShieldRegion,omitempty"`
					} `tfsdk:"origin_shield" json:"originShield,omitempty"`
					S3OriginConfig *struct {
						OriginAccessIdentity *string `tfsdk:"origin_access_identity" json:"originAccessIdentity,omitempty"`
					} `tfsdk:"s3_origin_config" json:"s3OriginConfig,omitempty"`
				} `tfsdk:"items" json:"items,omitempty"`
			} `tfsdk:"origins" json:"origins,omitempty"`
			PriceClass   *string `tfsdk:"price_class" json:"priceClass,omitempty"`
			Restrictions *struct {
				GeoRestriction *struct {
					Items           *[]string `tfsdk:"items" json:"items,omitempty"`
					RestrictionType *string   `tfsdk:"restriction_type" json:"restrictionType,omitempty"`
				} `tfsdk:"geo_restriction" json:"geoRestriction,omitempty"`
			} `tfsdk:"restrictions" json:"restrictions,omitempty"`
			Staging           *bool `tfsdk:"staging" json:"staging,omitempty"`
			ViewerCertificate *struct {
				AcmCertificateARN *string `tfsdk:"acm_certificate_arn" json:"acmCertificateARN,omitempty"`
				AcmCertificateRef *struct {
					From *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"from" json:"from,omitempty"`
				} `tfsdk:"acm_certificate_ref" json:"acmCertificateRef,omitempty"`
				Certificate                  *string `tfsdk:"certificate" json:"certificate,omitempty"`
				CertificateSource            *string `tfsdk:"certificate_source" json:"certificateSource,omitempty"`
				CloudFrontDefaultCertificate *bool   `tfsdk:"cloud_front_default_certificate" json:"cloudFrontDefaultCertificate,omitempty"`
				IamCertificateID             *string `tfsdk:"iam_certificate_id" json:"iamCertificateID,omitempty"`
				MinimumProtocolVersion       *string `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
				SslSupportMethod             *string `tfsdk:"ssl_support_method" json:"sslSupportMethod,omitempty"`
			} `tfsdk:"viewer_certificate" json:"viewerCertificate,omitempty"`
			WebACLID *string `tfsdk:"web_aclid" json:"webACLID,omitempty"`
		} `tfsdk:"distribution_config" json:"distributionConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudfrontServicesK8SAwsDistributionV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloudfront_services_k8s_aws_distribution_v1alpha1_manifest"
}

func (r *CloudfrontServicesK8SAwsDistributionV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Distribution is the Schema for the Distributions API",
		MarkdownDescription: "Distribution is the Schema for the Distributions API",
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
				Description:         "DistributionSpec defines the desired state of Distribution. A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.",
				MarkdownDescription: "DistributionSpec defines the desired state of Distribution. A distribution tells CloudFront where you want content to be delivered from, and the details about how to track and manage content delivery.",
				Attributes: map[string]schema.Attribute{
					"distribution_config": schema.SingleNestedAttribute{
						Description:         "The distribution's configuration information.",
						MarkdownDescription: "The distribution's configuration information.",
						Attributes: map[string]schema.Attribute{
							"aliases": schema.SingleNestedAttribute{
								Description:         "A complex type that contains information about CNAMEs (alternate domain names), if any, for this distribution.",
								MarkdownDescription: "A complex type that contains information about CNAMEs (alternate domain names), if any, for this distribution.",
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

							"cache_behaviors": schema.SingleNestedAttribute{
								Description:         "A complex type that contains zero or more CacheBehavior elements.",
								MarkdownDescription: "A complex type that contains zero or more CacheBehavior elements.",
								Attributes: map[string]schema.Attribute{
									"items": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"allowed_methods": schema.SingleNestedAttribute{
													Description:         "A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin. There are three choices: * CloudFront forwards only GET and HEAD requests. * CloudFront forwards only GET, HEAD, and OPTIONS requests. * CloudFront forwards GET, HEAD, OPTIONS, PUT, PATCH, POST, and DELETE requests. If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.",
													MarkdownDescription: "A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin. There are three choices: * CloudFront forwards only GET and HEAD requests. * CloudFront forwards only GET, HEAD, and OPTIONS requests. * CloudFront forwards GET, HEAD, OPTIONS, PUT, PATCH, POST, and DELETE requests. If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.",
													Attributes: map[string]schema.Attribute{
														"cached_methods": schema.SingleNestedAttribute{
															Description:         "A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods. There are two choices: * CloudFront caches responses to GET and HEAD requests. * CloudFront caches responses to GET, HEAD, and OPTIONS requests. If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.",
															MarkdownDescription: "A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods. There are two choices: * CloudFront caches responses to GET and HEAD requests. * CloudFront caches responses to GET, HEAD, and OPTIONS requests. If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.",
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

												"cache_policy_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"compress": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"default_ttl": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_level_encryption_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"forwarded_values": schema.SingleNestedAttribute{
													Description:         "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include values in the cache key, use a cache policy. For more information, see Creating cache policies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the Amazon CloudFront Developer Guide. If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see Creating origin request policies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the Amazon CloudFront Developer Guide. A complex type that specifies how CloudFront handles query strings, cookies, and HTTP headers.",
													MarkdownDescription: "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include values in the cache key, use a cache policy. For more information, see Creating cache policies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the Amazon CloudFront Developer Guide. If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see Creating origin request policies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the Amazon CloudFront Developer Guide. A complex type that specifies how CloudFront handles query strings, cookies, and HTTP headers.",
													Attributes: map[string]schema.Attribute{
														"cookies": schema.SingleNestedAttribute{
															Description:         "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include cookies in the cache key, use CookiesConfig in a cache policy. See CachePolicy. If you want to send cookies to the origin but not include them in the cache key, use CookiesConfig in an origin request policy. See OriginRequestPolicy. A complex type that specifies whether you want CloudFront to forward cookies to the origin and, if so, which ones. For more information about forwarding cookies to the origin, see Caching Content Based on Cookies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Cookies.html) in the Amazon CloudFront Developer Guide.",
															MarkdownDescription: "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include cookies in the cache key, use CookiesConfig in a cache policy. See CachePolicy. If you want to send cookies to the origin but not include them in the cache key, use CookiesConfig in an origin request policy. See OriginRequestPolicy. A complex type that specifies whether you want CloudFront to forward cookies to the origin and, if so, which ones. For more information about forwarding cookies to the origin, see Caching Content Based on Cookies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Cookies.html) in the Amazon CloudFront Developer Guide.",
															Attributes: map[string]schema.Attribute{
																"forward": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"whitelisted_names": schema.SingleNestedAttribute{
																	Description:         "Contains a list of cookie names.",
																	MarkdownDescription: "Contains a list of cookie names.",
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"headers": schema.SingleNestedAttribute{
															Description:         "Contains a list of HTTP header names.",
															MarkdownDescription: "Contains a list of HTTP header names.",
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

														"query_string": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"query_string_cache_keys": schema.SingleNestedAttribute{
															Description:         "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include query strings in the cache key, use QueryStringsConfig in a cache policy. See CachePolicy. If you want to send query strings to the origin but not include them in the cache key, use QueryStringsConfig in an origin request policy. See OriginRequestPolicy. A complex type that contains information about the query string parameters that you want CloudFront to use for caching for a cache behavior.",
															MarkdownDescription: "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include query strings in the cache key, use QueryStringsConfig in a cache policy. See CachePolicy. If you want to send query strings to the origin but not include them in the cache key, use QueryStringsConfig in an origin request policy. See OriginRequestPolicy. A complex type that contains information about the query string parameters that you want CloudFront to use for caching for a cache behavior.",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"function_associations": schema.SingleNestedAttribute{
													Description:         "A list of CloudFront functions that are associated with a cache behavior in a CloudFront distribution. CloudFront functions must be published to the LIVE stage to associate them with a cache behavior.",
													MarkdownDescription: "A list of CloudFront functions that are associated with a cache behavior in a CloudFront distribution. CloudFront functions must be published to the LIVE stage to associate them with a cache behavior.",
													Attributes: map[string]schema.Attribute{
														"items": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"event_type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"function_arn": schema.StringAttribute{
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

												"lambda_function_associations": schema.SingleNestedAttribute{
													Description:         "A complex type that specifies a list of Lambda@Edge functions associations for a cache behavior. If you want to invoke one or more Lambda@Edge functions triggered by requests that match the PathPattern of the cache behavior, specify the applicable values for Quantity and Items. Note that there can be up to 4 LambdaFunctionAssociation items in this list (one for each possible value of EventType) and each EventType can be associated with only one function. If you don't want to invoke any Lambda@Edge functions for the requests that match PathPattern, specify 0 for Quantity and omit Items.",
													MarkdownDescription: "A complex type that specifies a list of Lambda@Edge functions associations for a cache behavior. If you want to invoke one or more Lambda@Edge functions triggered by requests that match the PathPattern of the cache behavior, specify the applicable values for Quantity and Items. Note that there can be up to 4 LambdaFunctionAssociation items in this list (one for each possible value of EventType) and each EventType can be associated with only one function. If you don't want to invoke any Lambda@Edge functions for the requests that match PathPattern, specify 0 for Quantity and omit Items.",
													Attributes: map[string]schema.Attribute{
														"items": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"event_type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"include_body": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"lambda_function_arn": schema.StringAttribute{
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

												"max_ttl": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_ttl": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"origin_request_policy_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path_pattern": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"realtime_log_config_arn": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"response_headers_policy_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"smooth_streaming": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"target_origin_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"trusted_key_groups": schema.SingleNestedAttribute{
													Description:         "A list of key groups whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.",
													MarkdownDescription: "A list of key groups whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.",
													Attributes: map[string]schema.Attribute{
														"enabled": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

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

												"trusted_signers": schema.SingleNestedAttribute{
													Description:         "A list of Amazon Web Services accounts whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.",
													MarkdownDescription: "A list of Amazon Web Services accounts whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.",
													Attributes: map[string]schema.Attribute{
														"enabled": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

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

												"viewer_protocol_policy": schema.StringAttribute{
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

							"comment": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"continuous_deployment_policy_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_error_responses": schema.SingleNestedAttribute{
								Description:         "A complex type that controls: * Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer. * How long CloudFront caches HTTP status codes in the 4xx and 5xx range. For more information about custom error pages, see Customizing Error Responses (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the Amazon CloudFront Developer Guide.",
								MarkdownDescription: "A complex type that controls: * Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer. * How long CloudFront caches HTTP status codes in the 4xx and 5xx range. For more information about custom error pages, see Customizing Error Responses (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the Amazon CloudFront Developer Guide.",
								Attributes: map[string]schema.Attribute{
									"items": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"error_caching_min_ttl": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"error_code": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"response_code": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"response_page_path": schema.StringAttribute{
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

							"default_cache_behavior": schema.SingleNestedAttribute{
								Description:         "A complex type that describes the default cache behavior if you don't specify a CacheBehavior element or if request URLs don't match any of the values of PathPattern in CacheBehavior elements. You must create exactly one default cache behavior.",
								MarkdownDescription: "A complex type that describes the default cache behavior if you don't specify a CacheBehavior element or if request URLs don't match any of the values of PathPattern in CacheBehavior elements. You must create exactly one default cache behavior.",
								Attributes: map[string]schema.Attribute{
									"allowed_methods": schema.SingleNestedAttribute{
										Description:         "A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin. There are three choices: * CloudFront forwards only GET and HEAD requests. * CloudFront forwards only GET, HEAD, and OPTIONS requests. * CloudFront forwards GET, HEAD, OPTIONS, PUT, PATCH, POST, and DELETE requests. If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.",
										MarkdownDescription: "A complex type that controls which HTTP methods CloudFront processes and forwards to your Amazon S3 bucket or your custom origin. There are three choices: * CloudFront forwards only GET and HEAD requests. * CloudFront forwards only GET, HEAD, and OPTIONS requests. * CloudFront forwards GET, HEAD, OPTIONS, PUT, PATCH, POST, and DELETE requests. If you pick the third choice, you may need to restrict access to your Amazon S3 bucket or to your custom origin so users can't perform operations that you don't want them to. For example, you might not want users to have permissions to delete objects from your origin.",
										Attributes: map[string]schema.Attribute{
											"cached_methods": schema.SingleNestedAttribute{
												Description:         "A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods. There are two choices: * CloudFront caches responses to GET and HEAD requests. * CloudFront caches responses to GET, HEAD, and OPTIONS requests. If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.",
												MarkdownDescription: "A complex type that controls whether CloudFront caches the response to requests using the specified HTTP methods. There are two choices: * CloudFront caches responses to GET and HEAD requests. * CloudFront caches responses to GET, HEAD, and OPTIONS requests. If you pick the second choice for your Amazon S3 Origin, you may need to forward Access-Control-Request-Method, Access-Control-Request-Headers, and Origin headers for the responses to be cached correctly.",
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

									"cache_policy_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"compress": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"default_ttl": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"field_level_encryption_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"forwarded_values": schema.SingleNestedAttribute{
										Description:         "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include values in the cache key, use a cache policy. For more information, see Creating cache policies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the Amazon CloudFront Developer Guide. If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see Creating origin request policies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the Amazon CloudFront Developer Guide. A complex type that specifies how CloudFront handles query strings, cookies, and HTTP headers.",
										MarkdownDescription: "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include values in the cache key, use a cache policy. For more information, see Creating cache policies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-the-cache-key.html#cache-key-create-cache-policy) in the Amazon CloudFront Developer Guide. If you want to send values to the origin but not include them in the cache key, use an origin request policy. For more information, see Creating origin request policies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/controlling-origin-requests.html#origin-request-create-origin-request-policy) in the Amazon CloudFront Developer Guide. A complex type that specifies how CloudFront handles query strings, cookies, and HTTP headers.",
										Attributes: map[string]schema.Attribute{
											"cookies": schema.SingleNestedAttribute{
												Description:         "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include cookies in the cache key, use CookiesConfig in a cache policy. See CachePolicy. If you want to send cookies to the origin but not include them in the cache key, use CookiesConfig in an origin request policy. See OriginRequestPolicy. A complex type that specifies whether you want CloudFront to forward cookies to the origin and, if so, which ones. For more information about forwarding cookies to the origin, see Caching Content Based on Cookies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Cookies.html) in the Amazon CloudFront Developer Guide.",
												MarkdownDescription: "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include cookies in the cache key, use CookiesConfig in a cache policy. See CachePolicy. If you want to send cookies to the origin but not include them in the cache key, use CookiesConfig in an origin request policy. See OriginRequestPolicy. A complex type that specifies whether you want CloudFront to forward cookies to the origin and, if so, which ones. For more information about forwarding cookies to the origin, see Caching Content Based on Cookies (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Cookies.html) in the Amazon CloudFront Developer Guide.",
												Attributes: map[string]schema.Attribute{
													"forward": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"whitelisted_names": schema.SingleNestedAttribute{
														Description:         "Contains a list of cookie names.",
														MarkdownDescription: "Contains a list of cookie names.",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"headers": schema.SingleNestedAttribute{
												Description:         "Contains a list of HTTP header names.",
												MarkdownDescription: "Contains a list of HTTP header names.",
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

											"query_string": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"query_string_cache_keys": schema.SingleNestedAttribute{
												Description:         "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include query strings in the cache key, use QueryStringsConfig in a cache policy. See CachePolicy. If you want to send query strings to the origin but not include them in the cache key, use QueryStringsConfig in an origin request policy. See OriginRequestPolicy. A complex type that contains information about the query string parameters that you want CloudFront to use for caching for a cache behavior.",
												MarkdownDescription: "This field is deprecated. We recommend that you use a cache policy or an origin request policy instead of this field. If you want to include query strings in the cache key, use QueryStringsConfig in a cache policy. See CachePolicy. If you want to send query strings to the origin but not include them in the cache key, use QueryStringsConfig in an origin request policy. See OriginRequestPolicy. A complex type that contains information about the query string parameters that you want CloudFront to use for caching for a cache behavior.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"function_associations": schema.SingleNestedAttribute{
										Description:         "A list of CloudFront functions that are associated with a cache behavior in a CloudFront distribution. CloudFront functions must be published to the LIVE stage to associate them with a cache behavior.",
										MarkdownDescription: "A list of CloudFront functions that are associated with a cache behavior in a CloudFront distribution. CloudFront functions must be published to the LIVE stage to associate them with a cache behavior.",
										Attributes: map[string]schema.Attribute{
											"items": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"event_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"function_arn": schema.StringAttribute{
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

									"lambda_function_associations": schema.SingleNestedAttribute{
										Description:         "A complex type that specifies a list of Lambda@Edge functions associations for a cache behavior. If you want to invoke one or more Lambda@Edge functions triggered by requests that match the PathPattern of the cache behavior, specify the applicable values for Quantity and Items. Note that there can be up to 4 LambdaFunctionAssociation items in this list (one for each possible value of EventType) and each EventType can be associated with only one function. If you don't want to invoke any Lambda@Edge functions for the requests that match PathPattern, specify 0 for Quantity and omit Items.",
										MarkdownDescription: "A complex type that specifies a list of Lambda@Edge functions associations for a cache behavior. If you want to invoke one or more Lambda@Edge functions triggered by requests that match the PathPattern of the cache behavior, specify the applicable values for Quantity and Items. Note that there can be up to 4 LambdaFunctionAssociation items in this list (one for each possible value of EventType) and each EventType can be associated with only one function. If you don't want to invoke any Lambda@Edge functions for the requests that match PathPattern, specify 0 for Quantity and omit Items.",
										Attributes: map[string]schema.Attribute{
											"items": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"event_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"include_body": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"lambda_function_arn": schema.StringAttribute{
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

									"max_ttl": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_ttl": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"origin_request_policy_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"realtime_log_config_arn": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_headers_policy_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"smooth_streaming": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"target_origin_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"trusted_key_groups": schema.SingleNestedAttribute{
										Description:         "A list of key groups whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.",
										MarkdownDescription: "A list of key groups whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

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

									"trusted_signers": schema.SingleNestedAttribute{
										Description:         "A list of Amazon Web Services accounts whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.",
										MarkdownDescription: "A list of Amazon Web Services accounts whose public keys CloudFront can use to verify the signatures of signed URLs and signed cookies.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

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

									"viewer_protocol_policy": schema.StringAttribute{
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

							"default_root_object": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"is_ipv6_enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logging": schema.SingleNestedAttribute{
								Description:         "A complex type that controls whether access logs are written for the distribution.",
								MarkdownDescription: "A complex type that controls whether access logs are written for the distribution.",
								Attributes: map[string]schema.Attribute{
									"bucket": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"include_cookies": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"prefix": schema.StringAttribute{
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

							"origin_groups": schema.SingleNestedAttribute{
								Description:         "A complex data type for the origin groups specified for a distribution.",
								MarkdownDescription: "A complex data type for the origin groups specified for a distribution.",
								Attributes: map[string]schema.Attribute{
									"items": schema.ListNestedAttribute{
										Description:         "List of origin groups for a distribution.",
										MarkdownDescription: "List of origin groups for a distribution.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"failover_criteria": schema.SingleNestedAttribute{
													Description:         "A complex data type that includes information about the failover criteria for an origin group, including the status codes for which CloudFront will failover from the primary origin to the second origin.",
													MarkdownDescription: "A complex data type that includes information about the failover criteria for an origin group, including the status codes for which CloudFront will failover from the primary origin to the second origin.",
													Attributes: map[string]schema.Attribute{
														"status_codes": schema.SingleNestedAttribute{
															Description:         "A complex data type for the status codes that you specify that, when returned by a primary origin, trigger CloudFront to failover to a second origin.",
															MarkdownDescription: "A complex data type for the status codes that you specify that, when returned by a primary origin, trigger CloudFront to failover to a second origin.",
															Attributes: map[string]schema.Attribute{
																"items": schema.ListAttribute{
																	Description:         "List of status codes for origin failover.",
																	MarkdownDescription: "List of status codes for origin failover.",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"members": schema.SingleNestedAttribute{
													Description:         "A complex data type for the origins included in an origin group.",
													MarkdownDescription: "A complex data type for the origins included in an origin group.",
													Attributes: map[string]schema.Attribute{
														"items": schema.ListNestedAttribute{
															Description:         "List of origins in an origin group.",
															MarkdownDescription: "List of origins in an origin group.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"origin_id": schema.StringAttribute{
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

							"origins": schema.SingleNestedAttribute{
								Description:         "Contains information about the origins for this distribution.",
								MarkdownDescription: "Contains information about the origins for this distribution.",
								Attributes: map[string]schema.Attribute{
									"items": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"connection_attempts": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"connection_timeout": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"custom_headers": schema.SingleNestedAttribute{
													Description:         "A complex type that contains the list of Custom Headers for each origin.",
													MarkdownDescription: "A complex type that contains the list of Custom Headers for each origin.",
													Attributes: map[string]schema.Attribute{
														"items": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"header_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"header_value": schema.StringAttribute{
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

												"custom_origin_config": schema.SingleNestedAttribute{
													Description:         "A custom origin. A custom origin is any origin that is not an Amazon S3 bucket, with one exception. An Amazon S3 bucket that is configured with static website hosting (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) is a custom origin.",
													MarkdownDescription: "A custom origin. A custom origin is any origin that is not an Amazon S3 bucket, with one exception. An Amazon S3 bucket that is configured with static website hosting (https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html) is a custom origin.",
													Attributes: map[string]schema.Attribute{
														"http_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_s_port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin_keepalive_timeout": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin_protocol_policy": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin_read_timeout": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin_ssl_protocols": schema.SingleNestedAttribute{
															Description:         "A complex type that contains information about the SSL/TLS protocols that CloudFront can use when establishing an HTTPS connection with your origin.",
															MarkdownDescription: "A complex type that contains information about the SSL/TLS protocols that CloudFront can use when establishing an HTTPS connection with your origin.",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"domain_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"origin_access_control_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"origin_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"origin_shield": schema.SingleNestedAttribute{
													Description:         "CloudFront Origin Shield. Using Origin Shield can help reduce the load on your origin. For more information, see Using Origin Shield (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html) in the Amazon CloudFront Developer Guide.",
													MarkdownDescription: "CloudFront Origin Shield. Using Origin Shield can help reduce the load on your origin. For more information, see Using Origin Shield (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/origin-shield.html) in the Amazon CloudFront Developer Guide.",
													Attributes: map[string]schema.Attribute{
														"enabled": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"origin_shield_region": schema.StringAttribute{
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

												"s3_origin_config": schema.SingleNestedAttribute{
													Description:         "A complex type that contains information about the Amazon S3 origin. If the origin is a custom origin or an S3 bucket that is configured as a website endpoint, use the CustomOriginConfig element instead.",
													MarkdownDescription: "A complex type that contains information about the Amazon S3 origin. If the origin is a custom origin or an S3 bucket that is configured as a website endpoint, use the CustomOriginConfig element instead.",
													Attributes: map[string]schema.Attribute{
														"origin_access_identity": schema.StringAttribute{
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

							"price_class": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"restrictions": schema.SingleNestedAttribute{
								Description:         "A complex type that identifies ways in which you want to restrict distribution of your content.",
								MarkdownDescription: "A complex type that identifies ways in which you want to restrict distribution of your content.",
								Attributes: map[string]schema.Attribute{
									"geo_restriction": schema.SingleNestedAttribute{
										Description:         "A complex type that controls the countries in which your content is distributed. CloudFront determines the location of your users using MaxMind GeoIP databases.",
										MarkdownDescription: "A complex type that controls the countries in which your content is distributed. CloudFront determines the location of your users using MaxMind GeoIP databases.",
										Attributes: map[string]schema.Attribute{
											"items": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"restriction_type": schema.StringAttribute{
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

							"staging": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"viewer_certificate": schema.SingleNestedAttribute{
								Description:         "A complex type that determines the distribution's SSL/TLS configuration for communicating with viewers. If the distribution doesn't use Aliases (also known as alternate domain names or CNAMEs)—that is, if the distribution uses the CloudFront domain name such as d111111abcdef8.cloudfront.net—set CloudFrontDefaultCertificate to true and leave all other fields empty. If the distribution uses Aliases (alternate domain names or CNAMEs), use the fields in this type to specify the following settings: * Which viewers the distribution accepts HTTPS connections from: only viewers that support server name indication (SNI) (https://en.wikipedia.org/wiki/Server_Name_Indication) (recommended), or all viewers including those that don't support SNI. To accept HTTPS connections from only viewers that support SNI, set SSLSupportMethod to sni-only. This is recommended. Most browsers and clients support SNI. To accept HTTPS connections from all viewers, including those that don't support SNI, set SSLSupportMethod to vip. This is not recommended, and results in additional monthly charges from CloudFront. * The minimum SSL/TLS protocol version that the distribution can use to communicate with viewers. To specify a minimum version, choose a value for MinimumProtocolVersion. For more information, see Security Policy (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValues-security-policy) in the Amazon CloudFront Developer Guide. * The location of the SSL/TLS certificate, Certificate Manager (ACM) (https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html) (recommended) or Identity and Access Management (IAM) (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html). You specify the location by setting a value in one of the following fields (not both): ACMCertificateArn IAMCertificateId All distributions support HTTPS connections from viewers. To require viewers to use HTTPS only, or to redirect them from HTTP to HTTPS, use ViewerProtocolPolicy in the CacheBehavior or DefaultCacheBehavior. To specify how CloudFront should use SSL/TLS to communicate with your custom origin, use CustomOriginConfig. For more information, see Using HTTPS with CloudFront (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https.html) and Using Alternate Domain Names and HTTPS (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-alternate-domain-names.html) in the Amazon CloudFront Developer Guide.",
								MarkdownDescription: "A complex type that determines the distribution's SSL/TLS configuration for communicating with viewers. If the distribution doesn't use Aliases (also known as alternate domain names or CNAMEs)—that is, if the distribution uses the CloudFront domain name such as d111111abcdef8.cloudfront.net—set CloudFrontDefaultCertificate to true and leave all other fields empty. If the distribution uses Aliases (alternate domain names or CNAMEs), use the fields in this type to specify the following settings: * Which viewers the distribution accepts HTTPS connections from: only viewers that support server name indication (SNI) (https://en.wikipedia.org/wiki/Server_Name_Indication) (recommended), or all viewers including those that don't support SNI. To accept HTTPS connections from only viewers that support SNI, set SSLSupportMethod to sni-only. This is recommended. Most browsers and clients support SNI. To accept HTTPS connections from all viewers, including those that don't support SNI, set SSLSupportMethod to vip. This is not recommended, and results in additional monthly charges from CloudFront. * The minimum SSL/TLS protocol version that the distribution can use to communicate with viewers. To specify a minimum version, choose a value for MinimumProtocolVersion. For more information, see Security Policy (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-web-values-specify.html#DownloadDistValues-security-policy) in the Amazon CloudFront Developer Guide. * The location of the SSL/TLS certificate, Certificate Manager (ACM) (https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html) (recommended) or Identity and Access Management (IAM) (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html). You specify the location by setting a value in one of the following fields (not both): ACMCertificateArn IAMCertificateId All distributions support HTTPS connections from viewers. To require viewers to use HTTPS only, or to redirect them from HTTP to HTTPS, use ViewerProtocolPolicy in the CacheBehavior or DefaultCacheBehavior. To specify how CloudFront should use SSL/TLS to communicate with your custom origin, use CustomOriginConfig. For more information, see Using HTTPS with CloudFront (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https.html) and Using Alternate Domain Names and HTTPS (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-https-alternate-domain-names.html) in the Amazon CloudFront Developer Guide.",
								Attributes: map[string]schema.Attribute{
									"acm_certificate_arn": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"acm_certificate_ref": schema.SingleNestedAttribute{
										Description:         "Reference field for ACMCertificateARN",
										MarkdownDescription: "Reference field for ACMCertificateARN",
										Attributes: map[string]schema.Attribute{
											"from": schema.SingleNestedAttribute{
												Description:         "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
												MarkdownDescription: "AWSResourceReference provides all the values necessary to reference another k8s resource for finding the identifier(Id/ARN/Name)",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
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

									"certificate": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"certificate_source": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cloud_front_default_certificate": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"iam_certificate_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"minimum_protocol_version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ssl_support_method": schema.StringAttribute{
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

							"web_aclid": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
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

func (r *CloudfrontServicesK8SAwsDistributionV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloudfront_services_k8s_aws_distribution_v1alpha1_manifest")

	var model CloudfrontServicesK8SAwsDistributionV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cloudfront.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Distribution")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
