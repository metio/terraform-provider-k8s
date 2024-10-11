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
				Counters *[]string `tfsdk:"counters" json:"counters,omitempty"`
				Rates    *[]struct {
					Duration *int64  `tfsdk:"duration" json:"duration,omitempty"`
					Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
					Unit     *string `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"rates" json:"rates,omitempty"`
				When *[]struct {
					Operator *string `tfsdk:"operator" json:"operator,omitempty"`
					Selector *string `tfsdk:"selector" json:"selector,omitempty"`
					Value    *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"when" json:"when,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
		} `tfsdk:"defaults" json:"defaults,omitempty"`
		Limits *struct {
			Counters *[]string `tfsdk:"counters" json:"counters,omitempty"`
			Rates    *[]struct {
				Duration *int64  `tfsdk:"duration" json:"duration,omitempty"`
				Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
				Unit     *string `tfsdk:"unit" json:"unit,omitempty"`
			} `tfsdk:"rates" json:"rates,omitempty"`
			When *[]struct {
				Operator *string `tfsdk:"operator" json:"operator,omitempty"`
				Selector *string `tfsdk:"selector" json:"selector,omitempty"`
				Value    *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"when" json:"when,omitempty"`
		} `tfsdk:"limits" json:"limits,omitempty"`
		Overrides *struct {
			Limits *struct {
				Counters *[]string `tfsdk:"counters" json:"counters,omitempty"`
				Rates    *[]struct {
					Duration *int64  `tfsdk:"duration" json:"duration,omitempty"`
					Limit    *int64  `tfsdk:"limit" json:"limit,omitempty"`
					Unit     *string `tfsdk:"unit" json:"unit,omitempty"`
				} `tfsdk:"rates" json:"rates,omitempty"`
				When *[]struct {
					Operator *string `tfsdk:"operator" json:"operator,omitempty"`
					Selector *string `tfsdk:"selector" json:"selector,omitempty"`
					Value    *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"when" json:"when,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
		} `tfsdk:"overrides" json:"overrides,omitempty"`
		TargetRef *struct {
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Kind  *string `tfsdk:"kind" json:"kind,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
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
				Description:         "RateLimitPolicySpec defines the desired state of RateLimitPolicy",
				MarkdownDescription: "RateLimitPolicySpec defines the desired state of RateLimitPolicy",
				Attributes: map[string]schema.Attribute{
					"defaults": schema.SingleNestedAttribute{
						Description:         "Defaults define explicit default values for this policy and for policies inheriting this policy. Defaults are mutually exclusive with implicit defaults defined by RateLimitPolicyCommonSpec.",
						MarkdownDescription: "Defaults define explicit default values for this policy and for policies inheriting this policy. Defaults are mutually exclusive with implicit defaults defined by RateLimitPolicyCommonSpec.",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "Limits holds the struct of limits indexed by a unique name",
								MarkdownDescription: "Limits holds the struct of limits indexed by a unique name",
								Attributes: map[string]schema.Attribute{
									"counters": schema.ListAttribute{
										Description:         "Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
										MarkdownDescription: "Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
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
													Description:         "Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'",
													MarkdownDescription: "Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'",
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

									"when": schema.ListNestedAttribute{
										Description:         "When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match",
										MarkdownDescription: "When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"operator": schema.StringAttribute{
													Description:         "The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)",
													MarkdownDescription: "The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("eq", "neq", "startswith", "endswith", "incl", "excl", "matches"),
													},
												},

												"selector": schema.StringAttribute{
													Description:         "Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
													MarkdownDescription: "Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"limits": schema.SingleNestedAttribute{
						Description:         "Limits holds the struct of limits indexed by a unique name",
						MarkdownDescription: "Limits holds the struct of limits indexed by a unique name",
						Attributes: map[string]schema.Attribute{
							"counters": schema.ListAttribute{
								Description:         "Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
								MarkdownDescription: "Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
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
											Description:         "Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'",
											MarkdownDescription: "Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'",
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

							"when": schema.ListNestedAttribute{
								Description:         "When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match",
								MarkdownDescription: "When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"operator": schema.StringAttribute{
											Description:         "The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)",
											MarkdownDescription: "The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("eq", "neq", "startswith", "endswith", "incl", "excl", "matches"),
											},
										},

										"selector": schema.StringAttribute{
											Description:         "Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
											MarkdownDescription: "Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
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

					"overrides": schema.SingleNestedAttribute{
						Description:         "Overrides define override values for this policy and for policies inheriting this policy. Overrides are mutually exclusive with implicit defaults and explicit Defaults defined by RateLimitPolicyCommonSpec.",
						MarkdownDescription: "Overrides define override values for this policy and for policies inheriting this policy. Overrides are mutually exclusive with implicit defaults and explicit Defaults defined by RateLimitPolicyCommonSpec.",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "Limits holds the struct of limits indexed by a unique name",
								MarkdownDescription: "Limits holds the struct of limits indexed by a unique name",
								Attributes: map[string]schema.Attribute{
									"counters": schema.ListAttribute{
										Description:         "Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
										MarkdownDescription: "Counters defines additional rate limit counters based on context qualifiers and well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
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
													Description:         "Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'",
													MarkdownDescription: "Duration defines the time uni Possible values are: 'second', 'minute', 'hour', 'day'",
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

									"when": schema.ListNestedAttribute{
										Description:         "When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match",
										MarkdownDescription: "When holds the list of conditions for the policy to be enforced. Called also 'soft' conditions as route selectors must also match",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"operator": schema.StringAttribute{
													Description:         "The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)",
													MarkdownDescription: "The binary operator to be applied to the content fetched from the selector Possible values are: 'eq' (equal to), 'neq' (not equal to)",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("eq", "neq", "startswith", "endswith", "incl", "excl", "matches"),
													},
												},

												"selector": schema.StringAttribute{
													Description:         "Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
													MarkdownDescription: "Selector defines one item from the well known selectors TODO Document properly 'Well-known selector' https://github.com/Kuadrant/architecture/blob/main/rfcs/0001-rlp-v2.md#well-known-selectors",
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
