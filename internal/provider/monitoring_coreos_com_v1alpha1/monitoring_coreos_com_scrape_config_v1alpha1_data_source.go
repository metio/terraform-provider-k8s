/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package monitoring_coreos_com_v1alpha1

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
	_ datasource.DataSource              = &MonitoringCoreosComScrapeConfigV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &MonitoringCoreosComScrapeConfigV1Alpha1DataSource{}
)

func NewMonitoringCoreosComScrapeConfigV1Alpha1DataSource() datasource.DataSource {
	return &MonitoringCoreosComScrapeConfigV1Alpha1DataSource{}
}

type MonitoringCoreosComScrapeConfigV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type MonitoringCoreosComScrapeConfigV1Alpha1DataSourceData struct {
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
		Authorization *struct {
			Credentials *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"credentials" json:"credentials,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"authorization" json:"authorization,omitempty"`
		BasicAuth *struct {
			Password *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			Username *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
		ConsulSDConfigs *[]struct {
			Allow_stale   *bool `tfsdk:"allow_stale" json:"allow_stale,omitempty"`
			Authorization *struct {
				Credentials *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			BasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Username *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			Datacenter       *string            `tfsdk:"datacenter" json:"datacenter,omitempty"`
			Enable_http2     *bool              `tfsdk:"enable_http2" json:"enable_http2,omitempty"`
			Follow_redirects *bool              `tfsdk:"follow_redirects" json:"follow_redirects,omitempty"`
			Namespace        *string            `tfsdk:"namespace" json:"namespace,omitempty"`
			No_proxy         *string            `tfsdk:"no_proxy" json:"no_proxy,omitempty"`
			Node_meta        *map[string]string `tfsdk:"node_meta" json:"node_meta,omitempty"`
			Oauth2           *struct {
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
				} `tfsdk:"client_id" json:"clientId,omitempty"`
				ClientSecret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
				EndpointParams *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
				Scopes         *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
				TokenUrl       *string            `tfsdk:"token_url" json:"tokenUrl,omitempty"`
			} `tfsdk:"oauth2" json:"oauth2,omitempty"`
			Partition            *string `tfsdk:"partition" json:"partition,omitempty"`
			Proxy_connect_header *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"proxy_connect_header" json:"proxy_connect_header,omitempty"`
			Proxy_from_environment *bool     `tfsdk:"proxy_from_environment" json:"proxy_from_environment,omitempty"`
			Proxy_url              *string   `tfsdk:"proxy_url" json:"proxy_url,omitempty"`
			Refresh_interval       *string   `tfsdk:"refresh_interval" json:"refresh_interval,omitempty"`
			Scheme                 *string   `tfsdk:"scheme" json:"scheme,omitempty"`
			Server                 *string   `tfsdk:"server" json:"server,omitempty"`
			Services               *[]string `tfsdk:"services" json:"services,omitempty"`
			Tag_separator          *string   `tfsdk:"tag_separator" json:"tag_separator,omitempty"`
			Tags                   *[]string `tfsdk:"tags" json:"tags,omitempty"`
			TlsConfig              *struct {
				Ca *struct {
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
				} `tfsdk:"ca" json:"ca,omitempty"`
				Cert *struct {
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
				} `tfsdk:"cert" json:"cert,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				KeySecret          *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"key_secret" json:"keySecret,omitempty"`
				ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
			} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			TokenRef *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"token_ref" json:"tokenRef,omitempty"`
		} `tfsdk:"consul_sd_configs" json:"consulSDConfigs,omitempty"`
		FileSDConfigs *[]struct {
			Files           *[]string `tfsdk:"files" json:"files,omitempty"`
			RefreshInterval *string   `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
		} `tfsdk:"file_sd_configs" json:"fileSDConfigs,omitempty"`
		HonorLabels     *bool `tfsdk:"honor_labels" json:"honorLabels,omitempty"`
		HonorTimestamps *bool `tfsdk:"honor_timestamps" json:"honorTimestamps,omitempty"`
		HttpSDConfigs   *[]struct {
			Authorization *struct {
				Credentials *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			BasicAuth *struct {
				Password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"password" json:"password,omitempty"`
				Username *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			RefreshInterval *string `tfsdk:"refresh_interval" json:"refreshInterval,omitempty"`
			TlsConfig       *struct {
				Ca *struct {
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
				} `tfsdk:"ca" json:"ca,omitempty"`
				Cert *struct {
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
				} `tfsdk:"cert" json:"cert,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				KeySecret          *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"key_secret" json:"keySecret,omitempty"`
				ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
			} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
			Url *string `tfsdk:"url" json:"url,omitempty"`
		} `tfsdk:"http_sd_configs" json:"httpSDConfigs,omitempty"`
		KubernetesSDConfigs *[]struct {
			Role *string `tfsdk:"role" json:"role,omitempty"`
		} `tfsdk:"kubernetes_sd_configs" json:"kubernetesSDConfigs,omitempty"`
		LabelLimit            *int64               `tfsdk:"label_limit" json:"labelLimit,omitempty"`
		LabelNameLengthLimit  *int64               `tfsdk:"label_name_length_limit" json:"labelNameLengthLimit,omitempty"`
		LabelValueLengthLimit *int64               `tfsdk:"label_value_length_limit" json:"labelValueLengthLimit,omitempty"`
		MetricsPath           *string              `tfsdk:"metrics_path" json:"metricsPath,omitempty"`
		Params                *map[string][]string `tfsdk:"params" json:"params,omitempty"`
		Relabelings           *[]struct {
			Action       *string   `tfsdk:"action" json:"action,omitempty"`
			Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
			Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
			Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
			Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
			SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
			TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
		} `tfsdk:"relabelings" json:"relabelings,omitempty"`
		SampleLimit    *int64  `tfsdk:"sample_limit" json:"sampleLimit,omitempty"`
		Scheme         *string `tfsdk:"scheme" json:"scheme,omitempty"`
		ScrapeInterval *string `tfsdk:"scrape_interval" json:"scrapeInterval,omitempty"`
		ScrapeTimeout  *string `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
		StaticConfigs  *[]struct {
			Labels  *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Targets *[]string          `tfsdk:"targets" json:"targets,omitempty"`
		} `tfsdk:"static_configs" json:"staticConfigs,omitempty"`
		TargetLimit *int64 `tfsdk:"target_limit" json:"targetLimit,omitempty"`
		TlsConfig   *struct {
			Ca *struct {
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
			} `tfsdk:"ca" json:"ca,omitempty"`
			Cert *struct {
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
			} `tfsdk:"cert" json:"cert,omitempty"`
			InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			KeySecret          *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"key_secret" json:"keySecret,omitempty"`
			ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
		} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MonitoringCoreosComScrapeConfigV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_monitoring_coreos_com_scrape_config_v1alpha1"
}

func (r *MonitoringCoreosComScrapeConfigV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ScrapeConfig defines a namespaced Prometheus scrape_config to be aggregated across multiple namespaces into the Prometheus configuration.",
		MarkdownDescription: "ScrapeConfig defines a namespaced Prometheus scrape_config to be aggregated across multiple namespaces into the Prometheus configuration.",
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
				Description:         "ScrapeConfigSpec is a specification of the desired configuration for a scrape configuration.",
				MarkdownDescription: "ScrapeConfigSpec is a specification of the desired configuration for a scrape configuration.",
				Attributes: map[string]schema.Attribute{
					"authorization": schema.SingleNestedAttribute{
						Description:         "Authorization header to use on every scrape request.",
						MarkdownDescription: "Authorization header to use on every scrape request.",
						Attributes: map[string]schema.Attribute{
							"credentials": schema.SingleNestedAttribute{
								Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
								MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

							"type": schema.StringAttribute{
								Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
								MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"basic_auth": schema.SingleNestedAttribute{
						Description:         "BasicAuth information to use on every scrape request.",
						MarkdownDescription: "BasicAuth information to use on every scrape request.",
						Attributes: map[string]schema.Attribute{
							"password": schema.SingleNestedAttribute{
								Description:         "The secret in the service monitor namespace that contains the password for authentication.",
								MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

							"username": schema.SingleNestedAttribute{
								Description:         "The secret in the service monitor namespace that contains the username for authentication.",
								MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"consul_sd_configs": schema.ListNestedAttribute{
						Description:         "ConsulSDConfigs defines a list of Consul service discovery configurations.",
						MarkdownDescription: "ConsulSDConfigs defines a list of Consul service discovery configurations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"allow_stale": schema.BoolAttribute{
									Description:         "Allow stale Consul results (see https://www.consul.io/api/features/consistency.html). Will reduce load on Consul. If unset, Prometheus uses its default value.",
									MarkdownDescription: "Allow stale Consul results (see https://www.consul.io/api/features/consistency.html). Will reduce load on Consul. If unset, Prometheus uses its default value.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"authorization": schema.SingleNestedAttribute{
									Description:         "Authorization header configuration to authenticate against the Consul Server.",
									MarkdownDescription: "Authorization header configuration to authenticate against the Consul Server.",
									Attributes: map[string]schema.Attribute{
										"credentials": schema.SingleNestedAttribute{
											Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
											MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"type": schema.StringAttribute{
											Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
											MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"basic_auth": schema.SingleNestedAttribute{
									Description:         "BasicAuth information to authenticate against the Consul Server. More info: https://prometheus.io/docs/operating/configuration/#endpoints",
									MarkdownDescription: "BasicAuth information to authenticate against the Consul Server. More info: https://prometheus.io/docs/operating/configuration/#endpoints",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "The secret in the service monitor namespace that contains the password for authentication.",
											MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"username": schema.SingleNestedAttribute{
											Description:         "The secret in the service monitor namespace that contains the username for authentication.",
											MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"datacenter": schema.StringAttribute{
									Description:         "Consul Datacenter name, if not provided it will use the local Consul Agent Datacenter.",
									MarkdownDescription: "Consul Datacenter name, if not provided it will use the local Consul Agent Datacenter.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"enable_http2": schema.BoolAttribute{
									Description:         "Whether to enable HTTP2. If unset, Prometheus uses its default value.",
									MarkdownDescription: "Whether to enable HTTP2. If unset, Prometheus uses its default value.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"follow_redirects": schema.BoolAttribute{
									Description:         "Configure whether HTTP requests follow HTTP 3xx redirects. If unset, Prometheus uses its default value.",
									MarkdownDescription: "Configure whether HTTP requests follow HTTP 3xx redirects. If unset, Prometheus uses its default value.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"namespace": schema.StringAttribute{
									Description:         "Namespaces are only supported in Consul Enterprise.",
									MarkdownDescription: "Namespaces are only supported in Consul Enterprise.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"no_proxy": schema.StringAttribute{
									Description:         "Comma-separated string that can contain IPs, CIDR notation, domain names that should be excluded from proxying. IP and domain names can contain port numbers.",
									MarkdownDescription: "Comma-separated string that can contain IPs, CIDR notation, domain names that should be excluded from proxying. IP and domain names can contain port numbers.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"node_meta": schema.MapAttribute{
									Description:         "Node metadata key/value pairs to filter nodes for a given service.",
									MarkdownDescription: "Node metadata key/value pairs to filter nodes for a given service.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"oauth2": schema.SingleNestedAttribute{
									Description:         "Optional OAuth 2.0 configuration.",
									MarkdownDescription: "Optional OAuth 2.0 configuration.",
									Attributes: map[string]schema.Attribute{
										"client_id": schema.SingleNestedAttribute{
											Description:         "The secret or configmap containing the OAuth2 client id",
											MarkdownDescription: "The secret or configmap containing the OAuth2 client id",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap containing data to use for the targets.",
													MarkdownDescription: "ConfigMap containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
													Description:         "Secret containing data to use for the targets.",
													MarkdownDescription: "Secret containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from.  Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"client_secret": schema.SingleNestedAttribute{
											Description:         "The secret containing the OAuth2 client secret",
											MarkdownDescription: "The secret containing the OAuth2 client secret",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"endpoint_params": schema.MapAttribute{
											Description:         "Parameters to append to the token URL",
											MarkdownDescription: "Parameters to append to the token URL",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"scopes": schema.ListAttribute{
											Description:         "OAuth2 scopes used for the token request",
											MarkdownDescription: "OAuth2 scopes used for the token request",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"token_url": schema.StringAttribute{
											Description:         "The URL to fetch the token from",
											MarkdownDescription: "The URL to fetch the token from",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"partition": schema.StringAttribute{
									Description:         "Admin Partitions are only supported in Consul Enterprise.",
									MarkdownDescription: "Admin Partitions are only supported in Consul Enterprise.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"proxy_connect_header": schema.SingleNestedAttribute{
									Description:         "Specifies headers to send to proxies during CONNECT requests.",
									MarkdownDescription: "Specifies headers to send to proxies during CONNECT requests.",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from.  Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

								"proxy_from_environment": schema.BoolAttribute{
									Description:         "Use proxy URL indicated by environment variables (HTTP_PROXY, https_proxy, HTTPs_PROXY, https_proxy, and no_proxy) If unset, Prometheus uses its default value.",
									MarkdownDescription: "Use proxy URL indicated by environment variables (HTTP_PROXY, https_proxy, HTTPs_PROXY, https_proxy, and no_proxy) If unset, Prometheus uses its default value.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"proxy_url": schema.StringAttribute{
									Description:         "Optional proxy URL.",
									MarkdownDescription: "Optional proxy URL.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"refresh_interval": schema.StringAttribute{
									Description:         "The time after which the provided names are refreshed. On large setup it might be a good idea to increase this value because the catalog will change all the time. If unset, Prometheus uses its default value.",
									MarkdownDescription: "The time after which the provided names are refreshed. On large setup it might be a good idea to increase this value because the catalog will change all the time. If unset, Prometheus uses its default value.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"scheme": schema.StringAttribute{
									Description:         "HTTP Scheme default 'http'",
									MarkdownDescription: "HTTP Scheme default 'http'",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"server": schema.StringAttribute{
									Description:         "A valid string consisting of a hostname or IP followed by an optional port number.",
									MarkdownDescription: "A valid string consisting of a hostname or IP followed by an optional port number.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"services": schema.ListAttribute{
									Description:         "A list of services for which targets are retrieved. If omitted, all services are scraped.",
									MarkdownDescription: "A list of services for which targets are retrieved. If omitted, all services are scraped.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"tag_separator": schema.StringAttribute{
									Description:         "The string by which Consul tags are joined into the tag label. If unset, Prometheus uses its default value.",
									MarkdownDescription: "The string by which Consul tags are joined into the tag label. If unset, Prometheus uses its default value.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"tags": schema.ListAttribute{
									Description:         "An optional list of tags used to filter nodes for a given service. Services must contain all tags in the list.",
									MarkdownDescription: "An optional list of tags used to filter nodes for a given service. Services must contain all tags in the list.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"tls_config": schema.SingleNestedAttribute{
									Description:         "TLS Config",
									MarkdownDescription: "TLS Config",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "Certificate authority used when verifying server certificates.",
											MarkdownDescription: "Certificate authority used when verifying server certificates.",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap containing data to use for the targets.",
													MarkdownDescription: "ConfigMap containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
													Description:         "Secret containing data to use for the targets.",
													MarkdownDescription: "Secret containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from.  Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"cert": schema.SingleNestedAttribute{
											Description:         "Client certificate to present when doing client-authentication.",
											MarkdownDescription: "Client certificate to present when doing client-authentication.",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap containing data to use for the targets.",
													MarkdownDescription: "ConfigMap containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
													Description:         "Secret containing data to use for the targets.",
													MarkdownDescription: "Secret containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from.  Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"insecure_skip_verify": schema.BoolAttribute{
											Description:         "Disable target certificate validation.",
											MarkdownDescription: "Disable target certificate validation.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"key_secret": schema.SingleNestedAttribute{
											Description:         "Secret containing the client key file for the targets.",
											MarkdownDescription: "Secret containing the client key file for the targets.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"server_name": schema.StringAttribute{
											Description:         "Used to verify the hostname for the targets.",
											MarkdownDescription: "Used to verify the hostname for the targets.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"token_ref": schema.SingleNestedAttribute{
									Description:         "Consul ACL TokenRef, if not provided it will use the ACL from the local Consul Agent.",
									MarkdownDescription: "Consul ACL TokenRef, if not provided it will use the ACL from the local Consul Agent.",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from.  Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"file_sd_configs": schema.ListNestedAttribute{
						Description:         "FileSDConfigs defines a list of file service discovery configurations.",
						MarkdownDescription: "FileSDConfigs defines a list of file service discovery configurations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"files": schema.ListAttribute{
									Description:         "List of files to be used for file discovery. Recommendation: use absolute paths. While relative paths work, the prometheus-operator project makes no guarantees about the working directory where the configuration file is stored. Files must be mounted using Prometheus.ConfigMaps or Prometheus.Secrets.",
									MarkdownDescription: "List of files to be used for file discovery. Recommendation: use absolute paths. While relative paths work, the prometheus-operator project makes no guarantees about the working directory where the configuration file is stored. Files must be mounted using Prometheus.ConfigMaps or Prometheus.Secrets.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"refresh_interval": schema.StringAttribute{
									Description:         "RefreshInterval configures the refresh interval at which Prometheus will reload the content of the files.",
									MarkdownDescription: "RefreshInterval configures the refresh interval at which Prometheus will reload the content of the files.",
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

					"honor_labels": schema.BoolAttribute{
						Description:         "HonorLabels chooses the metric's labels on collisions with target labels.",
						MarkdownDescription: "HonorLabels chooses the metric's labels on collisions with target labels.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"honor_timestamps": schema.BoolAttribute{
						Description:         "HonorTimestamps controls whether Prometheus respects the timestamps present in scraped data.",
						MarkdownDescription: "HonorTimestamps controls whether Prometheus respects the timestamps present in scraped data.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"http_sd_configs": schema.ListNestedAttribute{
						Description:         "HTTPSDConfigs defines a list of HTTP service discovery configurations.",
						MarkdownDescription: "HTTPSDConfigs defines a list of HTTP service discovery configurations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"authorization": schema.SingleNestedAttribute{
									Description:         "Authorization header configuration to authenticate against the target HTTP endpoint.",
									MarkdownDescription: "Authorization header configuration to authenticate against the target HTTP endpoint.",
									Attributes: map[string]schema.Attribute{
										"credentials": schema.SingleNestedAttribute{
											Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
											MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"type": schema.StringAttribute{
											Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
											MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"basic_auth": schema.SingleNestedAttribute{
									Description:         "BasicAuth information to authenticate against the target HTTP endpoint. More info: https://prometheus.io/docs/operating/configuration/#endpoints",
									MarkdownDescription: "BasicAuth information to authenticate against the target HTTP endpoint. More info: https://prometheus.io/docs/operating/configuration/#endpoints",
									Attributes: map[string]schema.Attribute{
										"password": schema.SingleNestedAttribute{
											Description:         "The secret in the service monitor namespace that contains the password for authentication.",
											MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"username": schema.SingleNestedAttribute{
											Description:         "The secret in the service monitor namespace that contains the username for authentication.",
											MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"refresh_interval": schema.StringAttribute{
									Description:         "RefreshInterval configures the refresh interval at which Prometheus will re-query the endpoint to update the target list.",
									MarkdownDescription: "RefreshInterval configures the refresh interval at which Prometheus will re-query the endpoint to update the target list.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"tls_config": schema.SingleNestedAttribute{
									Description:         "TLS configuration applying to the target HTTP endpoint.",
									MarkdownDescription: "TLS configuration applying to the target HTTP endpoint.",
									Attributes: map[string]schema.Attribute{
										"ca": schema.SingleNestedAttribute{
											Description:         "Certificate authority used when verifying server certificates.",
											MarkdownDescription: "Certificate authority used when verifying server certificates.",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap containing data to use for the targets.",
													MarkdownDescription: "ConfigMap containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
													Description:         "Secret containing data to use for the targets.",
													MarkdownDescription: "Secret containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from.  Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"cert": schema.SingleNestedAttribute{
											Description:         "Client certificate to present when doing client-authentication.",
											MarkdownDescription: "Client certificate to present when doing client-authentication.",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap containing data to use for the targets.",
													MarkdownDescription: "ConfigMap containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select.",
															MarkdownDescription: "The key to select.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
													Description:         "Secret containing data to use for the targets.",
													MarkdownDescription: "Secret containing data to use for the targets.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key of the secret to select from.  Must be a valid secret key.",
															MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"insecure_skip_verify": schema.BoolAttribute{
											Description:         "Disable target certificate validation.",
											MarkdownDescription: "Disable target certificate validation.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"key_secret": schema.SingleNestedAttribute{
											Description:         "Secret containing the client key file for the targets.",
											MarkdownDescription: "Secret containing the client key file for the targets.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"server_name": schema.StringAttribute{
											Description:         "Used to verify the hostname for the targets.",
											MarkdownDescription: "Used to verify the hostname for the targets.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"url": schema.StringAttribute{
									Description:         "URL from which the targets are fetched.",
									MarkdownDescription: "URL from which the targets are fetched.",
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

					"kubernetes_sd_configs": schema.ListNestedAttribute{
						Description:         "KubernetesSDConfigs defines a list of Kubernetes service discovery configurations.",
						MarkdownDescription: "KubernetesSDConfigs defines a list of Kubernetes service discovery configurations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"role": schema.StringAttribute{
									Description:         "Role of the Kubernetes entities that should be discovered. Currently the only supported role is 'Node'.",
									MarkdownDescription: "Role of the Kubernetes entities that should be discovered. Currently the only supported role is 'Node'.",
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

					"label_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on number of labels that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"label_name_length_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on length of labels name that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"label_value_length_limit": schema.Int64Attribute{
						Description:         "Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on length of labels value that will be accepted for a sample. Only valid in Prometheus versions 2.27.0 and newer.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"metrics_path": schema.StringAttribute{
						Description:         "MetricsPath HTTP path to scrape for metrics. If empty, Prometheus uses the default value (e.g. /metrics).",
						MarkdownDescription: "MetricsPath HTTP path to scrape for metrics. If empty, Prometheus uses the default value (e.g. /metrics).",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"params": schema.MapAttribute{
						Description:         "Optional HTTP URL parameters",
						MarkdownDescription: "Optional HTTP URL parameters",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"relabelings": schema.ListNestedAttribute{
						Description:         "RelabelConfigs defines how to rewrite the target's labels before scraping. Prometheus Operator automatically adds relabelings for a few standard Kubernetes fields. The original scrape job's name is available via the '__tmp_prometheus_job_name' label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
						MarkdownDescription: "RelabelConfigs defines how to rewrite the target's labels before scraping. Prometheus Operator automatically adds relabelings for a few standard Kubernetes fields. The original scrape job's name is available via the '__tmp_prometheus_job_name' label. More info: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#relabel_config",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Description:         "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
									MarkdownDescription: "Action to perform based on the regex matching.  'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0. 'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.  Default: 'Replace'",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"modulus": schema.Int64Attribute{
									Description:         "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
									MarkdownDescription: "Modulus to take of the hash of the source label values.  Only applicable when the action is 'HashMod'.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"regex": schema.StringAttribute{
									Description:         "Regular expression against which the extracted value is matched.",
									MarkdownDescription: "Regular expression against which the extracted value is matched.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"replacement": schema.StringAttribute{
									Description:         "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
									MarkdownDescription: "Replacement value against which a Replace action is performed if the regular expression matches.  Regex capture groups are available.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"separator": schema.StringAttribute{
									Description:         "Separator is the string between concatenated SourceLabels.",
									MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"source_labels": schema.ListAttribute{
									Description:         "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
									MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured Separator and matched against the configured regular expression.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"target_label": schema.StringAttribute{
									Description:         "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
									MarkdownDescription: "Label to which the resulting string is written in a replacement.  It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase', 'KeepEqual' and 'DropEqual' actions.  Regex capture groups are available.",
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

					"sample_limit": schema.Int64Attribute{
						Description:         "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
						MarkdownDescription: "SampleLimit defines per-scrape limit on number of scraped samples that will be accepted.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scheme": schema.StringAttribute{
						Description:         "Configures the protocol scheme used for requests. If empty, Prometheus uses HTTP by default.",
						MarkdownDescription: "Configures the protocol scheme used for requests. If empty, Prometheus uses HTTP by default.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scrape_interval": schema.StringAttribute{
						Description:         "ScrapeInterval is the interval between consecutive scrapes.",
						MarkdownDescription: "ScrapeInterval is the interval between consecutive scrapes.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"scrape_timeout": schema.StringAttribute{
						Description:         "ScrapeTimeout is the number of seconds to wait until a scrape request times out.",
						MarkdownDescription: "ScrapeTimeout is the number of seconds to wait until a scrape request times out.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"static_configs": schema.ListNestedAttribute{
						Description:         "StaticConfigs defines a list of static targets with a common label set.",
						MarkdownDescription: "StaticConfigs defines a list of static targets with a common label set.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"labels": schema.MapAttribute{
									Description:         "Labels assigned to all metrics scraped from the targets.",
									MarkdownDescription: "Labels assigned to all metrics scraped from the targets.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"targets": schema.ListAttribute{
									Description:         "List of targets for this static configuration.",
									MarkdownDescription: "List of targets for this static configuration.",
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

					"target_limit": schema.Int64Attribute{
						Description:         "TargetLimit defines a limit on the number of scraped targets that will be accepted.",
						MarkdownDescription: "TargetLimit defines a limit on the number of scraped targets that will be accepted.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"tls_config": schema.SingleNestedAttribute{
						Description:         "TLS configuration to use on every scrape request",
						MarkdownDescription: "TLS configuration to use on every scrape request",
						Attributes: map[string]schema.Attribute{
							"ca": schema.SingleNestedAttribute{
								Description:         "Certificate authority used when verifying server certificates.",
								MarkdownDescription: "Certificate authority used when verifying server certificates.",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "ConfigMap containing data to use for the targets.",
										MarkdownDescription: "ConfigMap containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
										Description:         "Secret containing data to use for the targets.",
										MarkdownDescription: "Secret containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"cert": schema.SingleNestedAttribute{
								Description:         "Client certificate to present when doing client-authentication.",
								MarkdownDescription: "Client certificate to present when doing client-authentication.",
								Attributes: map[string]schema.Attribute{
									"config_map": schema.SingleNestedAttribute{
										Description:         "ConfigMap containing data to use for the targets.",
										MarkdownDescription: "ConfigMap containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select.",
												MarkdownDescription: "The key to select.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
										Description:         "Secret containing data to use for the targets.",
										MarkdownDescription: "Secret containing data to use for the targets.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key of the secret to select from.  Must be a valid secret key.",
												MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"insecure_skip_verify": schema.BoolAttribute{
								Description:         "Disable target certificate validation.",
								MarkdownDescription: "Disable target certificate validation.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"key_secret": schema.SingleNestedAttribute{
								Description:         "Secret containing the client key file for the targets.",
								MarkdownDescription: "Secret containing the client key file for the targets.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key of the secret to select from.  Must be a valid secret key.",
										MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

							"server_name": schema.StringAttribute{
								Description:         "Used to verify the hostname for the targets.",
								MarkdownDescription: "Used to verify the hostname for the targets.",
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
		},
	}
}

func (r *MonitoringCoreosComScrapeConfigV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *MonitoringCoreosComScrapeConfigV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_monitoring_coreos_com_scrape_config_v1alpha1")

	var data MonitoringCoreosComScrapeConfigV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1alpha1", Resource: "ScrapeConfig"}).
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

	var readResponse MonitoringCoreosComScrapeConfigV1Alpha1DataSourceData
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
	data.ApiVersion = pointer.String("monitoring.coreos.com/v1alpha1")
	data.Kind = pointer.String("ScrapeConfig")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
