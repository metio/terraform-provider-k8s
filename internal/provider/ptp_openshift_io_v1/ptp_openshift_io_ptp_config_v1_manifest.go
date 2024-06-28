/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ptp_openshift_io_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &PtpOpenshiftIoPtpConfigV1Manifest{}
)

func NewPtpOpenshiftIoPtpConfigV1Manifest() datasource.DataSource {
	return &PtpOpenshiftIoPtpConfigV1Manifest{}
}

type PtpOpenshiftIoPtpConfigV1Manifest struct{}

type PtpOpenshiftIoPtpConfigV1ManifestData struct {
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
		Profile *[]struct {
			Interface         *string            `tfsdk:"interface" json:"interface,omitempty"`
			Name              *string            `tfsdk:"name" json:"name,omitempty"`
			Phc2sysConf       *string            `tfsdk:"phc2sys_conf" json:"phc2sysConf,omitempty"`
			Phc2sysOpts       *string            `tfsdk:"phc2sys_opts" json:"phc2sysOpts,omitempty"`
			Plugins           *map[string]string `tfsdk:"plugins" json:"plugins,omitempty"`
			Ptp4lConf         *string            `tfsdk:"ptp4l_conf" json:"ptp4lConf,omitempty"`
			Ptp4lOpts         *string            `tfsdk:"ptp4l_opts" json:"ptp4lOpts,omitempty"`
			PtpClockThreshold *struct {
				HoldOverTimeout    *int64 `tfsdk:"hold_over_timeout" json:"holdOverTimeout,omitempty"`
				MaxOffsetThreshold *int64 `tfsdk:"max_offset_threshold" json:"maxOffsetThreshold,omitempty"`
				MinOffsetThreshold *int64 `tfsdk:"min_offset_threshold" json:"minOffsetThreshold,omitempty"`
			} `tfsdk:"ptp_clock_threshold" json:"ptpClockThreshold,omitempty"`
			PtpSchedulingPolicy   *string            `tfsdk:"ptp_scheduling_policy" json:"ptpSchedulingPolicy,omitempty"`
			PtpSchedulingPriority *int64             `tfsdk:"ptp_scheduling_priority" json:"ptpSchedulingPriority,omitempty"`
			PtpSettings           *map[string]string `tfsdk:"ptp_settings" json:"ptpSettings,omitempty"`
			Synce4lConf           *string            `tfsdk:"synce4l_conf" json:"synce4lConf,omitempty"`
			Synce4lOpts           *string            `tfsdk:"synce4l_opts" json:"synce4lOpts,omitempty"`
			Ts2phcConf            *string            `tfsdk:"ts2phc_conf" json:"ts2phcConf,omitempty"`
			Ts2phcOpts            *string            `tfsdk:"ts2phc_opts" json:"ts2phcOpts,omitempty"`
		} `tfsdk:"profile" json:"profile,omitempty"`
		Recommend *[]struct {
			Match *[]struct {
				NodeLabel *string `tfsdk:"node_label" json:"nodeLabel,omitempty"`
				NodeName  *string `tfsdk:"node_name" json:"nodeName,omitempty"`
			} `tfsdk:"match" json:"match,omitempty"`
			Priority *int64  `tfsdk:"priority" json:"priority,omitempty"`
			Profile  *string `tfsdk:"profile" json:"profile,omitempty"`
		} `tfsdk:"recommend" json:"recommend,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PtpOpenshiftIoPtpConfigV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ptp_openshift_io_ptp_config_v1_manifest"
}

func (r *PtpOpenshiftIoPtpConfigV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PtpConfig is the Schema for the ptpconfigs API",
		MarkdownDescription: "PtpConfig is the Schema for the ptpconfigs API",
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
				Description:         "PtpConfigSpec defines the desired state of PtpConfig",
				MarkdownDescription: "PtpConfigSpec defines the desired state of PtpConfig",
				Attributes: map[string]schema.Attribute{
					"profile": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"interface": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"phc2sys_conf": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"phc2sys_opts": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"plugins": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ptp4l_conf": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ptp4l_opts": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ptp_clock_threshold": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"hold_over_timeout": schema.Int64Attribute{
											Description:         "clock state to stay in holdover state in secs",
											MarkdownDescription: "clock state to stay in holdover state in secs",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_offset_threshold": schema.Int64Attribute{
											Description:         "max offset in nano secs",
											MarkdownDescription: "max offset in nano secs",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"min_offset_threshold": schema.Int64Attribute{
											Description:         "min offset in nano secs",
											MarkdownDescription: "min offset in nano secs",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ptp_scheduling_policy": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("SCHED_OTHER", "SCHED_FIFO"),
									},
								},

								"ptp_scheduling_priority": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(1),
										int64validator.AtMost(65),
									},
								},

								"ptp_settings": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"synce4l_conf": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"synce4l_opts": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ts2phc_conf": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ts2phc_opts": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"recommend": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"match": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"node_label": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_name": schema.StringAttribute{
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

								"priority": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"profile": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
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

func (r *PtpOpenshiftIoPtpConfigV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ptp_openshift_io_ptp_config_v1_manifest")

	var model PtpOpenshiftIoPtpConfigV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ptp.openshift.io/v1")
	model.Kind = pointer.String("PtpConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
