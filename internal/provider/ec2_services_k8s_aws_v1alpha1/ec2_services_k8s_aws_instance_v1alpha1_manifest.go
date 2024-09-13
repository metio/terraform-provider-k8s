/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package ec2_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &Ec2ServicesK8SAwsInstanceV1Alpha1Manifest{}
)

func NewEc2ServicesK8SAwsInstanceV1Alpha1Manifest() datasource.DataSource {
	return &Ec2ServicesK8SAwsInstanceV1Alpha1Manifest{}
}

type Ec2ServicesK8SAwsInstanceV1Alpha1Manifest struct{}

type Ec2ServicesK8SAwsInstanceV1Alpha1ManifestData struct {
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
		BlockDeviceMappings *[]struct {
			DeviceName *string `tfsdk:"device_name" json:"deviceName,omitempty"`
			Ebs        *struct {
				DeleteOnTermination *bool   `tfsdk:"delete_on_termination" json:"deleteOnTermination,omitempty"`
				Encrypted           *bool   `tfsdk:"encrypted" json:"encrypted,omitempty"`
				Iops                *int64  `tfsdk:"iops" json:"iops,omitempty"`
				KmsKeyID            *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
				OutpostARN          *string `tfsdk:"outpost_arn" json:"outpostARN,omitempty"`
				SnapshotID          *string `tfsdk:"snapshot_id" json:"snapshotID,omitempty"`
				Throughput          *int64  `tfsdk:"throughput" json:"throughput,omitempty"`
				VolumeSize          *int64  `tfsdk:"volume_size" json:"volumeSize,omitempty"`
				VolumeType          *string `tfsdk:"volume_type" json:"volumeType,omitempty"`
			} `tfsdk:"ebs" json:"ebs,omitempty"`
			NoDevice    *string `tfsdk:"no_device" json:"noDevice,omitempty"`
			VirtualName *string `tfsdk:"virtual_name" json:"virtualName,omitempty"`
		} `tfsdk:"block_device_mappings" json:"blockDeviceMappings,omitempty"`
		CapacityReservationSpecification *struct {
			CapacityReservationPreference *string `tfsdk:"capacity_reservation_preference" json:"capacityReservationPreference,omitempty"`
			CapacityReservationTarget     *struct {
				CapacityReservationID               *string `tfsdk:"capacity_reservation_id" json:"capacityReservationID,omitempty"`
				CapacityReservationResourceGroupARN *string `tfsdk:"capacity_reservation_resource_group_arn" json:"capacityReservationResourceGroupARN,omitempty"`
			} `tfsdk:"capacity_reservation_target" json:"capacityReservationTarget,omitempty"`
		} `tfsdk:"capacity_reservation_specification" json:"capacityReservationSpecification,omitempty"`
		CpuOptions *struct {
			CoreCount      *int64 `tfsdk:"core_count" json:"coreCount,omitempty"`
			ThreadsPerCore *int64 `tfsdk:"threads_per_core" json:"threadsPerCore,omitempty"`
		} `tfsdk:"cpu_options" json:"cpuOptions,omitempty"`
		CreditSpecification *struct {
			CpuCredits *string `tfsdk:"cpu_credits" json:"cpuCredits,omitempty"`
		} `tfsdk:"credit_specification" json:"creditSpecification,omitempty"`
		DisableAPIStop          *bool `tfsdk:"disable_api_stop" json:"disableAPIStop,omitempty"`
		DisableAPITermination   *bool `tfsdk:"disable_api_termination" json:"disableAPITermination,omitempty"`
		EbsOptimized            *bool `tfsdk:"ebs_optimized" json:"ebsOptimized,omitempty"`
		ElasticGPUSpecification *[]struct {
			Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
		} `tfsdk:"elastic_gpu_specification" json:"elasticGPUSpecification,omitempty"`
		ElasticInferenceAccelerators *[]struct {
			Count *int64  `tfsdk:"count" json:"count,omitempty"`
			Type_ *string `tfsdk:"type_" json:"type_,omitempty"`
		} `tfsdk:"elastic_inference_accelerators" json:"elasticInferenceAccelerators,omitempty"`
		EnclaveOptions *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"enclave_options" json:"enclaveOptions,omitempty"`
		HibernationOptions *struct {
			Configured *bool `tfsdk:"configured" json:"configured,omitempty"`
		} `tfsdk:"hibernation_options" json:"hibernationOptions,omitempty"`
		IamInstanceProfile *struct {
			Arn  *string `tfsdk:"arn" json:"arn,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"iam_instance_profile" json:"iamInstanceProfile,omitempty"`
		ImageID                           *string `tfsdk:"image_id" json:"imageID,omitempty"`
		InstanceInitiatedShutdownBehavior *string `tfsdk:"instance_initiated_shutdown_behavior" json:"instanceInitiatedShutdownBehavior,omitempty"`
		InstanceMarketOptions             *struct {
			MarketType  *string `tfsdk:"market_type" json:"marketType,omitempty"`
			SpotOptions *struct {
				BlockDurationMinutes         *int64  `tfsdk:"block_duration_minutes" json:"blockDurationMinutes,omitempty"`
				InstanceInterruptionBehavior *string `tfsdk:"instance_interruption_behavior" json:"instanceInterruptionBehavior,omitempty"`
				MaxPrice                     *string `tfsdk:"max_price" json:"maxPrice,omitempty"`
				SpotInstanceType             *string `tfsdk:"spot_instance_type" json:"spotInstanceType,omitempty"`
				ValidUntil                   *string `tfsdk:"valid_until" json:"validUntil,omitempty"`
			} `tfsdk:"spot_options" json:"spotOptions,omitempty"`
		} `tfsdk:"instance_market_options" json:"instanceMarketOptions,omitempty"`
		InstanceType     *string `tfsdk:"instance_type" json:"instanceType,omitempty"`
		Ipv6AddressCount *int64  `tfsdk:"ipv6_address_count" json:"ipv6AddressCount,omitempty"`
		Ipv6Addresses    *[]struct {
			Ipv6Address *string `tfsdk:"ipv6_address" json:"ipv6Address,omitempty"`
		} `tfsdk:"ipv6_addresses" json:"ipv6Addresses,omitempty"`
		KernelID       *string `tfsdk:"kernel_id" json:"kernelID,omitempty"`
		KeyName        *string `tfsdk:"key_name" json:"keyName,omitempty"`
		LaunchTemplate *struct {
			LaunchTemplateID   *string `tfsdk:"launch_template_id" json:"launchTemplateID,omitempty"`
			LaunchTemplateName *string `tfsdk:"launch_template_name" json:"launchTemplateName,omitempty"`
			Version            *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"launch_template" json:"launchTemplate,omitempty"`
		LicenseSpecifications *[]struct {
			LicenseConfigurationARN *string `tfsdk:"license_configuration_arn" json:"licenseConfigurationARN,omitempty"`
		} `tfsdk:"license_specifications" json:"licenseSpecifications,omitempty"`
		MaintenanceOptions *struct {
			AutoRecovery *string `tfsdk:"auto_recovery" json:"autoRecovery,omitempty"`
		} `tfsdk:"maintenance_options" json:"maintenanceOptions,omitempty"`
		MaxCount        *int64 `tfsdk:"max_count" json:"maxCount,omitempty"`
		MetadataOptions *struct {
			HttpEndpoint            *string `tfsdk:"http_endpoint" json:"httpEndpoint,omitempty"`
			HttpProtocolIPv6        *string `tfsdk:"http_protocol_i_pv6" json:"httpProtocolIPv6,omitempty"`
			HttpPutResponseHopLimit *int64  `tfsdk:"http_put_response_hop_limit" json:"httpPutResponseHopLimit,omitempty"`
			HttpTokens              *string `tfsdk:"http_tokens" json:"httpTokens,omitempty"`
			InstanceMetadataTags    *string `tfsdk:"instance_metadata_tags" json:"instanceMetadataTags,omitempty"`
		} `tfsdk:"metadata_options" json:"metadataOptions,omitempty"`
		MinCount   *int64 `tfsdk:"min_count" json:"minCount,omitempty"`
		Monitoring *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"monitoring" json:"monitoring,omitempty"`
		NetworkInterfaces *[]struct {
			AssociateCarrierIPAddress *bool   `tfsdk:"associate_carrier_ip_address" json:"associateCarrierIPAddress,omitempty"`
			AssociatePublicIPAddress  *bool   `tfsdk:"associate_public_ip_address" json:"associatePublicIPAddress,omitempty"`
			DeleteOnTermination       *bool   `tfsdk:"delete_on_termination" json:"deleteOnTermination,omitempty"`
			Description               *string `tfsdk:"description" json:"description,omitempty"`
			DeviceIndex               *int64  `tfsdk:"device_index" json:"deviceIndex,omitempty"`
			InterfaceType             *string `tfsdk:"interface_type" json:"interfaceType,omitempty"`
			Ipv4PrefixCount           *int64  `tfsdk:"ipv4_prefix_count" json:"ipv4PrefixCount,omitempty"`
			Ipv4Prefixes              *[]struct {
				Ipv4Prefix *string `tfsdk:"ipv4_prefix" json:"ipv4Prefix,omitempty"`
			} `tfsdk:"ipv4_prefixes" json:"ipv4Prefixes,omitempty"`
			Ipv6AddressCount *int64 `tfsdk:"ipv6_address_count" json:"ipv6AddressCount,omitempty"`
			Ipv6Addresses    *[]struct {
				Ipv6Address *string `tfsdk:"ipv6_address" json:"ipv6Address,omitempty"`
			} `tfsdk:"ipv6_addresses" json:"ipv6Addresses,omitempty"`
			Ipv6PrefixCount *int64 `tfsdk:"ipv6_prefix_count" json:"ipv6PrefixCount,omitempty"`
			Ipv6Prefixes    *[]struct {
				Ipv6Prefix *string `tfsdk:"ipv6_prefix" json:"ipv6Prefix,omitempty"`
			} `tfsdk:"ipv6_prefixes" json:"ipv6Prefixes,omitempty"`
			NetworkCardIndex   *int64  `tfsdk:"network_card_index" json:"networkCardIndex,omitempty"`
			NetworkInterfaceID *string `tfsdk:"network_interface_id" json:"networkInterfaceID,omitempty"`
			PrivateIPAddress   *string `tfsdk:"private_ip_address" json:"privateIPAddress,omitempty"`
			PrivateIPAddresses *[]struct {
				Primary          *bool   `tfsdk:"primary" json:"primary,omitempty"`
				PrivateIPAddress *string `tfsdk:"private_ip_address" json:"privateIPAddress,omitempty"`
			} `tfsdk:"private_ip_addresses" json:"privateIPAddresses,omitempty"`
			SecondaryPrivateIPAddressCount *int64  `tfsdk:"secondary_private_ip_address_count" json:"secondaryPrivateIPAddressCount,omitempty"`
			SubnetID                       *string `tfsdk:"subnet_id" json:"subnetID,omitempty"`
		} `tfsdk:"network_interfaces" json:"networkInterfaces,omitempty"`
		Placement *struct {
			Affinity             *string `tfsdk:"affinity" json:"affinity,omitempty"`
			AvailabilityZone     *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
			GroupName            *string `tfsdk:"group_name" json:"groupName,omitempty"`
			HostID               *string `tfsdk:"host_id" json:"hostID,omitempty"`
			HostResourceGroupARN *string `tfsdk:"host_resource_group_arn" json:"hostResourceGroupARN,omitempty"`
			PartitionNumber      *int64  `tfsdk:"partition_number" json:"partitionNumber,omitempty"`
			SpreadDomain         *string `tfsdk:"spread_domain" json:"spreadDomain,omitempty"`
			Tenancy              *string `tfsdk:"tenancy" json:"tenancy,omitempty"`
		} `tfsdk:"placement" json:"placement,omitempty"`
		PrivateDNSNameOptions *struct {
			EnableResourceNameDNSAAAARecord *bool   `tfsdk:"enable_resource_name_dnsaaaa_record" json:"enableResourceNameDNSAAAARecord,omitempty"`
			EnableResourceNameDNSARecord    *bool   `tfsdk:"enable_resource_name_dnsa_record" json:"enableResourceNameDNSARecord,omitempty"`
			HostnameType                    *string `tfsdk:"hostname_type" json:"hostnameType,omitempty"`
		} `tfsdk:"private_dns_name_options" json:"privateDNSNameOptions,omitempty"`
		PrivateIPAddress *string   `tfsdk:"private_ip_address" json:"privateIPAddress,omitempty"`
		RamDiskID        *string   `tfsdk:"ram_disk_id" json:"ramDiskID,omitempty"`
		SecurityGroupIDs *[]string `tfsdk:"security_group_i_ds" json:"securityGroupIDs,omitempty"`
		SecurityGroups   *[]string `tfsdk:"security_groups" json:"securityGroups,omitempty"`
		SubnetID         *string   `tfsdk:"subnet_id" json:"subnetID,omitempty"`
		Tags             *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		UserData *string `tfsdk:"user_data" json:"userData,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_ec2_services_k8s_aws_instance_v1alpha1_manifest"
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Instance is the Schema for the Instances API",
		MarkdownDescription: "Instance is the Schema for the Instances API",
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
				Description:         "InstanceSpec defines the desired state of Instance. Describes an instance.",
				MarkdownDescription: "InstanceSpec defines the desired state of Instance. Describes an instance.",
				Attributes: map[string]schema.Attribute{
					"block_device_mappings": schema.ListNestedAttribute{
						Description:         "The block device mapping, which defines the EBS volumes and instance store volumes to attach to the instance at launch. For more information, see Block device mappings (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "The block device mapping, which defines the EBS volumes and instance store volumes to attach to the instance at launch. For more information, see Block device mappings (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html) in the Amazon EC2 User Guide.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"device_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ebs": schema.SingleNestedAttribute{
									Description:         "Describes a block device for an EBS volume.",
									MarkdownDescription: "Describes a block device for an EBS volume.",
									Attributes: map[string]schema.Attribute{
										"delete_on_termination": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"encrypted": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"iops": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kms_key_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"outpost_arn": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"snapshot_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"throughput": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_size": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"no_device": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"virtual_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"capacity_reservation_specification": schema.SingleNestedAttribute{
						Description:         "Information about the Capacity Reservation targeting option. If you do not specify this parameter, the instance's Capacity Reservation preference defaults to open, which enables it to run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).",
						MarkdownDescription: "Information about the Capacity Reservation targeting option. If you do not specify this parameter, the instance's Capacity Reservation preference defaults to open, which enables it to run in any open Capacity Reservation that has matching attributes (instance type, platform, Availability Zone).",
						Attributes: map[string]schema.Attribute{
							"capacity_reservation_preference": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"capacity_reservation_target": schema.SingleNestedAttribute{
								Description:         "Describes a target Capacity Reservation or Capacity Reservation group.",
								MarkdownDescription: "Describes a target Capacity Reservation or Capacity Reservation group.",
								Attributes: map[string]schema.Attribute{
									"capacity_reservation_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"capacity_reservation_resource_group_arn": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

					"cpu_options": schema.SingleNestedAttribute{
						Description:         "The CPU options for the instance. For more information, see Optimize CPU options (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "The CPU options for the instance. For more information, see Optimize CPU options (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-optimize-cpu.html) in the Amazon EC2 User Guide.",
						Attributes: map[string]schema.Attribute{
							"core_count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"threads_per_core": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"credit_specification": schema.SingleNestedAttribute{
						Description:         "The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited. To change this attribute after launch, use ModifyInstanceCreditSpecification (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html). For more information, see Burstable performance instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the Amazon EC2 User Guide. Default: standard (T2 instances) or unlimited (T3/T3a/T4g instances) For T3 instances with host tenancy, only standard is supported.",
						MarkdownDescription: "The credit option for CPU usage of the burstable performance instance. Valid values are standard and unlimited. To change this attribute after launch, use ModifyInstanceCreditSpecification (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceCreditSpecification.html). For more information, see Burstable performance instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html) in the Amazon EC2 User Guide. Default: standard (T2 instances) or unlimited (T3/T3a/T4g instances) For T3 instances with host tenancy, only standard is supported.",
						Attributes: map[string]schema.Attribute{
							"cpu_credits": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_api_stop": schema.BoolAttribute{
						Description:         "Indicates whether an instance is enabled for stop protection. For more information, see Stop protection (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection).",
						MarkdownDescription: "Indicates whether an instance is enabled for stop protection. For more information, see Stop protection (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Stop_Start.html#Using_StopProtection).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_api_termination": schema.BoolAttribute{
						Description:         "If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute after launch, use ModifyInstanceAttribute (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html). Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate, you can terminate the instance by running the shutdown command from the instance. Default: false",
						MarkdownDescription: "If you set this parameter to true, you can't terminate the instance using the Amazon EC2 console, CLI, or API; otherwise, you can. To change this attribute after launch, use ModifyInstanceAttribute (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyInstanceAttribute.html). Alternatively, if you set InstanceInitiatedShutdownBehavior to terminate, you can terminate the instance by running the shutdown command from the instance. Default: false",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ebs_optimized": schema.BoolAttribute{
						Description:         "Indicates whether the instance is optimized for Amazon EBS I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance. Default: false",
						MarkdownDescription: "Indicates whether the instance is optimized for Amazon EBS I/O. This optimization provides dedicated throughput to Amazon EBS and an optimized configuration stack to provide optimal Amazon EBS I/O performance. This optimization isn't available with all instance types. Additional usage charges apply when using an EBS-optimized instance. Default: false",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"elastic_gpu_specification": schema.ListNestedAttribute{
						Description:         "An elastic GPU to associate with the instance. An Elastic GPU is a GPU resource that you can attach to your Windows instance to accelerate the graphics performance of your applications. For more information, see Amazon EC2 Elastic GPUs (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "An elastic GPU to associate with the instance. An Elastic GPU is a GPU resource that you can attach to your Windows instance to accelerate the graphics performance of your applications. For more information, see Amazon EC2 Elastic GPUs (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-graphics.html) in the Amazon EC2 User Guide.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type_": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"elastic_inference_accelerators": schema.ListNestedAttribute{
						Description:         "An elastic inference accelerator to associate with the instance. Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads. You cannot specify accelerators from different generations in the same request.",
						MarkdownDescription: "An elastic inference accelerator to associate with the instance. Elastic inference accelerators are a resource you can attach to your Amazon EC2 instances to accelerate your Deep Learning (DL) inference workloads. You cannot specify accelerators from different generations in the same request.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"count": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"type_": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"enclave_options": schema.SingleNestedAttribute{
						Description:         "Indicates whether the instance is enabled for Amazon Web Services Nitro Enclaves. For more information, see What is Amazon Web Services Nitro Enclaves? (https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the Amazon Web Services Nitro Enclaves User Guide. You can't enable Amazon Web Services Nitro Enclaves and hibernation on the same instance.",
						MarkdownDescription: "Indicates whether the instance is enabled for Amazon Web Services Nitro Enclaves. For more information, see What is Amazon Web Services Nitro Enclaves? (https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html) in the Amazon Web Services Nitro Enclaves User Guide. You can't enable Amazon Web Services Nitro Enclaves and hibernation on the same instance.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"hibernation_options": schema.SingleNestedAttribute{
						Description:         "Indicates whether an instance is enabled for hibernation. For more information, see Hibernate your instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the Amazon EC2 User Guide. You can't enable hibernation and Amazon Web Services Nitro Enclaves on the same instance.",
						MarkdownDescription: "Indicates whether an instance is enabled for hibernation. For more information, see Hibernate your instance (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Hibernate.html) in the Amazon EC2 User Guide. You can't enable hibernation and Amazon Web Services Nitro Enclaves on the same instance.",
						Attributes: map[string]schema.Attribute{
							"configured": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"iam_instance_profile": schema.SingleNestedAttribute{
						Description:         "The name or Amazon Resource Name (ARN) of an IAM instance profile.",
						MarkdownDescription: "The name or Amazon Resource Name (ARN) of an IAM instance profile.",
						Attributes: map[string]schema.Attribute{
							"arn": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_id": schema.StringAttribute{
						Description:         "The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.",
						MarkdownDescription: "The ID of the AMI. An AMI ID is required to launch an instance and must be specified here or in a launch template.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance_initiated_shutdown_behavior": schema.StringAttribute{
						Description:         "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown). Default: stop",
						MarkdownDescription: "Indicates whether an instance stops or terminates when you initiate shutdown from the instance (using the operating system command for system shutdown). Default: stop",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance_market_options": schema.SingleNestedAttribute{
						Description:         "The market (purchasing) option for the instances. For RunInstances, persistent Spot Instance requests are only supported when InstanceInterruptionBehavior is set to either hibernate or stop.",
						MarkdownDescription: "The market (purchasing) option for the instances. For RunInstances, persistent Spot Instance requests are only supported when InstanceInterruptionBehavior is set to either hibernate or stop.",
						Attributes: map[string]schema.Attribute{
							"market_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spot_options": schema.SingleNestedAttribute{
								Description:         "The options for Spot Instances.",
								MarkdownDescription: "The options for Spot Instances.",
								Attributes: map[string]schema.Attribute{
									"block_duration_minutes": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"instance_interruption_behavior": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_price": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"spot_instance_type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"valid_until": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.DateTime64Validator(),
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

					"instance_type": schema.StringAttribute{
						Description:         "The instance type. For more information, see Instance types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the Amazon EC2 User Guide. Default: m1.small",
						MarkdownDescription: "The instance type. For more information, see Instance types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the Amazon EC2 User Guide. Default: m1.small",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_address_count": schema.Int64Attribute{
						Description:         "[EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch. You cannot specify this option and the network interfaces option in the same request.",
						MarkdownDescription: "[EC2-VPC] The number of IPv6 addresses to associate with the primary network interface. Amazon EC2 chooses the IPv6 addresses from the range of your subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch. You cannot specify this option and the network interfaces option in the same request.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_addresses": schema.ListNestedAttribute{
						Description:         "[EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface. You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch. You cannot specify this option and the network interfaces option in the same request.",
						MarkdownDescription: "[EC2-VPC] The IPv6 addresses from the range of the subnet to associate with the primary network interface. You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch. You cannot specify this option and the network interfaces option in the same request.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ipv6_address": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"kernel_id": schema.StringAttribute{
						Description:         "The ID of the kernel. We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "The ID of the kernel. We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the Amazon EC2 User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_name": schema.StringAttribute{
						Description:         "The name of the key pair. You can create a key pair using CreateKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or ImportKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html). If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.",
						MarkdownDescription: "The name of the key pair. You can create a key pair using CreateKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateKeyPair.html) or ImportKeyPair (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ImportKeyPair.html). If you do not specify a key pair, you can't connect to the instance unless you choose an AMI that is configured to allow users another way to log in.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"launch_template": schema.SingleNestedAttribute{
						Description:         "The launch template to use to launch the instances. Any parameters that you specify in RunInstances override the same parameters in the launch template. You can specify either the name or ID of a launch template, but not both.",
						MarkdownDescription: "The launch template to use to launch the instances. Any parameters that you specify in RunInstances override the same parameters in the launch template. You can specify either the name or ID of a launch template, but not both.",
						Attributes: map[string]schema.Attribute{
							"launch_template_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"launch_template_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"license_specifications": schema.ListNestedAttribute{
						Description:         "The license configurations.",
						MarkdownDescription: "The license configurations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"license_configuration_arn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"maintenance_options": schema.SingleNestedAttribute{
						Description:         "The maintenance and recovery options for the instance.",
						MarkdownDescription: "The maintenance and recovery options for the instance.",
						Attributes: map[string]schema.Attribute{
							"auto_recovery": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_count": schema.Int64Attribute{
						Description:         "The maximum number of instances to launch. If you specify more instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches the largest possible number of instances above MinCount. Constraints: Between 1 and the maximum number you're allowed for the specified instance type. For more information about the default limits, and how to request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in the Amazon EC2 FAQ.",
						MarkdownDescription: "The maximum number of instances to launch. If you specify more instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches the largest possible number of instances above MinCount. Constraints: Between 1 and the maximum number you're allowed for the specified instance type. For more information about the default limits, and how to request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in the Amazon EC2 FAQ.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metadata_options": schema.SingleNestedAttribute{
						Description:         "The metadata options for the instance. For more information, see Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).",
						MarkdownDescription: "The metadata options for the instance. For more information, see Instance metadata and user data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).",
						Attributes: map[string]schema.Attribute{
							"http_endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_protocol_i_pv6": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_put_response_hop_limit": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_tokens": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"instance_metadata_tags": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"min_count": schema.Int64Attribute{
						Description:         "The minimum number of instances to launch. If you specify a minimum that is more instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches no instances. Constraints: Between 1 and the maximum number you're allowed for the specified instance type. For more information about the default limits, and how to request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in the Amazon EC2 General FAQ.",
						MarkdownDescription: "The minimum number of instances to launch. If you specify a minimum that is more instances than Amazon EC2 can launch in the target Availability Zone, Amazon EC2 launches no instances. Constraints: Between 1 and the maximum number you're allowed for the specified instance type. For more information about the default limits, and how to request an increase, see How many instances can I run in Amazon EC2 (http://aws.amazon.com/ec2/faqs/#How_many_instances_can_I_run_in_Amazon_EC2) in the Amazon EC2 General FAQ.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monitoring": schema.SingleNestedAttribute{
						Description:         "Specifies whether detailed monitoring is enabled for the instance.",
						MarkdownDescription: "Specifies whether detailed monitoring is enabled for the instance.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network_interfaces": schema.ListNestedAttribute{
						Description:         "The network interfaces to associate with the instance. If you specify a network interface, you must specify any security groups and subnets as part of the network interface.",
						MarkdownDescription: "The network interfaces to associate with the instance. If you specify a network interface, you must specify any security groups and subnets as part of the network interface.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"associate_carrier_ip_address": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"associate_public_ip_address": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"delete_on_termination": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"device_index": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"interface_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ipv4_prefix_count": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ipv4_prefixes": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ipv4_prefix": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

								"ipv6_address_count": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ipv6_addresses": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ipv6_address": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

								"ipv6_prefix_count": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ipv6_prefixes": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"ipv6_prefix": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

								"network_card_index": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"network_interface_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"private_ip_address": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"private_ip_addresses": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"primary": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"private_ip_address": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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

								"secondary_private_ip_address_count": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"subnet_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"placement": schema.SingleNestedAttribute{
						Description:         "The placement for the instance.",
						MarkdownDescription: "The placement for the instance.",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"availability_zone": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_resource_group_arn": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"partition_number": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spread_domain": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tenancy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"private_dns_name_options": schema.SingleNestedAttribute{
						Description:         "The options for the instance hostname. The default values are inherited from the subnet.",
						MarkdownDescription: "The options for the instance hostname. The default values are inherited from the subnet.",
						Attributes: map[string]schema.Attribute{
							"enable_resource_name_dnsaaaa_record": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_resource_name_dnsa_record": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"private_ip_address": schema.StringAttribute{
						Description:         "[EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet. Only one private IP address can be designated as primary. You can't specify this option if you've specified the option to designate a private IP address as the primary IP address in a network interface specification. You cannot specify this option if you're launching more than one instance in the request. You cannot specify this option and the network interfaces option in the same request.",
						MarkdownDescription: "[EC2-VPC] The primary IPv4 address. You must specify a value from the IPv4 address range of the subnet. Only one private IP address can be designated as primary. You can't specify this option if you've specified the option to designate a private IP address as the primary IP address in a network interface specification. You cannot specify this option if you're launching more than one instance in the request. You cannot specify this option and the network interfaces option in the same request.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ram_disk_id": schema.StringAttribute{
						Description:         "The ID of the RAM disk to select. Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, go to the Amazon Web Services Resource Center and search for the kernel ID. We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the Amazon EC2 User Guide.",
						MarkdownDescription: "The ID of the RAM disk to select. Some kernels require additional drivers at launch. Check the kernel requirements for information about whether you need to specify a RAM disk. To find kernel requirements, go to the Amazon Web Services Resource Center and search for the kernel ID. We recommend that you use PV-GRUB instead of kernels and RAM disks. For more information, see PV-GRUB (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/UserProvidedkernels.html) in the Amazon EC2 User Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_group_i_ds": schema.ListAttribute{
						Description:         "The IDs of the security groups. You can create a security group using CreateSecurityGroup (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html). If you specify a network interface, you must specify any security groups as part of the network interface.",
						MarkdownDescription: "The IDs of the security groups. You can create a security group using CreateSecurityGroup (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSecurityGroup.html). If you specify a network interface, you must specify any security groups as part of the network interface.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_groups": schema.ListAttribute{
						Description:         "[EC2-Classic, default VPC] The names of the security groups. For a nondefault VPC, you must use security group IDs instead. If you specify a network interface, you must specify any security groups as part of the network interface. Default: Amazon EC2 uses the default security group.",
						MarkdownDescription: "[EC2-Classic, default VPC] The names of the security groups. For a nondefault VPC, you must use security group IDs instead. If you specify a network interface, you must specify any security groups as part of the network interface. Default: Amazon EC2 uses the default security group.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subnet_id": schema.StringAttribute{
						Description:         "[EC2-VPC] The ID of the subnet to launch the instance into. If you specify a network interface, you must specify any subnets as part of the network interface.",
						MarkdownDescription: "[EC2-VPC] The ID of the subnet to launch the instance into. If you specify a network interface, you must specify any subnets as part of the network interface.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",
						MarkdownDescription: "The tags. The value parameter is required, but if you don't want the tag to have a value, specify the parameter with no value, and we set the value to an empty string.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"user_data": schema.StringAttribute{
						Description:         "The user data script to make available to the instance. For more information, see Run commands on your Linux instance at launch (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) and Run commands on your Windows instance at launch (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-windows-user-data.html). If you are using a command line tool, base64-encoding is performed for you, and you can load the text from a file. Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.",
						MarkdownDescription: "The user data script to make available to the instance. For more information, see Run commands on your Linux instance at launch (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html) and Run commands on your Windows instance at launch (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ec2-windows-user-data.html). If you are using a command line tool, base64-encoding is performed for you, and you can load the text from a file. Otherwise, you must provide base64-encoded text. User data is limited to 16 KB.",
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
	}
}

func (r *Ec2ServicesK8SAwsInstanceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ec2_services_k8s_aws_instance_v1alpha1_manifest")

	var model Ec2ServicesK8SAwsInstanceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("ec2.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("Instance")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
