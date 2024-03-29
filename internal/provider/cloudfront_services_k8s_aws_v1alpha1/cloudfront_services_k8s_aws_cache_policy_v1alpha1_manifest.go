/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package cloudfront_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &CloudfrontServicesK8SAwsCachePolicyV1Alpha1Manifest{}
)

func NewCloudfrontServicesK8SAwsCachePolicyV1Alpha1Manifest() datasource.DataSource {
	return &CloudfrontServicesK8SAwsCachePolicyV1Alpha1Manifest{}
}

type CloudfrontServicesK8SAwsCachePolicyV1Alpha1Manifest struct{}

type CloudfrontServicesK8SAwsCachePolicyV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		CachePolicyConfig *struct {
			Comment                                  *string `tfsdk:"comment" json:"comment,omitempty"`
			DefaultTTL                               *int64  `tfsdk:"default_ttl" json:"defaultTTL,omitempty"`
			MaxTTL                                   *int64  `tfsdk:"max_ttl" json:"maxTTL,omitempty"`
			MinTTL                                   *int64  `tfsdk:"min_ttl" json:"minTTL,omitempty"`
			Name                                     *string `tfsdk:"name" json:"name,omitempty"`
			ParametersInCacheKeyAndForwardedToOrigin *struct {
				CookiesConfig *struct {
					CookieBehavior *string `tfsdk:"cookie_behavior" json:"cookieBehavior,omitempty"`
					Cookies        *struct {
						Items *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"cookies" json:"cookies,omitempty"`
				} `tfsdk:"cookies_config" json:"cookiesConfig,omitempty"`
				EnableAcceptEncodingBrotli *bool `tfsdk:"enable_accept_encoding_brotli" json:"enableAcceptEncodingBrotli,omitempty"`
				EnableAcceptEncodingGzip   *bool `tfsdk:"enable_accept_encoding_gzip" json:"enableAcceptEncodingGzip,omitempty"`
				HeadersConfig              *struct {
					HeaderBehavior *string `tfsdk:"header_behavior" json:"headerBehavior,omitempty"`
					Headers        *struct {
						Items *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
				} `tfsdk:"headers_config" json:"headersConfig,omitempty"`
				QueryStringsConfig *struct {
					QueryStringBehavior *string `tfsdk:"query_string_behavior" json:"queryStringBehavior,omitempty"`
					QueryStrings        *struct {
						Items *[]string `tfsdk:"items" json:"items,omitempty"`
					} `tfsdk:"query_strings" json:"queryStrings,omitempty"`
				} `tfsdk:"query_strings_config" json:"queryStringsConfig,omitempty"`
			} `tfsdk:"parameters_in_cache_key_and_forwarded_to_origin" json:"parametersInCacheKeyAndForwardedToOrigin,omitempty"`
		} `tfsdk:"cache_policy_config" json:"cachePolicyConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudfrontServicesK8SAwsCachePolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest"
}

func (r *CloudfrontServicesK8SAwsCachePolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CachePolicy is the Schema for the CachePolicies API",
		MarkdownDescription: "CachePolicy is the Schema for the CachePolicies API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "CachePolicySpec defines the desired state of CachePolicy.A cache policy.When it's attached to a cache behavior, the cache policy determines the following:   * The values that CloudFront includes in the cache key. These values can   include HTTP headers, cookies, and URL query strings. CloudFront uses   the cache key to find an object in its cache that it can return to the   viewer.   * The default, minimum, and maximum time to live (TTL) values that you   want objects to stay in the CloudFront cache.The headers, cookies, and query strings that are included in the cache keyare also included in requests that CloudFront sends to the origin. CloudFrontsends a request when it can't find a valid object in its cache that matchesthe request's cache key. If you want to send values to the origin but notinclude them in the cache key, use OriginRequestPolicy.",
				MarkdownDescription: "CachePolicySpec defines the desired state of CachePolicy.A cache policy.When it's attached to a cache behavior, the cache policy determines the following:   * The values that CloudFront includes in the cache key. These values can   include HTTP headers, cookies, and URL query strings. CloudFront uses   the cache key to find an object in its cache that it can return to the   viewer.   * The default, minimum, and maximum time to live (TTL) values that you   want objects to stay in the CloudFront cache.The headers, cookies, and query strings that are included in the cache keyare also included in requests that CloudFront sends to the origin. CloudFrontsends a request when it can't find a valid object in its cache that matchesthe request's cache key. If you want to send values to the origin but notinclude them in the cache key, use OriginRequestPolicy.",
				Attributes: map[string]schema.Attribute{
					"cache_policy_config": schema.SingleNestedAttribute{
						Description:         "A cache policy configuration.",
						MarkdownDescription: "A cache policy configuration.",
						Attributes: map[string]schema.Attribute{
							"comment": schema.StringAttribute{
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

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"parameters_in_cache_key_and_forwarded_to_origin": schema.SingleNestedAttribute{
								Description:         "This object determines the values that CloudFront includes in the cache key.These values can include HTTP headers, cookies, and URL query strings. CloudFrontuses the cache key to find an object in its cache that it can return to theviewer.The headers, cookies, and query strings that are included in the cache keyare also included in requests that CloudFront sends to the origin. CloudFrontsends a request when it can't find an object in its cache that matches therequest's cache key. If you want to send values to the origin but not includethem in the cache key, use OriginRequestPolicy.",
								MarkdownDescription: "This object determines the values that CloudFront includes in the cache key.These values can include HTTP headers, cookies, and URL query strings. CloudFrontuses the cache key to find an object in its cache that it can return to theviewer.The headers, cookies, and query strings that are included in the cache keyare also included in requests that CloudFront sends to the origin. CloudFrontsends a request when it can't find an object in its cache that matches therequest's cache key. If you want to send values to the origin but not includethem in the cache key, use OriginRequestPolicy.",
								Attributes: map[string]schema.Attribute{
									"cookies_config": schema.SingleNestedAttribute{
										Description:         "An object that determines whether any cookies in viewer requests (and ifso, which cookies) are included in the cache key and in requests that CloudFrontsends to the origin.",
										MarkdownDescription: "An object that determines whether any cookies in viewer requests (and ifso, which cookies) are included in the cache key and in requests that CloudFrontsends to the origin.",
										Attributes: map[string]schema.Attribute{
											"cookie_behavior": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cookies": schema.SingleNestedAttribute{
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

									"enable_accept_encoding_brotli": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_accept_encoding_gzip": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"headers_config": schema.SingleNestedAttribute{
										Description:         "An object that determines whether any HTTP headers (and if so, which headers)are included in the cache key and in requests that CloudFront sends to theorigin.",
										MarkdownDescription: "An object that determines whether any HTTP headers (and if so, which headers)are included in the cache key and in requests that CloudFront sends to theorigin.",
										Attributes: map[string]schema.Attribute{
											"header_behavior": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_strings_config": schema.SingleNestedAttribute{
										Description:         "An object that determines whether any URL query strings in viewer requests(and if so, which query strings) are included in the cache key and in requeststhat CloudFront sends to the origin.",
										MarkdownDescription: "An object that determines whether any URL query strings in viewer requests(and if so, which query strings) are included in the cache key and in requeststhat CloudFront sends to the origin.",
										Attributes: map[string]schema.Attribute{
											"query_string_behavior": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"query_strings": schema.SingleNestedAttribute{
												Description:         "Contains a list of query string names.",
												MarkdownDescription: "Contains a list of query string names.",
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

func (r *CloudfrontServicesK8SAwsCachePolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest")

	var model CloudfrontServicesK8SAwsCachePolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("cloudfront.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("CachePolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
