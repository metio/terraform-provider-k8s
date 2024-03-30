/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package azure_microsoft_com_v1alpha1

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
	_ datasource.DataSource = &AzureMicrosoftComEventhubV1Alpha1Manifest{}
)

func NewAzureMicrosoftComEventhubV1Alpha1Manifest() datasource.DataSource {
	return &AzureMicrosoftComEventhubV1Alpha1Manifest{}
}

type AzureMicrosoftComEventhubV1Alpha1Manifest struct{}

type AzureMicrosoftComEventhubV1Alpha1ManifestData struct {
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
		AuthorizationRule *struct {
			Name   *string   `tfsdk:"name" json:"name,omitempty"`
			Rights *[]string `tfsdk:"rights" json:"rights,omitempty"`
		} `tfsdk:"authorization_rule" json:"authorizationRule,omitempty"`
		KeyVaultToStoreSecrets *string `tfsdk:"key_vault_to_store_secrets" json:"keyVaultToStoreSecrets,omitempty"`
		Location               *string `tfsdk:"location" json:"location,omitempty"`
		Namespace              *string `tfsdk:"namespace" json:"namespace,omitempty"`
		Properties             *struct {
			CaptureDescription *struct {
				Destination *struct {
					ArchiveNameFormat *string `tfsdk:"archive_name_format" json:"archiveNameFormat,omitempty"`
					BlobContainer     *string `tfsdk:"blob_container" json:"blobContainer,omitempty"`
					Name              *string `tfsdk:"name" json:"name,omitempty"`
					StorageAccount    *struct {
						AccountName   *string `tfsdk:"account_name" json:"accountName,omitempty"`
						ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
					} `tfsdk:"storage_account" json:"storageAccount,omitempty"`
				} `tfsdk:"destination" json:"destination,omitempty"`
				Enabled           *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
				IntervalInSeconds *int64 `tfsdk:"interval_in_seconds" json:"intervalInSeconds,omitempty"`
				SizeLimitInBytes  *int64 `tfsdk:"size_limit_in_bytes" json:"sizeLimitInBytes,omitempty"`
			} `tfsdk:"capture_description" json:"captureDescription,omitempty"`
			MessageRetentionInDays *int64 `tfsdk:"message_retention_in_days" json:"messageRetentionInDays,omitempty"`
			PartitionCount         *int64 `tfsdk:"partition_count" json:"partitionCount,omitempty"`
		} `tfsdk:"properties" json:"properties,omitempty"`
		ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		SecretName    *string `tfsdk:"secret_name" json:"secretName,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComEventhubV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_eventhub_v1alpha1_manifest"
}

func (r *AzureMicrosoftComEventhubV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Eventhub is the Schema for the eventhubs API",
		MarkdownDescription: "Eventhub is the Schema for the eventhubs API",
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
				Description:         "EventhubSpec defines the desired state of Eventhub",
				MarkdownDescription: "EventhubSpec defines the desired state of Eventhub",
				Attributes: map[string]schema.Attribute{
					"authorization_rule": schema.SingleNestedAttribute{
						Description:         "EventhubAuthorizationRule defines the name and rights of the access policy",
						MarkdownDescription: "EventhubAuthorizationRule defines the name and rights of the access policy",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name - Name of AuthorizationRule for eventhub",
								MarkdownDescription: "Name - Name of AuthorizationRule for eventhub",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rights": schema.ListAttribute{
								Description:         "Rights - Rights set on the AuthorizationRule",
								MarkdownDescription: "Rights - Rights set on the AuthorizationRule",
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

					"key_vault_to_store_secrets": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"location": schema.StringAttribute{
						Description:         "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'make' to regenerate code after modifying this file",
						MarkdownDescription: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run 'make' to regenerate code after modifying this file",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"namespace": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"properties": schema.SingleNestedAttribute{
						Description:         "EventhubProperties defines the namespace properties",
						MarkdownDescription: "EventhubProperties defines the namespace properties",
						Attributes: map[string]schema.Attribute{
							"capture_description": schema.SingleNestedAttribute{
								Description:         "CaptureDescription - Details specifying EventHub capture to persistent storage",
								MarkdownDescription: "CaptureDescription - Details specifying EventHub capture to persistent storage",
								Attributes: map[string]schema.Attribute{
									"destination": schema.SingleNestedAttribute{
										Description:         "Destination - Resource id of the storage account to be used to create the blobs",
										MarkdownDescription: "Destination - Resource id of the storage account to be used to create the blobs",
										Attributes: map[string]schema.Attribute{
											"archive_name_format": schema.StringAttribute{
												Description:         "ArchiveNameFormat - Blob naming convention for archive, e.g. {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order",
												MarkdownDescription: "ArchiveNameFormat - Blob naming convention for archive, e.g. {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters (Namespace,EventHub .. etc) are mandatory irrespective of order",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blob_container": schema.StringAttribute{
												Description:         "BlobContainer - Blob container Name",
												MarkdownDescription: "BlobContainer - Blob container Name",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name - Name for capture destination",
												MarkdownDescription: "Name - Name for capture destination",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("EventHubArchive.AzureBlockBlob", "EventHubArchive.AzureDataLake"),
												},
											},

											"storage_account": schema.SingleNestedAttribute{
												Description:         "StorageAccount - Details of the storage account",
												MarkdownDescription: "StorageAccount - Details of the storage account",
												Attributes: map[string]schema.Attribute{
													"account_name": schema.StringAttribute{
														Description:         "AccountName - Name of the storage account",
														MarkdownDescription: "AccountName - Name of the storage account",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(3),
															stringvalidator.LengthAtMost(24),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]+$`), ""),
														},
													},

													"resource_group": schema.StringAttribute{
														Description:         "ResourceGroup - Name of the storage account resource group",
														MarkdownDescription: "ResourceGroup - Name of the storage account resource group",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[-\w\._\(\)]+$`), ""),
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

									"enabled": schema.BoolAttribute{
										Description:         "Enabled - indicates whether capture is enabled",
										MarkdownDescription: "Enabled - indicates whether capture is enabled",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"interval_in_seconds": schema.Int64Attribute{
										Description:         "IntervalInSeconds - The time window allows you to set the frequency with which the capture to Azure Blobs will happen",
										MarkdownDescription: "IntervalInSeconds - The time window allows you to set the frequency with which the capture to Azure Blobs will happen",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(60),
											int64validator.AtMost(900),
										},
									},

									"size_limit_in_bytes": schema.Int64Attribute{
										Description:         "SizeLimitInBytes - The size window defines the amount of data built up in your Event Hub before an capture operation",
										MarkdownDescription: "SizeLimitInBytes - The size window defines the amount of data built up in your Event Hub before an capture operation",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1.048576e+07),
											int64validator.AtMost(5.24288e+08),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"message_retention_in_days": schema.Int64Attribute{
								Description:         "MessageRetentionInDays - Number of days to retain the events for this Event Hub, value should be 1 to 7 days",
								MarkdownDescription: "MessageRetentionInDays - Number of days to retain the events for this Event Hub, value should be 1 to 7 days",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
									int64validator.AtMost(7),
								},
							},

							"partition_count": schema.Int64Attribute{
								Description:         "PartitionCount - Number of partitions created for the Event Hub, allowed values are from 2 to 32 partitions.",
								MarkdownDescription: "PartitionCount - Number of partitions created for the Event Hub, allowed values are from 2 to 32 partitions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(2),
									int64validator.AtMost(32),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_group": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[-\w\._\(\)]+$`), ""),
						},
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName - Used to specify the name of the secret. Defaults to Event Hub name if omitted.",
						MarkdownDescription: "SecretName - Used to specify the name of the secret. Defaults to Event Hub name if omitted.",
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

func (r *AzureMicrosoftComEventhubV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_eventhub_v1alpha1_manifest")

	var model AzureMicrosoftComEventhubV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("azure.microsoft.com/v1alpha1")
	model.Kind = pointer.String("Eventhub")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
