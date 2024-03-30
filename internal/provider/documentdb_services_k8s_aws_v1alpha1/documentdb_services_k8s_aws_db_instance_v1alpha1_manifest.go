/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package documentdb_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &DocumentdbServicesK8SAwsDbinstanceV1Alpha1Manifest{}
)

func NewDocumentdbServicesK8SAwsDbinstanceV1Alpha1Manifest() datasource.DataSource {
	return &DocumentdbServicesK8SAwsDbinstanceV1Alpha1Manifest{}
}

type DocumentdbServicesK8SAwsDbinstanceV1Alpha1Manifest struct{}

type DocumentdbServicesK8SAwsDbinstanceV1Alpha1ManifestData struct {
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
		AutoMinorVersionUpgrade      *bool   `tfsdk:"auto_minor_version_upgrade" json:"autoMinorVersionUpgrade,omitempty"`
		AvailabilityZone             *string `tfsdk:"availability_zone" json:"availabilityZone,omitempty"`
		CaCertificateIdentifier      *string `tfsdk:"ca_certificate_identifier" json:"caCertificateIdentifier,omitempty"`
		CopyTagsToSnapshot           *bool   `tfsdk:"copy_tags_to_snapshot" json:"copyTagsToSnapshot,omitempty"`
		DbClusterIdentifier          *string `tfsdk:"db_cluster_identifier" json:"dbClusterIdentifier,omitempty"`
		DbInstanceClass              *string `tfsdk:"db_instance_class" json:"dbInstanceClass,omitempty"`
		DbInstanceIdentifier         *string `tfsdk:"db_instance_identifier" json:"dbInstanceIdentifier,omitempty"`
		Engine                       *string `tfsdk:"engine" json:"engine,omitempty"`
		PerformanceInsightsEnabled   *bool   `tfsdk:"performance_insights_enabled" json:"performanceInsightsEnabled,omitempty"`
		PerformanceInsightsKMSKeyID  *string `tfsdk:"performance_insights_kms_key_id" json:"performanceInsightsKMSKeyID,omitempty"`
		PerformanceInsightsKMSKeyRef *struct {
			From *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"from" json:"from,omitempty"`
		} `tfsdk:"performance_insights_kms_key_ref" json:"performanceInsightsKMSKeyRef,omitempty"`
		PreferredMaintenanceWindow *string `tfsdk:"preferred_maintenance_window" json:"preferredMaintenanceWindow,omitempty"`
		PromotionTier              *int64  `tfsdk:"promotion_tier" json:"promotionTier,omitempty"`
		Tags                       *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DocumentdbServicesK8SAwsDbinstanceV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_documentdb_services_k8s_aws_db_instance_v1alpha1_manifest"
}

func (r *DocumentdbServicesK8SAwsDbinstanceV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DBInstance is the Schema for the DBInstances API",
		MarkdownDescription: "DBInstance is the Schema for the DBInstances API",
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
				Description:         "DBInstanceSpec defines the desired state of DBInstance.Detailed information about an instance.",
				MarkdownDescription: "DBInstanceSpec defines the desired state of DBInstance.Detailed information about an instance.",
				Attributes: map[string]schema.Attribute{
					"auto_minor_version_upgrade": schema.BoolAttribute{
						Description:         "This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB doesnot perform minor version upgrades regardless of the value set.Default: false",
						MarkdownDescription: "This parameter does not apply to Amazon DocumentDB. Amazon DocumentDB doesnot perform minor version upgrades regardless of the value set.Default: false",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"availability_zone": schema.StringAttribute{
						Description:         "The Amazon EC2 Availability Zone that the instance is created in.Default: A random, system-chosen Availability Zone in the endpoint's AmazonWeb Services Region.Example: us-east-1d",
						MarkdownDescription: "The Amazon EC2 Availability Zone that the instance is created in.Default: A random, system-chosen Availability Zone in the endpoint's AmazonWeb Services Region.Example: us-east-1d",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ca_certificate_identifier": schema.StringAttribute{
						Description:         "The CA certificate identifier to use for the DB instance's server certificate.For more information, see Updating Your Amazon DocumentDB TLS Certificates(https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html)and Encrypting Data in Transit (https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html)in the Amazon DocumentDB Developer Guide.",
						MarkdownDescription: "The CA certificate identifier to use for the DB instance's server certificate.For more information, see Updating Your Amazon DocumentDB TLS Certificates(https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html)and Encrypting Data in Transit (https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html)in the Amazon DocumentDB Developer Guide.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"copy_tags_to_snapshot": schema.BoolAttribute{
						Description:         "A value that indicates whether to copy tags from the DB instance to snapshotsof the DB instance. By default, tags are not copied.",
						MarkdownDescription: "A value that indicates whether to copy tags from the DB instance to snapshotsof the DB instance. By default, tags are not copied.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"db_cluster_identifier": schema.StringAttribute{
						Description:         "The identifier of the cluster that the instance will belong to.",
						MarkdownDescription: "The identifier of the cluster that the instance will belong to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_instance_class": schema.StringAttribute{
						Description:         "The compute and memory capacity of the instance; for example, db.r5.large.",
						MarkdownDescription: "The compute and memory capacity of the instance; for example, db.r5.large.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"db_instance_identifier": schema.StringAttribute{
						Description:         "The instance identifier. This parameter is stored as a lowercase string.Constraints:   * Must contain from 1 to 63 letters, numbers, or hyphens.   * The first character must be a letter.   * Cannot end with a hyphen or contain two consecutive hyphens.Example: mydbinstance",
						MarkdownDescription: "The instance identifier. This parameter is stored as a lowercase string.Constraints:   * Must contain from 1 to 63 letters, numbers, or hyphens.   * The first character must be a letter.   * Cannot end with a hyphen or contain two consecutive hyphens.Example: mydbinstance",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"engine": schema.StringAttribute{
						Description:         "The name of the database engine to be used for this instance.Valid value: docdb",
						MarkdownDescription: "The name of the database engine to be used for this instance.Valid value: docdb",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"performance_insights_enabled": schema.BoolAttribute{
						Description:         "A value that indicates whether to enable Performance Insights for the DBInstance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html).",
						MarkdownDescription: "A value that indicates whether to enable Performance Insights for the DBInstance. For more information, see Using Amazon Performance Insights (https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_kms_key_id": schema.StringAttribute{
						Description:         "The KMS key identifier for encryption of Performance Insights data.The KMS key identifier is the key ARN, key ID, alias ARN, or alias name forthe KMS key.If you do not specify a value for PerformanceInsightsKMSKeyId, then AmazonDocumentDB uses your default KMS key. There is a default KMS key for yourAmazon Web Services account. Your Amazon Web Services account has a differentdefault KMS key for each Amazon Web Services region.",
						MarkdownDescription: "The KMS key identifier for encryption of Performance Insights data.The KMS key identifier is the key ARN, key ID, alias ARN, or alias name forthe KMS key.If you do not specify a value for PerformanceInsightsKMSKeyId, then AmazonDocumentDB uses your default KMS key. There is a default KMS key for yourAmazon Web Services account. Your Amazon Web Services account has a differentdefault KMS key for each Amazon Web Services region.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"performance_insights_kms_key_ref": schema.SingleNestedAttribute{
						Description:         "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						MarkdownDescription: "AWSResourceReferenceWrapper provides a wrapper around *AWSResourceReferencetype to provide more user friendly syntax for references using 'from' fieldEx:APIIDRef:	from:	  name: my-api",
						Attributes: map[string]schema.Attribute{
							"from": schema.SingleNestedAttribute{
								Description:         "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								MarkdownDescription: "AWSResourceReference provides all the values necessary to reference anotherk8s resource for finding the identifier(Id/ARN/Name)",
								Attributes: map[string]schema.Attribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"preferred_maintenance_window": schema.StringAttribute{
						Description:         "The time range each week during which system maintenance can occur, in UniversalCoordinated Time (UTC).Format: ddd:hh24:mi-ddd:hh24:miThe default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region, occurring on a random day ofthe week.Valid days: Mon, Tue, Wed, Thu, Fri, Sat, SunConstraints: Minimum 30-minute window.",
						MarkdownDescription: "The time range each week during which system maintenance can occur, in UniversalCoordinated Time (UTC).Format: ddd:hh24:mi-ddd:hh24:miThe default is a 30-minute window selected at random from an 8-hour blockof time for each Amazon Web Services Region, occurring on a random day ofthe week.Valid days: Mon, Tue, Wed, Thu, Fri, Sat, SunConstraints: Minimum 30-minute window.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"promotion_tier": schema.Int64Attribute{
						Description:         "A value that specifies the order in which an Amazon DocumentDB replica ispromoted to the primary instance after a failure of the existing primaryinstance.Default: 1Valid values: 0-15",
						MarkdownDescription: "A value that specifies the order in which an Amazon DocumentDB replica ispromoted to the primary instance after a failure of the existing primaryinstance.Default: 1Valid values: 0-15",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "The tags to be assigned to the instance. You can assign up to 10 tags toan instance.",
						MarkdownDescription: "The tags to be assigned to the instance. You can assign up to 10 tags toan instance.",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *DocumentdbServicesK8SAwsDbinstanceV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_documentdb_services_k8s_aws_db_instance_v1alpha1_manifest")

	var model DocumentdbServicesK8SAwsDbinstanceV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("documentdb.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("DBInstance")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
