/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package sloth_slok_dev_v1

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
	_ datasource.DataSource = &SlothSlokDevPrometheusServiceLevelV1Manifest{}
)

func NewSlothSlokDevPrometheusServiceLevelV1Manifest() datasource.DataSource {
	return &SlothSlokDevPrometheusServiceLevelV1Manifest{}
}

type SlothSlokDevPrometheusServiceLevelV1Manifest struct{}

type SlothSlokDevPrometheusServiceLevelV1ManifestData struct {
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
		Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Service *string            `tfsdk:"service" json:"service,omitempty"`
		Slos    *[]struct {
			Alerting *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				PageAlert   *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Disable     *bool              `tfsdk:"disable" json:"disable,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"page_alert" json:"pageAlert,omitempty"`
				TicketAlert *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Disable     *bool              `tfsdk:"disable" json:"disable,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"ticket_alert" json:"ticketAlert,omitempty"`
			} `tfsdk:"alerting" json:"alerting,omitempty"`
			Description *string            `tfsdk:"description" json:"description,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
			Objective   *float64           `tfsdk:"objective" json:"objective,omitempty"`
			Sli         *struct {
				Events *struct {
					ErrorQuery *string `tfsdk:"error_query" json:"errorQuery,omitempty"`
					TotalQuery *string `tfsdk:"total_query" json:"totalQuery,omitempty"`
				} `tfsdk:"events" json:"events,omitempty"`
				Plugin *struct {
					Id      *string            `tfsdk:"id" json:"id,omitempty"`
					Options *map[string]string `tfsdk:"options" json:"options,omitempty"`
				} `tfsdk:"plugin" json:"plugin,omitempty"`
				Raw *struct {
					ErrorRatioQuery *string `tfsdk:"error_ratio_query" json:"errorRatioQuery,omitempty"`
				} `tfsdk:"raw" json:"raw,omitempty"`
			} `tfsdk:"sli" json:"sli,omitempty"`
		} `tfsdk:"slos" json:"slos,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *SlothSlokDevPrometheusServiceLevelV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_sloth_slok_dev_prometheus_service_level_v1_manifest"
}

