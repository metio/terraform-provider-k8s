/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package datadoghq_com_v1alpha1

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
	_ datasource.DataSource = &DatadoghqComDatadogMetricV1Alpha1Manifest{}
)

func NewDatadoghqComDatadogMetricV1Alpha1Manifest() datasource.DataSource {
	return &DatadoghqComDatadogMetricV1Alpha1Manifest{}
}

type DatadoghqComDatadogMetricV1Alpha1Manifest struct{}

type DatadoghqComDatadogMetricV1Alpha1ManifestData struct {
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
		ExternalMetricName *string `tfsdk:"external_metric_name" json:"externalMetricName,omitempty"`
		MaxAge             *string `tfsdk:"max_age" json:"maxAge,omitempty"`
		Query              *string `tfsdk:"query" json:"query,omitempty"`
		TimeWindow         *string `tfsdk:"time_window" json:"timeWindow,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DatadoghqComDatadogMetricV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_datadoghq_com_datadog_metric_v1alpha1_manifest"
}

func (r *DatadoghqComDatadogMetricV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DatadogMetric allows autoscaling on arbitrary Datadog query",
		MarkdownDescription: "DatadogMetric allows autoscaling on arbitrary Datadog query",
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
				Description:         "DatadogMetricSpec defines the desired state of DatadogMetric",
				MarkdownDescription: "DatadogMetricSpec defines the desired state of DatadogMetric",
				Attributes: map[string]schema.Attribute{
					"external_metric_name": schema.StringAttribute{
						Description:         "ExternalMetricName is reserved for internal use",
						MarkdownDescription: "ExternalMetricName is reserved for internal use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_age": schema.StringAttribute{
						Description:         "MaxAge provides the max age for the metric query (overrides the default setting 'external_metrics_provider.max_age')",
						MarkdownDescription: "MaxAge provides the max age for the metric query (overrides the default setting 'external_metrics_provider.max_age')",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"query": schema.StringAttribute{
						Description:         "Query is the raw datadog query",
						MarkdownDescription: "Query is the raw datadog query",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"time_window": schema.StringAttribute{
						Description:         "TimeWindow provides the time window for the metric query, defaults to MaxAge.",
						MarkdownDescription: "TimeWindow provides the time window for the metric query, defaults to MaxAge.",
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
	}
}

func (r *DatadoghqComDatadogMetricV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_datadoghq_com_datadog_metric_v1alpha1_manifest")

	var model DatadoghqComDatadogMetricV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("datadoghq.com/v1alpha1")
	model.Kind = pointer.String("DatadogMetric")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
