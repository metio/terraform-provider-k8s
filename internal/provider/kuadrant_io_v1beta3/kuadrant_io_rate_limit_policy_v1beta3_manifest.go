/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuadrant_io_v1beta3

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KuadrantIoRateLimitPolicyV1Beta3Manifest{}
)

func NewKuadrantIoRateLimitPolicyV1Beta3Manifest() datasource.DataSource {
	return &KuadrantIoRateLimitPolicyV1Beta3Manifest{}
}

type KuadrantIoRateLimitPolicyV1Beta3Manifest struct{}

type KuadrantIoRateLimitPolicyV1Beta3ManifestData struct {
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
		Defaults *struct {
			Limits *struct {
				Counters *[]struct {
					Expression *string `tfsdk:"expression" json:"expression,omitempty"`
				} `tfsdk:"counters" json:"counters,omitempty"`
				Rates *[]struct {
					Limit  *int64  `tfsdk:"limit" json:"limit,omitempty"`
					Window *string `tfsdk:"window" json:"window,omitempty"`
				} `tfsdk:"rates" json:"rates,omitempty"`
				When *[]struct {
					Predicate *string `tfsdk:"predicate" json:"predicate,omitempty"`
				} `tfsdk:"when" json:"when,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
			When     *[]struct {
				Predicate *string `tfsdk:"predicate" json:"predicate,omitempty"`
			} `tfsdk:"when" json:"when,omitempty"`
		} `tfsdk:"defaults" json:"defaults,omitempty"`
		Limits *struct {
			Counters *[]struct {
				Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			} `tfsdk:"counters" json:"counters,omitempty"`
			Rates *[]struct {
				Limit  *int64  `tfsdk:"limit" json:"limit,omitempty"`
				Window *string `tfsdk:"window" json:"window,omitempty"`
			} `tfsdk:"rates" json:"rates,omitempty"`
			When *[]struct {
				Predicate *string `tfsdk:"predicate" json:"predicate,omitempty"`
			} `tfsdk:"when" json:"when,omitempty"`
		} `tfsdk:"limits" json:"limits,omitempty"`
		Overrides *struct {
			Limits *struct {
				Counters *[]struct {
					Expression *string `tfsdk:"expression" json:"expression,omitempty"`
				} `tfsdk:"counters" json:"counters,omitempty"`
				Rates *[]struct {
					Limit  *int64  `tfsdk:"limit" json:"limit,omitempty"`
					Window *string `tfsdk:"window" json:"window,omitempty"`
				} `tfsdk:"rates" json:"rates,omitempty"`
				When *[]struct {
					Predicate *string `tfsdk:"predicate" json:"predicate,omitempty"`
				} `tfsdk:"when" json:"when,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
			When     *[]struct {
				Predicate *string `tfsdk:"predicate" json:"predicate,omitempty"`
			} `tfsdk:"when" json:"when,omitempty"`
		} `tfsdk:"overrides" json:"overrides,omitempty"`
		TargetRef *struct {
			Group       *string `tfsdk:"group" json:"group,omitempty"`
			Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
			Name        *string `tfsdk:"name" json:"name,omitempty"`
			SectionName *string `tfsdk:"section_name" json:"sectionName,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		When *[]struct {
			Predicate *string `tfsdk:"predicate" json:"predicate,omitempty"`
		} `tfsdk:"when" json:"when,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KuadrantIoRateLimitPolicyV1Beta3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuadrant_io_rate_limit_policy_v1beta3_manifest"
}

func (r *KuadrantIoRateLimitPolicyV1Beta3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RateLimitPolicy enables rate limiting for service workloads in a Gateway API network",
		MarkdownDescription: "RateLimitPolicy enables rate limiting for service workloads in a Gateway API network",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"defaults": schema.SingleNestedAttribute{
						Description:         "Rules to apply as defaults. Can be overridden by more specific policiy rules lower in the hierarchy and by less specific policy overrides. Use one of: defaults, overrides, or bare set of policy rules (implicit defaults).",
						MarkdownDescription: "Rules to apply as defaults. Can be overridden by more specific policiy rules lower in the hierarchy and by less specific policy overrides. Use one of: defaults, overrides, or bare set of policy rules (implicit defaults).",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "Limits holds the struct of limits indexed by a unique name",
								MarkdownDescription: "Limits holds the struct of limits indexed by a unique name",
								Attributes: map[string]schema.Attribute{
									"counters": schema.ListNestedAttribute{
										Description:         "Counters defines additional rate limit counters based on CEL expressions which can reference well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
										MarkdownDescription: "Counters defines additional rate limit counters based on CEL expressions which can reference well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"expression": schema.StringAttribute{
													Description:         "Expression defines one CEL expression Expression can use well known attributes Attributes: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes Well-known selectors: https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors They are named by a dot-separated path (e.g. request.path) Example: 'request.path' -> The path portion of the URL",
													MarkdownDescription: "Expression defines one CEL expression Expression can use well known attributes Attributes: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes Well-known selectors: https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors They are named by a dot-separated path (e.g. request.path) Example: 'request.path' -> The path portion of the URL",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"rates": schema.ListNestedAttribute{
										Description:         "Rates holds the list of limit rates",
										MarkdownDescription: "Rates holds the list of limit rates",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"limit": schema.Int64Attribute{
													Description:         "Limit defines the max value allowed for a given period of time",
													MarkdownDescription: "Limit defines the max value allowed for a given period of time",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"window": schema.StringAttribute{
													Description:         "Window defines the time period for which the Limit specified above applies.",
													MarkdownDescription: "Window defines the time period for which the Limit specified above applies.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"when": schema.ListNestedAttribute{
										Description:         "When holds a list of 'limit-level' 'Predicate's Called also 'soft' conditions as route selectors must also match",
										MarkdownDescription: "When holds a list of 'limit-level' 'Predicate's Called also 'soft' conditions as route selectors must also match",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"predicate": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
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

							"strategy": schema.StringAttribute{
								Description:         "Strategy defines the merge strategy to apply when merging this policy with other policies.",
								MarkdownDescription: "Strategy defines the merge strategy to apply when merging this policy with other policies.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("atomic", "merge"),
								},
							},

							"when": schema.ListNestedAttribute{
								Description:         "Overall conditions for the policy to be enforced. If omitted, the policy will be enforced at all requests to the protected routes. If present, all conditions must match for the policy to be enforced.",
								MarkdownDescription: "Overall conditions for the policy to be enforced. If omitted, the policy will be enforced at all requests to the protected routes. If present, all conditions must match for the policy to be enforced.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"predicate": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
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

					"limits": schema.SingleNestedAttribute{
						Description:         "Limits holds the struct of limits indexed by a unique name",
						MarkdownDescription: "Limits holds the struct of limits indexed by a unique name",
						Attributes: map[string]schema.Attribute{
							"counters": schema.ListNestedAttribute{
								Description:         "Counters defines additional rate limit counters based on CEL expressions which can reference well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
								MarkdownDescription: "Counters defines additional rate limit counters based on CEL expressions which can reference well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"expression": schema.StringAttribute{
											Description:         "Expression defines one CEL expression Expression can use well known attributes Attributes: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes Well-known selectors: https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors They are named by a dot-separated path (e.g. request.path) Example: 'request.path' -> The path portion of the URL",
											MarkdownDescription: "Expression defines one CEL expression Expression can use well known attributes Attributes: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes Well-known selectors: https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors They are named by a dot-separated path (e.g. request.path) Example: 'request.path' -> The path portion of the URL",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rates": schema.ListNestedAttribute{
								Description:         "Rates holds the list of limit rates",
								MarkdownDescription: "Rates holds the list of limit rates",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"limit": schema.Int64Attribute{
											Description:         "Limit defines the max value allowed for a given period of time",
											MarkdownDescription: "Limit defines the max value allowed for a given period of time",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"window": schema.StringAttribute{
											Description:         "Window defines the time period for which the Limit specified above applies.",
											MarkdownDescription: "Window defines the time period for which the Limit specified above applies.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"when": schema.ListNestedAttribute{
								Description:         "When holds a list of 'limit-level' 'Predicate's Called also 'soft' conditions as route selectors must also match",
								MarkdownDescription: "When holds a list of 'limit-level' 'Predicate's Called also 'soft' conditions as route selectors must also match",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"predicate": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
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

					"overrides": schema.SingleNestedAttribute{
						Description:         "Rules to apply as overrides. Override all policy rules lower in the hierarchy. Can be overridden by less specific policy overrides. Use one of: defaults, overrides, or bare set of policy rules (implicit defaults).",
						MarkdownDescription: "Rules to apply as overrides. Override all policy rules lower in the hierarchy. Can be overridden by less specific policy overrides. Use one of: defaults, overrides, or bare set of policy rules (implicit defaults).",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "Limits holds the struct of limits indexed by a unique name",
								MarkdownDescription: "Limits holds the struct of limits indexed by a unique name",
								Attributes: map[string]schema.Attribute{
									"counters": schema.ListNestedAttribute{
										Description:         "Counters defines additional rate limit counters based on CEL expressions which can reference well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
										MarkdownDescription: "Counters defines additional rate limit counters based on CEL expressions which can reference well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"expression": schema.StringAttribute{
													Description:         "Expression defines one CEL expression Expression can use well known attributes Attributes: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes Well-known selectors: https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors They are named by a dot-separated path (e.g. request.path) Example: 'request.path' -> The path portion of the URL",
													MarkdownDescription: "Expression defines one CEL expression Expression can use well known attributes Attributes: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/advanced/attributes Well-known selectors: https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors They are named by a dot-separated path (e.g. request.path) Example: 'request.path' -> The path portion of the URL",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"rates": schema.ListNestedAttribute{
										Description:         "Rates holds the list of limit rates",
										MarkdownDescription: "Rates holds the list of limit rates",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"limit": schema.Int64Attribute{
													Description:         "Limit defines the max value allowed for a given period of time",
													MarkdownDescription: "Limit defines the max value allowed for a given period of time",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"window": schema.StringAttribute{
													Description:         "Window defines the time period for which the Limit specified above applies.",
													MarkdownDescription: "Window defines the time period for which the Limit specified above applies.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]{1,5}(h|m|s|ms)){1,4}$`), ""),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"when": schema.ListNestedAttribute{
										Description:         "When holds a list of 'limit-level' 'Predicate's Called also 'soft' conditions as route selectors must also match",
										MarkdownDescription: "When holds a list of 'limit-level' 'Predicate's Called also 'soft' conditions as route selectors must also match",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"predicate": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
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

							"strategy": schema.StringAttribute{
								Description:         "Strategy defines the merge strategy to apply when merging this policy with other policies.",
								MarkdownDescription: "Strategy defines the merge strategy to apply when merging this policy with other policies.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("atomic", "merge"),
								},
							},

							"when": schema.ListNestedAttribute{
								Description:         "Overall conditions for the policy to be enforced. If omitted, the policy will be enforced at all requests to the protected routes. If present, all conditions must match for the policy to be enforced.",
								MarkdownDescription: "Overall conditions for the policy to be enforced. If omitted, the policy will be enforced at all requests to the protected routes. If present, all conditions must match for the policy to be enforced.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"predicate": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
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
						Description:         "Reference to the object to which this policy applies.",
						MarkdownDescription: "Reference to the object to which this policy applies.",
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

							"section_name": schema.StringAttribute{
								Description:         "SectionName is the name of a section within the target resource. When unspecified, this targetRef targets the entire resource. In the following resources, SectionName is interpreted as the following: * Gateway: Listener name * HTTPRoute: HTTPRouteRule name * Service: Port name If a SectionName is specified, but does not exist on the targeted object, the Policy must fail to attach, and the policy implementation should record a 'ResolvedRefs' or similar Condition in the Policy's status.",
								MarkdownDescription: "SectionName is the name of a section within the target resource. When unspecified, this targetRef targets the entire resource. In the following resources, SectionName is interpreted as the following: * Gateway: Listener name * HTTPRoute: HTTPRouteRule name * Service: Port name If a SectionName is specified, but does not exist on the targeted object, the Policy must fail to attach, and the policy implementation should record a 'ResolvedRefs' or similar Condition in the Policy's status.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"when": schema.ListNestedAttribute{
						Description:         "Overall conditions for the policy to be enforced. If omitted, the policy will be enforced at all requests to the protected routes. If present, all conditions must match for the policy to be enforced.",
						MarkdownDescription: "Overall conditions for the policy to be enforced. If omitted, the policy will be enforced at all requests to the protected routes. If present, all conditions must match for the policy to be enforced.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"predicate": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
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

func (r *KuadrantIoRateLimitPolicyV1Beta3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuadrant_io_rate_limit_policy_v1beta3_manifest")

	var model KuadrantIoRateLimitPolicyV1Beta3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuadrant.io/v1beta3")
	model.Kind = pointer.String("RateLimitPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
