/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package operator_victoriametrics_com_v1beta1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &OperatorVictoriametricsComVmalertmanagerConfigV1Beta1Manifest{}
)

func NewOperatorVictoriametricsComVmalertmanagerConfigV1Beta1Manifest() datasource.DataSource {
	return &OperatorVictoriametricsComVmalertmanagerConfigV1Beta1Manifest{}
}

type OperatorVictoriametricsComVmalertmanagerConfigV1Beta1Manifest struct{}

type OperatorVictoriametricsComVmalertmanagerConfigV1Beta1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		Inhibit_rules *[]struct {
			Equal           *[]string `tfsdk:"equal" json:"equal,omitempty"`
			Source_matchers *[]string `tfsdk:"source_matchers" json:"source_matchers,omitempty"`
			Target_matchers *[]string `tfsdk:"target_matchers" json:"target_matchers,omitempty"`
		} `tfsdk:"inhibit_rules" json:"inhibit_rules,omitempty"`
		Mute_time_intervals *[]struct {
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Time_intervals *[]struct {
				Days_of_month *[]string `tfsdk:"days_of_month" json:"days_of_month,omitempty"`
				Location      *string   `tfsdk:"location" json:"location,omitempty"`
				Months        *[]string `tfsdk:"months" json:"months,omitempty"`
				Times         *[]struct {
					End_time   *string `tfsdk:"end_time" json:"end_time,omitempty"`
					Start_time *string `tfsdk:"start_time" json:"start_time,omitempty"`
				} `tfsdk:"times" json:"times,omitempty"`
				Weekdays *[]string `tfsdk:"weekdays" json:"weekdays,omitempty"`
				Years    *[]string `tfsdk:"years" json:"years,omitempty"`
			} `tfsdk:"time_intervals" json:"time_intervals,omitempty"`
		} `tfsdk:"mute_time_intervals" json:"mute_time_intervals,omitempty"`
		Receivers *[]struct {
			Discord_configs *[]struct {
				Http_config *struct {
					Basic_auth *struct {
						Password *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
						Username      *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basic_auth,omitempty"`
					Bearer_token_file   *string `tfsdk:"bearer_token_file" json:"bearer_token_file,omitempty"`
					Bearer_token_secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearer_token_secret,omitempty"`
					ProxyURL   *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					Tls_config *struct {
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
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
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
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
				} `tfsdk:"http_config" json:"http_config,omitempty"`
				Message            *string `tfsdk:"message" json:"message,omitempty"`
				Send_resolved      *bool   `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Title              *string `tfsdk:"title" json:"title,omitempty"`
				Webhook_url        *string `tfsdk:"webhook_url" json:"webhook_url,omitempty"`
				Webhook_url_secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"webhook_url_secret" json:"webhook_url_secret,omitempty"`
			} `tfsdk:"discord_configs" json:"discord_configs,omitempty"`
			Email_configs *[]struct {
				Auth_identity *string `tfsdk:"auth_identity" json:"auth_identity,omitempty"`
				Auth_password *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"auth_password" json:"auth_password,omitempty"`
				Auth_secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"auth_secret" json:"auth_secret,omitempty"`
				Auth_username *string            `tfsdk:"auth_username" json:"auth_username,omitempty"`
				From          *string            `tfsdk:"from" json:"from,omitempty"`
				Headers       *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
				Hello         *string            `tfsdk:"hello" json:"hello,omitempty"`
				Html          *string            `tfsdk:"html" json:"html,omitempty"`
				Require_tls   *bool              `tfsdk:"require_tls" json:"require_tls,omitempty"`
				Send_resolved *bool              `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Smarthost     *string            `tfsdk:"smarthost" json:"smarthost,omitempty"`
				Text          *string            `tfsdk:"text" json:"text,omitempty"`
				Tls_config    *struct {
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
					CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
					Cert   *struct {
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
					CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
					KeySecret          *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"key_secret" json:"keySecret,omitempty"`
					ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
				} `tfsdk:"tls_config" json:"tls_config,omitempty"`
				To *string `tfsdk:"to" json:"to,omitempty"`
			} `tfsdk:"email_configs" json:"email_configs,omitempty"`
			Msteams_configs *[]struct {
				Http_config *struct {
					Basic_auth *struct {
						Password *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
						Username      *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basic_auth,omitempty"`
					Bearer_token_file   *string `tfsdk:"bearer_token_file" json:"bearer_token_file,omitempty"`
					Bearer_token_secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearer_token_secret,omitempty"`
					ProxyURL   *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					Tls_config *struct {
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
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
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
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
				} `tfsdk:"http_config" json:"http_config,omitempty"`
				Send_resolved      *bool   `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Text               *string `tfsdk:"text" json:"text,omitempty"`
				Title              *string `tfsdk:"title" json:"title,omitempty"`
				Webhook_url        *string `tfsdk:"webhook_url" json:"webhook_url,omitempty"`
				Webhook_url_secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"webhook_url_secret" json:"webhook_url_secret,omitempty"`
			} `tfsdk:"msteams_configs" json:"msteams_configs,omitempty"`
			Name             *string `tfsdk:"name" json:"name,omitempty"`
			Opsgenie_configs *[]struct {
				Actions *string `tfsdk:"actions" json:"actions,omitempty"`
				ApiURL  *string `tfsdk:"api_url" json:"apiURL,omitempty"`
				Api_key *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"api_key" json:"api_key,omitempty"`
				Description *string            `tfsdk:"description" json:"description,omitempty"`
				Details     *map[string]string `tfsdk:"details" json:"details,omitempty"`
				Entity      *string            `tfsdk:"entity" json:"entity,omitempty"`
				Http_config *map[string]string `tfsdk:"http_config" json:"http_config,omitempty"`
				Message     *string            `tfsdk:"message" json:"message,omitempty"`
				Note        *string            `tfsdk:"note" json:"note,omitempty"`
				Priority    *string            `tfsdk:"priority" json:"priority,omitempty"`
				Responders  *[]struct {
					Id       *string `tfsdk:"id" json:"id,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Type     *string `tfsdk:"type" json:"type,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"responders" json:"responders,omitempty"`
				Send_resolved *bool   `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Source        *string `tfsdk:"source" json:"source,omitempty"`
				Tags          *string `tfsdk:"tags" json:"tags,omitempty"`
				Update_alerts *bool   `tfsdk:"update_alerts" json:"update_alerts,omitempty"`
			} `tfsdk:"opsgenie_configs" json:"opsgenie_configs,omitempty"`
			Pagerduty_configs *[]struct {
				Class       *string            `tfsdk:"class" json:"class,omitempty"`
				Client      *string            `tfsdk:"client" json:"client,omitempty"`
				Client_url  *string            `tfsdk:"client_url" json:"client_url,omitempty"`
				Component   *string            `tfsdk:"component" json:"component,omitempty"`
				Description *string            `tfsdk:"description" json:"description,omitempty"`
				Details     *map[string]string `tfsdk:"details" json:"details,omitempty"`
				Group       *string            `tfsdk:"group" json:"group,omitempty"`
				Http_config *map[string]string `tfsdk:"http_config" json:"http_config,omitempty"`
				Images      *[]struct {
					Alt    *string `tfsdk:"alt" json:"alt,omitempty"`
					Href   *string `tfsdk:"href" json:"href,omitempty"`
					Source *string `tfsdk:"source" json:"source,omitempty"`
				} `tfsdk:"images" json:"images,omitempty"`
				Links *[]struct {
					Href *string `tfsdk:"href" json:"href,omitempty"`
					Text *string `tfsdk:"text" json:"text,omitempty"`
				} `tfsdk:"links" json:"links,omitempty"`
				Routing_key *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"routing_key" json:"routing_key,omitempty"`
				Send_resolved *bool `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Service_key   *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"service_key" json:"service_key,omitempty"`
				Severity *string `tfsdk:"severity" json:"severity,omitempty"`
				Url      *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"pagerduty_configs" json:"pagerduty_configs,omitempty"`
			Pushover_configs *[]struct {
				Expire        *string            `tfsdk:"expire" json:"expire,omitempty"`
				Html          *bool              `tfsdk:"html" json:"html,omitempty"`
				Http_config   *map[string]string `tfsdk:"http_config" json:"http_config,omitempty"`
				Message       *string            `tfsdk:"message" json:"message,omitempty"`
				Priority      *string            `tfsdk:"priority" json:"priority,omitempty"`
				Retry         *string            `tfsdk:"retry" json:"retry,omitempty"`
				Send_resolved *bool              `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Sound         *string            `tfsdk:"sound" json:"sound,omitempty"`
				Title         *string            `tfsdk:"title" json:"title,omitempty"`
				Token         *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"token" json:"token,omitempty"`
				Url       *string `tfsdk:"url" json:"url,omitempty"`
				Url_title *string `tfsdk:"url_title" json:"url_title,omitempty"`
				User_key  *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"user_key" json:"user_key,omitempty"`
			} `tfsdk:"pushover_configs" json:"pushover_configs,omitempty"`
			Slack_configs *[]struct {
				Actions *[]struct {
					Confirm *struct {
						Dismiss_text *string `tfsdk:"dismiss_text" json:"dismiss_text,omitempty"`
						Ok_text      *string `tfsdk:"ok_text" json:"ok_text,omitempty"`
						Text         *string `tfsdk:"text" json:"text,omitempty"`
						Title        *string `tfsdk:"title" json:"title,omitempty"`
					} `tfsdk:"confirm" json:"confirm,omitempty"`
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Style *string `tfsdk:"style" json:"style,omitempty"`
					Text  *string `tfsdk:"text" json:"text,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Url   *string `tfsdk:"url" json:"url,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"actions" json:"actions,omitempty"`
				Api_url *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"api_url" json:"api_url,omitempty"`
				Callback_id *string `tfsdk:"callback_id" json:"callback_id,omitempty"`
				Channel     *string `tfsdk:"channel" json:"channel,omitempty"`
				Color       *string `tfsdk:"color" json:"color,omitempty"`
				Fallback    *string `tfsdk:"fallback" json:"fallback,omitempty"`
				Fields      *[]struct {
					Short *bool   `tfsdk:"short" json:"short,omitempty"`
					Title *string `tfsdk:"title" json:"title,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"fields" json:"fields,omitempty"`
				Footer        *string            `tfsdk:"footer" json:"footer,omitempty"`
				Http_config   *map[string]string `tfsdk:"http_config" json:"http_config,omitempty"`
				Icon_emoji    *string            `tfsdk:"icon_emoji" json:"icon_emoji,omitempty"`
				Icon_url      *string            `tfsdk:"icon_url" json:"icon_url,omitempty"`
				Image_url     *string            `tfsdk:"image_url" json:"image_url,omitempty"`
				Link_names    *bool              `tfsdk:"link_names" json:"link_names,omitempty"`
				Mrkdwn_in     *[]string          `tfsdk:"mrkdwn_in" json:"mrkdwn_in,omitempty"`
				Pretext       *string            `tfsdk:"pretext" json:"pretext,omitempty"`
				Send_resolved *bool              `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Short_fields  *bool              `tfsdk:"short_fields" json:"short_fields,omitempty"`
				Text          *string            `tfsdk:"text" json:"text,omitempty"`
				Thumb_url     *string            `tfsdk:"thumb_url" json:"thumb_url,omitempty"`
				Title         *string            `tfsdk:"title" json:"title,omitempty"`
				Title_link    *string            `tfsdk:"title_link" json:"title_link,omitempty"`
				Username      *string            `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"slack_configs" json:"slack_configs,omitempty"`
			Sns_configs *[]struct {
				Api_url     *string            `tfsdk:"api_url" json:"api_url,omitempty"`
				Attributes  *map[string]string `tfsdk:"attributes" json:"attributes,omitempty"`
				Http_config *struct {
					Basic_auth *struct {
						Password *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
						Username      *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basic_auth,omitempty"`
					Bearer_token_file   *string `tfsdk:"bearer_token_file" json:"bearer_token_file,omitempty"`
					Bearer_token_secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearer_token_secret,omitempty"`
					ProxyURL   *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					Tls_config *struct {
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
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
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
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
				} `tfsdk:"http_config" json:"http_config,omitempty"`
				Message       *string `tfsdk:"message" json:"message,omitempty"`
				Phone_number  *string `tfsdk:"phone_number" json:"phone_number,omitempty"`
				Send_resolved *bool   `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Sigv4         *struct {
					Access_key          *string `tfsdk:"access_key" json:"access_key,omitempty"`
					Access_key_selector *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"access_key_selector" json:"access_key_selector,omitempty"`
					Profile             *string `tfsdk:"profile" json:"profile,omitempty"`
					Region              *string `tfsdk:"region" json:"region,omitempty"`
					Role_arn            *string `tfsdk:"role_arn" json:"role_arn,omitempty"`
					Secret_key_selector *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"secret_key_selector" json:"secret_key_selector,omitempty"`
				} `tfsdk:"sigv4" json:"sigv4,omitempty"`
				Subject    *string `tfsdk:"subject" json:"subject,omitempty"`
				Target_arn *string `tfsdk:"target_arn" json:"target_arn,omitempty"`
				Topic_arn  *string `tfsdk:"topic_arn" json:"topic_arn,omitempty"`
			} `tfsdk:"sns_configs" json:"sns_configs,omitempty"`
			Telegram_configs *[]struct {
				Api_url   *string `tfsdk:"api_url" json:"api_url,omitempty"`
				Bot_token *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"bot_token" json:"bot_token,omitempty"`
				Chat_id               *int64             `tfsdk:"chat_id" json:"chat_id,omitempty"`
				Disable_notifications *bool              `tfsdk:"disable_notifications" json:"disable_notifications,omitempty"`
				Http_config           *map[string]string `tfsdk:"http_config" json:"http_config,omitempty"`
				Message               *string            `tfsdk:"message" json:"message,omitempty"`
				Parse_mode            *string            `tfsdk:"parse_mode" json:"parse_mode,omitempty"`
				Send_resolved         *bool              `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
			} `tfsdk:"telegram_configs" json:"telegram_configs,omitempty"`
			Victorops_configs *[]struct {
				Api_key *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"api_key" json:"api_key,omitempty"`
				Api_url             *string            `tfsdk:"api_url" json:"api_url,omitempty"`
				Custom_fields       *map[string]string `tfsdk:"custom_fields" json:"custom_fields,omitempty"`
				Entity_display_name *string            `tfsdk:"entity_display_name" json:"entity_display_name,omitempty"`
				Http_config         *struct {
					Basic_auth *struct {
						Password *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
						Username      *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basic_auth,omitempty"`
					Bearer_token_file   *string `tfsdk:"bearer_token_file" json:"bearer_token_file,omitempty"`
					Bearer_token_secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearer_token_secret,omitempty"`
					ProxyURL   *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					Tls_config *struct {
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
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
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
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
				} `tfsdk:"http_config" json:"http_config,omitempty"`
				Message_type    *string `tfsdk:"message_type" json:"message_type,omitempty"`
				Monitoring_tool *string `tfsdk:"monitoring_tool" json:"monitoring_tool,omitempty"`
				Routing_key     *string `tfsdk:"routing_key" json:"routing_key,omitempty"`
				Send_resolved   *bool   `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				State_message   *string `tfsdk:"state_message" json:"state_message,omitempty"`
			} `tfsdk:"victorops_configs" json:"victorops_configs,omitempty"`
			Webex_configs *[]struct {
				Api_url     *string `tfsdk:"api_url" json:"api_url,omitempty"`
				Http_config *struct {
					Basic_auth *struct {
						Password *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
						Username      *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basic_auth,omitempty"`
					Bearer_token_file   *string `tfsdk:"bearer_token_file" json:"bearer_token_file,omitempty"`
					Bearer_token_secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearer_token_secret,omitempty"`
					ProxyURL   *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					Tls_config *struct {
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
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
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
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
				} `tfsdk:"http_config" json:"http_config,omitempty"`
				Message       *string `tfsdk:"message" json:"message,omitempty"`
				Room_id       *string `tfsdk:"room_id" json:"room_id,omitempty"`
				Send_resolved *bool   `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
			} `tfsdk:"webex_configs" json:"webex_configs,omitempty"`
			Webhook_configs *[]struct {
				Http_config   *map[string]string `tfsdk:"http_config" json:"http_config,omitempty"`
				Max_alerts    *int64             `tfsdk:"max_alerts" json:"max_alerts,omitempty"`
				Send_resolved *bool              `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				Url           *string            `tfsdk:"url" json:"url,omitempty"`
				Url_secret    *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"url_secret" json:"url_secret,omitempty"`
			} `tfsdk:"webhook_configs" json:"webhook_configs,omitempty"`
			Wechat_configs *[]struct {
				Agent_id   *string `tfsdk:"agent_id" json:"agent_id,omitempty"`
				Api_secret *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"api_secret" json:"api_secret,omitempty"`
				Api_url     *string `tfsdk:"api_url" json:"api_url,omitempty"`
				Corp_id     *string `tfsdk:"corp_id" json:"corp_id,omitempty"`
				Http_config *struct {
					Basic_auth *struct {
						Password *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"password" json:"password,omitempty"`
						Password_file *string `tfsdk:"password_file" json:"password_file,omitempty"`
						Username      *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basic_auth,omitempty"`
					Bearer_token_file   *string `tfsdk:"bearer_token_file" json:"bearer_token_file,omitempty"`
					Bearer_token_secret *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bearer_token_secret" json:"bearer_token_secret,omitempty"`
					ProxyURL   *string `tfsdk:"proxy_url" json:"proxyURL,omitempty"`
					Tls_config *struct {
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
						CaFile *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert   *struct {
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
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						KeySecret          *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"key_secret" json:"keySecret,omitempty"`
						ServerName *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
				} `tfsdk:"http_config" json:"http_config,omitempty"`
				Message       *string `tfsdk:"message" json:"message,omitempty"`
				Message_type  *string `tfsdk:"message_type" json:"message_type,omitempty"`
				Send_resolved *bool   `tfsdk:"send_resolved" json:"send_resolved,omitempty"`
				To_party      *string `tfsdk:"to_party" json:"to_party,omitempty"`
				To_tag        *string `tfsdk:"to_tag" json:"to_tag,omitempty"`
				To_user       *string `tfsdk:"to_user" json:"to_user,omitempty"`
			} `tfsdk:"wechat_configs" json:"wechat_configs,omitempty"`
		} `tfsdk:"receivers" json:"receivers,omitempty"`
		Route *struct {
			Active_time_intervals *[]string `tfsdk:"active_time_intervals" json:"active_time_intervals,omitempty"`
			Continue              *bool     `tfsdk:"continue" json:"continue,omitempty"`
			Group_by              *[]string `tfsdk:"group_by" json:"group_by,omitempty"`
			Group_interval        *string   `tfsdk:"group_interval" json:"group_interval,omitempty"`
			Group_wait            *string   `tfsdk:"group_wait" json:"group_wait,omitempty"`
			Matchers              *[]string `tfsdk:"matchers" json:"matchers,omitempty"`
			Mute_time_intervals   *[]string `tfsdk:"mute_time_intervals" json:"mute_time_intervals,omitempty"`
			Receiver              *string   `tfsdk:"receiver" json:"receiver,omitempty"`
			Repeat_interval       *string   `tfsdk:"repeat_interval" json:"repeat_interval,omitempty"`
			Routes                *[]string `tfsdk:"routes" json:"routes,omitempty"`
		} `tfsdk:"route" json:"route,omitempty"`
		Time_intervals *[]struct {
			Name           *string `tfsdk:"name" json:"name,omitempty"`
			Time_intervals *[]struct {
				Days_of_month *[]string `tfsdk:"days_of_month" json:"days_of_month,omitempty"`
				Location      *string   `tfsdk:"location" json:"location,omitempty"`
				Months        *[]string `tfsdk:"months" json:"months,omitempty"`
				Times         *[]struct {
					End_time   *string `tfsdk:"end_time" json:"end_time,omitempty"`
					Start_time *string `tfsdk:"start_time" json:"start_time,omitempty"`
				} `tfsdk:"times" json:"times,omitempty"`
				Weekdays *[]string `tfsdk:"weekdays" json:"weekdays,omitempty"`
				Years    *[]string `tfsdk:"years" json:"years,omitempty"`
			} `tfsdk:"time_intervals" json:"time_intervals,omitempty"`
		} `tfsdk:"time_intervals" json:"time_intervals,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *OperatorVictoriametricsComVmalertmanagerConfigV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_operator_victoriametrics_com_vm_alertmanager_config_v1beta1_manifest"
}

func (r *OperatorVictoriametricsComVmalertmanagerConfigV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VMAlertmanagerConfig is the Schema for the vmalertmanagerconfigs API",
		MarkdownDescription: "VMAlertmanagerConfig is the Schema for the vmalertmanagerconfigs API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "VMAlertmanagerConfigSpec defines configuration for VMAlertmanagerConfig",
				MarkdownDescription: "VMAlertmanagerConfigSpec defines configuration for VMAlertmanagerConfig",
				Attributes: map[string]schema.Attribute{
					"inhibit_rules": schema.ListNestedAttribute{
						Description:         "InhibitRules will only apply for alerts matchingthe resource's namespace.",
						MarkdownDescription: "InhibitRules will only apply for alerts matchingthe resource's namespace.",
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

								"source_matchers": schema.ListAttribute{
									Description:         "SourceMatchers defines a list of matchers for which one or more alerts haveto exist for the inhibition to take effect.",
									MarkdownDescription: "SourceMatchers defines a list of matchers for which one or more alerts haveto exist for the inhibition to take effect.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_matchers": schema.ListAttribute{
									Description:         "TargetMatchers defines a list of matchers that have to be fulfilled by the targetalerts to be muted.",
									MarkdownDescription: "TargetMatchers defines a list of matchers that have to be fulfilled by the targetalerts to be muted.",
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

					"mute_time_intervals": schema.ListNestedAttribute{
						Description:         "MuteTimeInterval - global mute timeSee https://prometheus.io/docs/alerting/latest/configuration/#mute_time_interval",
						MarkdownDescription: "MuteTimeInterval - global mute timeSee https://prometheus.io/docs/alerting/latest/configuration/#mute_time_interval",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of interval",
									MarkdownDescription: "Name of interval",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"time_intervals": schema.ListNestedAttribute{
									Description:         "TimeIntervals interval configuration",
									MarkdownDescription: "TimeIntervals interval configuration",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"days_of_month": schema.ListAttribute{
												Description:         "DayOfMonth defines list of numerical days in the month. Days begin at 1. Negative values are also accepted.for example, ['1:5', '-3:-1']",
												MarkdownDescription: "DayOfMonth defines list of numerical days in the month. Days begin at 1. Negative values are also accepted.for example, ['1:5', '-3:-1']",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"location": schema.StringAttribute{
												Description:         "Location in golang time location form, e.g. UTC",
												MarkdownDescription: "Location in golang time location form, e.g. UTC",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"months": schema.ListAttribute{
												Description:         "Months  defines list of calendar months identified by a case-insentive name (e.g. ‘January’) or numeric 1.For example, ['1:3', 'may:august', 'december']",
												MarkdownDescription: "Months  defines list of calendar months identified by a case-insentive name (e.g. ‘January’) or numeric 1.For example, ['1:3', 'may:august', 'december']",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"times": schema.ListNestedAttribute{
												Description:         "Times defines time range for mute",
												MarkdownDescription: "Times defines time range for mute",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"end_time": schema.StringAttribute{
															Description:         "EndTime for example HH:MM",
															MarkdownDescription: "EndTime for example HH:MM",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"start_time": schema.StringAttribute{
															Description:         "StartTime for example  HH:MM",
															MarkdownDescription: "StartTime for example  HH:MM",
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

											"weekdays": schema.ListAttribute{
												Description:         "Weekdays defines list of days of the week, where the week begins on Sunday and ends on Saturday.",
												MarkdownDescription: "Weekdays defines list of days of the week, where the week begins on Sunday and ends on Saturday.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"years": schema.ListAttribute{
												Description:         "Years defines numerical list of years, ranges are accepted.For example, ['2020:2022', '2030']",
												MarkdownDescription: "Years defines numerical list of years, ranges are accepted.For example, ['2020:2022', '2030']",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
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

					"receivers": schema.ListNestedAttribute{
						Description:         "Receivers defines alert receivers.without defined Route, receivers will be skipped.",
						MarkdownDescription: "Receivers defines alert receivers.without defined Route, receivers will be skipped.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"discord_configs": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"basic_auth": schema.SingleNestedAttribute{
														Description:         "TODO oAuth2 supportBasicAuth for the client.",
														MarkdownDescription: "TODO oAuth2 supportBasicAuth for the client.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"password_file": schema.StringAttribute{
																Description:         "PasswordFile defines path to password file at disk",
																MarkdownDescription: "PasswordFile defines path to password file at disk",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_file": schema.StringAttribute{
														Description:         "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														MarkdownDescription: "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
														MarkdownDescription: "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
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
																Description:         "Stuct containing the CA cert to use for the targets.",
																MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"ca_file": schema.StringAttribute{
																Description:         "Path to the CA cert in the container to use for the targets.",
																MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert": schema.SingleNestedAttribute{
																Description:         "Struct containing the client cert file for the targets.",
																MarkdownDescription: "Struct containing the client cert file for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"cert_file": schema.StringAttribute{
																Description:         "Path to the client cert file in the container for the targets.",
																MarkdownDescription: "Path to the client cert file in the container for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_file": schema.StringAttribute{
																Description:         "Path to the client key file in the container for the targets.",
																MarkdownDescription: "Path to the client key file in the container for the targets.",
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
												Description:         "The message body template",
												MarkdownDescription: "The message body template",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title": schema.StringAttribute{
												Description:         "The message title template",
												MarkdownDescription: "The message title template",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"webhook_url": schema.StringAttribute{
												Description:         "The discord webhook URLone of 'urlSecret' and 'url' must be defined.",
												MarkdownDescription: "The discord webhook URLone of 'urlSecret' and 'url' must be defined.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"webhook_url_secret": schema.SingleNestedAttribute{
												Description:         "URLSecret defines secret name and key at the CRD namespace.It must contain the webhook URL.one of 'urlSecret' and 'url' must be defined.",
												MarkdownDescription: "URLSecret defines secret name and key at the CRD namespace.It must contain the webhook URL.one of 'urlSecret' and 'url' must be defined.",
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

								"email_configs": schema.ListNestedAttribute{
									Description:         "EmailConfigs defines email notification configurations.",
									MarkdownDescription: "EmailConfigs defines email notification configurations.",
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
												Description:         "AuthPassword defines secret name and key at CRD namespace.",
												MarkdownDescription: "AuthPassword defines secret name and key at CRD namespace.",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auth_secret": schema.SingleNestedAttribute{
												Description:         "AuthSecret defines secrent name and key at CRD namespace.It must contain the CRAM-MD5 secret.",
												MarkdownDescription: "AuthSecret defines secrent name and key at CRD namespace.It must contain the CRAM-MD5 secret.",
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

											"headers": schema.MapAttribute{
												Description:         "Further headers email header key/value pairs. Overrides any headerspreviously set by the notification implementation.",
												MarkdownDescription: "Further headers email header key/value pairs. Overrides any headerspreviously set by the notification implementation.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
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
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"smarthost": schema.StringAttribute{
												Description:         "The SMTP host through which emails are sent.",
												MarkdownDescription: "The SMTP host through which emails are sent.",
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
														Description:         "Stuct containing the CA cert to use for the targets.",
														MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
																		Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ca_file": schema.StringAttribute{
														Description:         "Path to the CA cert in the container to use for the targets.",
														MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert": schema.SingleNestedAttribute{
														Description:         "Struct containing the client cert file for the targets.",
														MarkdownDescription: "Struct containing the client cert file for the targets.",
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
																		Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"cert_file": schema.StringAttribute{
														Description:         "Path to the client cert file in the container for the targets.",
														MarkdownDescription: "Path to the client cert file in the container for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_file": schema.StringAttribute{
														Description:         "Path to the client key file in the container for the targets.",
														MarkdownDescription: "Path to the client key file in the container for the targets.",
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

								"msteams_configs": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"basic_auth": schema.SingleNestedAttribute{
														Description:         "TODO oAuth2 supportBasicAuth for the client.",
														MarkdownDescription: "TODO oAuth2 supportBasicAuth for the client.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"password_file": schema.StringAttribute{
																Description:         "PasswordFile defines path to password file at disk",
																MarkdownDescription: "PasswordFile defines path to password file at disk",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_file": schema.StringAttribute{
														Description:         "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														MarkdownDescription: "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
														MarkdownDescription: "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
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
																Description:         "Stuct containing the CA cert to use for the targets.",
																MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"ca_file": schema.StringAttribute{
																Description:         "Path to the CA cert in the container to use for the targets.",
																MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert": schema.SingleNestedAttribute{
																Description:         "Struct containing the client cert file for the targets.",
																MarkdownDescription: "Struct containing the client cert file for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"cert_file": schema.StringAttribute{
																Description:         "Path to the client cert file in the container for the targets.",
																MarkdownDescription: "Path to the client cert file in the container for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_file": schema.StringAttribute{
																Description:         "Path to the client key file in the container for the targets.",
																MarkdownDescription: "Path to the client key file in the container for the targets.",
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

											"send_resolved": schema.BoolAttribute{
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"text": schema.StringAttribute{
												Description:         "The text body of the teams notification.",
												MarkdownDescription: "The text body of the teams notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title": schema.StringAttribute{
												Description:         "The title of the teams notification.",
												MarkdownDescription: "The title of the teams notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"webhook_url": schema.StringAttribute{
												Description:         "The incoming webhook URLone of 'urlSecret' and 'url' must be defined.",
												MarkdownDescription: "The incoming webhook URLone of 'urlSecret' and 'url' must be defined.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"webhook_url_secret": schema.SingleNestedAttribute{
												Description:         "URLSecret defines secret name and key at the CRD namespace.It must contain the webhook URL.one of 'urlSecret' and 'url' must be defined.",
												MarkdownDescription: "URLSecret defines secret name and key at the CRD namespace.It must contain the webhook URL.one of 'urlSecret' and 'url' must be defined.",
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
									Description:         "OpsGenieConfigs defines ops genie notification configurations.",
									MarkdownDescription: "OpsGenieConfigs defines ops genie notification configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"actions": schema.StringAttribute{
												Description:         "Comma separated list of actions that will be available for the alert.",
												MarkdownDescription: "Comma separated list of actions that will be available for the alert.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"api_url": schema.StringAttribute{
												Description:         "The URL to send OpsGenie API requests to.",
												MarkdownDescription: "The URL to send OpsGenie API requests to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"api_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the OpsGenie API key.It must be at them same namespace as CRD",
												MarkdownDescription: "The secret's key that contains the OpsGenie API key.It must be at them same namespace as CRD",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"description": schema.StringAttribute{
												Description:         "Description of the incident.",
												MarkdownDescription: "Description of the incident.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"details": schema.MapAttribute{
												Description:         "A set of arbitrary key/value pairs that provide further detail about the incident.",
												MarkdownDescription: "A set of arbitrary key/value pairs that provide further detail about the incident.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"entity": schema.StringAttribute{
												Description:         "Optional field that can be used to specify which domain alert is related to.",
												MarkdownDescription: "Optional field that can be used to specify which domain alert is related to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.MapAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
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
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
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
									Description:         "PagerDutyConfigs defines pager duty notification configurations.",
									MarkdownDescription: "PagerDutyConfigs defines pager duty notification configurations.",
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

											"details": schema.MapAttribute{
												Description:         "Arbitrary key/value pairs that provide further detail about the incident.",
												MarkdownDescription: "Arbitrary key/value pairs that provide further detail about the incident.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"group": schema.StringAttribute{
												Description:         "A cluster or grouping of sources.",
												MarkdownDescription: "A cluster or grouping of sources.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.MapAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"images": schema.ListNestedAttribute{
												Description:         "Images to attach to the incident.",
												MarkdownDescription: "Images to attach to the incident.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"alt": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"href": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"source": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"links": schema.ListNestedAttribute{
												Description:         "Links to attach to the incident.",
												MarkdownDescription: "Links to attach to the incident.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"href": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"text": schema.StringAttribute{
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

											"routing_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the PagerDuty integration key (when usingEvents API v2). Either this field or 'serviceKey' needs to be defined.It must be at them same namespace as CRD",
												MarkdownDescription: "The secret's key that contains the PagerDuty integration key (when usingEvents API v2). Either this field or 'serviceKey' needs to be defined.It must be at them same namespace as CRD",
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
												Required: false,
												Optional: true,
												Computed: false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the PagerDuty service key (when usingintegration type 'Prometheus'). Either this field or 'routingKey' needs tobe defined.It must be at them same namespace as CRD",
												MarkdownDescription: "The secret's key that contains the PagerDuty service key (when usingintegration type 'Prometheus'). Either this field or 'routingKey' needs tobe defined.It must be at them same namespace as CRD",
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
									Description:         "PushoverConfigs defines push over notification configurations.",
									MarkdownDescription: "PushoverConfigs defines push over notification configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"expire": schema.StringAttribute{
												Description:         "How long your notification will continue to be retried for, unless the useracknowledges the notification.",
												MarkdownDescription: "How long your notification will continue to be retried for, unless the useracknowledges the notification.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"html": schema.BoolAttribute{
												Description:         "Whether notification message is HTML or plain text.",
												MarkdownDescription: "Whether notification message is HTML or plain text.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.MapAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
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
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
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
												Description:         "The secret's key that contains the registered application’s API token, see https://pushover.net/apps.It must be at them same namespace as CRD",
												MarkdownDescription: "The secret's key that contains the registered application’s API token, see https://pushover.net/apps.It must be at them same namespace as CRD",
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
												Description:         "The secret's key that contains the recipient user’s user key.It must be at them same namespace as CRD",
												MarkdownDescription: "The secret's key that contains the recipient user’s user key.It must be at them same namespace as CRD",
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
									Description:         "SlackConfigs defines slack notification configurations.",
									MarkdownDescription: "SlackConfigs defines slack notification configurations.",
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
												Description:         "The secret's key that contains the Slack webhook URL.It must be at them same namespace as CRD",
												MarkdownDescription: "The secret's key that contains the Slack webhook URL.It must be at them same namespace as CRD",
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

											"http_config": schema.MapAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
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
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
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
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_url": schema.StringAttribute{
												Description:         "The api URL",
												MarkdownDescription: "The api URL",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"attributes": schema.MapAttribute{
												Description:         "SNS message attributes",
												MarkdownDescription: "SNS message attributes",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												Attributes: map[string]schema.Attribute{
													"basic_auth": schema.SingleNestedAttribute{
														Description:         "TODO oAuth2 supportBasicAuth for the client.",
														MarkdownDescription: "TODO oAuth2 supportBasicAuth for the client.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"password_file": schema.StringAttribute{
																Description:         "PasswordFile defines path to password file at disk",
																MarkdownDescription: "PasswordFile defines path to password file at disk",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_file": schema.StringAttribute{
														Description:         "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														MarkdownDescription: "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
														MarkdownDescription: "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
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
																Description:         "Stuct containing the CA cert to use for the targets.",
																MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"ca_file": schema.StringAttribute{
																Description:         "Path to the CA cert in the container to use for the targets.",
																MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert": schema.SingleNestedAttribute{
																Description:         "Struct containing the client cert file for the targets.",
																MarkdownDescription: "Struct containing the client cert file for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"cert_file": schema.StringAttribute{
																Description:         "Path to the client cert file in the container for the targets.",
																MarkdownDescription: "Path to the client cert file in the container for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_file": schema.StringAttribute{
																Description:         "Path to the client key file in the container for the targets.",
																MarkdownDescription: "Path to the client key file in the container for the targets.",
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
												Description:         "Phone number if message is delivered via SMSSpecify this, topic_arn or target_arn",
												MarkdownDescription: "Phone number if message is delivered via SMSSpecify this, topic_arn or target_arn",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sigv4": schema.SingleNestedAttribute{
												Description:         "Configure the AWS Signature Verification 4 signing process",
												MarkdownDescription: "Configure the AWS Signature Verification 4 signing process",
												Attributes: map[string]schema.Attribute{
													"access_key": schema.StringAttribute{
														Description:         "The AWS API keys. Both access_key and secret_key must be supplied or both must be blank.If blank the environment variables 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY' are used.",
														MarkdownDescription: "The AWS API keys. Both access_key and secret_key must be supplied or both must be blank.If blank the environment variables 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY' are used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"access_key_selector": schema.SingleNestedAttribute{
														Description:         "secret key selector to get the keys from a Kubernetes Secret",
														MarkdownDescription: "secret key selector to get the keys from a Kubernetes Secret",
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"profile": schema.StringAttribute{
														Description:         "Named AWS profile used to authenticate",
														MarkdownDescription: "Named AWS profile used to authenticate",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"region": schema.StringAttribute{
														Description:         "AWS region, if blank the region from the default credentials chain is used",
														MarkdownDescription: "AWS region, if blank the region from the default credentials chain is used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"role_arn": schema.StringAttribute{
														Description:         "AWS Role ARN, an alternative to using AWS API keys",
														MarkdownDescription: "AWS Role ARN, an alternative to using AWS API keys",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_key_selector": schema.SingleNestedAttribute{
														Description:         "secret key selector to get the keys from a Kubernetes Secret",
														MarkdownDescription: "secret key selector to get the keys from a Kubernetes Secret",
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
												Description:         "The subject line if message is delivered to an email endpoint.",
												MarkdownDescription: "The subject line if message is delivered to an email endpoint.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_arn": schema.StringAttribute{
												Description:         "Mobile platform endpoint ARN if message is delivered via mobile notificationsSpecify this, topic_arn or phone_number",
												MarkdownDescription: "Mobile platform endpoint ARN if message is delivered via mobile notificationsSpecify this, topic_arn or phone_number",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"topic_arn": schema.StringAttribute{
												Description:         "SNS topic ARN, either specify this, phone_number or target_arn",
												MarkdownDescription: "SNS topic ARN, either specify this, phone_number or target_arn",
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
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_url": schema.StringAttribute{
												Description:         "APIUrl the Telegram API URL i.e. https://api.telegram.org.",
												MarkdownDescription: "APIUrl the Telegram API URL i.e. https://api.telegram.org.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"bot_token": schema.SingleNestedAttribute{
												Description:         "BotToken token for the bothttps://core.telegram.org/bots/api",
												MarkdownDescription: "BotToken token for the bothttps://core.telegram.org/bots/api",
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

											"chat_id": schema.Int64Attribute{
												Description:         "ChatID is ID of the chat where to send the messages.",
												MarkdownDescription: "ChatID is ID of the chat where to send the messages.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"disable_notifications": schema.BoolAttribute{
												Description:         "DisableNotifications",
												MarkdownDescription: "DisableNotifications",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.MapAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"message": schema.StringAttribute{
												Description:         "Message is templated message",
												MarkdownDescription: "Message is templated message",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"parse_mode": schema.StringAttribute{
												Description:         "ParseMode for telegram message,supported values are MarkdownV2, Markdown, Markdown and empty string for plain text.",
												MarkdownDescription: "ParseMode for telegram message,supported values are MarkdownV2, Markdown, Markdown and empty string for plain text.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
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
									Description:         "VictorOpsConfigs defines victor ops notification configurations.",
									MarkdownDescription: "VictorOpsConfigs defines victor ops notification configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_key": schema.SingleNestedAttribute{
												Description:         "The secret's key that contains the API key to use when talking to the VictorOps API.It must be at them same namespace as CRD",
												MarkdownDescription: "The secret's key that contains the API key to use when talking to the VictorOps API.It must be at them same namespace as CRD",
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

											"custom_fields": schema.MapAttribute{
												Description:         "Adds optional custom fieldshttps://github.com/prometheus/alertmanager/blob/v0.24.0/config/notifiers.go#L537",
												MarkdownDescription: "Adds optional custom fieldshttps://github.com/prometheus/alertmanager/blob/v0.24.0/config/notifiers.go#L537",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
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
													"basic_auth": schema.SingleNestedAttribute{
														Description:         "TODO oAuth2 supportBasicAuth for the client.",
														MarkdownDescription: "TODO oAuth2 supportBasicAuth for the client.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"password_file": schema.StringAttribute{
																Description:         "PasswordFile defines path to password file at disk",
																MarkdownDescription: "PasswordFile defines path to password file at disk",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_file": schema.StringAttribute{
														Description:         "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														MarkdownDescription: "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
														MarkdownDescription: "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
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
																Description:         "Stuct containing the CA cert to use for the targets.",
																MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"ca_file": schema.StringAttribute{
																Description:         "Path to the CA cert in the container to use for the targets.",
																MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert": schema.SingleNestedAttribute{
																Description:         "Struct containing the client cert file for the targets.",
																MarkdownDescription: "Struct containing the client cert file for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"cert_file": schema.StringAttribute{
																Description:         "Path to the client cert file in the container for the targets.",
																MarkdownDescription: "Path to the client cert file in the container for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_file": schema.StringAttribute{
																Description:         "Path to the client key file in the container for the targets.",
																MarkdownDescription: "Path to the client key file in the container for the targets.",
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
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
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
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"api_url": schema.StringAttribute{
												Description:         "The Webex Teams API URL, i.e. https://webexapis.com/v1/messages",
												MarkdownDescription: "The Webex Teams API URL, i.e. https://webexapis.com/v1/messages",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_config": schema.SingleNestedAttribute{
												Description:         "HTTP client configuration. You must use this configuration to supply the bot token as part of the HTTP 'Authorization' header.",
												MarkdownDescription: "HTTP client configuration. You must use this configuration to supply the bot token as part of the HTTP 'Authorization' header.",
												Attributes: map[string]schema.Attribute{
													"basic_auth": schema.SingleNestedAttribute{
														Description:         "TODO oAuth2 supportBasicAuth for the client.",
														MarkdownDescription: "TODO oAuth2 supportBasicAuth for the client.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"password_file": schema.StringAttribute{
																Description:         "PasswordFile defines path to password file at disk",
																MarkdownDescription: "PasswordFile defines path to password file at disk",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_file": schema.StringAttribute{
														Description:         "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														MarkdownDescription: "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
														MarkdownDescription: "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
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
																Description:         "Stuct containing the CA cert to use for the targets.",
																MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"ca_file": schema.StringAttribute{
																Description:         "Path to the CA cert in the container to use for the targets.",
																MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert": schema.SingleNestedAttribute{
																Description:         "Struct containing the client cert file for the targets.",
																MarkdownDescription: "Struct containing the client cert file for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"cert_file": schema.StringAttribute{
																Description:         "Path to the client cert file in the container for the targets.",
																MarkdownDescription: "Path to the client cert file in the container for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_file": schema.StringAttribute{
																Description:         "Path to the client key file in the container for the targets.",
																MarkdownDescription: "Path to the client key file in the container for the targets.",
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
												Description:         "The message body template",
												MarkdownDescription: "The message body template",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"room_id": schema.StringAttribute{
												Description:         "The ID of the Webex Teams room where to send the messages",
												MarkdownDescription: "The ID of the Webex Teams room where to send the messages",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"send_resolved": schema.BoolAttribute{
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
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
									Description:         "WebhookConfigs defines webhook notification configurations.",
									MarkdownDescription: "WebhookConfigs defines webhook notification configurations.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"http_config": schema.MapAttribute{
												Description:         "HTTP client configuration.",
												MarkdownDescription: "HTTP client configuration.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
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
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url": schema.StringAttribute{
												Description:         "URL to send requests to,one of 'urlSecret' and 'url' must be defined.",
												MarkdownDescription: "URL to send requests to,one of 'urlSecret' and 'url' must be defined.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url_secret": schema.SingleNestedAttribute{
												Description:         "URLSecret defines secret name and key at the CRD namespace.It must contain the webhook URL.one of 'urlSecret' and 'url' must be defined.",
												MarkdownDescription: "URLSecret defines secret name and key at the CRD namespace.It must contain the webhook URL.one of 'urlSecret' and 'url' must be defined.",
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
									Description:         "WeChatConfigs defines wechat notification configurations.",
									MarkdownDescription: "WeChatConfigs defines wechat notification configurations.",
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
													"basic_auth": schema.SingleNestedAttribute{
														Description:         "TODO oAuth2 supportBasicAuth for the client.",
														MarkdownDescription: "TODO oAuth2 supportBasicAuth for the client.",
														Attributes: map[string]schema.Attribute{
															"password": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the passwordfor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},

															"password_file": schema.StringAttribute{
																Description:         "PasswordFile defines path to password file at disk",
																MarkdownDescription: "PasswordFile defines path to password file at disk",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"username": schema.SingleNestedAttribute{
																Description:         "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
																MarkdownDescription: "The secret in the service scrape namespace that contains the usernamefor authentication.It must be at them same namespace as CRD",
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
																Required: false,
																Optional: true,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"bearer_token_file": schema.StringAttribute{
														Description:         "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														MarkdownDescription: "BearerTokenFile defines filename for bearer token, it must be mounted to pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"bearer_token_secret": schema.SingleNestedAttribute{
														Description:         "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
														MarkdownDescription: "The secret's key that contains the bearer tokenIt must be at them same namespace as CRD",
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
																Description:         "Stuct containing the CA cert to use for the targets.",
																MarkdownDescription: "Stuct containing the CA cert to use for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"ca_file": schema.StringAttribute{
																Description:         "Path to the CA cert in the container to use for the targets.",
																MarkdownDescription: "Path to the CA cert in the container to use for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"cert": schema.SingleNestedAttribute{
																Description:         "Struct containing the client cert file for the targets.",
																MarkdownDescription: "Struct containing the client cert file for the targets.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
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
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"cert_file": schema.StringAttribute{
																Description:         "Path to the client cert file in the container for the targets.",
																MarkdownDescription: "Path to the client cert file in the container for the targets.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "Disable target certificate validation.",
																MarkdownDescription: "Disable target certificate validation.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key_file": schema.StringAttribute{
																Description:         "Path to the client key file in the container for the targets.",
																MarkdownDescription: "Path to the client key file in the container for the targets.",
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
												Description:         "SendResolved controls notify about resolved alerts.",
												MarkdownDescription: "SendResolved controls notify about resolved alerts.",
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
						Description:         "Route definition for alertmanager, may include nested routes.",
						MarkdownDescription: "Route definition for alertmanager, may include nested routes.",
						Attributes: map[string]schema.Attribute{
							"active_time_intervals": schema.ListAttribute{
								Description:         "ActiveTimeIntervals Times when the route should be activeThese must match the name at time_intervals",
								MarkdownDescription: "ActiveTimeIntervals Times when the route should be activeThese must match the name at time_intervals",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"continue": schema.BoolAttribute{
								Description:         "Continue indicating whether an alert should continue matching subsequentsibling nodes. It will always be true for the first-level route if disableRouteContinueEnforce for vmalertmanager not set.",
								MarkdownDescription: "Continue indicating whether an alert should continue matching subsequentsibling nodes. It will always be true for the first-level route if disableRouteContinueEnforce for vmalertmanager not set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_by": schema.ListAttribute{
								Description:         "List of labels to group by.",
								MarkdownDescription: "List of labels to group by.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_interval": schema.StringAttribute{
								Description:         "How long to wait before sending an updated notification.",
								MarkdownDescription: "How long to wait before sending an updated notification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
								},
							},

							"group_wait": schema.StringAttribute{
								Description:         "How long to wait before sending the initial notification.",
								MarkdownDescription: "How long to wait before sending the initial notification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
								},
							},

							"matchers": schema.ListAttribute{
								Description:         "List of matchers that the alert’s labels should match. For the firstlevel route, the operator adds a namespace: 'CRD_NS' matcher.https://prometheus.io/docs/alerting/latest/configuration/#matcher",
								MarkdownDescription: "List of matchers that the alert’s labels should match. For the firstlevel route, the operator adds a namespace: 'CRD_NS' matcher.https://prometheus.io/docs/alerting/latest/configuration/#matcher",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mute_time_intervals": schema.ListAttribute{
								Description:         "MuteTimeIntervals for alerts",
								MarkdownDescription: "MuteTimeIntervals for alerts",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"receiver": schema.StringAttribute{
								Description:         "Name of the receiver for this route.",
								MarkdownDescription: "Name of the receiver for this route.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"repeat_interval": schema.StringAttribute{
								Description:         "How long to wait before repeating the last notification.",
								MarkdownDescription: "How long to wait before repeating the last notification.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`[0-9]+(ms|s|m|h)`), ""),
								},
							},

							"routes": schema.ListAttribute{
								Description:         "Child routes.https://prometheus.io/docs/alerting/latest/configuration/#route",
								MarkdownDescription: "Child routes.https://prometheus.io/docs/alerting/latest/configuration/#route",
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
						Description:         "ParsingError contents error with context if operator was failed to parse json object from kubernetes api serverTimeIntervals modern config option, use it instead of  mute_time_intervals",
						MarkdownDescription: "ParsingError contents error with context if operator was failed to parse json object from kubernetes api serverTimeIntervals modern config option, use it instead of  mute_time_intervals",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of interval",
									MarkdownDescription: "Name of interval",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"time_intervals": schema.ListNestedAttribute{
									Description:         "TimeIntervals interval configuration",
									MarkdownDescription: "TimeIntervals interval configuration",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"days_of_month": schema.ListAttribute{
												Description:         "DayOfMonth defines list of numerical days in the month. Days begin at 1. Negative values are also accepted.for example, ['1:5', '-3:-1']",
												MarkdownDescription: "DayOfMonth defines list of numerical days in the month. Days begin at 1. Negative values are also accepted.for example, ['1:5', '-3:-1']",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"location": schema.StringAttribute{
												Description:         "Location in golang time location form, e.g. UTC",
												MarkdownDescription: "Location in golang time location form, e.g. UTC",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"months": schema.ListAttribute{
												Description:         "Months  defines list of calendar months identified by a case-insentive name (e.g. ‘January’) or numeric 1.For example, ['1:3', 'may:august', 'december']",
												MarkdownDescription: "Months  defines list of calendar months identified by a case-insentive name (e.g. ‘January’) or numeric 1.For example, ['1:3', 'may:august', 'december']",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"times": schema.ListNestedAttribute{
												Description:         "Times defines time range for mute",
												MarkdownDescription: "Times defines time range for mute",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"end_time": schema.StringAttribute{
															Description:         "EndTime for example HH:MM",
															MarkdownDescription: "EndTime for example HH:MM",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"start_time": schema.StringAttribute{
															Description:         "StartTime for example  HH:MM",
															MarkdownDescription: "StartTime for example  HH:MM",
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

											"weekdays": schema.ListAttribute{
												Description:         "Weekdays defines list of days of the week, where the week begins on Sunday and ends on Saturday.",
												MarkdownDescription: "Weekdays defines list of days of the week, where the week begins on Sunday and ends on Saturday.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"years": schema.ListAttribute{
												Description:         "Years defines numerical list of years, ranges are accepted.For example, ['2020:2022', '2030']",
												MarkdownDescription: "Years defines numerical list of years, ranges are accepted.For example, ['2020:2022', '2030']",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *OperatorVictoriametricsComVmalertmanagerConfigV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_operator_victoriametrics_com_vm_alertmanager_config_v1beta1_manifest")

	var model OperatorVictoriametricsComVmalertmanagerConfigV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("operator.victoriametrics.com/v1beta1")
	model.Kind = pointer.String("VMAlertmanagerConfig")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
