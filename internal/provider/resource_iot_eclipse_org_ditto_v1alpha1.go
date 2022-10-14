/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"gopkg.in/yaml.v3"
	"time"
)

type IotEclipseOrgDittoV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*IotEclipseOrgDittoV1Alpha1Resource)(nil)
)

type IotEclipseOrgDittoV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type IotEclipseOrgDittoV1Alpha1GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		CreateDefaultUser *bool `tfsdk:"create_default_user" yaml:"createDefaultUser,omitempty"`

		Devops *struct {
			Expose *bool `tfsdk:"expose" yaml:"expose,omitempty"`

			Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

			Password *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"password" yaml:"password,omitempty"`

			StatusPassword *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"status_password" yaml:"statusPassword,omitempty"`
		} `tfsdk:"devops" yaml:"devops,omitempty"`

		DisableInfraProxy *bool `tfsdk:"disable_infra_proxy" yaml:"disableInfraProxy,omitempty"`

		DisableWelcomePage *bool `tfsdk:"disable_welcome_page" yaml:"disableWelcomePage,omitempty"`

		Ingress *struct {
			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			ClassName *string `tfsdk:"class_name" yaml:"className,omitempty"`

			Host *string `tfsdk:"host" yaml:"host,omitempty"`
		} `tfsdk:"ingress" yaml:"ingress,omitempty"`

		Kafka *struct {
			ConsumerThrottlingLimit *int64 `tfsdk:"consumer_throttling_limit" yaml:"consumerThrottlingLimit,omitempty"`
		} `tfsdk:"kafka" yaml:"kafka,omitempty"`

		Keycloak *struct {
			ClientId *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"client_id" yaml:"clientId,omitempty"`

			ClientSecret *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

			Description *string `tfsdk:"description" yaml:"description,omitempty"`

			DisableProxy *bool `tfsdk:"disable_proxy" yaml:"disableProxy,omitempty"`

			Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

			Label *string `tfsdk:"label" yaml:"label,omitempty"`

			Realm *string `tfsdk:"realm" yaml:"realm,omitempty"`

			RedirectUrl *string `tfsdk:"redirect_url" yaml:"redirectUrl,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"keycloak" yaml:"keycloak,omitempty"`

		Metrics *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
		} `tfsdk:"metrics" yaml:"metrics,omitempty"`

		MongoDb *struct {
			Database *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"database" yaml:"database,omitempty"`

			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Password *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"password" yaml:"password,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Username *struct {
				ConfigMap *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map" yaml:"configMap,omitempty"`

				Secret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"username" yaml:"username,omitempty"`
		} `tfsdk:"mongo_db" yaml:"mongoDb,omitempty"`

		Oauth *struct {
			Issuers *map[string]string `tfsdk:"issuers" yaml:"issuers,omitempty"`
		} `tfsdk:"oauth" yaml:"oauth,omitempty"`

		OpenApi *struct {
			ServerLabel *string `tfsdk:"server_label" yaml:"serverLabel,omitempty"`
		} `tfsdk:"open_api" yaml:"openApi,omitempty"`

		PullPolicy *string `tfsdk:"pull_policy" yaml:"pullPolicy,omitempty"`

		Registry *string `tfsdk:"registry" yaml:"registry,omitempty"`

		Services *struct {
			Concierge *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" yaml:"additionalProperties,omitempty"`

				AppLogLevel *string `tfsdk:"app_log_level" yaml:"appLogLevel,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				RootLogLevel *string `tfsdk:"root_log_level" yaml:"rootLogLevel,omitempty"`
			} `tfsdk:"concierge" yaml:"concierge,omitempty"`

			Connectivity *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" yaml:"additionalProperties,omitempty"`

				AppLogLevel *string `tfsdk:"app_log_level" yaml:"appLogLevel,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				RootLogLevel *string `tfsdk:"root_log_level" yaml:"rootLogLevel,omitempty"`
			} `tfsdk:"connectivity" yaml:"connectivity,omitempty"`

			Gateway *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" yaml:"additionalProperties,omitempty"`

				AppLogLevel *string `tfsdk:"app_log_level" yaml:"appLogLevel,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				RootLogLevel *string `tfsdk:"root_log_level" yaml:"rootLogLevel,omitempty"`
			} `tfsdk:"gateway" yaml:"gateway,omitempty"`

			Policies *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" yaml:"additionalProperties,omitempty"`

				AppLogLevel *string `tfsdk:"app_log_level" yaml:"appLogLevel,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				RootLogLevel *string `tfsdk:"root_log_level" yaml:"rootLogLevel,omitempty"`
			} `tfsdk:"policies" yaml:"policies,omitempty"`

			Things *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" yaml:"additionalProperties,omitempty"`

				AppLogLevel *string `tfsdk:"app_log_level" yaml:"appLogLevel,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				RootLogLevel *string `tfsdk:"root_log_level" yaml:"rootLogLevel,omitempty"`
			} `tfsdk:"things" yaml:"things,omitempty"`

			ThingsSearch *struct {
				AdditionalProperties *map[string]string `tfsdk:"additional_properties" yaml:"additionalProperties,omitempty"`

				AppLogLevel *string `tfsdk:"app_log_level" yaml:"appLogLevel,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				RootLogLevel *string `tfsdk:"root_log_level" yaml:"rootLogLevel,omitempty"`
			} `tfsdk:"things_search" yaml:"thingsSearch,omitempty"`
		} `tfsdk:"services" yaml:"services,omitempty"`

		SwaggerUi *struct {
			Disable *bool `tfsdk:"disable" yaml:"disable,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`
		} `tfsdk:"swagger_ui" yaml:"swaggerUi,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewIotEclipseOrgDittoV1Alpha1Resource() resource.Resource {
	return &IotEclipseOrgDittoV1Alpha1Resource{}
}

func (r *IotEclipseOrgDittoV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iot_eclipse_org_ditto_v1alpha1"
}

func (r *IotEclipseOrgDittoV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Auto-generated derived type for DittoSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for DittoSpec via 'CustomResource'",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Description:         "The timestamp of the last change to this resource.",
				MarkdownDescription: "The timestamp of the last change to this resource.",
				Type:                types.Int64Type,
				Computed:            true,
				Optional:            false,
			},

			"yaml": {
				Description:         "The generated manifest in YAML format.",
				MarkdownDescription: "The generated manifest in YAML format.",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"metadata": {
				Description:         "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				MarkdownDescription: "Data that helps uniquely identify this object. See https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#metadata for more details.",
				Required:            true,
				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
					"name": {
						Description:         "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						MarkdownDescription: "Unique identifier for this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names for more details.",
						Type:                types.StringType,
						Required:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.NameValidator(),
						},
					},

					"namespace": {
						Description:         "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						MarkdownDescription: "Namespaces provides a mechanism for isolating groups of resources within a single cluster. See https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ for more details.",
						Type:                types.StringType,
						Optional:            true,
					},

					"labels": {
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.LabelValidator(),
						},
					},
					"annotations": {
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						Type:                types.MapType{ElemType: types.StringType},
						Optional:            true,
						Validators: []tfsdk.AttributeValidator{
							validators.AnnotationValidator(),
						},
					},
				}),
			},

			"api_version": {
				Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"kind": {
				Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
				Type:                types.StringType,
				Computed:            true,
				Optional:            false,
			},

			"spec": {
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"create_default_user": {
						Description:         "Create the default 'ditto' user when initially deploying.This has no effect when using OAuth2.",
						MarkdownDescription: "Create the default 'ditto' user when initially deploying.This has no effect when using OAuth2.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"devops": {
						Description:         "Devops endpoint",
						MarkdownDescription: "Devops endpoint",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"expose": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"insecure": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"status_password": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_infra_proxy": {
						Description:         "Don't expose infra endpoints",
						MarkdownDescription: "Don't expose infra endpoints",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_welcome_page": {
						Description:         "Allow disabling the welcome page",
						MarkdownDescription: "Allow disabling the welcome page",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": {
						Description:         "Configure ingress optionsIf the field is missing, no ingress resource is being created.",
						MarkdownDescription: "Configure ingress optionsIf the field is missing, no ingress resource is being created.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotations": {
								Description:         "Annotations which should be applied to the ingress resources.The annotations will be set to the resource, not merged. All changes done on the ingress resource itself will be overridden.If no annotations are configured, reasonable defaults will be used instead. You can prevent this by setting a single dummy annotation.",
								MarkdownDescription: "Annotations which should be applied to the ingress resources.The annotations will be set to the resource, not merged. All changes done on the ingress resource itself will be overridden.If no annotations are configured, reasonable defaults will be used instead. You can prevent this by setting a single dummy annotation.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"class_name": {
								Description:         "The optional ingress class name.",
								MarkdownDescription: "The optional ingress class name.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host": {
								Description:         "The host of the ingress resource.This is required if the ingress resource should be created by the operator",
								MarkdownDescription: "The host of the ingress resource.This is required if the ingress resource should be created by the operator",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kafka": {
						Description:         "Kafka options",
						MarkdownDescription: "Kafka options",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"consumer_throttling_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"keycloak": {
						Description:         "Enable and configure keycloak integration.",
						MarkdownDescription: "Enable and configure keycloak integration.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"client_id": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"client_secret": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"description": {
								Description:         "Description of this login option.",
								MarkdownDescription: "Description of this login option.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_proxy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"groups": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"label": {
								Description:         "Label when referencing this login option.",
								MarkdownDescription: "Label when referencing this login option.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"realm": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"redirect_url": {
								Description:         "Allow overriding the redirect URL.",
								MarkdownDescription: "Allow overriding the redirect URL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"url": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"metrics": {
						Description:         "Metrics configuration",
						MarkdownDescription: "Metrics configuration",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enable metrics integration",
								MarkdownDescription: "Enable metrics integration",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mongo_db": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"database": {
								Description:         "The optional database name used to connect, defaults to 'ditto'.",
								MarkdownDescription: "The optional database name used to connect, defaults to 'ditto'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host": {
								Description:         "The hostname of the MongoDB instance.",
								MarkdownDescription: "The hostname of the MongoDB instance.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password": {
								Description:         "The password used to connect to the MongoDB instance.",
								MarkdownDescription: "The password used to connect to the MongoDB instance.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "The port name of the MongoDB instance.",
								MarkdownDescription: "The port name of the MongoDB instance.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),
								},
							},

							"username": {
								Description:         "The username used to connect to the MongoDB instance.",
								MarkdownDescription: "The username used to connect to the MongoDB instance.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "Selects a key from a ConfigMap.",
										MarkdownDescription: "Selects a key from a ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": {
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"name": {
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"optional": {
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"oauth": {
						Description:         "Provide additional OAuth configuration",
						MarkdownDescription: "Provide additional OAuth configuration",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"issuers": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"open_api": {
						Description:         "Influence some options of the hosted OpenAPI spec.",
						MarkdownDescription: "Influence some options of the hosted OpenAPI spec.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"server_label": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pull_policy": {
						Description:         "Override the imagePullPolicyBy default this will use Always if the image version is ':latest' and IfNotPresent otherwise",
						MarkdownDescription: "Override the imagePullPolicyBy default this will use Always if the image version is ':latest' and IfNotPresent otherwise",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"registry": {
						Description:         "Allow to override the Ditto container registry",
						MarkdownDescription: "Allow to override the Ditto container registry",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"services": {
						Description:         "Services configuration",
						MarkdownDescription: "Services configuration",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"concierge": {
								Description:         "The concierge service",
								MarkdownDescription: "The concierge service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_properties": {
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"app_log_level": {
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": {
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": {
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"resources": {
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_log_level": {
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"connectivity": {
								Description:         "The connectivity service",
								MarkdownDescription: "The connectivity service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_properties": {
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"app_log_level": {
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": {
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": {
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"resources": {
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_log_level": {
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gateway": {
								Description:         "The gateway service",
								MarkdownDescription: "The gateway service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_properties": {
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"app_log_level": {
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": {
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": {
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"resources": {
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_log_level": {
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"policies": {
								Description:         "The policies service",
								MarkdownDescription: "The policies service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_properties": {
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"app_log_level": {
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": {
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": {
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"resources": {
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_log_level": {
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"things": {
								Description:         "The things service",
								MarkdownDescription: "The things service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_properties": {
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"app_log_level": {
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": {
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": {
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"resources": {
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_log_level": {
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"things_search": {
								Description:         "The things search service",
								MarkdownDescription: "The things search service",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_properties": {
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"app_log_level": {
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"log_level": {
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},

									"replicas": {
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"resources": {
										Description:         "Service resource limits",
										MarkdownDescription: "Service resource limits",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_log_level": {
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("trace", "debug", "info", "warn", "error"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"swagger_ui": {
						Description:         "Influence some options of the hosted SwaggerUI.",
						MarkdownDescription: "Influence some options of the hosted SwaggerUI.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"disable": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "Allow to override the Ditto image version.",
						MarkdownDescription: "Allow to override the Ditto image version.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *IotEclipseOrgDittoV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_iot_eclipse_org_ditto_v1alpha1")

	var state IotEclipseOrgDittoV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel IotEclipseOrgDittoV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("iot.eclipse.org/v1alpha1")
	goModel.Kind = utilities.Ptr("Ditto")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *IotEclipseOrgDittoV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_iot_eclipse_org_ditto_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *IotEclipseOrgDittoV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_iot_eclipse_org_ditto_v1alpha1")

	var state IotEclipseOrgDittoV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel IotEclipseOrgDittoV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("iot.eclipse.org/v1alpha1")
	goModel.Kind = utilities.Ptr("Ditto")

	state.Id = types.Int64{Value: time.Now().UnixNano()}
	state.ApiVersion = types.String{Value: *goModel.ApiVersion}
	state.Kind = types.String{Value: *goModel.Kind}

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.String{Value: string(marshal)}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *IotEclipseOrgDittoV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_iot_eclipse_org_ditto_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
