/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package networking_istio_io_v1alpha3

import (
	"context"
	"encoding/json"
	"fmt"
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
	"strings"
)

var (
	_ resource.Resource                = &NetworkingIstioIoDestinationRuleV1Alpha3Resource{}
	_ resource.ResourceWithConfigure   = &NetworkingIstioIoDestinationRuleV1Alpha3Resource{}
	_ resource.ResourceWithImportState = &NetworkingIstioIoDestinationRuleV1Alpha3Resource{}
)

func NewNetworkingIstioIoDestinationRuleV1Alpha3Resource() resource.Resource {
	return &NetworkingIstioIoDestinationRuleV1Alpha3Resource{}
}

type NetworkingIstioIoDestinationRuleV1Alpha3Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type NetworkingIstioIoDestinationRuleV1Alpha3ResourceData struct {
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
						MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
						MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
						UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Tcp *struct {
						ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
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
							MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
							MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
							UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
						} `tfsdk:"http" json:"http,omitempty"`
						Tcp *struct {
							ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
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
						ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
						CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
						InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
						PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
						Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
						SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"port_level_settings" json:"portLevelSettings,omitempty"`
				Tls *struct {
					CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
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
					MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
					MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Tcp *struct {
					ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
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
						MaxRequestsPerConnection *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
						MaxRetries               *int64  `tfsdk:"max_retries" json:"maxRetries,omitempty"`
						UseClientProtocol        *bool   `tfsdk:"use_client_protocol" json:"useClientProtocol,omitempty"`
					} `tfsdk:"http" json:"http,omitempty"`
					Tcp *struct {
						ConnectTimeout        *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
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
					ClientCertificate  *string   `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
					CredentialName     *string   `tfsdk:"credential_name" json:"credentialName,omitempty"`
					InsecureSkipVerify *bool     `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
					Mode               *string   `tfsdk:"mode" json:"mode,omitempty"`
					PrivateKey         *string   `tfsdk:"private_key" json:"privateKey,omitempty"`
					Sni                *string   `tfsdk:"sni" json:"sni,omitempty"`
					SubjectAltNames    *[]string `tfsdk:"subject_alt_names" json:"subjectAltNames,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"port_level_settings" json:"portLevelSettings,omitempty"`
			Tls *struct {
				CaCertificates     *string   `tfsdk:"ca_certificates" json:"caCertificates,omitempty"`
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

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_networking_istio_io_destination_rule_v1alpha3"
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"subsets": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"labels": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name of the subset.",
									MarkdownDescription: "Name of the subset.",
									Required:            false,
									Optional:            true,
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
															Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
															MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
															},
														},

														"http1_max_pending_requests": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
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

														"max_requests_per_connection": schema.Int64Attribute{
															Description:         "Maximum number of requests per connection to a backend.",
															MarkdownDescription: "Maximum number of requests per connection to a backend.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_retries": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"time": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
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
													Description:         "",
													MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_health_percent": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
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
																		Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
																		MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
																		},
																	},

																	"http1_max_pending_requests": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
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

																	"max_requests_per_connection": schema.Int64Attribute{
																		Description:         "Maximum number of requests per connection to a backend.",
																		MarkdownDescription: "Maximum number of requests per connection to a backend.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"max_retries": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"time": schema.StringAttribute{
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
																				Required:            false,
																				Optional:            true,
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
																Description:         "",
																MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"min_health_percent": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
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
														Description:         "",
														MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"tunnel": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"target_port": schema.Int64Attribute{
													Description:         "Specifies a port to which the downstream connection is tunneled.",
													MarkdownDescription: "Specifies a port to which the downstream connection is tunneled.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"traffic_policy": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
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
												Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
												MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
												},
											},

											"http1_max_pending_requests": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
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

											"max_requests_per_connection": schema.Int64Attribute{
												Description:         "Maximum number of requests per connection to a backend.",
												MarkdownDescription: "Maximum number of requests per connection to a backend.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_retries": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
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
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"time": schema.StringAttribute{
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
														Required:            false,
														Optional:            true,
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
										Description:         "",
										MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_health_percent": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
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
															Description:         "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
															MarkdownDescription: "Specify if http1.1 connection should be upgraded to http2 for the associated destination.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("DEFAULT", "DO_NOT_UPGRADE", "UPGRADE"),
															},
														},

														"http1_max_pending_requests": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
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

														"max_requests_per_connection": schema.Int64Attribute{
															Description:         "Maximum number of requests per connection to a backend.",
															MarkdownDescription: "Maximum number of requests per connection to a backend.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_retries": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
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
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"time": schema.StringAttribute{
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
																	Required:            false,
																	Optional:            true,
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
													Description:         "",
													MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"min_health_percent": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
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
											Description:         "",
											MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"insecure_skip_verify": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
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
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure_skip_verify": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

							"tunnel": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"target_port": schema.Int64Attribute{
										Description:         "Specifies a port to which the downstream connection is tunneled.",
										MarkdownDescription: "Specifies a port to which the downstream connection is tunneled.",
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

					"workload_selector": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"match_labels": schema.MapAttribute{
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
		},
	}
}

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
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

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_networking_istio_io_destination_rule_v1alpha3")

	var model NetworkingIstioIoDestinationRuleV1Alpha3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("DestinationRule")

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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "destinationrules"}).
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

	var readResponse NetworkingIstioIoDestinationRuleV1Alpha3ResourceData
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

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_networking_istio_io_destination_rule_v1alpha3")

	var data NetworkingIstioIoDestinationRuleV1Alpha3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "destinationrules"}).
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

	var readResponse NetworkingIstioIoDestinationRuleV1Alpha3ResourceData
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

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_networking_istio_io_destination_rule_v1alpha3")

	var model NetworkingIstioIoDestinationRuleV1Alpha3ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("networking.istio.io/v1alpha3")
	model.Kind = pointer.String("DestinationRule")

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

	patchResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "destinationrules"}).
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

	var readResponse NetworkingIstioIoDestinationRuleV1Alpha3ResourceData
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

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_networking_istio_io_destination_rule_v1alpha3")

	var data NetworkingIstioIoDestinationRuleV1Alpha3ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "networking.istio.io", Version: "v1alpha3", Resource: "destinationrules"}).
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

func (r *NetworkingIstioIoDestinationRuleV1Alpha3Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
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
