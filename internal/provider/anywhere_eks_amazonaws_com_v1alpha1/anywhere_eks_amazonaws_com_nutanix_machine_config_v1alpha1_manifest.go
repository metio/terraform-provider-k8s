/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &AnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1ManifestData struct {
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
		AdditionalCategories *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"additional_categories" json:"additionalCategories,omitempty"`
		Cluster *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
			Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
		} `tfsdk:"cluster" json:"cluster,omitempty"`
		Gpus *[]struct {
			DeviceID *int64  `tfsdk:"device_id" json:"deviceID,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Type     *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"gpus" json:"gpus,omitempty"`
		Image *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
			Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		MemorySize *string `tfsdk:"memory_size" json:"memorySize,omitempty"`
		OsFamily   *string `tfsdk:"os_family" json:"osFamily,omitempty"`
		Project    *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
			Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
		} `tfsdk:"project" json:"project,omitempty"`
		Subnet *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
			Uuid *string `tfsdk:"uuid" json:"uuid,omitempty"`
		} `tfsdk:"subnet" json:"subnet,omitempty"`
		SystemDiskSize *string `tfsdk:"system_disk_size" json:"systemDiskSize,omitempty"`
		Users          *[]struct {
			Name              *string   `tfsdk:"name" json:"name,omitempty"`
			SshAuthorizedKeys *[]string `tfsdk:"ssh_authorized_keys" json:"sshAuthorizedKeys,omitempty"`
		} `tfsdk:"users" json:"users,omitempty"`
		VcpuSockets    *int64 `tfsdk:"vcpu_sockets" json:"vcpuSockets,omitempty"`
		VcpusPerSocket *int64 `tfsdk:"vcpus_per_socket" json:"vcpusPerSocket,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NutanixMachineConfig is the Schema for the nutanix machine configs API",
		MarkdownDescription: "NutanixMachineConfig is the Schema for the nutanix machine configs API",
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
				Description:         "NutanixMachineConfigSpec defines the desired state of NutanixMachineConfig.",
				MarkdownDescription: "NutanixMachineConfigSpec defines the desired state of NutanixMachineConfig.",
				Attributes: map[string]schema.Attribute{
					"additional_categories": schema.ListNestedAttribute{
						Description:         "additionalCategories is a list of optional categories to be added to the VM. Categories must be created in Prism Central before they can be used.",
						MarkdownDescription: "additionalCategories is a list of optional categories to be added to the VM. Categories must be created in Prism Central before they can be used.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "key is the Key of the category in the Prism Central.",
									MarkdownDescription: "key is the Key of the category in the Prism Central.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "value is the category value linked to the key in the Prism Central.",
									MarkdownDescription: "value is the category value linked to the key in the Prism Central.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster": schema.SingleNestedAttribute{
						Description:         "cluster is to identify the cluster (the Prism Element under management of the Prism Central), in which the Machine's VM will be created. The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
						MarkdownDescription: "cluster is to identify the cluster (the Prism Element under management of the Prism Central), in which the Machine's VM will be created. The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the prism_central API.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the resource name in the PC",
								MarkdownDescription: "name is the resource name in the PC",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the identifier type to use for this resource.",
								MarkdownDescription: "Type is the identifier type to use for this resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("uuid", "name"),
								},
							},

							"uuid": schema.StringAttribute{
								Description:         "uuid is the UUID of the resource in the PC.",
								MarkdownDescription: "uuid is the UUID of the resource in the PC.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"gpus": schema.ListNestedAttribute{
						Description:         "List of GPU devices that should be added to the VMs.",
						MarkdownDescription: "List of GPU devices that should be added to the VMs.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"device_id": schema.Int64Attribute{
									Description:         "deviceID is the device ID of the GPU device.",
									MarkdownDescription: "deviceID is the device ID of the GPU device.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "vendorID is the vendor ID of the GPU device.",
									MarkdownDescription: "vendorID is the vendor ID of the GPU device.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type": schema.StringAttribute{
									Description:         "type is the type of the GPU device.",
									MarkdownDescription: "type is the type of the GPU device.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("deviceID", "name"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "image is to identify the OS image uploaded to the Prism Central (PC) The image identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API. It must include the Kubernetes version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27.",
						MarkdownDescription: "image is to identify the OS image uploaded to the Prism Central (PC) The image identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API. It must include the Kubernetes version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the resource name in the PC",
								MarkdownDescription: "name is the resource name in the PC",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the identifier type to use for this resource.",
								MarkdownDescription: "Type is the identifier type to use for this resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("uuid", "name"),
								},
							},

							"uuid": schema.StringAttribute{
								Description:         "uuid is the UUID of the resource in the PC.",
								MarkdownDescription: "uuid is the UUID of the resource in the PC.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"memory_size": schema.StringAttribute{
						Description:         "memorySize is the memory size (in Quantity format) of the VM The minimum memorySize is 2Gi bytes",
						MarkdownDescription: "memorySize is the memory size (in Quantity format) of the VM The minimum memorySize is 2Gi bytes",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"os_family": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"project": schema.SingleNestedAttribute{
						Description:         "Project is an optional property that specifies the Prism Central project so that machine resources can be linked to it. The project identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API.",
						MarkdownDescription: "Project is an optional property that specifies the Prism Central project so that machine resources can be linked to it. The project identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the resource name in the PC",
								MarkdownDescription: "name is the resource name in the PC",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the identifier type to use for this resource.",
								MarkdownDescription: "Type is the identifier type to use for this resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("uuid", "name"),
								},
							},

							"uuid": schema.StringAttribute{
								Description:         "uuid is the UUID of the resource in the PC.",
								MarkdownDescription: "uuid is the UUID of the resource in the PC.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"subnet": schema.SingleNestedAttribute{
						Description:         "subnet is to identify the cluster's network subnet to use for the Machine's VM The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API.",
						MarkdownDescription: "subnet is to identify the cluster's network subnet to use for the Machine's VM The cluster identifier (uuid or name) can be obtained from the Prism Central console or using the Prism Central API.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is the resource name in the PC",
								MarkdownDescription: "name is the resource name in the PC",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the identifier type to use for this resource.",
								MarkdownDescription: "Type is the identifier type to use for this resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("uuid", "name"),
								},
							},

							"uuid": schema.StringAttribute{
								Description:         "uuid is the UUID of the resource in the PC.",
								MarkdownDescription: "uuid is the UUID of the resource in the PC.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"system_disk_size": schema.StringAttribute{
						Description:         "systemDiskSize is size (in Quantity format) of the system disk of the VM The minimum systemDiskSize is 20Gi bytes",
						MarkdownDescription: "systemDiskSize is size (in Quantity format) of the system disk of the VM The minimum systemDiskSize is 20Gi bytes",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"users": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"ssh_authorized_keys": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            true,
									Optional:            false,
									Computed:            false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vcpu_sockets": schema.Int64Attribute{
						Description:         "vcpuSockets is the number of vCPU sockets of the VM",
						MarkdownDescription: "vcpuSockets is the number of vCPU sockets of the VM",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"vcpus_per_socket": schema.Int64Attribute{
						Description:         "vcpusPerSocket is the number of vCPUs per socket of the VM",
						MarkdownDescription: "vcpusPerSocket is the number of vCPUs per socket of the VM",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_nutanix_machine_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComNutanixMachineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("NutanixMachineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
