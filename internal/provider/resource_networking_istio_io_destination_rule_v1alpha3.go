/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type NetworkingIstioIoDestinationRuleV1Alpha3Resource struct{}

var (
	_ resource.Resource = (*NetworkingIstioIoDestinationRuleV1Alpha3Resource)(nil)
)

type NetworkingIstioIoDestinationRuleV1Alpha3TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type NetworkingIstioIoDestinationRuleV1Alpha3GoModel struct {
	Id         *int64  `tfsdk:"id" yaml:",omitempty"`
	YAML       *string `tfsdk:"yaml" yaml:",omitempty"`
	ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion"`
	Kind       *string `tfsdk:"kind" yaml:"kind"`

	Metadata struct {
		Name string `tfsdk:"name" yaml:"name"`

		Namespace *string `tfsdk:"namespace" yaml:"namespace"`

		Labels      map[string]string `tfsdk:"labels" yaml:",omitempty"`
		Annotations map[string]string `tfsdk:"annotations" yaml:",omitempty"`
	} `tfsdk:"metadata" yaml:"metadata"`

	Spec *struct {
		ExportTo *[]string `tfsdk:"export_to" yaml:"exportTo,omitempty"`

		Host *string `tfsdk:"host" yaml:"host,omitempty"`

		Subsets *[]struct {
			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			TrafficPolicy *struct {
				ConnectionPool *struct {
					Http *struct {
						H2UpgradePolicy *string `tfsdk:"h2_upgrade_policy" yaml:"h2UpgradePolicy,omitempty"`

						Http1MaxPendingRequests *int64 `tfsdk:"http1_max_pending_requests" yaml:"http1MaxPendingRequests,omitempty"`

						Http2MaxRequests *int64 `tfsdk:"http2_max_requests" yaml:"http2MaxRequests,omitempty"`

						IdleTimeout *string `tfsdk:"idle_timeout" yaml:"idleTimeout,omitempty"`

						MaxRequestsPerConnection *int64 `tfsdk:"max_requests_per_connection" yaml:"maxRequestsPerConnection,omitempty"`

						MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`

						UseClientProtocol *bool `tfsdk:"use_client_protocol" yaml:"useClientProtocol,omitempty"`
					} `tfsdk:"http" yaml:"http,omitempty"`

					Tcp *struct {
						ConnectTimeout *string `tfsdk:"connect_timeout" yaml:"connectTimeout,omitempty"`

						MaxConnectionDuration *string `tfsdk:"max_connection_duration" yaml:"maxConnectionDuration,omitempty"`

						MaxConnections *int64 `tfsdk:"max_connections" yaml:"maxConnections,omitempty"`

						TcpKeepalive *struct {
							Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

							Probes *int64 `tfsdk:"probes" yaml:"probes,omitempty"`

							Time *string `tfsdk:"time" yaml:"time,omitempty"`
						} `tfsdk:"tcp_keepalive" yaml:"tcpKeepalive,omitempty"`
					} `tfsdk:"tcp" yaml:"tcp,omitempty"`
				} `tfsdk:"connection_pool" yaml:"connectionPool,omitempty"`

				LoadBalancer *struct {
					ConsistentHash *struct {
						HttpCookie *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
						} `tfsdk:"http_cookie" yaml:"httpCookie,omitempty"`

						HttpHeaderName *string `tfsdk:"http_header_name" yaml:"httpHeaderName,omitempty"`

						HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" yaml:"httpQueryParameterName,omitempty"`

						Maglev *struct {
							TableSize *int64 `tfsdk:"table_size" yaml:"tableSize,omitempty"`
						} `tfsdk:"maglev" yaml:"maglev,omitempty"`

						MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`

						RingHash *struct {
							MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`
						} `tfsdk:"ring_hash" yaml:"ringHash,omitempty"`

						UseSourceIp *bool `tfsdk:"use_source_ip" yaml:"useSourceIp,omitempty"`
					} `tfsdk:"consistent_hash" yaml:"consistentHash,omitempty"`

					LocalityLbSetting *struct {
						Distribute *[]struct {
							From *string `tfsdk:"from" yaml:"from,omitempty"`

							To *map[string]string `tfsdk:"to" yaml:"to,omitempty"`
						} `tfsdk:"distribute" yaml:"distribute,omitempty"`

						Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

						Failover *[]struct {
							From *string `tfsdk:"from" yaml:"from,omitempty"`

							To *string `tfsdk:"to" yaml:"to,omitempty"`
						} `tfsdk:"failover" yaml:"failover,omitempty"`

						FailoverPriority *[]string `tfsdk:"failover_priority" yaml:"failoverPriority,omitempty"`
					} `tfsdk:"locality_lb_setting" yaml:"localityLbSetting,omitempty"`

					Simple *string `tfsdk:"simple" yaml:"simple,omitempty"`

					WarmupDurationSecs *string `tfsdk:"warmup_duration_secs" yaml:"warmupDurationSecs,omitempty"`
				} `tfsdk:"load_balancer" yaml:"loadBalancer,omitempty"`

				OutlierDetection *struct {
					BaseEjectionTime *string `tfsdk:"base_ejection_time" yaml:"baseEjectionTime,omitempty"`

					Consecutive5xxErrors *int64 `tfsdk:"consecutive5xx_errors" yaml:"consecutive5xxErrors,omitempty"`

					ConsecutiveErrors *int64 `tfsdk:"consecutive_errors" yaml:"consecutiveErrors,omitempty"`

					ConsecutiveGatewayErrors *int64 `tfsdk:"consecutive_gateway_errors" yaml:"consecutiveGatewayErrors,omitempty"`

					ConsecutiveLocalOriginFailures *int64 `tfsdk:"consecutive_local_origin_failures" yaml:"consecutiveLocalOriginFailures,omitempty"`

					Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

					MaxEjectionPercent *int64 `tfsdk:"max_ejection_percent" yaml:"maxEjectionPercent,omitempty"`

					MinHealthPercent *int64 `tfsdk:"min_health_percent" yaml:"minHealthPercent,omitempty"`

					SplitExternalLocalOriginErrors *bool `tfsdk:"split_external_local_origin_errors" yaml:"splitExternalLocalOriginErrors,omitempty"`
				} `tfsdk:"outlier_detection" yaml:"outlierDetection,omitempty"`

				PortLevelSettings *[]struct {
					ConnectionPool *struct {
						Http *struct {
							H2UpgradePolicy *string `tfsdk:"h2_upgrade_policy" yaml:"h2UpgradePolicy,omitempty"`

							Http1MaxPendingRequests *int64 `tfsdk:"http1_max_pending_requests" yaml:"http1MaxPendingRequests,omitempty"`

							Http2MaxRequests *int64 `tfsdk:"http2_max_requests" yaml:"http2MaxRequests,omitempty"`

							IdleTimeout *string `tfsdk:"idle_timeout" yaml:"idleTimeout,omitempty"`

							MaxRequestsPerConnection *int64 `tfsdk:"max_requests_per_connection" yaml:"maxRequestsPerConnection,omitempty"`

							MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`

							UseClientProtocol *bool `tfsdk:"use_client_protocol" yaml:"useClientProtocol,omitempty"`
						} `tfsdk:"http" yaml:"http,omitempty"`

						Tcp *struct {
							ConnectTimeout *string `tfsdk:"connect_timeout" yaml:"connectTimeout,omitempty"`

							MaxConnectionDuration *string `tfsdk:"max_connection_duration" yaml:"maxConnectionDuration,omitempty"`

							MaxConnections *int64 `tfsdk:"max_connections" yaml:"maxConnections,omitempty"`

							TcpKeepalive *struct {
								Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

								Probes *int64 `tfsdk:"probes" yaml:"probes,omitempty"`

								Time *string `tfsdk:"time" yaml:"time,omitempty"`
							} `tfsdk:"tcp_keepalive" yaml:"tcpKeepalive,omitempty"`
						} `tfsdk:"tcp" yaml:"tcp,omitempty"`
					} `tfsdk:"connection_pool" yaml:"connectionPool,omitempty"`

					LoadBalancer *struct {
						ConsistentHash *struct {
							HttpCookie *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
							} `tfsdk:"http_cookie" yaml:"httpCookie,omitempty"`

							HttpHeaderName *string `tfsdk:"http_header_name" yaml:"httpHeaderName,omitempty"`

							HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" yaml:"httpQueryParameterName,omitempty"`

							Maglev *struct {
								TableSize *int64 `tfsdk:"table_size" yaml:"tableSize,omitempty"`
							} `tfsdk:"maglev" yaml:"maglev,omitempty"`

							MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`

							RingHash *struct {
								MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`
							} `tfsdk:"ring_hash" yaml:"ringHash,omitempty"`

							UseSourceIp *bool `tfsdk:"use_source_ip" yaml:"useSourceIp,omitempty"`
						} `tfsdk:"consistent_hash" yaml:"consistentHash,omitempty"`

						LocalityLbSetting *struct {
							Distribute *[]struct {
								From *string `tfsdk:"from" yaml:"from,omitempty"`

								To *map[string]string `tfsdk:"to" yaml:"to,omitempty"`
							} `tfsdk:"distribute" yaml:"distribute,omitempty"`

							Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

							Failover *[]struct {
								From *string `tfsdk:"from" yaml:"from,omitempty"`

								To *string `tfsdk:"to" yaml:"to,omitempty"`
							} `tfsdk:"failover" yaml:"failover,omitempty"`

							FailoverPriority *[]string `tfsdk:"failover_priority" yaml:"failoverPriority,omitempty"`
						} `tfsdk:"locality_lb_setting" yaml:"localityLbSetting,omitempty"`

						Simple *string `tfsdk:"simple" yaml:"simple,omitempty"`

						WarmupDurationSecs *string `tfsdk:"warmup_duration_secs" yaml:"warmupDurationSecs,omitempty"`
					} `tfsdk:"load_balancer" yaml:"loadBalancer,omitempty"`

					OutlierDetection *struct {
						BaseEjectionTime *string `tfsdk:"base_ejection_time" yaml:"baseEjectionTime,omitempty"`

						Consecutive5xxErrors *int64 `tfsdk:"consecutive5xx_errors" yaml:"consecutive5xxErrors,omitempty"`

						ConsecutiveErrors *int64 `tfsdk:"consecutive_errors" yaml:"consecutiveErrors,omitempty"`

						ConsecutiveGatewayErrors *int64 `tfsdk:"consecutive_gateway_errors" yaml:"consecutiveGatewayErrors,omitempty"`

						ConsecutiveLocalOriginFailures *int64 `tfsdk:"consecutive_local_origin_failures" yaml:"consecutiveLocalOriginFailures,omitempty"`

						Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

						MaxEjectionPercent *int64 `tfsdk:"max_ejection_percent" yaml:"maxEjectionPercent,omitempty"`

						MinHealthPercent *int64 `tfsdk:"min_health_percent" yaml:"minHealthPercent,omitempty"`

						SplitExternalLocalOriginErrors *bool `tfsdk:"split_external_local_origin_errors" yaml:"splitExternalLocalOriginErrors,omitempty"`
					} `tfsdk:"outlier_detection" yaml:"outlierDetection,omitempty"`

					Port *struct {
						Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
					} `tfsdk:"port" yaml:"port,omitempty"`

					Tls *struct {
						CaCertificates *string `tfsdk:"ca_certificates" yaml:"caCertificates,omitempty"`

						ClientCertificate *string `tfsdk:"client_certificate" yaml:"clientCertificate,omitempty"`

						CredentialName *string `tfsdk:"credential_name" yaml:"credentialName,omitempty"`

						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

						PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

						Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

						SubjectAltNames *[]string `tfsdk:"subject_alt_names" yaml:"subjectAltNames,omitempty"`
					} `tfsdk:"tls" yaml:"tls,omitempty"`
				} `tfsdk:"port_level_settings" yaml:"portLevelSettings,omitempty"`

				Tls *struct {
					CaCertificates *string `tfsdk:"ca_certificates" yaml:"caCertificates,omitempty"`

					ClientCertificate *string `tfsdk:"client_certificate" yaml:"clientCertificate,omitempty"`

					CredentialName *string `tfsdk:"credential_name" yaml:"credentialName,omitempty"`

					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

					SubjectAltNames *[]string `tfsdk:"subject_alt_names" yaml:"subjectAltNames,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`

				Tunnel *struct {
					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

					TargetHost *string `tfsdk:"target_host" yaml:"targetHost,omitempty"`

					TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
				} `tfsdk:"tunnel" yaml:"tunnel,omitempty"`
			} `tfsdk:"traffic_policy" yaml:"trafficPolicy,omitempty"`
		} `tfsdk:"subsets" yaml:"subsets,omitempty"`

		TrafficPolicy *struct {
			ConnectionPool *struct {
				Http *struct {
					H2UpgradePolicy *string `tfsdk:"h2_upgrade_policy" yaml:"h2UpgradePolicy,omitempty"`

					Http1MaxPendingRequests *int64 `tfsdk:"http1_max_pending_requests" yaml:"http1MaxPendingRequests,omitempty"`

					Http2MaxRequests *int64 `tfsdk:"http2_max_requests" yaml:"http2MaxRequests,omitempty"`

					IdleTimeout *string `tfsdk:"idle_timeout" yaml:"idleTimeout,omitempty"`

					MaxRequestsPerConnection *int64 `tfsdk:"max_requests_per_connection" yaml:"maxRequestsPerConnection,omitempty"`

					MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`

					UseClientProtocol *bool `tfsdk:"use_client_protocol" yaml:"useClientProtocol,omitempty"`
				} `tfsdk:"http" yaml:"http,omitempty"`

				Tcp *struct {
					ConnectTimeout *string `tfsdk:"connect_timeout" yaml:"connectTimeout,omitempty"`

					MaxConnectionDuration *string `tfsdk:"max_connection_duration" yaml:"maxConnectionDuration,omitempty"`

					MaxConnections *int64 `tfsdk:"max_connections" yaml:"maxConnections,omitempty"`

					TcpKeepalive *struct {
						Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

						Probes *int64 `tfsdk:"probes" yaml:"probes,omitempty"`

						Time *string `tfsdk:"time" yaml:"time,omitempty"`
					} `tfsdk:"tcp_keepalive" yaml:"tcpKeepalive,omitempty"`
				} `tfsdk:"tcp" yaml:"tcp,omitempty"`
			} `tfsdk:"connection_pool" yaml:"connectionPool,omitempty"`

			LoadBalancer *struct {
				ConsistentHash *struct {
					HttpCookie *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
					} `tfsdk:"http_cookie" yaml:"httpCookie,omitempty"`

					HttpHeaderName *string `tfsdk:"http_header_name" yaml:"httpHeaderName,omitempty"`

					HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" yaml:"httpQueryParameterName,omitempty"`

					Maglev *struct {
						TableSize *int64 `tfsdk:"table_size" yaml:"tableSize,omitempty"`
					} `tfsdk:"maglev" yaml:"maglev,omitempty"`

					MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`

					RingHash *struct {
						MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`
					} `tfsdk:"ring_hash" yaml:"ringHash,omitempty"`

					UseSourceIp *bool `tfsdk:"use_source_ip" yaml:"useSourceIp,omitempty"`
				} `tfsdk:"consistent_hash" yaml:"consistentHash,omitempty"`

				LocalityLbSetting *struct {
					Distribute *[]struct {
						From *string `tfsdk:"from" yaml:"from,omitempty"`

						To *map[string]string `tfsdk:"to" yaml:"to,omitempty"`
					} `tfsdk:"distribute" yaml:"distribute,omitempty"`

					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

					Failover *[]struct {
						From *string `tfsdk:"from" yaml:"from,omitempty"`

						To *string `tfsdk:"to" yaml:"to,omitempty"`
					} `tfsdk:"failover" yaml:"failover,omitempty"`

					FailoverPriority *[]string `tfsdk:"failover_priority" yaml:"failoverPriority,omitempty"`
				} `tfsdk:"locality_lb_setting" yaml:"localityLbSetting,omitempty"`

				Simple *string `tfsdk:"simple" yaml:"simple,omitempty"`

				WarmupDurationSecs *string `tfsdk:"warmup_duration_secs" yaml:"warmupDurationSecs,omitempty"`
			} `tfsdk:"load_balancer" yaml:"loadBalancer,omitempty"`

			OutlierDetection *struct {
				BaseEjectionTime *string `tfsdk:"base_ejection_time" yaml:"baseEjectionTime,omitempty"`

				Consecutive5xxErrors *int64 `tfsdk:"consecutive5xx_errors" yaml:"consecutive5xxErrors,omitempty"`

				ConsecutiveErrors *int64 `tfsdk:"consecutive_errors" yaml:"consecutiveErrors,omitempty"`

				ConsecutiveGatewayErrors *int64 `tfsdk:"consecutive_gateway_errors" yaml:"consecutiveGatewayErrors,omitempty"`

				ConsecutiveLocalOriginFailures *int64 `tfsdk:"consecutive_local_origin_failures" yaml:"consecutiveLocalOriginFailures,omitempty"`

				Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

				MaxEjectionPercent *int64 `tfsdk:"max_ejection_percent" yaml:"maxEjectionPercent,omitempty"`

				MinHealthPercent *int64 `tfsdk:"min_health_percent" yaml:"minHealthPercent,omitempty"`

				SplitExternalLocalOriginErrors *bool `tfsdk:"split_external_local_origin_errors" yaml:"splitExternalLocalOriginErrors,omitempty"`
			} `tfsdk:"outlier_detection" yaml:"outlierDetection,omitempty"`

			PortLevelSettings *[]struct {
				ConnectionPool *struct {
					Http *struct {
						H2UpgradePolicy *string `tfsdk:"h2_upgrade_policy" yaml:"h2UpgradePolicy,omitempty"`

						Http1MaxPendingRequests *int64 `tfsdk:"http1_max_pending_requests" yaml:"http1MaxPendingRequests,omitempty"`

						Http2MaxRequests *int64 `tfsdk:"http2_max_requests" yaml:"http2MaxRequests,omitempty"`

						IdleTimeout *string `tfsdk:"idle_timeout" yaml:"idleTimeout,omitempty"`

						MaxRequestsPerConnection *int64 `tfsdk:"max_requests_per_connection" yaml:"maxRequestsPerConnection,omitempty"`

						MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`

						UseClientProtocol *bool `tfsdk:"use_client_protocol" yaml:"useClientProtocol,omitempty"`
					} `tfsdk:"http" yaml:"http,omitempty"`

					Tcp *struct {
						ConnectTimeout *string `tfsdk:"connect_timeout" yaml:"connectTimeout,omitempty"`

						MaxConnectionDuration *string `tfsdk:"max_connection_duration" yaml:"maxConnectionDuration,omitempty"`

						MaxConnections *int64 `tfsdk:"max_connections" yaml:"maxConnections,omitempty"`

						TcpKeepalive *struct {
							Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

							Probes *int64 `tfsdk:"probes" yaml:"probes,omitempty"`

							Time *string `tfsdk:"time" yaml:"time,omitempty"`
						} `tfsdk:"tcp_keepalive" yaml:"tcpKeepalive,omitempty"`
					} `tfsdk:"tcp" yaml:"tcp,omitempty"`
				} `tfsdk:"connection_pool" yaml:"connectionPool,omitempty"`

				LoadBalancer *struct {
					ConsistentHash *struct {
						HttpCookie *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Ttl *string `tfsdk:"ttl" yaml:"ttl,omitempty"`
						} `tfsdk:"http_cookie" yaml:"httpCookie,omitempty"`

						HttpHeaderName *string `tfsdk:"http_header_name" yaml:"httpHeaderName,omitempty"`

						HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" yaml:"httpQueryParameterName,omitempty"`

						Maglev *struct {
							TableSize *int64 `tfsdk:"table_size" yaml:"tableSize,omitempty"`
						} `tfsdk:"maglev" yaml:"maglev,omitempty"`

						MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`

						RingHash *struct {
							MinimumRingSize *int64 `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`
						} `tfsdk:"ring_hash" yaml:"ringHash,omitempty"`

						UseSourceIp *bool `tfsdk:"use_source_ip" yaml:"useSourceIp,omitempty"`
					} `tfsdk:"consistent_hash" yaml:"consistentHash,omitempty"`

					LocalityLbSetting *struct {
						Distribute *[]struct {
							From *string `tfsdk:"from" yaml:"from,omitempty"`

							To *map[string]string `tfsdk:"to" yaml:"to,omitempty"`
						} `tfsdk:"distribute" yaml:"distribute,omitempty"`

						Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

						Failover *[]struct {
							From *string `tfsdk:"from" yaml:"from,omitempty"`

							To *string `tfsdk:"to" yaml:"to,omitempty"`
						} `tfsdk:"failover" yaml:"failover,omitempty"`

						FailoverPriority *[]string `tfsdk:"failover_priority" yaml:"failoverPriority,omitempty"`
					} `tfsdk:"locality_lb_setting" yaml:"localityLbSetting,omitempty"`

					Simple *string `tfsdk:"simple" yaml:"simple,omitempty"`

					WarmupDurationSecs *string `tfsdk:"warmup_duration_secs" yaml:"warmupDurationSecs,omitempty"`
				} `tfsdk:"load_balancer" yaml:"loadBalancer,omitempty"`

				OutlierDetection *struct {
					BaseEjectionTime *string `tfsdk:"base_ejection_time" yaml:"baseEjectionTime,omitempty"`

					Consecutive5xxErrors *int64 `tfsdk:"consecutive5xx_errors" yaml:"consecutive5xxErrors,omitempty"`

					ConsecutiveErrors *int64 `tfsdk:"consecutive_errors" yaml:"consecutiveErrors,omitempty"`

					ConsecutiveGatewayErrors *int64 `tfsdk:"consecutive_gateway_errors" yaml:"consecutiveGatewayErrors,omitempty"`

					ConsecutiveLocalOriginFailures *int64 `tfsdk:"consecutive_local_origin_failures" yaml:"consecutiveLocalOriginFailures,omitempty"`

					Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

					MaxEjectionPercent *int64 `tfsdk:"max_ejection_percent" yaml:"maxEjectionPercent,omitempty"`

					MinHealthPercent *int64 `tfsdk:"min_health_percent" yaml:"minHealthPercent,omitempty"`

					SplitExternalLocalOriginErrors *bool `tfsdk:"split_external_local_origin_errors" yaml:"splitExternalLocalOriginErrors,omitempty"`
				} `tfsdk:"outlier_detection" yaml:"outlierDetection,omitempty"`

				Port *struct {
					Number *int64 `tfsdk:"number" yaml:"number,omitempty"`
				} `tfsdk:"port" yaml:"port,omitempty"`

				Tls *struct {
					CaCertificates *string `tfsdk:"ca_certificates" yaml:"caCertificates,omitempty"`

					ClientCertificate *string `tfsdk:"client_certificate" yaml:"clientCertificate,omitempty"`

					CredentialName *string `tfsdk:"credential_name" yaml:"credentialName,omitempty"`

					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

					Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

					SubjectAltNames *[]string `tfsdk:"subject_alt_names" yaml:"subjectAltNames,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`
			} `tfsdk:"port_level_settings" yaml:"portLevelSettings,omitempty"`

			Tls *struct {
				CaCertificates *string `tfsdk:"ca_certificates" yaml:"caCertificates,omitempty"`

				ClientCertificate *string `tfsdk:"client_certificate" yaml:"clientCertificate,omitempty"`

				CredentialName *string `tfsdk:"credential_name" yaml:"credentialName,omitempty"`

				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

				Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

				PrivateKey *string `tfsdk:"private_key" yaml:"privateKey,omitempty"`

				Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

				SubjectAltNames *[]string `tfsdk:"subject_alt_names" yaml:"subjectAltNames,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`

			Tunnel *struct {
				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

				TargetHost *string `tfsdk:"target_host" yaml:"targetHost,omitempty"`

				TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
			} `tfsdk:"tunnel" yaml:"tunnel,omitempty"`
		} `tfsdk:"traffic_policy" yaml:"trafficPolicy,omitempty"`

		WorkloadSelector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"workload_selector" yaml:"workloadSelector,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewNetworkingIstioIoDestinationRuleV1Alpha3Resource() resource.Resource {
	return &NetworkingIstioIoDestinationRuleV1Alpha3Resource{}
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networking_istio_io_destination_rule_v1alpha3"
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html",
				MarkdownDescription: "Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"export_to": {
						Description:         "A list of namespaces to which this destination rule is exported.",
						MarkdownDescription: "A list of namespaces to which this destination rule is exported.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host": {
						Description:         "The name of a service from the service registry.",
						MarkdownDescription: "The name of a service from the service registry.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"subsets": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"labels": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the subset.",
								MarkdownDescription: "Name of the subset.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"traffic_policy": {
								Description:         "Traffic policies that apply to this subset.",
								MarkdownDescription: "Traffic policies that apply to this subset.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"connection_pool": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"http": {
												Description:         "HTTP connection pool settings.",
												MarkdownDescription: "HTTP connection pool settings.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"h2_upgrade_policy": {
														Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
														MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http1_max_pending_requests": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http2_max_requests": {
														Description:         "Maximum number of active requests to a destination.",
														MarkdownDescription: "Maximum number of active requests to a destination.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"idle_timeout": {
														Description:         "The idle timeout for upstream connection pool connections.",
														MarkdownDescription: "The idle timeout for upstream connection pool connections.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_requests_per_connection": {
														Description:         "Maximum number of requests per connection to a backend.",
														MarkdownDescription: "Maximum number of requests per connection to a backend.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_retries": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"use_client_protocol": {
														Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
														MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",

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

											"tcp": {
												Description:         "Settings common to both HTTP and TCP upstream connections.",
												MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"connect_timeout": {
														Description:         "TCP connection timeout.",
														MarkdownDescription: "TCP connection timeout.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_connection_duration": {
														Description:         "The maximum duration of a connection.",
														MarkdownDescription: "The maximum duration of a connection.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_connections": {
														Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
														MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_keepalive": {
														Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
														MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"interval": {
																Description:         "The time duration between keep-alive probes.",
																MarkdownDescription: "The time duration between keep-alive probes.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"probes": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"time": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"load_balancer": {
										Description:         "Settings controlling the load balancer algorithms.",
										MarkdownDescription: "Settings controlling the load balancer algorithms.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"consistent_hash": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"http_cookie": {
														Description:         "Hash based on HTTP cookie.",
														MarkdownDescription: "Hash based on HTTP cookie.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the cookie.",
																MarkdownDescription: "Name of the cookie.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "Path to set for the cookie.",
																MarkdownDescription: "Path to set for the cookie.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ttl": {
																Description:         "Lifetime of the cookie.",
																MarkdownDescription: "Lifetime of the cookie.",

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

													"http_header_name": {
														Description:         "Hash based on a specific HTTP header.",
														MarkdownDescription: "Hash based on a specific HTTP header.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http_query_parameter_name": {
														Description:         "Hash based on a specific HTTP query parameter.",
														MarkdownDescription: "Hash based on a specific HTTP query parameter.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"maglev": {
														Description:         "The Maglev load balancer implements consistent hashing to backend hosts.",
														MarkdownDescription: "The Maglev load balancer implements consistent hashing to backend hosts.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"table_size": {
																Description:         "The table size for Maglev hashing.",
																MarkdownDescription: "The table size for Maglev hashing.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"minimum_ring_size": {
														Description:         "Deprecated.",
														MarkdownDescription: "Deprecated.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ring_hash": {
														Description:         "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
														MarkdownDescription: "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"minimum_ring_size": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"use_source_ip": {
														Description:         "Hash based on the source IP address.",
														MarkdownDescription: "Hash based on the source IP address.",

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

											"locality_lb_setting": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"distribute": {
														Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
														MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"from": {
																Description:         "Originating locality, '/' separated, e.g.",
																MarkdownDescription: "Originating locality, '/' separated, e.g.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"to": {
																Description:         "Map of upstream localities to traffic distribution weights.",
																MarkdownDescription: "Map of upstream localities to traffic distribution weights.",

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

													"enabled": {
														Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
														MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"failover": {
														Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
														MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"from": {
																Description:         "Originating region.",
																MarkdownDescription: "Originating region.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"to": {
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

													"failover_priority": {
														Description:         "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
														MarkdownDescription: "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",

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

											"simple": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"warmup_duration_secs": {
												Description:         "Represents the warmup duration of Service.",
												MarkdownDescription: "Represents the warmup duration of Service.",

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

									"outlier_detection": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"base_ejection_time": {
												Description:         "Minimum ejection duration.",
												MarkdownDescription: "Minimum ejection duration.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive5xx_errors": {
												Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
												MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive_errors": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive_gateway_errors": {
												Description:         "Number of gateway errors before a host is ejected from the connection pool.",
												MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive_local_origin_failures": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"interval": {
												Description:         "Time interval between ejection sweep analysis.",
												MarkdownDescription: "Time interval between ejection sweep analysis.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_ejection_percent": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_health_percent": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"split_external_local_origin_errors": {
												Description:         "Determines whether to distinguish local origin failures from external errors.",
												MarkdownDescription: "Determines whether to distinguish local origin failures from external errors.",

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

									"port_level_settings": {
										Description:         "Traffic policies specific to individual ports.",
										MarkdownDescription: "Traffic policies specific to individual ports.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"connection_pool": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"http": {
														Description:         "HTTP connection pool settings.",
														MarkdownDescription: "HTTP connection pool settings.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"h2_upgrade_policy": {
																Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
																MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http1_max_pending_requests": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http2_max_requests": {
																Description:         "Maximum number of active requests to a destination.",
																MarkdownDescription: "Maximum number of active requests to a destination.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"idle_timeout": {
																Description:         "The idle timeout for upstream connection pool connections.",
																MarkdownDescription: "The idle timeout for upstream connection pool connections.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"max_requests_per_connection": {
																Description:         "Maximum number of requests per connection to a backend.",
																MarkdownDescription: "Maximum number of requests per connection to a backend.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"max_retries": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"use_client_protocol": {
																Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
																MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",

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

													"tcp": {
														Description:         "Settings common to both HTTP and TCP upstream connections.",
														MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"connect_timeout": {
																Description:         "TCP connection timeout.",
																MarkdownDescription: "TCP connection timeout.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"max_connection_duration": {
																Description:         "The maximum duration of a connection.",
																MarkdownDescription: "The maximum duration of a connection.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"max_connections": {
																Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
																MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_keepalive": {
																Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
																MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"interval": {
																		Description:         "The time duration between keep-alive probes.",
																		MarkdownDescription: "The time duration between keep-alive probes.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"probes": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"time": {
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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"load_balancer": {
												Description:         "Settings controlling the load balancer algorithms.",
												MarkdownDescription: "Settings controlling the load balancer algorithms.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"consistent_hash": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"http_cookie": {
																Description:         "Hash based on HTTP cookie.",
																MarkdownDescription: "Hash based on HTTP cookie.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name of the cookie.",
																		MarkdownDescription: "Name of the cookie.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "Path to set for the cookie.",
																		MarkdownDescription: "Path to set for the cookie.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ttl": {
																		Description:         "Lifetime of the cookie.",
																		MarkdownDescription: "Lifetime of the cookie.",

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

															"http_header_name": {
																Description:         "Hash based on a specific HTTP header.",
																MarkdownDescription: "Hash based on a specific HTTP header.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"http_query_parameter_name": {
																Description:         "Hash based on a specific HTTP query parameter.",
																MarkdownDescription: "Hash based on a specific HTTP query parameter.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"maglev": {
																Description:         "The Maglev load balancer implements consistent hashing to backend hosts.",
																MarkdownDescription: "The Maglev load balancer implements consistent hashing to backend hosts.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"table_size": {
																		Description:         "The table size for Maglev hashing.",
																		MarkdownDescription: "The table size for Maglev hashing.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"minimum_ring_size": {
																Description:         "Deprecated.",
																MarkdownDescription: "Deprecated.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ring_hash": {
																Description:         "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
																MarkdownDescription: "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"minimum_ring_size": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"use_source_ip": {
																Description:         "Hash based on the source IP address.",
																MarkdownDescription: "Hash based on the source IP address.",

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

													"locality_lb_setting": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"distribute": {
																Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
																MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"from": {
																		Description:         "Originating locality, '/' separated, e.g.",
																		MarkdownDescription: "Originating locality, '/' separated, e.g.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"to": {
																		Description:         "Map of upstream localities to traffic distribution weights.",
																		MarkdownDescription: "Map of upstream localities to traffic distribution weights.",

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

															"enabled": {
																Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
																MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"failover": {
																Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
																MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"from": {
																		Description:         "Originating region.",
																		MarkdownDescription: "Originating region.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"to": {
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

															"failover_priority": {
																Description:         "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
																MarkdownDescription: "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",

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

													"simple": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"warmup_duration_secs": {
														Description:         "Represents the warmup duration of Service.",
														MarkdownDescription: "Represents the warmup duration of Service.",

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

											"outlier_detection": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"base_ejection_time": {
														Description:         "Minimum ejection duration.",
														MarkdownDescription: "Minimum ejection duration.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"consecutive5xx_errors": {
														Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
														MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"consecutive_errors": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"consecutive_gateway_errors": {
														Description:         "Number of gateway errors before a host is ejected from the connection pool.",
														MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"consecutive_local_origin_failures": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"interval": {
														Description:         "Time interval between ejection sweep analysis.",
														MarkdownDescription: "Time interval between ejection sweep analysis.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_ejection_percent": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"min_health_percent": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"split_external_local_origin_errors": {
														Description:         "Determines whether to distinguish local origin failures from external errors.",
														MarkdownDescription: "Determines whether to distinguish local origin failures from external errors.",

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

											"port": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"number": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls": {
												Description:         "TLS related settings for connections to the upstream service.",
												MarkdownDescription: "TLS related settings for connections to the upstream service.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca_certificates": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"client_certificate": {
														Description:         "REQUIRED if mode is 'MUTUAL'.",
														MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"credential_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"mode": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"private_key": {
														Description:         "REQUIRED if mode is 'MUTUAL'.",
														MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sni": {
														Description:         "SNI string to present to the server during TLS handshake.",
														MarkdownDescription: "SNI string to present to the server during TLS handshake.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"subject_alt_names": {
														Description:         "",
														MarkdownDescription: "",

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

									"tls": {
										Description:         "TLS related settings for connections to the upstream service.",
										MarkdownDescription: "TLS related settings for connections to the upstream service.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_certificates": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_certificate": {
												Description:         "REQUIRED if mode is 'MUTUAL'.",
												MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"credential_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_skip_verify": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "REQUIRED if mode is 'MUTUAL'.",
												MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sni": {
												Description:         "SNI string to present to the server during TLS handshake.",
												MarkdownDescription: "SNI string to present to the server during TLS handshake.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subject_alt_names": {
												Description:         "",
												MarkdownDescription: "",

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

									"tunnel": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"protocol": {
												Description:         "Specifies which protocol to use for tunneling the downstream connection.",
												MarkdownDescription: "Specifies which protocol to use for tunneling the downstream connection.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target_host": {
												Description:         "Specifies a host to which the downstream connection is tunneled.",
												MarkdownDescription: "Specifies a host to which the downstream connection is tunneled.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target_port": {
												Description:         "Specifies a port to which the downstream connection is tunneled.",
												MarkdownDescription: "Specifies a port to which the downstream connection is tunneled.",

												Type: types.Int64Type,

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"traffic_policy": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"connection_pool": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"http": {
										Description:         "HTTP connection pool settings.",
										MarkdownDescription: "HTTP connection pool settings.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"h2_upgrade_policy": {
												Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
												MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http1_max_pending_requests": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http2_max_requests": {
												Description:         "Maximum number of active requests to a destination.",
												MarkdownDescription: "Maximum number of active requests to a destination.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"idle_timeout": {
												Description:         "The idle timeout for upstream connection pool connections.",
												MarkdownDescription: "The idle timeout for upstream connection pool connections.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_requests_per_connection": {
												Description:         "Maximum number of requests per connection to a backend.",
												MarkdownDescription: "Maximum number of requests per connection to a backend.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_retries": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"use_client_protocol": {
												Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
												MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",

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

									"tcp": {
										Description:         "Settings common to both HTTP and TCP upstream connections.",
										MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"connect_timeout": {
												Description:         "TCP connection timeout.",
												MarkdownDescription: "TCP connection timeout.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_connection_duration": {
												Description:         "The maximum duration of a connection.",
												MarkdownDescription: "The maximum duration of a connection.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_connections": {
												Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
												MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tcp_keepalive": {
												Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
												MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"interval": {
														Description:         "The time duration between keep-alive probes.",
														MarkdownDescription: "The time duration between keep-alive probes.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"probes": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"time": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"load_balancer": {
								Description:         "Settings controlling the load balancer algorithms.",
								MarkdownDescription: "Settings controlling the load balancer algorithms.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"consistent_hash": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"http_cookie": {
												Description:         "Hash based on HTTP cookie.",
												MarkdownDescription: "Hash based on HTTP cookie.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the cookie.",
														MarkdownDescription: "Name of the cookie.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Path to set for the cookie.",
														MarkdownDescription: "Path to set for the cookie.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ttl": {
														Description:         "Lifetime of the cookie.",
														MarkdownDescription: "Lifetime of the cookie.",

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

											"http_header_name": {
												Description:         "Hash based on a specific HTTP header.",
												MarkdownDescription: "Hash based on a specific HTTP header.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_query_parameter_name": {
												Description:         "Hash based on a specific HTTP query parameter.",
												MarkdownDescription: "Hash based on a specific HTTP query parameter.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"maglev": {
												Description:         "The Maglev load balancer implements consistent hashing to backend hosts.",
												MarkdownDescription: "The Maglev load balancer implements consistent hashing to backend hosts.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"table_size": {
														Description:         "The table size for Maglev hashing.",
														MarkdownDescription: "The table size for Maglev hashing.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum_ring_size": {
												Description:         "Deprecated.",
												MarkdownDescription: "Deprecated.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ring_hash": {
												Description:         "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
												MarkdownDescription: "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"minimum_ring_size": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"use_source_ip": {
												Description:         "Hash based on the source IP address.",
												MarkdownDescription: "Hash based on the source IP address.",

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

									"locality_lb_setting": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"distribute": {
												Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
												MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"from": {
														Description:         "Originating locality, '/' separated, e.g.",
														MarkdownDescription: "Originating locality, '/' separated, e.g.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"to": {
														Description:         "Map of upstream localities to traffic distribution weights.",
														MarkdownDescription: "Map of upstream localities to traffic distribution weights.",

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

											"enabled": {
												Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
												MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"failover": {
												Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
												MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"from": {
														Description:         "Originating region.",
														MarkdownDescription: "Originating region.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"to": {
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

											"failover_priority": {
												Description:         "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
												MarkdownDescription: "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",

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

									"simple": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"warmup_duration_secs": {
										Description:         "Represents the warmup duration of Service.",
										MarkdownDescription: "Represents the warmup duration of Service.",

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

							"outlier_detection": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"base_ejection_time": {
										Description:         "Minimum ejection duration.",
										MarkdownDescription: "Minimum ejection duration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"consecutive5xx_errors": {
										Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
										MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"consecutive_errors": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"consecutive_gateway_errors": {
										Description:         "Number of gateway errors before a host is ejected from the connection pool.",
										MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"consecutive_local_origin_failures": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"interval": {
										Description:         "Time interval between ejection sweep analysis.",
										MarkdownDescription: "Time interval between ejection sweep analysis.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_ejection_percent": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_health_percent": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"split_external_local_origin_errors": {
										Description:         "Determines whether to distinguish local origin failures from external errors.",
										MarkdownDescription: "Determines whether to distinguish local origin failures from external errors.",

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

							"port_level_settings": {
								Description:         "Traffic policies specific to individual ports.",
								MarkdownDescription: "Traffic policies specific to individual ports.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"connection_pool": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"http": {
												Description:         "HTTP connection pool settings.",
												MarkdownDescription: "HTTP connection pool settings.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"h2_upgrade_policy": {
														Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
														MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http1_max_pending_requests": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http2_max_requests": {
														Description:         "Maximum number of active requests to a destination.",
														MarkdownDescription: "Maximum number of active requests to a destination.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"idle_timeout": {
														Description:         "The idle timeout for upstream connection pool connections.",
														MarkdownDescription: "The idle timeout for upstream connection pool connections.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_requests_per_connection": {
														Description:         "Maximum number of requests per connection to a backend.",
														MarkdownDescription: "Maximum number of requests per connection to a backend.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_retries": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"use_client_protocol": {
														Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
														MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",

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

											"tcp": {
												Description:         "Settings common to both HTTP and TCP upstream connections.",
												MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"connect_timeout": {
														Description:         "TCP connection timeout.",
														MarkdownDescription: "TCP connection timeout.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_connection_duration": {
														Description:         "The maximum duration of a connection.",
														MarkdownDescription: "The maximum duration of a connection.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_connections": {
														Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
														MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp_keepalive": {
														Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
														MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"interval": {
																Description:         "The time duration between keep-alive probes.",
																MarkdownDescription: "The time duration between keep-alive probes.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"probes": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"time": {
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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"load_balancer": {
										Description:         "Settings controlling the load balancer algorithms.",
										MarkdownDescription: "Settings controlling the load balancer algorithms.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"consistent_hash": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"http_cookie": {
														Description:         "Hash based on HTTP cookie.",
														MarkdownDescription: "Hash based on HTTP cookie.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"name": {
																Description:         "Name of the cookie.",
																MarkdownDescription: "Name of the cookie.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "Path to set for the cookie.",
																MarkdownDescription: "Path to set for the cookie.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ttl": {
																Description:         "Lifetime of the cookie.",
																MarkdownDescription: "Lifetime of the cookie.",

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

													"http_header_name": {
														Description:         "Hash based on a specific HTTP header.",
														MarkdownDescription: "Hash based on a specific HTTP header.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http_query_parameter_name": {
														Description:         "Hash based on a specific HTTP query parameter.",
														MarkdownDescription: "Hash based on a specific HTTP query parameter.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"maglev": {
														Description:         "The Maglev load balancer implements consistent hashing to backend hosts.",
														MarkdownDescription: "The Maglev load balancer implements consistent hashing to backend hosts.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"table_size": {
																Description:         "The table size for Maglev hashing.",
																MarkdownDescription: "The table size for Maglev hashing.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"minimum_ring_size": {
														Description:         "Deprecated.",
														MarkdownDescription: "Deprecated.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ring_hash": {
														Description:         "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
														MarkdownDescription: "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"minimum_ring_size": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"use_source_ip": {
														Description:         "Hash based on the source IP address.",
														MarkdownDescription: "Hash based on the source IP address.",

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

											"locality_lb_setting": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"distribute": {
														Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
														MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"from": {
																Description:         "Originating locality, '/' separated, e.g.",
																MarkdownDescription: "Originating locality, '/' separated, e.g.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"to": {
																Description:         "Map of upstream localities to traffic distribution weights.",
																MarkdownDescription: "Map of upstream localities to traffic distribution weights.",

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

													"enabled": {
														Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
														MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"failover": {
														Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
														MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"from": {
																Description:         "Originating region.",
																MarkdownDescription: "Originating region.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"to": {
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

													"failover_priority": {
														Description:         "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
														MarkdownDescription: "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",

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

											"simple": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"warmup_duration_secs": {
												Description:         "Represents the warmup duration of Service.",
												MarkdownDescription: "Represents the warmup duration of Service.",

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

									"outlier_detection": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"base_ejection_time": {
												Description:         "Minimum ejection duration.",
												MarkdownDescription: "Minimum ejection duration.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive5xx_errors": {
												Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
												MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive_errors": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive_gateway_errors": {
												Description:         "Number of gateway errors before a host is ejected from the connection pool.",
												MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"consecutive_local_origin_failures": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"interval": {
												Description:         "Time interval between ejection sweep analysis.",
												MarkdownDescription: "Time interval between ejection sweep analysis.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_ejection_percent": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"min_health_percent": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"split_external_local_origin_errors": {
												Description:         "Determines whether to distinguish local origin failures from external errors.",
												MarkdownDescription: "Determines whether to distinguish local origin failures from external errors.",

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

									"port": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"number": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "TLS related settings for connections to the upstream service.",
										MarkdownDescription: "TLS related settings for connections to the upstream service.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_certificates": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_certificate": {
												Description:         "REQUIRED if mode is 'MUTUAL'.",
												MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"credential_name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_skip_verify": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"mode": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"private_key": {
												Description:         "REQUIRED if mode is 'MUTUAL'.",
												MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sni": {
												Description:         "SNI string to present to the server during TLS handshake.",
												MarkdownDescription: "SNI string to present to the server during TLS handshake.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"subject_alt_names": {
												Description:         "",
												MarkdownDescription: "",

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

							"tls": {
								Description:         "TLS related settings for connections to the upstream service.",
								MarkdownDescription: "TLS related settings for connections to the upstream service.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_certificates": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_certificate": {
										Description:         "REQUIRED if mode is 'MUTUAL'.",
										MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"credential_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mode": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"private_key": {
										Description:         "REQUIRED if mode is 'MUTUAL'.",
										MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sni": {
										Description:         "SNI string to present to the server during TLS handshake.",
										MarkdownDescription: "SNI string to present to the server during TLS handshake.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subject_alt_names": {
										Description:         "",
										MarkdownDescription: "",

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

							"tunnel": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"protocol": {
										Description:         "Specifies which protocol to use for tunneling the downstream connection.",
										MarkdownDescription: "Specifies which protocol to use for tunneling the downstream connection.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_host": {
										Description:         "Specifies a host to which the downstream connection is tunneled.",
										MarkdownDescription: "Specifies a host to which the downstream connection is tunneled.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_port": {
										Description:         "Specifies a port to which the downstream connection is tunneled.",
										MarkdownDescription: "Specifies a port to which the downstream connection is tunneled.",

										Type: types.Int64Type,

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

					"workload_selector": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"match_labels": {
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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_destination_rule_v1alpha3")

	var state NetworkingIstioIoDestinationRuleV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoDestinationRuleV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("DestinationRule")

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

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_destination_rule_v1alpha3")
	// NO-OP: All data is already in Terraform state
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_destination_rule_v1alpha3")

	var state NetworkingIstioIoDestinationRuleV1Alpha3TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel NetworkingIstioIoDestinationRuleV1Alpha3GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("networking.istio.io/v1alpha3")
	goModel.Kind = utilities.Ptr("DestinationRule")

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

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_destination_rule_v1alpha3")
	// NO-OP: Terraform removes the state automatically for us
}
