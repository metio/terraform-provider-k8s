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
	_ datasource.DataSource = &K8SMariadbComConnectionV1Alpha1Manifest{}
)

func NewK8SMariadbComConnectionV1Alpha1Manifest() datasource.DataSource {
	return &K8SMariadbComConnectionV1Alpha1Manifest{}
}

type K8SMariadbComConnectionV1Alpha1Manifest struct{}

type K8SMariadbComConnectionV1Alpha1ManifestData struct {
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
		Database    *string `tfsdk:"database" json:"database,omitempty"`
		HealthCheck *struct {
			Interval      *string `tfsdk:"interval" json:"interval,omitempty"`
			RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
		} `tfsdk:"health_check" json:"healthCheck,omitempty"`
		Host       *string `tfsdk:"host" json:"host,omitempty"`
		MariaDbRef *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			WaitForIt       *bool   `tfsdk:"wait_for_it" json:"waitForIt,omitempty"`
		} `tfsdk:"maria_db_ref" json:"mariaDbRef,omitempty"`
		MaxScaleRef *struct {
			ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
			FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
			Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
			ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
			Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"max_scale_ref" json:"maxScaleRef,omitempty"`
		Params               *map[string]string `tfsdk:"params" json:"params,omitempty"`
		PasswordSecretKeyRef *struct {
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
		} `tfsdk:"password_secret_key_ref" json:"passwordSecretKeyRef,omitempty"`
		Port           *int64  `tfsdk:"port" json:"port,omitempty"`
		SecretName     *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		SecretTemplate *struct {
			DatabaseKey *string `tfsdk:"database_key" json:"databaseKey,omitempty"`
			Format      *string `tfsdk:"format" json:"format,omitempty"`
			HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
			Key         *string `tfsdk:"key" json:"key,omitempty"`
			Metadata    *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
			PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
			UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
		} `tfsdk:"secret_template" json:"secretTemplate,omitempty"`
		ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
		Username    *string `tfsdk:"username" json:"username,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SMariadbComConnectionV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_mariadb_com_connection_v1alpha1_manifest"
}

func (r *K8SMariadbComConnectionV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Connection is the Schema for the connections API. It is used to configure connection strings for the applications connecting to MariaDB.",
		MarkdownDescription: "Connection is the Schema for the connections API. It is used to configure connection strings for the applications connecting to MariaDB.",
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
				Description:         "ConnectionSpec defines the desired state of Connection",
				MarkdownDescription: "ConnectionSpec defines the desired state of Connection",
				Attributes: map[string]schema.Attribute{
					"database": schema.StringAttribute{
						Description:         "Database to use when configuring the Connection.",
						MarkdownDescription: "Database to use when configuring the Connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"health_check": schema.SingleNestedAttribute{
						Description:         "HealthCheck to be used in the Connection.",
						MarkdownDescription: "HealthCheck to be used in the Connection.",
						Attributes: map[string]schema.Attribute{
							"interval": schema.StringAttribute{
								Description:         "Interval used to perform health checks.",
								MarkdownDescription: "Interval used to perform health checks.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retry_interval": schema.StringAttribute{
								Description:         "RetryInterval is the intervañ used to perform health check retries.",
								MarkdownDescription: "RetryInterval is the intervañ used to perform health check retries.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"host": schema.StringAttribute{
						Description:         "Host to connect to. If not provided, it defaults to the MariaDB host or to the MaxScale host.",
						MarkdownDescription: "Host to connect to. If not provided, it defaults to the MariaDB host or to the MaxScale host.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"maria_db_ref": schema.SingleNestedAttribute{
						Description:         "MariaDBRef is a reference to the MariaDB to connect to. Either MariaDBRef or MaxScaleRef must be provided.",
						MarkdownDescription: "MariaDBRef is a reference to the MariaDB to connect to. Either MariaDBRef or MaxScaleRef must be provided.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_scale_ref": schema.SingleNestedAttribute{
						Description:         "MaxScaleRef is a reference to the MaxScale to connect to. Either MariaDBRef or MaxScaleRef must be provided.",
						MarkdownDescription: "MaxScaleRef is a reference to the MaxScale to connect to. Either MariaDBRef or MaxScaleRef must be provided.",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this stringshould contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2].For example, if the object reference is to a container within a pod, this would take on a value like:'spec.containers{name}' (where 'name' refers to the name of the container that triggeredthe event) or if no container name is specified 'spec.containers[2]' (container withindex 2 in this pod). This syntax is chosen only to have some well-defined way ofreferencing a part of an object.TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"params": schema.MapAttribute{
						Description:         "Params to be used in the Connection.",
						MarkdownDescription: "Params to be used in the Connection.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"password_secret_key_ref": schema.SingleNestedAttribute{
						Description:         "PasswordSecretKeyRef is a reference to the password to use for configuring the Connection.",
						MarkdownDescription: "PasswordSecretKeyRef is a reference to the password to use for configuring the Connection.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "The key of the secret to select from.  Must be a valid secret key.",
								MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"optional": schema.BoolAttribute{
								Description:         "Specify whether the Secret or its key must be defined",
								MarkdownDescription: "Specify whether the Secret or its key must be defined",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"port": schema.Int64Attribute{
						Description:         "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
						MarkdownDescription: "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_name": schema.StringAttribute{
						Description:         "SecretName to be used in the Connection.",
						MarkdownDescription: "SecretName to be used in the Connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"secret_template": schema.SingleNestedAttribute{
						Description:         "SecretTemplate to be used in the Connection.",
						MarkdownDescription: "SecretTemplate to be used in the Connection.",
						Attributes: map[string]schema.Attribute{
							"database_key": schema.StringAttribute{
								Description:         "DatabaseKey to be used in the Secret.",
								MarkdownDescription: "DatabaseKey to be used in the Secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"format": schema.StringAttribute{
								Description:         "Format to be used in the Secret.",
								MarkdownDescription: "Format to be used in the Secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_key": schema.StringAttribute{
								Description:         "HostKey to be used in the Secret.",
								MarkdownDescription: "HostKey to be used in the Secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key": schema.StringAttribute{
								Description:         "Key to be used in the Secret.",
								MarkdownDescription: "Key to be used in the Secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata to be added to the Secret object.",
								MarkdownDescription: "Metadata to be added to the Secret object.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations to be added to children resources.",
										MarkdownDescription: "Annotations to be added to children resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels to be added to children resources.",
										MarkdownDescription: "Labels to be added to children resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"password_key": schema.StringAttribute{
								Description:         "PasswordKey to be used in the Secret.",
								MarkdownDescription: "PasswordKey to be used in the Secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port_key": schema.StringAttribute{
								Description:         "PortKey to be used in the Secret.",
								MarkdownDescription: "PortKey to be used in the Secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"username_key": schema.StringAttribute{
								Description:         "UsernameKey to be used in the Secret.",
								MarkdownDescription: "UsernameKey to be used in the Secret.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_name": schema.StringAttribute{
						Description:         "ServiceName to be used in the Connection.",
						MarkdownDescription: "ServiceName to be used in the Connection.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"username": schema.StringAttribute{
						Description:         "Username to use for configuring the Connection.",
						MarkdownDescription: "Username to use for configuring the Connection.",
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

func (r *K8SMariadbComConnectionV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_mariadb_com_connection_v1alpha1_manifest")

	var model K8SMariadbComConnectionV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.mariadb.com/v1alpha1")
	model.Kind = pointer.String("Connection")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
