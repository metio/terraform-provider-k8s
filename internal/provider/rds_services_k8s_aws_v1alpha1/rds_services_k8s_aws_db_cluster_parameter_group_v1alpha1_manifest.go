/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package rds_services_k8s_aws_v1alpha1

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &RdsServicesK8SAwsDBClusterParameterGroupV1Alpha1Manifest{}
)

func NewRdsServicesK8SAwsDBClusterParameterGroupV1Alpha1Manifest() datasource.DataSource {
	return &RdsServicesK8SAwsDBClusterParameterGroupV1Alpha1Manifest{}
}

type RdsServicesK8SAwsDBClusterParameterGroupV1Alpha1Manifest struct{}

type RdsServicesK8SAwsDBClusterParameterGroupV1Alpha1ManifestData struct {
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
		Description        *string            `tfsdk:"description" json:"description,omitempty"`
		Family             *string            `tfsdk:"family" json:"family,omitempty"`
		Name               *string            `tfsdk:"name" json:"name,omitempty"`
		ParameterOverrides *map[string]string `tfsdk:"parameter_overrides" json:"parameterOverrides,omitempty"`
		Parameters         *[]struct {
			AllowedValues        *string   `tfsdk:"allowed_values" json:"allowedValues,omitempty"`
			ApplyMethod          *string   `tfsdk:"apply_method" json:"applyMethod,omitempty"`
			ApplyType            *string   `tfsdk:"apply_type" json:"applyType,omitempty"`
			DataType             *string   `tfsdk:"data_type" json:"dataType,omitempty"`
			Description          *string   `tfsdk:"description" json:"description,omitempty"`
			IsModifiable         *bool     `tfsdk:"is_modifiable" json:"isModifiable,omitempty"`
			MinimumEngineVersion *string   `tfsdk:"minimum_engine_version" json:"minimumEngineVersion,omitempty"`
			ParameterName        *string   `tfsdk:"parameter_name" json:"parameterName,omitempty"`
			ParameterValue       *string   `tfsdk:"parameter_value" json:"parameterValue,omitempty"`
			Source               *string   `tfsdk:"source" json:"source,omitempty"`
			SupportedEngineModes *[]string `tfsdk:"supported_engine_modes" json:"supportedEngineModes,omitempty"`
		} `tfsdk:"parameters" json:"parameters,omitempty"`
		Tags *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RdsServicesK8SAwsDBClusterParameterGroupV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest"
}

func (r *RdsServicesK8SAwsDBClusterParameterGroupV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DBClusterParameterGroup is the Schema for the DBClusterParameterGroups API",
		MarkdownDescription: "DBClusterParameterGroup is the Schema for the DBClusterParameterGroups API",
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
				Description:         "DBClusterParameterGroupSpec defines the desired state of DBClusterParameterGroup.  Contains the details of an Amazon RDS DB cluster parameter group.  This data type is used as a response element in the DescribeDBClusterParameterGroups action.",
				MarkdownDescription: "DBClusterParameterGroupSpec defines the desired state of DBClusterParameterGroup.  Contains the details of an Amazon RDS DB cluster parameter group.  This data type is used as a response element in the DescribeDBClusterParameterGroups action.",
				Attributes: map[string]schema.Attribute{
					"description": schema.StringAttribute{
						Description:         "The description for the DB cluster parameter group.",
						MarkdownDescription: "The description for the DB cluster parameter group.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"family": schema.StringAttribute{
						Description:         "The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.  Aurora MySQL  Example: aurora5.6, aurora-mysql5.7, aurora-mysql8.0  Aurora PostgreSQL  Example: aurora-postgresql9.6  RDS for MySQL  Example: mysql8.0  RDS for PostgreSQL  Example: postgres12  To list all of the available parameter group families for a DB engine, use the following command:  aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine <engine>  For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:  aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine aurora-postgresql  The output contains duplicates.  The following are the valid DB engine values:  * aurora (for MySQL 5.6-compatible Aurora)  * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)  * aurora-postgresql  * mysql  * postgres",
						MarkdownDescription: "The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.  Aurora MySQL  Example: aurora5.6, aurora-mysql5.7, aurora-mysql8.0  Aurora PostgreSQL  Example: aurora-postgresql9.6  RDS for MySQL  Example: mysql8.0  RDS for PostgreSQL  Example: postgres12  To list all of the available parameter group families for a DB engine, use the following command:  aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine <engine>  For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:  aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine aurora-postgresql  The output contains duplicates.  The following are the valid DB engine values:  * aurora (for MySQL 5.6-compatible Aurora)  * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)  * aurora-postgresql  * mysql  * postgres",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the DB cluster parameter group.  Constraints:  * Must not match the name of an existing DB cluster parameter group.  This value is stored as a lowercase string.",
						MarkdownDescription: "The name of the DB cluster parameter group.  Constraints:  * Must not match the name of an existing DB cluster parameter group.  This value is stored as a lowercase string.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"parameter_overrides": schema.MapAttribute{
						Description:         "These are ONLY user-defined parameter overrides for the DB cluster parameter group. This does not contain default or system parameters.",
						MarkdownDescription: "These are ONLY user-defined parameter overrides for the DB cluster parameter group. This does not contain default or system parameters.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"parameters": schema.ListNestedAttribute{
						Description:         "A list of parameters in the DB cluster parameter group to modify.  Valid Values (for the application method): immediate | pending-reboot  You can use the immediate value with dynamic parameters only. You can use the pending-reboot value for both dynamic and static parameters.  When the application method is immediate, changes to dynamic parameters are applied immediately to the DB clusters associated with the parameter group. When the application method is pending-reboot, changes to dynamic and static parameters are applied after a reboot without failover to the DB clusters associated with the parameter group.  DEPRECATED - do not use.  Prefer ParameterOverrides instead.",
						MarkdownDescription: "A list of parameters in the DB cluster parameter group to modify.  Valid Values (for the application method): immediate | pending-reboot  You can use the immediate value with dynamic parameters only. You can use the pending-reboot value for both dynamic and static parameters.  When the application method is immediate, changes to dynamic parameters are applied immediately to the DB clusters associated with the parameter group. When the application method is pending-reboot, changes to dynamic and static parameters are applied after a reboot without failover to the DB clusters associated with the parameter group.  DEPRECATED - do not use.  Prefer ParameterOverrides instead.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"allowed_values": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"apply_method": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"apply_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"data_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"description": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"is_modifiable": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"minimum_engine_version": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"parameter_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"parameter_value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"supported_engine_modes": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"tags": schema.ListNestedAttribute{
						Description:         "Tags to assign to the DB cluster parameter group.",
						MarkdownDescription: "Tags to assign to the DB cluster parameter group.",
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

func (r *RdsServicesK8SAwsDBClusterParameterGroupV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1_manifest")

	var model RdsServicesK8SAwsDBClusterParameterGroupV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("DBClusterParameterGroup")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
