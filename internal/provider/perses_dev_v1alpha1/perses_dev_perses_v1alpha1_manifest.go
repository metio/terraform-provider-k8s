/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package perses_dev_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &PersesDevPersesV1Alpha1Manifest{}
)

func NewPersesDevPersesV1Alpha1Manifest() datasource.DataSource {
	return &PersesDevPersesV1Alpha1Manifest{}
}

type PersesDevPersesV1Alpha1Manifest struct{}

type PersesDevPersesV1Alpha1ManifestData struct {
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
		Affinity *struct {
			NodeAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					Preference *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchFields *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_fields" json:"matchFields,omitempty"`
					} `tfsdk:"preference" json:"preference,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *struct {
					NodeSelectorTerms *[]struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchFields *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_fields" json:"matchFields,omitempty"`
					} `tfsdk:"node_selector_terms" json:"nodeSelectorTerms,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
			PodAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
						MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
					MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
			PodAntiAffinity *struct {
				PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
					PodAffinityTerm *struct {
						LabelSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
						MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
						NamespaceSelector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
						Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
						TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
					} `tfsdk:"pod_affinity_term" json:"podAffinityTerm,omitempty"`
					Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
				} `tfsdk:"preferred_during_scheduling_ignored_during_execution" json:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
				RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
					MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					Namespaces  *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
					TopologyKey *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		Args   *[]string `tfsdk:"args" json:"args,omitempty"`
		Client *struct {
			BasicAuth *struct {
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Password_path *string `tfsdk:"password_path" json:"password_path,omitempty"`
				Type          *string `tfsdk:"type" json:"type,omitempty"`
				Username      *string `tfsdk:"username" json:"username,omitempty"`
			} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
			Oauth *struct {
				AuthStyle        *int64               `tfsdk:"auth_style" json:"authStyle,omitempty"`
				ClientIDPath     *string              `tfsdk:"client_id_path" json:"clientIDPath,omitempty"`
				ClientSecretPath *string              `tfsdk:"client_secret_path" json:"clientSecretPath,omitempty"`
				EndpointParams   *map[string][]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
				Name             *string              `tfsdk:"name" json:"name,omitempty"`
				Namespace        *string              `tfsdk:"namespace" json:"namespace,omitempty"`
				Scopes           *[]string            `tfsdk:"scopes" json:"scopes,omitempty"`
				TokenURL         *string              `tfsdk:"token_url" json:"tokenURL,omitempty"`
				Type             *string              `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"oauth" json:"oauth,omitempty"`
			Tls *struct {
				CaCert *struct {
					CertPath       *string `tfsdk:"cert_path" json:"certPath,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
					PrivateKeyPath *string `tfsdk:"private_key_path" json:"privateKeyPath,omitempty"`
					Type           *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"ca_cert" json:"caCert,omitempty"`
				Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				UserCert           *struct {
					CertPath       *string `tfsdk:"cert_path" json:"certPath,omitempty"`
					Name           *string `tfsdk:"name" json:"name,omitempty"`
					Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
					PrivateKeyPath *string `tfsdk:"private_key_path" json:"privateKeyPath,omitempty"`
					Type           *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"user_cert" json:"userCert,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"client" json:"client,omitempty"`
		Config *struct {
			Api_prefix *string `tfsdk:"api_prefix" json:"api_prefix,omitempty"`
			Database   *struct {
				File *struct {
					Case_sensitive *bool   `tfsdk:"case_sensitive" json:"case_sensitive,omitempty"`
					Extension      *string `tfsdk:"extension" json:"extension,omitempty"`
					Folder         *string `tfsdk:"folder" json:"folder,omitempty"`
				} `tfsdk:"file" json:"file,omitempty"`
				Sql *struct {
					Addr                        *string            `tfsdk:"addr" json:"addr,omitempty"`
					Allow_all_files             *bool              `tfsdk:"allow_all_files" json:"allow_all_files,omitempty"`
					Allow_cleartext_passwords   *bool              `tfsdk:"allow_cleartext_passwords" json:"allow_cleartext_passwords,omitempty"`
					Allow_fallback_to_plaintext *bool              `tfsdk:"allow_fallback_to_plaintext" json:"allow_fallback_to_plaintext,omitempty"`
					Allow_native_passwords      *bool              `tfsdk:"allow_native_passwords" json:"allow_native_passwords,omitempty"`
					Allow_old_passwords         *bool              `tfsdk:"allow_old_passwords" json:"allow_old_passwords,omitempty"`
					Case_sensitive              *bool              `tfsdk:"case_sensitive" json:"case_sensitive,omitempty"`
					Check_conn_liveness         *bool              `tfsdk:"check_conn_liveness" json:"check_conn_liveness,omitempty"`
					Client_found_rows           *bool              `tfsdk:"client_found_rows" json:"client_found_rows,omitempty"`
					Collation                   *string            `tfsdk:"collation" json:"collation,omitempty"`
					Columns_with_alias          *bool              `tfsdk:"columns_with_alias" json:"columns_with_alias,omitempty"`
					Db_name                     *string            `tfsdk:"db_name" json:"db_name,omitempty"`
					Interpolate_params          *bool              `tfsdk:"interpolate_params" json:"interpolate_params,omitempty"`
					Loc                         *map[string]string `tfsdk:"loc" json:"loc,omitempty"`
					Max_allowed_packet          *int64             `tfsdk:"max_allowed_packet" json:"max_allowed_packet,omitempty"`
					Multi_statements            *bool              `tfsdk:"multi_statements" json:"multi_statements,omitempty"`
					Net                         *string            `tfsdk:"net" json:"net,omitempty"`
					Parse_time                  *bool              `tfsdk:"parse_time" json:"parse_time,omitempty"`
					Password                    *string            `tfsdk:"password" json:"password,omitempty"`
					Password_file               *string            `tfsdk:"password_file" json:"password_file,omitempty"`
					Read_timeout                *string            `tfsdk:"read_timeout" json:"read_timeout,omitempty"`
					Reject_read_only            *bool              `tfsdk:"reject_read_only" json:"reject_read_only,omitempty"`
					Server_pub_key              *string            `tfsdk:"server_pub_key" json:"server_pub_key,omitempty"`
					Timeout                     *string            `tfsdk:"timeout" json:"timeout,omitempty"`
					Tls_config                  *struct {
						Ca                   *string `tfsdk:"ca" json:"ca,omitempty"`
						Ca_file              *string `tfsdk:"ca_file" json:"ca_file,omitempty"`
						Ca_ref               *string `tfsdk:"ca_ref" json:"ca_ref,omitempty"`
						Cert                 *string `tfsdk:"cert" json:"cert,omitempty"`
						Cert_file            *string `tfsdk:"cert_file" json:"cert_file,omitempty"`
						Cert_ref             *string `tfsdk:"cert_ref" json:"cert_ref,omitempty"`
						Insecure_skip_verify *bool   `tfsdk:"insecure_skip_verify" json:"insecure_skip_verify,omitempty"`
						Key                  *string `tfsdk:"key" json:"key,omitempty"`
						Key_file             *string `tfsdk:"key_file" json:"key_file,omitempty"`
						Key_ref              *string `tfsdk:"key_ref" json:"key_ref,omitempty"`
						Max_version          *int64  `tfsdk:"max_version" json:"max_version,omitempty"`
						Min_version          *int64  `tfsdk:"min_version" json:"min_version,omitempty"`
						Server_name          *string `tfsdk:"server_name" json:"server_name,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
					User          *string `tfsdk:"user" json:"user,omitempty"`
					Write_timeout *string `tfsdk:"write_timeout" json:"write_timeout,omitempty"`
				} `tfsdk:"sql" json:"sql,omitempty"`
			} `tfsdk:"database" json:"database,omitempty"`
			Ephemeral_dashboard *struct {
				Cleanup_interval *string `tfsdk:"cleanup_interval" json:"cleanup_interval,omitempty"`
				Enable           *bool   `tfsdk:"enable" json:"enable,omitempty"`
			} `tfsdk:"ephemeral_dashboard" json:"ephemeral_dashboard,omitempty"`
			Ephemeral_dashboards_cleanup_interval *string `tfsdk:"ephemeral_dashboards_cleanup_interval" json:"ephemeral_dashboards_cleanup_interval,omitempty"`
			Frontend                              *struct {
				Disable  *bool `tfsdk:"disable" json:"disable,omitempty"`
				Explorer *struct {
					Enable *bool `tfsdk:"enable" json:"enable,omitempty"`
				} `tfsdk:"explorer" json:"explorer,omitempty"`
				Important_dashboards *[]struct {
					Dashboard *string `tfsdk:"dashboard" json:"dashboard,omitempty"`
					Project   *string `tfsdk:"project" json:"project,omitempty"`
				} `tfsdk:"important_dashboards" json:"important_dashboards,omitempty"`
				Information *string `tfsdk:"information" json:"information,omitempty"`
				Time_range  *struct {
					Disable_custom *bool     `tfsdk:"disable_custom" json:"disable_custom,omitempty"`
					Options        *[]string `tfsdk:"options" json:"options,omitempty"`
				} `tfsdk:"time_range" json:"time_range,omitempty"`
			} `tfsdk:"frontend" json:"frontend,omitempty"`
			Global_datasource_discovery *[]struct {
				Discovery_name *string `tfsdk:"discovery_name" json:"discovery_name,omitempty"`
				Http_sd        *struct {
					Authorization *struct {
						Credentials     *string `tfsdk:"credentials" json:"credentials,omitempty"`
						CredentialsFile *string `tfsdk:"credentials_file" json:"credentialsFile,omitempty"`
						Type            *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"authorization" json:"authorization,omitempty"`
					Basic_auth *struct {
						Password     *string `tfsdk:"password" json:"password,omitempty"`
						PasswordFile *string `tfsdk:"password_file" json:"passwordFile,omitempty"`
						Username     *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"basic_auth" json:"basic_auth,omitempty"`
					Headers     *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
					Native_auth *struct {
						Login    *string `tfsdk:"login" json:"login,omitempty"`
						Password *string `tfsdk:"password" json:"password,omitempty"`
					} `tfsdk:"native_auth" json:"native_auth,omitempty"`
					Oauth *struct {
						AuthStyle        *int64               `tfsdk:"auth_style" json:"authStyle,omitempty"`
						ClientID         *string              `tfsdk:"client_id" json:"clientID,omitempty"`
						ClientSecret     *string              `tfsdk:"client_secret" json:"clientSecret,omitempty"`
						ClientSecretfile *string              `tfsdk:"client_secretfile" json:"clientSecretfile,omitempty"`
						EndpointParams   *map[string][]string `tfsdk:"endpoint_params" json:"endpointParams,omitempty"`
						Scopes           *[]string            `tfsdk:"scopes" json:"scopes,omitempty"`
						TokenURL         *string              `tfsdk:"token_url" json:"tokenURL,omitempty"`
					} `tfsdk:"oauth" json:"oauth,omitempty"`
					Tls_config *struct {
						Ca                 *string `tfsdk:"ca" json:"ca,omitempty"`
						CaFile             *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						Cert               *string `tfsdk:"cert" json:"cert,omitempty"`
						CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Key                *string `tfsdk:"key" json:"key,omitempty"`
						KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
						MaxVersion         *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
						MinVersion         *string `tfsdk:"min_version" json:"minVersion,omitempty"`
						ServerName         *string `tfsdk:"server_name" json:"serverName,omitempty"`
					} `tfsdk:"tls_config" json:"tls_config,omitempty"`
					Url *map[string]string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"http_sd" json:"http_sd,omitempty"`
				Kubernetes_sd *struct {
					Datasource_plugin_kind *string            `tfsdk:"datasource_plugin_kind" json:"datasource_plugin_kind,omitempty"`
					Labels                 *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Namespace              *string            `tfsdk:"namespace" json:"namespace,omitempty"`
					Pod_configuration      *struct {
						Container_name        *string `tfsdk:"container_name" json:"container_name,omitempty"`
						Container_port_name   *string `tfsdk:"container_port_name" json:"container_port_name,omitempty"`
						Container_port_number *int64  `tfsdk:"container_port_number" json:"container_port_number,omitempty"`
						Enable                *bool   `tfsdk:"enable" json:"enable,omitempty"`
					} `tfsdk:"pod_configuration" json:"pod_configuration,omitempty"`
					Service_configuration *struct {
						Enable       *bool   `tfsdk:"enable" json:"enable,omitempty"`
						Port_name    *string `tfsdk:"port_name" json:"port_name,omitempty"`
						Port_number  *int64  `tfsdk:"port_number" json:"port_number,omitempty"`
						Service_type *string `tfsdk:"service_type" json:"service_type,omitempty"`
					} `tfsdk:"service_configuration" json:"service_configuration,omitempty"`
				} `tfsdk:"kubernetes_sd" json:"kubernetes_sd,omitempty"`
				Refresh_interval *string `tfsdk:"refresh_interval" json:"refresh_interval,omitempty"`
			} `tfsdk:"global_datasource_discovery" json:"global_datasource_discovery,omitempty"`
			Provisioning *struct {
				Folders  *[]string `tfsdk:"folders" json:"folders,omitempty"`
				Interval *string   `tfsdk:"interval" json:"interval,omitempty"`
			} `tfsdk:"provisioning" json:"provisioning,omitempty"`
			Schemas *struct {
				Datasources_path *string `tfsdk:"datasources_path" json:"datasources_path,omitempty"`
				Interval         *string `tfsdk:"interval" json:"interval,omitempty"`
				Panels_path      *string `tfsdk:"panels_path" json:"panels_path,omitempty"`
				Queries_path     *string `tfsdk:"queries_path" json:"queries_path,omitempty"`
				Variables_path   *string `tfsdk:"variables_path" json:"variables_path,omitempty"`
			} `tfsdk:"schemas" json:"schemas,omitempty"`
			Security *struct {
				Authentication *struct {
					Access_token_ttl *string `tfsdk:"access_token_ttl" json:"access_token_ttl,omitempty"`
					Disable_sign_up  *bool   `tfsdk:"disable_sign_up" json:"disable_sign_up,omitempty"`
					Providers        *struct {
						Enable_native *bool `tfsdk:"enable_native" json:"enable_native,omitempty"`
						Oauth         *[]struct {
							Auth_url           *map[string]string `tfsdk:"auth_url" json:"auth_url,omitempty"`
							Client_credentials *struct {
								Client_id     *string   `tfsdk:"client_id" json:"client_id,omitempty"`
								Client_secret *string   `tfsdk:"client_secret" json:"client_secret,omitempty"`
								Scopes        *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
							} `tfsdk:"client_credentials" json:"client_credentials,omitempty"`
							Client_id             *string            `tfsdk:"client_id" json:"client_id,omitempty"`
							Client_secret         *string            `tfsdk:"client_secret" json:"client_secret,omitempty"`
							Custom_login_property *string            `tfsdk:"custom_login_property" json:"custom_login_property,omitempty"`
							Device_auth_url       *map[string]string `tfsdk:"device_auth_url" json:"device_auth_url,omitempty"`
							Device_code           *struct {
								Client_id     *string   `tfsdk:"client_id" json:"client_id,omitempty"`
								Client_secret *string   `tfsdk:"client_secret" json:"client_secret,omitempty"`
								Scopes        *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
							} `tfsdk:"device_code" json:"device_code,omitempty"`
							Http *struct {
								Timeout    *string `tfsdk:"timeout" json:"timeout,omitempty"`
								Tls_config *struct {
									Ca                 *string `tfsdk:"ca" json:"ca,omitempty"`
									CaFile             *string `tfsdk:"ca_file" json:"caFile,omitempty"`
									Cert               *string `tfsdk:"cert" json:"cert,omitempty"`
									CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
									InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
									Key                *string `tfsdk:"key" json:"key,omitempty"`
									KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
									MaxVersion         *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
									MinVersion         *string `tfsdk:"min_version" json:"minVersion,omitempty"`
									ServerName         *string `tfsdk:"server_name" json:"serverName,omitempty"`
								} `tfsdk:"tls_config" json:"tls_config,omitempty"`
							} `tfsdk:"http" json:"http,omitempty"`
							Name           *string            `tfsdk:"name" json:"name,omitempty"`
							Redirect_uri   *map[string]string `tfsdk:"redirect_uri" json:"redirect_uri,omitempty"`
							Scopes         *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
							Slug_id        *string            `tfsdk:"slug_id" json:"slug_id,omitempty"`
							Token_url      *map[string]string `tfsdk:"token_url" json:"token_url,omitempty"`
							User_infos_url *map[string]string `tfsdk:"user_infos_url" json:"user_infos_url,omitempty"`
						} `tfsdk:"oauth" json:"oauth,omitempty"`
						Oidc *[]struct {
							Client_credentials *struct {
								Client_id     *string   `tfsdk:"client_id" json:"client_id,omitempty"`
								Client_secret *string   `tfsdk:"client_secret" json:"client_secret,omitempty"`
								Scopes        *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
							} `tfsdk:"client_credentials" json:"client_credentials,omitempty"`
							Client_id     *string `tfsdk:"client_id" json:"client_id,omitempty"`
							Client_secret *string `tfsdk:"client_secret" json:"client_secret,omitempty"`
							Device_code   *struct {
								Client_id     *string   `tfsdk:"client_id" json:"client_id,omitempty"`
								Client_secret *string   `tfsdk:"client_secret" json:"client_secret,omitempty"`
								Scopes        *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
							} `tfsdk:"device_code" json:"device_code,omitempty"`
							Disable_pkce  *bool              `tfsdk:"disable_pkce" json:"disable_pkce,omitempty"`
							Discovery_url *map[string]string `tfsdk:"discovery_url" json:"discovery_url,omitempty"`
							Http          *struct {
								Timeout    *string `tfsdk:"timeout" json:"timeout,omitempty"`
								Tls_config *struct {
									Ca                 *string `tfsdk:"ca" json:"ca,omitempty"`
									CaFile             *string `tfsdk:"ca_file" json:"caFile,omitempty"`
									Cert               *string `tfsdk:"cert" json:"cert,omitempty"`
									CertFile           *string `tfsdk:"cert_file" json:"certFile,omitempty"`
									InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
									Key                *string `tfsdk:"key" json:"key,omitempty"`
									KeyFile            *string `tfsdk:"key_file" json:"keyFile,omitempty"`
									MaxVersion         *string `tfsdk:"max_version" json:"maxVersion,omitempty"`
									MinVersion         *string `tfsdk:"min_version" json:"minVersion,omitempty"`
									ServerName         *string `tfsdk:"server_name" json:"serverName,omitempty"`
								} `tfsdk:"tls_config" json:"tls_config,omitempty"`
							} `tfsdk:"http" json:"http,omitempty"`
							Issuer       *map[string]string `tfsdk:"issuer" json:"issuer,omitempty"`
							Name         *string            `tfsdk:"name" json:"name,omitempty"`
							Redirect_uri *map[string]string `tfsdk:"redirect_uri" json:"redirect_uri,omitempty"`
							Scopes       *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
							Slug_id      *string            `tfsdk:"slug_id" json:"slug_id,omitempty"`
							Url_params   *map[string]string `tfsdk:"url_params" json:"url_params,omitempty"`
						} `tfsdk:"oidc" json:"oidc,omitempty"`
					} `tfsdk:"providers" json:"providers,omitempty"`
					Refresh_token_ttl *string `tfsdk:"refresh_token_ttl" json:"refresh_token_ttl,omitempty"`
				} `tfsdk:"authentication" json:"authentication,omitempty"`
				Authorization *struct {
					Check_latest_update_interval *string `tfsdk:"check_latest_update_interval" json:"check_latest_update_interval,omitempty"`
					Guest_permissions            *[]struct {
						Actions *[]string `tfsdk:"actions" json:"actions,omitempty"`
						Scopes  *[]string `tfsdk:"scopes" json:"scopes,omitempty"`
					} `tfsdk:"guest_permissions" json:"guest_permissions,omitempty"`
				} `tfsdk:"authorization" json:"authorization,omitempty"`
				Cookie *struct {
					Same_site *int64 `tfsdk:"same_site" json:"same_site,omitempty"`
					Secure    *bool  `tfsdk:"secure" json:"secure,omitempty"`
				} `tfsdk:"cookie" json:"cookie,omitempty"`
				Enable_auth         *bool   `tfsdk:"enable_auth" json:"enable_auth,omitempty"`
				Encryption_key      *string `tfsdk:"encryption_key" json:"encryption_key,omitempty"`
				Encryption_key_file *string `tfsdk:"encryption_key_file" json:"encryption_key_file,omitempty"`
				Readonly            *bool   `tfsdk:"readonly" json:"readonly,omitempty"`
			} `tfsdk:"security" json:"security,omitempty"`
		} `tfsdk:"config" json:"config,omitempty"`
		ContainerPort *int64  `tfsdk:"container_port" json:"containerPort,omitempty"`
		Image         *string `tfsdk:"image" json:"image,omitempty"`
		LivenessProbe *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
			} `tfsdk:"exec" json:"exec,omitempty"`
			FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			Grpc             *struct {
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				Service *string `tfsdk:"service" json:"service,omitempty"`
			} `tfsdk:"grpc" json:"grpc,omitempty"`
			HttpGet *struct {
				Host        *string `tfsdk:"host" json:"host,omitempty"`
				HttpHeaders *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Port   *string `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"http_get" json:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
			PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
			SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
			TcpSocket           *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Port *string `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		Metadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"metadata" json:"metadata,omitempty"`
		NodeSelector   *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		ReadinessProbe *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
			} `tfsdk:"exec" json:"exec,omitempty"`
			FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			Grpc             *struct {
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				Service *string `tfsdk:"service" json:"service,omitempty"`
			} `tfsdk:"grpc" json:"grpc,omitempty"`
			HttpGet *struct {
				Host        *string `tfsdk:"host" json:"host,omitempty"`
				HttpHeaders *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"http_headers" json:"httpHeaders,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Port   *string `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"http_get" json:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
			PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
			SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
			TcpSocket           *struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Port *string `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"tcp_socket" json:"tcpSocket,omitempty"`
			TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
			TimeoutSeconds                *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
		Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Service  *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Name        *string            `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		Storage *struct {
			Size         *string `tfsdk:"size" json:"size,omitempty"`
			StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Tls *struct {
			CaCert *struct {
				CertPath       *string `tfsdk:"cert_path" json:"certPath,omitempty"`
				Name           *string `tfsdk:"name" json:"name,omitempty"`
				Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
				PrivateKeyPath *string `tfsdk:"private_key_path" json:"privateKeyPath,omitempty"`
				Type           *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"ca_cert" json:"caCert,omitempty"`
			Enable             *bool `tfsdk:"enable" json:"enable,omitempty"`
			InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
			UserCert           *struct {
				CertPath       *string `tfsdk:"cert_path" json:"certPath,omitempty"`
				Name           *string `tfsdk:"name" json:"name,omitempty"`
				Namespace      *string `tfsdk:"namespace" json:"namespace,omitempty"`
				PrivateKeyPath *string `tfsdk:"private_key_path" json:"privateKeyPath,omitempty"`
				Type           *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"user_cert" json:"userCert,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		Tolerations *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PersesDevPersesV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_perses_dev_perses_v1alpha1_manifest"
}

func (r *PersesDevPersesV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Perses is the Schema for the perses API",
		MarkdownDescription: "Perses is the Schema for the perses API",
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
				Description:         "PersesSpec defines the desired state of Perses",
				MarkdownDescription: "PersesSpec defines the desired state of Perses",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Affinity is a group of affinity scheduling rules.",
						MarkdownDescription: "Affinity is a group of affinity scheduling rules.",
						Attributes: map[string]schema.Attribute{
							"node_affinity": schema.SingleNestedAttribute{
								Description:         "Describes node affinity scheduling rules for the pod.",
								MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"preference": schema.SingleNestedAttribute{
													Description:         "A node selector term, associated with the corresponding weight.",
													MarkdownDescription: "A node selector term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's labels.",
															MarkdownDescription: "A list of node selector requirements by node's labels.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

														"match_fields": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
													Required: true,
													Optional: false,
													Computed: false,
												},

												"weight": schema.Int64Attribute{
													Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
													MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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

									"required_during_scheduling_ignored_during_execution": schema.SingleNestedAttribute{
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
										Attributes: map[string]schema.Attribute{
											"node_selector_terms": schema.ListNestedAttribute{
												Description:         "Required. A list of node selector terms. The terms are ORed.",
												MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's labels.",
															MarkdownDescription: "A list of node selector requirements by node's labels.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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

														"match_fields": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
												Required: true,
												Optional: false,
												Computed: false,
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

							"pod_affinity": schema.SingleNestedAttribute{
								Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Required. A pod affinity term, associated with the corresponding weight.",
													MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"match_label_keys": schema.ListAttribute{
															Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.",
															MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"weight": schema.Int64Attribute{
													Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"match_label_keys": schema.ListAttribute{
													Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.",
													MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mismatch_label_keys": schema.ListAttribute{
													Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.",
													MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"namespaces": schema.ListAttribute{
													Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_anti_affinity": schema.SingleNestedAttribute{
								Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
								MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Required. A pod affinity term, associated with the corresponding weight.",
													MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"match_label_keys": schema.ListAttribute{
															Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.",
															MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"namespaces": schema.ListAttribute{
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"weight": schema.Int64Attribute{
													Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods. If it's null, this PodAffinityTerm matches with no Pods.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"match_label_keys": schema.ListAttribute{
													Description:         "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.",
													MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both matchLabelKeys and labelSelector. Also, matchLabelKeys cannot be set when labelSelector isn't set.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mismatch_label_keys": schema.ListAttribute{
													Description:         "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.",
													MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods will be taken into consideration. The keys are used to lookup values from the incoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)' to select the group of existing pods which pods will be taken into consideration for the incoming pod's pod (anti) affinity. Keys that don't exist in the incoming pod labels will be ignored. The default value is empty. The same key is forbidden to exist in both mismatchLabelKeys and labelSelector. Also, mismatchLabelKeys cannot be set when labelSelector isn't set.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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

														"match_labels": schema.MapAttribute{
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"namespaces": schema.ListAttribute{
													Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

					"args": schema.ListAttribute{
						Description:         "Args extra arguments to pass to perses",
						MarkdownDescription: "Args extra arguments to pass to perses",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"client": schema.SingleNestedAttribute{
						Description:         "Perses client configuration",
						MarkdownDescription: "Perses client configuration",
						Attributes: map[string]schema.Attribute{
							"basic_auth": schema.SingleNestedAttribute{
								Description:         "BasicAuth basic auth config for datasource client",
								MarkdownDescription: "BasicAuth basic auth config for datasource client",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"password_path": schema.StringAttribute{
										Description:         "Path to password",
										MarkdownDescription: "Path to password",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type source type of secret",
										MarkdownDescription: "Type source type of secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("secret", "configmap", "file"),
										},
									},

									"username": schema.StringAttribute{
										Description:         "Username for basic auth",
										MarkdownDescription: "Username for basic auth",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"oauth": schema.SingleNestedAttribute{
								Description:         "OAuth configuration for datasource client",
								MarkdownDescription: "OAuth configuration for datasource client",
								Attributes: map[string]schema.Attribute{
									"auth_style": schema.Int64Attribute{
										Description:         "AuthStyle optionally specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect.",
										MarkdownDescription: "AuthStyle optionally specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_id_path": schema.StringAttribute{
										Description:         "Path to client id",
										MarkdownDescription: "Path to client id",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_secret_path": schema.StringAttribute{
										Description:         "Path to client secret",
										MarkdownDescription: "Path to client secret",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"endpoint_params": schema.MapAttribute{
										Description:         "EndpointParams specifies additional parameters for requests to the token endpoint.",
										MarkdownDescription: "EndpointParams specifies additional parameters for requests to the token endpoint.",
										ElementType:         types.ListType{ElemType: types.StringType},
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scopes": schema.ListAttribute{
										Description:         "Scope specifies optional requested permissions.",
										MarkdownDescription: "Scope specifies optional requested permissions.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"token_url": schema.StringAttribute{
										Description:         "TokenURL is the resource server's token endpoint URL. This is a constant specific to each server.",
										MarkdownDescription: "TokenURL is the resource server's token endpoint URL. This is a constant specific to each server.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type source type of secret",
										MarkdownDescription: "Type source type of secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("secret", "configmap", "file"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS the equivalent to the tls_config for perses client",
								MarkdownDescription: "TLS the equivalent to the tls_config for perses client",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.SingleNestedAttribute{
										Description:         "CaCert to verify the perses certificate",
										MarkdownDescription: "CaCert to verify the perses certificate",
										Attributes: map[string]schema.Attribute{
											"cert_path": schema.StringAttribute{
												Description:         "Path to Certificate",
												MarkdownDescription: "Path to Certificate",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
												MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
												MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"private_key_path": schema.StringAttribute{
												Description:         "Path to Private key certificate",
												MarkdownDescription: "Path to Private key certificate",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type source type of secret",
												MarkdownDescription: "Type source type of secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("secret", "configmap", "file"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable": schema.BoolAttribute{
										Description:         "Enable TLS connection to perses",
										MarkdownDescription: "Enable TLS connection to perses",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "InsecureSkipVerify skip verify of perses certificate",
										MarkdownDescription: "InsecureSkipVerify skip verify of perses certificate",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"user_cert": schema.SingleNestedAttribute{
										Description:         "UserCert client cert/key for mTLS",
										MarkdownDescription: "UserCert client cert/key for mTLS",
										Attributes: map[string]schema.Attribute{
											"cert_path": schema.StringAttribute{
												Description:         "Path to Certificate",
												MarkdownDescription: "Path to Certificate",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
												MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
												MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"private_key_path": schema.StringAttribute{
												Description:         "Path to Private key certificate",
												MarkdownDescription: "Path to Private key certificate",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Type source type of secret",
												MarkdownDescription: "Type source type of secret",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("secret", "configmap", "file"),
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_prefix": schema.StringAttribute{
								Description:         "Use it in case you want to prefix the API path.",
								MarkdownDescription: "Use it in case you want to prefix the API path.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"database": schema.SingleNestedAttribute{
								Description:         "Database contains the different configuration depending on the database you want to use",
								MarkdownDescription: "Database contains the different configuration depending on the database you want to use",
								Attributes: map[string]schema.Attribute{
									"file": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"case_sensitive": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"extension": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"folder": schema.StringAttribute{
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

									"sql": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"addr": schema.StringAttribute{
												Description:         "Network address (requires Net)",
												MarkdownDescription: "Network address (requires Net)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"allow_all_files": schema.BoolAttribute{
												Description:         "Allow all files to be used with LOAD DATA LOCAL INFILE",
												MarkdownDescription: "Allow all files to be used with LOAD DATA LOCAL INFILE",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"allow_cleartext_passwords": schema.BoolAttribute{
												Description:         "Allows the cleartext client side plugin",
												MarkdownDescription: "Allows the cleartext client side plugin",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"allow_fallback_to_plaintext": schema.BoolAttribute{
												Description:         "Allows fallback to unencrypted connection if server does not support TLS",
												MarkdownDescription: "Allows fallback to unencrypted connection if server does not support TLS",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"allow_native_passwords": schema.BoolAttribute{
												Description:         "Allows the native password authentication method",
												MarkdownDescription: "Allows the native password authentication method",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"allow_old_passwords": schema.BoolAttribute{
												Description:         "Allows the old insecure password method",
												MarkdownDescription: "Allows the old insecure password method",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"case_sensitive": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"check_conn_liveness": schema.BoolAttribute{
												Description:         "Check connections for liveness before using them",
												MarkdownDescription: "Check connections for liveness before using them",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"client_found_rows": schema.BoolAttribute{
												Description:         "Return number of matching rows instead of rows changed",
												MarkdownDescription: "Return number of matching rows instead of rows changed",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"collation": schema.StringAttribute{
												Description:         "Connection collation",
												MarkdownDescription: "Connection collation",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"columns_with_alias": schema.BoolAttribute{
												Description:         "Prepend table alias to column names",
												MarkdownDescription: "Prepend table alias to column names",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"db_name": schema.StringAttribute{
												Description:         "Database name",
												MarkdownDescription: "Database name",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"interpolate_params": schema.BoolAttribute{
												Description:         "Interpolate placeholders into query string",
												MarkdownDescription: "Interpolate placeholders into query string",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"loc": schema.MapAttribute{
												Description:         "Location for time.Time values",
												MarkdownDescription: "Location for time.Time values",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_allowed_packet": schema.Int64Attribute{
												Description:         "Max packet size allowed",
												MarkdownDescription: "Max packet size allowed",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"multi_statements": schema.BoolAttribute{
												Description:         "Allow multiple statements in one query",
												MarkdownDescription: "Allow multiple statements in one query",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"net": schema.StringAttribute{
												Description:         "Network type",
												MarkdownDescription: "Network type",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"parse_time": schema.BoolAttribute{
												Description:         "Parse time values to time.Time",
												MarkdownDescription: "Parse time values to time.Time",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"password": schema.StringAttribute{
												Description:         "Password (requires User)",
												MarkdownDescription: "Password (requires User)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"password_file": schema.StringAttribute{
												Description:         "PasswordFile is a path to a file that contains a password",
												MarkdownDescription: "PasswordFile is a path to a file that contains a password",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_timeout": schema.StringAttribute{
												Description:         "I/O read timeout",
												MarkdownDescription: "I/O read timeout",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
												},
											},

											"reject_read_only": schema.BoolAttribute{
												Description:         "Reject read-only connections",
												MarkdownDescription: "Reject read-only connections",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"server_pub_key": schema.StringAttribute{
												Description:         "Server public key name",
												MarkdownDescription: "Server public key name",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"timeout": schema.StringAttribute{
												Description:         "Dial timeout",
												MarkdownDescription: "Dial timeout",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
												},
											},

											"tls_config": schema.SingleNestedAttribute{
												Description:         "TLS configuration",
												MarkdownDescription: "TLS configuration",
												Attributes: map[string]schema.Attribute{
													"ca": schema.StringAttribute{
														Description:         "Text of the CA cert to use for the targets.",
														MarkdownDescription: "Text of the CA cert to use for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ca_file": schema.StringAttribute{
														Description:         "The CA cert to use for the targets.",
														MarkdownDescription: "The CA cert to use for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ca_ref": schema.StringAttribute{
														Description:         "CARef is the name of the secret within the secret manager to use as the CA cert for the targets.",
														MarkdownDescription: "CARef is the name of the secret within the secret manager to use as the CA cert for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert": schema.StringAttribute{
														Description:         "Text of the client cert file for the targets.",
														MarkdownDescription: "Text of the client cert file for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_file": schema.StringAttribute{
														Description:         "The client cert file for the targets.",
														MarkdownDescription: "The client cert file for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_ref": schema.StringAttribute{
														Description:         "CertRef is the name of the secret within the secret manager to use as the client cert for the targets.",
														MarkdownDescription: "CertRef is the name of the secret within the secret manager to use as the client cert for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "Disable target certificate validation.",
														MarkdownDescription: "Disable target certificate validation.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "Text of the client key file for the targets.",
														MarkdownDescription: "Text of the client key file for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_file": schema.StringAttribute{
														Description:         "The client key file for the targets.",
														MarkdownDescription: "The client key file for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_ref": schema.StringAttribute{
														Description:         "KeyRef is the name of the secret within the secret manager to use as the client key for the targets.",
														MarkdownDescription: "KeyRef is the name of the secret within the secret manager to use as the client key for the targets.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_version": schema.Int64Attribute{
														Description:         "Maximum TLS version.",
														MarkdownDescription: "Maximum TLS version.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_version": schema.Int64Attribute{
														Description:         "Minimum TLS version.",
														MarkdownDescription: "Minimum TLS version.",
														Required:            false,
														Optional:            true,
														Computed:            false,
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

											"user": schema.StringAttribute{
												Description:         "Username",
												MarkdownDescription: "Username",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_timeout": schema.StringAttribute{
												Description:         "I/O write timeout",
												MarkdownDescription: "I/O write timeout",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
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

							"ephemeral_dashboard": schema.SingleNestedAttribute{
								Description:         "EphemeralDashboard contains the config about the ephemeral dashboard feature",
								MarkdownDescription: "EphemeralDashboard contains the config about the ephemeral dashboard feature",
								Attributes: map[string]schema.Attribute{
									"cleanup_interval": schema.StringAttribute{
										Description:         "The interval at which to trigger the cleanup of ephemeral dashboards, based on their TTLs.",
										MarkdownDescription: "The interval at which to trigger the cleanup of ephemeral dashboards, based on their TTLs.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
										},
									},

									"enable": schema.BoolAttribute{
										Description:         "When true user will be able to use the ephemeral dashboard at project level.",
										MarkdownDescription: "When true user will be able to use the ephemeral dashboard at project level.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"ephemeral_dashboards_cleanup_interval": schema.StringAttribute{
								Description:         "EphemeralDashboardsCleanupInterval is the interval at which the ephemeral dashboards are cleaned up DEPRECATED. Please use the config EphemeralDashboard instead.",
								MarkdownDescription: "EphemeralDashboardsCleanupInterval is the interval at which the ephemeral dashboards are cleaned up DEPRECATED. Please use the config EphemeralDashboard instead.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
								},
							},

							"frontend": schema.SingleNestedAttribute{
								Description:         "Frontend contains any config that will be used by the frontend itself.",
								MarkdownDescription: "Frontend contains any config that will be used by the frontend itself.",
								Attributes: map[string]schema.Attribute{
									"disable": schema.BoolAttribute{
										Description:         "When it is true, Perses won't serve the frontend anymore, and any other config set here will be ignored",
										MarkdownDescription: "When it is true, Perses won't serve the frontend anymore, and any other config set here will be ignored",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"explorer": schema.SingleNestedAttribute{
										Description:         "Explorer is activating the different kind of explorer supported. Be sure you have installed an associated plugin for each explorer type.",
										MarkdownDescription: "Explorer is activating the different kind of explorer supported. Be sure you have installed an associated plugin for each explorer type.",
										Attributes: map[string]schema.Attribute{
											"enable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"important_dashboards": schema.ListNestedAttribute{
										Description:         "ImportantDashboards contains important dashboard selectors",
										MarkdownDescription: "ImportantDashboards contains important dashboard selectors",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"dashboard": schema.StringAttribute{
													Description:         "Dashboard is the name of the dashboard (dashboard.metadata.name)",
													MarkdownDescription: "Dashboard is the name of the dashboard (dashboard.metadata.name)",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"project": schema.StringAttribute{
													Description:         "Project is the name of the project (dashboard.metadata.project)",
													MarkdownDescription: "Project is the name of the project (dashboard.metadata.project)",
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

									"information": schema.StringAttribute{
										Description:         "Information contains markdown content to be display on the home page",
										MarkdownDescription: "Information contains markdown content to be display on the home page",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"time_range": schema.SingleNestedAttribute{
										Description:         "TimeRange contains the time range configuration for the dropdown",
										MarkdownDescription: "TimeRange contains the time range configuration for the dropdown",
										Attributes: map[string]schema.Attribute{
											"disable_custom": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"options": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"global_datasource_discovery": schema.ListNestedAttribute{
								Description:         "GlobalDatasourceDiscovery is the configuration that helps to generate a list of global datasource based on the discovery chosen. Be careful: the data coming from the discovery will totally override what exists in the database. Note that this is an experimental feature. Behavior and config may change in the future.",
								MarkdownDescription: "GlobalDatasourceDiscovery is the configuration that helps to generate a list of global datasource based on the discovery chosen. Be careful: the data coming from the discovery will totally override what exists in the database. Note that this is an experimental feature. Behavior and config may change in the future.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"discovery_name": schema.StringAttribute{
											Description:         "The name of the discovery config. It is used for logging purposes only",
											MarkdownDescription: "The name of the discovery config. It is used for logging purposes only",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"http_sd": schema.SingleNestedAttribute{
											Description:         "HTTP-based service discovery provides a more generic way to generate a set of global datasource and serves as an interface to plug in custom service discovery mechanisms. It fetches an HTTP endpoint containing a list of zero or more global datasources. The target must reply with an HTTP 200 response. The HTTP header Content-Type must be application/json, and the body must be valid array of JSON.",
											MarkdownDescription: "HTTP-based service discovery provides a more generic way to generate a set of global datasource and serves as an interface to plug in custom service discovery mechanisms. It fetches an HTTP endpoint containing a list of zero or more global datasources. The target must reply with an HTTP 200 response. The HTTP header Content-Type must be application/json, and the body must be valid array of JSON.",
											Attributes: map[string]schema.Attribute{
												"authorization": schema.SingleNestedAttribute{
													Description:         "The HTTP authorization credentials for the targets.",
													MarkdownDescription: "The HTTP authorization credentials for the targets.",
													Attributes: map[string]schema.Attribute{
														"credentials": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"credentials_file": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
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

												"basic_auth": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"password": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"password_file": schema.StringAttribute{
															Description:         "PasswordFile is a path to a file that contains a password",
															MarkdownDescription: "PasswordFile is a path to a file that contains a password",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"username": schema.StringAttribute{
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

												"headers": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"native_auth": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"login": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"password": schema.StringAttribute{
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

												"oauth": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auth_style": schema.Int64Attribute{
															Description:         "AuthStyle optionally specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect.",
															MarkdownDescription: "AuthStyle optionally specifies how the endpoint wants the client ID & client secret sent. The zero value means to auto-detect.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"client_id": schema.StringAttribute{
															Description:         "ClientID is the application's ID.",
															MarkdownDescription: "ClientID is the application's ID.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"client_secret": schema.StringAttribute{
															Description:         "ClientSecret is the application's secret.",
															MarkdownDescription: "ClientSecret is the application's secret.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"client_secretfile": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"endpoint_params": schema.MapAttribute{
															Description:         "EndpointParams specifies additional parameters for requests to the token endpoint.",
															MarkdownDescription: "EndpointParams specifies additional parameters for requests to the token endpoint.",
															ElementType:         types.ListType{ElemType: types.StringType},
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"scopes": schema.ListAttribute{
															Description:         "Scope specifies optional requested permissions.",
															MarkdownDescription: "Scope specifies optional requested permissions.",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"token_url": schema.StringAttribute{
															Description:         "TokenURL is the resource server's token endpoint URL. This is a constant specific to each server.",
															MarkdownDescription: "TokenURL is the resource server's token endpoint URL. This is a constant specific to each server.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"tls_config": schema.SingleNestedAttribute{
													Description:         "TLSConfig to use to connect to the targets.",
													MarkdownDescription: "TLSConfig to use to connect to the targets.",
													Attributes: map[string]schema.Attribute{
														"ca": schema.StringAttribute{
															Description:         "Text of the CA cert to use for the targets.",
															MarkdownDescription: "Text of the CA cert to use for the targets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ca_file": schema.StringAttribute{
															Description:         "The CA cert to use for the targets.",
															MarkdownDescription: "The CA cert to use for the targets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert": schema.StringAttribute{
															Description:         "Text of the client cert file for the targets.",
															MarkdownDescription: "Text of the client cert file for the targets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"cert_file": schema.StringAttribute{
															Description:         "The client cert file for the targets.",
															MarkdownDescription: "The client cert file for the targets.",
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

														"key": schema.StringAttribute{
															Description:         "Text of the client key file for the targets.",
															MarkdownDescription: "Text of the client key file for the targets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key_file": schema.StringAttribute{
															Description:         "The client key file for the targets.",
															MarkdownDescription: "The client key file for the targets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_version": schema.StringAttribute{
															Description:         "Maximum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default maximum version, which is TLS 1.3. See MaxVersion in https://pkg.go.dev/crypto/tls#Config.",
															MarkdownDescription: "Maximum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default maximum version, which is TLS 1.3. See MaxVersion in https://pkg.go.dev/crypto/tls#Config.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"min_version": schema.StringAttribute{
															Description:         "Minimum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default minimum version, which is TLS 1.2. See MinVersion in https://pkg.go.dev/crypto/tls#Config.",
															MarkdownDescription: "Minimum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default minimum version, which is TLS 1.2. See MinVersion in https://pkg.go.dev/crypto/tls#Config.",
															Required:            false,
															Optional:            true,
															Computed:            false,
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

												"url": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"kubernetes_sd": schema.SingleNestedAttribute{
											Description:         "Kubernetes SD configurations allow retrieving global datasource from Kubernetes' REST API and always staying synchronized with the cluster state.",
											MarkdownDescription: "Kubernetes SD configurations allow retrieving global datasource from Kubernetes' REST API and always staying synchronized with the cluster state.",
											Attributes: map[string]schema.Attribute{
												"datasource_plugin_kind": schema.StringAttribute{
													Description:         "DatasourcePluginKind is the name of the datasource plugin that should be filled when creating datasources found.",
													MarkdownDescription: "DatasourcePluginKind is the name of the datasource plugin that should be filled when creating datasources found.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"labels": schema.MapAttribute{
													Description:         "The labels used to filter the list of resource when contacting the Kubernetes API.",
													MarkdownDescription: "The labels used to filter the list of resource when contacting the Kubernetes API.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Kubernetes namespace to constraint the query to only one namespace. Leave empty if you are looking for datasource cross-namespace.",
													MarkdownDescription: "Kubernetes namespace to constraint the query to only one namespace. Leave empty if you are looking for datasource cross-namespace.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"pod_configuration": schema.SingleNestedAttribute{
													Description:         "Configuration when you want to discover the pods in Kubernetes",
													MarkdownDescription: "Configuration when you want to discover the pods in Kubernetes",
													Attributes: map[string]schema.Attribute{
														"container_name": schema.StringAttribute{
															Description:         "Name of the container the target address points to.",
															MarkdownDescription: "Name of the container the target address points to.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"container_port_name": schema.StringAttribute{
															Description:         "Name of the container port.",
															MarkdownDescription: "Name of the container port.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"container_port_number": schema.Int64Attribute{
															Description:         "Number of the container port.",
															MarkdownDescription: "Number of the container port.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"enable": schema.BoolAttribute{
															Description:         "If set to true, Perses server will discovery the pod",
															MarkdownDescription: "If set to true, Perses server will discovery the pod",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"service_configuration": schema.SingleNestedAttribute{
													Description:         "Configuration when you want to discover the services in Kubernetes",
													MarkdownDescription: "Configuration when you want to discover the services in Kubernetes",
													Attributes: map[string]schema.Attribute{
														"enable": schema.BoolAttribute{
															Description:         "If set to true, Perses server will discovery the service",
															MarkdownDescription: "If set to true, Perses server will discovery the service",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port_name": schema.StringAttribute{
															Description:         "Name of the service port for the target.",
															MarkdownDescription: "Name of the service port for the target.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port_number": schema.Int64Attribute{
															Description:         "Number of the service port for the target.",
															MarkdownDescription: "Number of the service port for the target.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service_type": schema.StringAttribute{
															Description:         "The type of the service.",
															MarkdownDescription: "The type of the service.",
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

										"refresh_interval": schema.StringAttribute{
											Description:         "Refresh interval to re-query the endpoint.",
											MarkdownDescription: "Refresh interval to re-query the endpoint.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"provisioning": schema.SingleNestedAttribute{
								Description:         "Provisioning contains the provisioning config that can be used if you want to provide default resources.",
								MarkdownDescription: "Provisioning contains the provisioning config that can be used if you want to provide default resources.",
								Attributes: map[string]schema.Attribute{
									"folders": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"interval": schema.StringAttribute{
										Description:         "Interval is the refresh frequency",
										MarkdownDescription: "Interval is the refresh frequency",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"schemas": schema.SingleNestedAttribute{
								Description:         "Schemas contain the configuration to get access to the CUE schemas",
								MarkdownDescription: "Schemas contain the configuration to get access to the CUE schemas",
								Attributes: map[string]schema.Attribute{
									"datasources_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"interval": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
										},
									},

									"panels_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"queries_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"variables_path": schema.StringAttribute{
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

							"security": schema.SingleNestedAttribute{
								Description:         "Security contains any configuration that changes the API behavior like the endpoints exposed or if the permissions are activated.",
								MarkdownDescription: "Security contains any configuration that changes the API behavior like the endpoints exposed or if the permissions are activated.",
								Attributes: map[string]schema.Attribute{
									"authentication": schema.SingleNestedAttribute{
										Description:         "Authentication contains configuration regarding management of access/refresh token",
										MarkdownDescription: "Authentication contains configuration regarding management of access/refresh token",
										Attributes: map[string]schema.Attribute{
											"access_token_ttl": schema.StringAttribute{
												Description:         "AccessTokenTTL is the time to live of the access token. By default, it is 15 minutes.",
												MarkdownDescription: "AccessTokenTTL is the time to live of the access token. By default, it is 15 minutes.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
												},
											},

											"disable_sign_up": schema.BoolAttribute{
												Description:         "DisableSignUp deactivates the Sign-up page in the UI. It also disables the endpoint that gives the possibility to create a user.",
												MarkdownDescription: "DisableSignUp deactivates the Sign-up page in the UI. It also disables the endpoint that gives the possibility to create a user.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"providers": schema.SingleNestedAttribute{
												Description:         "Providers configure the different authentication providers",
												MarkdownDescription: "Providers configure the different authentication providers",
												Attributes: map[string]schema.Attribute{
													"enable_native": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"oauth": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"auth_url": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"client_credentials": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"client_id": schema.StringAttribute{
																			Description:         "Hidden special type for storing secrets.",
																			MarkdownDescription: "Hidden special type for storing secrets.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"client_secret": schema.StringAttribute{
																			Description:         "Hidden special type for storing secrets.",
																			MarkdownDescription: "Hidden special type for storing secrets.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scopes": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"client_id": schema.StringAttribute{
																	Description:         "Hidden special type for storing secrets.",
																	MarkdownDescription: "Hidden special type for storing secrets.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"client_secret": schema.StringAttribute{
																	Description:         "Hidden special type for storing secrets.",
																	MarkdownDescription: "Hidden special type for storing secrets.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"custom_login_property": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"device_auth_url": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"device_code": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"client_id": schema.StringAttribute{
																			Description:         "Hidden special type for storing secrets.",
																			MarkdownDescription: "Hidden special type for storing secrets.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"client_secret": schema.StringAttribute{
																			Description:         "Hidden special type for storing secrets.",
																			MarkdownDescription: "Hidden special type for storing secrets.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scopes": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"http": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"timeout": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
																			},
																		},

																		"tls_config": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"ca": schema.StringAttribute{
																					Description:         "Text of the CA cert to use for the targets.",
																					MarkdownDescription: "Text of the CA cert to use for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"ca_file": schema.StringAttribute{
																					Description:         "The CA cert to use for the targets.",
																					MarkdownDescription: "The CA cert to use for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"cert": schema.StringAttribute{
																					Description:         "Text of the client cert file for the targets.",
																					MarkdownDescription: "Text of the client cert file for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"cert_file": schema.StringAttribute{
																					Description:         "The client cert file for the targets.",
																					MarkdownDescription: "The client cert file for the targets.",
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

																				"key": schema.StringAttribute{
																					Description:         "Text of the client key file for the targets.",
																					MarkdownDescription: "Text of the client key file for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"key_file": schema.StringAttribute{
																					Description:         "The client key file for the targets.",
																					MarkdownDescription: "The client key file for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"max_version": schema.StringAttribute{
																					Description:         "Maximum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default maximum version, which is TLS 1.3. See MaxVersion in https://pkg.go.dev/crypto/tls#Config.",
																					MarkdownDescription: "Maximum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default maximum version, which is TLS 1.3. See MaxVersion in https://pkg.go.dev/crypto/tls#Config.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"min_version": schema.StringAttribute{
																					Description:         "Minimum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default minimum version, which is TLS 1.2. See MinVersion in https://pkg.go.dev/crypto/tls#Config.",
																					MarkdownDescription: "Minimum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default minimum version, which is TLS 1.2. See MinVersion in https://pkg.go.dev/crypto/tls#Config.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"server_name": schema.StringAttribute{
																					Description:         "Used to verify the hostname for the targets.",
																					MarkdownDescription: "Used to verify the hostname for the targets.",
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
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"redirect_uri": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"scopes": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"slug_id": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"token_url": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"user_infos_url": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
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

													"oidc": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"client_credentials": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"client_id": schema.StringAttribute{
																			Description:         "Hidden special type for storing secrets.",
																			MarkdownDescription: "Hidden special type for storing secrets.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"client_secret": schema.StringAttribute{
																			Description:         "Hidden special type for storing secrets.",
																			MarkdownDescription: "Hidden special type for storing secrets.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scopes": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"client_id": schema.StringAttribute{
																	Description:         "Hidden special type for storing secrets.",
																	MarkdownDescription: "Hidden special type for storing secrets.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"client_secret": schema.StringAttribute{
																	Description:         "Hidden special type for storing secrets.",
																	MarkdownDescription: "Hidden special type for storing secrets.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"device_code": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"client_id": schema.StringAttribute{
																			Description:         "Hidden special type for storing secrets.",
																			MarkdownDescription: "Hidden special type for storing secrets.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"client_secret": schema.StringAttribute{
																			Description:         "Hidden special type for storing secrets.",
																			MarkdownDescription: "Hidden special type for storing secrets.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scopes": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"disable_pkce": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"discovery_url": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"timeout": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
																			},
																		},

																		"tls_config": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"ca": schema.StringAttribute{
																					Description:         "Text of the CA cert to use for the targets.",
																					MarkdownDescription: "Text of the CA cert to use for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"ca_file": schema.StringAttribute{
																					Description:         "The CA cert to use for the targets.",
																					MarkdownDescription: "The CA cert to use for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"cert": schema.StringAttribute{
																					Description:         "Text of the client cert file for the targets.",
																					MarkdownDescription: "Text of the client cert file for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"cert_file": schema.StringAttribute{
																					Description:         "The client cert file for the targets.",
																					MarkdownDescription: "The client cert file for the targets.",
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

																				"key": schema.StringAttribute{
																					Description:         "Text of the client key file for the targets.",
																					MarkdownDescription: "Text of the client key file for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"key_file": schema.StringAttribute{
																					Description:         "The client key file for the targets.",
																					MarkdownDescription: "The client key file for the targets.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"max_version": schema.StringAttribute{
																					Description:         "Maximum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default maximum version, which is TLS 1.3. See MaxVersion in https://pkg.go.dev/crypto/tls#Config.",
																					MarkdownDescription: "Maximum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default maximum version, which is TLS 1.3. See MaxVersion in https://pkg.go.dev/crypto/tls#Config.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"min_version": schema.StringAttribute{
																					Description:         "Minimum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default minimum version, which is TLS 1.2. See MinVersion in https://pkg.go.dev/crypto/tls#Config.",
																					MarkdownDescription: "Minimum acceptable TLS version. Accepted values: TLS10 (TLS 1.0), TLS11 (TLS 1.1), TLS12 (TLS 1.2), TLS13 (TLS 1.3). If unset, Perses will use Go default minimum version, which is TLS 1.2. See MinVersion in https://pkg.go.dev/crypto/tls#Config.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"server_name": schema.StringAttribute{
																					Description:         "Used to verify the hostname for the targets.",
																					MarkdownDescription: "Used to verify the hostname for the targets.",
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
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"issuer": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"redirect_uri": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"scopes": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"slug_id": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"url_params": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
												Required: true,
												Optional: false,
												Computed: false,
											},

											"refresh_token_ttl": schema.StringAttribute{
												Description:         "RefreshTokenTTL is the time to live of the refresh token. The refresh token is used to get a new access token when it is expired. By default, it is 24 hours.",
												MarkdownDescription: "RefreshTokenTTL is the time to live of the refresh token. The refresh token is used to get a new access token when it is expired. By default, it is 24 hours.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorization": schema.SingleNestedAttribute{
										Description:         "Authorization contains all configs around rbac (permissions and roles)",
										MarkdownDescription: "Authorization contains all configs around rbac (permissions and roles)",
										Attributes: map[string]schema.Attribute{
											"check_latest_update_interval": schema.StringAttribute{
												Description:         "CheckLatestUpdateInterval that checks if the RBAC cache needs to be refreshed with db content. Only for SQL database setup.",
												MarkdownDescription: "CheckLatestUpdateInterval that checks if the RBAC cache needs to be refreshed with db content. Only for SQL database setup.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?$`), ""),
												},
											},

											"guest_permissions": schema.ListNestedAttribute{
												Description:         "Default permissions for guest users (logged-in users)",
												MarkdownDescription: "Default permissions for guest users (logged-in users)",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"actions": schema.ListAttribute{
															Description:         "Actions of the permission (read, create, update, delete, ...)",
															MarkdownDescription: "Actions of the permission (read, create, update, delete, ...)",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"scopes": schema.ListAttribute{
															Description:         "The list of kind targeted by the permission. For example: 'Datasource', 'Dashboard', ... With Role, you can't target global kinds",
															MarkdownDescription: "The list of kind targeted by the permission. For example: 'Datasource', 'Dashboard', ... With Role, you can't target global kinds",
															ElementType:         types.StringType,
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cookie": schema.SingleNestedAttribute{
										Description:         "Cookie configuration",
										MarkdownDescription: "Cookie configuration",
										Attributes: map[string]schema.Attribute{
											"same_site": schema.Int64Attribute{
												Description:         "Set the SameSite cookie attribute and prevents the browser from sending the cookie along with cross-site requests. The main goal is to mitigate the risk of cross-origin information leakage. This setting also provides some protection against cross-site request forgery attacks (CSRF)",
												MarkdownDescription: "Set the SameSite cookie attribute and prevents the browser from sending the cookie along with cross-site requests. The main goal is to mitigate the risk of cross-origin information leakage. This setting also provides some protection against cross-site request forgery attacks (CSRF)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secure": schema.BoolAttribute{
												Description:         "Set to true if you host Perses behind HTTPS. Default is false",
												MarkdownDescription: "Set to true if you host Perses behind HTTPS. Default is false",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"enable_auth": schema.BoolAttribute{
										Description:         "When it is true, the authentication and authorization config are considered. And you will need a valid JWT token to contact most of the endpoints exposed by the API",
										MarkdownDescription: "When it is true, the authentication and authorization config are considered. And you will need a valid JWT token to contact most of the endpoints exposed by the API",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"encryption_key": schema.StringAttribute{
										Description:         "EncryptionKey is the secret key used to encrypt and decrypt sensitive data stored in the database such as the password of the basic auth for a datasource. Note that if it is not provided, it will use a default value. On a production instance, you should set this key. Also note the key size must be exactly 32 bytes long as we are using AES-256 to encrypt the data.",
										MarkdownDescription: "EncryptionKey is the secret key used to encrypt and decrypt sensitive data stored in the database such as the password of the basic auth for a datasource. Note that if it is not provided, it will use a default value. On a production instance, you should set this key. Also note the key size must be exactly 32 bytes long as we are using AES-256 to encrypt the data.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"encryption_key_file": schema.StringAttribute{
										Description:         "EncryptionKeyFile is the path to file containing the secret key",
										MarkdownDescription: "EncryptionKeyFile is the path to file containing the secret key",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"readonly": schema.BoolAttribute{
										Description:         "Readonly will deactivate any HTTP POST, PUT, DELETE endpoint",
										MarkdownDescription: "Readonly will deactivate any HTTP POST, PUT, DELETE endpoint",
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

					"container_port": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Image specifies the container image that should be used for the Perses deployment.",
						MarkdownDescription: "Image specifies the container image that should be used for the Perses deployment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"liveness_probe": schema.SingleNestedAttribute{
						Description:         "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
						MarkdownDescription: "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
						Attributes: map[string]schema.Attribute{
							"exec": schema.SingleNestedAttribute{
								Description:         "Exec specifies a command to execute in the container.",
								MarkdownDescription: "Exec specifies a command to execute in the container.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
										Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
										MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

							"failure_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"grpc": schema.SingleNestedAttribute{
								Description:         "GRPC specifies a GRPC HealthCheckRequest.",
								MarkdownDescription: "GRPC specifies a GRPC HealthCheckRequest.",
								Attributes: map[string]schema.Attribute{
									"port": schema.Int64Attribute{
										Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
										MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"service": schema.StringAttribute{
										Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
										MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_get": schema.SingleNestedAttribute{
								Description:         "HTTPGet specifies an HTTP GET request to perform.",
								MarkdownDescription: "HTTPGet specifies an HTTP GET request to perform.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
										MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"http_headers": schema.ListNestedAttribute{
										Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
										MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
													MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "The header field value",
													MarkdownDescription: "The header field value",
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

									"path": schema.StringAttribute{
										Description:         "Path to access on the HTTP server.",
										MarkdownDescription: "Path to access on the HTTP server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.StringAttribute{
										Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
										MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_delay_seconds": schema.Int64Attribute{
								Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"period_seconds": schema.Int64Attribute{
								Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
								MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"success_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tcp_socket": schema.SingleNestedAttribute{
								Description:         "TCPSocket specifies a connection to a TCP port.",
								MarkdownDescription: "TCPSocket specifies a connection to a TCP port.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Optional: Host name to connect to, defaults to the pod IP.",
										MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.StringAttribute{
										Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
								MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout_seconds": schema.Int64Attribute{
								Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"metadata": schema.SingleNestedAttribute{
						Description:         "Metadata to add to deployed pods",
						MarkdownDescription: "Metadata to add to deployed pods",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"node_selector": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"readiness_probe": schema.SingleNestedAttribute{
						Description:         "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
						MarkdownDescription: "Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.",
						Attributes: map[string]schema.Attribute{
							"exec": schema.SingleNestedAttribute{
								Description:         "Exec specifies a command to execute in the container.",
								MarkdownDescription: "Exec specifies a command to execute in the container.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
										Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
										MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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

							"failure_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"grpc": schema.SingleNestedAttribute{
								Description:         "GRPC specifies a GRPC HealthCheckRequest.",
								MarkdownDescription: "GRPC specifies a GRPC HealthCheckRequest.",
								Attributes: map[string]schema.Attribute{
									"port": schema.Int64Attribute{
										Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
										MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"service": schema.StringAttribute{
										Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
										MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"http_get": schema.SingleNestedAttribute{
								Description:         "HTTPGet specifies an HTTP GET request to perform.",
								MarkdownDescription: "HTTPGet specifies an HTTP GET request to perform.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
										MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"http_headers": schema.ListNestedAttribute{
										Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
										MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
													MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "The header field value",
													MarkdownDescription: "The header field value",
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

									"path": schema.StringAttribute{
										Description:         "Path to access on the HTTP server.",
										MarkdownDescription: "Path to access on the HTTP server.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.StringAttribute{
										Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
										MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_delay_seconds": schema.Int64Attribute{
								Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"period_seconds": schema.Int64Attribute{
								Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
								MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"success_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tcp_socket": schema.SingleNestedAttribute{
								Description:         "TCPSocket specifies a connection to a TCP port.",
								MarkdownDescription: "TCPSocket specifies a connection to a TCP port.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
										Description:         "Optional: Host name to connect to, defaults to the pod IP.",
										MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.StringAttribute{
										Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
								MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout_seconds": schema.Int64Attribute{
								Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "service specifies the service configuration for the perses instance",
						MarkdownDescription: "service specifies the service configuration for the perses instance",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
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

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage configuration used by the StatefulSet",
						MarkdownDescription: "Storage configuration used by the StatefulSet",
						Attributes: map[string]schema.Attribute{
							"size": schema.StringAttribute{
								Description:         "Size of the storage. cannot be decreased.",
								MarkdownDescription: "Size of the storage. cannot be decreased.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class": schema.StringAttribute{
								Description:         "StorageClass to use for PVCs. If not specified, will use the default storage class",
								MarkdownDescription: "StorageClass to use for PVCs. If not specified, will use the default storage class",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "tls specifies the tls configuration for the perses instance",
						MarkdownDescription: "tls specifies the tls configuration for the perses instance",
						Attributes: map[string]schema.Attribute{
							"ca_cert": schema.SingleNestedAttribute{
								Description:         "CaCert to verify the perses certificate",
								MarkdownDescription: "CaCert to verify the perses certificate",
								Attributes: map[string]schema.Attribute{
									"cert_path": schema.StringAttribute{
										Description:         "Path to Certificate",
										MarkdownDescription: "Path to Certificate",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"private_key_path": schema.StringAttribute{
										Description:         "Path to Private key certificate",
										MarkdownDescription: "Path to Private key certificate",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type source type of secret",
										MarkdownDescription: "Type source type of secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("secret", "configmap", "file"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable": schema.BoolAttribute{
								Description:         "Enable TLS connection to perses",
								MarkdownDescription: "Enable TLS connection to perses",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"insecure_skip_verify": schema.BoolAttribute{
								Description:         "InsecureSkipVerify skip verify of perses certificate",
								MarkdownDescription: "InsecureSkipVerify skip verify of perses certificate",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user_cert": schema.SingleNestedAttribute{
								Description:         "UserCert client cert/key for mTLS",
								MarkdownDescription: "UserCert client cert/key for mTLS",
								Attributes: map[string]schema.Attribute{
									"cert_path": schema.StringAttribute{
										Description:         "Path to Certificate",
										MarkdownDescription: "Path to Certificate",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name of basic auth k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Name of basic auth k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace": schema.StringAttribute{
										Description:         "Namsespace of certificate k8s resource (when type is secret or configmap)",
										MarkdownDescription: "Namsespace of certificate k8s resource (when type is secret or configmap)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"private_key_path": schema.StringAttribute{
										Description:         "Path to Private key certificate",
										MarkdownDescription: "Path to Private key certificate",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type source type of secret",
										MarkdownDescription: "Type source type of secret",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("secret", "configmap", "file"),
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

					"tolerations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *PersesDevPersesV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_perses_dev_perses_v1alpha1_manifest")

	var model PersesDevPersesV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("perses.dev/v1alpha1")
	model.Kind = pointer.String("Perses")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
