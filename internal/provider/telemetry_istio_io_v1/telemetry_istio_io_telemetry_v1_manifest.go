/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package telemetry_istio_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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
	_ datasource.DataSource = &TelemetryIstioIoTelemetryV1Manifest{}
)

func NewTelemetryIstioIoTelemetryV1Manifest() datasource.DataSource {
	return &TelemetryIstioIoTelemetryV1Manifest{}
}

type TelemetryIstioIoTelemetryV1Manifest struct{}

type TelemetryIstioIoTelemetryV1ManifestData struct {
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
		AccessLogging *[]struct {
			Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
			Filter   *struct {
				Expression *string `tfsdk:"expression" json:"expression,omitempty"`
			} `tfsdk:"filter" json:"filter,omitempty"`
			Match *struct {
				Mode *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Providers *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"providers" json:"providers,omitempty"`
		} `tfsdk:"access_logging" json:"accessLogging,omitempty"`
		Metrics *[]struct {
			Overrides *[]struct {
				Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				Match    *struct {
					CustomMetric *string `tfsdk:"custom_metric" json:"customMetric,omitempty"`
					Metric       *string `tfsdk:"metric" json:"metric,omitempty"`
					Mode         *string `tfsdk:"mode" json:"mode,omitempty"`
				} `tfsdk:"match" json:"match,omitempty"`
				TagOverrides *struct {
					Operation *string `tfsdk:"operation" json:"operation,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tag_overrides" json:"tagOverrides,omitempty"`
			} `tfsdk:"overrides" json:"overrides,omitempty"`
			Providers *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"providers" json:"providers,omitempty"`
			ReportingInterval *string `tfsdk:"reporting_interval" json:"reportingInterval,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		Selector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"selector" json:"selector,omitempty"`
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		TargetRefs *[]struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_refs" json:"targetRefs,omitempty"`
		Tracing *[]struct {
			CustomTags *struct {
				Environment *struct {
					DefaultValue *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
					Name         *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"environment" json:"environment,omitempty"`
				Header *struct {
					DefaultValue *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
					Name         *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"header" json:"header,omitempty"`
				Literal *struct {
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"literal" json:"literal,omitempty"`
			} `tfsdk:"custom_tags" json:"customTags,omitempty"`
			DisableSpanReporting *bool `tfsdk:"disable_span_reporting" json:"disableSpanReporting,omitempty"`
			Match                *struct {
				Mode *string `tfsdk:"mode" json:"mode,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Providers *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"providers" json:"providers,omitempty"`
			RandomSamplingPercentage     *float64 `tfsdk:"random_sampling_percentage" json:"randomSamplingPercentage,omitempty"`
			UseRequestIdForTraceSampling *bool    `tfsdk:"use_request_id_for_trace_sampling" json:"useRequestIdForTraceSampling,omitempty"`
		} `tfsdk:"tracing" json:"tracing,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *TelemetryIstioIoTelemetryV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_telemetry_istio_io_telemetry_v1_manifest"
}

func (r *TelemetryIstioIoTelemetryV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Telemetry configuration for workloads. See more details at: https://istio.io/docs/reference/config/telemetry.html",
				MarkdownDescription: "Telemetry configuration for workloads. See more details at: https://istio.io/docs/reference/config/telemetry.html",
				Attributes: map[string]schema.Attribute{
					"access_logging": schema.ListNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"disabled": schema.BoolAttribute{
									Description:         "Controls logging.",
									MarkdownDescription: "Controls logging.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"filter": schema.SingleNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									Attributes: map[string]schema.Attribute{
										"expression": schema.StringAttribute{
											Description:         "CEL expression for selecting when requests/connections should be logged.",
											MarkdownDescription: "CEL expression for selecting when requests/connections should be logged.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"match": schema.SingleNestedAttribute{
									Description:         "Allows tailoring of logging behavior to specific conditions.",
									MarkdownDescription: "Allows tailoring of logging behavior to specific conditions.",
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description:         "This determines whether or not to apply the access logging configuration based on the direction of traffic relative to the proxied workload.Valid Options: CLIENT_AND_SERVER, CLIENT, SERVER",
											MarkdownDescription: "This determines whether or not to apply the access logging configuration based on the direction of traffic relative to the proxied workload.Valid Options: CLIENT_AND_SERVER, CLIENT, SERVER",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("CLIENT_AND_SERVER", "CLIENT", "SERVER"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"providers": schema.ListNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Required.",
												MarkdownDescription: "Required.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": schema.ListNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"overrides": schema.ListNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"disabled": schema.BoolAttribute{
												Description:         "Optional.",
												MarkdownDescription: "Optional.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"match": schema.SingleNestedAttribute{
												Description:         "Match allows providing the scope of the override.",
												MarkdownDescription: "Match allows providing the scope of the override.",
												Attributes: map[string]schema.Attribute{
													"custom_metric": schema.StringAttribute{
														Description:         "Allows free-form specification of a metric.",
														MarkdownDescription: "Allows free-form specification of a metric.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"metric": schema.StringAttribute{
														Description:         "One of the well-known [Istio Standard Metrics](https://istio.io/latest/docs/reference/config/metrics/).Valid Options: ALL_METRICS, REQUEST_COUNT, REQUEST_DURATION, REQUEST_SIZE, RESPONSE_SIZE, TCP_OPENED_CONNECTIONS, TCP_CLOSED_CONNECTIONS, TCP_SENT_BYTES, TCP_RECEIVED_BYTES, GRPC_REQUEST_MESSAGES, GRPC_RESPONSE_MESSAGES",
														MarkdownDescription: "One of the well-known [Istio Standard Metrics](https://istio.io/latest/docs/reference/config/metrics/).Valid Options: ALL_METRICS, REQUEST_COUNT, REQUEST_DURATION, REQUEST_SIZE, RESPONSE_SIZE, TCP_OPENED_CONNECTIONS, TCP_CLOSED_CONNECTIONS, TCP_SENT_BYTES, TCP_RECEIVED_BYTES, GRPC_REQUEST_MESSAGES, GRPC_RESPONSE_MESSAGES",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ALL_METRICS", "REQUEST_COUNT", "REQUEST_DURATION", "REQUEST_SIZE", "RESPONSE_SIZE", "TCP_OPENED_CONNECTIONS", "TCP_CLOSED_CONNECTIONS", "TCP_SENT_BYTES", "TCP_RECEIVED_BYTES", "GRPC_REQUEST_MESSAGES", "GRPC_RESPONSE_MESSAGES"),
														},
													},

													"mode": schema.StringAttribute{
														Description:         "Controls which mode of metrics generation is selected: 'CLIENT', 'SERVER', or 'CLIENT_AND_SERVER'.Valid Options: CLIENT_AND_SERVER, CLIENT, SERVER",
														MarkdownDescription: "Controls which mode of metrics generation is selected: 'CLIENT', 'SERVER', or 'CLIENT_AND_SERVER'.Valid Options: CLIENT_AND_SERVER, CLIENT, SERVER",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("CLIENT_AND_SERVER", "CLIENT", "SERVER"),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"tag_overrides": schema.SingleNestedAttribute{
												Description:         "Optional.",
												MarkdownDescription: "Optional.",
												Attributes: map[string]schema.Attribute{
													"operation": schema.StringAttribute{
														Description:         "Operation controls whether or not to update/add a tag, or to remove it.Valid Options: UPSERT, REMOVE",
														MarkdownDescription: "Operation controls whether or not to update/add a tag, or to remove it.Valid Options: UPSERT, REMOVE",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("UPSERT", "REMOVE"),
														},
													},

													"value": schema.StringAttribute{
														Description:         "Value is only considered if the operation is 'UPSERT'.",
														MarkdownDescription: "Value is only considered if the operation is 'UPSERT'.",
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

								"providers": schema.ListNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Required.",
												MarkdownDescription: "Required.",
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

								"reporting_interval": schema.StringAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
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

					"selector": schema.SingleNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "One or more labels that indicate a specific set of pods/VMs on which a policy should be applied.",
								MarkdownDescription: "One or more labels that indicate a specific set of pods/VMs on which a policy should be applied.",
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

					"target_ref": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "group is the group of the target resource.",
								MarkdownDescription: "group is the group of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "kind is kind of the target resource.",
								MarkdownDescription: "kind is kind of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "name is the name of the target resource.",
								MarkdownDescription: "name is the name of the target resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace is the namespace of the referent.",
								MarkdownDescription: "namespace is the namespace of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_refs": schema.ListNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"group": schema.StringAttribute{
									Description:         "group is the group of the target resource.",
									MarkdownDescription: "group is the group of the target resource.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"kind": schema.StringAttribute{
									Description:         "kind is kind of the target resource.",
									MarkdownDescription: "kind is kind of the target resource.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "name is the name of the target resource.",
									MarkdownDescription: "name is the name of the target resource.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"namespace": schema.StringAttribute{
									Description:         "namespace is the namespace of the referent.",
									MarkdownDescription: "namespace is the namespace of the referent.",
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

					"tracing": schema.ListNestedAttribute{
						Description:         "Optional.",
						MarkdownDescription: "Optional.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"custom_tags": schema.SingleNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									Attributes: map[string]schema.Attribute{
										"environment": schema.SingleNestedAttribute{
											Description:         "Environment adds the value of an environment variable to each span.",
											MarkdownDescription: "Environment adds the value of an environment variable to each span.",
											Attributes: map[string]schema.Attribute{
												"default_value": schema.StringAttribute{
													Description:         "Optional.",
													MarkdownDescription: "Optional.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the environment variable from which to extract the tag value.",
													MarkdownDescription: "Name of the environment variable from which to extract the tag value.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"header": schema.SingleNestedAttribute{
											Description:         "RequestHeader adds the value of an header from the request to each span.",
											MarkdownDescription: "RequestHeader adds the value of an header from the request to each span.",
											Attributes: map[string]schema.Attribute{
												"default_value": schema.StringAttribute{
													Description:         "Optional.",
													MarkdownDescription: "Optional.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the header from which to extract the tag value.",
													MarkdownDescription: "Name of the header from which to extract the tag value.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"literal": schema.SingleNestedAttribute{
											Description:         "Literal adds the same, hard-coded value to each span.",
											MarkdownDescription: "Literal adds the same, hard-coded value to each span.",
											Attributes: map[string]schema.Attribute{
												"value": schema.StringAttribute{
													Description:         "The tag value to use.",
													MarkdownDescription: "The tag value to use.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
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

								"disable_span_reporting": schema.BoolAttribute{
									Description:         "Controls span reporting.",
									MarkdownDescription: "Controls span reporting.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"match": schema.SingleNestedAttribute{
									Description:         "Allows tailoring of behavior to specific conditions.",
									MarkdownDescription: "Allows tailoring of behavior to specific conditions.",
									Attributes: map[string]schema.Attribute{
										"mode": schema.StringAttribute{
											Description:         "This determines whether or not to apply the tracing configuration based on the direction of traffic relative to the proxied workload.Valid Options: CLIENT_AND_SERVER, CLIENT, SERVER",
											MarkdownDescription: "This determines whether or not to apply the tracing configuration based on the direction of traffic relative to the proxied workload.Valid Options: CLIENT_AND_SERVER, CLIENT, SERVER",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("CLIENT_AND_SERVER", "CLIENT", "SERVER"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"providers": schema.ListNestedAttribute{
									Description:         "Optional.",
									MarkdownDescription: "Optional.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Required.",
												MarkdownDescription: "Required.",
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

								"random_sampling_percentage": schema.Float64Attribute{
									Description:         "Controls the rate at which traffic will be selected for tracing if no prior sampling decision has been made.",
									MarkdownDescription: "Controls the rate at which traffic will be selected for tracing if no prior sampling decision has been made.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Float64{
										float64validator.AtLeast(0),
										float64validator.AtMost(100),
									},
								},

								"use_request_id_for_trace_sampling": schema.BoolAttribute{
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
	}
}

func (r *TelemetryIstioIoTelemetryV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_telemetry_istio_io_telemetry_v1_manifest")

	var model TelemetryIstioIoTelemetryV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("telemetry.istio.io/v1")
	model.Kind = pointer.String("Telemetry")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
