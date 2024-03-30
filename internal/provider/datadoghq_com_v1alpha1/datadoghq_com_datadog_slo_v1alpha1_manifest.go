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
	_ datasource.DataSource = &DatadoghqComDatadogSloV1Alpha1Manifest{}
)

func NewDatadoghqComDatadogSloV1Alpha1Manifest() datasource.DataSource {
	return &DatadoghqComDatadogSloV1Alpha1Manifest{}
}

type DatadoghqComDatadogSloV1Alpha1Manifest struct{}

type DatadoghqComDatadogSloV1Alpha1ManifestData struct {
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
		ControllerOptions *struct {
			DisableRequiredTags *bool `tfsdk:"disable_required_tags" json:"disableRequiredTags,omitempty"`
		} `tfsdk:"controller_options" json:"controllerOptions,omitempty"`
		Description *string   `tfsdk:"description" json:"description,omitempty"`
		Groups      *[]string `tfsdk:"groups" json:"groups,omitempty"`
		MonitorIDs  *[]string `tfsdk:"monitor_i_ds" json:"monitorIDs,omitempty"`
		Name        *string   `tfsdk:"name" json:"name,omitempty"`
		Query       *struct {
			Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
			Numerator   *string `tfsdk:"numerator" json:"numerator,omitempty"`
		} `tfsdk:"query" json:"query,omitempty"`
		Tags             *[]string `tfsdk:"tags" json:"tags,omitempty"`
		TargetThreshold  *string   `tfsdk:"target_threshold" json:"targetThreshold,omitempty"`
		Timeframe        *string   `tfsdk:"timeframe" json:"timeframe,omitempty"`
		Type             *string   `tfsdk:"type" json:"type,omitempty"`
		WarningThreshold *string   `tfsdk:"warning_threshold" json:"warningThreshold,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DatadoghqComDatadogSloV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_datadoghq_com_datadog_slo_v1alpha1_manifest"
}

func (r *DatadoghqComDatadogSloV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DatadogSLO allows a user to define and manage datadog SLOs from Kubernetes cluster.",
		MarkdownDescription: "DatadogSLO allows a user to define and manage datadog SLOs from Kubernetes cluster.",
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
					"controller_options": schema.SingleNestedAttribute{
						Description:         "ControllerOptions are the optional parameters in the DatadogSLO controller",
						MarkdownDescription: "ControllerOptions are the optional parameters in the DatadogSLO controller",
						Attributes: map[string]schema.Attribute{
							"disable_required_tags": schema.BoolAttribute{
								Description:         "DisableRequiredTags disables the automatic addition of required tags to SLOs.",
								MarkdownDescription: "DisableRequiredTags disables the automatic addition of required tags to SLOs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"description": schema.StringAttribute{
						Description:         "Description is a user-defined description of the service level objective. Always included in service level objective responses (but may be null). Optional in create/update requests.",
						MarkdownDescription: "Description is a user-defined description of the service level objective. Always included in service level objective responses (but may be null). Optional in create/update requests.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"groups": schema.ListAttribute{
						Description:         "Groups is a list of (up to 100) monitor groups that narrow the scope of a monitor service level objective. Included in service level objective responses if it is not empty. Optional in create/update requests for monitor service level objectives, but may only be used when the length of the monitor_ids field is one.",
						MarkdownDescription: "Groups is a list of (up to 100) monitor groups that narrow the scope of a monitor service level objective. Included in service level objective responses if it is not empty. Optional in create/update requests for monitor service level objectives, but may only be used when the length of the monitor_ids field is one.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monitor_i_ds": schema.ListAttribute{
						Description:         "MonitorIDs is a list of monitor IDs that defines the scope of a monitor service level objective. Required if type is monitor.",
						MarkdownDescription: "MonitorIDs is a list of monitor IDs that defines the scope of a monitor service level objective. Required if type is monitor.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name is the name of the service level objective.",
						MarkdownDescription: "Name is the name of the service level objective.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"query": schema.SingleNestedAttribute{
						Description:         "Query is the query for a metric-based SLO. Required if type is metric. Note that only the 'sum by' aggregator is allowed, which sums all request counts. 'Average', 'max', nor 'min' request aggregators are not supported.",
						MarkdownDescription: "Query is the query for a metric-based SLO. Required if type is metric. Note that only the 'sum by' aggregator is allowed, which sums all request counts. 'Average', 'max', nor 'min' request aggregators are not supported.",
						Attributes: map[string]schema.Attribute{
							"denominator": schema.StringAttribute{
								Description:         "Denominator is a Datadog metric query for total (valid) events.",
								MarkdownDescription: "Denominator is a Datadog metric query for total (valid) events.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"numerator": schema.StringAttribute{
								Description:         "Numerator is a Datadog metric query for good events.",
								MarkdownDescription: "Numerator is a Datadog metric query for good events.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": schema.ListAttribute{
						Description:         "Tags is a list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API.",
						MarkdownDescription: "Tags is a list of tags to associate with your service level objective. This can help you categorize and filter service level objectives in the service level objectives page of the UI. Note: it's not currently possible to filter by these tags when querying via the API.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"target_threshold": schema.StringAttribute{
						Description:         "TargetThreshold is the target threshold such that when the service level indicator is above this threshold over the given timeframe, the objective is being met.",
						MarkdownDescription: "TargetThreshold is the target threshold such that when the service level indicator is above this threshold over the given timeframe, the objective is being met.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"timeframe": schema.StringAttribute{
						Description:         "The SLO time window options.",
						MarkdownDescription: "The SLO time window options.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Type is the type of the service level objective.",
						MarkdownDescription: "Type is the type of the service level objective.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"warning_threshold": schema.StringAttribute{
						Description:         "WarningThreshold is a optional warning threshold such that when the service level indicator is below this value for the given threshold, but above the target threshold, the objective appears in a 'warning' state. This value must be greater than the target threshold.",
						MarkdownDescription: "WarningThreshold is a optional warning threshold such that when the service level indicator is below this value for the given threshold, but above the target threshold, the objective appears in a 'warning' state. This value must be greater than the target threshold.",
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

func (r *DatadoghqComDatadogSloV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_datadoghq_com_datadog_slo_v1alpha1_manifest")

	var model DatadoghqComDatadogSloV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("datadoghq.com/v1alpha1")
	model.Kind = pointer.String("DatadogSLO")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
