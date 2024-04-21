/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package infrastructure_cluster_x_k8s_io_v1beta2

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
	_ datasource.DataSource = &InfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2Manifest{}
)

func NewInfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2Manifest() datasource.DataSource {
	return &InfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2Manifest{}
}

type InfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2Manifest struct{}

type InfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2ManifestData struct {
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
		Template *struct {
			Spec *struct {
				BootVolume *struct {
					DeleteVolumeOnInstanceDelete *bool   `tfsdk:"delete_volume_on_instance_delete" json:"deleteVolumeOnInstanceDelete,omitempty"`
					EncryptionKeyCRN             *string `tfsdk:"encryption_key_crn" json:"encryptionKeyCRN,omitempty"`
					Iops                         *int64  `tfsdk:"iops" json:"iops,omitempty"`
					Name                         *string `tfsdk:"name" json:"name,omitempty"`
					Profile                      *string `tfsdk:"profile" json:"profile,omitempty"`
					SizeGiB                      *int64  `tfsdk:"size_gi_b" json:"sizeGiB,omitempty"`
				} `tfsdk:"boot_volume" json:"bootVolume,omitempty"`
				Image *struct {
					Id   *string `tfsdk:"id" json:"id,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"image" json:"image,omitempty"`
				Name                    *string `tfsdk:"name" json:"name,omitempty"`
				PrimaryNetworkInterface *struct {
					Subnet *string `tfsdk:"subnet" json:"subnet,omitempty"`
				} `tfsdk:"primary_network_interface" json:"primaryNetworkInterface,omitempty"`
				Profile    *string `tfsdk:"profile" json:"profile,omitempty"`
				ProviderID *string `tfsdk:"provider_id" json:"providerID,omitempty"`
				SshKeys    *[]struct {
					Id   *string `tfsdk:"id" json:"id,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"ssh_keys" json:"sshKeys,omitempty"`
				Zone *string `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *InfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta2_manifest"
}

func (r *InfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "IBMVPCMachineTemplate is the Schema for the ibmvpcmachinetemplates API.",
		MarkdownDescription: "IBMVPCMachineTemplate is the Schema for the ibmvpcmachinetemplates API.",
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
				Description:         "IBMVPCMachineTemplateSpec defines the desired state of IBMVPCMachineTemplate.",
				MarkdownDescription: "IBMVPCMachineTemplateSpec defines the desired state of IBMVPCMachineTemplate.",
				Attributes: map[string]schema.Attribute{
					"template": schema.SingleNestedAttribute{
						Description:         "IBMVPCMachineTemplateResource describes the data needed to create am IBMVPCMachine from a template.",
						MarkdownDescription: "IBMVPCMachineTemplateResource describes the data needed to create am IBMVPCMachine from a template.",
						Attributes: map[string]schema.Attribute{
							"spec": schema.SingleNestedAttribute{
								Description:         "Spec is the specification of the desired behavior of the machine.",
								MarkdownDescription: "Spec is the specification of the desired behavior of the machine.",
								Attributes: map[string]schema.Attribute{
									"boot_volume": schema.SingleNestedAttribute{
										Description:         "BootVolume contains machines's boot volume configurations like size, iops etc..",
										MarkdownDescription: "BootVolume contains machines's boot volume configurations like size, iops etc..",
										Attributes: map[string]schema.Attribute{
											"delete_volume_on_instance_delete": schema.BoolAttribute{
												Description:         "DeleteVolumeOnInstanceDelete If set to true, when deleting the instance the volume will also be deleted.Default is set as true",
												MarkdownDescription: "DeleteVolumeOnInstanceDelete If set to true, when deleting the instance the volume will also be deleted.Default is set as true",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"encryption_key_crn": schema.StringAttribute{
												Description:         "EncryptionKey is the root key to use to wrap the data encryption key for the volume and this points to the CRNand possible values are as follows.The CRN of the [Key Protect RootKey](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect CryptoService Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.If unspecified, the 'encryption' type for the volume will be 'provider_managed'.",
												MarkdownDescription: "EncryptionKey is the root key to use to wrap the data encryption key for the volume and this points to the CRNand possible values are as follows.The CRN of the [Key Protect RootKey](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect CryptoService Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.If unspecified, the 'encryption' type for the volume will be 'provider_managed'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"iops": schema.Int64Attribute{
												Description:         "Iops is the maximum I/O operations per second (IOPS) to use for the volume. Applicable only to volumes using a profilefamily of 'custom'.",
												MarkdownDescription: "Iops is the maximum I/O operations per second (IOPS) to use for the volume. Applicable only to volumes using a profilefamily of 'custom'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the unique user-defined name for this volume.Default will be autogenerated",
												MarkdownDescription: "Name is the unique user-defined name for this volume.Default will be autogenerated",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"profile": schema.StringAttribute{
												Description:         "Profile is the volume profile for the bootdisk, refer https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-profilesfor more information.Default to general-purpose",
												MarkdownDescription: "Profile is the volume profile for the bootdisk, refer https://cloud.ibm.com/docs/vpc?topic=vpc-block-storage-profilesfor more information.Default to general-purpose",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("general-purpose", "5iops-tier", "10iops-tier", "custom"),
												},
											},

											"size_gi_b": schema.Int64Attribute{
												Description:         "SizeGiB is the size of the virtual server's boot disk in GiB.Default to the size of the image's 'minimum_provisioned_size'.",
												MarkdownDescription: "SizeGiB is the size of the virtual server's boot disk in GiB.Default to the size of the image's 'minimum_provisioned_size'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": schema.SingleNestedAttribute{
										Description:         "Image is the OS image which would be install on the instance.ID will take higher precedence over Name if both specified.",
										MarkdownDescription: "Image is the OS image which would be install on the instance.ID will take higher precedence over Name if both specified.",
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Description:         "ID of resource",
												MarkdownDescription: "ID of resource",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Name of resource",
												MarkdownDescription: "Name of resource",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the instance.",
										MarkdownDescription: "Name of the instance.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"primary_network_interface": schema.SingleNestedAttribute{
										Description:         "PrimaryNetworkInterface is required to specify subnet.",
										MarkdownDescription: "PrimaryNetworkInterface is required to specify subnet.",
										Attributes: map[string]schema.Attribute{
											"subnet": schema.StringAttribute{
												Description:         "Subnet ID of the network interface.",
												MarkdownDescription: "Subnet ID of the network interface.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"profile": schema.StringAttribute{
										Description:         "Profile indicates the flavor of instance. Example: bx2-8x32	means 8 vCPUs	32 GB RAM	16 GbpsTODO: add a reference link of profile",
										MarkdownDescription: "Profile indicates the flavor of instance. Example: bx2-8x32	means 8 vCPUs	32 GB RAM	16 GbpsTODO: add a reference link of profile",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"provider_id": schema.StringAttribute{
										Description:         "ProviderID is the unique identifier as specified by the cloud provider.",
										MarkdownDescription: "ProviderID is the unique identifier as specified by the cloud provider.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ssh_keys": schema.ListNestedAttribute{
										Description:         "SSHKeys is the SSH pub keys that will be used to access VM.ID will take higher precedence over Name if both specified.",
										MarkdownDescription: "SSHKeys is the SSH pub keys that will be used to access VM.ID will take higher precedence over Name if both specified.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.StringAttribute{
													Description:         "ID of resource",
													MarkdownDescription: "ID of resource",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},

												"name": schema.StringAttribute{
													Description:         "Name of resource",
													MarkdownDescription: "Name of resource",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.LengthAtLeast(1),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"zone": schema.StringAttribute{
										Description:         "Zone is the place where the instance should be created. Example: us-south-3TODO: Actually zone is transparent to user. The field user can access is location. Example: Dallas 2",
										MarkdownDescription: "Zone is the place where the instance should be created. Example: us-south-3TODO: Actually zone is transparent to user. The field user can access is location. Example: Dallas 2",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},
						},
						Required: true,
						Optional: false,
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

func (r *InfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_infrastructure_cluster_x_k8s_io_ibmvpc_machine_template_v1beta2_manifest")

	var model InfrastructureClusterXK8SIoIbmvpcmachineTemplateV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("infrastructure.cluster.x-k8s.io/v1beta2")
	model.Kind = pointer.String("IBMVPCMachineTemplate")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
