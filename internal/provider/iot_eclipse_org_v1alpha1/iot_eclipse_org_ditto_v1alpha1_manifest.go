/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package iot_eclipse_org_v1alpha1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &IotEclipseOrgDittoV1Alpha1Manifest{}
)

func NewIotEclipseOrgDittoV1Alpha1Manifest() datasource.DataSource {
	return &IotEclipseOrgDittoV1Alpha1Manifest{}
}

type IotEclipseOrgDittoV1Alpha1Manifest struct{}

type IotEclipseOrgDittoV1Alpha1ManifestData struct {
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
		CreateDefaultUser *bool `tfsdk:"create_default_user" json:"createDefaultUser,omitempty"`
		Devops            *struct {
			Expose   *bool `tfsdk:"expose" json:"expose,omitempty"`
			Insecure *bool `tfsdk:"insecure" json:"insecure,omitempty"`
			Password *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			StatusPassword *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"status_password" json:"statusPassword,omitempty"`
		} `tfsdk:"devops" json:"devops,omitempty"`
		DisableInfraProxy  *bool `tfsdk:"disable_infra_proxy" json:"disableInfraProxy,omitempty"`
		DisableWelcomePage *bool `tfsdk:"disable_welcome_page" json:"disableWelcomePage,omitempty"`
		Ingress            *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			ClassName   *string            `tfsdk:"class_name" json:"className,omitempty"`
			Host        *string            `tfsdk:"host" json:"host,omitempty"`
		} `tfsdk:"ingress" json:"ingress,omitempty"`
		Kafka *struct {
			ConsumerThrottlingLimit *int64 `tfsdk:"consumer_throttling_limit" json:"consumerThrottlingLimit,omitempty"`
		} `tfsdk:"kafka" json:"kafka,omitempty"`
		Keycloak *struct {
			ClientId *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"client_id" json:"clientId,omitempty"`
			ClientSecret *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
			Description  *string   `tfsdk:"description" json:"description,omitempty"`
			DisableProxy *bool     `tfsdk:"disable_proxy" json:"disableProxy,omitempty"`
			Groups       *[]string `tfsdk:"groups" json:"groups,omitempty"`
			Label        *string   `tfsdk:"label" json:"label,omitempty"`
			Realm        *string   `tfsdk:"realm" json:"realm,omitempty"`
			RedirectUrl  *string   `tfsdk:"redirect_url" json:"redirectUrl,omitempty"`
			Url          *string   `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"keycloak" json:"keycloak,omitempty"`
		Metrics *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		MongoDb *struct {
			Database *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"database" json:"database,omitempty"`
			Host     *string `tfsdk:"host" json:"host,omitempty"`
			Password *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Port     *int64 `tfsdk:"port" json:"port,omitempty"`
			Username *struct {
				ConfigMap *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"mongo_db" json:"mongoDb,omitempty"`
		Oauth *struct {
			Issuers *struct {
				Subjects *[]string `tfsdk:"subjects" json:"subjects,omitempty"`
				Url      *string   `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"issuers" json:"issuers,omitempty"`
		} `tfsdk:"oauth" json:"oauth,omitempty"`
		OpenApi *struct {
			ServerLabel *string `tfsdk:"server_label" json:"serverLabel,omitempty"`
		} `tfsdk:"open_api" json:"openApi,omitempty"`
		PullPolicy *string `tfsdk:"pull_policy" json:"pullPolicy,omitempty"`
		Registry   *string `tfsdk:"registry" json:"registry,omitempty"`
		Services   *struct {
			Concierge *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" json:"additionalProperties,omitempty"`
				AppLogLevel          *string            `tfsdk:"app_log_level" json:"appLogLevel,omitempty"`
				LogLevel             *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
				Replicas             *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources            *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				RootLogLevel *string `tfsdk:"root_log_level" json:"rootLogLevel,omitempty"`
			} `tfsdk:"concierge" json:"concierge,omitempty"`
			Connectivity *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" json:"additionalProperties,omitempty"`
				AppLogLevel          *string            `tfsdk:"app_log_level" json:"appLogLevel,omitempty"`
				LogLevel             *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
				Replicas             *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources            *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				RootLogLevel *string `tfsdk:"root_log_level" json:"rootLogLevel,omitempty"`
			} `tfsdk:"connectivity" json:"connectivity,omitempty"`
			Gateway *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" json:"additionalProperties,omitempty"`
				AppLogLevel          *string            `tfsdk:"app_log_level" json:"appLogLevel,omitempty"`
				LogLevel             *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
				Replicas             *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources            *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				RootLogLevel *string `tfsdk:"root_log_level" json:"rootLogLevel,omitempty"`
			} `tfsdk:"gateway" json:"gateway,omitempty"`
			Policies *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" json:"additionalProperties,omitempty"`
				AppLogLevel          *string            `tfsdk:"app_log_level" json:"appLogLevel,omitempty"`
				LogLevel             *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
				Replicas             *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources            *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				RootLogLevel *string `tfsdk:"root_log_level" json:"rootLogLevel,omitempty"`
			} `tfsdk:"policies" json:"policies,omitempty"`
			Things *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" json:"additionalProperties,omitempty"`
				AppLogLevel          *string            `tfsdk:"app_log_level" json:"appLogLevel,omitempty"`
				LogLevel             *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
				Replicas             *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources            *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				RootLogLevel *string `tfsdk:"root_log_level" json:"rootLogLevel,omitempty"`
			} `tfsdk:"things" json:"things,omitempty"`
			ThingsSearch *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" json:"additionalProperties,omitempty"`
				AppLogLevel          *string            `tfsdk:"app_log_level" json:"appLogLevel,omitempty"`
				LogLevel             *string            `tfsdk:"log_level" json:"logLevel,omitempty"`
				Replicas             *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources            *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				RootLogLevel *string `tfsdk:"root_log_level" json:"rootLogLevel,omitempty"`
			} `tfsdk:"things_search" json:"thingsSearch,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		SwaggerUi *struct {
			Disable *bool   `tfsdk:"disable" json:"disable,omitempty"`
			Image   *string `tfsdk:"image" json:"image,omitempty"`
		} `tfsdk:"swagger_ui" json:"swaggerUi,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *IotEclipseOrgDittoV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_iot_eclipse_org_ditto_v1alpha1_manifest"
}

