/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &NetworkingIstioIoDestinationRuleV1Alpha3Manifest{}
)

func NewNetworkingIstioIoDestinationRuleV1Alpha3Manifest() datasource.DataSource {
	return &NetworkingIstioIoDestinationRuleV1Alpha3Manifest{}
}

type NetworkingIstioIoDestinationRuleV1Alpha3Manifest struct{}

type NetworkingIstioIoDestinationRuleV1Alpha3ManifestData struct {
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
		ExportTo *[]string `tfsdk:"export_to" json:"exportTo,omitempty"`
		Host     *string   `tfsdk:"host" json:"host,omitempty"`
		Subsets  *[]struct {
			Labels        *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			Name          *string            `tfsdk:"name" json:"name,omitempty"`
			TrafficPolicy *struct {
				ConnectionPool *struct {
					Http *struct {
						H2UpgradePolicy          *string `tfsdk:"h2_upgrade_policy" json:"h2UpgradePolicy,omitempty"`
						Http1MaxPendingRequests  *int64  `tfsdk:"http1_max_pending_requests" json:"http1MaxPendingRequests,omitempty"`
						Http2MaxRequests         *int64  `tfsdk:"http2_max_requests" json:"http2MaxRequests,omitempty"`
						IdleTimeout              *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
						MaxConcurrentStreams     *int64  `tfsdk:"max_concurrent_streams" json:"maxConcurrentStreams,omitempty"`
						MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
						MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
						UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Tcp *struct {
						ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
						IdleTimeout           *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
						MaxConnectionDuration *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
						MaxConnections        *int64  `tfsdk:"max_connections" json:"maxConnections,omitempty"`
						TcpKeepalive          *struct {
							Interval *string `tfsdk:"interval" json:"interval,omitempty"`
							Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
							Time     *string `tfsdk:"time" json:"time,omitempty"`
						} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
					} `tfsdk:"tcp" json:"tcp,omitempty"`
				} `tfsdk:"connection_pool" json:"connectionPool,omitempty"`
				LoadBalancer *struct {
					ConsistentHash *struct {
						HttpCookie *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
							Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
						} `tfsdk:"http_cookie" json:"httpCookie,omitempty"`
						HttpHeaderName         *string `tfsdk:"http_header_name" json:"httpHeaderName,omitempty"`
						HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" json:"httpQueryParameterName,omitempty"`
						Maglev                 *struct {
							TableSize *int64 `tfsdk:"table_size" json:"tableSize,omitempty"`
						} `tfsdk:"maglev" json:"maglev,omitempty"`
						MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
						RingHash        *struct {
							MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
						} `tfsdk:"ring_hash" json:"ringHash,omitempty"`
						UseSourceIp *bool `tfsdk:"use_source_ip" json:"useSourceIp,omitempty"`
					} `tfsdk:"consistent_hash" json:"consistentHash,omitempty"`
					LocalityLbSetting *struct {
						Distribute *[]struct {
							From *string            `tfsdk:"from" json:"from,omitempty"`
							To   *map[string]string `tfsdk:"to" json:"to,omitempty"`
						} `tfsdk:"distribute" json:"distribute,omitempty"`
						Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
						Failover *[]struct {
							From *string `tfsdk:"from" json:"from,omitempty"`
							To   *string `tfsdk:"to" json:"to,omitempty"`
						} `tfsdk:"failover" json:"failover,omitempty"`
						FailoverPriority *[]string `tfsdk:"failover_priority" json:"failoverPriority,omitempty"`
					} `tfsdk:"locality_lb_setting" json:"localityLbSetting,omitempty"`
					Simple             *string `tfsdk:"simple" json:"simple,omitempty"`
					WarmupDurationSecs *string `tfsdk:"warmup_duration_secs" json:"warmupDurationSecs,omitempty"`
				} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
				OutlierDetection *struct {
					BaseEjectionTime               *string `tfsdk:"base_ejection_time" json:"baseEjectionTime,omitempty"`
					Consecutive5xxErrors           *int64  `tfsdk:"consecutive5xx_errors" json:"consecutive5xxErrors,omitempty"`
					ConsecutiveErrors              *int64  `tfsdk:"consecutive_errors" json:"consecutiveErrors,omitempty"`
					ConsecutiveGatewayErrors       *int64  `tfsdk:"consecutive_gateway_errors" json:"consecutiveGatewayErrors,omitempty"`
					ConsecutiveLocalOriginFailures *int64  `tfsdk:"consecutive_local_origin_failures" json:"consecutiveLocalOriginFailures,omitempty"`
					Interval                       *string `tfsdk:"interval" json:"interval,omitempty"`
					MaxEjectionPercent             *int64  `tfsdk:"max_ejection_percent" json:"maxEjectionPercent,omitempty"`
					MinHealthPercent               *int64  `tfsdk:"min_health_percent" json:"minHealthPercent,omitempty"`
					SplitExternalLocalOriginErrors *bool   `tfsdk:"split_external_local_origin_errors" json:"splitExternalLocalOriginErrors,omitempty"`
				} `tfsdk:"outlier_detection" json:"outlierDetection,omitempty"`
				PortLevelSettings *[]struct {
					ConnectionPool *struct {
						Http *struct {
							H2UpgradePolicy          *string `tfsdk:"h2_upgrade_policy" json:"h2UpgradePolicy,omitempty"`
							Http1MaxPendingRequests  *int64  `tfsdk:"http1_max_pending_requests" json:"http1MaxPendingRequests,omitempty"`
							Http2MaxRequests         *int64  `tfsdk:"http2_max_requests" json:"http2MaxRequests,omitempty"`
							IdleTimeout              *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
							MaxConcurrentStreams     *int64  `tfsdk:"max_concurrent_streams" json:"maxConcurrentStreams,omitempty"`
							MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
							MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
							UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
						} `tfsdk:"http" json:"http,omitempty"`
						Tcp *struct {
							ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
							IdleTimeout           *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
							MaxConnectionDuration *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
							MaxConnections        *int64  `tfsdk:"max_connections" json:"maxConnections,omitempty"`
							TcpKeepalive          *struct {
								Interval *string `tfsdk:"interval" json:"interval,omitempty"`
								Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
								Time     *string `tfsdk:"time" json:"time,omitempty"`
							} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
						} `tfsdk:"tcp" json:"tcp,omitempty"`
					} `tfsdk:"connection_pool" json:"connectionPool,omitempty"`
					LoadBalancer *struct {
						ConsistentHash *struct {
							HttpCookie *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
								Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
							} `tfsdk:"http_cookie" json:"httpCookie,omitempty"`
							HttpHeaderName         *string `tfsdk:"http_header_name" json:"httpHeaderName,omitempty"`
							HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" json:"httpQueryParameterName,omitempty"`
							Maglev                 *struct {
								TableSize *int64 `tfsdk:"table_size" json:"tableSize,omitempty"`
							} `tfsdk:"maglev" json:"maglev,omitempty"`
							MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
							RingHash        *struct {
								MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
							} `tfsdk:"ring_hash" json:"ringHash,omitempty"`
							UseSourceIp *bool `tfsdk:"use_source_ip" json:"useSourceIp,omitempty"`
						} `tfsdk:"consistent_hash" json:"consistentHash,omitempty"`
						LocalityLbSetting *struct {
							Distribute *[]struct {
								From *string            `tfsdk:"from" json:"from,omitempty"`
								To   *map[string]string `tfsdk:"to" json:"to,omitempty"`
							} `tfsdk:"distribute" json:"distribute,omitempty"`
							Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
							Failover *[]struct {
								From *string `tfsdk:"from" json:"from,omitempty"`
								To   *string `tfsdk:"to" json:"to,omitempty"`
							} `tfsdk:"failover" json:"failover,omitempty"`
							FailoverPriority *[]string `tfsdk:"failover_priority" json:"failoverPriority,omitempty"`
						} `tfsdk:"locality_lb_setting" json:"localityLbSetting,omitempty"`
						Simple             *string `tfsdk:"simple" json:"simple,omitempty"`
						WarmupDurationSecs *string `tfsdk:"warmup_duration_secs" json:"warmupDurationSecs,omitempty"`
					} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
					OutlierDetection *struct {
						BaseEjectionTime               *string `tfsdk:"base_ejection_time" json:"baseEjectionTime,omitempty"`
						Consecutive5xxErrors           *int64  `tfsdk:"consecutive5xx_errors" json:"consecutive5xxErrors,omitempty"`
						ConsecutiveErrors              *int64  `tfsdk:"consecutive_errors" json:"consecutiveErrors,omitempty"`
						ConsecutiveGatewayErrors       *int64  `tfsdk:"consecutive_gateway_errors" json:"consecutiveGatewayErrors,omitempty"`
						ConsecutiveLocalOriginFailures *int64  `tfsdk:"consecutive_local_origin_failures" json:"consecutiveLocalOriginFailures,omitempty"`
						Interval                       *string `tfsdk:"interval" json:"interval,omitempty"`
						MaxEjectionPercent             *int64  `tfsdk:"max_ejection_percent" json:"maxEjectionPercent,omitempty"`
						MinHealthPercent               *int64  `tfsdk:"min_health_percent" json:"minHealthPercent,omitempty"`
						SplitExternalLocalOriginErrors *bool   `tfsdk:"split_external_local_origin_errors" json:"splitExternalLocalOriginErrors,omitempty"`
					} `tfsdk:"outlier_detection" json:"outlierDetection,omitempty"`
					Port *struct {
						Number *int64 `tfsdk:"number" json:"number,omitempty"`
					} `tfsdk:"port" json:"port,omitempty"`
					Tls *struct {
						CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
						CaCrl              *string   `tfsdk:"ca_crl" json:"caCrl,omitempty"`
						ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
						CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
						InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
						PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
						Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
						SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"port_level_settings" json:"portLevelSettings,omitempty"`
				ProxyProtocol *struct {
					Version *string `tfsdk:"version" json:"version,omitempty"`
				} `tfsdk:"proxy_protocol" json:"proxyProtocol,omitempty"`
				Tls *struct {
					CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
					CaCrl              *string   `tfsdk:"ca_crl" json:"caCrl,omitempty"`
					ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
					CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
					InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
					PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
					Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
					SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				Tunnel *struct {
					Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
					TargetHost *string `tfsdk:"target_host" json:"targetHost,omitempty"`
					TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"tunnel" json:"tunnel,omitempty"`
			} `tfsdk:"traffic_policy" json:"trafficPolicy,omitempty"`
		} `tfsdk:"subsets" json:"subsets,omitempty"`
		TrafficPolicy *struct {
			ConnectionPool *struct {
				Http *struct {
					H2UpgradePolicy          *string `tfsdk:"h2_upgrade_policy" json:"h2UpgradePolicy,omitempty"`
					Http1MaxPendingRequests  *int64  `tfsdk:"http1_max_pending_requests" json:"http1MaxPendingRequests,omitempty"`
					Http2MaxRequests         *int64  `tfsdk:"http2_max_requests" json:"http2MaxRequests,omitempty"`
					IdleTimeout              *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
					MaxConcurrentStreams     *int64  `tfsdk:"max_concurrent_streams" json:"maxConcurrentStreams,omitempty"`
					MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
					MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Tcp *struct {
					ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
					IdleTimeout           *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
					MaxConnectionDuration *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
					MaxConnections        *int64  `tfsdk:"max_connections" json:"maxConnections,omitempty"`
					TcpKeepalive          *struct {
						Interval *string `tfsdk:"interval" json:"interval,omitempty"`
						Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
						Time     *string `tfsdk:"time" json:"time,omitempty"`
					} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
				} `tfsdk:"tcp" json:"tcp,omitempty"`
			} `tfsdk:"connection_pool" json:"connectionPool,omitempty"`
			LoadBalancer *struct {
				ConsistentHash *struct {
					HttpCookie *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
						Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
					} `tfsdk:"http_cookie" json:"httpCookie,omitempty"`
					HttpHeaderName         *string `tfsdk:"http_header_name" json:"httpHeaderName,omitempty"`
					HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" json:"httpQueryParameterName,omitempty"`
					Maglev                 *struct {
						TableSize *int64 `tfsdk:"table_size" json:"tableSize,omitempty"`
					} `tfsdk:"maglev" json:"maglev,omitempty"`
					MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
					RingHash        *struct {
						MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
					} `tfsdk:"ring_hash" json:"ringHash,omitempty"`
					UseSourceIp *bool `tfsdk:"use_source_ip" json:"useSourceIp,omitempty"`
				} `tfsdk:"consistent_hash" json:"consistentHash,omitempty"`
				LocalityLbSetting *struct {
					Distribute *[]struct {
						From *string            `tfsdk:"from" json:"from,omitempty"`
						To   *map[string]string `tfsdk:"to" json:"to,omitempty"`
					} `tfsdk:"distribute" json:"distribute,omitempty"`
					Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					Failover *[]struct {
						From *string `tfsdk:"from" json:"from,omitempty"`
						To   *string `tfsdk:"to" json:"to,omitempty"`
					} `tfsdk:"failover" json:"failover,omitempty"`
					FailoverPriority *[]string `tfsdk:"failover_priority" json:"failoverPriority,omitempty"`
				} `tfsdk:"locality_lb_setting" json:"localityLbSetting,omitempty"`
				Simple             *string `tfsdk:"simple" json:"simple,omitempty"`
				WarmupDurationSecs *string `tfsdk:"warmup_duration_secs" json:"warmupDurationSecs,omitempty"`
			} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
			OutlierDetection *struct {
				BaseEjectionTime               *string `tfsdk:"base_ejection_time" json:"baseEjectionTime,omitempty"`
				Consecutive5xxErrors           *int64  `tfsdk:"consecutive5xx_errors" json:"consecutive5xxErrors,omitempty"`
				ConsecutiveErrors              *int64  `tfsdk:"consecutive_errors" json:"consecutiveErrors,omitempty"`
				ConsecutiveGatewayErrors       *int64  `tfsdk:"consecutive_gateway_errors" json:"consecutiveGatewayErrors,omitempty"`
				ConsecutiveLocalOriginFailures *int64  `tfsdk:"consecutive_local_origin_failures" json:"consecutiveLocalOriginFailures,omitempty"`
				Interval                       *string `tfsdk:"interval" json:"interval,omitempty"`
				MaxEjectionPercent             *int64  `tfsdk:"max_ejection_percent" json:"maxEjectionPercent,omitempty"`
				MinHealthPercent               *int64  `tfsdk:"min_health_percent" json:"minHealthPercent,omitempty"`
				SplitExternalLocalOriginErrors *bool   `tfsdk:"split_external_local_origin_errors" json:"splitExternalLocalOriginErrors,omitempty"`
			} `tfsdk:"outlier_detection" json:"outlierDetection,omitempty"`
			PortLevelSettings *[]struct {
				ConnectionPool *struct {
					Http *struct {
						H2UpgradePolicy          *string `tfsdk:"h2_upgrade_policy" json:"h2UpgradePolicy,omitempty"`
						Http1MaxPendingRequests  *int64  `tfsdk:"http1_max_pending_requests" json:"http1MaxPendingRequests,omitempty"`
						Http2MaxRequests         *int64  `tfsdk:"http2_max_requests" json:"http2MaxRequests,omitempty"`
						IdleTimeout              *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
						MaxConcurrentStreams     *int64  `tfsdk:"max_concurrent_streams" json:"maxConcurrentStreams,omitempty"`
						MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
						MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
						UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Tcp *struct {
						ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
						IdleTimeout           *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
						MaxConnectionDuration *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
						MaxConnections        *int64  `tfsdk:"max_connections" json:"maxConnections,omitempty"`
						TcpKeepalive          *struct {
							Interval *string `tfsdk:"interval" json:"interval,omitempty"`
							Probes   *int64  `tfsdk:"probes" json:"probes,omitempty"`
							Time     *string `tfsdk:"time" json:"time,omitempty"`
						} `tfsdk:"tcp_keepalive" json:"tcpKeepalive,omitempty"`
					} `tfsdk:"tcp" json:"tcp,omitempty"`
				} `tfsdk:"connection_pool" json:"connectionPool,omitempty"`
				LoadBalancer *struct {
					ConsistentHash *struct {
						HttpCookie *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
							Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
						} `tfsdk:"http_cookie" json:"httpCookie,omitempty"`
						HttpHeaderName         *string `tfsdk:"http_header_name" json:"httpHeaderName,omitempty"`
						HttpQueryParameterName *string `tfsdk:"http_query_parameter_name" json:"httpQueryParameterName,omitempty"`
						Maglev                 *struct {
							TableSize *int64 `tfsdk:"table_size" json:"tableSize,omitempty"`
						} `tfsdk:"maglev" json:"maglev,omitempty"`
						MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
						RingHash        *struct {
							MinimumRingSize *int64 `tfsdk:"minimum_ring_size" json:"minimumRingSize,omitempty"`
						} `tfsdk:"ring_hash" json:"ringHash,omitempty"`
						UseSourceIp *bool `tfsdk:"use_source_ip" json:"useSourceIp,omitempty"`
					} `tfsdk:"consistent_hash" json:"consistentHash,omitempty"`
					LocalityLbSetting *struct {
						Distribute *[]struct {
							From *string            `tfsdk:"from" json:"from,omitempty"`
							To   *map[string]string `tfsdk:"to" json:"to,omitempty"`
						} `tfsdk:"distribute" json:"distribute,omitempty"`
						Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
						Failover *[]struct {
							From *string `tfsdk:"from" json:"from,omitempty"`
							To   *string `tfsdk:"to" json:"to,omitempty"`
						} `tfsdk:"failover" json:"failover,omitempty"`
						FailoverPriority *[]string `tfsdk:"failover_priority" json:"failoverPriority,omitempty"`
					} `tfsdk:"locality_lb_setting" json:"localityLbSetting,omitempty"`
					Simple             *string `tfsdk:"simple" json:"simple,omitempty"`
					WarmupDurationSecs *string `tfsdk:"warmup_duration_secs" json:"warmupDurationSecs,omitempty"`
				} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
				OutlierDetection *struct {
					BaseEjectionTime               *string `tfsdk:"base_ejection_time" json:"baseEjectionTime,omitempty"`
					Consecutive5xxErrors           *int64  `tfsdk:"consecutive5xx_errors" json:"consecutive5xxErrors,omitempty"`
					ConsecutiveErrors              *int64  `tfsdk:"consecutive_errors" json:"consecutiveErrors,omitempty"`
					ConsecutiveGatewayErrors       *int64  `tfsdk:"consecutive_gateway_errors" json:"consecutiveGatewayErrors,omitempty"`
					ConsecutiveLocalOriginFailures *int64  `tfsdk:"consecutive_local_origin_failures" json:"consecutiveLocalOriginFailures,omitempty"`
					Interval                       *string `tfsdk:"interval" json:"interval,omitempty"`
					MaxEjectionPercent             *int64  `tfsdk:"max_ejection_percent" json:"maxEjectionPercent,omitempty"`
					MinHealthPercent               *int64  `tfsdk:"min_health_percent" json:"minHealthPercent,omitempty"`
					SplitExternalLocalOriginErrors *bool   `tfsdk:"split_external_local_origin_errors" json:"splitExternalLocalOriginErrors,omitempty"`
				} `tfsdk:"outlier_detection" json:"outlierDetection,omitempty"`
				Port *struct {
					Number *int64 `tfsdk:"number" json:"number,omitempty"`
				} `tfsdk:"port" json:"port,omitempty"`
				Tls *struct {
					CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
					CaCrl              *string   `tfsdk:"ca_crl" json:"caCrl,omitempty"`
					ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
					CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
					InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
					PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
					Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
					SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"port_level_settings" json:"portLevelSettings,omitempty"`
			ProxyProtocol *struct {
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"proxy_protocol" json:"proxyProtocol,omitempty"`
			Tls *struct {
				CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
				CaCrl              *string   `tfsdk:"ca_crl" json:"caCrl,omitempty"`
				ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
				CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
				InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
				Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
				PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
				Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
				SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
			Tunnel *struct {
				Protocol   *string `tfsdk:"protocol" json:"protocol,omitempty"`
				TargetHost *string `tfsdk:"target_host" json:"targetHost,omitempty"`
				TargetPort *int64  `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"tunnel" json:"tunnel,omitempty"`
		} `tfsdk:"traffic_policy" json:"trafficPolicy,omitempty"`
		WorkloadSelector *struct {
			MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
		} `tfsdk:"workload_selector" json:"workloadSelector,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_destination_rule_v1alpha3_manifest"
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html",
				MarkdownDescription: "Configuration affecting load balancing, outlier detection, etc. See more details at: https://istio.io/docs/reference/config/networking/destination-rule.html",
				Attributes: map[string]schema.Attribute{
					"export_to": schema.ListAttribute{
						Description:         "A list of namespaces to which this destination rule is exported.",
						MarkdownDescription: "A list of namespaces to which this destination rule is exported.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host": schema.StringAttribute{
						Description:         "The name of a service from the service registry.",
						MarkdownDescription: "The name of a service from the service registry.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"subsets": schema.ListNestedAttribute{
						Description:         "One or more named sets that represent individual versions of a service.",
						MarkdownDescription: "One or more named sets that represent individual versions of a service.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"labels": schema.MapAttribute{
									Description:         "Labels apply a filter over the endpoints of a service in the service registry.",
									MarkdownDescription: "Labels apply a filter over the endpoints of a service in the service registry.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the subset.",
									MarkdownDescription: "Name of the subset.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"traffic_policy": schema.SingleNestedAttribute{
									Description:         "Traffic policies that apply to this subset.",
									MarkdownDescription: "Traffic policies that apply to this subset.",
									Attributes: map[string]schema.Attribute{
										"connection_pool": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"http": schema.SingleNestedAttribute{
													Description:         "HTTP connection pool settings.",
													MarkdownDescription: "HTTP connection pool settings.",
													Attributes: map[string]schema.Attribute{
														"h2_upgrade_policy": schema.StringAttribute{
															Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
															MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
															},
														},

														"http1_max_pending_requests": schema.Int64Attribute{
															Description:         "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
															MarkdownDescription: "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http2_max_requests": schema.Int64Attribute{
															Description:         "Maximum number of active requests to a destination.",
															MarkdownDescription: "Maximum number of active requests to a destination.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"idle_timeout": schema.StringAttribute{
															Description:         "The idle timeout for upstream connection pool connections.",
															MarkdownDescription: "The idle timeout for upstream connection pool connections.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_concurrent_streams": schema.Int64Attribute{
															Description:         "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
															MarkdownDescription: "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_requests_per_connection": schema.Int64Attribute{
															Description:         "Maximum number of requests per connection to a backend.",
															MarkdownDescription: "Maximum number of requests per connection to a backend.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_retries": schema.Int64Attribute{
															Description:         "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
															MarkdownDescription: "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"use_client_protocol": schema.BoolAttribute{
															Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
															MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"tcp": schema.SingleNestedAttribute{
													Description:         "Settings common to both HTTP and TCP upstream connections.",
													MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",
													Attributes: map[string]schema.Attribute{
														"connect_timeout": schema.StringAttribute{
															Description:         "TCP connection timeout.",
															MarkdownDescription: "TCP connection timeout.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"idle_timeout": schema.StringAttribute{
															Description:         "The idle timeout for TCP connections.",
															MarkdownDescription: "The idle timeout for TCP connections.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_connection_duration": schema.StringAttribute{
															Description:         "The maximum duration of a connection.",
															MarkdownDescription: "The maximum duration of a connection.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_connections": schema.Int64Attribute{
															Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
															MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tcp_keepalive": schema.SingleNestedAttribute{
															Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
															MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
															Attributes: map[string]schema.Attribute{
																"interval": schema.StringAttribute{
																	Description:         "The time duration between keep-alive probes.",
																	MarkdownDescription: "The time duration between keep-alive probes.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"probes": schema.Int64Attribute{
																	Description:         "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
																	MarkdownDescription: "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"time": schema.StringAttribute{
																	Description:         "The time duration a connection needs to be idle before keep-alive probes start being sent.",
																	MarkdownDescription: "The time duration a connection needs to be idle before keep-alive probes start being sent.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"load_balancer": schema.SingleNestedAttribute{
											Description:         "Settings controlling the load balancer algorithms.",
											MarkdownDescription: "Settings controlling the load balancer algorithms.",
											Attributes: map[string]schema.Attribute{
												"consistent_hash": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"http_cookie": schema.SingleNestedAttribute{
															Description:         "Hash based on HTTP cookie.",
															MarkdownDescription: "Hash based on HTTP cookie.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the cookie.",
																	MarkdownDescription: "Name of the cookie.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "Path to set for the cookie.",
																	MarkdownDescription: "Path to set for the cookie.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"ttl": schema.StringAttribute{
																	Description:         "Lifetime of the cookie.",
																	MarkdownDescription: "Lifetime of the cookie.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"http_header_name": schema.StringAttribute{
															Description:         "Hash based on a specific HTTP header.",
															MarkdownDescription: "Hash based on a specific HTTP header.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_query_parameter_name": schema.StringAttribute{
															Description:         "Hash based on a specific HTTP query parameter.",
															MarkdownDescription: "Hash based on a specific HTTP query parameter.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"maglev": schema.SingleNestedAttribute{
															Description:         "The Maglev load balancer implements consistent hashing to backend hosts.",
															MarkdownDescription: "The Maglev load balancer implements consistent hashing to backend hosts.",
															Attributes: map[string]schema.Attribute{
																"table_size": schema.Int64Attribute{
																	Description:         "The table size for Maglev hashing.",
																	MarkdownDescription: "The table size for Maglev hashing.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"minimum_ring_size": schema.Int64Attribute{
															Description:         "Deprecated.",
															MarkdownDescription: "Deprecated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ring_hash": schema.SingleNestedAttribute{
															Description:         "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
															MarkdownDescription: "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
															Attributes: map[string]schema.Attribute{
																"minimum_ring_size": schema.Int64Attribute{
																	Description:         "The minimum number of virtual nodes to use for the hash ring.",
																	MarkdownDescription: "The minimum number of virtual nodes to use for the hash ring.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"use_source_ip": schema.BoolAttribute{
															Description:         "Hash based on the source IP address.",
															MarkdownDescription: "Hash based on the source IP address.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"locality_lb_setting": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"distribute": schema.ListNestedAttribute{
															Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
															MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"from": schema.StringAttribute{
																		Description:         "Originating locality, '/' separated, e.g.",
																		MarkdownDescription: "Originating locality, '/' separated, e.g.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"to": schema.MapAttribute{
																		Description:         "Map of upstream localities to traffic distribution weights.",
																		MarkdownDescription: "Map of upstream localities to traffic distribution weights.",
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

														"enabled": schema.BoolAttribute{
															Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
															MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"failover": schema.ListNestedAttribute{
															Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
															MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"from": schema.StringAttribute{
																		Description:         "Originating region.",
																		MarkdownDescription: "Originating region.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"to": schema.StringAttribute{
																		Description:         "Destination region the traffic will fail over to when endpoints in the 'from' region becomes unhealthy.",
																		MarkdownDescription: "Destination region the traffic will fail over to when endpoints in the 'from' region becomes unhealthy.",
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

														"failover_priority": schema.ListAttribute{
															Description:         "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
															MarkdownDescription: "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
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

												"simple": schema.StringAttribute{
													Description:         "Valid Options: LEAST_CONN, RANDOM, PASSTHROUGH, ROUND_ROBIN, LEAST_REQUEST",
													MarkdownDescription: "Valid Options: LEAST_CONN, RANDOM, PASSTHROUGH, ROUND_ROBIN, LEAST_REQUEST",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("UNSPECIFIED", "LEAST_CONN", "RANDOM", "PASSTHROUGH", "ROUND_ROBIN", "LEAST_REQUEST"),
													},
												},

												"warmup_duration_secs": schema.StringAttribute{
													Description:         "Represents the warmup duration of Service.",
													MarkdownDescription: "Represents the warmup duration of Service.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"outlier_detection": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"base_ejection_time": schema.StringAttribute{
													Description:         "Minimum ejection duration.",
													MarkdownDescription: "Minimum ejection duration.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"consecutive5xx_errors": schema.Int64Attribute{
													Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
													MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"consecutive_errors": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"consecutive_gateway_errors": schema.Int64Attribute{
													Description:         "Number of gateway errors before a host is ejected from the connection pool.",
													MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"consecutive_local_origin_failures": schema.Int64Attribute{
													Description:         "The number of consecutive locally originated failures before ejection occurs.",
													MarkdownDescription: "The number of consecutive locally originated failures before ejection occurs.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "Time interval between ejection sweep analysis.",
													MarkdownDescription: "Time interval between ejection sweep analysis.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_ejection_percent": schema.Int64Attribute{
													Description:         "Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.",
													MarkdownDescription: "Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_health_percent": schema.Int64Attribute{
													Description:         "Outlier detection will be enabled as long as the associated load balancing pool has at least min_health_percent hosts in healthy mode.",
													MarkdownDescription: "Outlier detection will be enabled as long as the associated load balancing pool has at least min_health_percent hosts in healthy mode.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"split_external_local_origin_errors": schema.BoolAttribute{
													Description:         "Determines whether to distinguish local origin failures from external errors.",
													MarkdownDescription: "Determines whether to distinguish local origin failures from external errors.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"port_level_settings": schema.ListNestedAttribute{
											Description:         "Traffic policies specific to individual ports.",
											MarkdownDescription: "Traffic policies specific to individual ports.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"connection_pool": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"http": schema.SingleNestedAttribute{
																Description:         "HTTP connection pool settings.",
																MarkdownDescription: "HTTP connection pool settings.",
																Attributes: map[string]schema.Attribute{
																	"h2_upgrade_policy": schema.StringAttribute{
																		Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
																		MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
																		},
																	},

																	"http1_max_pending_requests": schema.Int64Attribute{
																		Description:         "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
																		MarkdownDescription: "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"http2_max_requests": schema.Int64Attribute{
																		Description:         "Maximum number of active requests to a destination.",
																		MarkdownDescription: "Maximum number of active requests to a destination.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"idle_timeout": schema.StringAttribute{
																		Description:         "The idle timeout for upstream connection pool connections.",
																		MarkdownDescription: "The idle timeout for upstream connection pool connections.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"max_concurrent_streams": schema.Int64Attribute{
																		Description:         "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
																		MarkdownDescription: "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"max_requests_per_connection": schema.Int64Attribute{
																		Description:         "Maximum number of requests per connection to a backend.",
																		MarkdownDescription: "Maximum number of requests per connection to a backend.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"max_retries": schema.Int64Attribute{
																		Description:         "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
																		MarkdownDescription: "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"use_client_protocol": schema.BoolAttribute{
																		Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
																		MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp": schema.SingleNestedAttribute{
																Description:         "Settings common to both HTTP and TCP upstream connections.",
																MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",
																Attributes: map[string]schema.Attribute{
																	"connect_timeout": schema.StringAttribute{
																		Description:         "TCP connection timeout.",
																		MarkdownDescription: "TCP connection timeout.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"idle_timeout": schema.StringAttribute{
																		Description:         "The idle timeout for TCP connections.",
																		MarkdownDescription: "The idle timeout for TCP connections.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"max_connection_duration": schema.StringAttribute{
																		Description:         "The maximum duration of a connection.",
																		MarkdownDescription: "The maximum duration of a connection.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"max_connections": schema.Int64Attribute{
																		Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
																		MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"tcp_keepalive": schema.SingleNestedAttribute{
																		Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
																		MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
																		Attributes: map[string]schema.Attribute{
																			"interval": schema.StringAttribute{
																				Description:         "The time duration between keep-alive probes.",
																				MarkdownDescription: "The time duration between keep-alive probes.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"probes": schema.Int64Attribute{
																				Description:         "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
																				MarkdownDescription: "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"time": schema.StringAttribute{
																				Description:         "The time duration a connection needs to be idle before keep-alive probes start being sent.",
																				MarkdownDescription: "The time duration a connection needs to be idle before keep-alive probes start being sent.",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"load_balancer": schema.SingleNestedAttribute{
														Description:         "Settings controlling the load balancer algorithms.",
														MarkdownDescription: "Settings controlling the load balancer algorithms.",
														Attributes: map[string]schema.Attribute{
															"consistent_hash": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"http_cookie": schema.SingleNestedAttribute{
																		Description:         "Hash based on HTTP cookie.",
																		MarkdownDescription: "Hash based on HTTP cookie.",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "Name of the cookie.",
																				MarkdownDescription: "Name of the cookie.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"path": schema.StringAttribute{
																				Description:         "Path to set for the cookie.",
																				MarkdownDescription: "Path to set for the cookie.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"ttl": schema.StringAttribute{
																				Description:         "Lifetime of the cookie.",
																				MarkdownDescription: "Lifetime of the cookie.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"http_header_name": schema.StringAttribute{
																		Description:         "Hash based on a specific HTTP header.",
																		MarkdownDescription: "Hash based on a specific HTTP header.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"http_query_parameter_name": schema.StringAttribute{
																		Description:         "Hash based on a specific HTTP query parameter.",
																		MarkdownDescription: "Hash based on a specific HTTP query parameter.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"maglev": schema.SingleNestedAttribute{
																		Description:         "The Maglev load balancer implements consistent hashing to backend hosts.",
																		MarkdownDescription: "The Maglev load balancer implements consistent hashing to backend hosts.",
																		Attributes: map[string]schema.Attribute{
																			"table_size": schema.Int64Attribute{
																				Description:         "The table size for Maglev hashing.",
																				MarkdownDescription: "The table size for Maglev hashing.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"minimum_ring_size": schema.Int64Attribute{
																		Description:         "Deprecated.",
																		MarkdownDescription: "Deprecated.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"ring_hash": schema.SingleNestedAttribute{
																		Description:         "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
																		MarkdownDescription: "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
																		Attributes: map[string]schema.Attribute{
																			"minimum_ring_size": schema.Int64Attribute{
																				Description:         "The minimum number of virtual nodes to use for the hash ring.",
																				MarkdownDescription: "The minimum number of virtual nodes to use for the hash ring.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"use_source_ip": schema.BoolAttribute{
																		Description:         "Hash based on the source IP address.",
																		MarkdownDescription: "Hash based on the source IP address.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"locality_lb_setting": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"distribute": schema.ListNestedAttribute{
																		Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
																		MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"from": schema.StringAttribute{
																					Description:         "Originating locality, '/' separated, e.g.",
																					MarkdownDescription: "Originating locality, '/' separated, e.g.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"to": schema.MapAttribute{
																					Description:         "Map of upstream localities to traffic distribution weights.",
																					MarkdownDescription: "Map of upstream localities to traffic distribution weights.",
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

																	"enabled": schema.BoolAttribute{
																		Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
																		MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"failover": schema.ListNestedAttribute{
																		Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
																		MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"from": schema.StringAttribute{
																					Description:         "Originating region.",
																					MarkdownDescription: "Originating region.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"to": schema.StringAttribute{
																					Description:         "Destination region the traffic will fail over to when endpoints in the 'from' region becomes unhealthy.",
																					MarkdownDescription: "Destination region the traffic will fail over to when endpoints in the 'from' region becomes unhealthy.",
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

																	"failover_priority": schema.ListAttribute{
																		Description:         "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
																		MarkdownDescription: "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
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

															"simple": schema.StringAttribute{
																Description:         "Valid Options: LEAST_CONN, RANDOM, PASSTHROUGH, ROUND_ROBIN, LEAST_REQUEST",
																MarkdownDescription: "Valid Options: LEAST_CONN, RANDOM, PASSTHROUGH, ROUND_ROBIN, LEAST_REQUEST",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("UNSPECIFIED", "LEAST_CONN", "RANDOM", "PASSTHROUGH", "ROUND_ROBIN", "LEAST_REQUEST"),
																},
															},

															"warmup_duration_secs": schema.StringAttribute{
																Description:         "Represents the warmup duration of Service.",
																MarkdownDescription: "Represents the warmup duration of Service.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"outlier_detection": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"base_ejection_time": schema.StringAttribute{
																Description:         "Minimum ejection duration.",
																MarkdownDescription: "Minimum ejection duration.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"consecutive5xx_errors": schema.Int64Attribute{
																Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
																MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"consecutive_errors": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"consecutive_gateway_errors": schema.Int64Attribute{
																Description:         "Number of gateway errors before a host is ejected from the connection pool.",
																MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"consecutive_local_origin_failures": schema.Int64Attribute{
																Description:         "The number of consecutive locally originated failures before ejection occurs.",
																MarkdownDescription: "The number of consecutive locally originated failures before ejection occurs.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"interval": schema.StringAttribute{
																Description:         "Time interval between ejection sweep analysis.",
																MarkdownDescription: "Time interval between ejection sweep analysis.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"max_ejection_percent": schema.Int64Attribute{
																Description:         "Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.",
																MarkdownDescription: "Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"min_health_percent": schema.Int64Attribute{
																Description:         "Outlier detection will be enabled as long as the associated load balancing pool has at least min_health_percent hosts in healthy mode.",
																MarkdownDescription: "Outlier detection will be enabled as long as the associated load balancing pool has at least min_health_percent hosts in healthy mode.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"split_external_local_origin_errors": schema.BoolAttribute{
																Description:         "Determines whether to distinguish local origin failures from external errors.",
																MarkdownDescription: "Determines whether to distinguish local origin failures from external errors.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"port": schema.SingleNestedAttribute{
														Description:         "Specifies the number of a port on the destination service on which this policy is being applied.",
														MarkdownDescription: "Specifies the number of a port on the destination service on which this policy is being applied.",
														Attributes: map[string]schema.Attribute{
															"number": schema.Int64Attribute{
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

													"tls": schema.SingleNestedAttribute{
														Description:         "TLS related settings for connections to the upstream service.",
														MarkdownDescription: "TLS related settings for connections to the upstream service.",
														Attributes: map[string]schema.Attribute{
															"ca_certificates": schema.StringAttribute{
																Description:         "OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate.",
																MarkdownDescription: "OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"ca_crl": schema.StringAttribute{
																Description:         "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented server certificate.",
																MarkdownDescription: "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented server certificate.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"client_certificate": schema.StringAttribute{
																Description:         "REQUIRED if mode is 'MUTUAL'.",
																MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"credential_name": schema.StringAttribute{
																Description:         "The name of the secret that holds the TLS certs for the client including the CA certificates.",
																MarkdownDescription: "The name of the secret that holds the TLS certs for the client including the CA certificates.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "'insecureSkipVerify' specifies whether the proxy should skip verifying the CA signature and SAN for the server certificate corresponding to the host.",
																MarkdownDescription: "'insecureSkipVerify' specifies whether the proxy should skip verifying the CA signature and SAN for the server certificate corresponding to the host.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"mode": schema.StringAttribute{
																Description:         "Indicates whether connections to this port should be secured using TLS.Valid Options: DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL",
																MarkdownDescription: "Indicates whether connections to this port should be secured using TLS.Valid Options: DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
																},
															},

															"private_key": schema.StringAttribute{
																Description:         "REQUIRED if mode is 'MUTUAL'.",
																MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sni": schema.StringAttribute{
																Description:         "SNI string to present to the server during TLS handshake.",
																MarkdownDescription: "SNI string to present to the server during TLS handshake.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"subject_alt_names": schema.ListAttribute{
																Description:         "A list of alternate names to verify the subject identity in the certificate.",
																MarkdownDescription: "A list of alternate names to verify the subject identity in the certificate.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"proxy_protocol": schema.SingleNestedAttribute{
											Description:         "The upstream PROXY protocol settings.",
											MarkdownDescription: "The upstream PROXY protocol settings.",
											Attributes: map[string]schema.Attribute{
												"version": schema.StringAttribute{
													Description:         "The PROXY protocol version to use.Valid Options: V1, V2",
													MarkdownDescription: "The PROXY protocol version to use.Valid Options: V1, V2",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("V1", "V2"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls": schema.SingleNestedAttribute{
											Description:         "TLS related settings for connections to the upstream service.",
											MarkdownDescription: "TLS related settings for connections to the upstream service.",
											Attributes: map[string]schema.Attribute{
												"ca_certificates": schema.StringAttribute{
													Description:         "OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate.",
													MarkdownDescription: "OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_crl": schema.StringAttribute{
													Description:         "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented server certificate.",
													MarkdownDescription: "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented server certificate.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_certificate": schema.StringAttribute{
													Description:         "REQUIRED if mode is 'MUTUAL'.",
													MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"credential_name": schema.StringAttribute{
													Description:         "The name of the secret that holds the TLS certs for the client including the CA certificates.",
													MarkdownDescription: "The name of the secret that holds the TLS certs for the client including the CA certificates.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "'insecureSkipVerify' specifies whether the proxy should skip verifying the CA signature and SAN for the server certificate corresponding to the host.",
													MarkdownDescription: "'insecureSkipVerify' specifies whether the proxy should skip verifying the CA signature and SAN for the server certificate corresponding to the host.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Indicates whether connections to this port should be secured using TLS.Valid Options: DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL",
													MarkdownDescription: "Indicates whether connections to this port should be secured using TLS.Valid Options: DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
													},
												},

												"private_key": schema.StringAttribute{
													Description:         "REQUIRED if mode is 'MUTUAL'.",
													MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sni": schema.StringAttribute{
													Description:         "SNI string to present to the server during TLS handshake.",
													MarkdownDescription: "SNI string to present to the server during TLS handshake.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"subject_alt_names": schema.ListAttribute{
													Description:         "A list of alternate names to verify the subject identity in the certificate.",
													MarkdownDescription: "A list of alternate names to verify the subject identity in the certificate.",
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

										"tunnel": schema.SingleNestedAttribute{
											Description:         "Configuration of tunneling TCP over other transport or application layers for the host configured in the DestinationRule.",
											MarkdownDescription: "Configuration of tunneling TCP over other transport or application layers for the host configured in the DestinationRule.",
											Attributes: map[string]schema.Attribute{
												"protocol": schema.StringAttribute{
													Description:         "Specifies which protocol to use for tunneling the downstream connection.",
													MarkdownDescription: "Specifies which protocol to use for tunneling the downstream connection.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"target_host": schema.StringAttribute{
													Description:         "Specifies a host to which the downstream connection is tunneled.",
													MarkdownDescription: "Specifies a host to which the downstream connection is tunneled.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"target_port": schema.Int64Attribute{
													Description:         "Specifies a port to which the downstream connection is tunneled.",
													MarkdownDescription: "Specifies a port to which the downstream connection is tunneled.",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"traffic_policy": schema.SingleNestedAttribute{
						Description:         "Traffic policies to apply (load balancing policy, connection pool sizes, outlier detection).",
						MarkdownDescription: "Traffic policies to apply (load balancing policy, connection pool sizes, outlier detection).",
						Attributes: map[string]schema.Attribute{
							"connection_pool": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"http": schema.SingleNestedAttribute{
										Description:         "HTTP connection pool settings.",
										MarkdownDescription: "HTTP connection pool settings.",
										Attributes: map[string]schema.Attribute{
											"h2_upgrade_policy": schema.StringAttribute{
												Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
												MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
												},
											},

											"http1_max_pending_requests": schema.Int64Attribute{
												Description:         "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
												MarkdownDescription: "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http2_max_requests": schema.Int64Attribute{
												Description:         "Maximum number of active requests to a destination.",
												MarkdownDescription: "Maximum number of active requests to a destination.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"idle_timeout": schema.StringAttribute{
												Description:         "The idle timeout for upstream connection pool connections.",
												MarkdownDescription: "The idle timeout for upstream connection pool connections.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_concurrent_streams": schema.Int64Attribute{
												Description:         "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
												MarkdownDescription: "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_requests_per_connection": schema.Int64Attribute{
												Description:         "Maximum number of requests per connection to a backend.",
												MarkdownDescription: "Maximum number of requests per connection to a backend.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_retries": schema.Int64Attribute{
												Description:         "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
												MarkdownDescription: "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"use_client_protocol": schema.BoolAttribute{
												Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
												MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tcp": schema.SingleNestedAttribute{
										Description:         "Settings common to both HTTP and TCP upstream connections.",
										MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",
										Attributes: map[string]schema.Attribute{
											"connect_timeout": schema.StringAttribute{
												Description:         "TCP connection timeout.",
												MarkdownDescription: "TCP connection timeout.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"idle_timeout": schema.StringAttribute{
												Description:         "The idle timeout for TCP connections.",
												MarkdownDescription: "The idle timeout for TCP connections.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_connection_duration": schema.StringAttribute{
												Description:         "The maximum duration of a connection.",
												MarkdownDescription: "The maximum duration of a connection.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_connections": schema.Int64Attribute{
												Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
												MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tcp_keepalive": schema.SingleNestedAttribute{
												Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
												MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
												Attributes: map[string]schema.Attribute{
													"interval": schema.StringAttribute{
														Description:         "The time duration between keep-alive probes.",
														MarkdownDescription: "The time duration between keep-alive probes.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"probes": schema.Int64Attribute{
														Description:         "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
														MarkdownDescription: "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"time": schema.StringAttribute{
														Description:         "The time duration a connection needs to be idle before keep-alive probes start being sent.",
														MarkdownDescription: "The time duration a connection needs to be idle before keep-alive probes start being sent.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"load_balancer": schema.SingleNestedAttribute{
								Description:         "Settings controlling the load balancer algorithms.",
								MarkdownDescription: "Settings controlling the load balancer algorithms.",
								Attributes: map[string]schema.Attribute{
									"consistent_hash": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"http_cookie": schema.SingleNestedAttribute{
												Description:         "Hash based on HTTP cookie.",
												MarkdownDescription: "Hash based on HTTP cookie.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the cookie.",
														MarkdownDescription: "Name of the cookie.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"path": schema.StringAttribute{
														Description:         "Path to set for the cookie.",
														MarkdownDescription: "Path to set for the cookie.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ttl": schema.StringAttribute{
														Description:         "Lifetime of the cookie.",
														MarkdownDescription: "Lifetime of the cookie.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_header_name": schema.StringAttribute{
												Description:         "Hash based on a specific HTTP header.",
												MarkdownDescription: "Hash based on a specific HTTP header.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_query_parameter_name": schema.StringAttribute{
												Description:         "Hash based on a specific HTTP query parameter.",
												MarkdownDescription: "Hash based on a specific HTTP query parameter.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"maglev": schema.SingleNestedAttribute{
												Description:         "The Maglev load balancer implements consistent hashing to backend hosts.",
												MarkdownDescription: "The Maglev load balancer implements consistent hashing to backend hosts.",
												Attributes: map[string]schema.Attribute{
													"table_size": schema.Int64Attribute{
														Description:         "The table size for Maglev hashing.",
														MarkdownDescription: "The table size for Maglev hashing.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum_ring_size": schema.Int64Attribute{
												Description:         "Deprecated.",
												MarkdownDescription: "Deprecated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ring_hash": schema.SingleNestedAttribute{
												Description:         "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
												MarkdownDescription: "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
												Attributes: map[string]schema.Attribute{
													"minimum_ring_size": schema.Int64Attribute{
														Description:         "The minimum number of virtual nodes to use for the hash ring.",
														MarkdownDescription: "The minimum number of virtual nodes to use for the hash ring.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"use_source_ip": schema.BoolAttribute{
												Description:         "Hash based on the source IP address.",
												MarkdownDescription: "Hash based on the source IP address.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"locality_lb_setting": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"distribute": schema.ListNestedAttribute{
												Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
												MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"from": schema.StringAttribute{
															Description:         "Originating locality, '/' separated, e.g.",
															MarkdownDescription: "Originating locality, '/' separated, e.g.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"to": schema.MapAttribute{
															Description:         "Map of upstream localities to traffic distribution weights.",
															MarkdownDescription: "Map of upstream localities to traffic distribution weights.",
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

											"enabled": schema.BoolAttribute{
												Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
												MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"failover": schema.ListNestedAttribute{
												Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
												MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"from": schema.StringAttribute{
															Description:         "Originating region.",
															MarkdownDescription: "Originating region.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"to": schema.StringAttribute{
															Description:         "Destination region the traffic will fail over to when endpoints in the 'from' region becomes unhealthy.",
															MarkdownDescription: "Destination region the traffic will fail over to when endpoints in the 'from' region becomes unhealthy.",
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

											"failover_priority": schema.ListAttribute{
												Description:         "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
												MarkdownDescription: "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
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

									"simple": schema.StringAttribute{
										Description:         "Valid Options: LEAST_CONN, RANDOM, PASSTHROUGH, ROUND_ROBIN, LEAST_REQUEST",
										MarkdownDescription: "Valid Options: LEAST_CONN, RANDOM, PASSTHROUGH, ROUND_ROBIN, LEAST_REQUEST",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("UNSPECIFIED", "LEAST_CONN", "RANDOM", "PASSTHROUGH", "ROUND_ROBIN", "LEAST_REQUEST"),
										},
									},

									"warmup_duration_secs": schema.StringAttribute{
										Description:         "Represents the warmup duration of Service.",
										MarkdownDescription: "Represents the warmup duration of Service.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"outlier_detection": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"base_ejection_time": schema.StringAttribute{
										Description:         "Minimum ejection duration.",
										MarkdownDescription: "Minimum ejection duration.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"consecutive5xx_errors": schema.Int64Attribute{
										Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
										MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"consecutive_errors": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"consecutive_gateway_errors": schema.Int64Attribute{
										Description:         "Number of gateway errors before a host is ejected from the connection pool.",
										MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"consecutive_local_origin_failures": schema.Int64Attribute{
										Description:         "The number of consecutive locally originated failures before ejection occurs.",
										MarkdownDescription: "The number of consecutive locally originated failures before ejection occurs.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"interval": schema.StringAttribute{
										Description:         "Time interval between ejection sweep analysis.",
										MarkdownDescription: "Time interval between ejection sweep analysis.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_ejection_percent": schema.Int64Attribute{
										Description:         "Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.",
										MarkdownDescription: "Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_health_percent": schema.Int64Attribute{
										Description:         "Outlier detection will be enabled as long as the associated load balancing pool has at least min_health_percent hosts in healthy mode.",
										MarkdownDescription: "Outlier detection will be enabled as long as the associated load balancing pool has at least min_health_percent hosts in healthy mode.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"split_external_local_origin_errors": schema.BoolAttribute{
										Description:         "Determines whether to distinguish local origin failures from external errors.",
										MarkdownDescription: "Determines whether to distinguish local origin failures from external errors.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"port_level_settings": schema.ListNestedAttribute{
								Description:         "Traffic policies specific to individual ports.",
								MarkdownDescription: "Traffic policies specific to individual ports.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"connection_pool": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"http": schema.SingleNestedAttribute{
													Description:         "HTTP connection pool settings.",
													MarkdownDescription: "HTTP connection pool settings.",
													Attributes: map[string]schema.Attribute{
														"h2_upgrade_policy": schema.StringAttribute{
															Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
															MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.Valid Options: DEFAULT, DO_NOT_UPGRADE, UPGRADE",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
															},
														},

														"http1_max_pending_requests": schema.Int64Attribute{
															Description:         "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
															MarkdownDescription: "Maximum number of requests that will be queued while waiting for a ready connection pool connection.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http2_max_requests": schema.Int64Attribute{
															Description:         "Maximum number of active requests to a destination.",
															MarkdownDescription: "Maximum number of active requests to a destination.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"idle_timeout": schema.StringAttribute{
															Description:         "The idle timeout for upstream connection pool connections.",
															MarkdownDescription: "The idle timeout for upstream connection pool connections.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_concurrent_streams": schema.Int64Attribute{
															Description:         "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
															MarkdownDescription: "The maximum number of concurrent streams allowed for a peer on one HTTP/2 connection.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_requests_per_connection": schema.Int64Attribute{
															Description:         "Maximum number of requests per connection to a backend.",
															MarkdownDescription: "Maximum number of requests per connection to a backend.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_retries": schema.Int64Attribute{
															Description:         "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
															MarkdownDescription: "Maximum number of retries that can be outstanding to all hosts in a cluster at a given time.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"use_client_protocol": schema.BoolAttribute{
															Description:         "If set to true, client protocol will be preserved while initiating connection to backend.",
															MarkdownDescription: "If set to true, client protocol will be preserved while initiating connection to backend.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"tcp": schema.SingleNestedAttribute{
													Description:         "Settings common to both HTTP and TCP upstream connections.",
													MarkdownDescription: "Settings common to both HTTP and TCP upstream connections.",
													Attributes: map[string]schema.Attribute{
														"connect_timeout": schema.StringAttribute{
															Description:         "TCP connection timeout.",
															MarkdownDescription: "TCP connection timeout.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"idle_timeout": schema.StringAttribute{
															Description:         "The idle timeout for TCP connections.",
															MarkdownDescription: "The idle timeout for TCP connections.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_connection_duration": schema.StringAttribute{
															Description:         "The maximum duration of a connection.",
															MarkdownDescription: "The maximum duration of a connection.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_connections": schema.Int64Attribute{
															Description:         "Maximum number of HTTP1 /TCP connections to a destination host.",
															MarkdownDescription: "Maximum number of HTTP1 /TCP connections to a destination host.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tcp_keepalive": schema.SingleNestedAttribute{
															Description:         "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
															MarkdownDescription: "If set then set SO_KEEPALIVE on the socket to enable TCP Keepalives.",
															Attributes: map[string]schema.Attribute{
																"interval": schema.StringAttribute{
																	Description:         "The time duration between keep-alive probes.",
																	MarkdownDescription: "The time duration between keep-alive probes.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"probes": schema.Int64Attribute{
																	Description:         "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
																	MarkdownDescription: "Maximum number of keepalive probes to send without response before deciding the connection is dead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"time": schema.StringAttribute{
																	Description:         "The time duration a connection needs to be idle before keep-alive probes start being sent.",
																	MarkdownDescription: "The time duration a connection needs to be idle before keep-alive probes start being sent.",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"load_balancer": schema.SingleNestedAttribute{
											Description:         "Settings controlling the load balancer algorithms.",
											MarkdownDescription: "Settings controlling the load balancer algorithms.",
											Attributes: map[string]schema.Attribute{
												"consistent_hash": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"http_cookie": schema.SingleNestedAttribute{
															Description:         "Hash based on HTTP cookie.",
															MarkdownDescription: "Hash based on HTTP cookie.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the cookie.",
																	MarkdownDescription: "Name of the cookie.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "Path to set for the cookie.",
																	MarkdownDescription: "Path to set for the cookie.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"ttl": schema.StringAttribute{
																	Description:         "Lifetime of the cookie.",
																	MarkdownDescription: "Lifetime of the cookie.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"http_header_name": schema.StringAttribute{
															Description:         "Hash based on a specific HTTP header.",
															MarkdownDescription: "Hash based on a specific HTTP header.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_query_parameter_name": schema.StringAttribute{
															Description:         "Hash based on a specific HTTP query parameter.",
															MarkdownDescription: "Hash based on a specific HTTP query parameter.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"maglev": schema.SingleNestedAttribute{
															Description:         "The Maglev load balancer implements consistent hashing to backend hosts.",
															MarkdownDescription: "The Maglev load balancer implements consistent hashing to backend hosts.",
															Attributes: map[string]schema.Attribute{
																"table_size": schema.Int64Attribute{
																	Description:         "The table size for Maglev hashing.",
																	MarkdownDescription: "The table size for Maglev hashing.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"minimum_ring_size": schema.Int64Attribute{
															Description:         "Deprecated.",
															MarkdownDescription: "Deprecated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"ring_hash": schema.SingleNestedAttribute{
															Description:         "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
															MarkdownDescription: "The ring/modulo hash load balancer implements consistent hashing to backend hosts.",
															Attributes: map[string]schema.Attribute{
																"minimum_ring_size": schema.Int64Attribute{
																	Description:         "The minimum number of virtual nodes to use for the hash ring.",
																	MarkdownDescription: "The minimum number of virtual nodes to use for the hash ring.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"use_source_ip": schema.BoolAttribute{
															Description:         "Hash based on the source IP address.",
															MarkdownDescription: "Hash based on the source IP address.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"locality_lb_setting": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"distribute": schema.ListNestedAttribute{
															Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
															MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"from": schema.StringAttribute{
																		Description:         "Originating locality, '/' separated, e.g.",
																		MarkdownDescription: "Originating locality, '/' separated, e.g.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"to": schema.MapAttribute{
																		Description:         "Map of upstream localities to traffic distribution weights.",
																		MarkdownDescription: "Map of upstream localities to traffic distribution weights.",
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

														"enabled": schema.BoolAttribute{
															Description:         "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
															MarkdownDescription: "enable locality load balancing, this is DestinationRule-level and will override mesh wide settings in entirety.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"failover": schema.ListNestedAttribute{
															Description:         "Optional: only one of distribute, failover or failoverPriority can be set.",
															MarkdownDescription: "Optional: only one of distribute, failover or failoverPriority can be set.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"from": schema.StringAttribute{
																		Description:         "Originating region.",
																		MarkdownDescription: "Originating region.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"to": schema.StringAttribute{
																		Description:         "Destination region the traffic will fail over to when endpoints in the 'from' region becomes unhealthy.",
																		MarkdownDescription: "Destination region the traffic will fail over to when endpoints in the 'from' region becomes unhealthy.",
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

														"failover_priority": schema.ListAttribute{
															Description:         "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
															MarkdownDescription: "failoverPriority is an ordered list of labels used to sort endpoints to do priority based load balancing.",
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

												"simple": schema.StringAttribute{
													Description:         "Valid Options: LEAST_CONN, RANDOM, PASSTHROUGH, ROUND_ROBIN, LEAST_REQUEST",
													MarkdownDescription: "Valid Options: LEAST_CONN, RANDOM, PASSTHROUGH, ROUND_ROBIN, LEAST_REQUEST",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("UNSPECIFIED", "LEAST_CONN", "RANDOM", "PASSTHROUGH", "ROUND_ROBIN", "LEAST_REQUEST"),
													},
												},

												"warmup_duration_secs": schema.StringAttribute{
													Description:         "Represents the warmup duration of Service.",
													MarkdownDescription: "Represents the warmup duration of Service.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"outlier_detection": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"base_ejection_time": schema.StringAttribute{
													Description:         "Minimum ejection duration.",
													MarkdownDescription: "Minimum ejection duration.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"consecutive5xx_errors": schema.Int64Attribute{
													Description:         "Number of 5xx errors before a host is ejected from the connection pool.",
													MarkdownDescription: "Number of 5xx errors before a host is ejected from the connection pool.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"consecutive_errors": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"consecutive_gateway_errors": schema.Int64Attribute{
													Description:         "Number of gateway errors before a host is ejected from the connection pool.",
													MarkdownDescription: "Number of gateway errors before a host is ejected from the connection pool.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"consecutive_local_origin_failures": schema.Int64Attribute{
													Description:         "The number of consecutive locally originated failures before ejection occurs.",
													MarkdownDescription: "The number of consecutive locally originated failures before ejection occurs.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"interval": schema.StringAttribute{
													Description:         "Time interval between ejection sweep analysis.",
													MarkdownDescription: "Time interval between ejection sweep analysis.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_ejection_percent": schema.Int64Attribute{
													Description:         "Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.",
													MarkdownDescription: "Maximum % of hosts in the load balancing pool for the upstream service that can be ejected.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_health_percent": schema.Int64Attribute{
													Description:         "Outlier detection will be enabled as long as the associated load balancing pool has at least min_health_percent hosts in healthy mode.",
													MarkdownDescription: "Outlier detection will be enabled as long as the associated load balancing pool has at least min_health_percent hosts in healthy mode.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"split_external_local_origin_errors": schema.BoolAttribute{
													Description:         "Determines whether to distinguish local origin failures from external errors.",
													MarkdownDescription: "Determines whether to distinguish local origin failures from external errors.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"port": schema.SingleNestedAttribute{
											Description:         "Specifies the number of a port on the destination service on which this policy is being applied.",
											MarkdownDescription: "Specifies the number of a port on the destination service on which this policy is being applied.",
											Attributes: map[string]schema.Attribute{
												"number": schema.Int64Attribute{
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

										"tls": schema.SingleNestedAttribute{
											Description:         "TLS related settings for connections to the upstream service.",
											MarkdownDescription: "TLS related settings for connections to the upstream service.",
											Attributes: map[string]schema.Attribute{
												"ca_certificates": schema.StringAttribute{
													Description:         "OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate.",
													MarkdownDescription: "OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ca_crl": schema.StringAttribute{
													Description:         "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented server certificate.",
													MarkdownDescription: "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented server certificate.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_certificate": schema.StringAttribute{
													Description:         "REQUIRED if mode is 'MUTUAL'.",
													MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"credential_name": schema.StringAttribute{
													Description:         "The name of the secret that holds the TLS certs for the client including the CA certificates.",
													MarkdownDescription: "The name of the secret that holds the TLS certs for the client including the CA certificates.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "'insecureSkipVerify' specifies whether the proxy should skip verifying the CA signature and SAN for the server certificate corresponding to the host.",
													MarkdownDescription: "'insecureSkipVerify' specifies whether the proxy should skip verifying the CA signature and SAN for the server certificate corresponding to the host.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "Indicates whether connections to this port should be secured using TLS.Valid Options: DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL",
													MarkdownDescription: "Indicates whether connections to this port should be secured using TLS.Valid Options: DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
													},
												},

												"private_key": schema.StringAttribute{
													Description:         "REQUIRED if mode is 'MUTUAL'.",
													MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sni": schema.StringAttribute{
													Description:         "SNI string to present to the server during TLS handshake.",
													MarkdownDescription: "SNI string to present to the server during TLS handshake.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"subject_alt_names": schema.ListAttribute{
													Description:         "A list of alternate names to verify the subject identity in the certificate.",
													MarkdownDescription: "A list of alternate names to verify the subject identity in the certificate.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy_protocol": schema.SingleNestedAttribute{
								Description:         "The upstream PROXY protocol settings.",
								MarkdownDescription: "The upstream PROXY protocol settings.",
								Attributes: map[string]schema.Attribute{
									"version": schema.StringAttribute{
										Description:         "The PROXY protocol version to use.Valid Options: V1, V2",
										MarkdownDescription: "The PROXY protocol version to use.Valid Options: V1, V2",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("V1", "V2"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": schema.SingleNestedAttribute{
								Description:         "TLS related settings for connections to the upstream service.",
								MarkdownDescription: "TLS related settings for connections to the upstream service.",
								Attributes: map[string]schema.Attribute{
									"ca_certificates": schema.StringAttribute{
										Description:         "OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate.",
										MarkdownDescription: "OPTIONAL: The path to the file containing certificate authority certificates to use in verifying a presented server certificate.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_crl": schema.StringAttribute{
										Description:         "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented server certificate.",
										MarkdownDescription: "OPTIONAL: The path to the file containing the certificate revocation list (CRL) to use in verifying a presented server certificate.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_certificate": schema.StringAttribute{
										Description:         "REQUIRED if mode is 'MUTUAL'.",
										MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"credential_name": schema.StringAttribute{
										Description:         "The name of the secret that holds the TLS certs for the client including the CA certificates.",
										MarkdownDescription: "The name of the secret that holds the TLS certs for the client including the CA certificates.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "'insecureSkipVerify' specifies whether the proxy should skip verifying the CA signature and SAN for the server certificate corresponding to the host.",
										MarkdownDescription: "'insecureSkipVerify' specifies whether the proxy should skip verifying the CA signature and SAN for the server certificate corresponding to the host.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mode": schema.StringAttribute{
										Description:         "Indicates whether connections to this port should be secured using TLS.Valid Options: DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL",
										MarkdownDescription: "Indicates whether connections to this port should be secured using TLS.Valid Options: DISABLE, SIMPLE, MUTUAL, ISTIO_MUTUAL",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("DISABLE", "SIMPLE", "MUTUAL", "ISTIO_MUTUAL"),
										},
									},

									"private_key": schema.StringAttribute{
										Description:         "REQUIRED if mode is 'MUTUAL'.",
										MarkdownDescription: "REQUIRED if mode is 'MUTUAL'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sni": schema.StringAttribute{
										Description:         "SNI string to present to the server during TLS handshake.",
										MarkdownDescription: "SNI string to present to the server during TLS handshake.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"subject_alt_names": schema.ListAttribute{
										Description:         "A list of alternate names to verify the subject identity in the certificate.",
										MarkdownDescription: "A list of alternate names to verify the subject identity in the certificate.",
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

							"tunnel": schema.SingleNestedAttribute{
								Description:         "Configuration of tunneling TCP over other transport or application layers for the host configured in the DestinationRule.",
								MarkdownDescription: "Configuration of tunneling TCP over other transport or application layers for the host configured in the DestinationRule.",
								Attributes: map[string]schema.Attribute{
									"protocol": schema.StringAttribute{
										Description:         "Specifies which protocol to use for tunneling the downstream connection.",
										MarkdownDescription: "Specifies which protocol to use for tunneling the downstream connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"target_host": schema.StringAttribute{
										Description:         "Specifies a host to which the downstream connection is tunneled.",
										MarkdownDescription: "Specifies a host to which the downstream connection is tunneled.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"target_port": schema.Int64Attribute{
										Description:         "Specifies a port to which the downstream connection is tunneled.",
										MarkdownDescription: "Specifies a port to which the downstream connection is tunneled.",
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

					"workload_selector": schema.SingleNestedAttribute{
						Description:         "Criteria used to select the specific set of pods/VMs on which this 'DestinationRule' configuration should be applied.",
						MarkdownDescription: "Criteria used to select the specific set of pods/VMs on which this 'DestinationRule' configuration should be applied.",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
								Description:         "One or more labels that indicate a specific set of pods/VMs on which a policy should be applied.",
								MarkdownDescription: "One or more labels that indicate a specific set of pods/VMs on which a policy should be applied.",
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
		},
	}
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_destination_rule_v1alpha3_manifest")

	var model NetworkingIstioIoDestinationRuleV1Alpha3ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("DestinationRule")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
