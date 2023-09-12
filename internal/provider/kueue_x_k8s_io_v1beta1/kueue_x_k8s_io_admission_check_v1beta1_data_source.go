/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kueue_x_k8s_io_v1beta1

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
	_ datasource.DataSource              = &KueueXK8SIoAdmissionCheckV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &KueueXK8SIoAdmissionCheckV1Beta1DataSource{}
)

func NewKueueXK8SIoAdmissionCheckV1Beta1DataSource() datasource.DataSource {
	return &KueueXK8SIoAdmissionCheckV1Beta1DataSource{}
}

type KueueXK8SIoAdmissionCheckV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type KueueXK8SIoAdmissionCheckV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ControllerName *string `tfsdk:"controller_name" json:"controllerName,omitempty"`
		Parameters     *struct {
			ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"parameters" json:"parameters,omitempty"`
		RetryDelayMinutes *int64 `tfsdk:"retry_delay_minutes" json:"retryDelayMinutes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KueueXK8SIoAdmissionCheckV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kueue_x_k8s_io_admission_check_v1beta1"
}

func (r *KueueXK8SIoAdmissionCheckV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AdmissionCheck is the Schema for the admissionchecks API",
		MarkdownDescription: "AdmissionCheck is the Schema for the admissionchecks API",
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

			"spec": schema.SingleNestedAttribute{
				Description:         "AdmissionCheckSpec defines the desired state of AdmissionCheck",
				MarkdownDescription: "AdmissionCheckSpec defines the desired state of AdmissionCheck",
				Attributes: map[string]schema.Attribute{
					"controller_name": schema.StringAttribute{
						Description:         "controllerName is name of the controller which will actually perform the checks. This is the name with which controller identifies with, not necessarily a K8S Pod or Deployment name. Cannot be empty.",
						MarkdownDescription: "controllerName is name of the controller which will actually perform the checks. This is the name with which controller identifies with, not necessarily a K8S Pod or Deployment name. Cannot be empty.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"parameters": schema.SingleNestedAttribute{
						Description:         "Parameters identifies the resource providing additional check parameters.",
						MarkdownDescription: "Parameters identifies the resource providing additional check parameters.",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "ApiGroup is the group for the resource being referenced.",
								MarkdownDescription: "ApiGroup is the group for the resource being referenced.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is the type of the resource being referenced.",
								MarkdownDescription: "Kind is the type of the resource being referenced.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the resource being referenced.",
								MarkdownDescription: "Name is the name of the resource being referenced.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"retry_delay_minutes": schema.Int64Attribute{
						Description:         "RetryDelayMinutes specifies how long to keep the workload suspended after a failed check (after it transitioned to False). After that the check state goes to 'Unknown'. The default is 15 min.",
						MarkdownDescription: "RetryDelayMinutes specifies how long to keep the workload suspended after a failed check (after it transitioned to False). After that the check state goes to 'Unknown'. The default is 15 min.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *KueueXK8SIoAdmissionCheckV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *KueueXK8SIoAdmissionCheckV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kueue_x_k8s_io_admission_check_v1beta1")

	var data KueueXK8SIoAdmissionCheckV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kueue.x-k8s.io", Version: "v1beta1", Resource: "admissionchecks"}).
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

	var readResponse KueueXK8SIoAdmissionCheckV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("kueue.x-k8s.io/v1beta1")
	data.Kind = pointer.String("AdmissionCheck")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
