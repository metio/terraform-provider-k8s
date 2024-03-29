/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package karpenter_k8s_aws_v1beta1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KarpenterK8SAwsEc2NodeClassV1Beta1Manifest{}
)

func NewKarpenterK8SAwsEc2NodeClassV1Beta1Manifest() datasource.DataSource {
	return &KarpenterK8SAwsEc2NodeClassV1Beta1Manifest{}
}

type KarpenterK8SAwsEc2NodeClassV1Beta1Manifest struct{}

type KarpenterK8SAwsEc2NodeClassV1Beta1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		AmiFamily        *string `tfsdk:"ami_family" json:"amiFamily,omitempty"`
		AmiSelectorTerms *[]struct {
			Id    *string            `tfsdk:"id" json:"id,omitempty"`
			Name  *string            `tfsdk:"name" json:"name,omitempty"`
			Owner *string            `tfsdk:"owner" json:"owner,omitempty"`
			Tags  *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"ami_selector_terms" json:"amiSelectorTerms,omitempty"`
		AssociatePublicIPAddress *bool `tfsdk:"associate_public_ip_address" json:"associatePublicIPAddress,omitempty"`
		BlockDeviceMappings      *[]struct {
			DeviceName *string `tfsdk:"device_name" json:"deviceName,omitempty"`
			Ebs        *struct {
				DeleteOnTermination *bool   `tfsdk:"delete_on_termination" json:"deleteOnTermination,omitempty"`
				Encrypted           *bool   `tfsdk:"encrypted" json:"encrypted,omitempty"`
				Iops                *int64  `tfsdk:"iops" json:"iops,omitempty"`
				KmsKeyID            *string `tfsdk:"kms_key_id" json:"kmsKeyID,omitempty"`
				SnapshotID          *string `tfsdk:"snapshot_id" json:"snapshotID,omitempty"`
				Throughput          *int64  `tfsdk:"throughput" json:"throughput,omitempty"`
				VolumeSize          *string `tfsdk:"volume_size" json:"volumeSize,omitempty"`
				VolumeType          *string `tfsdk:"volume_type" json:"volumeType,omitempty"`
			} `tfsdk:"ebs" json:"ebs,omitempty"`
			RootVolume *bool `tfsdk:"root_volume" json:"rootVolume,omitempty"`
		} `tfsdk:"block_device_mappings" json:"blockDeviceMappings,omitempty"`
		Context             *string `tfsdk:"context" json:"context,omitempty"`
		DetailedMonitoring  *bool   `tfsdk:"detailed_monitoring" json:"detailedMonitoring,omitempty"`
		InstanceProfile     *string `tfsdk:"instance_profile" json:"instanceProfile,omitempty"`
		InstanceStorePolicy *string `tfsdk:"instance_store_policy" json:"instanceStorePolicy,omitempty"`
		MetadataOptions     *struct {
			HttpEndpoint            *string `tfsdk:"http_endpoint" json:"httpEndpoint,omitempty"`
			HttpProtocolIPv6        *string `tfsdk:"http_protocol_i_pv6" json:"httpProtocolIPv6,omitempty"`
			HttpPutResponseHopLimit *int64  `tfsdk:"http_put_response_hop_limit" json:"httpPutResponseHopLimit,omitempty"`
			HttpTokens              *string `tfsdk:"http_tokens" json:"httpTokens,omitempty"`
		} `tfsdk:"metadata_options" json:"metadataOptions,omitempty"`
		Role                       *string `tfsdk:"role" json:"role,omitempty"`
		SecurityGroupSelectorTerms *[]struct {
			Id   *string            `tfsdk:"id" json:"id,omitempty"`
			Name *string            `tfsdk:"name" json:"name,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"security_group_selector_terms" json:"securityGroupSelectorTerms,omitempty"`
		SubnetSelectorTerms *[]struct {
			Id   *string            `tfsdk:"id" json:"id,omitempty"`
			Tags *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		} `tfsdk:"subnet_selector_terms" json:"subnetSelectorTerms,omitempty"`
		Tags     *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
		UserData *string            `tfsdk:"user_data" json:"userData,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KarpenterK8SAwsEc2NodeClassV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_karpenter_k8s_aws_ec2_node_class_v1beta1_manifest"
}

func (r *KarpenterK8SAwsEc2NodeClassV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "EC2NodeClass is the Schema for the EC2NodeClass API",
		MarkdownDescription: "EC2NodeClass is the Schema for the EC2NodeClass API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "EC2NodeClassSpec is the top level specification for the AWS Karpenter Provider.This will contain configuration necessary to launch instances in AWS.",
				MarkdownDescription: "EC2NodeClassSpec is the top level specification for the AWS Karpenter Provider.This will contain configuration necessary to launch instances in AWS.",
				Attributes: map[string]schema.Attribute{
					"ami_family": schema.StringAttribute{
						Description:         "AMIFamily is the AMI family that instances use.",
						MarkdownDescription: "AMIFamily is the AMI family that instances use.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("AL2", "AL2023", "Bottlerocket", "Ubuntu", "Custom", "Windows2019", "Windows2022"),
						},
					},

					"ami_selector_terms": schema.ListNestedAttribute{
						Description:         "AMISelectorTerms is a list of or ami selector terms. The terms are ORed.",
						MarkdownDescription: "AMISelectorTerms is a list of or ami selector terms. The terms are ORed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "ID is the ami id in EC2",
									MarkdownDescription: "ID is the ami id in EC2",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`ami-[0-9a-z]+`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the ami name in EC2.This value is the name field, which is different from the name tag.",
									MarkdownDescription: "Name is the ami name in EC2.This value is the name field, which is different from the name tag.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"owner": schema.StringAttribute{
									Description:         "Owner is the owner for the ami.You can specify a combination of AWS account IDs, 'self', 'amazon', and 'aws-marketplace'",
									MarkdownDescription: "Owner is the owner for the ami.You can specify a combination of AWS account IDs, 'self', 'amazon', and 'aws-marketplace'",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tags": schema.MapAttribute{
									Description:         "Tags is a map of key/value tags used to select subnetsSpecifying '*' for a value selects all values for a given tag key.",
									MarkdownDescription: "Tags is a map of key/value tags used to select subnetsSpecifying '*' for a value selects all values for a given tag key.",
									ElementType:         types.StringType,
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

					"associate_public_ip_address": schema.BoolAttribute{
						Description:         "AssociatePublicIPAddress controls if public IP addresses are assigned to instances that are launched with the nodeclass.",
						MarkdownDescription: "AssociatePublicIPAddress controls if public IP addresses are assigned to instances that are launched with the nodeclass.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"block_device_mappings": schema.ListNestedAttribute{
						Description:         "BlockDeviceMappings to be applied to provisioned nodes.",
						MarkdownDescription: "BlockDeviceMappings to be applied to provisioned nodes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"device_name": schema.StringAttribute{
									Description:         "The device name (for example, /dev/sdh or xvdh).",
									MarkdownDescription: "The device name (for example, /dev/sdh or xvdh).",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ebs": schema.SingleNestedAttribute{
									Description:         "EBS contains parameters used to automatically set up EBS volumes when an instance is launched.",
									MarkdownDescription: "EBS contains parameters used to automatically set up EBS volumes when an instance is launched.",
									Attributes: map[string]schema.Attribute{
										"delete_on_termination": schema.BoolAttribute{
											Description:         "DeleteOnTermination indicates whether the EBS volume is deleted on instance termination.",
											MarkdownDescription: "DeleteOnTermination indicates whether the EBS volume is deleted on instance termination.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"encrypted": schema.BoolAttribute{
											Description:         "Encrypted indicates whether the EBS volume is encrypted. Encrypted volumes can onlybe attached to instances that support Amazon EBS encryption. If you are creatinga volume from a snapshot, you can't specify an encryption value.",
											MarkdownDescription: "Encrypted indicates whether the EBS volume is encrypted. Encrypted volumes can onlybe attached to instances that support Amazon EBS encryption. If you are creatinga volume from a snapshot, you can't specify an encryption value.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"iops": schema.Int64Attribute{
											Description:         "IOPS is the number of I/O operations per second (IOPS). For gp3, io1, and io2 volumes,this represents the number of IOPS that are provisioned for the volume. Forgp2 volumes, this represents the baseline performance of the volume and therate at which the volume accumulates I/O credits for bursting.The following are the supported values for each volume type:   * gp3: 3,000-16,000 IOPS   * io1: 100-64,000 IOPS   * io2: 100-64,000 IOPSFor io1 and io2 volumes, we guarantee 64,000 IOPS only for Instances builton the Nitro System (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).Other instance families guarantee performance up to 32,000 IOPS.This parameter is supported for io1, io2, and gp3 volumes only. This parameteris not supported for gp2, st1, sc1, or standard volumes.",
											MarkdownDescription: "IOPS is the number of I/O operations per second (IOPS). For gp3, io1, and io2 volumes,this represents the number of IOPS that are provisioned for the volume. Forgp2 volumes, this represents the baseline performance of the volume and therate at which the volume accumulates I/O credits for bursting.The following are the supported values for each volume type:   * gp3: 3,000-16,000 IOPS   * io1: 100-64,000 IOPS   * io2: 100-64,000 IOPSFor io1 and io2 volumes, we guarantee 64,000 IOPS only for Instances builton the Nitro System (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances).Other instance families guarantee performance up to 32,000 IOPS.This parameter is supported for io1, io2, and gp3 volumes only. This parameteris not supported for gp2, st1, sc1, or standard volumes.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kms_key_id": schema.StringAttribute{
											Description:         "KMSKeyID (ARN) of the symmetric Key Management Service (KMS) CMK used for encryption.",
											MarkdownDescription: "KMSKeyID (ARN) of the symmetric Key Management Service (KMS) CMK used for encryption.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"snapshot_id": schema.StringAttribute{
											Description:         "SnapshotID is the ID of an EBS snapshot",
											MarkdownDescription: "SnapshotID is the ID of an EBS snapshot",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"throughput": schema.Int64Attribute{
											Description:         "Throughput to provision for a gp3 volume, with a maximum of 1,000 MiB/s.Valid Range: Minimum value of 125. Maximum value of 1000.",
											MarkdownDescription: "Throughput to provision for a gp3 volume, with a maximum of 1,000 MiB/s.Valid Range: Minimum value of 125. Maximum value of 1000.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_size": schema.StringAttribute{
											Description:         "VolumeSize in 'Gi', 'G', 'Ti', or 'T'. You must specify either a snapshot ID ora volume size. The following are the supported volumes sizes for each volumetype:   * gp2 and gp3: 1-16,384   * io1 and io2: 4-16,384   * st1 and sc1: 125-16,384   * standard: 1-1,024",
											MarkdownDescription: "VolumeSize in 'Gi', 'G', 'Ti', or 'T'. You must specify either a snapshot ID ora volume size. The following are the supported volumes sizes for each volumetype:   * gp2 and gp3: 1-16,384   * io1 and io2: 4-16,384   * st1 and sc1: 125-16,384   * standard: 1-1,024",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_type": schema.StringAttribute{
											Description:         "VolumeType of the block device.For more information, see Amazon EBS volume types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)in the Amazon Elastic Compute Cloud User Guide.",
											MarkdownDescription: "VolumeType of the block device.For more information, see Amazon EBS volume types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html)in the Amazon Elastic Compute Cloud User Guide.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("standard", "io1", "io2", "gp2", "sc1", "st1", "gp3"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"root_volume": schema.BoolAttribute{
									Description:         "RootVolume is a flag indicating if this device is mounted as kubelet root dir. You canconfigure at most one root volume in BlockDeviceMappings.",
									MarkdownDescription: "RootVolume is a flag indicating if this device is mounted as kubelet root dir. You canconfigure at most one root volume in BlockDeviceMappings.",
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

					"context": schema.StringAttribute{
						Description:         "Context is a Reserved field in EC2 APIshttps://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet.html",
						MarkdownDescription: "Context is a Reserved field in EC2 APIshttps://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateFleet.html",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"detailed_monitoring": schema.BoolAttribute{
						Description:         "DetailedMonitoring controls if detailed monitoring is enabled for instances that are launched",
						MarkdownDescription: "DetailedMonitoring controls if detailed monitoring is enabled for instances that are launched",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance_profile": schema.StringAttribute{
						Description:         "InstanceProfile is the AWS entity that instances use.This field is mutually exclusive from role.The instance profile should already have a role assigned to it that Karpenter has PassRole permission on for instance launch using this instanceProfile to succeed.",
						MarkdownDescription: "InstanceProfile is the AWS entity that instances use.This field is mutually exclusive from role.The instance profile should already have a role assigned to it that Karpenter has PassRole permission on for instance launch using this instanceProfile to succeed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"instance_store_policy": schema.StringAttribute{
						Description:         "InstanceStorePolicy specifies how to handle instance-store disks.",
						MarkdownDescription: "InstanceStorePolicy specifies how to handle instance-store disks.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("RAID0"),
						},
					},

					"metadata_options": schema.SingleNestedAttribute{
						Description:         "MetadataOptions for the generated launch template of provisioned nodes.This specifies the exposure of the Instance Metadata Service toprovisioned EC2 nodes. For more information,see Instance Metadata and User Data(https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)in the Amazon Elastic Compute Cloud User Guide.Refer to recommended, security best practices(https://aws.github.io/aws-eks-best-practices/security/docs/iam/#restrict-access-to-the-instance-profile-assigned-to-the-worker-node)for limiting exposure of Instance Metadata and User Data to pods.If omitted, defaults to httpEndpoint enabled, with httpProtocolIPv6disabled, with httpPutResponseLimit of 2, and with httpTokensrequired.",
						MarkdownDescription: "MetadataOptions for the generated launch template of provisioned nodes.This specifies the exposure of the Instance Metadata Service toprovisioned EC2 nodes. For more information,see Instance Metadata and User Data(https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html)in the Amazon Elastic Compute Cloud User Guide.Refer to recommended, security best practices(https://aws.github.io/aws-eks-best-practices/security/docs/iam/#restrict-access-to-the-instance-profile-assigned-to-the-worker-node)for limiting exposure of Instance Metadata and User Data to pods.If omitted, defaults to httpEndpoint enabled, with httpProtocolIPv6disabled, with httpPutResponseLimit of 2, and with httpTokensrequired.",
						Attributes: map[string]schema.Attribute{
							"http_endpoint": schema.StringAttribute{
								Description:         "HTTPEndpoint enables or disables the HTTP metadata endpoint on provisionednodes. If metadata options is non-nil, but this parameter is not specified,the default state is 'enabled'.If you specify a value of 'disabled', instance metadata will not be accessibleon the node.",
								MarkdownDescription: "HTTPEndpoint enables or disables the HTTP metadata endpoint on provisionednodes. If metadata options is non-nil, but this parameter is not specified,the default state is 'enabled'.If you specify a value of 'disabled', instance metadata will not be accessibleon the node.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("enabled", "disabled"),
								},
							},

							"http_protocol_i_pv6": schema.StringAttribute{
								Description:         "HTTPProtocolIPv6 enables or disables the IPv6 endpoint for the instance metadataservice on provisioned nodes. If metadata options is non-nil, but this parameteris not specified, the default state is 'disabled'.",
								MarkdownDescription: "HTTPProtocolIPv6 enables or disables the IPv6 endpoint for the instance metadataservice on provisioned nodes. If metadata options is non-nil, but this parameteris not specified, the default state is 'disabled'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("enabled", "disabled"),
								},
							},

							"http_put_response_hop_limit": schema.Int64Attribute{
								Description:         "HTTPPutResponseHopLimit is the desired HTTP PUT response hop limit forinstance metadata requests. The larger the number, the further instancemetadata requests can travel. Possible values are integers from 1 to 64.If metadata options is non-nil, but this parameter is not specified, thedefault value is 2.",
								MarkdownDescription: "HTTPPutResponseHopLimit is the desired HTTP PUT response hop limit forinstance metadata requests. The larger the number, the further instancemetadata requests can travel. Possible values are integers from 1 to 64.If metadata options is non-nil, but this parameter is not specified, thedefault value is 2.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(64),
								},
							},

							"http_tokens": schema.StringAttribute{
								Description:         "HTTPTokens determines the state of token usage for instance metadatarequests. If metadata options is non-nil, but this parameter is notspecified, the default state is 'required'.If the state is optional, one can choose to retrieve instance metadata withor without a signed token header on the request. If one retrieves the IAMrole credentials without a token, the version 1.0 role credentials arereturned. If one retrieves the IAM role credentials using a valid signedtoken, the version 2.0 role credentials are returned.If the state is 'required', one must send a signed token header with anyinstance metadata retrieval requests. In this state, retrieving the IAMrole credentials always returns the version 2.0 credentials; the version1.0 credentials are not available.",
								MarkdownDescription: "HTTPTokens determines the state of token usage for instance metadatarequests. If metadata options is non-nil, but this parameter is notspecified, the default state is 'required'.If the state is optional, one can choose to retrieve instance metadata withor without a signed token header on the request. If one retrieves the IAMrole credentials without a token, the version 1.0 role credentials arereturned. If one retrieves the IAM role credentials using a valid signedtoken, the version 2.0 role credentials are returned.If the state is 'required', one must send a signed token header with anyinstance metadata retrieval requests. In this state, retrieving the IAMrole credentials always returns the version 2.0 credentials; the version1.0 credentials are not available.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("required", "optional"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"role": schema.StringAttribute{
						Description:         "Role is the AWS identity that nodes use. This field is immutable.This field is mutually exclusive from instanceProfile.Marking this field as immutable avoids concerns around terminating managed instance profiles from running instances.This field may be made mutable in the future, assuming the correct garbage collection and drift handling is implementedfor the old instance profiles on an update.",
						MarkdownDescription: "Role is the AWS identity that nodes use. This field is immutable.This field is mutually exclusive from instanceProfile.Marking this field as immutable avoids concerns around terminating managed instance profiles from running instances.This field may be made mutable in the future, assuming the correct garbage collection and drift handling is implementedfor the old instance profiles on an update.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"security_group_selector_terms": schema.ListNestedAttribute{
						Description:         "SecurityGroupSelectorTerms is a list of or security group selector terms. The terms are ORed.",
						MarkdownDescription: "SecurityGroupSelectorTerms is a list of or security group selector terms. The terms are ORed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "ID is the security group id in EC2",
									MarkdownDescription: "ID is the security group id in EC2",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`sg-[0-9a-z]+`), ""),
									},
								},

								"name": schema.StringAttribute{
									Description:         "Name is the security group name in EC2.This value is the name field, which is different from the name tag.",
									MarkdownDescription: "Name is the security group name in EC2.This value is the name field, which is different from the name tag.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tags": schema.MapAttribute{
									Description:         "Tags is a map of key/value tags used to select subnetsSpecifying '*' for a value selects all values for a given tag key.",
									MarkdownDescription: "Tags is a map of key/value tags used to select subnetsSpecifying '*' for a value selects all values for a given tag key.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"subnet_selector_terms": schema.ListNestedAttribute{
						Description:         "SubnetSelectorTerms is a list of or subnet selector terms. The terms are ORed.",
						MarkdownDescription: "SubnetSelectorTerms is a list of or subnet selector terms. The terms are ORed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "ID is the subnet id in EC2",
									MarkdownDescription: "ID is the subnet id in EC2",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.RegexMatches(regexp.MustCompile(`subnet-[0-9a-z]+`), ""),
									},
								},

								"tags": schema.MapAttribute{
									Description:         "Tags is a map of key/value tags used to select subnetsSpecifying '*' for a value selects all values for a given tag key.",
									MarkdownDescription: "Tags is a map of key/value tags used to select subnetsSpecifying '*' for a value selects all values for a given tag key.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"tags": schema.MapAttribute{
						Description:         "Tags to be applied on ec2 resources like instances and launch templates.",
						MarkdownDescription: "Tags to be applied on ec2 resources like instances and launch templates.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"user_data": schema.StringAttribute{
						Description:         "UserData to be applied to the provisioned nodes.It must be in the appropriate format based on the AMIFamily in use. Karpenter will merge certain fields intothis UserData to ensure nodes are being provisioned with the correct configuration.",
						MarkdownDescription: "UserData to be applied to the provisioned nodes.It must be in the appropriate format based on the AMIFamily in use. Karpenter will merge certain fields intothis UserData to ensure nodes are being provisioned with the correct configuration.",
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

func (r *KarpenterK8SAwsEc2NodeClassV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_karpenter_k8s_aws_ec2_node_class_v1beta1_manifest")

	var model KarpenterK8SAwsEc2NodeClassV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("karpenter.k8s.aws/v1beta1")
	model.Kind = pointer.String("EC2NodeClass")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
