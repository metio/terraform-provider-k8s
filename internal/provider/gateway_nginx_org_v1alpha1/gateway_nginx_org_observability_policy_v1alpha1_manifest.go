/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_nginx_org_v1alpha1

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
	_ datasource.DataSource = &GatewayNginxOrgObservabilityPolicyV1Alpha1Manifest{}
)

func NewGatewayNginxOrgObservabilityPolicyV1Alpha1Manifest() datasource.DataSource {
	return &GatewayNginxOrgObservabilityPolicyV1Alpha1Manifest{}
}

type GatewayNginxOrgObservabilityPolicyV1Alpha1Manifest struct{}

type GatewayNginxOrgObservabilityPolicyV1Alpha1ManifestData struct {
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
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		Tracing *struct {
			Context        *string `tfsdk:"context" json:"context,omitempty"`
			Ratio          *int64  `tfsdk:"ratio" json:"ratio,omitempty"`
			SpanAttributes *[]struct {
				Key   *string `tfsdk:"key" json:"key,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"span_attributes" json:"spanAttributes,omitempty"`
			SpanName *string `tfsdk:"span_name" json:"spanName,omitempty"`
			Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
		} `tfsdk:"tracing" json:"tracing,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewayNginxOrgObservabilityPolicyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_nginx_org_observability_policy_v1alpha1_manifest"
}

func (r *GatewayNginxOrgObservabilityPolicyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ObservabilityPolicy is a Direct Attached Policy. It provides a way to configure observability settings forthe NGINX Gateway Fabric data plane. Used in conjunction with the NginxProxy CRD that is attached to theGatewayClass parametersRef.",
		MarkdownDescription: "ObservabilityPolicy is a Direct Attached Policy. It provides a way to configure observability settings forthe NGINX Gateway Fabric data plane. Used in conjunction with the NginxProxy CRD that is attached to theGatewayClass parametersRef.",
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
				Description:         "Spec defines the desired state of the ObservabilityPolicy.",
				MarkdownDescription: "Spec defines the desired state of the ObservabilityPolicy.",
				Attributes: map[string]schema.Attribute{
					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef identifies an API object to apply the policy to.Object must be in the same namespace as the policy.Support: HTTPRoute",
						MarkdownDescription: "TargetRef identifies an API object to apply the policy to.Object must be in the same namespace as the policy.Support: HTTPRoute",
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

					"tracing": schema.SingleNestedAttribute{
						Description:         "Tracing allows for enabling and configuring tracing.",
						MarkdownDescription: "Tracing allows for enabling and configuring tracing.",
						Attributes: map[string]schema.Attribute{
							"context": schema.StringAttribute{
								Description:         "Context specifies how to propagate traceparent/tracestate headers.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_trace_context",
								MarkdownDescription: "Context specifies how to propagate traceparent/tracestate headers.Default: https://nginx.org/en/docs/ngx_otel_module.html#otel_trace_context",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("extract", "inject", "propagate", "ignore"),
								},
							},

							"ratio": schema.Int64Attribute{
								Description:         "Ratio is the percentage of traffic that should be sampled. Integer from 0 to 100.By default, 100% of http requests are traced. Not applicable for parent-based tracing.",
								MarkdownDescription: "Ratio is the percentage of traffic that should be sampled. Integer from 0 to 100.By default, 100% of http requests are traced. Not applicable for parent-based tracing.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(100),
								},
							},

							"span_attributes": schema.ListNestedAttribute{
								Description:         "SpanAttributes are custom key/value attributes that are added to each span.",
								MarkdownDescription: "SpanAttributes are custom key/value attributes that are added to each span.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "Key is the key for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''",
											MarkdownDescription: "Key is the key for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(255),
												stringvalidator.RegexMatches(regexp.MustCompile(`^([^"$\\]|\\[^$])*$`), ""),
											},
										},

										"value": schema.StringAttribute{
											Description:         "Value is the value for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''",
											MarkdownDescription: "Value is the value for a span attribute.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
												stringvalidator.LengthAtMost(255),
												stringvalidator.RegexMatches(regexp.MustCompile(`^([^"$\\]|\\[^$])*$`), ""),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"span_name": schema.StringAttribute{
								Description:         "SpanName defines the name of the Otel span. By default is the name of the location for a request.If specified, applies to all locations that are created for a route.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''Examples of invalid names: some-$value, quoted-'value'-name, unescaped",
								MarkdownDescription: "SpanName defines the name of the Otel span. By default is the name of the location for a request.If specified, applies to all locations that are created for a route.Format: must have all ''' escaped and must not contain any '$' or end with an unescaped ''Examples of invalid names: some-$value, quoted-'value'-name, unescaped",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(255),
									stringvalidator.RegexMatches(regexp.MustCompile(`^([^"$\\]|\\[^$])*$`), ""),
								},
							},

							"strategy": schema.StringAttribute{
								Description:         "Strategy defines if tracing is ratio-based or parent-based.",
								MarkdownDescription: "Strategy defines if tracing is ratio-based or parent-based.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ratio", "parent"),
								},
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
	}
}

func (r *GatewayNginxOrgObservabilityPolicyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gateway_nginx_org_observability_policy_v1alpha1_manifest")

	var model GatewayNginxOrgObservabilityPolicyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("gateway.nginx.org/v1alpha1")
	model.Kind = pointer.String("ObservabilityPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
