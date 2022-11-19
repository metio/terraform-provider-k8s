/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type TelemetryIstioIoTelemetryV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*TelemetryIstioIoTelemetryV1Alpha1Resource)(nil)
)

type TelemetryIstioIoTelemetryV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type TelemetryIstioIoTelemetryV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		AccessLogging *[]struct {
			Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

			Filter *struct {
				Expression *string `tfsdk:"expression" yaml:"expression,omitempty"`
			} `tfsdk:"filter" yaml:"filter,omitempty"`

			Match *struct {
				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			Providers *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"providers" yaml:"providers,omitempty"`
		} `tfsdk:"access_logging" yaml:"accessLogging,omitempty"`

		Metrics *[]struct {
			Overrides *[]struct {
				Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

				Match *struct {
					CustomMetric *string `tfsdk:"custom_metric" yaml:"customMetric,omitempty"`

					Metric *string `tfsdk:"metric" yaml:"metric,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
				} `tfsdk:"match" yaml:"match,omitempty"`

				TagOverrides *struct {
					Operation *string `tfsdk:"operation" yaml:"operation,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tag_overrides" yaml:"tagOverrides,omitempty"`
			} `tfsdk:"overrides" yaml:"overrides,omitempty"`

			Providers *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"providers" yaml:"providers,omitempty"`
		} `tfsdk:"metrics" yaml:"metrics,omitempty"`

		Selector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`

		Tracing *[]struct {
			CustomTags *struct {
				Environment *struct {
					DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"environment" yaml:"environment,omitempty"`

				Header *struct {
					DefaultValue *string `tfsdk:"default_value" yaml:"defaultValue,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"header" yaml:"header,omitempty"`

				Literal *struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"literal" yaml:"literal,omitempty"`
			} `tfsdk:"custom_tags" yaml:"customTags,omitempty"`

			DisableSpanReporting *bool `tfsdk:"disable_span_reporting" yaml:"disableSpanReporting,omitempty"`

			Match *struct {
				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`
			} `tfsdk:"match" yaml:"match,omitempty"`

			Providers *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"providers" yaml:"providers,omitempty"`

			RandomSamplingPercentage utilities.DynamicNumber `tfsdk:"random_sampling_percentage" yaml:"randomSamplingPercentage,omitempty"`

			UseRequestIdForTraceSampling *bool `tfsdk:"use_request_id_for_trace_sampling" yaml:"useRequestIdForTraceSampling,omitempty"`
		} `tfsdk:"tracing" yaml:"tracing,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewTelemetryIstioIoTelemetryV1Alpha1Resource() resource.Resource {
	return &TelemetryIstioIoTelemetryV1Alpha1Resource{}
}

func (r *TelemetryIstioIoTelemetryV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_telemetry_istio_io_telemetry_v1alpha1"
}

func (r *TelemetryIstioIoTelemetryV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "Telemetry configuration for workloads. See more details at: https://istio.io/docs/reference/config/telemetry.html",
				MarkdownDescription: "Telemetry configuration for workloads. See more details at: https://istio.io/docs/reference/config/telemetry.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"access_logging": {
						Description:         "Optional.",
						MarkdownDescription: "Optional.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"disabled": {
								Description:         "Controls logging.",
								MarkdownDescription: "Controls logging.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"filter": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"expression": {
										Description:         "CEL expression for selecting when requests/connections should be logged.",
										MarkdownDescription: "CEL expression for selecting when requests/connections should be logged.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match": {
								Description:         "Allows tailoring of logging behavior to specific conditions.",
								MarkdownDescription: "Allows tailoring of logging behavior to specific conditions.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"mode": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("CLIENT_AND_SERVER", "CLIENT", "SERVER"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"providers": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Required.",
										MarkdownDescription: "Required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": {
						Description:         "Optional.",
						MarkdownDescription: "Optional.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"overrides": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"disabled": {
										Description:         "Optional.",
										MarkdownDescription: "Optional.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"match": {
										Description:         "Match allows provides the scope of the override.",
										MarkdownDescription: "Match allows provides the scope of the override.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"custom_metric": {
												Description:         "Allows free-form specification of a metric.",
												MarkdownDescription: "Allows free-form specification of a metric.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"metric": {
												Description:         "One of the well-known Istio Standard Metrics.",
												MarkdownDescription: "One of the well-known Istio Standard Metrics.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ALL_METRICS", "REQUEST_COUNT", "REQUEST_DURATION", "REQUEST_SIZE", "RESPONSE_SIZE", "TCP_OPENED_CONNECTIONS", "TCP_CLOSED_CONNECTIONS", "TCP_SENT_BYTES", "TCP_RECEIVED_BYTES", "GRPC_REQUEST_MESSAGES", "GRPC_RESPONSE_MESSAGES"),
												},
											},

											"mode": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("CLIENT_AND_SERVER", "CLIENT", "SERVER"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tag_overrides": {
										Description:         "Optional.",
										MarkdownDescription: "Optional.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"operation": {
												Description:         "Operation controls whether or not to update/add a tag, or to remove it.",
												MarkdownDescription: "Operation controls whether or not to update/add a tag, or to remove it.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("UPSERT", "REMOVE"),
												},
											},

											"value": {
												Description:         "Value is only considered if the operation is 'UPSERT'.",
												MarkdownDescription: "Value is only considered if the operation is 'UPSERT'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"providers": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Required.",
										MarkdownDescription: "Required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "Optional.",
						MarkdownDescription: "Optional.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_labels": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tracing": {
						Description:         "Optional.",
						MarkdownDescription: "Optional.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"custom_tags": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"environment": {
										Description:         "Environment adds the value of an environment variable to each span.",
										MarkdownDescription: "Environment adds the value of an environment variable to each span.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default_value": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name of the environment variable from which to extract the tag value.",
												MarkdownDescription: "Name of the environment variable from which to extract the tag value.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"header": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default_value": {
												Description:         "Optional.",
												MarkdownDescription: "Optional.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name of the header from which to extract the tag value.",
												MarkdownDescription: "Name of the header from which to extract the tag value.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"literal": {
										Description:         "Literal adds the same, hard-coded value to each span.",
										MarkdownDescription: "Literal adds the same, hard-coded value to each span.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "The tag value to use.",
												MarkdownDescription: "The tag value to use.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_span_reporting": {
								Description:         "Controls span reporting.",
								MarkdownDescription: "Controls span reporting.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"match": {
								Description:         "Allows tailoring of behavior to specific conditions.",
								MarkdownDescription: "Allows tailoring of behavior to specific conditions.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"mode": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("CLIENT_AND_SERVER", "CLIENT", "SERVER"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"providers": {
								Description:         "Optional.",
								MarkdownDescription: "Optional.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Required.",
										MarkdownDescription: "Required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"random_sampling_percentage": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"use_request_id_for_trace_sampling": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *TelemetryIstioIoTelemetryV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_telemetry_istio_io_telemetry_v1alpha1")

	var state TelemetryIstioIoTelemetryV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TelemetryIstioIoTelemetryV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("telemetry.istio.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Telemetry")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *TelemetryIstioIoTelemetryV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_telemetry_istio_io_telemetry_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *TelemetryIstioIoTelemetryV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_telemetry_istio_io_telemetry_v1alpha1")

	var state TelemetryIstioIoTelemetryV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel TelemetryIstioIoTelemetryV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("telemetry.istio.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Telemetry")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *TelemetryIstioIoTelemetryV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_telemetry_istio_io_telemetry_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
