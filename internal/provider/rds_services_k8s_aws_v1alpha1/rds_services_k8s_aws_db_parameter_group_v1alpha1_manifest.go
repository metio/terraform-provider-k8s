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
	_ datasource.DataSource = &RdsServicesK8SAwsDbparameterGroupV1Alpha1Manifest{}
)

func NewRdsServicesK8SAwsDbparameterGroupV1Alpha1Manifest() datasource.DataSource {
	return &RdsServicesK8SAwsDbparameterGroupV1Alpha1Manifest{}
}

type RdsServicesK8SAwsDbparameterGroupV1Alpha1Manifest struct{}

type RdsServicesK8SAwsDbparameterGroupV1Alpha1ManifestData struct {
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
		Description        *string            `tfsdk:"description" json:"description,omitempty"`
		Family             *string            `tfsdk:"family" json:"family,omitempty"`
		Name               *string            `tfsdk:"name" json:"name,omitempty"`
		ParameterOverrides *map[string]string `tfsdk:"parameter_overrides" json:"parameterOverrides,omitempty"`
		Tags               *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RdsServicesK8SAwsDbparameterGroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rds_services_k8s_aws_db_parameter_group_v1alpha1_manifest"
}

func (r *RdsServicesK8SAwsDbparameterGroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DBParameterGroup is the Schema for the DBParameterGroups API",
		MarkdownDescription: "DBParameterGroup is the Schema for the DBParameterGroups API",
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
				Description:         "DBParameterGroupSpec defines the desired state of DBParameterGroup. Contains the details of an Amazon RDS DB parameter group. This data type is used as a response element in the DescribeDBParameterGroups action.",
				MarkdownDescription: "DBParameterGroupSpec defines the desired state of DBParameterGroup. Contains the details of an Amazon RDS DB parameter group. This data type is used as a response element in the DescribeDBParameterGroups action.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "The description for the DB parameter group.",
						MarkdownDescription: "The description for the DB parameter group.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"family": schema.StringAttribute{
						Description:         "The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family. To list all of the available parameter group families for a DB engine, use the following command: aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine <engine> For example, to list all of the available parameter group families for the MySQL DB engine, use the following command: aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine mysql The output contains duplicates. The following are the valid DB engine values: * aurora (for MySQL 5.6-compatible Aurora) * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora) * aurora-postgresql * mariadb * mysql * oracle-ee * oracle-ee-cdb * oracle-se2 * oracle-se2-cdb * postgres * sqlserver-ee * sqlserver-se * sqlserver-ex * sqlserver-web",
						MarkdownDescription: "The DB parameter group family name. A DB parameter group can be associated with one and only one DB parameter group family, and can be applied only to a DB instance running a database engine and engine version compatible with that DB parameter group family. To list all of the available parameter group families for a DB engine, use the following command: aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine <engine> For example, to list all of the available parameter group families for the MySQL DB engine, use the following command: aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine mysql The output contains duplicates. The following are the valid DB engine values: * aurora (for MySQL 5.6-compatible Aurora) * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora) * aurora-postgresql * mariadb * mysql * oracle-ee * oracle-ee-cdb * oracle-se2 * oracle-se2-cdb * postgres * sqlserver-ee * sqlserver-se * sqlserver-ex * sqlserver-web",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the DB parameter group. Constraints: * Must be 1 to 255 letters, numbers, or hyphens. * First character must be a letter * Can't end with a hyphen or contain two consecutive hyphens This value is stored as a lowercase string.",
						MarkdownDescription: "The name of the DB parameter group. Constraints: * Must be 1 to 255 letters, numbers, or hyphens. * First character must be a letter * Can't end with a hyphen or contain two consecutive hyphens This value is stored as a lowercase string.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"parameter_overrides": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "Tags to assign to the DB parameter group.",
						MarkdownDescription: "Tags to assign to the DB parameter group.",
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

func (r *RdsServicesK8SAwsDbparameterGroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rds_services_k8s_aws_db_parameter_group_v1alpha1_manifest")

	var model RdsServicesK8SAwsDbparameterGroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("DBParameterGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
