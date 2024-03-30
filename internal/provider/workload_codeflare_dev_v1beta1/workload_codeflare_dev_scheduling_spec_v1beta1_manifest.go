/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package workload_codeflare_dev_v1beta1

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
	_ datasource.DataSource = &WorkloadCodeflareDevSchedulingSpecV1Beta1Manifest{}
)

func NewWorkloadCodeflareDevSchedulingSpecV1Beta1Manifest() datasource.DataSource {
	return &WorkloadCodeflareDevSchedulingSpecV1Beta1Manifest{}
}

type WorkloadCodeflareDevSchedulingSpecV1Beta1Manifest struct{}

type WorkloadCodeflareDevSchedulingSpecV1Beta1ManifestData struct {
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
		DispatchDuration *struct {
			Expected *int64 `tfsdk:"expected" json:"expected,omitempty"`
			Limit    *int64 `tfsdk:"limit" json:"limit,omitempty"`
			Overrun  *bool  `tfsdk:"overrun" json:"overrun,omitempty"`
		} `tfsdk:"dispatch_duration" json:"dispatchDuration,omitempty"`
		MinAvailable *int64             `tfsdk:"min_available" json:"minAvailable,omitempty"`
		NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Requeuing    *struct {
			GrowthType           *string `tfsdk:"growth_type" json:"growthType,omitempty"`
			InitialTimeInSeconds *int64  `tfsdk:"initial_time_in_seconds" json:"initialTimeInSeconds,omitempty"`
			MaxNumRequeuings     *int64  `tfsdk:"max_num_requeuings" json:"maxNumRequeuings,omitempty"`
			MaxTimeInSeconds     *int64  `tfsdk:"max_time_in_seconds" json:"maxTimeInSeconds,omitempty"`
			NumRequeuings        *int64  `tfsdk:"num_requeuings" json:"numRequeuings,omitempty"`
			TimeInSeconds        *int64  `tfsdk:"time_in_seconds" json:"timeInSeconds,omitempty"`
		} `tfsdk:"requeuing" json:"requeuing,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *WorkloadCodeflareDevSchedulingSpecV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_workload_codeflare_dev_scheduling_spec_v1beta1_manifest"
}

func (r *WorkloadCodeflareDevSchedulingSpecV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"dispatch_duration": schema.SingleNestedAttribute{
						Description:         "Wall clock duration time of appwrapper in seconds.",
						MarkdownDescription: "Wall clock duration time of appwrapper in seconds.",
						Attributes: map[string]schema.Attribute{
							"expected": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"limit": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"overrun": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"min_available": schema.Int64Attribute{
						Description:         "Expected number of pods in running and/or completed state. Requeuing is triggered when the number of running/completed pods is not equal to this value. When not specified, requeuing is disabled and no check is performed.",
						MarkdownDescription: "Expected number of pods in running and/or completed state. Requeuing is triggered when the number of running/completed pods is not equal to this value. When not specified, requeuing is disabled and no check is performed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"requeuing": schema.SingleNestedAttribute{
						Description:         "Specification of the requeuing strategy based on waiting time. Values in this field control how often the pod check should happen, and if requeuing has reached its maximum number of times.",
						MarkdownDescription: "Specification of the requeuing strategy based on waiting time. Values in this field control how often the pod check should happen, and if requeuing has reached its maximum number of times.",
						Attributes: map[string]schema.Attribute{
							"growth_type": schema.StringAttribute{
								Description:         "Growth strategy to increase the waiting time between requeuing checks. The values available are 'exponential', 'linear', or 'none'. For example, 'exponential' growth would double the 'timeInSeconds' value every time a requeuing event is triggered. If the string value is misspelled or not one of the possible options, the growth behavior is defaulted to 'none'.",
								MarkdownDescription: "Growth strategy to increase the waiting time between requeuing checks. The values available are 'exponential', 'linear', or 'none'. For example, 'exponential' growth would double the 'timeInSeconds' value every time a requeuing event is triggered. If the string value is misspelled or not one of the possible options, the growth behavior is defaulted to 'none'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"initial_time_in_seconds": schema.Int64Attribute{
								Description:         "Value to keep track of the initial wait time. Users cannot set this as it is taken from 'timeInSeconds'.",
								MarkdownDescription: "Value to keep track of the initial wait time. Users cannot set this as it is taken from 'timeInSeconds'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_num_requeuings": schema.Int64Attribute{
								Description:         "Maximum number of requeuing events allowed. Once this value is reached (e.g., 'numRequeuings = maxNumRequeuings', no more requeuing checks are performed and the generic items are stopped and removed from the cluster (AppWrapper remains deployed).",
								MarkdownDescription: "Maximum number of requeuing events allowed. Once this value is reached (e.g., 'numRequeuings = maxNumRequeuings', no more requeuing checks are performed and the generic items are stopped and removed from the cluster (AppWrapper remains deployed).",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_time_in_seconds": schema.Int64Attribute{
								Description:         "Maximum waiting time for requeuing checks.",
								MarkdownDescription: "Maximum waiting time for requeuing checks.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"num_requeuings": schema.Int64Attribute{
								Description:         "Field to keep track of how many times a requeuing event has been triggered.",
								MarkdownDescription: "Field to keep track of how many times a requeuing event has been triggered.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"time_in_seconds": schema.Int64Attribute{
								Description:         "Initial waiting time before requeuing conditions are checked. This value is specified by the user, but it may grow as requeuing events happen.",
								MarkdownDescription: "Initial waiting time before requeuing conditions are checked. This value is specified by the user, but it may grow as requeuing events happen.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *WorkloadCodeflareDevSchedulingSpecV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_workload_codeflare_dev_scheduling_spec_v1beta1_manifest")

	var model WorkloadCodeflareDevSchedulingSpecV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("workload.codeflare.dev/v1beta1")
	model.Kind = pointer.String("SchedulingSpec")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