func (r *SlothSlokDevPrometheusServiceLevelV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PrometheusServiceLevel is the expected service quality level using Prometheus as the backend used by Sloth.",
		MarkdownDescription: "PrometheusServiceLevel is the expected service quality level using Prometheus as the backend used by Sloth.",
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
				Description:         "ServiceLevelSpec is the spec for a PrometheusServiceLevel.",
				MarkdownDescription: "ServiceLevelSpec is the spec for a PrometheusServiceLevel.",
				Attributes: map[string]schema.Attribute{
					"labels": schema.MapAttribute{
						Description:         "Labels are the Prometheus labels that will have all the recording and alerting rules generated for the service SLOs.",
						MarkdownDescription: "Labels are the Prometheus labels that will have all the recording and alerting rules generated for the service SLOs.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service": schema.StringAttribute{
						Description:         "Service is the application of the SLOs.",
						MarkdownDescription: "Service is the application of the SLOs.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"slos": schema.ListNestedAttribute{
						Description:         "SLOs are the SLOs of the service.",
						MarkdownDescription: "SLOs are the SLOs of the service.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"alerting": schema.SingleNestedAttribute{
									Description:         "Alerting is the configuration with all the things related with the SLO alerts.",
									MarkdownDescription: "Alerting is the configuration with all the things related with the SLO alerts.",
									Attributes: map[string]schema.Attribute{
										"annotations": schema.MapAttribute{
											Description:         "Annotations are the Prometheus annotations that will have all the alerts generated by this SLO.",
											MarkdownDescription: "Annotations are the Prometheus annotations that will have all the alerts generated by this SLO.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"labels": schema.MapAttribute{
											Description:         "Labels are the Prometheus labels that will have all the alerts generated by this SLO.",
											MarkdownDescription: "Labels are the Prometheus labels that will have all the alerts generated by this SLO.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the name used by the alerts generated for this SLO.",
											MarkdownDescription: "Name is the name used by the alerts generated for this SLO.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"page_alert": schema.SingleNestedAttribute{
											Description:         "Page alert refers to the critical alert (check multiwindow-multiburn alerts).",
											MarkdownDescription: "Page alert refers to the critical alert (check multiwindow-multiburn alerts).",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations are the Prometheus annotations for the specific alert.",
													MarkdownDescription: "Annotations are the Prometheus annotations for the specific alert.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"disable": schema.BoolAttribute{
													Description:         "Disable disables the alert and makes Sloth not generating this alert. This can be helpful for example to disable ticket(warning) alerts.",
													MarkdownDescription: "Disable disables the alert and makes Sloth not generating this alert. This can be helpful for example to disable ticket(warning) alerts.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"labels": schema.MapAttribute{
													Description:         "Labels are the Prometheus labels for the specific alert. For example can be useful to route the Page alert to specific Slack channel.",
													MarkdownDescription: "Labels are the Prometheus labels for the specific alert. For example can be useful to route the Page alert to specific Slack channel.",
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

										"ticket_alert": schema.SingleNestedAttribute{
											Description:         "TicketAlert alert refers to the warning alert (check multiwindow-multiburn alerts).",
											MarkdownDescription: "TicketAlert alert refers to the warning alert (check multiwindow-multiburn alerts).",
											Attributes: map[string]schema.Attribute{
												"annotations": schema.MapAttribute{
													Description:         "Annotations are the Prometheus annotations for the specific alert.",
													MarkdownDescription: "Annotations are the Prometheus annotations for the specific alert.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"disable": schema.BoolAttribute{
													Description:         "Disable disables the alert and makes Sloth not generating this alert. This can be helpful for example to disable ticket(warning) alerts.",
													MarkdownDescription: "Disable disables the alert and makes Sloth not generating this alert. This can be helpful for example to disable ticket(warning) alerts.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"labels": schema.MapAttribute{
													Description:         "Labels are the Prometheus labels for the specific alert. For example can be useful to route the Page alert to specific Slack channel.",
													MarkdownDescription: "Labels are the Prometheus labels for the specific alert. For example can be useful to route the Page alert to specific Slack channel.",
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
									Required: true,
									Optional: false,
									Computed: false,
								},

								"description": schema.StringAttribute{
									Description:         "Description is the description of the SLO.",
									MarkdownDescription: "Description is the description of the SLO.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"labels": schema.MapAttribute{
									Description:         "Labels are the Prometheus labels that will have all the recording and alerting rules for this specific SLO. These labels are merged with the previous level labels.",
									MarkdownDescription: "Labels are the Prometheus labels that will have all the recording and alerting rules for this specific SLO. These labels are merged with the previous level labels.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the name of the SLO.",
									MarkdownDescription: "Name is the name of the SLO.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(128),
									},
								},

								"objective": schema.Float64Attribute{
									Description:         "Objective is target of the SLO the percentage (0, 100] (e.g 99.9).",
									MarkdownDescription: "Objective is target of the SLO the percentage (0, 100] (e.g 99.9).",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"sli": schema.SingleNestedAttribute{
									Description:         "SLI is the indicator (service level indicator) for this specific SLO.",
									MarkdownDescription: "SLI is the indicator (service level indicator) for this specific SLO.",
									Attributes: map[string]schema.Attribute{
										"events": schema.SingleNestedAttribute{
											Description:         "Events is the events SLI type.",
											MarkdownDescription: "Events is the events SLI type.",
											Attributes: map[string]schema.Attribute{
												"error_query": schema.StringAttribute{
													Description:         "ErrorQuery is a Prometheus query that will get the number/count of events that we consider that are bad for the SLO (e.g 'http 5xx', 'latency > 250ms'...). Requires the usage of '{{.window}}' template variable.",
													MarkdownDescription: "ErrorQuery is a Prometheus query that will get the number/count of events that we consider that are bad for the SLO (e.g 'http 5xx', 'latency > 250ms'...). Requires the usage of '{{.window}}' template variable.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"total_query": schema.StringAttribute{
													Description:         "TotalQuery is a Prometheus query that will get the total number/count of events for the SLO (e.g 'all http requests'...). Requires the usage of '{{.window}}' template variable.",
													MarkdownDescription: "TotalQuery is a Prometheus query that will get the total number/count of events for the SLO (e.g 'all http requests'...). Requires the usage of '{{.window}}' template variable.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"plugin": schema.SingleNestedAttribute{
											Description:         "Plugin is the pluggable SLI type.",
											MarkdownDescription: "Plugin is the pluggable SLI type.",
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description:         "Name is the name of the plugin that needs to load.",
													MarkdownDescription: "Name is the name of the plugin that needs to load.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"options": schema.MapAttribute{
													Description:         "Options are the options used for the plugin.",
													MarkdownDescription: "Options are the options used for the plugin.",
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

										"raw": schema.SingleNestedAttribute{
											Description:         "Raw is the raw SLI type.",
											MarkdownDescription: "Raw is the raw SLI type.",
											Attributes: map[string]schema.Attribute{
												"error_ratio_query": schema.StringAttribute{
													Description:         "ErrorRatioQuery is a Prometheus query that will get the raw error ratio (0-1) for the SLO.",
													MarkdownDescription: "ErrorRatioQuery is a Prometheus query that will get the raw error ratio (0-1) for the SLO.",
													Required:            true,
													Optional:            false,
													Computed:            false,
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

func (r *SlothSlokDevPrometheusServiceLevelV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_sloth_slok_dev_prometheus_service_level_v1_manifest")

	var model SlothSlokDevPrometheusServiceLevelV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("sloth.slok.dev/v1")
	model.Kind = pointer.String("PrometheusServiceLevel")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
