/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kuadrant_io_v1beta2

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
	_ datasource.DataSource = &KuadrantIoAuthPolicyV1Beta2Manifest{}
)

func NewKuadrantIoAuthPolicyV1Beta2Manifest() datasource.DataSource {
	return &KuadrantIoAuthPolicyV1Beta2Manifest{}
}

type KuadrantIoAuthPolicyV1Beta2Manifest struct{}

type KuadrantIoAuthPolicyV1Beta2ManifestData struct {
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
		Patterns       *map[string]string `tfsdk:"patterns" json:"patterns,omitempty"`
		RouteSelectors *[]struct {
			Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
			Matches   *[]struct {
				Headers *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"headers" json:"headers,omitempty"`
				Method *string `tfsdk:"method" json:"method,omitempty"`
				Path   *struct {
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"path" json:"path,omitempty"`
				QueryParams *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Type  *string `tfsdk:"type" json:"type,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"query_params" json:"queryParams,omitempty"`
			} `tfsdk:"matches" json:"matches,omitempty"`
		} `tfsdk:"route_selectors" json:"routeSelectors,omitempty"`
		Rules *struct {
			Authentication *struct {
				Anonymous *map[string]string `tfsdk:"anonymous" json:"anonymous,omitempty"`
				ApiKey    *struct {
					AllNamespaces *bool `tfsdk:"all_namespaces" json:"allNamespaces,omitempty"`
					Selector      *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"api_key" json:"apiKey,omitempty"`
				Cache *struct {
					Key *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"key" json:"key,omitempty"`
					Ttl *int64 `tfsdk:"ttl" json:"ttl,omitempty"`
				} `tfsdk:"cache" json:"cache,omitempty"`
				Credentials *struct {
					AuthorizationHeader *struct {
						Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					} `tfsdk:"authorization_header" json:"authorizationHeader,omitempty"`
					Cookie *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"cookie" json:"cookie,omitempty"`
					CustomHeader *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"custom_header" json:"customHeader,omitempty"`
					QueryString *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"query_string" json:"queryString,omitempty"`
				} `tfsdk:"credentials" json:"credentials,omitempty"`
				Defaults *struct {
					Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"defaults" json:"defaults,omitempty"`
				Jwt *struct {
					IssuerUrl *string `tfsdk:"issuer_url" json:"issuerUrl,omitempty"`
					Ttl       *int64  `tfsdk:"ttl" json:"ttl,omitempty"`
				} `tfsdk:"jwt" json:"jwt,omitempty"`
				KubernetesTokenReview *struct {
					Audiences *[]string `tfsdk:"audiences" json:"audiences,omitempty"`
				} `tfsdk:"kubernetes_token_review" json:"kubernetesTokenReview,omitempty"`
				Metrics             *bool `tfsdk:"metrics" json:"metrics,omitempty"`
				Oauth2Introspection *struct {
					CredentialsRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"credentials_ref" json:"credentialsRef,omitempty"`
					Endpoint      *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
					TokenTypeHint *string `tfsdk:"token_type_hint" json:"tokenTypeHint,omitempty"`
				} `tfsdk:"oauth2_introspection" json:"oauth2Introspection,omitempty"`
				Overrides *struct {
					Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
					Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Plain *struct {
					Selector *string `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"plain" json:"plain,omitempty"`
				Priority       *int64 `tfsdk:"priority" json:"priority,omitempty"`
				RouteSelectors *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
					Matches   *[]struct {
						Headers *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						Method *string `tfsdk:"method" json:"method,omitempty"`
						Path   *struct {
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"path" json:"path,omitempty"`
						QueryParams *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"query_params" json:"queryParams,omitempty"`
					} `tfsdk:"matches" json:"matches,omitempty"`
				} `tfsdk:"route_selectors" json:"routeSelectors,omitempty"`
				When *[]struct {
					All        *[]map[string]string `tfsdk:"all" json:"all,omitempty"`
					Any        *[]map[string]string `tfsdk:"any" json:"any,omitempty"`
					Operator   *string              `tfsdk:"operator" json:"operator,omitempty"`
					PatternRef *string              `tfsdk:"pattern_ref" json:"patternRef,omitempty"`
					Selector   *string              `tfsdk:"selector" json:"selector,omitempty"`
					Value      *string              `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"when" json:"when,omitempty"`
				X509 *struct {
					AllNamespaces *bool `tfsdk:"all_namespaces" json:"allNamespaces,omitempty"`
					Selector      *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
				} `tfsdk:"x509" json:"x509,omitempty"`
			} `tfsdk:"authentication" json:"authentication,omitempty"`
			Authorization *struct {
				Cache *struct {
					Key *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"key" json:"key,omitempty"`
					Ttl *int64 `tfsdk:"ttl" json:"ttl,omitempty"`
				} `tfsdk:"cache" json:"cache,omitempty"`
				KubernetesSubjectAccessReview *struct {
					Groups             *[]string `tfsdk:"groups" json:"groups,omitempty"`
					ResourceAttributes *struct {
						Group *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"group" json:"group,omitempty"`
						Name *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"name" json:"name,omitempty"`
						Namespace *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"namespace" json:"namespace,omitempty"`
						Resource *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"resource" json:"resource,omitempty"`
						Subresource *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"subresource" json:"subresource,omitempty"`
						Verb *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"verb" json:"verb,omitempty"`
					} `tfsdk:"resource_attributes" json:"resourceAttributes,omitempty"`
					User *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"kubernetes_subject_access_review" json:"kubernetesSubjectAccessReview,omitempty"`
				Metrics *bool `tfsdk:"metrics" json:"metrics,omitempty"`
				Opa     *struct {
					AllValues      *bool `tfsdk:"all_values" json:"allValues,omitempty"`
					ExternalPolicy *struct {
						Body *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"body" json:"body,omitempty"`
						BodyParameters *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"body_parameters" json:"bodyParameters,omitempty"`
						ContentType *string `tfsdk:"content_type" json:"contentType,omitempty"`
						Credentials *struct {
							AuthorizationHeader *struct {
								Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
							} `tfsdk:"authorization_header" json:"authorizationHeader,omitempty"`
							Cookie *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"cookie" json:"cookie,omitempty"`
							CustomHeader *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"custom_header" json:"customHeader,omitempty"`
							QueryString *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"query_string" json:"queryString,omitempty"`
						} `tfsdk:"credentials" json:"credentials,omitempty"`
						Headers *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						Method *string `tfsdk:"method" json:"method,omitempty"`
						Oauth2 *struct {
							Cache           *bool   `tfsdk:"cache" json:"cache,omitempty"`
							ClientId        *string `tfsdk:"client_id" json:"clientId,omitempty"`
							ClientSecretRef *struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"client_secret_ref" json:"clientSecretRef,omitempty"`
							ExtraParams *map[string]string `tfsdk:"extra_params" json:"extraParams,omitempty"`
							Scopes      *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
							TokenUrl    *string            `tfsdk:"token_url" json:"tokenUrl,omitempty"`
						} `tfsdk:"oauth2" json:"oauth2,omitempty"`
						SharedSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"shared_secret_ref" json:"sharedSecretRef,omitempty"`
						Ttl *int64  `tfsdk:"ttl" json:"ttl,omitempty"`
						Url *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"external_policy" json:"externalPolicy,omitempty"`
					Rego *string `tfsdk:"rego" json:"rego,omitempty"`
				} `tfsdk:"opa" json:"opa,omitempty"`
				PatternMatching *struct {
					Patterns *[]struct {
						All        *[]map[string]string `tfsdk:"all" json:"all,omitempty"`
						Any        *[]map[string]string `tfsdk:"any" json:"any,omitempty"`
						Operator   *string              `tfsdk:"operator" json:"operator,omitempty"`
						PatternRef *string              `tfsdk:"pattern_ref" json:"patternRef,omitempty"`
						Selector   *string              `tfsdk:"selector" json:"selector,omitempty"`
						Value      *string              `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"patterns" json:"patterns,omitempty"`
				} `tfsdk:"pattern_matching" json:"patternMatching,omitempty"`
				Priority       *int64 `tfsdk:"priority" json:"priority,omitempty"`
				RouteSelectors *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
					Matches   *[]struct {
						Headers *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						Method *string `tfsdk:"method" json:"method,omitempty"`
						Path   *struct {
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"path" json:"path,omitempty"`
						QueryParams *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"query_params" json:"queryParams,omitempty"`
					} `tfsdk:"matches" json:"matches,omitempty"`
				} `tfsdk:"route_selectors" json:"routeSelectors,omitempty"`
				Spicedb *struct {
					Endpoint   *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
					Insecure   *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					Permission *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"permission" json:"permission,omitempty"`
					Resource *struct {
						Kind *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"kind" json:"kind,omitempty"`
						Name *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"resource" json:"resource,omitempty"`
					SharedSecretRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"shared_secret_ref" json:"sharedSecretRef,omitempty"`
					Subject *struct {
						Kind *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"kind" json:"kind,omitempty"`
						Name *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"subject" json:"subject,omitempty"`
				} `tfsdk:"spicedb" json:"spicedb,omitempty"`
				When *[]struct {
					All        *[]map[string]string `tfsdk:"all" json:"all,omitempty"`
					Any        *[]map[string]string `tfsdk:"any" json:"any,omitempty"`
					Operator   *string              `tfsdk:"operator" json:"operator,omitempty"`
					PatternRef *string              `tfsdk:"pattern_ref" json:"patternRef,omitempty"`
					Selector   *string              `tfsdk:"selector" json:"selector,omitempty"`
					Value      *string              `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"when" json:"when,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			Callbacks *struct {
				Cache *struct {
					Key *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"key" json:"key,omitempty"`
					Ttl *int64 `tfsdk:"ttl" json:"ttl,omitempty"`
				} `tfsdk:"cache" json:"cache,omitempty"`
				Http *struct {
					Body *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"body" json:"body,omitempty"`
					BodyParameters *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"body_parameters" json:"bodyParameters,omitempty"`
					ContentType *string `tfsdk:"content_type" json:"contentType,omitempty"`
					Credentials *struct {
						AuthorizationHeader *struct {
							Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
						} `tfsdk:"authorization_header" json:"authorizationHeader,omitempty"`
						Cookie *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"cookie" json:"cookie,omitempty"`
						CustomHeader *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"custom_header" json:"customHeader,omitempty"`
						QueryString *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"query_string" json:"queryString,omitempty"`
					} `tfsdk:"credentials" json:"credentials,omitempty"`
					Headers *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Method *string `tfsdk:"method" json:"method,omitempty"`
					Oauth2 *struct {
						Cache           *bool   `tfsdk:"cache" json:"cache,omitempty"`
						ClientId        *string `tfsdk:"client_id" json:"clientId,omitempty"`
						ClientSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"client_secret_ref" json:"clientSecretRef,omitempty"`
						ExtraParams *map[string]string `tfsdk:"extra_params" json:"extraParams,omitempty"`
						Scopes      *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TokenUrl    *string            `tfsdk:"token_url" json:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" json:"oauth2,omitempty"`
					SharedSecretRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"shared_secret_ref" json:"sharedSecretRef,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Metrics        *bool  `tfsdk:"metrics" json:"metrics,omitempty"`
				Priority       *int64 `tfsdk:"priority" json:"priority,omitempty"`
				RouteSelectors *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
					Matches   *[]struct {
						Headers *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						Method *string `tfsdk:"method" json:"method,omitempty"`
						Path   *struct {
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"path" json:"path,omitempty"`
						QueryParams *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"query_params" json:"queryParams,omitempty"`
					} `tfsdk:"matches" json:"matches,omitempty"`
				} `tfsdk:"route_selectors" json:"routeSelectors,omitempty"`
				When *[]struct {
					All        *[]map[string]string `tfsdk:"all" json:"all,omitempty"`
					Any        *[]map[string]string `tfsdk:"any" json:"any,omitempty"`
					Operator   *string              `tfsdk:"operator" json:"operator,omitempty"`
					PatternRef *string              `tfsdk:"pattern_ref" json:"patternRef,omitempty"`
					Selector   *string              `tfsdk:"selector" json:"selector,omitempty"`
					Value      *string              `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"when" json:"when,omitempty"`
			} `tfsdk:"callbacks" json:"callbacks,omitempty"`
			Metadata *struct {
				Cache *struct {
					Key *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"key" json:"key,omitempty"`
					Ttl *int64 `tfsdk:"ttl" json:"ttl,omitempty"`
				} `tfsdk:"cache" json:"cache,omitempty"`
				Http *struct {
					Body *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"body" json:"body,omitempty"`
					BodyParameters *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"body_parameters" json:"bodyParameters,omitempty"`
					ContentType *string `tfsdk:"content_type" json:"contentType,omitempty"`
					Credentials *struct {
						AuthorizationHeader *struct {
							Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
						} `tfsdk:"authorization_header" json:"authorizationHeader,omitempty"`
						Cookie *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"cookie" json:"cookie,omitempty"`
						CustomHeader *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"custom_header" json:"customHeader,omitempty"`
						QueryString *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"query_string" json:"queryString,omitempty"`
					} `tfsdk:"credentials" json:"credentials,omitempty"`
					Headers *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Method *string `tfsdk:"method" json:"method,omitempty"`
					Oauth2 *struct {
						Cache           *bool   `tfsdk:"cache" json:"cache,omitempty"`
						ClientId        *string `tfsdk:"client_id" json:"clientId,omitempty"`
						ClientSecretRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"client_secret_ref" json:"clientSecretRef,omitempty"`
						ExtraParams *map[string]string `tfsdk:"extra_params" json:"extraParams,omitempty"`
						Scopes      *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
						TokenUrl    *string            `tfsdk:"token_url" json:"tokenUrl,omitempty"`
					} `tfsdk:"oauth2" json:"oauth2,omitempty"`
					SharedSecretRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"shared_secret_ref" json:"sharedSecretRef,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Metrics        *bool  `tfsdk:"metrics" json:"metrics,omitempty"`
				Priority       *int64 `tfsdk:"priority" json:"priority,omitempty"`
				RouteSelectors *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
					Matches   *[]struct {
						Headers *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"headers" json:"headers,omitempty"`
						Method *string `tfsdk:"method" json:"method,omitempty"`
						Path   *struct {
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"path" json:"path,omitempty"`
						QueryParams *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Type  *string `tfsdk:"type" json:"type,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"query_params" json:"queryParams,omitempty"`
					} `tfsdk:"matches" json:"matches,omitempty"`
				} `tfsdk:"route_selectors" json:"routeSelectors,omitempty"`
				Uma *struct {
					CredentialsRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"credentials_ref" json:"credentialsRef,omitempty"`
					Endpoint *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				} `tfsdk:"uma" json:"uma,omitempty"`
				UserInfo *struct {
					IdentitySource *string `tfsdk:"identity_source" json:"identitySource,omitempty"`
				} `tfsdk:"user_info" json:"userInfo,omitempty"`
				When *[]struct {
					All        *[]map[string]string `tfsdk:"all" json:"all,omitempty"`
					Any        *[]map[string]string `tfsdk:"any" json:"any,omitempty"`
					Operator   *string              `tfsdk:"operator" json:"operator,omitempty"`
					PatternRef *string              `tfsdk:"pattern_ref" json:"patternRef,omitempty"`
					Selector   *string              `tfsdk:"selector" json:"selector,omitempty"`
					Value      *string              `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"when" json:"when,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			Response *struct {
				Success *struct {
					DynamicMetadata *struct {
						Cache *struct {
							Key *struct {
								Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
								Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"key" json:"key,omitempty"`
							Ttl *int64 `tfsdk:"ttl" json:"ttl,omitempty"`
						} `tfsdk:"cache" json:"cache,omitempty"`
						Json *struct {
							Properties *struct {
								Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
								Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"properties" json:"properties,omitempty"`
						} `tfsdk:"json" json:"json,omitempty"`
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Metrics *bool   `tfsdk:"metrics" json:"metrics,omitempty"`
						Plain   *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"plain" json:"plain,omitempty"`
						Priority       *int64 `tfsdk:"priority" json:"priority,omitempty"`
						RouteSelectors *[]struct {
							Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
							Matches   *[]struct {
								Headers *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Type  *string `tfsdk:"type" json:"type,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								Method *string `tfsdk:"method" json:"method,omitempty"`
								Path   *struct {
									Type  *string `tfsdk:"type" json:"type,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
								QueryParams *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Type  *string `tfsdk:"type" json:"type,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"query_params" json:"queryParams,omitempty"`
							} `tfsdk:"matches" json:"matches,omitempty"`
						} `tfsdk:"route_selectors" json:"routeSelectors,omitempty"`
						When *[]struct {
							All        *[]map[string]string `tfsdk:"all" json:"all,omitempty"`
							Any        *[]map[string]string `tfsdk:"any" json:"any,omitempty"`
							Operator   *string              `tfsdk:"operator" json:"operator,omitempty"`
							PatternRef *string              `tfsdk:"pattern_ref" json:"patternRef,omitempty"`
							Selector   *string              `tfsdk:"selector" json:"selector,omitempty"`
							Value      *string              `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"when" json:"when,omitempty"`
						Wristband *struct {
							CustomClaims *struct {
								Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
								Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"custom_claims" json:"customClaims,omitempty"`
							Issuer         *string `tfsdk:"issuer" json:"issuer,omitempty"`
							SigningKeyRefs *[]struct {
								Algorithm *string `tfsdk:"algorithm" json:"algorithm,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"signing_key_refs" json:"signingKeyRefs,omitempty"`
							TokenDuration *int64 `tfsdk:"token_duration" json:"tokenDuration,omitempty"`
						} `tfsdk:"wristband" json:"wristband,omitempty"`
					} `tfsdk:"dynamic_metadata" json:"dynamicMetadata,omitempty"`
					Headers *struct {
						Cache *struct {
							Key *struct {
								Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
								Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"key" json:"key,omitempty"`
							Ttl *int64 `tfsdk:"ttl" json:"ttl,omitempty"`
						} `tfsdk:"cache" json:"cache,omitempty"`
						Json *struct {
							Properties *struct {
								Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
								Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"properties" json:"properties,omitempty"`
						} `tfsdk:"json" json:"json,omitempty"`
						Key     *string `tfsdk:"key" json:"key,omitempty"`
						Metrics *bool   `tfsdk:"metrics" json:"metrics,omitempty"`
						Plain   *struct {
							Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
							Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"plain" json:"plain,omitempty"`
						Priority       *int64 `tfsdk:"priority" json:"priority,omitempty"`
						RouteSelectors *[]struct {
							Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
							Matches   *[]struct {
								Headers *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Type  *string `tfsdk:"type" json:"type,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								Method *string `tfsdk:"method" json:"method,omitempty"`
								Path   *struct {
									Type  *string `tfsdk:"type" json:"type,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"path" json:"path,omitempty"`
								QueryParams *[]struct {
									Name  *string `tfsdk:"name" json:"name,omitempty"`
									Type  *string `tfsdk:"type" json:"type,omitempty"`
									Value *string `tfsdk:"value" json:"value,omitempty"`
								} `tfsdk:"query_params" json:"queryParams,omitempty"`
							} `tfsdk:"matches" json:"matches,omitempty"`
						} `tfsdk:"route_selectors" json:"routeSelectors,omitempty"`
						When *[]struct {
							All        *[]map[string]string `tfsdk:"all" json:"all,omitempty"`
							Any        *[]map[string]string `tfsdk:"any" json:"any,omitempty"`
							Operator   *string              `tfsdk:"operator" json:"operator,omitempty"`
							PatternRef *string              `tfsdk:"pattern_ref" json:"patternRef,omitempty"`
							Selector   *string              `tfsdk:"selector" json:"selector,omitempty"`
							Value      *string              `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"when" json:"when,omitempty"`
						Wristband *struct {
							CustomClaims *struct {
								Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
								Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"custom_claims" json:"customClaims,omitempty"`
							Issuer         *string `tfsdk:"issuer" json:"issuer,omitempty"`
							SigningKeyRefs *[]struct {
								Algorithm *string `tfsdk:"algorithm" json:"algorithm,omitempty"`
								Name      *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"signing_key_refs" json:"signingKeyRefs,omitempty"`
							TokenDuration *int64 `tfsdk:"token_duration" json:"tokenDuration,omitempty"`
						} `tfsdk:"wristband" json:"wristband,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
				} `tfsdk:"success" json:"success,omitempty"`
				Unauthenticated *struct {
					Body *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"body" json:"body,omitempty"`
					Code    *int64 `tfsdk:"code" json:"code,omitempty"`
					Headers *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Message *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"message" json:"message,omitempty"`
				} `tfsdk:"unauthenticated" json:"unauthenticated,omitempty"`
				Unauthorized *struct {
					Body *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"body" json:"body,omitempty"`
					Code    *int64 `tfsdk:"code" json:"code,omitempty"`
					Headers *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"headers" json:"headers,omitempty"`
					Message *struct {
						Selector *string            `tfsdk:"selector" json:"selector,omitempty"`
						Value    *map[string]string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"message" json:"message,omitempty"`
				} `tfsdk:"unauthorized" json:"unauthorized,omitempty"`
			} `tfsdk:"response" json:"response,omitempty"`
		} `tfsdk:"rules" json:"rules,omitempty"`
		TargetRef *struct {
			Group     *string `tfsdk:"group" json:"group,omitempty"`
			Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"target_ref" json:"targetRef,omitempty"`
		When *[]struct {
			All        *[]map[string]string `tfsdk:"all" json:"all,omitempty"`
			Any        *[]map[string]string `tfsdk:"any" json:"any,omitempty"`
			Operator   *string              `tfsdk:"operator" json:"operator,omitempty"`
			PatternRef *string              `tfsdk:"pattern_ref" json:"patternRef,omitempty"`
			Selector   *string              `tfsdk:"selector" json:"selector,omitempty"`
			Value      *string              `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"when" json:"when,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KuadrantIoAuthPolicyV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kuadrant_io_auth_policy_v1beta2_manifest"
}

func (r *KuadrantIoAuthPolicyV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AuthPolicy enables authentication and authorization for service workloads in a Gateway API network",
		MarkdownDescription: "AuthPolicy enables authentication and authorization for service workloads in a Gateway API network",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"patterns": schema.MapAttribute{
						Description:         "Named sets of patterns that can be referred in 'when' conditions and in pattern-matching authorization policy rules.",
						MarkdownDescription: "Named sets of patterns that can be referred in 'when' conditions and in pattern-matching authorization policy rules.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_selectors": schema.ListNestedAttribute{
						Description:         "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the external authorization service.At least one selected HTTPRoute rule must match to trigger the AuthPolicy.If no route selectors are specified, the AuthPolicy will be enforced at all requests to the protected routes.",
						MarkdownDescription: "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the external authorization service.At least one selected HTTPRoute rule must match to trigger the AuthPolicy.If no route selectors are specified, the AuthPolicy will be enforced at all requests to the protected routes.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hostnames": schema.ListAttribute{
									Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
									MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"matches": schema.ListNestedAttribute{
									Description:         "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
									MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"headers": schema.ListNestedAttribute{
												Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
												MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
															MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(256),
																stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
															},
														},

														"type": schema.StringAttribute{
															Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
															MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Exact", "RegularExpression"),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value is the value of HTTP Header to be matched.",
															MarkdownDescription: "Value is the value of HTTP Header to be matched.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(4096),
															},
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"method": schema.StringAttribute{
												Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
												MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
												},
											},

											"path": schema.SingleNestedAttribute{
												Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
												MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
												Attributes: map[string]schema.Attribute{
													"type": schema.StringAttribute{
														Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
														MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
														},
													},

													"value": schema.StringAttribute{
														Description:         "Value of the HTTP path to match against.",
														MarkdownDescription: "Value of the HTTP path to match against.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(1024),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"query_params": schema.ListNestedAttribute{
												Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
												MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
															MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(256),
																stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
															},
														},

														"type": schema.StringAttribute{
															Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
															MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("Exact", "RegularExpression"),
															},
														},

														"value": schema.StringAttribute{
															Description:         "Value is the value of HTTP query param to be matched.",
															MarkdownDescription: "Value is the value of HTTP query param to be matched.",
															Required:            true,
															Optional:            false,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.LengthAtLeast(1),
																stringvalidator.LengthAtMost(1024),
															},
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rules": schema.SingleNestedAttribute{
						Description:         "The auth rules of the policy.See Authorino's AuthConfig CRD for more details.",
						MarkdownDescription: "The auth rules of the policy.See Authorino's AuthConfig CRD for more details.",
						Attributes: map[string]schema.Attribute{
							"authentication": schema.SingleNestedAttribute{
								Description:         "Authentication configs.At least one config MUST evaluate to a valid identity object for the auth request to be successful.",
								MarkdownDescription: "Authentication configs.At least one config MUST evaluate to a valid identity object for the auth request to be successful.",
								Attributes: map[string]schema.Attribute{
									"anonymous": schema.MapAttribute{
										Description:         "Anonymous access.",
										MarkdownDescription: "Anonymous access.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"api_key": schema.SingleNestedAttribute{
										Description:         "Authentication based on API keys stored in Kubernetes secrets.",
										MarkdownDescription: "Authentication based on API keys stored in Kubernetes secrets.",
										Attributes: map[string]schema.Attribute{
											"all_namespaces": schema.BoolAttribute{
												Description:         "Whether Authorino should look for API key secrets in all namespaces or only in the same namespace as the AuthConfig.Enabling this option in namespaced Authorino instances has no effect.",
												MarkdownDescription: "Whether Authorino should look for API key secrets in all namespaces or only in the same namespace as the AuthConfig.Enabling this option in namespaced Authorino instances has no effect.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"selector": schema.SingleNestedAttribute{
												Description:         "Label selector used by Authorino to match secrets from the cluster storing valid credentials to authenticate to this service",
												MarkdownDescription: "Label selector used by Authorino to match secrets from the cluster storing valid credentials to authenticate to this service",
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
																	Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																	MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"values": schema.ListAttribute{
																	Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																	MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														ElementType:         types.StringType,
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cache": schema.SingleNestedAttribute{
										Description:         "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
										MarkdownDescription: "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
										Attributes: map[string]schema.Attribute{
											"key": schema.SingleNestedAttribute{
												Description:         "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
												MarkdownDescription: "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"ttl": schema.Int64Attribute{
												Description:         "Duration (in seconds) of the external data in the cache before pulled again from the source.",
												MarkdownDescription: "Duration (in seconds) of the external data in the cache before pulled again from the source.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"credentials": schema.SingleNestedAttribute{
										Description:         "Defines where credentials are required to be passed in the request for authentication based on this config.If omitted, it defaults to credentials passed in the HTTP Authorization header and the 'Bearer' prefix prepended to the secret credential value.",
										MarkdownDescription: "Defines where credentials are required to be passed in the request for authentication based on this config.If omitted, it defaults to credentials passed in the HTTP Authorization header and the 'Bearer' prefix prepended to the secret credential value.",
										Attributes: map[string]schema.Attribute{
											"authorization_header": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"prefix": schema.StringAttribute{
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

											"cookie": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
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

											"custom_header": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
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

											"query_string": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"defaults": schema.SingleNestedAttribute{
										Description:         "Set default property values (claims) for the resolved identity object, that are set before appending the object tothe authorization JSON. If the property is already present in the resolved identity object, the default value is ignored.It requires the resolved identity object to always be a JSON object.Do not use this option with identity objects of other JSON types (array, string, etc).",
										MarkdownDescription: "Set default property values (claims) for the resolved identity object, that are set before appending the object tothe authorization JSON. If the property is already present in the resolved identity object, the default value is ignored.It requires the resolved identity object to always be a JSON object.Do not use this option with identity objects of other JSON types (array, string, etc).",
										Attributes: map[string]schema.Attribute{
											"selector": schema.StringAttribute{
												Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
												MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value": schema.MapAttribute{
												Description:         "Static value",
												MarkdownDescription: "Static value",
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

									"jwt": schema.SingleNestedAttribute{
										Description:         "Authentication based on JWT tokens.",
										MarkdownDescription: "Authentication based on JWT tokens.",
										Attributes: map[string]schema.Attribute{
											"issuer_url": schema.StringAttribute{
												Description:         "URL of the issuer of the JWT.If 'jwksUrl' is omitted, Authorino will append the path to the OpenID Connect Well-Known Discovery endpoint(i.e. '/.well-known/openid-configuration') to this URL, to discover the OIDC configuration where to obtainthe 'jkws_uri' claim from.The value must coincide with the value of  the 'iss' (issuer) claim of the discovered OpenID Connect configuration.",
												MarkdownDescription: "URL of the issuer of the JWT.If 'jwksUrl' is omitted, Authorino will append the path to the OpenID Connect Well-Known Discovery endpoint(i.e. '/.well-known/openid-configuration') to this URL, to discover the OIDC configuration where to obtainthe 'jkws_uri' claim from.The value must coincide with the value of  the 'iss' (issuer) claim of the discovered OpenID Connect configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ttl": schema.Int64Attribute{
												Description:         "Decides how long to wait before refreshing the JWKS (in seconds).If omitted, Authorino will never refresh the JWKS.",
												MarkdownDescription: "Decides how long to wait before refreshing the JWKS (in seconds).If omitted, Authorino will never refresh the JWKS.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kubernetes_token_review": schema.SingleNestedAttribute{
										Description:         "Authentication by Kubernetes token review.",
										MarkdownDescription: "Authentication by Kubernetes token review.",
										Attributes: map[string]schema.Attribute{
											"audiences": schema.ListAttribute{
												Description:         "The list of audiences (scopes) that must be claimed in a Kubernetes authentication token supplied in the request, and reviewed by Authorino.If omitted, Authorino will review tokens expecting the host name of the requested protected service amongst the audiences.",
												MarkdownDescription: "The list of audiences (scopes) that must be claimed in a Kubernetes authentication token supplied in the request, and reviewed by Authorino.If omitted, Authorino will review tokens expecting the host name of the requested protected service amongst the audiences.",
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

									"metrics": schema.BoolAttribute{
										Description:         "Whether this config should generate individual observability metrics",
										MarkdownDescription: "Whether this config should generate individual observability metrics",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"oauth2_introspection": schema.SingleNestedAttribute{
										Description:         "Authentication by OAuth2 token introspection.",
										MarkdownDescription: "Authentication by OAuth2 token introspection.",
										Attributes: map[string]schema.Attribute{
											"credentials_ref": schema.SingleNestedAttribute{
												Description:         "Reference to a Kubernetes secret in the same namespace, that stores client credentials to the OAuth2 server.",
												MarkdownDescription: "Reference to a Kubernetes secret in the same namespace, that stores client credentials to the OAuth2 server.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"endpoint": schema.StringAttribute{
												Description:         "The full URL of the token introspection endpoint.",
												MarkdownDescription: "The full URL of the token introspection endpoint.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"token_type_hint": schema.StringAttribute{
												Description:         "The token type hint for the token introspection.If omitted, it defaults to 'access_token'.",
												MarkdownDescription: "The token type hint for the token introspection.If omitted, it defaults to 'access_token'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"overrides": schema.SingleNestedAttribute{
										Description:         "Overrides the resolved identity object by setting the additional properties (claims) specified in this config,before appending the object to the authorization JSON.It requires the resolved identity object to always be a JSON object.Do not use this option with identity objects of other JSON types (array, string, etc).",
										MarkdownDescription: "Overrides the resolved identity object by setting the additional properties (claims) specified in this config,before appending the object to the authorization JSON.It requires the resolved identity object to always be a JSON object.Do not use this option with identity objects of other JSON types (array, string, etc).",
										Attributes: map[string]schema.Attribute{
											"selector": schema.StringAttribute{
												Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
												MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"value": schema.MapAttribute{
												Description:         "Static value",
												MarkdownDescription: "Static value",
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

									"plain": schema.SingleNestedAttribute{
										Description:         "Identity object extracted from the context.Use this method when authentication is performed beforehand by a proxy and the resulting object passed to Authorino as JSON in the auth request.",
										MarkdownDescription: "Identity object extracted from the context.Use this method when authentication is performed beforehand by a proxy and the resulting object passed to Authorino as JSON in the auth request.",
										Attributes: map[string]schema.Attribute{
											"selector": schema.StringAttribute{
												Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
												MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"priority": schema.Int64Attribute{
										Description:         "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
										MarkdownDescription: "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"route_selectors": schema.ListNestedAttribute{
										Description:         "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
										MarkdownDescription: "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"hostnames": schema.ListAttribute{
													Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"matches": schema.ListNestedAttribute{
													Description:         "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"headers": schema.ListNestedAttribute{
																Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																			MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(256),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																			MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Exact", "RegularExpression"),
																			},
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the value of HTTP Header to be matched.",
																			MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(4096),
																			},
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"method": schema.StringAttribute{
																Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
																},
															},

															"path": schema.SingleNestedAttribute{
																Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																Attributes: map[string]schema.Attribute{
																	"type": schema.StringAttribute{
																		Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																		MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
																		},
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value of the HTTP path to match against.",
																		MarkdownDescription: "Value of the HTTP path to match against.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(1024),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"query_params": schema.ListNestedAttribute{
																Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																			MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(256),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																			MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Exact", "RegularExpression"),
																			},
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the value of HTTP query param to be matched.",
																			MarkdownDescription: "Value is the value of HTTP query param to be matched.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(1024),
																			},
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
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"when": schema.ListNestedAttribute{
										Description:         "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
										MarkdownDescription: "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"all": schema.ListAttribute{
													Description:         "A list of pattern expressions to be evaluated as a logical AND.",
													MarkdownDescription: "A list of pattern expressions to be evaluated as a logical AND.",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"any": schema.ListAttribute{
													Description:         "A list of pattern expressions to be evaluated as a logical OR.",
													MarkdownDescription: "A list of pattern expressions to be evaluated as a logical OR.",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
													MarkdownDescription: "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("eq", "neq", "incl", "excl", "matches"),
													},
												},

												"pattern_ref": schema.StringAttribute{
													Description:         "Reference to a named set of pattern expressions",
													MarkdownDescription: "Reference to a named set of pattern expressions",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.StringAttribute{
													Description:         "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
													MarkdownDescription: "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
													MarkdownDescription: "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
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

									"x509": schema.SingleNestedAttribute{
										Description:         "Authentication based on client X.509 certificates.The certificates presented by the clients must be signed by a trusted CA whose certificates are stored in Kubernetes secrets.",
										MarkdownDescription: "Authentication based on client X.509 certificates.The certificates presented by the clients must be signed by a trusted CA whose certificates are stored in Kubernetes secrets.",
										Attributes: map[string]schema.Attribute{
											"all_namespaces": schema.BoolAttribute{
												Description:         "Whether Authorino should look for TLS secrets in all namespaces or only in the same namespace as the AuthConfig.Enabling this option in namespaced Authorino instances has no effect.",
												MarkdownDescription: "Whether Authorino should look for TLS secrets in all namespaces or only in the same namespace as the AuthConfig.Enabling this option in namespaced Authorino instances has no effect.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"selector": schema.SingleNestedAttribute{
												Description:         "Label selector used by Authorino to match secrets from the cluster storing trusted CA certificates to validateclients trying to authenticate to this service",
												MarkdownDescription: "Label selector used by Authorino to match secrets from the cluster storing trusted CA certificates to validateclients trying to authenticate to this service",
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
																	Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																	MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"values": schema.ListAttribute{
																	Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																	MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														ElementType:         types.StringType,
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
										Required: false,
										Optional: true,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"authorization": schema.SingleNestedAttribute{
								Description:         "Authorization policies.All policies MUST evaluate to 'allowed = true' for the auth request be successful.",
								MarkdownDescription: "Authorization policies.All policies MUST evaluate to 'allowed = true' for the auth request be successful.",
								Attributes: map[string]schema.Attribute{
									"cache": schema.SingleNestedAttribute{
										Description:         "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
										MarkdownDescription: "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
										Attributes: map[string]schema.Attribute{
											"key": schema.SingleNestedAttribute{
												Description:         "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
												MarkdownDescription: "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"ttl": schema.Int64Attribute{
												Description:         "Duration (in seconds) of the external data in the cache before pulled again from the source.",
												MarkdownDescription: "Duration (in seconds) of the external data in the cache before pulled again from the source.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kubernetes_subject_access_review": schema.SingleNestedAttribute{
										Description:         "Authorization by Kubernetes SubjectAccessReview",
										MarkdownDescription: "Authorization by Kubernetes SubjectAccessReview",
										Attributes: map[string]schema.Attribute{
											"groups": schema.ListAttribute{
												Description:         "Groups the user must be a member of or, if 'user' is omitted, the groups to check for authorization in the Kubernetes RBAC.",
												MarkdownDescription: "Groups the user must be a member of or, if 'user' is omitted, the groups to check for authorization in the Kubernetes RBAC.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resource_attributes": schema.SingleNestedAttribute{
												Description:         "Use resourceAttributes to check permissions on Kubernetes resources.If omitted, it performs a non-resource SubjectAccessReview, with verb and path inferred from the request.",
												MarkdownDescription: "Use resourceAttributes to check permissions on Kubernetes resources.If omitted, it performs a non-resource SubjectAccessReview, with verb and path inferred from the request.",
												Attributes: map[string]schema.Attribute{
													"group": schema.SingleNestedAttribute{
														Description:         "API group of the resource.Use '*' for all API groups.",
														MarkdownDescription: "API group of the resource.Use '*' for all API groups.",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"name": schema.SingleNestedAttribute{
														Description:         "Resource nameOmit it to check for authorization on all resources of the specified kind.",
														MarkdownDescription: "Resource nameOmit it to check for authorization on all resources of the specified kind.",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"namespace": schema.SingleNestedAttribute{
														Description:         "Namespace where the user must have permissions on the resource.",
														MarkdownDescription: "Namespace where the user must have permissions on the resource.",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"resource": schema.SingleNestedAttribute{
														Description:         "Resource kindUse '*' for all resource kinds.",
														MarkdownDescription: "Resource kindUse '*' for all resource kinds.",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"subresource": schema.SingleNestedAttribute{
														Description:         "Subresource kind",
														MarkdownDescription: "Subresource kind",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"verb": schema.SingleNestedAttribute{
														Description:         "Verb to check for authorization on the resource.Use '*' for all verbs.",
														MarkdownDescription: "Verb to check for authorization on the resource.Use '*' for all verbs.",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

											"user": schema.SingleNestedAttribute{
												Description:         "User to check for authorization in the Kubernetes RBAC.Omit it to check for group authorization only.",
												MarkdownDescription: "User to check for authorization in the Kubernetes RBAC.Omit it to check for group authorization only.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

									"metrics": schema.BoolAttribute{
										Description:         "Whether this config should generate individual observability metrics",
										MarkdownDescription: "Whether this config should generate individual observability metrics",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"opa": schema.SingleNestedAttribute{
										Description:         "Open Policy Agent (OPA) Rego policy.",
										MarkdownDescription: "Open Policy Agent (OPA) Rego policy.",
										Attributes: map[string]schema.Attribute{
											"all_values": schema.BoolAttribute{
												Description:         "Returns the value of all Rego rules in the virtual document. Values can be read in subsequent evaluators/phases of the Auth Pipeline.Otherwise, only the default 'allow' rule will be exposed.Returning all Rego rules can affect performance of OPA policies during reconciliation (policy precompile) and at runtime.",
												MarkdownDescription: "Returns the value of all Rego rules in the virtual document. Values can be read in subsequent evaluators/phases of the Auth Pipeline.Otherwise, only the default 'allow' rule will be exposed.Returning all Rego rules can affect performance of OPA policies during reconciliation (policy precompile) and at runtime.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_policy": schema.SingleNestedAttribute{
												Description:         "Settings for fetching the OPA policy from an external registry.Use it alternatively to 'rego'.For the configurations of the HTTP request, the following options are not implemented: 'method', 'body', 'bodyParameters','contentType', 'headers', 'oauth2'. Use it only with: 'url', 'sharedSecret', 'credentials'.",
												MarkdownDescription: "Settings for fetching the OPA policy from an external registry.Use it alternatively to 'rego'.For the configurations of the HTTP request, the following options are not implemented: 'method', 'body', 'bodyParameters','contentType', 'headers', 'oauth2'. Use it only with: 'url', 'sharedSecret', 'credentials'.",
												Attributes: map[string]schema.Attribute{
													"body": schema.SingleNestedAttribute{
														Description:         "Raw body of the HTTP request.Supersedes 'bodyParameters'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
														MarkdownDescription: "Raw body of the HTTP request.Supersedes 'bodyParameters'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"body_parameters": schema.SingleNestedAttribute{
														Description:         "Custom parameters to encode in the body of the HTTP request.Superseded by 'body'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
														MarkdownDescription: "Custom parameters to encode in the body of the HTTP request.Superseded by 'body'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"content_type": schema.StringAttribute{
														Description:         "Content-Type of the request body. Shapes how 'bodyParameters' are encoded.Use it with method=POST; for GET requests, Content-Type is automatically set to 'text/plain'.",
														MarkdownDescription: "Content-Type of the request body. Shapes how 'bodyParameters' are encoded.Use it with method=POST; for GET requests, Content-Type is automatically set to 'text/plain'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("application/x-www-form-urlencoded", "application/json"),
														},
													},

													"credentials": schema.SingleNestedAttribute{
														Description:         "Defines where client credentials will be passed in the request to the service.If omitted, it defaults to client credentials passed in the HTTP Authorization header and the 'Bearer' prefix expected prepended to the secret value.",
														MarkdownDescription: "Defines where client credentials will be passed in the request to the service.If omitted, it defaults to client credentials passed in the HTTP Authorization header and the 'Bearer' prefix expected prepended to the secret value.",
														Attributes: map[string]schema.Attribute{
															"authorization_header": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"prefix": schema.StringAttribute{
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

															"cookie": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
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

															"custom_header": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
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

															"query_string": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"headers": schema.SingleNestedAttribute{
														Description:         "Custom headers in the HTTP request.",
														MarkdownDescription: "Custom headers in the HTTP request.",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"method": schema.StringAttribute{
														Description:         "HTTP verb used in the request to the service. Accepted values: GET (default), POST.When the request method is POST, the authorization JSON is passed in the body of the request.",
														MarkdownDescription: "HTTP verb used in the request to the service. Accepted values: GET (default), POST.When the request method is POST, the authorization JSON is passed in the body of the request.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "CONNECT", "TRACE"),
														},
													},

													"oauth2": schema.SingleNestedAttribute{
														Description:         "Authentication with the HTTP service by OAuth2 Client Credentials grant.",
														MarkdownDescription: "Authentication with the HTTP service by OAuth2 Client Credentials grant.",
														Attributes: map[string]schema.Attribute{
															"cache": schema.BoolAttribute{
																Description:         "Caches and reuses the token until expired.Set it to false to force fetch the token at every authorization request regardless of expiration.",
																MarkdownDescription: "Caches and reuses the token until expired.Set it to false to force fetch the token at every authorization request regardless of expiration.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"client_id": schema.StringAttribute{
																Description:         "OAuth2 Client ID.",
																MarkdownDescription: "OAuth2 Client ID.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"client_secret_ref": schema.SingleNestedAttribute{
																Description:         "Reference to a Kuberentes Secret key that stores that OAuth2 Client Secret.",
																MarkdownDescription: "Reference to a Kuberentes Secret key that stores that OAuth2 Client Secret.",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the secret in the Authorino's namespace to select from.",
																		MarkdownDescription: "The name of the secret in the Authorino's namespace to select from.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"extra_params": schema.MapAttribute{
																Description:         "Optional extra parameters for the requests to the token URL.",
																MarkdownDescription: "Optional extra parameters for the requests to the token URL.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"scopes": schema.ListAttribute{
																Description:         "Optional scopes for the client credentials grant, if supported by he OAuth2 server.",
																MarkdownDescription: "Optional scopes for the client credentials grant, if supported by he OAuth2 server.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"token_url": schema.StringAttribute{
																Description:         "Token endpoint URL of the OAuth2 resource server.",
																MarkdownDescription: "Token endpoint URL of the OAuth2 resource server.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"shared_secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Secret key whose value will be passed by Authorino in the request.The HTTP service can use the shared secret to authenticate the origin of the request.Ignored if used together with oauth2.",
														MarkdownDescription: "Reference to a Secret key whose value will be passed by Authorino in the request.The HTTP service can use the shared secret to authenticate the origin of the request.Ignored if used together with oauth2.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the Authorino's namespace to select from.",
																MarkdownDescription: "The name of the secret in the Authorino's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"ttl": schema.Int64Attribute{
														Description:         "Duration (in seconds) of the external data in the cache before pulled again from the source.",
														MarkdownDescription: "Duration (in seconds) of the external data in the cache before pulled again from the source.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
														Description:         "Endpoint URL of the HTTP service.The value can include variable placeholders in the format '{selector}', where 'selector' is any pattern supportedby https://pkg.go.dev/github.com/tidwall/gjson and selects value from the authorization JSON.E.g. https://ext-auth-server.io/metadata?p={request.path}",
														MarkdownDescription: "Endpoint URL of the HTTP service.The value can include variable placeholders in the format '{selector}', where 'selector' is any pattern supportedby https://pkg.go.dev/github.com/tidwall/gjson and selects value from the authorization JSON.E.g. https://ext-auth-server.io/metadata?p={request.path}",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"rego": schema.StringAttribute{
												Description:         "Authorization policy as a Rego language document.The Rego document must include the 'allow' condition, set by Authorino to 'false' by default (i.e. requests are unauthorized unless changed).The Rego document must NOT include the 'package' declaration in line 1.",
												MarkdownDescription: "Authorization policy as a Rego language document.The Rego document must include the 'allow' condition, set by Authorino to 'false' by default (i.e. requests are unauthorized unless changed).The Rego document must NOT include the 'package' declaration in line 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pattern_matching": schema.SingleNestedAttribute{
										Description:         "Pattern-matching authorization rules.",
										MarkdownDescription: "Pattern-matching authorization rules.",
										Attributes: map[string]schema.Attribute{
											"patterns": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"all": schema.ListAttribute{
															Description:         "A list of pattern expressions to be evaluated as a logical AND.",
															MarkdownDescription: "A list of pattern expressions to be evaluated as a logical AND.",
															ElementType:         types.MapType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"any": schema.ListAttribute{
															Description:         "A list of pattern expressions to be evaluated as a logical OR.",
															MarkdownDescription: "A list of pattern expressions to be evaluated as a logical OR.",
															ElementType:         types.MapType{ElemType: types.StringType},
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
															MarkdownDescription: "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("eq", "neq", "incl", "excl", "matches"),
															},
														},

														"pattern_ref": schema.StringAttribute{
															Description:         "Reference to a named set of pattern expressions",
															MarkdownDescription: "Reference to a named set of pattern expressions",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"selector": schema.StringAttribute{
															Description:         "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
															MarkdownDescription: "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
															MarkdownDescription: "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"priority": schema.Int64Attribute{
										Description:         "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
										MarkdownDescription: "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"route_selectors": schema.ListNestedAttribute{
										Description:         "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
										MarkdownDescription: "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"hostnames": schema.ListAttribute{
													Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"matches": schema.ListNestedAttribute{
													Description:         "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"headers": schema.ListNestedAttribute{
																Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																			MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(256),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																			MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Exact", "RegularExpression"),
																			},
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the value of HTTP Header to be matched.",
																			MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(4096),
																			},
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"method": schema.StringAttribute{
																Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
																},
															},

															"path": schema.SingleNestedAttribute{
																Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																Attributes: map[string]schema.Attribute{
																	"type": schema.StringAttribute{
																		Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																		MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
																		},
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value of the HTTP path to match against.",
																		MarkdownDescription: "Value of the HTTP path to match against.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(1024),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"query_params": schema.ListNestedAttribute{
																Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																			MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(256),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																			MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Exact", "RegularExpression"),
																			},
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the value of HTTP query param to be matched.",
																			MarkdownDescription: "Value is the value of HTTP query param to be matched.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(1024),
																			},
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
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"spicedb": schema.SingleNestedAttribute{
										Description:         "Authorization decision delegated to external Authzed/SpiceDB server.",
										MarkdownDescription: "Authorization decision delegated to external Authzed/SpiceDB server.",
										Attributes: map[string]schema.Attribute{
											"endpoint": schema.StringAttribute{
												Description:         "Hostname and port number to the GRPC interface of the SpiceDB server (e.g. spicedb:50051).",
												MarkdownDescription: "Hostname and port number to the GRPC interface of the SpiceDB server (e.g. spicedb:50051).",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"insecure": schema.BoolAttribute{
												Description:         "Insecure HTTP connection (i.e. disables TLS verification)",
												MarkdownDescription: "Insecure HTTP connection (i.e. disables TLS verification)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"permission": schema.SingleNestedAttribute{
												Description:         "The name of the permission (or relation) on which to execute the check.",
												MarkdownDescription: "The name of the permission (or relation) on which to execute the check.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"resource": schema.SingleNestedAttribute{
												Description:         "The resource on which to check the permission or relation.",
												MarkdownDescription: "The resource on which to check the permission or relation.",
												Attributes: map[string]schema.Attribute{
													"kind": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"name": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

											"shared_secret_ref": schema.SingleNestedAttribute{
												Description:         "Reference to a Secret key whose value will be used by Authorino to authenticate with the Authzed service.",
												MarkdownDescription: "Reference to a Secret key whose value will be used by Authorino to authenticate with the Authzed service.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the Authorino's namespace to select from.",
														MarkdownDescription: "The name of the secret in the Authorino's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"subject": schema.SingleNestedAttribute{
												Description:         "The subject that will be checked for the permission or relation.",
												MarkdownDescription: "The subject that will be checked for the permission or relation.",
												Attributes: map[string]schema.Attribute{
													"kind": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"name": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"when": schema.ListNestedAttribute{
										Description:         "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
										MarkdownDescription: "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"all": schema.ListAttribute{
													Description:         "A list of pattern expressions to be evaluated as a logical AND.",
													MarkdownDescription: "A list of pattern expressions to be evaluated as a logical AND.",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"any": schema.ListAttribute{
													Description:         "A list of pattern expressions to be evaluated as a logical OR.",
													MarkdownDescription: "A list of pattern expressions to be evaluated as a logical OR.",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
													MarkdownDescription: "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("eq", "neq", "incl", "excl", "matches"),
													},
												},

												"pattern_ref": schema.StringAttribute{
													Description:         "Reference to a named set of pattern expressions",
													MarkdownDescription: "Reference to a named set of pattern expressions",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.StringAttribute{
													Description:         "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
													MarkdownDescription: "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
													MarkdownDescription: "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
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

							"callbacks": schema.SingleNestedAttribute{
								Description:         "Callback functions.Authorino sends callbacks at the end of the auth pipeline to the endpoints specified in this config.",
								MarkdownDescription: "Callback functions.Authorino sends callbacks at the end of the auth pipeline to the endpoints specified in this config.",
								Attributes: map[string]schema.Attribute{
									"cache": schema.SingleNestedAttribute{
										Description:         "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
										MarkdownDescription: "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
										Attributes: map[string]schema.Attribute{
											"key": schema.SingleNestedAttribute{
												Description:         "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
												MarkdownDescription: "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"ttl": schema.Int64Attribute{
												Description:         "Duration (in seconds) of the external data in the cache before pulled again from the source.",
												MarkdownDescription: "Duration (in seconds) of the external data in the cache before pulled again from the source.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"http": schema.SingleNestedAttribute{
										Description:         "Settings of the external HTTP request",
										MarkdownDescription: "Settings of the external HTTP request",
										Attributes: map[string]schema.Attribute{
											"body": schema.SingleNestedAttribute{
												Description:         "Raw body of the HTTP request.Supersedes 'bodyParameters'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
												MarkdownDescription: "Raw body of the HTTP request.Supersedes 'bodyParameters'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"body_parameters": schema.SingleNestedAttribute{
												Description:         "Custom parameters to encode in the body of the HTTP request.Superseded by 'body'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
												MarkdownDescription: "Custom parameters to encode in the body of the HTTP request.Superseded by 'body'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"content_type": schema.StringAttribute{
												Description:         "Content-Type of the request body. Shapes how 'bodyParameters' are encoded.Use it with method=POST; for GET requests, Content-Type is automatically set to 'text/plain'.",
												MarkdownDescription: "Content-Type of the request body. Shapes how 'bodyParameters' are encoded.Use it with method=POST; for GET requests, Content-Type is automatically set to 'text/plain'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("application/x-www-form-urlencoded", "application/json"),
												},
											},

											"credentials": schema.SingleNestedAttribute{
												Description:         "Defines where client credentials will be passed in the request to the service.If omitted, it defaults to client credentials passed in the HTTP Authorization header and the 'Bearer' prefix expected prepended to the secret value.",
												MarkdownDescription: "Defines where client credentials will be passed in the request to the service.If omitted, it defaults to client credentials passed in the HTTP Authorization header and the 'Bearer' prefix expected prepended to the secret value.",
												Attributes: map[string]schema.Attribute{
													"authorization_header": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"prefix": schema.StringAttribute{
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

													"cookie": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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

													"custom_header": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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

													"query_string": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"headers": schema.SingleNestedAttribute{
												Description:         "Custom headers in the HTTP request.",
												MarkdownDescription: "Custom headers in the HTTP request.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"method": schema.StringAttribute{
												Description:         "HTTP verb used in the request to the service. Accepted values: GET (default), POST.When the request method is POST, the authorization JSON is passed in the body of the request.",
												MarkdownDescription: "HTTP verb used in the request to the service. Accepted values: GET (default), POST.When the request method is POST, the authorization JSON is passed in the body of the request.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "CONNECT", "TRACE"),
												},
											},

											"oauth2": schema.SingleNestedAttribute{
												Description:         "Authentication with the HTTP service by OAuth2 Client Credentials grant.",
												MarkdownDescription: "Authentication with the HTTP service by OAuth2 Client Credentials grant.",
												Attributes: map[string]schema.Attribute{
													"cache": schema.BoolAttribute{
														Description:         "Caches and reuses the token until expired.Set it to false to force fetch the token at every authorization request regardless of expiration.",
														MarkdownDescription: "Caches and reuses the token until expired.Set it to false to force fetch the token at every authorization request regardless of expiration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"client_id": schema.StringAttribute{
														Description:         "OAuth2 Client ID.",
														MarkdownDescription: "OAuth2 Client ID.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"client_secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Kuberentes Secret key that stores that OAuth2 Client Secret.",
														MarkdownDescription: "Reference to a Kuberentes Secret key that stores that OAuth2 Client Secret.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the Authorino's namespace to select from.",
																MarkdownDescription: "The name of the secret in the Authorino's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"extra_params": schema.MapAttribute{
														Description:         "Optional extra parameters for the requests to the token URL.",
														MarkdownDescription: "Optional extra parameters for the requests to the token URL.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scopes": schema.ListAttribute{
														Description:         "Optional scopes for the client credentials grant, if supported by he OAuth2 server.",
														MarkdownDescription: "Optional scopes for the client credentials grant, if supported by he OAuth2 server.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"token_url": schema.StringAttribute{
														Description:         "Token endpoint URL of the OAuth2 resource server.",
														MarkdownDescription: "Token endpoint URL of the OAuth2 resource server.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"shared_secret_ref": schema.SingleNestedAttribute{
												Description:         "Reference to a Secret key whose value will be passed by Authorino in the request.The HTTP service can use the shared secret to authenticate the origin of the request.Ignored if used together with oauth2.",
												MarkdownDescription: "Reference to a Secret key whose value will be passed by Authorino in the request.The HTTP service can use the shared secret to authenticate the origin of the request.Ignored if used together with oauth2.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the Authorino's namespace to select from.",
														MarkdownDescription: "The name of the secret in the Authorino's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": schema.StringAttribute{
												Description:         "Endpoint URL of the HTTP service.The value can include variable placeholders in the format '{selector}', where 'selector' is any pattern supportedby https://pkg.go.dev/github.com/tidwall/gjson and selects value from the authorization JSON.E.g. https://ext-auth-server.io/metadata?p={request.path}",
												MarkdownDescription: "Endpoint URL of the HTTP service.The value can include variable placeholders in the format '{selector}', where 'selector' is any pattern supportedby https://pkg.go.dev/github.com/tidwall/gjson and selects value from the authorization JSON.E.g. https://ext-auth-server.io/metadata?p={request.path}",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"metrics": schema.BoolAttribute{
										Description:         "Whether this config should generate individual observability metrics",
										MarkdownDescription: "Whether this config should generate individual observability metrics",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"priority": schema.Int64Attribute{
										Description:         "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
										MarkdownDescription: "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"route_selectors": schema.ListNestedAttribute{
										Description:         "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
										MarkdownDescription: "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"hostnames": schema.ListAttribute{
													Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"matches": schema.ListNestedAttribute{
													Description:         "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"headers": schema.ListNestedAttribute{
																Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																			MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(256),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																			MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Exact", "RegularExpression"),
																			},
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the value of HTTP Header to be matched.",
																			MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(4096),
																			},
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"method": schema.StringAttribute{
																Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
																},
															},

															"path": schema.SingleNestedAttribute{
																Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																Attributes: map[string]schema.Attribute{
																	"type": schema.StringAttribute{
																		Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																		MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
																		},
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value of the HTTP path to match against.",
																		MarkdownDescription: "Value of the HTTP path to match against.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(1024),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"query_params": schema.ListNestedAttribute{
																Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																			MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(256),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																			MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Exact", "RegularExpression"),
																			},
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the value of HTTP query param to be matched.",
																			MarkdownDescription: "Value is the value of HTTP query param to be matched.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(1024),
																			},
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
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"when": schema.ListNestedAttribute{
										Description:         "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
										MarkdownDescription: "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"all": schema.ListAttribute{
													Description:         "A list of pattern expressions to be evaluated as a logical AND.",
													MarkdownDescription: "A list of pattern expressions to be evaluated as a logical AND.",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"any": schema.ListAttribute{
													Description:         "A list of pattern expressions to be evaluated as a logical OR.",
													MarkdownDescription: "A list of pattern expressions to be evaluated as a logical OR.",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
													MarkdownDescription: "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("eq", "neq", "incl", "excl", "matches"),
													},
												},

												"pattern_ref": schema.StringAttribute{
													Description:         "Reference to a named set of pattern expressions",
													MarkdownDescription: "Reference to a named set of pattern expressions",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.StringAttribute{
													Description:         "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
													MarkdownDescription: "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
													MarkdownDescription: "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
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

							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata sources.Authorino fetches auth metadata as JSON from sources specified in this config.",
								MarkdownDescription: "Metadata sources.Authorino fetches auth metadata as JSON from sources specified in this config.",
								Attributes: map[string]schema.Attribute{
									"cache": schema.SingleNestedAttribute{
										Description:         "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
										MarkdownDescription: "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
										Attributes: map[string]schema.Attribute{
											"key": schema.SingleNestedAttribute{
												Description:         "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
												MarkdownDescription: "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"ttl": schema.Int64Attribute{
												Description:         "Duration (in seconds) of the external data in the cache before pulled again from the source.",
												MarkdownDescription: "Duration (in seconds) of the external data in the cache before pulled again from the source.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"http": schema.SingleNestedAttribute{
										Description:         "External source of auth metadata via HTTP request",
										MarkdownDescription: "External source of auth metadata via HTTP request",
										Attributes: map[string]schema.Attribute{
											"body": schema.SingleNestedAttribute{
												Description:         "Raw body of the HTTP request.Supersedes 'bodyParameters'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
												MarkdownDescription: "Raw body of the HTTP request.Supersedes 'bodyParameters'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"body_parameters": schema.SingleNestedAttribute{
												Description:         "Custom parameters to encode in the body of the HTTP request.Superseded by 'body'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
												MarkdownDescription: "Custom parameters to encode in the body of the HTTP request.Superseded by 'body'; use either one or the other.Use it with method=POST; for GET requests, set parameters as query string in the 'endpoint' (placeholders can be used).",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"content_type": schema.StringAttribute{
												Description:         "Content-Type of the request body. Shapes how 'bodyParameters' are encoded.Use it with method=POST; for GET requests, Content-Type is automatically set to 'text/plain'.",
												MarkdownDescription: "Content-Type of the request body. Shapes how 'bodyParameters' are encoded.Use it with method=POST; for GET requests, Content-Type is automatically set to 'text/plain'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("application/x-www-form-urlencoded", "application/json"),
												},
											},

											"credentials": schema.SingleNestedAttribute{
												Description:         "Defines where client credentials will be passed in the request to the service.If omitted, it defaults to client credentials passed in the HTTP Authorization header and the 'Bearer' prefix expected prepended to the secret value.",
												MarkdownDescription: "Defines where client credentials will be passed in the request to the service.If omitted, it defaults to client credentials passed in the HTTP Authorization header and the 'Bearer' prefix expected prepended to the secret value.",
												Attributes: map[string]schema.Attribute{
													"authorization_header": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"prefix": schema.StringAttribute{
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

													"cookie": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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

													"custom_header": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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

													"query_string": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"headers": schema.SingleNestedAttribute{
												Description:         "Custom headers in the HTTP request.",
												MarkdownDescription: "Custom headers in the HTTP request.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"method": schema.StringAttribute{
												Description:         "HTTP verb used in the request to the service. Accepted values: GET (default), POST.When the request method is POST, the authorization JSON is passed in the body of the request.",
												MarkdownDescription: "HTTP verb used in the request to the service. Accepted values: GET (default), POST.When the request method is POST, the authorization JSON is passed in the body of the request.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS", "CONNECT", "TRACE"),
												},
											},

											"oauth2": schema.SingleNestedAttribute{
												Description:         "Authentication with the HTTP service by OAuth2 Client Credentials grant.",
												MarkdownDescription: "Authentication with the HTTP service by OAuth2 Client Credentials grant.",
												Attributes: map[string]schema.Attribute{
													"cache": schema.BoolAttribute{
														Description:         "Caches and reuses the token until expired.Set it to false to force fetch the token at every authorization request regardless of expiration.",
														MarkdownDescription: "Caches and reuses the token until expired.Set it to false to force fetch the token at every authorization request regardless of expiration.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"client_id": schema.StringAttribute{
														Description:         "OAuth2 Client ID.",
														MarkdownDescription: "OAuth2 Client ID.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"client_secret_ref": schema.SingleNestedAttribute{
														Description:         "Reference to a Kuberentes Secret key that stores that OAuth2 Client Secret.",
														MarkdownDescription: "Reference to a Kuberentes Secret key that stores that OAuth2 Client Secret.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key of the secret to select from.  Must be a valid secret key.",
																MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the secret in the Authorino's namespace to select from.",
																MarkdownDescription: "The name of the secret in the Authorino's namespace to select from.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"extra_params": schema.MapAttribute{
														Description:         "Optional extra parameters for the requests to the token URL.",
														MarkdownDescription: "Optional extra parameters for the requests to the token URL.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scopes": schema.ListAttribute{
														Description:         "Optional scopes for the client credentials grant, if supported by he OAuth2 server.",
														MarkdownDescription: "Optional scopes for the client credentials grant, if supported by he OAuth2 server.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"token_url": schema.StringAttribute{
														Description:         "Token endpoint URL of the OAuth2 resource server.",
														MarkdownDescription: "Token endpoint URL of the OAuth2 resource server.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"shared_secret_ref": schema.SingleNestedAttribute{
												Description:         "Reference to a Secret key whose value will be passed by Authorino in the request.The HTTP service can use the shared secret to authenticate the origin of the request.Ignored if used together with oauth2.",
												MarkdownDescription: "Reference to a Secret key whose value will be passed by Authorino in the request.The HTTP service can use the shared secret to authenticate the origin of the request.Ignored if used together with oauth2.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "The name of the secret in the Authorino's namespace to select from.",
														MarkdownDescription: "The name of the secret in the Authorino's namespace to select from.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": schema.StringAttribute{
												Description:         "Endpoint URL of the HTTP service.The value can include variable placeholders in the format '{selector}', where 'selector' is any pattern supportedby https://pkg.go.dev/github.com/tidwall/gjson and selects value from the authorization JSON.E.g. https://ext-auth-server.io/metadata?p={request.path}",
												MarkdownDescription: "Endpoint URL of the HTTP service.The value can include variable placeholders in the format '{selector}', where 'selector' is any pattern supportedby https://pkg.go.dev/github.com/tidwall/gjson and selects value from the authorization JSON.E.g. https://ext-auth-server.io/metadata?p={request.path}",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"metrics": schema.BoolAttribute{
										Description:         "Whether this config should generate individual observability metrics",
										MarkdownDescription: "Whether this config should generate individual observability metrics",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"priority": schema.Int64Attribute{
										Description:         "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
										MarkdownDescription: "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"route_selectors": schema.ListNestedAttribute{
										Description:         "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
										MarkdownDescription: "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"hostnames": schema.ListAttribute{
													Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"matches": schema.ListNestedAttribute{
													Description:         "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"headers": schema.ListNestedAttribute{
																Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																			MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(256),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																			MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Exact", "RegularExpression"),
																			},
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the value of HTTP Header to be matched.",
																			MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(4096),
																			},
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"method": schema.StringAttribute{
																Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
																},
															},

															"path": schema.SingleNestedAttribute{
																Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																Attributes: map[string]schema.Attribute{
																	"type": schema.StringAttribute{
																		Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																		MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
																		},
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value of the HTTP path to match against.",
																		MarkdownDescription: "Value of the HTTP path to match against.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtMost(1024),
																		},
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"query_params": schema.ListNestedAttribute{
																Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																			MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(256),
																				stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																			},
																		},

																		"type": schema.StringAttribute{
																			Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																			MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("Exact", "RegularExpression"),
																			},
																		},

																		"value": schema.StringAttribute{
																			Description:         "Value is the value of HTTP query param to be matched.",
																			MarkdownDescription: "Value is the value of HTTP query param to be matched.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.LengthAtLeast(1),
																				stringvalidator.LengthAtMost(1024),
																			},
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
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"uma": schema.SingleNestedAttribute{
										Description:         "User-Managed Access (UMA) source of resource data.",
										MarkdownDescription: "User-Managed Access (UMA) source of resource data.",
										Attributes: map[string]schema.Attribute{
											"credentials_ref": schema.SingleNestedAttribute{
												Description:         "Reference to a Kubernetes secret in the same namespace, that stores client credentials to the resource registration API of the UMA server.",
												MarkdownDescription: "Reference to a Kubernetes secret in the same namespace, that stores client credentials to the resource registration API of the UMA server.",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"endpoint": schema.StringAttribute{
												Description:         "The endpoint of the UMA server.The value must coincide with the 'issuer' claim of the UMA config discovered from the well-known uma configuration endpoint.",
												MarkdownDescription: "The endpoint of the UMA server.The value must coincide with the 'issuer' claim of the UMA config discovered from the well-known uma configuration endpoint.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_info": schema.SingleNestedAttribute{
										Description:         "OpendID Connect UserInfo linked to an OIDC authentication config specified in this same AuthConfig.",
										MarkdownDescription: "OpendID Connect UserInfo linked to an OIDC authentication config specified in this same AuthConfig.",
										Attributes: map[string]schema.Attribute{
											"identity_source": schema.StringAttribute{
												Description:         "The name of an OIDC-enabled JWT authentication config whose OpenID Connect configuration discovered includes the OIDC 'userinfo_endpoint' claim.",
												MarkdownDescription: "The name of an OIDC-enabled JWT authentication config whose OpenID Connect configuration discovered includes the OIDC 'userinfo_endpoint' claim.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"when": schema.ListNestedAttribute{
										Description:         "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
										MarkdownDescription: "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"all": schema.ListAttribute{
													Description:         "A list of pattern expressions to be evaluated as a logical AND.",
													MarkdownDescription: "A list of pattern expressions to be evaluated as a logical AND.",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"any": schema.ListAttribute{
													Description:         "A list of pattern expressions to be evaluated as a logical OR.",
													MarkdownDescription: "A list of pattern expressions to be evaluated as a logical OR.",
													ElementType:         types.MapType{ElemType: types.StringType},
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
													MarkdownDescription: "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("eq", "neq", "incl", "excl", "matches"),
													},
												},

												"pattern_ref": schema.StringAttribute{
													Description:         "Reference to a named set of pattern expressions",
													MarkdownDescription: "Reference to a named set of pattern expressions",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"selector": schema.StringAttribute{
													Description:         "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
													MarkdownDescription: "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
													MarkdownDescription: "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
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

							"response": schema.SingleNestedAttribute{
								Description:         "Response items.Authorino builds custom responses to the client of the auth request.",
								MarkdownDescription: "Response items.Authorino builds custom responses to the client of the auth request.",
								Attributes: map[string]schema.Attribute{
									"success": schema.SingleNestedAttribute{
										Description:         "Response items to be included in the auth response when the request is authenticated and authorized.For integration of Authorino via proxy, the proxy must use these settings to propagate dynamic metadata and/or inject data in the request.",
										MarkdownDescription: "Response items to be included in the auth response when the request is authenticated and authorized.For integration of Authorino via proxy, the proxy must use these settings to propagate dynamic metadata and/or inject data in the request.",
										Attributes: map[string]schema.Attribute{
											"dynamic_metadata": schema.SingleNestedAttribute{
												Description:         "Custom success response items wrapped as HTTP headers.For integration of Authorino via proxy, the proxy must use these settings to propagate dynamic metadata.See https://www.envoyproxy.io/docs/envoy/latest/configuration/advanced/well_known_dynamic_metadata",
												MarkdownDescription: "Custom success response items wrapped as HTTP headers.For integration of Authorino via proxy, the proxy must use these settings to propagate dynamic metadata.See https://www.envoyproxy.io/docs/envoy/latest/configuration/advanced/well_known_dynamic_metadata",
												Attributes: map[string]schema.Attribute{
													"cache": schema.SingleNestedAttribute{
														Description:         "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
														MarkdownDescription: "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
														Attributes: map[string]schema.Attribute{
															"key": schema.SingleNestedAttribute{
																Description:         "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
																MarkdownDescription: "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
																Attributes: map[string]schema.Attribute{
																	"selector": schema.StringAttribute{
																		Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.MapAttribute{
																		Description:         "Static value",
																		MarkdownDescription: "Static value",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"ttl": schema.Int64Attribute{
																Description:         "Duration (in seconds) of the external data in the cache before pulled again from the source.",
																MarkdownDescription: "Duration (in seconds) of the external data in the cache before pulled again from the source.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"json": schema.SingleNestedAttribute{
														Description:         "JSON objectSpecify it as the list of properties of the object, whose values can combine static values and values selected from the authorization JSON.",
														MarkdownDescription: "JSON objectSpecify it as the list of properties of the object, whose values can combine static values and values selected from the authorization JSON.",
														Attributes: map[string]schema.Attribute{
															"properties": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"selector": schema.StringAttribute{
																		Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.MapAttribute{
																		Description:         "Static value",
																		MarkdownDescription: "Static value",
																		ElementType:         types.StringType,
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": schema.StringAttribute{
														Description:         "The key used to add the custom response item (name of the HTTP header or root property of the Dynamic Metadata object).If omitted, it will be set to the name of the response config.",
														MarkdownDescription: "The key used to add the custom response item (name of the HTTP header or root property of the Dynamic Metadata object).If omitted, it will be set to the name of the response config.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metrics": schema.BoolAttribute{
														Description:         "Whether this config should generate individual observability metrics",
														MarkdownDescription: "Whether this config should generate individual observability metrics",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"plain": schema.SingleNestedAttribute{
														Description:         "Plain text content",
														MarkdownDescription: "Plain text content",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"priority": schema.Int64Attribute{
														Description:         "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
														MarkdownDescription: "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"route_selectors": schema.ListNestedAttribute{
														Description:         "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
														MarkdownDescription: "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"hostnames": schema.ListAttribute{
																	Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
																	MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"matches": schema.ListNestedAttribute{
																	Description:         "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
																	MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"headers": schema.ListNestedAttribute{
																				Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																				MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																							MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.LengthAtLeast(1),
																								stringvalidator.LengthAtMost(256),
																								stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																							},
																						},

																						"type": schema.StringAttribute{
																							Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																							MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.OneOf("Exact", "RegularExpression"),
																							},
																						},

																						"value": schema.StringAttribute{
																							Description:         "Value is the value of HTTP Header to be matched.",
																							MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.LengthAtLeast(1),
																								stringvalidator.LengthAtMost(4096),
																							},
																						},
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"method": schema.StringAttribute{
																				Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																				MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
																				},
																			},

																			"path": schema.SingleNestedAttribute{
																				Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																				MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																				Attributes: map[string]schema.Attribute{
																					"type": schema.StringAttribute{
																						Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																						MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
																						},
																					},

																					"value": schema.StringAttribute{
																						Description:         "Value of the HTTP path to match against.",
																						MarkdownDescription: "Value of the HTTP path to match against.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.LengthAtMost(1024),
																						},
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"query_params": schema.ListNestedAttribute{
																				Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																				MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																							MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.LengthAtLeast(1),
																								stringvalidator.LengthAtMost(256),
																								stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																							},
																						},

																						"type": schema.StringAttribute{
																							Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																							MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.OneOf("Exact", "RegularExpression"),
																							},
																						},

																						"value": schema.StringAttribute{
																							Description:         "Value is the value of HTTP query param to be matched.",
																							MarkdownDescription: "Value is the value of HTTP query param to be matched.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.LengthAtLeast(1),
																								stringvalidator.LengthAtMost(1024),
																							},
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"when": schema.ListNestedAttribute{
														Description:         "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
														MarkdownDescription: "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"all": schema.ListAttribute{
																	Description:         "A list of pattern expressions to be evaluated as a logical AND.",
																	MarkdownDescription: "A list of pattern expressions to be evaluated as a logical AND.",
																	ElementType:         types.MapType{ElemType: types.StringType},
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"any": schema.ListAttribute{
																	Description:         "A list of pattern expressions to be evaluated as a logical OR.",
																	MarkdownDescription: "A list of pattern expressions to be evaluated as a logical OR.",
																	ElementType:         types.MapType{ElemType: types.StringType},
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"operator": schema.StringAttribute{
																	Description:         "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
																	MarkdownDescription: "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("eq", "neq", "incl", "excl", "matches"),
																	},
																},

																"pattern_ref": schema.StringAttribute{
																	Description:         "Reference to a named set of pattern expressions",
																	MarkdownDescription: "Reference to a named set of pattern expressions",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"selector": schema.StringAttribute{
																	Description:         "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
																	MarkdownDescription: "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
																	MarkdownDescription: "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
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

													"wristband": schema.SingleNestedAttribute{
														Description:         "Authorino Festival Wristband token",
														MarkdownDescription: "Authorino Festival Wristband token",
														Attributes: map[string]schema.Attribute{
															"custom_claims": schema.SingleNestedAttribute{
																Description:         "Any claims to be added to the wristband token apart from the standard JWT claims (iss, iat, exp) added by default.",
																MarkdownDescription: "Any claims to be added to the wristband token apart from the standard JWT claims (iss, iat, exp) added by default.",
																Attributes: map[string]schema.Attribute{
																	"selector": schema.StringAttribute{
																		Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.MapAttribute{
																		Description:         "Static value",
																		MarkdownDescription: "Static value",
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

															"issuer": schema.StringAttribute{
																Description:         "The endpoint to the Authorino service that issues the wristband (format: <scheme>://<host>:<port>/<realm>, where <realm> = <namespace>/<authorino-auth-config-resource-name/wristband-config-name)",
																MarkdownDescription: "The endpoint to the Authorino service that issues the wristband (format: <scheme>://<host>:<port>/<realm>, where <realm> = <namespace>/<authorino-auth-config-resource-name/wristband-config-name)",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"signing_key_refs": schema.ListNestedAttribute{
																Description:         "Reference by name to Kubernetes secrets and corresponding signing algorithms.The secrets must contain a 'key.pem' entry whose value is the signing key formatted as PEM.",
																MarkdownDescription: "Reference by name to Kubernetes secrets and corresponding signing algorithms.The secrets must contain a 'key.pem' entry whose value is the signing key formatted as PEM.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"algorithm": schema.StringAttribute{
																			Description:         "Algorithm to sign the wristband token using the signing key provided",
																			MarkdownDescription: "Algorithm to sign the wristband token using the signing key provided",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("ES256", "ES384", "ES512", "RS256", "RS384", "RS512"),
																			},
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name of the signing key.The value is used to reference the Kubernetes secret that stores the key and in the 'kid' claim of the wristband token header.",
																			MarkdownDescription: "Name of the signing key.The value is used to reference the Kubernetes secret that stores the key and in the 'kid' claim of the wristband token header.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_duration": schema.Int64Attribute{
																Description:         "Time span of the wristband token, in seconds.",
																MarkdownDescription: "Time span of the wristband token, in seconds.",
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

											"headers": schema.SingleNestedAttribute{
												Description:         "Custom success response items wrapped as HTTP headers.For integration of Authorino via proxy, the proxy must use these settings to inject data in the request.",
												MarkdownDescription: "Custom success response items wrapped as HTTP headers.For integration of Authorino via proxy, the proxy must use these settings to inject data in the request.",
												Attributes: map[string]schema.Attribute{
													"cache": schema.SingleNestedAttribute{
														Description:         "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
														MarkdownDescription: "Caching options for the resolved object returned when applying this config.Omit it to avoid caching objects for this config.",
														Attributes: map[string]schema.Attribute{
															"key": schema.SingleNestedAttribute{
																Description:         "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
																MarkdownDescription: "Key used to store the entry in the cache.The resolved key must be unique within the scope of this particular config.",
																Attributes: map[string]schema.Attribute{
																	"selector": schema.StringAttribute{
																		Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.MapAttribute{
																		Description:         "Static value",
																		MarkdownDescription: "Static value",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"ttl": schema.Int64Attribute{
																Description:         "Duration (in seconds) of the external data in the cache before pulled again from the source.",
																MarkdownDescription: "Duration (in seconds) of the external data in the cache before pulled again from the source.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"json": schema.SingleNestedAttribute{
														Description:         "JSON objectSpecify it as the list of properties of the object, whose values can combine static values and values selected from the authorization JSON.",
														MarkdownDescription: "JSON objectSpecify it as the list of properties of the object, whose values can combine static values and values selected from the authorization JSON.",
														Attributes: map[string]schema.Attribute{
															"properties": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"selector": schema.StringAttribute{
																		Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.MapAttribute{
																		Description:         "Static value",
																		MarkdownDescription: "Static value",
																		ElementType:         types.StringType,
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
														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": schema.StringAttribute{
														Description:         "The key used to add the custom response item (name of the HTTP header or root property of the Dynamic Metadata object).If omitted, it will be set to the name of the response config.",
														MarkdownDescription: "The key used to add the custom response item (name of the HTTP header or root property of the Dynamic Metadata object).If omitted, it will be set to the name of the response config.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"metrics": schema.BoolAttribute{
														Description:         "Whether this config should generate individual observability metrics",
														MarkdownDescription: "Whether this config should generate individual observability metrics",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"plain": schema.SingleNestedAttribute{
														Description:         "Plain text content",
														MarkdownDescription: "Plain text content",
														Attributes: map[string]schema.Attribute{
															"selector": schema.StringAttribute{
																Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.MapAttribute{
																Description:         "Static value",
																MarkdownDescription: "Static value",
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

													"priority": schema.Int64Attribute{
														Description:         "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
														MarkdownDescription: "Priority group of the config.All configs in the same priority group are evaluated concurrently; consecutive priority groups are evaluated sequentially.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"route_selectors": schema.ListNestedAttribute{
														Description:         "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
														MarkdownDescription: "Top-level route selectors.If present, the elements will be used to select HTTPRoute rules that, when activated, trigger the auth rule.At least one selected HTTPRoute rule must match to trigger the auth rule.If no route selectors are specified, the auth rule will be evaluated at all requests to the protected routes.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"hostnames": schema.ListAttribute{
																	Description:         "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
																	MarkdownDescription: "Hostnames defines a set of hostname that should match against the HTTP Host header to select a HTTPRoute to process the requesthttps://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"matches": schema.ListNestedAttribute{
																	Description:         "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
																	MarkdownDescription: "Matches define conditions used for matching the rule against incoming HTTP requests.https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.HTTPRouteSpec",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"headers": schema.ListNestedAttribute{
																				Description:         "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																				MarkdownDescription: "Headers specifies HTTP request header matchers. Multiple match values areANDed together, meaning, a request must match all the specified headersto select the route.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																							MarkdownDescription: "Name is the name of the HTTP Header to be matched. Name matching MUST becase insensitive. (See https://tools.ietf.org/html/rfc7230#section-3.2).If multiple entries specify equivalent header names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent header name MUST be ignored. Due to thecase-insensitivity of header names, 'foo' and 'Foo' are consideredequivalent.When a header is repeated in an HTTP request, it isimplementation-specific behavior as to how this is represented.Generally, proxies should follow the guidance from the RFC:https://www.rfc-editor.org/rfc/rfc7230.html#section-3.2.2 regardingprocessing a repeated header, with special handling for 'Set-Cookie'.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.LengthAtLeast(1),
																								stringvalidator.LengthAtMost(256),
																								stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																							},
																						},

																						"type": schema.StringAttribute{
																							Description:         "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																							MarkdownDescription: "Type specifies how to match against the value of the header.Support: Core (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression HeaderMatchType has implementation-specificconformance, implementations can support POSIX, PCRE or any other dialectsof regular expressions. Please read the implementation's documentation todetermine the supported dialect.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.OneOf("Exact", "RegularExpression"),
																							},
																						},

																						"value": schema.StringAttribute{
																							Description:         "Value is the value of HTTP Header to be matched.",
																							MarkdownDescription: "Value is the value of HTTP Header to be matched.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.LengthAtLeast(1),
																								stringvalidator.LengthAtMost(4096),
																							},
																						},
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"method": schema.StringAttribute{
																				Description:         "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																				MarkdownDescription: "Method specifies HTTP method matcher.When specified, this route will be matched only if the request has thespecified method.Support: Extended",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.OneOf("GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"),
																				},
																			},

																			"path": schema.SingleNestedAttribute{
																				Description:         "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																				MarkdownDescription: "Path specifies a HTTP request path matcher. If this field is notspecified, a default prefix match on the '/' path is provided.",
																				Attributes: map[string]schema.Attribute{
																					"type": schema.StringAttribute{
																						Description:         "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																						MarkdownDescription: "Type specifies how to match against the path Value.Support: Core (Exact, PathPrefix)Support: Implementation-specific (RegularExpression)",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.OneOf("Exact", "PathPrefix", "RegularExpression"),
																						},
																					},

																					"value": schema.StringAttribute{
																						Description:         "Value of the HTTP path to match against.",
																						MarkdownDescription: "Value of the HTTP path to match against.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																						Validators: []validator.String{
																							stringvalidator.LengthAtMost(1024),
																						},
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"query_params": schema.ListNestedAttribute{
																				Description:         "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																				MarkdownDescription: "QueryParams specifies HTTP query parameter matchers. Multiple matchvalues are ANDed together, meaning, a request must match all thespecified query parameters to select the route.Support: Extended",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"name": schema.StringAttribute{
																							Description:         "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																							MarkdownDescription: "Name is the name of the HTTP query param to be matched. This must be anexact string match. (Seehttps://tools.ietf.org/html/rfc7230#section-2.7.3).If multiple entries specify equivalent query param names, only the firstentry with an equivalent name MUST be considered for a match. Subsequententries with an equivalent query param name MUST be ignored.If a query param is repeated in an HTTP request, the behavior ispurposely left undefined, since different data planes have differentcapabilities. However, it is *recommended* that implementations shouldmatch against the first value of the param if the data plane supports it,as this behavior is expected in other load balancing contexts outside ofthe Gateway API.Users SHOULD NOT route traffic based on repeated query params to guardthemselves against potential differences in the implementations.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.LengthAtLeast(1),
																								stringvalidator.LengthAtMost(256),
																								stringvalidator.RegexMatches(regexp.MustCompile(`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`), ""),
																							},
																						},

																						"type": schema.StringAttribute{
																							Description:         "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																							MarkdownDescription: "Type specifies how to match against the value of the query parameter.Support: Extended (Exact)Support: Implementation-specific (RegularExpression)Since RegularExpression QueryParamMatchType has Implementation-specificconformance, implementations can support POSIX, PCRE or any otherdialects of regular expressions. Please read the implementation'sdocumentation to determine the supported dialect.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.OneOf("Exact", "RegularExpression"),
																							},
																						},

																						"value": schema.StringAttribute{
																							Description:         "Value is the value of HTTP query param to be matched.",
																							MarkdownDescription: "Value is the value of HTTP query param to be matched.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																							Validators: []validator.String{
																								stringvalidator.LengthAtLeast(1),
																								stringvalidator.LengthAtMost(1024),
																							},
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"when": schema.ListNestedAttribute{
														Description:         "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
														MarkdownDescription: "Conditions for Authorino to enforce this config.If omitted, the config will be enforced for all requests.If present, all conditions must match for the config to be enforced; otherwise, the config will be skipped.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"all": schema.ListAttribute{
																	Description:         "A list of pattern expressions to be evaluated as a logical AND.",
																	MarkdownDescription: "A list of pattern expressions to be evaluated as a logical AND.",
																	ElementType:         types.MapType{ElemType: types.StringType},
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"any": schema.ListAttribute{
																	Description:         "A list of pattern expressions to be evaluated as a logical OR.",
																	MarkdownDescription: "A list of pattern expressions to be evaluated as a logical OR.",
																	ElementType:         types.MapType{ElemType: types.StringType},
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"operator": schema.StringAttribute{
																	Description:         "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
																	MarkdownDescription: "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																	Validators: []validator.String{
																		stringvalidator.OneOf("eq", "neq", "incl", "excl", "matches"),
																	},
																},

																"pattern_ref": schema.StringAttribute{
																	Description:         "Reference to a named set of pattern expressions",
																	MarkdownDescription: "Reference to a named set of pattern expressions",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"selector": schema.StringAttribute{
																	Description:         "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
																	MarkdownDescription: "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
																	MarkdownDescription: "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
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

													"wristband": schema.SingleNestedAttribute{
														Description:         "Authorino Festival Wristband token",
														MarkdownDescription: "Authorino Festival Wristband token",
														Attributes: map[string]schema.Attribute{
															"custom_claims": schema.SingleNestedAttribute{
																Description:         "Any claims to be added to the wristband token apart from the standard JWT claims (iss, iat, exp) added by default.",
																MarkdownDescription: "Any claims to be added to the wristband token apart from the standard JWT claims (iss, iat, exp) added by default.",
																Attributes: map[string]schema.Attribute{
																	"selector": schema.StringAttribute{
																		Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.MapAttribute{
																		Description:         "Static value",
																		MarkdownDescription: "Static value",
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

															"issuer": schema.StringAttribute{
																Description:         "The endpoint to the Authorino service that issues the wristband (format: <scheme>://<host>:<port>/<realm>, where <realm> = <namespace>/<authorino-auth-config-resource-name/wristband-config-name)",
																MarkdownDescription: "The endpoint to the Authorino service that issues the wristband (format: <scheme>://<host>:<port>/<realm>, where <realm> = <namespace>/<authorino-auth-config-resource-name/wristband-config-name)",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"signing_key_refs": schema.ListNestedAttribute{
																Description:         "Reference by name to Kubernetes secrets and corresponding signing algorithms.The secrets must contain a 'key.pem' entry whose value is the signing key formatted as PEM.",
																MarkdownDescription: "Reference by name to Kubernetes secrets and corresponding signing algorithms.The secrets must contain a 'key.pem' entry whose value is the signing key formatted as PEM.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"algorithm": schema.StringAttribute{
																			Description:         "Algorithm to sign the wristband token using the signing key provided",
																			MarkdownDescription: "Algorithm to sign the wristband token using the signing key provided",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																			Validators: []validator.String{
																				stringvalidator.OneOf("ES256", "ES384", "ES512", "RS256", "RS384", "RS512"),
																			},
																		},

																		"name": schema.StringAttribute{
																			Description:         "Name of the signing key.The value is used to reference the Kubernetes secret that stores the key and in the 'kid' claim of the wristband token header.",
																			MarkdownDescription: "Name of the signing key.The value is used to reference the Kubernetes secret that stores the key and in the 'kid' claim of the wristband token header.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"token_duration": schema.Int64Attribute{
																Description:         "Time span of the wristband token, in seconds.",
																MarkdownDescription: "Time span of the wristband token, in seconds.",
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

									"unauthenticated": schema.SingleNestedAttribute{
										Description:         "Customizations on the denial status attributes when the request is unauthenticated.For integration of Authorino via proxy, the proxy must honour the response status attributes specified in this config.Default: 401 Unauthorized",
										MarkdownDescription: "Customizations on the denial status attributes when the request is unauthenticated.For integration of Authorino via proxy, the proxy must honour the response status attributes specified in this config.Default: 401 Unauthorized",
										Attributes: map[string]schema.Attribute{
											"body": schema.SingleNestedAttribute{
												Description:         "HTTP response body to override the default denial body.",
												MarkdownDescription: "HTTP response body to override the default denial body.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"code": schema.Int64Attribute{
												Description:         "HTTP status code to override the default denial status code.",
												MarkdownDescription: "HTTP status code to override the default denial status code.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(300),
													int64validator.AtMost(599),
												},
											},

											"headers": schema.SingleNestedAttribute{
												Description:         "HTTP response headers to override the default denial headers.",
												MarkdownDescription: "HTTP response headers to override the default denial headers.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"message": schema.SingleNestedAttribute{
												Description:         "HTTP message to override the default denial message.",
												MarkdownDescription: "HTTP message to override the default denial message.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

									"unauthorized": schema.SingleNestedAttribute{
										Description:         "Customizations on the denial status attributes when the request is unauthorized.For integration of Authorino via proxy, the proxy must honour the response status attributes specified in this config.Default: 403 Forbidden",
										MarkdownDescription: "Customizations on the denial status attributes when the request is unauthorized.For integration of Authorino via proxy, the proxy must honour the response status attributes specified in this config.Default: 403 Forbidden",
										Attributes: map[string]schema.Attribute{
											"body": schema.SingleNestedAttribute{
												Description:         "HTTP response body to override the default denial body.",
												MarkdownDescription: "HTTP response body to override the default denial body.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"code": schema.Int64Attribute{
												Description:         "HTTP status code to override the default denial status code.",
												MarkdownDescription: "HTTP status code to override the default denial status code.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(300),
													int64validator.AtMost(599),
												},
											},

											"headers": schema.SingleNestedAttribute{
												Description:         "HTTP response headers to override the default denial headers.",
												MarkdownDescription: "HTTP response headers to override the default denial headers.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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

											"message": schema.SingleNestedAttribute{
												Description:         "HTTP message to override the default denial message.",
												MarkdownDescription: "HTTP message to override the default denial message.",
												Attributes: map[string]schema.Attribute{
													"selector": schema.StringAttribute{
														Description:         "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														MarkdownDescription: "Simple path selector to fetch content from the authorization JSON (e.g. 'request.method') or a string template with variables that resolve to patterns (e.g. 'Hello, {auth.identity.name}!').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.The following Authorino custom modifiers are supported: @extract:{sep:' ',pos:0}, @replace{old:'',new:''}, @case:upper|lower, @base64:encode|decode and @strip.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.MapAttribute{
														Description:         "Static value",
														MarkdownDescription: "Static value",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_ref": schema.SingleNestedAttribute{
						Description:         "TargetRef identifies an API object to apply policy to.",
						MarkdownDescription: "TargetRef identifies an API object to apply policy to.",
						Attributes: map[string]schema.Attribute{
							"group": schema.StringAttribute{
								Description:         "Group is the group of the target resource.",
								MarkdownDescription: "Group is the group of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtMost(253),
									stringvalidator.RegexMatches(regexp.MustCompile(`^$|^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`), ""),
								},
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is kind of the target resource.",
								MarkdownDescription: "Kind is kind of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z]([-a-zA-Z0-9]*[a-zA-Z0-9])?$`), ""),
								},
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of the target resource.",
								MarkdownDescription: "Name is the name of the target resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(253),
								},
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace is the namespace of the referent. When unspecified, the localnamespace is inferred. Even when policy targets a resource in a differentnamespace, it MUST only apply to traffic originating from the samenamespace as the policy.",
								MarkdownDescription: "Namespace is the namespace of the referent. When unspecified, the localnamespace is inferred. Even when policy targets a resource in a differentnamespace, it MUST only apply to traffic originating from the samenamespace as the policy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
									stringvalidator.LengthAtMost(63),
									stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`), ""),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"when": schema.ListNestedAttribute{
						Description:         "Overall conditions for the AuthPolicy to be enforced.If omitted, the AuthPolicy will be enforced at all requests to the protected routes.If present, all conditions must match for the AuthPolicy to be enforced; otherwise, the authorization service skips the AuthPolicy and returns to the auth request with status OK.",
						MarkdownDescription: "Overall conditions for the AuthPolicy to be enforced.If omitted, the AuthPolicy will be enforced at all requests to the protected routes.If present, all conditions must match for the AuthPolicy to be enforced; otherwise, the authorization service skips the AuthPolicy and returns to the auth request with status OK.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"all": schema.ListAttribute{
									Description:         "A list of pattern expressions to be evaluated as a logical AND.",
									MarkdownDescription: "A list of pattern expressions to be evaluated as a logical AND.",
									ElementType:         types.MapType{ElemType: types.StringType},
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"any": schema.ListAttribute{
									Description:         "A list of pattern expressions to be evaluated as a logical OR.",
									MarkdownDescription: "A list of pattern expressions to be evaluated as a logical OR.",
									ElementType:         types.MapType{ElemType: types.StringType},
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
									MarkdownDescription: "The binary operator to be applied to the content fetched from the authorization JSON, for comparison with 'value'.Possible values are: 'eq' (equal to), 'neq' (not equal to), 'incl' (includes; for arrays), 'excl' (excludes; for arrays), 'matches' (regex)",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("eq", "neq", "incl", "excl", "matches"),
									},
								},

								"pattern_ref": schema.StringAttribute{
									Description:         "Reference to a named set of pattern expressions",
									MarkdownDescription: "Reference to a named set of pattern expressions",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"selector": schema.StringAttribute{
									Description:         "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
									MarkdownDescription: "Path selector to fetch content from the authorization JSON (e.g. 'request.method').Any pattern supported by https://pkg.go.dev/github.com/tidwall/gjson can be used.Authorino custom JSON path modifiers are also supported.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
									MarkdownDescription: "The value of reference for the comparison with the content fetched from the authorization JSON.If used with the 'matches' operator, the value must compile to a valid Golang regex.",
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

func (r *KuadrantIoAuthPolicyV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kuadrant_io_auth_policy_v1beta2_manifest")

	var model KuadrantIoAuthPolicyV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kuadrant.io/v1beta2")
	model.Kind = pointer.String("AuthPolicy")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}