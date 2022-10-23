/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource)(nil)
)

type MonitoringCoreosComAlertmanagerConfigV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type MonitoringCoreosComAlertmanagerConfigV1Alpha1GoModel struct {
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
		InhibitRules *[]struct {
			Equal *[]string `tfsdk:"equal" yaml:"equal,omitempty"`

			SourceMatch *[]struct {
				MatchType *string `tfsdk:"match_type" yaml:"matchType,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"source_match" yaml:"sourceMatch,omitempty"`

			TargetMatch *[]struct {
				MatchType *string `tfsdk:"match_type" yaml:"matchType,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"target_match" yaml:"targetMatch,omitempty"`
		} `tfsdk:"inhibit_rules" yaml:"inhibitRules,omitempty"`

		MuteTimeIntervals *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			TimeIntervals *[]struct {
				DaysOfMonth *[]struct {
					End *int64 `tfsdk:"end" yaml:"end,omitempty"`

					Start *int64 `tfsdk:"start" yaml:"start,omitempty"`
				} `tfsdk:"days_of_month" yaml:"daysOfMonth,omitempty"`

				Months *[]string `tfsdk:"months" yaml:"months,omitempty"`

				Times *[]struct {
					EndTime *string `tfsdk:"end_time" yaml:"endTime,omitempty"`

					StartTime *string `tfsdk:"start_time" yaml:"startTime,omitempty"`
				} `tfsdk:"times" yaml:"times,omitempty"`

				Weekdays *[]string `tfsdk:"weekdays" yaml:"weekdays,omitempty"`

				Years *[]string `tfsdk:"years" yaml:"years,omitempty"`
			} `tfsdk:"time_intervals" yaml:"timeIntervals,omitempty"`
		} `tfsdk:"mute_time_intervals" yaml:"muteTimeIntervals,omitempty"`

		Receivers *[]struct {
			EmailConfigs *[]struct {
				AuthIdentity *string `tfsdk:"auth_identity" yaml:"authIdentity,omitempty"`

				AuthPassword *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"auth_password" yaml:"authPassword,omitempty"`

				AuthSecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"auth_secret" yaml:"authSecret,omitempty"`

				AuthUsername *string `tfsdk:"auth_username" yaml:"authUsername,omitempty"`

				From *string `tfsdk:"from" yaml:"from,omitempty"`

				Headers *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"headers" yaml:"headers,omitempty"`

				Hello *string `tfsdk:"hello" yaml:"hello,omitempty"`

				Html *string `tfsdk:"html" yaml:"html,omitempty"`

				RequireTLS *bool `tfsdk:"require_tls" yaml:"requireTLS,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				Smarthost *string `tfsdk:"smarthost" yaml:"smarthost,omitempty"`

				Text *string `tfsdk:"text" yaml:"text,omitempty"`

				TlsConfig *struct {
					Ca *struct {
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
					} `tfsdk:"ca" yaml:"ca,omitempty"`

					Cert *struct {
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
					} `tfsdk:"cert" yaml:"cert,omitempty"`

					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

					KeySecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

					ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
				} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`

				To *string `tfsdk:"to" yaml:"to,omitempty"`
			} `tfsdk:"email_configs" yaml:"emailConfigs,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			OpsgenieConfigs *[]struct {
				Actions *string `tfsdk:"actions" yaml:"actions,omitempty"`

				ApiKey *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"api_key" yaml:"apiKey,omitempty"`

				ApiURL *string `tfsdk:"api_url" yaml:"apiURL,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Details *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"details" yaml:"details,omitempty"`

				Entity *string `tfsdk:"entity" yaml:"entity,omitempty"`

				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				Message *string `tfsdk:"message" yaml:"message,omitempty"`

				Note *string `tfsdk:"note" yaml:"note,omitempty"`

				Priority *string `tfsdk:"priority" yaml:"priority,omitempty"`

				Responders *[]struct {
					Id *string `tfsdk:"id" yaml:"id,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					Username *string `tfsdk:"username" yaml:"username,omitempty"`
				} `tfsdk:"responders" yaml:"responders,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				Source *string `tfsdk:"source" yaml:"source,omitempty"`

				Tags *string `tfsdk:"tags" yaml:"tags,omitempty"`

				UpdateAlerts *bool `tfsdk:"update_alerts" yaml:"updateAlerts,omitempty"`
			} `tfsdk:"opsgenie_configs" yaml:"opsgenieConfigs,omitempty"`

			PagerdutyConfigs *[]struct {
				Class *string `tfsdk:"class" yaml:"class,omitempty"`

				Client *string `tfsdk:"client" yaml:"client,omitempty"`

				ClientURL *string `tfsdk:"client_url" yaml:"clientURL,omitempty"`

				Component *string `tfsdk:"component" yaml:"component,omitempty"`

				Description *string `tfsdk:"description" yaml:"description,omitempty"`

				Details *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"details" yaml:"details,omitempty"`

				Group *string `tfsdk:"group" yaml:"group,omitempty"`

				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				PagerDutyImageConfigs *[]struct {
					Alt *string `tfsdk:"alt" yaml:"alt,omitempty"`

					Href *string `tfsdk:"href" yaml:"href,omitempty"`

					Src *string `tfsdk:"src" yaml:"src,omitempty"`
				} `tfsdk:"pager_duty_image_configs" yaml:"pagerDutyImageConfigs,omitempty"`

				PagerDutyLinkConfigs *[]struct {
					Alt *string `tfsdk:"alt" yaml:"alt,omitempty"`

					Href *string `tfsdk:"href" yaml:"href,omitempty"`
				} `tfsdk:"pager_duty_link_configs" yaml:"pagerDutyLinkConfigs,omitempty"`

				RoutingKey *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"routing_key" yaml:"routingKey,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				ServiceKey *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"service_key" yaml:"serviceKey,omitempty"`

				Severity *string `tfsdk:"severity" yaml:"severity,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"pagerduty_configs" yaml:"pagerdutyConfigs,omitempty"`

			PushoverConfigs *[]struct {
				Expire *string `tfsdk:"expire" yaml:"expire,omitempty"`

				Html *bool `tfsdk:"html" yaml:"html,omitempty"`

				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				Message *string `tfsdk:"message" yaml:"message,omitempty"`

				Priority *string `tfsdk:"priority" yaml:"priority,omitempty"`

				Retry *string `tfsdk:"retry" yaml:"retry,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				Sound *string `tfsdk:"sound" yaml:"sound,omitempty"`

				Title *string `tfsdk:"title" yaml:"title,omitempty"`

				Token *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"token" yaml:"token,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				UrlTitle *string `tfsdk:"url_title" yaml:"urlTitle,omitempty"`

				UserKey *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"user_key" yaml:"userKey,omitempty"`
			} `tfsdk:"pushover_configs" yaml:"pushoverConfigs,omitempty"`

			SlackConfigs *[]struct {
				Actions *[]struct {
					Confirm *struct {
						DismissText *string `tfsdk:"dismiss_text" yaml:"dismissText,omitempty"`

						OkText *string `tfsdk:"ok_text" yaml:"okText,omitempty"`

						Text *string `tfsdk:"text" yaml:"text,omitempty"`

						Title *string `tfsdk:"title" yaml:"title,omitempty"`
					} `tfsdk:"confirm" yaml:"confirm,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Style *string `tfsdk:"style" yaml:"style,omitempty"`

					Text *string `tfsdk:"text" yaml:"text,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"actions" yaml:"actions,omitempty"`

				ApiURL *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"api_url" yaml:"apiURL,omitempty"`

				CallbackId *string `tfsdk:"callback_id" yaml:"callbackId,omitempty"`

				Channel *string `tfsdk:"channel" yaml:"channel,omitempty"`

				Color *string `tfsdk:"color" yaml:"color,omitempty"`

				Fallback *string `tfsdk:"fallback" yaml:"fallback,omitempty"`

				Fields *[]struct {
					Short *bool `tfsdk:"short" yaml:"short,omitempty"`

					Title *string `tfsdk:"title" yaml:"title,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"fields" yaml:"fields,omitempty"`

				Footer *string `tfsdk:"footer" yaml:"footer,omitempty"`

				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				IconEmoji *string `tfsdk:"icon_emoji" yaml:"iconEmoji,omitempty"`

				IconURL *string `tfsdk:"icon_url" yaml:"iconURL,omitempty"`

				ImageURL *string `tfsdk:"image_url" yaml:"imageURL,omitempty"`

				LinkNames *bool `tfsdk:"link_names" yaml:"linkNames,omitempty"`

				MrkdwnIn *[]string `tfsdk:"mrkdwn_in" yaml:"mrkdwnIn,omitempty"`

				Pretext *string `tfsdk:"pretext" yaml:"pretext,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				ShortFields *bool `tfsdk:"short_fields" yaml:"shortFields,omitempty"`

				Text *string `tfsdk:"text" yaml:"text,omitempty"`

				ThumbURL *string `tfsdk:"thumb_url" yaml:"thumbURL,omitempty"`

				Title *string `tfsdk:"title" yaml:"title,omitempty"`

				TitleLink *string `tfsdk:"title_link" yaml:"titleLink,omitempty"`

				Username *string `tfsdk:"username" yaml:"username,omitempty"`
			} `tfsdk:"slack_configs" yaml:"slackConfigs,omitempty"`

			SnsConfigs *[]struct {
				ApiURL *string `tfsdk:"api_url" yaml:"apiURL,omitempty"`

				Attributes *map[string]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				Message *string `tfsdk:"message" yaml:"message,omitempty"`

				PhoneNumber *string `tfsdk:"phone_number" yaml:"phoneNumber,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				Sigv4 *struct {
					AccessKey *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"access_key" yaml:"accessKey,omitempty"`

					Profile *string `tfsdk:"profile" yaml:"profile,omitempty"`

					Region *string `tfsdk:"region" yaml:"region,omitempty"`

					RoleArn *string `tfsdk:"role_arn" yaml:"roleArn,omitempty"`

					SecretKey *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key" yaml:"secretKey,omitempty"`
				} `tfsdk:"sigv4" yaml:"sigv4,omitempty"`

				Subject *string `tfsdk:"subject" yaml:"subject,omitempty"`

				TargetARN *string `tfsdk:"target_arn" yaml:"targetARN,omitempty"`

				TopicARN *string `tfsdk:"topic_arn" yaml:"topicARN,omitempty"`
			} `tfsdk:"sns_configs" yaml:"snsConfigs,omitempty"`

			TelegramConfigs *[]struct {
				ApiURL *string `tfsdk:"api_url" yaml:"apiURL,omitempty"`

				BotToken *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"bot_token" yaml:"botToken,omitempty"`

				ChatID *int64 `tfsdk:"chat_id" yaml:"chatID,omitempty"`

				DisableNotifications *bool `tfsdk:"disable_notifications" yaml:"disableNotifications,omitempty"`

				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				Message *string `tfsdk:"message" yaml:"message,omitempty"`

				ParseMode *string `tfsdk:"parse_mode" yaml:"parseMode,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`
			} `tfsdk:"telegram_configs" yaml:"telegramConfigs,omitempty"`

			VictoropsConfigs *[]struct {
				ApiKey *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"api_key" yaml:"apiKey,omitempty"`

				ApiUrl *string `tfsdk:"api_url" yaml:"apiUrl,omitempty"`

				CustomFields *[]struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"custom_fields" yaml:"customFields,omitempty"`

				EntityDisplayName *string `tfsdk:"entity_display_name" yaml:"entityDisplayName,omitempty"`

				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				MessageType *string `tfsdk:"message_type" yaml:"messageType,omitempty"`

				MonitoringTool *string `tfsdk:"monitoring_tool" yaml:"monitoringTool,omitempty"`

				RoutingKey *string `tfsdk:"routing_key" yaml:"routingKey,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				StateMessage *string `tfsdk:"state_message" yaml:"stateMessage,omitempty"`
			} `tfsdk:"victorops_configs" yaml:"victoropsConfigs,omitempty"`

			WebhookConfigs *[]struct {
				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				MaxAlerts *int64 `tfsdk:"max_alerts" yaml:"maxAlerts,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				UrlSecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"url_secret" yaml:"urlSecret,omitempty"`
			} `tfsdk:"webhook_configs" yaml:"webhookConfigs,omitempty"`

			WechatConfigs *[]struct {
				AgentID *string `tfsdk:"agent_id" yaml:"agentID,omitempty"`

				ApiSecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"api_secret" yaml:"apiSecret,omitempty"`

				ApiURL *string `tfsdk:"api_url" yaml:"apiURL,omitempty"`

				CorpID *string `tfsdk:"corp_id" yaml:"corpID,omitempty"`

				HttpConfig *struct {
					Authorization *struct {
						Credentials *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"credentials" yaml:"credentials,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"authorization" yaml:"authorization,omitempty"`

					BasicAuth *struct {
						Password *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"password" yaml:"password,omitempty"`

						Username *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"username" yaml:"username,omitempty"`
					} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

					BearerTokenSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" yaml:"bearerTokenSecret,omitempty"`

					FollowRedirects *bool `tfsdk:"follow_redirects" yaml:"followRedirects,omitempty"`

					Oauth2 *struct {
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
						} `tfsdk:"client_id" yaml:"clientId,omitempty"`

						ClientSecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

						EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

						Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

						TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

					ProxyURL *string `tfsdk:"proxy_url" yaml:"proxyURL,omitempty"`

					TlsConfig *struct {
						Ca *struct {
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
						} `tfsdk:"ca" yaml:"ca,omitempty"`

						Cert *struct {
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
						} `tfsdk:"cert" yaml:"cert,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						KeySecret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

						ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
					} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
				} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

				Message *string `tfsdk:"message" yaml:"message,omitempty"`

				MessageType *string `tfsdk:"message_type" yaml:"messageType,omitempty"`

				SendResolved *bool `tfsdk:"send_resolved" yaml:"sendResolved,omitempty"`

				ToParty *string `tfsdk:"to_party" yaml:"toParty,omitempty"`

				ToTag *string `tfsdk:"to_tag" yaml:"toTag,omitempty"`

				ToUser *string `tfsdk:"to_user" yaml:"toUser,omitempty"`
			} `tfsdk:"wechat_configs" yaml:"wechatConfigs,omitempty"`
		} `tfsdk:"receivers" yaml:"receivers,omitempty"`

		Route *struct {
			Continue *bool `tfsdk:"continue" yaml:"continue,omitempty"`

			GroupBy *[]string `tfsdk:"group_by" yaml:"groupBy,omitempty"`

			GroupInterval *string `tfsdk:"group_interval" yaml:"groupInterval,omitempty"`

			GroupWait *string `tfsdk:"group_wait" yaml:"groupWait,omitempty"`

			Matchers *[]struct {
				MatchType *string `tfsdk:"match_type" yaml:"matchType,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Regex *bool `tfsdk:"regex" yaml:"regex,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"matchers" yaml:"matchers,omitempty"`

			MuteTimeIntervals *[]string `tfsdk:"mute_time_intervals" yaml:"muteTimeIntervals,omitempty"`

			Receiver *string `tfsdk:"receiver" yaml:"receiver,omitempty"`

			RepeatInterval *string `tfsdk:"repeat_interval" yaml:"repeatInterval,omitempty"`

			Routes *[]string `tfsdk:"routes" yaml:"routes,omitempty"`
		} `tfsdk:"route" yaml:"route,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewMonitoringCoreosComAlertmanagerConfigV1Alpha1Resource() resource.Resource {
	return &MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource{}
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitoring_coreos_com_alertmanager_config_v1alpha1"
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "AlertmanagerConfig defines a namespaced AlertmanagerConfig to be aggregated across multiple namespaces configuring one Alertmanager cluster.",
		MarkdownDescription: "AlertmanagerConfig defines a namespaced AlertmanagerConfig to be aggregated across multiple namespaces configuring one Alertmanager cluster.",
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
				Description:         "AlertmanagerConfigSpec is a specification of the desired behavior of the Alertmanager configuration. By definition, the Alertmanager configuration only applies to alerts for which the 'namespace' label is equal to the namespace of the AlertmanagerConfig resource.",
				MarkdownDescription: "AlertmanagerConfigSpec is a specification of the desired behavior of the Alertmanager configuration. By definition, the Alertmanager configuration only applies to alerts for which the 'namespace' label is equal to the namespace of the AlertmanagerConfig resource.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"inhibit_rules": {
						Description:         "List of inhibition rules. The rules will only apply to alerts matching the resource's namespace.",
						MarkdownDescription: "List of inhibition rules. The rules will only apply to alerts matching the resource's namespace.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"equal": {
								Description:         "Labels that must have an equal value in the source and target alert for the inhibition to take effect.",
								MarkdownDescription: "Labels that must have an equal value in the source and target alert for the inhibition to take effect.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_match": {
								Description:         "Matchers for which one or more alerts have to exist for the inhibition to take effect. The operator enforces that the alert matches the resource's namespace.",
								MarkdownDescription: "Matchers for which one or more alerts have to exist for the inhibition to take effect. The operator enforces that the alert matches the resource's namespace.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_type": {
										Description:         "Match operation available with AlertManager >= v0.22.0 and takes precedence over Regex (deprecated) if non-empty.",
										MarkdownDescription: "Match operation available with AlertManager >= v0.22.0 and takes precedence over Regex (deprecated) if non-empty.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("!=", "=", "=~", "!~"),
										},
									},

									"name": {
										Description:         "Label to match.",
										MarkdownDescription: "Label to match.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),
										},
									},

									"regex": {
										Description:         "Whether to match on equality (false) or regular-expression (true). Deprecated as of AlertManager >= v0.22.0 where a user should use MatchType instead.",
										MarkdownDescription: "Whether to match on equality (false) or regular-expression (true). Deprecated as of AlertManager >= v0.22.0 where a user should use MatchType instead.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Label value to match.",
										MarkdownDescription: "Label value to match.",

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

							"target_match": {
								Description:         "Matchers that have to be fulfilled in the alerts to be muted. The operator enforces that the alert matches the resource's namespace.",
								MarkdownDescription: "Matchers that have to be fulfilled in the alerts to be muted. The operator enforces that the alert matches the resource's namespace.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_type": {
										Description:         "Match operation available with AlertManager >= v0.22.0 and takes precedence over Regex (deprecated) if non-empty.",
										MarkdownDescription: "Match operation available with AlertManager >= v0.22.0 and takes precedence over Regex (deprecated) if non-empty.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("!=", "=", "=~", "!~"),
										},
									},

									"name": {
										Description:         "Label to match.",
										MarkdownDescription: "Label to match.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),
										},
									},

									"regex": {
										Description:         "Whether to match on equality (false) or regular-expression (true). Deprecated as of AlertManager >= v0.22.0 where a user should use MatchType instead.",
										MarkdownDescription: "Whether to match on equality (false) or regular-expression (true). Deprecated as of AlertManager >= v0.22.0 where a user should use MatchType instead.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Label value to match.",
										MarkdownDescription: "Label value to match.",

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

					"mute_time_intervals": {
						Description:         "List of MuteTimeInterval specifying when the routes should be muted.",
						MarkdownDescription: "List of MuteTimeInterval specifying when the routes should be muted.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the time interval",
								MarkdownDescription: "Name of the time interval",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"time_intervals": {
								Description:         "TimeIntervals is a list of TimeInterval",
								MarkdownDescription: "TimeIntervals is a list of TimeInterval",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"days_of_month": {
										Description:         "DaysOfMonth is a list of DayOfMonthRange",
										MarkdownDescription: "DaysOfMonth is a list of DayOfMonthRange",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"end": {
												Description:         "End of the inclusive range",
												MarkdownDescription: "End of the inclusive range",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(-31),

													int64validator.AtMost(31),
												},
											},

											"start": {
												Description:         "Start of the inclusive range",
												MarkdownDescription: "Start of the inclusive range",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(-31),

													int64validator.AtMost(31),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"months": {
										Description:         "Months is a list of MonthRange",
										MarkdownDescription: "Months is a list of MonthRange",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"times": {
										Description:         "Times is a list of TimeRange",
										MarkdownDescription: "Times is a list of TimeRange",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"end_time": {
												Description:         "EndTime is the end time in 24hr format.",
												MarkdownDescription: "EndTime is the end time in 24hr format.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^((([01][0-9])|(2[0-3])):[0-5][0-9])$|(^24:00$)`), ""),
												},
											},

											"start_time": {
												Description:         "StartTime is the start time in 24hr format.",
												MarkdownDescription: "StartTime is the start time in 24hr format.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^((([01][0-9])|(2[0-3])):[0-5][0-9])$|(^24:00$)`), ""),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"weekdays": {
										Description:         "Weekdays is a list of WeekdayRange",
										MarkdownDescription: "Weekdays is a list of WeekdayRange",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"years": {
										Description:         "Years is a list of YearRange",
										MarkdownDescription: "Years is a list of YearRange",

										Type: types.ListType{ElemType: types.StringType},

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

					"receivers": {
						Description:         "List of receivers.",
						MarkdownDescription: "List of receivers.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"email_configs": {
								Description:         "List of Email configurations.",
								MarkdownDescription: "List of Email configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"auth_identity": {
										Description:         "The identity to use for authentication.",
										MarkdownDescription: "The identity to use for authentication.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auth_password": {
										Description:         "The secret's key that contains the password to use for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the password to use for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"auth_secret": {
										Description:         "The secret's key that contains the CRAM-MD5 secret. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the CRAM-MD5 secret. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"auth_username": {
										Description:         "The username to use for authentication.",
										MarkdownDescription: "The username to use for authentication.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"from": {
										Description:         "The sender address.",
										MarkdownDescription: "The sender address.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"headers": {
										Description:         "Further headers email header key/value pairs. Overrides any headers previously set by the notification implementation.",
										MarkdownDescription: "Further headers email header key/value pairs. Overrides any headers previously set by the notification implementation.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "Key of the tuple.",
												MarkdownDescription: "Key of the tuple.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},

											"value": {
												Description:         "Value of the tuple.",
												MarkdownDescription: "Value of the tuple.",

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

									"hello": {
										Description:         "The hostname to identify to the SMTP server.",
										MarkdownDescription: "The hostname to identify to the SMTP server.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"html": {
										Description:         "The HTML body of the email notification.",
										MarkdownDescription: "The HTML body of the email notification.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"require_tls": {
										Description:         "The SMTP TLS requirement. Note that Go does not support unencrypted connections to remote SMTP endpoints.",
										MarkdownDescription: "The SMTP TLS requirement. Note that Go does not support unencrypted connections to remote SMTP endpoints.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"smarthost": {
										Description:         "The SMTP host and port through which emails are sent. E.g. example.com:25",
										MarkdownDescription: "The SMTP host and port through which emails are sent. E.g. example.com:25",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"text": {
										Description:         "The text body of the email notification.",
										MarkdownDescription: "The text body of the email notification.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_config": {
										Description:         "TLS configuration",
										MarkdownDescription: "TLS configuration",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca": {
												Description:         "Struct containing the CA cert to use for the targets.",
												MarkdownDescription: "Struct containing the CA cert to use for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_map": {
														Description:         "ConfigMap containing data to use for the targets.",
														MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "Secret containing data to use for the targets.",
														MarkdownDescription: "Secret containing data to use for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cert": {
												Description:         "Struct containing the client cert file for the targets.",
												MarkdownDescription: "Struct containing the client cert file for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_map": {
														Description:         "ConfigMap containing data to use for the targets.",
														MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														Description:         "Secret containing data to use for the targets.",
														MarkdownDescription: "Secret containing data to use for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_skip_verify": {
												Description:         "Disable target certificate validation.",
												MarkdownDescription: "Disable target certificate validation.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key_secret": {
												Description:         "Secret containing the client key file for the targets.",
												MarkdownDescription: "Secret containing the client key file for the targets.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"server_name": {
												Description:         "Used to verify the hostname for the targets.",
												MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"to": {
										Description:         "The email address to send notifications to.",
										MarkdownDescription: "The email address to send notifications to.",

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

							"name": {
								Description:         "Name of the receiver. Must be unique across all items from the list.",
								MarkdownDescription: "Name of the receiver. Must be unique across all items from the list.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.LengthAtLeast(1),
								},
							},

							"opsgenie_configs": {
								Description:         "List of OpsGenie configurations.",
								MarkdownDescription: "List of OpsGenie configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"actions": {
										Description:         "Comma separated list of actions that will be available for the alert.",
										MarkdownDescription: "Comma separated list of actions that will be available for the alert.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"api_key": {
										Description:         "The secret's key that contains the OpsGenie API key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the OpsGenie API key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"api_url": {
										Description:         "The URL to send OpsGenie API requests to.",
										MarkdownDescription: "The URL to send OpsGenie API requests to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"description": {
										Description:         "Description of the incident.",
										MarkdownDescription: "Description of the incident.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"details": {
										Description:         "A set of arbitrary key/value pairs that provide further detail about the incident.",
										MarkdownDescription: "A set of arbitrary key/value pairs that provide further detail about the incident.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "Key of the tuple.",
												MarkdownDescription: "Key of the tuple.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},

											"value": {
												Description:         "Value of the tuple.",
												MarkdownDescription: "Value of the tuple.",

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

									"entity": {
										Description:         "Optional field that can be used to specify which domain alert is related to.",
										MarkdownDescription: "Optional field that can be used to specify which domain alert is related to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_config": {
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"message": {
										Description:         "Alert text limited to 130 characters.",
										MarkdownDescription: "Alert text limited to 130 characters.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"note": {
										Description:         "Additional alert note.",
										MarkdownDescription: "Additional alert note.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"priority": {
										Description:         "Priority level of alert. Possible values are P1, P2, P3, P4, and P5.",
										MarkdownDescription: "Priority level of alert. Possible values are P1, P2, P3, P4, and P5.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"responders": {
										Description:         "List of responders responsible for notifications.",
										MarkdownDescription: "List of responders responsible for notifications.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"id": {
												Description:         "ID of the responder.",
												MarkdownDescription: "ID of the responder.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name of the responder.",
												MarkdownDescription: "Name of the responder.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Type of responder.",
												MarkdownDescription: "Type of responder.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.OneOf("team", "teams", "user", "escalation", "schedule"),
												},
											},

											"username": {
												Description:         "Username of the responder.",
												MarkdownDescription: "Username of the responder.",

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

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source": {
										Description:         "Backlink to the sender of the notification.",
										MarkdownDescription: "Backlink to the sender of the notification.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tags": {
										Description:         "Comma separated list of tags attached to the notifications.",
										MarkdownDescription: "Comma separated list of tags attached to the notifications.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"update_alerts": {
										Description:         "Whether to update message and description of the alert in OpsGenie if it already exists By default, the alert is never updated in OpsGenie, the new message only appears in activity log.",
										MarkdownDescription: "Whether to update message and description of the alert in OpsGenie if it already exists By default, the alert is never updated in OpsGenie, the new message only appears in activity log.",

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

							"pagerduty_configs": {
								Description:         "List of PagerDuty configurations.",
								MarkdownDescription: "List of PagerDuty configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"class": {
										Description:         "The class/type of the event.",
										MarkdownDescription: "The class/type of the event.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client": {
										Description:         "Client identification.",
										MarkdownDescription: "Client identification.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_url": {
										Description:         "Backlink to the sender of notification.",
										MarkdownDescription: "Backlink to the sender of notification.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"component": {
										Description:         "The part or component of the affected system that is broken.",
										MarkdownDescription: "The part or component of the affected system that is broken.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"description": {
										Description:         "Description of the incident.",
										MarkdownDescription: "Description of the incident.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"details": {
										Description:         "Arbitrary key/value pairs that provide further detail about the incident.",
										MarkdownDescription: "Arbitrary key/value pairs that provide further detail about the incident.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "Key of the tuple.",
												MarkdownDescription: "Key of the tuple.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},

											"value": {
												Description:         "Value of the tuple.",
												MarkdownDescription: "Value of the tuple.",

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

									"group": {
										Description:         "A cluster or grouping of sources.",
										MarkdownDescription: "A cluster or grouping of sources.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_config": {
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"pager_duty_image_configs": {
										Description:         "A list of image details to attach that provide further detail about an incident.",
										MarkdownDescription: "A list of image details to attach that provide further detail about an incident.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"alt": {
												Description:         "Alt is the optional alternative text for the image.",
												MarkdownDescription: "Alt is the optional alternative text for the image.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"href": {
												Description:         "Optional URL; makes the image a clickable link.",
												MarkdownDescription: "Optional URL; makes the image a clickable link.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"src": {
												Description:         "Src of the image being attached to the incident",
												MarkdownDescription: "Src of the image being attached to the incident",

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

									"pager_duty_link_configs": {
										Description:         "A list of link details to attach that provide further detail about an incident.",
										MarkdownDescription: "A list of link details to attach that provide further detail about an incident.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"alt": {
												Description:         "Text that describes the purpose of the link, and can be used as the link's text.",
												MarkdownDescription: "Text that describes the purpose of the link, and can be used as the link's text.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"href": {
												Description:         "Href is the URL of the link to be attached",
												MarkdownDescription: "Href is the URL of the link to be attached",

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

									"routing_key": {
										Description:         "The secret's key that contains the PagerDuty integration key (when using Events API v2). Either this field or 'serviceKey' needs to be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the PagerDuty integration key (when using Events API v2). Either this field or 'serviceKey' needs to be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_key": {
										Description:         "The secret's key that contains the PagerDuty service key (when using integration type 'Prometheus'). Either this field or 'routingKey' needs to be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the PagerDuty service key (when using integration type 'Prometheus'). Either this field or 'routingKey' needs to be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"severity": {
										Description:         "Severity of the incident.",
										MarkdownDescription: "Severity of the incident.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "The URL to send requests to.",
										MarkdownDescription: "The URL to send requests to.",

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

							"pushover_configs": {
								Description:         "List of Pushover configurations.",
								MarkdownDescription: "List of Pushover configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"expire": {
										Description:         "How long your notification will continue to be retried for, unless the user acknowledges the notification.",
										MarkdownDescription: "How long your notification will continue to be retried for, unless the user acknowledges the notification.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
										},
									},

									"html": {
										Description:         "Whether notification message is HTML or plain text.",
										MarkdownDescription: "Whether notification message is HTML or plain text.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_config": {
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"message": {
										Description:         "Notification message.",
										MarkdownDescription: "Notification message.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"priority": {
										Description:         "Priority, see https://pushover.net/api#priority",
										MarkdownDescription: "Priority, see https://pushover.net/api#priority",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retry": {
										Description:         "How often the Pushover servers will send the same notification to the user. Must be at least 30 seconds.",
										MarkdownDescription: "How often the Pushover servers will send the same notification to the user. Must be at least 30 seconds.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
										},
									},

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sound": {
										Description:         "The name of one of the sounds supported by device clients to override the user's default sound choice",
										MarkdownDescription: "The name of one of the sounds supported by device clients to override the user's default sound choice",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"title": {
										Description:         "Notification title.",
										MarkdownDescription: "Notification title.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"token": {
										Description:         "The secret's key that contains the registered application's API token, see https://pushover.net/apps. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the registered application's API token, see https://pushover.net/apps. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"url": {
										Description:         "A supplementary URL shown alongside the message.",
										MarkdownDescription: "A supplementary URL shown alongside the message.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url_title": {
										Description:         "A title for supplementary URL, otherwise just the URL is shown",
										MarkdownDescription: "A title for supplementary URL, otherwise just the URL is shown",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_key": {
										Description:         "The secret's key that contains the recipient user's user key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the recipient user's user key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"slack_configs": {
								Description:         "List of Slack configurations.",
								MarkdownDescription: "List of Slack configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"actions": {
										Description:         "A list of Slack actions that are sent with each notification.",
										MarkdownDescription: "A list of Slack actions that are sent with each notification.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"confirm": {
												Description:         "SlackConfirmationField protect users from destructive actions or particularly distinguished decisions by asking them to confirm their button click one more time. See https://api.slack.com/docs/interactive-message-field-guide#confirmation_fields for more information.",
												MarkdownDescription: "SlackConfirmationField protect users from destructive actions or particularly distinguished decisions by asking them to confirm their button click one more time. See https://api.slack.com/docs/interactive-message-field-guide#confirmation_fields for more information.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dismiss_text": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ok_text": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"text": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},

													"title": {
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

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"style": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"text": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},

											"url": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

									"api_url": {
										Description:         "The secret's key that contains the Slack webhook URL. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the Slack webhook URL. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"callback_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"channel": {
										Description:         "The channel or user to send notifications to.",
										MarkdownDescription: "The channel or user to send notifications to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"color": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fallback": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fields": {
										Description:         "A list of Slack fields that are sent with each notification.",
										MarkdownDescription: "A list of Slack fields that are sent with each notification.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"short": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"title": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},

											"value": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"footer": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_config": {
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"icon_emoji": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"icon_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"link_names": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mrkdwn_in": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pretext": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"short_fields": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"text": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"thumb_url": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"title": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"title_link": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"username": {
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

							"sns_configs": {
								Description:         "List of SNS configurations",
								MarkdownDescription: "List of SNS configurations",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"api_url": {
										Description:         "The SNS API URL i.e. https://sns.us-east-2.amazonaws.com. If not specified, the SNS API URL from the SNS SDK will be used.",
										MarkdownDescription: "The SNS API URL i.e. https://sns.us-east-2.amazonaws.com. If not specified, the SNS API URL from the SNS SDK will be used.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"attributes": {
										Description:         "SNS message attributes.",
										MarkdownDescription: "SNS message attributes.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_config": {
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"message": {
										Description:         "The message content of the SNS notification.",
										MarkdownDescription: "The message content of the SNS notification.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"phone_number": {
										Description:         "Phone number if message is delivered via SMS in E.164 format. If you don't specify this value, you must specify a value for the TopicARN or TargetARN.",
										MarkdownDescription: "Phone number if message is delivered via SMS in E.164 format. If you don't specify this value, you must specify a value for the TopicARN or TargetARN.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sigv4": {
										Description:         "Configures AWS's Signature Verification 4 signing process to sign requests.",
										MarkdownDescription: "Configures AWS's Signature Verification 4 signing process to sign requests.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_key": {
												Description:         "AccessKey is the AWS API key. If blank, the environment variable 'AWS_ACCESS_KEY_ID' is used.",
												MarkdownDescription: "AccessKey is the AWS API key. If blank, the environment variable 'AWS_ACCESS_KEY_ID' is used.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"profile": {
												Description:         "Profile is the named AWS profile used to authenticate.",
												MarkdownDescription: "Profile is the named AWS profile used to authenticate.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"region": {
												Description:         "Region is the AWS region. If blank, the region from the default credentials chain used.",
												MarkdownDescription: "Region is the AWS region. If blank, the region from the default credentials chain used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"role_arn": {
												Description:         "RoleArn is the named AWS profile used to authenticate.",
												MarkdownDescription: "RoleArn is the named AWS profile used to authenticate.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_key": {
												Description:         "SecretKey is the AWS API secret. If blank, the environment variable 'AWS_SECRET_ACCESS_KEY' is used.",
												MarkdownDescription: "SecretKey is the AWS API secret. If blank, the environment variable 'AWS_SECRET_ACCESS_KEY' is used.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subject": {
										Description:         "Subject line when the message is delivered to email endpoints.",
										MarkdownDescription: "Subject line when the message is delivered to email endpoints.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_arn": {
										Description:         "The  mobile platform endpoint ARN if message is delivered via mobile notifications. If you don't specify this value, you must specify a value for the topic_arn or PhoneNumber.",
										MarkdownDescription: "The  mobile platform endpoint ARN if message is delivered via mobile notifications. If you don't specify this value, you must specify a value for the topic_arn or PhoneNumber.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"topic_arn": {
										Description:         "SNS topic ARN, i.e. arn:aws:sns:us-east-2:698519295917:My-Topic If you don't specify this value, you must specify a value for the PhoneNumber or TargetARN.",
										MarkdownDescription: "SNS topic ARN, i.e. arn:aws:sns:us-east-2:698519295917:My-Topic If you don't specify this value, you must specify a value for the PhoneNumber or TargetARN.",

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

							"telegram_configs": {
								Description:         "List of Telegram configurations.",
								MarkdownDescription: "List of Telegram configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"api_url": {
										Description:         "The Telegram API URL i.e. https://api.telegram.org. If not specified, default API URL will be used.",
										MarkdownDescription: "The Telegram API URL i.e. https://api.telegram.org. If not specified, default API URL will be used.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"bot_token": {
										Description:         "Telegram bot token The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "Telegram bot token The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"chat_id": {
										Description:         "The Telegram chat ID.",
										MarkdownDescription: "The Telegram chat ID.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_notifications": {
										Description:         "Disable telegram notifications",
										MarkdownDescription: "Disable telegram notifications",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_config": {
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"message": {
										Description:         "Message template",
										MarkdownDescription: "Message template",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"parse_mode": {
										Description:         "Parse mode for telegram message",
										MarkdownDescription: "Parse mode for telegram message",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("MarkdownV2", "Markdown", "HTML"),
										},
									},

									"send_resolved": {
										Description:         "Whether to notify about resolved alerts.",
										MarkdownDescription: "Whether to notify about resolved alerts.",

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

							"victorops_configs": {
								Description:         "List of VictorOps configurations.",
								MarkdownDescription: "List of VictorOps configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"api_key": {
										Description:         "The secret's key that contains the API key to use when talking to the VictorOps API. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the API key to use when talking to the VictorOps API. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"api_url": {
										Description:         "The VictorOps API URL.",
										MarkdownDescription: "The VictorOps API URL.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"custom_fields": {
										Description:         "Additional custom fields for notification.",
										MarkdownDescription: "Additional custom fields for notification.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "Key of the tuple.",
												MarkdownDescription: "Key of the tuple.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},

											"value": {
												Description:         "Value of the tuple.",
												MarkdownDescription: "Value of the tuple.",

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

									"entity_display_name": {
										Description:         "Contains summary of the alerted problem.",
										MarkdownDescription: "Contains summary of the alerted problem.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_config": {
										Description:         "The HTTP client's configuration.",
										MarkdownDescription: "The HTTP client's configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"message_type": {
										Description:         "Describes the behavior of the alert (CRITICAL, WARNING, INFO).",
										MarkdownDescription: "Describes the behavior of the alert (CRITICAL, WARNING, INFO).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"monitoring_tool": {
										Description:         "The monitoring tool the state message is from.",
										MarkdownDescription: "The monitoring tool the state message is from.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"routing_key": {
										Description:         "A key used to map the alert to a team.",
										MarkdownDescription: "A key used to map the alert to a team.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"state_message": {
										Description:         "Contains long explanation of the alerted problem.",
										MarkdownDescription: "Contains long explanation of the alerted problem.",

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

							"webhook_configs": {
								Description:         "List of webhook configurations.",
								MarkdownDescription: "List of webhook configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"http_config": {
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"max_alerts": {
										Description:         "Maximum number of alerts to be sent per webhook message. When 0, all alerts are included.",
										MarkdownDescription: "Maximum number of alerts to be sent per webhook message. When 0, all alerts are included.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "The URL to send HTTP POST requests to. 'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url' should be defined.",
										MarkdownDescription: "The URL to send HTTP POST requests to. 'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url' should be defined.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url_secret": {
										Description:         "The secret's key that contains the webhook URL to send HTTP requests to. 'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url' should be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the webhook URL to send HTTP requests to. 'urlSecret' takes precedence over 'url'. One of 'urlSecret' and 'url' should be defined. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"wechat_configs": {
								Description:         "List of WeChat configurations.",
								MarkdownDescription: "List of WeChat configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"agent_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"api_secret": {
										Description:         "The secret's key that contains the WeChat API key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
										MarkdownDescription: "The secret's key that contains the WeChat API key. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

									"api_url": {
										Description:         "The WeChat API URL.",
										MarkdownDescription: "The WeChat API URL.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"corp_id": {
										Description:         "The corp id for authentication.",
										MarkdownDescription: "The corp id for authentication.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_config": {
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"authorization": {
												Description:         "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",
												MarkdownDescription: "Authorization header configuration for the client. This is mutually exclusive with BasicAuth and is only available starting from Alertmanager v0.22+.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"credentials": {
														Description:         "The secret's key that contains the credentials of the request",
														MarkdownDescription: "The secret's key that contains the credentials of the request",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"type": {
														Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
														MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

											"basic_auth": {
												Description:         "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",
												MarkdownDescription: "BasicAuth for the client. This is mutually exclusive with Authorization. If both are defined, BasicAuth takes precedence.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"password": {
														Description:         "The secret in the service monitor namespace that contains the password for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"username": {
														Description:         "The secret in the service monitor namespace that contains the username for authentication.",
														MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"bearer_token_secret": {
												Description:         "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",
												MarkdownDescription: "The secret's key that contains the bearer token to be used by the client for authentication. The secret needs to be in the same namespace as the AlertmanagerConfig object and accessible by the Prometheus Operator.",

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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

											"follow_redirects": {
												Description:         "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",
												MarkdownDescription: "FollowRedirects specifies whether the client should follow HTTP 3xx redirects.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"oauth2": {
												Description:         "OAuth2 client credentials used to fetch a token for the targets.",
												MarkdownDescription: "OAuth2 client credentials used to fetch a token for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"client_id": {
														Description:         "The secret or configmap containing the OAuth2 client id",
														MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: true,
														Optional: false,
														Computed: false,
													},

													"client_secret": {
														Description:         "The secret containing the OAuth2 client secret",
														MarkdownDescription: "The secret containing the OAuth2 client secret",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

														Required: true,
														Optional: false,
														Computed: false,
													},

													"endpoint_params": {
														Description:         "Parameters to append to the token URL",
														MarkdownDescription: "Parameters to append to the token URL",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"scopes": {
														Description:         "OAuth2 scopes used for the token request",
														MarkdownDescription: "OAuth2 scopes used for the token request",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_url": {
														Description:         "The URL to fetch the token from",
														MarkdownDescription: "The URL to fetch the token from",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proxy_url": {
												Description:         "Optional proxy URL.",
												MarkdownDescription: "Optional proxy URL.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_config": {
												Description:         "TLS configuration for the client.",
												MarkdownDescription: "TLS configuration for the client.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca": {
														Description:         "Struct containing the CA cert to use for the targets.",
														MarkdownDescription: "Struct containing the CA cert to use for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert": {
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap containing data to use for the targets.",
																MarkdownDescription: "ConfigMap containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
																Description:         "Secret containing data to use for the targets.",
																MarkdownDescription: "Secret containing data to use for the targets.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key_secret": {
														Description:         "Secret containing the client key file for the targets.",
														MarkdownDescription: "Secret containing the client key file for the targets.",

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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",

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

													"server_name": {
														Description:         "Used to verify the hostname for the targets.",
														MarkdownDescription: "Used to verify the hostname for the targets.",

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

									"message": {
										Description:         "API request data as defined by the WeChat API.",
										MarkdownDescription: "API request data as defined by the WeChat API.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"message_type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"send_resolved": {
										Description:         "Whether or not to notify about resolved alerts.",
										MarkdownDescription: "Whether or not to notify about resolved alerts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_party": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_tag": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"to_user": {
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

					"route": {
						Description:         "The Alertmanager route definition for alerts matching the resource's namespace. If present, it will be added to the generated Alertmanager configuration as a first-level route.",
						MarkdownDescription: "The Alertmanager route definition for alerts matching the resource's namespace. If present, it will be added to the generated Alertmanager configuration as a first-level route.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"continue": {
								Description:         "Boolean indicating whether an alert should continue matching subsequent sibling nodes. It will always be overridden to true for the first-level route by the Prometheus operator.",
								MarkdownDescription: "Boolean indicating whether an alert should continue matching subsequent sibling nodes. It will always be overridden to true for the first-level route by the Prometheus operator.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"group_by": {
								Description:         "List of labels to group by. Labels must not be repeated (unique list). Special label '...' (aggregate by all possible labels), if provided, must be the only element in the list.",
								MarkdownDescription: "List of labels to group by. Labels must not be repeated (unique list). Special label '...' (aggregate by all possible labels), if provided, must be the only element in the list.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"group_interval": {
								Description:         "How long to wait before sending an updated notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '5m'",
								MarkdownDescription: "How long to wait before sending an updated notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '5m'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"group_wait": {
								Description:         "How long to wait before sending the initial notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '30s'",
								MarkdownDescription: "How long to wait before sending the initial notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '30s'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"matchers": {
								Description:         "List of matchers that the alert's labels should match. For the first level route, the operator removes any existing equality and regexp matcher on the 'namespace' label and adds a 'namespace: <object namespace>' matcher.",
								MarkdownDescription: "List of matchers that the alert's labels should match. For the first level route, the operator removes any existing equality and regexp matcher on the 'namespace' label and adds a 'namespace: <object namespace>' matcher.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"match_type": {
										Description:         "Match operation available with AlertManager >= v0.22.0 and takes precedence over Regex (deprecated) if non-empty.",
										MarkdownDescription: "Match operation available with AlertManager >= v0.22.0 and takes precedence over Regex (deprecated) if non-empty.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("!=", "=", "=~", "!~"),
										},
									},

									"name": {
										Description:         "Label to match.",
										MarkdownDescription: "Label to match.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),
										},
									},

									"regex": {
										Description:         "Whether to match on equality (false) or regular-expression (true). Deprecated as of AlertManager >= v0.22.0 where a user should use MatchType instead.",
										MarkdownDescription: "Whether to match on equality (false) or regular-expression (true). Deprecated as of AlertManager >= v0.22.0 where a user should use MatchType instead.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Label value to match.",
										MarkdownDescription: "Label value to match.",

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

							"mute_time_intervals": {
								Description:         "Note: this comment applies to the field definition above but appears below otherwise it gets included in the generated manifest. CRD schema doesn't support self-referential types for now (see https://github.com/kubernetes/kubernetes/issues/62872). We have to use an alternative type to circumvent the limitation. The downside is that the Kube API can't validate the data beyond the fact that it is a valid JSON representation. MuteTimeIntervals is a list of MuteTimeInterval names that will mute this route when matched,",
								MarkdownDescription: "Note: this comment applies to the field definition above but appears below otherwise it gets included in the generated manifest. CRD schema doesn't support self-referential types for now (see https://github.com/kubernetes/kubernetes/issues/62872). We have to use an alternative type to circumvent the limitation. The downside is that the Kube API can't validate the data beyond the fact that it is a valid JSON representation. MuteTimeIntervals is a list of MuteTimeInterval names that will mute this route when matched,",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"receiver": {
								Description:         "Name of the receiver for this route. If not empty, it should be listed in the 'receivers' field.",
								MarkdownDescription: "Name of the receiver for this route. If not empty, it should be listed in the 'receivers' field.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"repeat_interval": {
								Description:         "How long to wait before repeating the last notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '4h'",
								MarkdownDescription: "How long to wait before repeating the last notification. Must match the regular expression'^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$' Example: '4h'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"routes": {
								Description:         "Child routes.",
								MarkdownDescription: "Child routes.",

								Type: types.ListType{ElemType: types.StringType},

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

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_monitoring_coreos_com_alertmanager_config_v1alpha1")

	var state MonitoringCoreosComAlertmanagerConfigV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MonitoringCoreosComAlertmanagerConfigV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("monitoring.coreos.com/v1alpha1")
	goModel.Kind = utilities.Ptr("AlertmanagerConfig")

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

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_coreos_com_alertmanager_config_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_monitoring_coreos_com_alertmanager_config_v1alpha1")

	var state MonitoringCoreosComAlertmanagerConfigV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MonitoringCoreosComAlertmanagerConfigV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("monitoring.coreos.com/v1alpha1")
	goModel.Kind = utilities.Ptr("AlertmanagerConfig")

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

func (r *MonitoringCoreosComAlertmanagerConfigV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_monitoring_coreos_com_alertmanager_config_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
