/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package azure_microsoft_com_v1beta1

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
	_ datasource.DataSource = &AzureMicrosoftComAzureSqlDatabaseV1Beta1Manifest{}
)

func NewAzureMicrosoftComAzureSqlDatabaseV1Beta1Manifest() datasource.DataSource {
	return &AzureMicrosoftComAzureSqlDatabaseV1Beta1Manifest{}
}

type AzureMicrosoftComAzureSqlDatabaseV1Beta1Manifest struct{}

type AzureMicrosoftComAzureSqlDatabaseV1Beta1ManifestData struct {
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
		DbName                   *string `tfsdk:"db_name" json:"dbName,omitempty"`
		Edition                  *int64  `tfsdk:"edition" json:"edition,omitempty"`
		ElasticPoolId            *string `tfsdk:"elastic_pool_id" json:"elasticPoolId,omitempty"`
		Location                 *string `tfsdk:"location" json:"location,omitempty"`
		MaxSize                  *string `tfsdk:"max_size" json:"maxSize,omitempty"`
		MonthlyRetention         *string `tfsdk:"monthly_retention" json:"monthlyRetention,omitempty"`
		ResourceGroup            *string `tfsdk:"resource_group" json:"resourceGroup,omitempty"`
		Server                   *string `tfsdk:"server" json:"server,omitempty"`
		ShortTermRetentionPolicy *struct {
			RetentionDays *int64 `tfsdk:"retention_days" json:"retentionDays,omitempty"`
		} `tfsdk:"short_term_retention_policy" json:"shortTermRetentionPolicy,omitempty"`
		Sku *struct {
			Capacity *int64  `tfsdk:"capacity" json:"capacity,omitempty"`
			Family   *string `tfsdk:"family" json:"family,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Size     *string `tfsdk:"size" json:"size,omitempty"`
			Tier     *string `tfsdk:"tier" json:"tier,omitempty"`
		} `tfsdk:"sku" json:"sku,omitempty"`
		SubscriptionId  *string `tfsdk:"subscription_id" json:"subscriptionId,omitempty"`
		WeekOfYear      *int64  `tfsdk:"week_of_year" json:"weekOfYear,omitempty"`
		WeeklyRetention *string `tfsdk:"weekly_retention" json:"weeklyRetention,omitempty"`
		YearlyRetention *string `tfsdk:"yearly_retention" json:"yearlyRetention,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AzureMicrosoftComAzureSqlDatabaseV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_azure_microsoft_com_azure_sql_database_v1beta1_manifest"
}

func (r *AzureMicrosoftComAzureSqlDatabaseV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AzureSqlDatabase is the Schema for the azuresqldatabases API",
		MarkdownDescription: "AzureSqlDatabase is the Schema for the azuresqldatabases API",
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
				Description:         "AzureSqlDatabaseSpec defines the desired state of AzureSqlDatabase",
				MarkdownDescription: "AzureSqlDatabaseSpec defines the desired state of AzureSqlDatabase",
				Attributes: map[string]schema.Attribute{
					"db_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"edition": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"elastic_pool_id": schema.StringAttribute{
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
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"max_size": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"monthly_retention": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
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

					"server": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(1),
						},
					},

					"short_term_retention_policy": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"retention_days": schema.Int64Attribute{
								Description:         "RetentionDays is the backup retention period in days. This is how many days Point-in-Time Restore will be supported.",
								MarkdownDescription: "RetentionDays is the backup retention period in days. This is how many days Point-in-Time Restore will be supported.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"sku": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"capacity": schema.Int64Attribute{
								Description:         "Capacity - Capacity of the particular SKU.",
								MarkdownDescription: "Capacity - Capacity of the particular SKU.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"family": schema.StringAttribute{
								Description:         "Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.",
								MarkdownDescription: "Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name - The name of the SKU, typically, a letter + Number code, e.g. P3.",
								MarkdownDescription: "Name - The name of the SKU, typically, a letter + Number code, e.g. P3.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},

							"size": schema.StringAttribute{
								Description:         "Size - Size of the particular SKU",
								MarkdownDescription: "Size - Size of the particular SKU",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tier": schema.StringAttribute{
								Description:         "optional Tier - The tier or edition of the particular SKU, e.g. Basic, Premium.",
								MarkdownDescription: "optional Tier - The tier or edition of the particular SKU, e.g. Basic, Premium.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"subscription_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"week_of_year": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"weekly_retention": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"yearly_retention": schema.StringAttribute{
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

func (r *AzureMicrosoftComAzureSqlDatabaseV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_azure_microsoft_com_azure_sql_database_v1beta1_manifest")

	var model AzureMicrosoftComAzureSqlDatabaseV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("azure.microsoft.com/v1beta1")
	model.Kind = pointer.String("AzureSqlDatabase")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
