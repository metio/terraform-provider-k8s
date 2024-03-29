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
	_ datasource.DataSource = &AzureMicrosoftComCosmosDbV1Alpha1Manifest{}
)

func NewAzureMicrosoftComCosmosDbV1Alpha1Manifest() datasource.DataSource {
	return &AzureMicrosoftComCosmosDbV1Alpha1Manifest{}
}

type AzureMicrosoftComCosmosDbV1Alpha1Manifest struct{}

type AzureMicrosoftComCosmosDbV1Alpha1ManifestData struct {
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
		IpRules                *[]string `tfsdk:"ip_rules" json:"ipRules,omitempty"`
		KeyVaultToStoreSecrets *string   `tfsdk:"key_vault_to_store_secrets" json:"keyVaultToStoreSecrets,omitempty"`
		Kind                   *string   `tfsdk:"kind" json:"kind,omitempty"`
		Location               *string   `tfsdk:"location" json:"location,omitempty"`
		Locations              *[]struct {
			FailoverPriority *int64  `tfsdk:"failover_priority" json:"failoverPriority,omitempty"`
			IsZoneRedundant  *bool   `tfsdk:"is_zone_redundant" json:"isZoneRedundant,omitempty"`
			LocationName     *string `tfsdk:"location_name" json:"locationName,omitempty"`
		} `tfsdk:"locations" json:"locations,omitempty"`
		Properties *struct {
			Capabilities *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"capabilities" json:"capabilities,omitempty"`
			DatabaseAccountOfferType      *string `tfsdk:"database_account_offer_type" json:"databaseAccountOfferType,omitempty"`
			EnableMultipleWriteLocations  *bool   `tfsdk:"enable_multiple_write_locations" json:"enableMultipleWriteLocations,omitempty"`
			IsVirtualNetworkFilterEnabled *bool   `tfsdk:"is_virtual_network_filter_enabled" json:"isVirtualNetworkFilterEnabled,omitempty"`
			MongoDBVersion                *string `tfsdk:"mongo_db_version" json:"mongoDBVersion,omitempty"`
		} `tfsdk:"properties" json:"properties,omitempty"`
		ResourceGroup       *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		VirtualNetworkRules *[]struct {
			IgnoreMissingVNetServiceEndpoint *bool   `tfsdk:"ignore_missing_v_net_service_endpoint" json:"ignoreMissingVNetServiceEndpoint,omitempty"`
			SubnetID                         *string `tfsdk:"subnet_id" json:"subnetID,omitempty"`
		} `tfsdk:"virtual_network_rules" json:"virtualNetworkRules,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComCosmosDbV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_cosmos_db_v1alpha1_manifest"
}

func (r *AzureMicrosoftComCosmosDbV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "CosmosDB is the Schema for the cosmosdbs API",
		MarkdownDescription: "CosmosDB is the Schema for the cosmosdbs API",
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
				Description:         "CosmosDBSpec defines the desired state of CosmosDB",
				MarkdownDescription: "CosmosDBSpec defines the desired state of CosmosDB",
				Attributes: map[string]schema.Attribute{
					"ip_rules": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"key_vault_to_store_secrets": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kind": schema.StringAttribute{
						Description:         "CosmosDBKind enumerates the values for kind. Only one of the following kinds may be specified. If none of the following kinds is specified, the default one is GlobalDocumentDBKind.",
						MarkdownDescription: "CosmosDBKind enumerates the values for kind. Only one of the following kinds may be specified. If none of the following kinds is specified, the default one is GlobalDocumentDBKind.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("GlobalDocumentDB", "MongoDB"),
						},
					},

					"location": schema.StringAttribute{
						Description:         "Location is the Azure location where the CosmosDB exists",
						MarkdownDescription: "Location is the Azure location where the CosmosDB exists",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(0),
						},
					},

					"locations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"failover_priority": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"is_zone_redundant": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"location_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"properties": schema.SingleNestedAttribute{
						Description:         "CosmosDBProperties the CosmosDBProperties of CosmosDB.",
						MarkdownDescription: "CosmosDBProperties the CosmosDBProperties of CosmosDB.",
						Attributes: map[string]schema.Attribute{
							"capabilities": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name *CosmosCapability 'json:'name,omitempty''",
											MarkdownDescription: "Name *CosmosCapability 'json:'name,omitempty''",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("EnableCassandra", "EnableTable", "EnableGremlin", "EnableMongo"),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"database_account_offer_type": schema.StringAttribute{
								Description:         "DatabaseAccountOfferType - The offer type for the Cosmos DB database account.",
								MarkdownDescription: "DatabaseAccountOfferType - The offer type for the Cosmos DB database account.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Standard"),
								},
							},

							"enable_multiple_write_locations": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"is_virtual_network_filter_enabled": schema.BoolAttribute{
								Description:         "IsVirtualNetworkFilterEnabled - Flag to indicate whether to enable/disable Virtual Network ACL rules.",
								MarkdownDescription: "IsVirtualNetworkFilterEnabled - Flag to indicate whether to enable/disable Virtual Network ACL rules.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mongo_db_version": schema.StringAttribute{
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

					"virtual_network_rules": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ignore_missing_v_net_service_endpoint": schema.BoolAttribute{
									Description:         "IgnoreMissingVNetServiceEndpoint - Create firewall rule before the virtual network has vnet service endpoint enabled.",
									MarkdownDescription: "IgnoreMissingVNetServiceEndpoint - Create firewall rule before the virtual network has vnet service endpoint enabled.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"subnet_id": schema.StringAttribute{
									Description:         "ID - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.",
									MarkdownDescription: "ID - Resource ID of a subnet, for example: /subscriptions/{subscriptionId}/resourceGroups/{groupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}.",
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
		},
	}
}

func (r *AzureMicrosoftComCosmosDbV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_cosmos_db_v1alpha1_manifest")

	var model AzureMicrosoftComCosmosDbV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("azure.microsoft.com/v1alpha1")
	model.Kind = pointer.String("CosmosDB")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
