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
	_ datasource.DataSource = &CloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1Manifest{}
)

func NewCloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1Manifest() datasource.DataSource {
	return &CloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1Manifest{}
}

type CloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1Manifest struct{}

type CloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1ManifestData struct {
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
		OriginRequestPolicyConfig *struct {
			Comment       *string `tfsdk:"comment" json:"comment,omitempty"`
			CookiesConfig *struct {
				CookieBehavior *string `tfsdk:"cookie_behavior" json:"cookieBehavior,omitempty"`
				Cookies        *struct {
					Items *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"cookies" json:"cookies,omitempty"`
			} `tfsdk:"cookies_config" json:"cookiesConfig,omitempty"`
			HeadersConfig *struct {
				HeaderBehavior *string `tfsdk:"header_behavior" json:"headerBehavior,omitempty"`
				Headers        *struct {
					Items *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
			} `tfsdk:"headers_config" json:"headersConfig,omitempty"`
			Name               *string `tfsdk:"name" json:"name,omitempty"`
			QueryStringsConfig *struct {
				QueryStringBehavior *string `tfsdk:"query_string_behavior" json:"queryStringBehavior,omitempty"`
				QueryStrings        *struct {
					Items *[]string `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"query_strings" json:"queryStrings,omitempty"`
			} `tfsdk:"query_strings_config" json:"queryStringsConfig,omitempty"`
		} `tfsdk:"origin_request_policy_config" json:"originRequestPolicyConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest"
}

func (r *CloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OriginRequestPolicy is the Schema for the OriginRequestPolicies API",
		MarkdownDescription: "OriginRequestPolicy is the Schema for the OriginRequestPolicies API",
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
				Description:         "OriginRequestPolicySpec defines the desired state of OriginRequestPolicy.An origin request policy.When it's attached to a cache behavior, the origin request policy determinesthe values that CloudFront includes in requests that it sends to the origin.Each request that CloudFront sends to the origin includes the following:   * The request body and the URL path (without the domain name) from the   viewer request.   * The headers that CloudFront automatically includes in every origin request,   including Host, User-Agent, and X-Amz-Cf-Id.   * All HTTP headers, cookies, and URL query strings that are specified   in the cache policy or the origin request policy. These can include items   from the viewer request and, in the case of headers, additional ones that   are added by CloudFront.CloudFront sends a request when it can't find an object in its cache thatmatches the request. If you want to send values to the origin and also includethem in the cache key, use CachePolicy.",
				MarkdownDescription: "OriginRequestPolicySpec defines the desired state of OriginRequestPolicy.An origin request policy.When it's attached to a cache behavior, the origin request policy determinesthe values that CloudFront includes in requests that it sends to the origin.Each request that CloudFront sends to the origin includes the following:   * The request body and the URL path (without the domain name) from the   viewer request.   * The headers that CloudFront automatically includes in every origin request,   including Host, User-Agent, and X-Amz-Cf-Id.   * All HTTP headers, cookies, and URL query strings that are specified   in the cache policy or the origin request policy. These can include items   from the viewer request and, in the case of headers, additional ones that   are added by CloudFront.CloudFront sends a request when it can't find an object in its cache thatmatches the request. If you want to send values to the origin and also includethem in the cache key, use CachePolicy.",
				Attributes: map[string]schema.Attribute{
					"origin_request_policy_config": schema.SingleNestedAttribute{
						Description:         "An origin request policy configuration.",
						MarkdownDescription: "An origin request policy configuration.",
						Attributes: map[string]schema.Attribute{
							"comment": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cookies_config": schema.SingleNestedAttribute{
								Description:         "An object that determines whether any cookies in viewer requests (and ifso, which cookies) are included in requests that CloudFront sends to theorigin.",
								MarkdownDescription: "An object that determines whether any cookies in viewer requests (and ifso, which cookies) are included in requests that CloudFront sends to theorigin.",
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

							"headers_config": schema.SingleNestedAttribute{
								Description:         "An object that determines whether any HTTP headers (and if so, which headers)are included in requests that CloudFront sends to the origin.",
								MarkdownDescription: "An object that determines whether any HTTP headers (and if so, which headers)are included in requests that CloudFront sends to the origin.",
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

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"query_strings_config": schema.SingleNestedAttribute{
								Description:         "An object that determines whether any URL query strings in viewer requests(and if so, which query strings) are included in requests that CloudFrontsends to the origin.",
								MarkdownDescription: "An object that determines whether any URL query strings in viewer requests(and if so, which query strings) are included in requests that CloudFrontsends to the origin.",
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

func (r *CloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_cloudfront_services_k8s_aws_origin_request_policy_v1alpha1_manifest")

	var model CloudfrontServicesK8SAwsOriginRequestPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("cloudfront.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("OriginRequestPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
