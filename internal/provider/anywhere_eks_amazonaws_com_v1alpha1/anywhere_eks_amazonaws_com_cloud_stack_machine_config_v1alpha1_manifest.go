/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package anywhere_eks_amazonaws_com_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &AnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1Manifest{}
)

func NewAnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1Manifest() datasource.DataSource {
	return &AnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1Manifest{}
}

type AnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1Manifest struct{}

type AnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Affinity         *string   `tfsdk:"affinity" json:"affinity,omitempty"`
		AffinityGroupIds *[]string `tfsdk:"affinity_group_ids" json:"affinityGroupIds,omitempty"`
		ComputeOffering  *struct {
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"compute_offering" json:"computeOffering,omitempty"`
		DiskOffering *struct {
			CustomSizeInGB *int64  `tfsdk:"custom_size_in_gb" json:"customSizeInGB,omitempty"`
			Device         *string `tfsdk:"device" json:"device,omitempty"`
			Filesystem     *string `tfsdk:"filesystem" json:"filesystem,omitempty"`
			Id             *string `tfsdk:"id" json:"id,omitempty"`
			Label          *string `tfsdk:"label" json:"label,omitempty"`
			MountPath      *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			Name           *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"disk_offering" json:"diskOffering,omitempty"`
		Symlinks *map[string]string `tfsdk:"symlinks" json:"symlinks,omitempty"`
		Template *struct {
			Id   *string `tfsdk:"id" json:"id,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
		UserCustomDetails *map[string]string `tfsdk:"user_custom_details" json:"userCustomDetails,omitempty"`
		Users             *[]struct {
			Name              *string   `tfsdk:"name" json:"name,omitempty"`
			SshAuthorizedKeys *[]string `tfsdk:"ssh_authorized_keys" json:"sshAuthorizedKeys,omitempty"`
		} `tfsdk:"users" json:"users,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest"
}

func (r *AnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CloudStackMachineConfig is the Schema for the cloudstackmachineconfigs API.",
		MarkdownDescription: "CloudStackMachineConfig is the Schema for the cloudstackmachineconfigs API.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
				Description:         "CloudStackMachineConfigSpec defines the desired state of CloudStackMachineConfig.",
				MarkdownDescription: "CloudStackMachineConfigSpec defines the desired state of CloudStackMachineConfig.",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.StringAttribute{
						Description:         "Defaults to 'no'. Can be 'pro' or 'anti'. If set to 'pro' or 'anti', will create an affinity group per machine set of the corresponding type",
						MarkdownDescription: "Defaults to 'no'. Can be 'pro' or 'anti'. If set to 'pro' or 'anti', will create an affinity group per machine set of the corresponding type",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"affinity_group_ids": schema.ListAttribute{
						Description:         "AffinityGroupIds allows users to pass in a list of UUIDs for previously-created Affinity Groups. Any VM’s created with this spec will be added to the affinity group, which will dictate which physical host(s) they can be placed on. Affinity groups can be type “affinity” or “anti-affinity” in CloudStack. If they are type “anti-affinity”, all VM’s in the group must be on separate physical hosts for high availability. If they are type “affinity”, all VM’s in the group must be on the same physical host for improved performance",
						MarkdownDescription: "AffinityGroupIds allows users to pass in a list of UUIDs for previously-created Affinity Groups. Any VM’s created with this spec will be added to the affinity group, which will dictate which physical host(s) they can be placed on. Affinity groups can be type “affinity” or “anti-affinity” in CloudStack. If they are type “anti-affinity”, all VM’s in the group must be on separate physical hosts for high availability. If they are type “affinity”, all VM’s in the group must be on the same physical host for improved performance",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"compute_offering": schema.SingleNestedAttribute{
						Description:         "ComputeOffering refers to a compute offering which has been previously registered in CloudStack. It represents a VM’s instance size including number of CPU’s, memory, and CPU speed. It can either be specified as a UUID or name",
						MarkdownDescription: "ComputeOffering refers to a compute offering which has been previously registered in CloudStack. It represents a VM’s instance size including number of CPU’s, memory, and CPU speed. It can either be specified as a UUID or name",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
								MarkdownDescription: "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
								MarkdownDescription: "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"disk_offering": schema.SingleNestedAttribute{
						Description:         "DiskOffering refers to a disk offering which has been previously registered in CloudStack. It represents a disk offering with pre-defined size or custom specified disk size. It can either be specified as a UUID or name",
						MarkdownDescription: "DiskOffering refers to a disk offering which has been previously registered in CloudStack. It represents a disk offering with pre-defined size or custom specified disk size. It can either be specified as a UUID or name",
						Attributes: map[string]schema.Attribute{
							"custom_size_in_gb": schema.Int64Attribute{
								Description:         "disk size in GB, > 0 for customized disk offering; = 0 for non-customized disk offering",
								MarkdownDescription: "disk size in GB, > 0 for customized disk offering; = 0 for non-customized disk offering",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"device": schema.StringAttribute{
								Description:         "device name of the disk offering in VM, shows up in lsblk command",
								MarkdownDescription: "device name of the disk offering in VM, shows up in lsblk command",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"filesystem": schema.StringAttribute{
								Description:         "filesystem used to mkfs in disk offering partition",
								MarkdownDescription: "filesystem used to mkfs in disk offering partition",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"id": schema.StringAttribute{
								Description:         "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
								MarkdownDescription: "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label": schema.StringAttribute{
								Description:         "disk label used to label disk partition",
								MarkdownDescription: "disk label used to label disk partition",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"mount_path": schema.StringAttribute{
								Description:         "path the filesystem will use to mount in VM",
								MarkdownDescription: "path the filesystem will use to mount in VM",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
								MarkdownDescription: "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"symlinks": schema.MapAttribute{
						Description:         "Symlinks create soft symbolic links folders. One use case is to use data disk to store logs",
						MarkdownDescription: "Symlinks create soft symbolic links folders. One use case is to use data disk to store logs",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"template": schema.SingleNestedAttribute{
						Description:         "Template refers to a VM image template which has been previously registered in CloudStack. It can either be specified as a UUID or name. When using a template name it must include the Kubernetes version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27.",
						MarkdownDescription: "Template refers to a VM image template which has been previously registered in CloudStack. It can either be specified as a UUID or name. When using a template name it must include the Kubernetes version(s). For example, a template used for Kubernetes 1.27 could be ubuntu-2204-1.27.",
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Description:         "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
								MarkdownDescription: "Id of a resource in the CloudStack environment. Mutually exclusive with Name",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
								MarkdownDescription: "Name of a resource in the CloudStack environment. Mutually exclusive with Id",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"user_custom_details": schema.MapAttribute{
						Description:         "UserCustomDetails allows users to pass in non-standard key value inputs, outside those defined [here](https://github.com/shapeblue/cloudstack/blob/main/api/src/main/java/com/cloud/vm/VmDetailConstants.java)",
						MarkdownDescription: "UserCustomDetails allows users to pass in non-standard key value inputs, outside those defined [here](https://github.com/shapeblue/cloudstack/blob/main/api/src/main/java/com/cloud/vm/VmDetailConstants.java)",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"users": schema.ListNestedAttribute{
						Description:         "Users consists of an array of objects containing the username, as well as a list of their public keys. These users will be authorized to ssh into the machines",
						MarkdownDescription: "Users consists of an array of objects containing the username, as well as a list of their public keys. These users will be authorized to ssh into the machines",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_anywhere_eks_amazonaws_com_cloud_stack_machine_config_v1alpha1_manifest")

	var model AnywhereEksAmazonawsComCloudStackMachineConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("anywhere.eks.amazonaws.com/v1alpha1")
	model.Kind = pointer.String("CloudStackMachineConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
