/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package rds_services_k8s_aws_v1alpha1

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
	_ datasource.DataSource = &RdsServicesK8SAwsGlobalClusterV1Alpha1Manifest{}
)

func NewRdsServicesK8SAwsGlobalClusterV1Alpha1Manifest() datasource.DataSource {
	return &RdsServicesK8SAwsGlobalClusterV1Alpha1Manifest{}
}

type RdsServicesK8SAwsGlobalClusterV1Alpha1Manifest struct{}

type RdsServicesK8SAwsGlobalClusterV1Alpha1ManifestData struct {
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
		DatabaseName              *string `tfsdk:"database_name" json:"databaseName,omitempty"`
		DeletionProtection        *bool   `tfsdk:"deletion_protection" json:"deletionProtection,omitempty"`
		Engine                    *string `tfsdk:"engine" json:"engine,omitempty"`
		EngineVersion             *string `tfsdk:"engine_version" json:"engineVersion,omitempty"`
		GlobalClusterIdentifier   *string `tfsdk:"global_cluster_identifier" json:"globalClusterIdentifier,omitempty"`
		SourceDBClusterIdentifier *string `tfsdk:"source_db_cluster_identifier" json:"sourceDBClusterIdentifier,omitempty"`
		StorageEncrypted          *bool   `tfsdk:"storage_encrypted" json:"storageEncrypted,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RdsServicesK8SAwsGlobalClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rds_services_k8s_aws_global_cluster_v1alpha1_manifest"
}

func (r *RdsServicesK8SAwsGlobalClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "GlobalCluster is the Schema for the GlobalClusters API",
		MarkdownDescription: "GlobalCluster is the Schema for the GlobalClusters API",
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
				Description:         "GlobalClusterSpec defines the desired state of GlobalCluster. A data type representing an Aurora global database.",
				MarkdownDescription: "GlobalClusterSpec defines the desired state of GlobalCluster. A data type representing an Aurora global database.",
				Attributes: map[string]schema.Attribute{
					"database_name": schema.StringAttribute{
						Description:         "The name for your database of up to 64 alphanumeric characters. If you don't specify a name, Amazon Aurora doesn't create a database in the global database cluster. Constraints: * Can't be specified if SourceDBClusterIdentifier is specified. In this case, Amazon Aurora uses the database name from the source DB cluster.",
						MarkdownDescription: "The name for your database of up to 64 alphanumeric characters. If you don't specify a name, Amazon Aurora doesn't create a database in the global database cluster. Constraints: * Can't be specified if SourceDBClusterIdentifier is specified. In this case, Amazon Aurora uses the database name from the source DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deletion_protection": schema.BoolAttribute{
						Description:         "Specifies whether to enable deletion protection for the new global database cluster. The global database can't be deleted when deletion protection is enabled.",
						MarkdownDescription: "Specifies whether to enable deletion protection for the new global database cluster. The global database can't be deleted when deletion protection is enabled.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine": schema.StringAttribute{
						Description:         "The database engine to use for this global database cluster. Valid Values: aurora-mysql | aurora-postgresql Constraints: * Can't be specified if SourceDBClusterIdentifier is specified. In this case, Amazon Aurora uses the engine of the source DB cluster.",
						MarkdownDescription: "The database engine to use for this global database cluster. Valid Values: aurora-mysql | aurora-postgresql Constraints: * Can't be specified if SourceDBClusterIdentifier is specified. In this case, Amazon Aurora uses the engine of the source DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine_version": schema.StringAttribute{
						Description:         "The engine version to use for this global database cluster. Constraints: * Can't be specified if SourceDBClusterIdentifier is specified. In this case, Amazon Aurora uses the engine version of the source DB cluster.",
						MarkdownDescription: "The engine version to use for this global database cluster. Constraints: * Can't be specified if SourceDBClusterIdentifier is specified. In this case, Amazon Aurora uses the engine version of the source DB cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"global_cluster_identifier": schema.StringAttribute{
						Description:         "The cluster identifier for this global database cluster. This parameter is stored as a lowercase string.",
						MarkdownDescription: "The cluster identifier for this global database cluster. This parameter is stored as a lowercase string.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"source_db_cluster_identifier": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. If you provide a value for this parameter, don't specify values for the following settings because Amazon Aurora uses the values from the specified source DB cluster: * DatabaseName * Engine * EngineVersion * StorageEncrypted",
						MarkdownDescription: "The Amazon Resource Name (ARN) to use as the primary cluster of the global database. If you provide a value for this parameter, don't specify values for the following settings because Amazon Aurora uses the values from the specified source DB cluster: * DatabaseName * Engine * EngineVersion * StorageEncrypted",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage_encrypted": schema.BoolAttribute{
						Description:         "Specifies whether to enable storage encryption for the new global database cluster. Constraints: * Can't be specified if SourceDBClusterIdentifier is specified. In this case, Amazon Aurora uses the setting from the source DB cluster.",
						MarkdownDescription: "Specifies whether to enable storage encryption for the new global database cluster. Constraints: * Can't be specified if SourceDBClusterIdentifier is specified. In this case, Amazon Aurora uses the setting from the source DB cluster.",
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

func (r *RdsServicesK8SAwsGlobalClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rds_services_k8s_aws_global_cluster_v1alpha1_manifest")

	var model RdsServicesK8SAwsGlobalClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("GlobalCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
