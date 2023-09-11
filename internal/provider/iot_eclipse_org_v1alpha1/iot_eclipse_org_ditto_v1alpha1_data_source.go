/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package iot_eclipse_org_v1alpha1

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
	_ datasource.DataSource              = &IotEclipseOrgDittoV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &IotEclipseOrgDittoV1Alpha1DataSource{}
)

func NewIotEclipseOrgDittoV1Alpha1DataSource() datasource.DataSource {
	return &IotEclipseOrgDittoV1Alpha1DataSource{}
}

type IotEclipseOrgDittoV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type IotEclipseOrgDittoV1Alpha1DataSourceData struct {
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

func (r *IotEclipseOrgDittoV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_iot_eclipse_org_ditto_v1alpha1"
}

func (r *IotEclipseOrgDittoV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Auto-generated derived type for DittoSpec via 'CustomResource'",
		MarkdownDescription: "Auto-generated derived type for DittoSpec via 'CustomResource'",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"create_default_user": schema.BoolAttribute{
						Description:         "Create the default 'ditto' user when initially deploying.This has no effect when using OAuth2.",
						MarkdownDescription: "Create the default 'ditto' user when initially deploying.This has no effect when using OAuth2.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"devops": schema.SingleNestedAttribute{
						Description:         "Devops endpoint",
						MarkdownDescription: "Devops endpoint",
						Attributes: map[string]schema.Attribute{
							"expose": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"insecure": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
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

					"disable_infra_proxy": schema.BoolAttribute{
						Description:         "Don't expose infra endpoints",
						MarkdownDescription: "Don't expose infra endpoints",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"disable_welcome_page": schema.BoolAttribute{
						Description:         "Allow disabling the welcome page",
						MarkdownDescription: "Allow disabling the welcome page",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"ingress": schema.SingleNestedAttribute{
						Description:         "Configure ingress optionsIf the field is missing, no ingress resource is being created.",
						MarkdownDescription: "Configure ingress optionsIf the field is missing, no ingress resource is being created.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations which should be applied to the ingress resources.The annotations will be set to the resource, not merged. All changes done on the ingress resource itself will be overridden.If no annotations are configured, reasonable defaults will be used instead. You can prevent this by setting a single dummy annotation.",
								MarkdownDescription: "Annotations which should be applied to the ingress resources.The annotations will be set to the resource, not merged. All changes done on the ingress resource itself will be overridden.If no annotations are configured, reasonable defaults will be used instead. You can prevent this by setting a single dummy annotation.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"class_name": schema.StringAttribute{
								Description:         "The optional ingress class name.",
								MarkdownDescription: "The optional ingress class name.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"host": schema.StringAttribute{
								Description:         "The host of the ingress resource.This is required if the ingress resource should be created by the operator",
								MarkdownDescription: "The host of the ingress resource.This is required if the ingress resource should be created by the operator",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"kafka": schema.SingleNestedAttribute{
						Description:         "Kafka options",
						MarkdownDescription: "Kafka options",
						Attributes: map[string]schema.Attribute{
							"consumer_throttling_limit": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"description": schema.StringAttribute{
								Description:         "Description of this login option.",
								MarkdownDescription: "Description of this login option.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"disable_proxy": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"groups": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"label": schema.StringAttribute{
								Description:         "Label when referencing this login option.",
								MarkdownDescription: "Label when referencing this login option.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"realm": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"redirect_url": schema.StringAttribute{
								Description:         "Allow overriding the redirect URL.",
								MarkdownDescription: "Allow overriding the redirect URL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"metrics": schema.SingleNestedAttribute{
						Description:         "Metrics configuration",
						MarkdownDescription: "Metrics configuration",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enable metrics integration",
								MarkdownDescription: "Enable metrics integration",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
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
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"host": schema.StringAttribute{
								Description:         "The hostname of the MongoDB instance.",
								MarkdownDescription: "The hostname of the MongoDB instance.",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"port": schema.Int64Attribute{
								Description:         "The port name of the MongoDB instance.",
								MarkdownDescription: "The port name of the MongoDB instance.",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the ConfigMap or its key must be defined",
												MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "SecretKeySelector selects a key of a Secret.",
										MarkdownDescription: "SecretKeySelector selects a key of a Secret.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "Specify whether the Secret or its key must be defined",
												MarkdownDescription: "Specify whether the Secret or its key must be defined",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"value": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
										Optional:            false,
										Computed:            true,
									},

									"url": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
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

					"open_api": schema.SingleNestedAttribute{
						Description:         "Influence some options of the hosted OpenAPI spec.",
						MarkdownDescription: "Influence some options of the hosted OpenAPI spec.",
						Attributes: map[string]schema.Attribute{
							"server_label": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"pull_policy": schema.StringAttribute{
						Description:         "Override the imagePullPolicyBy default this will use Always if the image version is ':latest' and IfNotPresent otherwise",
						MarkdownDescription: "Override the imagePullPolicyBy default this will use Always if the image version is ':latest' and IfNotPresent otherwise",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"registry": schema.StringAttribute{
						Description:         "Allow to override the Ditto container registry",
						MarkdownDescription: "Allow to override the Ditto container registry",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"connectivity": schema.SingleNestedAttribute{
								Description:         "The connectivity service",
								MarkdownDescription: "The connectivity service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"gateway": schema.SingleNestedAttribute{
								Description:         "The gateway service",
								MarkdownDescription: "The gateway service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"policies": schema.SingleNestedAttribute{
								Description:         "The policies service",
								MarkdownDescription: "The policies service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"things": schema.SingleNestedAttribute{
								Description:         "The things service",
								MarkdownDescription: "The things service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"things_search": schema.SingleNestedAttribute{
								Description:         "The things search service",
								MarkdownDescription: "The things search service",
								Attributes: map[string]schema.Attribute{
									"additional_properties": schema.MapAttribute{
										Description:         "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										MarkdownDescription: "Additional system properties, which will be appended to the list of system properties.Note: Setting arbitrary system properties may break the deployment and may also not be compatible with future versions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"app_log_level": schema.StringAttribute{
										Description:         "Allow configuring the application log level.",
										MarkdownDescription: "Allow configuring the application log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"log_level": schema.StringAttribute{
										Description:         "Allow configuring all log levels.",
										MarkdownDescription: "Allow configuring all log levels.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"replicas": schema.Int64Attribute{
										Description:         "Number of replicas. Defaults to one.",
										MarkdownDescription: "Number of replicas. Defaults to one.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"root_log_level": schema.StringAttribute{
										Description:         "Allow configuring the root log level.",
										MarkdownDescription: "Allow configuring the root log level.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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

					"swagger_ui": schema.SingleNestedAttribute{
						Description:         "Influence some options of the hosted SwaggerUI.",
						MarkdownDescription: "Influence some options of the hosted SwaggerUI.",
						Attributes: map[string]schema.Attribute{
							"disable": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"version": schema.StringAttribute{
						Description:         "Allow to override the Ditto image version.",
						MarkdownDescription: "Allow to override the Ditto image version.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *IotEclipseOrgDittoV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *IotEclipseOrgDittoV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_iot_eclipse_org_ditto_v1alpha1")

	var data IotEclipseOrgDittoV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "iot.eclipse.org", Version: "v1alpha1", Resource: "dittos"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse IotEclipseOrgDittoV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("iot.eclipse.org/v1alpha1")
	data.Kind = pointer.String("Ditto")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
