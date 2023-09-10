/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package monitoring_coreos_com_v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"regexp"
	"strings"
)

var (
	_ resource.Resource                = &MonitoringCoreosComAlertmanagerConfigV1Beta1Resource{}
	_ resource.ResourceWithConfigure   = &MonitoringCoreosComAlertmanagerConfigV1Beta1Resource{}
	_ resource.ResourceWithImportState = &MonitoringCoreosComAlertmanagerConfigV1Beta1Resource{}
)

func NewMonitoringCoreosComAlertmanagerConfigV1Beta1Resource() resource.Resource {
	return &MonitoringCoreosComAlertmanagerConfigV1Beta1Resource{}
}

type MonitoringCoreosComAlertmanagerConfigV1Beta1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type MonitoringCoreosComAlertmanagerConfigV1Beta1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		InhibitRules *[]struct {
			Equal       *[]string `tfsdk:"equal" json:"equal,omitempty"`
			SourceMatch *[]struct {
				MatchType *string `tfsdk:"match_type" json:"matchType,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"source_match" json:"sourceMatch,omitempty"`
			TargetMatch *[]struct {
				MatchType *string `tfsdk:"match_type" json:"matchType,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"target_match" json:"targetMatch,omitempty"`
		} `tfsdk:"inhibit_rules" json:"inhibitRules,omitempty"`
		Receivers *[]struct {
			DiscordConfigs *[]struct {
				ApiURL *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"api_url" json:"apiURL,omitempty"`
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				Message      *string `tfsdk:"message" json:"message,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				Title        *string `tfsdk:"title" json:"title,omitempty"`
			} `tfsdk:"discord_configs" json:"discordConfigs,omitempty"`
			EmailConfigs *[]struct {
				AuthIdentity *string `tfsdk:"auth_identity" json:"authIdentity,omitempty"`
				AuthPassword *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"auth_password" json:"authPassword,omitempty"`
				AuthSecret *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"auth_secret" json:"authSecret,omitempty"`
				AuthUsername *string `tfsdk:"auth_username" json:"authUsername,omitempty"`
				From         *string `tfsdk:"from" json:"from,omitempty"`
				Headers      *[]struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				Hello        *string `tfsdk:"hello" json:"hello,omitempty"`
				Html         *string `tfsdk:"html" json:"html,omitempty"`
				RequireTLS   *bool   `tfsdk:"require_tls" json:"requireTLS,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				Smarthost    *string `tfsdk:"smarthost" json:"smarthost,omitempty"`
				Text         *string `tfsdk:"text" json:"text,omitempty"`
				TlsConfig    *struct {
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
				To *string `tfsdk:"to" json:"to,omitempty"`
			} `tfsdk:"email_configs" json:"emailConfigs,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			OpsgenieConfigs *[]struct {
				Actions *string `tfsdk:"actions" json:"actions,omitempty"`
				ApiKey  *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"api_key" json:"apiKey,omitempty"`
				ApiURL      *string `tfsdk:"api_url" json:"apiURL,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Details     *[]struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"details" json:"details,omitempty"`
				Entity     *string `tfsdk:"entity" json:"entity,omitempty"`
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				Message    *string `tfsdk:"message" json:"message,omitempty"`
				Note       *string `tfsdk:"note" json:"note,omitempty"`
				Priority   *string `tfsdk:"priority" json:"priority,omitempty"`
				Responders *[]struct {
					Id       *string `tfsdk:"id" json:"id,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Type     *string `tfsdk:"type" json:"type,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"responders" json:"responders,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				Source       *string `tfsdk:"source" json:"source,omitempty"`
				Tags         *string `tfsdk:"tags" json:"tags,omitempty"`
			} `tfsdk:"opsgenie_configs" json:"opsgenieConfigs,omitempty"`
			PagerdutyConfigs *[]struct {
				Class       *string `tfsdk:"class" json:"class,omitempty"`
				Client      *string `tfsdk:"client" json:"client,omitempty"`
				ClientURL   *string `tfsdk:"client_url" json:"clientURL,omitempty"`
				Component   *string `tfsdk:"component" json:"component,omitempty"`
				Description *string `tfsdk:"description" json:"description,omitempty"`
				Details     *[]struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"details" json:"details,omitempty"`
				Group      *string `tfsdk:"group" json:"group,omitempty"`
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				PagerDutyImageConfigs *[]struct {
					Alt  *string `tfsdk:"alt" json:"alt,omitempty"`
					Href *string `tfsdk:"href" json:"href,omitempty"`
					Src  *string `tfsdk:"src" json:"src,omitempty"`
				} `tfsdk:"pager_duty_image_configs" json:"pagerDutyImageConfigs,omitempty"`
				PagerDutyLinkConfigs *[]struct {
					Alt  *string `tfsdk:"alt" json:"alt,omitempty"`
					Href *string `tfsdk:"href" json:"href,omitempty"`
				} `tfsdk:"pager_duty_link_configs" json:"pagerDutyLinkConfigs,omitempty"`
				RoutingKey *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"routing_key" json:"routingKey,omitempty"`
				SendResolved *bool `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				ServiceKey   *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"service_key" json:"serviceKey,omitempty"`
				Severity *string `tfsdk:"severity" json:"severity,omitempty"`
				Url      *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"pagerduty_configs" json:"pagerdutyConfigs,omitempty"`
			PushoverConfigs *[]struct {
				Expire     *string `tfsdk:"expire" json:"expire,omitempty"`
				Html       *bool   `tfsdk:"html" json:"html,omitempty"`
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				Message      *string `tfsdk:"message" json:"message,omitempty"`
				Priority     *string `tfsdk:"priority" json:"priority,omitempty"`
				Retry        *string `tfsdk:"retry" json:"retry,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				Sound        *string `tfsdk:"sound" json:"sound,omitempty"`
				Title        *string `tfsdk:"title" json:"title,omitempty"`
				Token        *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"token" json:"token,omitempty"`
				Url      *string `tfsdk:"url" json:"url,omitempty"`
				UrlTitle *string `tfsdk:"url_title" json:"urlTitle,omitempty"`
				UserKey  *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"user_key" json:"userKey,omitempty"`
			} `tfsdk:"pushover_configs" json:"pushoverConfigs,omitempty"`
			SlackConfigs *[]struct {
				Actions *[]struct {
					Confirm *struct {
						DismissText *string `tfsdk:"dismiss_text" json:"dismissText,omitempty"`
						OkText      *string `tfsdk:"ok_text" json:"okText,omitempty"`
						Text        *string `tfsdk:"text" json:"text,omitempty"`
						Title       *string `tfsdk:"title" json:"title,omitempty"`
					} `tfsdk:"confirm" json:"confirm,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Style *string `tfsdk:"style" json:"style,omitempty"`
					Text  *string `tfsdk:"text" json:"text,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Url   *string `tfsdk:"url" json:"url,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"actions" json:"actions,omitempty"`
				ApiURL *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"api_url" json:"apiURL,omitempty"`
				CallbackId *string `tfsdk:"callback_id" json:"callbackId,omitempty"`
				Channel    *string `tfsdk:"channel" json:"channel,omitempty"`
				Color      *string `tfsdk:"color" json:"color,omitempty"`
				Fallback   *string `tfsdk:"fallback" json:"fallback,omitempty"`
				Fields     *[]struct {
					Short *bool   `tfsdk:"short" json:"short,omitempty"`
					Title *string `tfsdk:"title" json:"title,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
				Footer     *string `tfsdk:"footer" json:"footer,omitempty"`
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				IconEmoji    *string   `tfsdk:"icon_emoji" json:"iconEmoji,omitempty"`
				IconURL      *string   `tfsdk:"icon_url" json:"iconURL,omitempty"`
				ImageURL     *string   `tfsdk:"image_url" json:"imageURL,omitempty"`
				LinkNames    *bool     `tfsdk:"link_names" json:"linkNames,omitempty"`
				MrkdwnIn     *[]string `tfsdk:"mrkdwn_in" json:"mrkdwnIn,omitempty"`
				Pretext      *string   `tfsdk:"pretext" json:"pretext,omitempty"`
				SendResolved *bool     `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				ShortFields  *bool     `tfsdk:"short_fields" json:"shortFields,omitempty"`
				Text         *string   `tfsdk:"text" json:"text,omitempty"`
				ThumbURL     *string   `tfsdk:"thumb_url" json:"thumbURL,omitempty"`
				Title        *string   `tfsdk:"title" json:"title,omitempty"`
				TitleLink    *string   `tfsdk:"title_link" json:"titleLink,omitempty"`
				Username     *string   `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"slack_configs" json:"slackConfigs,omitempty"`
			SnsConfigs *[]struct {
				ApiURL     *string            `tfsdk:"api_url" json:"apiURL,omitempty"`
				Attributes *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				Message      *string `tfsdk:"message" json:"message,omitempty"`
				PhoneNumber  *string `tfsdk:"phone_number" json:"phoneNumber,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				Sigv4        *struct {
					AccessKey *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"access_key" json:"accessKey,omitempty"`
					Profile   *string `tfsdk:"profile" json:"profile,omitempty"`
					Region    *string `tfsdk:"region" json:"region,omitempty"`
					RoleArn   *string `tfsdk:"role_arn" json:"roleArn,omitempty"`
					SecretKey *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key" json:"secretKey,omitempty"`
				} `tfsdk:"sigv4" json:"sigv4,omitempty"`
				Subject   *string `tfsdk:"subject" json:"subject,omitempty"`
				TargetARN *string `tfsdk:"target_arn" json:"targetARN,omitempty"`
				TopicARN  *string `tfsdk:"topic_arn" json:"topicARN,omitempty"`
			} `tfsdk:"sns_configs" json:"snsConfigs,omitempty"`
			TelegramConfigs *[]struct {
				ApiURL   *string `tfsdk:"api_url" json:"apiURL,omitempty"`
				BotToken *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"bot_token" json:"botToken,omitempty"`
				BotTokenFile         *string `tfsdk:"bot_token_file" json:"botTokenFile,omitempty"`
				ChatID               *int64  `tfsdk:"chat_id" json:"chatID,omitempty"`
				DisableNotifications *bool   `tfsdk:"disable_notifications" json:"disableNotifications,omitempty"`
				HttpConfig           *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				Message      *string `tfsdk:"message" json:"message,omitempty"`
				ParseMode    *string `tfsdk:"parse_mode" json:"parseMode,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
			} `tfsdk:"telegram_configs" json:"telegramConfigs,omitempty"`
			VictoropsConfigs *[]struct {
				ApiKey *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"api_key" json:"apiKey,omitempty"`
				ApiUrl       *string `tfsdk:"api_url" json:"apiUrl,omitempty"`
				CustomFields *[]struct {
					Key   *string `tfsdk:"key" json:"key,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"custom_fields" json:"customFields,omitempty"`
				EntityDisplayName *string `tfsdk:"entity_display_name" json:"entityDisplayName,omitempty"`
				HttpConfig        *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				MessageType    *string `tfsdk:"message_type" json:"messageType,omitempty"`
				MonitoringTool *string `tfsdk:"monitoring_tool" json:"monitoringTool,omitempty"`
				RoutingKey     *string `tfsdk:"routing_key" json:"routingKey,omitempty"`
				SendResolved   *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				StateMessage   *string `tfsdk:"state_message" json:"stateMessage,omitempty"`
			} `tfsdk:"victorops_configs" json:"victoropsConfigs,omitempty"`
			WebexConfigs *[]struct {
				ApiURL     *string `tfsdk:"api_url" json:"apiURL,omitempty"`
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				Message      *string `tfsdk:"message" json:"message,omitempty"`
				RoomID       *string `tfsdk:"room_id" json:"roomID,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
			} `tfsdk:"webex_configs" json:"webexConfigs,omitempty"`
			WebhookConfigs *[]struct {
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				MaxAlerts    *int64  `tfsdk:"max_alerts" json:"maxAlerts,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				Url          *string `tfsdk:"url" json:"url,omitempty"`
				UrlSecret    *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"url_secret" json:"urlSecret,omitempty"`
			} `tfsdk:"webhook_configs" json:"webhookConfigs,omitempty"`
			WechatConfigs *[]struct {
				AgentID   *string `tfsdk:"agent_id" json:"agentID,omitempty"`
				ApiSecret *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"api_secret" json:"apiSecret,omitempty"`
				ApiURL     *string `tfsdk:"api_url" json:"apiURL,omitempty"`
				CorpID     *string `tfsdk:"corp_id" json:"corpID,omitempty"`
				HttpConfig *struct {
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
					BearerTokenSecret *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearerTokenSecret,omitempty"`
					FollowRedirects *bool `tfsdk:"follow_redirects" json:"followRedirects,omitempty"`
					Oauth2          *struct {
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
					ProxyURL  *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					TlsConfig *struct {
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
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				Message      *string `tfsdk:"message" json:"message,omitempty"`
				MessageType  *string `tfsdk:"message_type" json:"messageType,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				ToParty      *string `tfsdk:"to_party" json:"toParty,omitempty"`
				ToTag        *string `tfsdk:"to_tag" json:"toTag,omitempty"`
				ToUser       *string `tfsdk:"to_user" json:"toUser,omitempty"`
			} `tfsdk:"wechat_configs" json:"wechatConfigs,omitempty"`
		} `tfsdk:"receivers" json:"receivers,omitempty"`
		Route *struct {
			ActiveTimeIntervals *[]string `tfsdk:"active_time_intervals" json:"activeTimeIntervals,omitempty"`
			Continue            *bool     `tfsdk:"continue" json:"continue,omitempty"`
			GroupBy             *[]string `tfsdk:"group_by" json:"groupBy,omitempty"`
			GroupInterval       *string   `tfsdk:"group_interval" json:"groupInterval,omitempty"`
			GroupWait           *string   `tfsdk:"group_wait" json:"groupWait,omitempty"`
			Matchers            *[]struct {
				MatchType *string `tfsdk:"match_type" json:"matchType,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"matchers" json:"matchers,omitempty"`
			MuteTimeIntervals *[]string `tfsdk:"mute_time_intervals" json:"muteTimeIntervals,omitempty"`
			Receiver          *string   `tfsdk:"receiver" json:"receiver,omitempty"`
			RepeatInterval    *string   `tfsdk:"repeat_interval" json:"repeatInterval,omitempty"`
			Routes            *[]string `tfsdk:"routes" json:"routes,omitempty"`
		} `tfsdk:"route" json:"route,omitempty"`
		TimeIntervals *[]struct {
			Name          *string `tfsdk:"name" json:"name,omitempty"`
			TimeIntervals *[]struct {
				DaysOfMonth *[]struct {
					End   *int64 `tfsdk:"end" json:"end,omitempty"`
					Start *int64 `tfsdk:"start" json:"start,omitempty"`
				} `tfsdk:"days_of_month" json:"daysOfMonth,omitempty"`
				Months *[]string `tfsdk:"months" json:"months,omitempty"`
				Times  *[]struct {
					EndTime   *string `tfsdk:"end_time" json:"endTime,omitempty"`
					StartTime *string `tfsdk:"start_time" json:"startTime,omitempty"`
				} `tfsdk:"times" json:"times,omitempty"`
				Weekdays *[]string `tfsdk:"weekdays" json:"weekdays,omitempty"`
				Years    *[]string `tfsdk:"years" json:"years,omitempty"`
			} `tfsdk:"time_intervals" json:"timeIntervals,omitempty"`
		} `tfsdk:"time_intervals" json:"timeIntervals,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Beta1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_monitoring_coreos_com_alertmanager_config_v1beta1"
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Beta1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AlertmanagerConfig defines a namespaced AlertmanagerConfig to be aggregated across multiple namespaces configuring one Alertmanager cluster.",
		MarkdownDescription: "AlertmanagerConfig defines a namespaced AlertmanagerConfig to be aggregated across multiple namespaces configuring one Alertmanager cluster.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
						Validators: []validator.Map{
							validators.AnnotationValidator(),
						},
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "AlertmanagerConfigSpec is a specification of the desired behavior of the Alertmanager configuration. By definition, the Alertmanager configuration only applies to alerts for which the 'namespace' label is equal to the namespace of the AlertmanagerConfig resource.",
				MarkdownDescription: "AlertmanagerConfigSpec is a specification of the desired behavior of the Alertmanager configuration. By definition, the Alertmanager configuration only applies to alerts for which the 'namespace' label is equal to the namespace of the AlertmanagerConfig resource.",
				Attributes: map[string]schema.Attribute{
					"inhibit_rules": schema.ListNestedAttribute{
						Description:         "List of inhibition rules. The rules will only apply to alerts matching the resource's namespace.",
						MarkdownDescription: "List of inhibition rules. The rules will only apply to alerts matching the resource's namespace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"equal": schema.ListAttribute{
									Description:         "Labels that must have an equal value in the source and target alert for the inhibition to take effect.",
									MarkdownDescription: "Labels that must have an equal value in the source and target alert for the inhibition to take effect.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_match": schema.ListNestedAttribute{
									Description:         "Matchers for which one or more alerts have to exist for the inhibition to take effect. The operator enforces that the alert matches the resource's namespace.",
									MarkdownDescription: "Matchers for which one or more alerts have to exist for the inhibition to take effect. The operator enforces that the alert matches the resource's namespace.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_type": schema.StringAttribute{
												Description:         "Match operator, one of '=' (equal to), '!=' (not equal to), '=~' (regex match) or '!~' (not regex match). Negative operators ('!=' and '!~') require Alertmanager >= v0.22.0.",
												MarkdownDescription: "Match operator, one of '=' (equal to), '!=' (not equal to), '=~' (regex match) or '!~' (not regex match). Negative operators ('!=' and '!~') require Alertmanager >= v0.22.0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("!=", "=", "=~", "!~"),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Label to match.",
												MarkdownDescription: "Label to match.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"value": schema.StringAttribute{
												Description:         "Label value to match.",
												MarkdownDescription: "Label value to match.",
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

								"target_match": schema.ListNestedAttribute{
									Description:         "Matchers that have to be fulfilled in the alerts to be muted. The operator enforces that the alert matches the resource's namespace.",
									MarkdownDescription: "Matchers that have to be fulfilled in the alerts to be muted. The operator enforces that the alert matches the resource's namespace.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_type": schema.StringAttribute{
												Description:         "Match operator, one of '=' (equal to), '!=' (not equal to), '=~' (regex match) or '!~' (not regex match). Negative operators ('!=' and '!~') require Alertmanager >= v0.22.0.",
												MarkdownDescription: "Match operator, one of '=' (equal to), '!=' (not equal to), '=~' (regex match) or '!~' (not regex match). Negative operators ('!=' and '!~') require Alertmanager >= v0.22.0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("!=", "=", "=~", "!~"),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Label to match.",
												MarkdownDescription: "Label to match.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"value": schema.StringAttribute{
												Description:         "Label value to match.",
												MarkdownDescription: "Label value to match.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"receivers": schema.ListNestedAttribute{
						Description:         "List of receivers.",
						MarkdownDescription: "List of receivers.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"discord_configs": schema.ListNestedAttribute{
									Description:         "List of Slack configurations.",
									MarkdownDescription: "List of Slack configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_url": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the Discord webhook URL. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the Discord webhook URL. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"message": schema.StringAttribute{
												Description:         "The template of the message's body.",
												MarkdownDescription: "The template of the message's body.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title": schema.StringAttribute{
												Description:         "The template of the message's title.",
												MarkdownDescription: "The template of the message's title.",
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

								"email_configs": schema.ListNestedAttribute{
									Description:         "List of Email configurations.",
									MarkdownDescription: "List of Email configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"auth_identity": schema.StringAttribute{
												Description:         "The identity to use for authentication.",
												MarkdownDescription: "The identity to use for authentication.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"auth_password": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the password to use for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the password to use for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auth_secret": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the CRAM-MD5 secret. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the CRAM-MD5 secret. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auth_username": schema.StringAttribute{
												Description:         "The username to use for authentication.",
												MarkdownDescription: "The username to use for authentication.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"from": schema.StringAttribute{
												Description:         "The sender address.",
												MarkdownDescription: "The sender address.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"headers": schema.ListNestedAttribute{
												Description:         "Further headers email header key/value pairs. Overrides any headers previously set by the notification implementation.",
												MarkdownDescription: "Further headers email header key/value pairs. Overrides any headers previously set by the notification implementation.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key of the tuple.",
															MarkdownDescription: "Key of the tuple.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value of the tuple.",
															MarkdownDescription: "Value of the tuple.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"hello": schema.StringAttribute{
												Description:         "The hostname to identify to the SMTP server.",
												MarkdownDescription: "The hostname to identify to the SMTP server.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"html": schema.StringAttribute{
												Description:         "The HTML body of the email notification.",
												MarkdownDescription: "The HTML body of the email notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"require_tls": schema.BoolAttribute{
												Description:         "The SMTP TLS requirement. Note that Go does not support unencrypted connections to remote SMTP endpoints.",
												MarkdownDescription: "The SMTP TLS requirement. Note that Go does not support unencrypted connections to remote SMTP endpoints.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"smarthost": schema.StringAttribute{
												Description:         "The SMTP host and port through which emails are sent. E.g. example.com:25",
												MarkdownDescription: "The SMTP host and port through which emails are sent. E.g. example.com:25",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"text": schema.StringAttribute{
												Description:         "The text body of the email notification.",
												MarkdownDescription: "The text body of the email notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_config": schema.SingleNestedAttribute{
												Description:         "TLS configuration",
												MarkdownDescription: "TLS configuration",
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_secret": schema.SingleNestedAttribute{
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

													"server_name": schema.StringAttribute{
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"to": schema.StringAttribute{
												Description:         "The email address to send notifications to.",
												MarkdownDescription: "The email address to send notifications to.",
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

								"name": schema.StringAttribute{
									Description:         "Name of the receiver. Must be unique across all items from the list.",
									MarkdownDescription: "Name of the receiver. Must be unique across all items from the list.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
									},
								},

								"opsgenie_configs": schema.ListNestedAttribute{
									Description:         "List of OpsGenie configurations.",
									MarkdownDescription: "List of OpsGenie configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"actions": schema.StringAttribute{
												Description:         "Comma separated list of actions that will be available for the alert.",
												MarkdownDescription: "Comma separated list of actions that will be available for the alert.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"api_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the OpsGenie API key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the OpsGenie API key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"api_url": schema.StringAttribute{
												Description:         "The URL to send OpsGenie API requests to.",
												MarkdownDescription: "The URL to send OpsGenie API requests to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"description": schema.StringAttribute{
												Description:         "Description of the incident.",
												MarkdownDescription: "Description of the incident.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"details": schema.ListNestedAttribute{
												Description:         "A set of arbitrary key/value pairs that provide further detail about the incident.",
												MarkdownDescription: "A set of arbitrary key/value pairs that provide further detail about the incident.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key of the tuple.",
															MarkdownDescription: "Key of the tuple.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value of the tuple.",
															MarkdownDescription: "Value of the tuple.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"entity": schema.StringAttribute{
												Description:         "Optional field that can be used to specify which domain alert is related to.",
												MarkdownDescription: "Optional field that can be used to specify which domain alert is related to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"message": schema.StringAttribute{
												Description:         "Alert text limited to 130 characters.",
												MarkdownDescription: "Alert text limited to 130 characters.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"note": schema.StringAttribute{
												Description:         "Additional alert note.",
												MarkdownDescription: "Additional alert note.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority": schema.StringAttribute{
												Description:         "Priority level of alert. Possible values are P1, P2, P3, P4, and P5.",
												MarkdownDescription: "Priority level of alert. Possible values are P1, P2, P3, P4, and P5.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"responders": schema.ListNestedAttribute{
												Description:         "List of responders responsible for notifications.",
												MarkdownDescription: "List of responders responsible for notifications.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"id": schema.StringAttribute{
															Description:         "ID of the responder.",
															MarkdownDescription: "ID of the responder.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the responder.",
															MarkdownDescription: "Name of the responder.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "Type of responder.",
															MarkdownDescription: "Type of responder.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.OneOf("team", "teams", "user", "escalation", "schedule"),
															},
														},

														"username": schema.StringAttribute{
															Description:         "Username of the responder.",
															MarkdownDescription: "Username of the responder.",
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

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source": schema.StringAttribute{
												Description:         "Backlink to the sender of the notification.",
												MarkdownDescription: "Backlink to the sender of the notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tags": schema.StringAttribute{
												Description:         "Comma separated list of tags attached to the notifications.",
												MarkdownDescription: "Comma separated list of tags attached to the notifications.",
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

								"pagerduty_configs": schema.ListNestedAttribute{
									Description:         "List of PagerDuty configurations.",
									MarkdownDescription: "List of PagerDuty configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"class": schema.StringAttribute{
												Description:         "The class/type of the event.",
												MarkdownDescription: "The class/type of the event.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"client": schema.StringAttribute{
												Description:         "Client identification.",
												MarkdownDescription: "Client identification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"client_url": schema.StringAttribute{
												Description:         "Backlink to the sender of notification.",
												MarkdownDescription: "Backlink to the sender of notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"component": schema.StringAttribute{
												Description:         "The part or component of the affected system that is broken.",
												MarkdownDescription: "The part or component of the affected system that is broken.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"description": schema.StringAttribute{
												Description:         "Description of the incident.",
												MarkdownDescription: "Description of the incident.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"details": schema.ListNestedAttribute{
												Description:         "Arbitrary key/value pairs that provide further detail about the incident.",
												MarkdownDescription: "Arbitrary key/value pairs that provide further detail about the incident.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key of the tuple.",
															MarkdownDescription: "Key of the tuple.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value of the tuple.",
															MarkdownDescription: "Value of the tuple.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"group": schema.StringAttribute{
												Description:         "A cluster or grouping of sources.",
												MarkdownDescription: "A cluster or grouping of sources.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"pager_duty_image_configs": schema.ListNestedAttribute{
												Description:         "A list of image details to attach that provide further detail about an incident.",
												MarkdownDescription: "A list of image details to attach that provide further detail about an incident.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"alt": schema.StringAttribute{
															Description:         "Alt is the optional alternative text for the image.",
															MarkdownDescription: "Alt is the optional alternative text for the image.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"href": schema.StringAttribute{
															Description:         "Optional URL; makes the image a clickable link.",
															MarkdownDescription: "Optional URL; makes the image a clickable link.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"src": schema.StringAttribute{
															Description:         "Src of the image being attached to the incident",
															MarkdownDescription: "Src of the image being attached to the incident",
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

											"pager_duty_link_configs": schema.ListNestedAttribute{
												Description:         "A list of link details to attach that provide further detail about an incident.",
												MarkdownDescription: "A list of link details to attach that provide further detail about an incident.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"alt": schema.StringAttribute{
															Description:         "Text that describes the purpose of the link, and can be used as the link's text.",
															MarkdownDescription: "Text that describes the purpose of the link, and can be used as the link's text.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"href": schema.StringAttribute{
															Description:         "Href is the URL of the link to be attached",
															MarkdownDescription: "Href is the URL of the link to be attached",
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

											"routing_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the PagerDuty integration key (when using Events API v2). Either this field or 'serviceKey' needs to be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the PagerDuty integration key (when using Events API v2). Either this field or 'serviceKey' needs to be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the PagerDuty service key (when using integration type 'Prometheus'). Either this field or 'routingKey' needs to be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the PagerDuty service key (when using integration type 'Prometheus'). Either this field or 'routingKey' needs to be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"severity": schema.StringAttribute{
												Description:         "Severity of the incident.",
												MarkdownDescription: "Severity of the incident.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url": schema.StringAttribute{
												Description:         "The URL to send requests to.",
												MarkdownDescription: "The URL to send requests to.",
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

								"pushover_configs": schema.ListNestedAttribute{
									Description:         "List of Pushover configurations.",
									MarkdownDescription: "List of Pushover configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"expire": schema.StringAttribute{
												Description:         "How long your notification will continue to be retried for, unless the user acknowledges the notification.",
												MarkdownDescription: "How long your notification will continue to be retried for, unless the user acknowledges the notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
												},
											},

											"html": schema.BoolAttribute{
												Description:         "Whether notification message is HTML or plain text.",
												MarkdownDescription: "Whether notification message is HTML or plain text.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"message": schema.StringAttribute{
												Description:         "Notification message.",
												MarkdownDescription: "Notification message.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority": schema.StringAttribute{
												Description:         "Priority, see https://pushover.net/api#priority",
												MarkdownDescription: "Priority, see https://pushover.net/api#priority",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry": schema.StringAttribute{
												Description:         "How often the Pushover servers will send the same notification to the user. Must be at least 30 seconds.",
												MarkdownDescription: "How often the Pushover servers will send the same notification to the user. Must be at least 30 seconds.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
												},
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sound": schema.StringAttribute{
												Description:         "The name of one of the sounds supported by device clients to override the user's default sound choice",
												MarkdownDescription: "The name of one of the sounds supported by device clients to override the user's default sound choice",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title": schema.StringAttribute{
												Description:         "Notification title.",
												MarkdownDescription: "Notification title.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"token": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the registered application's API token, see https://pushover.net/apps. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the registered application's API token, see https://pushover.net/apps. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": schema.StringAttribute{
												Description:         "A supplementary URL shown alongside the message.",
												MarkdownDescription: "A supplementary URL shown alongside the message.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url_title": schema.StringAttribute{
												Description:         "A title for supplementary URL, otherwise just the URL is shown",
												MarkdownDescription: "A title for supplementary URL, otherwise just the URL is shown",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the recipient user's user key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the recipient user's user key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
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

								"slack_configs": schema.ListNestedAttribute{
									Description:         "List of Slack configurations.",
									MarkdownDescription: "List of Slack configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"actions": schema.ListNestedAttribute{
												Description:         "A list of Slack actions that are sent with each notification.",
												MarkdownDescription: "A list of Slack actions that are sent with each notification.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"confirm": schema.SingleNestedAttribute{
															Description:         "SlackConfirmationField protect users from destructive actions or particularly distinguished decisions by asking them to confirm their button click one more time. See https://api.slack.com/docs/interactive-message-field-guide#confirmation_fields for more information.",
															MarkdownDescription: "SlackConfirmationField protect users from destructive actions or particularly distinguished decisions by asking them to confirm their button click one more time. See https://api.slack.com/docs/interactive-message-field-guide#confirmation_fields for more information.",
															Attributes: map[string]schema.Attribute{
																"dismiss_text": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"ok_text": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"text": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.LengthAtLeast(1),
																	},
																},

																"title": schema.StringAttribute{
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

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"style": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"text": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"url": schema.StringAttribute{
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

											"api_url": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the Slack webhook URL. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the Slack webhook URL. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"callback_id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"channel": schema.StringAttribute{
												Description:         "The channel or user to send notifications to.",
												MarkdownDescription: "The channel or user to send notifications to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"color": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"fallback": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"fields": schema.ListNestedAttribute{
												Description:         "A list of Slack fields that are sent with each notification.",
												MarkdownDescription: "A list of Slack fields that are sent with each notification.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"short": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"title": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"value": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"footer": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"icon_emoji": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"icon_url": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_url": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"link_names": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mrkdwn_in": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pretext": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"short_fields": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"text": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"thumb_url": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title_link": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"username": schema.StringAttribute{
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

								"sns_configs": schema.ListNestedAttribute{
									Description:         "List of SNS configurations",
									MarkdownDescription: "List of SNS configurations",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_url": schema.StringAttribute{
												Description:         "The SNS API URL i.e. https://sns.us-east-2.amazonaws.com. If not specified, the SNS API URL from the SNS SDK will be used.",
												MarkdownDescription: "The SNS API URL i.e. https://sns.us-east-2.amazonaws.com. If not specified, the SNS API URL from the SNS SDK will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"attributes": schema.MapAttribute{
												Description:         "SNS message attributes.",
												MarkdownDescription: "SNS message attributes.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"message": schema.StringAttribute{
												Description:         "The message content of the SNS notification.",
												MarkdownDescription: "The message content of the SNS notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"phone_number": schema.StringAttribute{
												Description:         "Phone number if message is delivered via SMS in E.164 format. If you don't specify this value, you must specify a value for the TopicARN or TargetARN.",
												MarkdownDescription: "Phone number if message is delivered via SMS in E.164 format. If you don't specify this value, you must specify a value for the TopicARN or TargetARN.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sigv4": schema.SingleNestedAttribute{
												Description:         "Configures AWS's Signature Verification 4 signing process to sign requests.",
												MarkdownDescription: "Configures AWS's Signature Verification 4 signing process to sign requests.",
												Attributes: map[string]schema.Attribute{
													"access_key": schema.SingleNestedAttribute{
														Description:         "AccessKey is the AWS API key. If not specified, the environment variable 'AWS_ACCESS_KEY_ID' is used.",
														MarkdownDescription: "AccessKey is the AWS API key. If not specified, the environment variable 'AWS_ACCESS_KEY_ID' is used.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

													"profile": schema.StringAttribute{
														Description:         "Profile is the named AWS profile used to authenticate.",
														MarkdownDescription: "Profile is the named AWS profile used to authenticate.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "Region is the AWS region. If blank, the region from the default credentials chain used.",
														MarkdownDescription: "Region is the AWS region. If blank, the region from the default credentials chain used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"role_arn": schema.StringAttribute{
														Description:         "RoleArn is the named AWS profile used to authenticate.",
														MarkdownDescription: "RoleArn is the named AWS profile used to authenticate.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_key": schema.SingleNestedAttribute{
														Description:         "SecretKey is the AWS API secret. If not specified, the environment variable 'AWS_SECRET_ACCESS_KEY' is used.",
														MarkdownDescription: "SecretKey is the AWS API secret. If not specified, the environment variable 'AWS_SECRET_ACCESS_KEY' is used.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"subject": schema.StringAttribute{
												Description:         "Subject line when the message is delivered to email endpoints.",
												MarkdownDescription: "Subject line when the message is delivered to email endpoints.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_arn": schema.StringAttribute{
												Description:         "The  mobile platform endpoint ARN if message is delivered via mobile notifications. If you don't specify this value, you must specify a value for the topic_arn or PhoneNumber.",
												MarkdownDescription: "The  mobile platform endpoint ARN if message is delivered via mobile notifications. If you don't specify this value, you must specify a value for the topic_arn or PhoneNumber.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"topic_arn": schema.StringAttribute{
												Description:         "SNS topic ARN, i.e. arn:aws:sns:us-east-2:698519295917:My-Topic If you don't specify this value, you must specify a value for the PhoneNumber or TargetARN.",
												MarkdownDescription: "SNS topic ARN, i.e. arn:aws:sns:us-east-2:698519295917:My-Topic If you don't specify this value, you must specify a value for the PhoneNumber or TargetARN.",
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

								"telegram_configs": schema.ListNestedAttribute{
									Description:         "List of Telegram configurations.",
									MarkdownDescription: "List of Telegram configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_url": schema.StringAttribute{
												Description:         "The Telegram API URL i.e. https://api.telegram.org. If not specified, default API URL will be used.",
												MarkdownDescription: "The Telegram API URL i.e. https://api.telegram.org. If not specified, default API URL will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"bot_token": schema.SingleNestedAttribute{
												Description:         "Telegram bot token. It is mutually exclusive with 'botTokenFile'. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.  Either 'botToken' or 'botTokenFile' is required.",
												MarkdownDescription: "Telegram bot token. It is mutually exclusive with 'botTokenFile'. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.  Either 'botToken' or 'botTokenFile' is required.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"bot_token_file": schema.StringAttribute{
												Description:         "File to read the Telegram bot token from. It is mutually exclusive with 'botToken'. Either 'botToken' or 'botTokenFile' is required.  It requires Alertmanager >= v0.26.0.",
												MarkdownDescription: "File to read the Telegram bot token from. It is mutually exclusive with 'botToken'. Either 'botToken' or 'botTokenFile' is required.  It requires Alertmanager >= v0.26.0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"chat_id": schema.Int64Attribute{
												Description:         "The Telegram chat ID.",
												MarkdownDescription: "The Telegram chat ID.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disable_notifications": schema.BoolAttribute{
												Description:         "Disable telegram notifications",
												MarkdownDescription: "Disable telegram notifications",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"message": schema.StringAttribute{
												Description:         "Message template",
												MarkdownDescription: "Message template",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"parse_mode": schema.StringAttribute{
												Description:         "Parse mode for telegram message",
												MarkdownDescription: "Parse mode for telegram message",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("MarkdownV2", "Markdown", "HTML"),
												},
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether to notify about resolved alerts.",
												MarkdownDescription: "Whether to notify about resolved alerts.",
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

								"victorops_configs": schema.ListNestedAttribute{
									Description:         "List of VictorOps configurations.",
									MarkdownDescription: "List of VictorOps configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the API key to use when talking to the VictorOps API. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the API key to use when talking to the VictorOps API. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"api_url": schema.StringAttribute{
												Description:         "The VictorOps API URL.",
												MarkdownDescription: "The VictorOps API URL.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"custom_fields": schema.ListNestedAttribute{
												Description:         "Additional custom fields for notification.",
												MarkdownDescription: "Additional custom fields for notification.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "Key of the tuple.",
															MarkdownDescription: "Key of the tuple.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value of the tuple.",
															MarkdownDescription: "Value of the tuple.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"entity_display_name": schema.StringAttribute{
												Description:         "Contains summary of the alerted problem.",
												MarkdownDescription: "Contains summary of the alerted problem.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "The HTTP client's configuration.",
												MarkdownDescription: "The HTTP client's configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"message_type": schema.StringAttribute{
												Description:         "Describes the behavior of the alert (CRITICAL, WARNING, INFO).",
												MarkdownDescription: "Describes the behavior of the alert (CRITICAL, WARNING, INFO).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"monitoring_tool": schema.StringAttribute{
												Description:         "The monitoring tool the state message is from.",
												MarkdownDescription: "The monitoring tool the state message is from.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"routing_key": schema.StringAttribute{
												Description:         "A key used to map the alert to a team.",
												MarkdownDescription: "A key used to map the alert to a team.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"state_message": schema.StringAttribute{
												Description:         "Contains long explanation of the alerted problem.",
												MarkdownDescription: "Contains long explanation of the alerted problem.",
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

								"webex_configs": schema.ListNestedAttribute{
									Description:         "List of Webex configurations.",
									MarkdownDescription: "List of Webex configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_url": schema.StringAttribute{
												Description:         "The Webex Teams API URL i.e. https://webexapis.com/v1/messages",
												MarkdownDescription: "The Webex Teams API URL i.e. https://webexapis.com/v1/messages",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^https?://.+$`), ""),
												},
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "The HTTP client's configuration. You must use this configuration to supply the bot token as part of the HTTP 'Authorization' header.",
												MarkdownDescription: "The HTTP client's configuration. You must use this configuration to supply the bot token as part of the HTTP 'Authorization' header.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"message": schema.StringAttribute{
												Description:         "Message template",
												MarkdownDescription: "Message template",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"room_id": schema.StringAttribute{
												Description:         "ID of the Webex Teams room where to send the messages.",
												MarkdownDescription: "ID of the Webex Teams room where to send the messages.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether to notify about resolved alerts.",
												MarkdownDescription: "Whether to notify about resolved alerts.",
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

								"webhook_configs": schema.ListNestedAttribute{
									Description:         "List of webhook configurations.",
									MarkdownDescription: "List of webhook configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"max_alerts": schema.Int64Attribute{
												Description:         "Maximum number of alerts to be sent per webhook message. When 0, all alerts are included.",
												MarkdownDescription: "Maximum number of alerts to be sent per webhook message. When 0, all alerts are included.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url": schema.StringAttribute{
												Description:         "The URL to send HTTP POST requests to. 'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url' should be defined.",
												MarkdownDescription: "The URL to send HTTP POST requests to. 'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url' should be defined.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url_secret": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the webhook URL to send HTTP requests to. 'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url' should be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the webhook URL to send HTTP requests to. 'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url' should be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
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

								"wechat_configs": schema.ListNestedAttribute{
									Description:         "List of WeChat configurations.",
									MarkdownDescription: "List of WeChat configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"agent_id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"api_secret": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the WeChat API key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the WeChat API key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the object's namespace to select from.",
														MarkdownDescription: "The name of the secret in the object's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"api_url": schema.StringAttribute{
												Description:         "The WeChat API URL.",
												MarkdownDescription: "The WeChat API URL.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"corp_id": schema.StringAttribute{
												Description:         "The corp id for authentication.",
												MarkdownDescription: "The corp id for authentication.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														Attributes: map[string]schema.Attribute{
															"credentials": schema.SingleNestedAttribute{
																Description:         "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																MarkdownDescription: "Selects a key of a Secret in the namespace that contains the credentials for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"type": schema.StringAttribute{
																Description:         "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.  'Basic' is not a supported value.  Default: 'Bearer'",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"basic_auth": schema.SingleNestedAttribute{
														Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the password for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service monitor namespace that contains the username for authentication.",
																MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the object's namespace to select from.",
																MarkdownDescription: "The name of the secret in the object's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"follow_redirects": schema.BoolAttribute{
														Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "OAuth2 client credentials used to fetch a token for the targets.",
														MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"client_secret": schema.SingleNestedAttribute{
																Description:         "The secret containing the OAuth2 client secret",
																MarkdownDescription: "The secret containing the OAuth2 client secret",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"endpoint_params": schema.MapAttribute{
																Description:         "Parameters to append to the token URL",
																MarkdownDescription: "Parameters to append to the token URL",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "OAuth2 scopes used for the token request",
																MarkdownDescription: "OAuth2 scopes used for the token request",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "The URL to fetch the token from",
																MarkdownDescription: "The URL to fetch the token from",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"proxy_url": schema.StringAttribute{
														Description:         "Optional proxy URL.",
														MarkdownDescription: "Optional proxy URL.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_config": schema.SingleNestedAttribute{
														Description:         "TLS configuration for the client.",
														MarkdownDescription: "TLS configuration for the client.",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Description:         "Secret containing data to use for the targets.",
																		MarkdownDescription: "Secret containing data to use for the targets.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The key of the secret to select from.  Must be a valid secret key.",
																				MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_secret": schema.SingleNestedAttribute{
																Description:         "Secret containing the client key file for the targets.",
																MarkdownDescription: "Secret containing the client key file for the targets.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

															"server_name": schema.StringAttribute{
																Description:         "Used to verify the hostname for the targets.",
																MarkdownDescription: "Used to verify the hostname for the targets.",
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

											"message": schema.StringAttribute{
												Description:         "API request data as defined by the WeChat API.",
												MarkdownDescription: "API request data as defined by the WeChat API.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"message_type": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"to_party": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"to_tag": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"to_user": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"route": schema.SingleNestedAttribute{
						Description:         "The Alertmanager route definition for alerts matching the resource's namespace. If present, it will be added to the generated Alertmanager configuration as a first-level route.",
						MarkdownDescription: "The Alertmanager route definition for alerts matching the resource's namespace. If present, it will be added to the generated Alertmanager configuration as a first-level route.",
						Attributes: map[string]schema.Attribute{
							"active_time_intervals": schema.ListAttribute{
								Description:         "ActiveTimeIntervals is a list of TimeInterval names when this route should be active.",
								MarkdownDescription: "ActiveTimeIntervals is a list of TimeInterval names when this route should be active.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"continue": schema.BoolAttribute{
								Description:         "Boolean indicating whether an alert should continue matching subsequent sibling nodes. It will always be overridden to true for the first-level route by the Prometheus operator.",
								MarkdownDescription: "Boolean indicating whether an alert should continue matching subsequent sibling nodes. It will always be overridden to true for the first-level route by the Prometheus operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_by": schema.ListAttribute{
								Description:         "List of labels to group by. Labels must not be repeated (unique list). Special label '...' (aggregate by all possible labels), if provided, must be the only element in the list.",
								MarkdownDescription: "List of labels to group by. Labels must not be repeated (unique list). Special label '...' (aggregate by all possible labels), if provided, must be the only element in the list.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_interval": schema.StringAttribute{
								Description:         "How long to wait before sending an updated notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '5m'",
								MarkdownDescription: "How long to wait before sending an updated notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '5m'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_wait": schema.StringAttribute{
								Description:         "How long to wait before sending the initial notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '30s'",
								MarkdownDescription: "How long to wait before sending the initial notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '30s'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"matchers": schema.ListNestedAttribute{
								Description:         "List of matchers that the alert's labels should match. For the first level route, the operator removes any existing equality and regexp matcher on the 'namespace' label and adds a 'namespace: <object namespace>' matcher.",
								MarkdownDescription: "List of matchers that the alert's labels should match. For the first level route, the operator removes any existing equality and regexp matcher on the 'namespace' label and adds a 'namespace: <object namespace>' matcher.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match_type": schema.StringAttribute{
											Description:         "Match operator, one of '=' (equal to), '!=' (not equal to), '=~' (regex match) or '!~' (not regex match). Negative operators ('!=' and '!~') require Alertmanager >= v0.22.0.",
											MarkdownDescription: "Match operator, one of '=' (equal to), '!=' (not equal to), '=~' (regex match) or '!~' (not regex match). Negative operators ('!=' and '!~') require Alertmanager >= v0.22.0.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("!=", "=", "=~", "!~"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Label to match.",
											MarkdownDescription: "Label to match.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},

										"value": schema.StringAttribute{
											Description:         "Label value to match.",
											MarkdownDescription: "Label value to match.",
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

							"mute_time_intervals": schema.ListAttribute{
								Description:         "Note: this comment applies to the field definition above but appears below otherwise it gets included in the generated manifest. CRD schema doesn't support self-referential types for now (see https://github.com/kubernetes/kubernetes/issues/62872). We have to use an alternative type to circumvent the limitation. The downside is that the Kube API can't validate the data beyond the fact that it is a valid JSON representation. MuteTimeIntervals is a list of TimeInterval names that will mute this route when matched.",
								MarkdownDescription: "Note: this comment applies to the field definition above but appears below otherwise it gets included in the generated manifest. CRD schema doesn't support self-referential types for now (see https://github.com/kubernetes/kubernetes/issues/62872). We have to use an alternative type to circumvent the limitation. The downside is that the Kube API can't validate the data beyond the fact that it is a valid JSON representation. MuteTimeIntervals is a list of TimeInterval names that will mute this route when matched.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"receiver": schema.StringAttribute{
								Description:         "Name of the receiver for this route. If not empty, it should be listed in the 'receivers' field.",
								MarkdownDescription: "Name of the receiver for this route. If not empty, it should be listed in the 'receivers' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repeat_interval": schema.StringAttribute{
								Description:         "How long to wait before repeating the last notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '4h'",
								MarkdownDescription: "How long to wait before repeating the last notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '4h'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"routes": schema.ListAttribute{
								Description:         "Child routes.",
								MarkdownDescription: "Child routes.",
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

					"time_intervals": schema.ListNestedAttribute{
						Description:         "List of TimeInterval specifying when the routes should be muted or active.",
						MarkdownDescription: "List of TimeInterval specifying when the routes should be muted or active.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the time interval.",
									MarkdownDescription: "Name of the time interval.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"time_intervals": schema.ListNestedAttribute{
									Description:         "TimeIntervals is a list of TimePeriod.",
									MarkdownDescription: "TimeIntervals is a list of TimePeriod.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"days_of_month": schema.ListNestedAttribute{
												Description:         "DaysOfMonth is a list of DayOfMonthRange",
												MarkdownDescription: "DaysOfMonth is a list of DayOfMonthRange",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"end": schema.Int64Attribute{
															Description:         "End of the inclusive range",
															MarkdownDescription: "End of the inclusive range",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(-31),
																int64validator.AtMost(31),
															},
														},

														"start": schema.Int64Attribute{
															Description:         "Start of the inclusive range",
															MarkdownDescription: "Start of the inclusive range",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.Int64{
																int64validator.AtLeast(-31),
																int64validator.AtMost(31),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"months": schema.ListAttribute{
												Description:         "Months is a list of MonthRange",
												MarkdownDescription: "Months is a list of MonthRange",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"times": schema.ListNestedAttribute{
												Description:         "Times is a list of TimeRange",
												MarkdownDescription: "Times is a list of TimeRange",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"end_time": schema.StringAttribute{
															Description:         "EndTime is the end time in 24hr format.",
															MarkdownDescription: "EndTime is the end time in 24hr format.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^((([01][0-9])|(2[0-3])):[0-5][0-9])$|(^24:00$)`), ""),
															},
														},

														"start_time": schema.StringAttribute{
															Description:         "StartTime is the start time in 24hr format.",
															MarkdownDescription: "StartTime is the start time in 24hr format.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.RegexMatches(regexp.MustCompile(`^((([01][0-9])|(2[0-3])):[0-5][0-9])$|(^24:00$)`), ""),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"weekdays": schema.ListAttribute{
												Description:         "Weekdays is a list of WeekdayRange",
												MarkdownDescription: "Weekdays is a list of WeekdayRange",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"years": schema.ListAttribute{
												Description:         "Years is a list of YearRange",
												MarkdownDescription: "Years is a list of YearRange",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Beta1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Beta1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_monitoring_coreos_com_alertmanager_config_v1beta1")

	var model MonitoringCoreosComAlertmanagerConfigV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("monitoring.coreos.com/v1beta1")
	model.Kind = pointer.String("AlertmanagerConfig")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1beta1", Resource: "AlertmanagerConfig"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse MonitoringCoreosComAlertmanagerConfigV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Beta1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_coreos_com_alertmanager_config_v1beta1")

	var data MonitoringCoreosComAlertmanagerConfigV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1beta1", Resource: "AlertmanagerConfig"}).
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

	var readResponse MonitoringCoreosComAlertmanagerConfigV1Beta1ResourceData
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

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Beta1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_monitoring_coreos_com_alertmanager_config_v1beta1")

	var model MonitoringCoreosComAlertmanagerConfigV1Beta1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("monitoring.coreos.com/v1beta1")
	model.Kind = pointer.String("AlertmanagerConfig")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1beta1", Resource: "AlertmanagerConfig"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse MonitoringCoreosComAlertmanagerConfigV1Beta1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Beta1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_monitoring_coreos_com_alertmanager_config_v1beta1")

	var data MonitoringCoreosComAlertmanagerConfigV1Beta1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "monitoring.coreos.com", Version: "v1beta1", Resource: "AlertmanagerConfig"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Beta1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
