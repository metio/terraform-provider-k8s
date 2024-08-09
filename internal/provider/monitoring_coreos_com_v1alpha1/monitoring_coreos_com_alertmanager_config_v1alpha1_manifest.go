/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package monitoring_coreos_com_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &MonitoringCoreosComAlertmanagerConfigV1Alpha1Manifest{}
)

func NewMonitoringCoreosComAlertmanagerConfigV1Alpha1Manifest() datasource.DataSource {
	return &MonitoringCoreosComAlertmanagerConfigV1Alpha1Manifest{}
}

type MonitoringCoreosComAlertmanagerConfigV1Alpha1Manifest struct{}

type MonitoringCoreosComAlertmanagerConfigV1Alpha1ManifestData struct {
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
		InhibitRules *[]struct {
			Equal       *[]string `tfsdk:"equal" json:"equal,omitempty"`
			SourceMatch *[]struct {
				MatchType *string `tfsdk:"match_type" json:"matchType,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Regex     *bool   `tfsdk:"regex" json:"regex,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"source_match" json:"sourceMatch,omitempty"`
			TargetMatch *[]struct {
				MatchType *string `tfsdk:"match_type" json:"matchType,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Regex     *bool   `tfsdk:"regex" json:"regex,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"target_match" json:"targetMatch,omitempty"`
		} `tfsdk:"inhibit_rules" json:"inhibitRules,omitempty"`
		MuteTimeIntervals *[]struct {
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
		} `tfsdk:"mute_time_intervals" json:"muteTimeIntervals,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"auth_password" json:"authPassword,omitempty"`
				AuthSecret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
					MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
					MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				To *string `tfsdk:"to" json:"to,omitempty"`
			} `tfsdk:"email_configs" json:"emailConfigs,omitempty"`
			MsteamsConfigs *[]struct {
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				Summary      *string `tfsdk:"summary" json:"summary,omitempty"`
				Text         *string `tfsdk:"text" json:"text,omitempty"`
				Title        *string `tfsdk:"title" json:"title,omitempty"`
				WebhookUrl   *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"webhook_url" json:"webhookUrl,omitempty"`
			} `tfsdk:"msteams_configs" json:"msteamsConfigs,omitempty"`
			Name            *string `tfsdk:"name" json:"name,omitempty"`
			OpsgenieConfigs *[]struct {
				Actions *string `tfsdk:"actions" json:"actions,omitempty"`
				ApiKey  *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
				UpdateAlerts *bool   `tfsdk:"update_alerts" json:"updateAlerts,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"routing_key" json:"routingKey,omitempty"`
				SendResolved *bool `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				ServiceKey   *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"service_key" json:"serviceKey,omitempty"`
				Severity *string `tfsdk:"severity" json:"severity,omitempty"`
				Source   *string `tfsdk:"source" json:"source,omitempty"`
				Url      *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"pagerduty_configs" json:"pagerdutyConfigs,omitempty"`
			PushoverConfigs *[]struct {
				Device     *string `tfsdk:"device" json:"device,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"token" json:"token,omitempty"`
				TokenFile *string `tfsdk:"token_file" json:"tokenFile,omitempty"`
				Ttl       *string `tfsdk:"ttl" json:"ttl,omitempty"`
				Url       *string `tfsdk:"url" json:"url,omitempty"`
				UrlTitle  *string `tfsdk:"url_title" json:"urlTitle,omitempty"`
				UserKey   *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"user_key" json:"userKey,omitempty"`
				UserKeyFile *string `tfsdk:"user_key_file" json:"userKeyFile,omitempty"`
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
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				Message      *string `tfsdk:"message" json:"message,omitempty"`
				ParseMode    *string `tfsdk:"parse_mode" json:"parseMode,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
			} `tfsdk:"telegram_configs" json:"telegramConfigs,omitempty"`
			VictoropsConfigs *[]struct {
				ApiKey *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" json:"httpConfig,omitempty"`
				MaxAlerts    *int64  `tfsdk:"max_alerts" json:"maxAlerts,omitempty"`
				SendResolved *bool   `tfsdk:"send_resolved" json:"sendResolved,omitempty"`
				Url          *string `tfsdk:"url" json:"url,omitempty"`
				UrlSecret    *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"url_secret" json:"urlSecret,omitempty"`
			} `tfsdk:"webhook_configs" json:"webhookConfigs,omitempty"`
			WechatConfigs *[]struct {
				AgentID   *string `tfsdk:"agent_id" json:"agentID,omitempty"`
				ApiSecret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
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
						EndpointParams       *map[string]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						NoProxy              *string            `tfsdk:"no_proxy" json:"noProxy,omitempty"`
						ProxyConnectHeader   *map[string]string `tfsdk:"proxy_connect_header" json:"proxyConnectHeader,omitempty"`
						ProxyFromEnvironment *bool              `tfsdk:"proxy_from_environment" json:"proxyFromEnvironment,omitempty"`
						ProxyUrl             *string            `tfsdk:"proxy_url" json:"proxyUrl,omitempty"`
						Scopes               *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TlsConfig            *struct {
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
							MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
							MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
							ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
						} `tfsdk:"tls_config" json:"tlsConfig,omitempty"`
						TokenUrl *string `tfsdk:"token_url" json:"tokenUrl,omitempty"`
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
						MaxVersion *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion *string `tfsdk:"min_version" json:"minVersion,omitempty"`
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
				Regex     *bool   `tfsdk:"regex" json:"regex,omitempty"`
				Value     *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"matchers" json:"matchers,omitempty"`
			MuteTimeIntervals *[]string `tfsdk:"mute_time_intervals" json:"muteTimeIntervals,omitempty"`
			Receiver          *string   `tfsdk:"receiver" json:"receiver,omitempty"`
			RepeatInterval    *string   `tfsdk:"repeat_interval" json:"repeatInterval,omitempty"`
			Routes            *[]string `tfsdk:"routes" json:"routes,omitempty"`
		} `tfsdk:"route" json:"route,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_monitoring_coreos_com_alertmanager_config_v1alpha1_manifest"
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AlertmanagerConfig configures the Prometheus Alertmanager,specifying how alerts should be grouped, inhibited and notified to external systems.",
		MarkdownDescription: "AlertmanagerConfig configures the Prometheus Alertmanager,specifying how alerts should be grouped, inhibited and notified to external systems.",
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
				Description:         "AlertmanagerConfigSpec is a specification of the desired behavior of theAlertmanager configuration.By default, the Alertmanager configuration only applies to alerts for whichthe 'namespace' label is equal to the namespace of the AlertmanagerConfigresource (see the '.spec.alertmanagerConfigMatcherStrategy' field of theAlertmanager CRD).",
				MarkdownDescription: "AlertmanagerConfigSpec is a specification of the desired behavior of theAlertmanager configuration.By default, the Alertmanager configuration only applies to alerts for whichthe 'namespace' label is equal to the namespace of the AlertmanagerConfigresource (see the '.spec.alertmanagerConfigMatcherStrategy' field of theAlertmanager CRD).",
				Attributes: map[string]schema.Attribute{
					"inhibit_rules": schema.ListNestedAttribute{
						Description:         "List of inhibition rules. The rules will only apply to alerts matchingthe resource's namespace.",
						MarkdownDescription: "List of inhibition rules. The rules will only apply to alerts matchingthe resource's namespace.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"equal": schema.ListAttribute{
									Description:         "Labels that must have an equal value in the source and target alert forthe inhibition to take effect.",
									MarkdownDescription: "Labels that must have an equal value in the source and target alert forthe inhibition to take effect.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"source_match": schema.ListNestedAttribute{
									Description:         "Matchers for which one or more alerts have to exist for the inhibitionto take effect. The operator enforces that the alert matches theresource's namespace.",
									MarkdownDescription: "Matchers for which one or more alerts have to exist for the inhibitionto take effect. The operator enforces that the alert matches theresource's namespace.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_type": schema.StringAttribute{
												Description:         "Match operation available with AlertManager >= v0.22.0 andtakes precedence over Regex (deprecated) if non-empty.",
												MarkdownDescription: "Match operation available with AlertManager >= v0.22.0 andtakes precedence over Regex (deprecated) if non-empty.",
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

											"regex": schema.BoolAttribute{
												Description:         "Whether to match on equality (false) or regular-expression (true).Deprecated: for AlertManager >= v0.22.0, 'matchType' should be used instead.",
												MarkdownDescription: "Whether to match on equality (false) or regular-expression (true).Deprecated: for AlertManager >= v0.22.0, 'matchType' should be used instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
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
									Description:         "Matchers that have to be fulfilled in the alerts to be muted. Theoperator enforces that the alert matches the resource's namespace.",
									MarkdownDescription: "Matchers that have to be fulfilled in the alerts to be muted. Theoperator enforces that the alert matches the resource's namespace.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"match_type": schema.StringAttribute{
												Description:         "Match operation available with AlertManager >= v0.22.0 andtakes precedence over Regex (deprecated) if non-empty.",
												MarkdownDescription: "Match operation available with AlertManager >= v0.22.0 andtakes precedence over Regex (deprecated) if non-empty.",
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

											"regex": schema.BoolAttribute{
												Description:         "Whether to match on equality (false) or regular-expression (true).Deprecated: for AlertManager >= v0.22.0, 'matchType' should be used instead.",
												MarkdownDescription: "Whether to match on equality (false) or regular-expression (true).Deprecated: for AlertManager >= v0.22.0, 'matchType' should be used instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
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

					"mute_time_intervals": schema.ListNestedAttribute{
						Description:         "List of MuteTimeInterval specifying when the routes should be muted.",
						MarkdownDescription: "List of MuteTimeInterval specifying when the routes should be muted.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the time interval",
									MarkdownDescription: "Name of the time interval",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"time_intervals": schema.ListNestedAttribute{
									Description:         "TimeIntervals is a list of TimeInterval",
									MarkdownDescription: "TimeIntervals is a list of TimeInterval",
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

					"receivers": schema.ListNestedAttribute{
						Description:         "List of receivers.",
						MarkdownDescription: "List of receivers.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"discord_configs": schema.ListNestedAttribute{
									Description:         "List of Discord configurations.",
									MarkdownDescription: "List of Discord configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_url": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the Discord webhook URL.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the Discord webhook URL.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
												Description:         "The secret's key that contains the password to use for authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the password to use for authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

											"auth_secret": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the CRAM-MD5 secret.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the CRAM-MD5 secret.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
												Description:         "Further headers email header key/value pairs. Overrides any headerspreviously set by the notification implementation.",
												MarkdownDescription: "Further headers email header key/value pairs. Overrides any headerspreviously set by the notification implementation.",
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
												Description:         "The SMTP TLS requirement.Note that Go does not support unencrypted connections to remote SMTP endpoints.",
												MarkdownDescription: "The SMTP TLS requirement.Note that Go does not support unencrypted connections to remote SMTP endpoints.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

													"max_version": schema.StringAttribute{
														Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
														MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
														},
													},

													"min_version": schema.StringAttribute{
														Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
														MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
														},
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

								"msteams_configs": schema.ListNestedAttribute{
									Description:         "List of MSTeams configurations.It requires Alertmanager >= 0.26.0.",
									MarkdownDescription: "List of MSTeams configurations.It requires Alertmanager >= 0.26.0.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether to notify about resolved alerts.",
												MarkdownDescription: "Whether to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"summary": schema.StringAttribute{
												Description:         "Message summary template.It requires Alertmanager >= 0.27.0.",
												MarkdownDescription: "Message summary template.It requires Alertmanager >= 0.27.0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"text": schema.StringAttribute{
												Description:         "Message body template.",
												MarkdownDescription: "Message body template.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title": schema.StringAttribute{
												Description:         "Message title template.",
												MarkdownDescription: "Message title template.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"webhook_url": schema.SingleNestedAttribute{
												Description:         "MSTeams webhook URL.",
												MarkdownDescription: "MSTeams webhook URL.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
												Description:         "The secret's key that contains the OpsGenie API key.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the OpsGenie API key.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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

											"update_alerts": schema.BoolAttribute{
												Description:         "Whether to update message and description of the alert in OpsGenie if it already existsBy default, the alert is never updated in OpsGenie, the new message only appears in activity log.",
												MarkdownDescription: "Whether to update message and description of the alert in OpsGenie if it already existsBy default, the alert is never updated in OpsGenie, the new message only appears in activity log.",
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
												Description:         "The secret's key that contains the PagerDuty integration key (when usingEvents API v2). Either this field or 'serviceKey' needs to be defined.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the PagerDuty integration key (when usingEvents API v2). Either this field or 'serviceKey' needs to be defined.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

											"send_resolved": schema.BoolAttribute{
												Description:         "Whether or not to notify about resolved alerts.",
												MarkdownDescription: "Whether or not to notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the PagerDuty service key (when usingintegration type 'Prometheus'). Either this field or 'routingKey' needs tobe defined.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the PagerDuty service key (when usingintegration type 'Prometheus'). Either this field or 'routingKey' needs tobe defined.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

											"severity": schema.StringAttribute{
												Description:         "Severity of the incident.",
												MarkdownDescription: "Severity of the incident.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source": schema.StringAttribute{
												Description:         "Unique location of the affected system.",
												MarkdownDescription: "Unique location of the affected system.",
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
											"device": schema.StringAttribute{
												Description:         "The name of a device to send the notification to",
												MarkdownDescription: "The name of a device to send the notification to",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"expire": schema.StringAttribute{
												Description:         "How long your notification will continue to be retried for, unless the useracknowledges the notification.",
												MarkdownDescription: "How long your notification will continue to be retried for, unless the useracknowledges the notification.",
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
												Description:         "How often the Pushover servers will send the same notification to the user.Must be at least 30 seconds.",
												MarkdownDescription: "How often the Pushover servers will send the same notification to the user.Must be at least 30 seconds.",
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
												Description:         "The secret's key that contains the registered application's API token, see https://pushover.net/apps.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.Either 'token' or 'tokenFile' is required.",
												MarkdownDescription: "The secret's key that contains the registered application's API token, see https://pushover.net/apps.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.Either 'token' or 'tokenFile' is required.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

											"token_file": schema.StringAttribute{
												Description:         "The token file that contains the registered application's API token, see https://pushover.net/apps.Either 'token' or 'tokenFile' is required.It requires Alertmanager >= v0.26.0.",
												MarkdownDescription: "The token file that contains the registered application's API token, see https://pushover.net/apps.Either 'token' or 'tokenFile' is required.It requires Alertmanager >= v0.26.0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ttl": schema.StringAttribute{
												Description:         "The time to live definition for the alert notification",
												MarkdownDescription: "The time to live definition for the alert notification",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(0|(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?)$`), ""),
												},
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
												Description:         "The secret's key that contains the recipient user's user key.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.Either 'userKey' or 'userKeyFile' is required.",
												MarkdownDescription: "The secret's key that contains the recipient user's user key.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.Either 'userKey' or 'userKeyFile' is required.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

											"user_key_file": schema.StringAttribute{
												Description:         "The user key file that contains the recipient user's user key.Either 'userKey' or 'userKeyFile' is required.It requires Alertmanager >= v0.26.0.",
												MarkdownDescription: "The user key file that contains the recipient user's user key.Either 'userKey' or 'userKeyFile' is required.It requires Alertmanager >= v0.26.0.",
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
															Description:         "SlackConfirmationField protect users from destructive actions orparticularly distinguished decisions by asking them to confirm their buttonclick one more time.See https://api.slack.com/docs/interactive-message-field-guide#confirmation_fieldsfor more information.",
															MarkdownDescription: "SlackConfirmationField protect users from destructive actions orparticularly distinguished decisions by asking them to confirm their buttonclick one more time.See https://api.slack.com/docs/interactive-message-field-guide#confirmation_fieldsfor more information.",
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
												Description:         "The secret's key that contains the Slack webhook URL.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the Slack webhook URL.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
												Description:         "The SNS API URL i.e. https://sns.us-east-2.amazonaws.com.If not specified, the SNS API URL from the SNS SDK will be used.",
												MarkdownDescription: "The SNS API URL i.e. https://sns.us-east-2.amazonaws.com.If not specified, the SNS API URL from the SNS SDK will be used.",
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
												Description:         "Phone number if message is delivered via SMS in E.164 format.If you don't specify this value, you must specify a value for the TopicARN or TargetARN.",
												MarkdownDescription: "Phone number if message is delivered via SMS in E.164 format.If you don't specify this value, you must specify a value for the TopicARN or TargetARN.",
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
														Description:         "AccessKey is the AWS API key. If not specified, the environment variable'AWS_ACCESS_KEY_ID' is used.",
														MarkdownDescription: "AccessKey is the AWS API key. If not specified, the environment variable'AWS_ACCESS_KEY_ID' is used.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "SecretKey is the AWS API secret. If not specified, the environmentvariable 'AWS_SECRET_ACCESS_KEY' is used.",
														MarkdownDescription: "SecretKey is the AWS API secret. If not specified, the environmentvariable 'AWS_SECRET_ACCESS_KEY' is used.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
												Description:         "The  mobile platform endpoint ARN if message is delivered via mobile notifications.If you don't specify this value, you must specify a value for the topic_arn or PhoneNumber.",
												MarkdownDescription: "The  mobile platform endpoint ARN if message is delivered via mobile notifications.If you don't specify this value, you must specify a value for the topic_arn or PhoneNumber.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"topic_arn": schema.StringAttribute{
												Description:         "SNS topic ARN, i.e. arn:aws:sns:us-east-2:698519295917:My-TopicIf you don't specify this value, you must specify a value for the PhoneNumber or TargetARN.",
												MarkdownDescription: "SNS topic ARN, i.e. arn:aws:sns:us-east-2:698519295917:My-TopicIf you don't specify this value, you must specify a value for the PhoneNumber or TargetARN.",
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
												Description:         "The Telegram API URL i.e. https://api.telegram.org.If not specified, default API URL will be used.",
												MarkdownDescription: "The Telegram API URL i.e. https://api.telegram.org.If not specified, default API URL will be used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"bot_token": schema.SingleNestedAttribute{
												Description:         "Telegram bot token. It is mutually exclusive with 'botTokenFile'.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.Either 'botToken' or 'botTokenFile' is required.",
												MarkdownDescription: "Telegram bot token. It is mutually exclusive with 'botTokenFile'.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.Either 'botToken' or 'botTokenFile' is required.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

											"bot_token_file": schema.StringAttribute{
												Description:         "File to read the Telegram bot token from. It is mutually exclusive with 'botToken'.Either 'botToken' or 'botTokenFile' is required.It requires Alertmanager >= v0.26.0.",
												MarkdownDescription: "File to read the Telegram bot token from. It is mutually exclusive with 'botToken'.Either 'botToken' or 'botTokenFile' is required.It requires Alertmanager >= v0.26.0.",
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
												Description:         "The secret's key that contains the API key to use when talking to the VictorOps API.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the API key to use when talking to the VictorOps API.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
												Description:         "The Webex Teams API URL i.e. https://webexapis.com/v1/messagesProvide if different from the default API URL.",
												MarkdownDescription: "The Webex Teams API URL i.e. https://webexapis.com/v1/messagesProvide if different from the default API URL.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^https?://.+$`), ""),
												},
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "The HTTP client's configuration.You must supply the bot token via the 'httpConfig.authorization' field.",
												MarkdownDescription: "The HTTP client's configuration.You must supply the bot token via the 'httpConfig.authorization' field.",
												Attributes: map[string]schema.Attribute{
													"authorization": schema.SingleNestedAttribute{
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
												Description:         "The URL to send HTTP POST requests to. 'urlSecret' takes precedence over'url'. One of 'urlSecret' and 'url' should be defined.",
												MarkdownDescription: "The URL to send HTTP POST requests to. 'urlSecret' takes precedence over'url'. One of 'urlSecret' and 'url' should be defined.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url_secret": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the webhook URL to send HTTP requests to.'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url'should be defined.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the webhook URL to send HTTP requests to.'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url'should be defined.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
												Description:         "The secret's key that contains the WeChat API key.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the WeChat API key.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
														MarkdownDescription: "Authorization header configuration for the client.This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
																MarkdownDescription: "Defines the authentication type. The value is case-insensitive.'Basic' is not a supported value.Default: 'Bearer'",
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
														Description:         "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														MarkdownDescription: "BasicAuth for the client.This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "'password' specifies a key of a Secret containing the password forauthentication.",
																MarkdownDescription: "'password' specifies a key of a Secret containing the password forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'username' specifies a key of a Secret containing the username forauthentication.",
																MarkdownDescription: "'username' specifies a key of a Secret containing the username forauthentication.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
														Description:         "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														MarkdownDescription: "The secret's key that contains the bearer token to be used by the clientfor authentication.The secret needs to be in the same namespace as the AlertmanagerConfigobject and accessible by the Prometheus Operator.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
																MarkdownDescription: "'clientId' specifies a key of a Secret or ConfigMap containing theOAuth2 client's ID.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																MarkdownDescription: "'clientSecret' specifies a key of a Secret containing the OAuth2client's secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																Description:         "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																MarkdownDescription: "'endpointParams' configures the HTTP parameters to append to the tokenURL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"no_proxy": schema.StringAttribute{
																Description:         "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'noProxy' is a comma-separated string that can contain IPs, CIDR notation, domain namesthat should be excluded from proxying. IP and domain names cancontain port numbers.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_connect_header": schema.MapAttribute{
																Description:         "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "ProxyConnectHeader optionally specifies headers to send toproxies during CONNECT requests.It requires Prometheus >= v2.43.0.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_from_environment": schema.BoolAttribute{
																Description:         "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "Whether to use the proxy configuration defined by environment variables (HTTP_PROXY, HTTPS_PROXY, and NO_PROXY).If unset, Prometheus uses its default value.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"proxy_url": schema.StringAttribute{
																Description:         "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "'proxyURL' defines the HTTP proxy server to use.It requires Prometheus >= v2.43.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.RegexMatches(regexp.MustCompile(`^http(s)?://.+$`), ""),
																},
															},

															"scopes": schema.ListAttribute{
																Description:         "'scopes' defines the OAuth2 scopes used for the token request.",
																MarkdownDescription: "'scopes' defines the OAuth2 scopes used for the token request.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"tls_config": schema.SingleNestedAttribute{
																Description:         "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
																MarkdownDescription: "TLS configuration to use when connecting to the OAuth2 server.It requires Prometheus >= v2.43.0.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																						Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																						MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

																	"max_version": schema.StringAttribute{
																		Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
																	},

																	"min_version": schema.StringAttribute{
																		Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																		},
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

															"token_url": schema.StringAttribute{
																Description:         "'tokenURL' configures the URL to fetch the token from.",
																MarkdownDescription: "'tokenURL' configures the URL to fetch the token from.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																				Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																				MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
																		Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
																		MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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

															"max_version": schema.StringAttribute{
																Description:         "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																MarkdownDescription: "Maximum acceptable TLS version.It requires Prometheus >= v2.41.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
															},

															"min_version": schema.StringAttribute{
																Description:         "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																MarkdownDescription: "Minimum acceptable TLS version.It requires Prometheus >= v2.35.0.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("TLS10", "TLS11", "TLS12", "TLS13"),
																},
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
						Description:         "The Alertmanager route definition for alerts matching the resource'snamespace. If present, it will be added to the generated Alertmanagerconfiguration as a first-level route.",
						MarkdownDescription: "The Alertmanager route definition for alerts matching the resource'snamespace. If present, it will be added to the generated Alertmanagerconfiguration as a first-level route.",
						Attributes: map[string]schema.Attribute{
							"active_time_intervals": schema.ListAttribute{
								Description:         "ActiveTimeIntervals is a list of MuteTimeInterval names when this route should be active.",
								MarkdownDescription: "ActiveTimeIntervals is a list of MuteTimeInterval names when this route should be active.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"continue": schema.BoolAttribute{
								Description:         "Boolean indicating whether an alert should continue matching subsequentsibling nodes. It will always be overridden to true for the first-levelroute by the Prometheus operator.",
								MarkdownDescription: "Boolean indicating whether an alert should continue matching subsequentsibling nodes. It will always be overridden to true for the first-levelroute by the Prometheus operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_by": schema.ListAttribute{
								Description:         "List of labels to group by.Labels must not be repeated (unique list).Special label '...' (aggregate by all possible labels), if provided, must be the only element in the list.",
								MarkdownDescription: "List of labels to group by.Labels must not be repeated (unique list).Special label '...' (aggregate by all possible labels), if provided, must be the only element in the list.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_interval": schema.StringAttribute{
								Description:         "How long to wait before sending an updated notification.Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$'Example: '5m'",
								MarkdownDescription: "How long to wait before sending an updated notification.Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$'Example: '5m'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_wait": schema.StringAttribute{
								Description:         "How long to wait before sending the initial notification.Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$'Example: '30s'",
								MarkdownDescription: "How long to wait before sending the initial notification.Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$'Example: '30s'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"matchers": schema.ListNestedAttribute{
								Description:         "List of matchers that the alert's labels should match. For the firstlevel route, the operator removes any existing equality and regexpmatcher on the 'namespace' label and adds a 'namespace: <objectnamespace>' matcher.",
								MarkdownDescription: "List of matchers that the alert's labels should match. For the firstlevel route, the operator removes any existing equality and regexpmatcher on the 'namespace' label and adds a 'namespace: <objectnamespace>' matcher.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"match_type": schema.StringAttribute{
											Description:         "Match operation available with AlertManager >= v0.22.0 andtakes precedence over Regex (deprecated) if non-empty.",
											MarkdownDescription: "Match operation available with AlertManager >= v0.22.0 andtakes precedence over Regex (deprecated) if non-empty.",
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

										"regex": schema.BoolAttribute{
											Description:         "Whether to match on equality (false) or regular-expression (true).Deprecated: for AlertManager >= v0.22.0, 'matchType' should be used instead.",
											MarkdownDescription: "Whether to match on equality (false) or regular-expression (true).Deprecated: for AlertManager >= v0.22.0, 'matchType' should be used instead.",
											Required:            false,
											Optional:            true,
											Computed:            false,
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
								Description:         "Note: this comment applies to the field definition above but appearsbelow otherwise it gets included in the generated manifest.CRD schema doesn't support self-referential types for now (seehttps://github.com/kubernetes/kubernetes/issues/62872). We have to usean alternative type to circumvent the limitation. The downside is thatthe Kube API can't validate the data beyond the fact that it is a validJSON representation.MuteTimeIntervals is a list of MuteTimeInterval names that will mute this route when matched,",
								MarkdownDescription: "Note: this comment applies to the field definition above but appearsbelow otherwise it gets included in the generated manifest.CRD schema doesn't support self-referential types for now (seehttps://github.com/kubernetes/kubernetes/issues/62872). We have to usean alternative type to circumvent the limitation. The downside is thatthe Kube API can't validate the data beyond the fact that it is a validJSON representation.MuteTimeIntervals is a list of MuteTimeInterval names that will mute this route when matched,",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"receiver": schema.StringAttribute{
								Description:         "Name of the receiver for this route. If not empty, it should be listed inthe 'receivers' field.",
								MarkdownDescription: "Name of the receiver for this route. If not empty, it should be listed inthe 'receivers' field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repeat_interval": schema.StringAttribute{
								Description:         "How long to wait before repeating the last notification.Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$'Example: '4h'",
								MarkdownDescription: "How long to wait before repeating the last notification.Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$'Example: '4h'",
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
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_coreos_com_alertmanager_config_v1alpha1_manifest")

	var model MonitoringCoreosComAlertmanagerConfigV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("monitoring.coreos.com/v1alpha1")
	model.Kind = pointer.String("AlertmanagerConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
