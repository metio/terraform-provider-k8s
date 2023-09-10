/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package gateway_solo_io_v1

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
	_ datasource.DataSource              = &GatewaySoloIoVirtualServiceV1DataSource{}
	_ datasource.DataSourceWithConfigure = &GatewaySoloIoVirtualServiceV1DataSource{}
)

func NewGatewaySoloIoVirtualServiceV1DataSource() datasource.DataSource {
	return &GatewaySoloIoVirtualServiceV1DataSource{}
}

type GatewaySoloIoVirtualServiceV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type GatewaySoloIoVirtualServiceV1DataSourceData struct {
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
		DisplayName        *string `tfsdk:"display_name" json:"displayName,omitempty"`
		NamespacedStatuses *struct {
			Statuses *map[string]string `tfsdk:"statuses" json:"statuses,omitempty"`
		} `tfsdk:"namespaced_statuses" json:"namespacedStatuses,omitempty"`
		SslConfig *struct {
			AlpnProtocols               *[]string `tfsdk:"alpn_protocols" json:"alpnProtocols,omitempty"`
			DisableTlsSessionResumption *bool     `tfsdk:"disable_tls_session_resumption" json:"disableTlsSessionResumption,omitempty"`
			OcspStaplePolicy            *string   `tfsdk:"ocsp_staple_policy" json:"ocspStaplePolicy,omitempty"`
			OneWayTls                   *bool     `tfsdk:"one_way_tls" json:"oneWayTls,omitempty"`
			Parameters                  *struct {
				CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
				EcdhCurves             *[]string `tfsdk:"ecdh_curves" json:"ecdhCurves,omitempty"`
				MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
				MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
			} `tfsdk:"parameters" json:"parameters,omitempty"`
			Sds *struct {
				CallCredentials *struct {
					FileCredentialSource *struct {
						Header        *string `tfsdk:"header" json:"header,omitempty"`
						TokenFileName *string `tfsdk:"token_file_name" json:"tokenFileName,omitempty"`
					} `tfsdk:"file_credential_source" json:"fileCredentialSource,omitempty"`
				} `tfsdk:"call_credentials" json:"callCredentials,omitempty"`
				CertificatesSecretName *string `tfsdk:"certificates_secret_name" json:"certificatesSecretName,omitempty"`
				ClusterName            *string `tfsdk:"cluster_name" json:"clusterName,omitempty"`
				TargetUri              *string `tfsdk:"target_uri" json:"targetUri,omitempty"`
				ValidationContextName  *string `tfsdk:"validation_context_name" json:"validationContextName,omitempty"`
			} `tfsdk:"sds" json:"sds,omitempty"`
			SecretRef *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			SniDomains *[]string `tfsdk:"sni_domains" json:"sniDomains,omitempty"`
			SslFiles   *struct {
				OcspStaple *string `tfsdk:"ocsp_staple" json:"ocspStaple,omitempty"`
				RootCa     *string `tfsdk:"root_ca" json:"rootCa,omitempty"`
				TlsCert    *string `tfsdk:"tls_cert" json:"tlsCert,omitempty"`
				TlsKey     *string `tfsdk:"tls_key" json:"tlsKey,omitempty"`
			} `tfsdk:"ssl_files" json:"sslFiles,omitempty"`
			TransportSocketConnectTimeout *string   `tfsdk:"transport_socket_connect_timeout" json:"transportSocketConnectTimeout,omitempty"`
			VerifySubjectAltName          *[]string `tfsdk:"verify_subject_alt_name" json:"verifySubjectAltName,omitempty"`
		} `tfsdk:"ssl_config" json:"sslConfig,omitempty"`
		VirtualHost *struct {
			Domains *[]string `tfsdk:"domains" json:"domains,omitempty"`
			Options *struct {
				BufferPerRoute *struct {
					Buffer *struct {
						MaxRequestBytes *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
					} `tfsdk:"buffer" json:"buffer,omitempty"`
					Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				} `tfsdk:"buffer_per_route" json:"bufferPerRoute,omitempty"`
				Cors *struct {
					AllowCredentials *bool     `tfsdk:"allow_credentials" json:"allowCredentials,omitempty"`
					AllowHeaders     *[]string `tfsdk:"allow_headers" json:"allowHeaders,omitempty"`
					AllowMethods     *[]string `tfsdk:"allow_methods" json:"allowMethods,omitempty"`
					AllowOrigin      *[]string `tfsdk:"allow_origin" json:"allowOrigin,omitempty"`
					AllowOriginRegex *[]string `tfsdk:"allow_origin_regex" json:"allowOriginRegex,omitempty"`
					DisableForRoute  *bool     `tfsdk:"disable_for_route" json:"disableForRoute,omitempty"`
					ExposeHeaders    *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
					MaxAge           *string   `tfsdk:"max_age" json:"maxAge,omitempty"`
				} `tfsdk:"cors" json:"cors,omitempty"`
				Csrf *struct {
					AdditionalOrigins *[]struct {
						Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
						IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
						Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
						SafeRegex  *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"safe_regex" json:"safeRegex,omitempty"`
						Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
					} `tfsdk:"additional_origins" json:"additionalOrigins,omitempty"`
					FilterEnabled *struct {
						DefaultValue *struct {
							Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
							Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
						} `tfsdk:"default_value" json:"defaultValue,omitempty"`
						RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
					} `tfsdk:"filter_enabled" json:"filterEnabled,omitempty"`
					ShadowEnabled *struct {
						DefaultValue *struct {
							Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
							Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
						} `tfsdk:"default_value" json:"defaultValue,omitempty"`
						RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
					} `tfsdk:"shadow_enabled" json:"shadowEnabled,omitempty"`
				} `tfsdk:"csrf" json:"csrf,omitempty"`
				Dlp *struct {
					Actions *[]struct {
						ActionType   *string `tfsdk:"action_type" json:"actionType,omitempty"`
						CustomAction *struct {
							MaskChar *string `tfsdk:"mask_char" json:"maskChar,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Percent  *struct {
								Value *float64 `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"percent" json:"percent,omitempty"`
							Regex        *[]string `tfsdk:"regex" json:"regex,omitempty"`
							RegexActions *[]struct {
								Regex    *string `tfsdk:"regex" json:"regex,omitempty"`
								Subgroup *int64  `tfsdk:"subgroup" json:"subgroup,omitempty"`
							} `tfsdk:"regex_actions" json:"regexActions,omitempty"`
						} `tfsdk:"custom_action" json:"customAction,omitempty"`
						KeyValueAction *struct {
							KeyToMask *string `tfsdk:"key_to_mask" json:"keyToMask,omitempty"`
							MaskChar  *string `tfsdk:"mask_char" json:"maskChar,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Percent   *struct {
								Value *float64 `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"percent" json:"percent,omitempty"`
						} `tfsdk:"key_value_action" json:"keyValueAction,omitempty"`
						Shadow *bool `tfsdk:"shadow" json:"shadow,omitempty"`
					} `tfsdk:"actions" json:"actions,omitempty"`
					EnabledFor *string `tfsdk:"enabled_for" json:"enabledFor,omitempty"`
				} `tfsdk:"dlp" json:"dlp,omitempty"`
				ExtProc *struct {
					Disabled  *bool `tfsdk:"disabled" json:"disabled,omitempty"`
					Overrides *struct {
						AsyncMode   *bool `tfsdk:"async_mode" json:"asyncMode,omitempty"`
						GrpcService *struct {
							Authority        *string `tfsdk:"authority" json:"authority,omitempty"`
							ExtProcServerRef *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"ext_proc_server_ref" json:"extProcServerRef,omitempty"`
							InitialMetadata *[]struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"initial_metadata" json:"initialMetadata,omitempty"`
							RetryPolicy *struct {
								NumRetries   *int64 `tfsdk:"num_retries" json:"numRetries,omitempty"`
								RetryBackOff *struct {
									BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
									MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
								} `tfsdk:"retry_back_off" json:"retryBackOff,omitempty"`
							} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
							Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
						} `tfsdk:"grpc_service" json:"grpcService,omitempty"`
						MetadataContextNamespaces *[]string `tfsdk:"metadata_context_namespaces" json:"metadataContextNamespaces,omitempty"`
						ProcessingMode            *struct {
							RequestBodyMode     *string `tfsdk:"request_body_mode" json:"requestBodyMode,omitempty"`
							RequestHeaderMode   *string `tfsdk:"request_header_mode" json:"requestHeaderMode,omitempty"`
							RequestTrailerMode  *string `tfsdk:"request_trailer_mode" json:"requestTrailerMode,omitempty"`
							ResponseBodyMode    *string `tfsdk:"response_body_mode" json:"responseBodyMode,omitempty"`
							ResponseHeaderMode  *string `tfsdk:"response_header_mode" json:"responseHeaderMode,omitempty"`
							ResponseTrailerMode *string `tfsdk:"response_trailer_mode" json:"responseTrailerMode,omitempty"`
						} `tfsdk:"processing_mode" json:"processingMode,omitempty"`
						RequestAttributes              *[]string `tfsdk:"request_attributes" json:"requestAttributes,omitempty"`
						ResponseAttributes             *[]string `tfsdk:"response_attributes" json:"responseAttributes,omitempty"`
						TypedMetadataContextNamespaces *[]string `tfsdk:"typed_metadata_context_namespaces" json:"typedMetadataContextNamespaces,omitempty"`
					} `tfsdk:"overrides" json:"overrides,omitempty"`
				} `tfsdk:"ext_proc" json:"extProc,omitempty"`
				Extauth *struct {
					ConfigRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"config_ref" json:"configRef,omitempty"`
					CustomAuth *struct {
						ContextExtensions *map[string]string `tfsdk:"context_extensions" json:"contextExtensions,omitempty"`
						Name              *string            `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"custom_auth" json:"customAuth,omitempty"`
					Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
				} `tfsdk:"extauth" json:"extauth,omitempty"`
				Extensions *struct {
					Configs *map[string]string `tfsdk:"configs" json:"configs,omitempty"`
				} `tfsdk:"extensions" json:"extensions,omitempty"`
				HeaderManipulation *struct {
					RequestHeadersToAdd *[]struct {
						Append *bool `tfsdk:"append" json:"append,omitempty"`
						Header *struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"header" json:"header,omitempty"`
						HeaderSecretRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"header_secret_ref" json:"headerSecretRef,omitempty"`
					} `tfsdk:"request_headers_to_add" json:"requestHeadersToAdd,omitempty"`
					RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" json:"requestHeadersToRemove,omitempty"`
					ResponseHeadersToAdd   *[]struct {
						Append *bool `tfsdk:"append" json:"append,omitempty"`
						Header *struct {
							Key   *string `tfsdk:"key" json:"key,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"header" json:"header,omitempty"`
					} `tfsdk:"response_headers_to_add" json:"responseHeadersToAdd,omitempty"`
					ResponseHeadersToRemove *[]string `tfsdk:"response_headers_to_remove" json:"responseHeadersToRemove,omitempty"`
				} `tfsdk:"header_manipulation" json:"headerManipulation,omitempty"`
				IncludeAttemptCountInResponse *bool `tfsdk:"include_attempt_count_in_response" json:"includeAttemptCountInResponse,omitempty"`
				IncludeRequestAttemptCount    *bool `tfsdk:"include_request_attempt_count" json:"includeRequestAttemptCount,omitempty"`
				Jwt                           *struct {
					AllowMissingOrFailedJwt *bool `tfsdk:"allow_missing_or_failed_jwt" json:"allowMissingOrFailedJwt,omitempty"`
					Providers               *struct {
						Audiences       *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
						ClaimsToHeaders *[]struct {
							Append *bool   `tfsdk:"append" json:"append,omitempty"`
							Claim  *string `tfsdk:"claim" json:"claim,omitempty"`
							Header *string `tfsdk:"header" json:"header,omitempty"`
						} `tfsdk:"claims_to_headers" json:"claimsToHeaders,omitempty"`
						ClockSkewSeconds *int64  `tfsdk:"clock_skew_seconds" json:"clockSkewSeconds,omitempty"`
						Issuer           *string `tfsdk:"issuer" json:"issuer,omitempty"`
						Jwks             *struct {
							Local *struct {
								Key *string `tfsdk:"key" json:"key,omitempty"`
							} `tfsdk:"local" json:"local,omitempty"`
							Remote *struct {
								AsyncFetch *struct {
									FastListener *bool `tfsdk:"fast_listener" json:"fastListener,omitempty"`
								} `tfsdk:"async_fetch" json:"asyncFetch,omitempty"`
								CacheDuration *string `tfsdk:"cache_duration" json:"cacheDuration,omitempty"`
								UpstreamRef   *struct {
									Name      *string `tfsdk:"name" json:"name,omitempty"`
									Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"upstream_ref" json:"upstreamRef,omitempty"`
								Url *string `tfsdk:"url" json:"url,omitempty"`
							} `tfsdk:"remote" json:"remote,omitempty"`
						} `tfsdk:"jwks" json:"jwks,omitempty"`
						KeepToken   *bool `tfsdk:"keep_token" json:"keepToken,omitempty"`
						TokenSource *struct {
							Headers *[]struct {
								Header *string `tfsdk:"header" json:"header,omitempty"`
								Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							QueryParams *[]string `tfsdk:"query_params" json:"queryParams,omitempty"`
						} `tfsdk:"token_source" json:"tokenSource,omitempty"`
					} `tfsdk:"providers" json:"providers,omitempty"`
				} `tfsdk:"jwt" json:"jwt,omitempty"`
				JwtStaged *struct {
					AfterExtAuth *struct {
						AllowMissingOrFailedJwt *bool `tfsdk:"allow_missing_or_failed_jwt" json:"allowMissingOrFailedJwt,omitempty"`
						Providers               *struct {
							Audiences       *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							ClaimsToHeaders *[]struct {
								Append *bool   `tfsdk:"append" json:"append,omitempty"`
								Claim  *string `tfsdk:"claim" json:"claim,omitempty"`
								Header *string `tfsdk:"header" json:"header,omitempty"`
							} `tfsdk:"claims_to_headers" json:"claimsToHeaders,omitempty"`
							ClockSkewSeconds *int64  `tfsdk:"clock_skew_seconds" json:"clockSkewSeconds,omitempty"`
							Issuer           *string `tfsdk:"issuer" json:"issuer,omitempty"`
							Jwks             *struct {
								Local *struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"local" json:"local,omitempty"`
								Remote *struct {
									AsyncFetch *struct {
										FastListener *bool `tfsdk:"fast_listener" json:"fastListener,omitempty"`
									} `tfsdk:"async_fetch" json:"asyncFetch,omitempty"`
									CacheDuration *string `tfsdk:"cache_duration" json:"cacheDuration,omitempty"`
									UpstreamRef   *struct {
										Name      *string `tfsdk:"name" json:"name,omitempty"`
										Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
									} `tfsdk:"upstream_ref" json:"upstreamRef,omitempty"`
									Url *string `tfsdk:"url" json:"url,omitempty"`
								} `tfsdk:"remote" json:"remote,omitempty"`
							} `tfsdk:"jwks" json:"jwks,omitempty"`
							KeepToken   *bool `tfsdk:"keep_token" json:"keepToken,omitempty"`
							TokenSource *struct {
								Headers *[]struct {
									Header *string `tfsdk:"header" json:"header,omitempty"`
									Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								QueryParams *[]string `tfsdk:"query_params" json:"queryParams,omitempty"`
							} `tfsdk:"token_source" json:"tokenSource,omitempty"`
						} `tfsdk:"providers" json:"providers,omitempty"`
					} `tfsdk:"after_ext_auth" json:"afterExtAuth,omitempty"`
					BeforeExtAuth *struct {
						AllowMissingOrFailedJwt *bool `tfsdk:"allow_missing_or_failed_jwt" json:"allowMissingOrFailedJwt,omitempty"`
						Providers               *struct {
							Audiences       *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
							ClaimsToHeaders *[]struct {
								Append *bool   `tfsdk:"append" json:"append,omitempty"`
								Claim  *string `tfsdk:"claim" json:"claim,omitempty"`
								Header *string `tfsdk:"header" json:"header,omitempty"`
							} `tfsdk:"claims_to_headers" json:"claimsToHeaders,omitempty"`
							ClockSkewSeconds *int64  `tfsdk:"clock_skew_seconds" json:"clockSkewSeconds,omitempty"`
							Issuer           *string `tfsdk:"issuer" json:"issuer,omitempty"`
							Jwks             *struct {
								Local *struct {
									Key *string `tfsdk:"key" json:"key,omitempty"`
								} `tfsdk:"local" json:"local,omitempty"`
								Remote *struct {
									AsyncFetch *struct {
										FastListener *bool `tfsdk:"fast_listener" json:"fastListener,omitempty"`
									} `tfsdk:"async_fetch" json:"asyncFetch,omitempty"`
									CacheDuration *string `tfsdk:"cache_duration" json:"cacheDuration,omitempty"`
									UpstreamRef   *struct {
										Name      *string `tfsdk:"name" json:"name,omitempty"`
										Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
									} `tfsdk:"upstream_ref" json:"upstreamRef,omitempty"`
									Url *string `tfsdk:"url" json:"url,omitempty"`
								} `tfsdk:"remote" json:"remote,omitempty"`
							} `tfsdk:"jwks" json:"jwks,omitempty"`
							KeepToken   *bool `tfsdk:"keep_token" json:"keepToken,omitempty"`
							TokenSource *struct {
								Headers *[]struct {
									Header *string `tfsdk:"header" json:"header,omitempty"`
									Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								QueryParams *[]string `tfsdk:"query_params" json:"queryParams,omitempty"`
							} `tfsdk:"token_source" json:"tokenSource,omitempty"`
						} `tfsdk:"providers" json:"providers,omitempty"`
					} `tfsdk:"before_ext_auth" json:"beforeExtAuth,omitempty"`
				} `tfsdk:"jwt_staged" json:"jwtStaged,omitempty"`
				RateLimitConfigs *struct {
					Refs *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"refs" json:"refs,omitempty"`
				} `tfsdk:"rate_limit_configs" json:"rateLimitConfigs,omitempty"`
				RateLimitEarlyConfigs *struct {
					Refs *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"refs" json:"refs,omitempty"`
				} `tfsdk:"rate_limit_early_configs" json:"rateLimitEarlyConfigs,omitempty"`
				RateLimitRegularConfigs *struct {
					Refs *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"refs" json:"refs,omitempty"`
				} `tfsdk:"rate_limit_regular_configs" json:"rateLimitRegularConfigs,omitempty"`
				Ratelimit *struct {
					LocalRatelimit *struct {
						FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
						MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
						TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
					} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
					RateLimits *[]struct {
						Actions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
							GenericKey         *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" json:"genericKey,omitempty"`
							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
								Headers         *[]struct {
									ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
									InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name         *string `tfsdk:"name" json:"name,omitempty"`
									PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
									PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
									RangeMatch   *struct {
										End   *int64 `tfsdk:"end" json:"end,omitempty"`
										Start *int64 `tfsdk:"start" json:"start,omitempty"`
									} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
									RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
									SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
							} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
							Metadata *struct {
								DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								MetadataKey   *struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Path *[]struct {
										Key *string `tfsdk:"key" json:"key,omitempty"`
									} `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
								Source *string `tfsdk:"source" json:"source,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
							} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
							SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
						} `tfsdk:"actions" json:"actions,omitempty"`
						SetActions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
							GenericKey         *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" json:"genericKey,omitempty"`
							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
								Headers         *[]struct {
									ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
									InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name         *string `tfsdk:"name" json:"name,omitempty"`
									PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
									PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
									RangeMatch   *struct {
										End   *int64 `tfsdk:"end" json:"end,omitempty"`
										Start *int64 `tfsdk:"start" json:"start,omitempty"`
									} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
									RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
									SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
							} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
							Metadata *struct {
								DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								MetadataKey   *struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Path *[]struct {
										Key *string `tfsdk:"key" json:"key,omitempty"`
									} `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
								Source *string `tfsdk:"source" json:"source,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
							} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
							SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
						} `tfsdk:"set_actions" json:"setActions,omitempty"`
					} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
				} `tfsdk:"ratelimit" json:"ratelimit,omitempty"`
				RatelimitBasic *struct {
					AnonymousLimits *struct {
						RequestsPerUnit *int64  `tfsdk:"requests_per_unit" json:"requestsPerUnit,omitempty"`
						Unit            *string `tfsdk:"unit" json:"unit,omitempty"`
					} `tfsdk:"anonymous_limits" json:"anonymousLimits,omitempty"`
					AuthorizedLimits *struct {
						RequestsPerUnit *int64  `tfsdk:"requests_per_unit" json:"requestsPerUnit,omitempty"`
						Unit            *string `tfsdk:"unit" json:"unit,omitempty"`
					} `tfsdk:"authorized_limits" json:"authorizedLimits,omitempty"`
				} `tfsdk:"ratelimit_basic" json:"ratelimitBasic,omitempty"`
				RatelimitEarly *struct {
					LocalRatelimit *struct {
						FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
						MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
						TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
					} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
					RateLimits *[]struct {
						Actions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
							GenericKey         *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" json:"genericKey,omitempty"`
							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
								Headers         *[]struct {
									ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
									InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name         *string `tfsdk:"name" json:"name,omitempty"`
									PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
									PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
									RangeMatch   *struct {
										End   *int64 `tfsdk:"end" json:"end,omitempty"`
										Start *int64 `tfsdk:"start" json:"start,omitempty"`
									} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
									RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
									SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
							} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
							Metadata *struct {
								DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								MetadataKey   *struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Path *[]struct {
										Key *string `tfsdk:"key" json:"key,omitempty"`
									} `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
								Source *string `tfsdk:"source" json:"source,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
							} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
							SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
						} `tfsdk:"actions" json:"actions,omitempty"`
						SetActions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
							GenericKey         *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" json:"genericKey,omitempty"`
							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
								Headers         *[]struct {
									ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
									InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name         *string `tfsdk:"name" json:"name,omitempty"`
									PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
									PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
									RangeMatch   *struct {
										End   *int64 `tfsdk:"end" json:"end,omitempty"`
										Start *int64 `tfsdk:"start" json:"start,omitempty"`
									} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
									RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
									SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
							} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
							Metadata *struct {
								DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								MetadataKey   *struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Path *[]struct {
										Key *string `tfsdk:"key" json:"key,omitempty"`
									} `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
								Source *string `tfsdk:"source" json:"source,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
							} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
							SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
						} `tfsdk:"set_actions" json:"setActions,omitempty"`
					} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
				} `tfsdk:"ratelimit_early" json:"ratelimitEarly,omitempty"`
				RatelimitRegular *struct {
					LocalRatelimit *struct {
						FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
						MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
						TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
					} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
					RateLimits *[]struct {
						Actions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
							GenericKey         *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" json:"genericKey,omitempty"`
							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
								Headers         *[]struct {
									ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
									InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name         *string `tfsdk:"name" json:"name,omitempty"`
									PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
									PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
									RangeMatch   *struct {
										End   *int64 `tfsdk:"end" json:"end,omitempty"`
										Start *int64 `tfsdk:"start" json:"start,omitempty"`
									} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
									RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
									SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
							} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
							Metadata *struct {
								DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								MetadataKey   *struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Path *[]struct {
										Key *string `tfsdk:"key" json:"key,omitempty"`
									} `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
								Source *string `tfsdk:"source" json:"source,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
							} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
							SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
						} `tfsdk:"actions" json:"actions,omitempty"`
						SetActions *[]struct {
							DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
							GenericKey         *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
							} `tfsdk:"generic_key" json:"genericKey,omitempty"`
							HeaderValueMatch *struct {
								DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
								Headers         *[]struct {
									ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
									InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name         *string `tfsdk:"name" json:"name,omitempty"`
									PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
									PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
									RangeMatch   *struct {
										End   *int64 `tfsdk:"end" json:"end,omitempty"`
										Start *int64 `tfsdk:"start" json:"start,omitempty"`
									} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
									RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
									SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
							} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
							Metadata *struct {
								DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								MetadataKey   *struct {
									Key  *string `tfsdk:"key" json:"key,omitempty"`
									Path *[]struct {
										Key *string `tfsdk:"key" json:"key,omitempty"`
									} `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
								Source *string `tfsdk:"source" json:"source,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
							RequestHeaders *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
							} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
							SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
						} `tfsdk:"set_actions" json:"setActions,omitempty"`
					} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
				} `tfsdk:"ratelimit_regular" json:"ratelimitRegular,omitempty"`
				Rbac *struct {
					Disable  *bool `tfsdk:"disable" json:"disable,omitempty"`
					Policies *struct {
						NestedClaimDelimiter *string `tfsdk:"nested_claim_delimiter" json:"nestedClaimDelimiter,omitempty"`
						Permissions          *struct {
							Methods    *[]string `tfsdk:"methods" json:"methods,omitempty"`
							PathPrefix *string   `tfsdk:"path_prefix" json:"pathPrefix,omitempty"`
						} `tfsdk:"permissions" json:"permissions,omitempty"`
						Principals *[]struct {
							JwtPrincipal *struct {
								Claims   *map[string]string `tfsdk:"claims" json:"claims,omitempty"`
								Matcher  *string            `tfsdk:"matcher" json:"matcher,omitempty"`
								Provider *string            `tfsdk:"provider" json:"provider,omitempty"`
							} `tfsdk:"jwt_principal" json:"jwtPrincipal,omitempty"`
						} `tfsdk:"principals" json:"principals,omitempty"`
					} `tfsdk:"policies" json:"policies,omitempty"`
				} `tfsdk:"rbac" json:"rbac,omitempty"`
				Retries *struct {
					NumRetries    *int64  `tfsdk:"num_retries" json:"numRetries,omitempty"`
					PerTryTimeout *string `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
					RetryBackOff  *struct {
						BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
						MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
					} `tfsdk:"retry_back_off" json:"retryBackOff,omitempty"`
					RetryOn *string `tfsdk:"retry_on" json:"retryOn,omitempty"`
				} `tfsdk:"retries" json:"retries,omitempty"`
				StagedTransformations *struct {
					Early *struct {
						RequestTransforms *[]struct {
							ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
							Matcher         *struct {
								CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
								ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
								Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
								Headers        *[]struct {
									InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name        *string `tfsdk:"name" json:"name,omitempty"`
									Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
									Value       *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
								Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
								QueryParameters *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
								Regex *string `tfsdk:"regex" json:"regex,omitempty"`
							} `tfsdk:"matcher" json:"matcher,omitempty"`
							RequestTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
								LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
									Body              *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"body" json:"body,omitempty"`
									DynamicMetadataValues *[]struct {
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
									} `tfsdk:"extractors" json:"extractors,omitempty"`
									Headers *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									HeadersToAppend *[]struct {
										Key   *string `tfsdk:"key" json:"key,omitempty"`
										Value *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
									HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
									IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
									ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
									Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
								XsltTransformation *struct {
									NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
									SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
									Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
							} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
							ResponseTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
								LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
									Body              *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"body" json:"body,omitempty"`
									DynamicMetadataValues *[]struct {
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
									} `tfsdk:"extractors" json:"extractors,omitempty"`
									Headers *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									HeadersToAppend *[]struct {
										Key   *string `tfsdk:"key" json:"key,omitempty"`
										Value *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
									HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
									IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
									ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
									Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
								XsltTransformation *struct {
									NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
									SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
									Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
							} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
						} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
						ResponseTransforms *[]struct {
							Matchers *[]struct {
								InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
								Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value       *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"matchers" json:"matchers,omitempty"`
							ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
							ResponseTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
								LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
									Body              *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"body" json:"body,omitempty"`
									DynamicMetadataValues *[]struct {
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
									} `tfsdk:"extractors" json:"extractors,omitempty"`
									Headers *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									HeadersToAppend *[]struct {
										Key   *string `tfsdk:"key" json:"key,omitempty"`
										Value *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
									HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
									IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
									ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
									Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
								XsltTransformation *struct {
									NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
									SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
									Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
							} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
						} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
					} `tfsdk:"early" json:"early,omitempty"`
					EscapeCharacters       *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
					InheritTransformation  *bool `tfsdk:"inherit_transformation" json:"inheritTransformation,omitempty"`
					LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
					Regular                *struct {
						RequestTransforms *[]struct {
							ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
							Matcher         *struct {
								CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
								ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
								Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
								Headers        *[]struct {
									InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name        *string `tfsdk:"name" json:"name,omitempty"`
									Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
									Value       *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
								Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
								QueryParameters *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
								Regex *string `tfsdk:"regex" json:"regex,omitempty"`
							} `tfsdk:"matcher" json:"matcher,omitempty"`
							RequestTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
								LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
									Body              *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"body" json:"body,omitempty"`
									DynamicMetadataValues *[]struct {
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
									} `tfsdk:"extractors" json:"extractors,omitempty"`
									Headers *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									HeadersToAppend *[]struct {
										Key   *string `tfsdk:"key" json:"key,omitempty"`
										Value *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
									HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
									IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
									ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
									Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
								XsltTransformation *struct {
									NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
									SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
									Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
							} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
							ResponseTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
								LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
									Body              *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"body" json:"body,omitempty"`
									DynamicMetadataValues *[]struct {
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
									} `tfsdk:"extractors" json:"extractors,omitempty"`
									Headers *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									HeadersToAppend *[]struct {
										Key   *string `tfsdk:"key" json:"key,omitempty"`
										Value *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
									HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
									IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
									ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
									Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
								XsltTransformation *struct {
									NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
									SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
									Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
							} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
						} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
						ResponseTransforms *[]struct {
							Matchers *[]struct {
								InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
								Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
								Value       *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"matchers" json:"matchers,omitempty"`
							ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
							ResponseTransformation *struct {
								HeaderBodyTransform *struct {
									AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
								} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
								LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
								TransformationTemplate *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
									Body              *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"body" json:"body,omitempty"`
									DynamicMetadataValues *[]struct {
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
									} `tfsdk:"extractors" json:"extractors,omitempty"`
									Headers *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									HeadersToAppend *[]struct {
										Key   *string `tfsdk:"key" json:"key,omitempty"`
										Value *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
									HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
									IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
									ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
									Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
								} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
								XsltTransformation *struct {
									NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
									SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
									Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
								} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
							} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
						} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
					} `tfsdk:"regular" json:"regular,omitempty"`
				} `tfsdk:"staged_transformations" json:"stagedTransformations,omitempty"`
				Stats *struct {
					VirtualClusters *[]struct {
						Method  *string `tfsdk:"method" json:"method,omitempty"`
						Name    *string `tfsdk:"name" json:"name,omitempty"`
						Pattern *string `tfsdk:"pattern" json:"pattern,omitempty"`
					} `tfsdk:"virtual_clusters" json:"virtualClusters,omitempty"`
				} `tfsdk:"stats" json:"stats,omitempty"`
				Transformations *struct {
					ClearRouteCache       *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
					RequestTransformation *struct {
						HeaderBodyTransform *struct {
							AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
						} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
						LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
						TransformationTemplate *struct {
							AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
							Body              *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"body" json:"body,omitempty"`
							DynamicMetadataValues *[]struct {
								Key               *string `tfsdk:"key" json:"key,omitempty"`
								MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
								Value             *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
							EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
							Extractors       *struct {
								Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
								Header   *string            `tfsdk:"header" json:"header,omitempty"`
								Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
								Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
							} `tfsdk:"extractors" json:"extractors,omitempty"`
							Headers *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							HeadersToAppend *[]struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
							HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
							IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
							MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
							ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
							Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
						} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
						XsltTransformation *struct {
							NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
							SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
							Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
						} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
					} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
					ResponseTransformation *struct {
						HeaderBodyTransform *struct {
							AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
						} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
						LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
						TransformationTemplate *struct {
							AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
							Body              *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"body" json:"body,omitempty"`
							DynamicMetadataValues *[]struct {
								Key               *string `tfsdk:"key" json:"key,omitempty"`
								MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
								Value             *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
							EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
							Extractors       *struct {
								Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
								Header   *string            `tfsdk:"header" json:"header,omitempty"`
								Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
								Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
							} `tfsdk:"extractors" json:"extractors,omitempty"`
							Headers *struct {
								Text *string `tfsdk:"text" json:"text,omitempty"`
							} `tfsdk:"headers" json:"headers,omitempty"`
							HeadersToAppend *[]struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
							HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
							IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
							MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
							ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
							Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
						} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
						XsltTransformation *struct {
							NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
							SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
							Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
						} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
					} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
				} `tfsdk:"transformations" json:"transformations,omitempty"`
				Waf *struct {
					AuditLogging *struct {
						Action   *string `tfsdk:"action" json:"action,omitempty"`
						Location *string `tfsdk:"location" json:"location,omitempty"`
					} `tfsdk:"audit_logging" json:"auditLogging,omitempty"`
					ConfigMapRuleSets *[]struct {
						ConfigMapRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
						DataMapKeys *[]string `tfsdk:"data_map_keys" json:"dataMapKeys,omitempty"`
					} `tfsdk:"config_map_rule_sets" json:"configMapRuleSets,omitempty"`
					CoreRuleSet *struct {
						CustomSettingsFile   *string `tfsdk:"custom_settings_file" json:"customSettingsFile,omitempty"`
						CustomSettingsString *string `tfsdk:"custom_settings_string" json:"customSettingsString,omitempty"`
					} `tfsdk:"core_rule_set" json:"coreRuleSet,omitempty"`
					CustomInterventionMessage *string `tfsdk:"custom_intervention_message" json:"customInterventionMessage,omitempty"`
					Disabled                  *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
					RequestHeadersOnly        *bool   `tfsdk:"request_headers_only" json:"requestHeadersOnly,omitempty"`
					ResponseHeadersOnly       *bool   `tfsdk:"response_headers_only" json:"responseHeadersOnly,omitempty"`
					RuleSets                  *[]struct {
						Directory *string   `tfsdk:"directory" json:"directory,omitempty"`
						Files     *[]string `tfsdk:"files" json:"files,omitempty"`
						RuleStr   *string   `tfsdk:"rule_str" json:"ruleStr,omitempty"`
					} `tfsdk:"rule_sets" json:"ruleSets,omitempty"`
				} `tfsdk:"waf" json:"waf,omitempty"`
			} `tfsdk:"options" json:"options,omitempty"`
			OptionsConfigRefs *struct {
				DelegateOptions *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"delegate_options" json:"delegateOptions,omitempty"`
			} `tfsdk:"options_config_refs" json:"optionsConfigRefs,omitempty"`
			Routes *[]struct {
				DelegateAction *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					Ref       *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"ref" json:"ref,omitempty"`
					Selector *struct {
						Expressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"expressions" json:"expressions,omitempty"`
						Labels     *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						Namespaces *[]string          `tfsdk:"namespaces" json:"namespaces,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"delegate_action" json:"delegateAction,omitempty"`
				DirectResponseAction *struct {
					Body   *string `tfsdk:"body" json:"body,omitempty"`
					Status *int64  `tfsdk:"status" json:"status,omitempty"`
				} `tfsdk:"direct_response_action" json:"directResponseAction,omitempty"`
				GraphqlApiRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"graphql_api_ref" json:"graphqlApiRef,omitempty"`
				InheritableMatchers     *bool `tfsdk:"inheritable_matchers" json:"inheritableMatchers,omitempty"`
				InheritablePathMatchers *bool `tfsdk:"inheritable_path_matchers" json:"inheritablePathMatchers,omitempty"`
				Matchers                *[]struct {
					CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
					ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
					Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
					Headers        *[]struct {
						InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
						Name        *string `tfsdk:"name" json:"name,omitempty"`
						Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
						Value       *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
					Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
					QueryParameters *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
					Regex *string `tfsdk:"regex" json:"regex,omitempty"`
				} `tfsdk:"matchers" json:"matchers,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Options *struct {
					AppendXForwardedHost *bool `tfsdk:"append_x_forwarded_host" json:"appendXForwardedHost,omitempty"`
					AutoHostRewrite      *bool `tfsdk:"auto_host_rewrite" json:"autoHostRewrite,omitempty"`
					BufferPerRoute       *struct {
						Buffer *struct {
							MaxRequestBytes *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
						} `tfsdk:"buffer" json:"buffer,omitempty"`
						Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
					} `tfsdk:"buffer_per_route" json:"bufferPerRoute,omitempty"`
					Cors *struct {
						AllowCredentials *bool     `tfsdk:"allow_credentials" json:"allowCredentials,omitempty"`
						AllowHeaders     *[]string `tfsdk:"allow_headers" json:"allowHeaders,omitempty"`
						AllowMethods     *[]string `tfsdk:"allow_methods" json:"allowMethods,omitempty"`
						AllowOrigin      *[]string `tfsdk:"allow_origin" json:"allowOrigin,omitempty"`
						AllowOriginRegex *[]string `tfsdk:"allow_origin_regex" json:"allowOriginRegex,omitempty"`
						DisableForRoute  *bool     `tfsdk:"disable_for_route" json:"disableForRoute,omitempty"`
						ExposeHeaders    *[]string `tfsdk:"expose_headers" json:"exposeHeaders,omitempty"`
						MaxAge           *string   `tfsdk:"max_age" json:"maxAge,omitempty"`
					} `tfsdk:"cors" json:"cors,omitempty"`
					Csrf *struct {
						AdditionalOrigins *[]struct {
							Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
							IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
							Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
							SafeRegex  *struct {
								GoogleRe2 *struct {
									MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
								} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
								Regex *string `tfsdk:"regex" json:"regex,omitempty"`
							} `tfsdk:"safe_regex" json:"safeRegex,omitempty"`
							Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
						} `tfsdk:"additional_origins" json:"additionalOrigins,omitempty"`
						FilterEnabled *struct {
							DefaultValue *struct {
								Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
								Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
							} `tfsdk:"default_value" json:"defaultValue,omitempty"`
							RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
						} `tfsdk:"filter_enabled" json:"filterEnabled,omitempty"`
						ShadowEnabled *struct {
							DefaultValue *struct {
								Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
								Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
							} `tfsdk:"default_value" json:"defaultValue,omitempty"`
							RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
						} `tfsdk:"shadow_enabled" json:"shadowEnabled,omitempty"`
					} `tfsdk:"csrf" json:"csrf,omitempty"`
					Dlp *struct {
						Actions *[]struct {
							ActionType   *string `tfsdk:"action_type" json:"actionType,omitempty"`
							CustomAction *struct {
								MaskChar *string `tfsdk:"mask_char" json:"maskChar,omitempty"`
								Name     *string `tfsdk:"name" json:"name,omitempty"`
								Percent  *struct {
									Value *float64 `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"percent" json:"percent,omitempty"`
								Regex        *[]string `tfsdk:"regex" json:"regex,omitempty"`
								RegexActions *[]struct {
									Regex    *string `tfsdk:"regex" json:"regex,omitempty"`
									Subgroup *int64  `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"regex_actions" json:"regexActions,omitempty"`
							} `tfsdk:"custom_action" json:"customAction,omitempty"`
							KeyValueAction *struct {
								KeyToMask *string `tfsdk:"key_to_mask" json:"keyToMask,omitempty"`
								MaskChar  *string `tfsdk:"mask_char" json:"maskChar,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Percent   *struct {
									Value *float64 `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"percent" json:"percent,omitempty"`
							} `tfsdk:"key_value_action" json:"keyValueAction,omitempty"`
							Shadow *bool `tfsdk:"shadow" json:"shadow,omitempty"`
						} `tfsdk:"actions" json:"actions,omitempty"`
						EnabledFor *string `tfsdk:"enabled_for" json:"enabledFor,omitempty"`
					} `tfsdk:"dlp" json:"dlp,omitempty"`
					EnvoyMetadata *map[string]string `tfsdk:"envoy_metadata" json:"envoyMetadata,omitempty"`
					ExtProc       *struct {
						Disabled  *bool `tfsdk:"disabled" json:"disabled,omitempty"`
						Overrides *struct {
							AsyncMode   *bool `tfsdk:"async_mode" json:"asyncMode,omitempty"`
							GrpcService *struct {
								Authority        *string `tfsdk:"authority" json:"authority,omitempty"`
								ExtProcServerRef *struct {
									Name      *string `tfsdk:"name" json:"name,omitempty"`
									Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"ext_proc_server_ref" json:"extProcServerRef,omitempty"`
								InitialMetadata *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"initial_metadata" json:"initialMetadata,omitempty"`
								RetryPolicy *struct {
									NumRetries   *int64 `tfsdk:"num_retries" json:"numRetries,omitempty"`
									RetryBackOff *struct {
										BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
										MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
									} `tfsdk:"retry_back_off" json:"retryBackOff,omitempty"`
								} `tfsdk:"retry_policy" json:"retryPolicy,omitempty"`
								Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
							} `tfsdk:"grpc_service" json:"grpcService,omitempty"`
							MetadataContextNamespaces *[]string `tfsdk:"metadata_context_namespaces" json:"metadataContextNamespaces,omitempty"`
							ProcessingMode            *struct {
								RequestBodyMode     *string `tfsdk:"request_body_mode" json:"requestBodyMode,omitempty"`
								RequestHeaderMode   *string `tfsdk:"request_header_mode" json:"requestHeaderMode,omitempty"`
								RequestTrailerMode  *string `tfsdk:"request_trailer_mode" json:"requestTrailerMode,omitempty"`
								ResponseBodyMode    *string `tfsdk:"response_body_mode" json:"responseBodyMode,omitempty"`
								ResponseHeaderMode  *string `tfsdk:"response_header_mode" json:"responseHeaderMode,omitempty"`
								ResponseTrailerMode *string `tfsdk:"response_trailer_mode" json:"responseTrailerMode,omitempty"`
							} `tfsdk:"processing_mode" json:"processingMode,omitempty"`
							RequestAttributes              *[]string `tfsdk:"request_attributes" json:"requestAttributes,omitempty"`
							ResponseAttributes             *[]string `tfsdk:"response_attributes" json:"responseAttributes,omitempty"`
							TypedMetadataContextNamespaces *[]string `tfsdk:"typed_metadata_context_namespaces" json:"typedMetadataContextNamespaces,omitempty"`
						} `tfsdk:"overrides" json:"overrides,omitempty"`
					} `tfsdk:"ext_proc" json:"extProc,omitempty"`
					Extauth *struct {
						ConfigRef *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"config_ref" json:"configRef,omitempty"`
						CustomAuth *struct {
							ContextExtensions *map[string]string `tfsdk:"context_extensions" json:"contextExtensions,omitempty"`
							Name              *string            `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"custom_auth" json:"customAuth,omitempty"`
						Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
					} `tfsdk:"extauth" json:"extauth,omitempty"`
					Extensions *struct {
						Configs *map[string]string `tfsdk:"configs" json:"configs,omitempty"`
					} `tfsdk:"extensions" json:"extensions,omitempty"`
					Faults *struct {
						Abort *struct {
							HttpStatus *int64   `tfsdk:"http_status" json:"httpStatus,omitempty"`
							Percentage *float64 `tfsdk:"percentage" json:"percentage,omitempty"`
						} `tfsdk:"abort" json:"abort,omitempty"`
						Delay *struct {
							FixedDelay *string  `tfsdk:"fixed_delay" json:"fixedDelay,omitempty"`
							Percentage *float64 `tfsdk:"percentage" json:"percentage,omitempty"`
						} `tfsdk:"delay" json:"delay,omitempty"`
					} `tfsdk:"faults" json:"faults,omitempty"`
					HeaderManipulation *struct {
						RequestHeadersToAdd *[]struct {
							Append *bool `tfsdk:"append" json:"append,omitempty"`
							Header *struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"header" json:"header,omitempty"`
							HeaderSecretRef *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"header_secret_ref" json:"headerSecretRef,omitempty"`
						} `tfsdk:"request_headers_to_add" json:"requestHeadersToAdd,omitempty"`
						RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" json:"requestHeadersToRemove,omitempty"`
						ResponseHeadersToAdd   *[]struct {
							Append *bool `tfsdk:"append" json:"append,omitempty"`
							Header *struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"header" json:"header,omitempty"`
						} `tfsdk:"response_headers_to_add" json:"responseHeadersToAdd,omitempty"`
						ResponseHeadersToRemove *[]string `tfsdk:"response_headers_to_remove" json:"responseHeadersToRemove,omitempty"`
					} `tfsdk:"header_manipulation" json:"headerManipulation,omitempty"`
					HostRewrite          *string `tfsdk:"host_rewrite" json:"hostRewrite,omitempty"`
					HostRewritePathRegex *struct {
						Pattern *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"pattern" json:"pattern,omitempty"`
						Substitution *string `tfsdk:"substitution" json:"substitution,omitempty"`
					} `tfsdk:"host_rewrite_path_regex" json:"hostRewritePathRegex,omitempty"`
					IdleTimeout *string `tfsdk:"idle_timeout" json:"idleTimeout,omitempty"`
					Jwt         *struct {
						Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
					} `tfsdk:"jwt" json:"jwt,omitempty"`
					JwtStaged *struct {
						AfterExtAuth *struct {
							Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
						} `tfsdk:"after_ext_auth" json:"afterExtAuth,omitempty"`
						BeforeExtAuth *struct {
							Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
						} `tfsdk:"before_ext_auth" json:"beforeExtAuth,omitempty"`
					} `tfsdk:"jwt_staged" json:"jwtStaged,omitempty"`
					LbHash *struct {
						HashPolicies *[]struct {
							Cookie *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
								Ttl  *string `tfsdk:"ttl" json:"ttl,omitempty"`
							} `tfsdk:"cookie" json:"cookie,omitempty"`
							Header   *string `tfsdk:"header" json:"header,omitempty"`
							SourceIp *bool   `tfsdk:"source_ip" json:"sourceIp,omitempty"`
							Terminal *bool   `tfsdk:"terminal" json:"terminal,omitempty"`
						} `tfsdk:"hash_policies" json:"hashPolicies,omitempty"`
					} `tfsdk:"lb_hash" json:"lbHash,omitempty"`
					MaxStreamDuration *struct {
						GrpcTimeoutHeaderMax    *string `tfsdk:"grpc_timeout_header_max" json:"grpcTimeoutHeaderMax,omitempty"`
						GrpcTimeoutHeaderOffset *string `tfsdk:"grpc_timeout_header_offset" json:"grpcTimeoutHeaderOffset,omitempty"`
						MaxStreamDuration       *string `tfsdk:"max_stream_duration" json:"maxStreamDuration,omitempty"`
					} `tfsdk:"max_stream_duration" json:"maxStreamDuration,omitempty"`
					PrefixRewrite    *string `tfsdk:"prefix_rewrite" json:"prefixRewrite,omitempty"`
					RateLimitConfigs *struct {
						Refs *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"refs" json:"refs,omitempty"`
					} `tfsdk:"rate_limit_configs" json:"rateLimitConfigs,omitempty"`
					RateLimitEarlyConfigs *struct {
						Refs *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"refs" json:"refs,omitempty"`
					} `tfsdk:"rate_limit_early_configs" json:"rateLimitEarlyConfigs,omitempty"`
					RateLimitRegularConfigs *struct {
						Refs *[]struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"refs" json:"refs,omitempty"`
					} `tfsdk:"rate_limit_regular_configs" json:"rateLimitRegularConfigs,omitempty"`
					Ratelimit *struct {
						IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" json:"includeVhRateLimits,omitempty"`
						LocalRatelimit      *struct {
							FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
							MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
							TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
						} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
						RateLimits *[]struct {
							Actions *[]struct {
								DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
								GenericKey         *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								} `tfsdk:"generic_key" json:"genericKey,omitempty"`
								HeaderValueMatch *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
									ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
									Headers         *[]struct {
										ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
										InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
										Name         *string `tfsdk:"name" json:"name,omitempty"`
										PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
										PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
										RangeMatch   *struct {
											End   *int64 `tfsdk:"end" json:"end,omitempty"`
											Start *int64 `tfsdk:"start" json:"start,omitempty"`
										} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
										RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
										SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
								} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
								Metadata *struct {
									DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									MetadataKey   *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Path *[]struct {
											Key *string `tfsdk:"key" json:"key,omitempty"`
										} `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
									Source *string `tfsdk:"source" json:"source,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
								RequestHeaders *struct {
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
								} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
								SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
							} `tfsdk:"actions" json:"actions,omitempty"`
							SetActions *[]struct {
								DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
								GenericKey         *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								} `tfsdk:"generic_key" json:"genericKey,omitempty"`
								HeaderValueMatch *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
									ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
									Headers         *[]struct {
										ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
										InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
										Name         *string `tfsdk:"name" json:"name,omitempty"`
										PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
										PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
										RangeMatch   *struct {
											End   *int64 `tfsdk:"end" json:"end,omitempty"`
											Start *int64 `tfsdk:"start" json:"start,omitempty"`
										} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
										RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
										SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
								} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
								Metadata *struct {
									DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									MetadataKey   *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Path *[]struct {
											Key *string `tfsdk:"key" json:"key,omitempty"`
										} `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
									Source *string `tfsdk:"source" json:"source,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
								RequestHeaders *struct {
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
								} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
								SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
							} `tfsdk:"set_actions" json:"setActions,omitempty"`
						} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
					} `tfsdk:"ratelimit" json:"ratelimit,omitempty"`
					RatelimitBasic *struct {
						AnonymousLimits *struct {
							RequestsPerUnit *int64  `tfsdk:"requests_per_unit" json:"requestsPerUnit,omitempty"`
							Unit            *string `tfsdk:"unit" json:"unit,omitempty"`
						} `tfsdk:"anonymous_limits" json:"anonymousLimits,omitempty"`
						AuthorizedLimits *struct {
							RequestsPerUnit *int64  `tfsdk:"requests_per_unit" json:"requestsPerUnit,omitempty"`
							Unit            *string `tfsdk:"unit" json:"unit,omitempty"`
						} `tfsdk:"authorized_limits" json:"authorizedLimits,omitempty"`
					} `tfsdk:"ratelimit_basic" json:"ratelimitBasic,omitempty"`
					RatelimitEarly *struct {
						IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" json:"includeVhRateLimits,omitempty"`
						LocalRatelimit      *struct {
							FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
							MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
							TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
						} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
						RateLimits *[]struct {
							Actions *[]struct {
								DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
								GenericKey         *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								} `tfsdk:"generic_key" json:"genericKey,omitempty"`
								HeaderValueMatch *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
									ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
									Headers         *[]struct {
										ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
										InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
										Name         *string `tfsdk:"name" json:"name,omitempty"`
										PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
										PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
										RangeMatch   *struct {
											End   *int64 `tfsdk:"end" json:"end,omitempty"`
											Start *int64 `tfsdk:"start" json:"start,omitempty"`
										} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
										RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
										SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
								} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
								Metadata *struct {
									DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									MetadataKey   *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Path *[]struct {
											Key *string `tfsdk:"key" json:"key,omitempty"`
										} `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
									Source *string `tfsdk:"source" json:"source,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
								RequestHeaders *struct {
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
								} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
								SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
							} `tfsdk:"actions" json:"actions,omitempty"`
							SetActions *[]struct {
								DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
								GenericKey         *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								} `tfsdk:"generic_key" json:"genericKey,omitempty"`
								HeaderValueMatch *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
									ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
									Headers         *[]struct {
										ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
										InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
										Name         *string `tfsdk:"name" json:"name,omitempty"`
										PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
										PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
										RangeMatch   *struct {
											End   *int64 `tfsdk:"end" json:"end,omitempty"`
											Start *int64 `tfsdk:"start" json:"start,omitempty"`
										} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
										RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
										SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
								} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
								Metadata *struct {
									DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									MetadataKey   *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Path *[]struct {
											Key *string `tfsdk:"key" json:"key,omitempty"`
										} `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
									Source *string `tfsdk:"source" json:"source,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
								RequestHeaders *struct {
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
								} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
								SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
							} `tfsdk:"set_actions" json:"setActions,omitempty"`
						} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
					} `tfsdk:"ratelimit_early" json:"ratelimitEarly,omitempty"`
					RatelimitRegular *struct {
						IncludeVhRateLimits *bool `tfsdk:"include_vh_rate_limits" json:"includeVhRateLimits,omitempty"`
						LocalRatelimit      *struct {
							FillInterval  *string `tfsdk:"fill_interval" json:"fillInterval,omitempty"`
							MaxTokens     *int64  `tfsdk:"max_tokens" json:"maxTokens,omitempty"`
							TokensPerFill *int64  `tfsdk:"tokens_per_fill" json:"tokensPerFill,omitempty"`
						} `tfsdk:"local_ratelimit" json:"localRatelimit,omitempty"`
						RateLimits *[]struct {
							Actions *[]struct {
								DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
								GenericKey         *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								} `tfsdk:"generic_key" json:"genericKey,omitempty"`
								HeaderValueMatch *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
									ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
									Headers         *[]struct {
										ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
										InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
										Name         *string `tfsdk:"name" json:"name,omitempty"`
										PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
										PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
										RangeMatch   *struct {
											End   *int64 `tfsdk:"end" json:"end,omitempty"`
											Start *int64 `tfsdk:"start" json:"start,omitempty"`
										} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
										RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
										SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
								} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
								Metadata *struct {
									DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									MetadataKey   *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Path *[]struct {
											Key *string `tfsdk:"key" json:"key,omitempty"`
										} `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
									Source *string `tfsdk:"source" json:"source,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
								RequestHeaders *struct {
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
								} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
								SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
							} `tfsdk:"actions" json:"actions,omitempty"`
							SetActions *[]struct {
								DestinationCluster *map[string]string `tfsdk:"destination_cluster" json:"destinationCluster,omitempty"`
								GenericKey         *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
								} `tfsdk:"generic_key" json:"genericKey,omitempty"`
								HeaderValueMatch *struct {
									DescriptorValue *string `tfsdk:"descriptor_value" json:"descriptorValue,omitempty"`
									ExpectMatch     *bool   `tfsdk:"expect_match" json:"expectMatch,omitempty"`
									Headers         *[]struct {
										ExactMatch   *string `tfsdk:"exact_match" json:"exactMatch,omitempty"`
										InvertMatch  *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
										Name         *string `tfsdk:"name" json:"name,omitempty"`
										PrefixMatch  *string `tfsdk:"prefix_match" json:"prefixMatch,omitempty"`
										PresentMatch *bool   `tfsdk:"present_match" json:"presentMatch,omitempty"`
										RangeMatch   *struct {
											End   *int64 `tfsdk:"end" json:"end,omitempty"`
											Start *int64 `tfsdk:"start" json:"start,omitempty"`
										} `tfsdk:"range_match" json:"rangeMatch,omitempty"`
										RegexMatch  *string `tfsdk:"regex_match" json:"regexMatch,omitempty"`
										SuffixMatch *string `tfsdk:"suffix_match" json:"suffixMatch,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
								} `tfsdk:"header_value_match" json:"headerValueMatch,omitempty"`
								Metadata *struct {
									DefaultValue  *string `tfsdk:"default_value" json:"defaultValue,omitempty"`
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									MetadataKey   *struct {
										Key  *string `tfsdk:"key" json:"key,omitempty"`
										Path *[]struct {
											Key *string `tfsdk:"key" json:"key,omitempty"`
										} `tfsdk:"path" json:"path,omitempty"`
									} `tfsdk:"metadata_key" json:"metadataKey,omitempty"`
									Source *string `tfsdk:"source" json:"source,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								RemoteAddress  *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
								RequestHeaders *struct {
									DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
									HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
								} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
								SourceCluster *map[string]string `tfsdk:"source_cluster" json:"sourceCluster,omitempty"`
							} `tfsdk:"set_actions" json:"setActions,omitempty"`
						} `tfsdk:"rate_limits" json:"rateLimits,omitempty"`
					} `tfsdk:"ratelimit_regular" json:"ratelimitRegular,omitempty"`
					Rbac *struct {
						Disable  *bool `tfsdk:"disable" json:"disable,omitempty"`
						Policies *struct {
							NestedClaimDelimiter *string `tfsdk:"nested_claim_delimiter" json:"nestedClaimDelimiter,omitempty"`
							Permissions          *struct {
								Methods    *[]string `tfsdk:"methods" json:"methods,omitempty"`
								PathPrefix *string   `tfsdk:"path_prefix" json:"pathPrefix,omitempty"`
							} `tfsdk:"permissions" json:"permissions,omitempty"`
							Principals *[]struct {
								JwtPrincipal *struct {
									Claims   *map[string]string `tfsdk:"claims" json:"claims,omitempty"`
									Matcher  *string            `tfsdk:"matcher" json:"matcher,omitempty"`
									Provider *string            `tfsdk:"provider" json:"provider,omitempty"`
								} `tfsdk:"jwt_principal" json:"jwtPrincipal,omitempty"`
							} `tfsdk:"principals" json:"principals,omitempty"`
						} `tfsdk:"policies" json:"policies,omitempty"`
					} `tfsdk:"rbac" json:"rbac,omitempty"`
					RegexRewrite *struct {
						Pattern *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"pattern" json:"pattern,omitempty"`
						Substitution *string `tfsdk:"substitution" json:"substitution,omitempty"`
					} `tfsdk:"regex_rewrite" json:"regexRewrite,omitempty"`
					Retries *struct {
						NumRetries    *int64  `tfsdk:"num_retries" json:"numRetries,omitempty"`
						PerTryTimeout *string `tfsdk:"per_try_timeout" json:"perTryTimeout,omitempty"`
						RetryBackOff  *struct {
							BaseInterval *string `tfsdk:"base_interval" json:"baseInterval,omitempty"`
							MaxInterval  *string `tfsdk:"max_interval" json:"maxInterval,omitempty"`
						} `tfsdk:"retry_back_off" json:"retryBackOff,omitempty"`
						RetryOn *string `tfsdk:"retry_on" json:"retryOn,omitempty"`
					} `tfsdk:"retries" json:"retries,omitempty"`
					Shadowing *struct {
						Percentage *float64 `tfsdk:"percentage" json:"percentage,omitempty"`
						Upstream   *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"upstream" json:"upstream,omitempty"`
					} `tfsdk:"shadowing" json:"shadowing,omitempty"`
					StagedTransformations *struct {
						Early *struct {
							RequestTransforms *[]struct {
								ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
								Matcher         *struct {
									CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
									ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
									Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
									Headers        *[]struct {
										InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
										Name        *string `tfsdk:"name" json:"name,omitempty"`
										Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
										Value       *string `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
									Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
									QueryParameters *[]struct {
										Name  *string `tfsdk:"name" json:"name,omitempty"`
										Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
										Value *string `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
									Regex *string `tfsdk:"regex" json:"regex,omitempty"`
								} `tfsdk:"matcher" json:"matcher,omitempty"`
								RequestTransformation *struct {
									HeaderBodyTransform *struct {
										AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
									} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
									LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
									TransformationTemplate *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
										Body              *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"body" json:"body,omitempty"`
										DynamicMetadataValues *[]struct {
											Key               *string `tfsdk:"key" json:"key,omitempty"`
											MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
											Value             *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
										EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
										Extractors       *struct {
											Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
											Header   *string            `tfsdk:"header" json:"header,omitempty"`
											Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
											Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
										} `tfsdk:"extractors" json:"extractors,omitempty"`
										Headers *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"headers" json:"headers,omitempty"`
										HeadersToAppend *[]struct {
											Key   *string `tfsdk:"key" json:"key,omitempty"`
											Value *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
										HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
										IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
										ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
										Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
									} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
									XsltTransformation *struct {
										NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
										SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
										Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
									} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
								} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
								ResponseTransformation *struct {
									HeaderBodyTransform *struct {
										AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
									} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
									LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
									TransformationTemplate *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
										Body              *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"body" json:"body,omitempty"`
										DynamicMetadataValues *[]struct {
											Key               *string `tfsdk:"key" json:"key,omitempty"`
											MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
											Value             *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
										EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
										Extractors       *struct {
											Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
											Header   *string            `tfsdk:"header" json:"header,omitempty"`
											Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
											Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
										} `tfsdk:"extractors" json:"extractors,omitempty"`
										Headers *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"headers" json:"headers,omitempty"`
										HeadersToAppend *[]struct {
											Key   *string `tfsdk:"key" json:"key,omitempty"`
											Value *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
										HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
										IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
										ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
										Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
									} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
									XsltTransformation *struct {
										NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
										SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
										Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
									} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
								} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
							} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
							ResponseTransforms *[]struct {
								Matchers *[]struct {
									InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name        *string `tfsdk:"name" json:"name,omitempty"`
									Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
									Value       *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"matchers" json:"matchers,omitempty"`
								ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
								ResponseTransformation *struct {
									HeaderBodyTransform *struct {
										AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
									} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
									LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
									TransformationTemplate *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
										Body              *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"body" json:"body,omitempty"`
										DynamicMetadataValues *[]struct {
											Key               *string `tfsdk:"key" json:"key,omitempty"`
											MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
											Value             *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
										EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
										Extractors       *struct {
											Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
											Header   *string            `tfsdk:"header" json:"header,omitempty"`
											Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
											Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
										} `tfsdk:"extractors" json:"extractors,omitempty"`
										Headers *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"headers" json:"headers,omitempty"`
										HeadersToAppend *[]struct {
											Key   *string `tfsdk:"key" json:"key,omitempty"`
											Value *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
										HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
										IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
										ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
										Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
									} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
									XsltTransformation *struct {
										NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
										SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
										Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
									} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
								} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
							} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
						} `tfsdk:"early" json:"early,omitempty"`
						EscapeCharacters       *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
						InheritTransformation  *bool `tfsdk:"inherit_transformation" json:"inheritTransformation,omitempty"`
						LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
						Regular                *struct {
							RequestTransforms *[]struct {
								ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
								Matcher         *struct {
									CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
									ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
									Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
									Headers        *[]struct {
										InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
										Name        *string `tfsdk:"name" json:"name,omitempty"`
										Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
										Value       *string `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
									Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
									QueryParameters *[]struct {
										Name  *string `tfsdk:"name" json:"name,omitempty"`
										Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
										Value *string `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
									Regex *string `tfsdk:"regex" json:"regex,omitempty"`
								} `tfsdk:"matcher" json:"matcher,omitempty"`
								RequestTransformation *struct {
									HeaderBodyTransform *struct {
										AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
									} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
									LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
									TransformationTemplate *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
										Body              *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"body" json:"body,omitempty"`
										DynamicMetadataValues *[]struct {
											Key               *string `tfsdk:"key" json:"key,omitempty"`
											MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
											Value             *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
										EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
										Extractors       *struct {
											Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
											Header   *string            `tfsdk:"header" json:"header,omitempty"`
											Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
											Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
										} `tfsdk:"extractors" json:"extractors,omitempty"`
										Headers *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"headers" json:"headers,omitempty"`
										HeadersToAppend *[]struct {
											Key   *string `tfsdk:"key" json:"key,omitempty"`
											Value *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
										HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
										IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
										ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
										Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
									} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
									XsltTransformation *struct {
										NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
										SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
										Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
									} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
								} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
								ResponseTransformation *struct {
									HeaderBodyTransform *struct {
										AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
									} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
									LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
									TransformationTemplate *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
										Body              *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"body" json:"body,omitempty"`
										DynamicMetadataValues *[]struct {
											Key               *string `tfsdk:"key" json:"key,omitempty"`
											MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
											Value             *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
										EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
										Extractors       *struct {
											Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
											Header   *string            `tfsdk:"header" json:"header,omitempty"`
											Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
											Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
										} `tfsdk:"extractors" json:"extractors,omitempty"`
										Headers *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"headers" json:"headers,omitempty"`
										HeadersToAppend *[]struct {
											Key   *string `tfsdk:"key" json:"key,omitempty"`
											Value *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
										HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
										IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
										ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
										Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
									} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
									XsltTransformation *struct {
										NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
										SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
										Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
									} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
								} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
							} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
							ResponseTransforms *[]struct {
								Matchers *[]struct {
									InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
									Name        *string `tfsdk:"name" json:"name,omitempty"`
									Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
									Value       *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"matchers" json:"matchers,omitempty"`
								ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
								ResponseTransformation *struct {
									HeaderBodyTransform *struct {
										AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
									} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
									LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
									TransformationTemplate *struct {
										AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
										Body              *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"body" json:"body,omitempty"`
										DynamicMetadataValues *[]struct {
											Key               *string `tfsdk:"key" json:"key,omitempty"`
											MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
											Value             *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
										EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
										Extractors       *struct {
											Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
											Header   *string            `tfsdk:"header" json:"header,omitempty"`
											Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
											Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
										} `tfsdk:"extractors" json:"extractors,omitempty"`
										Headers *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"headers" json:"headers,omitempty"`
										HeadersToAppend *[]struct {
											Key   *string `tfsdk:"key" json:"key,omitempty"`
											Value *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
										HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
										IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
										MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
										ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
										Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
									} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
									XsltTransformation *struct {
										NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
										SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
										Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
									} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
								} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
							} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
						} `tfsdk:"regular" json:"regular,omitempty"`
					} `tfsdk:"staged_transformations" json:"stagedTransformations,omitempty"`
					Timeout *string `tfsdk:"timeout" json:"timeout,omitempty"`
					Tracing *struct {
						Propagate        *bool   `tfsdk:"propagate" json:"propagate,omitempty"`
						RouteDescriptor  *string `tfsdk:"route_descriptor" json:"routeDescriptor,omitempty"`
						TracePercentages *struct {
							ClientSamplePercentage  *float64 `tfsdk:"client_sample_percentage" json:"clientSamplePercentage,omitempty"`
							OverallSamplePercentage *float64 `tfsdk:"overall_sample_percentage" json:"overallSamplePercentage,omitempty"`
							RandomSamplePercentage  *float64 `tfsdk:"random_sample_percentage" json:"randomSamplePercentage,omitempty"`
						} `tfsdk:"trace_percentages" json:"tracePercentages,omitempty"`
					} `tfsdk:"tracing" json:"tracing,omitempty"`
					Transformations *struct {
						ClearRouteCache       *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
						RequestTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header   *string            `tfsdk:"header" json:"header,omitempty"`
									Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
									Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
						ResponseTransformation *struct {
							HeaderBodyTransform *struct {
								AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
							} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
							LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
							TransformationTemplate *struct {
								AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
								Body              *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"body" json:"body,omitempty"`
								DynamicMetadataValues *[]struct {
									Key               *string `tfsdk:"key" json:"key,omitempty"`
									MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
									Value             *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
								EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
								Extractors       *struct {
									Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
									Header   *string            `tfsdk:"header" json:"header,omitempty"`
									Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
									Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
								} `tfsdk:"extractors" json:"extractors,omitempty"`
								Headers *struct {
									Text *string `tfsdk:"text" json:"text,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								HeadersToAppend *[]struct {
									Key   *string `tfsdk:"key" json:"key,omitempty"`
									Value *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
								HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
								IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
								MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
								ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
								Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
							} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
							XsltTransformation *struct {
								NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
								SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
								Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
							} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
						} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
					} `tfsdk:"transformations" json:"transformations,omitempty"`
					Upgrades *[]struct {
						Connect *struct {
							Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
						} `tfsdk:"connect" json:"connect,omitempty"`
						Websocket *struct {
							Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
						} `tfsdk:"websocket" json:"websocket,omitempty"`
					} `tfsdk:"upgrades" json:"upgrades,omitempty"`
					Waf *struct {
						AuditLogging *struct {
							Action   *string `tfsdk:"action" json:"action,omitempty"`
							Location *string `tfsdk:"location" json:"location,omitempty"`
						} `tfsdk:"audit_logging" json:"auditLogging,omitempty"`
						ConfigMapRuleSets *[]struct {
							ConfigMapRef *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
							DataMapKeys *[]string `tfsdk:"data_map_keys" json:"dataMapKeys,omitempty"`
						} `tfsdk:"config_map_rule_sets" json:"configMapRuleSets,omitempty"`
						CoreRuleSet *struct {
							CustomSettingsFile   *string `tfsdk:"custom_settings_file" json:"customSettingsFile,omitempty"`
							CustomSettingsString *string `tfsdk:"custom_settings_string" json:"customSettingsString,omitempty"`
						} `tfsdk:"core_rule_set" json:"coreRuleSet,omitempty"`
						CustomInterventionMessage *string `tfsdk:"custom_intervention_message" json:"customInterventionMessage,omitempty"`
						Disabled                  *bool   `tfsdk:"disabled" json:"disabled,omitempty"`
						RequestHeadersOnly        *bool   `tfsdk:"request_headers_only" json:"requestHeadersOnly,omitempty"`
						ResponseHeadersOnly       *bool   `tfsdk:"response_headers_only" json:"responseHeadersOnly,omitempty"`
						RuleSets                  *[]struct {
							Directory *string   `tfsdk:"directory" json:"directory,omitempty"`
							Files     *[]string `tfsdk:"files" json:"files,omitempty"`
							RuleStr   *string   `tfsdk:"rule_str" json:"ruleStr,omitempty"`
						} `tfsdk:"rule_sets" json:"ruleSets,omitempty"`
					} `tfsdk:"waf" json:"waf,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				OptionsConfigRefs *struct {
					DelegateOptions *[]struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"delegate_options" json:"delegateOptions,omitempty"`
				} `tfsdk:"options_config_refs" json:"optionsConfigRefs,omitempty"`
				RedirectAction *struct {
					HostRedirect  *string `tfsdk:"host_redirect" json:"hostRedirect,omitempty"`
					HttpsRedirect *bool   `tfsdk:"https_redirect" json:"httpsRedirect,omitempty"`
					PathRedirect  *string `tfsdk:"path_redirect" json:"pathRedirect,omitempty"`
					PrefixRewrite *string `tfsdk:"prefix_rewrite" json:"prefixRewrite,omitempty"`
					RegexRewrite  *struct {
						Pattern *struct {
							GoogleRe2 *struct {
								MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
							} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
							Regex *string `tfsdk:"regex" json:"regex,omitempty"`
						} `tfsdk:"pattern" json:"pattern,omitempty"`
						Substitution *string `tfsdk:"substitution" json:"substitution,omitempty"`
					} `tfsdk:"regex_rewrite" json:"regexRewrite,omitempty"`
					ResponseCode *string `tfsdk:"response_code" json:"responseCode,omitempty"`
					StripQuery   *bool   `tfsdk:"strip_query" json:"stripQuery,omitempty"`
				} `tfsdk:"redirect_action" json:"redirectAction,omitempty"`
				RouteAction *struct {
					ClusterHeader       *string `tfsdk:"cluster_header" json:"clusterHeader,omitempty"`
					DynamicForwardProxy *struct {
						AutoHostRewriteHeader *string `tfsdk:"auto_host_rewrite_header" json:"autoHostRewriteHeader,omitempty"`
						HostRewrite           *string `tfsdk:"host_rewrite" json:"hostRewrite,omitempty"`
					} `tfsdk:"dynamic_forward_proxy" json:"dynamicForwardProxy,omitempty"`
					Multi *struct {
						Destinations *[]struct {
							Destination *struct {
								Consul *struct {
									DataCenters *[]string `tfsdk:"data_centers" json:"dataCenters,omitempty"`
									ServiceName *string   `tfsdk:"service_name" json:"serviceName,omitempty"`
									Tags        *[]string `tfsdk:"tags" json:"tags,omitempty"`
								} `tfsdk:"consul" json:"consul,omitempty"`
								DestinationSpec *struct {
									Aws *struct {
										InvocationStyle        *string `tfsdk:"invocation_style" json:"invocationStyle,omitempty"`
										LogicalName            *string `tfsdk:"logical_name" json:"logicalName,omitempty"`
										RequestTransformation  *bool   `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
										ResponseTransformation *bool   `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
										UnwrapAsAlb            *bool   `tfsdk:"unwrap_as_alb" json:"unwrapAsAlb,omitempty"`
										UnwrapAsApiGateway     *bool   `tfsdk:"unwrap_as_api_gateway" json:"unwrapAsApiGateway,omitempty"`
										WrapAsApiGateway       *bool   `tfsdk:"wrap_as_api_gateway" json:"wrapAsApiGateway,omitempty"`
									} `tfsdk:"aws" json:"aws,omitempty"`
									Azure *struct {
										FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
									} `tfsdk:"azure" json:"azure,omitempty"`
									Grpc *struct {
										Function   *string `tfsdk:"function" json:"function,omitempty"`
										Package    *string `tfsdk:"package" json:"package,omitempty"`
										Parameters *struct {
											Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
											Path    *string            `tfsdk:"path" json:"path,omitempty"`
										} `tfsdk:"parameters" json:"parameters,omitempty"`
										Service *string `tfsdk:"service" json:"service,omitempty"`
									} `tfsdk:"grpc" json:"grpc,omitempty"`
									Rest *struct {
										FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
										Parameters   *struct {
											Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
											Path    *string            `tfsdk:"path" json:"path,omitempty"`
										} `tfsdk:"parameters" json:"parameters,omitempty"`
										ResponseTransformation *struct {
											AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
											Body              *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"body" json:"body,omitempty"`
											DynamicMetadataValues *[]struct {
												Key               *string `tfsdk:"key" json:"key,omitempty"`
												MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
												Value             *struct {
													Text *string `tfsdk:"text" json:"text,omitempty"`
												} `tfsdk:"value" json:"value,omitempty"`
											} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
											EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
											Extractors       *struct {
												Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
												Header   *string            `tfsdk:"header" json:"header,omitempty"`
												Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
												Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
											} `tfsdk:"extractors" json:"extractors,omitempty"`
											Headers *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"headers" json:"headers,omitempty"`
											HeadersToAppend *[]struct {
												Key   *string `tfsdk:"key" json:"key,omitempty"`
												Value *struct {
													Text *string `tfsdk:"text" json:"text,omitempty"`
												} `tfsdk:"value" json:"value,omitempty"`
											} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
											HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
											IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
											MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
											ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
											Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
										} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
									} `tfsdk:"rest" json:"rest,omitempty"`
								} `tfsdk:"destination_spec" json:"destinationSpec,omitempty"`
								Kube *struct {
									Port *int64 `tfsdk:"port" json:"port,omitempty"`
									Ref  *struct {
										Name      *string `tfsdk:"name" json:"name,omitempty"`
										Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
									} `tfsdk:"ref" json:"ref,omitempty"`
								} `tfsdk:"kube" json:"kube,omitempty"`
								Subset *struct {
									Values *map[string]string `tfsdk:"values" json:"values,omitempty"`
								} `tfsdk:"subset" json:"subset,omitempty"`
								Upstream *struct {
									Name      *string `tfsdk:"name" json:"name,omitempty"`
									Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"upstream" json:"upstream,omitempty"`
							} `tfsdk:"destination" json:"destination,omitempty"`
							Options *struct {
								BufferPerRoute *struct {
									Buffer *struct {
										MaxRequestBytes *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
									} `tfsdk:"buffer" json:"buffer,omitempty"`
									Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
								} `tfsdk:"buffer_per_route" json:"bufferPerRoute,omitempty"`
								Csrf *struct {
									AdditionalOrigins *[]struct {
										Exact      *string `tfsdk:"exact" json:"exact,omitempty"`
										IgnoreCase *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
										Prefix     *string `tfsdk:"prefix" json:"prefix,omitempty"`
										SafeRegex  *struct {
											GoogleRe2 *struct {
												MaxProgramSize *int64 `tfsdk:"max_program_size" json:"maxProgramSize,omitempty"`
											} `tfsdk:"google_re2" json:"googleRe2,omitempty"`
											Regex *string `tfsdk:"regex" json:"regex,omitempty"`
										} `tfsdk:"safe_regex" json:"safeRegex,omitempty"`
										Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
									} `tfsdk:"additional_origins" json:"additionalOrigins,omitempty"`
									FilterEnabled *struct {
										DefaultValue *struct {
											Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
											Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
										} `tfsdk:"default_value" json:"defaultValue,omitempty"`
										RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
									} `tfsdk:"filter_enabled" json:"filterEnabled,omitempty"`
									ShadowEnabled *struct {
										DefaultValue *struct {
											Denominator *string `tfsdk:"denominator" json:"denominator,omitempty"`
											Numerator   *int64  `tfsdk:"numerator" json:"numerator,omitempty"`
										} `tfsdk:"default_value" json:"defaultValue,omitempty"`
										RuntimeKey *string `tfsdk:"runtime_key" json:"runtimeKey,omitempty"`
									} `tfsdk:"shadow_enabled" json:"shadowEnabled,omitempty"`
								} `tfsdk:"csrf" json:"csrf,omitempty"`
								Extauth *struct {
									ConfigRef *struct {
										Name      *string `tfsdk:"name" json:"name,omitempty"`
										Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
									} `tfsdk:"config_ref" json:"configRef,omitempty"`
									CustomAuth *struct {
										ContextExtensions *map[string]string `tfsdk:"context_extensions" json:"contextExtensions,omitempty"`
										Name              *string            `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"custom_auth" json:"customAuth,omitempty"`
									Disable *bool `tfsdk:"disable" json:"disable,omitempty"`
								} `tfsdk:"extauth" json:"extauth,omitempty"`
								Extensions *struct {
									Configs *map[string]string `tfsdk:"configs" json:"configs,omitempty"`
								} `tfsdk:"extensions" json:"extensions,omitempty"`
								HeaderManipulation *struct {
									RequestHeadersToAdd *[]struct {
										Append *bool `tfsdk:"append" json:"append,omitempty"`
										Header *struct {
											Key   *string `tfsdk:"key" json:"key,omitempty"`
											Value *string `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"header" json:"header,omitempty"`
										HeaderSecretRef *struct {
											Name      *string `tfsdk:"name" json:"name,omitempty"`
											Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
										} `tfsdk:"header_secret_ref" json:"headerSecretRef,omitempty"`
									} `tfsdk:"request_headers_to_add" json:"requestHeadersToAdd,omitempty"`
									RequestHeadersToRemove *[]string `tfsdk:"request_headers_to_remove" json:"requestHeadersToRemove,omitempty"`
									ResponseHeadersToAdd   *[]struct {
										Append *bool `tfsdk:"append" json:"append,omitempty"`
										Header *struct {
											Key   *string `tfsdk:"key" json:"key,omitempty"`
											Value *string `tfsdk:"value" json:"value,omitempty"`
										} `tfsdk:"header" json:"header,omitempty"`
									} `tfsdk:"response_headers_to_add" json:"responseHeadersToAdd,omitempty"`
									ResponseHeadersToRemove *[]string `tfsdk:"response_headers_to_remove" json:"responseHeadersToRemove,omitempty"`
								} `tfsdk:"header_manipulation" json:"headerManipulation,omitempty"`
								StagedTransformations *struct {
									Early *struct {
										RequestTransforms *[]struct {
											ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
											Matcher         *struct {
												CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
												ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
												Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
												Headers        *[]struct {
													InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
													Name        *string `tfsdk:"name" json:"name,omitempty"`
													Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
													Value       *string `tfsdk:"value" json:"value,omitempty"`
												} `tfsdk:"headers" json:"headers,omitempty"`
												Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
												Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
												QueryParameters *[]struct {
													Name  *string `tfsdk:"name" json:"name,omitempty"`
													Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
													Value *string `tfsdk:"value" json:"value,omitempty"`
												} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
												Regex *string `tfsdk:"regex" json:"regex,omitempty"`
											} `tfsdk:"matcher" json:"matcher,omitempty"`
											RequestTransformation *struct {
												HeaderBodyTransform *struct {
													AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
												} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
												LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
												TransformationTemplate *struct {
													AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
													Body              *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"body" json:"body,omitempty"`
													DynamicMetadataValues *[]struct {
														Key               *string `tfsdk:"key" json:"key,omitempty"`
														MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
														Value             *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
													EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
													Extractors       *struct {
														Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
														Header   *string            `tfsdk:"header" json:"header,omitempty"`
														Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
														Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
													} `tfsdk:"extractors" json:"extractors,omitempty"`
													Headers *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"headers" json:"headers,omitempty"`
													HeadersToAppend *[]struct {
														Key   *string `tfsdk:"key" json:"key,omitempty"`
														Value *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
													HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
													IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
													MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
													ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
													Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
												} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
												XsltTransformation *struct {
													NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
													SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
													Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
												} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
											} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
											ResponseTransformation *struct {
												HeaderBodyTransform *struct {
													AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
												} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
												LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
												TransformationTemplate *struct {
													AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
													Body              *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"body" json:"body,omitempty"`
													DynamicMetadataValues *[]struct {
														Key               *string `tfsdk:"key" json:"key,omitempty"`
														MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
														Value             *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
													EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
													Extractors       *struct {
														Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
														Header   *string            `tfsdk:"header" json:"header,omitempty"`
														Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
														Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
													} `tfsdk:"extractors" json:"extractors,omitempty"`
													Headers *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"headers" json:"headers,omitempty"`
													HeadersToAppend *[]struct {
														Key   *string `tfsdk:"key" json:"key,omitempty"`
														Value *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
													HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
													IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
													MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
													ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
													Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
												} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
												XsltTransformation *struct {
													NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
													SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
													Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
												} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
											} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
										} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
										ResponseTransforms *[]struct {
											Matchers *[]struct {
												InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
												Name        *string `tfsdk:"name" json:"name,omitempty"`
												Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
												Value       *string `tfsdk:"value" json:"value,omitempty"`
											} `tfsdk:"matchers" json:"matchers,omitempty"`
											ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
											ResponseTransformation *struct {
												HeaderBodyTransform *struct {
													AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
												} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
												LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
												TransformationTemplate *struct {
													AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
													Body              *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"body" json:"body,omitempty"`
													DynamicMetadataValues *[]struct {
														Key               *string `tfsdk:"key" json:"key,omitempty"`
														MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
														Value             *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
													EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
													Extractors       *struct {
														Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
														Header   *string            `tfsdk:"header" json:"header,omitempty"`
														Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
														Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
													} `tfsdk:"extractors" json:"extractors,omitempty"`
													Headers *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"headers" json:"headers,omitempty"`
													HeadersToAppend *[]struct {
														Key   *string `tfsdk:"key" json:"key,omitempty"`
														Value *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
													HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
													IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
													MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
													ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
													Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
												} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
												XsltTransformation *struct {
													NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
													SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
													Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
												} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
											} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
										} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
									} `tfsdk:"early" json:"early,omitempty"`
									EscapeCharacters       *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									InheritTransformation  *bool `tfsdk:"inherit_transformation" json:"inheritTransformation,omitempty"`
									LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
									Regular                *struct {
										RequestTransforms *[]struct {
											ClearRouteCache *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
											Matcher         *struct {
												CaseSensitive  *bool              `tfsdk:"case_sensitive" json:"caseSensitive,omitempty"`
												ConnectMatcher *map[string]string `tfsdk:"connect_matcher" json:"connectMatcher,omitempty"`
												Exact          *string            `tfsdk:"exact" json:"exact,omitempty"`
												Headers        *[]struct {
													InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
													Name        *string `tfsdk:"name" json:"name,omitempty"`
													Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
													Value       *string `tfsdk:"value" json:"value,omitempty"`
												} `tfsdk:"headers" json:"headers,omitempty"`
												Methods         *[]string `tfsdk:"methods" json:"methods,omitempty"`
												Prefix          *string   `tfsdk:"prefix" json:"prefix,omitempty"`
												QueryParameters *[]struct {
													Name  *string `tfsdk:"name" json:"name,omitempty"`
													Regex *bool   `tfsdk:"regex" json:"regex,omitempty"`
													Value *string `tfsdk:"value" json:"value,omitempty"`
												} `tfsdk:"query_parameters" json:"queryParameters,omitempty"`
												Regex *string `tfsdk:"regex" json:"regex,omitempty"`
											} `tfsdk:"matcher" json:"matcher,omitempty"`
											RequestTransformation *struct {
												HeaderBodyTransform *struct {
													AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
												} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
												LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
												TransformationTemplate *struct {
													AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
													Body              *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"body" json:"body,omitempty"`
													DynamicMetadataValues *[]struct {
														Key               *string `tfsdk:"key" json:"key,omitempty"`
														MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
														Value             *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
													EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
													Extractors       *struct {
														Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
														Header   *string            `tfsdk:"header" json:"header,omitempty"`
														Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
														Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
													} `tfsdk:"extractors" json:"extractors,omitempty"`
													Headers *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"headers" json:"headers,omitempty"`
													HeadersToAppend *[]struct {
														Key   *string `tfsdk:"key" json:"key,omitempty"`
														Value *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
													HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
													IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
													MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
													ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
													Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
												} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
												XsltTransformation *struct {
													NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
													SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
													Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
												} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
											} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
											ResponseTransformation *struct {
												HeaderBodyTransform *struct {
													AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
												} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
												LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
												TransformationTemplate *struct {
													AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
													Body              *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"body" json:"body,omitempty"`
													DynamicMetadataValues *[]struct {
														Key               *string `tfsdk:"key" json:"key,omitempty"`
														MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
														Value             *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
													EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
													Extractors       *struct {
														Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
														Header   *string            `tfsdk:"header" json:"header,omitempty"`
														Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
														Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
													} `tfsdk:"extractors" json:"extractors,omitempty"`
													Headers *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"headers" json:"headers,omitempty"`
													HeadersToAppend *[]struct {
														Key   *string `tfsdk:"key" json:"key,omitempty"`
														Value *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
													HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
													IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
													MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
													ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
													Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
												} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
												XsltTransformation *struct {
													NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
													SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
													Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
												} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
											} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
										} `tfsdk:"request_transforms" json:"requestTransforms,omitempty"`
										ResponseTransforms *[]struct {
											Matchers *[]struct {
												InvertMatch *bool   `tfsdk:"invert_match" json:"invertMatch,omitempty"`
												Name        *string `tfsdk:"name" json:"name,omitempty"`
												Regex       *bool   `tfsdk:"regex" json:"regex,omitempty"`
												Value       *string `tfsdk:"value" json:"value,omitempty"`
											} `tfsdk:"matchers" json:"matchers,omitempty"`
											ResponseCodeDetails    *string `tfsdk:"response_code_details" json:"responseCodeDetails,omitempty"`
											ResponseTransformation *struct {
												HeaderBodyTransform *struct {
													AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
												} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
												LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
												TransformationTemplate *struct {
													AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
													Body              *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"body" json:"body,omitempty"`
													DynamicMetadataValues *[]struct {
														Key               *string `tfsdk:"key" json:"key,omitempty"`
														MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
														Value             *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
													EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
													Extractors       *struct {
														Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
														Header   *string            `tfsdk:"header" json:"header,omitempty"`
														Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
														Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
													} `tfsdk:"extractors" json:"extractors,omitempty"`
													Headers *struct {
														Text *string `tfsdk:"text" json:"text,omitempty"`
													} `tfsdk:"headers" json:"headers,omitempty"`
													HeadersToAppend *[]struct {
														Key   *string `tfsdk:"key" json:"key,omitempty"`
														Value *struct {
															Text *string `tfsdk:"text" json:"text,omitempty"`
														} `tfsdk:"value" json:"value,omitempty"`
													} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
													HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
													IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
													MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
													ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
													Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
												} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
												XsltTransformation *struct {
													NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
													SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
													Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
												} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
											} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
										} `tfsdk:"response_transforms" json:"responseTransforms,omitempty"`
									} `tfsdk:"regular" json:"regular,omitempty"`
								} `tfsdk:"staged_transformations" json:"stagedTransformations,omitempty"`
								Transformations *struct {
									ClearRouteCache       *bool `tfsdk:"clear_route_cache" json:"clearRouteCache,omitempty"`
									RequestTransformation *struct {
										HeaderBodyTransform *struct {
											AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
										} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
										LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
										TransformationTemplate *struct {
											AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
											Body              *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"body" json:"body,omitempty"`
											DynamicMetadataValues *[]struct {
												Key               *string `tfsdk:"key" json:"key,omitempty"`
												MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
												Value             *struct {
													Text *string `tfsdk:"text" json:"text,omitempty"`
												} `tfsdk:"value" json:"value,omitempty"`
											} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
											EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
											Extractors       *struct {
												Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
												Header   *string            `tfsdk:"header" json:"header,omitempty"`
												Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
												Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
											} `tfsdk:"extractors" json:"extractors,omitempty"`
											Headers *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"headers" json:"headers,omitempty"`
											HeadersToAppend *[]struct {
												Key   *string `tfsdk:"key" json:"key,omitempty"`
												Value *struct {
													Text *string `tfsdk:"text" json:"text,omitempty"`
												} `tfsdk:"value" json:"value,omitempty"`
											} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
											HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
											IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
											MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
											ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
											Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
										} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
										XsltTransformation *struct {
											NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
											SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
											Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
										} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
									} `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
									ResponseTransformation *struct {
										HeaderBodyTransform *struct {
											AddRequestMetadata *bool `tfsdk:"add_request_metadata" json:"addRequestMetadata,omitempty"`
										} `tfsdk:"header_body_transform" json:"headerBodyTransform,omitempty"`
										LogRequestResponseInfo *bool `tfsdk:"log_request_response_info" json:"logRequestResponseInfo,omitempty"`
										TransformationTemplate *struct {
											AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
											Body              *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"body" json:"body,omitempty"`
											DynamicMetadataValues *[]struct {
												Key               *string `tfsdk:"key" json:"key,omitempty"`
												MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
												Value             *struct {
													Text *string `tfsdk:"text" json:"text,omitempty"`
												} `tfsdk:"value" json:"value,omitempty"`
											} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
											EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
											Extractors       *struct {
												Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
												Header   *string            `tfsdk:"header" json:"header,omitempty"`
												Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
												Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
											} `tfsdk:"extractors" json:"extractors,omitempty"`
											Headers *struct {
												Text *string `tfsdk:"text" json:"text,omitempty"`
											} `tfsdk:"headers" json:"headers,omitempty"`
											HeadersToAppend *[]struct {
												Key   *string `tfsdk:"key" json:"key,omitempty"`
												Value *struct {
													Text *string `tfsdk:"text" json:"text,omitempty"`
												} `tfsdk:"value" json:"value,omitempty"`
											} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
											HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
											IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
											MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
											ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
											Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
										} `tfsdk:"transformation_template" json:"transformationTemplate,omitempty"`
										XsltTransformation *struct {
											NonXmlTransform *bool   `tfsdk:"non_xml_transform" json:"nonXmlTransform,omitempty"`
											SetContentType  *string `tfsdk:"set_content_type" json:"setContentType,omitempty"`
											Xslt            *string `tfsdk:"xslt" json:"xslt,omitempty"`
										} `tfsdk:"xslt_transformation" json:"xsltTransformation,omitempty"`
									} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
								} `tfsdk:"transformations" json:"transformations,omitempty"`
							} `tfsdk:"options" json:"options,omitempty"`
							Weight *int64 `tfsdk:"weight" json:"weight,omitempty"`
						} `tfsdk:"destinations" json:"destinations,omitempty"`
					} `tfsdk:"multi" json:"multi,omitempty"`
					Single *struct {
						Consul *struct {
							DataCenters *[]string `tfsdk:"data_centers" json:"dataCenters,omitempty"`
							ServiceName *string   `tfsdk:"service_name" json:"serviceName,omitempty"`
							Tags        *[]string `tfsdk:"tags" json:"tags,omitempty"`
						} `tfsdk:"consul" json:"consul,omitempty"`
						DestinationSpec *struct {
							Aws *struct {
								InvocationStyle        *string `tfsdk:"invocation_style" json:"invocationStyle,omitempty"`
								LogicalName            *string `tfsdk:"logical_name" json:"logicalName,omitempty"`
								RequestTransformation  *bool   `tfsdk:"request_transformation" json:"requestTransformation,omitempty"`
								ResponseTransformation *bool   `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
								UnwrapAsAlb            *bool   `tfsdk:"unwrap_as_alb" json:"unwrapAsAlb,omitempty"`
								UnwrapAsApiGateway     *bool   `tfsdk:"unwrap_as_api_gateway" json:"unwrapAsApiGateway,omitempty"`
								WrapAsApiGateway       *bool   `tfsdk:"wrap_as_api_gateway" json:"wrapAsApiGateway,omitempty"`
							} `tfsdk:"aws" json:"aws,omitempty"`
							Azure *struct {
								FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
							} `tfsdk:"azure" json:"azure,omitempty"`
							Grpc *struct {
								Function   *string `tfsdk:"function" json:"function,omitempty"`
								Package    *string `tfsdk:"package" json:"package,omitempty"`
								Parameters *struct {
									Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
									Path    *string            `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"parameters" json:"parameters,omitempty"`
								Service *string `tfsdk:"service" json:"service,omitempty"`
							} `tfsdk:"grpc" json:"grpc,omitempty"`
							Rest *struct {
								FunctionName *string `tfsdk:"function_name" json:"functionName,omitempty"`
								Parameters   *struct {
									Headers *map[string]string `tfsdk:"headers" json:"headers,omitempty"`
									Path    *string            `tfsdk:"path" json:"path,omitempty"`
								} `tfsdk:"parameters" json:"parameters,omitempty"`
								ResponseTransformation *struct {
									AdvancedTemplates *bool `tfsdk:"advanced_templates" json:"advancedTemplates,omitempty"`
									Body              *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"body" json:"body,omitempty"`
									DynamicMetadataValues *[]struct {
										Key               *string `tfsdk:"key" json:"key,omitempty"`
										MetadataNamespace *string `tfsdk:"metadata_namespace" json:"metadataNamespace,omitempty"`
										Value             *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"dynamic_metadata_values" json:"dynamicMetadataValues,omitempty"`
									EscapeCharacters *bool `tfsdk:"escape_characters" json:"escapeCharacters,omitempty"`
									Extractors       *struct {
										Body     *map[string]string `tfsdk:"body" json:"body,omitempty"`
										Header   *string            `tfsdk:"header" json:"header,omitempty"`
										Regex    *string            `tfsdk:"regex" json:"regex,omitempty"`
										Subgroup *int64             `tfsdk:"subgroup" json:"subgroup,omitempty"`
									} `tfsdk:"extractors" json:"extractors,omitempty"`
									Headers *struct {
										Text *string `tfsdk:"text" json:"text,omitempty"`
									} `tfsdk:"headers" json:"headers,omitempty"`
									HeadersToAppend *[]struct {
										Key   *string `tfsdk:"key" json:"key,omitempty"`
										Value *struct {
											Text *string `tfsdk:"text" json:"text,omitempty"`
										} `tfsdk:"value" json:"value,omitempty"`
									} `tfsdk:"headers_to_append" json:"headersToAppend,omitempty"`
									HeadersToRemove       *[]string          `tfsdk:"headers_to_remove" json:"headersToRemove,omitempty"`
									IgnoreErrorOnParse    *bool              `tfsdk:"ignore_error_on_parse" json:"ignoreErrorOnParse,omitempty"`
									MergeExtractorsToBody *map[string]string `tfsdk:"merge_extractors_to_body" json:"mergeExtractorsToBody,omitempty"`
									ParseBodyBehavior     *string            `tfsdk:"parse_body_behavior" json:"parseBodyBehavior,omitempty"`
									Passthrough           *map[string]string `tfsdk:"passthrough" json:"passthrough,omitempty"`
								} `tfsdk:"response_transformation" json:"responseTransformation,omitempty"`
							} `tfsdk:"rest" json:"rest,omitempty"`
						} `tfsdk:"destination_spec" json:"destinationSpec,omitempty"`
						Kube *struct {
							Port *int64 `tfsdk:"port" json:"port,omitempty"`
							Ref  *struct {
								Name      *string `tfsdk:"name" json:"name,omitempty"`
								Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"ref" json:"ref,omitempty"`
						} `tfsdk:"kube" json:"kube,omitempty"`
						Subset *struct {
							Values *map[string]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"subset" json:"subset,omitempty"`
						Upstream *struct {
							Name      *string `tfsdk:"name" json:"name,omitempty"`
							Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						} `tfsdk:"upstream" json:"upstream,omitempty"`
					} `tfsdk:"single" json:"single,omitempty"`
					UpstreamGroup *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"upstream_group" json:"upstreamGroup,omitempty"`
				} `tfsdk:"route_action" json:"routeAction,omitempty"`
			} `tfsdk:"routes" json:"routes,omitempty"`
		} `tfsdk:"virtual_host" json:"virtualHost,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *GatewaySoloIoVirtualServiceV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_gateway_solo_io_virtual_service_v1"
}

func (r *GatewaySoloIoVirtualServiceV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"display_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"namespaced_statuses": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"statuses": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"ssl_config": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"alpn_protocols": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"disable_tls_session_resumption": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ocsp_staple_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"one_way_tls": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"parameters": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cipher_suites": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ecdh_curves": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"maximum_protocol_version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"minimum_protocol_version": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"sds": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"call_credentials": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"file_credential_source": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"header": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"token_file_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

									"certificates_secret_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"cluster_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"target_uri": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"validation_context_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"secret_ref": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"sni_domains": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"ssl_files": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"ocsp_staple": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"root_ca": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_cert": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"transport_socket_connect_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"verify_subject_alt_name": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"virtual_host": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"domains": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"options": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"buffer_per_route": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"buffer": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"max_request_bytes": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"disabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"cors": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"allow_credentials": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"allow_headers": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"allow_methods": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"allow_origin": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"allow_origin_regex": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"disable_for_route": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"expose_headers": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"max_age": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"csrf": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"additional_origins": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"exact": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"ignore_case": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"prefix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"safe_regex": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"google_re2": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"max_program_size": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"regex": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"suffix": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"filter_enabled": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"default_value": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"denominator": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"numerator": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"runtime_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"shadow_enabled": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"default_value": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"denominator": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"numerator": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"runtime_key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

									"dlp": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"actions": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"action_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"custom_action": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"mask_char": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"percent": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"value": schema.Float64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"regex": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"regex_actions": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"regex": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"subgroup": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
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
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"key_value_action": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key_to_mask": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"mask_char": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"percent": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"value": schema.Float64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
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

														"shadow": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"enabled_for": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ext_proc": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"disabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"overrides": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"async_mode": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"grpc_service": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"authority": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"ext_proc_server_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"initial_metadata": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"value": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"retry_policy": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"num_retries": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"retry_back_off": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"base_interval": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"max_interval": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

															"timeout": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"metadata_context_namespaces": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"processing_mode": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"request_body_mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"request_header_mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"request_trailer_mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"response_body_mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"response_header_mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"response_trailer_mode": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"request_attributes": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"response_attributes": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"typed_metadata_context_namespaces": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
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

									"extauth": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"config_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"namespace": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"custom_auth": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"context_extensions": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"disable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"extensions": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"configs": schema.MapAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"header_manipulation": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"request_headers_to_add": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"append": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"header": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"header_secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

											"request_headers_to_remove": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_headers_to_add": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"append": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"header": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

											"response_headers_to_remove": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"include_attempt_count_in_response": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"include_request_attempt_count": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"jwt": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"allow_missing_or_failed_jwt": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"providers": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"audiences": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"claims_to_headers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"append": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"claim": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"header": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"clock_skew_seconds": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"issuer": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"jwks": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"local": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"remote": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"async_fetch": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"fast_listener": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"cache_duration": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"upstream_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"namespace": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																		Description:         "",
																		MarkdownDescription: "",
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

													"keep_token": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"token_source": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"headers": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"header": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"prefix": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"query_params": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"jwt_staged": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"after_ext_auth": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"allow_missing_or_failed_jwt": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"providers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"claims_to_headers": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"append": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"claim": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"header": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"clock_skew_seconds": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"issuer": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"jwks": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"local": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"remote": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"async_fetch": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"fast_listener": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"cache_duration": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"upstream_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
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

															"keep_token": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"token_source": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"headers": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"prefix": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"query_params": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
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
												Required: false,
												Optional: false,
												Computed: true,
											},

											"before_ext_auth": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"allow_missing_or_failed_jwt": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"providers": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"audiences": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"claims_to_headers": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"append": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"claim": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"header": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"clock_skew_seconds": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"issuer": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"jwks": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"local": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"remote": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"async_fetch": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"fast_listener": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"cache_duration": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"upstream_ref": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																				Description:         "",
																				MarkdownDescription: "",
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

															"keep_token": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"token_source": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"headers": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"prefix": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"query_params": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
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
												Required: false,
												Optional: false,
												Computed: true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"rate_limit_configs": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"refs": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"rate_limit_early_configs": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"refs": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"rate_limit_regular_configs": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"refs": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ratelimit": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"local_ratelimit": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"fill_interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_tokens": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"tokens_per_fill": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"rate_limits": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"destination_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"generic_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header_value_match": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"expect_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"headers": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"invert_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"prefix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"present_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"range_match": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"end": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"start": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"regex_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"suffix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"metadata": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"default_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"metadata_key": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"source": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"remote_address": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"request_headers": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"header_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"source_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"set_actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"destination_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"generic_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header_value_match": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"expect_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"headers": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"invert_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"prefix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"present_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"range_match": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"end": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"start": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"regex_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"suffix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"metadata": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"default_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"metadata_key": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"source": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"remote_address": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"request_headers": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"header_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"source_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

									"ratelimit_basic": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"anonymous_limits": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"requests_per_unit": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"unit": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"authorized_limits": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"requests_per_unit": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"unit": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

									"ratelimit_early": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"local_ratelimit": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"fill_interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_tokens": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"tokens_per_fill": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"rate_limits": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"destination_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"generic_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header_value_match": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"expect_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"headers": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"invert_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"prefix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"present_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"range_match": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"end": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"start": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"regex_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"suffix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"metadata": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"default_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"metadata_key": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"source": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"remote_address": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"request_headers": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"header_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"source_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"set_actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"destination_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"generic_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header_value_match": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"expect_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"headers": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"invert_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"prefix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"present_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"range_match": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"end": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"start": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"regex_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"suffix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"metadata": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"default_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"metadata_key": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"source": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"remote_address": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"request_headers": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"header_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"source_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

									"ratelimit_regular": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"local_ratelimit": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"fill_interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_tokens": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"tokens_per_fill": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"rate_limits": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"destination_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"generic_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header_value_match": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"expect_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"headers": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"invert_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"prefix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"present_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"range_match": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"end": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"start": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"regex_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"suffix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"metadata": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"default_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"metadata_key": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"source": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"remote_address": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"request_headers": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"header_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"source_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"set_actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"destination_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"generic_key": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header_value_match": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"expect_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"headers": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"exact_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"invert_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"prefix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"present_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"range_match": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"end": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"start": schema.Int64Attribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"regex_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"suffix_match": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"metadata": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"default_value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"metadata_key": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"path": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"source": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"remote_address": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"request_headers": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"header_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"source_cluster": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

									"rbac": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"disable": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"policies": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"nested_claim_delimiter": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"permissions": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"methods": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path_prefix": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"principals": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"jwt_principal": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"matcher": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"provider": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

									"retries": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"num_retries": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"per_try_timeout": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"retry_back_off": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"base_interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"max_interval": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"retry_on": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"staged_transformations": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"early": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"request_transforms": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"clear_route_cache": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"matcher": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"case_sensitive": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"connect_matcher": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"exact": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"headers": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"invert_match": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"regex": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"methods": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"prefix": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"query_parameters": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"regex": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"regex": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"request_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"header_body_transform": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add_request_metadata": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"log_request_response_info": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"transformation_template": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"advanced_templates": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"body": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"dynamic_metadata_values": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"metadata_namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"escape_characters": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"extractors": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"body": schema.MapAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers_to_append": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"headers_to_remove": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"ignore_error_on_parse": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"merge_extractors_to_body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"parse_body_behavior": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"passthrough": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"xslt_transformation": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"non_xml_transform": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"set_content_type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"xslt": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																"response_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"header_body_transform": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add_request_metadata": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"log_request_response_info": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"transformation_template": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"advanced_templates": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"body": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"dynamic_metadata_values": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"metadata_namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"escape_characters": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"extractors": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"body": schema.MapAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers_to_append": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"headers_to_remove": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"ignore_error_on_parse": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"merge_extractors_to_body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"parse_body_behavior": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"passthrough": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"xslt_transformation": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"non_xml_transform": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"set_content_type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"xslt": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"response_transforms": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"matchers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"invert_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																"response_code_details": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"response_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"header_body_transform": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add_request_metadata": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"log_request_response_info": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"transformation_template": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"advanced_templates": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"body": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"dynamic_metadata_values": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"metadata_namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"escape_characters": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"extractors": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"body": schema.MapAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers_to_append": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"headers_to_remove": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"ignore_error_on_parse": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"merge_extractors_to_body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"parse_body_behavior": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"passthrough": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"xslt_transformation": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"non_xml_transform": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"set_content_type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"xslt": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

											"escape_characters": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"inherit_transformation": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"log_request_response_info": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"regular": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"request_transforms": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"clear_route_cache": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"matcher": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"case_sensitive": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"connect_matcher": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"exact": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"headers": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"invert_match": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"regex": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"methods": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"prefix": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"query_parameters": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"regex": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"regex": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"request_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"header_body_transform": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add_request_metadata": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"log_request_response_info": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"transformation_template": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"advanced_templates": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"body": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"dynamic_metadata_values": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"metadata_namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"escape_characters": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"extractors": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"body": schema.MapAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers_to_append": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"headers_to_remove": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"ignore_error_on_parse": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"merge_extractors_to_body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"parse_body_behavior": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"passthrough": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"xslt_transformation": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"non_xml_transform": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"set_content_type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"xslt": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																"response_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"header_body_transform": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add_request_metadata": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"log_request_response_info": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"transformation_template": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"advanced_templates": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"body": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"dynamic_metadata_values": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"metadata_namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"escape_characters": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"extractors": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"body": schema.MapAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers_to_append": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"headers_to_remove": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"ignore_error_on_parse": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"merge_extractors_to_body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"parse_body_behavior": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"passthrough": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"xslt_transformation": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"non_xml_transform": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"set_content_type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"xslt": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"response_transforms": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"matchers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"invert_match": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"regex": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

																"response_code_details": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"response_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"header_body_transform": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"add_request_metadata": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"log_request_response_info": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"transformation_template": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"advanced_templates": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"body": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"dynamic_metadata_values": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"metadata_namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"escape_characters": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"extractors": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"body": schema.MapAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers_to_append": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"headers_to_remove": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"ignore_error_on_parse": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"merge_extractors_to_body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"parse_body_behavior": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"passthrough": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"xslt_transformation": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"non_xml_transform": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"set_content_type": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"xslt": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"stats": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"virtual_clusters": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"method": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"pattern": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"transformations": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"clear_route_cache": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"request_transformation": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"header_body_transform": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"add_request_metadata": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"log_request_response_info": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"transformation_template": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"advanced_templates": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"body": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"text": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"dynamic_metadata_values": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"metadata_namespace": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"value": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

															"escape_characters": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"extractors": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"body": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"header": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"regex": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"subgroup": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"headers": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"text": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"headers_to_append": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"value": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

															"headers_to_remove": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"ignore_error_on_parse": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"merge_extractors_to_body": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"parse_body_behavior": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"passthrough": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"xslt_transformation": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"non_xml_transform": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"set_content_type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"xslt": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

											"response_transformation": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"header_body_transform": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"add_request_metadata": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"log_request_response_info": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"transformation_template": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"advanced_templates": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"body": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"text": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"dynamic_metadata_values": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"metadata_namespace": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"value": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

															"escape_characters": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"extractors": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"body": schema.MapAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"header": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"regex": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"subgroup": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"headers": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"text": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"headers_to_append": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"value": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

															"headers_to_remove": schema.ListAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"ignore_error_on_parse": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"merge_extractors_to_body": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"parse_body_behavior": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"passthrough": schema.MapAttribute{
																Description:         "",
																MarkdownDescription: "",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"xslt_transformation": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"non_xml_transform": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"set_content_type": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"xslt": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
										Required: false,
										Optional: false,
										Computed: true,
									},

									"waf": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"audit_logging": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"action": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"location": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"config_map_rule_sets": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"config_map_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"data_map_keys": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
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

											"core_rule_set": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"custom_settings_file": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"custom_settings_string": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},
												},
												Required: false,
												Optional: false,
												Computed: true,
											},

											"custom_intervention_message": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"disabled": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"request_headers_only": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"response_headers_only": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"rule_sets": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"directory": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"files": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"rule_str": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

							"options_config_refs": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"delegate_options": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"routes": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"delegate_action": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"selector": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"expressions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"labels": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespaces": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
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

										"direct_response_action": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"body": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"status": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"graphql_api_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"namespace": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"inheritable_matchers": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"inheritable_path_matchers": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"matchers": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"case_sensitive": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"connect_matcher": schema.MapAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"exact": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"headers": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"invert_match": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"regex": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"methods": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"prefix": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"query_parameters": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"regex": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"value": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"regex": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"options": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"append_x_forwarded_host": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"auto_host_rewrite": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"buffer_per_route": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"buffer": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"max_request_bytes": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"disabled": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"cors": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"allow_credentials": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"allow_headers": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"allow_methods": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"allow_origin": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"allow_origin_regex": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"disable_for_route": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"expose_headers": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"max_age": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"csrf": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"additional_origins": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"exact": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"ignore_case": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"prefix": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"safe_regex": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"google_re2": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"max_program_size": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"regex": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"suffix": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"filter_enabled": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"default_value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"denominator": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"numerator": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"runtime_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"shadow_enabled": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"default_value": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"denominator": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"numerator": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"runtime_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"dlp": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"actions": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"action_type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"custom_action": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"mask_char": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"percent": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"value": schema.Float64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"regex": schema.ListAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"regex_actions": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"key_value_action": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key_to_mask": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"mask_char": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"percent": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"value": schema.Float64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																	"shadow": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"enabled_for": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"envoy_metadata": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ext_proc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"disabled": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"overrides": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"async_mode": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"grpc_service": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"authority": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"ext_proc_server_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"namespace": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"initial_metadata": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"retry_policy": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"num_retries": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"retry_back_off": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"base_interval": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"max_interval": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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

																		"timeout": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"metadata_context_namespaces": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"processing_mode": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"request_body_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"request_header_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"request_trailer_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"response_body_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"response_header_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"response_trailer_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"request_attributes": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"response_attributes": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"typed_metadata_context_namespaces": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
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

												"extauth": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"config_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"custom_auth": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"context_extensions": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"disable": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"extensions": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"configs": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"faults": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"abort": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"http_status": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"percentage": schema.Float64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"delay": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"fixed_delay": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"percentage": schema.Float64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"header_manipulation": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"request_headers_to_add": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"append": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"header": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header_secret_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"namespace": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

														"request_headers_to_remove": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"response_headers_to_add": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"append": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"header": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"value": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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

														"response_headers_to_remove": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"host_rewrite": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"host_rewrite_path_regex": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"pattern": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"google_re2": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"max_program_size": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"regex": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"substitution": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"idle_timeout": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"jwt": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"disable": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"jwt_staged": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"after_ext_auth": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"disable": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"before_ext_auth": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"disable": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"lb_hash": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"hash_policies": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"cookie": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"ttl": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"header": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"source_ip": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"terminal": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"max_stream_duration": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"grpc_timeout_header_max": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"grpc_timeout_header_offset": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"max_stream_duration": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"prefix_rewrite": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"rate_limit_configs": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"refs": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"rate_limit_early_configs": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"refs": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"rate_limit_regular_configs": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"refs": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"namespace": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"ratelimit": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"include_vh_rate_limits": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"local_ratelimit": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"fill_interval": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_tokens": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tokens_per_fill": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"rate_limits": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"actions": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"destination_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"generic_key": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"header_value_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"expect_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"exact_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"invert_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"prefix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"present_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"range_match": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"end": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"start": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"regex_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"suffix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"metadata": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"default_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"metadata_key": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"path": schema.ListNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"key": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"source": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"remote_address": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"request_headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"source_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"set_actions": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"destination_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"generic_key": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"header_value_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"expect_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"exact_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"invert_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"prefix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"present_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"range_match": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"end": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"start": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"regex_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"suffix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"metadata": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"default_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"metadata_key": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"path": schema.ListNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"key": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"source": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"remote_address": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"request_headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"source_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

												"ratelimit_basic": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"anonymous_limits": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"requests_per_unit": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"unit": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"authorized_limits": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"requests_per_unit": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"unit": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"ratelimit_early": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"include_vh_rate_limits": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"local_ratelimit": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"fill_interval": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_tokens": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tokens_per_fill": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"rate_limits": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"actions": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"destination_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"generic_key": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"header_value_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"expect_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"exact_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"invert_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"prefix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"present_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"range_match": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"end": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"start": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"regex_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"suffix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"metadata": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"default_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"metadata_key": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"path": schema.ListNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"key": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"source": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"remote_address": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"request_headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"source_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"set_actions": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"destination_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"generic_key": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"header_value_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"expect_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"exact_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"invert_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"prefix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"present_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"range_match": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"end": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"start": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"regex_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"suffix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"metadata": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"default_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"metadata_key": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"path": schema.ListNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"key": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"source": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"remote_address": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"request_headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"source_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

												"ratelimit_regular": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"include_vh_rate_limits": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"local_ratelimit": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"fill_interval": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_tokens": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tokens_per_fill": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"rate_limits": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"actions": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"destination_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"generic_key": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"header_value_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"expect_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"exact_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"invert_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"prefix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"present_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"range_match": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"end": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"start": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"regex_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"suffix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"metadata": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"default_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"metadata_key": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"path": schema.ListNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"key": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"source": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"remote_address": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"request_headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"source_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

																	"set_actions": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"destination_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"generic_key": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"header_value_match": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"expect_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"headers": schema.ListNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"exact_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"invert_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"name": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"prefix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"present_match": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"range_match": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"end": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"start": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"regex_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"suffix_match": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"metadata": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"default_value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"metadata_key": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"path": schema.ListNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"key": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
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
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"source": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"remote_address": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"request_headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"descriptor_key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"source_cluster": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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

												"rbac": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"disable": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"policies": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"nested_claim_delimiter": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"permissions": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"methods": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"path_prefix": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"principals": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"jwt_principal": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"claims": schema.MapAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"matcher": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"provider": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

												"regex_rewrite": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"pattern": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"google_re2": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"max_program_size": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"regex": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"substitution": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"retries": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"num_retries": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"per_try_timeout": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"retry_back_off": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"base_interval": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"max_interval": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"retry_on": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"shadowing": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"percentage": schema.Float64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"upstream": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"staged_transformations": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"early": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"request_transforms": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"clear_route_cache": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"matcher": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"case_sensitive": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"connect_matcher": schema.MapAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"exact": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"headers": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"invert_match": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"name": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"regex": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"value": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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

																					"methods": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"prefix": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"query_parameters": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"regex": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"value": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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

																					"regex": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"request_transformation": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"header_body_transform": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"add_request_metadata": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"log_request_response_info": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"transformation_template": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"advanced_templates": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"body": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"dynamic_metadata_values": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"metadata_namespace": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"escape_characters": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"extractors": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"header": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"regex": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"subgroup": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers_to_append": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"headers_to_remove": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"ignore_error_on_parse": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"merge_extractors_to_body": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"parse_body_behavior": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"passthrough": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"xslt_transformation": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"non_xml_transform": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"set_content_type": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"xslt": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																			"response_transformation": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"header_body_transform": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"add_request_metadata": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"log_request_response_info": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"transformation_template": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"advanced_templates": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"body": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"dynamic_metadata_values": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"metadata_namespace": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"escape_characters": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"extractors": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"header": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"regex": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"subgroup": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers_to_append": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"headers_to_remove": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"ignore_error_on_parse": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"merge_extractors_to_body": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"parse_body_behavior": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"passthrough": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"xslt_transformation": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"non_xml_transform": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"set_content_type": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"xslt": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"response_transforms": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"matchers": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"invert_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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

																			"response_code_details": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"response_transformation": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"header_body_transform": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"add_request_metadata": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"log_request_response_info": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"transformation_template": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"advanced_templates": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"body": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"dynamic_metadata_values": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"metadata_namespace": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"escape_characters": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"extractors": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"header": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"regex": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"subgroup": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers_to_append": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"headers_to_remove": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"ignore_error_on_parse": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"merge_extractors_to_body": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"parse_body_behavior": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"passthrough": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"xslt_transformation": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"non_xml_transform": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"set_content_type": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"xslt": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

														"escape_characters": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"inherit_transformation": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"log_request_response_info": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"regular": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"request_transforms": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"clear_route_cache": schema.BoolAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"matcher": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"case_sensitive": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"connect_matcher": schema.MapAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"exact": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"headers": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"invert_match": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"name": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"regex": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"value": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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

																					"methods": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"prefix": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"query_parameters": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"name": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"regex": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"value": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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

																					"regex": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"request_transformation": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"header_body_transform": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"add_request_metadata": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"log_request_response_info": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"transformation_template": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"advanced_templates": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"body": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"dynamic_metadata_values": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"metadata_namespace": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"escape_characters": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"extractors": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"header": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"regex": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"subgroup": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers_to_append": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"headers_to_remove": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"ignore_error_on_parse": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"merge_extractors_to_body": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"parse_body_behavior": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"passthrough": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"xslt_transformation": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"non_xml_transform": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"set_content_type": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"xslt": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																			"response_transformation": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"header_body_transform": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"add_request_metadata": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"log_request_response_info": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"transformation_template": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"advanced_templates": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"body": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"dynamic_metadata_values": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"metadata_namespace": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"escape_characters": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"extractors": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"header": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"regex": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"subgroup": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers_to_append": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"headers_to_remove": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"ignore_error_on_parse": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"merge_extractors_to_body": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"parse_body_behavior": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"passthrough": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"xslt_transformation": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"non_xml_transform": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"set_content_type": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"xslt": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"response_transforms": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"matchers": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"invert_match": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.BoolAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"value": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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

																			"response_code_details": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"response_transformation": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"header_body_transform": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"add_request_metadata": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"log_request_response_info": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"transformation_template": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"advanced_templates": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"body": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"dynamic_metadata_values": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"metadata_namespace": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"escape_characters": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"extractors": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"header": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"regex": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"subgroup": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"headers_to_append": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"text": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
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

																							"headers_to_remove": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"ignore_error_on_parse": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"merge_extractors_to_body": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"parse_body_behavior": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"passthrough": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"xslt_transformation": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"non_xml_transform": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"set_content_type": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"xslt": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
													Required: false,
													Optional: false,
													Computed: true,
												},

												"timeout": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"tracing": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"propagate": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"route_descriptor": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"trace_percentages": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"client_sample_percentage": schema.Float64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"overall_sample_percentage": schema.Float64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"random_sample_percentage": schema.Float64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"transformations": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"clear_route_cache": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"request_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"headers_to_append": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"passthrough": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"xslt": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

														"response_transformation": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"header_body_transform": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"add_request_metadata": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"log_request_response_info": schema.BoolAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"transformation_template": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"advanced_templates": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"body": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"dynamic_metadata_values": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"metadata_namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																		"escape_characters": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"extractors": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"header": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"regex": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"subgroup": schema.Int64Attribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"headers": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"text": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"headers_to_append": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"value": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"text": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																		"headers_to_remove": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"ignore_error_on_parse": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"merge_extractors_to_body": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"parse_body_behavior": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"passthrough": schema.MapAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"xslt_transformation": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"non_xml_transform": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"set_content_type": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"xslt": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
													Required: false,
													Optional: false,
													Computed: true,
												},

												"upgrades": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"connect": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"enabled": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},
																},
																Required: false,
																Optional: false,
																Computed: true,
															},

															"websocket": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"enabled": schema.BoolAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

												"waf": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"audit_logging": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"action": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"location": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"config_map_rule_sets": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"config_map_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"namespace": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"data_map_keys": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

														"core_rule_set": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"custom_settings_file": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"custom_settings_string": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"custom_intervention_message": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"disabled": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"request_headers_only": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"response_headers_only": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"rule_sets": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"directory": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"files": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"rule_str": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

										"options_config_refs": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"delegate_options": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"namespace": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"redirect_action": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host_redirect": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"https_redirect": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"path_redirect": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"prefix_rewrite": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"regex_rewrite": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"pattern": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"google_re2": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"max_program_size": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"regex": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"substitution": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"response_code": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"strip_query": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"route_action": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"cluster_header": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"dynamic_forward_proxy": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"auto_host_rewrite_header": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"host_rewrite": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"multi": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"destinations": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"destination": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"consul": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"data_centers": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"service_name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"tags": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"destination_spec": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"aws": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"invocation_style": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"logical_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"request_transformation": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"response_transformation": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"unwrap_as_alb": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"unwrap_as_api_gateway": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"wrap_as_api_gateway": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"azure": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"function_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"grpc": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"function": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"package": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"parameters": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"headers": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"path": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"service": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"rest": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"function_name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"parameters": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"headers": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"path": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"response_transformation": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"advanced_templates": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"body": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"text": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"dynamic_metadata_values": schema.ListNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										NestedObject: schema.NestedAttributeObject{
																											Attributes: map[string]schema.Attribute{
																												"key": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"metadata_namespace": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"value": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"text": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																									"escape_characters": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"extractors": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"body": schema.MapAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												ElementType:         types.StringType,
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"header": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"regex": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"subgroup": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"headers": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"text": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"headers_to_append": schema.ListNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										NestedObject: schema.NestedAttributeObject{
																											Attributes: map[string]schema.Attribute{
																												"key": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"value": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"text": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																									"headers_to_remove": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"ignore_error_on_parse": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"merge_extractors_to_body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"parse_body_behavior": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"passthrough": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
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
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"kube": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"port": schema.Int64Attribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																			"subset": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"values": schema.MapAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"upstream": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"namespace": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																	"options": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"buffer_per_route": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"buffer": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"max_request_bytes": schema.Int64Attribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"disabled": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"csrf": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"additional_origins": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"exact": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"ignore_case": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"prefix": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"safe_regex": schema.SingleNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Attributes: map[string]schema.Attribute{
																										"google_re2": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"max_program_size": schema.Int64Attribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},
																											},
																											Required: false,
																											Optional: false,
																											Computed: true,
																										},

																										"regex": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},
																									},
																									Required: false,
																									Optional: false,
																									Computed: true,
																								},

																								"suffix": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
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

																					"filter_enabled": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"default_value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"denominator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"numerator": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"runtime_key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"shadow_enabled": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"default_value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"denominator": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"numerator": schema.Int64Attribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"runtime_key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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

																			"extauth": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"config_ref": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"custom_auth": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"context_extensions": schema.MapAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								ElementType:         types.StringType,
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"name": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},
																						},
																						Required: false,
																						Optional: false,
																						Computed: true,
																					},

																					"disable": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"extensions": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"configs": schema.MapAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"header_manipulation": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"request_headers_to_add": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"append": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"header": schema.SingleNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},
																									},
																									Required: false,
																									Optional: false,
																									Computed: true,
																								},

																								"header_secret_ref": schema.SingleNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Attributes: map[string]schema.Attribute{
																										"name": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"namespace": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
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

																					"request_headers_to_remove": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"response_headers_to_add": schema.ListNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"append": schema.BoolAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"header": schema.SingleNestedAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Attributes: map[string]schema.Attribute{
																										"key": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"value": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
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

																					"response_headers_to_remove": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},
																				},
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"staged_transformations": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"early": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"request_transforms": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"clear_route_cache": schema.BoolAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"matcher": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"case_sensitive": schema.BoolAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"connect_matcher": schema.MapAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													ElementType:         types.StringType,
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"exact": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"headers": schema.ListNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													NestedObject: schema.NestedAttributeObject{
																														Attributes: map[string]schema.Attribute{
																															"invert_match": schema.BoolAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"name": schema.StringAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"regex": schema.BoolAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"value": schema.StringAttribute{
																																Description:         "",
																																MarkdownDescription: "",
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

																												"methods": schema.ListAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													ElementType:         types.StringType,
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"prefix": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"query_parameters": schema.ListNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													NestedObject: schema.NestedAttributeObject{
																														Attributes: map[string]schema.Attribute{
																															"name": schema.StringAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"regex": schema.BoolAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"value": schema.StringAttribute{
																																Description:         "",
																																MarkdownDescription: "",
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

																												"regex": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},
																											},
																											Required: false,
																											Optional: false,
																											Computed: true,
																										},

																										"request_transformation": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"header_body_transform": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"add_request_metadata": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"log_request_response_info": schema.BoolAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"transformation_template": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"advanced_templates": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"body": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"dynamic_metadata_values": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"metadata_namespace": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"escape_characters": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"extractors": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"body": schema.MapAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	ElementType:         types.StringType,
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"header": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"regex": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"subgroup": schema.Int64Attribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers_to_append": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"headers_to_remove": schema.ListAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"ignore_error_on_parse": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"merge_extractors_to_body": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"parse_body_behavior": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"passthrough": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"xslt_transformation": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"non_xml_transform": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"set_content_type": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"xslt": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																										"response_transformation": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"header_body_transform": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"add_request_metadata": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"log_request_response_info": schema.BoolAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"transformation_template": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"advanced_templates": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"body": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"dynamic_metadata_values": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"metadata_namespace": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"escape_characters": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"extractors": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"body": schema.MapAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	ElementType:         types.StringType,
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"header": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"regex": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"subgroup": schema.Int64Attribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers_to_append": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"headers_to_remove": schema.ListAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"ignore_error_on_parse": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"merge_extractors_to_body": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"parse_body_behavior": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"passthrough": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"xslt_transformation": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"non_xml_transform": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"set_content_type": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"xslt": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"response_transforms": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"matchers": schema.ListNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											NestedObject: schema.NestedAttributeObject{
																												Attributes: map[string]schema.Attribute{
																													"invert_match": schema.BoolAttribute{
																														Description:         "",
																														MarkdownDescription: "",
																														Required:            false,
																														Optional:            false,
																														Computed:            true,
																													},

																													"name": schema.StringAttribute{
																														Description:         "",
																														MarkdownDescription: "",
																														Required:            false,
																														Optional:            false,
																														Computed:            true,
																													},

																													"regex": schema.BoolAttribute{
																														Description:         "",
																														MarkdownDescription: "",
																														Required:            false,
																														Optional:            false,
																														Computed:            true,
																													},

																													"value": schema.StringAttribute{
																														Description:         "",
																														MarkdownDescription: "",
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

																										"response_code_details": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"response_transformation": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"header_body_transform": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"add_request_metadata": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"log_request_response_info": schema.BoolAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"transformation_template": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"advanced_templates": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"body": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"dynamic_metadata_values": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"metadata_namespace": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"escape_characters": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"extractors": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"body": schema.MapAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	ElementType:         types.StringType,
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"header": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"regex": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"subgroup": schema.Int64Attribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers_to_append": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"headers_to_remove": schema.ListAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"ignore_error_on_parse": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"merge_extractors_to_body": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"parse_body_behavior": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"passthrough": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"xslt_transformation": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"non_xml_transform": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"set_content_type": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"xslt": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																					"escape_characters": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"inherit_transformation": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"log_request_response_info": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"regular": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"request_transforms": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"clear_route_cache": schema.BoolAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"matcher": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"case_sensitive": schema.BoolAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"connect_matcher": schema.MapAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													ElementType:         types.StringType,
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"exact": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"headers": schema.ListNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													NestedObject: schema.NestedAttributeObject{
																														Attributes: map[string]schema.Attribute{
																															"invert_match": schema.BoolAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"name": schema.StringAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"regex": schema.BoolAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"value": schema.StringAttribute{
																																Description:         "",
																																MarkdownDescription: "",
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

																												"methods": schema.ListAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													ElementType:         types.StringType,
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"prefix": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"query_parameters": schema.ListNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													NestedObject: schema.NestedAttributeObject{
																														Attributes: map[string]schema.Attribute{
																															"name": schema.StringAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"regex": schema.BoolAttribute{
																																Description:         "",
																																MarkdownDescription: "",
																																Required:            false,
																																Optional:            false,
																																Computed:            true,
																															},

																															"value": schema.StringAttribute{
																																Description:         "",
																																MarkdownDescription: "",
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

																												"regex": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},
																											},
																											Required: false,
																											Optional: false,
																											Computed: true,
																										},

																										"request_transformation": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"header_body_transform": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"add_request_metadata": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"log_request_response_info": schema.BoolAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"transformation_template": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"advanced_templates": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"body": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"dynamic_metadata_values": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"metadata_namespace": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"escape_characters": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"extractors": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"body": schema.MapAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	ElementType:         types.StringType,
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"header": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"regex": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"subgroup": schema.Int64Attribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers_to_append": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"headers_to_remove": schema.ListAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"ignore_error_on_parse": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"merge_extractors_to_body": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"parse_body_behavior": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"passthrough": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"xslt_transformation": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"non_xml_transform": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"set_content_type": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"xslt": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																										"response_transformation": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"header_body_transform": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"add_request_metadata": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"log_request_response_info": schema.BoolAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"transformation_template": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"advanced_templates": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"body": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"dynamic_metadata_values": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"metadata_namespace": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"escape_characters": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"extractors": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"body": schema.MapAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	ElementType:         types.StringType,
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"header": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"regex": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"subgroup": schema.Int64Attribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers_to_append": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"headers_to_remove": schema.ListAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"ignore_error_on_parse": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"merge_extractors_to_body": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"parse_body_behavior": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"passthrough": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"xslt_transformation": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"non_xml_transform": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"set_content_type": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"xslt": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"response_transforms": schema.ListNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								NestedObject: schema.NestedAttributeObject{
																									Attributes: map[string]schema.Attribute{
																										"matchers": schema.ListNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											NestedObject: schema.NestedAttributeObject{
																												Attributes: map[string]schema.Attribute{
																													"invert_match": schema.BoolAttribute{
																														Description:         "",
																														MarkdownDescription: "",
																														Required:            false,
																														Optional:            false,
																														Computed:            true,
																													},

																													"name": schema.StringAttribute{
																														Description:         "",
																														MarkdownDescription: "",
																														Required:            false,
																														Optional:            false,
																														Computed:            true,
																													},

																													"regex": schema.BoolAttribute{
																														Description:         "",
																														MarkdownDescription: "",
																														Required:            false,
																														Optional:            false,
																														Computed:            true,
																													},

																													"value": schema.StringAttribute{
																														Description:         "",
																														MarkdownDescription: "",
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

																										"response_code_details": schema.StringAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Required:            false,
																											Optional:            false,
																											Computed:            true,
																										},

																										"response_transformation": schema.SingleNestedAttribute{
																											Description:         "",
																											MarkdownDescription: "",
																											Attributes: map[string]schema.Attribute{
																												"header_body_transform": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"add_request_metadata": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"log_request_response_info": schema.BoolAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"transformation_template": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"advanced_templates": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"body": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"dynamic_metadata_values": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"metadata_namespace": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"escape_characters": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"extractors": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"body": schema.MapAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	ElementType:         types.StringType,
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"header": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"regex": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},

																																"subgroup": schema.Int64Attribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers": schema.SingleNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Attributes: map[string]schema.Attribute{
																																"text": schema.StringAttribute{
																																	Description:         "",
																																	MarkdownDescription: "",
																																	Required:            false,
																																	Optional:            false,
																																	Computed:            true,
																																},
																															},
																															Required: false,
																															Optional: false,
																															Computed: true,
																														},

																														"headers_to_append": schema.ListNestedAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															NestedObject: schema.NestedAttributeObject{
																																Attributes: map[string]schema.Attribute{
																																	"key": schema.StringAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Required:            false,
																																		Optional:            false,
																																		Computed:            true,
																																	},

																																	"value": schema.SingleNestedAttribute{
																																		Description:         "",
																																		MarkdownDescription: "",
																																		Attributes: map[string]schema.Attribute{
																																			"text": schema.StringAttribute{
																																				Description:         "",
																																				MarkdownDescription: "",
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

																														"headers_to_remove": schema.ListAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"ignore_error_on_parse": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"merge_extractors_to_body": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"parse_body_behavior": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"passthrough": schema.MapAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															ElementType:         types.StringType,
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},
																													},
																													Required: false,
																													Optional: false,
																													Computed: true,
																												},

																												"xslt_transformation": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"non_xml_transform": schema.BoolAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"set_content_type": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
																															Required:            false,
																															Optional:            false,
																															Computed:            true,
																														},

																														"xslt": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},

																			"transformations": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"clear_route_cache": schema.BoolAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"request_transformation": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"header_body_transform": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"add_request_metadata": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"log_request_response_info": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"transformation_template": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"advanced_templates": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"body": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"text": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"dynamic_metadata_values": schema.ListNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										NestedObject: schema.NestedAttributeObject{
																											Attributes: map[string]schema.Attribute{
																												"key": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"metadata_namespace": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"value": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"text": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																									"escape_characters": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"extractors": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"body": schema.MapAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												ElementType:         types.StringType,
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"header": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"regex": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"subgroup": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"headers": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"text": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"headers_to_append": schema.ListNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										NestedObject: schema.NestedAttributeObject{
																											Attributes: map[string]schema.Attribute{
																												"key": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"value": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"text": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																									"headers_to_remove": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"ignore_error_on_parse": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"merge_extractors_to_body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"parse_body_behavior": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"passthrough": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"xslt_transformation": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"non_xml_transform": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"set_content_type": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"xslt": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																					"response_transformation": schema.SingleNestedAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Attributes: map[string]schema.Attribute{
																							"header_body_transform": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"add_request_metadata": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"log_request_response_info": schema.BoolAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"transformation_template": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"advanced_templates": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"body": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"text": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"dynamic_metadata_values": schema.ListNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										NestedObject: schema.NestedAttributeObject{
																											Attributes: map[string]schema.Attribute{
																												"key": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"metadata_namespace": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"value": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"text": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																									"escape_characters": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"extractors": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"body": schema.MapAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												ElementType:         types.StringType,
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"header": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"regex": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},

																											"subgroup": schema.Int64Attribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"headers": schema.SingleNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Attributes: map[string]schema.Attribute{
																											"text": schema.StringAttribute{
																												Description:         "",
																												MarkdownDescription: "",
																												Required:            false,
																												Optional:            false,
																												Computed:            true,
																											},
																										},
																										Required: false,
																										Optional: false,
																										Computed: true,
																									},

																									"headers_to_append": schema.ListNestedAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										NestedObject: schema.NestedAttributeObject{
																											Attributes: map[string]schema.Attribute{
																												"key": schema.StringAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Required:            false,
																													Optional:            false,
																													Computed:            true,
																												},

																												"value": schema.SingleNestedAttribute{
																													Description:         "",
																													MarkdownDescription: "",
																													Attributes: map[string]schema.Attribute{
																														"text": schema.StringAttribute{
																															Description:         "",
																															MarkdownDescription: "",
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

																									"headers_to_remove": schema.ListAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"ignore_error_on_parse": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"merge_extractors_to_body": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"parse_body_behavior": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"passthrough": schema.MapAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										ElementType:         types.StringType,
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},
																								},
																								Required: false,
																								Optional: false,
																								Computed: true,
																							},

																							"xslt_transformation": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"non_xml_transform": schema.BoolAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"set_content_type": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
																										Required:            false,
																										Optional:            false,
																										Computed:            true,
																									},

																									"xslt": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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
																				Required: false,
																				Optional: false,
																				Computed: true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"weight": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
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
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"single": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"consul": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"data_centers": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"service_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"tags": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"destination_spec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"aws": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"invocation_style": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"logical_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"request_transformation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"response_transformation": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"unwrap_as_alb": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"unwrap_as_api_gateway": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"wrap_as_api_gateway": schema.BoolAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"azure": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"function_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"function": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"package": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"parameters": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"headers": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"service": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"rest": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"function_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"parameters": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"headers": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"path": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"response_transformation": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"advanced_templates": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"body": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"dynamic_metadata_values": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"metadata_namespace": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"escape_characters": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"extractors": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"body": schema.MapAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"header": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"subgroup": schema.Int64Attribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"text": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},
																					},
																					Required: false,
																					Optional: false,
																					Computed: true,
																				},

																				"headers_to_append": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"value": schema.SingleNestedAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Attributes: map[string]schema.Attribute{
																									"text": schema.StringAttribute{
																										Description:         "",
																										MarkdownDescription: "",
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

																				"headers_to_remove": schema.ListAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"ignore_error_on_parse": schema.BoolAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"merge_extractors_to_body": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"parse_body_behavior": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"passthrough": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
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
															Required: false,
															Optional: false,
															Computed: true,
														},

														"kube": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"namespace": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

														"subset": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"values": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"upstream": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"namespace": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

												"upstream_group": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"namespace": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *GatewaySoloIoVirtualServiceV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *GatewaySoloIoVirtualServiceV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_gateway_solo_io_virtual_service_v1")

	var data GatewaySoloIoVirtualServiceV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "gateway.solo.io", Version: "v1", Resource: "VirtualService"}).
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

	var readResponse GatewaySoloIoVirtualServiceV1DataSourceData
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
	data.ApiVersion = pointer.String("gateway.solo.io/v1")
	data.Kind = pointer.String("VirtualService")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
