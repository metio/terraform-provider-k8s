/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type Ec2ServicesK8SAwsInstanceV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*Ec2ServicesK8SAwsInstanceV1Alpha1Resource)(nil)
)

type Ec2ServicesK8SAwsInstanceV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type Ec2ServicesK8SAwsInstanceV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		BlockDeviceMappings *[]struct {
			DeviceName *string `tfsdk:"device_name" yaml:"deviceName,omitempty"`

			Ebs *struct {
				DeleteOnTermination *bool `tfsdk:"delete_on_termination" yaml:"deleteOnTermination,omitempty"`

				Encrypted *bool `tfsdk:"encrypted" yaml:"encrypted,omitempty"`

				Iops *int64 `tfsdk:"iops" yaml:"iops,omitempty"`

				KmsKeyID *string `tfsdk:"kms_key_id" yaml:"kmsKeyID,omitempty"`

				OutpostARN *string `tfsdk:"outpost_arn" yaml:"outpostARN,omitempty"`

				SnapshotID *string `tfsdk:"snapshot_id" yaml:"snapshotID,omitempty"`

				Throughput *int64 `tfsdk:"throughput" yaml:"throughput,omitempty"`

				VolumeSize *int64 `tfsdk:"volume_size" yaml:"volumeSize,omitempty"`

				VolumeType *string `tfsdk:"volume_type" yaml:"volumeType,omitempty"`
			} `tfsdk:"ebs" yaml:"ebs,omitempty"`

			NoDevice *string `tfsdk:"no_device" yaml:"noDevice,omitempty"`

			VirtualName *string `tfsdk:"virtual_name" yaml:"virtualName,omitempty"`
		} `tfsdk:"block_device_mappings" yaml:"blockDeviceMappings,omitempty"`

		CapacityReservationSpecification *struct {
			CapacityReservationPreference *string `tfsdk:"capacity_reservation_preference" yaml:"capacityReservationPreference,omitempty"`

			CapacityReservationTarget *struct {
				CapacityReservationID *string `tfsdk:"capacity_reservation_id" yaml:"capacityReservationID,omitempty"`

				CapacityReservationResourceGroupARN *string `tfsdk:"capacity_reservation_resource_group_arn" yaml:"capacityReservationResourceGroupARN,omitempty"`
			} `tfsdk:"capacity_reservation_target" yaml:"capacityReservationTarget,omitempty"`
		} `tfsdk:"capacity_reservation_specification" yaml:"capacityReservationSpecification,omitempty"`

		CpuOptions *struct {
			CoreCount *int64 `tfsdk:"core_count" yaml:"coreCount,omitempty"`

			ThreadsPerCore *int64 `tfsdk:"threads_per_core" yaml:"threadsPerCore,omitempty"`
		} `tfsdk:"cpu_options" yaml:"cpuOptions,omitempty"`

		CreditSpecification *struct {
			CpuCredits *string `tfsdk:"cpu_credits" yaml:"cpuCredits,omitempty"`
		} `tfsdk:"credit_specification" yaml:"creditSpecification,omitempty"`

		DisableAPIStop *bool `tfsdk:"disable_api_stop" yaml:"disableAPIStop,omitempty"`

		DisableAPITermination *bool `tfsdk:"disable_api_termination" yaml:"disableAPITermination,omitempty"`

		EbsOptimized *bool `tfsdk:"ebs_optimized" yaml:"ebsOptimized,omitempty"`

		ElasticGPUSpecification *[]struct {
			Type_ *string `tfsdk:"type_" yaml:"type_,omitempty"`
		} `tfsdk:"elastic_gpu_specification" yaml:"elasticGPUSpecification,omitempty"`

		ElasticInferenceAccelerators *[]struct {
			Count *int64 `tfsdk:"count" yaml:"count,omitempty"`

			Type_ *string `tfsdk:"type_" yaml:"type_,omitempty"`
		} `tfsdk:"elastic_inference_accelerators" yaml:"elasticInferenceAccelerators,omitempty"`

		EnclaveOptions *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"enclave_options" yaml:"enclaveOptions,omitempty"`

		HibernationOptions *struct {
			Configured *bool `tfsdk:"configured" yaml:"configured,omitempty"`
		} `tfsdk:"hibernation_options" yaml:"hibernationOptions,omitempty"`

		IamInstanceProfile *struct {
			Arn *string `tfsdk:"arn" yaml:"arn,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"iam_instance_profile" yaml:"iamInstanceProfile,omitempty"`

		ImageID *string `tfsdk:"image_id" yaml:"imageID,omitempty"`

		InstanceInitiatedShutdownBehavior *string `tfsdk:"instance_initiated_shutdown_behavior" yaml:"instanceInitiatedShutdownBehavior,omitempty"`

		InstanceMarketOptions *struct {
			MarketType *string `tfsdk:"market_type" yaml:"marketType,omitempty"`

			SpotOptions *struct {
				BlockDurationMinutes *int64 `tfsdk:"block_duration_minutes" yaml:"blockDurationMinutes,omitempty"`

				InstanceInterruptionBehavior *string `tfsdk:"instance_interruption_behavior" yaml:"instanceInterruptionBehavior,omitempty"`

				MaxPrice *string `tfsdk:"max_price" yaml:"maxPrice,omitempty"`

				SpotInstanceType *string `tfsdk:"spot_instance_type" yaml:"spotInstanceType,omitempty"`

				ValidUntil *string `tfsdk:"valid_until" yaml:"validUntil,omitempty"`
			} `tfsdk:"spot_options" yaml:"spotOptions,omitempty"`
		} `tfsdk:"instance_market_options" yaml:"instanceMarketOptions,omitempty"`

		InstanceType *string `tfsdk:"instance_type" yaml:"instanceType,omitempty"`

		Ipv6AddressCount *int64 `tfsdk:"ipv6_address_count" yaml:"ipv6AddressCount,omitempty"`

		Ipv6Addresses *[]struct {
			Ipv6Address *string `tfsdk:"ipv6_address" yaml:"ipv6Address,omitempty"`
		} `tfsdk:"ipv6_addresses" yaml:"ipv6Addresses,omitempty"`

		KernelID *string `tfsdk:"kernel_id" yaml:"kernelID,omitempty"`

		KeyName *string `tfsdk:"key_name" yaml:"keyName,omitempty"`

		LaunchTemplate *struct {
			LaunchTemplateID *string `tfsdk:"launch_template_id" yaml:"launchTemplateID,omitempty"`

			LaunchTemplateName *string `tfsdk:"launch_template_name" yaml:"launchTemplateName,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"launch_template" yaml:"launchTemplate,omitempty"`

		LicenseSpecifications *[]struct {
			LicenseConfigurationARN *string `tfsdk:"license_configuration_arn" yaml:"licenseConfigurationARN,omitempty"`
		} `tfsdk:"license_specifications" yaml:"licenseSpecifications,omitempty"`

		MaintenanceOptions *struct {
			AutoRecovery *string `tfsdk:"auto_recovery" yaml:"autoRecovery,omitempty"`
		} `tfsdk:"maintenance_options" yaml:"maintenanceOptions,omitempty"`

		MaxCount *int64 `tfsdk:"max_count" yaml:"maxCount,omitempty"`

		MetadataOptions *struct {
			HttpEndpoint *string `tfsdk:"http_endpoint" yaml:"httpEndpoint,omitempty"`

			HttpProtocolIPv6 *string `tfsdk:"http_protocol_i_pv6" yaml:"httpProtocolIPv6,omitempty"`

			HttpPutResponseHopLimit *int64 `tfsdk:"http_put_response_hop_limit" yaml:"httpPutResponseHopLimit,omitempty"`

			HttpTokens *string `tfsdk:"http_tokens" yaml:"httpTokens,omitempty"`

			InstanceMetadataTags *string `tfsdk:"instance_metadata_tags" yaml:"instanceMetadataTags,omitempty"`
		} `tfsdk:"metadata_options" yaml:"metadataOptions,omitempty"`

		MinCount *int64 `tfsdk:"min_count" yaml:"minCount,omitempty"`

		Monitoring *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"monitoring" yaml:"monitoring,omitempty"`

		NetworkInterfaces *[]struct {
			AssociateCarrierIPAddress *bool `tfsdk:"associate_carrier_ip_address" yaml:"associateCarrierIPAddress,omitempty"`

			AssociatePublicIPAddress *bool `tfsdk:"associate_public_ip_address" yaml:"associatePublicIPAddress,omitempty"`

			DeleteOnTermination *bool `tfsdk:"delete_on_termination" yaml:"deleteOnTermination,omitempty"`

			Description *string `tfsdk:"description" yaml:"description,omitempty"`

			DeviceIndex *int64 `tfsdk:"device_index" yaml:"deviceIndex,omitempty"`

			InterfaceType *string `tfsdk:"interface_type" yaml:"interfaceType,omitempty"`

			Ipv4PrefixCount *int64 `tfsdk:"ipv4_prefix_count" yaml:"ipv4PrefixCount,omitempty"`

			Ipv4Prefixes *[]struct {
				Ipv4Prefix *string `tfsdk:"ipv4_prefix" yaml:"ipv4Prefix,omitempty"`
			} `tfsdk:"ipv4_prefixes" yaml:"ipv4Prefixes,omitempty"`

			Ipv6AddressCount *int64 `tfsdk:"ipv6_address_count" yaml:"ipv6AddressCount,omitempty"`

			Ipv6Addresses *[]struct {
				Ipv6Address *string `tfsdk:"ipv6_address" yaml:"ipv6Address,omitempty"`
			} `tfsdk:"ipv6_addresses" yaml:"ipv6Addresses,omitempty"`

			Ipv6PrefixCount *int64 `tfsdk:"ipv6_prefix_count" yaml:"ipv6PrefixCount,omitempty"`

			Ipv6Prefixes *[]struct {
				Ipv6Prefix *string `tfsdk:"ipv6_prefix" yaml:"ipv6Prefix,omitempty"`
			} `tfsdk:"ipv6_prefixes" yaml:"ipv6Prefixes,omitempty"`

			NetworkCardIndex *int64 `tfsdk:"network_card_index" yaml:"networkCardIndex,omitempty"`

			NetworkInterfaceID *string `tfsdk:"network_interface_id" yaml:"networkInterfaceID,omitempty"`

			PrivateIPAddress *string `tfsdk:"private_ip_address" yaml:"privateIPAddress,omitempty"`

			PrivateIPAddresses *[]struct {
				Primary *bool `tfsdk:"primary" yaml:"primary,omitempty"`

				PrivateIPAddress *string `tfsdk:"private_ip_address" yaml:"privateIPAddress,omitempty"`
			} `tfsdk:"private_ip_addresses" yaml:"privateIPAddresses,omitempty"`

			SecondaryPrivateIPAddressCount *int64 `tfsdk:"secondary_private_ip_address_count" yaml:"secondaryPrivateIPAddressCount,omitempty"`

			SubnetID *string `tfsdk:"subnet_id" yaml:"subnetID,omitempty"`
		} `tfsdk:"network_interfaces" yaml:"networkInterfaces,omitempty"`

		Placement *struct {
			Affinity *string `tfsdk:"affinity" yaml:"affinity,omitempty"`

			AvailabilityZone *string `tfsdk:"availability_zone" yaml:"availabilityZone,omitempty"`

			GroupName *string `tfsdk:"group_name" yaml:"groupName,omitempty"`

			HostID *string `tfsdk:"host_id" yaml:"hostID,omitempty"`

			HostResourceGroupARN *string `tfsdk:"host_resource_group_arn" yaml:"hostResourceGroupARN,omitempty"`

			PartitionNumber *int64 `tfsdk:"partition_number" yaml:"partitionNumber,omitempty"`

			SpreadDomain *string `tfsdk:"spread_domain" yaml:"spreadDomain,omitempty"`

			Tenancy *string `tfsdk:"tenancy" yaml:"tenancy,omitempty"`
		} `tfsdk:"placement" yaml:"placement,omitempty"`

		PrivateDNSNameOptions *struct {
			EnableResourceNameDNSAAAARecord *bool `tfsdk:"enable_resource_name_dnsaaaa_record" yaml:"enableResourceNameDNSAAAARecord,omitempty"`

			EnableResourceNameDNSARecord *bool `tfsdk:"enable_resource_name_dnsa_record" yaml:"enableResourceNameDNSARecord,omitempty"`

			HostnameType *string `tfsdk:"hostname_type" yaml:"hostnameType,omitempty"`
		} `tfsdk:"private_dns_name_options" yaml:"privateDNSNameOptions,omitempty"`

		PrivateIPAddress *string `tfsdk:"private_ip_address" yaml:"privateIPAddress,omitempty"`

		RamDiskID *string `tfsdk:"ram_disk_id" yaml:"ramDiskID,omitempty"`

		SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" yaml:"securityGroupIDs,omitempty"`

		SecurityGroups *[]string `tfsdk:"security_groups" yaml:"securityGroups,omitempty"`

		SubnetID *string `tfsdk:"subnet_id" yaml:"subnetID,omitempty"`

		Tags *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tags" yaml:"tags,omitempty"`

		UserData *string `tfsdk:"user_data" yaml:"userData,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewEc2ServicesK8SAwsInstanceV1Alpha1Resource() resource.Resource {
	return &Ec2ServicesK8SAwsInstanceV1Alpha1Resource{}
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ec2_services_k8s_aws_instance_v1alpha1"
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Instance is the Schema for the Instances API",
		MarkdownDescription: "Instance is the Schema for the Instances API",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "InstanceSpec defines the desired state of Instance.  Describes an instance.",
				MarkdownDescription: "InstanceSpec defines the desired state of Instance.  Describes an instance.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"block_device_mappings": {
						Description:         "The block device mapping, which defines the EBS volumes and instance store volumes to attach to the instance at launch. For more information, see Block device mappings (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "The block device mapping, which defines the EBS volumes and instance store volumes to attach to the instance at launch. For more information, see Block device mappings (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) in the Amazon EC2 User Guide.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"device_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ebs": {
								Description:         "Describes a block device for an EBS volume.",
								MarkdownDescription: "Describes a block device for an EBS volume.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"delete_on_termination": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"encrypted": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"iops": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kms_key_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"outpost_arn": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"snapshot_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"throughput": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_size": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"no_device": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"virtual_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"capacity_reservation_specification": {
						Description:         "Information about the Capacity Reservation targeting option. If you do not specify this parameter, the instance's Capacity Reservation preference defaults to open, which enables it to run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).",
						MarkdownDescription: "Information about the Capacity Reservation targeting option. If you do not specify this parameter, the instance's Capacity Reservation preference defaults to open, which enables it to run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"capacity_reservation_preference": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"capacity_reservation_target": {
								Description:         "Describes a target Capacity Reservation or Capacity Reservation group.",
								MarkdownDescription: "Describes a target Capacity Reservation or Capacity Reservation group.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"capacity_reservation_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"capacity_reservation_resource_group_arn": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cpu_options": {
						Description:         "The CPU options for the instance. For more information, see Optimize CPU options (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "The CPU options for the instance. For more information, see Optimize CPU options (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the Amazon EC2 User Guide.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"core_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"threads_per_core": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"credit_specification": {
						Description:         "The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited. To change this attribute after launch, use ModifyInstanceCreditSpecification (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html). For more information, see Burstable performance instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the Amazon EC2 User Guide.  Default: standard (T2 instances) or unlimited (T3/T3a/T4g instances)  For T3 instances with host tenancy, only standard is supported.",
						MarkdownDescription: "The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited. To change this attribute after launch, use ModifyInstanceCreditSpecification (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html). For more information, see Burstable performance instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the Amazon EC2 User Guide.  Default: standard (T2 instances) or unlimited (T3/T3a/T4g instances)  For T3 instances with host tenancy, only standard is supported.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cpu_credits": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_api_stop": {
						Description:         "Indicates whether an instance is enabled for stop protection. For more information, see Stop protection (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection).",
						MarkdownDescription: "Indicates whether an instance is enabled for stop protection. For more information, see Stop protection (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection).",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_api_termination": {
						Description:         "If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute after launch, use ModifyInstanceAttribute (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html). Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate, you can terminate the instance by running the shutdown command from the instance.  Default: false",
						MarkdownDescription: "If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute after launch, use ModifyInstanceAttribute (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html). Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate, you can terminate the instance by running the shutdown command from the instance.  Default: false",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ebs_optimized": {
						Description:         "Indicates whether the instance is optimized for Amazon EBS I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.  Default: false",
						MarkdownDescription: "Indicates whether the instance is optimized for Amazon EBS I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance.  Default: false",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"elastic_gpu_specification": {
						Description:         "An elastic GPU to associate with the instance. An Elastic GPU is a GPU resource that you can attach to your Windows instance to accelerate the graphics performance of your applications. For more information, see Amazon EC2 Elastic GPUs (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "An elastic GPU to associate with the instance. An Elastic GPU is a GPU resource that you can attach to your Windows instance to accelerate the graphics performance of your applications. For more information, see Amazon EC2 Elastic GPUs (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html) in the Amazon EC2 User Guide.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"type_": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"elastic_inference_accelerators": {
						Description:         "An elastic inference accelerator to associate with the instance. Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads.  You cannot specify accelerators from different generations in the same request.",
						MarkdownDescription: "An elastic inference accelerator to associate with the instance. Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads.  You cannot specify accelerators from different generations in the same request.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"type_": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enclave_options": {
						Description:         "Indicates whether the instance is enabled for Amazon Web Services Nitro Enclaves. For more information, see What is Amazon Web Services Nitro Enclaves? (https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the Amazon Web Services Nitro Enclaves User Guide.  You can't enable Amazon Web Services Nitro Enclaves and hibernation on the same instance.",
						MarkdownDescription: "Indicates whether the instance is enabled for Amazon Web Services Nitro Enclaves. For more information, see What is Amazon Web Services Nitro Enclaves? (https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the Amazon Web Services Nitro Enclaves User Guide.  You can't enable Amazon Web Services Nitro Enclaves and hibernation on the same instance.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"hibernation_options": {
						Description:         "Indicates whether an instance is enabled for hibernation. For more information, see Hibernate your instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the Amazon EC2 User Guide.  You can't enable hibernation and Amazon Web Services Nitro Enclaves on the same instance.",
						MarkdownDescription: "Indicates whether an instance is enabled for hibernation. For more information, see Hibernate your instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the Amazon EC2 User Guide.  You can't enable hibernation and Amazon Web Services Nitro Enclaves on the same instance.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"configured": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"iam_instance_profile": {
						Description:         "The name or Amazon Resource Name (ARN) of an IAM instance profile.",
						MarkdownDescription: "The name or Amazon Resource Name (ARN) of an IAM instance profile.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_id": {
						Description:         "The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.",
						MarkdownDescription: "The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"instance_initiated_shutdown_behavior": {
						Description:         "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).  Default: stop",
						MarkdownDescription: "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown).  Default: stop",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"instance_market_options": {
						Description:         "The market (purchasing) option for the instances.  For RunInstances, persistent Spot Instance requests are only supported when InstanceInterruptionBehavior is set to either hibernate or stop.",
						MarkdownDescription: "The market (purchasing) option for the instances.  For RunInstances, persistent Spot Instance requests are only supported when InstanceInterruptionBehavior is set to either hibernate or stop.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"market_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spot_options": {
								Description:         "The options for Spot Instances.",
								MarkdownDescription: "The options for Spot Instances.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"block_duration_minutes": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"instance_interruption_behavior": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_price": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"spot_instance_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"valid_until": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											validators.DateTime64Validator(),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"instance_type": {
						Description:         "The instance type. For more information, see Instance types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the Amazon EC2 User Guide.  Default: m1.small",
						MarkdownDescription: "The instance type. For more information, see Instance types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the Amazon EC2 User Guide.  Default: m1.small",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipv6_address_count": {
						Description:         "[EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch.  You cannot specify this option and the network interfaces option in the same request.",
						MarkdownDescription: "[EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch.  You cannot specify this option and the network interfaces option in the same request.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ipv6_addresses": {
						Description:         "[EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface. You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch.  You cannot specify this option and the network interfaces option in the same request.",
						MarkdownDescription: "[EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface. You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch.  You cannot specify this option and the network interfaces option in the same request.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"ipv6_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kernel_id": {
						Description:         "The ID of the kernel.  We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "The ID of the kernel.  We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the Amazon EC2 User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"key_name": {
						Description:         "The name of the key pair. You can create a key pair using CreateKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or ImportKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html).  If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.",
						MarkdownDescription: "The name of the key pair. You can create a key pair using CreateKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or ImportKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html).  If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"launch_template": {
						Description:         "The launch template to use to launch the instances. Any parameters that you specify in RunInstances override the same parameters in the launch template. You can specify either the name or ID of a launch template, but not both.",
						MarkdownDescription: "The launch template to use to launch the instances. Any parameters that you specify in RunInstances override the same parameters in the launch template. You can specify either the name or ID of a launch template, but not both.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"launch_template_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"launch_template_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"license_specifications": {
						Description:         "The license configurations.",
						MarkdownDescription: "The license configurations.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"license_configuration_arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"maintenance_options": {
						Description:         "The maintenance and recovery options for the instance.",
						MarkdownDescription: "The maintenance and recovery options for the instance.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"auto_recovery": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_count": {
						Description:         "The maximum number of instances to launch. If you specify more instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches the largest possible number of instances above MinCount.  Constraints: Between 1 and the maximum number you're allowed for the specified instance type. For more information about the default limits, and how to request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in the Amazon EC2 FAQ.",
						MarkdownDescription: "The maximum number of instances to launch. If you specify more instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches the largest possible number of instances above MinCount.  Constraints: Between 1 and the maximum number you're allowed for the specified instance type. For more information about the default limits, and how to request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in the Amazon EC2 FAQ.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metadata_options": {
						Description:         "The metadata options for the instance. For more information, see Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).",
						MarkdownDescription: "The metadata options for the instance. For more information, see Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"http_endpoint": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_protocol_i_pv6": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_put_response_hop_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_tokens": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance_metadata_tags": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"min_count": {
						Description:         "The minimum number of instances to launch. If you specify a minimum that is more instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches no instances.  Constraints: Between 1 and the maximum number you're allowed for the specified instance type. For more information about the default limits, and how to request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in the Amazon EC2 General FAQ.",
						MarkdownDescription: "The minimum number of instances to launch. If you specify a minimum that is more instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches no instances.  Constraints: Between 1 and the maximum number you're allowed for the specified instance type. For more information about the default limits, and how to request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in the Amazon EC2 General FAQ.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"monitoring": {
						Description:         "Specifies whether detailed monitoring is enabled for the instance.",
						MarkdownDescription: "Specifies whether detailed monitoring is enabled for the instance.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_interfaces": {
						Description:         "The network interfaces to associate with the instance. If you specify a network interface, you must specify any security groups and subnets as part of the network interface.",
						MarkdownDescription: "The network interfaces to associate with the instance. If you specify a network interface, you must specify any security groups and subnets as part of the network interface.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"associate_carrier_ip_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"associate_public_ip_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"delete_on_termination": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"description": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"device_index": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interface_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipv4_prefix_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipv4_prefixes": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"ipv4_prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipv6_address_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipv6_addresses": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"ipv6_address": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipv6_prefix_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ipv6_prefixes": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"ipv6_prefix": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"network_card_index": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"network_interface_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"private_ip_address": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"private_ip_addresses": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"primary": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"private_ip_address": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secondary_private_ip_address_count": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"subnet_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"placement": {
						Description:         "The placement for the instance.",
						MarkdownDescription: "The placement for the instance.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"affinity": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"availability_zone": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"group_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_resource_group_arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"partition_number": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spread_domain": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tenancy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"private_dns_name_options": {
						Description:         "The options for the instance hostname. The default values are inherited from the subnet.",
						MarkdownDescription: "The options for the instance hostname. The default values are inherited from the subnet.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enable_resource_name_dnsaaaa_record": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_resource_name_dnsa_record": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname_type": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"private_ip_address": {
						Description:         "[EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.  Only one private IP address can be designated as primary. You can't specify this option if you've specified the option to designate a private IP address as the primary IP address in a network interface specification. You cannot specify this option if you're launching more than one instance in the request.  You cannot specify this option and the network interfaces option in the same request.",
						MarkdownDescription: "[EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet.  Only one private IP address can be designated as primary. You can't specify this option if you've specified the option to designate a private IP address as the primary IP address in a network interface specification. You cannot specify this option if you're launching more than one instance in the request.  You cannot specify this option and the network interfaces option in the same request.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ram_disk_id": {
						Description:         "The ID of the RAM disk to select. Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, go to the Amazon Web Services Resource Center and search for the kernel ID.  We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "The ID of the RAM disk to select. Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, go to the Amazon Web Services Resource Center and search for the kernel ID.  We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the Amazon EC2 User Guide.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"security_group_i_ds": {
						Description:         "The IDs of the security groups. You can create a security group using CreateSecurityGroup (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html).  If you specify a network interface, you must specify any security groups as part of the network interface.",
						MarkdownDescription: "The IDs of the security groups. You can create a security group using CreateSecurityGroup (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html).  If you specify a network interface, you must specify any security groups as part of the network interface.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"security_groups": {
						Description:         "[EC2-Classic, default VPC] The names of the security groups. For a nondefault VPC, you must use security group IDs instead.  If you specify a network interface, you must specify any security groups as part of the network interface.  Default: Amazon EC2 uses the default security group.",
						MarkdownDescription: "[EC2-Classic, default VPC] The names of the security groups. For a nondefault VPC, you must use security group IDs instead.  If you specify a network interface, you must specify any security groups as part of the network interface.  Default: Amazon EC2 uses the default security group.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"subnet_id": {
						Description:         "[EC2-VPC] The ID of the subnet to launch the instance into.  If you specify a network interface, you must specify any subnets as part of the network interface.",
						MarkdownDescription: "[EC2-VPC] The ID of the subnet to launch the instance into.  If you specify a network interface, you must specify any subnets as part of the network interface.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tags": {
						Description:         "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",
						MarkdownDescription: "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"value": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"user_data": {
						Description:         "The user data script to make available to the instance. For more information, see Run commands on your Linux instance at launch (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) and Run commands on your Windows instance at launch (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-windows-user-data.html). If you are using a command line tool, base64-encoding is performed for you, and you can load the text from a file. Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.",
						MarkdownDescription: "The user data script to make available to the instance. For more information, see Run commands on your Linux instance at launch (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) and Run commands on your Windows instance at launch (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-windows-user-data.html). If you are using a command line tool, base64-encoding is performed for you, and you can load the text from a file. Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_ec2_services_k8s_aws_instance_v1alpha1")

	var state Ec2ServicesK8SAwsInstanceV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel Ec2ServicesK8SAwsInstanceV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ec2.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Instance")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ec2_services_k8s_aws_instance_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_ec2_services_k8s_aws_instance_v1alpha1")

	var state Ec2ServicesK8SAwsInstanceV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel Ec2ServicesK8SAwsInstanceV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ec2.services.k8s.aws/v1alpha1")
	goModel.Kind = utilities.Ptr("Instance")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_ec2_services_k8s_aws_instance_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
