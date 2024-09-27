/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_mariadb_com_v1alpha1

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
	_ datasource.DataSource = &K8SMariadbComUserV1Alpha1Manifest{}
)

func NewK8SMariadbComUserV1Alpha1Manifest() datasource.DataSource {
	return &K8SMariadbComUserV1Alpha1Manifest{}
}

type K8SMariadbComUserV1Alpha1Manifest struct{}

type K8SMariadbComUserV1Alpha1ManifestData struct {
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
		CleanupPolicy *string `tfsdk:"cleanup_policy" json:"cleanupPolicy,omitempty"`
		Host          *string `tfsdk:"host" json:"host,omitempty"`
		MariaDbRef    *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			WaitForIt *bool   `tfsdk:"wait_for_it" json:"waitForIt,omitempty"`
		} `tfsdk:"maria_db_ref" json:"mariaDbRef,omitempty"`
		MaxUserConnections       *int64  `tfsdk:"max_user_connections" json:"maxUserConnections,omitempty"`
		Name                     *string `tfsdk:"name" json:"name,omitempty"`
		PasswordHashSecretKeyRef *struct {
			Key  *string `tfsdk:"key" json:"key,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"password_hash_secret_key_ref" json:"passwordHashSecretKeyRef,omitempty"`
		PasswordPlugin *struct {
			PluginArgSecretKeyRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"plugin_arg_secret_key_ref" json:"pluginArgSecretKeyRef,omitempty"`
			PluginNameSecretKeyRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"plugin_name_secret_key_ref" json:"pluginNameSecretKeyRef,omitempty"`
		} `tfsdk:"password_plugin" json:"passwordPlugin,omitempty"`
		PasswordSecretKeyRef *struct {
			Key  *string `tfsdk:"key" json:"key,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"password_secret_key_ref" json:"passwordSecretKeyRef,omitempty"`
		RequeueInterval *string `tfsdk:"requeue_interval" json:"requeueInterval,omitempty"`
		RetryInterval   *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SMariadbComUserV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_mariadb_com_user_v1alpha1_manifest"
}

func (r *K8SMariadbComUserV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "User is the Schema for the users API. It is used to define grants as if you were running a 'CREATE USER' statement.",
		MarkdownDescription: "User is the Schema for the users API. It is used to define grants as if you were running a 'CREATE USER' statement.",
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
				Description:         "UserSpec defines the desired state of User",
				MarkdownDescription: "UserSpec defines the desired state of User",
				Attributes: map[string]schema.Attribute{
					"cleanup_policy": schema.StringAttribute{
						Description:         "CleanupPolicy defines the behavior for cleaning up a SQL resource.",
						MarkdownDescription: "CleanupPolicy defines the behavior for cleaning up a SQL resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Skip", "Delete"),
						},
					},

					"host": schema.StringAttribute{
						Description:         "Host related to the User.",
						MarkdownDescription: "Host related to the User.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
						},
					},

					"maria_db_ref": schema.SingleNestedAttribute{
						Description:         "MariaDBRef is a reference to a MariaDB object.",
						MarkdownDescription: "MariaDBRef is a reference to a MariaDB object.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wait_for_it": schema.BoolAttribute{
								Description:         "WaitForIt indicates whether the controller using this reference should wait for MariaDB to be ready.",
								MarkdownDescription: "WaitForIt indicates whether the controller using this reference should wait for MariaDB to be ready.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"max_user_connections": schema.Int64Attribute{
						Description:         "MaxUserConnections defines the maximum number of connections that the User can establish.",
						MarkdownDescription: "MaxUserConnections defines the maximum number of connections that the User can establish.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"name": schema.StringAttribute{
						Description:         "Name overrides the default name provided by metadata.name.",
						MarkdownDescription: "Name overrides the default name provided by metadata.name.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(80),
						},
					},

					"password_hash_secret_key_ref": schema.SingleNestedAttribute{
						Description:         "PasswordHashSecretKeyRef is a reference to the password hash to be used by the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password hash.",
						MarkdownDescription: "PasswordHashSecretKeyRef is a reference to the password hash to be used by the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password hash.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

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

					"password_plugin": schema.SingleNestedAttribute{
						Description:         "PasswordPlugin is a reference to the password plugin and arguments to be used by the User.",
						MarkdownDescription: "PasswordPlugin is a reference to the password plugin and arguments to be used by the User.",
						Attributes: map[string]schema.Attribute{
							"plugin_arg_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "PluginArgSecretKeyRef is a reference to the arguments to be provided to the authentication plugin for the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the authentication plugin arguments.",
								MarkdownDescription: "PluginArgSecretKeyRef is a reference to the arguments to be provided to the authentication plugin for the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the authentication plugin arguments.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

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

							"plugin_name_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "PluginNameSecretKeyRef is a reference to the authentication plugin to be used by the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the authentication plugin.",
								MarkdownDescription: "PluginNameSecretKeyRef is a reference to the authentication plugin to be used by the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the authentication plugin.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

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

					"password_secret_key_ref": schema.SingleNestedAttribute{
						Description:         "PasswordSecretKeyRef is a reference to the password to be used by the User. If not provided, the account will be locked and the password will expire. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
						MarkdownDescription: "PasswordSecretKeyRef is a reference to the password to be used by the User. If not provided, the account will be locked and the password will expire. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

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

					"requeue_interval": schema.StringAttribute{
						Description:         "RequeueInterval is used to perform requeue reconciliations.",
						MarkdownDescription: "RequeueInterval is used to perform requeue reconciliations.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"retry_interval": schema.StringAttribute{
						Description:         "RetryInterval is the interval used to perform retries.",
						MarkdownDescription: "RetryInterval is the interval used to perform retries.",
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

func (r *K8SMariadbComUserV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_mariadb_com_user_v1alpha1_manifest")

	var model K8SMariadbComUserV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.mariadb.com/v1alpha1")
	model.Kind = pointer.String("User")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
