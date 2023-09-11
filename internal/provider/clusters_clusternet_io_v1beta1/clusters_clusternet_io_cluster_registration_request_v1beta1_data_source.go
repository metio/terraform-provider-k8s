/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package clusters_clusternet_io_v1beta1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"net/http"
)

var (
	_ datasource.DataSource              = &ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource{}
)

func NewClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource() datasource.DataSource {
	return &ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource{}
}

type ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		ClusterId        *string            `tfsdk:"cluster_id" json:"clusterId,omitempty"`
		ClusterLabels    *map[string]string `tfsdk:"cluster_labels" json:"clusterLabels,omitempty"`
		ClusterName      *string            `tfsdk:"cluster_name" json:"clusterName,omitempty"`
		ClusterNamespace *string            `tfsdk:"cluster_namespace" json:"clusterNamespace,omitempty"`
		ClusterType      *string            `tfsdk:"cluster_type" json:"clusterType,omitempty"`
		SyncMode         *string            `tfsdk:"sync_mode" json:"syncMode,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_clusters_clusternet_io_cluster_registration_request_v1beta1"
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterRegistrationRequest is the Schema for the clusterregistrationrequests API",
		MarkdownDescription: "ClusterRegistrationRequest is the Schema for the clusterregistrationrequests API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "ClusterRegistrationRequestSpec defines the desired state of ClusterRegistrationRequest",
				MarkdownDescription: "ClusterRegistrationRequestSpec defines the desired state of ClusterRegistrationRequest",
				Attributes: map[string]schema.Attribute{
					"cluster_id": schema.StringAttribute{
						Description:         "ClusterID, a Random (Version 4) UUID, is a unique value in time and space value representing for child cluster. It is typically generated by the clusternet agent on the successful creation of a 'clusternet-agent' Lease in the child cluster. Also it is not allowed to change on PUT operations.",
						MarkdownDescription: "ClusterID, a Random (Version 4) UUID, is a unique value in time and space value representing for child cluster. It is typically generated by the clusternet agent on the successful creation of a 'clusternet-agent' Lease in the child cluster. Also it is not allowed to change on PUT operations.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_labels": schema.MapAttribute{
						Description:         "ClusterLabels is the labels of the child cluster.",
						MarkdownDescription: "ClusterLabels is the labels of the child cluster.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_name": schema.StringAttribute{
						Description:         "ClusterName is the cluster name. a lower case alphanumeric characters or '-', and must start and end with an alphanumeric character",
						MarkdownDescription: "ClusterName is the cluster name. a lower case alphanumeric characters or '-', and must start and end with an alphanumeric character",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_namespace": schema.StringAttribute{
						Description:         "ClusterNamespace is the dedicated namespace of the cluster.",
						MarkdownDescription: "ClusterNamespace is the dedicated namespace of the cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"cluster_type": schema.StringAttribute{
						Description:         "ClusterType denotes the type of the child cluster.",
						MarkdownDescription: "ClusterType denotes the type of the child cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"sync_mode": schema.StringAttribute{
						Description:         "SyncMode decides how to sync resources from parent cluster to child cluster.",
						MarkdownDescription: "SyncMode decides how to sync resources from parent cluster to child cluster.",
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

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_clusters_clusternet_io_cluster_registration_request_v1beta1")

	var data ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "clusters.clusternet.io", Version: "v1beta1", Resource: "clusterregistrationrequests"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		var statusError *k8sErrors.StatusError
		if errors.As(err, &statusError) {
			if statusError.Status().Code == http.StatusNotFound {
				response.Diagnostics.AddError(
					"Unable to find resource",
					fmt.Sprintf("The requested resource cannot be found. "+
						"Make sure that it does exist in your cluster and you have set the correct name configured.\n\n"+
						"Name: %s", data.Metadata.Name),
				)
				return
			}
		} else {
			response.Diagnostics.AddError(
				"Unable to GET resource",
				fmt.Sprintf("An unexpected error occurred while reading the resource. "+
					"Please report this issue to the provider developers.\n\n"+
					"GET Error (%T): %s", err, err.Error()),
			)
		}
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse ClustersClusternetIoClusterRegistrationRequestV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(data.Metadata.Name)
	data.ApiVersion = pointer.String("clusters.clusternet.io/v1beta1")
	data.Kind = pointer.String("ClusterRegistrationRequest")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
