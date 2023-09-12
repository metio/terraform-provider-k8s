/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package scheduling_k8s_io_v1

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &SchedulingK8SIoPriorityClassV1DataSource{}
	_ datasource.DataSourceWithConfigure = &SchedulingK8SIoPriorityClassV1DataSource{}
)

func NewSchedulingK8SIoPriorityClassV1DataSource() datasource.DataSource {
	return &SchedulingK8SIoPriorityClassV1DataSource{}
}

type SchedulingK8SIoPriorityClassV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type SchedulingK8SIoPriorityClassV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *SchedulingK8SIoPriorityClassV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_scheduling_k8s_io_priority_class_v1"
}

func (r *SchedulingK8SIoPriorityClassV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"description": schema.StringAttribute{
				Description:         "description is an arbitrary string that usually provides guidelines on when this priority class should be used.",
				MarkdownDescription: "description is an arbitrary string that usually provides guidelines on when this priority class should be used.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"global_default": schema.BoolAttribute{
				Description:         "globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as 'globalDefault'. However, if more than one PriorityClasses exists with their 'globalDefault' field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.",
				MarkdownDescription: "globalDefault specifies whether this PriorityClass should be considered as the default priority for pods that do not have any priority class. Only one PriorityClass can be marked as 'globalDefault'. However, if more than one PriorityClasses exists with their 'globalDefault' field set to true, the smallest value of such global default PriorityClasses will be used as the default priority.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"preemption_policy": schema.StringAttribute{
				Description:         "preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.",
				MarkdownDescription: "preemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"value": schema.Int64Attribute{
				Description:         "value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.",
				MarkdownDescription: "value represents the integer value of this priority class. This is the actual priority that pods receive when they have the name of this class in their pod spec.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},
		},
	}
}

func (r *SchedulingK8SIoPriorityClassV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *SchedulingK8SIoPriorityClassV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_scheduling_k8s_io_priority_class_v1")

	var data SchedulingK8SIoPriorityClassV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "scheduling.k8s.io", Version: "v1", Resource: "priorityclasses"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetResourceError(err, data.Metadata.Name))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse SchedulingK8SIoPriorityClassV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("scheduling.k8s.io/v1")
	data.Kind = pointer.String("PriorityClass")
	data.Metadata = readResponse.Metadata
	data.Description = readResponse.Description
	data.GlobalDefault = readResponse.GlobalDefault
	data.PreemptionPolicy = readResponse.PreemptionPolicy
	data.Value = readResponse.Value

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
