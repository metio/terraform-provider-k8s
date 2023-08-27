/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package scheduling_k8s_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &SchedulingK8SIoPriorityClassV1Manifest{}
)

func NewSchedulingK8SIoPriorityClassV1Manifest() datasource.DataSource {
	return &SchedulingK8SIoPriorityClassV1Manifest{}
}

type SchedulingK8SIoPriorityClassV1Manifest struct{}

type SchedulingK8SIoPriorityClassV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Description      *string `tfsdk:"description" json:"description,omitempty"`
	GlobalDefault    *bool   `tfsdk:"global_default" json:"globalDefault,omitempty"`
	PreemptionPolicy *string `tfsdk:"preemption_policy" json:"preemptionPolicy,omitempty"`
	Value            *int64  `tfsdk:"value" json:"value,omitempty"`
}

func (r *SchedulingK8SIoPriorityClassV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_scheduling_k8s_io_priority_class_v1_manifest"
}

func (r *SchedulingK8SIoPriorityClassV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.",
		MarkdownDescription: "PriorityClass defines mapping from a priority class name to the priority integer value. The value can be any valid integer.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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

			"description": schema.StringAttribute{
				Description:         "description is an arbitrary string that usually provides guidelines on when this priority class should be used.",
				MarkdownDescription: "description is an arbitrary string that usually provides guidelines on when this priority class should be used.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"global_default": schema.BoolAttribute{
				Description:         "globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as 'globalDefault'. However, if more than one PriorityClasses exists with their 'globalDefault' field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.",
				MarkdownDescription: "globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as 'globalDefault'. However, if more than one PriorityClasses exists with their 'globalDefault' field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"preemption_policy": schema.StringAttribute{
				Description:         "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.",
				MarkdownDescription: "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.",
				Required:            false,
				Optional:            true,
				Computed:            false,
			},

			"value": schema.Int64Attribute{
				Description:         "The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.",
				MarkdownDescription: "The value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.",
				Required:            true,
				Optional:            false,
				Computed:            false,
			},
		},
	}
}

func (r *SchedulingK8SIoPriorityClassV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_scheduling_k8s_io_priority_class_v1_manifest")

	var model SchedulingK8SIoPriorityClassV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("scheduling.k8s.io/v1")
	model.Kind = pointer.String("PriorityClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
