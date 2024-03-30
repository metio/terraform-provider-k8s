/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package oracle_db_anthosapis_com_v1alpha1

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
	_ datasource.DataSource = &OracleDbAnthosapisComDatabaseV1Alpha1Manifest{}
)

func NewOracleDbAnthosapisComDatabaseV1Alpha1Manifest() datasource.DataSource {
	return &OracleDbAnthosapisComDatabaseV1Alpha1Manifest{}
}

type OracleDbAnthosapisComDatabaseV1Alpha1Manifest struct{}

type OracleDbAnthosapisComDatabaseV1Alpha1ManifestData struct {
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
		AdminPasswordGsmSecretRef *struct {
			ProjectId *string `tfsdk:"project_id" json:"projectId,omitempty"`
			SecretId  *string `tfsdk:"secret_id" json:"secretId,omitempty"`
			Version   *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"admin_password_gsm_secret_ref" json:"adminPasswordGsmSecretRef,omitempty"`
		Admin_password *string `tfsdk:"admin_password" json:"admin_password,omitempty"`
		Instance       *string `tfsdk:"instance" json:"instance,omitempty"`
		Name           *string `tfsdk:"name" json:"name,omitempty"`
		Users          *[]struct {
			GsmSecretRef *struct {
				ProjectId *string `tfsdk:"project_id" json:"projectId,omitempty"`
				SecretId  *string `tfsdk:"secret_id" json:"secretId,omitempty"`
				Version   *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"gsm_secret_ref" json:"gsmSecretRef,omitempty"`
			Name       *string   `tfsdk:"name" json:"name,omitempty"`
			Password   *string   `tfsdk:"password" json:"password,omitempty"`
			Privileges *[]string `tfsdk:"privileges" json:"privileges,omitempty"`
			SecretRef  *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"users" json:"users,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OracleDbAnthosapisComDatabaseV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_oracle_db_anthosapis_com_database_v1alpha1_manifest"
}

func (r *OracleDbAnthosapisComDatabaseV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Database is the Schema for the databases API.",
		MarkdownDescription: "Database is the Schema for the databases API.",
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
				Description:         "DatabaseSpec defines the desired state of Database.",
				MarkdownDescription: "DatabaseSpec defines the desired state of Database.",
				Attributes: map[string]schema.Attribute{
					"admin_password_gsm_secret_ref": schema.SingleNestedAttribute{
						Description:         "AdminPasswordGsmSecretRef is a reference to the secret object containing sensitive information to pass to config agent. This field is optional, and may be empty if plaintext password is used.",
						MarkdownDescription: "AdminPasswordGsmSecretRef is a reference to the secret object containing sensitive information to pass to config agent. This field is optional, and may be empty if plaintext password is used.",
						Attributes: map[string]schema.Attribute{
							"project_id": schema.StringAttribute{
								Description:         "ProjectId identifies the project where the secret resource is.",
								MarkdownDescription: "ProjectId identifies the project where the secret resource is.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_id": schema.StringAttribute{
								Description:         "SecretId identifies the secret.",
								MarkdownDescription: "SecretId identifies the secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "Version is the version of the secret. If 'latest' is specified, underlying the latest SecretId is used.",
								MarkdownDescription: "Version is the version of the secret. If 'latest' is specified, underlying the latest SecretId is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"admin_password": schema.StringAttribute{
						Description:         "AdminPassword is the password for the sys admin of the database.",
						MarkdownDescription: "AdminPassword is the password for the sys admin of the database.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtLeast(5),
							stringvalidator.LengthAtMost(30),
						},
					},

					"instance": schema.StringAttribute{
						Description:         "Name of the instance that the database belongs to.",
						MarkdownDescription: "Name of the instance that the database belongs to.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name of the database.",
						MarkdownDescription: "Name of the database.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"users": schema.ListNestedAttribute{
						Description:         "Users specifies an optional list of users to be created in this database.",
						MarkdownDescription: "Users specifies an optional list of users to be created in this database.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"gsm_secret_ref": schema.SingleNestedAttribute{
									Description:         "A reference to a GSM secret.",
									MarkdownDescription: "A reference to a GSM secret.",
									Attributes: map[string]schema.Attribute{
										"project_id": schema.StringAttribute{
											Description:         "ProjectId identifies the project where the secret resource is.",
											MarkdownDescription: "ProjectId identifies the project where the secret resource is.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_id": schema.StringAttribute{
											Description:         "SecretId identifies the secret.",
											MarkdownDescription: "SecretId identifies the secret.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"version": schema.StringAttribute{
											Description:         "Version is the version of the secret. If 'latest' is specified, underlying the latest SecretId is used.",
											MarkdownDescription: "Version is the version of the secret. If 'latest' is specified, underlying the latest SecretId is used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the User.",
									MarkdownDescription: "Name of the User.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"password": schema.StringAttribute{
									Description:         "Plaintext password.",
									MarkdownDescription: "Plaintext password.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"privileges": schema.ListAttribute{
									Description:         "Privileges specifies an optional list of privileges to grant to the user.",
									MarkdownDescription: "Privileges specifies an optional list of privileges to grant to the user.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "A reference to a k8s secret.",
									MarkdownDescription: "A reference to a k8s secret.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "name is unique within a namespace to reference a secret resource.",
											MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "namespace defines the space within which the secret name must be unique.",
											MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

func (r *OracleDbAnthosapisComDatabaseV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_oracle_db_anthosapis_com_database_v1alpha1_manifest")

	var model OracleDbAnthosapisComDatabaseV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("oracle.db.anthosapis.com/v1alpha1")
	model.Kind = pointer.String("Database")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
