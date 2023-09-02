/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package metal3_io_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &Metal3IoBareMetalHostV1Alpha1Manifest{}
)

func NewMetal3IoBareMetalHostV1Alpha1Manifest() datasource.DataSource {
	return &Metal3IoBareMetalHostV1Alpha1Manifest{}
}

type Metal3IoBareMetalHostV1Alpha1Manifest struct{}

type Metal3IoBareMetalHostV1Alpha1ManifestData struct {
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
		Architecture          *string `tfsdk:"architecture" json:"architecture,omitempty"`
		AutomatedCleaningMode *string `tfsdk:"automated_cleaning_mode" json:"automatedCleaningMode,omitempty"`
		Bmc                   *struct {
			Address                        *string `tfsdk:"address" json:"address,omitempty"`
			CredentialsName                *string `tfsdk:"credentials_name" json:"credentialsName,omitempty"`
			DisableCertificateVerification *bool   `tfsdk:"disable_certificate_verification" json:"disableCertificateVerification,omitempty"`
		} `tfsdk:"bmc" json:"bmc,omitempty"`
		BootMACAddress *string `tfsdk:"boot_mac_address" json:"bootMACAddress,omitempty"`
		BootMode       *string `tfsdk:"boot_mode" json:"bootMode,omitempty"`
		ConsumerRef    *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"consumer_ref" json:"consumerRef,omitempty"`
		CustomDeploy *struct {
			Method *string `tfsdk:"method" json:"method,omitempty"`
		} `tfsdk:"custom_deploy" json:"customDeploy,omitempty"`
		Description           *string `tfsdk:"description" json:"description,omitempty"`
		ExternallyProvisioned *bool   `tfsdk:"externally_provisioned" json:"externallyProvisioned,omitempty"`
		Firmware              *struct {
			SimultaneousMultithreadingEnabled *bool `tfsdk:"simultaneous_multithreading_enabled" json:"simultaneousMultithreadingEnabled,omitempty"`
			SriovEnabled                      *bool `tfsdk:"sriov_enabled" json:"sriovEnabled,omitempty"`
			VirtualizationEnabled             *bool `tfsdk:"virtualization_enabled" json:"virtualizationEnabled,omitempty"`
		} `tfsdk:"firmware" json:"firmware,omitempty"`
		HardwareProfile *string `tfsdk:"hardware_profile" json:"hardwareProfile,omitempty"`
		Image           *struct {
			Checksum     *string `tfsdk:"checksum" json:"checksum,omitempty"`
			ChecksumType *string `tfsdk:"checksum_type" json:"checksumType,omitempty"`
			Format       *string `tfsdk:"format" json:"format,omitempty"`
			Url          *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"image" json:"image,omitempty"`
		MetaData *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"meta_data" json:"metaData,omitempty"`
		NetworkData *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"network_data" json:"networkData,omitempty"`
		Online                         *bool   `tfsdk:"online" json:"online,omitempty"`
		PreprovisioningNetworkDataName *string `tfsdk:"preprovisioning_network_data_name" json:"preprovisioningNetworkDataName,omitempty"`
		Raid                           *struct {
			HardwareRAIDVolumes *[]struct {
				Controller            *string   `tfsdk:"controller" json:"controller,omitempty"`
				Level                 *string   `tfsdk:"level" json:"level,omitempty"`
				Name                  *string   `tfsdk:"name" json:"name,omitempty"`
				NumberOfPhysicalDisks *int64    `tfsdk:"number_of_physical_disks" json:"numberOfPhysicalDisks,omitempty"`
				PhysicalDisks         *[]string `tfsdk:"physical_disks" json:"physicalDisks,omitempty"`
				Rotational            *bool     `tfsdk:"rotational" json:"rotational,omitempty"`
				SizeGibibytes         *int64    `tfsdk:"size_gibibytes" json:"sizeGibibytes,omitempty"`
			} `tfsdk:"hardware_raid_volumes" json:"hardwareRAIDVolumes,omitempty"`
			SoftwareRAIDVolumes *[]struct {
				Level         *string `tfsdk:"level" json:"level,omitempty"`
				PhysicalDisks *[]struct {
					DeviceName         *string `tfsdk:"device_name" json:"deviceName,omitempty"`
					Hctl               *string `tfsdk:"hctl" json:"hctl,omitempty"`
					MinSizeGigabytes   *int64  `tfsdk:"min_size_gigabytes" json:"minSizeGigabytes,omitempty"`
					Model              *string `tfsdk:"model" json:"model,omitempty"`
					Rotational         *bool   `tfsdk:"rotational" json:"rotational,omitempty"`
					SerialNumber       *string `tfsdk:"serial_number" json:"serialNumber,omitempty"`
					Vendor             *string `tfsdk:"vendor" json:"vendor,omitempty"`
					Wwn                *string `tfsdk:"wwn" json:"wwn,omitempty"`
					WwnVendorExtension *string `tfsdk:"wwn_vendor_extension" json:"wwnVendorExtension,omitempty"`
					WwnWithExtension   *string `tfsdk:"wwn_with_extension" json:"wwnWithExtension,omitempty"`
				} `tfsdk:"physical_disks" json:"physicalDisks,omitempty"`
				SizeGibibytes *int64 `tfsdk:"size_gibibytes" json:"sizeGibibytes,omitempty"`
			} `tfsdk:"software_raid_volumes" json:"softwareRAIDVolumes,omitempty"`
		} `tfsdk:"raid" json:"raid,omitempty"`
		RootDeviceHints *struct {
			DeviceName         *string `tfsdk:"device_name" json:"deviceName,omitempty"`
			Hctl               *string `tfsdk:"hctl" json:"hctl,omitempty"`
			MinSizeGigabytes   *int64  `tfsdk:"min_size_gigabytes" json:"minSizeGigabytes,omitempty"`
			Model              *string `tfsdk:"model" json:"model,omitempty"`
			Rotational         *bool   `tfsdk:"rotational" json:"rotational,omitempty"`
			SerialNumber       *string `tfsdk:"serial_number" json:"serialNumber,omitempty"`
			Vendor             *string `tfsdk:"vendor" json:"vendor,omitempty"`
			Wwn                *string `tfsdk:"wwn" json:"wwn,omitempty"`
			WwnVendorExtension *string `tfsdk:"wwn_vendor_extension" json:"wwnVendorExtension,omitempty"`
			WwnWithExtension   *string `tfsdk:"wwn_with_extension" json:"wwnWithExtension,omitempty"`
		} `tfsdk:"root_device_hints" json:"rootDeviceHints,omitempty"`
		Taints *[]struct {
			Effect    *string `tfsdk:"effect" json:"effect,omitempty"`
			Key       *string `tfsdk:"key" json:"key,omitempty"`
			TimeAdded *string `tfsdk:"time_added" json:"timeAdded,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"taints" json:"taints,omitempty"`
		UserData *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"user_data" json:"userData,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Metal3IoBareMetalHostV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_metal3_io_bare_metal_host_v1alpha1_manifest"
}

func (r *Metal3IoBareMetalHostV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "BareMetalHost is the Schema for the baremetalhosts API",
		MarkdownDescription: "BareMetalHost is the Schema for the baremetalhosts API",
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
				Description:         "BareMetalHostSpec defines the desired state of BareMetalHost",
				MarkdownDescription: "BareMetalHostSpec defines the desired state of BareMetalHost",
				Attributes: map[string]schema.Attribute{
					"architecture": schema.StringAttribute{
						Description:         "CPU architecture of the host, e.g. 'x86_64' or 'aarch64'. If unset, eventually populated by inspection.",
						MarkdownDescription: "CPU architecture of the host, e.g. 'x86_64' or 'aarch64'. If unset, eventually populated by inspection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"automated_cleaning_mode": schema.StringAttribute{
						Description:         "When set to disabled, automated cleaning will be avoided during provisioning and deprovisioning.",
						MarkdownDescription: "When set to disabled, automated cleaning will be avoided during provisioning and deprovisioning.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("metadata", "disabled"),
						},
					},

					"bmc": schema.SingleNestedAttribute{
						Description:         "How do we connect to the BMC?",
						MarkdownDescription: "How do we connect to the BMC?",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Address holds the URL for accessing the controller on the network.",
								MarkdownDescription: "Address holds the URL for accessing the controller on the network.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"credentials_name": schema.StringAttribute{
								Description:         "The name of the secret containing the BMC credentials (requires keys 'username' and 'password').",
								MarkdownDescription: "The name of the secret containing the BMC credentials (requires keys 'username' and 'password').",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"disable_certificate_verification": schema.BoolAttribute{
								Description:         "DisableCertificateVerification disables verification of server certificates when using HTTPS to connect to the BMC. This is required when the server certificate is self-signed, but is insecure because it allows a man-in-the-middle to intercept the connection.",
								MarkdownDescription: "DisableCertificateVerification disables verification of server certificates when using HTTPS to connect to the BMC. This is required when the server certificate is self-signed, but is insecure because it allows a man-in-the-middle to intercept the connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"boot_mac_address": schema.StringAttribute{
						Description:         "Which MAC address will PXE boot? This is optional for some types, but required for libvirt VMs driven by vbmc.",
						MarkdownDescription: "Which MAC address will PXE boot? This is optional for some types, but required for libvirt VMs driven by vbmc.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}`), ""),
						},
					},

					"boot_mode": schema.StringAttribute{
						Description:         "Select the method of initializing the hardware during boot. Defaults to UEFI.",
						MarkdownDescription: "Select the method of initializing the hardware during boot. Defaults to UEFI.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("UEFI", "UEFISecureBoot", "legacy"),
						},
					},

					"consumer_ref": schema.SingleNestedAttribute{
						Description:         "ConsumerRef can be used to store information about something that is using a host. When it is not empty, the host is considered 'in use'.",
						MarkdownDescription: "ConsumerRef can be used to store information about something that is using a host. When it is not empty, the host is considered 'in use'.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"custom_deploy": schema.SingleNestedAttribute{
						Description:         "A custom deploy procedure.",
						MarkdownDescription: "A custom deploy procedure.",
						Attributes: map[string]schema.Attribute{
							"method": schema.StringAttribute{
								Description:         "Custom deploy method name. This name is specific to the deploy ramdisk used. If you don't have a custom deploy ramdisk, you shouldn't use CustomDeploy.",
								MarkdownDescription: "Custom deploy method name. This name is specific to the deploy ramdisk used. If you don't have a custom deploy ramdisk, you shouldn't use CustomDeploy.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"description": schema.StringAttribute{
						Description:         "Description is a human-entered text used to help identify the host",
						MarkdownDescription: "Description is a human-entered text used to help identify the host",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"externally_provisioned": schema.BoolAttribute{
						Description:         "ExternallyProvisioned means something else is managing the image running on the host and the operator should only manage the power status and hardware inventory inspection. If the Image field is filled in, this field is ignored.",
						MarkdownDescription: "ExternallyProvisioned means something else is managing the image running on the host and the operator should only manage the power status and hardware inventory inspection. If the Image field is filled in, this field is ignored.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"firmware": schema.SingleNestedAttribute{
						Description:         "BIOS configuration for bare metal server",
						MarkdownDescription: "BIOS configuration for bare metal server",
						Attributes: map[string]schema.Attribute{
							"simultaneous_multithreading_enabled": schema.BoolAttribute{
								Description:         "Allows a single physical processor core to appear as several logical processors. This supports following options: true, false.",
								MarkdownDescription: "Allows a single physical processor core to appear as several logical processors. This supports following options: true, false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sriov_enabled": schema.BoolAttribute{
								Description:         "SR-IOV support enables a hypervisor to create virtual instances of a PCI-express device, potentially increasing performance. This supports following options: true, false.",
								MarkdownDescription: "SR-IOV support enables a hypervisor to create virtual instances of a PCI-express device, potentially increasing performance. This supports following options: true, false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"virtualization_enabled": schema.BoolAttribute{
								Description:         "Supports the virtualization of platform hardware. This supports following options: true, false.",
								MarkdownDescription: "Supports the virtualization of platform hardware. This supports following options: true, false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"hardware_profile": schema.StringAttribute{
						Description:         "What is the name of the hardware profile for this host? It should only be necessary to set this when inspection cannot automatically determine the profile.",
						MarkdownDescription: "What is the name of the hardware profile for this host? It should only be necessary to set this when inspection cannot automatically determine the profile.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.SingleNestedAttribute{
						Description:         "Image holds the details of the image to be provisioned.",
						MarkdownDescription: "Image holds the details of the image to be provisioned.",
						Attributes: map[string]schema.Attribute{
							"checksum": schema.StringAttribute{
								Description:         "Checksum is the checksum for the image.",
								MarkdownDescription: "Checksum is the checksum for the image.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"checksum_type": schema.StringAttribute{
								Description:         "ChecksumType is the checksum algorithm for the image. e.g md5, sha256, sha512",
								MarkdownDescription: "ChecksumType is the checksum algorithm for the image. e.g md5, sha256, sha512",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("md5", "sha256", "sha512"),
								},
							},

							"format": schema.StringAttribute{
								Description:         "DiskFormat contains the format of the image (raw, qcow2, ...). Needs to be set to raw for raw images streaming. Note live-iso means an iso referenced by the url will be live-booted and not deployed to disk, and in this case the checksum options are not required and if specified will be ignored.",
								MarkdownDescription: "DiskFormat contains the format of the image (raw, qcow2, ...). Needs to be set to raw for raw images streaming. Note live-iso means an iso referenced by the url will be live-booted and not deployed to disk, and in this case the checksum options are not required and if specified will be ignored.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("raw", "qcow2", "vdi", "vmdk", "live-iso"),
								},
							},

							"url": schema.StringAttribute{
								Description:         "URL is a location of an image to deploy.",
								MarkdownDescription: "URL is a location of an image to deploy.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"meta_data": schema.SingleNestedAttribute{
						Description:         "MetaData holds the reference to the Secret containing host metadata (e.g. meta_data.json) which is passed to the Config Drive.",
						MarkdownDescription: "MetaData holds the reference to the Secret containing host metadata (e.g. meta_data.json) which is passed to the Config Drive.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_data": schema.SingleNestedAttribute{
						Description:         "NetworkData holds the reference to the Secret containing network configuration (e.g content of network_data.json) which is passed to the Config Drive.",
						MarkdownDescription: "NetworkData holds the reference to the Secret containing network configuration (e.g content of network_data.json) which is passed to the Config Drive.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"online": schema.BoolAttribute{
						Description:         "Should the server be online?",
						MarkdownDescription: "Should the server be online?",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"preprovisioning_network_data_name": schema.StringAttribute{
						Description:         "PreprovisioningNetworkDataName is the name of the Secret in the local namespace containing network configuration (e.g content of network_data.json) which is passed to the preprovisioning image, and to the Config Drive if not overridden by specifying NetworkData.",
						MarkdownDescription: "PreprovisioningNetworkDataName is the name of the Secret in the local namespace containing network configuration (e.g content of network_data.json) which is passed to the preprovisioning image, and to the Config Drive if not overridden by specifying NetworkData.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"raid": schema.SingleNestedAttribute{
						Description:         "RAID configuration for bare metal server",
						MarkdownDescription: "RAID configuration for bare metal server",
						Attributes: map[string]schema.Attribute{
							"hardware_raid_volumes": schema.ListNestedAttribute{
								Description:         "The list of logical disks for hardware RAID, if rootDeviceHints isn't used, first volume is root volume. You can set the value of this field to '[]' to clear all the hardware RAID configurations.",
								MarkdownDescription: "The list of logical disks for hardware RAID, if rootDeviceHints isn't used, first volume is root volume. You can set the value of this field to '[]' to clear all the hardware RAID configurations.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"controller": schema.StringAttribute{
											Description:         "The name of the RAID controller to use",
											MarkdownDescription: "The name of the RAID controller to use",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"level": schema.StringAttribute{
											Description:         "RAID level for the logical disk. The following levels are supported: 0;1;2;5;6;1+0;5+0;6+0.",
											MarkdownDescription: "RAID level for the logical disk. The following levels are supported: 0;1;2;5;6;1+0;5+0;6+0.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("0", "1", "2", "5", "6", "1+0", "5+0", "6+0"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name of the volume. Should be unique within the Node. If not specified, volume name will be auto-generated.",
											MarkdownDescription: "Name of the volume. Should be unique within the Node. If not specified, volume name will be auto-generated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(64),
											},
										},

										"number_of_physical_disks": schema.Int64Attribute{
											Description:         "Integer, number of physical disks to use for the logical disk. Defaults to minimum number of disks required for the particular RAID level.",
											MarkdownDescription: "Integer, number of physical disks to use for the logical disk. Defaults to minimum number of disks required for the particular RAID level.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(1),
											},
										},

										"physical_disks": schema.ListAttribute{
											Description:         "Optional list of physical disk names to be used for the Hardware RAID volumes. The disk names are interpreted by the Hardware RAID controller, and the format is hardware specific.",
											MarkdownDescription: "Optional list of physical disk names to be used for the Hardware RAID volumes. The disk names are interpreted by the Hardware RAID controller, and the format is hardware specific.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"rotational": schema.BoolAttribute{
											Description:         "Select disks with only rotational or solid-state storage",
											MarkdownDescription: "Select disks with only rotational or solid-state storage",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size_gibibytes": schema.Int64Attribute{
											Description:         "Size (Integer) of the logical disk to be created in GiB. If unspecified or set be 0, the maximum capacity of disk will be used for logical disk.",
											MarkdownDescription: "Size (Integer) of the logical disk to be created in GiB. If unspecified or set be 0, the maximum capacity of disk will be used for logical disk.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"software_raid_volumes": schema.ListNestedAttribute{
								Description:         "The list of logical disks for software RAID, if rootDeviceHints isn't used, first volume is root volume. If HardwareRAIDVolumes is set this item will be invalid. The number of created Software RAID devices must be 1 or 2. If there is only one Software RAID device, it has to be a RAID-1. If there are two, the first one has to be a RAID-1, while the RAID level for the second one can be 0, 1, or 1+0. As the first RAID device will be the deployment device, enforcing a RAID-1 reduces the risk of ending up with a non-booting node in case of a disk failure. Software RAID will always be deleted.",
								MarkdownDescription: "The list of logical disks for software RAID, if rootDeviceHints isn't used, first volume is root volume. If HardwareRAIDVolumes is set this item will be invalid. The number of created Software RAID devices must be 1 or 2. If there is only one Software RAID device, it has to be a RAID-1. If there are two, the first one has to be a RAID-1, while the RAID level for the second one can be 0, 1, or 1+0. As the first RAID device will be the deployment device, enforcing a RAID-1 reduces the risk of ending up with a non-booting node in case of a disk failure. Software RAID will always be deleted.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"level": schema.StringAttribute{
											Description:         "RAID level for the logical disk. The following levels are supported: 0;1;1+0.",
											MarkdownDescription: "RAID level for the logical disk. The following levels are supported: 0;1;1+0.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("0", "1", "1+0"),
											},
										},

										"physical_disks": schema.ListNestedAttribute{
											Description:         "A list of device hints, the number of items should be greater than or equal to 2.",
											MarkdownDescription: "A list of device hints, the number of items should be greater than or equal to 2.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"device_name": schema.StringAttribute{
														Description:         "A Linux device name like '/dev/vda', or a by-path link to it like '/dev/disk/by-path/pci-0000:01:00.0-scsi-0:2:0:0'. The hint must match the actual value exactly.",
														MarkdownDescription: "A Linux device name like '/dev/vda', or a by-path link to it like '/dev/disk/by-path/pci-0000:01:00.0-scsi-0:2:0:0'. The hint must match the actual value exactly.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"hctl": schema.StringAttribute{
														Description:         "A SCSI bus address like 0:0:0:0. The hint must match the actual value exactly.",
														MarkdownDescription: "A SCSI bus address like 0:0:0:0. The hint must match the actual value exactly.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_size_gigabytes": schema.Int64Attribute{
														Description:         "The minimum size of the device in Gigabytes.",
														MarkdownDescription: "The minimum size of the device in Gigabytes.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"model": schema.StringAttribute{
														Description:         "A vendor-specific device identifier. The hint can be a substring of the actual value.",
														MarkdownDescription: "A vendor-specific device identifier. The hint can be a substring of the actual value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"rotational": schema.BoolAttribute{
														Description:         "True if the device should use spinning media, false otherwise.",
														MarkdownDescription: "True if the device should use spinning media, false otherwise.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"serial_number": schema.StringAttribute{
														Description:         "Device serial number. The hint must match the actual value exactly.",
														MarkdownDescription: "Device serial number. The hint must match the actual value exactly.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"vendor": schema.StringAttribute{
														Description:         "The name of the vendor or manufacturer of the device. The hint can be a substring of the actual value.",
														MarkdownDescription: "The name of the vendor or manufacturer of the device. The hint can be a substring of the actual value.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"wwn": schema.StringAttribute{
														Description:         "Unique storage identifier. The hint must match the actual value exactly.",
														MarkdownDescription: "Unique storage identifier. The hint must match the actual value exactly.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"wwn_vendor_extension": schema.StringAttribute{
														Description:         "Unique vendor storage identifier. The hint must match the actual value exactly.",
														MarkdownDescription: "Unique vendor storage identifier. The hint must match the actual value exactly.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"wwn_with_extension": schema.StringAttribute{
														Description:         "Unique storage identifier with the vendor extension appended. The hint must match the actual value exactly.",
														MarkdownDescription: "Unique storage identifier with the vendor extension appended. The hint must match the actual value exactly.",
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

										"size_gibibytes": schema.Int64Attribute{
											Description:         "Size (Integer) of the logical disk to be created in GiB. If unspecified or set be 0, the maximum capacity of disk will be used for logical disk.",
											MarkdownDescription: "Size (Integer) of the logical disk to be created in GiB. If unspecified or set be 0, the maximum capacity of disk will be used for logical disk.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
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

					"root_device_hints": schema.SingleNestedAttribute{
						Description:         "Provide guidance about how to choose the device for the image being provisioned.",
						MarkdownDescription: "Provide guidance about how to choose the device for the image being provisioned.",
						Attributes: map[string]schema.Attribute{
							"device_name": schema.StringAttribute{
								Description:         "A Linux device name like '/dev/vda', or a by-path link to it like '/dev/disk/by-path/pci-0000:01:00.0-scsi-0:2:0:0'. The hint must match the actual value exactly.",
								MarkdownDescription: "A Linux device name like '/dev/vda', or a by-path link to it like '/dev/disk/by-path/pci-0000:01:00.0-scsi-0:2:0:0'. The hint must match the actual value exactly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hctl": schema.StringAttribute{
								Description:         "A SCSI bus address like 0:0:0:0. The hint must match the actual value exactly.",
								MarkdownDescription: "A SCSI bus address like 0:0:0:0. The hint must match the actual value exactly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_size_gigabytes": schema.Int64Attribute{
								Description:         "The minimum size of the device in Gigabytes.",
								MarkdownDescription: "The minimum size of the device in Gigabytes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"model": schema.StringAttribute{
								Description:         "A vendor-specific device identifier. The hint can be a substring of the actual value.",
								MarkdownDescription: "A vendor-specific device identifier. The hint can be a substring of the actual value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rotational": schema.BoolAttribute{
								Description:         "True if the device should use spinning media, false otherwise.",
								MarkdownDescription: "True if the device should use spinning media, false otherwise.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"serial_number": schema.StringAttribute{
								Description:         "Device serial number. The hint must match the actual value exactly.",
								MarkdownDescription: "Device serial number. The hint must match the actual value exactly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"vendor": schema.StringAttribute{
								Description:         "The name of the vendor or manufacturer of the device. The hint can be a substring of the actual value.",
								MarkdownDescription: "The name of the vendor or manufacturer of the device. The hint can be a substring of the actual value.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wwn": schema.StringAttribute{
								Description:         "Unique storage identifier. The hint must match the actual value exactly.",
								MarkdownDescription: "Unique storage identifier. The hint must match the actual value exactly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wwn_vendor_extension": schema.StringAttribute{
								Description:         "Unique vendor storage identifier. The hint must match the actual value exactly.",
								MarkdownDescription: "Unique vendor storage identifier. The hint must match the actual value exactly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wwn_with_extension": schema.StringAttribute{
								Description:         "Unique storage identifier with the vendor extension appended. The hint must match the actual value exactly.",
								MarkdownDescription: "Unique storage identifier with the vendor extension appended. The hint must match the actual value exactly.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"taints": schema.ListNestedAttribute{
						Description:         "Taints is the full, authoritative list of taints to apply to the corresponding Machine. This list will overwrite any modifications made to the Machine on an ongoing basis.",
						MarkdownDescription: "Taints is the full, authoritative list of taints to apply to the corresponding Machine. This list will overwrite any modifications made to the Machine on an ongoing basis.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Required. The taint key to be applied to a node.",
									MarkdownDescription: "Required. The taint key to be applied to a node.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"time_added": schema.StringAttribute{
									Description:         "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
									MarkdownDescription: "TimeAdded represents the time at which the taint was added. It is only written for NoExecute taints.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										validators.DateTime64Validator(),
									},
								},

								"value": schema.StringAttribute{
									Description:         "The taint value corresponding to the taint key.",
									MarkdownDescription: "The taint value corresponding to the taint key.",
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

					"user_data": schema.SingleNestedAttribute{
						Description:         "UserData holds the reference to the Secret containing the user data to be passed to the host before it boots.",
						MarkdownDescription: "UserData holds the reference to the Secret containing the user data to be passed to the host before it boots.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "name is unique within a namespace to reference a secret resource.",
								MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "namespace defines the space within which the secret name must be unique.",
								MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
								Required:            false,
								Optional:            true,
								Computed:            false,
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

func (r *Metal3IoBareMetalHostV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_metal3_io_bare_metal_host_v1alpha1_manifest")

	var model Metal3IoBareMetalHostV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("metal3.io/v1alpha1")
	model.Kind = pointer.String("BareMetalHost")

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