func (r *IotEclipseOrgDittoV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for DittoSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for DittoSpec via 'CustomResource'",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"create_default_user": schema.BoolAttribute{
						Description:         "Create the default 'ditto' user when initially deploying. This has no effect when using OAuth2.",
						MarkdownDescription: "Create the default 'ditto' user when initially deploying. This has no effect when using OAuth2.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"devops": schema.SingleNestedAttribute{
						Description:         "Devops endpoint",
						MarkdownDescription: "Devops endpoint",
						Attributes: map[string]schema.Attribute{
							"expose": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"insecure": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": schema.StringAttribute{
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

							"status_password": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": schema.StringAttribute{
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

					"disable_infra_proxy": schema.BoolAttribute{
						Description:         "Don't expose infra endpoints",
						MarkdownDescription: "Don't expose infra endpoints",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"disable_welcome_page": schema.BoolAttribute{
						Description:         "Allow disabling the welcome page",
						MarkdownDescription: "Allow disabling the welcome page",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress": schema.SingleNestedAttribute{
						Description:         "Configure ingress options If the field is missing, no ingress resource is being created.",
						MarkdownDescription: "Configure ingress options If the field is missing, no ingress resource is being created.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations which should be applied to the ingress resources. The annotations will be set to the resource, not merged. All changes done on the ingress resource itself will be overridden. If no annotations are configured, reasonable defaults will be used instead. You can prevent this by setting a single dummy annotation.",
								MarkdownDescription: "Annotations which should be applied to the ingress resources. The annotations will be set to the resource, not merged. All changes done on the ingress resource itself will be overridden. If no annotations are configured, reasonable defaults will be used instead. You can prevent this by setting a single dummy annotation.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"class_name": schema.StringAttribute{
								Description:         "The optional ingress class name.",
								MarkdownDescription: "The optional ingress class name.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host": schema.StringAttribute{
								Description:         "The host of the ingress resource. This is required if the ingress resource should be created by the operator",
								MarkdownDescription: "The host of the ingress resource. This is required if the ingress resource should be created by the operator",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka": schema.SingleNestedAttribute{
						Description:         "Kafka options",
						MarkdownDescription: "Kafka options",
						Attributes: map[string]schema.Attribute{
							"consumer_throttling_limit": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"keycloak": schema.SingleNestedAttribute{
						Description:         "Enable and configure keycloak integration.",
						MarkdownDescription: "Enable and configure keycloak integration.",
						Attributes: map[string]schema.Attribute{
							"client_id": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"client_secret": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"description": schema.StringAttribute{
								Description:         "Description of this login option.",
								MarkdownDescription: "Description of this login option.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"disable_proxy": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"groups": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"label": schema.StringAttribute{
								Description:         "Label when referencing this login option.",
								MarkdownDescription: "Label when referencing this login option.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"realm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"redirect_url": schema.StringAttribute{
								Description:         "Allow overriding the redirect URL.",
								MarkdownDescription: "Allow overriding the redirect URL.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": schema.SingleNestedAttribute{
						Description:         "Metrics configuration",
						MarkdownDescription: "Metrics configuration",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enable metrics integration",
								MarkdownDescription: "Enable metrics integration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"mongo_db": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"database": schema.SingleNestedAttribute{
								Description:         "The optional database name used to connect, defaults to 'ditto'.",
								MarkdownDescription: "The optional database name used to connect, defaults to 'ditto'.",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": schema.StringAttribute{
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

							"host": schema.StringAttribute{
								Description:         "The hostname of the MongoDB instance.",
								MarkdownDescription: "The hostname of the MongoDB instance.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password": schema.SingleNestedAttribute{
								Description:         "The password used to connect to the MongoDB instance.",
								MarkdownDescription: "The password used to connect to the MongoDB instance.",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": schema.StringAttribute{
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

							"port": schema.Int64Attribute{
								Description:         "The port name of the MongoDB instance.",
								MarkdownDescription: "The port name of the MongoDB instance.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"username": schema.SingleNestedAttribute{
								Description:         "The username used to connect to the MongoDB instance.",
								MarkdownDescription: "The username used to connect to the MongoDB instance.",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from. Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": schema.StringAttribute{
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

					"oauth": schema.SingleNestedAttribute{
						Description:         "Provide additional OAuth configuration",
						MarkdownDescription: "Provide additional OAuth configuration",
						Attributes: map[string]schema.Attribute{
							"issuers": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"subjects": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"open_api": schema.SingleNestedAttribute{
						Description:         "Influence some options of the hosted OpenAPI spec.",
						MarkdownDescription: "Influence some options of the hosted OpenAPI spec.",
						Attributes: map[string]schema.Attribute{
							"server_label": schema.StringAttribute{
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

					"pull_policy": schema.StringAttribute{
						Description:         "Override the imagePullPolicy By default this will use Always if the image version is ':latest' and IfNotPresent otherwise",
						MarkdownDescription: "Override the imagePullPolicy By default this will use Always if the image version is ':latest' and IfNotPresent otherwise",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"registry": schema.StringAttribute{
						Description:         "Allow to override the Ditto container registry",
						MarkdownDescription: "Allow to override the Ditto container registry",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"services": schema.SingleNestedAttribute{
						Description:         "Services configuration",
						MarkdownDescription: "Services configuration",
						Attributes: map[string]schema.Attribute{
							"concierge": schema.SingleNestedAttribute{
								Description:         "The concierge service",
								MarkdownDescription: "The concierge service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
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

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"connectivity": schema.SingleNestedAttribute{
								Description:         "The connectivity service",
								MarkdownDescription: "The connectivity service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
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

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway": schema.SingleNestedAttribute{
								Description:         "The gateway service",
								MarkdownDescription: "The gateway service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
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

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"policies": schema.SingleNestedAttribute{
								Description:         "The policies service",
								MarkdownDescription: "The policies service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
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

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"things": schema.SingleNestedAttribute{
								Description:         "The things service",
								MarkdownDescription: "The things service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
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

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"things_search": schema.SingleNestedAttribute{
								Description:         "The things search service",
								MarkdownDescription: "The things search service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties. Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
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

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
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

					"swagger_ui": schema.SingleNestedAttribute{
						Description:         "Influence some options of the hosted SwaggerUI.",
						MarkdownDescription: "Influence some options of the hosted SwaggerUI.",
						Attributes: map[string]schema.Attribute{
							"disable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
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

					"version": schema.StringAttribute{
						Description:         "Allow to override the Ditto image version.",
						MarkdownDescription: "Allow to override the Ditto image version.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *IotEclipseOrgDittoV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_iot_eclipse_org_ditto_v1alpha1_manifest")

	var model IotEclipseOrgDittoV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("iot.eclipse.org/v1alpha1")
	model.Kind = pointer.String("Ditto")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
