/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type GlooSoloIoUpstreamV1Resource struct{}

var (
	_ resource.Resource = (*GlooSoloIoUpstreamV1Resource)(nil)
)

type GlooSoloIoUpstreamV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type GlooSoloIoUpstreamV1GoModel struct {
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
		Aws *struct {
			AwsAccountId *string `tfsdk:"aws_account_id" yaml:"awsAccountId,omitempty"`

			DisableRoleChaining *bool `tfsdk:"disable_role_chaining" yaml:"disableRoleChaining,omitempty"`

			LambdaFunctions *[]struct {
				LambdaFunctionName *string `tfsdk:"lambda_function_name" yaml:"lambdaFunctionName,omitempty"`

				LogicalName *string `tfsdk:"logical_name" yaml:"logicalName,omitempty"`

				Qualifier *string `tfsdk:"qualifier" yaml:"qualifier,omitempty"`
			} `tfsdk:"lambda_functions" yaml:"lambdaFunctions,omitempty"`

			Region *string `tfsdk:"region" yaml:"region,omitempty"`

			RoleArn *string `tfsdk:"role_arn" yaml:"roleArn,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
		} `tfsdk:"aws" yaml:"aws,omitempty"`

		AwsEc2 *struct {
			Filters *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				KvPair *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"kv_pair" yaml:"kvPair,omitempty"`
			} `tfsdk:"filters" yaml:"filters,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			PublicIp *bool `tfsdk:"public_ip" yaml:"publicIp,omitempty"`

			Region *string `tfsdk:"region" yaml:"region,omitempty"`

			RoleArn *string `tfsdk:"role_arn" yaml:"roleArn,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
		} `tfsdk:"aws_ec2" yaml:"awsEc2,omitempty"`

		Azure *struct {
			FunctionAppName *string `tfsdk:"function_app_name" yaml:"functionAppName,omitempty"`

			Functions *[]struct {
				AuthLevel utilities.IntOrString `tfsdk:"auth_level" yaml:"authLevel,omitempty"`

				FunctionName *string `tfsdk:"function_name" yaml:"functionName,omitempty"`
			} `tfsdk:"functions" yaml:"functions,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
		} `tfsdk:"azure" yaml:"azure,omitempty"`

		CircuitBreakers *struct {
			MaxConnections *int64 `tfsdk:"max_connections" yaml:"maxConnections,omitempty"`

			MaxPendingRequests *int64 `tfsdk:"max_pending_requests" yaml:"maxPendingRequests,omitempty"`

			MaxRequests *int64 `tfsdk:"max_requests" yaml:"maxRequests,omitempty"`

			MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`
		} `tfsdk:"circuit_breakers" yaml:"circuitBreakers,omitempty"`

		ConnectionConfig *struct {
			CommonHttpProtocolOptions *struct {
				HeadersWithUnderscoresAction utilities.IntOrString `tfsdk:"headers_with_underscores_action" yaml:"headersWithUnderscoresAction,omitempty"`

				IdleTimeout *string `tfsdk:"idle_timeout" yaml:"idleTimeout,omitempty"`

				MaxHeadersCount *int64 `tfsdk:"max_headers_count" yaml:"maxHeadersCount,omitempty"`

				MaxStreamDuration *string `tfsdk:"max_stream_duration" yaml:"maxStreamDuration,omitempty"`
			} `tfsdk:"common_http_protocol_options" yaml:"commonHttpProtocolOptions,omitempty"`

			ConnectTimeout *string `tfsdk:"connect_timeout" yaml:"connectTimeout,omitempty"`

			Http1ProtocolOptions *struct {
				EnableTrailers *bool `tfsdk:"enable_trailers" yaml:"enableTrailers,omitempty"`

				OverrideStreamErrorOnInvalidHttpMessage *bool `tfsdk:"override_stream_error_on_invalid_http_message" yaml:"overrideStreamErrorOnInvalidHttpMessage,omitempty"`

				PreserveCaseHeaderKeyFormat *bool `tfsdk:"preserve_case_header_key_format" yaml:"preserveCaseHeaderKeyFormat,omitempty"`

				ProperCaseHeaderKeyFormat *bool `tfsdk:"proper_case_header_key_format" yaml:"properCaseHeaderKeyFormat,omitempty"`
			} `tfsdk:"http1_protocol_options" yaml:"http1ProtocolOptions,omitempty"`

			MaxRequestsPerConnection *int64 `tfsdk:"max_requests_per_connection" yaml:"maxRequestsPerConnection,omitempty"`

			PerConnectionBufferLimitBytes *int64 `tfsdk:"per_connection_buffer_limit_bytes" yaml:"perConnectionBufferLimitBytes,omitempty"`

			TcpKeepalive *struct {
				KeepaliveInterval *string `tfsdk:"keepalive_interval" yaml:"keepaliveInterval,omitempty"`

				KeepaliveProbes *int64 `tfsdk:"keepalive_probes" yaml:"keepaliveProbes,omitempty"`

				KeepaliveTime *string `tfsdk:"keepalive_time" yaml:"keepaliveTime,omitempty"`
			} `tfsdk:"tcp_keepalive" yaml:"tcpKeepalive,omitempty"`
		} `tfsdk:"connection_config" yaml:"connectionConfig,omitempty"`

		Consul *struct {
			ConnectEnabled *bool `tfsdk:"connect_enabled" yaml:"connectEnabled,omitempty"`

			ConsistencyMode utilities.IntOrString `tfsdk:"consistency_mode" yaml:"consistencyMode,omitempty"`

			DataCenters *[]string `tfsdk:"data_centers" yaml:"dataCenters,omitempty"`

			InstanceBlacklistTags *[]string `tfsdk:"instance_blacklist_tags" yaml:"instanceBlacklistTags,omitempty"`

			InstanceTags *[]string `tfsdk:"instance_tags" yaml:"instanceTags,omitempty"`

			QueryOptions *struct {
				UseCache *bool `tfsdk:"use_cache" yaml:"useCache,omitempty"`
			} `tfsdk:"query_options" yaml:"queryOptions,omitempty"`

			ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`

			ServiceSpec *struct {
				Grpc *struct {
					Descriptors *string `tfsdk:"descriptors" yaml:"descriptors,omitempty"`

					GrpcServices *[]struct {
						FunctionNames *[]string `tfsdk:"function_names" yaml:"functionNames,omitempty"`

						PackageName *string `tfsdk:"package_name" yaml:"packageName,omitempty"`

						ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
					} `tfsdk:"grpc_services" yaml:"grpcServices,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				Rest *struct {
					SwaggerInfo *struct {
						Inline *string `tfsdk:"inline" yaml:"inline,omitempty"`

						Url *string `tfsdk:"url" yaml:"url,omitempty"`
					} `tfsdk:"swagger_info" yaml:"swaggerInfo,omitempty"`

					Transformations *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

						Body *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"body" yaml:"body,omitempty"`

						DynamicMetadataValues *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

						Extractors *struct {
							Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

							Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
						} `tfsdk:"extractors" yaml:"extractors,omitempty"`

						Headers *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"headers" yaml:"headers,omitempty"`

						HeadersToAppend *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

						HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

						IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

						ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

						Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
					} `tfsdk:"transformations" yaml:"transformations,omitempty"`
				} `tfsdk:"rest" yaml:"rest,omitempty"`
			} `tfsdk:"service_spec" yaml:"serviceSpec,omitempty"`

			ServiceTags *[]string `tfsdk:"service_tags" yaml:"serviceTags,omitempty"`

			SubsetTags *[]string `tfsdk:"subset_tags" yaml:"subsetTags,omitempty"`
		} `tfsdk:"consul" yaml:"consul,omitempty"`

		DiscoveryMetadata *struct {
			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
		} `tfsdk:"discovery_metadata" yaml:"discoveryMetadata,omitempty"`

		Failover *struct {
			Policy *struct {
				OverprovisioningFactor *int64 `tfsdk:"overprovisioning_factor" yaml:"overprovisioningFactor,omitempty"`
			} `tfsdk:"policy" yaml:"policy,omitempty"`

			PrioritizedLocalities *[]struct {
				LocalityEndpoints *[]struct {
					LbEndpoints *[]struct {
						Address *string `tfsdk:"address" yaml:"address,omitempty"`

						HealthCheckConfig *struct {
							Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

							Method *string `tfsdk:"method" yaml:"method,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							PortValue *int64 `tfsdk:"port_value" yaml:"portValue,omitempty"`
						} `tfsdk:"health_check_config" yaml:"healthCheckConfig,omitempty"`

						LoadBalancingWeight *int64 `tfsdk:"load_balancing_weight" yaml:"loadBalancingWeight,omitempty"`

						Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

						UpstreamSslConfig *struct {
							AllowRenegotiation *bool `tfsdk:"allow_renegotiation" yaml:"allowRenegotiation,omitempty"`

							AlpnProtocols *[]string `tfsdk:"alpn_protocols" yaml:"alpnProtocols,omitempty"`

							Parameters *struct {
								CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

								EcdhCurves *[]string `tfsdk:"ecdh_curves" yaml:"ecdhCurves,omitempty"`

								MaximumProtocolVersion utilities.IntOrString `tfsdk:"maximum_protocol_version" yaml:"maximumProtocolVersion,omitempty"`

								MinimumProtocolVersion utilities.IntOrString `tfsdk:"minimum_protocol_version" yaml:"minimumProtocolVersion,omitempty"`
							} `tfsdk:"parameters" yaml:"parameters,omitempty"`

							Sds *struct {
								CallCredentials *struct {
									FileCredentialSource *struct {
										Header *string `tfsdk:"header" yaml:"header,omitempty"`

										TokenFileName *string `tfsdk:"token_file_name" yaml:"tokenFileName,omitempty"`
									} `tfsdk:"file_credential_source" yaml:"fileCredentialSource,omitempty"`
								} `tfsdk:"call_credentials" yaml:"callCredentials,omitempty"`

								CertificatesSecretName *string `tfsdk:"certificates_secret_name" yaml:"certificatesSecretName,omitempty"`

								ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

								TargetUri *string `tfsdk:"target_uri" yaml:"targetUri,omitempty"`

								ValidationContextName *string `tfsdk:"validation_context_name" yaml:"validationContextName,omitempty"`
							} `tfsdk:"sds" yaml:"sds,omitempty"`

							SecretRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
							} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

							Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

							SslFiles *struct {
								RootCa *string `tfsdk:"root_ca" yaml:"rootCa,omitempty"`

								TlsCert *string `tfsdk:"tls_cert" yaml:"tlsCert,omitempty"`

								TlsKey *string `tfsdk:"tls_key" yaml:"tlsKey,omitempty"`
							} `tfsdk:"ssl_files" yaml:"sslFiles,omitempty"`

							VerifySubjectAltName *[]string `tfsdk:"verify_subject_alt_name" yaml:"verifySubjectAltName,omitempty"`
						} `tfsdk:"upstream_ssl_config" yaml:"upstreamSslConfig,omitempty"`
					} `tfsdk:"lb_endpoints" yaml:"lbEndpoints,omitempty"`

					LoadBalancingWeight *int64 `tfsdk:"load_balancing_weight" yaml:"loadBalancingWeight,omitempty"`

					Locality *struct {
						Region *string `tfsdk:"region" yaml:"region,omitempty"`

						SubZone *string `tfsdk:"sub_zone" yaml:"subZone,omitempty"`

						Zone *string `tfsdk:"zone" yaml:"zone,omitempty"`
					} `tfsdk:"locality" yaml:"locality,omitempty"`
				} `tfsdk:"locality_endpoints" yaml:"localityEndpoints,omitempty"`
			} `tfsdk:"prioritized_localities" yaml:"prioritizedLocalities,omitempty"`
		} `tfsdk:"failover" yaml:"failover,omitempty"`

		HealthChecks *[]struct {
			AlwaysLogHealthCheckFailures *bool `tfsdk:"always_log_health_check_failures" yaml:"alwaysLogHealthCheckFailures,omitempty"`

			CustomHealthCheck *struct {
				Config utilities.Dynamic `tfsdk:"config" yaml:"config,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				TypedConfig utilities.Dynamic `tfsdk:"typed_config" yaml:"typedConfig,omitempty"`
			} `tfsdk:"custom_health_check" yaml:"customHealthCheck,omitempty"`

			EventLogPath *string `tfsdk:"event_log_path" yaml:"eventLogPath,omitempty"`

			GrpcHealthCheck *struct {
				Authority *string `tfsdk:"authority" yaml:"authority,omitempty"`

				ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
			} `tfsdk:"grpc_health_check" yaml:"grpcHealthCheck,omitempty"`

			HealthyEdgeInterval *string `tfsdk:"healthy_edge_interval" yaml:"healthyEdgeInterval,omitempty"`

			HealthyThreshold *int64 `tfsdk:"healthy_threshold" yaml:"healthyThreshold,omitempty"`

			HttpHealthCheck *struct {
				ExpectedStatuses *[]struct {
					End utilities.IntOrString `tfsdk:"end" yaml:"end,omitempty"`

					Start utilities.IntOrString `tfsdk:"start" yaml:"start,omitempty"`
				} `tfsdk:"expected_statuses" yaml:"expectedStatuses,omitempty"`

				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				RequestHeadersToAdd *[]struct {
					Append *bool `tfsdk:"append" yaml:"append,omitempty"`

					Header *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"header" yaml:"header,omitempty"`

					HeaderSecretRef *struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"header_secret_ref" yaml:"headerSecretRef,omitempty"`
				} `tfsdk:"request_headers_to_add" yaml:"requestHeadersToAdd,omitempty"`

				RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" yaml:"requestHeadersToRemove,omitempty"`

				ResponseAssertions *struct {
					NoMatchHealth utilities.IntOrString `tfsdk:"no_match_health" yaml:"noMatchHealth,omitempty"`

					ResponseMatchers *[]struct {
						MatchHealth utilities.IntOrString `tfsdk:"match_health" yaml:"matchHealth,omitempty"`

						ResponseMatch *struct {
							Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

							JsonKey *struct {
								Path *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`
								} `tfsdk:"path" yaml:"path,omitempty"`
							} `tfsdk:"json_key" yaml:"jsonKey,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`
						} `tfsdk:"response_match" yaml:"responseMatch,omitempty"`
					} `tfsdk:"response_matchers" yaml:"responseMatchers,omitempty"`
				} `tfsdk:"response_assertions" yaml:"responseAssertions,omitempty"`

				ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`

				UseHttp2 *bool `tfsdk:"use_http2" yaml:"useHttp2,omitempty"`
			} `tfsdk:"http_health_check" yaml:"httpHealthCheck,omitempty"`

			InitialJitter *string `tfsdk:"initial_jitter" yaml:"initialJitter,omitempty"`

			Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

			IntervalJitter *string `tfsdk:"interval_jitter" yaml:"intervalJitter,omitempty"`

			IntervalJitterPercent *int64 `tfsdk:"interval_jitter_percent" yaml:"intervalJitterPercent,omitempty"`

			NoTrafficInterval *string `tfsdk:"no_traffic_interval" yaml:"noTrafficInterval,omitempty"`

			ReuseConnection *bool `tfsdk:"reuse_connection" yaml:"reuseConnection,omitempty"`

			TcpHealthCheck *struct {
				Receive *[]struct {
					Text *string `tfsdk:"text" yaml:"text,omitempty"`
				} `tfsdk:"receive" yaml:"receive,omitempty"`

				Send *struct {
					Text *string `tfsdk:"text" yaml:"text,omitempty"`
				} `tfsdk:"send" yaml:"send,omitempty"`
			} `tfsdk:"tcp_health_check" yaml:"tcpHealthCheck,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

			UnhealthyEdgeInterval *string `tfsdk:"unhealthy_edge_interval" yaml:"unhealthyEdgeInterval,omitempty"`

			UnhealthyInterval *string `tfsdk:"unhealthy_interval" yaml:"unhealthyInterval,omitempty"`

			UnhealthyThreshold *int64 `tfsdk:"unhealthy_threshold" yaml:"unhealthyThreshold,omitempty"`
		} `tfsdk:"health_checks" yaml:"healthChecks,omitempty"`

		HttpConnectHeaders *[]struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"http_connect_headers" yaml:"httpConnectHeaders,omitempty"`

		HttpConnectSslConfig *struct {
			AllowRenegotiation *bool `tfsdk:"allow_renegotiation" yaml:"allowRenegotiation,omitempty"`

			AlpnProtocols *[]string `tfsdk:"alpn_protocols" yaml:"alpnProtocols,omitempty"`

			Parameters *struct {
				CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

				EcdhCurves *[]string `tfsdk:"ecdh_curves" yaml:"ecdhCurves,omitempty"`

				MaximumProtocolVersion utilities.IntOrString `tfsdk:"maximum_protocol_version" yaml:"maximumProtocolVersion,omitempty"`

				MinimumProtocolVersion utilities.IntOrString `tfsdk:"minimum_protocol_version" yaml:"minimumProtocolVersion,omitempty"`
			} `tfsdk:"parameters" yaml:"parameters,omitempty"`

			Sds *struct {
				CallCredentials *struct {
					FileCredentialSource *struct {
						Header *string `tfsdk:"header" yaml:"header,omitempty"`

						TokenFileName *string `tfsdk:"token_file_name" yaml:"tokenFileName,omitempty"`
					} `tfsdk:"file_credential_source" yaml:"fileCredentialSource,omitempty"`
				} `tfsdk:"call_credentials" yaml:"callCredentials,omitempty"`

				CertificatesSecretName *string `tfsdk:"certificates_secret_name" yaml:"certificatesSecretName,omitempty"`

				ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

				TargetUri *string `tfsdk:"target_uri" yaml:"targetUri,omitempty"`

				ValidationContextName *string `tfsdk:"validation_context_name" yaml:"validationContextName,omitempty"`
			} `tfsdk:"sds" yaml:"sds,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

			SslFiles *struct {
				RootCa *string `tfsdk:"root_ca" yaml:"rootCa,omitempty"`

				TlsCert *string `tfsdk:"tls_cert" yaml:"tlsCert,omitempty"`

				TlsKey *string `tfsdk:"tls_key" yaml:"tlsKey,omitempty"`
			} `tfsdk:"ssl_files" yaml:"sslFiles,omitempty"`

			VerifySubjectAltName *[]string `tfsdk:"verify_subject_alt_name" yaml:"verifySubjectAltName,omitempty"`
		} `tfsdk:"http_connect_ssl_config" yaml:"httpConnectSslConfig,omitempty"`

		HttpProxyHostname *string `tfsdk:"http_proxy_hostname" yaml:"httpProxyHostname,omitempty"`

		IgnoreHealthOnHostRemoval *bool `tfsdk:"ignore_health_on_host_removal" yaml:"ignoreHealthOnHostRemoval,omitempty"`

		InitialConnectionWindowSize *int64 `tfsdk:"initial_connection_window_size" yaml:"initialConnectionWindowSize,omitempty"`

		InitialStreamWindowSize *int64 `tfsdk:"initial_stream_window_size" yaml:"initialStreamWindowSize,omitempty"`

		Kube *struct {
			Selector *map[string]string `tfsdk:"selector" yaml:"selector,omitempty"`

			ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`

			ServiceNamespace *string `tfsdk:"service_namespace" yaml:"serviceNamespace,omitempty"`

			ServicePort *int64 `tfsdk:"service_port" yaml:"servicePort,omitempty"`

			ServiceSpec *struct {
				Grpc *struct {
					Descriptors *string `tfsdk:"descriptors" yaml:"descriptors,omitempty"`

					GrpcServices *[]struct {
						FunctionNames *[]string `tfsdk:"function_names" yaml:"functionNames,omitempty"`

						PackageName *string `tfsdk:"package_name" yaml:"packageName,omitempty"`

						ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
					} `tfsdk:"grpc_services" yaml:"grpcServices,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				Rest *struct {
					SwaggerInfo *struct {
						Inline *string `tfsdk:"inline" yaml:"inline,omitempty"`

						Url *string `tfsdk:"url" yaml:"url,omitempty"`
					} `tfsdk:"swagger_info" yaml:"swaggerInfo,omitempty"`

					Transformations *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

						Body *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"body" yaml:"body,omitempty"`

						DynamicMetadataValues *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

						Extractors *struct {
							Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

							Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
						} `tfsdk:"extractors" yaml:"extractors,omitempty"`

						Headers *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"headers" yaml:"headers,omitempty"`

						HeadersToAppend *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

						HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

						IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

						ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

						Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
					} `tfsdk:"transformations" yaml:"transformations,omitempty"`
				} `tfsdk:"rest" yaml:"rest,omitempty"`
			} `tfsdk:"service_spec" yaml:"serviceSpec,omitempty"`

			SubsetSpec *struct {
				DefaultSubset *struct {
					Values *map[string]string `tfsdk:"values" yaml:"values,omitempty"`
				} `tfsdk:"default_subset" yaml:"defaultSubset,omitempty"`

				FallbackPolicy utilities.IntOrString `tfsdk:"fallback_policy" yaml:"fallbackPolicy,omitempty"`

				Selectors *[]struct {
					Keys *[]string `tfsdk:"keys" yaml:"keys,omitempty"`

					SingleHostPerSubset *bool `tfsdk:"single_host_per_subset" yaml:"singleHostPerSubset,omitempty"`
				} `tfsdk:"selectors" yaml:"selectors,omitempty"`
			} `tfsdk:"subset_spec" yaml:"subsetSpec,omitempty"`
		} `tfsdk:"kube" yaml:"kube,omitempty"`

		LoadBalancerConfig *struct {
			HealthyPanicThreshold utilities.DynamicNumber `tfsdk:"healthy_panic_threshold" yaml:"healthyPanicThreshold,omitempty"`

			LeastRequest *struct {
				ChoiceCount *int64 `tfsdk:"choice_count" yaml:"choiceCount,omitempty"`
			} `tfsdk:"least_request" yaml:"leastRequest,omitempty"`

			LocalityWeightedLbConfig *map[string]string `tfsdk:"locality_weighted_lb_config" yaml:"localityWeightedLbConfig,omitempty"`

			Maglev *map[string]string `tfsdk:"maglev" yaml:"maglev,omitempty"`

			Random *map[string]string `tfsdk:"random" yaml:"random,omitempty"`

			RingHash *struct {
				RingHashConfig *struct {
					MaximumRingSize utilities.IntOrString `tfsdk:"maximum_ring_size" yaml:"maximumRingSize,omitempty"`

					MinimumRingSize utilities.IntOrString `tfsdk:"minimum_ring_size" yaml:"minimumRingSize,omitempty"`
				} `tfsdk:"ring_hash_config" yaml:"ringHashConfig,omitempty"`
			} `tfsdk:"ring_hash" yaml:"ringHash,omitempty"`

			RoundRobin *map[string]string `tfsdk:"round_robin" yaml:"roundRobin,omitempty"`

			UpdateMergeWindow *string `tfsdk:"update_merge_window" yaml:"updateMergeWindow,omitempty"`
		} `tfsdk:"load_balancer_config" yaml:"loadBalancerConfig,omitempty"`

		MaxConcurrentStreams *int64 `tfsdk:"max_concurrent_streams" yaml:"maxConcurrentStreams,omitempty"`

		NamespacedStatuses *struct {
			Statuses utilities.Dynamic `tfsdk:"statuses" yaml:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" yaml:"namespacedStatuses,omitempty"`

		OutlierDetection *struct {
			BaseEjectionTime *string `tfsdk:"base_ejection_time" yaml:"baseEjectionTime,omitempty"`

			Consecutive5xx *int64 `tfsdk:"consecutive5xx" yaml:"consecutive5xx,omitempty"`

			ConsecutiveGatewayFailure *int64 `tfsdk:"consecutive_gateway_failure" yaml:"consecutiveGatewayFailure,omitempty"`

			ConsecutiveLocalOriginFailure *int64 `tfsdk:"consecutive_local_origin_failure" yaml:"consecutiveLocalOriginFailure,omitempty"`

			EnforcingConsecutive5xx *int64 `tfsdk:"enforcing_consecutive5xx" yaml:"enforcingConsecutive5xx,omitempty"`

			EnforcingConsecutiveGatewayFailure *int64 `tfsdk:"enforcing_consecutive_gateway_failure" yaml:"enforcingConsecutiveGatewayFailure,omitempty"`

			EnforcingConsecutiveLocalOriginFailure *int64 `tfsdk:"enforcing_consecutive_local_origin_failure" yaml:"enforcingConsecutiveLocalOriginFailure,omitempty"`

			EnforcingLocalOriginSuccessRate *int64 `tfsdk:"enforcing_local_origin_success_rate" yaml:"enforcingLocalOriginSuccessRate,omitempty"`

			EnforcingSuccessRate *int64 `tfsdk:"enforcing_success_rate" yaml:"enforcingSuccessRate,omitempty"`

			Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

			MaxEjectionPercent *int64 `tfsdk:"max_ejection_percent" yaml:"maxEjectionPercent,omitempty"`

			SplitExternalLocalOriginErrors *bool `tfsdk:"split_external_local_origin_errors" yaml:"splitExternalLocalOriginErrors,omitempty"`

			SuccessRateMinimumHosts *int64 `tfsdk:"success_rate_minimum_hosts" yaml:"successRateMinimumHosts,omitempty"`

			SuccessRateRequestVolume *int64 `tfsdk:"success_rate_request_volume" yaml:"successRateRequestVolume,omitempty"`

			SuccessRateStdevFactor *int64 `tfsdk:"success_rate_stdev_factor" yaml:"successRateStdevFactor,omitempty"`
		} `tfsdk:"outlier_detection" yaml:"outlierDetection,omitempty"`

		OverrideStreamErrorOnInvalidHttpMessage *bool `tfsdk:"override_stream_error_on_invalid_http_message" yaml:"overrideStreamErrorOnInvalidHttpMessage,omitempty"`

		Pipe *struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			ServiceSpec *struct {
				Grpc *struct {
					Descriptors *string `tfsdk:"descriptors" yaml:"descriptors,omitempty"`

					GrpcServices *[]struct {
						FunctionNames *[]string `tfsdk:"function_names" yaml:"functionNames,omitempty"`

						PackageName *string `tfsdk:"package_name" yaml:"packageName,omitempty"`

						ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
					} `tfsdk:"grpc_services" yaml:"grpcServices,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				Rest *struct {
					SwaggerInfo *struct {
						Inline *string `tfsdk:"inline" yaml:"inline,omitempty"`

						Url *string `tfsdk:"url" yaml:"url,omitempty"`
					} `tfsdk:"swagger_info" yaml:"swaggerInfo,omitempty"`

					Transformations *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

						Body *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"body" yaml:"body,omitempty"`

						DynamicMetadataValues *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

						Extractors *struct {
							Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

							Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
						} `tfsdk:"extractors" yaml:"extractors,omitempty"`

						Headers *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"headers" yaml:"headers,omitempty"`

						HeadersToAppend *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

						HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

						IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

						ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

						Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
					} `tfsdk:"transformations" yaml:"transformations,omitempty"`
				} `tfsdk:"rest" yaml:"rest,omitempty"`
			} `tfsdk:"service_spec" yaml:"serviceSpec,omitempty"`
		} `tfsdk:"pipe" yaml:"pipe,omitempty"`

		ProtocolSelection utilities.IntOrString `tfsdk:"protocol_selection" yaml:"protocolSelection,omitempty"`

		SslConfig *struct {
			AllowRenegotiation *bool `tfsdk:"allow_renegotiation" yaml:"allowRenegotiation,omitempty"`

			AlpnProtocols *[]string `tfsdk:"alpn_protocols" yaml:"alpnProtocols,omitempty"`

			Parameters *struct {
				CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

				EcdhCurves *[]string `tfsdk:"ecdh_curves" yaml:"ecdhCurves,omitempty"`

				MaximumProtocolVersion utilities.IntOrString `tfsdk:"maximum_protocol_version" yaml:"maximumProtocolVersion,omitempty"`

				MinimumProtocolVersion utilities.IntOrString `tfsdk:"minimum_protocol_version" yaml:"minimumProtocolVersion,omitempty"`
			} `tfsdk:"parameters" yaml:"parameters,omitempty"`

			Sds *struct {
				CallCredentials *struct {
					FileCredentialSource *struct {
						Header *string `tfsdk:"header" yaml:"header,omitempty"`

						TokenFileName *string `tfsdk:"token_file_name" yaml:"tokenFileName,omitempty"`
					} `tfsdk:"file_credential_source" yaml:"fileCredentialSource,omitempty"`
				} `tfsdk:"call_credentials" yaml:"callCredentials,omitempty"`

				CertificatesSecretName *string `tfsdk:"certificates_secret_name" yaml:"certificatesSecretName,omitempty"`

				ClusterName *string `tfsdk:"cluster_name" yaml:"clusterName,omitempty"`

				TargetUri *string `tfsdk:"target_uri" yaml:"targetUri,omitempty"`

				ValidationContextName *string `tfsdk:"validation_context_name" yaml:"validationContextName,omitempty"`
			} `tfsdk:"sds" yaml:"sds,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			Sni *string `tfsdk:"sni" yaml:"sni,omitempty"`

			SslFiles *struct {
				RootCa *string `tfsdk:"root_ca" yaml:"rootCa,omitempty"`

				TlsCert *string `tfsdk:"tls_cert" yaml:"tlsCert,omitempty"`

				TlsKey *string `tfsdk:"tls_key" yaml:"tlsKey,omitempty"`
			} `tfsdk:"ssl_files" yaml:"sslFiles,omitempty"`

			VerifySubjectAltName *[]string `tfsdk:"verify_subject_alt_name" yaml:"verifySubjectAltName,omitempty"`
		} `tfsdk:"ssl_config" yaml:"sslConfig,omitempty"`

		Static *struct {
			AutoSniRewrite *bool `tfsdk:"auto_sni_rewrite" yaml:"autoSniRewrite,omitempty"`

			Hosts *[]struct {
				Addr *string `tfsdk:"addr" yaml:"addr,omitempty"`

				HealthCheckConfig *struct {
					Method *string `tfsdk:"method" yaml:"method,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"health_check_config" yaml:"healthCheckConfig,omitempty"`

				LoadBalancingWeight *int64 `tfsdk:"load_balancing_weight" yaml:"loadBalancingWeight,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				SniAddr *string `tfsdk:"sni_addr" yaml:"sniAddr,omitempty"`
			} `tfsdk:"hosts" yaml:"hosts,omitempty"`

			ServiceSpec *struct {
				Grpc *struct {
					Descriptors *string `tfsdk:"descriptors" yaml:"descriptors,omitempty"`

					GrpcServices *[]struct {
						FunctionNames *[]string `tfsdk:"function_names" yaml:"functionNames,omitempty"`

						PackageName *string `tfsdk:"package_name" yaml:"packageName,omitempty"`

						ServiceName *string `tfsdk:"service_name" yaml:"serviceName,omitempty"`
					} `tfsdk:"grpc_services" yaml:"grpcServices,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				Rest *struct {
					SwaggerInfo *struct {
						Inline *string `tfsdk:"inline" yaml:"inline,omitempty"`

						Url *string `tfsdk:"url" yaml:"url,omitempty"`
					} `tfsdk:"swagger_info" yaml:"swaggerInfo,omitempty"`

					Transformations *struct {
						AdvancedTemplates *bool `tfsdk:"advanced_templates" yaml:"advancedTemplates,omitempty"`

						Body *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"body" yaml:"body,omitempty"`

						DynamicMetadataValues *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							MetadataNamespace *string `tfsdk:"metadata_namespace" yaml:"metadataNamespace,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"dynamic_metadata_values" yaml:"dynamicMetadataValues,omitempty"`

						Extractors *struct {
							Body *map[string]string `tfsdk:"body" yaml:"body,omitempty"`

							Header *string `tfsdk:"header" yaml:"header,omitempty"`

							Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

							Subgroup *int64 `tfsdk:"subgroup" yaml:"subgroup,omitempty"`
						} `tfsdk:"extractors" yaml:"extractors,omitempty"`

						Headers *struct {
							Text *string `tfsdk:"text" yaml:"text,omitempty"`
						} `tfsdk:"headers" yaml:"headers,omitempty"`

						HeadersToAppend *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Value *struct {
								Text *string `tfsdk:"text" yaml:"text,omitempty"`
							} `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"headers_to_append" yaml:"headersToAppend,omitempty"`

						HeadersToRemove *[]string `tfsdk:"headers_to_remove" yaml:"headersToRemove,omitempty"`

						IgnoreErrorOnParse *bool `tfsdk:"ignore_error_on_parse" yaml:"ignoreErrorOnParse,omitempty"`

						MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" yaml:"mergeExtractorsToBody,omitempty"`

						ParseBodyBehavior utilities.IntOrString `tfsdk:"parse_body_behavior" yaml:"parseBodyBehavior,omitempty"`

						Passthrough *map[string]string `tfsdk:"passthrough" yaml:"passthrough,omitempty"`
					} `tfsdk:"transformations" yaml:"transformations,omitempty"`
				} `tfsdk:"rest" yaml:"rest,omitempty"`
			} `tfsdk:"service_spec" yaml:"serviceSpec,omitempty"`

			UseTls *bool `tfsdk:"use_tls" yaml:"useTls,omitempty"`
		} `tfsdk:"static" yaml:"static,omitempty"`

		UseHttp2 *bool `tfsdk:"use_http2" yaml:"useHttp2,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewGlooSoloIoUpstreamV1Resource() resource.Resource {
	return &GlooSoloIoUpstreamV1Resource{}
}

func (r *GlooSoloIoUpstreamV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gloo_solo_io_upstream_v1"
}

func (r *GlooSoloIoUpstreamV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"aws": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"aws_account_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_role_chaining": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"lambda_functions": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"lambda_function_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"logical_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"qualifier": {
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

							"region": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_ref": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
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

					"aws_ec2": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"filters": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"key": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kv_pair": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"public_ip": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"region": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_arn": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_ref": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
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

					"azure": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"function_app_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"functions": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"auth_level": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"function_name": {
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

							"secret_ref": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
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

					"circuit_breakers": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_connections": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"max_pending_requests": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"max_requests": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"max_retries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"connection_config": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"common_http_protocol_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"headers_with_underscores_action": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"idle_timeout": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_headers_count": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_stream_duration": {
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

							"connect_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"http1_protocol_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enable_trailers": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"override_stream_error_on_invalid_http_message": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"preserve_case_header_key_format": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"proper_case_header_key_format": {
										Description:         "",
										MarkdownDescription: "",

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

							"max_requests_per_connection": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"per_connection_buffer_limit_bytes": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"tcp_keepalive": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"keepalive_interval": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"keepalive_probes": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"keepalive_time": {
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

					"consul": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"connect_enabled": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"consistency_mode": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.IntOrStringType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"data_centers": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance_blacklist_tags": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance_tags": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"query_options": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"use_cache": {
										Description:         "",
										MarkdownDescription: "",

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

							"service_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_spec": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"grpc": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"descriptors": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"grpc_services": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"function_names": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"package_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_name": {
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

									"rest": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"swagger_info": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"inline": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"url": {
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

											"transformations": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"advanced_templates": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"body": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
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

													"dynamic_metadata_values": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_namespace": {
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

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
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

													"extractors": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"subgroup": {
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

													"headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
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

													"headers_to_append": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
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

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
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

													"headers_to_remove": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_error_on_parse": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"merge_extractors_to_body": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parse_body_behavior": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"passthrough": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_tags": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"subset_tags": {
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

					"discovery_metadata": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"labels": {
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

					"failover": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"policy": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"overprovisioning_factor": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),

											int64validator.AtMost(4.294967295e+09),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"prioritized_localities": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"locality_endpoints": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"lb_endpoints": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"address": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"health_check_config": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"hostname": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"method": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"port_value": {
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

													"load_balancing_weight": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															int64validator.AtLeast(0),

															int64validator.AtMost(4.294967295e+09),
														},
													},

													"port": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"upstream_ssl_config": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"allow_renegotiation": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"alpn_protocols": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"parameters": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"cipher_suites": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"ecdh_curves": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"maximum_protocol_version": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"minimum_protocol_version": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: utilities.IntOrStringType{},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"sds": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"call_credentials": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"file_credential_source": {
																				Description:         "",
																				MarkdownDescription: "",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"header": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"token_file_name": {
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

																	"certificates_secret_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"cluster_name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"target_uri": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"validation_context_name": {
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

															"secret_ref": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"namespace": {
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

															"sni": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ssl_files": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"root_ca": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"tls_cert": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"tls_key": {
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

															"verify_subject_alt_name": {
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

											"load_balancing_weight": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(4.294967295e+09),
												},
											},

											"locality": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"region": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_zone": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"zone": {
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"health_checks": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"always_log_health_check_failures": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_health_check": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

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

									"typed_config": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"event_log_path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"grpc_health_check": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"authority": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_name": {
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

							"healthy_edge_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"healthy_threshold": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"http_health_check": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"expected_statuses": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"end": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"start": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"request_headers_to_add": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"append": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"header": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
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

											"header_secret_ref": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
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

									"request_headers_to_remove": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"response_assertions": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"no_match_health": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"response_matchers": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"match_health": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"response_match": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"ignore_error_on_parse": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"json_key": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
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

															"regex": {
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

									"service_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_http2": {
										Description:         "",
										MarkdownDescription: "",

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

							"initial_jitter": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interval_jitter": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interval_jitter_percent": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"no_traffic_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"reuse_connection": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tcp_health_check": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"receive": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"text": {
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

									"send": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"text": {
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

							"timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"unhealthy_edge_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"unhealthy_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"unhealthy_threshold": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"http_connect_headers": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"key": {
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

					"http_connect_ssl_config": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allow_renegotiation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"alpn_protocols": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"parameters": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cipher_suites": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ecdh_curves": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"maximum_protocol_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"minimum_protocol_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sds": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"call_credentials": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"file_credential_source": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"header": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_file_name": {
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

									"certificates_secret_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cluster_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_uri": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"validation_context_name": {
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

							"secret_ref": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
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

							"sni": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_files": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"root_ca": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_cert": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_key": {
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

							"verify_subject_alt_name": {
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

					"http_proxy_hostname": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ignore_health_on_host_removal": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"initial_connection_window_size": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(0),

							int64validator.AtMost(4.294967295e+09),
						},
					},

					"initial_stream_window_size": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(0),

							int64validator.AtMost(4.294967295e+09),
						},
					},

					"kube": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"selector": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_namespace": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_spec": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"grpc": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"descriptors": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"grpc_services": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"function_names": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"package_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_name": {
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

									"rest": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"swagger_info": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"inline": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"url": {
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

											"transformations": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"advanced_templates": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"body": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
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

													"dynamic_metadata_values": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_namespace": {
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

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
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

													"extractors": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"subgroup": {
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

													"headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
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

													"headers_to_append": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
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

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
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

													"headers_to_remove": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_error_on_parse": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"merge_extractors_to_body": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parse_body_behavior": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"passthrough": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"subset_spec": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default_subset": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"values": {
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

									"fallback_policy": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"selectors": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"keys": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"single_host_per_subset": {
												Description:         "",
												MarkdownDescription: "",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"load_balancer_config": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"healthy_panic_threshold": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicNumberType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"least_request": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"choice_count": {
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

							"locality_weighted_lb_config": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"maglev": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"random": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ring_hash": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ring_hash_config": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"maximum_ring_size": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"minimum_ring_size": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.IntOrStringType{},

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

							"round_robin": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"update_merge_window": {
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

					"max_concurrent_streams": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(0),

							int64validator.AtMost(4.294967295e+09),
						},
					},

					"namespaced_statuses": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"statuses": {
								Description:         "",
								MarkdownDescription: "",

								Type: utilities.DynamicType{},

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
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"consecutive5xx": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"consecutive_gateway_failure": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"consecutive_local_origin_failure": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_consecutive5xx": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_consecutive_gateway_failure": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_consecutive_local_origin_failure": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_local_origin_success_rate": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"enforcing_success_rate": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"interval": {
								Description:         "",
								MarkdownDescription: "",

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

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"split_external_local_origin_errors": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"success_rate_minimum_hosts": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"success_rate_request_volume": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},

							"success_rate_stdev_factor": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(0),

									int64validator.AtMost(4.294967295e+09),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"override_stream_error_on_invalid_http_message": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pipe": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_spec": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"grpc": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"descriptors": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"grpc_services": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"function_names": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"package_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_name": {
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

									"rest": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"swagger_info": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"inline": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"url": {
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

											"transformations": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"advanced_templates": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"body": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
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

													"dynamic_metadata_values": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_namespace": {
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

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
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

													"extractors": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"subgroup": {
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

													"headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
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

													"headers_to_append": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
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

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
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

													"headers_to_remove": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_error_on_parse": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"merge_extractors_to_body": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parse_body_behavior": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"passthrough": {
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

					"protocol_selection": {
						Description:         "",
						MarkdownDescription: "",

						Type: utilities.IntOrStringType{},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ssl_config": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allow_renegotiation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"alpn_protocols": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"parameters": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cipher_suites": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ecdh_curves": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"maximum_protocol_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"minimum_protocol_version": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sds": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"call_credentials": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"file_credential_source": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"header": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"token_file_name": {
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

									"certificates_secret_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cluster_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_uri": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"validation_context_name": {
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

							"secret_ref": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
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

							"sni": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ssl_files": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"root_ca": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_cert": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_key": {
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

							"verify_subject_alt_name": {
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

					"static": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"auto_sni_rewrite": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hosts": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"addr": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"health_check_config": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"method": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
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

									"load_balancing_weight": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),

											int64validator.AtMost(4.294967295e+09),
										},
									},

									"port": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sni_addr": {
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

							"service_spec": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"grpc": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"descriptors": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.Base64Validator(),
												},
											},

											"grpc_services": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"function_names": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"package_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service_name": {
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

									"rest": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"swagger_info": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"inline": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"url": {
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

											"transformations": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"advanced_templates": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"body": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
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

													"dynamic_metadata_values": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"metadata_namespace": {
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

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
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

													"extractors": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"body": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"header": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"regex": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"subgroup": {
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

													"headers": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"text": {
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

													"headers_to_append": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
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

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"text": {
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

													"headers_to_remove": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ignore_error_on_parse": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"merge_extractors_to_body": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"parse_body_behavior": {
														Description:         "",
														MarkdownDescription: "",

														Type: utilities.IntOrStringType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"passthrough": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"use_tls": {
								Description:         "",
								MarkdownDescription: "",

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

					"use_http2": {
						Description:         "",
						MarkdownDescription: "",

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
		},
	}, nil
}

func (r *GlooSoloIoUpstreamV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_gloo_solo_io_upstream_v1")

	var state GlooSoloIoUpstreamV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GlooSoloIoUpstreamV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("Upstream")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GlooSoloIoUpstreamV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_gloo_solo_io_upstream_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *GlooSoloIoUpstreamV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_gloo_solo_io_upstream_v1")

	var state GlooSoloIoUpstreamV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel GlooSoloIoUpstreamV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("gloo.solo.io/v1")
	goModel.Kind = utilities.Ptr("Upstream")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *GlooSoloIoUpstreamV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_gloo_solo_io_upstream_v1")
	// NO-OP: Terraform removes the state automatically for us
}
