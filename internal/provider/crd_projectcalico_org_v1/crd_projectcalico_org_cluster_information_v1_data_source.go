/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package crd_projectcalico_org_v1

import (
	"context"
	"encoding/json"
	"fmt"
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
	_ datasource.DataSource              = &CrdProjectcalicoOrgClusterInformationV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CrdProjectcalicoOrgClusterInformationV1DataSource{}
)

func NewCrdProjectcalicoOrgClusterInformationV1DataSource() datasource.DataSource {
	return &CrdProjectcalicoOrgClusterInformationV1DataSource{}
}

type CrdProjectcalicoOrgClusterInformationV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CrdProjectcalicoOrgClusterInformationV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		CalicoVersion  *string `tfsdk:"calico_version" json:"calicoVersion,omitempty"`
		ClusterGUID    *string `tfsdk:"cluster_guid" json:"clusterGUID,omitempty"`
		ClusterType    *string `tfsdk:"cluster_type" json:"clusterType,omitempty"`
		DatastoreReady *bool   `tfsdk:"datastore_ready" json:"datastoreReady,omitempty"`
		Variant        *string `tfsdk:"variant" json:"variant,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CrdProjectcalicoOrgClusterInformationV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_crd_projectcalico_org_cluster_information_v1"
}

func (r *CrdProjectcalicoOrgClusterInformationV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterInformation contains the cluster specific information.",
		MarkdownDescription: "ClusterInformation contains the cluster specific information.",
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
				Description:         "ClusterInformationSpec contains the values of describing the cluster.",
				MarkdownDescription: "ClusterInformationSpec contains the values of describing the cluster.",
				Attributes: map[string]schema.Attribute{
					"calico_version": schema.StringAttribute{
						Description:         "CalicoVersion is the version of Calico that the cluster is running",
						MarkdownDescription: "CalicoVersion is the version of Calico that the cluster is running",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_guid": schema.StringAttribute{
						Description:         "ClusterGUID is the GUID of the cluster",
						MarkdownDescription: "ClusterGUID is the GUID of the cluster",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_type": schema.StringAttribute{
						Description:         "ClusterType describes the type of the cluster",
						MarkdownDescription: "ClusterType describes the type of the cluster",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"datastore_ready": schema.BoolAttribute{
						Description:         "DatastoreReady is used during significant datastore migrations to signal to components such as Felix that it should wait before accessing the datastore.",
						MarkdownDescription: "DatastoreReady is used during significant datastore migrations to signal to components such as Felix that it should wait before accessing the datastore.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"variant": schema.StringAttribute{
						Description:         "Variant declares which variant of Calico should be active.",
						MarkdownDescription: "Variant declares which variant of Calico should be active.",
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

func (r *CrdProjectcalicoOrgClusterInformationV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CrdProjectcalicoOrgClusterInformationV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_crd_projectcalico_org_cluster_information_v1")

	var data CrdProjectcalicoOrgClusterInformationV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "v1", Resource: "clusterinformations"}).
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

	var readResponse CrdProjectcalicoOrgClusterInformationV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("crd.projectcalico.org/v1")
	data.Kind = pointer.String("ClusterInformation")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
