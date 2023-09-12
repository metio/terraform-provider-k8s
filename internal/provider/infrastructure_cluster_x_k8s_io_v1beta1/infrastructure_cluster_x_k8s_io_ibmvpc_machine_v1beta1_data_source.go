/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta1

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
	_ datasource.DataSource              = &InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource{}
	_ datasource.DataSourceWithConfigure = &InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource{}
)

func NewInfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource{}
}

type InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource struct {
	kubernetesClient dynamic.Interface
}

type InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		BootVolume *struct {
			DeleteVolumeOnInstanceDelete *bool   `tfsdk:"delete_volume_on_instance_delete" json:"deleteVolumeOnInstanceDelete,omitempty"`
			EncryptionKeyCRN             *string `tfsdk:"encryption_key_crn" json:"encryptionKeyCRN,omitempty"`
			Iops                         *int64  `tfsdk:"iops" json:"iops,omitempty"`
			Name                         *string `tfsdk:"name" json:"name,omitempty"`
			Profile                      *string `tfsdk:"profile" json:"profile,omitempty"`
			SizeGiB                      *int64  `tfsdk:"size_gi_b" json:"sizeGiB,omitempty"`
		} `tfsdk:"boot_volume" json:"bootVolume,omitempty"`
		Image                   *string `tfsdk:"image" json:"image,omitempty"`
		ImageName               *string `tfsdk:"image_name" json:"imageName,omitempty"`
		Name                    *string `tfsdk:"name" json:"name,omitempty"`
		PrimaryNetworkInterface *struct {
			Subnet *string `tfsdk:"subnet" json:"subnet,omitempty"`
		} `tfsdk:"primary_network_interface" json:"primaryNetworkInterface,omitempty"`
		Profile     *string   `tfsdk:"profile" json:"profile,omitempty"`
		ProviderID  *string   `tfsdk:"provider_id" json:"providerID,omitempty"`
		SshKeyNames *[]string `tfsdk:"ssh_key_names" json:"sshKeyNames,omitempty"`
		SshKeys     *[]string `tfsdk:"ssh_keys" json:"sshKeys,omitempty"`
		Zone        *string   `tfsdk:"zone" json:"zone,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta1"
}

func (r *InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMVPCMachine is the Schema for the ibmvpcmachines API.",
		MarkdownDescription: "IBMVPCMachine is the Schema for the ibmvpcmachines API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "IBMVPCMachineSpec defines the desired state of IBMVPCMachine.",
				MarkdownDescription: "IBMVPCMachineSpec defines the desired state of IBMVPCMachine.",
				Attributes: map[string]schema.Attribute{
					"boot_volume": schema.SingleNestedAttribute{
						Description:         "BootVolume contains machines's boot volume configurations like size, iops etc..",
						MarkdownDescription: "BootVolume contains machines's boot volume configurations like size, iops etc..",
						Attributes: map[string]schema.Attribute{
							"delete_volume_on_instance_delete": schema.BoolAttribute{
								Description:         "DeleteVolumeOnInstanceDelete If set to true, when deleting the instance the volume will also be deleted. Default is set as true",
								MarkdownDescription: "DeleteVolumeOnInstanceDelete If set to true, when deleting the instance the volume will also be deleted. Default is set as true",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"encryption_key_crn": schema.StringAttribute{
								Description:         "EncryptionKey is the root key to use to wrap the data encryption key for the volume and this points to the CRN and possible values are as follows. The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource. If unspecified, the 'encryption' type for the volume will be 'provider_managed'.",
								MarkdownDescription: "EncryptionKey is the root key to use to wrap the data encryption key for the volume and this points to the CRN and possible values are as follows. The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Service Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource. If unspecified, the 'encryption' type for the volume will be 'provider_managed'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"iops": schema.Int64Attribute{
								Description:         "Iops is the maximum I/O operations per second (IOPS) to use for the volume. Applicable only to volumes using a profile family of 'custom'.",
								MarkdownDescription: "Iops is the maximum I/O operations per second (IOPS) to use for the volume. Applicable only to volumes using a profile family of 'custom'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the unique user-defined name for this volume. Default will be autogenerated",
								MarkdownDescription: "Name is the unique user-defined name for this volume. Default will be autogenerated",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"profile": schema.StringAttribute{
								Description:         "Profile is the volume profile for the bootdisk, refer https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-profiles for more information. Default to general-purpose",
								MarkdownDescription: "Profile is the volume profile for the bootdisk, refer https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-profiles for more information. Default to general-purpose",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"size_gi_b": schema.Int64Attribute{
								Description:         "SizeGiB is the size of the virtual server's boot disk in GiB. Default to the size of the image's 'minimum_provisioned_size'.",
								MarkdownDescription: "SizeGiB is the size of the virtual server's boot disk in GiB. Default to the size of the image's 'minimum_provisioned_size'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image": schema.StringAttribute{
						Description:         "Image is the id of OS image which would be install on the instance. Example: r134-ed3f775f-ad7e-4e37-ae62-7199b4988b00",
						MarkdownDescription: "Image is the id of OS image which would be install on the instance. Example: r134-ed3f775f-ad7e-4e37-ae62-7199b4988b00",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"image_name": schema.StringAttribute{
						Description:         "ImageName is the name of OS image which would be install on the instance.",
						MarkdownDescription: "ImageName is the name of OS image which would be install on the instance.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the instance.",
						MarkdownDescription: "Name of the instance.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"primary_network_interface": schema.SingleNestedAttribute{
						Description:         "PrimaryNetworkInterface is required to specify subnet.",
						MarkdownDescription: "PrimaryNetworkInterface is required to specify subnet.",
						Attributes: map[string]schema.Attribute{
							"subnet": schema.StringAttribute{
								Description:         "Subnet ID of the network interface.",
								MarkdownDescription: "Subnet ID of the network interface.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"profile": schema.StringAttribute{
						Description:         "Profile indicates the flavor of instance. Example: bx2-8x32	means 8 vCPUs	32 GB RAM	16 Gbps TODO: add a reference link of profile",
						MarkdownDescription: "Profile indicates the flavor of instance. Example: bx2-8x32	means 8 vCPUs	32 GB RAM	16 Gbps TODO: add a reference link of profile",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"provider_id": schema.StringAttribute{
						Description:         "ProviderID is the unique identifier as specified by the cloud provider.",
						MarkdownDescription: "ProviderID is the unique identifier as specified by the cloud provider.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ssh_key_names": schema.ListAttribute{
						Description:         "SSHKeysNames is the SSH pub key names that will be used to access VM.",
						MarkdownDescription: "SSHKeysNames is the SSH pub key names that will be used to access VM.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ssh_keys": schema.ListAttribute{
						Description:         "SSHKeys is the SSH pub keys that will be used to access VM.",
						MarkdownDescription: "SSHKeys is the SSH pub keys that will be used to access VM.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"zone": schema.StringAttribute{
						Description:         "Zone is the place where the instance should be created. Example: us-south-3 TODO: Actually zone is transparent to user. The field user can access is location. Example: Dallas 2",
						MarkdownDescription: "Zone is the place where the instance should be created. Example: us-south-3 TODO: Actually zone is transparent to user. The field user can access is location. Example: Dallas 2",
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

func (r *InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_infrastructure_cluster_x_k8s_io_ibmvpc_machine_v1beta1")

	var data InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "infrastructure.cluster.x-k8s.io", Version: "v1beta1", Resource: "ibmvpcmachines"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse InfrastructureClusterXK8SIoIbmvpcmachineV1Beta1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta1")
	data.Kind = pointer.String("IBMVPCMachine")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
