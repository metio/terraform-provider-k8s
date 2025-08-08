/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package edp_epam_com_v1alpha1

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
	_ datasource.DataSource = &EdpEpamComNexusRepositoryV1Alpha1Manifest{}
)

func NewEdpEpamComNexusRepositoryV1Alpha1Manifest() datasource.DataSource {
	return &EdpEpamComNexusRepositoryV1Alpha1Manifest{}
}

type EdpEpamComNexusRepositoryV1Alpha1Manifest struct{}

type EdpEpamComNexusRepositoryV1Alpha1ManifestData struct {
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
		Apt *struct {
			Hosted *struct {
				Apt *struct {
					Distribution *string `tfsdk:"distribution" json:"distribution,omitempty"`
				} `tfsdk:"apt" json:"apt,omitempty"`
				AptSigning *struct {
					Keypair    *string `tfsdk:"keypair" json:"keypair,omitempty"`
					Passphrase *string `tfsdk:"passphrase" json:"passphrase,omitempty"`
				} `tfsdk:"apt_signing" json:"aptSigning,omitempty"`
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Apt *struct {
					Distribution *string `tfsdk:"distribution" json:"distribution,omitempty"`
					Flat         *bool   `tfsdk:"flat" json:"flat,omitempty"`
				} `tfsdk:"apt" json:"apt,omitempty"`
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"apt" json:"apt,omitempty"`
		Bower *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Bower *struct {
					RewritePackageUrls *bool `tfsdk:"rewrite_package_urls" json:"rewritePackageUrls,omitempty"`
				} `tfsdk:"bower" json:"bower,omitempty"`
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"bower" json:"bower,omitempty"`
		Cocoapods *struct {
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"cocoapods" json:"cocoapods,omitempty"`
		Conan *struct {
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"conan" json:"conan,omitempty"`
		Conda *struct {
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"conda" json:"conda,omitempty"`
		Docker *struct {
			Group *struct {
				Docker *struct {
					ForceBasicAuth *bool  `tfsdk:"force_basic_auth" json:"forceBasicAuth,omitempty"`
					HttpPort       *int64 `tfsdk:"http_port" json:"httpPort,omitempty"`
					HttpsPort      *int64 `tfsdk:"https_port" json:"httpsPort,omitempty"`
					V1Enabled      *bool  `tfsdk:"v1_enabled" json:"v1Enabled,omitempty"`
				} `tfsdk:"docker" json:"docker,omitempty"`
				Group *struct {
					MemberNames    *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
					WritableMember *string   `tfsdk:"writable_member" json:"writableMember,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Docker *struct {
					ForceBasicAuth *bool  `tfsdk:"force_basic_auth" json:"forceBasicAuth,omitempty"`
					HttpPort       *int64 `tfsdk:"http_port" json:"httpPort,omitempty"`
					HttpsPort      *int64 `tfsdk:"https_port" json:"httpsPort,omitempty"`
					V1Enabled      *bool  `tfsdk:"v1_enabled" json:"v1Enabled,omitempty"`
				} `tfsdk:"docker" json:"docker,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Docker *struct {
					ForceBasicAuth *bool  `tfsdk:"force_basic_auth" json:"forceBasicAuth,omitempty"`
					HttpPort       *int64 `tfsdk:"http_port" json:"httpPort,omitempty"`
					HttpsPort      *int64 `tfsdk:"https_port" json:"httpsPort,omitempty"`
					V1Enabled      *bool  `tfsdk:"v1_enabled" json:"v1Enabled,omitempty"`
				} `tfsdk:"docker" json:"docker,omitempty"`
				DockerProxy *struct {
					IndexType *string `tfsdk:"index_type" json:"indexType,omitempty"`
					IndexUrl  *string `tfsdk:"index_url" json:"indexUrl,omitempty"`
				} `tfsdk:"docker_proxy" json:"dockerProxy,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"docker" json:"docker,omitempty"`
		GitLfs *struct {
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
		} `tfsdk:"git_lfs" json:"gitLfs,omitempty"`
		Go *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"go" json:"go,omitempty"`
		Helm *struct {
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"helm" json:"helm,omitempty"`
		Maven *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Maven *struct {
					ContentDisposition *string `tfsdk:"content_disposition" json:"contentDisposition,omitempty"`
					LayoutPolicy       *string `tfsdk:"layout_policy" json:"layoutPolicy,omitempty"`
					VersionPolicy      *string `tfsdk:"version_policy" json:"versionPolicy,omitempty"`
				} `tfsdk:"maven" json:"maven,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Maven *struct {
					ContentDisposition *string `tfsdk:"content_disposition" json:"contentDisposition,omitempty"`
					LayoutPolicy       *string `tfsdk:"layout_policy" json:"layoutPolicy,omitempty"`
					VersionPolicy      *string `tfsdk:"version_policy" json:"versionPolicy,omitempty"`
				} `tfsdk:"maven" json:"maven,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Preemptive *bool   `tfsdk:"preemptive" json:"preemptive,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Maven *struct {
					ContentDisposition *string `tfsdk:"content_disposition" json:"contentDisposition,omitempty"`
					LayoutPolicy       *string `tfsdk:"layout_policy" json:"layoutPolicy,omitempty"`
					VersionPolicy      *string `tfsdk:"version_policy" json:"versionPolicy,omitempty"`
				} `tfsdk:"maven" json:"maven,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"maven" json:"maven,omitempty"`
		NexusRef *struct {
			Kind *string `tfsdk:"kind" json:"kind,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"nexus_ref" json:"nexusRef,omitempty"`
		Npm *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Npm *struct {
					RemoveNonCataloged *bool `tfsdk:"remove_non_cataloged" json:"removeNonCataloged,omitempty"`
					RemoveQuarantined  *bool `tfsdk:"remove_quarantined" json:"removeQuarantined,omitempty"`
				} `tfsdk:"npm" json:"npm,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"npm" json:"npm,omitempty"`
		Nuget *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				NugetProxy *struct {
					NugetVersion         *string `tfsdk:"nuget_version" json:"nugetVersion,omitempty"`
					QueryCacheItemMaxAge *int64  `tfsdk:"query_cache_item_max_age" json:"queryCacheItemMaxAge,omitempty"`
				} `tfsdk:"nuget_proxy" json:"nugetProxy,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"nuget" json:"nuget,omitempty"`
		P2 *struct {
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"p2" json:"p2,omitempty"`
		Pypi *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"pypi" json:"pypi,omitempty"`
		R *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"r" json:"r,omitempty"`
		Raw *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Online *bool   `tfsdk:"online" json:"online,omitempty"`
				Raw    *struct {
					ContentDisposition *string `tfsdk:"content_disposition" json:"contentDisposition,omitempty"`
				} `tfsdk:"raw" json:"raw,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Online *bool   `tfsdk:"online" json:"online,omitempty"`
				Raw    *struct {
					ContentDisposition *string `tfsdk:"content_disposition" json:"contentDisposition,omitempty"`
				} `tfsdk:"raw" json:"raw,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				Raw *struct {
					ContentDisposition *string `tfsdk:"content_disposition" json:"contentDisposition,omitempty"`
				} `tfsdk:"raw" json:"raw,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"raw" json:"raw,omitempty"`
		RubyGems *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"ruby_gems" json:"rubyGems,omitempty"`
		Yum *struct {
			Group *struct {
				Group *struct {
					MemberNames *[]string `tfsdk:"member_names" json:"memberNames,omitempty"`
				} `tfsdk:"group" json:"group,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
				YumSigning *struct {
					Keypair    *string `tfsdk:"keypair" json:"keypair,omitempty"`
					Passphrase *string `tfsdk:"passphrase" json:"passphrase,omitempty"`
				} `tfsdk:"yum_signing" json:"yumSigning,omitempty"`
			} `tfsdk:"group" json:"group,omitempty"`
			Hosted *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				Component *struct {
					ProprietaryComponents *bool `tfsdk:"proprietary_components" json:"proprietaryComponents,omitempty"`
				} `tfsdk:"component" json:"component,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Online  *bool   `tfsdk:"online" json:"online,omitempty"`
				Storage *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
					WritePolicy                 *string `tfsdk:"write_policy" json:"writePolicy,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
				Yum *struct {
					DeployPolicy  *string `tfsdk:"deploy_policy" json:"deployPolicy,omitempty"`
					RepodataDepth *int64  `tfsdk:"repodata_depth" json:"repodataDepth,omitempty"`
				} `tfsdk:"yum" json:"yum,omitempty"`
			} `tfsdk:"hosted" json:"hosted,omitempty"`
			Proxy *struct {
				Cleanup *struct {
					PolicyNames *[]string `tfsdk:"policy_names" json:"policyNames,omitempty"`
				} `tfsdk:"cleanup" json:"cleanup,omitempty"`
				HttpClient *struct {
					Authentication *struct {
						NtlmDomain *string `tfsdk:"ntlm_domain" json:"ntlmDomain,omitempty"`
						NtlmHost   *string `tfsdk:"ntlm_host" json:"ntlmHost,omitempty"`
						Password   *string `tfsdk:"password" json:"password,omitempty"`
						Type       *string `tfsdk:"type" json:"type,omitempty"`
						Username   *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"authentication" json:"authentication,omitempty"`
					AutoBlock  *bool `tfsdk:"auto_block" json:"autoBlock,omitempty"`
					Blocked    *bool `tfsdk:"blocked" json:"blocked,omitempty"`
					Connection *struct {
						EnableCircularRedirects *bool   `tfsdk:"enable_circular_redirects" json:"enableCircularRedirects,omitempty"`
						EnableCookies           *bool   `tfsdk:"enable_cookies" json:"enableCookies,omitempty"`
						Retries                 *int64  `tfsdk:"retries" json:"retries,omitempty"`
						Timeout                 *int64  `tfsdk:"timeout" json:"timeout,omitempty"`
						UseTrustStore           *bool   `tfsdk:"use_trust_store" json:"useTrustStore,omitempty"`
						UserAgentSuffix         *string `tfsdk:"user_agent_suffix" json:"userAgentSuffix,omitempty"`
					} `tfsdk:"connection" json:"connection,omitempty"`
				} `tfsdk:"http_client" json:"httpClient,omitempty"`
				Name          *string `tfsdk:"name" json:"name,omitempty"`
				NegativeCache *struct {
					Enabled    *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					TimeToLive *int64 `tfsdk:"time_to_live" json:"timeToLive,omitempty"`
				} `tfsdk:"negative_cache" json:"negativeCache,omitempty"`
				Online *bool `tfsdk:"online" json:"online,omitempty"`
				Proxy  *struct {
					ContentMaxAge  *int64  `tfsdk:"content_max_age" json:"contentMaxAge,omitempty"`
					MetadataMaxAge *int64  `tfsdk:"metadata_max_age" json:"metadataMaxAge,omitempty"`
					RemoteUrl      *string `tfsdk:"remote_url" json:"remoteUrl,omitempty"`
				} `tfsdk:"proxy" json:"proxy,omitempty"`
				RoutingRule *string `tfsdk:"routing_rule" json:"routingRule,omitempty"`
				Storage     *struct {
					BlobStoreName               *string `tfsdk:"blob_store_name" json:"blobStoreName,omitempty"`
					StrictContentTypeValidation *bool   `tfsdk:"strict_content_type_validation" json:"strictContentTypeValidation,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
				YumSigning *struct {
					Keypair    *string `tfsdk:"keypair" json:"keypair,omitempty"`
					Passphrase *string `tfsdk:"passphrase" json:"passphrase,omitempty"`
				} `tfsdk:"yum_signing" json:"yumSigning,omitempty"`
			} `tfsdk:"proxy" json:"proxy,omitempty"`
		} `tfsdk:"yum" json:"yum,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *EdpEpamComNexusRepositoryV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_edp_epam_com_nexus_repository_v1alpha1_manifest"
}

func (r *EdpEpamComNexusRepositoryV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "NexusRepository is the Schema for the nexusrepositories API.",
		MarkdownDescription: "NexusRepository is the Schema for the nexusrepositories API.",
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
				Description:         "NexusRepositorySpec defines the desired state of NexusRepository. It should contain only one format of repository - go, maven, npm, etc. and only one type - proxy, hosted or group.",
				MarkdownDescription: "NexusRepositorySpec defines the desired state of NexusRepository. It should contain only one format of repository - go, maven, npm, etc. and only one type - proxy, hosted or group.",
				Attributes: map[string]schema.Attribute{
					"apt": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"apt": schema.SingleNestedAttribute{
										Description:         "Apt contains data of hosted repositories of format Apt.",
										MarkdownDescription: "Apt contains data of hosted repositories of format Apt.",
										Attributes: map[string]schema.Attribute{
											"distribution": schema.StringAttribute{
												Description:         "Distribution to fetch",
												MarkdownDescription: "Distribution to fetch",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"apt_signing": schema.SingleNestedAttribute{
										Description:         "AptSigning contains signing data of hosted repositores of format Apt.",
										MarkdownDescription: "AptSigning contains signing data of hosted repositores of format Apt.",
										Attributes: map[string]schema.Attribute{
											"keypair": schema.StringAttribute{
												Description:         "PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)",
												MarkdownDescription: "PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"passphrase": schema.StringAttribute{
												Description:         "Passphrase to access PGP signing key",
												MarkdownDescription: "Passphrase to access PGP signing key",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"apt": schema.SingleNestedAttribute{
										Description:         "Apt configuration.",
										MarkdownDescription: "Apt configuration.",
										Attributes: map[string]schema.Attribute{
											"distribution": schema.StringAttribute{
												Description:         "Distribution to fetch.",
												MarkdownDescription: "Distribution to fetch.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"flat": schema.BoolAttribute{
												Description:         "Whether this repository is flat.",
												MarkdownDescription: "Whether this repository is flat.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"bower": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"bower": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"rewrite_package_urls": schema.BoolAttribute{
												Description:         "Whether to force Bower to retrieve packages through this proxy repository",
												MarkdownDescription: "Whether to force Bower to retrieve packages through this proxy repository",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"cocoapods": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"conan": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"conda": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"docker": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"docker": schema.SingleNestedAttribute{
										Description:         "Docker contains data of a Docker Repositoriy.",
										MarkdownDescription: "Docker contains data of a Docker Repositoriy.",
										Attributes: map[string]schema.Attribute{
											"force_basic_auth": schema.BoolAttribute{
												Description:         "Whether to force authentication (Docker Bearer Token Realm required if false)",
												MarkdownDescription: "Whether to force authentication (Docker Bearer Token Realm required if false)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"http_port": schema.Int64Attribute{
												Description:         "Create an HTTP connector at specified port",
												MarkdownDescription: "Create an HTTP connector at specified port",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"https_port": schema.Int64Attribute{
												Description:         "Create an HTTPS connector at specified port",
												MarkdownDescription: "Create an HTTPS connector at specified port",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"v1_enabled": schema.BoolAttribute{
												Description:         "Whether to allow clients to use the V1 API to interact with this repository",
												MarkdownDescription: "Whether to allow clients to use the V1 API to interact with this repository",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"writable_member": schema.StringAttribute{
												Description:         "Pro-only: This field is for the Group Deployment feature available in NXRM Pro.",
												MarkdownDescription: "Pro-only: This field is for the Group Deployment feature available in NXRM Pro.",
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
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"docker": schema.SingleNestedAttribute{
										Description:         "Docker contains data of a Docker Repositoriy.",
										MarkdownDescription: "Docker contains data of a Docker Repositoriy.",
										Attributes: map[string]schema.Attribute{
											"force_basic_auth": schema.BoolAttribute{
												Description:         "Whether to force authentication (Docker Bearer Token Realm required if false)",
												MarkdownDescription: "Whether to force authentication (Docker Bearer Token Realm required if false)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"http_port": schema.Int64Attribute{
												Description:         "Create an HTTP connector at specified port",
												MarkdownDescription: "Create an HTTP connector at specified port",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"https_port": schema.Int64Attribute{
												Description:         "Create an HTTPS connector at specified port",
												MarkdownDescription: "Create an HTTPS connector at specified port",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"v1_enabled": schema.BoolAttribute{
												Description:         "Whether to allow clients to use the V1 API to interact with this repository",
												MarkdownDescription: "Whether to allow clients to use the V1 API to interact with this repository",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"docker": schema.SingleNestedAttribute{
										Description:         "Docker contains data of a Docker Repositoriy.",
										MarkdownDescription: "Docker contains data of a Docker Repositoriy.",
										Attributes: map[string]schema.Attribute{
											"force_basic_auth": schema.BoolAttribute{
												Description:         "Whether to force authentication (Docker Bearer Token Realm required if false)",
												MarkdownDescription: "Whether to force authentication (Docker Bearer Token Realm required if false)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"http_port": schema.Int64Attribute{
												Description:         "Create an HTTP connector at specified port",
												MarkdownDescription: "Create an HTTP connector at specified port",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"https_port": schema.Int64Attribute{
												Description:         "Create an HTTPS connector at specified port",
												MarkdownDescription: "Create an HTTPS connector at specified port",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"v1_enabled": schema.BoolAttribute{
												Description:         "Whether to allow clients to use the V1 API to interact with this repository",
												MarkdownDescription: "Whether to allow clients to use the V1 API to interact with this repository",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"docker_proxy": schema.SingleNestedAttribute{
										Description:         "DockerProxy contains data of a Docker Proxy Repository.",
										MarkdownDescription: "DockerProxy contains data of a Docker Proxy Repository.",
										Attributes: map[string]schema.Attribute{
											"index_type": schema.StringAttribute{
												Description:         "Type of Docker Index.",
												MarkdownDescription: "Type of Docker Index.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("HUB", "REGISTRY", "CUSTOM"),
												},
											},

											"index_url": schema.StringAttribute{
												Description:         "Url of Docker Index to use.",
												MarkdownDescription: "Url of Docker Index to use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"git_lfs": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

					"go": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"helm": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"maven": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"maven": schema.SingleNestedAttribute{
										Description:         "Maven contains additional data of maven repository.",
										MarkdownDescription: "Maven contains additional data of maven repository.",
										Attributes: map[string]schema.Attribute{
											"content_disposition": schema.StringAttribute{
												Description:         "Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browser.",
												MarkdownDescription: "Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browser.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("INLINE", "ATTACHMENT"),
												},
											},

											"layout_policy": schema.StringAttribute{
												Description:         "Validate that all paths are maven artifact or metadata paths.",
												MarkdownDescription: "Validate that all paths are maven artifact or metadata paths.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("STRICT", "PERMISSIVE"),
												},
											},

											"version_policy": schema.StringAttribute{
												Description:         "VersionPolicy is a type of artifact that this repository stores.",
												MarkdownDescription: "VersionPolicy is a type of artifact that this repository stores.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RELEASE", "SNAPSHOT", "MIXED"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"maven": schema.SingleNestedAttribute{
										Description:         "Maven contains additional data of maven repository.",
										MarkdownDescription: "Maven contains additional data of maven repository.",
										Attributes: map[string]schema.Attribute{
											"content_disposition": schema.StringAttribute{
												Description:         "Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browser.",
												MarkdownDescription: "Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browser.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("INLINE", "ATTACHMENT"),
												},
											},

											"layout_policy": schema.StringAttribute{
												Description:         "Validate that all paths are maven artifact or metadata paths.",
												MarkdownDescription: "Validate that all paths are maven artifact or metadata paths.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("STRICT", "PERMISSIVE"),
												},
											},

											"version_policy": schema.StringAttribute{
												Description:         "VersionPolicy is a type of artifact that this repository stores.",
												MarkdownDescription: "VersionPolicy is a type of artifact that this repository stores.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RELEASE", "SNAPSHOT", "MIXED"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthenticationWithPreemptive contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthenticationWithPreemptive contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"preemptive": schema.BoolAttribute{
														Description:         "Whether to use pre-emptive authentication. Use with caution. Defaults to false.",
														MarkdownDescription: "Whether to use pre-emptive authentication. Use with caution. Defaults to false.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
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

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Whether to block outbound connections on the repository.",
												MarkdownDescription: "Whether to block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"maven": schema.SingleNestedAttribute{
										Description:         "Maven contains additional data of maven repository.",
										MarkdownDescription: "Maven contains additional data of maven repository.",
										Attributes: map[string]schema.Attribute{
											"content_disposition": schema.StringAttribute{
												Description:         "Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browser.",
												MarkdownDescription: "Add Content-Disposition header as 'Attachment' to disable some content from being inline in a browser.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("INLINE", "ATTACHMENT"),
												},
											},

											"layout_policy": schema.StringAttribute{
												Description:         "Validate that all paths are maven artifact or metadata paths.",
												MarkdownDescription: "Validate that all paths are maven artifact or metadata paths.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("STRICT", "PERMISSIVE"),
												},
											},

											"version_policy": schema.StringAttribute{
												Description:         "VersionPolicy is a type of artifact that this repository stores.",
												MarkdownDescription: "VersionPolicy is a type of artifact that this repository stores.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RELEASE", "SNAPSHOT", "MIXED"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"nexus_ref": schema.SingleNestedAttribute{
						Description:         "NexusRef is a reference to Nexus custom resource.",
						MarkdownDescription: "NexusRef is a reference to Nexus custom resource.",
						Attributes: map[string]schema.Attribute{
							"kind": schema.StringAttribute{
								Description:         "Kind specifies the kind of the Nexus resource.",
								MarkdownDescription: "Kind specifies the kind of the Nexus resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name specifies the name of the Nexus resource.",
								MarkdownDescription: "Name specifies the name of the Nexus resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"npm": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"npm": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"remove_non_cataloged": schema.BoolAttribute{
												Description:         "Remove Non-Cataloged Versions",
												MarkdownDescription: "Remove Non-Cataloged Versions",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"remove_quarantined": schema.BoolAttribute{
												Description:         "Remove Quarantined Versions",
												MarkdownDescription: "Remove Quarantined Versions",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"nuget": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"nuget_proxy": schema.SingleNestedAttribute{
										Description:         "NugetProxy contains data specific to proxy repositories of format Nuget.",
										MarkdownDescription: "NugetProxy contains data specific to proxy repositories of format Nuget.",
										Attributes: map[string]schema.Attribute{
											"nuget_version": schema.StringAttribute{
												Description:         "NugetVersion is the used Nuget protocol version.",
												MarkdownDescription: "NugetVersion is the used Nuget protocol version.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("V2", "V3"),
												},
											},

											"query_cache_item_max_age": schema.Int64Attribute{
												Description:         "How long to cache query results from the proxied repository (in seconds)",
												MarkdownDescription: "How long to cache query results from the proxied repository (in seconds)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"p2": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"pypi": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"r": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"raw": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"raw": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"content_disposition": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("INLINE", "ATTACHMENT"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"raw": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"content_disposition": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("INLINE", "ATTACHMENT"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"raw": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"content_disposition": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("INLINE", "ATTACHMENT"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"ruby_gems": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
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

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
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

					"yum": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"group": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"group": schema.SingleNestedAttribute{
										Description:         "Group configuration.",
										MarkdownDescription: "Group configuration.",
										Attributes: map[string]schema.Attribute{
											"member_names": schema.ListAttribute{
												Description:         "Member repositories' names.",
												MarkdownDescription: "Member repositories' names.",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"yum_signing": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"keypair": schema.StringAttribute{
												Description:         "PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)",
												MarkdownDescription: "PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"passphrase": schema.StringAttribute{
												Description:         "Passphrase to access PGP signing key",
												MarkdownDescription: "Passphrase to access PGP signing key",
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

							"hosted": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"component": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"proprietary_components": schema.BoolAttribute{
												Description:         "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												MarkdownDescription: "Components in this repository count as proprietary for namespace conflict attacks (requires Sonatype Nexus Firewall)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"write_policy": schema.StringAttribute{
												Description:         "WritePolicy controls if deployments of and updates to assets are allowed.",
												MarkdownDescription: "WritePolicy controls if deployments of and updates to assets are allowed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ALLOW", "ALLOW_ONCE", "DENY", "REPLICATION_ONLY"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"yum": schema.SingleNestedAttribute{
										Description:         "Yum contains data of hosted repositories of format Yum.",
										MarkdownDescription: "Yum contains data of hosted repositories of format Yum.",
										Attributes: map[string]schema.Attribute{
											"deploy_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("PERMISSIVE", "STRICT"),
												},
											},

											"repodata_depth": schema.Int64Attribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"proxy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cleanup": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"policy_names": schema.ListAttribute{
												Description:         " Components that match any of the applied policies will be deleted.",
												MarkdownDescription: " Components that match any of the applied policies will be deleted.",
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

									"http_client": schema.SingleNestedAttribute{
										Description:         "HTTP client configuration.",
										MarkdownDescription: "HTTP client configuration.",
										Attributes: map[string]schema.Attribute{
											"authentication": schema.SingleNestedAttribute{
												Description:         "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												MarkdownDescription: "HTTPClientAuthentication contains HTTP client authentication configuration data.",
												Attributes: map[string]schema.Attribute{
													"ntlm_domain": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"ntlm_host": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password for authentication.",
														MarkdownDescription: "Password for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "Type of authentication to use.",
														MarkdownDescription: "Type of authentication to use.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("username", "ntlm"),
														},
													},

													"username": schema.StringAttribute{
														Description:         "Username for authentication.",
														MarkdownDescription: "Username for authentication.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"auto_block": schema.BoolAttribute{
												Description:         "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												MarkdownDescription: "Auto-block outbound connections on the repository if remote peer is detected as unreachable/unresponsive",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"blocked": schema.BoolAttribute{
												Description:         "Block outbound connections on the repository.",
												MarkdownDescription: "Block outbound connections on the repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection": schema.SingleNestedAttribute{
												Description:         "HTTPClientConnection contains HTTP client connection configuration data.",
												MarkdownDescription: "HTTPClientConnection contains HTTP client connection configuration data.",
												Attributes: map[string]schema.Attribute{
													"enable_circular_redirects": schema.BoolAttribute{
														Description:         "Whether to enable redirects to the same location (required by some servers)",
														MarkdownDescription: "Whether to enable redirects to the same location (required by some servers)",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enable_cookies": schema.BoolAttribute{
														Description:         "Whether to allow cookies to be stored and used",
														MarkdownDescription: "Whether to allow cookies to be stored and used",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retries": schema.Int64Attribute{
														Description:         "Total retries if the initial connection attempt suffers a timeout",
														MarkdownDescription: "Total retries if the initial connection attempt suffers a timeout",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"timeout": schema.Int64Attribute{
														Description:         "Seconds to wait for activity before stopping and retrying the connection',",
														MarkdownDescription: "Seconds to wait for activity before stopping and retrying the connection',",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_trust_store": schema.BoolAttribute{
														Description:         "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														MarkdownDescription: "Use certificates stored in the Nexus Repository Manager truststore to connect to external systems",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"user_agent_suffix": schema.StringAttribute{
														Description:         "Custom fragment to append to User-Agent header in HTTP requests",
														MarkdownDescription: "Custom fragment to append to User-Agent header in HTTP requests",
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

									"name": schema.StringAttribute{
										Description:         "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										MarkdownDescription: "A unique identifier for this repository. Only letters, digits, underscores(_), hyphens(-), and dots(.) are allowed and may not start with underscore or dot.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9_.-]*$`), ""),
										},
									},

									"negative_cache": schema.SingleNestedAttribute{
										Description:         "Negative cache configuration.",
										MarkdownDescription: "Negative cache configuration.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Whether to cache responses for content not present in the proxied repository.",
												MarkdownDescription: "Whether to cache responses for content not present in the proxied repository.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_to_live": schema.Int64Attribute{
												Description:         "How long to cache the fact that a file was not found in the repository (in minutes).",
												MarkdownDescription: "How long to cache the fact that a file was not found in the repository (in minutes).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"online": schema.BoolAttribute{
										Description:         "Online determines if the repository accepts incoming requests.",
										MarkdownDescription: "Online determines if the repository accepts incoming requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"proxy": schema.SingleNestedAttribute{
										Description:         "Proxy configuration.",
										MarkdownDescription: "Proxy configuration.",
										Attributes: map[string]schema.Attribute{
											"content_max_age": schema.Int64Attribute{
												Description:         "How long to cache artifacts before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache artifacts before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata_max_age": schema.Int64Attribute{
												Description:         "How long to cache metadata before rechecking the remote repository (in minutes)",
												MarkdownDescription: "How long to cache metadata before rechecking the remote repository (in minutes)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"remote_url": schema.StringAttribute{
												Description:         "Location of the remote repository being proxied.",
												MarkdownDescription: "Location of the remote repository being proxied.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"routing_rule": schema.StringAttribute{
										Description:         "The name of the routing rule assigned to this repository.",
										MarkdownDescription: "The name of the routing rule assigned to this repository.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.SingleNestedAttribute{
										Description:         "Storage configuration.",
										MarkdownDescription: "Storage configuration.",
										Attributes: map[string]schema.Attribute{
											"blob_store_name": schema.StringAttribute{
												Description:         "Blob store used to store repository contents.",
												MarkdownDescription: "Blob store used to store repository contents.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strict_content_type_validation": schema.BoolAttribute{
												Description:         "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												MarkdownDescription: "StrictContentTypeValidation: Whether to validate uploaded content's MIME type appropriate for the repository format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"yum_signing": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"keypair": schema.StringAttribute{
												Description:         "PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)",
												MarkdownDescription: "PGP signing key pair (armored private key e.g. gpg --export-secret-key --armor)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"passphrase": schema.StringAttribute{
												Description:         "Passphrase to access PGP signing key",
												MarkdownDescription: "Passphrase to access PGP signing key",
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *EdpEpamComNexusRepositoryV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_edp_epam_com_nexus_repository_v1alpha1_manifest")

	var model EdpEpamComNexusRepositoryV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("edp.epam.com/v1alpha1")
	model.Kind = pointer.String("NexusRepository")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
