/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package azure_microsoft_com_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AzureMicrosoftComStorageAccountV1Alpha1Manifest{}
)

func NewAzureMicrosoftComStorageAccountV1Alpha1Manifest() datasource.DataSource {
	return &AzureMicrosoftComStorageAccountV1Alpha1Manifest{}
}

type AzureMicrosoftComStorageAccountV1Alpha1Manifest struct{}

type AzureMicrosoftComStorageAccountV1Alpha1ManifestData struct {
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

	AdditionalResources *struct {
		Secrets *[]string `tfsdk:"secrets" json:"secrets,omitempty"`
	} `tfsdk:"additional_resources" json:"additionalResources,omitempty"`
	Output *struct {
		ConnectionString1  *string `tfsdk:"connection_string1" json:"connectionString1,omitempty"`
		ConnectionString2  *string `tfsdk:"connection_string2" json:"connectionString2,omitempty"`
		Key1               *string `tfsdk:"key1" json:"key1,omitempty"`
		Key2               *string `tfsdk:"key2" json:"key2,omitempty"`
		StorageAccountName *string `tfsdk:"storage_account_name" json:"storageAccountName,omitempty"`
	} `tfsdk:"output" json:"output,omitempty"`
	Spec *struct {
		AccessTier      *string `tfsdk:"access_tier" json:"accessTier,omitempty"`
		DataLakeEnabled *bool   `tfsdk:"data_lake_enabled" json:"dataLakeEnabled,omitempty"`
		Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
		Location        *string `tfsdk:"location" json:"location,omitempty"`
		NetworkRule     *struct {
			Bypass        *string `tfsdk:"bypass" json:"bypass,omitempty"`
			DefaultAction *string `tfsdk:"default_action" json:"defaultAction,omitempty"`
			IpRules       *[]struct {
				IpAddressOrRange *string `tfsdk:"ip_address_or_range" json:"ipAddressOrRange,omitempty"`
			} `tfsdk:"ip_rules" json:"ipRules,omitempty"`
			VirtualNetworkRules *[]struct {
				SubnetId *string `tfsdk:"subnet_id" json:"subnetId,omitempty"`
			} `tfsdk:"virtual_network_rules" json:"virtualNetworkRules,omitempty"`
		} `tfsdk:"network_rule" json:"networkRule,omitempty"`
		ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		Sku           *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"sku" json:"sku,omitempty"`
		SupportsHttpsTrafficOnly *bool `tfsdk:"supports_https_traffic_only" json:"supportsHttpsTrafficOnly,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComStorageAccountV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_storage_account_v1alpha1_manifest"
}

func (r *AzureMicrosoftComStorageAccountV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "StorageAccount is the Schema for the storages API",
		MarkdownDescription: "StorageAccount is the Schema for the storages API",
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

			"additional_resources": schema.SingleNestedAttribute{
				Description:         "StorageAccountAdditionalResources holds the additional resources",
				MarkdownDescription: "StorageAccountAdditionalResources holds the additional resources",
				Attributes: map[string]schema.Attribute{
					"secrets": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
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

			"output": schema.SingleNestedAttribute{
				Description:         "StorageAccountOutput is the object that contains the output from creating a Storage Account object",
				MarkdownDescription: "StorageAccountOutput is the object that contains the output from creating a Storage Account object",
				Attributes: map[string]schema.Attribute{
					"connection_string1": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"connection_string2": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key1": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key2": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_account_name": schema.StringAttribute{
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

			"spec": schema.SingleNestedAttribute{
				Description:         "StorageAccountSpec defines the desired state of Storage",
				MarkdownDescription: "StorageAccountSpec defines the desired state of Storage",
				Attributes: map[string]schema.Attribute{
					"access_tier": schema.StringAttribute{
						Description:         "StorageAccountAccessTier enumerates the values for access tier. Only one of the following access tiers may be specified. If none of the following access tiers is specified, the default one is Hot.",
						MarkdownDescription: "StorageAccountAccessTier enumerates the values for access tier. Only one of the following access tiers may be specified. If none of the following access tiers is specified, the default one is Hot.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Cool", "Hot"),
						},
					},

					"data_lake_enabled": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kind": schema.StringAttribute{
						Description:         "StorageAccountKind enumerates the values for kind. Only one of the following kinds may be specified. If none of the following kinds is specified, the default one is StorageV2.",
						MarkdownDescription: "StorageAccountKind enumerates the values for kind. Only one of the following kinds may be specified. If none of the following kinds is specified, the default one is StorageV2.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("BlobStorage", "BlockBlobStorage", "FileStorage", "Storage", "StorageV2"),
						},
					},

					"location": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(0),
						},
					},

					"network_rule": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"bypass": schema.StringAttribute{
								Description:         "Bypass - Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices (For example, 'Logging, Metrics'), or None to bypass none of those traffics. Possible values include: 'None', 'Logging', 'Metrics', 'AzureServices'",
								MarkdownDescription: "Bypass - Specifies whether traffic is bypassed for Logging/Metrics/AzureServices. Possible values are any combination of Logging|Metrics|AzureServices (For example, 'Logging, Metrics'), or None to bypass none of those traffics. Possible values include: 'None', 'Logging', 'Metrics', 'AzureServices'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"default_action": schema.StringAttribute{
								Description:         "DefaultAction - Specifies the default action of allow or deny when no other rules match. Possible values include: 'DefaultActionAllow', 'DefaultActionDeny'",
								MarkdownDescription: "DefaultAction - Specifies the default action of allow or deny when no other rules match. Possible values include: 'DefaultActionAllow', 'DefaultActionDeny'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ip_rules": schema.ListNestedAttribute{
								Description:         "IPRules - Sets the IP ACL rules",
								MarkdownDescription: "IPRules - Sets the IP ACL rules",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip_address_or_range": schema.StringAttribute{
											Description:         "IPAddressOrRange - Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.",
											MarkdownDescription: "IPAddressOrRange - Specifies the IP or IP range in CIDR format. Only IPV4 address is allowed.",
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

							"virtual_network_rules": schema.ListNestedAttribute{
								Description:         "VirtualNetworkRules - Sets the virtual network rules",
								MarkdownDescription: "VirtualNetworkRules - Sets the virtual network rules",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"subnet_id": schema.StringAttribute{
											Description:         "SubnetId - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.",
											MarkdownDescription: "SubnetId - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{vnetName}/subnets/{subnetName}.",
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

					"sku": schema.SingleNestedAttribute{
						Description:         "StorageAccountSku the SKU of the storage account.",
						MarkdownDescription: "StorageAccountSku the SKU of the storage account.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name - The SKU name. Required for account creation; optional for update. Possible values include: 'Standard_LRS', 'Standard_GRS', 'Standard_RAGRS', 'Standard_ZRS', 'Premium_LRS', 'Premium_ZRS', 'Standard_GZRS', 'Standard_RAGZRS'. For the full list of allowed options, see: https://docs.microsoft.com/en-us/rest/api/storagerp/storageaccounts/create#skuname",
								MarkdownDescription: "Name - The SKU name. Required for account creation; optional for update. Possible values include: 'Standard_LRS', 'Standard_GRS', 'Standard_RAGRS', 'Standard_ZRS', 'Premium_LRS', 'Premium_ZRS', 'Standard_GZRS', 'Standard_RAGZRS'. For the full list of allowed options, see: https://docs.microsoft.com/en-us/rest/api/storagerp/storageaccounts/create#skuname",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Premium_LRS", "Premium_ZRS", "Standard_GRS", "Standard_GZRS", "Standard_LRS", "Standard_RAGRS", "Standard_RAGZRS", "Standard_ZRS"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"supports_https_traffic_only": schema.BoolAttribute{
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
	}
}

func (r *AzureMicrosoftComStorageAccountV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_storage_account_v1alpha1_manifest")

	var model AzureMicrosoftComStorageAccountV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("azure.microsoft.com/v1alpha1")
	model.Kind = pointer.String("StorageAccount")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
