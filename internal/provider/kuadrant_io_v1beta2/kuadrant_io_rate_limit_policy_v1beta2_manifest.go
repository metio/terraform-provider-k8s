/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuadrant_io_v1beta2

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KuadrantIoRateLimitPolicyV1Beta2Manifest{}
)

func NewKuadrantIoRateLimitPolicyV1Beta2Manifest() datasource.DataSource {
	return &KuadrantIoRateLimitPolicyV1Beta2Manifest{}
}

type KuadrantIoRateLimitPolicyV1Beta2Manifest struct{}

type KuadrantIoRateLimitPolicyV1Beta2ManifestData struct {
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
		Limits *struct {
			Counters *[]string `tfsdk:"counters" json:"counters,omitempty"`
			Rates    *[]struct {
				Duration *int64  `tfsdk:"duration" json:"duration,omitempty"`
				Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
				Unit     *string `tfsdk:"unit" json:"unit,omitempty"`
			} `tfsdk:"rates" json:"rates,omitempty"`
			RouteSelectors *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Matches   *[]struct {
					Headers *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Method *string `tfsdk:"method" json:"method,omitempty"`
					Path   *struct {
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"path" json:"path,omitempty"`
					QueryParams *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Type  *string `tfsdk:"type" json:"type,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"query_params" json:"queryParams,omitempty"`
				} `tfsdk:"matches" json:"matches,omitempty"`
			} `tfsdk:"route_selectors" json:"routeSelectors,omitempty"`
			When *[]struct {
				Operator *string `tfsdk:"operator" json:"operator,omitempty"`
				Selector *string `tfsdk:"selector" json:"selector,omitempty"`
				Value    *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"when" json:"when,omitempty"`
		} `tfsdk:"limits" json:"limits,omitempty"`
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KuadrantIoRateLimitPolicyV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuadrant_io_rate_limit_policy_v1beta2_manifest"
}

func (r *KuadrantIoRateLimitPolicyV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RateLimitPolicy enables rate limiting for service workloads in a Gateway API network",
		MarkdownDescription: "RateLimitPolicy enables rate limiting for service workloads in a Gateway API network",
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
				Description:         "RateLimitPolicySpec defines the desired state of RateLimitPolicy",
				MarkdownDescription: "RateLimitPolicySpec defines the desired state of RateLimitPolicy",
				Attributes: map[string]schema.Attribute{
					"limits": schema.SingleNestedAttribute{
						Description:         "Limits holds the struct of limits indexed by a unique name",
						MarkdownDescription: "Limits holds the struct of limits indexed by a unique name",
						Attributes: map[string]schema.Attribute{
							"counters": schema.ListAttribute{
								Description:         "Counters defines additional rate limit counters based on context qualifiers and well known selectorsTODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
								MarkdownDescription: "Counters defines additional rate limit counters based on context qualifiers and well known selectorsTODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rates": schema.ListNestedAttribute{
								Description:         "Rates holds the list of limit rates",
								MarkdownDescription: "Rates holds the list of limit rates",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"duration": schema.Int64Attribute{
											Description:         "Duration defines the time period for which the Limit specified above applies.",
											MarkdownDescription: "Duration defines the time period for which the Limit specified above applies.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"limit": schema.Int64Attribute{
											Description:         "Limit defines the max value allowed for a given period of time",
											MarkdownDescription: "Limit defines the max value allowed for a given period of time",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"unit": schema.StringAttribute{
											Description:         "Duration defines the time uniPossible values are: 'second', 'minute', 'hour', 'day'",
											MarkdownDescription: "Duration defines the time uniPossible values are: 'second', 'minute', 'hour', 'day'",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("second", "minute", "hour", "day"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"route_selectors": schema.ListNestedAttribute{
								Description:         "RouteSelectors defines semantics for matching an HTTP request based on conditions",
								MarkdownDescription: "RouteSelectors defines semantics for matching an HTTP request based on conditions",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"hostnames": schema.ListAttribute{
											Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
											MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"matches": schema.ListNestedAttribute{
											Description:         "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
											MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"headers": schema.ListNestedAttribute{
														Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
														MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																	MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(256),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																	},
																},

																"type": schema.StringAttribute{
																	Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																	MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Exact", "RegularExpression"),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value of HTTP Header to be matched.",
																	MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(4096),
																	},
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": schema.StringAttribute{
														Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
														MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
														},
													},

													"path": schema.SingleNestedAttribute{
														Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
														MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
														Attributes: map[string]schema.Attribute{
															"type": schema.StringAttribute{
																Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
																},
															},

															"value": schema.StringAttribute{
																Description:         "Value of the HTTP path to match against.",
																MarkdownDescription: "Value of the HTTP path to match against.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(1024),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"query_params": schema.ListNestedAttribute{
														Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
														MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																	MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(256),
																		stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																	},
																},

																"type": schema.StringAttribute{
																	Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																	MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("Exact", "RegularExpression"),
																	},
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the value of HTTP query param to be matched.",
																	MarkdownDescription: "Value is the value of HTTP query param to be matched.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																		stringvalidator.LengthAtMost(1024),
																	},
																},
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"when": schema.ListNestedAttribute{
								Description:         "When holds the list of conditions for the policy to be enforced.Called also 'soft' conditions as route selectors must also match",
								MarkdownDescription: "When holds the list of conditions for the policy to be enforced.Called also 'soft' conditions as route selectors must also match",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"operator": schema.StringAttribute{
											Description:         "The binary operator to be applied to the content fetched from the selectorPossible values are: 'eq' (equal to), 'neq' (not equal to)",
											MarkdownDescription: "The binary operator to be applied to the content fetched from the selectorPossible values are: 'eq' (equal to), 'neq' (not equal to)",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("eq", "neq", "startswith", "endswith", "incl", "excl", "matches"),
											},
										},

										"selector": schema.StringAttribute{
											Description:         "Selector defines one item from the well known selectorsTODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
											MarkdownDescription: "Selector defines one item from the well known selectorsTODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(253),
											},
										},

										"value": schema.StringAttribute{
											Description:         "The value of reference for the comparison.",
											MarkdownDescription: "The value of reference for the comparison.",
											Required:            true,
											Optional:            false,
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef identifies an API object to apply policy to.",
						MarkdownDescription: "TargetRef identifies an API object to apply policy to.",
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

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the referent. When unspecified, the localnamespace is inferred. Even when policy targets a resource in a differentnamespace, it MUST only apply to traffic originating from the samenamespace as the policy.",
								MarkdownDescription: "Namespace is the namespace of the referent. When unspecified, the localnamespace is inferred. Even when policy targets a resource in a differentnamespace, it MUST only apply to traffic originating from the samenamespace as the policy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
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

func (r *KuadrantIoRateLimitPolicyV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuadrant_io_rate_limit_policy_v1beta2_manifest")

	var model KuadrantIoRateLimitPolicyV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("kuadrant.io/v1beta2")
	model.Kind = pointer.String("RateLimitPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
