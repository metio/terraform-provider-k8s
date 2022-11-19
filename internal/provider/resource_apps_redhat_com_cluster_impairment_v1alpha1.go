/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type AppsRedhatComClusterImpairmentV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*AppsRedhatComClusterImpairmentV1Alpha1Resource)(nil)
)

type AppsRedhatComClusterImpairmentV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AppsRedhatComClusterImpairmentV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		Duration *int64 `tfsdk:"duration" yaml:"duration,omitempty"`

		Egress *struct {
			Bandwidth *int64 `tfsdk:"bandwidth" yaml:"bandwidth,omitempty"`

			Corruption utilities.DynamicNumber `tfsdk:"corruption" yaml:"corruption,omitempty"`

			CorruptionOptions *struct {
				Correlation utilities.DynamicNumber `tfsdk:"correlation" yaml:"correlation,omitempty"`
			} `tfsdk:"corruption_options" yaml:"corruptionOptions,omitempty"`

			Duplication utilities.DynamicNumber `tfsdk:"duplication" yaml:"duplication,omitempty"`

			DuplicationOptions *struct {
				Correlation utilities.DynamicNumber `tfsdk:"correlation" yaml:"correlation,omitempty"`
			} `tfsdk:"duplication_options" yaml:"duplicationOptions,omitempty"`

			Latency utilities.DynamicNumber `tfsdk:"latency" yaml:"latency,omitempty"`

			LatencyOptions *struct {
				Distribution *string `tfsdk:"distribution" yaml:"distribution,omitempty"`

				Jitter utilities.DynamicNumber `tfsdk:"jitter" yaml:"jitter,omitempty"`

				JitterCorrelation utilities.DynamicNumber `tfsdk:"jitter_correlation" yaml:"jitterCorrelation,omitempty"`

				Reorder utilities.DynamicNumber `tfsdk:"reorder" yaml:"reorder,omitempty"`

				ReorderCorrelation utilities.DynamicNumber `tfsdk:"reorder_correlation" yaml:"reorderCorrelation,omitempty"`
			} `tfsdk:"latency_options" yaml:"latencyOptions,omitempty"`

			Loss utilities.DynamicNumber `tfsdk:"loss" yaml:"loss,omitempty"`

			LossOptions *struct {
				Correlation utilities.DynamicNumber `tfsdk:"correlation" yaml:"correlation,omitempty"`
			} `tfsdk:"loss_options" yaml:"lossOptions,omitempty"`
		} `tfsdk:"egress" yaml:"egress,omitempty"`

		Ingress *struct {
			Bandwidth *int64 `tfsdk:"bandwidth" yaml:"bandwidth,omitempty"`

			Corruption utilities.DynamicNumber `tfsdk:"corruption" yaml:"corruption,omitempty"`

			CorruptionOptions *struct {
				Correlation utilities.DynamicNumber `tfsdk:"correlation" yaml:"correlation,omitempty"`
			} `tfsdk:"corruption_options" yaml:"corruptionOptions,omitempty"`

			Duplication utilities.DynamicNumber `tfsdk:"duplication" yaml:"duplication,omitempty"`

			DuplicationOptions *struct {
				Correlation utilities.DynamicNumber `tfsdk:"correlation" yaml:"correlation,omitempty"`
			} `tfsdk:"duplication_options" yaml:"duplicationOptions,omitempty"`

			Latency utilities.DynamicNumber `tfsdk:"latency" yaml:"latency,omitempty"`

			LatencyOptions *struct {
				Distribution *string `tfsdk:"distribution" yaml:"distribution,omitempty"`

				Jitter utilities.DynamicNumber `tfsdk:"jitter" yaml:"jitter,omitempty"`

				JitterCorrelation utilities.DynamicNumber `tfsdk:"jitter_correlation" yaml:"jitterCorrelation,omitempty"`

				Reorder utilities.DynamicNumber `tfsdk:"reorder" yaml:"reorder,omitempty"`

				ReorderCorrelation utilities.DynamicNumber `tfsdk:"reorder_correlation" yaml:"reorderCorrelation,omitempty"`
			} `tfsdk:"latency_options" yaml:"latencyOptions,omitempty"`

			Loss utilities.DynamicNumber `tfsdk:"loss" yaml:"loss,omitempty"`

			LossOptions *struct {
				Correlation utilities.DynamicNumber `tfsdk:"correlation" yaml:"correlation,omitempty"`
			} `tfsdk:"loss_options" yaml:"lossOptions,omitempty"`
		} `tfsdk:"ingress" yaml:"ingress,omitempty"`

		Interfaces *[]string `tfsdk:"interfaces" yaml:"interfaces,omitempty"`

		LinkFlapping *struct {
			DownTime *int64 `tfsdk:"down_time" yaml:"downTime,omitempty"`

			Enable *bool `tfsdk:"enable" yaml:"enable,omitempty"`

			UpTime *int64 `tfsdk:"up_time" yaml:"upTime,omitempty"`
		} `tfsdk:"link_flapping" yaml:"linkFlapping,omitempty"`

		NodeSelector *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		StartDelay *int64 `tfsdk:"start_delay" yaml:"startDelay,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAppsRedhatComClusterImpairmentV1Alpha1Resource() resource.Resource {
	return &AppsRedhatComClusterImpairmentV1Alpha1Resource{}
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_apps_redhat_com_cluster_impairment_v1alpha1"
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ClusterImpairment is the Schema for the clusterimpairments API",
		MarkdownDescription: "ClusterImpairment is the Schema for the clusterimpairments API",
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
				Description:         "Spec defines the desired state of ClusterImpairment",
				MarkdownDescription: "Spec defines the desired state of ClusterImpairment",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"duration": {
						Description:         "The duration of the impairment in seconds.",
						MarkdownDescription: "The duration of the impairment in seconds.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"egress": {
						Description:         "The configuration section that specifies the egress impairments.",
						MarkdownDescription: "The configuration section that specifies the egress impairments.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bandwidth": {
								Description:         "The bandwidth limit in kbit/sec",
								MarkdownDescription: "The bandwidth limit in kbit/sec",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"corruption": {
								Description:         "The percent of packets that are corrupted",
								MarkdownDescription: "The percent of packets that are corrupted",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"corruption_options": {
								Description:         "Advanced corruption options",
								MarkdownDescription: "Advanced corruption options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "The correlation between sequential corruption values",
										MarkdownDescription: "The correlation between sequential corruption values",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplication": {
								Description:         "The percent of packets duplicated",
								MarkdownDescription: "The percent of packets duplicated",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplication_options": {
								Description:         "Advanced duplication options",
								MarkdownDescription: "Advanced duplication options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "The correlation between sequential duplication values",
										MarkdownDescription: "The correlation between sequential duplication values",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency": {
								Description:         "The latency applied in ms",
								MarkdownDescription: "The latency applied in ms",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency_options": {
								Description:         "Advanced latency options. Example: jitter",
								MarkdownDescription: "Advanced latency options. Example: jitter",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"distribution": {
										Description:         "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										MarkdownDescription: "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jitter": {
										Description:         "Variation in the latency that follows the specified distribution.",
										MarkdownDescription: "Variation in the latency that follows the specified distribution.",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jitter_correlation": {
										Description:         "The correlation between sequential jitter values",
										MarkdownDescription: "The correlation between sequential jitter values",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reorder": {
										Description:         "The percentage of packets that are not delayed, causing reordering",
										MarkdownDescription: "The percentage of packets that are not delayed, causing reordering",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reorder_correlation": {
										Description:         "The correlation between sequential reorder values",
										MarkdownDescription: "The correlation between sequential reorder values",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"loss": {
								Description:         "The packet loss in percent",
								MarkdownDescription: "The packet loss in percent",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"loss_options": {
								Description:         "Advanced packet loss options",
								MarkdownDescription: "Advanced packet loss options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "The correlation between sequential packet loss values",
										MarkdownDescription: "The correlation between sequential packet loss values",

										Type: utilities.DynamicNumberType{},

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

					"ingress": {
						Description:         "The configuration section that specifies the ingress impairments.",
						MarkdownDescription: "The configuration section that specifies the ingress impairments.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bandwidth": {
								Description:         "The bandwidth limit in kbit/sec",
								MarkdownDescription: "The bandwidth limit in kbit/sec",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"corruption": {
								Description:         "The percent of packets that are corrupted",
								MarkdownDescription: "The percent of packets that are corrupted",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"corruption_options": {
								Description:         "Advanced corruption options",
								MarkdownDescription: "Advanced corruption options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "The correlation between sequential corruption values",
										MarkdownDescription: "The correlation between sequential corruption values",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplication": {
								Description:         "The percent of packets duplicated",
								MarkdownDescription: "The percent of packets duplicated",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"duplication_options": {
								Description:         "Advanced duplication options",
								MarkdownDescription: "Advanced duplication options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "The correlation between sequential duplication values",
										MarkdownDescription: "The correlation between sequential duplication values",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency": {
								Description:         "The latency applied in ms",
								MarkdownDescription: "The latency applied in ms",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"latency_options": {
								Description:         "Advanced latency options. Example: jitter",
								MarkdownDescription: "Advanced latency options. Example: jitter",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"distribution": {
										Description:         "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",
										MarkdownDescription: "The way the jitter is distributed. Options: Normal, Uniform, Pareto, Paretonormal",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jitter": {
										Description:         "Variation in the latency that follows the specified distribution.",
										MarkdownDescription: "Variation in the latency that follows the specified distribution.",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jitter_correlation": {
										Description:         "The correlation between sequential jitter values",
										MarkdownDescription: "The correlation between sequential jitter values",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reorder": {
										Description:         "The percentage of packets that are not delayed, causing reordering",
										MarkdownDescription: "The percentage of packets that are not delayed, causing reordering",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reorder_correlation": {
										Description:         "The correlation between sequential reorder values",
										MarkdownDescription: "The correlation between sequential reorder values",

										Type: utilities.DynamicNumberType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"loss": {
								Description:         "The packet loss in percent",
								MarkdownDescription: "The packet loss in percent",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"loss_options": {
								Description:         "Advanced packet loss options",
								MarkdownDescription: "Advanced packet loss options",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"correlation": {
										Description:         "The correlation between sequential packet loss values",
										MarkdownDescription: "The correlation between sequential packet loss values",

										Type: utilities.DynamicNumberType{},

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

					"interfaces": {
						Description:         "All interfaces that the impairments should be applied to. Must be valid interfaces or the impairments will fail to apply.",
						MarkdownDescription: "All interfaces that the impairments should be applied to. Must be valid interfaces or the impairments will fail to apply.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"link_flapping": {
						Description:         "The configuration section that specifies the link flapping impairments.",
						MarkdownDescription: "The configuration section that specifies the link flapping impairments.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"down_time": {
								Description:         "The duration that the link should be disabled.",
								MarkdownDescription: "The duration that the link should be disabled.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": {
								Description:         "Whether to enable link flapping.",
								MarkdownDescription: "Whether to enable link flapping.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"up_time": {
								Description:         "The duration that the link should be enabled.",
								MarkdownDescription: "The duration that the link should be enabled.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "The configuration section that specifies the node selector that should be applied to the daemonset. Default: worker nodes.",
						MarkdownDescription: "The configuration section that specifies the node selector that should be applied to the daemonset. Default: worker nodes.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "The key for the node selector",
								MarkdownDescription: "The key for the node selector",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "The value for the node selector",
								MarkdownDescription: "The value for the node selector",

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

					"start_delay": {
						Description:         "The delay (in seconds) before starting the impairments. At least 5 seconds recommended for Kubernetes and for synchronization of the impairments.",
						MarkdownDescription: "The delay (in seconds) before starting the impairments. At least 5 seconds recommended for Kubernetes and for synchronization of the impairments.",

						Type: types.Int64Type,

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

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_apps_redhat_com_cluster_impairment_v1alpha1")

	var state AppsRedhatComClusterImpairmentV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppsRedhatComClusterImpairmentV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.redhat.com/v1alpha1")
	goModel.Kind = utilities.Ptr("ClusterImpairment")

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

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_redhat_com_cluster_impairment_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_apps_redhat_com_cluster_impairment_v1alpha1")

	var state AppsRedhatComClusterImpairmentV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppsRedhatComClusterImpairmentV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("apps.redhat.com/v1alpha1")
	goModel.Kind = utilities.Ptr("ClusterImpairment")

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

func (r *AppsRedhatComClusterImpairmentV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_apps_redhat_com_cluster_impairment_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
