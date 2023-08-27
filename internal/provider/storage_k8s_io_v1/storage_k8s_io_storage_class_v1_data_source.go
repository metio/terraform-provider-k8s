/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package storage_k8s_io_v1

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
	_ datasource.DataSource              = &StorageK8SIoStorageClassV1DataSource{}
	_ datasource.DataSourceWithConfigure = &StorageK8SIoStorageClassV1DataSource{}
)

func NewStorageK8SIoStorageClassV1DataSource() datasource.DataSource {
	return &StorageK8SIoStorageClassV1DataSource{}
}

type StorageK8SIoStorageClassV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type StorageK8SIoStorageClassV1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	AllowVolumeExpansion *bool `tfsdk:"allow_volume_expansion" json:"allowVolumeExpansion,omitempty"`
	AllowedTopologies    *[]struct {
		MatchLabelExpressions *[]struct {
			Key    *string   `tfsdk:"key" json:"key,omitempty"`
			Values *[]string `tfsdk:"values" json:"values,omitempty"`
		} `tfsdk:"match_label_expressions" json:"matchLabelExpressions,omitempty"`
	} `tfsdk:"allowed_topologies" json:"allowedTopologies,omitempty"`
	MountOptions      *[]string          `tfsdk:"mount_options" json:"mountOptions,omitempty"`
	Parameters        *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
	Provisioner       *string            `tfsdk:"k8s_provisioner" json:"provisioner,omitempty"`
	ReclaimPolicy     *string            `tfsdk:"reclaim_policy" json:"reclaimPolicy,omitempty"`
	VolumeBindingMode *string            `tfsdk:"volume_binding_mode" json:"volumeBindingMode,omitempty"`
}

func (r *StorageK8SIoStorageClassV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_storage_k8s_io_storage_class_v1"
}

func (r *StorageK8SIoStorageClassV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
		MarkdownDescription: "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
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

			"allow_volume_expansion": schema.BoolAttribute{
				Description:         "AllowVolumeExpansion shows whether the storage class allow volume expand",
				MarkdownDescription: "AllowVolumeExpansion shows whether the storage class allow volume expand",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"allowed_topologies": schema.ListNestedAttribute{
				Description:         "Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.",
				MarkdownDescription: "Restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"match_label_expressions": schema.ListNestedAttribute{
							Description:         "A list of topology selector requirements by labels.",
							MarkdownDescription: "A list of topology selector requirements by labels.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The label key that the selector applies to.",
										MarkdownDescription: "The label key that the selector applies to.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"values": schema.ListAttribute{
										Description:         "An array of string values. One value must match the label to be selected. Each entry in Values is ORed.",
										MarkdownDescription: "An array of string values. One value must match the label to be selected. Each entry in Values is ORed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
							},
							Required: false,
							Optional: false,
							Computed: true,
						},
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},

			"mount_options": schema.ListAttribute{
				Description:         "Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ['ro', 'soft']. Not validated - mount of the PVs will simply fail if one is invalid.",
				MarkdownDescription: "Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. ['ro', 'soft']. Not validated - mount of the PVs will simply fail if one is invalid.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"parameters": schema.MapAttribute{
				Description:         "Parameters holds the parameters for the provisioner that should create volumes of this storage class.",
				MarkdownDescription: "Parameters holds the parameters for the provisioner that should create volumes of this storage class.",
				ElementType:         types.StringType,
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"k8s_provisioner": schema.StringAttribute{
				Description:         "Provisioner indicates the type of the provisioner.",
				MarkdownDescription: "Provisioner indicates the type of the provisioner.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"reclaim_policy": schema.StringAttribute{
				Description:         "Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.",
				MarkdownDescription: "Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"volume_binding_mode": schema.StringAttribute{
				Description:         "VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",
				MarkdownDescription: "VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},
		},
	}
}

func (r *StorageK8SIoStorageClassV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *StorageK8SIoStorageClassV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_storage_k8s_io_storage_class_v1")

	var data StorageK8SIoStorageClassV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "storage.k8s.io", Version: "v1", Resource: "StorageClass"}).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
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

	var readResponse StorageK8SIoStorageClassV1DataSourceData
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
	data.ApiVersion = pointer.String("storage.k8s.io/v1")
	data.Kind = pointer.String("StorageClass")
	data.Metadata = readResponse.Metadata
	data.AllowVolumeExpansion = readResponse.AllowVolumeExpansion
	data.AllowedTopologies = readResponse.AllowedTopologies
	data.MountOptions = readResponse.MountOptions
	data.Parameters = readResponse.Parameters
	data.Provisioner = readResponse.Provisioner
	data.ReclaimPolicy = readResponse.ReclaimPolicy
	data.VolumeBindingMode = readResponse.VolumeBindingMode

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
