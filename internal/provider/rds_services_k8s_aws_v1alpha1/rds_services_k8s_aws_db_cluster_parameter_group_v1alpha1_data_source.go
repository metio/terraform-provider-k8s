/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package rds_services_k8s_aws_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource{}
)

func NewRdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource() datasource.DataSource {
	return &RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource{}
}

type RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

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

func (r *RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1"
}

func (r *RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
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
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"family": schema.StringAttribute{
						Description:         "The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.  Aurora MySQL  Example: aurora5.6, aurora-mysql5.7, aurora-mysql8.0  Aurora PostgreSQL  Example: aurora-postgresql9.6  RDS for MySQL  Example: mysql8.0  RDS for PostgreSQL  Example: postgres12  To list all of the available parameter group families for a DB engine, use the following command:  aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine <engine>  For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:  aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine aurora-postgresql  The output contains duplicates.  The following are the valid DB engine values:  * aurora (for MySQL 5.6-compatible Aurora)  * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)  * aurora-postgresql  * mysql  * postgres",
						MarkdownDescription: "The DB cluster parameter group family name. A DB cluster parameter group can be associated with one and only one DB cluster parameter group family, and can be applied only to a DB cluster running a database engine and engine version compatible with that DB cluster parameter group family.  Aurora MySQL  Example: aurora5.6, aurora-mysql5.7, aurora-mysql8.0  Aurora PostgreSQL  Example: aurora-postgresql9.6  RDS for MySQL  Example: mysql8.0  RDS for PostgreSQL  Example: postgres12  To list all of the available parameter group families for a DB engine, use the following command:  aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine <engine>  For example, to list all of the available parameter group families for the Aurora PostgreSQL DB engine, use the following command:  aws rds describe-db-engine-versions --query 'DBEngineVersions[].DBParameterGroupFamily' --engine aurora-postgresql  The output contains duplicates.  The following are the valid DB engine values:  * aurora (for MySQL 5.6-compatible Aurora)  * aurora-mysql (for MySQL 5.7-compatible and MySQL 8.0-compatible Aurora)  * aurora-postgresql  * mysql  * postgres",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"name": schema.StringAttribute{
						Description:         "The name of the DB cluster parameter group.  Constraints:  * Must not match the name of an existing DB cluster parameter group.  This value is stored as a lowercase string.",
						MarkdownDescription: "The name of the DB cluster parameter group.  Constraints:  * Must not match the name of an existing DB cluster parameter group.  This value is stored as a lowercase string.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"parameter_overrides": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"parameters": schema.ListNestedAttribute{
						Description:         "A list of parameters in the DB cluster parameter group to modify.  Valid Values (for the application method): immediate | pending-reboot  You can use the immediate value with dynamic parameters only. You can use the pending-reboot value for both dynamic and static parameters.  When the application method is immediate, changes to dynamic parameters are applied immediately to the DB clusters associated with the parameter group. When the application method is pending-reboot, changes to dynamic and static parameters are applied after a reboot without failover to the DB clusters associated with the parameter group.",
						MarkdownDescription: "A list of parameters in the DB cluster parameter group to modify.  Valid Values (for the application method): immediate | pending-reboot  You can use the immediate value with dynamic parameters only. You can use the pending-reboot value for both dynamic and static parameters.  When the application method is immediate, changes to dynamic parameters are applied immediately to the DB clusters associated with the parameter group. When the application method is pending-reboot, changes to dynamic and static parameters are applied after a reboot without failover to the DB clusters associated with the parameter group.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"allowed_values": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"apply_method": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"apply_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"data_type": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"description": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"is_modifiable": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"minimum_engine_version": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"parameter_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"parameter_value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"source": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"supported_engine_modes": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_rds_services_k8s_aws_db_cluster_parameter_group_v1alpha1")

	var data RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "rds.services.k8s.aws", Version: "v1alpha1", Resource: "DBClusterParameterGroup"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse RdsServicesK8SAwsDbclusterParameterGroupV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	data.Kind = pointer.String("DBClusterParameterGroup")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
