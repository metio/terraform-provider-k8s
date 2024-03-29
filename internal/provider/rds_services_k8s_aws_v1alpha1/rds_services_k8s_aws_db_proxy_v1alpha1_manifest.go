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
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &RdsServicesK8SAwsDbproxyV1Alpha1Manifest{}
)

func NewRdsServicesK8SAwsDbproxyV1Alpha1Manifest() datasource.DataSource {
	return &RdsServicesK8SAwsDbproxyV1Alpha1Manifest{}
}

type RdsServicesK8SAwsDbproxyV1Alpha1Manifest struct{}

type RdsServicesK8SAwsDbproxyV1Alpha1ManifestData struct {
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
		Auth *[]struct {
			AuthScheme             *string `tfsdk:"auth_scheme" json:"authScheme,omitempty"`
			ClientPasswordAuthType *string `tfsdk:"client_password_auth_type" json:"clientPasswordAuthType,omitempty"`
			Description            *string `tfsdk:"description" json:"description,omitempty"`
			IamAuth                *string `tfsdk:"iam_auth" json:"iamAuth,omitempty"`
			SecretARN              *string `tfsdk:"secret_arn" json:"secretARN,omitempty"`
			UserName               *string `tfsdk:"user_name" json:"userName,omitempty"`
		} `tfsdk:"auth" json:"auth,omitempty"`
		DebugLogging      *bool   `tfsdk:"debug_logging" json:"debugLogging,omitempty"`
		EngineFamily      *string `tfsdk:"engine_family" json:"engineFamily,omitempty"`
		IdleClientTimeout *int64  `tfsdk:"idle_client_timeout" json:"idleClientTimeout,omitempty"`
		Name              *string `tfsdk:"name" json:"name,omitempty"`
		RequireTLS        *bool   `tfsdk:"require_tls" json:"requireTLS,omitempty"`
		RoleARN           *string `tfsdk:"role_arn" json:"roleARN,omitempty"`
		Tags              *[]struct {
			Key   *string `tfsdk:"key" json:"key,omitempty"`
			Value *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tags" json:"tags,omitempty"`
		VpcSecurityGroupIDs *[]string `tfsdk:"vpc_security_group_i_ds" json:"vpcSecurityGroupIDs,omitempty"`
		VpcSubnetIDs        *[]string `tfsdk:"vpc_subnet_i_ds" json:"vpcSubnetIDs,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *RdsServicesK8SAwsDbproxyV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_rds_services_k8s_aws_db_proxy_v1alpha1_manifest"
}

func (r *RdsServicesK8SAwsDbproxyV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "DBProxy is the Schema for the DBProxies API",
		MarkdownDescription: "DBProxy is the Schema for the DBProxies API",
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
				Description:         "DBProxySpec defines the desired state of DBProxy.The data structure representing a proxy managed by the RDS Proxy.This data type is used as a response element in the DescribeDBProxies action.",
				MarkdownDescription: "DBProxySpec defines the desired state of DBProxy.The data structure representing a proxy managed by the RDS Proxy.This data type is used as a response element in the DescribeDBProxies action.",
				Attributes: map[string]schema.Attribute{
					"auth": schema.ListNestedAttribute{
						Description:         "The authorization mechanism that the proxy uses.",
						MarkdownDescription: "The authorization mechanism that the proxy uses.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"auth_scheme": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"client_password_auth_type": schema.StringAttribute{
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

								"iam_auth": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_arn": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"user_name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"debug_logging": schema.BoolAttribute{
						Description:         "Whether the proxy includes detailed information about SQL statements in itslogs. This information helps you to debug issues involving SQL behavior orthe performance and scalability of the proxy connections. The debug informationincludes the text of SQL statements that you submit through the proxy. Thus,only enable this setting when needed for debugging, and only when you havesecurity measures in place to safeguard any sensitive information that appearsin the logs.",
						MarkdownDescription: "Whether the proxy includes detailed information about SQL statements in itslogs. This information helps you to debug issues involving SQL behavior orthe performance and scalability of the proxy connections. The debug informationincludes the text of SQL statements that you submit through the proxy. Thus,only enable this setting when needed for debugging, and only when you havesecurity measures in place to safeguard any sensitive information that appearsin the logs.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine_family": schema.StringAttribute{
						Description:         "The kinds of databases that the proxy can connect to. This value determineswhich database network protocol the proxy recognizes when it interprets networktraffic to and from the database. For Aurora MySQL, RDS for MariaDB, andRDS for MySQL databases, specify MYSQL. For Aurora PostgreSQL and RDS forPostgreSQL databases, specify POSTGRESQL. For RDS for Microsoft SQL Server,specify SQLSERVER.",
						MarkdownDescription: "The kinds of databases that the proxy can connect to. This value determineswhich database network protocol the proxy recognizes when it interprets networktraffic to and from the database. For Aurora MySQL, RDS for MariaDB, andRDS for MySQL databases, specify MYSQL. For Aurora PostgreSQL and RDS forPostgreSQL databases, specify POSTGRESQL. For RDS for Microsoft SQL Server,specify SQLSERVER.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"idle_client_timeout": schema.Int64Attribute{
						Description:         "The number of seconds that a connection to the proxy can be inactive beforethe proxy disconnects it. You can set this value higher or lower than theconnection timeout limit for the associated database.",
						MarkdownDescription: "The number of seconds that a connection to the proxy can be inactive beforethe proxy disconnects it. You can set this value higher or lower than theconnection timeout limit for the associated database.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "The identifier for the proxy. This name must be unique for all proxies ownedby your Amazon Web Services account in the specified Amazon Web ServicesRegion. An identifier must begin with a letter and must contain only ASCIIletters, digits, and hyphens; it can't end with a hyphen or contain two consecutivehyphens.",
						MarkdownDescription: "The identifier for the proxy. This name must be unique for all proxies ownedby your Amazon Web Services account in the specified Amazon Web ServicesRegion. An identifier must begin with a letter and must contain only ASCIIletters, digits, and hyphens; it can't end with a hyphen or contain two consecutivehyphens.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"require_tls": schema.BoolAttribute{
						Description:         "A Boolean parameter that specifies whether Transport Layer Security (TLS)encryption is required for connections to the proxy. By enabling this setting,you can enforce encrypted TLS connections to the proxy.",
						MarkdownDescription: "A Boolean parameter that specifies whether Transport Layer Security (TLS)encryption is required for connections to the proxy. By enabling this setting,you can enforce encrypted TLS connections to the proxy.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"role_arn": schema.StringAttribute{
						Description:         "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to accesssecrets in Amazon Web Services Secrets Manager.",
						MarkdownDescription: "The Amazon Resource Name (ARN) of the IAM role that the proxy uses to accesssecrets in Amazon Web Services Secrets Manager.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tags": schema.ListNestedAttribute{
						Description:         "An optional set of key-value pairs to associate arbitrary data of your choosingwith the proxy.",
						MarkdownDescription: "An optional set of key-value pairs to associate arbitrary data of your choosingwith the proxy.",
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

					"vpc_security_group_i_ds": schema.ListAttribute{
						Description:         "One or more VPC security group IDs to associate with the new proxy.",
						MarkdownDescription: "One or more VPC security group IDs to associate with the new proxy.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vpc_subnet_i_ds": schema.ListAttribute{
						Description:         "One or more VPC subnet IDs to associate with the new proxy.",
						MarkdownDescription: "One or more VPC subnet IDs to associate with the new proxy.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
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

func (r *RdsServicesK8SAwsDbproxyV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_rds_services_k8s_aws_db_proxy_v1alpha1_manifest")

	var model RdsServicesK8SAwsDbproxyV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("rds.services.k8s.aws/v1alpha1")
	model.Kind = pointer.String("DBProxy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
