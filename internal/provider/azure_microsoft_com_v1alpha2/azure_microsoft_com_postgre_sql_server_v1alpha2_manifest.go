/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package azure_microsoft_com_v1alpha2

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
	_ datasource.DataSource = &AzureMicrosoftComPostgreSqlserverV1Alpha2Manifest{}
)

func NewAzureMicrosoftComPostgreSqlserverV1Alpha2Manifest() datasource.DataSource {
	return &AzureMicrosoftComPostgreSqlserverV1Alpha2Manifest{}
}

type AzureMicrosoftComPostgreSqlserverV1Alpha2Manifest struct{}

type AzureMicrosoftComPostgreSqlserverV1Alpha2ManifestData struct {
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
		CreateMode             *string `tfsdk:"create_mode" json:"createMode,omitempty"`
		KeyVaultToStoreSecrets *string `tfsdk:"key_vault_to_store_secrets" json:"keyVaultToStoreSecrets,omitempty"`
		Location               *string `tfsdk:"location" json:"location,omitempty"`
		ReplicaProperties      *struct {
			SourceServerId *string `tfsdk:"source_server_id" json:"sourceServerId,omitempty"`
		} `tfsdk:"replica_properties" json:"replicaProperties,omitempty"`
		ResourceGroup *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		ServerVersion *string `tfsdk:"server_version" json:"serverVersion,omitempty"`
		Sku           *struct {
			Capacity *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
			Family   *string `tfsdk:"family" json:"family,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Size     *string `tfsdk:"size" json:"size,omitempty"`
			Tier     *string `tfsdk:"tier" json:"tier,omitempty"`
		} `tfsdk:"sku" json:"sku,omitempty"`
		SslEnforcement *string `tfsdk:"ssl_enforcement" json:"sslEnforcement,omitempty"`
		StorageProfile *struct {
			BackupRetentionDays *int64  `tfsdk:"backup_retention_days" json:"backupRetentionDays,omitempty"`
			GeoRedundantBackup  *string `tfsdk:"geo_redundant_backup" json:"geoRedundantBackup,omitempty"`
			StorageAutogrow     *string `tfsdk:"storage_autogrow" json:"storageAutogrow,omitempty"`
			StorageMB           *int64  `tfsdk:"storage_mb" json:"storageMB,omitempty"`
		} `tfsdk:"storage_profile" json:"storageProfile,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComPostgreSqlserverV1Alpha2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_postgre_sql_server_v1alpha2_manifest"
}

func (r *AzureMicrosoftComPostgreSqlserverV1Alpha2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "PostgreSQLServer is the Schema for the postgresqlservers API",
		MarkdownDescription: "PostgreSQLServer is the Schema for the postgresqlservers API",
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
				Description:         "PostgreSQLServerSpec defines the desired state of PostgreSQLServer",
				MarkdownDescription: "PostgreSQLServerSpec defines the desired state of PostgreSQLServer",
				Attributes: map[string]schema.Attribute{
					"create_mode": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
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

					"location": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"replica_properties": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"source_server_id": schema.StringAttribute{
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

					"server_version": schema.StringAttribute{
						Description:         "ServerVersion enumerates the values for server version.",
						MarkdownDescription: "ServerVersion enumerates the values for server version.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sku": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"capacity": schema.Int64Attribute{
								Description:         "Capacity - The scale up/out capacity, representing server's compute units.",
								MarkdownDescription: "Capacity - The scale up/out capacity, representing server's compute units.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"family": schema.StringAttribute{
								Description:         "Family - The family of hardware.",
								MarkdownDescription: "Family - The family of hardware.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name - The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.",
								MarkdownDescription: "Name - The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size - The size code, to be interpreted by resource as appropriate.",
								MarkdownDescription: "Size - The size code, to be interpreted by resource as appropriate.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tier": schema.StringAttribute{
								Description:         "Tier - The tier of the particular SKU, e.g. Basic. Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'",
								MarkdownDescription: "Tier - The tier of the particular SKU, e.g. Basic. Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Basic", "GeneralPurpose", "MemoryOptimized"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ssl_enforcement": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Enabled", "Disabled"),
						},
					},

					"storage_profile": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"backup_retention_days": schema.Int64Attribute{
								Description:         "BackupRetentionDays - Backup retention days for the server.",
								MarkdownDescription: "BackupRetentionDays - Backup retention days for the server.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"geo_redundant_backup": schema.StringAttribute{
								Description:         "GeoRedundantBackup - Enable Geo-redundant or not for server backup. Possible values include: 'Enabled', 'Disabled'",
								MarkdownDescription: "GeoRedundantBackup - Enable Geo-redundant or not for server backup. Possible values include: 'Enabled', 'Disabled'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_autogrow": schema.StringAttribute{
								Description:         "StorageAutogrow - Enable Storage Auto Grow. Possible values include: 'Enabled', 'Disabled'",
								MarkdownDescription: "StorageAutogrow - Enable Storage Auto Grow. Possible values include: 'Enabled', 'Disabled'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Enabled", "Disabled"),
								},
							},

							"storage_mb": schema.Int64Attribute{
								Description:         "StorageMB - Max storage allowed for a server.",
								MarkdownDescription: "StorageMB - Max storage allowed for a server.",
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

func (r *AzureMicrosoftComPostgreSqlserverV1Alpha2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_postgre_sql_server_v1alpha2_manifest")

	var model AzureMicrosoftComPostgreSqlserverV1Alpha2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("azure.microsoft.com/v1alpha2")
	model.Kind = pointer.String("PostgreSQLServer")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
