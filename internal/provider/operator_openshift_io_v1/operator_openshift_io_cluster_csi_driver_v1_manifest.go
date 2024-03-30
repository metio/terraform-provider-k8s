/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_openshift_io_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &OperatorOpenshiftIoClusterCsidriverV1Manifest{}
)

func NewOperatorOpenshiftIoClusterCsidriverV1Manifest() datasource.DataSource {
	return &OperatorOpenshiftIoClusterCsidriverV1Manifest{}
}

type OperatorOpenshiftIoClusterCsidriverV1Manifest struct{}

type OperatorOpenshiftIoClusterCsidriverV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		DriverConfig *struct {
			Aws *struct {
				KmsKeyARN *string `tfsdk:"kms_key_arn" json:"kmsKeyARN,omitempty"`
			} `tfsdk:"aws" json:"aws,omitempty"`
			Azure *struct {
				DiskEncryptionSet *struct {
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					ResourceGroup  *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
					SubscriptionID *string `tfsdk:"subscription_id" json:"subscriptionID,omitempty"`
				} `tfsdk:"disk_encryption_set" json:"diskEncryptionSet,omitempty"`
			} `tfsdk:"azure" json:"azure,omitempty"`
			DriverType *string `tfsdk:"driver_type" json:"driverType,omitempty"`
			Gcp        *struct {
				KmsKey *struct {
					KeyRing   *string `tfsdk:"key_ring" json:"keyRing,omitempty"`
					Location  *string `tfsdk:"location" json:"location,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					ProjectID *string `tfsdk:"project_id" json:"projectID,omitempty"`
				} `tfsdk:"kms_key" json:"kmsKey,omitempty"`
			} `tfsdk:"gcp" json:"gcp,omitempty"`
			Ibmcloud *struct {
				EncryptionKeyCRN *string `tfsdk:"encryption_key_crn" json:"encryptionKeyCRN,omitempty"`
			} `tfsdk:"ibmcloud" json:"ibmcloud,omitempty"`
			VSphere *struct {
				TopologyCategories *[]string `tfsdk:"topology_categories" json:"topologyCategories,omitempty"`
			} `tfsdk:"v_sphere" json:"vSphere,omitempty"`
		} `tfsdk:"driver_config" json:"driverConfig,omitempty"`
		LogLevel                   *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
		ManagementState            *string            `tfsdk:"management_state" json:"managementState,omitempty"`
		ObservedConfig             *map[string]string `tfsdk:"observed_config" json:"observedConfig,omitempty"`
		OperatorLogLevel           *string            `tfsdk:"operator_log_level" json:"operatorLogLevel,omitempty"`
		StorageClassState          *string            `tfsdk:"storage_class_state" json:"storageClassState,omitempty"`
		UnsupportedConfigOverrides *map[string]string `tfsdk:"unsupported_config_overrides" json:"unsupportedConfigOverrides,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorOpenshiftIoClusterCsidriverV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_openshift_io_cluster_csi_driver_v1_manifest"
}

func (r *OperatorOpenshiftIoClusterCsidriverV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ClusterCSIDriver object allows management and configuration of a CSI driver operator installed by default in OpenShift. Name of the object must be name of the CSI driver it operates. See CSIDriverName type for list of allowed values.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
		MarkdownDescription: "ClusterCSIDriver object allows management and configuration of a CSI driver operator installed by default in OpenShift. Name of the object must be name of the CSI driver it operates. See CSIDriverName type for list of allowed values.  Compatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
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
				Description:         "spec holds user settable values for configuration",
				MarkdownDescription: "spec holds user settable values for configuration",
				Attributes: map[string]schema.Attribute{
					"driver_config": schema.SingleNestedAttribute{
						Description:         "driverConfig can be used to specify platform specific driver configuration. When omitted, this means no opinion and the platform is left to choose reasonable defaults. These defaults are subject to change over time.",
						MarkdownDescription: "driverConfig can be used to specify platform specific driver configuration. When omitted, this means no opinion and the platform is left to choose reasonable defaults. These defaults are subject to change over time.",
						Attributes: map[string]schema.Attribute{
							"aws": schema.SingleNestedAttribute{
								Description:         "aws is used to configure the AWS CSI driver.",
								MarkdownDescription: "aws is used to configure the AWS CSI driver.",
								Attributes: map[string]schema.Attribute{
									"kms_key_arn": schema.StringAttribute{
										Description:         "kmsKeyARN sets the cluster default storage class to encrypt volumes with a user-defined KMS key, rather than the default KMS key used by AWS. The value may be either the ARN or Alias ARN of a KMS key.",
										MarkdownDescription: "kmsKeyARN sets the cluster default storage class to encrypt volumes with a user-defined KMS key, rather than the default KMS key used by AWS. The value may be either the ARN or Alias ARN of a KMS key.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b|aws-iso-e|aws-iso-f):kms:[a-z0-9-]+:[0-9]{12}:(key|alias)\/.*$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"azure": schema.SingleNestedAttribute{
								Description:         "azure is used to configure the Azure CSI driver.",
								MarkdownDescription: "azure is used to configure the Azure CSI driver.",
								Attributes: map[string]schema.Attribute{
									"disk_encryption_set": schema.SingleNestedAttribute{
										Description:         "diskEncryptionSet sets the cluster default storage class to encrypt volumes with a customer-managed encryption set, rather than the default platform-managed keys.",
										MarkdownDescription: "diskEncryptionSet sets the cluster default storage class to encrypt volumes with a customer-managed encryption set, rather than the default platform-managed keys.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "name is the name of the disk encryption set that will be set on the default storage class. The value should consist of only alphanumberic characters, underscores (_), hyphens, and be at most 80 characters in length.",
												MarkdownDescription: "name is the name of the disk encryption set that will be set on the default storage class. The value should consist of only alphanumberic characters, underscores (_), hyphens, and be at most 80 characters in length.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(80),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9\_-]+$`), ""),
												},
											},

											"resource_group": schema.StringAttribute{
												Description:         "resourceGroup defines the Azure resource group that contains the disk encryption set. The value should consist of only alphanumberic characters, underscores (_), parentheses, hyphens and periods. The value should not end in a period and be at most 90 characters in length.",
												MarkdownDescription: "resourceGroup defines the Azure resource group that contains the disk encryption set. The value should consist of only alphanumberic characters, underscores (_), parentheses, hyphens and periods. The value should not end in a period and be at most 90 characters in length.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(90),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[\w\.\-\(\)]*[\w\-\(\)]$`), ""),
												},
											},

											"subscription_id": schema.StringAttribute{
												Description:         "subscriptionID defines the Azure subscription that contains the disk encryption set. The value should meet the following conditions: 1. It should be a 128-bit number. 2. It should be 36 characters (32 hexadecimal characters and 4 hyphens) long. 3. It should be displayed in five groups separated by hyphens (-). 4. The first group should be 8 characters long. 5. The second, third, and fourth groups should be 4 characters long. 6. The fifth group should be 12 characters long. An Example SubscrionID: f2007bbf-f802-4a47-9336-cf7c6b89b378",
												MarkdownDescription: "subscriptionID defines the Azure subscription that contains the disk encryption set. The value should meet the following conditions: 1. It should be a 128-bit number. 2. It should be 36 characters (32 hexadecimal characters and 4 hyphens) long. 3. It should be displayed in five groups separated by hyphens (-). 4. The first group should be 8 characters long. 5. The second, third, and fourth groups should be 4 characters long. 6. The fifth group should be 12 characters long. An Example SubscrionID: f2007bbf-f802-4a47-9336-cf7c6b89b378",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(36),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`), ""),
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

							"driver_type": schema.StringAttribute{
								Description:         "driverType indicates type of CSI driver for which the driverConfig is being applied to. Valid values are: AWS, Azure, GCP, IBMCloud, vSphere and omitted. Consumers should treat unknown values as a NO-OP.",
								MarkdownDescription: "driverType indicates type of CSI driver for which the driverConfig is being applied to. Valid values are: AWS, Azure, GCP, IBMCloud, vSphere and omitted. Consumers should treat unknown values as a NO-OP.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "AWS", "Azure", "GCP", "IBMCloud", "vSphere"),
								},
							},

							"gcp": schema.SingleNestedAttribute{
								Description:         "gcp is used to configure the GCP CSI driver.",
								MarkdownDescription: "gcp is used to configure the GCP CSI driver.",
								Attributes: map[string]schema.Attribute{
									"kms_key": schema.SingleNestedAttribute{
										Description:         "kmsKey sets the cluster default storage class to encrypt volumes with customer-supplied encryption keys, rather than the default keys managed by GCP.",
										MarkdownDescription: "kmsKey sets the cluster default storage class to encrypt volumes with customer-supplied encryption keys, rather than the default keys managed by GCP.",
										Attributes: map[string]schema.Attribute{
											"key_ring": schema.StringAttribute{
												Description:         "keyRing is the name of the KMS Key Ring which the KMS Key belongs to. The value should correspond to an existing KMS key ring and should consist of only alphanumeric characters, hyphens (-) and underscores (_), and be at most 63 characters in length.",
												MarkdownDescription: "keyRing is the name of the KMS Key Ring which the KMS Key belongs to. The value should correspond to an existing KMS key ring and should consist of only alphanumeric characters, hyphens (-) and underscores (_), and be at most 63 characters in length.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9\_-]+$`), ""),
												},
											},

											"location": schema.StringAttribute{
												Description:         "location is the GCP location in which the Key Ring exists. The value must match an existing GCP location, or 'global'. Defaults to global, if not set.",
												MarkdownDescription: "location is the GCP location in which the Key Ring exists. The value must match an existing GCP location, or 'global'. Defaults to global, if not set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9\_-]+$`), ""),
												},
											},

											"name": schema.StringAttribute{
												Description:         "name is the name of the customer-managed encryption key to be used for disk encryption. The value should correspond to an existing KMS key and should consist of only alphanumeric characters, hyphens (-) and underscores (_), and be at most 63 characters in length.",
												MarkdownDescription: "name is the name of the customer-managed encryption key to be used for disk encryption. The value should correspond to an existing KMS key and should consist of only alphanumeric characters, hyphens (-) and underscores (_), and be at most 63 characters in length.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9\_-]+$`), ""),
												},
											},

											"project_id": schema.StringAttribute{
												Description:         "projectID is the ID of the Project in which the KMS Key Ring exists. It must be 6 to 30 lowercase letters, digits, or hyphens. It must start with a letter. Trailing hyphens are prohibited.",
												MarkdownDescription: "projectID is the ID of the Project in which the KMS Key Ring exists. It must be 6 to 30 lowercase letters, digits, or hyphens. It must start with a letter. Trailing hyphens are prohibited.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(6),
													stringvalidator.LengthAtMost(30),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z][a-z0-9-]+[a-z0-9]$`), ""),
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

							"ibmcloud": schema.SingleNestedAttribute{
								Description:         "ibmcloud is used to configure the IBM Cloud CSI driver.",
								MarkdownDescription: "ibmcloud is used to configure the IBM Cloud CSI driver.",
								Attributes: map[string]schema.Attribute{
									"encryption_key_crn": schema.StringAttribute{
										Description:         "encryptionKeyCRN is the IBM Cloud CRN of the customer-managed root key to use for disk encryption of volumes for the default storage classes.",
										MarkdownDescription: "encryptionKeyCRN is the IBM Cloud CRN of the customer-managed root key to use for disk encryption of volumes for the default storage classes.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(144),
											stringvalidator.LengthAtMost(154),
											stringvalidator.RegexMatches(regexp.MustCompile(`^crn:v[0-9]+:bluemix:(public|private):(kms|hs-crypto):[a-z-]+:a/[0-9a-f]+:[0-9a-f-]{36}:key:[0-9a-f-]{36}$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"v_sphere": schema.SingleNestedAttribute{
								Description:         "vsphere is used to configure the vsphere CSI driver.",
								MarkdownDescription: "vsphere is used to configure the vsphere CSI driver.",
								Attributes: map[string]schema.Attribute{
									"topology_categories": schema.ListAttribute{
										Description:         "topologyCategories indicates tag categories with which vcenter resources such as hostcluster or datacenter were tagged with. If cluster Infrastructure object has a topology, values specified in Infrastructure object will be used and modifications to topologyCategories will be rejected.",
										MarkdownDescription: "topologyCategories indicates tag categories with which vcenter resources such as hostcluster or datacenter were tagged with. If cluster Infrastructure object has a topology, values specified in Infrastructure object will be used and modifications to topologyCategories will be rejected.",
										ElementType:         types.StringType,
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

					"log_level": schema.StringAttribute{
						Description:         "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						MarkdownDescription: "logLevel is an intent based logging for an overall component.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for their operands.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Normal", "Debug", "Trace", "TraceAll"),
						},
					},

					"management_state": schema.StringAttribute{
						Description:         "managementState indicates whether and how the operator should manage the component",
						MarkdownDescription: "managementState indicates whether and how the operator should manage the component",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(Managed|Unmanaged|Force|Removed)$`), ""),
						},
					},

					"observed_config": schema.MapAttribute{
						Description:         "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
						MarkdownDescription: "observedConfig holds a sparse config that controller has observed from the cluster state.  It exists in spec because it is an input to the level for the operator",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operator_log_level": schema.StringAttribute{
						Description:         "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						MarkdownDescription: "operatorLogLevel is an intent based logging for the operator itself.  It does not give fine grained control, but it is a simple way to manage coarse grained logging choices that operators have to interpret for themselves.  Valid values are: 'Normal', 'Debug', 'Trace', 'TraceAll'. Defaults to 'Normal'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Normal", "Debug", "Trace", "TraceAll"),
						},
					},

					"storage_class_state": schema.StringAttribute{
						Description:         "StorageClassState determines if CSI operator should create and manage storage classes. If this field value is empty or Managed - CSI operator will continuously reconcile storage class and create if necessary. If this field value is Unmanaged - CSI operator will not reconcile any previously created storage class. If this field value is Removed - CSI operator will delete the storage class it created previously. When omitted, this means the user has no opinion and the platform chooses a reasonable default, which is subject to change over time. The current default behaviour is Managed.",
						MarkdownDescription: "StorageClassState determines if CSI operator should create and manage storage classes. If this field value is empty or Managed - CSI operator will continuously reconcile storage class and create if necessary. If this field value is Unmanaged - CSI operator will not reconcile any previously created storage class. If this field value is Removed - CSI operator will delete the storage class it created previously. When omitted, this means the user has no opinion and the platform chooses a reasonable default, which is subject to change over time. The current default behaviour is Managed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("", "Managed", "Unmanaged", "Removed"),
						},
					},

					"unsupported_config_overrides": schema.MapAttribute{
						Description:         "unsupportedConfigOverrides overrides the final configuration that was computed by the operator. Red Hat does not support the use of this field. Misuse of this field could lead to unexpected behavior or conflict with other configuration options. Seek guidance from the Red Hat support before using this field. Use of this property blocks cluster upgrades, it must be removed before upgrading your cluster.",
						MarkdownDescription: "unsupportedConfigOverrides overrides the final configuration that was computed by the operator. Red Hat does not support the use of this field. Misuse of this field could lead to unexpected behavior or conflict with other configuration options. Seek guidance from the Red Hat support before using this field. Use of this property blocks cluster upgrades, it must be removed before upgrading your cluster.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *OperatorOpenshiftIoClusterCsidriverV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_openshift_io_cluster_csi_driver_v1_manifest")

	var model OperatorOpenshiftIoClusterCsidriverV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("operator.openshift.io/v1")
	model.Kind = pointer.String("ClusterCSIDriver")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
